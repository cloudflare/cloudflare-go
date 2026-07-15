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

// AccessibilityTreeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessibilityTreeService] method instead.
type AccessibilityTreeService struct {
	Options []option.RequestOption
}

// NewAccessibilityTreeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessibilityTreeService(opts ...option.RequestOption) (r *AccessibilityTreeService) {
	r = &AccessibilityTreeService{}
	r.Options = opts
	return
}

// Returns the page's accessibility tree. Use `interestingOnly` to only return
// semantically meaningful nodes; use `root` to scope the tree to a
// CSS-selector-anchored subtree. Control page loading with `gotoOptions` and
// `waitFor*` options.
func (r *AccessibilityTreeService) New(ctx context.Context, params AccessibilityTreeNewParams, opts ...option.RequestOption) (res *AccessibilityTreeNewResponse, err error) {
	var env AccessibilityTreeNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/accessibilityTree", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AccessibilityTreeNewResponse struct {
	// Accessibility tree node
	AccessibilityTree AccessibilityTreeNewResponseAccessibilityTree `json:"accessibilityTree" api:"required,nullable"`
	JSON              accessibilityTreeNewResponseJSON              `json:"-"`
}

// accessibilityTreeNewResponseJSON contains the JSON metadata for the struct
// [AccessibilityTreeNewResponse]
type accessibilityTreeNewResponseJSON struct {
	AccessibilityTree apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessibilityTreeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessibilityTreeNewResponseJSON) RawJSON() string {
	return r.raw
}

// Accessibility tree node
type AccessibilityTreeNewResponseAccessibilityTree struct {
	Role            string                                                    `json:"role" api:"required"`
	Autocomplete    string                                                    `json:"autocomplete"`
	Checked         AccessibilityTreeNewResponseAccessibilityTreeCheckedUnion `json:"checked"`
	Children        []interface{}                                             `json:"children"`
	Description     string                                                    `json:"description"`
	Disabled        bool                                                      `json:"disabled"`
	Expanded        bool                                                      `json:"expanded"`
	Focused         bool                                                      `json:"focused"`
	Haspopup        string                                                    `json:"haspopup"`
	Invalid         string                                                    `json:"invalid"`
	Keyshortcuts    string                                                    `json:"keyshortcuts"`
	Level           float64                                                   `json:"level"`
	Modal           bool                                                      `json:"modal"`
	Multiline       bool                                                      `json:"multiline"`
	Multiselectable bool                                                      `json:"multiselectable"`
	Name            string                                                    `json:"name"`
	Orientation     string                                                    `json:"orientation"`
	Pressed         AccessibilityTreeNewResponseAccessibilityTreePressedUnion `json:"pressed"`
	Readonly        bool                                                      `json:"readonly"`
	Required        bool                                                      `json:"required"`
	Roledescription string                                                    `json:"roledescription"`
	Selected        bool                                                      `json:"selected"`
	Value           AccessibilityTreeNewResponseAccessibilityTreeValueUnion   `json:"value"`
	Valuemax        float64                                                   `json:"valuemax"`
	Valuemin        float64                                                   `json:"valuemin"`
	Valuetext       string                                                    `json:"valuetext"`
	JSON            accessibilityTreeNewResponseAccessibilityTreeJSON         `json:"-"`
}

// accessibilityTreeNewResponseAccessibilityTreeJSON contains the JSON metadata for
// the struct [AccessibilityTreeNewResponseAccessibilityTree]
type accessibilityTreeNewResponseAccessibilityTreeJSON struct {
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

func (r *AccessibilityTreeNewResponseAccessibilityTree) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessibilityTreeNewResponseAccessibilityTreeJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool] or
// [AccessibilityTreeNewResponseAccessibilityTreeCheckedString].
type AccessibilityTreeNewResponseAccessibilityTreeCheckedUnion interface {
	ImplementsAccessibilityTreeNewResponseAccessibilityTreeCheckedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessibilityTreeNewResponseAccessibilityTreeCheckedUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AccessibilityTreeNewResponseAccessibilityTreeCheckedString("")),
		},
	)
}

type AccessibilityTreeNewResponseAccessibilityTreeCheckedString string

const (
	AccessibilityTreeNewResponseAccessibilityTreeCheckedStringMixed AccessibilityTreeNewResponseAccessibilityTreeCheckedString = "mixed"
)

func (r AccessibilityTreeNewResponseAccessibilityTreeCheckedString) IsKnown() bool {
	switch r {
	case AccessibilityTreeNewResponseAccessibilityTreeCheckedStringMixed:
		return true
	}
	return false
}

func (r AccessibilityTreeNewResponseAccessibilityTreeCheckedString) ImplementsAccessibilityTreeNewResponseAccessibilityTreeCheckedUnion() {
}

// Union satisfied by [shared.UnionBool] or
// [AccessibilityTreeNewResponseAccessibilityTreePressedString].
type AccessibilityTreeNewResponseAccessibilityTreePressedUnion interface {
	ImplementsAccessibilityTreeNewResponseAccessibilityTreePressedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessibilityTreeNewResponseAccessibilityTreePressedUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AccessibilityTreeNewResponseAccessibilityTreePressedString("")),
		},
	)
}

type AccessibilityTreeNewResponseAccessibilityTreePressedString string

const (
	AccessibilityTreeNewResponseAccessibilityTreePressedStringMixed AccessibilityTreeNewResponseAccessibilityTreePressedString = "mixed"
)

func (r AccessibilityTreeNewResponseAccessibilityTreePressedString) IsKnown() bool {
	switch r {
	case AccessibilityTreeNewResponseAccessibilityTreePressedStringMixed:
		return true
	}
	return false
}

func (r AccessibilityTreeNewResponseAccessibilityTreePressedString) ImplementsAccessibilityTreeNewResponseAccessibilityTreePressedUnion() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type AccessibilityTreeNewResponseAccessibilityTreeValueUnion interface {
	ImplementsAccessibilityTreeNewResponseAccessibilityTreeValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessibilityTreeNewResponseAccessibilityTreeValueUnion)(nil)).Elem(),
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

type AccessibilityTreeNewParams struct {
	// Account ID.
	AccountID param.Field[string]                 `path:"account_id" api:"required"`
	Body      AccessibilityTreeNewParamsBodyUnion `json:"body" api:"required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
}

func (r AccessibilityTreeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccessibilityTreeNewParams]'s query parameters as
// `url.Values`.
func (r AccessibilityTreeNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessibilityTreeNewParamsBody struct {
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
	InterestingOnly      param.Field[bool]        `json:"interestingOnly"`
	RejectRequestPattern param.Field[interface{}] `json:"rejectRequestPattern"`
	RejectResourceTypes  param.Field[interface{}] `json:"rejectResourceTypes"`
	Root                 param.Field[string]      `json:"root"`
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

func (r AccessibilityTreeNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessibilityTreeNewParamsBody) implementsAccessibilityTreeNewParamsBodyUnion() {}

// Satisfied by [browser_rendering.AccessibilityTreeNewParamsBodyObject],
// [browser_rendering.AccessibilityTreeNewParamsBodyObject],
// [AccessibilityTreeNewParamsBody].
type AccessibilityTreeNewParamsBodyUnion interface {
	implementsAccessibilityTreeNewParamsBodyUnion()
}

type AccessibilityTreeNewParamsBodyObject struct {
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html" api:"required"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]AccessibilityTreeNewParamsBodyObjectAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]AccessibilityTreeNewParamsBodyObjectAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]AccessibilityTreeNewParamsBodyObjectAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[AccessibilityTreeNewParamsBodyObjectAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]AccessibilityTreeNewParamsBodyObjectCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                       `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions     param.Field[AccessibilityTreeNewParamsBodyObjectGotoOptions] `json:"gotoOptions"`
	InterestingOnly param.Field[bool]                                            `json:"interestingOnly"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]AccessibilityTreeNewParamsBodyObjectRejectResourceType] `json:"rejectResourceTypes"`
	Root                 param.Field[string]                                                   `json:"root"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                        `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                                     `json:"setJavaScriptEnabled"`
	UserAgent            param.Field[string]                                                   `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[AccessibilityTreeNewParamsBodyObjectViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[AccessibilityTreeNewParamsBodyObjectWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r AccessibilityTreeNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessibilityTreeNewParamsBodyObject) implementsAccessibilityTreeNewParamsBodyUnion() {}

type AccessibilityTreeNewParamsBodyObjectAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url" format:"uri"`
}

func (r AccessibilityTreeNewParamsBodyObjectAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessibilityTreeNewParamsBodyObjectAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url" format:"uri"`
}

func (r AccessibilityTreeNewParamsBodyObjectAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessibilityTreeNewParamsBodyObjectAllowResourceType string

const (
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeDocument           AccessibilityTreeNewParamsBodyObjectAllowResourceType = "document"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeStylesheet         AccessibilityTreeNewParamsBodyObjectAllowResourceType = "stylesheet"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeImage              AccessibilityTreeNewParamsBodyObjectAllowResourceType = "image"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeMedia              AccessibilityTreeNewParamsBodyObjectAllowResourceType = "media"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeFont               AccessibilityTreeNewParamsBodyObjectAllowResourceType = "font"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeScript             AccessibilityTreeNewParamsBodyObjectAllowResourceType = "script"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeTexttrack          AccessibilityTreeNewParamsBodyObjectAllowResourceType = "texttrack"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeXHR                AccessibilityTreeNewParamsBodyObjectAllowResourceType = "xhr"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeFetch              AccessibilityTreeNewParamsBodyObjectAllowResourceType = "fetch"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypePrefetch           AccessibilityTreeNewParamsBodyObjectAllowResourceType = "prefetch"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeEventsource        AccessibilityTreeNewParamsBodyObjectAllowResourceType = "eventsource"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeWebsocket          AccessibilityTreeNewParamsBodyObjectAllowResourceType = "websocket"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeManifest           AccessibilityTreeNewParamsBodyObjectAllowResourceType = "manifest"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeSignedexchange     AccessibilityTreeNewParamsBodyObjectAllowResourceType = "signedexchange"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypePing               AccessibilityTreeNewParamsBodyObjectAllowResourceType = "ping"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeCspviolationreport AccessibilityTreeNewParamsBodyObjectAllowResourceType = "cspviolationreport"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypePreflight          AccessibilityTreeNewParamsBodyObjectAllowResourceType = "preflight"
	AccessibilityTreeNewParamsBodyObjectAllowResourceTypeOther              AccessibilityTreeNewParamsBodyObjectAllowResourceType = "other"
)

func (r AccessibilityTreeNewParamsBodyObjectAllowResourceType) IsKnown() bool {
	switch r {
	case AccessibilityTreeNewParamsBodyObjectAllowResourceTypeDocument, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeStylesheet, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeImage, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeMedia, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeFont, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeScript, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeTexttrack, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeXHR, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeFetch, AccessibilityTreeNewParamsBodyObjectAllowResourceTypePrefetch, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeEventsource, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeWebsocket, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeManifest, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeSignedexchange, AccessibilityTreeNewParamsBodyObjectAllowResourceTypePing, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeCspviolationreport, AccessibilityTreeNewParamsBodyObjectAllowResourceTypePreflight, AccessibilityTreeNewParamsBodyObjectAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type AccessibilityTreeNewParamsBodyObjectAuthenticate struct {
	Password param.Field[string] `json:"password" api:"required"`
	Username param.Field[string] `json:"username" api:"required"`
}

func (r AccessibilityTreeNewParamsBodyObjectAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessibilityTreeNewParamsBodyObjectCookie struct {
	// Cookie name.
	Name         param.Field[string]                                                  `json:"name" api:"required"`
	Value        param.Field[string]                                                  `json:"value" api:"required"`
	Domain       param.Field[string]                                                  `json:"domain"`
	Expires      param.Field[float64]                                                 `json:"expires"`
	HTTPOnly     param.Field[bool]                                                    `json:"httpOnly"`
	PartitionKey param.Field[string]                                                  `json:"partitionKey"`
	Path         param.Field[string]                                                  `json:"path"`
	Priority     param.Field[AccessibilityTreeNewParamsBodyObjectCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                                    `json:"sameParty"`
	SameSite     param.Field[AccessibilityTreeNewParamsBodyObjectCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                                    `json:"secure"`
	SourcePort   param.Field[float64]                                                 `json:"sourcePort"`
	SourceScheme param.Field[AccessibilityTreeNewParamsBodyObjectCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                                  `json:"url"`
}

func (r AccessibilityTreeNewParamsBodyObjectCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessibilityTreeNewParamsBodyObjectCookiesPriority string

const (
	AccessibilityTreeNewParamsBodyObjectCookiesPriorityLow    AccessibilityTreeNewParamsBodyObjectCookiesPriority = "Low"
	AccessibilityTreeNewParamsBodyObjectCookiesPriorityMedium AccessibilityTreeNewParamsBodyObjectCookiesPriority = "Medium"
	AccessibilityTreeNewParamsBodyObjectCookiesPriorityHigh   AccessibilityTreeNewParamsBodyObjectCookiesPriority = "High"
)

func (r AccessibilityTreeNewParamsBodyObjectCookiesPriority) IsKnown() bool {
	switch r {
	case AccessibilityTreeNewParamsBodyObjectCookiesPriorityLow, AccessibilityTreeNewParamsBodyObjectCookiesPriorityMedium, AccessibilityTreeNewParamsBodyObjectCookiesPriorityHigh:
		return true
	}
	return false
}

type AccessibilityTreeNewParamsBodyObjectCookiesSameSite string

const (
	AccessibilityTreeNewParamsBodyObjectCookiesSameSiteStrict AccessibilityTreeNewParamsBodyObjectCookiesSameSite = "Strict"
	AccessibilityTreeNewParamsBodyObjectCookiesSameSiteLax    AccessibilityTreeNewParamsBodyObjectCookiesSameSite = "Lax"
	AccessibilityTreeNewParamsBodyObjectCookiesSameSiteNone   AccessibilityTreeNewParamsBodyObjectCookiesSameSite = "None"
)

func (r AccessibilityTreeNewParamsBodyObjectCookiesSameSite) IsKnown() bool {
	switch r {
	case AccessibilityTreeNewParamsBodyObjectCookiesSameSiteStrict, AccessibilityTreeNewParamsBodyObjectCookiesSameSiteLax, AccessibilityTreeNewParamsBodyObjectCookiesSameSiteNone:
		return true
	}
	return false
}

type AccessibilityTreeNewParamsBodyObjectCookiesSourceScheme string

const (
	AccessibilityTreeNewParamsBodyObjectCookiesSourceSchemeUnset     AccessibilityTreeNewParamsBodyObjectCookiesSourceScheme = "Unset"
	AccessibilityTreeNewParamsBodyObjectCookiesSourceSchemeNonSecure AccessibilityTreeNewParamsBodyObjectCookiesSourceScheme = "NonSecure"
	AccessibilityTreeNewParamsBodyObjectCookiesSourceSchemeSecure    AccessibilityTreeNewParamsBodyObjectCookiesSourceScheme = "Secure"
)

func (r AccessibilityTreeNewParamsBodyObjectCookiesSourceScheme) IsKnown() bool {
	switch r {
	case AccessibilityTreeNewParamsBodyObjectCookiesSourceSchemeUnset, AccessibilityTreeNewParamsBodyObjectCookiesSourceSchemeNonSecure, AccessibilityTreeNewParamsBodyObjectCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type AccessibilityTreeNewParamsBodyObjectGotoOptions struct {
	Referer        param.Field[string]                                                        `json:"referer"`
	ReferrerPolicy param.Field[string]                                                        `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                       `json:"timeout"`
	WaitUntil      param.Field[AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r AccessibilityTreeNewParamsBodyObjectGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilString],
// [browser_rendering.AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArray].
type AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilUnion interface {
	implementsAccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilUnion()
}

type AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilString string

const (
	AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilStringLoad             AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilString = "load"
	AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilString = "domcontentloaded"
	AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0     AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle0"
	AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2     AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle2"
)

func (r AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilStringLoad, AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded, AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0, AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilString) implementsAccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArray []AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem

func (r AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArray) implementsAccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem string

const (
	AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad             AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "load"
	AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0     AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle0"
	AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2     AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad, AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded, AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0, AccessibilityTreeNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type AccessibilityTreeNewParamsBodyObjectRejectResourceType string

const (
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeDocument           AccessibilityTreeNewParamsBodyObjectRejectResourceType = "document"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeStylesheet         AccessibilityTreeNewParamsBodyObjectRejectResourceType = "stylesheet"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeImage              AccessibilityTreeNewParamsBodyObjectRejectResourceType = "image"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeMedia              AccessibilityTreeNewParamsBodyObjectRejectResourceType = "media"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeFont               AccessibilityTreeNewParamsBodyObjectRejectResourceType = "font"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeScript             AccessibilityTreeNewParamsBodyObjectRejectResourceType = "script"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeTexttrack          AccessibilityTreeNewParamsBodyObjectRejectResourceType = "texttrack"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeXHR                AccessibilityTreeNewParamsBodyObjectRejectResourceType = "xhr"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeFetch              AccessibilityTreeNewParamsBodyObjectRejectResourceType = "fetch"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypePrefetch           AccessibilityTreeNewParamsBodyObjectRejectResourceType = "prefetch"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeEventsource        AccessibilityTreeNewParamsBodyObjectRejectResourceType = "eventsource"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeWebsocket          AccessibilityTreeNewParamsBodyObjectRejectResourceType = "websocket"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeManifest           AccessibilityTreeNewParamsBodyObjectRejectResourceType = "manifest"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeSignedexchange     AccessibilityTreeNewParamsBodyObjectRejectResourceType = "signedexchange"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypePing               AccessibilityTreeNewParamsBodyObjectRejectResourceType = "ping"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeCspviolationreport AccessibilityTreeNewParamsBodyObjectRejectResourceType = "cspviolationreport"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypePreflight          AccessibilityTreeNewParamsBodyObjectRejectResourceType = "preflight"
	AccessibilityTreeNewParamsBodyObjectRejectResourceTypeOther              AccessibilityTreeNewParamsBodyObjectRejectResourceType = "other"
)

func (r AccessibilityTreeNewParamsBodyObjectRejectResourceType) IsKnown() bool {
	switch r {
	case AccessibilityTreeNewParamsBodyObjectRejectResourceTypeDocument, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeStylesheet, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeImage, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeMedia, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeFont, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeScript, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeTexttrack, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeXHR, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeFetch, AccessibilityTreeNewParamsBodyObjectRejectResourceTypePrefetch, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeEventsource, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeWebsocket, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeManifest, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeSignedexchange, AccessibilityTreeNewParamsBodyObjectRejectResourceTypePing, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeCspviolationreport, AccessibilityTreeNewParamsBodyObjectRejectResourceTypePreflight, AccessibilityTreeNewParamsBodyObjectRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type AccessibilityTreeNewParamsBodyObjectViewport struct {
	Height            param.Field[float64] `json:"height" api:"required"`
	Width             param.Field[float64] `json:"width" api:"required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r AccessibilityTreeNewParamsBodyObjectViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type AccessibilityTreeNewParamsBodyObjectWaitForSelector struct {
	Selector param.Field[string]                                                     `json:"selector" api:"required"`
	Hidden   param.Field[AccessibilityTreeNewParamsBodyObjectWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                                    `json:"timeout"`
	Visible  param.Field[AccessibilityTreeNewParamsBodyObjectWaitForSelectorVisible] `json:"visible"`
}

func (r AccessibilityTreeNewParamsBodyObjectWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessibilityTreeNewParamsBodyObjectWaitForSelectorHidden bool

const (
	AccessibilityTreeNewParamsBodyObjectWaitForSelectorHiddenTrue AccessibilityTreeNewParamsBodyObjectWaitForSelectorHidden = true
)

func (r AccessibilityTreeNewParamsBodyObjectWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case AccessibilityTreeNewParamsBodyObjectWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type AccessibilityTreeNewParamsBodyObjectWaitForSelectorVisible bool

const (
	AccessibilityTreeNewParamsBodyObjectWaitForSelectorVisibleTrue AccessibilityTreeNewParamsBodyObjectWaitForSelectorVisible = true
)

func (r AccessibilityTreeNewParamsBodyObjectWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case AccessibilityTreeNewParamsBodyObjectWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type AccessibilityTreeNewResponseEnvelope struct {
	Meta   AccessibilityTreeNewResponseEnvelopeMeta `json:"meta" api:"required"`
	Result AccessibilityTreeNewResponse             `json:"result" api:"required"`
	// Response status.
	Success bool                                         `json:"success" api:"required"`
	Errors  []AccessibilityTreeNewResponseEnvelopeErrors `json:"errors"`
	JSON    accessibilityTreeNewResponseEnvelopeJSON     `json:"-"`
}

// accessibilityTreeNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessibilityTreeNewResponseEnvelope]
type accessibilityTreeNewResponseEnvelopeJSON struct {
	Meta        apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessibilityTreeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessibilityTreeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessibilityTreeNewResponseEnvelopeMeta struct {
	Status float64                                      `json:"status"`
	Title  string                                       `json:"title"`
	JSON   accessibilityTreeNewResponseEnvelopeMetaJSON `json:"-"`
}

// accessibilityTreeNewResponseEnvelopeMetaJSON contains the JSON metadata for the
// struct [AccessibilityTreeNewResponseEnvelopeMeta]
type accessibilityTreeNewResponseEnvelopeMetaJSON struct {
	Status      apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessibilityTreeNewResponseEnvelopeMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessibilityTreeNewResponseEnvelopeMetaJSON) RawJSON() string {
	return r.raw
}

type AccessibilityTreeNewResponseEnvelopeErrors struct {
	// Error code.
	Code float64 `json:"code" api:"required"`
	// Error message.
	Message string                                         `json:"message" api:"required"`
	JSON    accessibilityTreeNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessibilityTreeNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessibilityTreeNewResponseEnvelopeErrors]
type accessibilityTreeNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessibilityTreeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessibilityTreeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
