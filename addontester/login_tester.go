package addontester

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/bitrise-io/bitrise-addon-test/addonprovisioner"
	"github.com/bitrise-io/bitrise-addon-test/utils"
	"github.com/pkg/errors"
)

// LoginTesterParams ...
type LoginTesterParams struct {
	AppSlug   string
	AppTitle  string
	BuildSlug string
	Timestamp int64
}

func prepareNormalizeLoginTesterParams(params *LoginTesterParams) error {
	if len(params.AppSlug) == 0 {
		var err error
		params.AppSlug, err = utils.RandomHex(8)
		if err != nil {
			return fmt.Errorf("Failed to generate app slug: %+v", err)
		}
	}

	if len(params.AppTitle) == 0 {
		params.AppTitle = fmt.Sprintf("Title of app #%s", params.AppSlug)
	}

	if len(params.BuildSlug) == 0 {
		var err error
		params.BuildSlug, err = utils.RandomHex(8)
		if err != nil {
			return fmt.Errorf("Failed to generate build slug: %+v", err)
		}
	}

	if params.Timestamp == 0 {
		params.Timestamp = time.Now().Unix()
	}

	return nil
}

// Login ...
func (t *Tester) Login(params LoginTesterParams, remainingRetries int) error {
	if err := prepareNormalizeLoginTesterParams(&params); err != nil {
		return errors.WithStack(err)
	}

	t.logger.Printf("\nLogin details:")
	t.logger.Printf("App slug: %s", params.AppSlug)
	t.logger.Printf("App title: %s", params.AppTitle)
	t.logger.Printf("Build slug: %s", params.BuildSlug)
	t.logger.Printf("Timestamp: %d", params.Timestamp)

	status, body, err := t.addonClient.Login(addonprovisioner.LoginRequestParams{
		AppSlug:   params.AppSlug,
		AppTitle:  params.AppTitle,
		BuildSlug: params.BuildSlug,
		Timestamp: fmt.Sprintf("%d", params.Timestamp),
	})

	if err != nil {
		return fmt.Errorf("Login failed: %+v", err)
	}

	t.logger.Printf("\nResponse status: %d", status)
	t.logger.Printf("Response body: %v\n", body)

	if status < 200 || status > 299 {
		return fmt.Errorf("Login request resulted in a non-2xx response (%d)", status)
	}

	t.logger.Println("\nLogin success.")

	r := strings.NewReader(body)
	d := xml.NewDecoder(r)
	d.Strict = true
	d.Entity = xml.HTMLEntity
	var nodes []html.Node
	err = d.Decode(&nodes)

	if err != nil {
		return fmt.Errorf("Login request responded with invalid HTML: %+v", err)
	}

	return nil
}

// LoginRequestInfos ...
func (t *Tester) LoginRequestInfos(params LoginTesterParams) (addonprovisioner.LoginRequestInfos, error) {
	if err := prepareNormalizeLoginTesterParams(&params); err != nil {
		return addonprovisioner.LoginRequestInfos{}, errors.WithStack(err)
	}

	t.logger.Printf("\nLogin details:")
	t.logger.Printf("App slug: %s", params.AppSlug)
	t.logger.Printf("App title: %s", params.AppTitle)
	t.logger.Printf("Build slug: %s", params.BuildSlug)
	t.logger.Printf("Timestamp: %d", params.Timestamp)

	lri, err := t.addonClient.LoginRequestInfos(addonprovisioner.LoginRequestParams{
		AppSlug:   params.AppSlug,
		AppTitle:  params.AppTitle,
		BuildSlug: params.BuildSlug,
		Timestamp: fmt.Sprintf("%d", params.Timestamp),
	})
	return lri, errors.WithStack(err)
}
