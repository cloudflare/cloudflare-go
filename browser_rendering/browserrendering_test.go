// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/browser_rendering"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestBrowserRenderingContentWithOptionalParams(t *testing.T) {
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
	_, err := client.BrowserRendering.Content(
		context.TODO(),
		"accountId",
		browser_rendering.BrowserRenderingContentParams{
			CacheTTL: cloudflare.F(86400.000000),
			AddScriptTag: cloudflare.F([]browser_rendering.BrowserRenderingContentParamsAddScriptTag{{
				ID:      cloudflare.F("id"),
				Content: cloudflare.F("content"),
				Type:    cloudflare.F("type"),
				URL:     cloudflare.F("url"),
			}}),
			AddStyleTag: cloudflare.F([]browser_rendering.BrowserRenderingContentParamsAddStyleTag{{
				Content: cloudflare.F("content"),
				URL:     cloudflare.F("url"),
			}}),
			AllowRequestPattern: cloudflare.F([]string{"string"}),
			AllowResourceTypes:  cloudflare.F([]browser_rendering.BrowserRenderingContentParamsAllowResourceType{browser_rendering.BrowserRenderingContentParamsAllowResourceTypeDocument}),
			Authenticate: cloudflare.F(browser_rendering.BrowserRenderingContentParamsAuthenticate{
				Password: cloudflare.F("x"),
				Username: cloudflare.F("x"),
			}),
			BestAttempt: cloudflare.F(true),
			Cookies: cloudflare.F([]browser_rendering.BrowserRenderingContentParamsCookie{{
				Name:         cloudflare.F("name"),
				Value:        cloudflare.F("value"),
				Domain:       cloudflare.F("domain"),
				Expires:      cloudflare.F(0.000000),
				HTTPOnly:     cloudflare.F(true),
				PartitionKey: cloudflare.F("partitionKey"),
				Path:         cloudflare.F("path"),
				Priority:     cloudflare.F(browser_rendering.BrowserRenderingContentParamsCookiesPriorityLow),
				SameParty:    cloudflare.F(true),
				SameSite:     cloudflare.F(browser_rendering.BrowserRenderingContentParamsCookiesSameSiteStrict),
				Secure:       cloudflare.F(true),
				SourcePort:   cloudflare.F(0.000000),
				SourceScheme: cloudflare.F(browser_rendering.BrowserRenderingContentParamsCookiesSourceSchemeUnset),
				URL:          cloudflare.F("url"),
			}}),
			EmulateMediaType: cloudflare.F("emulateMediaType"),
			GotoOptions: cloudflare.F(browser_rendering.BrowserRenderingContentParamsGotoOptions{
				Referer:        cloudflare.F("referer"),
				ReferrerPolicy: cloudflare.F("referrerPolicy"),
				Timeout:        cloudflare.F(60000.000000),
				WaitUntil:      cloudflare.F[browser_rendering.BrowserRenderingContentParamsGotoOptionsWaitUntilUnion](browser_rendering.BrowserRenderingContentParamsGotoOptionsWaitUntilString(browser_rendering.BrowserRenderingContentParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cloudflare.F("x"),
			RejectRequestPattern: cloudflare.F([]string{"string"}),
			RejectResourceTypes:  cloudflare.F([]browser_rendering.BrowserRenderingContentParamsRejectResourceType{browser_rendering.BrowserRenderingContentParamsRejectResourceTypeDocument}),
			SetExtraHTTPHeaders: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cloudflare.F(true),
			URL:                  cloudflare.F("https://example.com"),
			UserAgent:            cloudflare.F("userAgent"),
			Viewport: cloudflare.F(browser_rendering.BrowserRenderingContentParamsViewport{
				Height:            cloudflare.F(0.000000),
				Width:             cloudflare.F(0.000000),
				DeviceScaleFactor: cloudflare.F(0.000000),
				HasTouch:          cloudflare.F(true),
				IsLandscape:       cloudflare.F(true),
				IsMobile:          cloudflare.F(true),
			}),
			WaitForSelector: cloudflare.F(browser_rendering.BrowserRenderingContentParamsWaitForSelector{
				Selector: cloudflare.F("selector"),
				Hidden:   cloudflare.F(browser_rendering.BrowserRenderingContentParamsWaitForSelectorHiddenTrue),
				Timeout:  cloudflare.F(60000.000000),
				Visible:  cloudflare.F(browser_rendering.BrowserRenderingContentParamsWaitForSelectorVisibleTrue),
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

func TestBrowserRenderingPDFWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	resp, err := client.BrowserRendering.PDF(
		context.TODO(),
		"accountId",
		browser_rendering.BrowserRenderingPDFParams{
			CacheTTL: cloudflare.F(86400.000000),
			AddScriptTag: cloudflare.F([]browser_rendering.BrowserRenderingPDFParamsAddScriptTag{{
				ID:      cloudflare.F("id"),
				Content: cloudflare.F("content"),
				Type:    cloudflare.F("type"),
				URL:     cloudflare.F("url"),
			}}),
			AddStyleTag: cloudflare.F([]browser_rendering.BrowserRenderingPDFParamsAddStyleTag{{
				Content: cloudflare.F("content"),
				URL:     cloudflare.F("url"),
			}}),
			AllowRequestPattern: cloudflare.F([]string{"string"}),
			AllowResourceTypes:  cloudflare.F([]browser_rendering.BrowserRenderingPDFParamsAllowResourceType{browser_rendering.BrowserRenderingPDFParamsAllowResourceTypeDocument}),
			Authenticate: cloudflare.F(browser_rendering.BrowserRenderingPDFParamsAuthenticate{
				Password: cloudflare.F("x"),
				Username: cloudflare.F("x"),
			}),
			BestAttempt: cloudflare.F(true),
			Cookies: cloudflare.F([]browser_rendering.BrowserRenderingPDFParamsCookie{{
				Name:         cloudflare.F("name"),
				Value:        cloudflare.F("value"),
				Domain:       cloudflare.F("domain"),
				Expires:      cloudflare.F(0.000000),
				HTTPOnly:     cloudflare.F(true),
				PartitionKey: cloudflare.F("partitionKey"),
				Path:         cloudflare.F("path"),
				Priority:     cloudflare.F(browser_rendering.BrowserRenderingPDFParamsCookiesPriorityLow),
				SameParty:    cloudflare.F(true),
				SameSite:     cloudflare.F(browser_rendering.BrowserRenderingPDFParamsCookiesSameSiteStrict),
				Secure:       cloudflare.F(true),
				SourcePort:   cloudflare.F(0.000000),
				SourceScheme: cloudflare.F(browser_rendering.BrowserRenderingPDFParamsCookiesSourceSchemeUnset),
				URL:          cloudflare.F("url"),
			}}),
			EmulateMediaType: cloudflare.F("emulateMediaType"),
			GotoOptions: cloudflare.F(browser_rendering.BrowserRenderingPDFParamsGotoOptions{
				Referer:        cloudflare.F("referer"),
				ReferrerPolicy: cloudflare.F("referrerPolicy"),
				Timeout:        cloudflare.F(60000.000000),
				WaitUntil:      cloudflare.F[browser_rendering.BrowserRenderingPDFParamsGotoOptionsWaitUntilUnion](browser_rendering.BrowserRenderingPDFParamsGotoOptionsWaitUntilString(browser_rendering.BrowserRenderingPDFParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cloudflare.F("x"),
			RejectRequestPattern: cloudflare.F([]string{"string"}),
			RejectResourceTypes:  cloudflare.F([]browser_rendering.BrowserRenderingPDFParamsRejectResourceType{browser_rendering.BrowserRenderingPDFParamsRejectResourceTypeDocument}),
			SetExtraHTTPHeaders: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cloudflare.F(true),
			URL:                  cloudflare.F("https://example.com"),
			UserAgent:            cloudflare.F("userAgent"),
			Viewport: cloudflare.F(browser_rendering.BrowserRenderingPDFParamsViewport{
				Height:            cloudflare.F(0.000000),
				Width:             cloudflare.F(0.000000),
				DeviceScaleFactor: cloudflare.F(0.000000),
				HasTouch:          cloudflare.F(true),
				IsLandscape:       cloudflare.F(true),
				IsMobile:          cloudflare.F(true),
			}),
			WaitForSelector: cloudflare.F(browser_rendering.BrowserRenderingPDFParamsWaitForSelector{
				Selector: cloudflare.F("selector"),
				Hidden:   cloudflare.F(browser_rendering.BrowserRenderingPDFParamsWaitForSelectorHiddenTrue),
				Timeout:  cloudflare.F(60000.000000),
				Visible:  cloudflare.F(browser_rendering.BrowserRenderingPDFParamsWaitForSelectorVisibleTrue),
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestBrowserRenderingScrapeWithOptionalParams(t *testing.T) {
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
	_, err := client.BrowserRendering.Scrape(
		context.TODO(),
		"accountId",
		browser_rendering.BrowserRenderingScrapeParams{
			Elements: cloudflare.F([]browser_rendering.BrowserRenderingScrapeParamsElement{{
				Selector: cloudflare.F("selector"),
			}}),
			CacheTTL: cloudflare.F(86400.000000),
			AddScriptTag: cloudflare.F([]browser_rendering.BrowserRenderingScrapeParamsAddScriptTag{{
				ID:      cloudflare.F("id"),
				Content: cloudflare.F("content"),
				Type:    cloudflare.F("type"),
				URL:     cloudflare.F("url"),
			}}),
			AddStyleTag: cloudflare.F([]browser_rendering.BrowserRenderingScrapeParamsAddStyleTag{{
				Content: cloudflare.F("content"),
				URL:     cloudflare.F("url"),
			}}),
			AllowRequestPattern: cloudflare.F([]string{"string"}),
			AllowResourceTypes:  cloudflare.F([]browser_rendering.BrowserRenderingScrapeParamsAllowResourceType{browser_rendering.BrowserRenderingScrapeParamsAllowResourceTypeDocument}),
			Authenticate: cloudflare.F(browser_rendering.BrowserRenderingScrapeParamsAuthenticate{
				Password: cloudflare.F("x"),
				Username: cloudflare.F("x"),
			}),
			BestAttempt: cloudflare.F(true),
			Cookies: cloudflare.F([]browser_rendering.BrowserRenderingScrapeParamsCookie{{
				Name:         cloudflare.F("name"),
				Value:        cloudflare.F("value"),
				Domain:       cloudflare.F("domain"),
				Expires:      cloudflare.F(0.000000),
				HTTPOnly:     cloudflare.F(true),
				PartitionKey: cloudflare.F("partitionKey"),
				Path:         cloudflare.F("path"),
				Priority:     cloudflare.F(browser_rendering.BrowserRenderingScrapeParamsCookiesPriorityLow),
				SameParty:    cloudflare.F(true),
				SameSite:     cloudflare.F(browser_rendering.BrowserRenderingScrapeParamsCookiesSameSiteStrict),
				Secure:       cloudflare.F(true),
				SourcePort:   cloudflare.F(0.000000),
				SourceScheme: cloudflare.F(browser_rendering.BrowserRenderingScrapeParamsCookiesSourceSchemeUnset),
				URL:          cloudflare.F("url"),
			}}),
			EmulateMediaType: cloudflare.F("emulateMediaType"),
			GotoOptions: cloudflare.F(browser_rendering.BrowserRenderingScrapeParamsGotoOptions{
				Referer:        cloudflare.F("referer"),
				ReferrerPolicy: cloudflare.F("referrerPolicy"),
				Timeout:        cloudflare.F(60000.000000),
				WaitUntil:      cloudflare.F[browser_rendering.BrowserRenderingScrapeParamsGotoOptionsWaitUntilUnion](browser_rendering.BrowserRenderingScrapeParamsGotoOptionsWaitUntilString(browser_rendering.BrowserRenderingScrapeParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cloudflare.F("x"),
			RejectRequestPattern: cloudflare.F([]string{"string"}),
			RejectResourceTypes:  cloudflare.F([]browser_rendering.BrowserRenderingScrapeParamsRejectResourceType{browser_rendering.BrowserRenderingScrapeParamsRejectResourceTypeDocument}),
			SetExtraHTTPHeaders: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cloudflare.F(true),
			URL:                  cloudflare.F("https://example.com"),
			UserAgent:            cloudflare.F("userAgent"),
			Viewport: cloudflare.F(browser_rendering.BrowserRenderingScrapeParamsViewport{
				Height:            cloudflare.F(0.000000),
				Width:             cloudflare.F(0.000000),
				DeviceScaleFactor: cloudflare.F(0.000000),
				HasTouch:          cloudflare.F(true),
				IsLandscape:       cloudflare.F(true),
				IsMobile:          cloudflare.F(true),
			}),
			WaitForSelector: cloudflare.F(browser_rendering.BrowserRenderingScrapeParamsWaitForSelector{
				Selector: cloudflare.F("selector"),
				Hidden:   cloudflare.F(browser_rendering.BrowserRenderingScrapeParamsWaitForSelectorHiddenTrue),
				Timeout:  cloudflare.F(60000.000000),
				Visible:  cloudflare.F(browser_rendering.BrowserRenderingScrapeParamsWaitForSelectorVisibleTrue),
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

func TestBrowserRenderingScreenshotWithOptionalParams(t *testing.T) {
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
	_, err := client.BrowserRendering.Screenshot(
		context.TODO(),
		"accountId",
		browser_rendering.BrowserRenderingScreenshotParams{
			CacheTTL: cloudflare.F(86400.000000),
			AddScriptTag: cloudflare.F([]browser_rendering.BrowserRenderingScreenshotParamsAddScriptTag{{
				ID:      cloudflare.F("id"),
				Content: cloudflare.F("content"),
				Type:    cloudflare.F("type"),
				URL:     cloudflare.F("url"),
			}}),
			AddStyleTag: cloudflare.F([]browser_rendering.BrowserRenderingScreenshotParamsAddStyleTag{{
				Content: cloudflare.F("content"),
				URL:     cloudflare.F("url"),
			}}),
			AllowRequestPattern: cloudflare.F([]string{"string"}),
			AllowResourceTypes:  cloudflare.F([]browser_rendering.BrowserRenderingScreenshotParamsAllowResourceType{browser_rendering.BrowserRenderingScreenshotParamsAllowResourceTypeDocument}),
			Authenticate: cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsAuthenticate{
				Password: cloudflare.F("x"),
				Username: cloudflare.F("x"),
			}),
			BestAttempt: cloudflare.F(true),
			Cookies: cloudflare.F([]browser_rendering.BrowserRenderingScreenshotParamsCookie{{
				Name:         cloudflare.F("name"),
				Value:        cloudflare.F("value"),
				Domain:       cloudflare.F("domain"),
				Expires:      cloudflare.F(0.000000),
				HTTPOnly:     cloudflare.F(true),
				PartitionKey: cloudflare.F("partitionKey"),
				Path:         cloudflare.F("path"),
				Priority:     cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsCookiesPriorityLow),
				SameParty:    cloudflare.F(true),
				SameSite:     cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsCookiesSameSiteStrict),
				Secure:       cloudflare.F(true),
				SourcePort:   cloudflare.F(0.000000),
				SourceScheme: cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsCookiesSourceSchemeUnset),
				URL:          cloudflare.F("url"),
			}}),
			EmulateMediaType: cloudflare.F("emulateMediaType"),
			GotoOptions: cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsGotoOptions{
				Referer:        cloudflare.F("referer"),
				ReferrerPolicy: cloudflare.F("referrerPolicy"),
				Timeout:        cloudflare.F(60000.000000),
				WaitUntil:      cloudflare.F[browser_rendering.BrowserRenderingScreenshotParamsGotoOptionsWaitUntilUnion](browser_rendering.BrowserRenderingScreenshotParamsGotoOptionsWaitUntilString(browser_rendering.BrowserRenderingScreenshotParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cloudflare.F("x"),
			RejectRequestPattern: cloudflare.F([]string{"string"}),
			RejectResourceTypes:  cloudflare.F([]browser_rendering.BrowserRenderingScreenshotParamsRejectResourceType{browser_rendering.BrowserRenderingScreenshotParamsRejectResourceTypeDocument}),
			ScreenshotOptions: cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsScreenshotOptions{
				CaptureBeyondViewport: cloudflare.F(true),
				Clip: cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsScreenshotOptionsClip{
					Height: cloudflare.F(0.000000),
					Width:  cloudflare.F(0.000000),
					X:      cloudflare.F(0.000000),
					Y:      cloudflare.F(0.000000),
					Scale:  cloudflare.F(0.000000),
				}),
				Encoding:         cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsScreenshotOptionsEncodingBinary),
				FromSurface:      cloudflare.F(true),
				FullPage:         cloudflare.F(true),
				OmitBackground:   cloudflare.F(true),
				OptimizeForSpeed: cloudflare.F(true),
				Quality:          cloudflare.F(0.000000),
				Type:             cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsScreenshotOptionsTypePNG),
			}),
			ScrollPage: cloudflare.F(true),
			Selector:   cloudflare.F("selector"),
			SetExtraHTTPHeaders: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cloudflare.F(true),
			URL:                  cloudflare.F("https://example.com"),
			UserAgent:            cloudflare.F("userAgent"),
			Viewport: cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsViewport{
				Height:            cloudflare.F(0.000000),
				Width:             cloudflare.F(0.000000),
				DeviceScaleFactor: cloudflare.F(0.000000),
				HasTouch:          cloudflare.F(true),
				IsLandscape:       cloudflare.F(true),
				IsMobile:          cloudflare.F(true),
			}),
			WaitForSelector: cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsWaitForSelector{
				Selector: cloudflare.F("selector"),
				Hidden:   cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsWaitForSelectorHiddenTrue),
				Timeout:  cloudflare.F(60000.000000),
				Visible:  cloudflare.F(browser_rendering.BrowserRenderingScreenshotParamsWaitForSelectorVisibleTrue),
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

func TestBrowserRenderingSnapshotWithOptionalParams(t *testing.T) {
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
	_, err := client.BrowserRendering.Snapshot(
		context.TODO(),
		"accountId",
		browser_rendering.BrowserRenderingSnapshotParams{
			CacheTTL: cloudflare.F(86400.000000),
			AddScriptTag: cloudflare.F([]browser_rendering.BrowserRenderingSnapshotParamsAddScriptTag{{
				ID:      cloudflare.F("id"),
				Content: cloudflare.F("content"),
				Type:    cloudflare.F("type"),
				URL:     cloudflare.F("url"),
			}}),
			AddStyleTag: cloudflare.F([]browser_rendering.BrowserRenderingSnapshotParamsAddStyleTag{{
				Content: cloudflare.F("content"),
				URL:     cloudflare.F("url"),
			}}),
			AllowRequestPattern: cloudflare.F([]string{"string"}),
			AllowResourceTypes:  cloudflare.F([]browser_rendering.BrowserRenderingSnapshotParamsAllowResourceType{browser_rendering.BrowserRenderingSnapshotParamsAllowResourceTypeDocument}),
			Authenticate: cloudflare.F(browser_rendering.BrowserRenderingSnapshotParamsAuthenticate{
				Password: cloudflare.F("x"),
				Username: cloudflare.F("x"),
			}),
			BestAttempt: cloudflare.F(true),
			Cookies: cloudflare.F([]browser_rendering.BrowserRenderingSnapshotParamsCookie{{
				Name:         cloudflare.F("name"),
				Value:        cloudflare.F("value"),
				Domain:       cloudflare.F("domain"),
				Expires:      cloudflare.F(0.000000),
				HTTPOnly:     cloudflare.F(true),
				PartitionKey: cloudflare.F("partitionKey"),
				Path:         cloudflare.F("path"),
				Priority:     cloudflare.F(browser_rendering.BrowserRenderingSnapshotParamsCookiesPriorityLow),
				SameParty:    cloudflare.F(true),
				SameSite:     cloudflare.F(browser_rendering.BrowserRenderingSnapshotParamsCookiesSameSiteStrict),
				Secure:       cloudflare.F(true),
				SourcePort:   cloudflare.F(0.000000),
				SourceScheme: cloudflare.F(browser_rendering.BrowserRenderingSnapshotParamsCookiesSourceSchemeUnset),
				URL:          cloudflare.F("url"),
			}}),
			EmulateMediaType: cloudflare.F("emulateMediaType"),
			GotoOptions: cloudflare.F(browser_rendering.BrowserRenderingSnapshotParamsGotoOptions{
				Referer:        cloudflare.F("referer"),
				ReferrerPolicy: cloudflare.F("referrerPolicy"),
				Timeout:        cloudflare.F(60000.000000),
				WaitUntil:      cloudflare.F[browser_rendering.BrowserRenderingSnapshotParamsGotoOptionsWaitUntilUnion](browser_rendering.BrowserRenderingSnapshotParamsGotoOptionsWaitUntilString(browser_rendering.BrowserRenderingSnapshotParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cloudflare.F("x"),
			RejectRequestPattern: cloudflare.F([]string{"string"}),
			RejectResourceTypes:  cloudflare.F([]browser_rendering.BrowserRenderingSnapshotParamsRejectResourceType{browser_rendering.BrowserRenderingSnapshotParamsRejectResourceTypeDocument}),
			SetExtraHTTPHeaders: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cloudflare.F(true),
			URL:                  cloudflare.F("https://example.com"),
			UserAgent:            cloudflare.F("userAgent"),
			Viewport: cloudflare.F(browser_rendering.BrowserRenderingSnapshotParamsViewport{
				Height:            cloudflare.F(0.000000),
				Width:             cloudflare.F(0.000000),
				DeviceScaleFactor: cloudflare.F(0.000000),
				HasTouch:          cloudflare.F(true),
				IsLandscape:       cloudflare.F(true),
				IsMobile:          cloudflare.F(true),
			}),
			WaitForSelector: cloudflare.F(browser_rendering.BrowserRenderingSnapshotParamsWaitForSelector{
				Selector: cloudflare.F("selector"),
				Hidden:   cloudflare.F(browser_rendering.BrowserRenderingSnapshotParamsWaitForSelectorHiddenTrue),
				Timeout:  cloudflare.F(60000.000000),
				Visible:  cloudflare.F(browser_rendering.BrowserRenderingSnapshotParamsWaitForSelectorVisibleTrue),
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
