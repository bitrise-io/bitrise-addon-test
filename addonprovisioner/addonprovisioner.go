// Ideally we would share this  package with the actual addons-service
// It would have a logger input and an optional debug mode to print out requests.

package addonprovisioner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-team/bitrise-api/utils"
	"github.com/moul/http2curl"
	"github.com/pkg/errors"
)

// ClientConfig ...
type ClientConfig struct {
	Client    *http.Client
	AddonURL  string
	AuthToken string
	SSOSecret string
}

// Client ...
type Client struct {
	client    *http.Client
	addonURL  string
	authToken string
	ssoSecret string
}

func (c *Client) validate() error {
	if c.client == nil {
		c.client = &http.Client{}
	}
	if c.addonURL == "" {
		return errors.New("No Add-on URL specified")
	}
	if c.authToken == "" {
		return errors.New("No authorization token specified")
	}
	if c.ssoSecret == "" {
		return errors.New("No SSO secret specified")
	}
	return nil
}

// NewClient ...
func NewClient(config *ClientConfig) (*Client, error) {
	client := Client{
		client:    config.Client,
		addonURL:  config.AddonURL,
		authToken: config.AuthToken,
		ssoSecret: config.SSOSecret,
	}

	if err := client.validate(); err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *Client) doRequest(method, path string, payload interface{}) (*http.Response, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to JSON serialize")
	}
	req, err := http.NewRequest(method, c.addonURL+path, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authentication", c.authToken)

	command, err := http2curl.GetCurlCommand(req)
	if err != nil {
		return nil, err
	}

	fmt.Println("\nMaking request:")
	fmt.Println(command)

	return c.client.Do(req)
}

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
	_ = string(bodyBytes)

	return resp.StatusCode, "bodyString", nil
}
