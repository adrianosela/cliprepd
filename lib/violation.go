package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mozilla.org/iprepd"
)

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
	req.Header.Set("Authorization", c.authTk)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not send http request: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non 200 status code received: %d", resp.StatusCode)
	}
	return nil
}
