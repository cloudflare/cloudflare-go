// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// BrowserRenderingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrowserRenderingService] method instead.
type BrowserRenderingService struct {
	Options []option.RequestOption
}

// NewBrowserRenderingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBrowserRenderingService(opts ...option.RequestOption) (r *BrowserRenderingService) {
	r = &BrowserRenderingService{}
	r.Options = opts
	return
}

// Fetches rendered HTML content from provided URL or HTML. Check available options
// like `goToOptions` and `waitFor*` to control page load behaviour.
func (r *BrowserRenderingService) Content(ctx context.Context, accountID string, params BrowserRenderingContentParams, opts ...option.RequestOption) (res *string, err error) {
	var env BrowserRenderingContentResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/content", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches rendered PDF from provided URL or HTML. Check available options like
// `goToOptions` and `waitFor*` to control page load behaviour.
func (r *BrowserRenderingService) PDF(ctx context.Context, accountID string, params BrowserRenderingPDFParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/pdf", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get meta attributes like height, width, text and others of selected elements.
func (r *BrowserRenderingService) Scrape(ctx context.Context, accountID string, params BrowserRenderingScrapeParams, opts ...option.RequestOption) (res *[]BrowserRenderingScrapeResponse, err error) {
	var env BrowserRenderingScrapeResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/scrape", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Takes a screenshot of a webpage from provided URL or HTML. Control page loading
// with `goToOptions` and `waitFor*` options. Customize screenshots with
// `viewport`, `fullPage`, `clip` and others.
func (r *BrowserRenderingService) Screenshot(ctx context.Context, accountID string, params BrowserRenderingScreenshotParams, opts ...option.RequestOption) (res *BrowserRenderingScreenshotResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/screenshot", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns the page's HTML content and screenshot. Control page loading with
// `goToOptions` and `waitFor*` options. Customize screenshots with `viewport`,
// `fullPage`, `clip` and others.
func (r *BrowserRenderingService) Snapshot(ctx context.Context, accountID string, params BrowserRenderingSnapshotParams, opts ...option.RequestOption) (res *BrowserRenderingSnapshotResponse, err error) {
	var env BrowserRenderingSnapshotResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/snapshot", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BrowserRenderingScrapeResponse struct {
	Result BrowserRenderingScrapeResponseResult `json:"result,required"`
	// Selector
	Selector string                             `json:"selector,required"`
	JSON     browserRenderingScrapeResponseJSON `json:"-"`
}

// browserRenderingScrapeResponseJSON contains the JSON metadata for the struct
// [BrowserRenderingScrapeResponse]
type browserRenderingScrapeResponseJSON struct {
	Result      apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserRenderingScrapeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserRenderingScrapeResponseJSON) RawJSON() string {
	return r.raw
}

type BrowserRenderingScrapeResponseResult struct {
	Attributes []BrowserRenderingScrapeResponseResultAttribute `json:"attributes,required"`
	// Element height
	Height float64 `json:"height,required"`
	// Html content
	HTML string `json:"html,required"`
	// Element left
	Left float64 `json:"left,required"`
	// Text content
	Text string `json:"text,required"`
	// Element top
	Top float64 `json:"top,required"`
	// Element width
	Width float64                                  `json:"width,required"`
	JSON  browserRenderingScrapeResponseResultJSON `json:"-"`
}

// browserRenderingScrapeResponseResultJSON contains the JSON metadata for the
// struct [BrowserRenderingScrapeResponseResult]
type browserRenderingScrapeResponseResultJSON struct {
	Attributes  apijson.Field
	Height      apijson.Field
	HTML        apijson.Field
	Left        apijson.Field
	Text        apijson.Field
	Top         apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserRenderingScrapeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserRenderingScrapeResponseResultJSON) RawJSON() string {
	return r.raw
}

type BrowserRenderingScrapeResponseResultAttribute struct {
	// Attribute name
	Name string `json:"name,required"`
	// Attribute value
	Value string                                            `json:"value,required"`
	JSON  browserRenderingScrapeResponseResultAttributeJSON `json:"-"`
}

// browserRenderingScrapeResponseResultAttributeJSON contains the JSON metadata for
// the struct [BrowserRenderingScrapeResponseResultAttribute]
type browserRenderingScrapeResponseResultAttributeJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserRenderingScrapeResponseResultAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserRenderingScrapeResponseResultAttributeJSON) RawJSON() string {
	return r.raw
}

type BrowserRenderingScreenshotResponse struct {
	// Response status
	Status bool                                      `json:"status,required"`
	Errors []BrowserRenderingScreenshotResponseError `json:"errors"`
	JSON   browserRenderingScreenshotResponseJSON    `json:"-"`
}

// browserRenderingScreenshotResponseJSON contains the JSON metadata for the struct
// [BrowserRenderingScreenshotResponse]
type browserRenderingScreenshotResponseJSON struct {
	Status      apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserRenderingScreenshotResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserRenderingScreenshotResponseJSON) RawJSON() string {
	return r.raw
}

type BrowserRenderingScreenshotResponseError struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                      `json:"message,required"`
	JSON    browserRenderingScreenshotResponseErrorJSON `json:"-"`
}

// browserRenderingScreenshotResponseErrorJSON contains the JSON metadata for the
// struct [BrowserRenderingScreenshotResponseError]
type browserRenderingScreenshotResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserRenderingScreenshotResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserRenderingScreenshotResponseErrorJSON) RawJSON() string {
	return r.raw
}

type BrowserRenderingSnapshotResponse struct {
	// HTML content
	Content string `json:"content,required"`
	// Base64 encoded image
	Screenshot string                               `json:"screenshot,required"`
	JSON       browserRenderingSnapshotResponseJSON `json:"-"`
}

// browserRenderingSnapshotResponseJSON contains the JSON metadata for the struct
// [BrowserRenderingSnapshotResponse]
type browserRenderingSnapshotResponseJSON struct {
	Content     apijson.Field
	Screenshot  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserRenderingSnapshotResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserRenderingSnapshotResponseJSON) RawJSON() string {
	return r.raw
}

type BrowserRenderingContentParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]BrowserRenderingContentParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]BrowserRenderingContentParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]BrowserRenderingContentParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[BrowserRenderingContentParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]BrowserRenderingContentParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[BrowserRenderingContentParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]BrowserRenderingContentParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                 `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                              `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[BrowserRenderingContentParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[BrowserRenderingContentParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r BrowserRenderingContentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [BrowserRenderingContentParams]'s query parameters as
// `url.Values`.
func (r BrowserRenderingContentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BrowserRenderingContentParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r BrowserRenderingContentParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingContentParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r BrowserRenderingContentParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingContentParamsAllowResourceType string

const (
	BrowserRenderingContentParamsAllowResourceTypeDocument           BrowserRenderingContentParamsAllowResourceType = "document"
	BrowserRenderingContentParamsAllowResourceTypeStylesheet         BrowserRenderingContentParamsAllowResourceType = "stylesheet"
	BrowserRenderingContentParamsAllowResourceTypeImage              BrowserRenderingContentParamsAllowResourceType = "image"
	BrowserRenderingContentParamsAllowResourceTypeMedia              BrowserRenderingContentParamsAllowResourceType = "media"
	BrowserRenderingContentParamsAllowResourceTypeFont               BrowserRenderingContentParamsAllowResourceType = "font"
	BrowserRenderingContentParamsAllowResourceTypeScript             BrowserRenderingContentParamsAllowResourceType = "script"
	BrowserRenderingContentParamsAllowResourceTypeTexttrack          BrowserRenderingContentParamsAllowResourceType = "texttrack"
	BrowserRenderingContentParamsAllowResourceTypeXHR                BrowserRenderingContentParamsAllowResourceType = "xhr"
	BrowserRenderingContentParamsAllowResourceTypeFetch              BrowserRenderingContentParamsAllowResourceType = "fetch"
	BrowserRenderingContentParamsAllowResourceTypePrefetch           BrowserRenderingContentParamsAllowResourceType = "prefetch"
	BrowserRenderingContentParamsAllowResourceTypeEventsource        BrowserRenderingContentParamsAllowResourceType = "eventsource"
	BrowserRenderingContentParamsAllowResourceTypeWebsocket          BrowserRenderingContentParamsAllowResourceType = "websocket"
	BrowserRenderingContentParamsAllowResourceTypeManifest           BrowserRenderingContentParamsAllowResourceType = "manifest"
	BrowserRenderingContentParamsAllowResourceTypeSignedexchange     BrowserRenderingContentParamsAllowResourceType = "signedexchange"
	BrowserRenderingContentParamsAllowResourceTypePing               BrowserRenderingContentParamsAllowResourceType = "ping"
	BrowserRenderingContentParamsAllowResourceTypeCspviolationreport BrowserRenderingContentParamsAllowResourceType = "cspviolationreport"
	BrowserRenderingContentParamsAllowResourceTypePreflight          BrowserRenderingContentParamsAllowResourceType = "preflight"
	BrowserRenderingContentParamsAllowResourceTypeOther              BrowserRenderingContentParamsAllowResourceType = "other"
)

func (r BrowserRenderingContentParamsAllowResourceType) IsKnown() bool {
	switch r {
	case BrowserRenderingContentParamsAllowResourceTypeDocument, BrowserRenderingContentParamsAllowResourceTypeStylesheet, BrowserRenderingContentParamsAllowResourceTypeImage, BrowserRenderingContentParamsAllowResourceTypeMedia, BrowserRenderingContentParamsAllowResourceTypeFont, BrowserRenderingContentParamsAllowResourceTypeScript, BrowserRenderingContentParamsAllowResourceTypeTexttrack, BrowserRenderingContentParamsAllowResourceTypeXHR, BrowserRenderingContentParamsAllowResourceTypeFetch, BrowserRenderingContentParamsAllowResourceTypePrefetch, BrowserRenderingContentParamsAllowResourceTypeEventsource, BrowserRenderingContentParamsAllowResourceTypeWebsocket, BrowserRenderingContentParamsAllowResourceTypeManifest, BrowserRenderingContentParamsAllowResourceTypeSignedexchange, BrowserRenderingContentParamsAllowResourceTypePing, BrowserRenderingContentParamsAllowResourceTypeCspviolationreport, BrowserRenderingContentParamsAllowResourceTypePreflight, BrowserRenderingContentParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type BrowserRenderingContentParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r BrowserRenderingContentParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingContentParamsCookie struct {
	Name         param.Field[string]                                           `json:"name,required"`
	Value        param.Field[string]                                           `json:"value,required"`
	Domain       param.Field[string]                                           `json:"domain"`
	Expires      param.Field[float64]                                          `json:"expires"`
	HTTPOnly     param.Field[bool]                                             `json:"httpOnly"`
	PartitionKey param.Field[string]                                           `json:"partitionKey"`
	Path         param.Field[string]                                           `json:"path"`
	Priority     param.Field[BrowserRenderingContentParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                             `json:"sameParty"`
	SameSite     param.Field[BrowserRenderingContentParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                             `json:"secure"`
	SourcePort   param.Field[float64]                                          `json:"sourcePort"`
	SourceScheme param.Field[BrowserRenderingContentParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                           `json:"url"`
}

func (r BrowserRenderingContentParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingContentParamsCookiesPriority string

const (
	BrowserRenderingContentParamsCookiesPriorityLow    BrowserRenderingContentParamsCookiesPriority = "Low"
	BrowserRenderingContentParamsCookiesPriorityMedium BrowserRenderingContentParamsCookiesPriority = "Medium"
	BrowserRenderingContentParamsCookiesPriorityHigh   BrowserRenderingContentParamsCookiesPriority = "High"
)

func (r BrowserRenderingContentParamsCookiesPriority) IsKnown() bool {
	switch r {
	case BrowserRenderingContentParamsCookiesPriorityLow, BrowserRenderingContentParamsCookiesPriorityMedium, BrowserRenderingContentParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type BrowserRenderingContentParamsCookiesSameSite string

const (
	BrowserRenderingContentParamsCookiesSameSiteStrict BrowserRenderingContentParamsCookiesSameSite = "Strict"
	BrowserRenderingContentParamsCookiesSameSiteLax    BrowserRenderingContentParamsCookiesSameSite = "Lax"
	BrowserRenderingContentParamsCookiesSameSiteNone   BrowserRenderingContentParamsCookiesSameSite = "None"
)

func (r BrowserRenderingContentParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case BrowserRenderingContentParamsCookiesSameSiteStrict, BrowserRenderingContentParamsCookiesSameSiteLax, BrowserRenderingContentParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type BrowserRenderingContentParamsCookiesSourceScheme string

const (
	BrowserRenderingContentParamsCookiesSourceSchemeUnset     BrowserRenderingContentParamsCookiesSourceScheme = "Unset"
	BrowserRenderingContentParamsCookiesSourceSchemeNonSecure BrowserRenderingContentParamsCookiesSourceScheme = "NonSecure"
	BrowserRenderingContentParamsCookiesSourceSchemeSecure    BrowserRenderingContentParamsCookiesSourceScheme = "Secure"
)

func (r BrowserRenderingContentParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case BrowserRenderingContentParamsCookiesSourceSchemeUnset, BrowserRenderingContentParamsCookiesSourceSchemeNonSecure, BrowserRenderingContentParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type BrowserRenderingContentParamsGotoOptions struct {
	Referer        param.Field[string]                                                 `json:"referer"`
	ReferrerPolicy param.Field[string]                                                 `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                `json:"timeout"`
	WaitUntil      param.Field[BrowserRenderingContentParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r BrowserRenderingContentParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.BrowserRenderingContentParamsGotoOptionsWaitUntilString],
// [browser_rendering.BrowserRenderingContentParamsGotoOptionsWaitUntilArray].
type BrowserRenderingContentParamsGotoOptionsWaitUntilUnion interface {
	implementsBrowserRenderingContentParamsGotoOptionsWaitUntilUnion()
}

type BrowserRenderingContentParamsGotoOptionsWaitUntilString string

const (
	BrowserRenderingContentParamsGotoOptionsWaitUntilStringLoad             BrowserRenderingContentParamsGotoOptionsWaitUntilString = "load"
	BrowserRenderingContentParamsGotoOptionsWaitUntilStringDomcontentloaded BrowserRenderingContentParamsGotoOptionsWaitUntilString = "domcontentloaded"
	BrowserRenderingContentParamsGotoOptionsWaitUntilStringNetworkidle0     BrowserRenderingContentParamsGotoOptionsWaitUntilString = "networkidle0"
	BrowserRenderingContentParamsGotoOptionsWaitUntilStringNetworkidle2     BrowserRenderingContentParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r BrowserRenderingContentParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case BrowserRenderingContentParamsGotoOptionsWaitUntilStringLoad, BrowserRenderingContentParamsGotoOptionsWaitUntilStringDomcontentloaded, BrowserRenderingContentParamsGotoOptionsWaitUntilStringNetworkidle0, BrowserRenderingContentParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r BrowserRenderingContentParamsGotoOptionsWaitUntilString) implementsBrowserRenderingContentParamsGotoOptionsWaitUntilUnion() {
}

type BrowserRenderingContentParamsGotoOptionsWaitUntilArray []BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItem

func (r BrowserRenderingContentParamsGotoOptionsWaitUntilArray) implementsBrowserRenderingContentParamsGotoOptionsWaitUntilUnion() {
}

type BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItem string

const (
	BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItemLoad             BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItem = "load"
	BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItemDomcontentloaded BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItemNetworkidle0     BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItemNetworkidle2     BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItemLoad, BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItemNetworkidle0, BrowserRenderingContentParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type BrowserRenderingContentParamsRejectResourceType string

const (
	BrowserRenderingContentParamsRejectResourceTypeDocument           BrowserRenderingContentParamsRejectResourceType = "document"
	BrowserRenderingContentParamsRejectResourceTypeStylesheet         BrowserRenderingContentParamsRejectResourceType = "stylesheet"
	BrowserRenderingContentParamsRejectResourceTypeImage              BrowserRenderingContentParamsRejectResourceType = "image"
	BrowserRenderingContentParamsRejectResourceTypeMedia              BrowserRenderingContentParamsRejectResourceType = "media"
	BrowserRenderingContentParamsRejectResourceTypeFont               BrowserRenderingContentParamsRejectResourceType = "font"
	BrowserRenderingContentParamsRejectResourceTypeScript             BrowserRenderingContentParamsRejectResourceType = "script"
	BrowserRenderingContentParamsRejectResourceTypeTexttrack          BrowserRenderingContentParamsRejectResourceType = "texttrack"
	BrowserRenderingContentParamsRejectResourceTypeXHR                BrowserRenderingContentParamsRejectResourceType = "xhr"
	BrowserRenderingContentParamsRejectResourceTypeFetch              BrowserRenderingContentParamsRejectResourceType = "fetch"
	BrowserRenderingContentParamsRejectResourceTypePrefetch           BrowserRenderingContentParamsRejectResourceType = "prefetch"
	BrowserRenderingContentParamsRejectResourceTypeEventsource        BrowserRenderingContentParamsRejectResourceType = "eventsource"
	BrowserRenderingContentParamsRejectResourceTypeWebsocket          BrowserRenderingContentParamsRejectResourceType = "websocket"
	BrowserRenderingContentParamsRejectResourceTypeManifest           BrowserRenderingContentParamsRejectResourceType = "manifest"
	BrowserRenderingContentParamsRejectResourceTypeSignedexchange     BrowserRenderingContentParamsRejectResourceType = "signedexchange"
	BrowserRenderingContentParamsRejectResourceTypePing               BrowserRenderingContentParamsRejectResourceType = "ping"
	BrowserRenderingContentParamsRejectResourceTypeCspviolationreport BrowserRenderingContentParamsRejectResourceType = "cspviolationreport"
	BrowserRenderingContentParamsRejectResourceTypePreflight          BrowserRenderingContentParamsRejectResourceType = "preflight"
	BrowserRenderingContentParamsRejectResourceTypeOther              BrowserRenderingContentParamsRejectResourceType = "other"
)

func (r BrowserRenderingContentParamsRejectResourceType) IsKnown() bool {
	switch r {
	case BrowserRenderingContentParamsRejectResourceTypeDocument, BrowserRenderingContentParamsRejectResourceTypeStylesheet, BrowserRenderingContentParamsRejectResourceTypeImage, BrowserRenderingContentParamsRejectResourceTypeMedia, BrowserRenderingContentParamsRejectResourceTypeFont, BrowserRenderingContentParamsRejectResourceTypeScript, BrowserRenderingContentParamsRejectResourceTypeTexttrack, BrowserRenderingContentParamsRejectResourceTypeXHR, BrowserRenderingContentParamsRejectResourceTypeFetch, BrowserRenderingContentParamsRejectResourceTypePrefetch, BrowserRenderingContentParamsRejectResourceTypeEventsource, BrowserRenderingContentParamsRejectResourceTypeWebsocket, BrowserRenderingContentParamsRejectResourceTypeManifest, BrowserRenderingContentParamsRejectResourceTypeSignedexchange, BrowserRenderingContentParamsRejectResourceTypePing, BrowserRenderingContentParamsRejectResourceTypeCspviolationreport, BrowserRenderingContentParamsRejectResourceTypePreflight, BrowserRenderingContentParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type BrowserRenderingContentParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r BrowserRenderingContentParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type BrowserRenderingContentParamsWaitForSelector struct {
	Selector param.Field[string]                                              `json:"selector,required"`
	Hidden   param.Field[BrowserRenderingContentParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                             `json:"timeout"`
	Visible  param.Field[BrowserRenderingContentParamsWaitForSelectorVisible] `json:"visible"`
}

func (r BrowserRenderingContentParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingContentParamsWaitForSelectorHidden bool

const (
	BrowserRenderingContentParamsWaitForSelectorHiddenTrue BrowserRenderingContentParamsWaitForSelectorHidden = true
)

func (r BrowserRenderingContentParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case BrowserRenderingContentParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type BrowserRenderingContentParamsWaitForSelectorVisible bool

const (
	BrowserRenderingContentParamsWaitForSelectorVisibleTrue BrowserRenderingContentParamsWaitForSelectorVisible = true
)

func (r BrowserRenderingContentParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case BrowserRenderingContentParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type BrowserRenderingContentResponseEnvelope struct {
	// Response status
	Status bool                                            `json:"status,required"`
	Errors []BrowserRenderingContentResponseEnvelopeErrors `json:"errors"`
	// HTML content
	Result string                                      `json:"result"`
	JSON   browserRenderingContentResponseEnvelopeJSON `json:"-"`
}

// browserRenderingContentResponseEnvelopeJSON contains the JSON metadata for the
// struct [BrowserRenderingContentResponseEnvelope]
type browserRenderingContentResponseEnvelopeJSON struct {
	Status      apijson.Field
	Errors      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserRenderingContentResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserRenderingContentResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BrowserRenderingContentResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                            `json:"message,required"`
	JSON    browserRenderingContentResponseEnvelopeErrorsJSON `json:"-"`
}

// browserRenderingContentResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BrowserRenderingContentResponseEnvelopeErrors]
type browserRenderingContentResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserRenderingContentResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserRenderingContentResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BrowserRenderingPDFParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]BrowserRenderingPDFParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]BrowserRenderingPDFParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]BrowserRenderingPDFParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[BrowserRenderingPDFParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]BrowserRenderingPDFParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                            `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[BrowserRenderingPDFParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]BrowserRenderingPDFParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                             `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                          `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[BrowserRenderingPDFParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[BrowserRenderingPDFParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r BrowserRenderingPDFParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [BrowserRenderingPDFParams]'s query parameters as
// `url.Values`.
func (r BrowserRenderingPDFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BrowserRenderingPDFParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r BrowserRenderingPDFParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingPDFParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r BrowserRenderingPDFParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingPDFParamsAllowResourceType string

const (
	BrowserRenderingPDFParamsAllowResourceTypeDocument           BrowserRenderingPDFParamsAllowResourceType = "document"
	BrowserRenderingPDFParamsAllowResourceTypeStylesheet         BrowserRenderingPDFParamsAllowResourceType = "stylesheet"
	BrowserRenderingPDFParamsAllowResourceTypeImage              BrowserRenderingPDFParamsAllowResourceType = "image"
	BrowserRenderingPDFParamsAllowResourceTypeMedia              BrowserRenderingPDFParamsAllowResourceType = "media"
	BrowserRenderingPDFParamsAllowResourceTypeFont               BrowserRenderingPDFParamsAllowResourceType = "font"
	BrowserRenderingPDFParamsAllowResourceTypeScript             BrowserRenderingPDFParamsAllowResourceType = "script"
	BrowserRenderingPDFParamsAllowResourceTypeTexttrack          BrowserRenderingPDFParamsAllowResourceType = "texttrack"
	BrowserRenderingPDFParamsAllowResourceTypeXHR                BrowserRenderingPDFParamsAllowResourceType = "xhr"
	BrowserRenderingPDFParamsAllowResourceTypeFetch              BrowserRenderingPDFParamsAllowResourceType = "fetch"
	BrowserRenderingPDFParamsAllowResourceTypePrefetch           BrowserRenderingPDFParamsAllowResourceType = "prefetch"
	BrowserRenderingPDFParamsAllowResourceTypeEventsource        BrowserRenderingPDFParamsAllowResourceType = "eventsource"
	BrowserRenderingPDFParamsAllowResourceTypeWebsocket          BrowserRenderingPDFParamsAllowResourceType = "websocket"
	BrowserRenderingPDFParamsAllowResourceTypeManifest           BrowserRenderingPDFParamsAllowResourceType = "manifest"
	BrowserRenderingPDFParamsAllowResourceTypeSignedexchange     BrowserRenderingPDFParamsAllowResourceType = "signedexchange"
	BrowserRenderingPDFParamsAllowResourceTypePing               BrowserRenderingPDFParamsAllowResourceType = "ping"
	BrowserRenderingPDFParamsAllowResourceTypeCspviolationreport BrowserRenderingPDFParamsAllowResourceType = "cspviolationreport"
	BrowserRenderingPDFParamsAllowResourceTypePreflight          BrowserRenderingPDFParamsAllowResourceType = "preflight"
	BrowserRenderingPDFParamsAllowResourceTypeOther              BrowserRenderingPDFParamsAllowResourceType = "other"
)

func (r BrowserRenderingPDFParamsAllowResourceType) IsKnown() bool {
	switch r {
	case BrowserRenderingPDFParamsAllowResourceTypeDocument, BrowserRenderingPDFParamsAllowResourceTypeStylesheet, BrowserRenderingPDFParamsAllowResourceTypeImage, BrowserRenderingPDFParamsAllowResourceTypeMedia, BrowserRenderingPDFParamsAllowResourceTypeFont, BrowserRenderingPDFParamsAllowResourceTypeScript, BrowserRenderingPDFParamsAllowResourceTypeTexttrack, BrowserRenderingPDFParamsAllowResourceTypeXHR, BrowserRenderingPDFParamsAllowResourceTypeFetch, BrowserRenderingPDFParamsAllowResourceTypePrefetch, BrowserRenderingPDFParamsAllowResourceTypeEventsource, BrowserRenderingPDFParamsAllowResourceTypeWebsocket, BrowserRenderingPDFParamsAllowResourceTypeManifest, BrowserRenderingPDFParamsAllowResourceTypeSignedexchange, BrowserRenderingPDFParamsAllowResourceTypePing, BrowserRenderingPDFParamsAllowResourceTypeCspviolationreport, BrowserRenderingPDFParamsAllowResourceTypePreflight, BrowserRenderingPDFParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type BrowserRenderingPDFParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r BrowserRenderingPDFParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingPDFParamsCookie struct {
	Name         param.Field[string]                                       `json:"name,required"`
	Value        param.Field[string]                                       `json:"value,required"`
	Domain       param.Field[string]                                       `json:"domain"`
	Expires      param.Field[float64]                                      `json:"expires"`
	HTTPOnly     param.Field[bool]                                         `json:"httpOnly"`
	PartitionKey param.Field[string]                                       `json:"partitionKey"`
	Path         param.Field[string]                                       `json:"path"`
	Priority     param.Field[BrowserRenderingPDFParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                         `json:"sameParty"`
	SameSite     param.Field[BrowserRenderingPDFParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                         `json:"secure"`
	SourcePort   param.Field[float64]                                      `json:"sourcePort"`
	SourceScheme param.Field[BrowserRenderingPDFParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                       `json:"url"`
}

func (r BrowserRenderingPDFParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingPDFParamsCookiesPriority string

const (
	BrowserRenderingPDFParamsCookiesPriorityLow    BrowserRenderingPDFParamsCookiesPriority = "Low"
	BrowserRenderingPDFParamsCookiesPriorityMedium BrowserRenderingPDFParamsCookiesPriority = "Medium"
	BrowserRenderingPDFParamsCookiesPriorityHigh   BrowserRenderingPDFParamsCookiesPriority = "High"
)

func (r BrowserRenderingPDFParamsCookiesPriority) IsKnown() bool {
	switch r {
	case BrowserRenderingPDFParamsCookiesPriorityLow, BrowserRenderingPDFParamsCookiesPriorityMedium, BrowserRenderingPDFParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type BrowserRenderingPDFParamsCookiesSameSite string

const (
	BrowserRenderingPDFParamsCookiesSameSiteStrict BrowserRenderingPDFParamsCookiesSameSite = "Strict"
	BrowserRenderingPDFParamsCookiesSameSiteLax    BrowserRenderingPDFParamsCookiesSameSite = "Lax"
	BrowserRenderingPDFParamsCookiesSameSiteNone   BrowserRenderingPDFParamsCookiesSameSite = "None"
)

func (r BrowserRenderingPDFParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case BrowserRenderingPDFParamsCookiesSameSiteStrict, BrowserRenderingPDFParamsCookiesSameSiteLax, BrowserRenderingPDFParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type BrowserRenderingPDFParamsCookiesSourceScheme string

const (
	BrowserRenderingPDFParamsCookiesSourceSchemeUnset     BrowserRenderingPDFParamsCookiesSourceScheme = "Unset"
	BrowserRenderingPDFParamsCookiesSourceSchemeNonSecure BrowserRenderingPDFParamsCookiesSourceScheme = "NonSecure"
	BrowserRenderingPDFParamsCookiesSourceSchemeSecure    BrowserRenderingPDFParamsCookiesSourceScheme = "Secure"
)

func (r BrowserRenderingPDFParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case BrowserRenderingPDFParamsCookiesSourceSchemeUnset, BrowserRenderingPDFParamsCookiesSourceSchemeNonSecure, BrowserRenderingPDFParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type BrowserRenderingPDFParamsGotoOptions struct {
	Referer        param.Field[string]                                             `json:"referer"`
	ReferrerPolicy param.Field[string]                                             `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                            `json:"timeout"`
	WaitUntil      param.Field[BrowserRenderingPDFParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r BrowserRenderingPDFParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.BrowserRenderingPDFParamsGotoOptionsWaitUntilString],
// [browser_rendering.BrowserRenderingPDFParamsGotoOptionsWaitUntilArray].
type BrowserRenderingPDFParamsGotoOptionsWaitUntilUnion interface {
	implementsBrowserRenderingPDFParamsGotoOptionsWaitUntilUnion()
}

type BrowserRenderingPDFParamsGotoOptionsWaitUntilString string

const (
	BrowserRenderingPDFParamsGotoOptionsWaitUntilStringLoad             BrowserRenderingPDFParamsGotoOptionsWaitUntilString = "load"
	BrowserRenderingPDFParamsGotoOptionsWaitUntilStringDomcontentloaded BrowserRenderingPDFParamsGotoOptionsWaitUntilString = "domcontentloaded"
	BrowserRenderingPDFParamsGotoOptionsWaitUntilStringNetworkidle0     BrowserRenderingPDFParamsGotoOptionsWaitUntilString = "networkidle0"
	BrowserRenderingPDFParamsGotoOptionsWaitUntilStringNetworkidle2     BrowserRenderingPDFParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r BrowserRenderingPDFParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case BrowserRenderingPDFParamsGotoOptionsWaitUntilStringLoad, BrowserRenderingPDFParamsGotoOptionsWaitUntilStringDomcontentloaded, BrowserRenderingPDFParamsGotoOptionsWaitUntilStringNetworkidle0, BrowserRenderingPDFParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r BrowserRenderingPDFParamsGotoOptionsWaitUntilString) implementsBrowserRenderingPDFParamsGotoOptionsWaitUntilUnion() {
}

type BrowserRenderingPDFParamsGotoOptionsWaitUntilArray []BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItem

func (r BrowserRenderingPDFParamsGotoOptionsWaitUntilArray) implementsBrowserRenderingPDFParamsGotoOptionsWaitUntilUnion() {
}

type BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItem string

const (
	BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItemLoad             BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItem = "load"
	BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItemDomcontentloaded BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItemNetworkidle0     BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItemNetworkidle2     BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItemLoad, BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItemNetworkidle0, BrowserRenderingPDFParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type BrowserRenderingPDFParamsRejectResourceType string

const (
	BrowserRenderingPDFParamsRejectResourceTypeDocument           BrowserRenderingPDFParamsRejectResourceType = "document"
	BrowserRenderingPDFParamsRejectResourceTypeStylesheet         BrowserRenderingPDFParamsRejectResourceType = "stylesheet"
	BrowserRenderingPDFParamsRejectResourceTypeImage              BrowserRenderingPDFParamsRejectResourceType = "image"
	BrowserRenderingPDFParamsRejectResourceTypeMedia              BrowserRenderingPDFParamsRejectResourceType = "media"
	BrowserRenderingPDFParamsRejectResourceTypeFont               BrowserRenderingPDFParamsRejectResourceType = "font"
	BrowserRenderingPDFParamsRejectResourceTypeScript             BrowserRenderingPDFParamsRejectResourceType = "script"
	BrowserRenderingPDFParamsRejectResourceTypeTexttrack          BrowserRenderingPDFParamsRejectResourceType = "texttrack"
	BrowserRenderingPDFParamsRejectResourceTypeXHR                BrowserRenderingPDFParamsRejectResourceType = "xhr"
	BrowserRenderingPDFParamsRejectResourceTypeFetch              BrowserRenderingPDFParamsRejectResourceType = "fetch"
	BrowserRenderingPDFParamsRejectResourceTypePrefetch           BrowserRenderingPDFParamsRejectResourceType = "prefetch"
	BrowserRenderingPDFParamsRejectResourceTypeEventsource        BrowserRenderingPDFParamsRejectResourceType = "eventsource"
	BrowserRenderingPDFParamsRejectResourceTypeWebsocket          BrowserRenderingPDFParamsRejectResourceType = "websocket"
	BrowserRenderingPDFParamsRejectResourceTypeManifest           BrowserRenderingPDFParamsRejectResourceType = "manifest"
	BrowserRenderingPDFParamsRejectResourceTypeSignedexchange     BrowserRenderingPDFParamsRejectResourceType = "signedexchange"
	BrowserRenderingPDFParamsRejectResourceTypePing               BrowserRenderingPDFParamsRejectResourceType = "ping"
	BrowserRenderingPDFParamsRejectResourceTypeCspviolationreport BrowserRenderingPDFParamsRejectResourceType = "cspviolationreport"
	BrowserRenderingPDFParamsRejectResourceTypePreflight          BrowserRenderingPDFParamsRejectResourceType = "preflight"
	BrowserRenderingPDFParamsRejectResourceTypeOther              BrowserRenderingPDFParamsRejectResourceType = "other"
)

func (r BrowserRenderingPDFParamsRejectResourceType) IsKnown() bool {
	switch r {
	case BrowserRenderingPDFParamsRejectResourceTypeDocument, BrowserRenderingPDFParamsRejectResourceTypeStylesheet, BrowserRenderingPDFParamsRejectResourceTypeImage, BrowserRenderingPDFParamsRejectResourceTypeMedia, BrowserRenderingPDFParamsRejectResourceTypeFont, BrowserRenderingPDFParamsRejectResourceTypeScript, BrowserRenderingPDFParamsRejectResourceTypeTexttrack, BrowserRenderingPDFParamsRejectResourceTypeXHR, BrowserRenderingPDFParamsRejectResourceTypeFetch, BrowserRenderingPDFParamsRejectResourceTypePrefetch, BrowserRenderingPDFParamsRejectResourceTypeEventsource, BrowserRenderingPDFParamsRejectResourceTypeWebsocket, BrowserRenderingPDFParamsRejectResourceTypeManifest, BrowserRenderingPDFParamsRejectResourceTypeSignedexchange, BrowserRenderingPDFParamsRejectResourceTypePing, BrowserRenderingPDFParamsRejectResourceTypeCspviolationreport, BrowserRenderingPDFParamsRejectResourceTypePreflight, BrowserRenderingPDFParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type BrowserRenderingPDFParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r BrowserRenderingPDFParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type BrowserRenderingPDFParamsWaitForSelector struct {
	Selector param.Field[string]                                          `json:"selector,required"`
	Hidden   param.Field[BrowserRenderingPDFParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                         `json:"timeout"`
	Visible  param.Field[BrowserRenderingPDFParamsWaitForSelectorVisible] `json:"visible"`
}

func (r BrowserRenderingPDFParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingPDFParamsWaitForSelectorHidden bool

const (
	BrowserRenderingPDFParamsWaitForSelectorHiddenTrue BrowserRenderingPDFParamsWaitForSelectorHidden = true
)

func (r BrowserRenderingPDFParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case BrowserRenderingPDFParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type BrowserRenderingPDFParamsWaitForSelectorVisible bool

const (
	BrowserRenderingPDFParamsWaitForSelectorVisibleTrue BrowserRenderingPDFParamsWaitForSelectorVisible = true
)

func (r BrowserRenderingPDFParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case BrowserRenderingPDFParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type BrowserRenderingScrapeParams struct {
	Elements param.Field[[]BrowserRenderingScrapeParamsElement] `json:"elements,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]BrowserRenderingScrapeParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]BrowserRenderingScrapeParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]BrowserRenderingScrapeParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[BrowserRenderingScrapeParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]BrowserRenderingScrapeParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                               `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[BrowserRenderingScrapeParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]BrowserRenderingScrapeParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                             `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[BrowserRenderingScrapeParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[BrowserRenderingScrapeParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r BrowserRenderingScrapeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [BrowserRenderingScrapeParams]'s query parameters as
// `url.Values`.
func (r BrowserRenderingScrapeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BrowserRenderingScrapeParamsElement struct {
	Selector param.Field[string] `json:"selector,required"`
}

func (r BrowserRenderingScrapeParamsElement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScrapeParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r BrowserRenderingScrapeParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScrapeParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r BrowserRenderingScrapeParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScrapeParamsAllowResourceType string

const (
	BrowserRenderingScrapeParamsAllowResourceTypeDocument           BrowserRenderingScrapeParamsAllowResourceType = "document"
	BrowserRenderingScrapeParamsAllowResourceTypeStylesheet         BrowserRenderingScrapeParamsAllowResourceType = "stylesheet"
	BrowserRenderingScrapeParamsAllowResourceTypeImage              BrowserRenderingScrapeParamsAllowResourceType = "image"
	BrowserRenderingScrapeParamsAllowResourceTypeMedia              BrowserRenderingScrapeParamsAllowResourceType = "media"
	BrowserRenderingScrapeParamsAllowResourceTypeFont               BrowserRenderingScrapeParamsAllowResourceType = "font"
	BrowserRenderingScrapeParamsAllowResourceTypeScript             BrowserRenderingScrapeParamsAllowResourceType = "script"
	BrowserRenderingScrapeParamsAllowResourceTypeTexttrack          BrowserRenderingScrapeParamsAllowResourceType = "texttrack"
	BrowserRenderingScrapeParamsAllowResourceTypeXHR                BrowserRenderingScrapeParamsAllowResourceType = "xhr"
	BrowserRenderingScrapeParamsAllowResourceTypeFetch              BrowserRenderingScrapeParamsAllowResourceType = "fetch"
	BrowserRenderingScrapeParamsAllowResourceTypePrefetch           BrowserRenderingScrapeParamsAllowResourceType = "prefetch"
	BrowserRenderingScrapeParamsAllowResourceTypeEventsource        BrowserRenderingScrapeParamsAllowResourceType = "eventsource"
	BrowserRenderingScrapeParamsAllowResourceTypeWebsocket          BrowserRenderingScrapeParamsAllowResourceType = "websocket"
	BrowserRenderingScrapeParamsAllowResourceTypeManifest           BrowserRenderingScrapeParamsAllowResourceType = "manifest"
	BrowserRenderingScrapeParamsAllowResourceTypeSignedexchange     BrowserRenderingScrapeParamsAllowResourceType = "signedexchange"
	BrowserRenderingScrapeParamsAllowResourceTypePing               BrowserRenderingScrapeParamsAllowResourceType = "ping"
	BrowserRenderingScrapeParamsAllowResourceTypeCspviolationreport BrowserRenderingScrapeParamsAllowResourceType = "cspviolationreport"
	BrowserRenderingScrapeParamsAllowResourceTypePreflight          BrowserRenderingScrapeParamsAllowResourceType = "preflight"
	BrowserRenderingScrapeParamsAllowResourceTypeOther              BrowserRenderingScrapeParamsAllowResourceType = "other"
)

func (r BrowserRenderingScrapeParamsAllowResourceType) IsKnown() bool {
	switch r {
	case BrowserRenderingScrapeParamsAllowResourceTypeDocument, BrowserRenderingScrapeParamsAllowResourceTypeStylesheet, BrowserRenderingScrapeParamsAllowResourceTypeImage, BrowserRenderingScrapeParamsAllowResourceTypeMedia, BrowserRenderingScrapeParamsAllowResourceTypeFont, BrowserRenderingScrapeParamsAllowResourceTypeScript, BrowserRenderingScrapeParamsAllowResourceTypeTexttrack, BrowserRenderingScrapeParamsAllowResourceTypeXHR, BrowserRenderingScrapeParamsAllowResourceTypeFetch, BrowserRenderingScrapeParamsAllowResourceTypePrefetch, BrowserRenderingScrapeParamsAllowResourceTypeEventsource, BrowserRenderingScrapeParamsAllowResourceTypeWebsocket, BrowserRenderingScrapeParamsAllowResourceTypeManifest, BrowserRenderingScrapeParamsAllowResourceTypeSignedexchange, BrowserRenderingScrapeParamsAllowResourceTypePing, BrowserRenderingScrapeParamsAllowResourceTypeCspviolationreport, BrowserRenderingScrapeParamsAllowResourceTypePreflight, BrowserRenderingScrapeParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type BrowserRenderingScrapeParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r BrowserRenderingScrapeParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScrapeParamsCookie struct {
	Name         param.Field[string]                                          `json:"name,required"`
	Value        param.Field[string]                                          `json:"value,required"`
	Domain       param.Field[string]                                          `json:"domain"`
	Expires      param.Field[float64]                                         `json:"expires"`
	HTTPOnly     param.Field[bool]                                            `json:"httpOnly"`
	PartitionKey param.Field[string]                                          `json:"partitionKey"`
	Path         param.Field[string]                                          `json:"path"`
	Priority     param.Field[BrowserRenderingScrapeParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                            `json:"sameParty"`
	SameSite     param.Field[BrowserRenderingScrapeParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                            `json:"secure"`
	SourcePort   param.Field[float64]                                         `json:"sourcePort"`
	SourceScheme param.Field[BrowserRenderingScrapeParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                          `json:"url"`
}

func (r BrowserRenderingScrapeParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScrapeParamsCookiesPriority string

const (
	BrowserRenderingScrapeParamsCookiesPriorityLow    BrowserRenderingScrapeParamsCookiesPriority = "Low"
	BrowserRenderingScrapeParamsCookiesPriorityMedium BrowserRenderingScrapeParamsCookiesPriority = "Medium"
	BrowserRenderingScrapeParamsCookiesPriorityHigh   BrowserRenderingScrapeParamsCookiesPriority = "High"
)

func (r BrowserRenderingScrapeParamsCookiesPriority) IsKnown() bool {
	switch r {
	case BrowserRenderingScrapeParamsCookiesPriorityLow, BrowserRenderingScrapeParamsCookiesPriorityMedium, BrowserRenderingScrapeParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type BrowserRenderingScrapeParamsCookiesSameSite string

const (
	BrowserRenderingScrapeParamsCookiesSameSiteStrict BrowserRenderingScrapeParamsCookiesSameSite = "Strict"
	BrowserRenderingScrapeParamsCookiesSameSiteLax    BrowserRenderingScrapeParamsCookiesSameSite = "Lax"
	BrowserRenderingScrapeParamsCookiesSameSiteNone   BrowserRenderingScrapeParamsCookiesSameSite = "None"
)

func (r BrowserRenderingScrapeParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case BrowserRenderingScrapeParamsCookiesSameSiteStrict, BrowserRenderingScrapeParamsCookiesSameSiteLax, BrowserRenderingScrapeParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type BrowserRenderingScrapeParamsCookiesSourceScheme string

const (
	BrowserRenderingScrapeParamsCookiesSourceSchemeUnset     BrowserRenderingScrapeParamsCookiesSourceScheme = "Unset"
	BrowserRenderingScrapeParamsCookiesSourceSchemeNonSecure BrowserRenderingScrapeParamsCookiesSourceScheme = "NonSecure"
	BrowserRenderingScrapeParamsCookiesSourceSchemeSecure    BrowserRenderingScrapeParamsCookiesSourceScheme = "Secure"
)

func (r BrowserRenderingScrapeParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case BrowserRenderingScrapeParamsCookiesSourceSchemeUnset, BrowserRenderingScrapeParamsCookiesSourceSchemeNonSecure, BrowserRenderingScrapeParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type BrowserRenderingScrapeParamsGotoOptions struct {
	Referer        param.Field[string]                                                `json:"referer"`
	ReferrerPolicy param.Field[string]                                                `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                               `json:"timeout"`
	WaitUntil      param.Field[BrowserRenderingScrapeParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r BrowserRenderingScrapeParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.BrowserRenderingScrapeParamsGotoOptionsWaitUntilString],
// [browser_rendering.BrowserRenderingScrapeParamsGotoOptionsWaitUntilArray].
type BrowserRenderingScrapeParamsGotoOptionsWaitUntilUnion interface {
	implementsBrowserRenderingScrapeParamsGotoOptionsWaitUntilUnion()
}

type BrowserRenderingScrapeParamsGotoOptionsWaitUntilString string

const (
	BrowserRenderingScrapeParamsGotoOptionsWaitUntilStringLoad             BrowserRenderingScrapeParamsGotoOptionsWaitUntilString = "load"
	BrowserRenderingScrapeParamsGotoOptionsWaitUntilStringDomcontentloaded BrowserRenderingScrapeParamsGotoOptionsWaitUntilString = "domcontentloaded"
	BrowserRenderingScrapeParamsGotoOptionsWaitUntilStringNetworkidle0     BrowserRenderingScrapeParamsGotoOptionsWaitUntilString = "networkidle0"
	BrowserRenderingScrapeParamsGotoOptionsWaitUntilStringNetworkidle2     BrowserRenderingScrapeParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r BrowserRenderingScrapeParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case BrowserRenderingScrapeParamsGotoOptionsWaitUntilStringLoad, BrowserRenderingScrapeParamsGotoOptionsWaitUntilStringDomcontentloaded, BrowserRenderingScrapeParamsGotoOptionsWaitUntilStringNetworkidle0, BrowserRenderingScrapeParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r BrowserRenderingScrapeParamsGotoOptionsWaitUntilString) implementsBrowserRenderingScrapeParamsGotoOptionsWaitUntilUnion() {
}

type BrowserRenderingScrapeParamsGotoOptionsWaitUntilArray []BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItem

func (r BrowserRenderingScrapeParamsGotoOptionsWaitUntilArray) implementsBrowserRenderingScrapeParamsGotoOptionsWaitUntilUnion() {
}

type BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItem string

const (
	BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItemLoad             BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItem = "load"
	BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItemDomcontentloaded BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItemNetworkidle0     BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItemNetworkidle2     BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItemLoad, BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItemNetworkidle0, BrowserRenderingScrapeParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type BrowserRenderingScrapeParamsRejectResourceType string

const (
	BrowserRenderingScrapeParamsRejectResourceTypeDocument           BrowserRenderingScrapeParamsRejectResourceType = "document"
	BrowserRenderingScrapeParamsRejectResourceTypeStylesheet         BrowserRenderingScrapeParamsRejectResourceType = "stylesheet"
	BrowserRenderingScrapeParamsRejectResourceTypeImage              BrowserRenderingScrapeParamsRejectResourceType = "image"
	BrowserRenderingScrapeParamsRejectResourceTypeMedia              BrowserRenderingScrapeParamsRejectResourceType = "media"
	BrowserRenderingScrapeParamsRejectResourceTypeFont               BrowserRenderingScrapeParamsRejectResourceType = "font"
	BrowserRenderingScrapeParamsRejectResourceTypeScript             BrowserRenderingScrapeParamsRejectResourceType = "script"
	BrowserRenderingScrapeParamsRejectResourceTypeTexttrack          BrowserRenderingScrapeParamsRejectResourceType = "texttrack"
	BrowserRenderingScrapeParamsRejectResourceTypeXHR                BrowserRenderingScrapeParamsRejectResourceType = "xhr"
	BrowserRenderingScrapeParamsRejectResourceTypeFetch              BrowserRenderingScrapeParamsRejectResourceType = "fetch"
	BrowserRenderingScrapeParamsRejectResourceTypePrefetch           BrowserRenderingScrapeParamsRejectResourceType = "prefetch"
	BrowserRenderingScrapeParamsRejectResourceTypeEventsource        BrowserRenderingScrapeParamsRejectResourceType = "eventsource"
	BrowserRenderingScrapeParamsRejectResourceTypeWebsocket          BrowserRenderingScrapeParamsRejectResourceType = "websocket"
	BrowserRenderingScrapeParamsRejectResourceTypeManifest           BrowserRenderingScrapeParamsRejectResourceType = "manifest"
	BrowserRenderingScrapeParamsRejectResourceTypeSignedexchange     BrowserRenderingScrapeParamsRejectResourceType = "signedexchange"
	BrowserRenderingScrapeParamsRejectResourceTypePing               BrowserRenderingScrapeParamsRejectResourceType = "ping"
	BrowserRenderingScrapeParamsRejectResourceTypeCspviolationreport BrowserRenderingScrapeParamsRejectResourceType = "cspviolationreport"
	BrowserRenderingScrapeParamsRejectResourceTypePreflight          BrowserRenderingScrapeParamsRejectResourceType = "preflight"
	BrowserRenderingScrapeParamsRejectResourceTypeOther              BrowserRenderingScrapeParamsRejectResourceType = "other"
)

func (r BrowserRenderingScrapeParamsRejectResourceType) IsKnown() bool {
	switch r {
	case BrowserRenderingScrapeParamsRejectResourceTypeDocument, BrowserRenderingScrapeParamsRejectResourceTypeStylesheet, BrowserRenderingScrapeParamsRejectResourceTypeImage, BrowserRenderingScrapeParamsRejectResourceTypeMedia, BrowserRenderingScrapeParamsRejectResourceTypeFont, BrowserRenderingScrapeParamsRejectResourceTypeScript, BrowserRenderingScrapeParamsRejectResourceTypeTexttrack, BrowserRenderingScrapeParamsRejectResourceTypeXHR, BrowserRenderingScrapeParamsRejectResourceTypeFetch, BrowserRenderingScrapeParamsRejectResourceTypePrefetch, BrowserRenderingScrapeParamsRejectResourceTypeEventsource, BrowserRenderingScrapeParamsRejectResourceTypeWebsocket, BrowserRenderingScrapeParamsRejectResourceTypeManifest, BrowserRenderingScrapeParamsRejectResourceTypeSignedexchange, BrowserRenderingScrapeParamsRejectResourceTypePing, BrowserRenderingScrapeParamsRejectResourceTypeCspviolationreport, BrowserRenderingScrapeParamsRejectResourceTypePreflight, BrowserRenderingScrapeParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type BrowserRenderingScrapeParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r BrowserRenderingScrapeParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type BrowserRenderingScrapeParamsWaitForSelector struct {
	Selector param.Field[string]                                             `json:"selector,required"`
	Hidden   param.Field[BrowserRenderingScrapeParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                            `json:"timeout"`
	Visible  param.Field[BrowserRenderingScrapeParamsWaitForSelectorVisible] `json:"visible"`
}

func (r BrowserRenderingScrapeParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScrapeParamsWaitForSelectorHidden bool

const (
	BrowserRenderingScrapeParamsWaitForSelectorHiddenTrue BrowserRenderingScrapeParamsWaitForSelectorHidden = true
)

func (r BrowserRenderingScrapeParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case BrowserRenderingScrapeParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type BrowserRenderingScrapeParamsWaitForSelectorVisible bool

const (
	BrowserRenderingScrapeParamsWaitForSelectorVisibleTrue BrowserRenderingScrapeParamsWaitForSelectorVisible = true
)

func (r BrowserRenderingScrapeParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case BrowserRenderingScrapeParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type BrowserRenderingScrapeResponseEnvelope struct {
	Result []BrowserRenderingScrapeResponse `json:"result,required"`
	// Response status
	Status bool                                           `json:"status,required"`
	Errors []BrowserRenderingScrapeResponseEnvelopeErrors `json:"errors"`
	JSON   browserRenderingScrapeResponseEnvelopeJSON     `json:"-"`
}

// browserRenderingScrapeResponseEnvelopeJSON contains the JSON metadata for the
// struct [BrowserRenderingScrapeResponseEnvelope]
type browserRenderingScrapeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Status      apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserRenderingScrapeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserRenderingScrapeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BrowserRenderingScrapeResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                           `json:"message,required"`
	JSON    browserRenderingScrapeResponseEnvelopeErrorsJSON `json:"-"`
}

// browserRenderingScrapeResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BrowserRenderingScrapeResponseEnvelopeErrors]
type browserRenderingScrapeResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserRenderingScrapeResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserRenderingScrapeResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BrowserRenderingScreenshotParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]BrowserRenderingScreenshotParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]BrowserRenderingScreenshotParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]BrowserRenderingScreenshotParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[BrowserRenderingScreenshotParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]BrowserRenderingScreenshotParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                   `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[BrowserRenderingScreenshotParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes param.Field[[]BrowserRenderingScreenshotParamsRejectResourceType] `json:"rejectResourceTypes"`
	// Check [options](https://pptr.dev/api/puppeteer.screenshotoptions).
	ScreenshotOptions    param.Field[BrowserRenderingScreenshotParamsScreenshotOptions] `json:"screenshotOptions"`
	ScrollPage           param.Field[bool]                                              `json:"scrollPage"`
	Selector             param.Field[string]                                            `json:"selector"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                 `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                              `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[BrowserRenderingScreenshotParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[BrowserRenderingScreenshotParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r BrowserRenderingScreenshotParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [BrowserRenderingScreenshotParams]'s query parameters as
// `url.Values`.
func (r BrowserRenderingScreenshotParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BrowserRenderingScreenshotParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r BrowserRenderingScreenshotParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScreenshotParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r BrowserRenderingScreenshotParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScreenshotParamsAllowResourceType string

const (
	BrowserRenderingScreenshotParamsAllowResourceTypeDocument           BrowserRenderingScreenshotParamsAllowResourceType = "document"
	BrowserRenderingScreenshotParamsAllowResourceTypeStylesheet         BrowserRenderingScreenshotParamsAllowResourceType = "stylesheet"
	BrowserRenderingScreenshotParamsAllowResourceTypeImage              BrowserRenderingScreenshotParamsAllowResourceType = "image"
	BrowserRenderingScreenshotParamsAllowResourceTypeMedia              BrowserRenderingScreenshotParamsAllowResourceType = "media"
	BrowserRenderingScreenshotParamsAllowResourceTypeFont               BrowserRenderingScreenshotParamsAllowResourceType = "font"
	BrowserRenderingScreenshotParamsAllowResourceTypeScript             BrowserRenderingScreenshotParamsAllowResourceType = "script"
	BrowserRenderingScreenshotParamsAllowResourceTypeTexttrack          BrowserRenderingScreenshotParamsAllowResourceType = "texttrack"
	BrowserRenderingScreenshotParamsAllowResourceTypeXHR                BrowserRenderingScreenshotParamsAllowResourceType = "xhr"
	BrowserRenderingScreenshotParamsAllowResourceTypeFetch              BrowserRenderingScreenshotParamsAllowResourceType = "fetch"
	BrowserRenderingScreenshotParamsAllowResourceTypePrefetch           BrowserRenderingScreenshotParamsAllowResourceType = "prefetch"
	BrowserRenderingScreenshotParamsAllowResourceTypeEventsource        BrowserRenderingScreenshotParamsAllowResourceType = "eventsource"
	BrowserRenderingScreenshotParamsAllowResourceTypeWebsocket          BrowserRenderingScreenshotParamsAllowResourceType = "websocket"
	BrowserRenderingScreenshotParamsAllowResourceTypeManifest           BrowserRenderingScreenshotParamsAllowResourceType = "manifest"
	BrowserRenderingScreenshotParamsAllowResourceTypeSignedexchange     BrowserRenderingScreenshotParamsAllowResourceType = "signedexchange"
	BrowserRenderingScreenshotParamsAllowResourceTypePing               BrowserRenderingScreenshotParamsAllowResourceType = "ping"
	BrowserRenderingScreenshotParamsAllowResourceTypeCspviolationreport BrowserRenderingScreenshotParamsAllowResourceType = "cspviolationreport"
	BrowserRenderingScreenshotParamsAllowResourceTypePreflight          BrowserRenderingScreenshotParamsAllowResourceType = "preflight"
	BrowserRenderingScreenshotParamsAllowResourceTypeOther              BrowserRenderingScreenshotParamsAllowResourceType = "other"
)

func (r BrowserRenderingScreenshotParamsAllowResourceType) IsKnown() bool {
	switch r {
	case BrowserRenderingScreenshotParamsAllowResourceTypeDocument, BrowserRenderingScreenshotParamsAllowResourceTypeStylesheet, BrowserRenderingScreenshotParamsAllowResourceTypeImage, BrowserRenderingScreenshotParamsAllowResourceTypeMedia, BrowserRenderingScreenshotParamsAllowResourceTypeFont, BrowserRenderingScreenshotParamsAllowResourceTypeScript, BrowserRenderingScreenshotParamsAllowResourceTypeTexttrack, BrowserRenderingScreenshotParamsAllowResourceTypeXHR, BrowserRenderingScreenshotParamsAllowResourceTypeFetch, BrowserRenderingScreenshotParamsAllowResourceTypePrefetch, BrowserRenderingScreenshotParamsAllowResourceTypeEventsource, BrowserRenderingScreenshotParamsAllowResourceTypeWebsocket, BrowserRenderingScreenshotParamsAllowResourceTypeManifest, BrowserRenderingScreenshotParamsAllowResourceTypeSignedexchange, BrowserRenderingScreenshotParamsAllowResourceTypePing, BrowserRenderingScreenshotParamsAllowResourceTypeCspviolationreport, BrowserRenderingScreenshotParamsAllowResourceTypePreflight, BrowserRenderingScreenshotParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type BrowserRenderingScreenshotParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r BrowserRenderingScreenshotParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScreenshotParamsCookie struct {
	Name         param.Field[string]                                              `json:"name,required"`
	Value        param.Field[string]                                              `json:"value,required"`
	Domain       param.Field[string]                                              `json:"domain"`
	Expires      param.Field[float64]                                             `json:"expires"`
	HTTPOnly     param.Field[bool]                                                `json:"httpOnly"`
	PartitionKey param.Field[string]                                              `json:"partitionKey"`
	Path         param.Field[string]                                              `json:"path"`
	Priority     param.Field[BrowserRenderingScreenshotParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                                `json:"sameParty"`
	SameSite     param.Field[BrowserRenderingScreenshotParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                                `json:"secure"`
	SourcePort   param.Field[float64]                                             `json:"sourcePort"`
	SourceScheme param.Field[BrowserRenderingScreenshotParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                              `json:"url"`
}

func (r BrowserRenderingScreenshotParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScreenshotParamsCookiesPriority string

const (
	BrowserRenderingScreenshotParamsCookiesPriorityLow    BrowserRenderingScreenshotParamsCookiesPriority = "Low"
	BrowserRenderingScreenshotParamsCookiesPriorityMedium BrowserRenderingScreenshotParamsCookiesPriority = "Medium"
	BrowserRenderingScreenshotParamsCookiesPriorityHigh   BrowserRenderingScreenshotParamsCookiesPriority = "High"
)

func (r BrowserRenderingScreenshotParamsCookiesPriority) IsKnown() bool {
	switch r {
	case BrowserRenderingScreenshotParamsCookiesPriorityLow, BrowserRenderingScreenshotParamsCookiesPriorityMedium, BrowserRenderingScreenshotParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type BrowserRenderingScreenshotParamsCookiesSameSite string

const (
	BrowserRenderingScreenshotParamsCookiesSameSiteStrict BrowserRenderingScreenshotParamsCookiesSameSite = "Strict"
	BrowserRenderingScreenshotParamsCookiesSameSiteLax    BrowserRenderingScreenshotParamsCookiesSameSite = "Lax"
	BrowserRenderingScreenshotParamsCookiesSameSiteNone   BrowserRenderingScreenshotParamsCookiesSameSite = "None"
)

func (r BrowserRenderingScreenshotParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case BrowserRenderingScreenshotParamsCookiesSameSiteStrict, BrowserRenderingScreenshotParamsCookiesSameSiteLax, BrowserRenderingScreenshotParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type BrowserRenderingScreenshotParamsCookiesSourceScheme string

const (
	BrowserRenderingScreenshotParamsCookiesSourceSchemeUnset     BrowserRenderingScreenshotParamsCookiesSourceScheme = "Unset"
	BrowserRenderingScreenshotParamsCookiesSourceSchemeNonSecure BrowserRenderingScreenshotParamsCookiesSourceScheme = "NonSecure"
	BrowserRenderingScreenshotParamsCookiesSourceSchemeSecure    BrowserRenderingScreenshotParamsCookiesSourceScheme = "Secure"
)

func (r BrowserRenderingScreenshotParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case BrowserRenderingScreenshotParamsCookiesSourceSchemeUnset, BrowserRenderingScreenshotParamsCookiesSourceSchemeNonSecure, BrowserRenderingScreenshotParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type BrowserRenderingScreenshotParamsGotoOptions struct {
	Referer        param.Field[string]                                                    `json:"referer"`
	ReferrerPolicy param.Field[string]                                                    `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                   `json:"timeout"`
	WaitUntil      param.Field[BrowserRenderingScreenshotParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r BrowserRenderingScreenshotParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.BrowserRenderingScreenshotParamsGotoOptionsWaitUntilString],
// [browser_rendering.BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArray].
type BrowserRenderingScreenshotParamsGotoOptionsWaitUntilUnion interface {
	implementsBrowserRenderingScreenshotParamsGotoOptionsWaitUntilUnion()
}

type BrowserRenderingScreenshotParamsGotoOptionsWaitUntilString string

const (
	BrowserRenderingScreenshotParamsGotoOptionsWaitUntilStringLoad             BrowserRenderingScreenshotParamsGotoOptionsWaitUntilString = "load"
	BrowserRenderingScreenshotParamsGotoOptionsWaitUntilStringDomcontentloaded BrowserRenderingScreenshotParamsGotoOptionsWaitUntilString = "domcontentloaded"
	BrowserRenderingScreenshotParamsGotoOptionsWaitUntilStringNetworkidle0     BrowserRenderingScreenshotParamsGotoOptionsWaitUntilString = "networkidle0"
	BrowserRenderingScreenshotParamsGotoOptionsWaitUntilStringNetworkidle2     BrowserRenderingScreenshotParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r BrowserRenderingScreenshotParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case BrowserRenderingScreenshotParamsGotoOptionsWaitUntilStringLoad, BrowserRenderingScreenshotParamsGotoOptionsWaitUntilStringDomcontentloaded, BrowserRenderingScreenshotParamsGotoOptionsWaitUntilStringNetworkidle0, BrowserRenderingScreenshotParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r BrowserRenderingScreenshotParamsGotoOptionsWaitUntilString) implementsBrowserRenderingScreenshotParamsGotoOptionsWaitUntilUnion() {
}

type BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArray []BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItem

func (r BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArray) implementsBrowserRenderingScreenshotParamsGotoOptionsWaitUntilUnion() {
}

type BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItem string

const (
	BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItemLoad             BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItem = "load"
	BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItemDomcontentloaded BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItemNetworkidle0     BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItemNetworkidle2     BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItemLoad, BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItemNetworkidle0, BrowserRenderingScreenshotParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type BrowserRenderingScreenshotParamsRejectResourceType string

const (
	BrowserRenderingScreenshotParamsRejectResourceTypeDocument           BrowserRenderingScreenshotParamsRejectResourceType = "document"
	BrowserRenderingScreenshotParamsRejectResourceTypeStylesheet         BrowserRenderingScreenshotParamsRejectResourceType = "stylesheet"
	BrowserRenderingScreenshotParamsRejectResourceTypeImage              BrowserRenderingScreenshotParamsRejectResourceType = "image"
	BrowserRenderingScreenshotParamsRejectResourceTypeMedia              BrowserRenderingScreenshotParamsRejectResourceType = "media"
	BrowserRenderingScreenshotParamsRejectResourceTypeFont               BrowserRenderingScreenshotParamsRejectResourceType = "font"
	BrowserRenderingScreenshotParamsRejectResourceTypeScript             BrowserRenderingScreenshotParamsRejectResourceType = "script"
	BrowserRenderingScreenshotParamsRejectResourceTypeTexttrack          BrowserRenderingScreenshotParamsRejectResourceType = "texttrack"
	BrowserRenderingScreenshotParamsRejectResourceTypeXHR                BrowserRenderingScreenshotParamsRejectResourceType = "xhr"
	BrowserRenderingScreenshotParamsRejectResourceTypeFetch              BrowserRenderingScreenshotParamsRejectResourceType = "fetch"
	BrowserRenderingScreenshotParamsRejectResourceTypePrefetch           BrowserRenderingScreenshotParamsRejectResourceType = "prefetch"
	BrowserRenderingScreenshotParamsRejectResourceTypeEventsource        BrowserRenderingScreenshotParamsRejectResourceType = "eventsource"
	BrowserRenderingScreenshotParamsRejectResourceTypeWebsocket          BrowserRenderingScreenshotParamsRejectResourceType = "websocket"
	BrowserRenderingScreenshotParamsRejectResourceTypeManifest           BrowserRenderingScreenshotParamsRejectResourceType = "manifest"
	BrowserRenderingScreenshotParamsRejectResourceTypeSignedexchange     BrowserRenderingScreenshotParamsRejectResourceType = "signedexchange"
	BrowserRenderingScreenshotParamsRejectResourceTypePing               BrowserRenderingScreenshotParamsRejectResourceType = "ping"
	BrowserRenderingScreenshotParamsRejectResourceTypeCspviolationreport BrowserRenderingScreenshotParamsRejectResourceType = "cspviolationreport"
	BrowserRenderingScreenshotParamsRejectResourceTypePreflight          BrowserRenderingScreenshotParamsRejectResourceType = "preflight"
	BrowserRenderingScreenshotParamsRejectResourceTypeOther              BrowserRenderingScreenshotParamsRejectResourceType = "other"
)

func (r BrowserRenderingScreenshotParamsRejectResourceType) IsKnown() bool {
	switch r {
	case BrowserRenderingScreenshotParamsRejectResourceTypeDocument, BrowserRenderingScreenshotParamsRejectResourceTypeStylesheet, BrowserRenderingScreenshotParamsRejectResourceTypeImage, BrowserRenderingScreenshotParamsRejectResourceTypeMedia, BrowserRenderingScreenshotParamsRejectResourceTypeFont, BrowserRenderingScreenshotParamsRejectResourceTypeScript, BrowserRenderingScreenshotParamsRejectResourceTypeTexttrack, BrowserRenderingScreenshotParamsRejectResourceTypeXHR, BrowserRenderingScreenshotParamsRejectResourceTypeFetch, BrowserRenderingScreenshotParamsRejectResourceTypePrefetch, BrowserRenderingScreenshotParamsRejectResourceTypeEventsource, BrowserRenderingScreenshotParamsRejectResourceTypeWebsocket, BrowserRenderingScreenshotParamsRejectResourceTypeManifest, BrowserRenderingScreenshotParamsRejectResourceTypeSignedexchange, BrowserRenderingScreenshotParamsRejectResourceTypePing, BrowserRenderingScreenshotParamsRejectResourceTypeCspviolationreport, BrowserRenderingScreenshotParamsRejectResourceTypePreflight, BrowserRenderingScreenshotParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.screenshotoptions).
type BrowserRenderingScreenshotParamsScreenshotOptions struct {
	CaptureBeyondViewport param.Field[bool]                                                      `json:"captureBeyondViewport"`
	Clip                  param.Field[BrowserRenderingScreenshotParamsScreenshotOptionsClip]     `json:"clip"`
	Encoding              param.Field[BrowserRenderingScreenshotParamsScreenshotOptionsEncoding] `json:"encoding"`
	FromSurface           param.Field[bool]                                                      `json:"fromSurface"`
	FullPage              param.Field[bool]                                                      `json:"fullPage"`
	OmitBackground        param.Field[bool]                                                      `json:"omitBackground"`
	OptimizeForSpeed      param.Field[bool]                                                      `json:"optimizeForSpeed"`
	Quality               param.Field[float64]                                                   `json:"quality"`
	Type                  param.Field[BrowserRenderingScreenshotParamsScreenshotOptionsType]     `json:"type"`
}

func (r BrowserRenderingScreenshotParamsScreenshotOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScreenshotParamsScreenshotOptionsClip struct {
	Height param.Field[float64] `json:"height,required"`
	Width  param.Field[float64] `json:"width,required"`
	X      param.Field[float64] `json:"x,required"`
	Y      param.Field[float64] `json:"y,required"`
	Scale  param.Field[float64] `json:"scale"`
}

func (r BrowserRenderingScreenshotParamsScreenshotOptionsClip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScreenshotParamsScreenshotOptionsEncoding string

const (
	BrowserRenderingScreenshotParamsScreenshotOptionsEncodingBinary BrowserRenderingScreenshotParamsScreenshotOptionsEncoding = "binary"
	BrowserRenderingScreenshotParamsScreenshotOptionsEncodingBase64 BrowserRenderingScreenshotParamsScreenshotOptionsEncoding = "base64"
)

func (r BrowserRenderingScreenshotParamsScreenshotOptionsEncoding) IsKnown() bool {
	switch r {
	case BrowserRenderingScreenshotParamsScreenshotOptionsEncodingBinary, BrowserRenderingScreenshotParamsScreenshotOptionsEncodingBase64:
		return true
	}
	return false
}

type BrowserRenderingScreenshotParamsScreenshotOptionsType string

const (
	BrowserRenderingScreenshotParamsScreenshotOptionsTypePNG  BrowserRenderingScreenshotParamsScreenshotOptionsType = "png"
	BrowserRenderingScreenshotParamsScreenshotOptionsTypeJPEG BrowserRenderingScreenshotParamsScreenshotOptionsType = "jpeg"
	BrowserRenderingScreenshotParamsScreenshotOptionsTypeWebP BrowserRenderingScreenshotParamsScreenshotOptionsType = "webp"
)

func (r BrowserRenderingScreenshotParamsScreenshotOptionsType) IsKnown() bool {
	switch r {
	case BrowserRenderingScreenshotParamsScreenshotOptionsTypePNG, BrowserRenderingScreenshotParamsScreenshotOptionsTypeJPEG, BrowserRenderingScreenshotParamsScreenshotOptionsTypeWebP:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type BrowserRenderingScreenshotParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r BrowserRenderingScreenshotParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type BrowserRenderingScreenshotParamsWaitForSelector struct {
	Selector param.Field[string]                                                 `json:"selector,required"`
	Hidden   param.Field[BrowserRenderingScreenshotParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                                `json:"timeout"`
	Visible  param.Field[BrowserRenderingScreenshotParamsWaitForSelectorVisible] `json:"visible"`
}

func (r BrowserRenderingScreenshotParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingScreenshotParamsWaitForSelectorHidden bool

const (
	BrowserRenderingScreenshotParamsWaitForSelectorHiddenTrue BrowserRenderingScreenshotParamsWaitForSelectorHidden = true
)

func (r BrowserRenderingScreenshotParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case BrowserRenderingScreenshotParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type BrowserRenderingScreenshotParamsWaitForSelectorVisible bool

const (
	BrowserRenderingScreenshotParamsWaitForSelectorVisibleTrue BrowserRenderingScreenshotParamsWaitForSelectorVisible = true
)

func (r BrowserRenderingScreenshotParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case BrowserRenderingScreenshotParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type BrowserRenderingSnapshotParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]BrowserRenderingSnapshotParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]BrowserRenderingSnapshotParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]BrowserRenderingSnapshotParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[BrowserRenderingSnapshotParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]BrowserRenderingSnapshotParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                 `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[BrowserRenderingSnapshotParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]BrowserRenderingSnapshotParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                  `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                               `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[BrowserRenderingSnapshotParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[BrowserRenderingSnapshotParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r BrowserRenderingSnapshotParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [BrowserRenderingSnapshotParams]'s query parameters as
// `url.Values`.
func (r BrowserRenderingSnapshotParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BrowserRenderingSnapshotParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r BrowserRenderingSnapshotParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingSnapshotParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r BrowserRenderingSnapshotParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingSnapshotParamsAllowResourceType string

const (
	BrowserRenderingSnapshotParamsAllowResourceTypeDocument           BrowserRenderingSnapshotParamsAllowResourceType = "document"
	BrowserRenderingSnapshotParamsAllowResourceTypeStylesheet         BrowserRenderingSnapshotParamsAllowResourceType = "stylesheet"
	BrowserRenderingSnapshotParamsAllowResourceTypeImage              BrowserRenderingSnapshotParamsAllowResourceType = "image"
	BrowserRenderingSnapshotParamsAllowResourceTypeMedia              BrowserRenderingSnapshotParamsAllowResourceType = "media"
	BrowserRenderingSnapshotParamsAllowResourceTypeFont               BrowserRenderingSnapshotParamsAllowResourceType = "font"
	BrowserRenderingSnapshotParamsAllowResourceTypeScript             BrowserRenderingSnapshotParamsAllowResourceType = "script"
	BrowserRenderingSnapshotParamsAllowResourceTypeTexttrack          BrowserRenderingSnapshotParamsAllowResourceType = "texttrack"
	BrowserRenderingSnapshotParamsAllowResourceTypeXHR                BrowserRenderingSnapshotParamsAllowResourceType = "xhr"
	BrowserRenderingSnapshotParamsAllowResourceTypeFetch              BrowserRenderingSnapshotParamsAllowResourceType = "fetch"
	BrowserRenderingSnapshotParamsAllowResourceTypePrefetch           BrowserRenderingSnapshotParamsAllowResourceType = "prefetch"
	BrowserRenderingSnapshotParamsAllowResourceTypeEventsource        BrowserRenderingSnapshotParamsAllowResourceType = "eventsource"
	BrowserRenderingSnapshotParamsAllowResourceTypeWebsocket          BrowserRenderingSnapshotParamsAllowResourceType = "websocket"
	BrowserRenderingSnapshotParamsAllowResourceTypeManifest           BrowserRenderingSnapshotParamsAllowResourceType = "manifest"
	BrowserRenderingSnapshotParamsAllowResourceTypeSignedexchange     BrowserRenderingSnapshotParamsAllowResourceType = "signedexchange"
	BrowserRenderingSnapshotParamsAllowResourceTypePing               BrowserRenderingSnapshotParamsAllowResourceType = "ping"
	BrowserRenderingSnapshotParamsAllowResourceTypeCspviolationreport BrowserRenderingSnapshotParamsAllowResourceType = "cspviolationreport"
	BrowserRenderingSnapshotParamsAllowResourceTypePreflight          BrowserRenderingSnapshotParamsAllowResourceType = "preflight"
	BrowserRenderingSnapshotParamsAllowResourceTypeOther              BrowserRenderingSnapshotParamsAllowResourceType = "other"
)

func (r BrowserRenderingSnapshotParamsAllowResourceType) IsKnown() bool {
	switch r {
	case BrowserRenderingSnapshotParamsAllowResourceTypeDocument, BrowserRenderingSnapshotParamsAllowResourceTypeStylesheet, BrowserRenderingSnapshotParamsAllowResourceTypeImage, BrowserRenderingSnapshotParamsAllowResourceTypeMedia, BrowserRenderingSnapshotParamsAllowResourceTypeFont, BrowserRenderingSnapshotParamsAllowResourceTypeScript, BrowserRenderingSnapshotParamsAllowResourceTypeTexttrack, BrowserRenderingSnapshotParamsAllowResourceTypeXHR, BrowserRenderingSnapshotParamsAllowResourceTypeFetch, BrowserRenderingSnapshotParamsAllowResourceTypePrefetch, BrowserRenderingSnapshotParamsAllowResourceTypeEventsource, BrowserRenderingSnapshotParamsAllowResourceTypeWebsocket, BrowserRenderingSnapshotParamsAllowResourceTypeManifest, BrowserRenderingSnapshotParamsAllowResourceTypeSignedexchange, BrowserRenderingSnapshotParamsAllowResourceTypePing, BrowserRenderingSnapshotParamsAllowResourceTypeCspviolationreport, BrowserRenderingSnapshotParamsAllowResourceTypePreflight, BrowserRenderingSnapshotParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type BrowserRenderingSnapshotParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r BrowserRenderingSnapshotParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingSnapshotParamsCookie struct {
	Name         param.Field[string]                                            `json:"name,required"`
	Value        param.Field[string]                                            `json:"value,required"`
	Domain       param.Field[string]                                            `json:"domain"`
	Expires      param.Field[float64]                                           `json:"expires"`
	HTTPOnly     param.Field[bool]                                              `json:"httpOnly"`
	PartitionKey param.Field[string]                                            `json:"partitionKey"`
	Path         param.Field[string]                                            `json:"path"`
	Priority     param.Field[BrowserRenderingSnapshotParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                              `json:"sameParty"`
	SameSite     param.Field[BrowserRenderingSnapshotParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                              `json:"secure"`
	SourcePort   param.Field[float64]                                           `json:"sourcePort"`
	SourceScheme param.Field[BrowserRenderingSnapshotParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                            `json:"url"`
}

func (r BrowserRenderingSnapshotParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingSnapshotParamsCookiesPriority string

const (
	BrowserRenderingSnapshotParamsCookiesPriorityLow    BrowserRenderingSnapshotParamsCookiesPriority = "Low"
	BrowserRenderingSnapshotParamsCookiesPriorityMedium BrowserRenderingSnapshotParamsCookiesPriority = "Medium"
	BrowserRenderingSnapshotParamsCookiesPriorityHigh   BrowserRenderingSnapshotParamsCookiesPriority = "High"
)

func (r BrowserRenderingSnapshotParamsCookiesPriority) IsKnown() bool {
	switch r {
	case BrowserRenderingSnapshotParamsCookiesPriorityLow, BrowserRenderingSnapshotParamsCookiesPriorityMedium, BrowserRenderingSnapshotParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type BrowserRenderingSnapshotParamsCookiesSameSite string

const (
	BrowserRenderingSnapshotParamsCookiesSameSiteStrict BrowserRenderingSnapshotParamsCookiesSameSite = "Strict"
	BrowserRenderingSnapshotParamsCookiesSameSiteLax    BrowserRenderingSnapshotParamsCookiesSameSite = "Lax"
	BrowserRenderingSnapshotParamsCookiesSameSiteNone   BrowserRenderingSnapshotParamsCookiesSameSite = "None"
)

func (r BrowserRenderingSnapshotParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case BrowserRenderingSnapshotParamsCookiesSameSiteStrict, BrowserRenderingSnapshotParamsCookiesSameSiteLax, BrowserRenderingSnapshotParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type BrowserRenderingSnapshotParamsCookiesSourceScheme string

const (
	BrowserRenderingSnapshotParamsCookiesSourceSchemeUnset     BrowserRenderingSnapshotParamsCookiesSourceScheme = "Unset"
	BrowserRenderingSnapshotParamsCookiesSourceSchemeNonSecure BrowserRenderingSnapshotParamsCookiesSourceScheme = "NonSecure"
	BrowserRenderingSnapshotParamsCookiesSourceSchemeSecure    BrowserRenderingSnapshotParamsCookiesSourceScheme = "Secure"
)

func (r BrowserRenderingSnapshotParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case BrowserRenderingSnapshotParamsCookiesSourceSchemeUnset, BrowserRenderingSnapshotParamsCookiesSourceSchemeNonSecure, BrowserRenderingSnapshotParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type BrowserRenderingSnapshotParamsGotoOptions struct {
	Referer        param.Field[string]                                                  `json:"referer"`
	ReferrerPolicy param.Field[string]                                                  `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                 `json:"timeout"`
	WaitUntil      param.Field[BrowserRenderingSnapshotParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r BrowserRenderingSnapshotParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.BrowserRenderingSnapshotParamsGotoOptionsWaitUntilString],
// [browser_rendering.BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArray].
type BrowserRenderingSnapshotParamsGotoOptionsWaitUntilUnion interface {
	implementsBrowserRenderingSnapshotParamsGotoOptionsWaitUntilUnion()
}

type BrowserRenderingSnapshotParamsGotoOptionsWaitUntilString string

const (
	BrowserRenderingSnapshotParamsGotoOptionsWaitUntilStringLoad             BrowserRenderingSnapshotParamsGotoOptionsWaitUntilString = "load"
	BrowserRenderingSnapshotParamsGotoOptionsWaitUntilStringDomcontentloaded BrowserRenderingSnapshotParamsGotoOptionsWaitUntilString = "domcontentloaded"
	BrowserRenderingSnapshotParamsGotoOptionsWaitUntilStringNetworkidle0     BrowserRenderingSnapshotParamsGotoOptionsWaitUntilString = "networkidle0"
	BrowserRenderingSnapshotParamsGotoOptionsWaitUntilStringNetworkidle2     BrowserRenderingSnapshotParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r BrowserRenderingSnapshotParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case BrowserRenderingSnapshotParamsGotoOptionsWaitUntilStringLoad, BrowserRenderingSnapshotParamsGotoOptionsWaitUntilStringDomcontentloaded, BrowserRenderingSnapshotParamsGotoOptionsWaitUntilStringNetworkidle0, BrowserRenderingSnapshotParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r BrowserRenderingSnapshotParamsGotoOptionsWaitUntilString) implementsBrowserRenderingSnapshotParamsGotoOptionsWaitUntilUnion() {
}

type BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArray []BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItem

func (r BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArray) implementsBrowserRenderingSnapshotParamsGotoOptionsWaitUntilUnion() {
}

type BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItem string

const (
	BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItemLoad             BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItem = "load"
	BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItemDomcontentloaded BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItemNetworkidle0     BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItemNetworkidle2     BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItemLoad, BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItemNetworkidle0, BrowserRenderingSnapshotParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type BrowserRenderingSnapshotParamsRejectResourceType string

const (
	BrowserRenderingSnapshotParamsRejectResourceTypeDocument           BrowserRenderingSnapshotParamsRejectResourceType = "document"
	BrowserRenderingSnapshotParamsRejectResourceTypeStylesheet         BrowserRenderingSnapshotParamsRejectResourceType = "stylesheet"
	BrowserRenderingSnapshotParamsRejectResourceTypeImage              BrowserRenderingSnapshotParamsRejectResourceType = "image"
	BrowserRenderingSnapshotParamsRejectResourceTypeMedia              BrowserRenderingSnapshotParamsRejectResourceType = "media"
	BrowserRenderingSnapshotParamsRejectResourceTypeFont               BrowserRenderingSnapshotParamsRejectResourceType = "font"
	BrowserRenderingSnapshotParamsRejectResourceTypeScript             BrowserRenderingSnapshotParamsRejectResourceType = "script"
	BrowserRenderingSnapshotParamsRejectResourceTypeTexttrack          BrowserRenderingSnapshotParamsRejectResourceType = "texttrack"
	BrowserRenderingSnapshotParamsRejectResourceTypeXHR                BrowserRenderingSnapshotParamsRejectResourceType = "xhr"
	BrowserRenderingSnapshotParamsRejectResourceTypeFetch              BrowserRenderingSnapshotParamsRejectResourceType = "fetch"
	BrowserRenderingSnapshotParamsRejectResourceTypePrefetch           BrowserRenderingSnapshotParamsRejectResourceType = "prefetch"
	BrowserRenderingSnapshotParamsRejectResourceTypeEventsource        BrowserRenderingSnapshotParamsRejectResourceType = "eventsource"
	BrowserRenderingSnapshotParamsRejectResourceTypeWebsocket          BrowserRenderingSnapshotParamsRejectResourceType = "websocket"
	BrowserRenderingSnapshotParamsRejectResourceTypeManifest           BrowserRenderingSnapshotParamsRejectResourceType = "manifest"
	BrowserRenderingSnapshotParamsRejectResourceTypeSignedexchange     BrowserRenderingSnapshotParamsRejectResourceType = "signedexchange"
	BrowserRenderingSnapshotParamsRejectResourceTypePing               BrowserRenderingSnapshotParamsRejectResourceType = "ping"
	BrowserRenderingSnapshotParamsRejectResourceTypeCspviolationreport BrowserRenderingSnapshotParamsRejectResourceType = "cspviolationreport"
	BrowserRenderingSnapshotParamsRejectResourceTypePreflight          BrowserRenderingSnapshotParamsRejectResourceType = "preflight"
	BrowserRenderingSnapshotParamsRejectResourceTypeOther              BrowserRenderingSnapshotParamsRejectResourceType = "other"
)

func (r BrowserRenderingSnapshotParamsRejectResourceType) IsKnown() bool {
	switch r {
	case BrowserRenderingSnapshotParamsRejectResourceTypeDocument, BrowserRenderingSnapshotParamsRejectResourceTypeStylesheet, BrowserRenderingSnapshotParamsRejectResourceTypeImage, BrowserRenderingSnapshotParamsRejectResourceTypeMedia, BrowserRenderingSnapshotParamsRejectResourceTypeFont, BrowserRenderingSnapshotParamsRejectResourceTypeScript, BrowserRenderingSnapshotParamsRejectResourceTypeTexttrack, BrowserRenderingSnapshotParamsRejectResourceTypeXHR, BrowserRenderingSnapshotParamsRejectResourceTypeFetch, BrowserRenderingSnapshotParamsRejectResourceTypePrefetch, BrowserRenderingSnapshotParamsRejectResourceTypeEventsource, BrowserRenderingSnapshotParamsRejectResourceTypeWebsocket, BrowserRenderingSnapshotParamsRejectResourceTypeManifest, BrowserRenderingSnapshotParamsRejectResourceTypeSignedexchange, BrowserRenderingSnapshotParamsRejectResourceTypePing, BrowserRenderingSnapshotParamsRejectResourceTypeCspviolationreport, BrowserRenderingSnapshotParamsRejectResourceTypePreflight, BrowserRenderingSnapshotParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type BrowserRenderingSnapshotParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r BrowserRenderingSnapshotParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type BrowserRenderingSnapshotParamsWaitForSelector struct {
	Selector param.Field[string]                                               `json:"selector,required"`
	Hidden   param.Field[BrowserRenderingSnapshotParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                              `json:"timeout"`
	Visible  param.Field[BrowserRenderingSnapshotParamsWaitForSelectorVisible] `json:"visible"`
}

func (r BrowserRenderingSnapshotParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrowserRenderingSnapshotParamsWaitForSelectorHidden bool

const (
	BrowserRenderingSnapshotParamsWaitForSelectorHiddenTrue BrowserRenderingSnapshotParamsWaitForSelectorHidden = true
)

func (r BrowserRenderingSnapshotParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case BrowserRenderingSnapshotParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type BrowserRenderingSnapshotParamsWaitForSelectorVisible bool

const (
	BrowserRenderingSnapshotParamsWaitForSelectorVisibleTrue BrowserRenderingSnapshotParamsWaitForSelectorVisible = true
)

func (r BrowserRenderingSnapshotParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case BrowserRenderingSnapshotParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type BrowserRenderingSnapshotResponseEnvelope struct {
	// Response status
	Status bool                                             `json:"status,required"`
	Errors []BrowserRenderingSnapshotResponseEnvelopeErrors `json:"errors"`
	Result BrowserRenderingSnapshotResponse                 `json:"result"`
	JSON   browserRenderingSnapshotResponseEnvelopeJSON     `json:"-"`
}

// browserRenderingSnapshotResponseEnvelopeJSON contains the JSON metadata for the
// struct [BrowserRenderingSnapshotResponseEnvelope]
type browserRenderingSnapshotResponseEnvelopeJSON struct {
	Status      apijson.Field
	Errors      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserRenderingSnapshotResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserRenderingSnapshotResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BrowserRenderingSnapshotResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                             `json:"message,required"`
	JSON    browserRenderingSnapshotResponseEnvelopeErrorsJSON `json:"-"`
}

// browserRenderingSnapshotResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [BrowserRenderingSnapshotResponseEnvelopeErrors]
type browserRenderingSnapshotResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserRenderingSnapshotResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserRenderingSnapshotResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
