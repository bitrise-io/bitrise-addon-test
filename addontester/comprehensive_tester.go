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
func (t *Tester) Comprehensive(params ComprehensiveTesterParams) error {
	if len(params.AppSlug) == 0 {
		params.AppSlug, _ = utils.RandomHex(8)
	}
	if len(params.APIToken) == 0 {
		params.APIToken, _ = utils.RandomHex(8)
	}

	t.logger.Printf("\nDeprovisioning details:")
	t.logger.Printf("App slug: %s", params.AppSlug)
	t.logger.Printf("Build slug: %s", params.BuildSlug)
	t.logger.Printf("API token: %s", params.APIToken)
	t.logger.Printf("Initial plan: %s", params.InitialPlan)
	t.logger.Printf("Plan change to: %s", params.PlanChangeTo)
	t.logger.Printf("Timestamp for SSO: %d", params.Timestamp)

	t.Provision(ProvisionTesterParams{
		AppSlug:   params.AppSlug,
		APIToken:  params.APIToken,
		Plan:      params.InitialPlan,
		WithRetry: true,
	}, numberOfTestsWithRetry)

	t.ChangePlan(ChangePlanTesterParams{
		AppSlug: params.AppSlug,
		Plan:    params.PlanChangeTo,
	}, 0)

	t.Login(LoginTesterParams{
		AppSlug:   params.AppSlug,
		BuildSlug: params.BuildSlug,
		Timestamp: params.Timestamp,
	}, 0)

	t.Deprovision(DeprovisionTesterParams{
		AppSlug:   params.AppSlug,
		WithRetry: true,
	}, 3)

	t.logger.Println("\nComprehensive test success.")

	return nil
}
