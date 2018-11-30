package cmd

import (
	"github.com/bitrise-team/bitrise-addon-test/addontester"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	deprovisionAppSlug   string
	deprovisionWithRetry bool
)

var deprovisionCmd = &cobra.Command{
	Use:   "deprovision",
	Short: "Test for deprovision request",
	Long: `Test whether the developed add-on is capable of handling the deprovisioning request.

The test sends a DELETE request to the add-on's /provision/{app_slug} endpoint and expects a success response.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := deprovision()
		if err != nil {
			fail(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deprovisionCmd)

	deprovisionCmd.PersistentFlags().StringVar(&deprovisionAppSlug, "app-slug", "", "The slug of the app the add-on gets deprovisioned to. It gets randomly generated if not given.")
	deprovisionCmd.PersistentFlags().BoolVarP(&deprovisionWithRetry, "retry", "r", false, "Retry deprovisioning to test idempotency.")
}

func deprovision() error {
	tester, err := addonTesterFromConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	return tester.Deprovision(addontester.DeprovisionTesterParams{
		AppSlug:   deprovisionAppSlug,
		WithRetry: deprovisionWithRetry,
	}, numberOfRetries)
}
