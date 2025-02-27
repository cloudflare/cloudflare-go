// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browsing_rendering_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/browsing_rendering"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestPdfNewWithOptionalParams(t *testing.T) {
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
	resp, err := client.BrowsingRendering.Pdf.New(
		context.TODO(),
		"accountId",
		browsing_rendering.PdfNewParams{
			CacheTTL: cloudflare.F(86400.000000),
			AddScriptTag: cloudflare.F([]browsing_rendering.PdfNewParamsAddScriptTag{{
				ID:      cloudflare.F("id"),
				Content: cloudflare.F("content"),
				Type:    cloudflare.F("type"),
				URL:     cloudflare.F("url"),
			}}),
			AddStyleTag: cloudflare.F([]browsing_rendering.PdfNewParamsAddStyleTag{{
				Content: cloudflare.F("content"),
				URL:     cloudflare.F("url"),
			}}),
			AllowRequestPattern: cloudflare.F([]string{"string"}),
			AllowResourceTypes:  cloudflare.F([]browsing_rendering.PdfNewParamsAllowResourceType{browsing_rendering.PdfNewParamsAllowResourceTypeDocument}),
			Authenticate: cloudflare.F(browsing_rendering.PdfNewParamsAuthenticate{
				Password: cloudflare.F("x"),
				Username: cloudflare.F("x"),
			}),
			BestAttempt: cloudflare.F(true),
			Cookies: cloudflare.F([]browsing_rendering.PdfNewParamsCookie{{
				Name:         cloudflare.F("name"),
				Value:        cloudflare.F("value"),
				Domain:       cloudflare.F("domain"),
				Expires:      cloudflare.F(0.000000),
				HTTPOnly:     cloudflare.F(true),
				PartitionKey: cloudflare.F("partitionKey"),
				Path:         cloudflare.F("path"),
				Priority:     cloudflare.F(browsing_rendering.PdfNewParamsCookiesPriorityLow),
				SameParty:    cloudflare.F(true),
				SameSite:     cloudflare.F(browsing_rendering.PdfNewParamsCookiesSameSiteStrict),
				Secure:       cloudflare.F(true),
				SourcePort:   cloudflare.F(0.000000),
				SourceScheme: cloudflare.F(browsing_rendering.PdfNewParamsCookiesSourceSchemeUnset),
				URL:          cloudflare.F("url"),
			}}),
			EmulateMediaType: cloudflare.F("emulateMediaType"),
			GotoOptions: cloudflare.F(browsing_rendering.PdfNewParamsGotoOptions{
				Referer:        cloudflare.F("referer"),
				ReferrerPolicy: cloudflare.F("referrerPolicy"),
				Timeout:        cloudflare.F(60000.000000),
				WaitUntil:      cloudflare.F[browsing_rendering.PdfNewParamsGotoOptionsWaitUntilUnion](browsing_rendering.PdfNewParamsGotoOptionsWaitUntilString(browsing_rendering.PdfNewParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cloudflare.F("x"),
			RejectRequestPattern: cloudflare.F([]string{"string"}),
			RejectResourceTypes:  cloudflare.F([]browsing_rendering.PdfNewParamsRejectResourceType{browsing_rendering.PdfNewParamsRejectResourceTypeDocument}),
			SetExtraHTTPHeaders: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cloudflare.F(true),
			URL:                  cloudflare.F("https://example.com"),
			UserAgent:            cloudflare.F("userAgent"),
			Viewport: cloudflare.F(browsing_rendering.PdfNewParamsViewport{
				Height:            cloudflare.F(0.000000),
				Width:             cloudflare.F(0.000000),
				DeviceScaleFactor: cloudflare.F(0.000000),
				HasTouch:          cloudflare.F(true),
				IsLandscape:       cloudflare.F(true),
				IsMobile:          cloudflare.F(true),
			}),
			WaitForSelector: cloudflare.F(browsing_rendering.PdfNewParamsWaitForSelector{
				Selector: cloudflare.F("selector"),
				Hidden:   cloudflare.F(browsing_rendering.PdfNewParamsWaitForSelectorHiddenTrue),
				Timeout:  cloudflare.F(60000.000000),
				Visible:  cloudflare.F(browsing_rendering.PdfNewParamsWaitForSelectorVisibleTrue),
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
