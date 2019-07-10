package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	cli "gopkg.in/urfave/cli.v1"
)

// ConfigCmd is the CLI command object for the config operation
var ConfigCmd = cli.Command{
	Name: "config",
	Subcommands: []cli.Command{
		cli.Command{
			Name: "set",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path, p",
					Usage: "where to save config file",
					Value: defaultConfigFilePath,
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
			Before: setFlagValidator,
			Action: setHandler,
		},
		cli.Command{
			Name: "show",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path, p",
					Usage: "configuration file path",
					Value: defaultConfigFilePath,
				},
			},
			Action: showHandler,
		},
	},
}

func setFlagValidator(ctx *cli.Context) error {
	if ctx.String("url") == "" {
		return errors.New("missing [mandatory] argument \"url\"")
	}
	if ctx.String("token") == "" {
		return errors.New("missing [mandatory] argument \"token\"")
	}
	return nil
}

func setHandler(ctx *cli.Context) error {
	if err := setConfig(
		ctx.String("url"),
		ctx.String("token"),
		ctx.String("path"),
	); err != nil {
		return fmt.Errorf("could not set configuration: %s", err)
	}
	return nil
}

func showHandler(ctx *cli.Context) error {
	path := ctx.String("path")
	c, err := readFSConfig(path)
	if err != nil {
		return fmt.Errorf("could not retrive configuration from %s: %s", path, err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Append([]string{"HOST_URL", c.HostURL})
	table.Append([]string{"AUTH_TK", c.AuthTK})
	table.Render()
	return nil
}
