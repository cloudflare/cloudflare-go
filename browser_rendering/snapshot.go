// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
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
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/snapshot", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SnapshotNewResponse struct {
	// Accessibility tree node
	AccessibilityTree SnapshotNewResponseAccessibilityTree `json:"accessibilityTree"`
	// HTML content.
	Content string `json:"content"`
	// Markdown content. Prefixed with YAML frontmatter (e.g. `title`) when the page
	// provides that metadata.
	Markdown string `json:"markdown"`
	// Base64 encoded image.
	Screenshot string                  `json:"screenshot"`
	JSON       snapshotNewResponseJSON `json:"-"`
}

// snapshotNewResponseJSON contains the JSON metadata for the struct
// [SnapshotNewResponse]
type snapshotNewResponseJSON struct {
	AccessibilityTree apijson.Field
	Content           apijson.Field
	Markdown          apijson.Field
	Screenshot        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SnapshotNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snapshotNewResponseJSON) RawJSON() string {
	return r.raw
}

// Accessibility tree node
type SnapshotNewResponseAccessibilityTree struct {
	Role            string                                           `json:"role" api:"required"`
	Autocomplete    string                                           `json:"autocomplete"`
	Checked         SnapshotNewResponseAccessibilityTreeCheckedUnion `json:"checked"`
	Children        []interface{}                                    `json:"children"`
	Description     string                                           `json:"description"`
	Disabled        bool                                             `json:"disabled"`
	Expanded        bool                                             `json:"expanded"`
	Focused         bool                                             `json:"focused"`
	Haspopup        string                                           `json:"haspopup"`
	Invalid         string                                           `json:"invalid"`
	Keyshortcuts    string                                           `json:"keyshortcuts"`
	Level           float64                                          `json:"level"`
	Modal           bool                                             `json:"modal"`
	Multiline       bool                                             `json:"multiline"`
	Multiselectable bool                                             `json:"multiselectable"`
	Name            string                                           `json:"name"`
	Orientation     string                                           `json:"orientation"`
	Pressed         SnapshotNewResponseAccessibilityTreePressedUnion `json:"pressed"`
	Readonly        bool                                             `json:"readonly"`
	Required        bool                                             `json:"required"`
	Roledescription string                                           `json:"roledescription"`
	Selected        bool                                             `json:"selected"`
	Value           SnapshotNewResponseAccessibilityTreeValueUnion   `json:"value"`
	Valuemax        float64                                          `json:"valuemax"`
	Valuemin        float64                                          `json:"valuemin"`
	Valuetext       string                                           `json:"valuetext"`
	JSON            snapshotNewResponseAccessibilityTreeJSON         `json:"-"`
}

// snapshotNewResponseAccessibilityTreeJSON contains the JSON metadata for the
// struct [SnapshotNewResponseAccessibilityTree]
type snapshotNewResponseAccessibilityTreeJSON struct {
	Role            apijson.Field
	Autocomplete    apijson.Field
	Checked         apijson.Field
	Children        apijson.Field
	Description     apijson.Field
	Disabled        apijson.Field
	Expanded        apijson.Field
	Focused         apijson.Field
	Haspopup        apijson.Field
	Invalid         apijson.Field
	Keyshortcuts    apijson.Field
	Level           apijson.Field
	Modal           apijson.Field
	Multiline       apijson.Field
	Multiselectable apijson.Field
	Name            apijson.Field
	Orientation     apijson.Field
	Pressed         apijson.Field
	Readonly        apijson.Field
	Required        apijson.Field
	Roledescription apijson.Field
	Selected        apijson.Field
	Value           apijson.Field
	Valuemax        apijson.Field
	Valuemin        apijson.Field
	Valuetext       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SnapshotNewResponseAccessibilityTree) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snapshotNewResponseAccessibilityTreeJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool] or
// [SnapshotNewResponseAccessibilityTreeCheckedString].
type SnapshotNewResponseAccessibilityTreeCheckedUnion interface {
	ImplementsSnapshotNewResponseAccessibilityTreeCheckedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SnapshotNewResponseAccessibilityTreeCheckedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(SnapshotNewResponseAccessibilityTreeCheckedString("")),
		},
	)
}

type SnapshotNewResponseAccessibilityTreeCheckedString string

const (
	SnapshotNewResponseAccessibilityTreeCheckedStringMixed SnapshotNewResponseAccessibilityTreeCheckedString = "mixed"
)

func (r SnapshotNewResponseAccessibilityTreeCheckedString) IsKnown() bool {
	switch r {
	case SnapshotNewResponseAccessibilityTreeCheckedStringMixed:
		return true
	}
	return false
}

func (r SnapshotNewResponseAccessibilityTreeCheckedString) ImplementsSnapshotNewResponseAccessibilityTreeCheckedUnion() {
}

// Union satisfied by [shared.UnionBool] or
// [SnapshotNewResponseAccessibilityTreePressedString].
type SnapshotNewResponseAccessibilityTreePressedUnion interface {
	ImplementsSnapshotNewResponseAccessibilityTreePressedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SnapshotNewResponseAccessibilityTreePressedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(SnapshotNewResponseAccessibilityTreePressedString("")),
		},
	)
}

type SnapshotNewResponseAccessibilityTreePressedString string

const (
	SnapshotNewResponseAccessibilityTreePressedStringMixed SnapshotNewResponseAccessibilityTreePressedString = "mixed"
)

func (r SnapshotNewResponseAccessibilityTreePressedString) IsKnown() bool {
	switch r {
	case SnapshotNewResponseAccessibilityTreePressedStringMixed:
		return true
	}
	return false
}

func (r SnapshotNewResponseAccessibilityTreePressedString) ImplementsSnapshotNewResponseAccessibilityTreePressedUnion() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type SnapshotNewResponseAccessibilityTreeValueUnion interface {
	ImplementsSnapshotNewResponseAccessibilityTreeValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SnapshotNewResponseAccessibilityTreeValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type SnapshotNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]SnapshotNewParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]SnapshotNewParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	// Reject rules are applied first.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'. Reject rules are applied first.
	AllowResourceTypes param.Field[[]SnapshotNewParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[SnapshotNewParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]SnapshotNewParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                    `json:"emulateMediaType"`
	Formats          param.Field[[]SnapshotNewParamsFormat] `json:"formats"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[SnapshotNewParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]SnapshotNewParamsRejectResourceType] `json:"rejectResourceTypes"`
	ScreenshotOptions    param.Field[SnapshotNewParamsScreenshotOptions]    `json:"screenshotOptions"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                     `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                  `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[SnapshotNewParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[SnapshotNewParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r SnapshotNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SnapshotNewParams]'s query parameters as `url.Values`.
func (r SnapshotNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SnapshotNewParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url" format:"uri"`
}

func (r SnapshotNewParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url" format:"uri"`
}

func (r SnapshotNewParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsAllowResourceType string

const (
	SnapshotNewParamsAllowResourceTypeDocument           SnapshotNewParamsAllowResourceType = "document"
	SnapshotNewParamsAllowResourceTypeStylesheet         SnapshotNewParamsAllowResourceType = "stylesheet"
	SnapshotNewParamsAllowResourceTypeImage              SnapshotNewParamsAllowResourceType = "image"
	SnapshotNewParamsAllowResourceTypeMedia              SnapshotNewParamsAllowResourceType = "media"
	SnapshotNewParamsAllowResourceTypeFont               SnapshotNewParamsAllowResourceType = "font"
	SnapshotNewParamsAllowResourceTypeScript             SnapshotNewParamsAllowResourceType = "script"
	SnapshotNewParamsAllowResourceTypeTexttrack          SnapshotNewParamsAllowResourceType = "texttrack"
	SnapshotNewParamsAllowResourceTypeXHR                SnapshotNewParamsAllowResourceType = "xhr"
	SnapshotNewParamsAllowResourceTypeFetch              SnapshotNewParamsAllowResourceType = "fetch"
	SnapshotNewParamsAllowResourceTypePrefetch           SnapshotNewParamsAllowResourceType = "prefetch"
	SnapshotNewParamsAllowResourceTypeEventsource        SnapshotNewParamsAllowResourceType = "eventsource"
	SnapshotNewParamsAllowResourceTypeWebsocket          SnapshotNewParamsAllowResourceType = "websocket"
	SnapshotNewParamsAllowResourceTypeManifest           SnapshotNewParamsAllowResourceType = "manifest"
	SnapshotNewParamsAllowResourceTypeSignedexchange     SnapshotNewParamsAllowResourceType = "signedexchange"
	SnapshotNewParamsAllowResourceTypePing               SnapshotNewParamsAllowResourceType = "ping"
	SnapshotNewParamsAllowResourceTypeCspviolationreport SnapshotNewParamsAllowResourceType = "cspviolationreport"
	SnapshotNewParamsAllowResourceTypePreflight          SnapshotNewParamsAllowResourceType = "preflight"
	SnapshotNewParamsAllowResourceTypeOther              SnapshotNewParamsAllowResourceType = "other"
)

func (r SnapshotNewParamsAllowResourceType) IsKnown() bool {
	switch r {
	case SnapshotNewParamsAllowResourceTypeDocument, SnapshotNewParamsAllowResourceTypeStylesheet, SnapshotNewParamsAllowResourceTypeImage, SnapshotNewParamsAllowResourceTypeMedia, SnapshotNewParamsAllowResourceTypeFont, SnapshotNewParamsAllowResourceTypeScript, SnapshotNewParamsAllowResourceTypeTexttrack, SnapshotNewParamsAllowResourceTypeXHR, SnapshotNewParamsAllowResourceTypeFetch, SnapshotNewParamsAllowResourceTypePrefetch, SnapshotNewParamsAllowResourceTypeEventsource, SnapshotNewParamsAllowResourceTypeWebsocket, SnapshotNewParamsAllowResourceTypeManifest, SnapshotNewParamsAllowResourceTypeSignedexchange, SnapshotNewParamsAllowResourceTypePing, SnapshotNewParamsAllowResourceTypeCspviolationreport, SnapshotNewParamsAllowResourceTypePreflight, SnapshotNewParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type SnapshotNewParamsAuthenticate struct {
	Password param.Field[string] `json:"password" api:"required"`
	Username param.Field[string] `json:"username" api:"required"`
}

func (r SnapshotNewParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsCookie struct {
	// Cookie name.
	Name         param.Field[string]                               `json:"name" api:"required"`
	Value        param.Field[string]                               `json:"value" api:"required"`
	Domain       param.Field[string]                               `json:"domain"`
	Expires      param.Field[float64]                              `json:"expires"`
	HTTPOnly     param.Field[bool]                                 `json:"httpOnly"`
	PartitionKey param.Field[string]                               `json:"partitionKey"`
	Path         param.Field[string]                               `json:"path"`
	Priority     param.Field[SnapshotNewParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                 `json:"sameParty"`
	SameSite     param.Field[SnapshotNewParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                 `json:"secure"`
	SourcePort   param.Field[float64]                              `json:"sourcePort"`
	SourceScheme param.Field[SnapshotNewParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                               `json:"url"`
}

func (r SnapshotNewParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsCookiesPriority string

const (
	SnapshotNewParamsCookiesPriorityLow    SnapshotNewParamsCookiesPriority = "Low"
	SnapshotNewParamsCookiesPriorityMedium SnapshotNewParamsCookiesPriority = "Medium"
	SnapshotNewParamsCookiesPriorityHigh   SnapshotNewParamsCookiesPriority = "High"
)

func (r SnapshotNewParamsCookiesPriority) IsKnown() bool {
	switch r {
	case SnapshotNewParamsCookiesPriorityLow, SnapshotNewParamsCookiesPriorityMedium, SnapshotNewParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type SnapshotNewParamsCookiesSameSite string

const (
	SnapshotNewParamsCookiesSameSiteStrict SnapshotNewParamsCookiesSameSite = "Strict"
	SnapshotNewParamsCookiesSameSiteLax    SnapshotNewParamsCookiesSameSite = "Lax"
	SnapshotNewParamsCookiesSameSiteNone   SnapshotNewParamsCookiesSameSite = "None"
)

func (r SnapshotNewParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case SnapshotNewParamsCookiesSameSiteStrict, SnapshotNewParamsCookiesSameSiteLax, SnapshotNewParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type SnapshotNewParamsCookiesSourceScheme string

const (
	SnapshotNewParamsCookiesSourceSchemeUnset     SnapshotNewParamsCookiesSourceScheme = "Unset"
	SnapshotNewParamsCookiesSourceSchemeNonSecure SnapshotNewParamsCookiesSourceScheme = "NonSecure"
	SnapshotNewParamsCookiesSourceSchemeSecure    SnapshotNewParamsCookiesSourceScheme = "Secure"
)

func (r SnapshotNewParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case SnapshotNewParamsCookiesSourceSchemeUnset, SnapshotNewParamsCookiesSourceSchemeNonSecure, SnapshotNewParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

type SnapshotNewParamsFormat string

const (
	SnapshotNewParamsFormatContent           SnapshotNewParamsFormat = "content"
	SnapshotNewParamsFormatScreenshot        SnapshotNewParamsFormat = "screenshot"
	SnapshotNewParamsFormatMarkdown          SnapshotNewParamsFormat = "markdown"
	SnapshotNewParamsFormatAccessibilityTree SnapshotNewParamsFormat = "accessibilityTree"
)

func (r SnapshotNewParamsFormat) IsKnown() bool {
	switch r {
	case SnapshotNewParamsFormatContent, SnapshotNewParamsFormatScreenshot, SnapshotNewParamsFormatMarkdown, SnapshotNewParamsFormatAccessibilityTree:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type SnapshotNewParamsGotoOptions struct {
	Referer        param.Field[string]                                     `json:"referer"`
	ReferrerPolicy param.Field[string]                                     `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                    `json:"timeout"`
	WaitUntil      param.Field[SnapshotNewParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r SnapshotNewParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [browser_rendering.SnapshotNewParamsGotoOptionsWaitUntilString],
// [browser_rendering.SnapshotNewParamsGotoOptionsWaitUntilArray].
type SnapshotNewParamsGotoOptionsWaitUntilUnion interface {
	implementsSnapshotNewParamsGotoOptionsWaitUntilUnion()
}

type SnapshotNewParamsGotoOptionsWaitUntilString string

const (
	SnapshotNewParamsGotoOptionsWaitUntilStringLoad             SnapshotNewParamsGotoOptionsWaitUntilString = "load"
	SnapshotNewParamsGotoOptionsWaitUntilStringDomcontentloaded SnapshotNewParamsGotoOptionsWaitUntilString = "domcontentloaded"
	SnapshotNewParamsGotoOptionsWaitUntilStringNetworkidle0     SnapshotNewParamsGotoOptionsWaitUntilString = "networkidle0"
	SnapshotNewParamsGotoOptionsWaitUntilStringNetworkidle2     SnapshotNewParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r SnapshotNewParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case SnapshotNewParamsGotoOptionsWaitUntilStringLoad, SnapshotNewParamsGotoOptionsWaitUntilStringDomcontentloaded, SnapshotNewParamsGotoOptionsWaitUntilStringNetworkidle0, SnapshotNewParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r SnapshotNewParamsGotoOptionsWaitUntilString) implementsSnapshotNewParamsGotoOptionsWaitUntilUnion() {
}

type SnapshotNewParamsGotoOptionsWaitUntilArray []SnapshotNewParamsGotoOptionsWaitUntilArrayItem

func (r SnapshotNewParamsGotoOptionsWaitUntilArray) implementsSnapshotNewParamsGotoOptionsWaitUntilUnion() {
}

type SnapshotNewParamsGotoOptionsWaitUntilArrayItem string

const (
	SnapshotNewParamsGotoOptionsWaitUntilArrayItemLoad             SnapshotNewParamsGotoOptionsWaitUntilArrayItem = "load"
	SnapshotNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded SnapshotNewParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	SnapshotNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0     SnapshotNewParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	SnapshotNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2     SnapshotNewParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r SnapshotNewParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case SnapshotNewParamsGotoOptionsWaitUntilArrayItemLoad, SnapshotNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, SnapshotNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0, SnapshotNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type SnapshotNewParamsRejectResourceType string

const (
	SnapshotNewParamsRejectResourceTypeDocument           SnapshotNewParamsRejectResourceType = "document"
	SnapshotNewParamsRejectResourceTypeStylesheet         SnapshotNewParamsRejectResourceType = "stylesheet"
	SnapshotNewParamsRejectResourceTypeImage              SnapshotNewParamsRejectResourceType = "image"
	SnapshotNewParamsRejectResourceTypeMedia              SnapshotNewParamsRejectResourceType = "media"
	SnapshotNewParamsRejectResourceTypeFont               SnapshotNewParamsRejectResourceType = "font"
	SnapshotNewParamsRejectResourceTypeScript             SnapshotNewParamsRejectResourceType = "script"
	SnapshotNewParamsRejectResourceTypeTexttrack          SnapshotNewParamsRejectResourceType = "texttrack"
	SnapshotNewParamsRejectResourceTypeXHR                SnapshotNewParamsRejectResourceType = "xhr"
	SnapshotNewParamsRejectResourceTypeFetch              SnapshotNewParamsRejectResourceType = "fetch"
	SnapshotNewParamsRejectResourceTypePrefetch           SnapshotNewParamsRejectResourceType = "prefetch"
	SnapshotNewParamsRejectResourceTypeEventsource        SnapshotNewParamsRejectResourceType = "eventsource"
	SnapshotNewParamsRejectResourceTypeWebsocket          SnapshotNewParamsRejectResourceType = "websocket"
	SnapshotNewParamsRejectResourceTypeManifest           SnapshotNewParamsRejectResourceType = "manifest"
	SnapshotNewParamsRejectResourceTypeSignedexchange     SnapshotNewParamsRejectResourceType = "signedexchange"
	SnapshotNewParamsRejectResourceTypePing               SnapshotNewParamsRejectResourceType = "ping"
	SnapshotNewParamsRejectResourceTypeCspviolationreport SnapshotNewParamsRejectResourceType = "cspviolationreport"
	SnapshotNewParamsRejectResourceTypePreflight          SnapshotNewParamsRejectResourceType = "preflight"
	SnapshotNewParamsRejectResourceTypeOther              SnapshotNewParamsRejectResourceType = "other"
)

func (r SnapshotNewParamsRejectResourceType) IsKnown() bool {
	switch r {
	case SnapshotNewParamsRejectResourceTypeDocument, SnapshotNewParamsRejectResourceTypeStylesheet, SnapshotNewParamsRejectResourceTypeImage, SnapshotNewParamsRejectResourceTypeMedia, SnapshotNewParamsRejectResourceTypeFont, SnapshotNewParamsRejectResourceTypeScript, SnapshotNewParamsRejectResourceTypeTexttrack, SnapshotNewParamsRejectResourceTypeXHR, SnapshotNewParamsRejectResourceTypeFetch, SnapshotNewParamsRejectResourceTypePrefetch, SnapshotNewParamsRejectResourceTypeEventsource, SnapshotNewParamsRejectResourceTypeWebsocket, SnapshotNewParamsRejectResourceTypeManifest, SnapshotNewParamsRejectResourceTypeSignedexchange, SnapshotNewParamsRejectResourceTypePing, SnapshotNewParamsRejectResourceTypeCspviolationreport, SnapshotNewParamsRejectResourceTypePreflight, SnapshotNewParamsRejectResourceTypeOther:
		return true
	}
	return false
}

type SnapshotNewParamsScreenshotOptions struct {
	CaptureBeyondViewport param.Field[bool]                                   `json:"captureBeyondViewport"`
	Clip                  param.Field[SnapshotNewParamsScreenshotOptionsClip] `json:"clip"`
	FromSurface           param.Field[bool]                                   `json:"fromSurface"`
	FullPage              param.Field[bool]                                   `json:"fullPage"`
	OmitBackground        param.Field[bool]                                   `json:"omitBackground"`
	OptimizeForSpeed      param.Field[bool]                                   `json:"optimizeForSpeed"`
	Quality               param.Field[float64]                                `json:"quality"`
	Type                  param.Field[SnapshotNewParamsScreenshotOptionsType] `json:"type"`
}

func (r SnapshotNewParamsScreenshotOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsScreenshotOptionsClip struct {
	Height param.Field[float64] `json:"height" api:"required"`
	Width  param.Field[float64] `json:"width" api:"required"`
	X      param.Field[float64] `json:"x" api:"required"`
	Y      param.Field[float64] `json:"y" api:"required"`
	Scale  param.Field[float64] `json:"scale"`
}

func (r SnapshotNewParamsScreenshotOptionsClip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsScreenshotOptionsType string

const (
	SnapshotNewParamsScreenshotOptionsTypePNG  SnapshotNewParamsScreenshotOptionsType = "png"
	SnapshotNewParamsScreenshotOptionsTypeJPEG SnapshotNewParamsScreenshotOptionsType = "jpeg"
	SnapshotNewParamsScreenshotOptionsTypeWebP SnapshotNewParamsScreenshotOptionsType = "webp"
)

func (r SnapshotNewParamsScreenshotOptionsType) IsKnown() bool {
	switch r {
	case SnapshotNewParamsScreenshotOptionsTypePNG, SnapshotNewParamsScreenshotOptionsTypeJPEG, SnapshotNewParamsScreenshotOptionsTypeWebP:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type SnapshotNewParamsViewport struct {
	Height            param.Field[float64] `json:"height" api:"required"`
	Width             param.Field[float64] `json:"width" api:"required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r SnapshotNewParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type SnapshotNewParamsWaitForSelector struct {
	Selector param.Field[string]                                  `json:"selector" api:"required"`
	Hidden   param.Field[SnapshotNewParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                 `json:"timeout"`
	Visible  param.Field[SnapshotNewParamsWaitForSelectorVisible] `json:"visible"`
}

func (r SnapshotNewParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnapshotNewParamsWaitForSelectorHidden bool

const (
	SnapshotNewParamsWaitForSelectorHiddenTrue SnapshotNewParamsWaitForSelectorHidden = true
)

func (r SnapshotNewParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case SnapshotNewParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type SnapshotNewParamsWaitForSelectorVisible bool

const (
	SnapshotNewParamsWaitForSelectorVisibleTrue SnapshotNewParamsWaitForSelectorVisible = true
)

func (r SnapshotNewParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case SnapshotNewParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type SnapshotNewResponseEnvelope struct {
	Meta SnapshotNewResponseEnvelopeMeta `json:"meta" api:"required"`
	// Response status.
	Success bool                                `json:"success" api:"required"`
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
	// URL that served the response, after any redirects the browser followed.
	FinalURL string `json:"finalUrl"`
	// Origin response headers, lowercased. Repeated headers are joined with a newline.
	// Credential and transport-only headers that do not survive rendering are omitted.
	Headers map[string]string `json:"headers"`
	// HTTP redirects followed to reach `finalUrl`, oldest first. Omitted for direct
	// navigation and for client-side redirects such as meta refresh. An empty array
	// means redirects occurred but their intermediate responses could not be read.
	RedirectChain []SnapshotNewResponseEnvelopeMetaRedirectChain `json:"redirectChain"`
	// HTTP status returned by the origin.
	Status float64 `json:"status"`
	// Page title.
	Title string                              `json:"title"`
	JSON  snapshotNewResponseEnvelopeMetaJSON `json:"-"`
}

// snapshotNewResponseEnvelopeMetaJSON contains the JSON metadata for the struct
// [SnapshotNewResponseEnvelopeMeta]
type snapshotNewResponseEnvelopeMetaJSON struct {
	FinalURL      apijson.Field
	Headers       apijson.Field
	RedirectChain apijson.Field
	Status        apijson.Field
	Title         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SnapshotNewResponseEnvelopeMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snapshotNewResponseEnvelopeMetaJSON) RawJSON() string {
	return r.raw
}

type SnapshotNewResponseEnvelopeMetaRedirectChain struct {
	// Redirect response headers, including `location`.
	Headers map[string]string `json:"headers" api:"required"`
	// HTTP status of the redirect.
	Status float64 `json:"status" api:"required"`
	// URL that returned the redirect.
	URL  string                                           `json:"url" api:"required"`
	JSON snapshotNewResponseEnvelopeMetaRedirectChainJSON `json:"-"`
}

// snapshotNewResponseEnvelopeMetaRedirectChainJSON contains the JSON metadata for
// the struct [SnapshotNewResponseEnvelopeMetaRedirectChain]
type snapshotNewResponseEnvelopeMetaRedirectChainJSON struct {
	Headers     apijson.Field
	Status      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnapshotNewResponseEnvelopeMetaRedirectChain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snapshotNewResponseEnvelopeMetaRedirectChainJSON) RawJSON() string {
	return r.raw
}

type SnapshotNewResponseEnvelopeErrors struct {
	// Error code.
	Code float64 `json:"code" api:"required"`
	// Error message.
	Message string                                `json:"message" api:"required"`
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
