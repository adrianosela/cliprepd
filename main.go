package main

import (
	"log"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

// injected at build time
var buildVersion string

func main() {

	app := cli.NewApp()
	app.Version = buildVersion
	app.EnableBashCompletion = true
	app.Usage = "cli to manage iprepd server"
	app.CommandNotFound = func(c *cli.Context, command string) {
		log.Printf("error: the command provided is not supported: %s", command)
		c.App.Run([]string{"help"})
	}

	app.Commands = []cli.Command{}

	if err := app.Run(os.Args); err != nil {
		log.Printf("error: %s", err)
		os.Exit(1)
	}
}
