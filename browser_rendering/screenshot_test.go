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
)

func TestScreenshotNewWithOptionalParams(t *testing.T) {
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
	_, err := client.BrowserRendering.Screenshot.New(context.TODO(), browser_rendering.ScreenshotNewParams{
		AccountID: cloudflare.F("account_id"),
		Body: browser_rendering.ScreenshotNewParamsBodyObject{
			URL:           cloudflare.F("https://www.example.com/"),
			ActionTimeout: cloudflare.F(120000.000000),
			AddScriptTag: cloudflare.F([]browser_rendering.ScreenshotNewParamsBodyObjectAddScriptTag{{
				ID:      cloudflare.F("id"),
				Content: cloudflare.F("content"),
				Type:    cloudflare.F("type"),
				URL:     cloudflare.F("url"),
			}}),
			AddStyleTag: cloudflare.F([]browser_rendering.ScreenshotNewParamsBodyObjectAddStyleTag{{
				Content: cloudflare.F("content"),
				URL:     cloudflare.F("url"),
			}}),
			AllowRequestPattern: cloudflare.F([]string{"string"}),
			AllowResourceTypes:  cloudflare.F([]browser_rendering.ScreenshotNewParamsBodyObjectAllowResourceType{browser_rendering.ScreenshotNewParamsBodyObjectAllowResourceTypeDocument}),
			Authenticate: cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectAuthenticate{
				Password: cloudflare.F("x"),
				Username: cloudflare.F("x"),
			}),
			BestAttempt: cloudflare.F(true),
			Cookies: cloudflare.F([]browser_rendering.ScreenshotNewParamsBodyObjectCookie{{
				Name:         cloudflare.F("name"),
				Value:        cloudflare.F("value"),
				Domain:       cloudflare.F("domain"),
				Expires:      cloudflare.F(0.000000),
				HTTPOnly:     cloudflare.F(true),
				PartitionKey: cloudflare.F("partitionKey"),
				Path:         cloudflare.F("path"),
				Priority:     cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectCookiesPriorityLow),
				SameParty:    cloudflare.F(true),
				SameSite:     cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectCookiesSameSiteStrict),
				Secure:       cloudflare.F(true),
				SourcePort:   cloudflare.F(0.000000),
				SourceScheme: cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectCookiesSourceSchemeUnset),
				URL:          cloudflare.F("url"),
			}}),
			EmulateMediaType: cloudflare.F("emulateMediaType"),
			GotoOptions: cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectGotoOptions{
				Referer:        cloudflare.F("referer"),
				ReferrerPolicy: cloudflare.F("referrerPolicy"),
				Timeout:        cloudflare.F(60000.000000),
				WaitUntil:      cloudflare.F[browser_rendering.ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilUnion](browser_rendering.ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilString(browser_rendering.ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilStringLoad)),
			}),
			RejectRequestPattern: cloudflare.F([]string{"string"}),
			RejectResourceTypes:  cloudflare.F([]browser_rendering.ScreenshotNewParamsBodyObjectRejectResourceType{browser_rendering.ScreenshotNewParamsBodyObjectRejectResourceTypeDocument}),
			ScreenshotOptions: cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectScreenshotOptions{
				CaptureBeyondViewport: cloudflare.F(true),
				Clip: cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectScreenshotOptionsClip{
					Height: cloudflare.F(0.000000),
					Width:  cloudflare.F(0.000000),
					X:      cloudflare.F(0.000000),
					Y:      cloudflare.F(0.000000),
					Scale:  cloudflare.F(0.000000),
				}),
				Encoding:         cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectScreenshotOptionsEncodingBinary),
				FromSurface:      cloudflare.F(true),
				FullPage:         cloudflare.F(true),
				OmitBackground:   cloudflare.F(true),
				OptimizeForSpeed: cloudflare.F(true),
				Quality:          cloudflare.F(0.000000),
				Type:             cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectScreenshotOptionsTypePNG),
			}),
			ScrollPage: cloudflare.F(true),
			Selector:   cloudflare.F("selector"),
			SetExtraHTTPHeaders: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cloudflare.F(true),
			UserAgent:            cloudflare.F("userAgent"),
			Viewport: cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectViewport{
				Height:            cloudflare.F(0.000000),
				Width:             cloudflare.F(0.000000),
				DeviceScaleFactor: cloudflare.F(0.000000),
				HasTouch:          cloudflare.F(true),
				IsLandscape:       cloudflare.F(true),
				IsMobile:          cloudflare.F(true),
			}),
			WaitForSelector: cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectWaitForSelector{
				Selector: cloudflare.F("selector"),
				Hidden:   cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectWaitForSelectorHiddenTrue),
				Timeout:  cloudflare.F(120000.000000),
				Visible:  cloudflare.F(browser_rendering.ScreenshotNewParamsBodyObjectWaitForSelectorVisibleTrue),
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
