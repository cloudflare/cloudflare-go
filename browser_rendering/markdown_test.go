// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/browser_rendering"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

func TestMarkdownNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.BrowserRendering.Markdown.New(context.TODO(), browser_rendering.MarkdownNewParams{
		AccountID:     cloudflare.F("account_id"),
		CacheTTL:      cloudflare.F(86400.000000),
		ActionTimeout: cloudflare.F(300000.000000),
		AddScriptTag: cloudflare.F([]browser_rendering.MarkdownNewParamsAddScriptTag{{
			ID:      cloudflare.F("id"),
			Content: cloudflare.F("content"),
			Type:    cloudflare.F("type"),
			URL:     cloudflare.F("url"),
		}}),
		AddStyleTag: cloudflare.F([]browser_rendering.MarkdownNewParamsAddStyleTag{{
			Content: cloudflare.F("content"),
			URL:     cloudflare.F("url"),
		}}),
		AllowRequestPattern: cloudflare.F([]string{"string"}),
		AllowResourceTypes:  cloudflare.F([]browser_rendering.MarkdownNewParamsAllowResourceType{browser_rendering.MarkdownNewParamsAllowResourceTypeDocument}),
		Authenticate: cloudflare.F(browser_rendering.MarkdownNewParamsAuthenticate{
			Password: cloudflare.F("x"),
			Username: cloudflare.F("x"),
		}),
		BestAttempt: cloudflare.F(true),
		Cookies: cloudflare.F([]browser_rendering.MarkdownNewParamsCookie{{
			Name:         cloudflare.F("name"),
			Value:        cloudflare.F("value"),
			Domain:       cloudflare.F("domain"),
			Expires:      cloudflare.F(0.000000),
			HTTPOnly:     cloudflare.F(true),
			PartitionKey: cloudflare.F("partitionKey"),
			Path:         cloudflare.F("path"),
			Priority:     cloudflare.F(browser_rendering.MarkdownNewParamsCookiesPriorityLow),
			SameParty:    cloudflare.F(true),
			SameSite:     cloudflare.F(browser_rendering.MarkdownNewParamsCookiesSameSiteStrict),
			Secure:       cloudflare.F(true),
			SourcePort:   cloudflare.F(0.000000),
			SourceScheme: cloudflare.F(browser_rendering.MarkdownNewParamsCookiesSourceSchemeUnset),
			URL:          cloudflare.F("url"),
		}}),
		EmulateMediaType: cloudflare.F("emulateMediaType"),
		GotoOptions: cloudflare.F(browser_rendering.MarkdownNewParamsGotoOptions{
			Referer:        cloudflare.F("referer"),
			ReferrerPolicy: cloudflare.F("referrerPolicy"),
			Timeout:        cloudflare.F(60000.000000),
			WaitUntil:      cloudflare.F[browser_rendering.MarkdownNewParamsGotoOptionsWaitUntilUnion](browser_rendering.MarkdownNewParamsGotoOptionsWaitUntilString(browser_rendering.MarkdownNewParamsGotoOptionsWaitUntilStringLoad)),
		}),
		HTML:                 cloudflare.F("x"),
		RejectRequestPattern: cloudflare.F([]string{"string"}),
		RejectResourceTypes:  cloudflare.F([]browser_rendering.MarkdownNewParamsRejectResourceType{browser_rendering.MarkdownNewParamsRejectResourceTypeDocument}),
		SetExtraHTTPHeaders: cloudflare.F(map[string]string{
			"foo": "string",
		}),
		SetJavaScriptEnabled: cloudflare.F(true),
		URL:                  cloudflare.F("https://example.com"),
		UserAgent:            cloudflare.F("userAgent"),
		Viewport: cloudflare.F(browser_rendering.MarkdownNewParamsViewport{
			Height:            cloudflare.F(0.000000),
			Width:             cloudflare.F(0.000000),
			DeviceScaleFactor: cloudflare.F(0.000000),
			HasTouch:          cloudflare.F(true),
			IsLandscape:       cloudflare.F(true),
			IsMobile:          cloudflare.F(true),
		}),
		WaitForSelector: cloudflare.F(browser_rendering.MarkdownNewParamsWaitForSelector{
			Selector: cloudflare.F("selector"),
			Hidden:   cloudflare.F(browser_rendering.MarkdownNewParamsWaitForSelectorHiddenTrue),
			Timeout:  cloudflare.F(60000.000000),
			Visible:  cloudflare.F(browser_rendering.MarkdownNewParamsWaitForSelectorVisibleTrue),
		}),
		WaitForTimeout: cloudflare.F(60000.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
