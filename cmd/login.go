package cmd

import (
	"github.com/bitrise-team/bitrise-addon-test/addontester"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	loginAppSlug   string
	loginBuildSlug string
	loginTimestamp int64
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Test for SSO login request",
	Long: `Test whether the developed add-on is capable of handling the SSO login request.

The test sends a POST request to the add-on's /login endpoint with an URL encoded form body. Expects an HTML response with 200 code.`,
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
	loginCmd.PersistentFlags().Int64Var(&loginTimestamp, "timestamp", 0, "Timestamp for SSO login token generation")
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
	}, numberOfRetries)
}
