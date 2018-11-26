package cmd

import (
	"github.com/bitrise-team/bitrise-add-on-testing-kit/addontester"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	provisionAppSlug   string
	provisionAPIToken  string
	provisionPlan      string
	provisionWithRetry bool
)

var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := provision()
		if err != nil {
			fail(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(provisionCmd)

	provisionCmd.PersistentFlags().StringVar(&provisionAppSlug, "app-slug", "", "The slug of the app the add-on gets provisioned to. It gets randomly generated if not given.")
	provisionCmd.PersistentFlags().StringVar(&provisionAPIToken, "api-token", "", "An API token of the app the add-on gets provisioned to. The add-on can behave on behalf of the app using the Bitrise API. It gets randomly generated if not given.")
	provisionCmd.PersistentFlags().StringVar(&provisionPlan, "plan", "free", "The plan of the provisioned add-on.")
	provisionCmd.PersistentFlags().BoolVarP(&provisionWithRetry, "retry", "r", false, "Retry provisioning to test idempotency.")
}

func provision() error {
	tester, err := addonTesterFromConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	return tester.Provision(addontester.ProvisionTesterParams{
		AppSlug:   provisionAppSlug,
		APIToken:  provisionAPIToken,
		Plan:      provisionPlan,
		WithRetry: provisionWithRetry,
	}, numberOfRetries)
}
