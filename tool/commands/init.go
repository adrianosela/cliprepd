package commands

import (
	"github.com/adrianosela/cliprepd/tool/config"
	cli "gopkg.in/urfave/cli.v1"
)

// InitCmd is the CLI command object for the init operation
var InitCmd = cli.Command{
	Name:   "init",
	Action: initHandler,
}

func initHandler(ctx *cli.Context) error {
	return config.SetConfig(
		ctx.GlobalString("url"),
		ctx.GlobalString("token"),
		ctx.GlobalString("config"),
	)
}
