package commands

import (
	"github.com/adrianosela/cliprepd/config"
	"github.com/adrianosela/cliprepd/lib"
	cli "gopkg.in/urfave/cli.v1"
)

func getClient(ctx *cli.Context) (*lib.IPrepd, error) {
	cPath := ctx.GlobalString("config")
	config, err := config.GetConfig(cPath)
	if err != nil {
		return nil, err
	}
	return lib.NewIPrepd(config.HostURL, config.AuthTK, nil)
}
