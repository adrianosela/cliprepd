package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mozilla.org/iprepd"
)

// Dump retrieves all reputation entries
func (c *IPrepd) Dump() ([]iprepd.Reputation, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/dump", c.hostURL), nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", clientErrBuildRequest, err)
	}
	c.addAuth(req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", clientErrSendRequest, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %d", clientErrNon200, resp.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", clientErrReadResponse, err)
	}
	var ret []iprepd.Reputation
	if err = json.Unmarshal(bodyBytes, &ret); err != nil {
		return nil, fmt.Errorf("%s: %s", clientErrUnmarshal, err)
	}
	return ret, nil
}
