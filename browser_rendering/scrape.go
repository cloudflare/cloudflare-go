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

// ScrapeService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScrapeService] method instead.
type ScrapeService struct {
	Options []option.RequestOption
}

// NewScrapeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScrapeService(opts ...option.RequestOption) (r *ScrapeService) {
	r = &ScrapeService{}
	r.Options = opts
	return
}

// Get meta attributes like height, width, text and others of selected elements.
func (r *ScrapeService) New(ctx context.Context, params ScrapeNewParams, opts ...option.RequestOption) (res *[]ScrapeNewResponse, err error) {
	var env ScrapeNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/scrape", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScrapeNewResponse struct {
	Results ScrapeNewResponseResults `json:"results,required"`
	// Selector
	Selector string                `json:"selector,required"`
	JSON     scrapeNewResponseJSON `json:"-"`
}

// scrapeNewResponseJSON contains the JSON metadata for the struct
// [ScrapeNewResponse]
type scrapeNewResponseJSON struct {
	Results     apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScrapeNewResponseResults struct {
	Attributes []ScrapeNewResponseResultsAttribute `json:"attributes,required"`
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
	Width float64                      `json:"width,required"`
	JSON  scrapeNewResponseResultsJSON `json:"-"`
}

// scrapeNewResponseResultsJSON contains the JSON metadata for the struct
// [ScrapeNewResponseResults]
type scrapeNewResponseResultsJSON struct {
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

func (r *ScrapeNewResponseResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeNewResponseResultsJSON) RawJSON() string {
	return r.raw
}

type ScrapeNewResponseResultsAttribute struct {
	// Attribute name
	Name string `json:"name,required"`
	// Attribute value
	Value string                                `json:"value,required"`
	JSON  scrapeNewResponseResultsAttributeJSON `json:"-"`
}

// scrapeNewResponseResultsAttributeJSON contains the JSON metadata for the struct
// [ScrapeNewResponseResultsAttribute]
type scrapeNewResponseResultsAttributeJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeNewResponseResultsAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeNewResponseResultsAttributeJSON) RawJSON() string {
	return r.raw
}

type ScrapeNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64]     `query:"cacheTTL"`
	Body     ScrapeNewParamsBodyUnion `json:"body"`
}

func (r ScrapeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [ScrapeNewParams]'s query parameters as `url.Values`.
func (r ScrapeNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScrapeNewParamsBody struct {
	Elements param.Field[interface{}] `json:"elements,required"`
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

func (r ScrapeNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScrapeNewParamsBody) implementsScrapeNewParamsBodyUnion() {}

// Satisfied by [browser_rendering.ScrapeNewParamsBodyObject],
// [browser_rendering.ScrapeNewParamsBodyObject], [ScrapeNewParamsBody].
type ScrapeNewParamsBodyUnion interface {
	implementsScrapeNewParamsBodyUnion()
}

type ScrapeNewParamsBodyObject struct {
	Elements param.Field[[]ScrapeNewParamsBodyObjectElement] `json:"elements,required"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html,required"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]ScrapeNewParamsBodyObjectAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]ScrapeNewParamsBodyObjectAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]ScrapeNewParamsBodyObjectAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[ScrapeNewParamsBodyObjectAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]ScrapeNewParamsBodyObjectCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                            `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[ScrapeNewParamsBodyObjectGotoOptions] `json:"gotoOptions"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]ScrapeNewParamsBodyObjectRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                             `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                          `json:"setJavaScriptEnabled"`
	UserAgent            param.Field[string]                                        `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[ScrapeNewParamsBodyObjectViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[ScrapeNewParamsBodyObjectWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r ScrapeNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScrapeNewParamsBodyObject) implementsScrapeNewParamsBodyUnion() {}

type ScrapeNewParamsBodyObjectElement struct {
	Selector param.Field[string] `json:"selector,required"`
}

func (r ScrapeNewParamsBodyObjectElement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsBodyObjectAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r ScrapeNewParamsBodyObjectAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsBodyObjectAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r ScrapeNewParamsBodyObjectAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsBodyObjectAllowResourceType string

const (
	ScrapeNewParamsBodyObjectAllowResourceTypeDocument           ScrapeNewParamsBodyObjectAllowResourceType = "document"
	ScrapeNewParamsBodyObjectAllowResourceTypeStylesheet         ScrapeNewParamsBodyObjectAllowResourceType = "stylesheet"
	ScrapeNewParamsBodyObjectAllowResourceTypeImage              ScrapeNewParamsBodyObjectAllowResourceType = "image"
	ScrapeNewParamsBodyObjectAllowResourceTypeMedia              ScrapeNewParamsBodyObjectAllowResourceType = "media"
	ScrapeNewParamsBodyObjectAllowResourceTypeFont               ScrapeNewParamsBodyObjectAllowResourceType = "font"
	ScrapeNewParamsBodyObjectAllowResourceTypeScript             ScrapeNewParamsBodyObjectAllowResourceType = "script"
	ScrapeNewParamsBodyObjectAllowResourceTypeTexttrack          ScrapeNewParamsBodyObjectAllowResourceType = "texttrack"
	ScrapeNewParamsBodyObjectAllowResourceTypeXHR                ScrapeNewParamsBodyObjectAllowResourceType = "xhr"
	ScrapeNewParamsBodyObjectAllowResourceTypeFetch              ScrapeNewParamsBodyObjectAllowResourceType = "fetch"
	ScrapeNewParamsBodyObjectAllowResourceTypePrefetch           ScrapeNewParamsBodyObjectAllowResourceType = "prefetch"
	ScrapeNewParamsBodyObjectAllowResourceTypeEventsource        ScrapeNewParamsBodyObjectAllowResourceType = "eventsource"
	ScrapeNewParamsBodyObjectAllowResourceTypeWebsocket          ScrapeNewParamsBodyObjectAllowResourceType = "websocket"
	ScrapeNewParamsBodyObjectAllowResourceTypeManifest           ScrapeNewParamsBodyObjectAllowResourceType = "manifest"
	ScrapeNewParamsBodyObjectAllowResourceTypeSignedexchange     ScrapeNewParamsBodyObjectAllowResourceType = "signedexchange"
	ScrapeNewParamsBodyObjectAllowResourceTypePing               ScrapeNewParamsBodyObjectAllowResourceType = "ping"
	ScrapeNewParamsBodyObjectAllowResourceTypeCspviolationreport ScrapeNewParamsBodyObjectAllowResourceType = "cspviolationreport"
	ScrapeNewParamsBodyObjectAllowResourceTypePreflight          ScrapeNewParamsBodyObjectAllowResourceType = "preflight"
	ScrapeNewParamsBodyObjectAllowResourceTypeOther              ScrapeNewParamsBodyObjectAllowResourceType = "other"
)

func (r ScrapeNewParamsBodyObjectAllowResourceType) IsKnown() bool {
	switch r {
	case ScrapeNewParamsBodyObjectAllowResourceTypeDocument, ScrapeNewParamsBodyObjectAllowResourceTypeStylesheet, ScrapeNewParamsBodyObjectAllowResourceTypeImage, ScrapeNewParamsBodyObjectAllowResourceTypeMedia, ScrapeNewParamsBodyObjectAllowResourceTypeFont, ScrapeNewParamsBodyObjectAllowResourceTypeScript, ScrapeNewParamsBodyObjectAllowResourceTypeTexttrack, ScrapeNewParamsBodyObjectAllowResourceTypeXHR, ScrapeNewParamsBodyObjectAllowResourceTypeFetch, ScrapeNewParamsBodyObjectAllowResourceTypePrefetch, ScrapeNewParamsBodyObjectAllowResourceTypeEventsource, ScrapeNewParamsBodyObjectAllowResourceTypeWebsocket, ScrapeNewParamsBodyObjectAllowResourceTypeManifest, ScrapeNewParamsBodyObjectAllowResourceTypeSignedexchange, ScrapeNewParamsBodyObjectAllowResourceTypePing, ScrapeNewParamsBodyObjectAllowResourceTypeCspviolationreport, ScrapeNewParamsBodyObjectAllowResourceTypePreflight, ScrapeNewParamsBodyObjectAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type ScrapeNewParamsBodyObjectAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r ScrapeNewParamsBodyObjectAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsBodyObjectCookie struct {
	Name         param.Field[string]                                       `json:"name,required"`
	Value        param.Field[string]                                       `json:"value,required"`
	Domain       param.Field[string]                                       `json:"domain"`
	Expires      param.Field[float64]                                      `json:"expires"`
	HTTPOnly     param.Field[bool]                                         `json:"httpOnly"`
	PartitionKey param.Field[string]                                       `json:"partitionKey"`
	Path         param.Field[string]                                       `json:"path"`
	Priority     param.Field[ScrapeNewParamsBodyObjectCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                         `json:"sameParty"`
	SameSite     param.Field[ScrapeNewParamsBodyObjectCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                         `json:"secure"`
	SourcePort   param.Field[float64]                                      `json:"sourcePort"`
	SourceScheme param.Field[ScrapeNewParamsBodyObjectCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                       `json:"url"`
}

func (r ScrapeNewParamsBodyObjectCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsBodyObjectCookiesPriority string

const (
	ScrapeNewParamsBodyObjectCookiesPriorityLow    ScrapeNewParamsBodyObjectCookiesPriority = "Low"
	ScrapeNewParamsBodyObjectCookiesPriorityMedium ScrapeNewParamsBodyObjectCookiesPriority = "Medium"
	ScrapeNewParamsBodyObjectCookiesPriorityHigh   ScrapeNewParamsBodyObjectCookiesPriority = "High"
)

func (r ScrapeNewParamsBodyObjectCookiesPriority) IsKnown() bool {
	switch r {
	case ScrapeNewParamsBodyObjectCookiesPriorityLow, ScrapeNewParamsBodyObjectCookiesPriorityMedium, ScrapeNewParamsBodyObjectCookiesPriorityHigh:
		return true
	}
	return false
}

type ScrapeNewParamsBodyObjectCookiesSameSite string

const (
	ScrapeNewParamsBodyObjectCookiesSameSiteStrict ScrapeNewParamsBodyObjectCookiesSameSite = "Strict"
	ScrapeNewParamsBodyObjectCookiesSameSiteLax    ScrapeNewParamsBodyObjectCookiesSameSite = "Lax"
	ScrapeNewParamsBodyObjectCookiesSameSiteNone   ScrapeNewParamsBodyObjectCookiesSameSite = "None"
)

func (r ScrapeNewParamsBodyObjectCookiesSameSite) IsKnown() bool {
	switch r {
	case ScrapeNewParamsBodyObjectCookiesSameSiteStrict, ScrapeNewParamsBodyObjectCookiesSameSiteLax, ScrapeNewParamsBodyObjectCookiesSameSiteNone:
		return true
	}
	return false
}

type ScrapeNewParamsBodyObjectCookiesSourceScheme string

const (
	ScrapeNewParamsBodyObjectCookiesSourceSchemeUnset     ScrapeNewParamsBodyObjectCookiesSourceScheme = "Unset"
	ScrapeNewParamsBodyObjectCookiesSourceSchemeNonSecure ScrapeNewParamsBodyObjectCookiesSourceScheme = "NonSecure"
	ScrapeNewParamsBodyObjectCookiesSourceSchemeSecure    ScrapeNewParamsBodyObjectCookiesSourceScheme = "Secure"
)

func (r ScrapeNewParamsBodyObjectCookiesSourceScheme) IsKnown() bool {
	switch r {
	case ScrapeNewParamsBodyObjectCookiesSourceSchemeUnset, ScrapeNewParamsBodyObjectCookiesSourceSchemeNonSecure, ScrapeNewParamsBodyObjectCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type ScrapeNewParamsBodyObjectGotoOptions struct {
	Referer        param.Field[string]                                             `json:"referer"`
	ReferrerPolicy param.Field[string]                                             `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                            `json:"timeout"`
	WaitUntil      param.Field[ScrapeNewParamsBodyObjectGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r ScrapeNewParamsBodyObjectGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.ScrapeNewParamsBodyObjectGotoOptionsWaitUntilString],
// [browser_rendering.ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArray].
type ScrapeNewParamsBodyObjectGotoOptionsWaitUntilUnion interface {
	implementsScrapeNewParamsBodyObjectGotoOptionsWaitUntilUnion()
}

type ScrapeNewParamsBodyObjectGotoOptionsWaitUntilString string

const (
	ScrapeNewParamsBodyObjectGotoOptionsWaitUntilStringLoad             ScrapeNewParamsBodyObjectGotoOptionsWaitUntilString = "load"
	ScrapeNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded ScrapeNewParamsBodyObjectGotoOptionsWaitUntilString = "domcontentloaded"
	ScrapeNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0     ScrapeNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle0"
	ScrapeNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2     ScrapeNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle2"
)

func (r ScrapeNewParamsBodyObjectGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case ScrapeNewParamsBodyObjectGotoOptionsWaitUntilStringLoad, ScrapeNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded, ScrapeNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0, ScrapeNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r ScrapeNewParamsBodyObjectGotoOptionsWaitUntilString) implementsScrapeNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArray []ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem

func (r ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArray) implementsScrapeNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem string

const (
	ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad             ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "load"
	ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0     ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle0"
	ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2     ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad, ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded, ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0, ScrapeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type ScrapeNewParamsBodyObjectRejectResourceType string

const (
	ScrapeNewParamsBodyObjectRejectResourceTypeDocument           ScrapeNewParamsBodyObjectRejectResourceType = "document"
	ScrapeNewParamsBodyObjectRejectResourceTypeStylesheet         ScrapeNewParamsBodyObjectRejectResourceType = "stylesheet"
	ScrapeNewParamsBodyObjectRejectResourceTypeImage              ScrapeNewParamsBodyObjectRejectResourceType = "image"
	ScrapeNewParamsBodyObjectRejectResourceTypeMedia              ScrapeNewParamsBodyObjectRejectResourceType = "media"
	ScrapeNewParamsBodyObjectRejectResourceTypeFont               ScrapeNewParamsBodyObjectRejectResourceType = "font"
	ScrapeNewParamsBodyObjectRejectResourceTypeScript             ScrapeNewParamsBodyObjectRejectResourceType = "script"
	ScrapeNewParamsBodyObjectRejectResourceTypeTexttrack          ScrapeNewParamsBodyObjectRejectResourceType = "texttrack"
	ScrapeNewParamsBodyObjectRejectResourceTypeXHR                ScrapeNewParamsBodyObjectRejectResourceType = "xhr"
	ScrapeNewParamsBodyObjectRejectResourceTypeFetch              ScrapeNewParamsBodyObjectRejectResourceType = "fetch"
	ScrapeNewParamsBodyObjectRejectResourceTypePrefetch           ScrapeNewParamsBodyObjectRejectResourceType = "prefetch"
	ScrapeNewParamsBodyObjectRejectResourceTypeEventsource        ScrapeNewParamsBodyObjectRejectResourceType = "eventsource"
	ScrapeNewParamsBodyObjectRejectResourceTypeWebsocket          ScrapeNewParamsBodyObjectRejectResourceType = "websocket"
	ScrapeNewParamsBodyObjectRejectResourceTypeManifest           ScrapeNewParamsBodyObjectRejectResourceType = "manifest"
	ScrapeNewParamsBodyObjectRejectResourceTypeSignedexchange     ScrapeNewParamsBodyObjectRejectResourceType = "signedexchange"
	ScrapeNewParamsBodyObjectRejectResourceTypePing               ScrapeNewParamsBodyObjectRejectResourceType = "ping"
	ScrapeNewParamsBodyObjectRejectResourceTypeCspviolationreport ScrapeNewParamsBodyObjectRejectResourceType = "cspviolationreport"
	ScrapeNewParamsBodyObjectRejectResourceTypePreflight          ScrapeNewParamsBodyObjectRejectResourceType = "preflight"
	ScrapeNewParamsBodyObjectRejectResourceTypeOther              ScrapeNewParamsBodyObjectRejectResourceType = "other"
)

func (r ScrapeNewParamsBodyObjectRejectResourceType) IsKnown() bool {
	switch r {
	case ScrapeNewParamsBodyObjectRejectResourceTypeDocument, ScrapeNewParamsBodyObjectRejectResourceTypeStylesheet, ScrapeNewParamsBodyObjectRejectResourceTypeImage, ScrapeNewParamsBodyObjectRejectResourceTypeMedia, ScrapeNewParamsBodyObjectRejectResourceTypeFont, ScrapeNewParamsBodyObjectRejectResourceTypeScript, ScrapeNewParamsBodyObjectRejectResourceTypeTexttrack, ScrapeNewParamsBodyObjectRejectResourceTypeXHR, ScrapeNewParamsBodyObjectRejectResourceTypeFetch, ScrapeNewParamsBodyObjectRejectResourceTypePrefetch, ScrapeNewParamsBodyObjectRejectResourceTypeEventsource, ScrapeNewParamsBodyObjectRejectResourceTypeWebsocket, ScrapeNewParamsBodyObjectRejectResourceTypeManifest, ScrapeNewParamsBodyObjectRejectResourceTypeSignedexchange, ScrapeNewParamsBodyObjectRejectResourceTypePing, ScrapeNewParamsBodyObjectRejectResourceTypeCspviolationreport, ScrapeNewParamsBodyObjectRejectResourceTypePreflight, ScrapeNewParamsBodyObjectRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type ScrapeNewParamsBodyObjectViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r ScrapeNewParamsBodyObjectViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type ScrapeNewParamsBodyObjectWaitForSelector struct {
	Selector param.Field[string]                                          `json:"selector,required"`
	Hidden   param.Field[ScrapeNewParamsBodyObjectWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                         `json:"timeout"`
	Visible  param.Field[ScrapeNewParamsBodyObjectWaitForSelectorVisible] `json:"visible"`
}

func (r ScrapeNewParamsBodyObjectWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsBodyObjectWaitForSelectorHidden bool

const (
	ScrapeNewParamsBodyObjectWaitForSelectorHiddenTrue ScrapeNewParamsBodyObjectWaitForSelectorHidden = true
)

func (r ScrapeNewParamsBodyObjectWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case ScrapeNewParamsBodyObjectWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type ScrapeNewParamsBodyObjectWaitForSelectorVisible bool

const (
	ScrapeNewParamsBodyObjectWaitForSelectorVisibleTrue ScrapeNewParamsBodyObjectWaitForSelectorVisible = true
)

func (r ScrapeNewParamsBodyObjectWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case ScrapeNewParamsBodyObjectWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type ScrapeNewResponseEnvelope struct {
	Result []ScrapeNewResponse `json:"result,required"`
	// Response status
	Success bool                              `json:"success,required"`
	Errors  []ScrapeNewResponseEnvelopeErrors `json:"errors"`
	JSON    scrapeNewResponseEnvelopeJSON     `json:"-"`
}

// scrapeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScrapeNewResponseEnvelope]
type scrapeNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScrapeNewResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                              `json:"message,required"`
	JSON    scrapeNewResponseEnvelopeErrorsJSON `json:"-"`
}

// scrapeNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ScrapeNewResponseEnvelopeErrors]
type scrapeNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
