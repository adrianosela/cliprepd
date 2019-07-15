package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

// DefaultConfigFilename is the default filename for the config file
const DefaultConfigFilename = ".repd"

// Config is the iprepd cli configuration
type Config struct {
	HostURL string `json:"host_url"`
	AuthTK  string `json:"auth_token"`
}

// SetConfig writes a configuration file to the given path
func SetConfig(url, tk, path string) error {
	if url == "" {
		return errors.New("url cannot be empty")
	}
	if tk == "" {
		return errors.New("token cannot be empty")
	}
	if path == "" {
		usr, err := user.Current()
		if err != nil {
			return fmt.Errorf("could not get user's home directory: %s", err)
		}
		path = fmt.Sprintf("%s/%s", usr.HomeDir, DefaultConfigFilename)
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create new file %s: %s", path, err)
	}
	byt, err := json.Marshal(&Config{HostURL: url, AuthTK: tk})
	if err != nil {
		return fmt.Errorf("could not marshal configuration file: %s", err)
	}
	if _, err := f.Write(byt); err != nil {
		return fmt.Errorf("could not write configuration file: %s", err)
	}
	return nil
}

// GetConfig returns the configuration at a given path
func GetConfig(path string) (*Config, error) {
	if path == "" {
		usr, err := user.Current()
		if err != nil {
			return nil, fmt.Errorf("could not get user's home directory: %s", err)
		}
		path = fmt.Sprintf("%s/%s", usr.HomeDir, DefaultConfigFilename)
	}
	return readFSConfig(path)
}

func readFSConfig(path string) (*Config, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read configuration file %s: %s", path, err)
	}
	var c *Config
	if err = json.Unmarshal(dat, &c); err != nil {
		return nil, fmt.Errorf("could not unmarshal config: %s", err)
	}
	return c, nil
}
