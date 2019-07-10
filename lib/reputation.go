package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mozilla.org/iprepd"
)

// GetReputation fetches the reputation of a given object and type
func (c *IPrepd) GetReputation(objectType, object string) (*iprepd.Reputation, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/type/%s/%s", c.hostURL, objectType, object), nil)
	if err != nil {
		return nil, fmt.Errorf("could not build http request: %s", err)
	}
	c.addAuth(req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send http request: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non 200 status code received: %d", resp.StatusCode)
	}
	byt, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %s", err)
	}
	var r *iprepd.Reputation
	if err := json.Unmarshal(byt, &r); err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err)
	}
	return r, nil
}

// DeleteReputation deletes the reputation of a given object and type
func (c *IPrepd) DeleteReputation(objectType, object string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/type/%s/%s", c.hostURL, objectType, object), nil)
	if err != nil {
		return fmt.Errorf("could not build http request: %s", err)
	}
	c.addAuth(req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not send http request: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non 200 status code received: %d", resp.StatusCode)
	}
	return nil
}
