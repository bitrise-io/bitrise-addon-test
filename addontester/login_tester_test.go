package addontester_test

import (
	"net/http"
	"testing"

	"github.com/bitrise-team/bitrise-add-on-testing-kit/addontester"
)

func Test_Login(t *testing.T) {
	for _, tc := range []TesterTestCase{
		{
			responseStatusCode: http.StatusOK,
			responseBody:       "<html><body><p>Hello Bitrise!</p></body></html>",
			testCaseID:         "ok",
			testTimestamp:      1542292740,
		},
		{
			responseStatusCode: http.StatusOK,
			responseBody:       "<html><body><p>Hello Bitrise!</p></body></html>",
			testCaseID:         "ok_with_retry",
			testWithRetry:      true,
			testTimestamp:      1542292740,
		},
		{
			responseStatusCode: http.StatusInternalServerError,
			responseBody:       `{"message":"Internal Server Error"}`,
			expectedError:      "Login request resulted in a non-2xx response",
		},
		{
			requestError:  "Some HTTP request issue",
			expectedError: "Login failed: Some HTTP request issue",
		},
	} {
		tc.testerMethodToCall = func(tester *addontester.Tester) error {
			return tester.Login(addontester.LoginTesterParams{
				AppSlug:   "app-slug",
				BuildSlug: "build-slug",
				Timestamp: tc.testTimestamp,
				WithRetry: tc.testWithRetry,
			}, numberOfRetryTests)
		}
		performTesterTest(t, tc)
	}
}
