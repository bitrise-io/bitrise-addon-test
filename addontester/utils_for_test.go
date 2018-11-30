package addontester_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/bitrise-team/bitrise-addon-test/addontester"
	"github.com/stretchr/testify/require"
)

type TesterTestCase struct {
	responseStatusCode int
	responseBody       string
	requestError       string
	expectedError      string
	testCaseID         string
	testWithRetry      bool
	testerMethodToCall func(tester *addontester.Tester) error

	//
	testTimestamp int64
}

func performTesterTest(t *testing.T, tc TesterTestCase) {
	t.Run(tc.testCaseID, func(t *testing.T) {
		var buf bytes.Buffer

		requestError := (error)(nil)
		if tc.requestError != "" {
			requestError = errors.New(tc.requestError)
		}
		tester, _ := addontester.New(
			&testAddonClient{
				addonURL:           "http://example.com",
				authToken:          "auth-token",
				ssoSecret:          "sso-secret",
				responseStatusCode: tc.responseStatusCode,
				responseBody:       tc.responseBody,
				err:                requestError,
			},
			log.New(&buf, "", 0),
		)

		if tc.testerMethodToCall == nil {
			panic("Specify a func to test")
		}

		err := tc.testerMethodToCall(tester)

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
