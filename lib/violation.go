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

// GetViolations gets all existing violations on the server
func (c *IPrepd) GetViolations() ([]iprepd.Violation, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/violations", c.hostURL), nil)
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
	bodyByt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", clientErrReadResponse, err)
	}
	var v []iprepd.Violation
	if err = json.Unmarshal(bodyByt, &v); err != nil {
		return nil, fmt.Errorf("%s: %s", clientErrUnmarshal, err)
	}
	return v, nil
}

// ApplyViolation submits a ViolationRequest to iprepd
func (c *IPrepd) ApplyViolation(vr *iprepd.ViolationRequest) error {
	if vr == nil {
		return errors.New(clientErrViolationRequestNil)
	}
	if vr.Object == "" {
		return errors.New(clientErrObjectEmpty)
	}
	if vr.Type == "" {
		return errors.New(clientErrObjectTypeEmpty)
	}
	if vr.Violation == "" {
		return errors.New(clientErrViolationEmpty)
	}
	byt, err := json.Marshal(&vr)
	if err != nil {
		return fmt.Errorf("%s: %s", clientErrMarshal, err)
	}
	req, err := http.NewRequest(http.MethodPut,
		fmt.Sprintf("%s/violations/type/%s/%s", c.hostURL, vr.Type, vr.Object), bytes.NewBuffer(byt))
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

// BatchApplyViolation submits a batch of ViolationRequests to iprepd
func (c *IPrepd) BatchApplyViolation(typ string, vrs []iprepd.ViolationRequest) error {
	if typ == "" {
		return errors.New(clientErrObjectTypeEmpty)
	}
	if len(vrs) == 0 {
		return nil
	}
	byt, err := json.Marshal(&vrs)
	if err != nil {
		return fmt.Errorf("%s: %s", clientErrMarshal, err)
	}
	req, err := http.NewRequest(http.MethodPut,
		fmt.Sprintf("%s/violations/type/%s", c.hostURL, typ), bytes.NewBuffer(byt))
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
