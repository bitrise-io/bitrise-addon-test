package addontester_test

import (
	"bytes"
	"log"
	"net/http"
	"testing"

	"github.com/bitrise-team/bitrise-add-on-testing-kit/addontester"
	"github.com/stretchr/testify/require"
)

func Test_Provision(t *testing.T) {
	var buf bytes.Buffer

	t.Log("when client response has non-success response")
	{
		tester, _ := addontester.New(
			&testAddonClient{
				addonURL:           "http://example.com",
				authToken:          "auth-token",
				ssoSecret:          "sso-secret",
				responseStatusCode: http.StatusInternalServerError,
			},
			log.New(&buf, "", 0),
		)

		err := tester.Provision(addontester.ProvisionParams{
			AppSlug:  "app-slug",
			APIToken: "api-token",
			Plan:     "plan",
		})

		require.Equal(t, "Provisioning request resulted in a non-2xx response", err.Error())
	}
}
