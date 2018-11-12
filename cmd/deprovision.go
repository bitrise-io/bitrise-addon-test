package cmd

import (
	"github.com/bitrise-team/bitrise-add-on-testing-kit/addontester"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	deprovisionAppSlug   string
	deprovisionWithRetry bool
)

var deprovisionCmd = &cobra.Command{
	Use:   "deprovision",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	deprovisionCmd.PersistentFlags().BoolVarP(&deprovisionWithRetry, "retry", "r", false, "Retry deprovisioning  to test idempotency")
}

func deprovision() error {
	tester, err := addonTesterFromConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	return tester.Deprovision(addontester.DeprovisionTesterParams{AppSlug: deprovisionAppSlug})
}
