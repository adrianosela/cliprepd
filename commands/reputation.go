package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/adrianosela/cliprepd/config"
	"github.com/olekukonko/tablewriter"
	"go.mozilla.org/iprepd"
	cli "gopkg.in/urfave/cli.v1"
)

// ReputationCmd is the CLI command object for reputation operations
var ReputationCmd = cli.Command{
	Name: "reputation",
	Subcommands: []cli.Command{
		cli.Command{
			Name:   "get",
			Flags:  append(reputationBaseFlags, reputationGetFlags...),
			Before: reputationBaseFlagValidator,
			Action: reputationGetHandler,
		},
		cli.Command{
			Name:   "clear",
			Flags:  reputationBaseFlags,
			Before: reputationBaseFlagValidator,
			Action: reputationClearHandler,
		},
		cli.Command{
			Name:   "set",
			Flags:  append(reputationBaseFlags, reputationSetFlags...),
			Before: reputationSetFlagValidator,
			Action: reputationSetHandler,
		},
	},
}

var reputationBaseFlags = []cli.Flag{
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

var reputationGetFlags = []cli.Flag{
	jsonFlag,
}

var reputationSetFlags = []cli.Flag{
	cli.IntFlag{
		Name:  "score, s",
		Usage: "[mandatory] reputation score to assign ",
	},
	cli.IntFlag{
		Name:  "decay-after, d",
		Usage: "seconds after which reputation should begin to recover",
	},
}

func reputationBaseFlagValidator(ctx *cli.Context) error {
	if ctx.String("object") == "" {
		return errors.New("missing [mandatory] argument \"object\"")
	}
	return nil
}

func reputationSetFlagValidator(ctx *cli.Context) error {
	if !ctx.IsSet("score") {
		return errors.New("missing [mandatory] argument \"score\"")
	}
	return reputationBaseFlagValidator(ctx)
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
		return fmt.Errorf("could not delete reputation for %s %s: %s", typ, obj, err)
	}
	fmt.Printf("reputation for %s %s deleted succesfully!\n", typ, obj)
	return nil
}

func reputationSetHandler(ctx *cli.Context) error {
	typ := ctx.String("type")
	obj := ctx.String("object")
	rep := ctx.Int("score")
	da := ctx.Int("decay-after")

	client, err := config.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	if err := client.SetReputation(&iprepd.Reputation{
		Type:       typ,
		Object:     obj,
		Reputation: rep,
		DecayAfter: time.Now().Add(time.Second * time.Duration(da)),
	}); err != nil {
		return fmt.Errorf("could not update reputation for %s %s: %s", typ, obj, err)
	}
	fmt.Printf("reputation for %s %s updated succesfully!\n", typ, obj)
	return nil
}
