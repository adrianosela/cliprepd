package commands

import (
	"fmt"
	"os"

	"github.com/adrianosela/cliprepd/config"
	"github.com/olekukonko/tablewriter"
	cli "gopkg.in/urfave/cli.v1"
)

// ConfigCmd is the CLI command object for the config operation
var ConfigCmd = cli.Command{
	Name:  "config",
	Usage: "configure command line tool",
	Subcommands: []cli.Command{
		{
			Name:  "set",
			Usage: "create configuration file with given options",
			Flags: []cli.Flag{
				withDefault(pathFlag, config.DefaultConfigFilePath),
				asMandatory(urlFlag),
				asMandatory(tokenFlag),
			},
			Before: configSetValidator,
			Action: configSetHandler,
		},
		{
			Name:  "show",
			Usage: "show contents of set configuration file",
			Flags: []cli.Flag{
				withDefault(pathFlag, config.DefaultConfigFilePath),
			},
			Action: configShowHandler,
		},
	},
}

func configSetValidator(ctx *cli.Context) error {
	return assertSet(ctx, "url", "token", "path")
}

func configSetHandler(ctx *cli.Context) error {
	if err := config.SetConfig(
		ctx.String("url"),
		ctx.String("token"),
		ctx.String("path"),
	); err != nil {
		return fmt.Errorf("could not set configuration: %s", err)
	}
	return nil
}

func configShowHandler(ctx *cli.Context) error {
	path := ctx.String("path")
	c, err := config.GetConfig(path)
	if err != nil {
		return fmt.Errorf("could not retrive configuration from %s: %s", path, err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Append([]string{"HOST_URL", c.HostURL})
	table.Append([]string{"AUTH_TK", c.AuthTK})
	table.Render()
	return nil
}
