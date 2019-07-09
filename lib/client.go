package lib

import (
	"errors"
	"net/http"
)

// IPrepd is the connection IPrepd service client
type IPrepd struct {
	url        string
	authToken  string
	httpClient *http.Client
}

// NewIPrepd is the default constructor for the Iprepd client
func NewIPrepd(url, token string) (*IPrepd, error) {
	return newIPrepd(url, token, http.DefaultClient)
}

// NewIPrepdWithCustomHTTPClient is the custom-http-client
// constructor for the IPrepd client
func NewIPrepdWithCustomHTTPClient(url, token string, httpClient http.Client) (*IPrepd, error) {
	return newIPrepd(url, token, nil)
}

func newIPrepd(url, token string, httpClient *http.Client) (*IPrepd, error) {
	if url == "" {
		return nil, errors.New("url cannot be empty")
	}
	if token == "" {
		return nil, errors.New("auth token cannot be empty")
	}
	if httpClient == nil {
		return nil, errors.New("http client cannot be empty")
	}
	return &IPrepd{
		url:        url,
		authToken:  token,
		httpClient: httpClient,
	}, nil
}
