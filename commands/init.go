package commands

import (
	"errors"
	"fmt"

	"github.com/adrianosela/cliprepd/config"
	cli "gopkg.in/urfave/cli.v1"
)

// InitCmd is the CLI command object for the init operation
var InitCmd = cli.Command{
	Name: "init",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "path, p",
			Usage: "path where to save config file",
		},
		cli.StringFlag{
			Name:  "url, u",
			Usage: "[mandatory] host URL to use",
		},
		cli.StringFlag{
			Name:  "token, t",
			Usage: "[mandatory] auth token to use",
		},
	},
	Before: initFlagValidator,
	Action: initHandler,
}

func initFlagValidator(ctx *cli.Context) error {
	if ctx.String("url") == "" {
		return errors.New("missing [mandatory] argument \"url\"")
	}
	if ctx.String("token") == "" {
		return errors.New("missing [mandatory] argument \"token\"")
	}
	return nil
}

func initHandler(ctx *cli.Context) error {
	if err := config.SetConfig(
		ctx.String("url"),
		ctx.String("token"),
		ctx.String("path"),
	); err != nil {
		return fmt.Errorf("could not set configuration: %s", err)
	}
	return nil
}
