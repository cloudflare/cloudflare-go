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

// ContentService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContentService] method instead.
type ContentService struct {
	Options []option.RequestOption
}

// NewContentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContentService(opts ...option.RequestOption) (r *ContentService) {
	r = &ContentService{}
	r.Options = opts
	return
}

// Fetches rendered HTML content from provided URL or HTML. Check available options
// like `gotoOptions` and `waitFor*` to control page load behaviour.
func (r *ContentService) New(ctx context.Context, params ContentNewParams, opts ...option.RequestOption) (res *string, err error) {
	var env ContentNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/content", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ContentNewParams struct {
	// Account ID.
	AccountID param.Field[string]       `path:"account_id,required"`
	Body      ContentNewParamsBodyUnion `json:"body,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
}

func (r ContentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [ContentNewParams]'s query parameters as `url.Values`.
func (r ContentNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ContentNewParamsBody struct {
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

func (r ContentNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ContentNewParamsBody) implementsContentNewParamsBodyUnion() {}

// Satisfied by [browser_rendering.ContentNewParamsBodyObject],
// [browser_rendering.ContentNewParamsBodyObject], [ContentNewParamsBody].
type ContentNewParamsBodyUnion interface {
	implementsContentNewParamsBodyUnion()
}

type ContentNewParamsBodyObject struct {
	// URL to navigate to, eg. `https://example.com`.
	URL param.Field[string] `json:"url,required" format:"uri"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]ContentNewParamsBodyObjectAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]ContentNewParamsBodyObjectAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]ContentNewParamsBodyObjectAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[ContentNewParamsBodyObjectAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]ContentNewParamsBodyObjectCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                             `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[ContentNewParamsBodyObjectGotoOptions] `json:"gotoOptions"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]ContentNewParamsBodyObjectRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                              `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                           `json:"setJavaScriptEnabled"`
	UserAgent            param.Field[string]                                         `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[ContentNewParamsBodyObjectViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[ContentNewParamsBodyObjectWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r ContentNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ContentNewParamsBodyObject) implementsContentNewParamsBodyUnion() {}

type ContentNewParamsBodyObjectAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r ContentNewParamsBodyObjectAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContentNewParamsBodyObjectAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r ContentNewParamsBodyObjectAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContentNewParamsBodyObjectAllowResourceType string

const (
	ContentNewParamsBodyObjectAllowResourceTypeDocument           ContentNewParamsBodyObjectAllowResourceType = "document"
	ContentNewParamsBodyObjectAllowResourceTypeStylesheet         ContentNewParamsBodyObjectAllowResourceType = "stylesheet"
	ContentNewParamsBodyObjectAllowResourceTypeImage              ContentNewParamsBodyObjectAllowResourceType = "image"
	ContentNewParamsBodyObjectAllowResourceTypeMedia              ContentNewParamsBodyObjectAllowResourceType = "media"
	ContentNewParamsBodyObjectAllowResourceTypeFont               ContentNewParamsBodyObjectAllowResourceType = "font"
	ContentNewParamsBodyObjectAllowResourceTypeScript             ContentNewParamsBodyObjectAllowResourceType = "script"
	ContentNewParamsBodyObjectAllowResourceTypeTexttrack          ContentNewParamsBodyObjectAllowResourceType = "texttrack"
	ContentNewParamsBodyObjectAllowResourceTypeXHR                ContentNewParamsBodyObjectAllowResourceType = "xhr"
	ContentNewParamsBodyObjectAllowResourceTypeFetch              ContentNewParamsBodyObjectAllowResourceType = "fetch"
	ContentNewParamsBodyObjectAllowResourceTypePrefetch           ContentNewParamsBodyObjectAllowResourceType = "prefetch"
	ContentNewParamsBodyObjectAllowResourceTypeEventsource        ContentNewParamsBodyObjectAllowResourceType = "eventsource"
	ContentNewParamsBodyObjectAllowResourceTypeWebsocket          ContentNewParamsBodyObjectAllowResourceType = "websocket"
	ContentNewParamsBodyObjectAllowResourceTypeManifest           ContentNewParamsBodyObjectAllowResourceType = "manifest"
	ContentNewParamsBodyObjectAllowResourceTypeSignedexchange     ContentNewParamsBodyObjectAllowResourceType = "signedexchange"
	ContentNewParamsBodyObjectAllowResourceTypePing               ContentNewParamsBodyObjectAllowResourceType = "ping"
	ContentNewParamsBodyObjectAllowResourceTypeCspviolationreport ContentNewParamsBodyObjectAllowResourceType = "cspviolationreport"
	ContentNewParamsBodyObjectAllowResourceTypePreflight          ContentNewParamsBodyObjectAllowResourceType = "preflight"
	ContentNewParamsBodyObjectAllowResourceTypeOther              ContentNewParamsBodyObjectAllowResourceType = "other"
)

func (r ContentNewParamsBodyObjectAllowResourceType) IsKnown() bool {
	switch r {
	case ContentNewParamsBodyObjectAllowResourceTypeDocument, ContentNewParamsBodyObjectAllowResourceTypeStylesheet, ContentNewParamsBodyObjectAllowResourceTypeImage, ContentNewParamsBodyObjectAllowResourceTypeMedia, ContentNewParamsBodyObjectAllowResourceTypeFont, ContentNewParamsBodyObjectAllowResourceTypeScript, ContentNewParamsBodyObjectAllowResourceTypeTexttrack, ContentNewParamsBodyObjectAllowResourceTypeXHR, ContentNewParamsBodyObjectAllowResourceTypeFetch, ContentNewParamsBodyObjectAllowResourceTypePrefetch, ContentNewParamsBodyObjectAllowResourceTypeEventsource, ContentNewParamsBodyObjectAllowResourceTypeWebsocket, ContentNewParamsBodyObjectAllowResourceTypeManifest, ContentNewParamsBodyObjectAllowResourceTypeSignedexchange, ContentNewParamsBodyObjectAllowResourceTypePing, ContentNewParamsBodyObjectAllowResourceTypeCspviolationreport, ContentNewParamsBodyObjectAllowResourceTypePreflight, ContentNewParamsBodyObjectAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type ContentNewParamsBodyObjectAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r ContentNewParamsBodyObjectAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContentNewParamsBodyObjectCookie struct {
	Name         param.Field[string]                                        `json:"name,required"`
	Value        param.Field[string]                                        `json:"value,required"`
	Domain       param.Field[string]                                        `json:"domain"`
	Expires      param.Field[float64]                                       `json:"expires"`
	HTTPOnly     param.Field[bool]                                          `json:"httpOnly"`
	PartitionKey param.Field[string]                                        `json:"partitionKey"`
	Path         param.Field[string]                                        `json:"path"`
	Priority     param.Field[ContentNewParamsBodyObjectCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                          `json:"sameParty"`
	SameSite     param.Field[ContentNewParamsBodyObjectCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                          `json:"secure"`
	SourcePort   param.Field[float64]                                       `json:"sourcePort"`
	SourceScheme param.Field[ContentNewParamsBodyObjectCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                        `json:"url"`
}

func (r ContentNewParamsBodyObjectCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContentNewParamsBodyObjectCookiesPriority string

const (
	ContentNewParamsBodyObjectCookiesPriorityLow    ContentNewParamsBodyObjectCookiesPriority = "Low"
	ContentNewParamsBodyObjectCookiesPriorityMedium ContentNewParamsBodyObjectCookiesPriority = "Medium"
	ContentNewParamsBodyObjectCookiesPriorityHigh   ContentNewParamsBodyObjectCookiesPriority = "High"
)

func (r ContentNewParamsBodyObjectCookiesPriority) IsKnown() bool {
	switch r {
	case ContentNewParamsBodyObjectCookiesPriorityLow, ContentNewParamsBodyObjectCookiesPriorityMedium, ContentNewParamsBodyObjectCookiesPriorityHigh:
		return true
	}
	return false
}

type ContentNewParamsBodyObjectCookiesSameSite string

const (
	ContentNewParamsBodyObjectCookiesSameSiteStrict ContentNewParamsBodyObjectCookiesSameSite = "Strict"
	ContentNewParamsBodyObjectCookiesSameSiteLax    ContentNewParamsBodyObjectCookiesSameSite = "Lax"
	ContentNewParamsBodyObjectCookiesSameSiteNone   ContentNewParamsBodyObjectCookiesSameSite = "None"
)

func (r ContentNewParamsBodyObjectCookiesSameSite) IsKnown() bool {
	switch r {
	case ContentNewParamsBodyObjectCookiesSameSiteStrict, ContentNewParamsBodyObjectCookiesSameSiteLax, ContentNewParamsBodyObjectCookiesSameSiteNone:
		return true
	}
	return false
}

type ContentNewParamsBodyObjectCookiesSourceScheme string

const (
	ContentNewParamsBodyObjectCookiesSourceSchemeUnset     ContentNewParamsBodyObjectCookiesSourceScheme = "Unset"
	ContentNewParamsBodyObjectCookiesSourceSchemeNonSecure ContentNewParamsBodyObjectCookiesSourceScheme = "NonSecure"
	ContentNewParamsBodyObjectCookiesSourceSchemeSecure    ContentNewParamsBodyObjectCookiesSourceScheme = "Secure"
)

func (r ContentNewParamsBodyObjectCookiesSourceScheme) IsKnown() bool {
	switch r {
	case ContentNewParamsBodyObjectCookiesSourceSchemeUnset, ContentNewParamsBodyObjectCookiesSourceSchemeNonSecure, ContentNewParamsBodyObjectCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type ContentNewParamsBodyObjectGotoOptions struct {
	Referer        param.Field[string]                                              `json:"referer"`
	ReferrerPolicy param.Field[string]                                              `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                             `json:"timeout"`
	WaitUntil      param.Field[ContentNewParamsBodyObjectGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r ContentNewParamsBodyObjectGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.ContentNewParamsBodyObjectGotoOptionsWaitUntilString],
// [browser_rendering.ContentNewParamsBodyObjectGotoOptionsWaitUntilArray].
type ContentNewParamsBodyObjectGotoOptionsWaitUntilUnion interface {
	implementsContentNewParamsBodyObjectGotoOptionsWaitUntilUnion()
}

type ContentNewParamsBodyObjectGotoOptionsWaitUntilString string

const (
	ContentNewParamsBodyObjectGotoOptionsWaitUntilStringLoad             ContentNewParamsBodyObjectGotoOptionsWaitUntilString = "load"
	ContentNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded ContentNewParamsBodyObjectGotoOptionsWaitUntilString = "domcontentloaded"
	ContentNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0     ContentNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle0"
	ContentNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2     ContentNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle2"
)

func (r ContentNewParamsBodyObjectGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case ContentNewParamsBodyObjectGotoOptionsWaitUntilStringLoad, ContentNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded, ContentNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0, ContentNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r ContentNewParamsBodyObjectGotoOptionsWaitUntilString) implementsContentNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type ContentNewParamsBodyObjectGotoOptionsWaitUntilArray []ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItem

func (r ContentNewParamsBodyObjectGotoOptionsWaitUntilArray) implementsContentNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItem string

const (
	ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad             ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "load"
	ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0     ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle0"
	ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2     ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad, ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded, ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0, ContentNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type ContentNewParamsBodyObjectRejectResourceType string

const (
	ContentNewParamsBodyObjectRejectResourceTypeDocument           ContentNewParamsBodyObjectRejectResourceType = "document"
	ContentNewParamsBodyObjectRejectResourceTypeStylesheet         ContentNewParamsBodyObjectRejectResourceType = "stylesheet"
	ContentNewParamsBodyObjectRejectResourceTypeImage              ContentNewParamsBodyObjectRejectResourceType = "image"
	ContentNewParamsBodyObjectRejectResourceTypeMedia              ContentNewParamsBodyObjectRejectResourceType = "media"
	ContentNewParamsBodyObjectRejectResourceTypeFont               ContentNewParamsBodyObjectRejectResourceType = "font"
	ContentNewParamsBodyObjectRejectResourceTypeScript             ContentNewParamsBodyObjectRejectResourceType = "script"
	ContentNewParamsBodyObjectRejectResourceTypeTexttrack          ContentNewParamsBodyObjectRejectResourceType = "texttrack"
	ContentNewParamsBodyObjectRejectResourceTypeXHR                ContentNewParamsBodyObjectRejectResourceType = "xhr"
	ContentNewParamsBodyObjectRejectResourceTypeFetch              ContentNewParamsBodyObjectRejectResourceType = "fetch"
	ContentNewParamsBodyObjectRejectResourceTypePrefetch           ContentNewParamsBodyObjectRejectResourceType = "prefetch"
	ContentNewParamsBodyObjectRejectResourceTypeEventsource        ContentNewParamsBodyObjectRejectResourceType = "eventsource"
	ContentNewParamsBodyObjectRejectResourceTypeWebsocket          ContentNewParamsBodyObjectRejectResourceType = "websocket"
	ContentNewParamsBodyObjectRejectResourceTypeManifest           ContentNewParamsBodyObjectRejectResourceType = "manifest"
	ContentNewParamsBodyObjectRejectResourceTypeSignedexchange     ContentNewParamsBodyObjectRejectResourceType = "signedexchange"
	ContentNewParamsBodyObjectRejectResourceTypePing               ContentNewParamsBodyObjectRejectResourceType = "ping"
	ContentNewParamsBodyObjectRejectResourceTypeCspviolationreport ContentNewParamsBodyObjectRejectResourceType = "cspviolationreport"
	ContentNewParamsBodyObjectRejectResourceTypePreflight          ContentNewParamsBodyObjectRejectResourceType = "preflight"
	ContentNewParamsBodyObjectRejectResourceTypeOther              ContentNewParamsBodyObjectRejectResourceType = "other"
)

func (r ContentNewParamsBodyObjectRejectResourceType) IsKnown() bool {
	switch r {
	case ContentNewParamsBodyObjectRejectResourceTypeDocument, ContentNewParamsBodyObjectRejectResourceTypeStylesheet, ContentNewParamsBodyObjectRejectResourceTypeImage, ContentNewParamsBodyObjectRejectResourceTypeMedia, ContentNewParamsBodyObjectRejectResourceTypeFont, ContentNewParamsBodyObjectRejectResourceTypeScript, ContentNewParamsBodyObjectRejectResourceTypeTexttrack, ContentNewParamsBodyObjectRejectResourceTypeXHR, ContentNewParamsBodyObjectRejectResourceTypeFetch, ContentNewParamsBodyObjectRejectResourceTypePrefetch, ContentNewParamsBodyObjectRejectResourceTypeEventsource, ContentNewParamsBodyObjectRejectResourceTypeWebsocket, ContentNewParamsBodyObjectRejectResourceTypeManifest, ContentNewParamsBodyObjectRejectResourceTypeSignedexchange, ContentNewParamsBodyObjectRejectResourceTypePing, ContentNewParamsBodyObjectRejectResourceTypeCspviolationreport, ContentNewParamsBodyObjectRejectResourceTypePreflight, ContentNewParamsBodyObjectRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type ContentNewParamsBodyObjectViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r ContentNewParamsBodyObjectViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type ContentNewParamsBodyObjectWaitForSelector struct {
	Selector param.Field[string]                                           `json:"selector,required"`
	Hidden   param.Field[ContentNewParamsBodyObjectWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                          `json:"timeout"`
	Visible  param.Field[ContentNewParamsBodyObjectWaitForSelectorVisible] `json:"visible"`
}

func (r ContentNewParamsBodyObjectWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContentNewParamsBodyObjectWaitForSelectorHidden bool

const (
	ContentNewParamsBodyObjectWaitForSelectorHiddenTrue ContentNewParamsBodyObjectWaitForSelectorHidden = true
)

func (r ContentNewParamsBodyObjectWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case ContentNewParamsBodyObjectWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type ContentNewParamsBodyObjectWaitForSelectorVisible bool

const (
	ContentNewParamsBodyObjectWaitForSelectorVisibleTrue ContentNewParamsBodyObjectWaitForSelectorVisible = true
)

func (r ContentNewParamsBodyObjectWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case ContentNewParamsBodyObjectWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type ContentNewResponseEnvelope struct {
	Meta ContentNewResponseEnvelopeMeta `json:"meta,required"`
	// Response status
	Success bool                               `json:"success,required"`
	Errors  []ContentNewResponseEnvelopeErrors `json:"errors"`
	// HTML content
	Result string                         `json:"result"`
	JSON   contentNewResponseEnvelopeJSON `json:"-"`
}

// contentNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ContentNewResponseEnvelope]
type contentNewResponseEnvelopeJSON struct {
	Meta        apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ContentNewResponseEnvelopeMeta struct {
	Status float64                            `json:"status"`
	Title  string                             `json:"title"`
	JSON   contentNewResponseEnvelopeMetaJSON `json:"-"`
}

// contentNewResponseEnvelopeMetaJSON contains the JSON metadata for the struct
// [ContentNewResponseEnvelopeMeta]
type contentNewResponseEnvelopeMetaJSON struct {
	Status      apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentNewResponseEnvelopeMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentNewResponseEnvelopeMetaJSON) RawJSON() string {
	return r.raw
}

type ContentNewResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                               `json:"message,required"`
	JSON    contentNewResponseEnvelopeErrorsJSON `json:"-"`
}

// contentNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ContentNewResponseEnvelopeErrors]
type contentNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
