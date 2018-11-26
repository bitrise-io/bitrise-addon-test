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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
