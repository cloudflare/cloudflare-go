// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browsing_rendering

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
// `goToOptions` and `waitFor*` to control page load behaviour.
func (r *PDFService) New(ctx context.Context, accountID string, params PDFNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
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

type PDFNewParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]PDFNewParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]PDFNewParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]PDFNewParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[PDFNewParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]PDFNewParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]               `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[PDFNewParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]PDFNewParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                             `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[PDFNewParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[PDFNewParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r PDFNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [PDFNewParams]'s query parameters as `url.Values`.
func (r PDFNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PDFNewParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r PDFNewParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PDFNewParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r PDFNewParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PDFNewParamsAllowResourceType string

const (
	PDFNewParamsAllowResourceTypeDocument           PDFNewParamsAllowResourceType = "document"
	PDFNewParamsAllowResourceTypeStylesheet         PDFNewParamsAllowResourceType = "stylesheet"
	PDFNewParamsAllowResourceTypeImage              PDFNewParamsAllowResourceType = "image"
	PDFNewParamsAllowResourceTypeMedia              PDFNewParamsAllowResourceType = "media"
	PDFNewParamsAllowResourceTypeFont               PDFNewParamsAllowResourceType = "font"
	PDFNewParamsAllowResourceTypeScript             PDFNewParamsAllowResourceType = "script"
	PDFNewParamsAllowResourceTypeTexttrack          PDFNewParamsAllowResourceType = "texttrack"
	PDFNewParamsAllowResourceTypeXhr                PDFNewParamsAllowResourceType = "xhr"
	PDFNewParamsAllowResourceTypeFetch              PDFNewParamsAllowResourceType = "fetch"
	PDFNewParamsAllowResourceTypePrefetch           PDFNewParamsAllowResourceType = "prefetch"
	PDFNewParamsAllowResourceTypeEventsource        PDFNewParamsAllowResourceType = "eventsource"
	PDFNewParamsAllowResourceTypeWebsocket          PDFNewParamsAllowResourceType = "websocket"
	PDFNewParamsAllowResourceTypeManifest           PDFNewParamsAllowResourceType = "manifest"
	PDFNewParamsAllowResourceTypeSignedexchange     PDFNewParamsAllowResourceType = "signedexchange"
	PDFNewParamsAllowResourceTypePing               PDFNewParamsAllowResourceType = "ping"
	PDFNewParamsAllowResourceTypeCspviolationreport PDFNewParamsAllowResourceType = "cspviolationreport"
	PDFNewParamsAllowResourceTypePreflight          PDFNewParamsAllowResourceType = "preflight"
	PDFNewParamsAllowResourceTypeOther              PDFNewParamsAllowResourceType = "other"
)

func (r PDFNewParamsAllowResourceType) IsKnown() bool {
	switch r {
	case PDFNewParamsAllowResourceTypeDocument, PDFNewParamsAllowResourceTypeStylesheet, PDFNewParamsAllowResourceTypeImage, PDFNewParamsAllowResourceTypeMedia, PDFNewParamsAllowResourceTypeFont, PDFNewParamsAllowResourceTypeScript, PDFNewParamsAllowResourceTypeTexttrack, PDFNewParamsAllowResourceTypeXhr, PDFNewParamsAllowResourceTypeFetch, PDFNewParamsAllowResourceTypePrefetch, PDFNewParamsAllowResourceTypeEventsource, PDFNewParamsAllowResourceTypeWebsocket, PDFNewParamsAllowResourceTypeManifest, PDFNewParamsAllowResourceTypeSignedexchange, PDFNewParamsAllowResourceTypePing, PDFNewParamsAllowResourceTypeCspviolationreport, PDFNewParamsAllowResourceTypePreflight, PDFNewParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type PDFNewParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r PDFNewParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PDFNewParamsCookie struct {
	Name         param.Field[string]                          `json:"name,required"`
	Value        param.Field[string]                          `json:"value,required"`
	Domain       param.Field[string]                          `json:"domain"`
	Expires      param.Field[float64]                         `json:"expires"`
	HTTPOnly     param.Field[bool]                            `json:"httpOnly"`
	PartitionKey param.Field[string]                          `json:"partitionKey"`
	Path         param.Field[string]                          `json:"path"`
	Priority     param.Field[PDFNewParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                            `json:"sameParty"`
	SameSite     param.Field[PDFNewParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                            `json:"secure"`
	SourcePort   param.Field[float64]                         `json:"sourcePort"`
	SourceScheme param.Field[PDFNewParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                          `json:"url"`
}

func (r PDFNewParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PDFNewParamsCookiesPriority string

const (
	PDFNewParamsCookiesPriorityLow    PDFNewParamsCookiesPriority = "Low"
	PDFNewParamsCookiesPriorityMedium PDFNewParamsCookiesPriority = "Medium"
	PDFNewParamsCookiesPriorityHigh   PDFNewParamsCookiesPriority = "High"
)

func (r PDFNewParamsCookiesPriority) IsKnown() bool {
	switch r {
	case PDFNewParamsCookiesPriorityLow, PDFNewParamsCookiesPriorityMedium, PDFNewParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type PDFNewParamsCookiesSameSite string

const (
	PDFNewParamsCookiesSameSiteStrict PDFNewParamsCookiesSameSite = "Strict"
	PDFNewParamsCookiesSameSiteLax    PDFNewParamsCookiesSameSite = "Lax"
	PDFNewParamsCookiesSameSiteNone   PDFNewParamsCookiesSameSite = "None"
)

func (r PDFNewParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case PDFNewParamsCookiesSameSiteStrict, PDFNewParamsCookiesSameSiteLax, PDFNewParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type PDFNewParamsCookiesSourceScheme string

const (
	PDFNewParamsCookiesSourceSchemeUnset     PDFNewParamsCookiesSourceScheme = "Unset"
	PDFNewParamsCookiesSourceSchemeNonSecure PDFNewParamsCookiesSourceScheme = "NonSecure"
	PDFNewParamsCookiesSourceSchemeSecure    PDFNewParamsCookiesSourceScheme = "Secure"
)

func (r PDFNewParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case PDFNewParamsCookiesSourceSchemeUnset, PDFNewParamsCookiesSourceSchemeNonSecure, PDFNewParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type PDFNewParamsGotoOptions struct {
	Referer        param.Field[string]                                `json:"referer"`
	ReferrerPolicy param.Field[string]                                `json:"referrerPolicy"`
	Timeout        param.Field[float64]                               `json:"timeout"`
	WaitUntil      param.Field[PDFNewParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r PDFNewParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [browsing_rendering.PDFNewParamsGotoOptionsWaitUntilString],
// [browsing_rendering.PDFNewParamsGotoOptionsWaitUntilArray].
type PDFNewParamsGotoOptionsWaitUntilUnion interface {
	implementsPDFNewParamsGotoOptionsWaitUntilUnion()
}

type PDFNewParamsGotoOptionsWaitUntilString string

const (
	PDFNewParamsGotoOptionsWaitUntilStringLoad             PDFNewParamsGotoOptionsWaitUntilString = "load"
	PDFNewParamsGotoOptionsWaitUntilStringDomcontentloaded PDFNewParamsGotoOptionsWaitUntilString = "domcontentloaded"
	PDFNewParamsGotoOptionsWaitUntilStringNetworkidle0     PDFNewParamsGotoOptionsWaitUntilString = "networkidle0"
	PDFNewParamsGotoOptionsWaitUntilStringNetworkidle2     PDFNewParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r PDFNewParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case PDFNewParamsGotoOptionsWaitUntilStringLoad, PDFNewParamsGotoOptionsWaitUntilStringDomcontentloaded, PDFNewParamsGotoOptionsWaitUntilStringNetworkidle0, PDFNewParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r PDFNewParamsGotoOptionsWaitUntilString) implementsPDFNewParamsGotoOptionsWaitUntilUnion() {}

type PDFNewParamsGotoOptionsWaitUntilArray []PDFNewParamsGotoOptionsWaitUntilArrayItem

func (r PDFNewParamsGotoOptionsWaitUntilArray) implementsPDFNewParamsGotoOptionsWaitUntilUnion() {}

type PDFNewParamsGotoOptionsWaitUntilArrayItem string

const (
	PDFNewParamsGotoOptionsWaitUntilArrayItemLoad             PDFNewParamsGotoOptionsWaitUntilArrayItem = "load"
	PDFNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded PDFNewParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	PDFNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0     PDFNewParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	PDFNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2     PDFNewParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r PDFNewParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case PDFNewParamsGotoOptionsWaitUntilArrayItemLoad, PDFNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, PDFNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0, PDFNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type PDFNewParamsRejectResourceType string

const (
	PDFNewParamsRejectResourceTypeDocument           PDFNewParamsRejectResourceType = "document"
	PDFNewParamsRejectResourceTypeStylesheet         PDFNewParamsRejectResourceType = "stylesheet"
	PDFNewParamsRejectResourceTypeImage              PDFNewParamsRejectResourceType = "image"
	PDFNewParamsRejectResourceTypeMedia              PDFNewParamsRejectResourceType = "media"
	PDFNewParamsRejectResourceTypeFont               PDFNewParamsRejectResourceType = "font"
	PDFNewParamsRejectResourceTypeScript             PDFNewParamsRejectResourceType = "script"
	PDFNewParamsRejectResourceTypeTexttrack          PDFNewParamsRejectResourceType = "texttrack"
	PDFNewParamsRejectResourceTypeXhr                PDFNewParamsRejectResourceType = "xhr"
	PDFNewParamsRejectResourceTypeFetch              PDFNewParamsRejectResourceType = "fetch"
	PDFNewParamsRejectResourceTypePrefetch           PDFNewParamsRejectResourceType = "prefetch"
	PDFNewParamsRejectResourceTypeEventsource        PDFNewParamsRejectResourceType = "eventsource"
	PDFNewParamsRejectResourceTypeWebsocket          PDFNewParamsRejectResourceType = "websocket"
	PDFNewParamsRejectResourceTypeManifest           PDFNewParamsRejectResourceType = "manifest"
	PDFNewParamsRejectResourceTypeSignedexchange     PDFNewParamsRejectResourceType = "signedexchange"
	PDFNewParamsRejectResourceTypePing               PDFNewParamsRejectResourceType = "ping"
	PDFNewParamsRejectResourceTypeCspviolationreport PDFNewParamsRejectResourceType = "cspviolationreport"
	PDFNewParamsRejectResourceTypePreflight          PDFNewParamsRejectResourceType = "preflight"
	PDFNewParamsRejectResourceTypeOther              PDFNewParamsRejectResourceType = "other"
)

func (r PDFNewParamsRejectResourceType) IsKnown() bool {
	switch r {
	case PDFNewParamsRejectResourceTypeDocument, PDFNewParamsRejectResourceTypeStylesheet, PDFNewParamsRejectResourceTypeImage, PDFNewParamsRejectResourceTypeMedia, PDFNewParamsRejectResourceTypeFont, PDFNewParamsRejectResourceTypeScript, PDFNewParamsRejectResourceTypeTexttrack, PDFNewParamsRejectResourceTypeXhr, PDFNewParamsRejectResourceTypeFetch, PDFNewParamsRejectResourceTypePrefetch, PDFNewParamsRejectResourceTypeEventsource, PDFNewParamsRejectResourceTypeWebsocket, PDFNewParamsRejectResourceTypeManifest, PDFNewParamsRejectResourceTypeSignedexchange, PDFNewParamsRejectResourceTypePing, PDFNewParamsRejectResourceTypeCspviolationreport, PDFNewParamsRejectResourceTypePreflight, PDFNewParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type PDFNewParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r PDFNewParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type PDFNewParamsWaitForSelector struct {
	Selector param.Field[string]                             `json:"selector,required"`
	Hidden   param.Field[PDFNewParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                            `json:"timeout"`
	Visible  param.Field[PDFNewParamsWaitForSelectorVisible] `json:"visible"`
}

func (r PDFNewParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PDFNewParamsWaitForSelectorHidden bool

const (
	PDFNewParamsWaitForSelectorHiddenTrue PDFNewParamsWaitForSelectorHidden = true
)

func (r PDFNewParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case PDFNewParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type PDFNewParamsWaitForSelectorVisible bool

const (
	PDFNewParamsWaitForSelectorVisibleTrue PDFNewParamsWaitForSelectorVisible = true
)

func (r PDFNewParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case PDFNewParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}
