package tool

import (
	"log"

	"github.com/adrianosela/cliprepd/tool/commands"
	cli "gopkg.in/urfave/cli.v1"
)

// GetApp returns the urfave/cli app for the IPrepd client
func GetApp(version string) *cli.App {
	app := cli.NewApp()
	app.Version = version
	app.EnableBashCompletion = true
	app.Usage = "cli to manage iprepd server"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "override the config file's path",
		},
		cli.StringFlag{
			Name:   "url, u",
			EnvVar: "IPREPD_HOST_URL",
			Usage:  "override the config file's host URL value",
		},
		cli.StringFlag{
			Name:   "token, t",
			EnvVar: "IPREPD_AUTH_TOKEN",
			Usage:  "override the config file's auth token value",
		},
		cli.BoolFlag{
			Name:  "dry-run, d",
			Usage: "show, dont do",
		},
	}
	app.Commands = []cli.Command{
		commands.InitCmd,
		commands.HeartbeatCmd,
		commands.LBHeartbeatCmd,
		commands.VersionCmd,
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		log.Printf("error: the command provided is not supported: %s", command)
		c.App.Run([]string{"help"})
	}
	return app
}
