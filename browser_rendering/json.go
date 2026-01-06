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

// JsonService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJsonService] method instead.
type JsonService struct {
	Options []option.RequestOption
}

// NewJsonService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJsonService(opts ...option.RequestOption) (r *JsonService) {
	r = &JsonService{}
	r.Options = opts
	return
}

// Gets json from a webpage from a provided URL or HTML. Pass `prompt` or `schema`
// in the body. Control page loading with `gotoOptions` and `waitFor*` options.
func (r *JsonService) New(ctx context.Context, params JsonNewParams, opts ...option.RequestOption) (res *JsonNewResponse, err error) {
	var env JsonNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/json", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type JsonNewResponse map[string]interface{}

type JsonNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64]   `query:"cacheTTL"`
	Body     JsonNewParamsBodyUnion `json:"body"`
}

func (r JsonNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [JsonNewParams]'s query parameters as `url.Values`.
func (r JsonNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type JsonNewParamsBody struct {
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
	CustomAI         param.Field[interface{}] `json:"custom_ai"`
	EmulateMediaType param.Field[string]      `json:"emulateMediaType"`
	GotoOptions      param.Field[interface{}] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML                 param.Field[string]      `json:"html"`
	Prompt               param.Field[string]      `json:"prompt"`
	RejectRequestPattern param.Field[interface{}] `json:"rejectRequestPattern"`
	RejectResourceTypes  param.Field[interface{}] `json:"rejectResourceTypes"`
	ResponseFormat       param.Field[interface{}] `json:"response_format"`
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

func (r JsonNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r JsonNewParamsBody) implementsJsonNewParamsBodyUnion() {}

// Satisfied by [browser_rendering.JsonNewParamsBodyObject],
// [browser_rendering.JsonNewParamsBodyObject], [JsonNewParamsBody].
type JsonNewParamsBodyUnion interface {
	implementsJsonNewParamsBodyUnion()
}

type JsonNewParamsBodyObject struct {
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html,required"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]JsonNewParamsBodyObjectAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]JsonNewParamsBodyObjectAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]JsonNewParamsBodyObjectAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[JsonNewParamsBodyObjectAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies param.Field[[]JsonNewParamsBodyObjectCookie] `json:"cookies"`
	// Optional list of custom AI models to use for the request. The models will be
	// tried in the order provided, and in case a model returns an error, the next one
	// will be used as fallback.
	CustomAI         param.Field[[]JsonNewParamsBodyObjectCustomAI] `json:"custom_ai"`
	EmulateMediaType param.Field[string]                            `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[JsonNewParamsBodyObjectGotoOptions] `json:"gotoOptions"`
	Prompt      param.Field[string]                             `json:"prompt"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]JsonNewParamsBodyObjectRejectResourceType] `json:"rejectResourceTypes"`
	ResponseFormat       param.Field[JsonNewParamsBodyObjectResponseFormat]       `json:"response_format"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                           `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                        `json:"setJavaScriptEnabled"`
	UserAgent            param.Field[string]                                      `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[JsonNewParamsBodyObjectViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[JsonNewParamsBodyObjectWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r JsonNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r JsonNewParamsBodyObject) implementsJsonNewParamsBodyUnion() {}

type JsonNewParamsBodyObjectAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r JsonNewParamsBodyObjectAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JsonNewParamsBodyObjectAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r JsonNewParamsBodyObjectAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JsonNewParamsBodyObjectAllowResourceType string

const (
	JsonNewParamsBodyObjectAllowResourceTypeDocument           JsonNewParamsBodyObjectAllowResourceType = "document"
	JsonNewParamsBodyObjectAllowResourceTypeStylesheet         JsonNewParamsBodyObjectAllowResourceType = "stylesheet"
	JsonNewParamsBodyObjectAllowResourceTypeImage              JsonNewParamsBodyObjectAllowResourceType = "image"
	JsonNewParamsBodyObjectAllowResourceTypeMedia              JsonNewParamsBodyObjectAllowResourceType = "media"
	JsonNewParamsBodyObjectAllowResourceTypeFont               JsonNewParamsBodyObjectAllowResourceType = "font"
	JsonNewParamsBodyObjectAllowResourceTypeScript             JsonNewParamsBodyObjectAllowResourceType = "script"
	JsonNewParamsBodyObjectAllowResourceTypeTexttrack          JsonNewParamsBodyObjectAllowResourceType = "texttrack"
	JsonNewParamsBodyObjectAllowResourceTypeXHR                JsonNewParamsBodyObjectAllowResourceType = "xhr"
	JsonNewParamsBodyObjectAllowResourceTypeFetch              JsonNewParamsBodyObjectAllowResourceType = "fetch"
	JsonNewParamsBodyObjectAllowResourceTypePrefetch           JsonNewParamsBodyObjectAllowResourceType = "prefetch"
	JsonNewParamsBodyObjectAllowResourceTypeEventsource        JsonNewParamsBodyObjectAllowResourceType = "eventsource"
	JsonNewParamsBodyObjectAllowResourceTypeWebsocket          JsonNewParamsBodyObjectAllowResourceType = "websocket"
	JsonNewParamsBodyObjectAllowResourceTypeManifest           JsonNewParamsBodyObjectAllowResourceType = "manifest"
	JsonNewParamsBodyObjectAllowResourceTypeSignedexchange     JsonNewParamsBodyObjectAllowResourceType = "signedexchange"
	JsonNewParamsBodyObjectAllowResourceTypePing               JsonNewParamsBodyObjectAllowResourceType = "ping"
	JsonNewParamsBodyObjectAllowResourceTypeCspviolationreport JsonNewParamsBodyObjectAllowResourceType = "cspviolationreport"
	JsonNewParamsBodyObjectAllowResourceTypePreflight          JsonNewParamsBodyObjectAllowResourceType = "preflight"
	JsonNewParamsBodyObjectAllowResourceTypeOther              JsonNewParamsBodyObjectAllowResourceType = "other"
)

func (r JsonNewParamsBodyObjectAllowResourceType) IsKnown() bool {
	switch r {
	case JsonNewParamsBodyObjectAllowResourceTypeDocument, JsonNewParamsBodyObjectAllowResourceTypeStylesheet, JsonNewParamsBodyObjectAllowResourceTypeImage, JsonNewParamsBodyObjectAllowResourceTypeMedia, JsonNewParamsBodyObjectAllowResourceTypeFont, JsonNewParamsBodyObjectAllowResourceTypeScript, JsonNewParamsBodyObjectAllowResourceTypeTexttrack, JsonNewParamsBodyObjectAllowResourceTypeXHR, JsonNewParamsBodyObjectAllowResourceTypeFetch, JsonNewParamsBodyObjectAllowResourceTypePrefetch, JsonNewParamsBodyObjectAllowResourceTypeEventsource, JsonNewParamsBodyObjectAllowResourceTypeWebsocket, JsonNewParamsBodyObjectAllowResourceTypeManifest, JsonNewParamsBodyObjectAllowResourceTypeSignedexchange, JsonNewParamsBodyObjectAllowResourceTypePing, JsonNewParamsBodyObjectAllowResourceTypeCspviolationreport, JsonNewParamsBodyObjectAllowResourceTypePreflight, JsonNewParamsBodyObjectAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type JsonNewParamsBodyObjectAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r JsonNewParamsBodyObjectAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JsonNewParamsBodyObjectCookie struct {
	Name         param.Field[string]                                     `json:"name,required"`
	Value        param.Field[string]                                     `json:"value,required"`
	Domain       param.Field[string]                                     `json:"domain"`
	Expires      param.Field[float64]                                    `json:"expires"`
	HTTPOnly     param.Field[bool]                                       `json:"httpOnly"`
	PartitionKey param.Field[string]                                     `json:"partitionKey"`
	Path         param.Field[string]                                     `json:"path"`
	Priority     param.Field[JsonNewParamsBodyObjectCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                       `json:"sameParty"`
	SameSite     param.Field[JsonNewParamsBodyObjectCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                       `json:"secure"`
	SourcePort   param.Field[float64]                                    `json:"sourcePort"`
	SourceScheme param.Field[JsonNewParamsBodyObjectCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                     `json:"url"`
}

func (r JsonNewParamsBodyObjectCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JsonNewParamsBodyObjectCookiesPriority string

const (
	JsonNewParamsBodyObjectCookiesPriorityLow    JsonNewParamsBodyObjectCookiesPriority = "Low"
	JsonNewParamsBodyObjectCookiesPriorityMedium JsonNewParamsBodyObjectCookiesPriority = "Medium"
	JsonNewParamsBodyObjectCookiesPriorityHigh   JsonNewParamsBodyObjectCookiesPriority = "High"
)

func (r JsonNewParamsBodyObjectCookiesPriority) IsKnown() bool {
	switch r {
	case JsonNewParamsBodyObjectCookiesPriorityLow, JsonNewParamsBodyObjectCookiesPriorityMedium, JsonNewParamsBodyObjectCookiesPriorityHigh:
		return true
	}
	return false
}

type JsonNewParamsBodyObjectCookiesSameSite string

const (
	JsonNewParamsBodyObjectCookiesSameSiteStrict JsonNewParamsBodyObjectCookiesSameSite = "Strict"
	JsonNewParamsBodyObjectCookiesSameSiteLax    JsonNewParamsBodyObjectCookiesSameSite = "Lax"
	JsonNewParamsBodyObjectCookiesSameSiteNone   JsonNewParamsBodyObjectCookiesSameSite = "None"
)

func (r JsonNewParamsBodyObjectCookiesSameSite) IsKnown() bool {
	switch r {
	case JsonNewParamsBodyObjectCookiesSameSiteStrict, JsonNewParamsBodyObjectCookiesSameSiteLax, JsonNewParamsBodyObjectCookiesSameSiteNone:
		return true
	}
	return false
}

type JsonNewParamsBodyObjectCookiesSourceScheme string

const (
	JsonNewParamsBodyObjectCookiesSourceSchemeUnset     JsonNewParamsBodyObjectCookiesSourceScheme = "Unset"
	JsonNewParamsBodyObjectCookiesSourceSchemeNonSecure JsonNewParamsBodyObjectCookiesSourceScheme = "NonSecure"
	JsonNewParamsBodyObjectCookiesSourceSchemeSecure    JsonNewParamsBodyObjectCookiesSourceScheme = "Secure"
)

func (r JsonNewParamsBodyObjectCookiesSourceScheme) IsKnown() bool {
	switch r {
	case JsonNewParamsBodyObjectCookiesSourceSchemeUnset, JsonNewParamsBodyObjectCookiesSourceSchemeNonSecure, JsonNewParamsBodyObjectCookiesSourceSchemeSecure:
		return true
	}
	return false
}

type JsonNewParamsBodyObjectCustomAI struct {
	// Authorization token for the AI model: `Bearer <token>`.
	Authorization param.Field[string] `json:"authorization,required"`
	// AI model to use for the request. Must be formed as `<provider>/<model_name>`,
	// e.g. `workers-ai/@cf/meta/llama-3.3-70b-instruct-fp8-fast`
	Model param.Field[string] `json:"model,required"`
}

func (r JsonNewParamsBodyObjectCustomAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type JsonNewParamsBodyObjectGotoOptions struct {
	Referer        param.Field[string]                                           `json:"referer"`
	ReferrerPolicy param.Field[string]                                           `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                          `json:"timeout"`
	WaitUntil      param.Field[JsonNewParamsBodyObjectGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r JsonNewParamsBodyObjectGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.JsonNewParamsBodyObjectGotoOptionsWaitUntilString],
// [browser_rendering.JsonNewParamsBodyObjectGotoOptionsWaitUntilArray].
type JsonNewParamsBodyObjectGotoOptionsWaitUntilUnion interface {
	implementsJsonNewParamsBodyObjectGotoOptionsWaitUntilUnion()
}

type JsonNewParamsBodyObjectGotoOptionsWaitUntilString string

const (
	JsonNewParamsBodyObjectGotoOptionsWaitUntilStringLoad             JsonNewParamsBodyObjectGotoOptionsWaitUntilString = "load"
	JsonNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded JsonNewParamsBodyObjectGotoOptionsWaitUntilString = "domcontentloaded"
	JsonNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0     JsonNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle0"
	JsonNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2     JsonNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle2"
)

func (r JsonNewParamsBodyObjectGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case JsonNewParamsBodyObjectGotoOptionsWaitUntilStringLoad, JsonNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded, JsonNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0, JsonNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r JsonNewParamsBodyObjectGotoOptionsWaitUntilString) implementsJsonNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type JsonNewParamsBodyObjectGotoOptionsWaitUntilArray []JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItem

func (r JsonNewParamsBodyObjectGotoOptionsWaitUntilArray) implementsJsonNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItem string

const (
	JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad             JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "load"
	JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0     JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle0"
	JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2     JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad, JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded, JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0, JsonNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type JsonNewParamsBodyObjectRejectResourceType string

const (
	JsonNewParamsBodyObjectRejectResourceTypeDocument           JsonNewParamsBodyObjectRejectResourceType = "document"
	JsonNewParamsBodyObjectRejectResourceTypeStylesheet         JsonNewParamsBodyObjectRejectResourceType = "stylesheet"
	JsonNewParamsBodyObjectRejectResourceTypeImage              JsonNewParamsBodyObjectRejectResourceType = "image"
	JsonNewParamsBodyObjectRejectResourceTypeMedia              JsonNewParamsBodyObjectRejectResourceType = "media"
	JsonNewParamsBodyObjectRejectResourceTypeFont               JsonNewParamsBodyObjectRejectResourceType = "font"
	JsonNewParamsBodyObjectRejectResourceTypeScript             JsonNewParamsBodyObjectRejectResourceType = "script"
	JsonNewParamsBodyObjectRejectResourceTypeTexttrack          JsonNewParamsBodyObjectRejectResourceType = "texttrack"
	JsonNewParamsBodyObjectRejectResourceTypeXHR                JsonNewParamsBodyObjectRejectResourceType = "xhr"
	JsonNewParamsBodyObjectRejectResourceTypeFetch              JsonNewParamsBodyObjectRejectResourceType = "fetch"
	JsonNewParamsBodyObjectRejectResourceTypePrefetch           JsonNewParamsBodyObjectRejectResourceType = "prefetch"
	JsonNewParamsBodyObjectRejectResourceTypeEventsource        JsonNewParamsBodyObjectRejectResourceType = "eventsource"
	JsonNewParamsBodyObjectRejectResourceTypeWebsocket          JsonNewParamsBodyObjectRejectResourceType = "websocket"
	JsonNewParamsBodyObjectRejectResourceTypeManifest           JsonNewParamsBodyObjectRejectResourceType = "manifest"
	JsonNewParamsBodyObjectRejectResourceTypeSignedexchange     JsonNewParamsBodyObjectRejectResourceType = "signedexchange"
	JsonNewParamsBodyObjectRejectResourceTypePing               JsonNewParamsBodyObjectRejectResourceType = "ping"
	JsonNewParamsBodyObjectRejectResourceTypeCspviolationreport JsonNewParamsBodyObjectRejectResourceType = "cspviolationreport"
	JsonNewParamsBodyObjectRejectResourceTypePreflight          JsonNewParamsBodyObjectRejectResourceType = "preflight"
	JsonNewParamsBodyObjectRejectResourceTypeOther              JsonNewParamsBodyObjectRejectResourceType = "other"
)

func (r JsonNewParamsBodyObjectRejectResourceType) IsKnown() bool {
	switch r {
	case JsonNewParamsBodyObjectRejectResourceTypeDocument, JsonNewParamsBodyObjectRejectResourceTypeStylesheet, JsonNewParamsBodyObjectRejectResourceTypeImage, JsonNewParamsBodyObjectRejectResourceTypeMedia, JsonNewParamsBodyObjectRejectResourceTypeFont, JsonNewParamsBodyObjectRejectResourceTypeScript, JsonNewParamsBodyObjectRejectResourceTypeTexttrack, JsonNewParamsBodyObjectRejectResourceTypeXHR, JsonNewParamsBodyObjectRejectResourceTypeFetch, JsonNewParamsBodyObjectRejectResourceTypePrefetch, JsonNewParamsBodyObjectRejectResourceTypeEventsource, JsonNewParamsBodyObjectRejectResourceTypeWebsocket, JsonNewParamsBodyObjectRejectResourceTypeManifest, JsonNewParamsBodyObjectRejectResourceTypeSignedexchange, JsonNewParamsBodyObjectRejectResourceTypePing, JsonNewParamsBodyObjectRejectResourceTypeCspviolationreport, JsonNewParamsBodyObjectRejectResourceTypePreflight, JsonNewParamsBodyObjectRejectResourceTypeOther:
		return true
	}
	return false
}

type JsonNewParamsBodyObjectResponseFormat struct {
	Type param.Field[string] `json:"type,required"`
	// Schema for the response format. More information here:
	// https://developers.cloudflare.com/workers-ai/json-mode/
	JsonSchema param.Field[map[string]JsonNewParamsBodyObjectResponseFormatJsonSchemaUnion] `json:"json_schema"`
}

func (r JsonNewParamsBodyObjectResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [browser_rendering.JsonNewParamsBodyObjectResponseFormatJsonSchemaArray].
//
// Use [Raw()] to specify an arbitrary value for this param
type JsonNewParamsBodyObjectResponseFormatJsonSchemaUnion interface {
	ImplementsJsonNewParamsBodyObjectResponseFormatJsonSchemaUnion()
}

type JsonNewParamsBodyObjectResponseFormatJsonSchemaArray []string

func (r JsonNewParamsBodyObjectResponseFormatJsonSchemaArray) ImplementsJsonNewParamsBodyObjectResponseFormatJsonSchemaUnion() {
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type JsonNewParamsBodyObjectViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r JsonNewParamsBodyObjectViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type JsonNewParamsBodyObjectWaitForSelector struct {
	Selector param.Field[string]                                        `json:"selector,required"`
	Hidden   param.Field[JsonNewParamsBodyObjectWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                       `json:"timeout"`
	Visible  param.Field[JsonNewParamsBodyObjectWaitForSelectorVisible] `json:"visible"`
}

func (r JsonNewParamsBodyObjectWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JsonNewParamsBodyObjectWaitForSelectorHidden bool

const (
	JsonNewParamsBodyObjectWaitForSelectorHiddenTrue JsonNewParamsBodyObjectWaitForSelectorHidden = true
)

func (r JsonNewParamsBodyObjectWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case JsonNewParamsBodyObjectWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type JsonNewParamsBodyObjectWaitForSelectorVisible bool

const (
	JsonNewParamsBodyObjectWaitForSelectorVisibleTrue JsonNewParamsBodyObjectWaitForSelectorVisible = true
)

func (r JsonNewParamsBodyObjectWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case JsonNewParamsBodyObjectWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type JsonNewResponseEnvelope struct {
	Result JsonNewResponse `json:"result,required"`
	// Response status
	Success bool                            `json:"success,required"`
	Errors  []JsonNewResponseEnvelopeErrors `json:"errors"`
	JSON    jsonNewResponseEnvelopeJSON     `json:"-"`
}

// jsonNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [JsonNewResponseEnvelope]
type jsonNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JsonNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jsonNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type JsonNewResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                            `json:"message,required"`
	JSON    jsonNewResponseEnvelopeErrorsJSON `json:"-"`
}

// jsonNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [JsonNewResponseEnvelopeErrors]
type jsonNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JsonNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jsonNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
