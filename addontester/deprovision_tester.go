package addontester

import (
	"fmt"

	"github.com/bitrise-team/bitrise-add-on-testing-kit/utils"
)

// DeprovisionTesterParams ...
type DeprovisionTesterParams struct {
	AppSlug   string
	WithRetry bool
}

// Deprovision ...
func (c *Tester) Deprovision(params DeprovisionTesterParams, remainingRetries int) error {
	if len(params.AppSlug) == 0 {
		params.AppSlug, _ = utils.RandomHex(8)
	}

	c.logger.Printf("\nDeprovisioning details:")
	c.logger.Printf("App slug: %s", params.AppSlug)
	c.logger.Printf("Should retry: %v", params.WithRetry)
	if params.WithRetry {
		c.logger.Printf("No. of test: %d.", 3-remainingRetries)
	}

	status, body, err := c.addonClient.Deprovision(params.AppSlug)

	if err != nil {
		return fmt.Errorf("Deprovisioning failed: %s", err)
	}

	c.logger.Printf("\nResponse status: %d", status)
	c.logger.Printf("Response body: %v\n", body)

	if status < 200 || status > 299 {
		return fmt.Errorf("Deprovisioning request resulted in a non-2xx response")
	}

	c.logger.Println("\nDeprovisioning success.")

	if params.WithRetry && remainingRetries > 0 {
		return c.Deprovision(params, remainingRetries-1)
	}

	return nil
}
