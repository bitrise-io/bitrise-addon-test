package addontester

import (
	"fmt"

	"github.com/bitrise-team/bitrise-add-on-testing-kit/addonprovisioner"
	"github.com/bitrise-team/bitrise-add-on-testing-kit/utils"
)

// ChangePlanTesterParams ...
type ChangePlanTesterParams struct {
	AppSlug   string
	Plan      string
	WithRetry bool
}

// ChangePlan ...
func (c *Tester) ChangePlan(params ChangePlanTesterParams) error {
	if len(params.AppSlug) == 0 {
		params.AppSlug, _ = utils.RandomHex(8)
	}

	c.logger.Printf("\nPlan changing details:")
	c.logger.Printf("App slug: %s", params.AppSlug)
	c.logger.Printf("Plan: %s", params.Plan)
	c.logger.Printf("Should retry: %v", params.WithRetry)

	status, body, err := c.addonClient.ChangePlan(addonprovisioner.ChangePlanRequestParams{
		Plan: params.Plan,
	}, params.AppSlug)

	if err != nil {
		return fmt.Errorf("Plan changing failed: %s", err)
	}

	c.logger.Printf("\nResponse status: %d", status)
	c.logger.Printf("Response body: %v\n", body)

	if status < 200 || status > 299 {
		return fmt.Errorf("Plan changing request resulted in a non-2xx response")
	}

	c.logger.Println("\nPlan changing success.")

	return nil
	//TODO: retry logic, sending with wrong header
}
