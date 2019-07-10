package commands

import cli "gopkg.in/urfave/cli.v1"

var jsonFlag = cli.BoolFlag{
	Name:  "json, j",
	Usage: "print raw json -- don't pretty print",
}
