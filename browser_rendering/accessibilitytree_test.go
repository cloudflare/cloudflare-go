// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/browser_rendering"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestAccessibilityTreeNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.BrowserRendering.AccessibilityTree.New(context.TODO(), browser_rendering.AccessibilityTreeNewParams{
		AccountID: cloudflare.F("account_id"),
		Body: browser_rendering.AccessibilityTreeNewParamsBodyObject{
			URL:           cloudflare.F("https://www.example.com/"),
			ActionTimeout: cloudflare.F(120000.000000),
			AddScriptTag: cloudflare.F([]browser_rendering.AccessibilityTreeNewParamsBodyObjectAddScriptTag{{
				ID:      cloudflare.F("id"),
				Content: cloudflare.F("content"),
				Type:    cloudflare.F("type"),
				URL:     cloudflare.F("https://example.com"),
			}}),
			AddStyleTag: cloudflare.F([]browser_rendering.AccessibilityTreeNewParamsBodyObjectAddStyleTag{{
				Content: cloudflare.F("content"),
				URL:     cloudflare.F("https://example.com"),
			}}),
			AllowRequestPattern: cloudflare.F([]string{"string"}),
			AllowResourceTypes:  cloudflare.F([]browser_rendering.AccessibilityTreeNewParamsBodyObjectAllowResourceType{browser_rendering.AccessibilityTreeNewParamsBodyObjectAllowResourceTypeDocument}),
			Authenticate: cloudflare.F(browser_rendering.AccessibilityTreeNewParamsBodyObjectAuthenticate{
				Password: cloudflare.F("x"),
				Username: cloudflare.F("x"),
			}),
			BestAttempt: cloudflare.F(true),
			Cookies: cloudflare.F([]browser_rendering.AccessibilityTreeNewParamsBodyObjectCookie{{
				Name:         cloudflare.F("name"),
				Value:        cloudflare.F("value"),
				Domain:       cloudflare.F("domain"),
				Expires:      cloudflare.F(0.000000),
				HTTPOnly:     cloudflare.F(true),
				PartitionKey: cloudflare.F("partitionKey"),
				Path:         cloudflare.F("path"),
				Priority:     cloudflare.F(browser_rendering.AccessibilityTreeNewParamsBodyObjectCookiesPriorityLow),
				SameParty:    cloudflare.F(true),
				SameSite:     cloudflare.F(browser_rendering.AccessibilityTreeNewParamsBodyObjectCookiesSameSiteStrict),
				Secure:       cloudflare.F(true),
				SourcePort:   cloudflare.F(0.000000),
				SourceScheme: cloudflare.F(browser_rendering.AccessibilityTreeNewParamsBodyObjectCookiesSourceSchemeUnset),
				URL:          cloudflare.F("url"),
			}}),
			EmulateMediaType: cloudflare.F("emulateMediaType"),
			GotoOptions: cloudflare.F(browser_rendering.AccessibilityTreeNewParamsBodyObjectGotoOptions{
				Referer:        cloudflare.F("referer"),
				ReferrerPolicy: cloudflare.F("referrerPolicy"),
				Timeout:        cloudflare.F(60000.000000),
				WaitUntil:      cloudflare.F[browser_rendering.AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilUnion](browser_rendering.AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilString(browser_rendering.AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cloudflare.F("html"),
			InterestingOnly:      cloudflare.F(true),
			RejectRequestPattern: cloudflare.F([]string{"string"}),
			RejectResourceTypes:  cloudflare.F([]browser_rendering.AccessibilityTreeNewParamsBodyObjectRejectResourceType{browser_rendering.AccessibilityTreeNewParamsBodyObjectRejectResourceTypeDocument}),
			Root:                 cloudflare.F("root"),
			SetExtraHTTPHeaders: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cloudflare.F(true),
			UserAgent:            cloudflare.F("userAgent"),
			Viewport: cloudflare.F(browser_rendering.AccessibilityTreeNewParamsBodyObjectViewport{
				Height:            cloudflare.F(0.000000),
				Width:             cloudflare.F(0.000000),
				DeviceScaleFactor: cloudflare.F(0.000000),
				HasTouch:          cloudflare.F(true),
				IsLandscape:       cloudflare.F(true),
				IsMobile:          cloudflare.F(true),
			}),
			WaitForSelector: cloudflare.F(browser_rendering.AccessibilityTreeNewParamsBodyObjectWaitForSelector{
				Selector: cloudflare.F("selector"),
				Hidden:   cloudflare.F(browser_rendering.AccessibilityTreeNewParamsBodyObjectWaitForSelectorHiddenTrue),
				Timeout:  cloudflare.F(120000.000000),
				Visible:  cloudflare.F(browser_rendering.AccessibilityTreeNewParamsBodyObjectWaitForSelectorVisibleTrue),
			}),
			WaitForTimeout: cloudflare.F(120000.000000),
		},
		CacheTTL: cloudflare.F(0.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
