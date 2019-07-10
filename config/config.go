package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/adrianosela/cliprepd/lib"
	cli "gopkg.in/urfave/cli.v1"
)

const defaultConfigFilePath = "/.repd"

type config struct {
	HostURL string `json:"host_url"`
	AuthTK  string `json:"auth_token"`
}

// GetClient returns an IPrepd client populated with the correct config
func GetClient(ctx *cli.Context) (*lib.IPrepd, error) {
	cPath := ctx.GlobalString("config")
	if cPath == "" {
		cPath = defaultConfigFilePath
	}
	config, err := readFSConfig(cPath)
	if err != nil {
		return nil, err
	}
	return lib.NewIPrepd(config.HostURL, config.AuthTK, nil)
}

func readFSConfig(path string) (*config, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read configuration file %s: %s", path, err)
	}
	var c *config
	if err = json.Unmarshal(dat, &c); err != nil {
		return nil, fmt.Errorf("could not unmarshal config: %s", err)
	}
	return c, nil
}

func setConfig(url, tk, path string) error {
	if url == "" {
		return errors.New("url cannot be empty")
	}
	if tk == "" {
		return errors.New("token cannot be empty")
	}
	if path == "" {
		path = fmt.Sprintf("%s", defaultConfigFilePath)
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create new file %s: %s", path, err)
	}
	byt, err := json.Marshal(&config{HostURL: url, AuthTK: tk})
	if err != nil {
		return fmt.Errorf("could not marshal configuration file: %s", err)
	}
	if _, err := f.Write(byt); err != nil {
		return fmt.Errorf("could not write configuration file: %s", err)
	}
	return nil
}
