package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mozilla.org/iprepd"
)

// GetViolations gets all existing violations on the server
func (c *IPrepd) GetViolations() ([]iprepd.Violation, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/violations", c.hostURL), nil)
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
	bodyByt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %s", err)
	}
	var v []iprepd.Violation
	if err = json.Unmarshal(bodyByt, &v); err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err)
	}
	return v, nil
}

// ApplyViolation submits a ViolationRequest to iprepd
func (c *IPrepd) ApplyViolation(vr *iprepd.ViolationRequest) error {
	byt, err := json.Marshal(&vr)
	if err != nil {
		return fmt.Errorf("could not marshal payload: %s", err)
	}
	req, err := http.NewRequest(http.MethodPut,
		fmt.Sprintf("%s/violations/type/%s/%s", c.hostURL, vr.Type, vr.Object), bytes.NewBuffer(byt))
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

// BatchApplyViolation submits a batch of ViolationRequests to iprepd
func (c *IPrepd) BatchApplyViolation(typ string, vrs []iprepd.ViolationRequest) error {
	byt, err := json.Marshal(&vrs)
	if err != nil {
		return fmt.Errorf("could not marshal payload: %s", err)
	}
	req, err := http.NewRequest(http.MethodPut,
		fmt.Sprintf("%s/violations/type/%s", c.hostURL, typ), bytes.NewBuffer(byt))
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
