package cmd

import (
	"github.com/bitrise-team/bitrise-add-on-testing-kit/addontester"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	loginAppSlug   string
	loginBuildSlug string
	loginTimestamp int64
	loginWithRetry bool
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := login()
		if err != nil {
			fail(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.PersistentFlags().StringVar(&loginAppSlug, "app-slug", "", "The slug of the app the add-on gets provisioned to. It gets randomly generated if not given.")
	loginCmd.PersistentFlags().StringVar(&loginBuildSlug, "build-slug", "", "The slug of the build")
	loginCmd.PersistentFlags().Int64Var(&loginTimestamp, "timestamp", 0, "The slug of the build")
	loginCmd.PersistentFlags().BoolVarP(&loginWithRetry, "retry", "r", false, "Retry provisioning to test idempotency.")
}

func login() error {
	tester, err := addonTesterFromConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	return tester.Login(addontester.LoginTesterParams{
		AppSlug:   loginAppSlug,
		BuildSlug: loginBuildSlug,
		Timestamp: loginTimestamp,
		WithRetry: loginWithRetry,
	}, numberOfRetries)
}
