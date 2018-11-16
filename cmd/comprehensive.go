package cmd

import (
	"github.com/bitrise-team/bitrise-add-on-testing-kit/addontester"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	comprehensiveAppSlug      string
	comprehensiveBuildSlug    string
	comprehensiveAPIToken     string
	comprehensiveInitialPlan  string
	comprehensivePlanChangeTo string
	comprehensiveTimestamp    int64
)

var comprehensiveCmd = &cobra.Command{
	Use:   "comprehensive",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := comprehensive()
		if err != nil {
			fail(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(comprehensiveCmd)
	comprehensiveCmd.PersistentFlags().StringVar(&comprehensiveAppSlug, "app-slug", "", "The slug of the app the add-on gets provisioned to. It gets randomly generated if not given.")
	comprehensiveCmd.PersistentFlags().StringVar(&comprehensiveBuildSlug, "build-slug", "", "The slug of the build")
	comprehensiveCmd.PersistentFlags().StringVar(&comprehensiveAPIToken, "api-token", "", "An API token of the app the add-on gets provisioned to. The add-on can behave on behalf of the app using the Bitrise API. It gets randomly generated if not given.")
	comprehensiveCmd.PersistentFlags().StringVar(&comprehensiveInitialPlan, "plan", "free", "The plan of the provisioned add-on.")
	comprehensiveCmd.PersistentFlags().StringVar(&comprehensivePlanChangeTo, "plan-change-to", "pro", "The plan the add-on gets changed to.")
	comprehensiveCmd.PersistentFlags().Int64Var(&comprehensiveTimestamp, "timestamp", 0, "The slug of the build")
}

func comprehensive() error {
	tester, err := addonTesterFromConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	return tester.Comprehensive(addontester.ComprehensiveTesterParams{
		AppSlug:      comprehensiveAppSlug,
		BuildSlug:    comprehensiveBuildSlug,
		APIToken:     comprehensiveAPIToken,
		InitialPlan:  comprehensiveInitialPlan,
		PlanChangeTo: comprehensivePlanChangeTo,
		Timestamp:    comprehensiveTimestamp,
	})
}
