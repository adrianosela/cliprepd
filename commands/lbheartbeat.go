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
	Hidden: true,
}

func lbheartbeatHandler(ctx *cli.Context) error {
	client, err := config.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize client: %s", err)
	}
	ok, err := client.LBHeartbeat()
	if err != nil {
		return fmt.Errorf("could not check heartbeat: %s", err)
	}
	if ok {
		fmt.Println("[OK] the server is reachable via HTTP")
	} else {
		fmt.Println("[PROBLEM] the server is NOT reachable via HTTP")
	}
	return nil
}
