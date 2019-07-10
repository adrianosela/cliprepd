package main

import (
	"fmt"
	"os"

	"github.com/adrianosela/cliprepd/commands"
	"github.com/adrianosela/cliprepd/config"
	cli "gopkg.in/urfave/cli.v1"
)

var version string // injected at build-time

var appflags = []cli.Flag{
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
	commands.ViolationCmd,
	commands.ReputationCmd,
	commands.HeartbeatCmd,
	commands.LBHeartbeatCmd,
	commands.DumpCmd,
	commands.VersionCmd,
	config.ConfigCmd,
}

func main() {
	app := cli.NewApp()
	app.Version = version
	app.EnableBashCompletion = true
	app.Usage = "cli to manage iprepd server"
	app.Flags = appflags
	app.Commands = appcmds
	app.CommandNotFound = func(c *cli.Context, command string) {
		c.App.Run([]string{"help"})
		fmt.Printf("\ncommand \"%s\" does not exist\n", command)
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
