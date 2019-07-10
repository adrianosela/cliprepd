package lib

import (
	"fmt"
	"net/http"
)

// Heartbeat checks whether an IPrepd deployment is healthy / reachable
func (c *IPrepd) Heartbeat() (bool, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/__heartbeat__", c.hostURL), nil)
	if err != nil {
		return false, fmt.Errorf("could not build http request: %s", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("could not send http request: %s", err)
	}
	return (resp.StatusCode == http.StatusOK), nil
}
