package addontester

import (
	"fmt"

	"github.com/bitrise-team/bitrise-addon-test/utils"
	"github.com/pkg/errors"
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
		var err error
		params.AppSlug, err = utils.RandomHex(8)
		if err != nil {
			return fmt.Errorf("Failed to generate app slug: %s", err)
		}
	}
	if len(params.APIToken) == 0 {
		var err error
		params.APIToken, err = utils.RandomHex(8)
		if err != nil {
			return fmt.Errorf("Failed to generate api token: %s", err)
		}
	}

	t.logger.Printf("\nDeprovisioning details:")
	t.logger.Printf("App slug: %s", params.AppSlug)
	t.logger.Printf("Build slug: %s", params.BuildSlug)
	t.logger.Printf("API token: %s", params.APIToken)
	t.logger.Printf("Initial plan: %s", params.InitialPlan)
	t.logger.Printf("Plan change to: %s", params.PlanChangeTo)
	t.logger.Printf("Timestamp for SSO: %d", params.Timestamp)

	err := t.Provision(ProvisionTesterParams{
		AppSlug:   params.AppSlug,
		APIToken:  params.APIToken,
		Plan:      params.InitialPlan,
		WithRetry: true,
	}, numberOfTestsWithRetry)
	if err != nil {
		return errors.WithStack(err)
	}

	err = t.ChangePlan(ChangePlanTesterParams{
		AppSlug: params.AppSlug,
		Plan:    params.PlanChangeTo,
	}, 0)
	if err != nil {
		return errors.WithStack(err)
	}

	err = t.Login(LoginTesterParams{
		AppSlug:   params.AppSlug,
		BuildSlug: params.BuildSlug,
		Timestamp: params.Timestamp,
	}, 0)
	if err != nil {
		return errors.WithStack(err)
	}

	err = t.Deprovision(DeprovisionTesterParams{
		AppSlug:   params.AppSlug,
		WithRetry: true,
	}, 3)
	if err != nil {
		return errors.WithStack(err)
	}

	t.logger.Println("\nComprehensive test success.")

	return nil
}
