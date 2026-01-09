// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/browser_rendering"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

func TestJsonNewWithOptionalParams(t *testing.T) {
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
	_, err := client.BrowserRendering.Json.New(context.TODO(), browser_rendering.JsonNewParams{
		AccountID: cloudflare.F("account_id"),
		Body: browser_rendering.JsonNewParamsBodyObject{
			URL:           cloudflare.F("https://www.example.com/"),
			ActionTimeout: cloudflare.F(120000.000000),
			AddScriptTag: cloudflare.F([]browser_rendering.JsonNewParamsBodyObjectAddScriptTag{{
				ID:      cloudflare.F("id"),
				Content: cloudflare.F("content"),
				Type:    cloudflare.F("type"),
				URL:     cloudflare.F("url"),
			}}),
			AddStyleTag: cloudflare.F([]browser_rendering.JsonNewParamsBodyObjectAddStyleTag{{
				Content: cloudflare.F("content"),
				URL:     cloudflare.F("url"),
			}}),
			AllowRequestPattern: cloudflare.F([]string{"string"}),
			AllowResourceTypes:  cloudflare.F([]browser_rendering.JsonNewParamsBodyObjectAllowResourceType{browser_rendering.JsonNewParamsBodyObjectAllowResourceTypeDocument}),
			Authenticate: cloudflare.F(browser_rendering.JsonNewParamsBodyObjectAuthenticate{
				Password: cloudflare.F("x"),
				Username: cloudflare.F("x"),
			}),
			BestAttempt: cloudflare.F(true),
			Cookies: cloudflare.F([]browser_rendering.JsonNewParamsBodyObjectCookie{{
				Name:         cloudflare.F("name"),
				Value:        cloudflare.F("value"),
				Domain:       cloudflare.F("domain"),
				Expires:      cloudflare.F(0.000000),
				HTTPOnly:     cloudflare.F(true),
				PartitionKey: cloudflare.F("partitionKey"),
				Path:         cloudflare.F("path"),
				Priority:     cloudflare.F(browser_rendering.JsonNewParamsBodyObjectCookiesPriorityLow),
				SameParty:    cloudflare.F(true),
				SameSite:     cloudflare.F(browser_rendering.JsonNewParamsBodyObjectCookiesSameSiteStrict),
				Secure:       cloudflare.F(true),
				SourcePort:   cloudflare.F(0.000000),
				SourceScheme: cloudflare.F(browser_rendering.JsonNewParamsBodyObjectCookiesSourceSchemeUnset),
				URL:          cloudflare.F("url"),
			}}),
			CustomAI: cloudflare.F([]browser_rendering.JsonNewParamsBodyObjectCustomAI{{
				Authorization: cloudflare.F("authorization"),
				Model:         cloudflare.F("model"),
			}}),
			EmulateMediaType: cloudflare.F("emulateMediaType"),
			GotoOptions: cloudflare.F(browser_rendering.JsonNewParamsBodyObjectGotoOptions{
				Referer:        cloudflare.F("referer"),
				ReferrerPolicy: cloudflare.F("referrerPolicy"),
				Timeout:        cloudflare.F(60000.000000),
				WaitUntil:      cloudflare.F[browser_rendering.JsonNewParamsBodyObjectGotoOptionsWaitUntilUnion](browser_rendering.JsonNewParamsBodyObjectGotoOptionsWaitUntilString(browser_rendering.JsonNewParamsBodyObjectGotoOptionsWaitUntilStringLoad)),
			}),
			Prompt:               cloudflare.F("prompt"),
			RejectRequestPattern: cloudflare.F([]string{"string"}),
			RejectResourceTypes:  cloudflare.F([]browser_rendering.JsonNewParamsBodyObjectRejectResourceType{browser_rendering.JsonNewParamsBodyObjectRejectResourceTypeDocument}),
			ResponseFormat: cloudflare.F(browser_rendering.JsonNewParamsBodyObjectResponseFormat{
				Type: cloudflare.F("type"),
				JsonSchema: cloudflare.F(map[string]browser_rendering.JsonNewParamsBodyObjectResponseFormatJsonSchemaUnion{
					"foo": shared.UnionString("string"),
				}),
			}),
			SetExtraHTTPHeaders: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cloudflare.F(true),
			UserAgent:            cloudflare.F("userAgent"),
			Viewport: cloudflare.F(browser_rendering.JsonNewParamsBodyObjectViewport{
				Height:            cloudflare.F(0.000000),
				Width:             cloudflare.F(0.000000),
				DeviceScaleFactor: cloudflare.F(0.000000),
				HasTouch:          cloudflare.F(true),
				IsLandscape:       cloudflare.F(true),
				IsMobile:          cloudflare.F(true),
			}),
			WaitForSelector: cloudflare.F(browser_rendering.JsonNewParamsBodyObjectWaitForSelector{
				Selector: cloudflare.F("selector"),
				Hidden:   cloudflare.F(browser_rendering.JsonNewParamsBodyObjectWaitForSelectorHiddenTrue),
				Timeout:  cloudflare.F(120000.000000),
				Visible:  cloudflare.F(browser_rendering.JsonNewParamsBodyObjectWaitForSelectorVisibleTrue),
			}),
			WaitForTimeout: cloudflare.F(120000.000000),
		},
		CacheTTL: cloudflare.F(86400.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
