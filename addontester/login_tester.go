package addontester

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/bitrise-team/bitrise-addon-test/addonprovisioner"
	"github.com/bitrise-team/bitrise-addon-test/utils"
)

// LoginTesterParams ...
type LoginTesterParams struct {
	AppSlug   string
	BuildSlug string
	Timestamp int64
}

// Login ...
func (t *Tester) Login(params LoginTesterParams, remainingRetries int) error {
	if len(params.AppSlug) == 0 {
		var err error
		params.AppSlug, err = utils.RandomHex(8)
		if err != nil {
			return fmt.Errorf("Failed to generate app slug: %s", err)
		}
	}

	if len(params.BuildSlug) == 0 {
		var err error
		params.BuildSlug, err = utils.RandomHex(8)
		if err != nil {
			return fmt.Errorf("Failed to generate build slug: %s", err)
		}
	}

	if params.Timestamp == 0 {
		params.Timestamp = time.Now().Unix()
	}

	t.logger.Printf("\nLogin details:")
	t.logger.Printf("App slug: %s", params.AppSlug)
	t.logger.Printf("Build slug: %s", params.BuildSlug)
	t.logger.Printf("Timestamp: %d", params.Timestamp)

	status, body, err := t.addonClient.Login(addonprovisioner.LoginRequestParams{
		AppSlug:   params.AppSlug,
		BuildSlug: params.BuildSlug,
		Timestamp: fmt.Sprintf("%d", params.Timestamp),
	})

	if err != nil {
		return fmt.Errorf("Login failed: %s", err)
	}

	t.logger.Printf("\nResponse status: %d", status)
	t.logger.Printf("Response body: %v\n", body)

	if status < 200 || status > 299 {
		return fmt.Errorf("Login request resulted in a non-2xx response")
	}

	t.logger.Println("\nLogin success.")

	r := strings.NewReader(body)
	d := xml.NewDecoder(r)
	d.Strict = true
	d.Entity = xml.HTMLEntity
	var nodes []html.Node
	err = d.Decode(&nodes)

	if err != nil {
		return fmt.Errorf("Login request responded with invalid HTML")
	}

	return nil
}
