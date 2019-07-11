package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"go.mozilla.org/iprepd"
	cli "gopkg.in/urfave/cli.v1"
)

// ReputationCmd is the CLI command object for reputation operations
var ReputationCmd = cli.Command{
	Name:  "reputation",
	Usage: "reputation entry related commands",
	Subcommands: []cli.Command{
		{
			Name:   "get",
			Usage:  "get the entry for a given object",
			Flags:  append(reputationBaseFlags, reputationGetFlags...),
			Before: reputationBaseValidator,
			Action: reputationGetHandler,
		},
		{
			Name:   "clear",
			Usage:  "delete the entry for a given object",
			Flags:  reputationBaseFlags,
			Before: reputationBaseValidator,
			Action: reputationClearHandler,
		},
		{
			Name:   "set",
			Usage:  "update the entry for a given object",
			Flags:  append(reputationBaseFlags, reputationSetFlags...),
			Before: reputationSetValidator,
			Action: reputationSetHandler,
		},
	},
}

var reputationBaseFlags = []cli.Flag{
	withDefault(typeFlag, "ip"),
	asMandatory(objectFlag),
}

var reputationGetFlags = []cli.Flag{
	jsonFlag,
}

var reputationSetFlags = []cli.Flag{
	asMandatoryInt(scoreFlag),
	decayAfterFlag,
}

func reputationBaseValidator(ctx *cli.Context) error {
	return assertSet(ctx,
		objectFlag,
	)
}

func reputationSetValidator(ctx *cli.Context) error {
	if err := assertSet(ctx, scoreFlag); err != nil {
		return err
	}
	return reputationBaseValidator(ctx)
}

func reputationGetHandler(ctx *cli.Context) error {
	typ := ctx.String(name(typeFlag))
	obj := ctx.String(name(objectFlag))

	client, err := getClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	rept, err := client.GetReputation(typ, obj)
	if err != nil {
		return fmt.Errorf("could not get reputation for %s=%s", typ, obj)
	}
	if ctx.BoolT(name(jsonFlag)) {
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
	typ := ctx.String(name(typeFlag))
	obj := ctx.String(name(objectFlag))
	client, err := getClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	if err := client.DeleteReputation(typ, obj); err != nil {
		return fmt.Errorf("could not delete reputation for %s %s: %s", typ, obj, err)
	}
	fmt.Printf("reputation for %s %s deleted successfully!\n", typ, obj)
	return nil
}

func reputationSetHandler(ctx *cli.Context) error {
	typ := ctx.String(name(typeFlag))
	obj := ctx.String(name(objectFlag))
	rep := ctx.Int(name(scoreFlag))
	da := ctx.Int(name(decayAfterFlag))

	client, err := getClient(ctx)
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
	fmt.Printf("reputation for %s %s updated successfully!\n", typ, obj)
	return nil
}
