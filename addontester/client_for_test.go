package addontester_test

import (
	"github.com/bitrise-team/bitrise-add-on-testing-kit/addonprovisioner"
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

func (p testAddonClient) AddonURL() string {
	return p.addonURL
}
func (p testAddonClient) AuthToken() string {
	return p.authToken
}
func (p testAddonClient) SSOSecret() string {
	return p.ssoSecret
}
func (p testAddonClient) Provision(params addonprovisioner.ProvisionRequestParams) (int, string, error) {
	return p.responseStatusCode, p.responseBody, p.err
}
func (p testAddonClient) Deprovision(appSlug string) (int, string, error) {
	return p.responseStatusCode, p.responseBody, p.err
}
func (p testAddonClient) ChangePlan(params addonprovisioner.ChangePlanRequestParams, appSlug string) (int, string, error) {
	return p.responseStatusCode, p.responseBody, p.err
}
func (p testAddonClient) Login(params addonprovisioner.LoginRequestParams) (int, string, error) {
	return p.responseStatusCode, p.responseBody, p.err
}
