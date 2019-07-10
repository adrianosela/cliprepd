package commands

import (
	"errors"
	"fmt"

	"github.com/adrianosela/cliprepd/config"
	"go.mozilla.org/iprepd"
	cli "gopkg.in/urfave/cli.v1"
)

// ViolationCmd is the CLI command object for the violation operation
var ViolationCmd = cli.Command{
	Name: "violation",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "violation, vn",
			Usage: "[mandatory] name of violation to be applied",
		},
		cli.StringFlag{
			Name:  "object, o",
			Usage: "[mandatory] object to apply violation to",
		},
		cli.StringFlag{
			Name:  "type, t",
			Usage: "[mandatory] type of object e.g. ip/email",
		},
		cli.IntFlag{
			Name:  "suppress-recovery, sr",
			Usage: "seconds to wait before reputation begins to heal for object",
			Value: 0,
		},
	},
	Before: violationFlagValidator,
	Action: violationHandler,
}

func violationFlagValidator(ctx *cli.Context) error {
	if ctx.String("violation") == "" {
		return errors.New("missing [mandatory] argument \"violation\"")
	}
	if ctx.String("object") == "" {
		return errors.New("missing [mandatory] argument \"object\"")
	}
	if ctx.String("type") == "" {
		return errors.New("missing [mandatory] argument \"type\"")
	}
	return nil
}

func violationHandler(ctx *cli.Context) error {
	client, err := config.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	if err = client.ApplyViolation(&iprepd.ViolationRequest{
		Object:           ctx.String("object"),
		Type:             ctx.String("type"),
		Violation:        ctx.String("violation"),
		SuppressRecovery: ctx.Int("suppress-recovery"),
	}); err != nil {
		return fmt.Errorf("could not apply violation: %s", err)
	}
	return nil
}
