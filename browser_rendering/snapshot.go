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

// SnapshotService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSnapshotService] method instead.
type SnapshotService struct {
	Options []option.RequestOption
}

// NewSnapshotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSnapshotService(opts ...option.RequestOption) (r *SnapshotService) {
	r = &SnapshotService{}
	r.Options = opts
	return
}

// Returns the page's HTML content and screenshot. Control page loading with
// `gotoOptions` and `waitFor*` options. Customize screenshots with `viewport`,
// `fullPage`, `clip` and others.
func (r *SnapshotService) New(ctx context.Context, params SnapshotNewParams, opts ...option.RequestOption) (res *SnapshotNewResponse, err error) {
	var env SnapshotNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/snapshot", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SnapshotNewResponse struct {
	// HTML content
	Content string `json:"content,required"`
	// Base64 encoded image
	Screenshot string                  `json:"screenshot,required"`
	JSON       snapshotNewResponseJSON `json:"-"`
}

// snapshotNewResponseJSON contains the JSON metadata for the struct
// [SnapshotNewResponse]
type snapshotNewResponseJSON struct {
	Content     apijson.Field
	Screenshot  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnapshotNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snapshotNewResponseJSON) RawJSON() string {
	return r.raw
}

type SnapshotNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64]       `query:"cacheTTL"`
	Body     SnapshotNewParamsBodyUnion `json:"body"`
}

func (r SnapshotNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [SnapshotNewParams]'s query parameters as `url.Values`.
func (r SnapshotNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SnapshotNewParamsBody struct {
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

func (r SnapshotNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SnapshotNewParamsBody) implementsSnapshotNewParamsBodyUnion() {}

// Satisfied by [browser_rendering.SnapshotNewParamsBodyObject],
// [browser_rendering.SnapshotNewParamsBodyObject], [SnapshotNewParamsBody].
type SnapshotNewParamsBodyUnion interface {
	implementsSnapshotNewParamsBodyUnion()
}

type SnapshotNewParamsBodyObject struct {
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html,required"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]SnapshotNewParamsBodyObjectAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]SnapshotNewParamsBodyObjectAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]SnapshotNewParamsBodyObjectAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[SnapshotNewParamsBodyObjectAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]SnapshotNewParamsBodyObjectCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                              `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[SnapshotNewParamsBodyObjectGotoOptions] `json:"gotoOptions"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]SnapshotNewParamsBodyObjectRejectResourceType] `json:"rejectResourceTypes"`
	ScreenshotOptions    param.Field[SnapshotNewParamsBodyObjectScreenshotOptions]    `json:"screenshotOptions"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                               `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                            `json:"setJavaScriptEnabled"`
	UserAgent            param.Field[string]                                          `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[SnapshotNewParamsBodyObjectViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[SnapshotNewParamsBodyObjectWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r SnapshotNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SnapshotNewParamsBodyObject) implementsSnapshotNewParamsBodyUnion() {}

type SnapshotNewParamsBodyObjectAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r SnapshotNewParamsBodyObjectAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsBodyObjectAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r SnapshotNewParamsBodyObjectAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsBodyObjectAllowResourceType string

const (
	SnapshotNewParamsBodyObjectAllowResourceTypeDocument           SnapshotNewParamsBodyObjectAllowResourceType = "document"
	SnapshotNewParamsBodyObjectAllowResourceTypeStylesheet         SnapshotNewParamsBodyObjectAllowResourceType = "stylesheet"
	SnapshotNewParamsBodyObjectAllowResourceTypeImage              SnapshotNewParamsBodyObjectAllowResourceType = "image"
	SnapshotNewParamsBodyObjectAllowResourceTypeMedia              SnapshotNewParamsBodyObjectAllowResourceType = "media"
	SnapshotNewParamsBodyObjectAllowResourceTypeFont               SnapshotNewParamsBodyObjectAllowResourceType = "font"
	SnapshotNewParamsBodyObjectAllowResourceTypeScript             SnapshotNewParamsBodyObjectAllowResourceType = "script"
	SnapshotNewParamsBodyObjectAllowResourceTypeTexttrack          SnapshotNewParamsBodyObjectAllowResourceType = "texttrack"
	SnapshotNewParamsBodyObjectAllowResourceTypeXHR                SnapshotNewParamsBodyObjectAllowResourceType = "xhr"
	SnapshotNewParamsBodyObjectAllowResourceTypeFetch              SnapshotNewParamsBodyObjectAllowResourceType = "fetch"
	SnapshotNewParamsBodyObjectAllowResourceTypePrefetch           SnapshotNewParamsBodyObjectAllowResourceType = "prefetch"
	SnapshotNewParamsBodyObjectAllowResourceTypeEventsource        SnapshotNewParamsBodyObjectAllowResourceType = "eventsource"
	SnapshotNewParamsBodyObjectAllowResourceTypeWebsocket          SnapshotNewParamsBodyObjectAllowResourceType = "websocket"
	SnapshotNewParamsBodyObjectAllowResourceTypeManifest           SnapshotNewParamsBodyObjectAllowResourceType = "manifest"
	SnapshotNewParamsBodyObjectAllowResourceTypeSignedexchange     SnapshotNewParamsBodyObjectAllowResourceType = "signedexchange"
	SnapshotNewParamsBodyObjectAllowResourceTypePing               SnapshotNewParamsBodyObjectAllowResourceType = "ping"
	SnapshotNewParamsBodyObjectAllowResourceTypeCspviolationreport SnapshotNewParamsBodyObjectAllowResourceType = "cspviolationreport"
	SnapshotNewParamsBodyObjectAllowResourceTypePreflight          SnapshotNewParamsBodyObjectAllowResourceType = "preflight"
	SnapshotNewParamsBodyObjectAllowResourceTypeOther              SnapshotNewParamsBodyObjectAllowResourceType = "other"
)

func (r SnapshotNewParamsBodyObjectAllowResourceType) IsKnown() bool {
	switch r {
	case SnapshotNewParamsBodyObjectAllowResourceTypeDocument, SnapshotNewParamsBodyObjectAllowResourceTypeStylesheet, SnapshotNewParamsBodyObjectAllowResourceTypeImage, SnapshotNewParamsBodyObjectAllowResourceTypeMedia, SnapshotNewParamsBodyObjectAllowResourceTypeFont, SnapshotNewParamsBodyObjectAllowResourceTypeScript, SnapshotNewParamsBodyObjectAllowResourceTypeTexttrack, SnapshotNewParamsBodyObjectAllowResourceTypeXHR, SnapshotNewParamsBodyObjectAllowResourceTypeFetch, SnapshotNewParamsBodyObjectAllowResourceTypePrefetch, SnapshotNewParamsBodyObjectAllowResourceTypeEventsource, SnapshotNewParamsBodyObjectAllowResourceTypeWebsocket, SnapshotNewParamsBodyObjectAllowResourceTypeManifest, SnapshotNewParamsBodyObjectAllowResourceTypeSignedexchange, SnapshotNewParamsBodyObjectAllowResourceTypePing, SnapshotNewParamsBodyObjectAllowResourceTypeCspviolationreport, SnapshotNewParamsBodyObjectAllowResourceTypePreflight, SnapshotNewParamsBodyObjectAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type SnapshotNewParamsBodyObjectAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r SnapshotNewParamsBodyObjectAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsBodyObjectCookie struct {
	Name         param.Field[string]                                         `json:"name,required"`
	Value        param.Field[string]                                         `json:"value,required"`
	Domain       param.Field[string]                                         `json:"domain"`
	Expires      param.Field[float64]                                        `json:"expires"`
	HTTPOnly     param.Field[bool]                                           `json:"httpOnly"`
	PartitionKey param.Field[string]                                         `json:"partitionKey"`
	Path         param.Field[string]                                         `json:"path"`
	Priority     param.Field[SnapshotNewParamsBodyObjectCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                           `json:"sameParty"`
	SameSite     param.Field[SnapshotNewParamsBodyObjectCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                           `json:"secure"`
	SourcePort   param.Field[float64]                                        `json:"sourcePort"`
	SourceScheme param.Field[SnapshotNewParamsBodyObjectCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                         `json:"url"`
}

func (r SnapshotNewParamsBodyObjectCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsBodyObjectCookiesPriority string

const (
	SnapshotNewParamsBodyObjectCookiesPriorityLow    SnapshotNewParamsBodyObjectCookiesPriority = "Low"
	SnapshotNewParamsBodyObjectCookiesPriorityMedium SnapshotNewParamsBodyObjectCookiesPriority = "Medium"
	SnapshotNewParamsBodyObjectCookiesPriorityHigh   SnapshotNewParamsBodyObjectCookiesPriority = "High"
)

func (r SnapshotNewParamsBodyObjectCookiesPriority) IsKnown() bool {
	switch r {
	case SnapshotNewParamsBodyObjectCookiesPriorityLow, SnapshotNewParamsBodyObjectCookiesPriorityMedium, SnapshotNewParamsBodyObjectCookiesPriorityHigh:
		return true
	}
	return false
}

type SnapshotNewParamsBodyObjectCookiesSameSite string

const (
	SnapshotNewParamsBodyObjectCookiesSameSiteStrict SnapshotNewParamsBodyObjectCookiesSameSite = "Strict"
	SnapshotNewParamsBodyObjectCookiesSameSiteLax    SnapshotNewParamsBodyObjectCookiesSameSite = "Lax"
	SnapshotNewParamsBodyObjectCookiesSameSiteNone   SnapshotNewParamsBodyObjectCookiesSameSite = "None"
)

func (r SnapshotNewParamsBodyObjectCookiesSameSite) IsKnown() bool {
	switch r {
	case SnapshotNewParamsBodyObjectCookiesSameSiteStrict, SnapshotNewParamsBodyObjectCookiesSameSiteLax, SnapshotNewParamsBodyObjectCookiesSameSiteNone:
		return true
	}
	return false
}

type SnapshotNewParamsBodyObjectCookiesSourceScheme string

const (
	SnapshotNewParamsBodyObjectCookiesSourceSchemeUnset     SnapshotNewParamsBodyObjectCookiesSourceScheme = "Unset"
	SnapshotNewParamsBodyObjectCookiesSourceSchemeNonSecure SnapshotNewParamsBodyObjectCookiesSourceScheme = "NonSecure"
	SnapshotNewParamsBodyObjectCookiesSourceSchemeSecure    SnapshotNewParamsBodyObjectCookiesSourceScheme = "Secure"
)

func (r SnapshotNewParamsBodyObjectCookiesSourceScheme) IsKnown() bool {
	switch r {
	case SnapshotNewParamsBodyObjectCookiesSourceSchemeUnset, SnapshotNewParamsBodyObjectCookiesSourceSchemeNonSecure, SnapshotNewParamsBodyObjectCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type SnapshotNewParamsBodyObjectGotoOptions struct {
	Referer        param.Field[string]                                               `json:"referer"`
	ReferrerPolicy param.Field[string]                                               `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                              `json:"timeout"`
	WaitUntil      param.Field[SnapshotNewParamsBodyObjectGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r SnapshotNewParamsBodyObjectGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.SnapshotNewParamsBodyObjectGotoOptionsWaitUntilString],
// [browser_rendering.SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArray].
type SnapshotNewParamsBodyObjectGotoOptionsWaitUntilUnion interface {
	implementsSnapshotNewParamsBodyObjectGotoOptionsWaitUntilUnion()
}

type SnapshotNewParamsBodyObjectGotoOptionsWaitUntilString string

const (
	SnapshotNewParamsBodyObjectGotoOptionsWaitUntilStringLoad             SnapshotNewParamsBodyObjectGotoOptionsWaitUntilString = "load"
	SnapshotNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded SnapshotNewParamsBodyObjectGotoOptionsWaitUntilString = "domcontentloaded"
	SnapshotNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0     SnapshotNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle0"
	SnapshotNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2     SnapshotNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle2"
)

func (r SnapshotNewParamsBodyObjectGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case SnapshotNewParamsBodyObjectGotoOptionsWaitUntilStringLoad, SnapshotNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded, SnapshotNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0, SnapshotNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r SnapshotNewParamsBodyObjectGotoOptionsWaitUntilString) implementsSnapshotNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArray []SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem

func (r SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArray) implementsSnapshotNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem string

const (
	SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad             SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "load"
	SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0     SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle0"
	SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2     SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad, SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded, SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0, SnapshotNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type SnapshotNewParamsBodyObjectRejectResourceType string

const (
	SnapshotNewParamsBodyObjectRejectResourceTypeDocument           SnapshotNewParamsBodyObjectRejectResourceType = "document"
	SnapshotNewParamsBodyObjectRejectResourceTypeStylesheet         SnapshotNewParamsBodyObjectRejectResourceType = "stylesheet"
	SnapshotNewParamsBodyObjectRejectResourceTypeImage              SnapshotNewParamsBodyObjectRejectResourceType = "image"
	SnapshotNewParamsBodyObjectRejectResourceTypeMedia              SnapshotNewParamsBodyObjectRejectResourceType = "media"
	SnapshotNewParamsBodyObjectRejectResourceTypeFont               SnapshotNewParamsBodyObjectRejectResourceType = "font"
	SnapshotNewParamsBodyObjectRejectResourceTypeScript             SnapshotNewParamsBodyObjectRejectResourceType = "script"
	SnapshotNewParamsBodyObjectRejectResourceTypeTexttrack          SnapshotNewParamsBodyObjectRejectResourceType = "texttrack"
	SnapshotNewParamsBodyObjectRejectResourceTypeXHR                SnapshotNewParamsBodyObjectRejectResourceType = "xhr"
	SnapshotNewParamsBodyObjectRejectResourceTypeFetch              SnapshotNewParamsBodyObjectRejectResourceType = "fetch"
	SnapshotNewParamsBodyObjectRejectResourceTypePrefetch           SnapshotNewParamsBodyObjectRejectResourceType = "prefetch"
	SnapshotNewParamsBodyObjectRejectResourceTypeEventsource        SnapshotNewParamsBodyObjectRejectResourceType = "eventsource"
	SnapshotNewParamsBodyObjectRejectResourceTypeWebsocket          SnapshotNewParamsBodyObjectRejectResourceType = "websocket"
	SnapshotNewParamsBodyObjectRejectResourceTypeManifest           SnapshotNewParamsBodyObjectRejectResourceType = "manifest"
	SnapshotNewParamsBodyObjectRejectResourceTypeSignedexchange     SnapshotNewParamsBodyObjectRejectResourceType = "signedexchange"
	SnapshotNewParamsBodyObjectRejectResourceTypePing               SnapshotNewParamsBodyObjectRejectResourceType = "ping"
	SnapshotNewParamsBodyObjectRejectResourceTypeCspviolationreport SnapshotNewParamsBodyObjectRejectResourceType = "cspviolationreport"
	SnapshotNewParamsBodyObjectRejectResourceTypePreflight          SnapshotNewParamsBodyObjectRejectResourceType = "preflight"
	SnapshotNewParamsBodyObjectRejectResourceTypeOther              SnapshotNewParamsBodyObjectRejectResourceType = "other"
)

func (r SnapshotNewParamsBodyObjectRejectResourceType) IsKnown() bool {
	switch r {
	case SnapshotNewParamsBodyObjectRejectResourceTypeDocument, SnapshotNewParamsBodyObjectRejectResourceTypeStylesheet, SnapshotNewParamsBodyObjectRejectResourceTypeImage, SnapshotNewParamsBodyObjectRejectResourceTypeMedia, SnapshotNewParamsBodyObjectRejectResourceTypeFont, SnapshotNewParamsBodyObjectRejectResourceTypeScript, SnapshotNewParamsBodyObjectRejectResourceTypeTexttrack, SnapshotNewParamsBodyObjectRejectResourceTypeXHR, SnapshotNewParamsBodyObjectRejectResourceTypeFetch, SnapshotNewParamsBodyObjectRejectResourceTypePrefetch, SnapshotNewParamsBodyObjectRejectResourceTypeEventsource, SnapshotNewParamsBodyObjectRejectResourceTypeWebsocket, SnapshotNewParamsBodyObjectRejectResourceTypeManifest, SnapshotNewParamsBodyObjectRejectResourceTypeSignedexchange, SnapshotNewParamsBodyObjectRejectResourceTypePing, SnapshotNewParamsBodyObjectRejectResourceTypeCspviolationreport, SnapshotNewParamsBodyObjectRejectResourceTypePreflight, SnapshotNewParamsBodyObjectRejectResourceTypeOther:
		return true
	}
	return false
}

type SnapshotNewParamsBodyObjectScreenshotOptions struct {
	CaptureBeyondViewport param.Field[bool]                                             `json:"captureBeyondViewport"`
	Clip                  param.Field[SnapshotNewParamsBodyObjectScreenshotOptionsClip] `json:"clip"`
	FromSurface           param.Field[bool]                                             `json:"fromSurface"`
	FullPage              param.Field[bool]                                             `json:"fullPage"`
	OmitBackground        param.Field[bool]                                             `json:"omitBackground"`
	OptimizeForSpeed      param.Field[bool]                                             `json:"optimizeForSpeed"`
	Quality               param.Field[float64]                                          `json:"quality"`
	Type                  param.Field[SnapshotNewParamsBodyObjectScreenshotOptionsType] `json:"type"`
}

func (r SnapshotNewParamsBodyObjectScreenshotOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsBodyObjectScreenshotOptionsClip struct {
	Height param.Field[float64] `json:"height,required"`
	Width  param.Field[float64] `json:"width,required"`
	X      param.Field[float64] `json:"x,required"`
	Y      param.Field[float64] `json:"y,required"`
	Scale  param.Field[float64] `json:"scale"`
}

func (r SnapshotNewParamsBodyObjectScreenshotOptionsClip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsBodyObjectScreenshotOptionsType string

const (
	SnapshotNewParamsBodyObjectScreenshotOptionsTypePNG  SnapshotNewParamsBodyObjectScreenshotOptionsType = "png"
	SnapshotNewParamsBodyObjectScreenshotOptionsTypeJPEG SnapshotNewParamsBodyObjectScreenshotOptionsType = "jpeg"
	SnapshotNewParamsBodyObjectScreenshotOptionsTypeWebP SnapshotNewParamsBodyObjectScreenshotOptionsType = "webp"
)

func (r SnapshotNewParamsBodyObjectScreenshotOptionsType) IsKnown() bool {
	switch r {
	case SnapshotNewParamsBodyObjectScreenshotOptionsTypePNG, SnapshotNewParamsBodyObjectScreenshotOptionsTypeJPEG, SnapshotNewParamsBodyObjectScreenshotOptionsTypeWebP:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type SnapshotNewParamsBodyObjectViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r SnapshotNewParamsBodyObjectViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type SnapshotNewParamsBodyObjectWaitForSelector struct {
	Selector param.Field[string]                                            `json:"selector,required"`
	Hidden   param.Field[SnapshotNewParamsBodyObjectWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                           `json:"timeout"`
	Visible  param.Field[SnapshotNewParamsBodyObjectWaitForSelectorVisible] `json:"visible"`
}

func (r SnapshotNewParamsBodyObjectWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsBodyObjectWaitForSelectorHidden bool

const (
	SnapshotNewParamsBodyObjectWaitForSelectorHiddenTrue SnapshotNewParamsBodyObjectWaitForSelectorHidden = true
)

func (r SnapshotNewParamsBodyObjectWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case SnapshotNewParamsBodyObjectWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type SnapshotNewParamsBodyObjectWaitForSelectorVisible bool

const (
	SnapshotNewParamsBodyObjectWaitForSelectorVisibleTrue SnapshotNewParamsBodyObjectWaitForSelectorVisible = true
)

func (r SnapshotNewParamsBodyObjectWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case SnapshotNewParamsBodyObjectWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type SnapshotNewResponseEnvelope struct {
	Meta SnapshotNewResponseEnvelopeMeta `json:"meta,required"`
	// Response status
	Success bool                                `json:"success,required"`
	Errors  []SnapshotNewResponseEnvelopeErrors `json:"errors"`
	Result  SnapshotNewResponse                 `json:"result"`
	JSON    snapshotNewResponseEnvelopeJSON     `json:"-"`
}

// snapshotNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnapshotNewResponseEnvelope]
type snapshotNewResponseEnvelopeJSON struct {
	Meta        apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnapshotNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snapshotNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SnapshotNewResponseEnvelopeMeta struct {
	Status float64                             `json:"status"`
	Title  string                              `json:"title"`
	JSON   snapshotNewResponseEnvelopeMetaJSON `json:"-"`
}

// snapshotNewResponseEnvelopeMetaJSON contains the JSON metadata for the struct
// [SnapshotNewResponseEnvelopeMeta]
type snapshotNewResponseEnvelopeMetaJSON struct {
	Status      apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnapshotNewResponseEnvelopeMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snapshotNewResponseEnvelopeMetaJSON) RawJSON() string {
	return r.raw
}

type SnapshotNewResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                `json:"message,required"`
	JSON    snapshotNewResponseEnvelopeErrorsJSON `json:"-"`
}

// snapshotNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SnapshotNewResponseEnvelopeErrors]
type snapshotNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnapshotNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snapshotNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
