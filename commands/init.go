package commands

import (
	"github.com/adrianosela/cliprepd/config"
	cli "gopkg.in/urfave/cli.v1"
)

// InitCmd is the CLI command object for the init operation
var InitCmd = cli.Command{
	Name: "init",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "path, p",
			Usage: "path where to save config",
		},
		cli.StringFlag{
			Name:  "url, u",
			Usage: "host URL to use",
		},
		cli.StringFlag{
			Name:  "token, t",
			Usage: "auth token to use",
		},
	},
	Action: initHandler,
}

func initHandler(ctx *cli.Context) error {
	return config.SetConfig(
		ctx.String("url"),
		ctx.String("token"),
		ctx.String("path"),
	)
}
