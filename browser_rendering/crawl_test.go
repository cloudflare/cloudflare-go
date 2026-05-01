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

func TestCrawlNewWithOptionalParams(t *testing.T) {
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
	_, err := client.BrowserRendering.Crawl.New(context.TODO(), browser_rendering.CrawlNewParams{
		AccountID: cloudflare.F("account_id"),
		Body: browser_rendering.CrawlNewParamsBodyObject{
			URL:           cloudflare.F("https://example.com"),
			ActionTimeout: cloudflare.F(120000.000000),
			AddScriptTag: cloudflare.F([]browser_rendering.CrawlNewParamsBodyObjectAddScriptTag{{
				ID:      cloudflare.F("id"),
				Content: cloudflare.F("content"),
				Type:    cloudflare.F("type"),
				URL:     cloudflare.F("url"),
			}}),
			AddStyleTag: cloudflare.F([]browser_rendering.CrawlNewParamsBodyObjectAddStyleTag{{
				Content: cloudflare.F("content"),
				URL:     cloudflare.F("url"),
			}}),
			AllowRequestPattern: cloudflare.F([]string{"string"}),
			AllowResourceTypes:  cloudflare.F([]browser_rendering.CrawlNewParamsBodyObjectAllowResourceType{browser_rendering.CrawlNewParamsBodyObjectAllowResourceTypeDocument}),
			Authenticate: cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectAuthenticate{
				Password: cloudflare.F("x"),
				Username: cloudflare.F("x"),
			}),
			BestAttempt: cloudflare.F(true),
			Cookies: cloudflare.F([]browser_rendering.CrawlNewParamsBodyObjectCookie{{
				Name:         cloudflare.F("name"),
				Value:        cloudflare.F("value"),
				Domain:       cloudflare.F("domain"),
				Expires:      cloudflare.F(0.000000),
				HTTPOnly:     cloudflare.F(true),
				PartitionKey: cloudflare.F("partitionKey"),
				Path:         cloudflare.F("path"),
				Priority:     cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectCookiesPriorityLow),
				SameParty:    cloudflare.F(true),
				SameSite:     cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectCookiesSameSiteStrict),
				Secure:       cloudflare.F(true),
				SourcePort:   cloudflare.F(0.000000),
				SourceScheme: cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectCookiesSourceSchemeUnset),
				URL:          cloudflare.F("url"),
			}}),
			CrawlPurposes:    cloudflare.F([]browser_rendering.CrawlNewParamsBodyObjectCrawlPurpose{browser_rendering.CrawlNewParamsBodyObjectCrawlPurposeSearch}),
			Depth:            cloudflare.F(1.000000),
			EmulateMediaType: cloudflare.F("emulateMediaType"),
			Formats:          cloudflare.F([]browser_rendering.CrawlNewParamsBodyObjectFormat{browser_rendering.CrawlNewParamsBodyObjectFormatHTML}),
			GotoOptions: cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectGotoOptions{
				Referer:        cloudflare.F("referer"),
				ReferrerPolicy: cloudflare.F("referrerPolicy"),
				Timeout:        cloudflare.F(60000.000000),
				WaitUntil:      cloudflare.F[browser_rendering.CrawlNewParamsBodyObjectGotoOptionsWaitUntilUnion](browser_rendering.CrawlNewParamsBodyObjectGotoOptionsWaitUntilString(browser_rendering.CrawlNewParamsBodyObjectGotoOptionsWaitUntilStringLoad)),
			}),
			JsonOptions: cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectJsonOptions{
				CustomAI: cloudflare.F([]browser_rendering.CrawlNewParamsBodyObjectJsonOptionsCustomAI{{
					Model:         cloudflare.F("model"),
					Authorization: cloudflare.F("authorization"),
				}}),
				Prompt: cloudflare.F("prompt"),
				ResponseFormat: cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectJsonOptionsResponseFormat{
					Type: cloudflare.F("type"),
					JsonSchema: cloudflare.F(map[string]browser_rendering.CrawlNewParamsBodyObjectJsonOptionsResponseFormatJsonSchemaUnion{
						"foo": shared.UnionString("string"),
					}),
				}),
			}),
			Limit:         cloudflare.F(1.000000),
			MaxAge:        cloudflare.F(0.000000),
			ModifiedSince: cloudflare.F(int64(1)),
			Options: cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectOptions{
				ExcludePatterns:      cloudflare.F([]string{"x"}),
				IncludeExternalLinks: cloudflare.F(true),
				IncludePatterns:      cloudflare.F([]string{"x"}),
				IncludeSubdomains:    cloudflare.F(true),
			}),
			RejectRequestPattern: cloudflare.F([]string{"string"}),
			RejectResourceTypes:  cloudflare.F([]browser_rendering.CrawlNewParamsBodyObjectRejectResourceType{browser_rendering.CrawlNewParamsBodyObjectRejectResourceTypeDocument}),
			Render:               cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectRenderTrue),
			SetExtraHTTPHeaders: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cloudflare.F(true),
			Source:               cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectSourceSitemaps),
			Viewport: cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectViewport{
				Height:            cloudflare.F(0.000000),
				Width:             cloudflare.F(0.000000),
				DeviceScaleFactor: cloudflare.F(0.000000),
				HasTouch:          cloudflare.F(true),
				IsLandscape:       cloudflare.F(true),
				IsMobile:          cloudflare.F(true),
			}),
			WaitForSelector: cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectWaitForSelector{
				Selector: cloudflare.F("selector"),
				Hidden:   cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectWaitForSelectorHiddenTrue),
				Timeout:  cloudflare.F(120000.000000),
				Visible:  cloudflare.F(browser_rendering.CrawlNewParamsBodyObjectWaitForSelectorVisibleTrue),
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

func TestCrawlDelete(t *testing.T) {
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
	_, err := client.BrowserRendering.Crawl.Delete(
		context.TODO(),
		"job_id",
		browser_rendering.CrawlDeleteParams{
			AccountID: cloudflare.F("account_id"),
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

func TestCrawlGetWithOptionalParams(t *testing.T) {
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
	_, err := client.BrowserRendering.Crawl.Get(
		context.TODO(),
		"x",
		browser_rendering.CrawlGetParams{
			AccountID: cloudflare.F("account_id"),
			CacheTTL:  cloudflare.F(0.000000),
			Cursor:    cloudflare.F(0.000000),
			Limit:     cloudflare.F(0.000000),
			Status:    cloudflare.F(browser_rendering.CrawlGetParamsStatusQueued),
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
