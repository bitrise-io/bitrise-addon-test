package addontester

import (
	"github.com/bitrise-team/bitrise-add-on-testing-kit/utils"
)

// ComprehensiveTesterParams ...
type ComprehensiveTesterParams struct {
	AppSlug      string
	BuildSlug    string
	InitialPlan  string
	PlanChangeTo string
	APIToken     string
	Timestamp    int64
}

// Comprehensive ...
func (c *Tester) Comprehensive(params ComprehensiveTesterParams) error {
	if len(params.AppSlug) == 0 {
		params.AppSlug, _ = utils.RandomHex(8)
	}
	if len(params.APIToken) == 0 {
		params.APIToken, _ = utils.RandomHex(8)
	}

	c.logger.Printf("\nDeprovisioning details:")
	c.logger.Printf("App slug: %s", params.AppSlug)
	c.logger.Printf("Build slug: %s", params.BuildSlug)
	c.logger.Printf("API token: %s", params.APIToken)
	c.logger.Printf("Initial plan: %s", params.InitialPlan)
	c.logger.Printf("Plan change to: %s", params.PlanChangeTo)
	c.logger.Printf("Timestamp for SSO: %d", params.Timestamp)

	c.Provision(ProvisionTesterParams{
		AppSlug:   params.AppSlug,
		APIToken:  params.APIToken,
		Plan:      params.InitialPlan,
		WithRetry: true,
	}, numberOfTestsWithRetry)

	c.ChangePlan(ChangePlanTesterParams{
		AppSlug: params.AppSlug,
		Plan:    params.PlanChangeTo,
	}, 0)

	c.Login(LoginTesterParams{
		AppSlug:   params.AppSlug,
		BuildSlug: params.BuildSlug,
		Timestamp: params.Timestamp,
	}, 0)

	c.Deprovision(DeprovisionTesterParams{
		AppSlug:   params.AppSlug,
		WithRetry: true,
	}, 3)

	c.logger.Println("\nComprehensive test success.")

	return nil
}
