package lib

import (
	"errors"
	"net/http"
)

// IPrepd is the connection IPrepd service client
type IPrepd struct {
	hostURL    string
	authTk     string
	httpClient *http.Client
}

// NewIPrepd is the default constructor for the Iprepd client
func NewIPrepd(url, token string, httpClient *http.Client) (*IPrepd, error) {
	if url == "" {
		return nil, errors.New("url cannot be empty")
	}
	if token == "" {
		return nil, errors.New("auth token cannot be empty")
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &IPrepd{
		hostURL:    url,
		authTk:     token,
		httpClient: httpClient,
	}, nil
}

func (c *IPrepd) addAuth(r *http.Request) {
	r.Header.Set("Authorization", c.authTk)
}
