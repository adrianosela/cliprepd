package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/adrianosela/cliprepd/config"
	"github.com/olekukonko/tablewriter"
	"go.mozilla.org/iprepd"
	cli "gopkg.in/urfave/cli.v1"
)

// ViolationCmd is the CLI command object for the violation operation
var ViolationCmd = cli.Command{
	Name:  "violation",
	Usage: "violation related commands",
	Subcommands: []cli.Command{
		{
			Name:  "get",
			Usage: "show all available violations",
			Flags: []cli.Flag{
				jsonFlag,
			},
			Action: violationGetHandler,
		},
		{
			Name:  "apply",
			Usage: "apply a single violation to a given object",
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
			Before: violationApplyValidator,
			Action: violationApplyHandler,
		},
		{
			Name:        "batch-apply",
			Description: "see https://github.com/mozilla-services/iprepd#put-violationstypeip for payload format",
			Usage:       "batch-apply violations in a json file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "payload, p",
					Usage: "[mandatory] path to payload file",
				},
				cli.StringFlag{
					Name:  "type, t",
					Usage: "type of objects e.g. ip/email",
					Value: "ip",
				},
			},
			Before: violationBatchApplyValidator,
			Action: violationBatchApplyHandler,
		},
	},
}

func violationApplyValidator(ctx *cli.Context) error {
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

func violationBatchApplyValidator(ctx *cli.Context) error {
	path := ctx.String("payload")
	if path == "" {
		return errors.New("missing [mandatory] argument \"payload\"")
	}
	if _, err := readPayloadFile(path); err != nil {
		return fmt.Errorf("could not validate payload file: %s", err)
	}
	return nil
}

func readPayloadFile(path string) ([]iprepd.ViolationRequest, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read payload file %s: %s", path, err)
	}
	var vrs []iprepd.ViolationRequest
	if err = json.Unmarshal(dat, &vrs); err != nil {
		return nil, fmt.Errorf("could not unmarshal payload file: %s", err)
	}
	return vrs, nil
}

func violationGetHandler(ctx *cli.Context) error {
	client, err := config.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	vs, err := client.GetViolations()
	if err != nil {
		return fmt.Errorf("could not retrieve available violation: %s", err)
	}
	if ctx.BoolT("json") {
		if len(vs) == 0 {
			// ensure array format, i.e. ensure we dont print "nil"
			fmt.Println("[]")
			return nil
		}
		raw, err := json.Marshal(vs)
		if err != nil {
			return fmt.Errorf("could not format response payload: %s", err)
		}
		fmt.Println(string(raw))
		return nil
	}

	if len(vs) == 0 {
		fmt.Println("-- no violations to show --")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NAME", "PENALTY", "DECREASE LIMIT"})
	for _, viol := range vs {
		table.Append([]string{viol.Name, strconv.Itoa(viol.Penalty), strconv.Itoa(viol.DecreaseLimit)})
	}
	table.Render()

	return nil
}

func violationApplyHandler(ctx *cli.Context) error {
	obj := ctx.String("object")
	typ := ctx.String("type")
	vio := ctx.String("violation")

	client, err := config.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	if err = client.ApplyViolation(&iprepd.ViolationRequest{
		Object:           obj,
		Type:             typ,
		Violation:        vio,
		SuppressRecovery: ctx.Int("suppress-recovery"),
	}); err != nil {
		return fmt.Errorf("could not apply violation: %s", err)
	}
	fmt.Printf("violation %s successfully applied to %s %s!\n", vio, typ, obj)
	return nil
}

func violationBatchApplyHandler(ctx *cli.Context) error {
	typ := ctx.String("type")
	path := ctx.String("payload")
	if path == "" {
		return errors.New("missing [mandatory] argument \"payload\"")
	}
	violreqs, err := readPayloadFile(path)
	if err != nil {
		return fmt.Errorf("could not validate payload file: %s", err)
	}
	client, err := config.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	err = client.BatchApplyViolation(typ, violreqs)
	if err != nil {
		return fmt.Errorf("could not batch apply violations: %s", err)
	}
	return nil
}
