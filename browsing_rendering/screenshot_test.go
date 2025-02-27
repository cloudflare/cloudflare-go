// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browsing_rendering_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/browsing_rendering"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
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
	_, err := client.BrowsingRendering.Screenshot.New(
		context.TODO(),
		"accountId",
		browsing_rendering.ScreenshotNewParams{
			CacheTTL: cloudflare.F(86400.000000),
			AddScriptTag: cloudflare.F([]browsing_rendering.ScreenshotNewParamsAddScriptTag{{
				ID:      cloudflare.F("id"),
				Content: cloudflare.F("content"),
				Type:    cloudflare.F("type"),
				URL:     cloudflare.F("url"),
			}}),
			AddStyleTag: cloudflare.F([]browsing_rendering.ScreenshotNewParamsAddStyleTag{{
				Content: cloudflare.F("content"),
				URL:     cloudflare.F("url"),
			}}),
			AllowRequestPattern: cloudflare.F([]string{"string"}),
			AllowResourceTypes:  cloudflare.F([]browsing_rendering.ScreenshotNewParamsAllowResourceType{browsing_rendering.ScreenshotNewParamsAllowResourceTypeDocument}),
			Authenticate: cloudflare.F(browsing_rendering.ScreenshotNewParamsAuthenticate{
				Password: cloudflare.F("x"),
				Username: cloudflare.F("x"),
			}),
			BestAttempt: cloudflare.F(true),
			Cookies: cloudflare.F([]browsing_rendering.ScreenshotNewParamsCookie{{
				Name:         cloudflare.F("name"),
				Value:        cloudflare.F("value"),
				Domain:       cloudflare.F("domain"),
				Expires:      cloudflare.F(0.000000),
				HTTPOnly:     cloudflare.F(true),
				PartitionKey: cloudflare.F("partitionKey"),
				Path:         cloudflare.F("path"),
				Priority:     cloudflare.F(browsing_rendering.ScreenshotNewParamsCookiesPriorityLow),
				SameParty:    cloudflare.F(true),
				SameSite:     cloudflare.F(browsing_rendering.ScreenshotNewParamsCookiesSameSiteStrict),
				Secure:       cloudflare.F(true),
				SourcePort:   cloudflare.F(0.000000),
				SourceScheme: cloudflare.F(browsing_rendering.ScreenshotNewParamsCookiesSourceSchemeUnset),
				URL:          cloudflare.F("url"),
			}}),
			EmulateMediaType: cloudflare.F("emulateMediaType"),
			GotoOptions: cloudflare.F(browsing_rendering.ScreenshotNewParamsGotoOptions{
				Referer:        cloudflare.F("referer"),
				ReferrerPolicy: cloudflare.F("referrerPolicy"),
				Timeout:        cloudflare.F(60000.000000),
				WaitUntil:      cloudflare.F[browsing_rendering.ScreenshotNewParamsGotoOptionsWaitUntilUnion](browsing_rendering.ScreenshotNewParamsGotoOptionsWaitUntilString(browsing_rendering.ScreenshotNewParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cloudflare.F("x"),
			RejectRequestPattern: cloudflare.F([]string{"string"}),
			RejectResourceTypes:  cloudflare.F([]browsing_rendering.ScreenshotNewParamsRejectResourceType{browsing_rendering.ScreenshotNewParamsRejectResourceTypeDocument}),
			ScreenshotOptions: cloudflare.F(browsing_rendering.ScreenshotNewParamsScreenshotOptions{
				CaptureBeyondViewport: cloudflare.F(true),
				Clip: cloudflare.F(browsing_rendering.ScreenshotNewParamsScreenshotOptionsClip{
					Height: cloudflare.F(0.000000),
					Width:  cloudflare.F(0.000000),
					X:      cloudflare.F(0.000000),
					Y:      cloudflare.F(0.000000),
					Scale:  cloudflare.F(0.000000),
				}),
				Encoding:         cloudflare.F(browsing_rendering.ScreenshotNewParamsScreenshotOptionsEncodingBinary),
				FromSurface:      cloudflare.F(true),
				FullPage:         cloudflare.F(true),
				OmitBackground:   cloudflare.F(true),
				OptimizeForSpeed: cloudflare.F(true),
				Quality:          cloudflare.F(0.000000),
				Type:             cloudflare.F(browsing_rendering.ScreenshotNewParamsScreenshotOptionsTypePNG),
			}),
			ScrollPage: cloudflare.F(true),
			Selector:   cloudflare.F("selector"),
			SetExtraHTTPHeaders: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cloudflare.F(true),
			URL:                  cloudflare.F("https://example.com"),
			UserAgent:            cloudflare.F("userAgent"),
			Viewport: cloudflare.F(browsing_rendering.ScreenshotNewParamsViewport{
				Height:            cloudflare.F(0.000000),
				Width:             cloudflare.F(0.000000),
				DeviceScaleFactor: cloudflare.F(0.000000),
				HasTouch:          cloudflare.F(true),
				IsLandscape:       cloudflare.F(true),
				IsMobile:          cloudflare.F(true),
			}),
			WaitForSelector: cloudflare.F(browsing_rendering.ScreenshotNewParamsWaitForSelector{
				Selector: cloudflare.F("selector"),
				Hidden:   cloudflare.F(browsing_rendering.ScreenshotNewParamsWaitForSelectorHiddenTrue),
				Timeout:  cloudflare.F(60000.000000),
				Visible:  cloudflare.F(browsing_rendering.ScreenshotNewParamsWaitForSelectorVisibleTrue),
			}),
			WaitForTimeout: cloudflare.F(60000.000000),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
