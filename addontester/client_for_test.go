package addontester_test

import (
	"github.com/bitrise-team/bitrise-addon-test/addonprovisioner"
)

const numberOfRetryTests = 2

type testAddonClient struct {
	addonURL           string
	authToken          string
	ssoSecret          string
	responseStatusCode int
	responseBody       string
	err                error
}

func (c testAddonClient) AddonURL() string {
	return c.addonURL
}
func (c testAddonClient) AuthToken() string {
	return c.authToken
}
func (c testAddonClient) SSOSecret() string {
	return c.ssoSecret
}
func (c testAddonClient) Provision(params addonprovisioner.ProvisionRequestParams) (int, string, error) {
	return c.responseStatusCode, c.responseBody, c.err
}
func (c testAddonClient) Deprovision(appSlug string) (int, string, error) {
	return c.responseStatusCode, c.responseBody, c.err
}
func (c testAddonClient) ChangePlan(params addonprovisioner.ChangePlanRequestParams, appSlug string) (int, string, error) {
	return c.responseStatusCode, c.responseBody, c.err
}
func (c testAddonClient) Login(params addonprovisioner.LoginRequestParams) (int, string, error) {
	return c.responseStatusCode, c.responseBody, c.err
}
