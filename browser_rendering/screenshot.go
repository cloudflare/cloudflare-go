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

// ScreenshotService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScreenshotService] method instead.
type ScreenshotService struct {
	Options []option.RequestOption
}

// NewScreenshotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScreenshotService(opts ...option.RequestOption) (r *ScreenshotService) {
	r = &ScreenshotService{}
	r.Options = opts
	return
}

// Takes a screenshot of a webpage from provided URL or HTML. Control page loading
// with `gotoOptions` and `waitFor*` options. Customize screenshots with
// `viewport`, `fullPage`, `clip` and others.
func (r *ScreenshotService) New(ctx context.Context, params ScreenshotNewParams, opts ...option.RequestOption) (res *ScreenshotNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/screenshot", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ScreenshotNewResponse struct {
	// Response status
	Success bool                         `json:"success,required"`
	Errors  []ScreenshotNewResponseError `json:"errors"`
	JSON    screenshotNewResponseJSON    `json:"-"`
}

// screenshotNewResponseJSON contains the JSON metadata for the struct
// [ScreenshotNewResponse]
type screenshotNewResponseJSON struct {
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScreenshotNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r screenshotNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScreenshotNewResponseError struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                         `json:"message,required"`
	JSON    screenshotNewResponseErrorJSON `json:"-"`
}

// screenshotNewResponseErrorJSON contains the JSON metadata for the struct
// [ScreenshotNewResponseError]
type screenshotNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScreenshotNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r screenshotNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ScreenshotNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64]         `query:"cacheTTL"`
	Body     ScreenshotNewParamsBodyUnion `json:"body"`
}

func (r ScreenshotNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [ScreenshotNewParams]'s query parameters as `url.Values`.
func (r ScreenshotNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScreenshotNewParamsBody struct {
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
	RejectRequestPattern param.Field[interface{}] `json:"rejectRequestPattern"`
	RejectResourceTypes  param.Field[interface{}] `json:"rejectResourceTypes"`
	ScreenshotOptions    param.Field[interface{}] `json:"screenshotOptions"`
	ScrollPage           param.Field[bool]        `json:"scrollPage"`
	Selector             param.Field[string]      `json:"selector"`
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

func (r ScreenshotNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScreenshotNewParamsBody) implementsScreenshotNewParamsBodyUnion() {}

// Satisfied by [browser_rendering.ScreenshotNewParamsBodyObject],
// [browser_rendering.ScreenshotNewParamsBodyObject], [ScreenshotNewParamsBody].
type ScreenshotNewParamsBodyUnion interface {
	implementsScreenshotNewParamsBodyUnion()
}

type ScreenshotNewParamsBodyObject struct {
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html,required"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]ScreenshotNewParamsBodyObjectAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]ScreenshotNewParamsBodyObjectAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]ScreenshotNewParamsBodyObjectAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[ScreenshotNewParamsBodyObjectAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]ScreenshotNewParamsBodyObjectCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[ScreenshotNewParamsBodyObjectGotoOptions] `json:"gotoOptions"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes param.Field[[]ScreenshotNewParamsBodyObjectRejectResourceType] `json:"rejectResourceTypes"`
	// Check [options](https://pptr.dev/api/puppeteer.screenshotoptions).
	ScreenshotOptions    param.Field[ScreenshotNewParamsBodyObjectScreenshotOptions] `json:"screenshotOptions"`
	ScrollPage           param.Field[bool]                                           `json:"scrollPage"`
	Selector             param.Field[string]                                         `json:"selector"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                              `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                           `json:"setJavaScriptEnabled"`
	UserAgent            param.Field[string]                                         `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[ScreenshotNewParamsBodyObjectViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[ScreenshotNewParamsBodyObjectWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r ScreenshotNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScreenshotNewParamsBodyObject) implementsScreenshotNewParamsBodyUnion() {}

type ScreenshotNewParamsBodyObjectAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r ScreenshotNewParamsBodyObjectAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsBodyObjectAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r ScreenshotNewParamsBodyObjectAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsBodyObjectAllowResourceType string

const (
	ScreenshotNewParamsBodyObjectAllowResourceTypeDocument           ScreenshotNewParamsBodyObjectAllowResourceType = "document"
	ScreenshotNewParamsBodyObjectAllowResourceTypeStylesheet         ScreenshotNewParamsBodyObjectAllowResourceType = "stylesheet"
	ScreenshotNewParamsBodyObjectAllowResourceTypeImage              ScreenshotNewParamsBodyObjectAllowResourceType = "image"
	ScreenshotNewParamsBodyObjectAllowResourceTypeMedia              ScreenshotNewParamsBodyObjectAllowResourceType = "media"
	ScreenshotNewParamsBodyObjectAllowResourceTypeFont               ScreenshotNewParamsBodyObjectAllowResourceType = "font"
	ScreenshotNewParamsBodyObjectAllowResourceTypeScript             ScreenshotNewParamsBodyObjectAllowResourceType = "script"
	ScreenshotNewParamsBodyObjectAllowResourceTypeTexttrack          ScreenshotNewParamsBodyObjectAllowResourceType = "texttrack"
	ScreenshotNewParamsBodyObjectAllowResourceTypeXHR                ScreenshotNewParamsBodyObjectAllowResourceType = "xhr"
	ScreenshotNewParamsBodyObjectAllowResourceTypeFetch              ScreenshotNewParamsBodyObjectAllowResourceType = "fetch"
	ScreenshotNewParamsBodyObjectAllowResourceTypePrefetch           ScreenshotNewParamsBodyObjectAllowResourceType = "prefetch"
	ScreenshotNewParamsBodyObjectAllowResourceTypeEventsource        ScreenshotNewParamsBodyObjectAllowResourceType = "eventsource"
	ScreenshotNewParamsBodyObjectAllowResourceTypeWebsocket          ScreenshotNewParamsBodyObjectAllowResourceType = "websocket"
	ScreenshotNewParamsBodyObjectAllowResourceTypeManifest           ScreenshotNewParamsBodyObjectAllowResourceType = "manifest"
	ScreenshotNewParamsBodyObjectAllowResourceTypeSignedexchange     ScreenshotNewParamsBodyObjectAllowResourceType = "signedexchange"
	ScreenshotNewParamsBodyObjectAllowResourceTypePing               ScreenshotNewParamsBodyObjectAllowResourceType = "ping"
	ScreenshotNewParamsBodyObjectAllowResourceTypeCspviolationreport ScreenshotNewParamsBodyObjectAllowResourceType = "cspviolationreport"
	ScreenshotNewParamsBodyObjectAllowResourceTypePreflight          ScreenshotNewParamsBodyObjectAllowResourceType = "preflight"
	ScreenshotNewParamsBodyObjectAllowResourceTypeOther              ScreenshotNewParamsBodyObjectAllowResourceType = "other"
)

func (r ScreenshotNewParamsBodyObjectAllowResourceType) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsBodyObjectAllowResourceTypeDocument, ScreenshotNewParamsBodyObjectAllowResourceTypeStylesheet, ScreenshotNewParamsBodyObjectAllowResourceTypeImage, ScreenshotNewParamsBodyObjectAllowResourceTypeMedia, ScreenshotNewParamsBodyObjectAllowResourceTypeFont, ScreenshotNewParamsBodyObjectAllowResourceTypeScript, ScreenshotNewParamsBodyObjectAllowResourceTypeTexttrack, ScreenshotNewParamsBodyObjectAllowResourceTypeXHR, ScreenshotNewParamsBodyObjectAllowResourceTypeFetch, ScreenshotNewParamsBodyObjectAllowResourceTypePrefetch, ScreenshotNewParamsBodyObjectAllowResourceTypeEventsource, ScreenshotNewParamsBodyObjectAllowResourceTypeWebsocket, ScreenshotNewParamsBodyObjectAllowResourceTypeManifest, ScreenshotNewParamsBodyObjectAllowResourceTypeSignedexchange, ScreenshotNewParamsBodyObjectAllowResourceTypePing, ScreenshotNewParamsBodyObjectAllowResourceTypeCspviolationreport, ScreenshotNewParamsBodyObjectAllowResourceTypePreflight, ScreenshotNewParamsBodyObjectAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type ScreenshotNewParamsBodyObjectAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r ScreenshotNewParamsBodyObjectAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsBodyObjectCookie struct {
	Name         param.Field[string]                                           `json:"name,required"`
	Value        param.Field[string]                                           `json:"value,required"`
	Domain       param.Field[string]                                           `json:"domain"`
	Expires      param.Field[float64]                                          `json:"expires"`
	HTTPOnly     param.Field[bool]                                             `json:"httpOnly"`
	PartitionKey param.Field[string]                                           `json:"partitionKey"`
	Path         param.Field[string]                                           `json:"path"`
	Priority     param.Field[ScreenshotNewParamsBodyObjectCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                             `json:"sameParty"`
	SameSite     param.Field[ScreenshotNewParamsBodyObjectCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                             `json:"secure"`
	SourcePort   param.Field[float64]                                          `json:"sourcePort"`
	SourceScheme param.Field[ScreenshotNewParamsBodyObjectCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                           `json:"url"`
}

func (r ScreenshotNewParamsBodyObjectCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsBodyObjectCookiesPriority string

const (
	ScreenshotNewParamsBodyObjectCookiesPriorityLow    ScreenshotNewParamsBodyObjectCookiesPriority = "Low"
	ScreenshotNewParamsBodyObjectCookiesPriorityMedium ScreenshotNewParamsBodyObjectCookiesPriority = "Medium"
	ScreenshotNewParamsBodyObjectCookiesPriorityHigh   ScreenshotNewParamsBodyObjectCookiesPriority = "High"
)

func (r ScreenshotNewParamsBodyObjectCookiesPriority) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsBodyObjectCookiesPriorityLow, ScreenshotNewParamsBodyObjectCookiesPriorityMedium, ScreenshotNewParamsBodyObjectCookiesPriorityHigh:
		return true
	}
	return false
}

type ScreenshotNewParamsBodyObjectCookiesSameSite string

const (
	ScreenshotNewParamsBodyObjectCookiesSameSiteStrict ScreenshotNewParamsBodyObjectCookiesSameSite = "Strict"
	ScreenshotNewParamsBodyObjectCookiesSameSiteLax    ScreenshotNewParamsBodyObjectCookiesSameSite = "Lax"
	ScreenshotNewParamsBodyObjectCookiesSameSiteNone   ScreenshotNewParamsBodyObjectCookiesSameSite = "None"
)

func (r ScreenshotNewParamsBodyObjectCookiesSameSite) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsBodyObjectCookiesSameSiteStrict, ScreenshotNewParamsBodyObjectCookiesSameSiteLax, ScreenshotNewParamsBodyObjectCookiesSameSiteNone:
		return true
	}
	return false
}

type ScreenshotNewParamsBodyObjectCookiesSourceScheme string

const (
	ScreenshotNewParamsBodyObjectCookiesSourceSchemeUnset     ScreenshotNewParamsBodyObjectCookiesSourceScheme = "Unset"
	ScreenshotNewParamsBodyObjectCookiesSourceSchemeNonSecure ScreenshotNewParamsBodyObjectCookiesSourceScheme = "NonSecure"
	ScreenshotNewParamsBodyObjectCookiesSourceSchemeSecure    ScreenshotNewParamsBodyObjectCookiesSourceScheme = "Secure"
)

func (r ScreenshotNewParamsBodyObjectCookiesSourceScheme) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsBodyObjectCookiesSourceSchemeUnset, ScreenshotNewParamsBodyObjectCookiesSourceSchemeNonSecure, ScreenshotNewParamsBodyObjectCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type ScreenshotNewParamsBodyObjectGotoOptions struct {
	Referer        param.Field[string]                                                 `json:"referer"`
	ReferrerPolicy param.Field[string]                                                 `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                `json:"timeout"`
	WaitUntil      param.Field[ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r ScreenshotNewParamsBodyObjectGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilString],
// [browser_rendering.ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArray].
type ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilUnion interface {
	implementsScreenshotNewParamsBodyObjectGotoOptionsWaitUntilUnion()
}

type ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilString string

const (
	ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilStringLoad             ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilString = "load"
	ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilString = "domcontentloaded"
	ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0     ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle0"
	ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2     ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle2"
)

func (r ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilStringLoad, ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded, ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0, ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilString) implementsScreenshotNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArray []ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem

func (r ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArray) implementsScreenshotNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem string

const (
	ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad             ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "load"
	ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0     ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle0"
	ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2     ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad, ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded, ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0, ScreenshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type ScreenshotNewParamsBodyObjectRejectResourceType string

const (
	ScreenshotNewParamsBodyObjectRejectResourceTypeDocument           ScreenshotNewParamsBodyObjectRejectResourceType = "document"
	ScreenshotNewParamsBodyObjectRejectResourceTypeStylesheet         ScreenshotNewParamsBodyObjectRejectResourceType = "stylesheet"
	ScreenshotNewParamsBodyObjectRejectResourceTypeImage              ScreenshotNewParamsBodyObjectRejectResourceType = "image"
	ScreenshotNewParamsBodyObjectRejectResourceTypeMedia              ScreenshotNewParamsBodyObjectRejectResourceType = "media"
	ScreenshotNewParamsBodyObjectRejectResourceTypeFont               ScreenshotNewParamsBodyObjectRejectResourceType = "font"
	ScreenshotNewParamsBodyObjectRejectResourceTypeScript             ScreenshotNewParamsBodyObjectRejectResourceType = "script"
	ScreenshotNewParamsBodyObjectRejectResourceTypeTexttrack          ScreenshotNewParamsBodyObjectRejectResourceType = "texttrack"
	ScreenshotNewParamsBodyObjectRejectResourceTypeXHR                ScreenshotNewParamsBodyObjectRejectResourceType = "xhr"
	ScreenshotNewParamsBodyObjectRejectResourceTypeFetch              ScreenshotNewParamsBodyObjectRejectResourceType = "fetch"
	ScreenshotNewParamsBodyObjectRejectResourceTypePrefetch           ScreenshotNewParamsBodyObjectRejectResourceType = "prefetch"
	ScreenshotNewParamsBodyObjectRejectResourceTypeEventsource        ScreenshotNewParamsBodyObjectRejectResourceType = "eventsource"
	ScreenshotNewParamsBodyObjectRejectResourceTypeWebsocket          ScreenshotNewParamsBodyObjectRejectResourceType = "websocket"
	ScreenshotNewParamsBodyObjectRejectResourceTypeManifest           ScreenshotNewParamsBodyObjectRejectResourceType = "manifest"
	ScreenshotNewParamsBodyObjectRejectResourceTypeSignedexchange     ScreenshotNewParamsBodyObjectRejectResourceType = "signedexchange"
	ScreenshotNewParamsBodyObjectRejectResourceTypePing               ScreenshotNewParamsBodyObjectRejectResourceType = "ping"
	ScreenshotNewParamsBodyObjectRejectResourceTypeCspviolationreport ScreenshotNewParamsBodyObjectRejectResourceType = "cspviolationreport"
	ScreenshotNewParamsBodyObjectRejectResourceTypePreflight          ScreenshotNewParamsBodyObjectRejectResourceType = "preflight"
	ScreenshotNewParamsBodyObjectRejectResourceTypeOther              ScreenshotNewParamsBodyObjectRejectResourceType = "other"
)

func (r ScreenshotNewParamsBodyObjectRejectResourceType) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsBodyObjectRejectResourceTypeDocument, ScreenshotNewParamsBodyObjectRejectResourceTypeStylesheet, ScreenshotNewParamsBodyObjectRejectResourceTypeImage, ScreenshotNewParamsBodyObjectRejectResourceTypeMedia, ScreenshotNewParamsBodyObjectRejectResourceTypeFont, ScreenshotNewParamsBodyObjectRejectResourceTypeScript, ScreenshotNewParamsBodyObjectRejectResourceTypeTexttrack, ScreenshotNewParamsBodyObjectRejectResourceTypeXHR, ScreenshotNewParamsBodyObjectRejectResourceTypeFetch, ScreenshotNewParamsBodyObjectRejectResourceTypePrefetch, ScreenshotNewParamsBodyObjectRejectResourceTypeEventsource, ScreenshotNewParamsBodyObjectRejectResourceTypeWebsocket, ScreenshotNewParamsBodyObjectRejectResourceTypeManifest, ScreenshotNewParamsBodyObjectRejectResourceTypeSignedexchange, ScreenshotNewParamsBodyObjectRejectResourceTypePing, ScreenshotNewParamsBodyObjectRejectResourceTypeCspviolationreport, ScreenshotNewParamsBodyObjectRejectResourceTypePreflight, ScreenshotNewParamsBodyObjectRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.screenshotoptions).
type ScreenshotNewParamsBodyObjectScreenshotOptions struct {
	CaptureBeyondViewport param.Field[bool]                                                   `json:"captureBeyondViewport"`
	Clip                  param.Field[ScreenshotNewParamsBodyObjectScreenshotOptionsClip]     `json:"clip"`
	Encoding              param.Field[ScreenshotNewParamsBodyObjectScreenshotOptionsEncoding] `json:"encoding"`
	FromSurface           param.Field[bool]                                                   `json:"fromSurface"`
	FullPage              param.Field[bool]                                                   `json:"fullPage"`
	OmitBackground        param.Field[bool]                                                   `json:"omitBackground"`
	OptimizeForSpeed      param.Field[bool]                                                   `json:"optimizeForSpeed"`
	Quality               param.Field[float64]                                                `json:"quality"`
	Type                  param.Field[ScreenshotNewParamsBodyObjectScreenshotOptionsType]     `json:"type"`
}

func (r ScreenshotNewParamsBodyObjectScreenshotOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsBodyObjectScreenshotOptionsClip struct {
	Height param.Field[float64] `json:"height,required"`
	Width  param.Field[float64] `json:"width,required"`
	X      param.Field[float64] `json:"x,required"`
	Y      param.Field[float64] `json:"y,required"`
	Scale  param.Field[float64] `json:"scale"`
}

func (r ScreenshotNewParamsBodyObjectScreenshotOptionsClip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsBodyObjectScreenshotOptionsEncoding string

const (
	ScreenshotNewParamsBodyObjectScreenshotOptionsEncodingBinary ScreenshotNewParamsBodyObjectScreenshotOptionsEncoding = "binary"
	ScreenshotNewParamsBodyObjectScreenshotOptionsEncodingBase64 ScreenshotNewParamsBodyObjectScreenshotOptionsEncoding = "base64"
)

func (r ScreenshotNewParamsBodyObjectScreenshotOptionsEncoding) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsBodyObjectScreenshotOptionsEncodingBinary, ScreenshotNewParamsBodyObjectScreenshotOptionsEncodingBase64:
		return true
	}
	return false
}

type ScreenshotNewParamsBodyObjectScreenshotOptionsType string

const (
	ScreenshotNewParamsBodyObjectScreenshotOptionsTypePNG  ScreenshotNewParamsBodyObjectScreenshotOptionsType = "png"
	ScreenshotNewParamsBodyObjectScreenshotOptionsTypeJPEG ScreenshotNewParamsBodyObjectScreenshotOptionsType = "jpeg"
	ScreenshotNewParamsBodyObjectScreenshotOptionsTypeWebP ScreenshotNewParamsBodyObjectScreenshotOptionsType = "webp"
)

func (r ScreenshotNewParamsBodyObjectScreenshotOptionsType) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsBodyObjectScreenshotOptionsTypePNG, ScreenshotNewParamsBodyObjectScreenshotOptionsTypeJPEG, ScreenshotNewParamsBodyObjectScreenshotOptionsTypeWebP:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type ScreenshotNewParamsBodyObjectViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r ScreenshotNewParamsBodyObjectViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type ScreenshotNewParamsBodyObjectWaitForSelector struct {
	Selector param.Field[string]                                              `json:"selector,required"`
	Hidden   param.Field[ScreenshotNewParamsBodyObjectWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                             `json:"timeout"`
	Visible  param.Field[ScreenshotNewParamsBodyObjectWaitForSelectorVisible] `json:"visible"`
}

func (r ScreenshotNewParamsBodyObjectWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsBodyObjectWaitForSelectorHidden bool

const (
	ScreenshotNewParamsBodyObjectWaitForSelectorHiddenTrue ScreenshotNewParamsBodyObjectWaitForSelectorHidden = true
)

func (r ScreenshotNewParamsBodyObjectWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsBodyObjectWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type ScreenshotNewParamsBodyObjectWaitForSelectorVisible bool

const (
	ScreenshotNewParamsBodyObjectWaitForSelectorVisibleTrue ScreenshotNewParamsBodyObjectWaitForSelectorVisible = true
)

func (r ScreenshotNewParamsBodyObjectWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsBodyObjectWaitForSelectorVisibleTrue:
		return true
	}
	return false
}
