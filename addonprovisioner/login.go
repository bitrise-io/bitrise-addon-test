package addonprovisioner

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/bitrise-io/bitrise-addon-test/utils"
	"github.com/bitrise-io/go-utils/log"
	"github.com/moul/http2curl"
	"github.com/pkg/errors"
)

// LoginRequestParams ...
type LoginRequestParams struct {
	AppSlug   string `json:"app_slug"`
	BuildSlug string `json:"build_slug"`
	Timestamp string `json:"timestamp"`
}

// Login ...
func (c *Client) Login(params LoginRequestParams) (int, string, error) {
	s := sha1.New()
	_, err := s.Write([]byte(fmt.Sprintf("%s:%s:%s", params.AppSlug, c.SSOSecret(), params.Timestamp)))
	if err != nil {
		log.Errorf("Failed to write into sha1 buffer, error: %s", err)
		return 0, "", errors.Wrap(err, "Failed to submit form")
	}
	refToken := fmt.Sprintf("%x", s.Sum(nil))

	formData := strings.NewReader(url.Values{
		"app_slug":  {params.AppSlug},
		"token":     {refToken},
		"timestamp": {params.Timestamp},
	}.Encode())
	req, err := http.NewRequest("POST", c.addonURL+fmt.Sprintf("/login?build_slug=%s", params.BuildSlug), formData)
	if err != nil {
		return 0, "", errors.Wrap(err, "Failed to submit form")
	}
	req.Header.Add("Authentication", c.authToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	command, err := http2curl.GetCurlCommand(req)
	if err != nil {
		return 0, "", err
	}

	fmt.Println("\nMaking request:")
	fmt.Println(command)
	resp, err := c.client.Do(req)

	defer utils.ResponseBodyCloseWithErrorLog(resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to read response body: %s", err)
		os.Exit(1)
	}
	bodyString := string(bodyBytes)

	return resp.StatusCode, bodyString, nil
}
