package addonprovisioner

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bitrise-io/bitrise-addon-test/utils"
	"github.com/bitrise-io/go-utils/log"
	"github.com/pkg/errors"
)

// Deprovision ...
func (c *Client) Deprovision(appSlug string) (int, string, error) {
	resp, err := c.doRequest("DELETE", fmt.Sprintf("/provision/%s", appSlug), nil)
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
