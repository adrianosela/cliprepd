package commands

import (
	"fmt"

	"github.com/adrianosela/cliprepd/tool/config"
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
		return fmt.Errorf("could not initialize client: %s", err)
	}
	return client.Version()
}
