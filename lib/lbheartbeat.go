package lib

import (
	"fmt"
	"net/http"
)

// LBHeartbeat checks whether an IPrepd LB is healthy / reachable
func (c *IPrepd) LBHeartbeat() (bool, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/__lbheartbeat__", c.hostURL), nil)
	if err != nil {
		return false, fmt.Errorf("%s: %s", clientErrBuildRequest, err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("%s: %s", clientErrSendRequest, err)
	}
	return (resp.StatusCode == http.StatusOK), nil
}
