package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mozilla.org/iprepd"
)

// GetReputation fetches the reputation of a given object and type
func (c *IPrepd) GetReputation(objectType, object string) (*iprepd.Reputation, error) {
	if object == "" {
		return nil, errors.New(clientErrObjectEmpty)
	}
	if objectType == "" {
		return nil, errors.New(clientErrObjectTypeEmpty)
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/type/%s/%s", c.hostURL, objectType, object), nil)
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
	byt, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", clientErrReadResponse, err)
	}
	var r *iprepd.Reputation
	if err := json.Unmarshal(byt, &r); err != nil {
		return nil, fmt.Errorf("%s: %s", clientErrUnmarshal, err)
	}
	return r, nil
}

// SetReputation updates the reputation of a given object and type to a given score
func (c *IPrepd) SetReputation(r *iprepd.Reputation) error {
	if r == nil {
		return errors.New(clientErrReputationNil)
	}
	if r.Object == "" {
		return errors.New(clientErrObjectEmpty)
	}
	if r.Type == "" {
		return errors.New(clientErrObjectTypeEmpty)
	}
	byt, err := json.Marshal(&r)
	if err != nil {
		return fmt.Errorf("%s: %s", clientErrMarshal, err)
	}
	req, err := http.NewRequest(http.MethodPut,
		fmt.Sprintf("%s/type/%s/%s", c.hostURL, r.Type, r.Object), bytes.NewBuffer(byt))
	if err != nil {
		return fmt.Errorf("%s: %s", clientErrMarshal, err)
	}
	c.addAuth(req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s: %s", clientErrSendRequest, err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %d", clientErrNon200, resp.StatusCode)
	}
	return nil
}

// DeleteReputation deletes the reputation of a given object and type
func (c *IPrepd) DeleteReputation(objectType, object string) error {
	if object == "" {
		return errors.New(clientErrObjectEmpty)
	}
	if objectType == "" {
		return errors.New(clientErrObjectTypeEmpty)
	}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/type/%s/%s", c.hostURL, objectType, object), nil)
	if err != nil {
		return fmt.Errorf("%s: %s", clientErrBuildRequest, err)
	}
	c.addAuth(req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s: %s", clientErrSendRequest, err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %d", clientErrNon200, resp.StatusCode)
	}
	return nil
}
