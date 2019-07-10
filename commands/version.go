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
	Name:  "version",
	Usage: "get the iprepd version of the server",
	Flags: []cli.Flag{
		jsonFlag,
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
		{"COMMIT", resp.Commit},
		{"VERSION", resp.Version},
		{"SOURCE", resp.Source},
		{"BUILD", resp.Build},
	}
	table := tablewriter.NewWriter(os.Stdout)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return nil
}
