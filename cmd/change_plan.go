package cmd

import (
	"github.com/bitrise-team/bitrise-add-on-testing-kit/addontester"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	changePlanAppSlug   string
	changePlanPlan      string
	changePlanWithRetry bool
)

var changePlanCmd = &cobra.Command{
	Use:   "change-plan",
	Short: "Test for plan change request",
	Long: `Test whether the developed add-on is capable of handling the plan change request.

The test sends a PUT request to the add-on's /provision/{app_slug} endpoint and expects a success response.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := changePlan()
		if err != nil {
			fail(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(changePlanCmd)

	changePlanCmd.PersistentFlags().StringVar(&changePlanAppSlug, "app-slug", "", "The slug of the app whose add-on's plan gets changed. It gets randomly generated if not given.")
	changePlanCmd.PersistentFlags().StringVar(&changePlanPlan, "plan", "free", "The plan the add-on gets changed to.")
	changePlanCmd.PersistentFlags().BoolVarP(&changePlanWithRetry, "retry", "r", false, "Retry provisioning to test idempotency.")
}

func changePlan() error {
	tester, err := addonTesterFromConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	return tester.ChangePlan(addontester.ChangePlanTesterParams{
		AppSlug:   changePlanAppSlug,
		Plan:      changePlanPlan,
		WithRetry: changePlanWithRetry,
	}, numberOfRetries)
}
