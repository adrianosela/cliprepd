package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mozilla.org/iprepd"
)

// Dump retrieves all reputation entries
func (c *IPrepd) Dump() ([]iprepd.Reputation, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/dump", c.hostURL), nil)
	if err != nil {
		return nil, fmt.Errorf("could not build http request: %s", err)
	}
	req.Header.Set("Authorization", c.authTk)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send http request: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("non 200 status code: %s")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %s", err)
	}
	var ret []iprepd.Reputation
	if err = json.Unmarshal(bodyBytes, &ret); err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err)
	}
	return ret, nil
}
