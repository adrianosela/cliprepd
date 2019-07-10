package app

import (
	"log"

	"github.com/adrianosela/cliprepd/commands"
	cli "gopkg.in/urfave/cli.v1"
)

var appflags = []cli.Flag{
	cli.BoolFlag{
		Name:  "verbose, vv",
		Usage: "display all errors fully",
	},
	cli.BoolFlag{
		Name:  "dry-run, d",
		Usage: "show, dont do",
	},
	cli.StringFlag{
		Name:  "config, c",
		Usage: "override default config file path",
	},
}

var appcmds = []cli.Command{
	commands.InitCmd,
	commands.HeartbeatCmd,
	commands.LBHeartbeatCmd,
	commands.VersionCmd,
}

// New returns the urfave/cli app for the IPrepd client
func New(version string) *cli.App {
	app := cli.NewApp()
	app.Version = version
	app.EnableBashCompletion = true
	app.Usage = "cli to manage iprepd server"
	app.Flags = appflags
	app.Commands = appcmds
	app.CommandNotFound = func(c *cli.Context, command string) {
		log.Printf("error: the command provided is not supported: %s", command)
		c.App.Run([]string{"help"})
	}
	return app
}
