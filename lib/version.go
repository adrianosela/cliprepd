package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// VersionResponse is the response payload from the /__version__ endpoint
type VersionResponse struct {
	Commit  string `json:"commit"`
	Version string `json:"version"`
	Source  string `json:"source"`
	Build   string `json:"build"`
}

// Version retrieves the version of the IPrepd deployment
func (c *IPrepd) Version() (*VersionResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/__version__", c.hostURL), nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", clientErrBuildRequest, err)
	}
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
	var vr *VersionResponse
	if err = json.Unmarshal(bodyBytes, &vr); err != nil {
		return nil, fmt.Errorf("%s: %s", clientErrUnmarshal, err)
	}
	return vr, nil
}
