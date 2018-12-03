package addonprovisioner

import (
	"io/ioutil"
	"os"

	"github.com/bitrise-io/bitrise-addon-test/utils"
	"github.com/bitrise-io/go-utils/log"
	"github.com/pkg/errors"
)

// ProvisionRequestParams ...
type ProvisionRequestParams struct {
	AppSlug  string `json:"app_slug"`
	AppTitle string `json:"app_title"`
	APIToken string `json:"api_token"`
	Plan     string `json:"plan"`
}

// Provision ...
func (c *Client) Provision(params ProvisionRequestParams) (int, string, error) {
	resp, err := c.doRequest("POST", "/provision", params)
	if err != nil {
		return 0, "", errors.Wrap(err, "Failed to send request")
	}
	defer utils.ResponseBodyCloseWithErrorLog(resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to read response body: %s", err)
		os.Exit(1)
	}
	bodyString := string(bodyBytes)

	return resp.StatusCode, bodyString, nil
}
