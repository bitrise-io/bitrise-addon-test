package addontester_test

import (
	"net/http"
	"testing"

	"github.com/bitrise-team/bitrise-add-on-testing-kit/addontester"
)

func Test_Deprovision(t *testing.T) {
	for _, tc := range []TesterTestCase{
		{
			responseStatusCode: http.StatusOK,
			testCaseID:         "ok",
		},
		{
			responseStatusCode: http.StatusOK,
			testCaseID:         "ok_with_retry",
			testWithRetry:      true,
		},
		{
			responseStatusCode: http.StatusInternalServerError,
			responseBody:       `{"message":"Internal Server Error"}`,
			expectedError:      "Deprovisioning request resulted in a non-2xx response",
		},
		{
			requestError:  "Some HTTP request issue",
			expectedError: "Deprovisioning failed: Some HTTP request issue",
		},
	} {
		tc.testerMethodToCall = func(tester *addontester.Tester) error {
			return tester.Deprovision(addontester.DeprovisionTesterParams{
				AppSlug:   "app-slug",
				WithRetry: tc.testWithRetry,
			}, numberOfRetryTests)
		}
		performTesterTest(t, tc)
	}
}
