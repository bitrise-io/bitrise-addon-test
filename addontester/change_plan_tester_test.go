package addontester_test

import (
	"net/http"
	"testing"

	"github.com/bitrise-team/bitrise-add-on-testing-kit/addontester"
)

func Test_ChangePlan(t *testing.T) {
	for _, tc := range []TesterTestCase{
		{
			responseStatusCode: http.StatusOK,
			testCaseID:         "ok",
		},
		{
			responseStatusCode: http.StatusInternalServerError,
			responseBody:       `{"message":"Internal Server Error"}`,
			expectedError:      "Plan changing request resulted in a non-2xx response",
		},
		{
			requestError:  "Some HTTP request issue",
			expectedError: "Plan changing failed: Some HTTP request issue",
		},
	} {
		tc.testerMethodToCall = func(tester *addontester.Tester) error {
			return tester.ChangePlan(addontester.ChangePlanTesterParams{
				AppSlug: "app-slug",
				Plan:    "plan",
			})
		}
		performTesterTest(t, tc)
	}
}
