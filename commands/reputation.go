package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/adrianosela/cliprepd/config"
	"github.com/olekukonko/tablewriter"
	cli "gopkg.in/urfave/cli.v1"
)

// ReputationCmd is the CLI command object for reputation operations
var ReputationCmd = cli.Command{
	Name: "reputation",
	Subcommands: []cli.Command{
		cli.Command{
			Name: "get",
			Flags: append(reputationSubcmdFlags, cli.BoolFlag{
				Name:  "json, j",
				Usage: "print raw json -- don't pretty print",
			}),
			Before: reputationFlagValidator,
			Action: reputationGetHandler,
		},
		cli.Command{
			Name:   "clear",
			Flags:  reputationSubcmdFlags,
			Before: reputationFlagValidator,
			Action: reputationClearHandler,
		},
		cli.Command{
			Name: "set",
		},
	},
}

var reputationSubcmdFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "object, o",
		Usage: "[mandatory] object to apply violation to",
	},
	cli.StringFlag{
		Name:  "type, t",
		Usage: "type of object",
		Value: "ip", // default to IP
	},
}

func reputationFlagValidator(ctx *cli.Context) error {
	if ctx.String("object") == "" {
		return errors.New("missing [mandatory] argument \"object\"")
	}
	return nil
}

func reputationGetHandler(ctx *cli.Context) error {
	typ := ctx.String("type")
	obj := ctx.String("object")

	client, err := config.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}

	rept, err := client.GetReputation(typ, obj)
	if err != nil {
		return fmt.Errorf("could not get reputation for %s=%s", typ, obj)
	}

	if ctx.BoolT("json") {
		raw, err := json.Marshal(rept)
		if err != nil {
			return fmt.Errorf("could not format response payload: %s", err)
		}
		fmt.Println(string(raw))
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Append([]string{"OBJECT", rept.Object})
	table.Append([]string{"REPUTATION", strconv.Itoa(rept.Reputation)})
	table.Append([]string{"TYPE", rept.Type})
	table.Append([]string{"REVIEWED", strconv.FormatBool(rept.Reviewed)})
	table.Append([]string{"LAST UPDATED", rept.LastUpdated.String()})
	table.Append([]string{"DECAY AFTER", rept.DecayAfter.String()})
	table.Render()

	return nil
}

func reputationClearHandler(ctx *cli.Context) error {
	typ := ctx.String("type")
	obj := ctx.String("object")

	client, err := config.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	if err := client.DeleteReputation(typ, obj); err != nil {
		return fmt.Errorf("could not delete reputation for %s=%s", typ, obj)
	}
	fmt.Printf("reputation for %s %s deleted succesfully!\n", typ, obj)
	return nil
}
