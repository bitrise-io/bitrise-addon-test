package cmd

import (
	"log"
	"os"

	"github.com/bitrise-team/bitrise-add-on-testing-kit/addonprovisiontester"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		tester, _ := addonprovisiontester.New(
			&addonprovisiontester.Config{
				AddonURL:  viper.Get("addon-url").(string),
				AuthToken: viper.Get("auth-token").(string),
				SSOSecret: viper.Get("sso-secret").(string),
				Logger:    log.New(os.Stdout, "", 0),
			})

		err := tester.Provision(addonprovisiontester.ProvisionParams{
			AppSlug:  appSlug,
			APIToken: apiToken,
			Plan:     plan,
		})

		if err != nil {
			fail(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(provisionCmd)

	provisionCmd.PersistentFlags().StringVar(&appSlug, "app-slug", "", "The slug of the app the add-on gets provisioned to. It gets randomly generated if not given.")
	provisionCmd.PersistentFlags().StringVar(&apiToken, "api-token", "", "An API token of the app the add-on gets provisioned to. The add-on can behave on behalf of the app using the Bitrise API. It gets randomly generated if not given.")
	provisionCmd.PersistentFlags().StringVar(&plan, "plan", "free", "The plan of the provisioned add-on.")
	provisionCmd.PersistentFlags().BoolVarP(&withRetry, "retry", "r", false, "Retry provisioning  to test idempotency")
}
