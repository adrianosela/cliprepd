package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Version retrieves the version of the IPrepd deployment
func (c *IPrepd) Version() error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/__version__", c.hostURL), nil)
	if err != nil {
		return fmt.Errorf("could not build http request: %s", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not send http request: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non 200 status code: %s", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("could not read response body: %s", err)
	}
	fmt.Print(string(bodyBytes))
	return nil
}
