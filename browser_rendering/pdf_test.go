// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/browser_rendering"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

func TestPDFNewWithOptionalParams(t *testing.T) {
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
	resp, err := client.BrowserRendering.PDF.New(context.TODO(), browser_rendering.PDFNewParams{
		AccountID: cloudflare.F("account_id"),
		Body: browser_rendering.PDFNewParamsBodyObject{
			URL:           cloudflare.F("https://www.example.com/"),
			ActionTimeout: cloudflare.F(120000.000000),
			AddScriptTag: cloudflare.F([]browser_rendering.PDFNewParamsBodyObjectAddScriptTag{{
				ID:      cloudflare.F("id"),
				Content: cloudflare.F("content"),
				Type:    cloudflare.F("type"),
				URL:     cloudflare.F("url"),
			}}),
			AddStyleTag: cloudflare.F([]browser_rendering.PDFNewParamsBodyObjectAddStyleTag{{
				Content: cloudflare.F("content"),
				URL:     cloudflare.F("url"),
			}}),
			AllowRequestPattern: cloudflare.F([]string{"string"}),
			AllowResourceTypes:  cloudflare.F([]browser_rendering.PDFNewParamsBodyObjectAllowResourceType{browser_rendering.PDFNewParamsBodyObjectAllowResourceTypeDocument}),
			Authenticate: cloudflare.F(browser_rendering.PDFNewParamsBodyObjectAuthenticate{
				Password: cloudflare.F("x"),
				Username: cloudflare.F("x"),
			}),
			BestAttempt: cloudflare.F(true),
			Cookies: cloudflare.F([]browser_rendering.PDFNewParamsBodyObjectCookie{{
				Name:         cloudflare.F("name"),
				Value:        cloudflare.F("value"),
				Domain:       cloudflare.F("domain"),
				Expires:      cloudflare.F(0.000000),
				HTTPOnly:     cloudflare.F(true),
				PartitionKey: cloudflare.F("partitionKey"),
				Path:         cloudflare.F("path"),
				Priority:     cloudflare.F(browser_rendering.PDFNewParamsBodyObjectCookiesPriorityLow),
				SameParty:    cloudflare.F(true),
				SameSite:     cloudflare.F(browser_rendering.PDFNewParamsBodyObjectCookiesSameSiteStrict),
				Secure:       cloudflare.F(true),
				SourcePort:   cloudflare.F(0.000000),
				SourceScheme: cloudflare.F(browser_rendering.PDFNewParamsBodyObjectCookiesSourceSchemeUnset),
				URL:          cloudflare.F("url"),
			}}),
			EmulateMediaType: cloudflare.F("emulateMediaType"),
			GotoOptions: cloudflare.F(browser_rendering.PDFNewParamsBodyObjectGotoOptions{
				Referer:        cloudflare.F("referer"),
				ReferrerPolicy: cloudflare.F("referrerPolicy"),
				Timeout:        cloudflare.F(60000.000000),
				WaitUntil:      cloudflare.F[browser_rendering.PDFNewParamsBodyObjectGotoOptionsWaitUntilUnion](browser_rendering.PDFNewParamsBodyObjectGotoOptionsWaitUntilString(browser_rendering.PDFNewParamsBodyObjectGotoOptionsWaitUntilStringLoad)),
			}),
			PDFOptions: cloudflare.F(browser_rendering.PDFNewParamsBodyObjectPDFOptions{
				DisplayHeaderFooter: cloudflare.F(true),
				FooterTemplate:      cloudflare.F("footerTemplate"),
				Format:              cloudflare.F(browser_rendering.PDFNewParamsBodyObjectPDFOptionsFormatLetter),
				HeaderTemplate:      cloudflare.F("headerTemplate"),
				Height:              cloudflare.F[browser_rendering.PDFNewParamsBodyObjectPDFOptionsHeightUnion](shared.UnionString("string")),
				Landscape:           cloudflare.F(true),
				Margin: cloudflare.F(browser_rendering.PDFNewParamsBodyObjectPDFOptionsMargin{
					Bottom: cloudflare.F[browser_rendering.PDFNewParamsBodyObjectPDFOptionsMarginBottomUnion](shared.UnionString("string")),
					Left:   cloudflare.F[browser_rendering.PDFNewParamsBodyObjectPDFOptionsMarginLeftUnion](shared.UnionString("string")),
					Right:  cloudflare.F[browser_rendering.PDFNewParamsBodyObjectPDFOptionsMarginRightUnion](shared.UnionString("string")),
					Top:    cloudflare.F[browser_rendering.PDFNewParamsBodyObjectPDFOptionsMarginTopUnion](shared.UnionString("string")),
				}),
				OmitBackground:    cloudflare.F(true),
				Outline:           cloudflare.F(true),
				PageRanges:        cloudflare.F("pageRanges"),
				PreferCSSPageSize: cloudflare.F(true),
				PrintBackground:   cloudflare.F(true),
				Scale:             cloudflare.F(0.100000),
				Tagged:            cloudflare.F(true),
				Timeout:           cloudflare.F(0.000000),
				Width:             cloudflare.F[browser_rendering.PDFNewParamsBodyObjectPDFOptionsWidthUnion](shared.UnionString("string")),
			}),
			RejectRequestPattern: cloudflare.F([]string{"string"}),
			RejectResourceTypes:  cloudflare.F([]browser_rendering.PDFNewParamsBodyObjectRejectResourceType{browser_rendering.PDFNewParamsBodyObjectRejectResourceTypeDocument}),
			SetExtraHTTPHeaders: cloudflare.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cloudflare.F(true),
			UserAgent:            cloudflare.F("userAgent"),
			Viewport: cloudflare.F(browser_rendering.PDFNewParamsBodyObjectViewport{
				Height:            cloudflare.F(0.000000),
				Width:             cloudflare.F(0.000000),
				DeviceScaleFactor: cloudflare.F(0.000000),
				HasTouch:          cloudflare.F(true),
				IsLandscape:       cloudflare.F(true),
				IsMobile:          cloudflare.F(true),
			}),
			WaitForSelector: cloudflare.F(browser_rendering.PDFNewParamsBodyObjectWaitForSelector{
				Selector: cloudflare.F("selector"),
				Hidden:   cloudflare.F(browser_rendering.PDFNewParamsBodyObjectWaitForSelectorHiddenTrue),
				Timeout:  cloudflare.F(120000.000000),
				Visible:  cloudflare.F(browser_rendering.PDFNewParamsBodyObjectWaitForSelectorVisibleTrue),
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
