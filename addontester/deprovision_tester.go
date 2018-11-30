package addontester

import (
	"fmt"

	"github.com/bitrise-team/bitrise-addon-test/utils"
)

// DeprovisionTesterParams ...
type DeprovisionTesterParams struct {
	AppSlug   string
	WithRetry bool
}

// Deprovision ...
func (t *Tester) Deprovision(params DeprovisionTesterParams, remainingRetries int) error {
	if len(params.AppSlug) == 0 {
		params.AppSlug, _ = utils.RandomHex(8)
	}

	t.logger.Printf("\nDeprovisioning details:")
	t.logger.Printf("App slug: %s", params.AppSlug)
	t.logger.Printf("Should retry: %v", params.WithRetry)
	if params.WithRetry {
		t.logger.Printf("No. of test: %d.", numberOfTestsWithRetry-remainingRetries)
	}

	status, body, err := t.addonClient.Deprovision(params.AppSlug)

	if err != nil {
		return fmt.Errorf("Deprovisioning failed: %s", err)
	}

	t.logger.Printf("\nResponse status: %d", status)
	t.logger.Printf("Response body: %v\n", body)

	if status < 200 || status > 299 {
		return fmt.Errorf("Deprovisioning request resulted in a non-2xx response")
	}

	t.logger.Println("\nDeprovisioning success.")

	if params.WithRetry && remainingRetries > 0 {
		return t.Deprovision(params, remainingRetries-1)
	}

	return nil
}
