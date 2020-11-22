package addontester_test

import (
	"github.com/bitrise-io/bitrise-addon-test/addonprovisioner"
)

const numberOfRetryTests = 2

type testAddonClient struct {
	numberOfProvisionRequestsCalled int
	addonURL                        string
	authToken                       string
	ssoSecret                       string
	responseStatusCode              int
	retryResponseStatusCode         int
	responseBody                    string
	err                             error
	loginRequestInfos               addonprovisioner.LoginRequestInfos
}

func (c *testAddonClient) AddonURL() string {
	return c.addonURL
}

func (c *testAddonClient) AuthToken() string {
	return c.authToken
}

func (c *testAddonClient) SSOSecret() string {
	return c.ssoSecret
}

func (c *testAddonClient) Provision(params addonprovisioner.ProvisionRequestParams) (int, string, error) {
	c.numberOfProvisionRequestsCalled = c.numberOfProvisionRequestsCalled + 1
	if c.numberOfProvisionRequestsCalled == 1 {
		return c.responseStatusCode, c.responseBody, c.err
	}
	return c.retryResponseStatusCode, c.responseBody, c.err
}

func (c *testAddonClient) Deprovision(appSlug string) (int, string, error) {
	return c.responseStatusCode, c.responseBody, c.err
}

func (c *testAddonClient) ChangePlan(params addonprovisioner.ChangePlanRequestParams, appSlug string) (int, string, error) {
	return c.responseStatusCode, c.responseBody, c.err
}

func (c *testAddonClient) Login(params addonprovisioner.LoginRequestParams) (int, string, error) {
	return c.responseStatusCode, c.responseBody, c.err
}

func (c *testAddonClient) Login(params addonprovisioner.LoginRequestParams) (int, string, error) {
	return c.responseStatusCode, c.responseBody, c.err
}

func (c *testAddonClient) LoginRequestInfos(params addonprovisioner.LoginRequestParams) (addonprovisioner.LoginRequestInfos, error) {
	return c.loginRequestInfos, c.err
}
