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

// LinkService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLinkService] method instead.
type LinkService struct {
	Options []option.RequestOption
}

// NewLinkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLinkService(opts ...option.RequestOption) (r *LinkService) {
	r = &LinkService{}
	r.Options = opts
	return
}

// Get links from a web page.
func (r *LinkService) New(ctx context.Context, params LinkNewParams, opts ...option.RequestOption) (res *[]string, err error) {
	var env LinkNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/links", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LinkNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64]   `query:"cacheTTL"`
	Body     LinkNewParamsBodyUnion `json:"body"`
}

func (r LinkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [LinkNewParams]'s query parameters as `url.Values`.
func (r LinkNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LinkNewParamsBody struct {
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
	BestAttempt          param.Field[bool]        `json:"bestAttempt"`
	Cookies              param.Field[interface{}] `json:"cookies"`
	EmulateMediaType     param.Field[string]      `json:"emulateMediaType"`
	ExcludeExternalLinks param.Field[bool]        `json:"excludeExternalLinks"`
	GotoOptions          param.Field[interface{}] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML                 param.Field[string]      `json:"html"`
	RejectRequestPattern param.Field[interface{}] `json:"rejectRequestPattern"`
	RejectResourceTypes  param.Field[interface{}] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[interface{}] `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]        `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL              param.Field[string]      `json:"url" format:"uri"`
	UserAgent        param.Field[string]      `json:"userAgent"`
	Viewport         param.Field[interface{}] `json:"viewport"`
	VisibleLinksOnly param.Field[bool]        `json:"visibleLinksOnly"`
	WaitForSelector  param.Field[interface{}] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r LinkNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LinkNewParamsBody) implementsLinkNewParamsBodyUnion() {}

// Satisfied by [browser_rendering.LinkNewParamsBodyObject],
// [browser_rendering.LinkNewParamsBodyObject], [LinkNewParamsBody].
type LinkNewParamsBodyUnion interface {
	implementsLinkNewParamsBodyUnion()
}

type LinkNewParamsBodyObject struct {
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html,required"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]LinkNewParamsBodyObjectAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]LinkNewParamsBodyObjectAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]LinkNewParamsBodyObjectAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[LinkNewParamsBodyObjectAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies              param.Field[[]LinkNewParamsBodyObjectCookie] `json:"cookies"`
	EmulateMediaType     param.Field[string]                          `json:"emulateMediaType"`
	ExcludeExternalLinks param.Field[bool]                            `json:"excludeExternalLinks"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[LinkNewParamsBodyObjectGotoOptions] `json:"gotoOptions"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]LinkNewParamsBodyObjectRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                           `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                        `json:"setJavaScriptEnabled"`
	UserAgent            param.Field[string]                                      `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport         param.Field[LinkNewParamsBodyObjectViewport] `json:"viewport"`
	VisibleLinksOnly param.Field[bool]                            `json:"visibleLinksOnly"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[LinkNewParamsBodyObjectWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r LinkNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LinkNewParamsBodyObject) implementsLinkNewParamsBodyUnion() {}

type LinkNewParamsBodyObjectAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r LinkNewParamsBodyObjectAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LinkNewParamsBodyObjectAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r LinkNewParamsBodyObjectAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LinkNewParamsBodyObjectAllowResourceType string

const (
	LinkNewParamsBodyObjectAllowResourceTypeDocument           LinkNewParamsBodyObjectAllowResourceType = "document"
	LinkNewParamsBodyObjectAllowResourceTypeStylesheet         LinkNewParamsBodyObjectAllowResourceType = "stylesheet"
	LinkNewParamsBodyObjectAllowResourceTypeImage              LinkNewParamsBodyObjectAllowResourceType = "image"
	LinkNewParamsBodyObjectAllowResourceTypeMedia              LinkNewParamsBodyObjectAllowResourceType = "media"
	LinkNewParamsBodyObjectAllowResourceTypeFont               LinkNewParamsBodyObjectAllowResourceType = "font"
	LinkNewParamsBodyObjectAllowResourceTypeScript             LinkNewParamsBodyObjectAllowResourceType = "script"
	LinkNewParamsBodyObjectAllowResourceTypeTexttrack          LinkNewParamsBodyObjectAllowResourceType = "texttrack"
	LinkNewParamsBodyObjectAllowResourceTypeXHR                LinkNewParamsBodyObjectAllowResourceType = "xhr"
	LinkNewParamsBodyObjectAllowResourceTypeFetch              LinkNewParamsBodyObjectAllowResourceType = "fetch"
	LinkNewParamsBodyObjectAllowResourceTypePrefetch           LinkNewParamsBodyObjectAllowResourceType = "prefetch"
	LinkNewParamsBodyObjectAllowResourceTypeEventsource        LinkNewParamsBodyObjectAllowResourceType = "eventsource"
	LinkNewParamsBodyObjectAllowResourceTypeWebsocket          LinkNewParamsBodyObjectAllowResourceType = "websocket"
	LinkNewParamsBodyObjectAllowResourceTypeManifest           LinkNewParamsBodyObjectAllowResourceType = "manifest"
	LinkNewParamsBodyObjectAllowResourceTypeSignedexchange     LinkNewParamsBodyObjectAllowResourceType = "signedexchange"
	LinkNewParamsBodyObjectAllowResourceTypePing               LinkNewParamsBodyObjectAllowResourceType = "ping"
	LinkNewParamsBodyObjectAllowResourceTypeCspviolationreport LinkNewParamsBodyObjectAllowResourceType = "cspviolationreport"
	LinkNewParamsBodyObjectAllowResourceTypePreflight          LinkNewParamsBodyObjectAllowResourceType = "preflight"
	LinkNewParamsBodyObjectAllowResourceTypeOther              LinkNewParamsBodyObjectAllowResourceType = "other"
)

func (r LinkNewParamsBodyObjectAllowResourceType) IsKnown() bool {
	switch r {
	case LinkNewParamsBodyObjectAllowResourceTypeDocument, LinkNewParamsBodyObjectAllowResourceTypeStylesheet, LinkNewParamsBodyObjectAllowResourceTypeImage, LinkNewParamsBodyObjectAllowResourceTypeMedia, LinkNewParamsBodyObjectAllowResourceTypeFont, LinkNewParamsBodyObjectAllowResourceTypeScript, LinkNewParamsBodyObjectAllowResourceTypeTexttrack, LinkNewParamsBodyObjectAllowResourceTypeXHR, LinkNewParamsBodyObjectAllowResourceTypeFetch, LinkNewParamsBodyObjectAllowResourceTypePrefetch, LinkNewParamsBodyObjectAllowResourceTypeEventsource, LinkNewParamsBodyObjectAllowResourceTypeWebsocket, LinkNewParamsBodyObjectAllowResourceTypeManifest, LinkNewParamsBodyObjectAllowResourceTypeSignedexchange, LinkNewParamsBodyObjectAllowResourceTypePing, LinkNewParamsBodyObjectAllowResourceTypeCspviolationreport, LinkNewParamsBodyObjectAllowResourceTypePreflight, LinkNewParamsBodyObjectAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type LinkNewParamsBodyObjectAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r LinkNewParamsBodyObjectAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LinkNewParamsBodyObjectCookie struct {
	Name         param.Field[string]                                     `json:"name,required"`
	Value        param.Field[string]                                     `json:"value,required"`
	Domain       param.Field[string]                                     `json:"domain"`
	Expires      param.Field[float64]                                    `json:"expires"`
	HTTPOnly     param.Field[bool]                                       `json:"httpOnly"`
	PartitionKey param.Field[string]                                     `json:"partitionKey"`
	Path         param.Field[string]                                     `json:"path"`
	Priority     param.Field[LinkNewParamsBodyObjectCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                       `json:"sameParty"`
	SameSite     param.Field[LinkNewParamsBodyObjectCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                       `json:"secure"`
	SourcePort   param.Field[float64]                                    `json:"sourcePort"`
	SourceScheme param.Field[LinkNewParamsBodyObjectCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                     `json:"url"`
}

func (r LinkNewParamsBodyObjectCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LinkNewParamsBodyObjectCookiesPriority string

const (
	LinkNewParamsBodyObjectCookiesPriorityLow    LinkNewParamsBodyObjectCookiesPriority = "Low"
	LinkNewParamsBodyObjectCookiesPriorityMedium LinkNewParamsBodyObjectCookiesPriority = "Medium"
	LinkNewParamsBodyObjectCookiesPriorityHigh   LinkNewParamsBodyObjectCookiesPriority = "High"
)

func (r LinkNewParamsBodyObjectCookiesPriority) IsKnown() bool {
	switch r {
	case LinkNewParamsBodyObjectCookiesPriorityLow, LinkNewParamsBodyObjectCookiesPriorityMedium, LinkNewParamsBodyObjectCookiesPriorityHigh:
		return true
	}
	return false
}

type LinkNewParamsBodyObjectCookiesSameSite string

const (
	LinkNewParamsBodyObjectCookiesSameSiteStrict LinkNewParamsBodyObjectCookiesSameSite = "Strict"
	LinkNewParamsBodyObjectCookiesSameSiteLax    LinkNewParamsBodyObjectCookiesSameSite = "Lax"
	LinkNewParamsBodyObjectCookiesSameSiteNone   LinkNewParamsBodyObjectCookiesSameSite = "None"
)

func (r LinkNewParamsBodyObjectCookiesSameSite) IsKnown() bool {
	switch r {
	case LinkNewParamsBodyObjectCookiesSameSiteStrict, LinkNewParamsBodyObjectCookiesSameSiteLax, LinkNewParamsBodyObjectCookiesSameSiteNone:
		return true
	}
	return false
}

type LinkNewParamsBodyObjectCookiesSourceScheme string

const (
	LinkNewParamsBodyObjectCookiesSourceSchemeUnset     LinkNewParamsBodyObjectCookiesSourceScheme = "Unset"
	LinkNewParamsBodyObjectCookiesSourceSchemeNonSecure LinkNewParamsBodyObjectCookiesSourceScheme = "NonSecure"
	LinkNewParamsBodyObjectCookiesSourceSchemeSecure    LinkNewParamsBodyObjectCookiesSourceScheme = "Secure"
)

func (r LinkNewParamsBodyObjectCookiesSourceScheme) IsKnown() bool {
	switch r {
	case LinkNewParamsBodyObjectCookiesSourceSchemeUnset, LinkNewParamsBodyObjectCookiesSourceSchemeNonSecure, LinkNewParamsBodyObjectCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type LinkNewParamsBodyObjectGotoOptions struct {
	Referer        param.Field[string]                                           `json:"referer"`
	ReferrerPolicy param.Field[string]                                           `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                          `json:"timeout"`
	WaitUntil      param.Field[LinkNewParamsBodyObjectGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r LinkNewParamsBodyObjectGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.LinkNewParamsBodyObjectGotoOptionsWaitUntilString],
// [browser_rendering.LinkNewParamsBodyObjectGotoOptionsWaitUntilArray].
type LinkNewParamsBodyObjectGotoOptionsWaitUntilUnion interface {
	implementsLinkNewParamsBodyObjectGotoOptionsWaitUntilUnion()
}

type LinkNewParamsBodyObjectGotoOptionsWaitUntilString string

const (
	LinkNewParamsBodyObjectGotoOptionsWaitUntilStringLoad             LinkNewParamsBodyObjectGotoOptionsWaitUntilString = "load"
	LinkNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded LinkNewParamsBodyObjectGotoOptionsWaitUntilString = "domcontentloaded"
	LinkNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0     LinkNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle0"
	LinkNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2     LinkNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle2"
)

func (r LinkNewParamsBodyObjectGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case LinkNewParamsBodyObjectGotoOptionsWaitUntilStringLoad, LinkNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded, LinkNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0, LinkNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r LinkNewParamsBodyObjectGotoOptionsWaitUntilString) implementsLinkNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type LinkNewParamsBodyObjectGotoOptionsWaitUntilArray []LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItem

func (r LinkNewParamsBodyObjectGotoOptionsWaitUntilArray) implementsLinkNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItem string

const (
	LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad             LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "load"
	LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0     LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle0"
	LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2     LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad, LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded, LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0, LinkNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type LinkNewParamsBodyObjectRejectResourceType string

const (
	LinkNewParamsBodyObjectRejectResourceTypeDocument           LinkNewParamsBodyObjectRejectResourceType = "document"
	LinkNewParamsBodyObjectRejectResourceTypeStylesheet         LinkNewParamsBodyObjectRejectResourceType = "stylesheet"
	LinkNewParamsBodyObjectRejectResourceTypeImage              LinkNewParamsBodyObjectRejectResourceType = "image"
	LinkNewParamsBodyObjectRejectResourceTypeMedia              LinkNewParamsBodyObjectRejectResourceType = "media"
	LinkNewParamsBodyObjectRejectResourceTypeFont               LinkNewParamsBodyObjectRejectResourceType = "font"
	LinkNewParamsBodyObjectRejectResourceTypeScript             LinkNewParamsBodyObjectRejectResourceType = "script"
	LinkNewParamsBodyObjectRejectResourceTypeTexttrack          LinkNewParamsBodyObjectRejectResourceType = "texttrack"
	LinkNewParamsBodyObjectRejectResourceTypeXHR                LinkNewParamsBodyObjectRejectResourceType = "xhr"
	LinkNewParamsBodyObjectRejectResourceTypeFetch              LinkNewParamsBodyObjectRejectResourceType = "fetch"
	LinkNewParamsBodyObjectRejectResourceTypePrefetch           LinkNewParamsBodyObjectRejectResourceType = "prefetch"
	LinkNewParamsBodyObjectRejectResourceTypeEventsource        LinkNewParamsBodyObjectRejectResourceType = "eventsource"
	LinkNewParamsBodyObjectRejectResourceTypeWebsocket          LinkNewParamsBodyObjectRejectResourceType = "websocket"
	LinkNewParamsBodyObjectRejectResourceTypeManifest           LinkNewParamsBodyObjectRejectResourceType = "manifest"
	LinkNewParamsBodyObjectRejectResourceTypeSignedexchange     LinkNewParamsBodyObjectRejectResourceType = "signedexchange"
	LinkNewParamsBodyObjectRejectResourceTypePing               LinkNewParamsBodyObjectRejectResourceType = "ping"
	LinkNewParamsBodyObjectRejectResourceTypeCspviolationreport LinkNewParamsBodyObjectRejectResourceType = "cspviolationreport"
	LinkNewParamsBodyObjectRejectResourceTypePreflight          LinkNewParamsBodyObjectRejectResourceType = "preflight"
	LinkNewParamsBodyObjectRejectResourceTypeOther              LinkNewParamsBodyObjectRejectResourceType = "other"
)

func (r LinkNewParamsBodyObjectRejectResourceType) IsKnown() bool {
	switch r {
	case LinkNewParamsBodyObjectRejectResourceTypeDocument, LinkNewParamsBodyObjectRejectResourceTypeStylesheet, LinkNewParamsBodyObjectRejectResourceTypeImage, LinkNewParamsBodyObjectRejectResourceTypeMedia, LinkNewParamsBodyObjectRejectResourceTypeFont, LinkNewParamsBodyObjectRejectResourceTypeScript, LinkNewParamsBodyObjectRejectResourceTypeTexttrack, LinkNewParamsBodyObjectRejectResourceTypeXHR, LinkNewParamsBodyObjectRejectResourceTypeFetch, LinkNewParamsBodyObjectRejectResourceTypePrefetch, LinkNewParamsBodyObjectRejectResourceTypeEventsource, LinkNewParamsBodyObjectRejectResourceTypeWebsocket, LinkNewParamsBodyObjectRejectResourceTypeManifest, LinkNewParamsBodyObjectRejectResourceTypeSignedexchange, LinkNewParamsBodyObjectRejectResourceTypePing, LinkNewParamsBodyObjectRejectResourceTypeCspviolationreport, LinkNewParamsBodyObjectRejectResourceTypePreflight, LinkNewParamsBodyObjectRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type LinkNewParamsBodyObjectViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r LinkNewParamsBodyObjectViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type LinkNewParamsBodyObjectWaitForSelector struct {
	Selector param.Field[string]                                        `json:"selector,required"`
	Hidden   param.Field[LinkNewParamsBodyObjectWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                       `json:"timeout"`
	Visible  param.Field[LinkNewParamsBodyObjectWaitForSelectorVisible] `json:"visible"`
}

func (r LinkNewParamsBodyObjectWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LinkNewParamsBodyObjectWaitForSelectorHidden bool

const (
	LinkNewParamsBodyObjectWaitForSelectorHiddenTrue LinkNewParamsBodyObjectWaitForSelectorHidden = true
)

func (r LinkNewParamsBodyObjectWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case LinkNewParamsBodyObjectWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type LinkNewParamsBodyObjectWaitForSelectorVisible bool

const (
	LinkNewParamsBodyObjectWaitForSelectorVisibleTrue LinkNewParamsBodyObjectWaitForSelectorVisible = true
)

func (r LinkNewParamsBodyObjectWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case LinkNewParamsBodyObjectWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type LinkNewResponseEnvelope struct {
	Result []string `json:"result,required"`
	// Response status
	Success bool                            `json:"success,required"`
	Errors  []LinkNewResponseEnvelopeErrors `json:"errors"`
	JSON    linkNewResponseEnvelopeJSON     `json:"-"`
}

// linkNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LinkNewResponseEnvelope]
type linkNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LinkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LinkNewResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                            `json:"message,required"`
	JSON    linkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// linkNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [LinkNewResponseEnvelopeErrors]
type linkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LinkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
