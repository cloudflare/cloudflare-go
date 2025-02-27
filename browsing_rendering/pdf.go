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

// PdfService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPdfService] method instead.
type PdfService struct {
	Options []option.RequestOption
}

// NewPdfService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPdfService(opts ...option.RequestOption) (r *PdfService) {
	r = &PdfService{}
	r.Options = opts
	return
}

// Fetches rendered PDF from provided URL or HTML. Check available options like
// `goToOptions` and `waitFor*` to control page load behaviour.
func (r *PdfService) New(ctx context.Context, accountID string, params PdfNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
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

type PdfNewParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]PdfNewParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]PdfNewParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]PdfNewParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[PdfNewParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]PdfNewParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]               `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[PdfNewParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]PdfNewParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                             `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[PdfNewParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[PdfNewParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r PdfNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [PdfNewParams]'s query parameters as `url.Values`.
func (r PdfNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PdfNewParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r PdfNewParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PdfNewParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r PdfNewParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PdfNewParamsAllowResourceType string

const (
	PdfNewParamsAllowResourceTypeDocument           PdfNewParamsAllowResourceType = "document"
	PdfNewParamsAllowResourceTypeStylesheet         PdfNewParamsAllowResourceType = "stylesheet"
	PdfNewParamsAllowResourceTypeImage              PdfNewParamsAllowResourceType = "image"
	PdfNewParamsAllowResourceTypeMedia              PdfNewParamsAllowResourceType = "media"
	PdfNewParamsAllowResourceTypeFont               PdfNewParamsAllowResourceType = "font"
	PdfNewParamsAllowResourceTypeScript             PdfNewParamsAllowResourceType = "script"
	PdfNewParamsAllowResourceTypeTexttrack          PdfNewParamsAllowResourceType = "texttrack"
	PdfNewParamsAllowResourceTypeXhr                PdfNewParamsAllowResourceType = "xhr"
	PdfNewParamsAllowResourceTypeFetch              PdfNewParamsAllowResourceType = "fetch"
	PdfNewParamsAllowResourceTypePrefetch           PdfNewParamsAllowResourceType = "prefetch"
	PdfNewParamsAllowResourceTypeEventsource        PdfNewParamsAllowResourceType = "eventsource"
	PdfNewParamsAllowResourceTypeWebsocket          PdfNewParamsAllowResourceType = "websocket"
	PdfNewParamsAllowResourceTypeManifest           PdfNewParamsAllowResourceType = "manifest"
	PdfNewParamsAllowResourceTypeSignedexchange     PdfNewParamsAllowResourceType = "signedexchange"
	PdfNewParamsAllowResourceTypePing               PdfNewParamsAllowResourceType = "ping"
	PdfNewParamsAllowResourceTypeCspviolationreport PdfNewParamsAllowResourceType = "cspviolationreport"
	PdfNewParamsAllowResourceTypePreflight          PdfNewParamsAllowResourceType = "preflight"
	PdfNewParamsAllowResourceTypeOther              PdfNewParamsAllowResourceType = "other"
)

func (r PdfNewParamsAllowResourceType) IsKnown() bool {
	switch r {
	case PdfNewParamsAllowResourceTypeDocument, PdfNewParamsAllowResourceTypeStylesheet, PdfNewParamsAllowResourceTypeImage, PdfNewParamsAllowResourceTypeMedia, PdfNewParamsAllowResourceTypeFont, PdfNewParamsAllowResourceTypeScript, PdfNewParamsAllowResourceTypeTexttrack, PdfNewParamsAllowResourceTypeXhr, PdfNewParamsAllowResourceTypeFetch, PdfNewParamsAllowResourceTypePrefetch, PdfNewParamsAllowResourceTypeEventsource, PdfNewParamsAllowResourceTypeWebsocket, PdfNewParamsAllowResourceTypeManifest, PdfNewParamsAllowResourceTypeSignedexchange, PdfNewParamsAllowResourceTypePing, PdfNewParamsAllowResourceTypeCspviolationreport, PdfNewParamsAllowResourceTypePreflight, PdfNewParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type PdfNewParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r PdfNewParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PdfNewParamsCookie struct {
	Name         param.Field[string]                          `json:"name,required"`
	Value        param.Field[string]                          `json:"value,required"`
	Domain       param.Field[string]                          `json:"domain"`
	Expires      param.Field[float64]                         `json:"expires"`
	HTTPOnly     param.Field[bool]                            `json:"httpOnly"`
	PartitionKey param.Field[string]                          `json:"partitionKey"`
	Path         param.Field[string]                          `json:"path"`
	Priority     param.Field[PdfNewParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                            `json:"sameParty"`
	SameSite     param.Field[PdfNewParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                            `json:"secure"`
	SourcePort   param.Field[float64]                         `json:"sourcePort"`
	SourceScheme param.Field[PdfNewParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                          `json:"url"`
}

func (r PdfNewParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PdfNewParamsCookiesPriority string

const (
	PdfNewParamsCookiesPriorityLow    PdfNewParamsCookiesPriority = "Low"
	PdfNewParamsCookiesPriorityMedium PdfNewParamsCookiesPriority = "Medium"
	PdfNewParamsCookiesPriorityHigh   PdfNewParamsCookiesPriority = "High"
)

func (r PdfNewParamsCookiesPriority) IsKnown() bool {
	switch r {
	case PdfNewParamsCookiesPriorityLow, PdfNewParamsCookiesPriorityMedium, PdfNewParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type PdfNewParamsCookiesSameSite string

const (
	PdfNewParamsCookiesSameSiteStrict PdfNewParamsCookiesSameSite = "Strict"
	PdfNewParamsCookiesSameSiteLax    PdfNewParamsCookiesSameSite = "Lax"
	PdfNewParamsCookiesSameSiteNone   PdfNewParamsCookiesSameSite = "None"
)

func (r PdfNewParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case PdfNewParamsCookiesSameSiteStrict, PdfNewParamsCookiesSameSiteLax, PdfNewParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type PdfNewParamsCookiesSourceScheme string

const (
	PdfNewParamsCookiesSourceSchemeUnset     PdfNewParamsCookiesSourceScheme = "Unset"
	PdfNewParamsCookiesSourceSchemeNonSecure PdfNewParamsCookiesSourceScheme = "NonSecure"
	PdfNewParamsCookiesSourceSchemeSecure    PdfNewParamsCookiesSourceScheme = "Secure"
)

func (r PdfNewParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case PdfNewParamsCookiesSourceSchemeUnset, PdfNewParamsCookiesSourceSchemeNonSecure, PdfNewParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type PdfNewParamsGotoOptions struct {
	Referer        param.Field[string]                                `json:"referer"`
	ReferrerPolicy param.Field[string]                                `json:"referrerPolicy"`
	Timeout        param.Field[float64]                               `json:"timeout"`
	WaitUntil      param.Field[PdfNewParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r PdfNewParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [browsing_rendering.PdfNewParamsGotoOptionsWaitUntilString],
// [browsing_rendering.PdfNewParamsGotoOptionsWaitUntilArray].
type PdfNewParamsGotoOptionsWaitUntilUnion interface {
	implementsPdfNewParamsGotoOptionsWaitUntilUnion()
}

type PdfNewParamsGotoOptionsWaitUntilString string

const (
	PdfNewParamsGotoOptionsWaitUntilStringLoad             PdfNewParamsGotoOptionsWaitUntilString = "load"
	PdfNewParamsGotoOptionsWaitUntilStringDomcontentloaded PdfNewParamsGotoOptionsWaitUntilString = "domcontentloaded"
	PdfNewParamsGotoOptionsWaitUntilStringNetworkidle0     PdfNewParamsGotoOptionsWaitUntilString = "networkidle0"
	PdfNewParamsGotoOptionsWaitUntilStringNetworkidle2     PdfNewParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r PdfNewParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case PdfNewParamsGotoOptionsWaitUntilStringLoad, PdfNewParamsGotoOptionsWaitUntilStringDomcontentloaded, PdfNewParamsGotoOptionsWaitUntilStringNetworkidle0, PdfNewParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r PdfNewParamsGotoOptionsWaitUntilString) implementsPdfNewParamsGotoOptionsWaitUntilUnion() {}

type PdfNewParamsGotoOptionsWaitUntilArray []PdfNewParamsGotoOptionsWaitUntilArrayItem

func (r PdfNewParamsGotoOptionsWaitUntilArray) implementsPdfNewParamsGotoOptionsWaitUntilUnion() {}

type PdfNewParamsGotoOptionsWaitUntilArrayItem string

const (
	PdfNewParamsGotoOptionsWaitUntilArrayItemLoad             PdfNewParamsGotoOptionsWaitUntilArrayItem = "load"
	PdfNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded PdfNewParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	PdfNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0     PdfNewParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	PdfNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2     PdfNewParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r PdfNewParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case PdfNewParamsGotoOptionsWaitUntilArrayItemLoad, PdfNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, PdfNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0, PdfNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type PdfNewParamsRejectResourceType string

const (
	PdfNewParamsRejectResourceTypeDocument           PdfNewParamsRejectResourceType = "document"
	PdfNewParamsRejectResourceTypeStylesheet         PdfNewParamsRejectResourceType = "stylesheet"
	PdfNewParamsRejectResourceTypeImage              PdfNewParamsRejectResourceType = "image"
	PdfNewParamsRejectResourceTypeMedia              PdfNewParamsRejectResourceType = "media"
	PdfNewParamsRejectResourceTypeFont               PdfNewParamsRejectResourceType = "font"
	PdfNewParamsRejectResourceTypeScript             PdfNewParamsRejectResourceType = "script"
	PdfNewParamsRejectResourceTypeTexttrack          PdfNewParamsRejectResourceType = "texttrack"
	PdfNewParamsRejectResourceTypeXhr                PdfNewParamsRejectResourceType = "xhr"
	PdfNewParamsRejectResourceTypeFetch              PdfNewParamsRejectResourceType = "fetch"
	PdfNewParamsRejectResourceTypePrefetch           PdfNewParamsRejectResourceType = "prefetch"
	PdfNewParamsRejectResourceTypeEventsource        PdfNewParamsRejectResourceType = "eventsource"
	PdfNewParamsRejectResourceTypeWebsocket          PdfNewParamsRejectResourceType = "websocket"
	PdfNewParamsRejectResourceTypeManifest           PdfNewParamsRejectResourceType = "manifest"
	PdfNewParamsRejectResourceTypeSignedexchange     PdfNewParamsRejectResourceType = "signedexchange"
	PdfNewParamsRejectResourceTypePing               PdfNewParamsRejectResourceType = "ping"
	PdfNewParamsRejectResourceTypeCspviolationreport PdfNewParamsRejectResourceType = "cspviolationreport"
	PdfNewParamsRejectResourceTypePreflight          PdfNewParamsRejectResourceType = "preflight"
	PdfNewParamsRejectResourceTypeOther              PdfNewParamsRejectResourceType = "other"
)

func (r PdfNewParamsRejectResourceType) IsKnown() bool {
	switch r {
	case PdfNewParamsRejectResourceTypeDocument, PdfNewParamsRejectResourceTypeStylesheet, PdfNewParamsRejectResourceTypeImage, PdfNewParamsRejectResourceTypeMedia, PdfNewParamsRejectResourceTypeFont, PdfNewParamsRejectResourceTypeScript, PdfNewParamsRejectResourceTypeTexttrack, PdfNewParamsRejectResourceTypeXhr, PdfNewParamsRejectResourceTypeFetch, PdfNewParamsRejectResourceTypePrefetch, PdfNewParamsRejectResourceTypeEventsource, PdfNewParamsRejectResourceTypeWebsocket, PdfNewParamsRejectResourceTypeManifest, PdfNewParamsRejectResourceTypeSignedexchange, PdfNewParamsRejectResourceTypePing, PdfNewParamsRejectResourceTypeCspviolationreport, PdfNewParamsRejectResourceTypePreflight, PdfNewParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type PdfNewParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r PdfNewParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type PdfNewParamsWaitForSelector struct {
	Selector param.Field[string]                             `json:"selector,required"`
	Hidden   param.Field[PdfNewParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                            `json:"timeout"`
	Visible  param.Field[PdfNewParamsWaitForSelectorVisible] `json:"visible"`
}

func (r PdfNewParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PdfNewParamsWaitForSelectorHidden bool

const (
	PdfNewParamsWaitForSelectorHiddenTrue PdfNewParamsWaitForSelectorHidden = true
)

func (r PdfNewParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case PdfNewParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type PdfNewParamsWaitForSelectorVisible bool

const (
	PdfNewParamsWaitForSelectorVisibleTrue PdfNewParamsWaitForSelectorVisible = true
)

func (r PdfNewParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case PdfNewParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}
