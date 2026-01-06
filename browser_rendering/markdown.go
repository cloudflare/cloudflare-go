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

// MarkdownService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarkdownService] method instead.
type MarkdownService struct {
	Options []option.RequestOption
}

// NewMarkdownService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMarkdownService(opts ...option.RequestOption) (r *MarkdownService) {
	r = &MarkdownService{}
	r.Options = opts
	return
}

// Gets markdown of a webpage from provided URL or HTML. Control page loading with
// `gotoOptions` and `waitFor*` options.
func (r *MarkdownService) New(ctx context.Context, params MarkdownNewParams, opts ...option.RequestOption) (res *string, err error) {
	var env MarkdownNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/markdown", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MarkdownNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64]       `query:"cacheTTL"`
	Body     MarkdownNewParamsBodyUnion `json:"body"`
}

func (r MarkdownNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [MarkdownNewParams]'s query parameters as `url.Values`.
func (r MarkdownNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type MarkdownNewParamsBody struct {
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

func (r MarkdownNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MarkdownNewParamsBody) implementsMarkdownNewParamsBodyUnion() {}

// Satisfied by [browser_rendering.MarkdownNewParamsBodyObject],
// [browser_rendering.MarkdownNewParamsBodyObject], [MarkdownNewParamsBody].
type MarkdownNewParamsBodyUnion interface {
	implementsMarkdownNewParamsBodyUnion()
}

type MarkdownNewParamsBodyObject struct {
	// URL to navigate to, eg. `https://example.com`.
	URL param.Field[string] `json:"url,required" format:"uri"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]MarkdownNewParamsBodyObjectAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]MarkdownNewParamsBodyObjectAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]MarkdownNewParamsBodyObjectAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[MarkdownNewParamsBodyObjectAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]MarkdownNewParamsBodyObjectCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                              `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[MarkdownNewParamsBodyObjectGotoOptions] `json:"gotoOptions"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]MarkdownNewParamsBodyObjectRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                               `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                            `json:"setJavaScriptEnabled"`
	UserAgent            param.Field[string]                                          `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[MarkdownNewParamsBodyObjectViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[MarkdownNewParamsBodyObjectWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r MarkdownNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MarkdownNewParamsBodyObject) implementsMarkdownNewParamsBodyUnion() {}

type MarkdownNewParamsBodyObjectAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r MarkdownNewParamsBodyObjectAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarkdownNewParamsBodyObjectAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r MarkdownNewParamsBodyObjectAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarkdownNewParamsBodyObjectAllowResourceType string

const (
	MarkdownNewParamsBodyObjectAllowResourceTypeDocument           MarkdownNewParamsBodyObjectAllowResourceType = "document"
	MarkdownNewParamsBodyObjectAllowResourceTypeStylesheet         MarkdownNewParamsBodyObjectAllowResourceType = "stylesheet"
	MarkdownNewParamsBodyObjectAllowResourceTypeImage              MarkdownNewParamsBodyObjectAllowResourceType = "image"
	MarkdownNewParamsBodyObjectAllowResourceTypeMedia              MarkdownNewParamsBodyObjectAllowResourceType = "media"
	MarkdownNewParamsBodyObjectAllowResourceTypeFont               MarkdownNewParamsBodyObjectAllowResourceType = "font"
	MarkdownNewParamsBodyObjectAllowResourceTypeScript             MarkdownNewParamsBodyObjectAllowResourceType = "script"
	MarkdownNewParamsBodyObjectAllowResourceTypeTexttrack          MarkdownNewParamsBodyObjectAllowResourceType = "texttrack"
	MarkdownNewParamsBodyObjectAllowResourceTypeXHR                MarkdownNewParamsBodyObjectAllowResourceType = "xhr"
	MarkdownNewParamsBodyObjectAllowResourceTypeFetch              MarkdownNewParamsBodyObjectAllowResourceType = "fetch"
	MarkdownNewParamsBodyObjectAllowResourceTypePrefetch           MarkdownNewParamsBodyObjectAllowResourceType = "prefetch"
	MarkdownNewParamsBodyObjectAllowResourceTypeEventsource        MarkdownNewParamsBodyObjectAllowResourceType = "eventsource"
	MarkdownNewParamsBodyObjectAllowResourceTypeWebsocket          MarkdownNewParamsBodyObjectAllowResourceType = "websocket"
	MarkdownNewParamsBodyObjectAllowResourceTypeManifest           MarkdownNewParamsBodyObjectAllowResourceType = "manifest"
	MarkdownNewParamsBodyObjectAllowResourceTypeSignedexchange     MarkdownNewParamsBodyObjectAllowResourceType = "signedexchange"
	MarkdownNewParamsBodyObjectAllowResourceTypePing               MarkdownNewParamsBodyObjectAllowResourceType = "ping"
	MarkdownNewParamsBodyObjectAllowResourceTypeCspviolationreport MarkdownNewParamsBodyObjectAllowResourceType = "cspviolationreport"
	MarkdownNewParamsBodyObjectAllowResourceTypePreflight          MarkdownNewParamsBodyObjectAllowResourceType = "preflight"
	MarkdownNewParamsBodyObjectAllowResourceTypeOther              MarkdownNewParamsBodyObjectAllowResourceType = "other"
)

func (r MarkdownNewParamsBodyObjectAllowResourceType) IsKnown() bool {
	switch r {
	case MarkdownNewParamsBodyObjectAllowResourceTypeDocument, MarkdownNewParamsBodyObjectAllowResourceTypeStylesheet, MarkdownNewParamsBodyObjectAllowResourceTypeImage, MarkdownNewParamsBodyObjectAllowResourceTypeMedia, MarkdownNewParamsBodyObjectAllowResourceTypeFont, MarkdownNewParamsBodyObjectAllowResourceTypeScript, MarkdownNewParamsBodyObjectAllowResourceTypeTexttrack, MarkdownNewParamsBodyObjectAllowResourceTypeXHR, MarkdownNewParamsBodyObjectAllowResourceTypeFetch, MarkdownNewParamsBodyObjectAllowResourceTypePrefetch, MarkdownNewParamsBodyObjectAllowResourceTypeEventsource, MarkdownNewParamsBodyObjectAllowResourceTypeWebsocket, MarkdownNewParamsBodyObjectAllowResourceTypeManifest, MarkdownNewParamsBodyObjectAllowResourceTypeSignedexchange, MarkdownNewParamsBodyObjectAllowResourceTypePing, MarkdownNewParamsBodyObjectAllowResourceTypeCspviolationreport, MarkdownNewParamsBodyObjectAllowResourceTypePreflight, MarkdownNewParamsBodyObjectAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type MarkdownNewParamsBodyObjectAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r MarkdownNewParamsBodyObjectAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarkdownNewParamsBodyObjectCookie struct {
	Name         param.Field[string]                                         `json:"name,required"`
	Value        param.Field[string]                                         `json:"value,required"`
	Domain       param.Field[string]                                         `json:"domain"`
	Expires      param.Field[float64]                                        `json:"expires"`
	HTTPOnly     param.Field[bool]                                           `json:"httpOnly"`
	PartitionKey param.Field[string]                                         `json:"partitionKey"`
	Path         param.Field[string]                                         `json:"path"`
	Priority     param.Field[MarkdownNewParamsBodyObjectCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                           `json:"sameParty"`
	SameSite     param.Field[MarkdownNewParamsBodyObjectCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                           `json:"secure"`
	SourcePort   param.Field[float64]                                        `json:"sourcePort"`
	SourceScheme param.Field[MarkdownNewParamsBodyObjectCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                         `json:"url"`
}

func (r MarkdownNewParamsBodyObjectCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarkdownNewParamsBodyObjectCookiesPriority string

const (
	MarkdownNewParamsBodyObjectCookiesPriorityLow    MarkdownNewParamsBodyObjectCookiesPriority = "Low"
	MarkdownNewParamsBodyObjectCookiesPriorityMedium MarkdownNewParamsBodyObjectCookiesPriority = "Medium"
	MarkdownNewParamsBodyObjectCookiesPriorityHigh   MarkdownNewParamsBodyObjectCookiesPriority = "High"
)

func (r MarkdownNewParamsBodyObjectCookiesPriority) IsKnown() bool {
	switch r {
	case MarkdownNewParamsBodyObjectCookiesPriorityLow, MarkdownNewParamsBodyObjectCookiesPriorityMedium, MarkdownNewParamsBodyObjectCookiesPriorityHigh:
		return true
	}
	return false
}

type MarkdownNewParamsBodyObjectCookiesSameSite string

const (
	MarkdownNewParamsBodyObjectCookiesSameSiteStrict MarkdownNewParamsBodyObjectCookiesSameSite = "Strict"
	MarkdownNewParamsBodyObjectCookiesSameSiteLax    MarkdownNewParamsBodyObjectCookiesSameSite = "Lax"
	MarkdownNewParamsBodyObjectCookiesSameSiteNone   MarkdownNewParamsBodyObjectCookiesSameSite = "None"
)

func (r MarkdownNewParamsBodyObjectCookiesSameSite) IsKnown() bool {
	switch r {
	case MarkdownNewParamsBodyObjectCookiesSameSiteStrict, MarkdownNewParamsBodyObjectCookiesSameSiteLax, MarkdownNewParamsBodyObjectCookiesSameSiteNone:
		return true
	}
	return false
}

type MarkdownNewParamsBodyObjectCookiesSourceScheme string

const (
	MarkdownNewParamsBodyObjectCookiesSourceSchemeUnset     MarkdownNewParamsBodyObjectCookiesSourceScheme = "Unset"
	MarkdownNewParamsBodyObjectCookiesSourceSchemeNonSecure MarkdownNewParamsBodyObjectCookiesSourceScheme = "NonSecure"
	MarkdownNewParamsBodyObjectCookiesSourceSchemeSecure    MarkdownNewParamsBodyObjectCookiesSourceScheme = "Secure"
)

func (r MarkdownNewParamsBodyObjectCookiesSourceScheme) IsKnown() bool {
	switch r {
	case MarkdownNewParamsBodyObjectCookiesSourceSchemeUnset, MarkdownNewParamsBodyObjectCookiesSourceSchemeNonSecure, MarkdownNewParamsBodyObjectCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type MarkdownNewParamsBodyObjectGotoOptions struct {
	Referer        param.Field[string]                                               `json:"referer"`
	ReferrerPolicy param.Field[string]                                               `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                              `json:"timeout"`
	WaitUntil      param.Field[MarkdownNewParamsBodyObjectGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r MarkdownNewParamsBodyObjectGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.MarkdownNewParamsBodyObjectGotoOptionsWaitUntilString],
// [browser_rendering.MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArray].
type MarkdownNewParamsBodyObjectGotoOptionsWaitUntilUnion interface {
	implementsMarkdownNewParamsBodyObjectGotoOptionsWaitUntilUnion()
}

type MarkdownNewParamsBodyObjectGotoOptionsWaitUntilString string

const (
	MarkdownNewParamsBodyObjectGotoOptionsWaitUntilStringLoad             MarkdownNewParamsBodyObjectGotoOptionsWaitUntilString = "load"
	MarkdownNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded MarkdownNewParamsBodyObjectGotoOptionsWaitUntilString = "domcontentloaded"
	MarkdownNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0     MarkdownNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle0"
	MarkdownNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2     MarkdownNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle2"
)

func (r MarkdownNewParamsBodyObjectGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case MarkdownNewParamsBodyObjectGotoOptionsWaitUntilStringLoad, MarkdownNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded, MarkdownNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0, MarkdownNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r MarkdownNewParamsBodyObjectGotoOptionsWaitUntilString) implementsMarkdownNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArray []MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItem

func (r MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArray) implementsMarkdownNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItem string

const (
	MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad             MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "load"
	MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0     MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle0"
	MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2     MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad, MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded, MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0, MarkdownNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type MarkdownNewParamsBodyObjectRejectResourceType string

const (
	MarkdownNewParamsBodyObjectRejectResourceTypeDocument           MarkdownNewParamsBodyObjectRejectResourceType = "document"
	MarkdownNewParamsBodyObjectRejectResourceTypeStylesheet         MarkdownNewParamsBodyObjectRejectResourceType = "stylesheet"
	MarkdownNewParamsBodyObjectRejectResourceTypeImage              MarkdownNewParamsBodyObjectRejectResourceType = "image"
	MarkdownNewParamsBodyObjectRejectResourceTypeMedia              MarkdownNewParamsBodyObjectRejectResourceType = "media"
	MarkdownNewParamsBodyObjectRejectResourceTypeFont               MarkdownNewParamsBodyObjectRejectResourceType = "font"
	MarkdownNewParamsBodyObjectRejectResourceTypeScript             MarkdownNewParamsBodyObjectRejectResourceType = "script"
	MarkdownNewParamsBodyObjectRejectResourceTypeTexttrack          MarkdownNewParamsBodyObjectRejectResourceType = "texttrack"
	MarkdownNewParamsBodyObjectRejectResourceTypeXHR                MarkdownNewParamsBodyObjectRejectResourceType = "xhr"
	MarkdownNewParamsBodyObjectRejectResourceTypeFetch              MarkdownNewParamsBodyObjectRejectResourceType = "fetch"
	MarkdownNewParamsBodyObjectRejectResourceTypePrefetch           MarkdownNewParamsBodyObjectRejectResourceType = "prefetch"
	MarkdownNewParamsBodyObjectRejectResourceTypeEventsource        MarkdownNewParamsBodyObjectRejectResourceType = "eventsource"
	MarkdownNewParamsBodyObjectRejectResourceTypeWebsocket          MarkdownNewParamsBodyObjectRejectResourceType = "websocket"
	MarkdownNewParamsBodyObjectRejectResourceTypeManifest           MarkdownNewParamsBodyObjectRejectResourceType = "manifest"
	MarkdownNewParamsBodyObjectRejectResourceTypeSignedexchange     MarkdownNewParamsBodyObjectRejectResourceType = "signedexchange"
	MarkdownNewParamsBodyObjectRejectResourceTypePing               MarkdownNewParamsBodyObjectRejectResourceType = "ping"
	MarkdownNewParamsBodyObjectRejectResourceTypeCspviolationreport MarkdownNewParamsBodyObjectRejectResourceType = "cspviolationreport"
	MarkdownNewParamsBodyObjectRejectResourceTypePreflight          MarkdownNewParamsBodyObjectRejectResourceType = "preflight"
	MarkdownNewParamsBodyObjectRejectResourceTypeOther              MarkdownNewParamsBodyObjectRejectResourceType = "other"
)

func (r MarkdownNewParamsBodyObjectRejectResourceType) IsKnown() bool {
	switch r {
	case MarkdownNewParamsBodyObjectRejectResourceTypeDocument, MarkdownNewParamsBodyObjectRejectResourceTypeStylesheet, MarkdownNewParamsBodyObjectRejectResourceTypeImage, MarkdownNewParamsBodyObjectRejectResourceTypeMedia, MarkdownNewParamsBodyObjectRejectResourceTypeFont, MarkdownNewParamsBodyObjectRejectResourceTypeScript, MarkdownNewParamsBodyObjectRejectResourceTypeTexttrack, MarkdownNewParamsBodyObjectRejectResourceTypeXHR, MarkdownNewParamsBodyObjectRejectResourceTypeFetch, MarkdownNewParamsBodyObjectRejectResourceTypePrefetch, MarkdownNewParamsBodyObjectRejectResourceTypeEventsource, MarkdownNewParamsBodyObjectRejectResourceTypeWebsocket, MarkdownNewParamsBodyObjectRejectResourceTypeManifest, MarkdownNewParamsBodyObjectRejectResourceTypeSignedexchange, MarkdownNewParamsBodyObjectRejectResourceTypePing, MarkdownNewParamsBodyObjectRejectResourceTypeCspviolationreport, MarkdownNewParamsBodyObjectRejectResourceTypePreflight, MarkdownNewParamsBodyObjectRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type MarkdownNewParamsBodyObjectViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r MarkdownNewParamsBodyObjectViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type MarkdownNewParamsBodyObjectWaitForSelector struct {
	Selector param.Field[string]                                            `json:"selector,required"`
	Hidden   param.Field[MarkdownNewParamsBodyObjectWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                           `json:"timeout"`
	Visible  param.Field[MarkdownNewParamsBodyObjectWaitForSelectorVisible] `json:"visible"`
}

func (r MarkdownNewParamsBodyObjectWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarkdownNewParamsBodyObjectWaitForSelectorHidden bool

const (
	MarkdownNewParamsBodyObjectWaitForSelectorHiddenTrue MarkdownNewParamsBodyObjectWaitForSelectorHidden = true
)

func (r MarkdownNewParamsBodyObjectWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case MarkdownNewParamsBodyObjectWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type MarkdownNewParamsBodyObjectWaitForSelectorVisible bool

const (
	MarkdownNewParamsBodyObjectWaitForSelectorVisibleTrue MarkdownNewParamsBodyObjectWaitForSelectorVisible = true
)

func (r MarkdownNewParamsBodyObjectWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case MarkdownNewParamsBodyObjectWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type MarkdownNewResponseEnvelope struct {
	// Response status
	Success bool                                `json:"success,required"`
	Errors  []MarkdownNewResponseEnvelopeErrors `json:"errors"`
	// Markdown
	Result string                          `json:"result"`
	JSON   markdownNewResponseEnvelopeJSON `json:"-"`
}

// markdownNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MarkdownNewResponseEnvelope]
type markdownNewResponseEnvelopeJSON struct {
	Success     apijson.Field
	Errors      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarkdownNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r markdownNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MarkdownNewResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                `json:"message,required"`
	JSON    markdownNewResponseEnvelopeErrorsJSON `json:"-"`
}

// markdownNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MarkdownNewResponseEnvelopeErrors]
type markdownNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarkdownNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r markdownNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
