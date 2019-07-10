package commands

import (
	"errors"
	"fmt"

	"github.com/adrianosela/cliprepd/config"
	cli "gopkg.in/urfave/cli.v1"
)

// VersionCmd is the CLI command object for the Version operation
var VersionCmd = cli.Command{
	Name:   "version",
	Action: versionHandler,
}

func versionHandler(ctx *cli.Context) error {
	vv := ctx.GlobalBool("verbose")

	client, err := config.GetClient(ctx)
	if err != nil {
		msg := "error initializing IPrepd client"
		if vv {
			return fmt.Errorf("%s: %s", msg, err)
		}
		return errors.New(msg)
	}
	resp, err := client.Version()
	if err != nil {
		return fmt.Errorf("could not retrieve version")
	}

	fmt.Println(*resp)
	return nil
}
