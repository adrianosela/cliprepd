package lib

import (
	"fmt"
	"net/http"
)

// LBHeartbeat checks whether an IPrepd LB is healthy / reachable
func (c *IPrepd) LBHeartbeat() error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/__lbheartbeat__", c.hostURL), nil)
	if err != nil {
		return fmt.Errorf("could not build http request: %s", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not send http request: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: expected 200, got %d", resp.StatusCode)
	}
	fmt.Println("iprepd deployment is healthy!")
	return nil
}
