package addontester_test

import (
	"net/http"
	"testing"

	"github.com/bitrise-io/bitrise-addon-test/addontester"
)

func Test_Provision(t *testing.T) {
	for _, tc := range []TesterTestCase{
		{
			responseStatusCode: http.StatusOK,
			responseBody:       `{"envs":[{"key":"key1","value":"value1"},{"key":"key2","value":"value2"}]}`,
			testCaseID:         "ok",
		},
		{
			responseStatusCode: http.StatusOK,
			responseBody:       `{"envs":[{"key":"key1","value":"value1"},{"key":"key2","value":"value2"}]}`,
			testCaseID:         "ok_with_retry",
			testWithRetry:      true,
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
			requestError:  "Some HTTP request issue",
			expectedError: "Provisioning failed: Some HTTP request issue",
		},
	} {
		tc.testerMethodToCall = func(tester *addontester.Tester) error {
			return tester.Provision(addontester.ProvisionTesterParams{
				AppSlug:   "app-slug",
				AppTitle:  "Title of the app",
				APIToken:  "api-token",
				Plan:      "plan",
				WithRetry: tc.testWithRetry,
			}, numberOfRetryTests)
		}
		performTesterTest(t, tc)
	}
}
