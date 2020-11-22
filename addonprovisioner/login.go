package addonprovisioner

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/bitrise-io/bitrise-addon-test/utils"
	"github.com/bitrise-io/go-utils/log"
	"github.com/moul/http2curl"
	"github.com/pkg/errors"
)

// LoginRequestParams ...
type LoginRequestParams struct {
	AppSlug   string `json:"app_slug"`
	AppTitle  string `json:"app_title"`
	BuildSlug string `json:"build_slug"`
	Timestamp string `json:"timestamp"`
}

// NewLoginRequestParams ...
func NewLoginRequestParams(appSlug, appTitle, buildSlug, timestamp string) LoginRequestParams {
	return LoginRequestParams{
		AppSlug:   appSlug,
		AppTitle:  appTitle,
		BuildSlug: buildSlug,
		Timestamp: timestamp,
	}
}

// LoginFormData ..
type LoginFormData struct {
	AppSlug   string
	Token     string
	Timestamp string
}

// LoginFormData ...
func (loginRequestParams LoginRequestParams) LoginFormData(ssoSecret string) (LoginFormData, error) {
	s := sha1.New()
	_, err := s.Write([]byte(fmt.Sprintf("%s:%s:%s", loginRequestParams.AppSlug, ssoSecret, loginRequestParams.Timestamp)))
	if err != nil {
		log.Errorf("Failed to write into sha1 buffer, error: %s", err)
		return LoginFormData{}, errors.Wrap(err, "Failed to submit form")
	}
	refToken := fmt.Sprintf("%x", s.Sum(nil))

	return LoginFormData{
		AppSlug:   loginRequestParams.AppSlug,
		Token:     refToken,
		Timestamp: loginRequestParams.Timestamp,
	}, nil
}

// LoginRequestHeaders ...
type LoginRequestHeaders struct {
	Authentication string
	ContentType    string
}

// LoginRequestInfos ...
type LoginRequestInfos struct {
	Method   string
	URL      string
	FormData LoginFormData
	Headers  LoginRequestHeaders
}

// LoginRequestInfos ...
func (c *Client) LoginRequestInfos(params LoginRequestParams) (LoginRequestInfos, error) {
	loginFormData, err := params.LoginFormData(c.SSOSecret())
	if err != nil {
		return LoginRequestInfos{}, errors.Wrap(err, "Failed to generate form")
	}

	v := url.Values{}
	v.Set("build_slug", params.BuildSlug)
	v.Set("app_title", params.AppTitle)

	return LoginRequestInfos{
		Method:   "POST",
		URL:      c.addonURL + "/login?" + v.Encode(),
		FormData: loginFormData,
		Headers: LoginRequestHeaders{
			Authentication: c.authToken,
			ContentType:    "application/x-www-form-urlencoded",
		},
	}, nil
}

// Login ...
func (c *Client) Login(params LoginRequestParams) (int, string, error) {
	loginRequestInfos, err := c.LoginRequestInfos(params)

	formData := strings.NewReader(url.Values{
		"app_slug":  {loginRequestInfos.FormData.AppSlug},
		"token":     {loginRequestInfos.FormData.Token},
		"timestamp": {loginRequestInfos.FormData.Timestamp},
	}.Encode())

	req, err := http.NewRequest(
		loginRequestInfos.Method,
		loginRequestInfos.URL,
		formData,
	)
	if err != nil {
		return 0, "", errors.Wrap(err, "Failed to submit form")
	}
	req.Header.Add("Authentication", loginRequestInfos.Headers.Authentication)
	req.Header.Set("Content-Type", loginRequestInfos.Headers.ContentType)

	command, err := http2curl.GetCurlCommand(req)
	if err != nil {
		return 0, "", err
	}

	fmt.Println("\nMaking request:")
	fmt.Println(command)
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, "", errors.Wrap(err, "Failed to send request")
	}

	defer utils.ResponseBodyCloseWithErrorLog(resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to read response body: %s", err)
		os.Exit(1)
	}
	bodyString := string(bodyBytes)

	return resp.StatusCode, bodyString, nil
}
