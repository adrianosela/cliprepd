package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/adrianosela/cliprepd/config"
	"github.com/olekukonko/tablewriter"
	cli "gopkg.in/urfave/cli.v1"
)

// VersionCmd is the CLI command object for the Version operation
var VersionCmd = cli.Command{
	Name: "version",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "json, j",
			Usage: "print raw json -- don't pretty print",
		},
	},
	Action: versionHandler,
}

func versionHandler(ctx *cli.Context) error {
	client, err := config.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	resp, err := client.Version()
	if err != nil {
		return fmt.Errorf("could not get server version: %s", err)
	}

	if ctx.BoolT("json") {
		raw, err := json.Marshal(resp)
		if err != nil {
			return fmt.Errorf("could not format response payload: %s", err)
		}
		fmt.Println(string(raw))
		return nil
	}

	data := [][]string{
		[]string{"COMMIT", resp.Commit},
		[]string{"VERSION", resp.Version},
		[]string{"SOURCE", resp.Source},
		[]string{"BUILD", resp.Build},
	}
	table := tablewriter.NewWriter(os.Stdout)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return nil
}
