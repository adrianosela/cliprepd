package commands

import (
	"fmt"
	"os/user"

	"github.com/adrianosela/cliprepd/config"
	"github.com/adrianosela/cliprepd/lib"
	cli "gopkg.in/urfave/cli.v1"
)

func getClient(ctx *cli.Context) (*lib.IPrepd, error) {
	cPath := ctx.GlobalString("config")
	if cPath == "" {
		usr, err := user.Current()
		if err != nil {
			return nil, fmt.Errorf("could not get user's home directory: %s", err)
		}
		cPath = fmt.Sprintf("%s/%s", usr.HomeDir, config.DefaultConfigFilename)
	}
	config, err := config.GetConfig(cPath)
	if err != nil {
		return nil, err
	}
	return lib.NewIPrepd(config.HostURL, config.AuthTK, nil)
}
