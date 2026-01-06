// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// PDFService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPDFService] method instead.
type PDFService struct {
	Options []option.RequestOption
}

// NewPDFService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPDFService(opts ...option.RequestOption) (r *PDFService) {
	r = &PDFService{}
	r.Options = opts
	return
}

// Fetches rendered PDF from provided URL or HTML. Check available options like
// `gotoOptions` and `waitFor*` to control page load behaviour.
func (r *PDFService) New(ctx context.Context, params PDFNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/pdf", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type PDFNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64]  `query:"cacheTTL"`
	Body     PDFNewParamsBodyUnion `json:"body"`
}

func (r PDFNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [PDFNewParams]'s query parameters as `url.Values`.
func (r PDFNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PDFNewParamsBody struct {
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout       param.Field[float64]     `json:"actionTimeout"`
	AddScriptTag        param.Field[interface{}] `json:"addScriptTag"`
	AddStyleTag         param.Field[interface{}] `json:"addStyleTag"`
	AllowRequestPattern param.Field[interface{}] `json:"allowRequestPattern"`
	AllowResourceTypes  param.Field[interface{}] `json:"allowResourceTypes"`
	Authenticate        param.Field[interface{}] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt      param.Field[bool]        `json:"bestAttempt"`
	Cookies          param.Field[interface{}] `json:"cookies"`
	EmulateMediaType param.Field[string]      `json:"emulateMediaType"`
	GotoOptions      param.Field[interface{}] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML                 param.Field[string]      `json:"html"`
	PDFOptions           param.Field[interface{}] `json:"pdfOptions"`
	RejectRequestPattern param.Field[interface{}] `json:"rejectRequestPattern"`
	RejectResourceTypes  param.Field[interface{}] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[interface{}] `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]        `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL             param.Field[string]      `json:"url" format:"uri"`
	UserAgent       param.Field[string]      `json:"userAgent"`
	Viewport        param.Field[interface{}] `json:"viewport"`
	WaitForSelector param.Field[interface{}] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r PDFNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PDFNewParamsBody) implementsPDFNewParamsBodyUnion() {}

// Satisfied by [browser_rendering.PDFNewParamsBodyObject],
// [browser_rendering.PDFNewParamsBodyObject], [PDFNewParamsBody].
type PDFNewParamsBodyUnion interface {
	implementsPDFNewParamsBodyUnion()
}

type PDFNewParamsBodyObject struct {
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html,required"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]PDFNewParamsBodyObjectAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]PDFNewParamsBodyObjectAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]PDFNewParamsBodyObjectAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[PDFNewParamsBodyObjectAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]PDFNewParamsBodyObjectCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                         `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[PDFNewParamsBodyObjectGotoOptions] `json:"gotoOptions"`
	// Check [options](https://pptr.dev/api/puppeteer.pdfoptions).
	PDFOptions param.Field[PDFNewParamsBodyObjectPDFOptions] `json:"pdfOptions"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]PDFNewParamsBodyObjectRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                          `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                       `json:"setJavaScriptEnabled"`
	UserAgent            param.Field[string]                                     `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[PDFNewParamsBodyObjectViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[PDFNewParamsBodyObjectWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r PDFNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PDFNewParamsBodyObject) implementsPDFNewParamsBodyUnion() {}

type PDFNewParamsBodyObjectAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r PDFNewParamsBodyObjectAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PDFNewParamsBodyObjectAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r PDFNewParamsBodyObjectAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PDFNewParamsBodyObjectAllowResourceType string

const (
	PDFNewParamsBodyObjectAllowResourceTypeDocument           PDFNewParamsBodyObjectAllowResourceType = "document"
	PDFNewParamsBodyObjectAllowResourceTypeStylesheet         PDFNewParamsBodyObjectAllowResourceType = "stylesheet"
	PDFNewParamsBodyObjectAllowResourceTypeImage              PDFNewParamsBodyObjectAllowResourceType = "image"
	PDFNewParamsBodyObjectAllowResourceTypeMedia              PDFNewParamsBodyObjectAllowResourceType = "media"
	PDFNewParamsBodyObjectAllowResourceTypeFont               PDFNewParamsBodyObjectAllowResourceType = "font"
	PDFNewParamsBodyObjectAllowResourceTypeScript             PDFNewParamsBodyObjectAllowResourceType = "script"
	PDFNewParamsBodyObjectAllowResourceTypeTexttrack          PDFNewParamsBodyObjectAllowResourceType = "texttrack"
	PDFNewParamsBodyObjectAllowResourceTypeXHR                PDFNewParamsBodyObjectAllowResourceType = "xhr"
	PDFNewParamsBodyObjectAllowResourceTypeFetch              PDFNewParamsBodyObjectAllowResourceType = "fetch"
	PDFNewParamsBodyObjectAllowResourceTypePrefetch           PDFNewParamsBodyObjectAllowResourceType = "prefetch"
	PDFNewParamsBodyObjectAllowResourceTypeEventsource        PDFNewParamsBodyObjectAllowResourceType = "eventsource"
	PDFNewParamsBodyObjectAllowResourceTypeWebsocket          PDFNewParamsBodyObjectAllowResourceType = "websocket"
	PDFNewParamsBodyObjectAllowResourceTypeManifest           PDFNewParamsBodyObjectAllowResourceType = "manifest"
	PDFNewParamsBodyObjectAllowResourceTypeSignedexchange     PDFNewParamsBodyObjectAllowResourceType = "signedexchange"
	PDFNewParamsBodyObjectAllowResourceTypePing               PDFNewParamsBodyObjectAllowResourceType = "ping"
	PDFNewParamsBodyObjectAllowResourceTypeCspviolationreport PDFNewParamsBodyObjectAllowResourceType = "cspviolationreport"
	PDFNewParamsBodyObjectAllowResourceTypePreflight          PDFNewParamsBodyObjectAllowResourceType = "preflight"
	PDFNewParamsBodyObjectAllowResourceTypeOther              PDFNewParamsBodyObjectAllowResourceType = "other"
)

func (r PDFNewParamsBodyObjectAllowResourceType) IsKnown() bool {
	switch r {
	case PDFNewParamsBodyObjectAllowResourceTypeDocument, PDFNewParamsBodyObjectAllowResourceTypeStylesheet, PDFNewParamsBodyObjectAllowResourceTypeImage, PDFNewParamsBodyObjectAllowResourceTypeMedia, PDFNewParamsBodyObjectAllowResourceTypeFont, PDFNewParamsBodyObjectAllowResourceTypeScript, PDFNewParamsBodyObjectAllowResourceTypeTexttrack, PDFNewParamsBodyObjectAllowResourceTypeXHR, PDFNewParamsBodyObjectAllowResourceTypeFetch, PDFNewParamsBodyObjectAllowResourceTypePrefetch, PDFNewParamsBodyObjectAllowResourceTypeEventsource, PDFNewParamsBodyObjectAllowResourceTypeWebsocket, PDFNewParamsBodyObjectAllowResourceTypeManifest, PDFNewParamsBodyObjectAllowResourceTypeSignedexchange, PDFNewParamsBodyObjectAllowResourceTypePing, PDFNewParamsBodyObjectAllowResourceTypeCspviolationreport, PDFNewParamsBodyObjectAllowResourceTypePreflight, PDFNewParamsBodyObjectAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type PDFNewParamsBodyObjectAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r PDFNewParamsBodyObjectAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PDFNewParamsBodyObjectCookie struct {
	Name         param.Field[string]                                    `json:"name,required"`
	Value        param.Field[string]                                    `json:"value,required"`
	Domain       param.Field[string]                                    `json:"domain"`
	Expires      param.Field[float64]                                   `json:"expires"`
	HTTPOnly     param.Field[bool]                                      `json:"httpOnly"`
	PartitionKey param.Field[string]                                    `json:"partitionKey"`
	Path         param.Field[string]                                    `json:"path"`
	Priority     param.Field[PDFNewParamsBodyObjectCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                      `json:"sameParty"`
	SameSite     param.Field[PDFNewParamsBodyObjectCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                      `json:"secure"`
	SourcePort   param.Field[float64]                                   `json:"sourcePort"`
	SourceScheme param.Field[PDFNewParamsBodyObjectCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                    `json:"url"`
}

func (r PDFNewParamsBodyObjectCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PDFNewParamsBodyObjectCookiesPriority string

const (
	PDFNewParamsBodyObjectCookiesPriorityLow    PDFNewParamsBodyObjectCookiesPriority = "Low"
	PDFNewParamsBodyObjectCookiesPriorityMedium PDFNewParamsBodyObjectCookiesPriority = "Medium"
	PDFNewParamsBodyObjectCookiesPriorityHigh   PDFNewParamsBodyObjectCookiesPriority = "High"
)

func (r PDFNewParamsBodyObjectCookiesPriority) IsKnown() bool {
	switch r {
	case PDFNewParamsBodyObjectCookiesPriorityLow, PDFNewParamsBodyObjectCookiesPriorityMedium, PDFNewParamsBodyObjectCookiesPriorityHigh:
		return true
	}
	return false
}

type PDFNewParamsBodyObjectCookiesSameSite string

const (
	PDFNewParamsBodyObjectCookiesSameSiteStrict PDFNewParamsBodyObjectCookiesSameSite = "Strict"
	PDFNewParamsBodyObjectCookiesSameSiteLax    PDFNewParamsBodyObjectCookiesSameSite = "Lax"
	PDFNewParamsBodyObjectCookiesSameSiteNone   PDFNewParamsBodyObjectCookiesSameSite = "None"
)

func (r PDFNewParamsBodyObjectCookiesSameSite) IsKnown() bool {
	switch r {
	case PDFNewParamsBodyObjectCookiesSameSiteStrict, PDFNewParamsBodyObjectCookiesSameSiteLax, PDFNewParamsBodyObjectCookiesSameSiteNone:
		return true
	}
	return false
}

type PDFNewParamsBodyObjectCookiesSourceScheme string

const (
	PDFNewParamsBodyObjectCookiesSourceSchemeUnset     PDFNewParamsBodyObjectCookiesSourceScheme = "Unset"
	PDFNewParamsBodyObjectCookiesSourceSchemeNonSecure PDFNewParamsBodyObjectCookiesSourceScheme = "NonSecure"
	PDFNewParamsBodyObjectCookiesSourceSchemeSecure    PDFNewParamsBodyObjectCookiesSourceScheme = "Secure"
)

func (r PDFNewParamsBodyObjectCookiesSourceScheme) IsKnown() bool {
	switch r {
	case PDFNewParamsBodyObjectCookiesSourceSchemeUnset, PDFNewParamsBodyObjectCookiesSourceSchemeNonSecure, PDFNewParamsBodyObjectCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type PDFNewParamsBodyObjectGotoOptions struct {
	Referer        param.Field[string]                                          `json:"referer"`
	ReferrerPolicy param.Field[string]                                          `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                         `json:"timeout"`
	WaitUntil      param.Field[PDFNewParamsBodyObjectGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r PDFNewParamsBodyObjectGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.PDFNewParamsBodyObjectGotoOptionsWaitUntilString],
// [browser_rendering.PDFNewParamsBodyObjectGotoOptionsWaitUntilArray].
type PDFNewParamsBodyObjectGotoOptionsWaitUntilUnion interface {
	implementsPDFNewParamsBodyObjectGotoOptionsWaitUntilUnion()
}

type PDFNewParamsBodyObjectGotoOptionsWaitUntilString string

const (
	PDFNewParamsBodyObjectGotoOptionsWaitUntilStringLoad             PDFNewParamsBodyObjectGotoOptionsWaitUntilString = "load"
	PDFNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded PDFNewParamsBodyObjectGotoOptionsWaitUntilString = "domcontentloaded"
	PDFNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0     PDFNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle0"
	PDFNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2     PDFNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle2"
)

func (r PDFNewParamsBodyObjectGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case PDFNewParamsBodyObjectGotoOptionsWaitUntilStringLoad, PDFNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded, PDFNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0, PDFNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r PDFNewParamsBodyObjectGotoOptionsWaitUntilString) implementsPDFNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type PDFNewParamsBodyObjectGotoOptionsWaitUntilArray []PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItem

func (r PDFNewParamsBodyObjectGotoOptionsWaitUntilArray) implementsPDFNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItem string

const (
	PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad             PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "load"
	PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0     PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle0"
	PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2     PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad, PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded, PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0, PDFNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.pdfoptions).
type PDFNewParamsBodyObjectPDFOptions struct {
	// Whether to show the header and footer.
	DisplayHeaderFooter param.Field[bool] `json:"displayHeaderFooter"`
	// HTML template for the print footer.
	FooterTemplate param.Field[string] `json:"footerTemplate"`
	// Paper format. Takes priority over width and height if set.
	Format param.Field[PDFNewParamsBodyObjectPDFOptionsFormat] `json:"format"`
	// HTML template for the print header.
	HeaderTemplate param.Field[string] `json:"headerTemplate"`
	// Sets the height of paper. Can be a number or string with unit.
	Height param.Field[PDFNewParamsBodyObjectPDFOptionsHeightUnion] `json:"height"`
	// Whether to print in landscape orientation.
	Landscape param.Field[bool] `json:"landscape"`
	// Set the PDF margins. Useful when setting header and footer.
	Margin param.Field[PDFNewParamsBodyObjectPDFOptionsMargin] `json:"margin"`
	// Hides default white background and allows generating pdfs with transparency.
	OmitBackground param.Field[bool] `json:"omitBackground"`
	// Generate document outline.
	Outline param.Field[bool] `json:"outline"`
	// Paper ranges to print, e.g. '1-5, 8, 11-13'.
	PageRanges param.Field[string] `json:"pageRanges"`
	// Give CSS @page size priority over other size declarations.
	PreferCSSPageSize param.Field[bool] `json:"preferCSSPageSize"`
	// Set to true to print background graphics.
	PrintBackground param.Field[bool] `json:"printBackground"`
	// Scales the rendering of the web page. Amount must be between 0.1 and 2.
	Scale param.Field[float64] `json:"scale"`
	// Generate tagged (accessible) PDF.
	Tagged param.Field[bool] `json:"tagged"`
	// Timeout in milliseconds.
	Timeout param.Field[float64] `json:"timeout"`
	// Sets the width of paper. Can be a number or string with unit.
	Width param.Field[PDFNewParamsBodyObjectPDFOptionsWidthUnion] `json:"width"`
}

func (r PDFNewParamsBodyObjectPDFOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Paper format. Takes priority over width and height if set.
type PDFNewParamsBodyObjectPDFOptionsFormat string

const (
	PDFNewParamsBodyObjectPDFOptionsFormatLetter  PDFNewParamsBodyObjectPDFOptionsFormat = "letter"
	PDFNewParamsBodyObjectPDFOptionsFormatLegal   PDFNewParamsBodyObjectPDFOptionsFormat = "legal"
	PDFNewParamsBodyObjectPDFOptionsFormatTabloid PDFNewParamsBodyObjectPDFOptionsFormat = "tabloid"
	PDFNewParamsBodyObjectPDFOptionsFormatLedger  PDFNewParamsBodyObjectPDFOptionsFormat = "ledger"
	PDFNewParamsBodyObjectPDFOptionsFormatA0      PDFNewParamsBodyObjectPDFOptionsFormat = "a0"
	PDFNewParamsBodyObjectPDFOptionsFormatA1      PDFNewParamsBodyObjectPDFOptionsFormat = "a1"
	PDFNewParamsBodyObjectPDFOptionsFormatA2      PDFNewParamsBodyObjectPDFOptionsFormat = "a2"
	PDFNewParamsBodyObjectPDFOptionsFormatA3      PDFNewParamsBodyObjectPDFOptionsFormat = "a3"
	PDFNewParamsBodyObjectPDFOptionsFormatA4      PDFNewParamsBodyObjectPDFOptionsFormat = "a4"
	PDFNewParamsBodyObjectPDFOptionsFormatA5      PDFNewParamsBodyObjectPDFOptionsFormat = "a5"
	PDFNewParamsBodyObjectPDFOptionsFormatA6      PDFNewParamsBodyObjectPDFOptionsFormat = "a6"
)

func (r PDFNewParamsBodyObjectPDFOptionsFormat) IsKnown() bool {
	switch r {
	case PDFNewParamsBodyObjectPDFOptionsFormatLetter, PDFNewParamsBodyObjectPDFOptionsFormatLegal, PDFNewParamsBodyObjectPDFOptionsFormatTabloid, PDFNewParamsBodyObjectPDFOptionsFormatLedger, PDFNewParamsBodyObjectPDFOptionsFormatA0, PDFNewParamsBodyObjectPDFOptionsFormatA1, PDFNewParamsBodyObjectPDFOptionsFormatA2, PDFNewParamsBodyObjectPDFOptionsFormatA3, PDFNewParamsBodyObjectPDFOptionsFormatA4, PDFNewParamsBodyObjectPDFOptionsFormatA5, PDFNewParamsBodyObjectPDFOptionsFormatA6:
		return true
	}
	return false
}

// Sets the height of paper. Can be a number or string with unit.
//
// Satisfied by [shared.UnionString], [shared.UnionFloat].
type PDFNewParamsBodyObjectPDFOptionsHeightUnion interface {
	ImplementsPDFNewParamsBodyObjectPDFOptionsHeightUnion()
}

// Set the PDF margins. Useful when setting header and footer.
type PDFNewParamsBodyObjectPDFOptionsMargin struct {
	Bottom param.Field[PDFNewParamsBodyObjectPDFOptionsMarginBottomUnion] `json:"bottom"`
	Left   param.Field[PDFNewParamsBodyObjectPDFOptionsMarginLeftUnion]   `json:"left"`
	Right  param.Field[PDFNewParamsBodyObjectPDFOptionsMarginRightUnion]  `json:"right"`
	Top    param.Field[PDFNewParamsBodyObjectPDFOptionsMarginTopUnion]    `json:"top"`
}

func (r PDFNewParamsBodyObjectPDFOptionsMargin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type PDFNewParamsBodyObjectPDFOptionsMarginBottomUnion interface {
	ImplementsPDFNewParamsBodyObjectPDFOptionsMarginBottomUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type PDFNewParamsBodyObjectPDFOptionsMarginLeftUnion interface {
	ImplementsPDFNewParamsBodyObjectPDFOptionsMarginLeftUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type PDFNewParamsBodyObjectPDFOptionsMarginRightUnion interface {
	ImplementsPDFNewParamsBodyObjectPDFOptionsMarginRightUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type PDFNewParamsBodyObjectPDFOptionsMarginTopUnion interface {
	ImplementsPDFNewParamsBodyObjectPDFOptionsMarginTopUnion()
}

// Sets the width of paper. Can be a number or string with unit.
//
// Satisfied by [shared.UnionString], [shared.UnionFloat].
type PDFNewParamsBodyObjectPDFOptionsWidthUnion interface {
	ImplementsPDFNewParamsBodyObjectPDFOptionsWidthUnion()
}

type PDFNewParamsBodyObjectRejectResourceType string

const (
	PDFNewParamsBodyObjectRejectResourceTypeDocument           PDFNewParamsBodyObjectRejectResourceType = "document"
	PDFNewParamsBodyObjectRejectResourceTypeStylesheet         PDFNewParamsBodyObjectRejectResourceType = "stylesheet"
	PDFNewParamsBodyObjectRejectResourceTypeImage              PDFNewParamsBodyObjectRejectResourceType = "image"
	PDFNewParamsBodyObjectRejectResourceTypeMedia              PDFNewParamsBodyObjectRejectResourceType = "media"
	PDFNewParamsBodyObjectRejectResourceTypeFont               PDFNewParamsBodyObjectRejectResourceType = "font"
	PDFNewParamsBodyObjectRejectResourceTypeScript             PDFNewParamsBodyObjectRejectResourceType = "script"
	PDFNewParamsBodyObjectRejectResourceTypeTexttrack          PDFNewParamsBodyObjectRejectResourceType = "texttrack"
	PDFNewParamsBodyObjectRejectResourceTypeXHR                PDFNewParamsBodyObjectRejectResourceType = "xhr"
	PDFNewParamsBodyObjectRejectResourceTypeFetch              PDFNewParamsBodyObjectRejectResourceType = "fetch"
	PDFNewParamsBodyObjectRejectResourceTypePrefetch           PDFNewParamsBodyObjectRejectResourceType = "prefetch"
	PDFNewParamsBodyObjectRejectResourceTypeEventsource        PDFNewParamsBodyObjectRejectResourceType = "eventsource"
	PDFNewParamsBodyObjectRejectResourceTypeWebsocket          PDFNewParamsBodyObjectRejectResourceType = "websocket"
	PDFNewParamsBodyObjectRejectResourceTypeManifest           PDFNewParamsBodyObjectRejectResourceType = "manifest"
	PDFNewParamsBodyObjectRejectResourceTypeSignedexchange     PDFNewParamsBodyObjectRejectResourceType = "signedexchange"
	PDFNewParamsBodyObjectRejectResourceTypePing               PDFNewParamsBodyObjectRejectResourceType = "ping"
	PDFNewParamsBodyObjectRejectResourceTypeCspviolationreport PDFNewParamsBodyObjectRejectResourceType = "cspviolationreport"
	PDFNewParamsBodyObjectRejectResourceTypePreflight          PDFNewParamsBodyObjectRejectResourceType = "preflight"
	PDFNewParamsBodyObjectRejectResourceTypeOther              PDFNewParamsBodyObjectRejectResourceType = "other"
)

func (r PDFNewParamsBodyObjectRejectResourceType) IsKnown() bool {
	switch r {
	case PDFNewParamsBodyObjectRejectResourceTypeDocument, PDFNewParamsBodyObjectRejectResourceTypeStylesheet, PDFNewParamsBodyObjectRejectResourceTypeImage, PDFNewParamsBodyObjectRejectResourceTypeMedia, PDFNewParamsBodyObjectRejectResourceTypeFont, PDFNewParamsBodyObjectRejectResourceTypeScript, PDFNewParamsBodyObjectRejectResourceTypeTexttrack, PDFNewParamsBodyObjectRejectResourceTypeXHR, PDFNewParamsBodyObjectRejectResourceTypeFetch, PDFNewParamsBodyObjectRejectResourceTypePrefetch, PDFNewParamsBodyObjectRejectResourceTypeEventsource, PDFNewParamsBodyObjectRejectResourceTypeWebsocket, PDFNewParamsBodyObjectRejectResourceTypeManifest, PDFNewParamsBodyObjectRejectResourceTypeSignedexchange, PDFNewParamsBodyObjectRejectResourceTypePing, PDFNewParamsBodyObjectRejectResourceTypeCspviolationreport, PDFNewParamsBodyObjectRejectResourceTypePreflight, PDFNewParamsBodyObjectRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type PDFNewParamsBodyObjectViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r PDFNewParamsBodyObjectViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type PDFNewParamsBodyObjectWaitForSelector struct {
	Selector param.Field[string]                                       `json:"selector,required"`
	Hidden   param.Field[PDFNewParamsBodyObjectWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                      `json:"timeout"`
	Visible  param.Field[PDFNewParamsBodyObjectWaitForSelectorVisible] `json:"visible"`
}

func (r PDFNewParamsBodyObjectWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PDFNewParamsBodyObjectWaitForSelectorHidden bool

const (
	PDFNewParamsBodyObjectWaitForSelectorHiddenTrue PDFNewParamsBodyObjectWaitForSelectorHidden = true
)

func (r PDFNewParamsBodyObjectWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case PDFNewParamsBodyObjectWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type PDFNewParamsBodyObjectWaitForSelectorVisible bool

const (
	PDFNewParamsBodyObjectWaitForSelectorVisibleTrue PDFNewParamsBodyObjectWaitForSelectorVisible = true
)

func (r PDFNewParamsBodyObjectWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case PDFNewParamsBodyObjectWaitForSelectorVisibleTrue:
		return true
	}
	return false
}
