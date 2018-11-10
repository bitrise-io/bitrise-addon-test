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

func Test_Provision(t *testing.T) {
	for _, tc := range []struct {
		responseStatusCode    int
		responseBody          string
		provisionRequestError string
		expectedError         string
		testCaseID            string
	}{
		{
			responseStatusCode: http.StatusOK,
			responseBody:       `{"envs":[{"key":"key1","value":"value1"}]}`,
			testCaseID:         "ok",
		},
		{
			responseStatusCode: http.StatusOK,
			responseBody:       `{"envs":[]}`,
			testCaseID:         "ok_no_envs",
		},
		{
			responseStatusCode: http.StatusOK,
			responseBody:       `{"envs":[{"value":"value1"}]}`,
			expectedError:      "ENV var key is not present: { value1}",
		},
		{
			responseStatusCode: http.StatusOK,
			responseBody:       `{"envs":[{"key":"key1"}]}`,
			expectedError:      "ENV var value is not present: {key1 }",
		},
		{
			responseStatusCode: http.StatusInternalServerError,
			responseBody:       `{"message":"Internal Server Error"}`,
			expectedError:      "Provisioning request resulted in a non-2xx response",
		},
		{
			provisionRequestError: "Some HTTP request issue",
			expectedError:         "Provisioning failed: Some HTTP request issue",
		},
	} {
		t.Run(tc.testCaseID, func(t *testing.T) {
			var buf bytes.Buffer

			provisionRequestError := (error)(nil)
			if tc.provisionRequestError != "" {
				provisionRequestError = errors.New(tc.provisionRequestError)
			}
			tester, _ := addontester.New(
				&testAddonClient{
					addonURL:           "http://example.com",
					authToken:          "auth-token",
					ssoSecret:          "sso-secret",
					responseStatusCode: tc.responseStatusCode,
					responseBody:       tc.responseBody,
					err:                provisionRequestError,
				},
				log.New(&buf, "", 0),
			)

			err := tester.Provision(addontester.ProvisionParams{
				AppSlug:  "app-slug",
				APIToken: "api-token",
				Plan:     "plan",
			})

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
		})
	}

	// t.Log("OK")
	// {
	// 	tester, _ := addontester.New(
	// 		&testAddonClient{
	// 			addonURL:           "http://example.com",
	// 			authToken:          "auth-token",
	// 			ssoSecret:          "sso-secret",
	// 			responseStatusCode: http.StatusOK,
	// 			responseBody:       `{"envs":[{"key":"key1","value":"value1"}]}`,
	// 		},
	// 		log.New(&buf, "", 0),
	// 	)

	// 	err := tester.Provision(addontester.ProvisionParams{
	// 		AppSlug:  "app-slug",
	// 		APIToken: "api-token",
	// 		Plan:     "plan",
	// 	})
	// 	require.NoError(t, err)
	// 	require.Equal(t, testDatas["case1"], buf.String())
	// }

	// t.Log("OK")
	// {
	// 	tester, _ := addontester.New(
	// 		&testAddonClient{
	// 			addonURL:           "http://example.com",
	// 			authToken:          "auth-token",
	// 			ssoSecret:          "sso-secret",
	// 			responseStatusCode: http.StatusOK,
	// 			responseBody:       `{"envs":[]}`,
	// 		},
	// 		log.New(&buf, "", 0),
	// 	)

	// 	err := tester.Provision(addontester.ProvisionParams{
	// 		AppSlug:  "app-slug",
	// 		APIToken: "api-token",
	// 		Plan:     "plan",
	// 	})
	// 	require.NoError(t, err)
	// 	// require.Equal(t, testDatas["case1"], buf.String())
	// }

	// t.Log("when client response has non-success response")
	// {
	// 	tester, _ := addontester.New(
	// 		&testAddonClient{
	// 			addonURL:           "http://example.com",
	// 			authToken:          "auth-token",
	// 			ssoSecret:          "sso-secret",
	// 			responseStatusCode: http.StatusInternalServerError,
	// 		},
	// 		log.New(&buf, "", 0),
	// 	)

	// 	err := tester.Provision(addontester.ProvisionParams{
	// 		AppSlug:  "app-slug",
	// 		APIToken: "api-token",
	// 		Plan:     "plan",
	// 	})

	// 	require.Equal(t, "Provisioning request resulted in a non-2xx response", err.Error())
	// }
}
