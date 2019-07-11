package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

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
			Name:  "list",
			Usage: "list all available violations",
			Flags: []cli.Flag{
				jsonFlag,
			},
			Action: violationListHandler,
		},
		{
			Name:  "apply",
			Usage: "apply a violation to a single object",
			Flags: []cli.Flag{
				asMandatory(violationFlag),
				asMandatory(objectFlag),
				withDefault(typeFlag, "ip"),
				suppressRecoveryFlag,
			},
			Before: violationApplyValidator,
			Action: violationApplyHandler,
		},
		{
			Name:        "batch-apply",
			Description: "see https://github.com/mozilla-services/iprepd#put-violationstypeip for payload format",
			Usage:       "batch-apply violations in a json file",
			Flags: []cli.Flag{
				asMandatory(payloadFlag),
				withDefault(typeFlag, "ip"),
			},
			Before: violationBatchApplyValidator,
			Action: violationBatchApplyHandler,
		},
	},
}

func violationApplyValidator(ctx *cli.Context) error {
	return assertSet(ctx,
		"violation",
		"object",
	)
}

func violationBatchApplyValidator(ctx *cli.Context) error {
	if err := assertSet(ctx, "payload"); err != nil {
		return err
	}
	path := ctx.String("payload")
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

func violationListHandler(ctx *cli.Context) error {
	client, err := getClient(ctx)
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
	sr := ctx.Int("suppress-recovery")

	client, err := getClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	if err = client.ApplyViolation(&iprepd.ViolationRequest{
		Object:           obj,
		Type:             typ,
		Violation:        vio,
		SuppressRecovery: sr,
	}); err != nil {
		return fmt.Errorf("could not apply violation: %s", err)
	}
	fmt.Printf("violation %s successfully applied to %s %s!\n", vio, typ, obj)
	return nil
}

func violationBatchApplyHandler(ctx *cli.Context) error {
	typ := ctx.String("type")
	path := ctx.String("payload")
	violreqs, err := readPayloadFile(path)
	if err != nil {
		return fmt.Errorf("could not validate payload file: %s", err)
	}
	client, err := getClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	err = client.BatchApplyViolation(typ, violreqs)
	if err != nil {
		return fmt.Errorf("could not batch apply violations: %s", err)
	}
	return nil
}
