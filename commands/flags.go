package commands

import (
	"fmt"

	cli "gopkg.in/urfave/cli.v1"
)

const mandatory = "[mandatory]"

var (
	// config flags
	pathFlag = cli.StringFlag{
		Name:  "path, p",
		Usage: "where to save config file",
	}
	urlFlag = cli.StringFlag{
		Name:  "url, u",
		Usage: "host URL to use",
	}
	tokenFlag = cli.StringFlag{
		Name:  "token, t",
		Usage: "auth token to use",
	}

	// i/o option flags
	jsonFlag = cli.BoolFlag{
		Name:  "json, j",
		Usage: "print raw json -- don't pretty print",
	}

	// input flags
	objectFlag = cli.StringFlag{
		Name:  "object, o",
		Usage: "object to apply violation to",
	}
	typeFlag = cli.StringFlag{
		Name:  "type, t",
		Usage: "type of object e.g. ip/email",
	}
	suppressRecoveryFlag = cli.IntFlag{
		Name:  "suppress-recovery, r",
		Usage: "seconds before object's reputation begins to heal",
	}
	scoreFlag = cli.IntFlag{
		Name:  "score, s",
		Usage: "reputation score to assign",
	}
	payloadFlag = cli.StringFlag{
		Name:  "payload, p",
		Usage: "path to payload file",
	}
	violationFlag = cli.StringFlag{
		Name:  "violation, v",
		Usage: "name of violation to be applied",
	}
	decayAfterFlag = cli.IntFlag{
		Name:  "decay-after, d",
		Usage: "seconds after which reputation should begin to recover",
	}
)

func withDefault(f cli.StringFlag, def string) cli.StringFlag {
	f.Value = def
	return f
}

func withDefaultInt(f cli.IntFlag, def int) cli.IntFlag {
	f.Value = def
	return f
}

func asMandatory(f cli.StringFlag) cli.StringFlag {
	f.Usage = fmt.Sprintf("%s %s", mandatory, f.Usage)
	return f
}

func asMandatoryInt(f cli.IntFlag) cli.IntFlag {
	f.Usage = fmt.Sprintf("%s %s", mandatory, f.Usage)
	return f
}

func assertSet(ctx *cli.Context, flags ...string) error {
	for _, f := range flags {
		if !ctx.IsSet(f) {
			return fmt.Errorf("missing %s argument \"%s\"", mandatory, f)
		}
	}
	return nil
}
