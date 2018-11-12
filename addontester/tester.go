package addontester

import (
	"log"

	"github.com/bitrise-team/bitrise-add-on-testing-kit/addonprovisioner"
)

const numberOfTestsWithRetry = 3

// Config ...
type Config struct {
	Logger    *log.Logger
	AddonURL  string
	AuthToken string
	SSOSecret string
}

// Tester ...
type Tester struct {
	logger      *log.Logger
	addonClient addonprovisioner.AddonClientInterface
	addonURL    string
	authToken   string
	ssoSecret   string
}

// New ...
func New(addonProvClient addonprovisioner.AddonClientInterface, logger *log.Logger) (*Tester, error) {
	client := Tester{
		logger:      logger,
		addonURL:    addonProvClient.AddonURL(),
		authToken:   addonProvClient.AuthToken(),
		ssoSecret:   addonProvClient.SSOSecret(),
		addonClient: addonProvClient,
	}

	return &client, nil
}
