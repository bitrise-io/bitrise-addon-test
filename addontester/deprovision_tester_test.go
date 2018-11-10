package addontester_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/bitrise-team/bitrise-add-on-testing-kit/addontester"
	"github.com/stretchr/testify/require"
)

func Test_Deprovision(t *testing.T) {
	for _, tc := range []struct {
		responseStatusCode      int
		responseBody            string
		deprovisionRequestError string
		expectedError           string
		testCaseID              string
	}{
		{
			responseStatusCode: http.StatusOK,
			testCaseID:         "ok",
		},
		{
			responseStatusCode: http.StatusInternalServerError,
			responseBody:       `{"message":"Internal Server Error"}`,
			expectedError:      "Deprovisioning request resulted in a non-2xx response",
		},
		{
			deprovisionRequestError: "Some HTTP request issue",
			expectedError:           "Deprovisioning failed: Some HTTP request issue",
		},
	} {
		t.Run(tc.testCaseID, func(t *testing.T) {
			var buf bytes.Buffer

			deprovisionRequestError := (error)(nil)
			if tc.deprovisionRequestError != "" {
				deprovisionRequestError = errors.New(tc.deprovisionRequestError)
			}
			tester, _ := addontester.New(
				&testAddonClient{
					addonURL:           "http://example.com",
					authToken:          "auth-token",
					ssoSecret:          "sso-secret",
					responseStatusCode: tc.responseStatusCode,
					responseBody:       tc.responseBody,
					err:                deprovisionRequestError,
				},
				log.New(&buf, "", 0),
			)

			err := tester.Deprovision(addontester.DeprovisionParams{AppSlug: "app-slug"})

			if tc.expectedError == "" {
				require.NoError(t, err)
			} else {
				require.Equal(t, tc.expectedError, err.Error())
			}
			if tc.testCaseID != "" {
				expectedTestData, err := ioutil.ReadFile(filepath.Join("../testdata", filepath.FromSlash(t.Name()+".golden")))
				if err != nil {
					t.Fatalf("failed reading .golden: %s", err)
				}
				require.Equal(t, string(expectedTestData), buf.String())
			}
			t.FailNow()
		})
	}
}
