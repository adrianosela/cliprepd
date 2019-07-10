package commands

import (
	"encoding/json"
	"fmt"

	"github.com/adrianosela/cliprepd/clierr"
	"github.com/adrianosela/cliprepd/config"
	cli "gopkg.in/urfave/cli.v1"
)

// VersionCmd is the CLI command object for the Version operation
var VersionCmd = cli.Command{
	Name:   "version",
	Action: versionHandler,
}

func versionHandler(ctx *cli.Context) error {
	client, err := config.GetClient(ctx)
	if err != nil {
		return clierr.Handle(ctx, clierr.ErrCouldNotInit, err)
	}
	resp, err := client.Version()
	if err != nil {
		return fmt.Errorf("could not retrieve version")
	}

	if ctx.GlobalBool("json") {
		raw, err := json.Marshal(resp)
		if err != nil {
			return err // fixme
		}
		fmt.Println(string(raw))
		return nil
	}

	// FIXME
	return nil
}
