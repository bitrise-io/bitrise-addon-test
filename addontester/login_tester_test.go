package addontester_test

import (
	"net/http"
	"testing"

	"github.com/bitrise-io/bitrise-addon-test/addontester"
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
			responseStatusCode: http.StatusInternalServerError,
			responseBody:       `{"message":"Internal Server Error"}`,
			expectedError:      "Login request resulted in a non-2xx response",
		},
		{
			responseStatusCode: http.StatusOK,
			responseBody:       `definately not an HTML`,
			expectedError:      "Login request responded with invalid HTML",
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
			}, numberOfRetryTests)
		}
		performTesterTest(t, tc)
	}
}
