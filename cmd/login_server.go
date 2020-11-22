package cmd

import (
	"github.com/bitrise-io/bitrise-addon-test/addonprovisioner"
	"github.com/bitrise-io/bitrise-addon-test/addontester"
	"github.com/bitrise-io/bitrise-addon-test/loginserver"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	serverCmdPort string
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a login server",
	Long: `Start a login server.
Open the URL of the login server in your browser
to simulate a Bitrise Add-on Login for your Add-on.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.WithStack(loginServerFn())
	},
}

func init() {
	loginCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	serverCmd.PersistentFlags().StringVar(&serverCmdPort, "port", "3000", "Port used by the login server")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func loginServerFn() error {
	tester, err := addonTesterFromConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	ls := loginserver.NewLoginServer(serverCmdPort, func() (addonprovisioner.LoginRequestInfos, error) {
		loginRequest, err := tester.LoginRequestInfos(addontester.LoginTesterParams{
			AppSlug:   loginAppSlug,
			AppTitle:  loginAppTitle,
			BuildSlug: loginBuildSlug,
			Timestamp: loginTimestamp,
		})
		return loginRequest, errors.WithStack(err)
	})
	return errors.WithStack(ls.Start())
}
