package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/adrianosela/cliprepd/config"
	"github.com/olekukonko/tablewriter"
	cli "gopkg.in/urfave/cli.v1"
)

// DumpCmd is the CLI command object for the dump operation
var DumpCmd = cli.Command{
	Name:  "dump",
	Usage: "dump all available reputation entries",
	Flags: []cli.Flag{
		jsonFlag,
	},
	Action: dumpCmdHandler,
}

func dumpCmdHandler(ctx *cli.Context) error {
	client, err := config.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}

	entries, err := client.Dump()
	if err != nil {
		return fmt.Errorf("could not retrieve reputation entries: %s", err)
	}

	if ctx.BoolT("json") {
		if len(entries) == 0 {
			// ensure array format, i.e. ensure we dont print "nil"
			fmt.Println("[]")
			return nil
		}
		raw, err := json.Marshal(entries)
		if err != nil {
			return fmt.Errorf("could not format response payload: %s", err)
		}
		fmt.Println(string(raw))
		return nil
	}

	if len(entries) == 0 {
		fmt.Println("-- no entries to show --")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"TYPE", "OBJECT", "SCORE"})
	for _, entry := range entries {
		table.Append([]string{entry.Type, entry.Object, strconv.Itoa(entry.Reputation)})
	}
	table.Render()

	return nil
}
