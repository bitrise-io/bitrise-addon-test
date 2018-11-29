package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/bitrise-team/bitrise-add-on-testing-kit/addonprovisioner"
	"github.com/bitrise-team/bitrise-add-on-testing-kit/addontester"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	colorRed        = "\x1b[31;1m"
	colorNeutral    = "\x1b[0m"
	numberOfRetries = 2
)

var (
	cfgFile string
	//
	rootAppSlug      string
	rootBuildSlug    string
	rootAPIToken     string
	rootInitialPlan  string
	rootPlanChangeTo string
	rootTimestamp    int64
)

func fail(err error) {
	fmt.Printf("\n%s%s%s\n", colorRed, err, colorNeutral)
	os.Exit(1)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bitrise-addon-test",
	Short: "Testing application for Bitrise Add-on Developers",
	Long: `Application for testing add-ons written for Bitrise.

Running this root command (bitrise-addon-test) will make a comprehensive testing, which consists of testing provisioning request (with 2 retries), change plan request, login request and deprovisioning request (with 2 retries). You can run these tests separately, please find the available commands below.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := comprehensive()
		if err != nil {
			fail(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file to use (default is ./config.yaml)")
	rootCmd.PersistentFlags().StringVar(&rootAppSlug, "app-slug", "", "The slug of the app the add-on gets provisioned to. It gets randomly generated if not given.")
	rootCmd.PersistentFlags().StringVar(&rootBuildSlug, "build-slug", "", "The slug of the build")
	rootCmd.PersistentFlags().StringVar(&rootAPIToken, "api-token", "", "An API token of the app the add-on gets provisioned to. The add-on can behave on behalf of the app using the Bitrise API. It gets randomly generated if not given.")
	rootCmd.PersistentFlags().StringVar(&rootInitialPlan, "plan", "free", "The plan of the provisioned add-on.")
	rootCmd.PersistentFlags().StringVar(&rootPlanChangeTo, "plan-change-to", "pro", "The plan the add-on gets changed to.")
	rootCmd.PersistentFlags().Int64Var(&rootTimestamp, "timestamp", 0, "The slug of the build")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name "config" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Reading config file:", viper.ConfigFileUsed())
	} else {
		fmt.Printf("Failed to read config file: %s", err.Error())
		os.Exit(1)
	}

	validateConfig()
}

func validateConfig() {
	requiredConfigs := []string{"addon-url", "auth-token", "sso-secret"}

	fmt.Println("\nConfigs:")
	for _, config := range requiredConfigs {
		if !viper.IsSet(config) {
			fmt.Printf("Config %s is required but not set\n", config)
			os.Exit(1)
		}
		fmt.Printf("%s: %s\n", config, viper.Get(config))
	}
}

func addonTesterFromConfig() (*addontester.Tester, error) {
	addonClient, err := addonprovisioner.NewClient(&addonprovisioner.ClientConfig{
		AddonURL:  viper.Get("addon-url").(string),
		AuthToken: viper.Get("auth-token").(string),
		SSOSecret: viper.Get("sso-secret").(string),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return addontester.New(addonClient, log.New(os.Stdout, "", 0))
}

func comprehensive() error {
	tester, err := addonTesterFromConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	return tester.Comprehensive(addontester.ComprehensiveTesterParams{
		AppSlug:      rootAppSlug,
		BuildSlug:    rootBuildSlug,
		APIToken:     rootAPIToken,
		InitialPlan:  rootInitialPlan,
		PlanChangeTo: rootPlanChangeTo,
		Timestamp:    rootTimestamp,
	})
}
