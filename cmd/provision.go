package cmd

import (
	"github.com/bitrise-io/bitrise-addon-test/addontester"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	provisionAppSlug   string
	provisionAppTitle  string
	provisionAPIToken  string
	provisionPlan      string
	provisionWithRetry bool
)

var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Test for deprovision request",
	Long: `Test whether the developed add-on is capable of handling the provisioning request.

The test sends a POST request to the add-on's /provision endpoint and expects a success response.`,
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
	provisionCmd.PersistentFlags().StringVar(&provisionAppSlug, "app-title", "", "The title of the app the add-on gets provisioned to. It gets randomly generated if not given.")
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
		AppTitle:  provisionAppTitle,
		APIToken:  provisionAPIToken,
		Plan:      provisionPlan,
		WithRetry: provisionWithRetry,
	}, numberOfRetries)
}
