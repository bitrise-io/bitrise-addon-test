package addontester

import (
	"encoding/json"
	"fmt"

	"github.com/bitrise-io/bitrise-addon-test/addonprovisioner"
	"github.com/bitrise-io/bitrise-addon-test/utils"
)

// ProvisionTesterParams ...
type ProvisionTesterParams struct {
	AppSlug   string
	AppTitle  string
	APIToken  string
	Plan      string
	WithRetry bool
}

type env struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type provisionResp struct {
	Envs []env `json:"envs"`
}

// Provision ...
func (t *Tester) Provision(params ProvisionTesterParams, remainingRetries int) error {

	if len(params.AppSlug) == 0 {
		var err error
		params.AppSlug, err = utils.RandomHex(8)
		if err != nil {
			return fmt.Errorf("Failed to generate app slug: %s", err)
		}
	}

	if len(params.AppTitle) == 0 {
		params.AppTitle = fmt.Sprintf("Title of app #%s", params.AppSlug)
	}

	if len(params.APIToken) == 0 {
		var err error
		params.APIToken, err = utils.RandomHex(8)
		if err != nil {
			return fmt.Errorf("Failed to generate api token: %s", err)
		}
	}

	t.logger.Printf("\nProvisioning details:")
	t.logger.Printf("App slug: %s", params.AppSlug)
	t.logger.Printf("App title: %s", params.AppTitle)
	t.logger.Printf("API token: %s", params.APIToken)
	t.logger.Printf("Plan: %s", params.Plan)
	t.logger.Printf("Should retry: %v", params.WithRetry)
	if params.WithRetry {
		t.logger.Printf("No. of test: %d.", numberOfTestsWithRetry-remainingRetries)
	}

	status, body, err := t.addonClient.Provision(addonprovisioner.ProvisionRequestParams{
		AppSlug:  params.AppSlug,
		AppTitle: params.AppTitle,
		APIToken: params.APIToken,
		Plan:     params.Plan,
	})

	if err != nil {
		return fmt.Errorf("Provisioning failed: %s", err)
	}

	t.logger.Printf("\nResponse status: %d", status)
	t.logger.Printf("Response body: %v\n", body)

	if status < 200 || status > 299 {
		return fmt.Errorf("Provisioning request resulted in a non-2xx response")
	}

	var pr provisionResp

	if err := json.Unmarshal([]byte(body), &pr); err != nil {
		return fmt.Errorf("JSON parsing of response failed: %s", err)
	}

	if len(pr.Envs) > 0 {
		for _, e := range pr.Envs {
			if len(e.Key) == 0 {
				return fmt.Errorf("ENV var key is not present: %v", e)
			}

			if len(e.Value) == 0 {
				return fmt.Errorf("ENV var value is not present: %v", e)
			}

			t.logger.Printf("ENV var processed succesfully: %s: %s", e.Key, e.Value)
		}
	} else {
		t.logger.Printf("No ENV vars to check in response")
	}

	t.logger.Println("\nProvisioning success.")

	if params.WithRetry && remainingRetries > 0 {
		return t.Provision(params, remainingRetries-1)
	}

	return nil
}
