package commands

import (
	"fmt"

	"github.com/adrianosela/cliprepd/config"
	cli "gopkg.in/urfave/cli.v1"
)

// HeartbeatCmd is the CLI command object for the Heartbeat operation
var HeartbeatCmd = cli.Command{
	Name:   "heartbeat",
	Action: heartbeatHandler,
}

func heartbeatHandler(ctx *cli.Context) error {
	client, err := config.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	return client.Heartbeat()
}
