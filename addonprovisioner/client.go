// Ideally we would share this  package with the actual addons-service
// It would have a logger input and an optional debug mode to print out requests.

package addonprovisioner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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

// AddonClientInterface ...
type AddonClientInterface interface {
	AddonURL() string
	AuthToken() string
	SSOSecret() string
	Provision(params ProvisionRequestParams) (int, string, error)
	Deprovision(appSlug string) (int, string, error)
	ChangePlan(params ChangePlanRequestParams, appSlug string) (int, string, error)
	Login(params LoginRequestParams) (int, string, error)
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

// AddonURL ...
func (c Client) AddonURL() string {
	return c.addonURL
}

// AuthToken ...
func (c Client) AuthToken() string {
	return c.authToken
}

// SSOSecret ...
func (c Client) SSOSecret() string {
	return c.ssoSecret
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
