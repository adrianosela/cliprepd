package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adrianosela/cliprepd/commands"
	cli "gopkg.in/urfave/cli.v1"
)

var version string // injected at build-time

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

func main() {
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
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
