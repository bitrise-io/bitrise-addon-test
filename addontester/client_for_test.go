package addontester_test

import (
	"fmt"

	"github.com/bitrise-team/bitrise-add-on-testing-kit/addonprovisioner"
)

type testAddonClient struct {
	addonURL           string
	authToken          string
	ssoSecret          string
	responseStatusCode int
	responseBody       string
	err                error
}

func (p testAddonClient) AddonURL() string {
	return p.addonURL
}
func (p testAddonClient) AuthToken() string {
	return p.authToken
}
func (p testAddonClient) SSOSecret() string {
	return p.ssoSecret
}
func (p testAddonClient) Provision(params addonprovisioner.ProvisionParams) (int, string, error) {
	fmt.Println("asdf")
	return p.responseStatusCode, p.responseBody, p.err
}
