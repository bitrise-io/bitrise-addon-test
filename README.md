# Bitrise Add-on Test

Testing kit for Bitrise Add-on developers, you can use this CLI tool to test your add-on for Bitrise. For detailed information please use the help flag for the commands.

```
Application for testing add-ons written for Bitrise.

Running this root command (bitrise-addon-test) will make a comprehensive testing, which consists of testing provisioning request (with 2 retries), change plan request, login request and deprovisioning request (with 2 retries). You can run these tests separately, please find the available commands below.

Usage:
  bitrise-addon-test [flags]
  bitrise-addon-test [command]

Available Commands:
  change-plan Test for plan change request
  deprovision Test for deprovision request
  help        Help about any command
  login       Test for SSO login request
  provision   Test for deprovision request

Flags:
      --api-token string        An API token of the app the add-on gets provisioned to. The add-on can behave on behalf of the app using the Bitrise API. It gets randomly generated if not given.
      --app-slug string         The slug of the app the add-on gets provisioned to. It gets randomly generated if not given.
      --build-slug string       The slug of the build
      --config string           config file to use (default is ./config.yaml)
  -h, --help                    help for bitrise-addon-test
      --plan string             The plan of the provisioned add-on. (default "free")
      --plan-change-to string   The plan the add-on gets changed to. (default "pro")
      --timestamp int           Timestamp for SSO login token generation

Use "bitrise-addon-test [command] --help" for more information about a command.
```