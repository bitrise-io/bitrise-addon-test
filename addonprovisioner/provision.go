package addonprovisioner

import (
	"io/ioutil"
	"os"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-team/bitrise-add-on-testing-kit/utils"
	"github.com/pkg/errors"
)

// ProvisionParams ...
type ProvisionParams struct {
	AppSlug  string `json:"app_slug"`
	APIToken string `json:"api_token"`
	Plan     string `json:"plan"`
}

// Provision ...
func (c *Client) Provision(params ProvisionParams) (int, string, error) {
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
