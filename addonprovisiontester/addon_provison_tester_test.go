package addonprovisiontester_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/bitrise-team/bitrise-add-on-testing-kit/addonprovisiontester"
	"github.com/stretchr/testify/require"
)

func Test_Provision(t *testing.T) {

	var buf bytes.Buffer
	tester, _ := addonprovisiontester.New(
		&addonprovisiontester.Config{
			AddonURL:  "http://example.com",
			AuthToken: "auth-token",
			SSOSecret: "sso-secret",
			Logger:    log.New(&buf, "", 0),
		})

	err := tester.Provision(addonprovisiontester.ProvisionParams{
		AppSlug:  "app-slug",
		APIToken: "api-token",
		Plan:     "plan",
	})

	require.Equal(t, "nemez", buf.String())
	require.Equal(t, "Provisioning request resulted in a non-2xx response", err.Error())
}
