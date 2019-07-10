package commands

import (
	"fmt"

	"github.com/adrianosela/cliprepd/config"
	cli "gopkg.in/urfave/cli.v1"
)

// LBHeartbeatCmd is the CLI command object for the LBHeartbeat operation
var LBHeartbeatCmd = cli.Command{
	Name:   "lbheartbeat",
	Action: lbheartbeatHandler,
}

func lbheartbeatHandler(ctx *cli.Context) error {
	client, err := config.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	return client.LBHeartbeat()
}
