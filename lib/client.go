package lib

import (
	"errors"
	"net/http"
)

// IPrepd is the iprepd service client
type IPrepd struct {
	hostURL    string
	authStr    string
	httpClient *http.Client
}

// NewIPrepd is the default constructor for the client
func NewIPrepd(url, token string, httpClient *http.Client) (*IPrepd, error) {
	if url == "" {
		return nil, errors.New(clientErrURLEmpty)
	}
	if token == "" {
		return nil, errors.New(clientErrAuthEmpty)
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &IPrepd{
		hostURL:    url,
		authStr:    token,
		httpClient: httpClient,
	}, nil
}

func (c *IPrepd) addAuth(r *http.Request) {
	r.Header.Set("Authorization", c.authStr)
}
