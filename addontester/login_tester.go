package addontester

import (
	"github.com/bitrise-team/bitrise-add-on-testing-kit/utils"
)

// LoginTesterParams ...
type LoginTesterParams struct {
	AppSlug   string
	BuildSlug string
	WithRetry bool
}

// Login ...
func (c *Tester) Login(params LoginTesterParams, remainingRetries int) error {
	if len(params.AppSlug) == 0 {
		params.AppSlug, _ = utils.RandomHex(8)
	}

	if len(params.BuildSlug) == 0 {
		params.BuildSlug, _ = utils.RandomHex(8)
	}

	c.logger.Printf("\nLogin details:")
	c.logger.Printf("App slug: %s", params.AppSlug)
	c.logger.Printf("Build slug: %s", params.BuildSlug)
	c.logger.Printf("Should retry: %v", params.WithRetry)
	if params.WithRetry {
		c.logger.Printf("No. of test: %d.", numberOfTestsWithRetry-remainingRetries)
	}

	return nil
}
