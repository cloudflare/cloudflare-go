// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagerules

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/cloudflare/cloudflare-go/v4/zones"
	"github.com/tidwall/gjson"
)

// PageruleService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageruleService] method instead.
type PageruleService struct {
	Options []option.RequestOption
}

// NewPageruleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageruleService(opts ...option.RequestOption) (r *PageruleService) {
	r = &PageruleService{}
	r.Options = opts
	return
}

// Creates a new Page Rule.
func (r *PageruleService) New(ctx context.Context, params PageruleNewParams, opts ...option.RequestOption) (res *PageRule, err error) {
	var env PageruleNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Replaces the configuration of an existing Page Rule. The configuration of the
// updated Page Rule will exactly match the data passed in the API request.
func (r *PageruleService) Update(ctx context.Context, pageruleID string, params PageruleUpdateParams, opts ...option.RequestOption) (res *PageRule, err error) {
	var env PageruleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", params.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Page Rules in a zone.
func (r *PageruleService) List(ctx context.Context, params PageruleListParams, opts ...option.RequestOption) (res *[]PageRule, err error) {
	var env PageruleListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing Page Rule.
func (r *PageruleService) Delete(ctx context.Context, pageruleID string, body PageruleDeleteParams, opts ...option.RequestOption) (res *PageruleDeleteResponse, err error) {
	var env PageruleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", body.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates one or more fields of an existing Page Rule.
func (r *PageruleService) Edit(ctx context.Context, pageruleID string, params PageruleEditParams, opts ...option.RequestOption) (res *PageRule, err error) {
	var env PageruleEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", params.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a Page Rule.
func (r *PageruleService) Get(ctx context.Context, pageruleID string, query PageruleGetParams, opts ...option.RequestOption) (res *PageRule, err error) {
	var env PageruleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", query.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageRule struct {
	// Identifier
	ID string `json:"id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions []PageRuleAction `json:"actions,required"`
	// The timestamp of when the Page Rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The timestamp of when the Page Rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority int64 `json:"priority,required"`
	// The status of the Page Rule.
	Status PageRuleStatus `json:"status,required"`
	// The rule targets to evaluate on each request.
	Targets []Target     `json:"targets,required"`
	JSON    pageRuleJSON `json:"-"`
}

// pageRuleJSON contains the JSON metadata for the struct [PageRule]
type pageRuleJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Priority    apijson.Field
	Status      apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleJSON) RawJSON() string {
	return r.raw
}

type PageRuleAction struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID PageRuleActionsID `json:"id"`
	// This field can have the runtime type of [zones.AutomaticHTTPSRewritesValue],
	// [int64], [zones.BrowserCheckValue], [string],
	// [PageRuleActionsCacheByDeviceTypeValue],
	// [PageRuleActionsCacheDeceptionArmorValue], [PageRuleActionsCacheKeyValue],
	// [zones.CacheLevelValue], [zones.EmailObfuscationValue],
	// [PageRuleActionsExplicitCacheControlValue], [PageRuleActionsForwardingURLValue],
	// [zones.IPGeolocationValue], [zones.MirageValue],
	// [zones.OpportunisticEncryptionValue], [zones.OriginErrorPagePassThruValue],
	// [zones.PolishValue], [PageRuleActionsRespectStrongEtagValue],
	// [zones.ResponseBufferingValue], [zones.RocketLoaderValue],
	// [zones.SecurityLevelValue], [zones.SortQueryStringForCacheValue],
	// [zones.SSLValue], [zones.TrueClientIPHeaderValue], [zones.WAFValue].
	Value interface{}        `json:"value"`
	JSON  pageRuleActionJSON `json:"-"`
	union PageRuleActionsUnion
}

// pageRuleActionJSON contains the JSON metadata for the struct [PageRuleAction]
type pageRuleActionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r pageRuleActionJSON) RawJSON() string {
	return r.raw
}

func (r *PageRuleAction) UnmarshalJSON(data []byte) (err error) {
	*r = PageRuleAction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PageRuleActionsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [zones.AlwaysUseHTTPS],
// [zones.AutomaticHTTPSRewrites], [zones.BrowserCacheTTL], [zones.BrowserCheck],
// [pagerules.PageRuleActionsBypassCacheOnCookie],
// [pagerules.PageRuleActionsCacheByDeviceType],
// [pagerules.PageRuleActionsCacheDeceptionArmor],
// [pagerules.PageRuleActionsCacheKey], [zones.CacheLevel],
// [pagerules.PageRuleActionsCacheOnCookie],
// [pagerules.PageRuleActionsDisableApps],
// [pagerules.PageRuleActionsDisablePerformance],
// [pagerules.PageRuleActionsDisableSecurity],
// [pagerules.PageRuleActionsDisableZaraz],
// [pagerules.PageRuleActionsEdgeCacheTTL], [zones.EmailObfuscation],
// [pagerules.PageRuleActionsExplicitCacheControl],
// [pagerules.PageRuleActionsForwardingURL],
// [pagerules.PageRuleActionsHostHeaderOverride], [zones.IPGeolocation],
// [zones.Mirage], [zones.OpportunisticEncryption],
// [zones.OriginErrorPagePassThru], [zones.Polish],
// [pagerules.PageRuleActionsResolveOverride],
// [pagerules.PageRuleActionsRespectStrongEtag], [zones.ResponseBuffering],
// [zones.RocketLoader], [zones.SecurityLevel], [zones.SortQueryStringForCache],
// [zones.SSL], [zones.TrueClientIPHeader], [zones.WAF].
func (r PageRuleAction) AsUnion() PageRuleActionsUnion {
	return r.union
}

// Union satisfied by [zones.AlwaysUseHTTPS], [zones.AutomaticHTTPSRewrites],
// [zones.BrowserCacheTTL], [zones.BrowserCheck],
// [pagerules.PageRuleActionsBypassCacheOnCookie],
// [pagerules.PageRuleActionsCacheByDeviceType],
// [pagerules.PageRuleActionsCacheDeceptionArmor],
// [pagerules.PageRuleActionsCacheKey], [zones.CacheLevel],
// [pagerules.PageRuleActionsCacheOnCookie],
// [pagerules.PageRuleActionsDisableApps],
// [pagerules.PageRuleActionsDisablePerformance],
// [pagerules.PageRuleActionsDisableSecurity],
// [pagerules.PageRuleActionsDisableZaraz],
// [pagerules.PageRuleActionsEdgeCacheTTL], [zones.EmailObfuscation],
// [pagerules.PageRuleActionsExplicitCacheControl],
// [pagerules.PageRuleActionsForwardingURL],
// [pagerules.PageRuleActionsHostHeaderOverride], [zones.IPGeolocation],
// [zones.Mirage], [zones.OpportunisticEncryption],
// [zones.OriginErrorPagePassThru], [zones.Polish],
// [pagerules.PageRuleActionsResolveOverride],
// [pagerules.PageRuleActionsRespectStrongEtag], [zones.ResponseBuffering],
// [zones.RocketLoader], [zones.SecurityLevel], [zones.SortQueryStringForCache],
// [zones.SSL], [zones.TrueClientIPHeader] or [zones.WAF].
type PageRuleActionsUnion interface {
	ImplementsPagerulesPageRuleAction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageRuleActionsUnion)(nil)).Elem(),
		"id",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.AlwaysUseHTTPS{}),
			DiscriminatorValue: "always_use_https",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.AutomaticHTTPSRewrites{}),
			DiscriminatorValue: "automatic_https_rewrites",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.BrowserCacheTTL{}),
			DiscriminatorValue: "browser_cache_ttl",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.BrowserCheck{}),
			DiscriminatorValue: "browser_check",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsBypassCacheOnCookie{}),
			DiscriminatorValue: "bypass_cache_on_cookie",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsCacheByDeviceType{}),
			DiscriminatorValue: "cache_by_device_type",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsCacheDeceptionArmor{}),
			DiscriminatorValue: "cache_deception_armor",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsCacheKey{}),
			DiscriminatorValue: "cache_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.CacheLevel{}),
			DiscriminatorValue: "cache_level",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsCacheOnCookie{}),
			DiscriminatorValue: "cache_on_cookie",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsDisableApps{}),
			DiscriminatorValue: "disable_apps",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsDisablePerformance{}),
			DiscriminatorValue: "disable_performance",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsDisableSecurity{}),
			DiscriminatorValue: "disable_security",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsDisableZaraz{}),
			DiscriminatorValue: "disable_zaraz",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsEdgeCacheTTL{}),
			DiscriminatorValue: "edge_cache_ttl",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.EmailObfuscation{}),
			DiscriminatorValue: "email_obfuscation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsExplicitCacheControl{}),
			DiscriminatorValue: "explicit_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsForwardingURL{}),
			DiscriminatorValue: "forwarding_url",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsHostHeaderOverride{}),
			DiscriminatorValue: "host_header_override",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.IPGeolocation{}),
			DiscriminatorValue: "ip_geolocation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.Mirage{}),
			DiscriminatorValue: "mirage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.OpportunisticEncryption{}),
			DiscriminatorValue: "opportunistic_encryption",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.OriginErrorPagePassThru{}),
			DiscriminatorValue: "origin_error_page_pass_thru",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.Polish{}),
			DiscriminatorValue: "polish",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsResolveOverride{}),
			DiscriminatorValue: "resolve_override",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PageRuleActionsRespectStrongEtag{}),
			DiscriminatorValue: "respect_strong_etag",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.ResponseBuffering{}),
			DiscriminatorValue: "response_buffering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.RocketLoader{}),
			DiscriminatorValue: "rocket_loader",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.SecurityLevel{}),
			DiscriminatorValue: "security_level",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.SortQueryStringForCache{}),
			DiscriminatorValue: "sort_query_string_for_cache",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.SSL{}),
			DiscriminatorValue: "ssl",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.TrueClientIPHeader{}),
			DiscriminatorValue: "true_client_ip_header",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(zones.WAF{}),
			DiscriminatorValue: "waf",
		},
	)
}

type PageRuleActionsBypassCacheOnCookie struct {
	// Bypass cache and fetch resources from the origin server if a regular expression
	// matches against a cookie name present in the request.
	ID PageRuleActionsBypassCacheOnCookieID `json:"id"`
	// The regular expression to use for matching cookie names in the request. Refer to
	// [Bypass Cache on Cookie setting](https://developers.cloudflare.com/rules/page-rules/reference/additional-reference/#bypass-cache-on-cookie-setting)
	// to learn about limited regular expression support.
	Value string                                 `json:"value"`
	JSON  pageRuleActionsBypassCacheOnCookieJSON `json:"-"`
}

// pageRuleActionsBypassCacheOnCookieJSON contains the JSON metadata for the struct
// [PageRuleActionsBypassCacheOnCookie]
type pageRuleActionsBypassCacheOnCookieJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsBypassCacheOnCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsBypassCacheOnCookieJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsBypassCacheOnCookie) ImplementsPagerulesPageRuleAction() {}

// Bypass cache and fetch resources from the origin server if a regular expression
// matches against a cookie name present in the request.
type PageRuleActionsBypassCacheOnCookieID string

const (
	PageRuleActionsBypassCacheOnCookieIDBypassCacheOnCookie PageRuleActionsBypassCacheOnCookieID = "bypass_cache_on_cookie"
)

func (r PageRuleActionsBypassCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageRuleActionsBypassCacheOnCookieIDBypassCacheOnCookie:
		return true
	}
	return false
}

type PageRuleActionsCacheByDeviceType struct {
	// Separate cached content based on the visitor's device type.
	ID PageRuleActionsCacheByDeviceTypeID `json:"id"`
	// The status of Cache By Device Type.
	Value PageRuleActionsCacheByDeviceTypeValue `json:"value"`
	JSON  pageRuleActionsCacheByDeviceTypeJSON  `json:"-"`
}

// pageRuleActionsCacheByDeviceTypeJSON contains the JSON metadata for the struct
// [PageRuleActionsCacheByDeviceType]
type pageRuleActionsCacheByDeviceTypeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsCacheByDeviceType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsCacheByDeviceTypeJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsCacheByDeviceType) ImplementsPagerulesPageRuleAction() {}

// Separate cached content based on the visitor's device type.
type PageRuleActionsCacheByDeviceTypeID string

const (
	PageRuleActionsCacheByDeviceTypeIDCacheByDeviceType PageRuleActionsCacheByDeviceTypeID = "cache_by_device_type"
)

func (r PageRuleActionsCacheByDeviceTypeID) IsKnown() bool {
	switch r {
	case PageRuleActionsCacheByDeviceTypeIDCacheByDeviceType:
		return true
	}
	return false
}

// The status of Cache By Device Type.
type PageRuleActionsCacheByDeviceTypeValue string

const (
	PageRuleActionsCacheByDeviceTypeValueOn  PageRuleActionsCacheByDeviceTypeValue = "on"
	PageRuleActionsCacheByDeviceTypeValueOff PageRuleActionsCacheByDeviceTypeValue = "off"
)

func (r PageRuleActionsCacheByDeviceTypeValue) IsKnown() bool {
	switch r {
	case PageRuleActionsCacheByDeviceTypeValueOn, PageRuleActionsCacheByDeviceTypeValueOff:
		return true
	}
	return false
}

type PageRuleActionsCacheDeceptionArmor struct {
	// Protect from web cache deception attacks while still allowing static assets to
	// be cached. This setting verifies that the URL's extension matches the returned
	// `Content-Type`.
	ID PageRuleActionsCacheDeceptionArmorID `json:"id"`
	// The status of Cache Deception Armor.
	Value PageRuleActionsCacheDeceptionArmorValue `json:"value"`
	JSON  pageRuleActionsCacheDeceptionArmorJSON  `json:"-"`
}

// pageRuleActionsCacheDeceptionArmorJSON contains the JSON metadata for the struct
// [PageRuleActionsCacheDeceptionArmor]
type pageRuleActionsCacheDeceptionArmorJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsCacheDeceptionArmor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsCacheDeceptionArmorJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsCacheDeceptionArmor) ImplementsPagerulesPageRuleAction() {}

// Protect from web cache deception attacks while still allowing static assets to
// be cached. This setting verifies that the URL's extension matches the returned
// `Content-Type`.
type PageRuleActionsCacheDeceptionArmorID string

const (
	PageRuleActionsCacheDeceptionArmorIDCacheDeceptionArmor PageRuleActionsCacheDeceptionArmorID = "cache_deception_armor"
)

func (r PageRuleActionsCacheDeceptionArmorID) IsKnown() bool {
	switch r {
	case PageRuleActionsCacheDeceptionArmorIDCacheDeceptionArmor:
		return true
	}
	return false
}

// The status of Cache Deception Armor.
type PageRuleActionsCacheDeceptionArmorValue string

const (
	PageRuleActionsCacheDeceptionArmorValueOn  PageRuleActionsCacheDeceptionArmorValue = "on"
	PageRuleActionsCacheDeceptionArmorValueOff PageRuleActionsCacheDeceptionArmorValue = "off"
)

func (r PageRuleActionsCacheDeceptionArmorValue) IsKnown() bool {
	switch r {
	case PageRuleActionsCacheDeceptionArmorValueOn, PageRuleActionsCacheDeceptionArmorValueOff:
		return true
	}
	return false
}

type PageRuleActionsCacheKey struct {
	// Control specifically what variables to include when deciding which resources to
	// cache. This allows customers to determine what to cache based on something other
	// than just the URL.
	ID    PageRuleActionsCacheKeyID    `json:"id"`
	Value PageRuleActionsCacheKeyValue `json:"value"`
	JSON  pageRuleActionsCacheKeyJSON  `json:"-"`
}

// pageRuleActionsCacheKeyJSON contains the JSON metadata for the struct
// [PageRuleActionsCacheKey]
type pageRuleActionsCacheKeyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsCacheKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsCacheKeyJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsCacheKey) ImplementsPagerulesPageRuleAction() {}

// Control specifically what variables to include when deciding which resources to
// cache. This allows customers to determine what to cache based on something other
// than just the URL.
type PageRuleActionsCacheKeyID string

const (
	PageRuleActionsCacheKeyIDCacheKey PageRuleActionsCacheKeyID = "cache_key"
)

func (r PageRuleActionsCacheKeyID) IsKnown() bool {
	switch r {
	case PageRuleActionsCacheKeyIDCacheKey:
		return true
	}
	return false
}

type PageRuleActionsCacheKeyValue struct {
	// Controls which cookies appear in the Cache Key.
	Cookie PageRuleActionsCacheKeyValueCookie `json:"cookie"`
	// Controls which headers go into the Cache Key. Exactly one of `include` or
	// `exclude` is expected.
	Header PageRuleActionsCacheKeyValueHeader `json:"header"`
	// Determines which host header to include in the Cache Key.
	Host PageRuleActionsCacheKeyValueHost `json:"host"`
	// Controls which URL query string parameters go into the Cache Key. Exactly one of
	// `include` or `exclude` is expected.
	QueryString PageRuleActionsCacheKeyValueQueryString `json:"query_string"`
	// Feature fields to add features about the end-user (client) into the Cache Key.
	User PageRuleActionsCacheKeyValueUser `json:"user"`
	JSON pageRuleActionsCacheKeyValueJSON `json:"-"`
}

// pageRuleActionsCacheKeyValueJSON contains the JSON metadata for the struct
// [PageRuleActionsCacheKeyValue]
type pageRuleActionsCacheKeyValueJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsCacheKeyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsCacheKeyValueJSON) RawJSON() string {
	return r.raw
}

// Controls which cookies appear in the Cache Key.
type PageRuleActionsCacheKeyValueCookie struct {
	// A list of cookies to check for the presence of, without including their actual
	// values.
	CheckPresence []string `json:"check_presence"`
	// A list of cookies to include.
	Include []string                               `json:"include"`
	JSON    pageRuleActionsCacheKeyValueCookieJSON `json:"-"`
}

// pageRuleActionsCacheKeyValueCookieJSON contains the JSON metadata for the struct
// [PageRuleActionsCacheKeyValueCookie]
type pageRuleActionsCacheKeyValueCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PageRuleActionsCacheKeyValueCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsCacheKeyValueCookieJSON) RawJSON() string {
	return r.raw
}

// Controls which headers go into the Cache Key. Exactly one of `include` or
// `exclude` is expected.
type PageRuleActionsCacheKeyValueHeader struct {
	// A list of headers to check for the presence of, without including their actual
	// values.
	CheckPresence []string `json:"check_presence"`
	// A list of headers to ignore.
	Exclude []string `json:"exclude"`
	// A list of headers to include.
	Include []string                               `json:"include"`
	JSON    pageRuleActionsCacheKeyValueHeaderJSON `json:"-"`
}

// pageRuleActionsCacheKeyValueHeaderJSON contains the JSON metadata for the struct
// [PageRuleActionsCacheKeyValueHeader]
type pageRuleActionsCacheKeyValueHeaderJSON struct {
	CheckPresence apijson.Field
	Exclude       apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PageRuleActionsCacheKeyValueHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsCacheKeyValueHeaderJSON) RawJSON() string {
	return r.raw
}

// Determines which host header to include in the Cache Key.
type PageRuleActionsCacheKeyValueHost struct {
	// Whether to include the Host header in the HTTP request sent to the origin.
	Resolved bool                                 `json:"resolved"`
	JSON     pageRuleActionsCacheKeyValueHostJSON `json:"-"`
}

// pageRuleActionsCacheKeyValueHostJSON contains the JSON metadata for the struct
// [PageRuleActionsCacheKeyValueHost]
type pageRuleActionsCacheKeyValueHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsCacheKeyValueHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsCacheKeyValueHostJSON) RawJSON() string {
	return r.raw
}

// Controls which URL query string parameters go into the Cache Key. Exactly one of
// `include` or `exclude` is expected.
type PageRuleActionsCacheKeyValueQueryString struct {
	// Ignore all query string parameters.
	Exclude PageRuleActionsCacheKeyValueQueryStringExcludeUnion `json:"exclude"`
	// Include all query string parameters.
	Include PageRuleActionsCacheKeyValueQueryStringIncludeUnion `json:"include"`
	JSON    pageRuleActionsCacheKeyValueQueryStringJSON         `json:"-"`
}

// pageRuleActionsCacheKeyValueQueryStringJSON contains the JSON metadata for the
// struct [PageRuleActionsCacheKeyValueQueryString]
type pageRuleActionsCacheKeyValueQueryStringJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsCacheKeyValueQueryString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsCacheKeyValueQueryStringJSON) RawJSON() string {
	return r.raw
}

// Ignore all query string parameters.
//
// Union satisfied by
// [pagerules.PageRuleActionsCacheKeyValueQueryStringExcludeString] or
// [pagerules.PageRuleActionsCacheKeyValueQueryStringExcludeArray].
type PageRuleActionsCacheKeyValueQueryStringExcludeUnion interface {
	implementsPagerulesPageRuleActionsCacheKeyValueQueryStringExcludeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageRuleActionsCacheKeyValueQueryStringExcludeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PageRuleActionsCacheKeyValueQueryStringExcludeString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PageRuleActionsCacheKeyValueQueryStringExcludeArray{}),
		},
	)
}

// Ignore all query string parameters.
type PageRuleActionsCacheKeyValueQueryStringExcludeString string

const (
	PageRuleActionsCacheKeyValueQueryStringExcludeStringStar PageRuleActionsCacheKeyValueQueryStringExcludeString = "*"
)

func (r PageRuleActionsCacheKeyValueQueryStringExcludeString) IsKnown() bool {
	switch r {
	case PageRuleActionsCacheKeyValueQueryStringExcludeStringStar:
		return true
	}
	return false
}

func (r PageRuleActionsCacheKeyValueQueryStringExcludeString) implementsPagerulesPageRuleActionsCacheKeyValueQueryStringExcludeUnion() {
}

type PageRuleActionsCacheKeyValueQueryStringExcludeArray []string

func (r PageRuleActionsCacheKeyValueQueryStringExcludeArray) implementsPagerulesPageRuleActionsCacheKeyValueQueryStringExcludeUnion() {
}

// Include all query string parameters.
//
// Union satisfied by
// [pagerules.PageRuleActionsCacheKeyValueQueryStringIncludeString] or
// [pagerules.PageRuleActionsCacheKeyValueQueryStringIncludeArray].
type PageRuleActionsCacheKeyValueQueryStringIncludeUnion interface {
	implementsPagerulesPageRuleActionsCacheKeyValueQueryStringIncludeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageRuleActionsCacheKeyValueQueryStringIncludeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PageRuleActionsCacheKeyValueQueryStringIncludeString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PageRuleActionsCacheKeyValueQueryStringIncludeArray{}),
		},
	)
}

// Include all query string parameters.
type PageRuleActionsCacheKeyValueQueryStringIncludeString string

const (
	PageRuleActionsCacheKeyValueQueryStringIncludeStringStar PageRuleActionsCacheKeyValueQueryStringIncludeString = "*"
)

func (r PageRuleActionsCacheKeyValueQueryStringIncludeString) IsKnown() bool {
	switch r {
	case PageRuleActionsCacheKeyValueQueryStringIncludeStringStar:
		return true
	}
	return false
}

func (r PageRuleActionsCacheKeyValueQueryStringIncludeString) implementsPagerulesPageRuleActionsCacheKeyValueQueryStringIncludeUnion() {
}

type PageRuleActionsCacheKeyValueQueryStringIncludeArray []string

func (r PageRuleActionsCacheKeyValueQueryStringIncludeArray) implementsPagerulesPageRuleActionsCacheKeyValueQueryStringIncludeUnion() {
}

// Feature fields to add features about the end-user (client) into the Cache Key.
type PageRuleActionsCacheKeyValueUser struct {
	// Classifies a request as `mobile`, `desktop`, or `tablet` based on the User
	// Agent.
	DeviceType bool `json:"device_type"`
	// Includes the client's country, derived from the IP address.
	Geo bool `json:"geo"`
	// Includes the first language code contained in the `Accept-Language` header sent
	// by the client.
	Lang bool                                 `json:"lang"`
	JSON pageRuleActionsCacheKeyValueUserJSON `json:"-"`
}

// pageRuleActionsCacheKeyValueUserJSON contains the JSON metadata for the struct
// [PageRuleActionsCacheKeyValueUser]
type pageRuleActionsCacheKeyValueUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsCacheKeyValueUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsCacheKeyValueUserJSON) RawJSON() string {
	return r.raw
}

type PageRuleActionsCacheOnCookie struct {
	// Apply the Cache Everything option (Cache Level setting) based on a regular
	// expression match against a cookie name.
	ID PageRuleActionsCacheOnCookieID `json:"id"`
	// The regular expression to use for matching cookie names in the request.
	Value string                           `json:"value"`
	JSON  pageRuleActionsCacheOnCookieJSON `json:"-"`
}

// pageRuleActionsCacheOnCookieJSON contains the JSON metadata for the struct
// [PageRuleActionsCacheOnCookie]
type pageRuleActionsCacheOnCookieJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsCacheOnCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsCacheOnCookieJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsCacheOnCookie) ImplementsPagerulesPageRuleAction() {}

// Apply the Cache Everything option (Cache Level setting) based on a regular
// expression match against a cookie name.
type PageRuleActionsCacheOnCookieID string

const (
	PageRuleActionsCacheOnCookieIDCacheOnCookie PageRuleActionsCacheOnCookieID = "cache_on_cookie"
)

func (r PageRuleActionsCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageRuleActionsCacheOnCookieIDCacheOnCookie:
		return true
	}
	return false
}

type PageRuleActionsDisableApps struct {
	// Turn off all active
	// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
	// (deprecated).
	ID   PageRuleActionsDisableAppsID   `json:"id"`
	JSON pageRuleActionsDisableAppsJSON `json:"-"`
}

// pageRuleActionsDisableAppsJSON contains the JSON metadata for the struct
// [PageRuleActionsDisableApps]
type pageRuleActionsDisableAppsJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsDisableApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsDisableAppsJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsDisableApps) ImplementsPagerulesPageRuleAction() {}

// Turn off all active
// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
// (deprecated).
type PageRuleActionsDisableAppsID string

const (
	PageRuleActionsDisableAppsIDDisableApps PageRuleActionsDisableAppsID = "disable_apps"
)

func (r PageRuleActionsDisableAppsID) IsKnown() bool {
	switch r {
	case PageRuleActionsDisableAppsIDDisableApps:
		return true
	}
	return false
}

type PageRuleActionsDisablePerformance struct {
	// Turn off
	// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
	// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
	// and [Polish](https://developers.cloudflare.com/images/polish/).
	ID   PageRuleActionsDisablePerformanceID   `json:"id"`
	JSON pageRuleActionsDisablePerformanceJSON `json:"-"`
}

// pageRuleActionsDisablePerformanceJSON contains the JSON metadata for the struct
// [PageRuleActionsDisablePerformance]
type pageRuleActionsDisablePerformanceJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsDisablePerformance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsDisablePerformanceJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsDisablePerformance) ImplementsPagerulesPageRuleAction() {}

// Turn off
// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
// and [Polish](https://developers.cloudflare.com/images/polish/).
type PageRuleActionsDisablePerformanceID string

const (
	PageRuleActionsDisablePerformanceIDDisablePerformance PageRuleActionsDisablePerformanceID = "disable_performance"
)

func (r PageRuleActionsDisablePerformanceID) IsKnown() bool {
	switch r {
	case PageRuleActionsDisablePerformanceIDDisablePerformance:
		return true
	}
	return false
}

type PageRuleActionsDisableSecurity struct {
	// Turn off
	// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
	// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
	// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
	// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
	// and
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	ID   PageRuleActionsDisableSecurityID   `json:"id"`
	JSON pageRuleActionsDisableSecurityJSON `json:"-"`
}

// pageRuleActionsDisableSecurityJSON contains the JSON metadata for the struct
// [PageRuleActionsDisableSecurity]
type pageRuleActionsDisableSecurityJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsDisableSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsDisableSecurityJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsDisableSecurity) ImplementsPagerulesPageRuleAction() {}

// Turn off
// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
// and
// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
type PageRuleActionsDisableSecurityID string

const (
	PageRuleActionsDisableSecurityIDDisableSecurity PageRuleActionsDisableSecurityID = "disable_security"
)

func (r PageRuleActionsDisableSecurityID) IsKnown() bool {
	switch r {
	case PageRuleActionsDisableSecurityIDDisableSecurity:
		return true
	}
	return false
}

type PageRuleActionsDisableZaraz struct {
	// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
	ID   PageRuleActionsDisableZarazID   `json:"id"`
	JSON pageRuleActionsDisableZarazJSON `json:"-"`
}

// pageRuleActionsDisableZarazJSON contains the JSON metadata for the struct
// [PageRuleActionsDisableZaraz]
type pageRuleActionsDisableZarazJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsDisableZaraz) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsDisableZarazJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsDisableZaraz) ImplementsPagerulesPageRuleAction() {}

// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
type PageRuleActionsDisableZarazID string

const (
	PageRuleActionsDisableZarazIDDisableZaraz PageRuleActionsDisableZarazID = "disable_zaraz"
)

func (r PageRuleActionsDisableZarazID) IsKnown() bool {
	switch r {
	case PageRuleActionsDisableZarazIDDisableZaraz:
		return true
	}
	return false
}

type PageRuleActionsEdgeCacheTTL struct {
	// Specify how long to cache a resource in the Cloudflare global network. _Edge
	// Cache TTL_ is not visible in response headers.
	ID    PageRuleActionsEdgeCacheTTLID   `json:"id"`
	Value int64                           `json:"value"`
	JSON  pageRuleActionsEdgeCacheTTLJSON `json:"-"`
}

// pageRuleActionsEdgeCacheTTLJSON contains the JSON metadata for the struct
// [PageRuleActionsEdgeCacheTTL]
type pageRuleActionsEdgeCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsEdgeCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsEdgeCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsEdgeCacheTTL) ImplementsPagerulesPageRuleAction() {}

// Specify how long to cache a resource in the Cloudflare global network. _Edge
// Cache TTL_ is not visible in response headers.
type PageRuleActionsEdgeCacheTTLID string

const (
	PageRuleActionsEdgeCacheTTLIDEdgeCacheTTL PageRuleActionsEdgeCacheTTLID = "edge_cache_ttl"
)

func (r PageRuleActionsEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case PageRuleActionsEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

type PageRuleActionsExplicitCacheControl struct {
	// Origin Cache Control is enabled by default for Free, Pro, and Business domains
	// and disabled by default for Enterprise domains.
	ID PageRuleActionsExplicitCacheControlID `json:"id"`
	// The status of Origin Cache Control.
	Value PageRuleActionsExplicitCacheControlValue `json:"value"`
	JSON  pageRuleActionsExplicitCacheControlJSON  `json:"-"`
}

// pageRuleActionsExplicitCacheControlJSON contains the JSON metadata for the
// struct [PageRuleActionsExplicitCacheControl]
type pageRuleActionsExplicitCacheControlJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsExplicitCacheControl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsExplicitCacheControlJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsExplicitCacheControl) ImplementsPagerulesPageRuleAction() {}

// Origin Cache Control is enabled by default for Free, Pro, and Business domains
// and disabled by default for Enterprise domains.
type PageRuleActionsExplicitCacheControlID string

const (
	PageRuleActionsExplicitCacheControlIDExplicitCacheControl PageRuleActionsExplicitCacheControlID = "explicit_cache_control"
)

func (r PageRuleActionsExplicitCacheControlID) IsKnown() bool {
	switch r {
	case PageRuleActionsExplicitCacheControlIDExplicitCacheControl:
		return true
	}
	return false
}

// The status of Origin Cache Control.
type PageRuleActionsExplicitCacheControlValue string

const (
	PageRuleActionsExplicitCacheControlValueOn  PageRuleActionsExplicitCacheControlValue = "on"
	PageRuleActionsExplicitCacheControlValueOff PageRuleActionsExplicitCacheControlValue = "off"
)

func (r PageRuleActionsExplicitCacheControlValue) IsKnown() bool {
	switch r {
	case PageRuleActionsExplicitCacheControlValueOn, PageRuleActionsExplicitCacheControlValueOff:
		return true
	}
	return false
}

type PageRuleActionsForwardingURL struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    PageRuleActionsForwardingURLID    `json:"id"`
	Value PageRuleActionsForwardingURLValue `json:"value"`
	JSON  pageRuleActionsForwardingURLJSON  `json:"-"`
}

// pageRuleActionsForwardingURLJSON contains the JSON metadata for the struct
// [PageRuleActionsForwardingURL]
type pageRuleActionsForwardingURLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsForwardingURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsForwardingURLJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsForwardingURL) ImplementsPagerulesPageRuleAction() {}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageRuleActionsForwardingURLID string

const (
	PageRuleActionsForwardingURLIDForwardingURL PageRuleActionsForwardingURLID = "forwarding_url"
)

func (r PageRuleActionsForwardingURLID) IsKnown() bool {
	switch r {
	case PageRuleActionsForwardingURLIDForwardingURL:
		return true
	}
	return false
}

type PageRuleActionsForwardingURLValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode PageRuleActionsForwardingURLValueStatusCode `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string                                `json:"url"`
	JSON pageRuleActionsForwardingURLValueJSON `json:"-"`
}

// pageRuleActionsForwardingURLValueJSON contains the JSON metadata for the struct
// [PageRuleActionsForwardingURLValue]
type pageRuleActionsForwardingURLValueJSON struct {
	StatusCode  apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsForwardingURLValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsForwardingURLValueJSON) RawJSON() string {
	return r.raw
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageRuleActionsForwardingURLValueStatusCode int64

const (
	PageRuleActionsForwardingURLValueStatusCode301 PageRuleActionsForwardingURLValueStatusCode = 301
	PageRuleActionsForwardingURLValueStatusCode302 PageRuleActionsForwardingURLValueStatusCode = 302
)

func (r PageRuleActionsForwardingURLValueStatusCode) IsKnown() bool {
	switch r {
	case PageRuleActionsForwardingURLValueStatusCode301, PageRuleActionsForwardingURLValueStatusCode302:
		return true
	}
	return false
}

type PageRuleActionsHostHeaderOverride struct {
	// Apply a specific host header.
	ID PageRuleActionsHostHeaderOverrideID `json:"id"`
	// The hostname to use in the `Host` header
	Value string                                `json:"value"`
	JSON  pageRuleActionsHostHeaderOverrideJSON `json:"-"`
}

// pageRuleActionsHostHeaderOverrideJSON contains the JSON metadata for the struct
// [PageRuleActionsHostHeaderOverride]
type pageRuleActionsHostHeaderOverrideJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsHostHeaderOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsHostHeaderOverrideJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsHostHeaderOverride) ImplementsPagerulesPageRuleAction() {}

// Apply a specific host header.
type PageRuleActionsHostHeaderOverrideID string

const (
	PageRuleActionsHostHeaderOverrideIDHostHeaderOverride PageRuleActionsHostHeaderOverrideID = "host_header_override"
)

func (r PageRuleActionsHostHeaderOverrideID) IsKnown() bool {
	switch r {
	case PageRuleActionsHostHeaderOverrideIDHostHeaderOverride:
		return true
	}
	return false
}

type PageRuleActionsResolveOverride struct {
	// Change the origin address to the value specified in this setting.
	ID PageRuleActionsResolveOverrideID `json:"id"`
	// The origin address you want to override with.
	Value string                             `json:"value"`
	JSON  pageRuleActionsResolveOverrideJSON `json:"-"`
}

// pageRuleActionsResolveOverrideJSON contains the JSON metadata for the struct
// [PageRuleActionsResolveOverride]
type pageRuleActionsResolveOverrideJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsResolveOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsResolveOverrideJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsResolveOverride) ImplementsPagerulesPageRuleAction() {}

// Change the origin address to the value specified in this setting.
type PageRuleActionsResolveOverrideID string

const (
	PageRuleActionsResolveOverrideIDResolveOverride PageRuleActionsResolveOverrideID = "resolve_override"
)

func (r PageRuleActionsResolveOverrideID) IsKnown() bool {
	switch r {
	case PageRuleActionsResolveOverrideIDResolveOverride:
		return true
	}
	return false
}

type PageRuleActionsRespectStrongEtag struct {
	// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
	// the origin server.
	ID PageRuleActionsRespectStrongEtagID `json:"id"`
	// The status of Respect Strong ETags
	Value PageRuleActionsRespectStrongEtagValue `json:"value"`
	JSON  pageRuleActionsRespectStrongEtagJSON  `json:"-"`
}

// pageRuleActionsRespectStrongEtagJSON contains the JSON metadata for the struct
// [PageRuleActionsRespectStrongEtag]
type pageRuleActionsRespectStrongEtagJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleActionsRespectStrongEtag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleActionsRespectStrongEtagJSON) RawJSON() string {
	return r.raw
}

func (r PageRuleActionsRespectStrongEtag) ImplementsPagerulesPageRuleAction() {}

// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
// the origin server.
type PageRuleActionsRespectStrongEtagID string

const (
	PageRuleActionsRespectStrongEtagIDRespectStrongEtag PageRuleActionsRespectStrongEtagID = "respect_strong_etag"
)

func (r PageRuleActionsRespectStrongEtagID) IsKnown() bool {
	switch r {
	case PageRuleActionsRespectStrongEtagIDRespectStrongEtag:
		return true
	}
	return false
}

// The status of Respect Strong ETags
type PageRuleActionsRespectStrongEtagValue string

const (
	PageRuleActionsRespectStrongEtagValueOn  PageRuleActionsRespectStrongEtagValue = "on"
	PageRuleActionsRespectStrongEtagValueOff PageRuleActionsRespectStrongEtagValue = "off"
)

func (r PageRuleActionsRespectStrongEtagValue) IsKnown() bool {
	switch r {
	case PageRuleActionsRespectStrongEtagValueOn, PageRuleActionsRespectStrongEtagValueOff:
		return true
	}
	return false
}

// If enabled, any ` http://“ URL is converted to  `https://` through a 301
// redirect.
type PageRuleActionsID string

const (
	PageRuleActionsIDAlwaysUseHTTPS          PageRuleActionsID = "always_use_https"
	PageRuleActionsIDAutomaticHTTPSRewrites  PageRuleActionsID = "automatic_https_rewrites"
	PageRuleActionsIDBrowserCacheTTL         PageRuleActionsID = "browser_cache_ttl"
	PageRuleActionsIDBrowserCheck            PageRuleActionsID = "browser_check"
	PageRuleActionsIDBypassCacheOnCookie     PageRuleActionsID = "bypass_cache_on_cookie"
	PageRuleActionsIDCacheByDeviceType       PageRuleActionsID = "cache_by_device_type"
	PageRuleActionsIDCacheDeceptionArmor     PageRuleActionsID = "cache_deception_armor"
	PageRuleActionsIDCacheKey                PageRuleActionsID = "cache_key"
	PageRuleActionsIDCacheLevel              PageRuleActionsID = "cache_level"
	PageRuleActionsIDCacheOnCookie           PageRuleActionsID = "cache_on_cookie"
	PageRuleActionsIDDisableApps             PageRuleActionsID = "disable_apps"
	PageRuleActionsIDDisablePerformance      PageRuleActionsID = "disable_performance"
	PageRuleActionsIDDisableSecurity         PageRuleActionsID = "disable_security"
	PageRuleActionsIDDisableZaraz            PageRuleActionsID = "disable_zaraz"
	PageRuleActionsIDEdgeCacheTTL            PageRuleActionsID = "edge_cache_ttl"
	PageRuleActionsIDEmailObfuscation        PageRuleActionsID = "email_obfuscation"
	PageRuleActionsIDExplicitCacheControl    PageRuleActionsID = "explicit_cache_control"
	PageRuleActionsIDForwardingURL           PageRuleActionsID = "forwarding_url"
	PageRuleActionsIDHostHeaderOverride      PageRuleActionsID = "host_header_override"
	PageRuleActionsIDIPGeolocation           PageRuleActionsID = "ip_geolocation"
	PageRuleActionsIDMirage                  PageRuleActionsID = "mirage"
	PageRuleActionsIDOpportunisticEncryption PageRuleActionsID = "opportunistic_encryption"
	PageRuleActionsIDOriginErrorPagePassThru PageRuleActionsID = "origin_error_page_pass_thru"
	PageRuleActionsIDPolish                  PageRuleActionsID = "polish"
	PageRuleActionsIDResolveOverride         PageRuleActionsID = "resolve_override"
	PageRuleActionsIDRespectStrongEtag       PageRuleActionsID = "respect_strong_etag"
	PageRuleActionsIDResponseBuffering       PageRuleActionsID = "response_buffering"
	PageRuleActionsIDRocketLoader            PageRuleActionsID = "rocket_loader"
	PageRuleActionsIDSecurityLevel           PageRuleActionsID = "security_level"
	PageRuleActionsIDSortQueryStringForCache PageRuleActionsID = "sort_query_string_for_cache"
	PageRuleActionsIDSSL                     PageRuleActionsID = "ssl"
	PageRuleActionsIDTrueClientIPHeader      PageRuleActionsID = "true_client_ip_header"
	PageRuleActionsIDWAF                     PageRuleActionsID = "waf"
)

func (r PageRuleActionsID) IsKnown() bool {
	switch r {
	case PageRuleActionsIDAlwaysUseHTTPS, PageRuleActionsIDAutomaticHTTPSRewrites, PageRuleActionsIDBrowserCacheTTL, PageRuleActionsIDBrowserCheck, PageRuleActionsIDBypassCacheOnCookie, PageRuleActionsIDCacheByDeviceType, PageRuleActionsIDCacheDeceptionArmor, PageRuleActionsIDCacheKey, PageRuleActionsIDCacheLevel, PageRuleActionsIDCacheOnCookie, PageRuleActionsIDDisableApps, PageRuleActionsIDDisablePerformance, PageRuleActionsIDDisableSecurity, PageRuleActionsIDDisableZaraz, PageRuleActionsIDEdgeCacheTTL, PageRuleActionsIDEmailObfuscation, PageRuleActionsIDExplicitCacheControl, PageRuleActionsIDForwardingURL, PageRuleActionsIDHostHeaderOverride, PageRuleActionsIDIPGeolocation, PageRuleActionsIDMirage, PageRuleActionsIDOpportunisticEncryption, PageRuleActionsIDOriginErrorPagePassThru, PageRuleActionsIDPolish, PageRuleActionsIDResolveOverride, PageRuleActionsIDRespectStrongEtag, PageRuleActionsIDResponseBuffering, PageRuleActionsIDRocketLoader, PageRuleActionsIDSecurityLevel, PageRuleActionsIDSortQueryStringForCache, PageRuleActionsIDSSL, PageRuleActionsIDTrueClientIPHeader, PageRuleActionsIDWAF:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageRuleStatus string

const (
	PageRuleStatusActive   PageRuleStatus = "active"
	PageRuleStatusDisabled PageRuleStatus = "disabled"
)

func (r PageRuleStatus) IsKnown() bool {
	switch r {
	case PageRuleStatusActive, PageRuleStatusDisabled:
		return true
	}
	return false
}

// URL target.
type Target struct {
	// String constraint.
	Constraint TargetConstraint `json:"constraint"`
	// A target based on the URL of the request.
	Target TargetTarget `json:"target"`
	JSON   targetJSON   `json:"-"`
}

// targetJSON contains the JSON metadata for the struct [Target]
type targetJSON struct {
	Constraint  apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Target) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r targetJSON) RawJSON() string {
	return r.raw
}

// String constraint.
type TargetConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator TargetConstraintOperator `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value string               `json:"value,required"`
	JSON  targetConstraintJSON `json:"-"`
}

// targetConstraintJSON contains the JSON metadata for the struct
// [TargetConstraint]
type targetConstraintJSON struct {
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TargetConstraint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r targetConstraintJSON) RawJSON() string {
	return r.raw
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type TargetConstraintOperator string

const (
	TargetConstraintOperatorMatches    TargetConstraintOperator = "matches"
	TargetConstraintOperatorContains   TargetConstraintOperator = "contains"
	TargetConstraintOperatorEquals     TargetConstraintOperator = "equals"
	TargetConstraintOperatorNotEqual   TargetConstraintOperator = "not_equal"
	TargetConstraintOperatorNotContain TargetConstraintOperator = "not_contain"
)

func (r TargetConstraintOperator) IsKnown() bool {
	switch r {
	case TargetConstraintOperatorMatches, TargetConstraintOperatorContains, TargetConstraintOperatorEquals, TargetConstraintOperatorNotEqual, TargetConstraintOperatorNotContain:
		return true
	}
	return false
}

// A target based on the URL of the request.
type TargetTarget string

const (
	TargetTargetURL TargetTarget = "url"
)

func (r TargetTarget) IsKnown() bool {
	switch r {
	case TargetTargetURL:
		return true
	}
	return false
}

// URL target.
type TargetParam struct {
	// String constraint.
	Constraint param.Field[TargetConstraintParam] `json:"constraint"`
	// A target based on the URL of the request.
	Target param.Field[TargetTarget] `json:"target"`
}

func (r TargetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// String constraint.
type TargetConstraintParam struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator param.Field[TargetConstraintOperator] `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value param.Field[string] `json:"value,required"`
}

func (r TargetConstraintParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleDeleteResponse struct {
	// Identifier
	ID   string                     `json:"id,required"`
	JSON pageruleDeleteResponseJSON `json:"-"`
}

// pageruleDeleteResponseJSON contains the JSON metadata for the struct
// [PageruleDeleteResponse]
type pageruleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PageruleNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]PageruleNewParamsActionUnion] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargetParam] `json:"targets,required"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageruleNewParamsStatus] `json:"status"`
}

func (r PageruleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleNewParamsAction struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID    param.Field[PageruleNewParamsActionsID] `json:"id"`
	Value param.Field[interface{}]                `json:"value"`
}

func (r PageruleNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsAction) ImplementsPagerulesPageruleNewParamsActionUnion() {}

// Satisfied by [zones.AlwaysUseHTTPSParam], [zones.AutomaticHTTPSRewritesParam],
// [zones.BrowserCacheTTLParam], [zones.BrowserCheckParam],
// [pagerules.PageruleNewParamsActionsBypassCacheOnCookie],
// [pagerules.PageruleNewParamsActionsCacheByDeviceType],
// [pagerules.PageruleNewParamsActionsCacheDeceptionArmor],
// [pagerules.PageruleNewParamsActionsCacheKey], [zones.CacheLevelParam],
// [pagerules.PageruleNewParamsActionsCacheOnCookie],
// [pagerules.PageruleNewParamsActionsDisableApps],
// [pagerules.PageruleNewParamsActionsDisablePerformance],
// [pagerules.PageruleNewParamsActionsDisableSecurity],
// [pagerules.PageruleNewParamsActionsDisableZaraz],
// [pagerules.PageruleNewParamsActionsEdgeCacheTTL], [zones.EmailObfuscationParam],
// [pagerules.PageruleNewParamsActionsExplicitCacheControl],
// [pagerules.PageruleNewParamsActionsForwardingURL],
// [pagerules.PageruleNewParamsActionsHostHeaderOverride],
// [zones.IPGeolocationParam], [zones.MirageParam],
// [zones.OpportunisticEncryptionParam], [zones.OriginErrorPagePassThruParam],
// [zones.PolishParam], [pagerules.PageruleNewParamsActionsResolveOverride],
// [pagerules.PageruleNewParamsActionsRespectStrongEtag],
// [zones.ResponseBufferingParam], [zones.RocketLoaderParam],
// [zones.SecurityLevelParam], [zones.SortQueryStringForCacheParam],
// [zones.SSLParam], [zones.TrueClientIPHeaderParam], [zones.WAFParam],
// [PageruleNewParamsAction].
type PageruleNewParamsActionUnion interface {
	ImplementsPagerulesPageruleNewParamsActionUnion()
}

type PageruleNewParamsActionsBypassCacheOnCookie struct {
	// Bypass cache and fetch resources from the origin server if a regular expression
	// matches against a cookie name present in the request.
	ID param.Field[PageruleNewParamsActionsBypassCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request. Refer to
	// [Bypass Cache on Cookie setting](https://developers.cloudflare.com/rules/page-rules/reference/additional-reference/#bypass-cache-on-cookie-setting)
	// to learn about limited regular expression support.
	Value param.Field[string] `json:"value"`
}

func (r PageruleNewParamsActionsBypassCacheOnCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsBypassCacheOnCookie) ImplementsPagerulesPageruleNewParamsActionUnion() {
}

// Bypass cache and fetch resources from the origin server if a regular expression
// matches against a cookie name present in the request.
type PageruleNewParamsActionsBypassCacheOnCookieID string

const (
	PageruleNewParamsActionsBypassCacheOnCookieIDBypassCacheOnCookie PageruleNewParamsActionsBypassCacheOnCookieID = "bypass_cache_on_cookie"
)

func (r PageruleNewParamsActionsBypassCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsBypassCacheOnCookieIDBypassCacheOnCookie:
		return true
	}
	return false
}

type PageruleNewParamsActionsCacheByDeviceType struct {
	// Separate cached content based on the visitor's device type.
	ID param.Field[PageruleNewParamsActionsCacheByDeviceTypeID] `json:"id"`
	// The status of Cache By Device Type.
	Value param.Field[PageruleNewParamsActionsCacheByDeviceTypeValue] `json:"value"`
}

func (r PageruleNewParamsActionsCacheByDeviceType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsCacheByDeviceType) ImplementsPagerulesPageruleNewParamsActionUnion() {
}

// Separate cached content based on the visitor's device type.
type PageruleNewParamsActionsCacheByDeviceTypeID string

const (
	PageruleNewParamsActionsCacheByDeviceTypeIDCacheByDeviceType PageruleNewParamsActionsCacheByDeviceTypeID = "cache_by_device_type"
)

func (r PageruleNewParamsActionsCacheByDeviceTypeID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsCacheByDeviceTypeIDCacheByDeviceType:
		return true
	}
	return false
}

// The status of Cache By Device Type.
type PageruleNewParamsActionsCacheByDeviceTypeValue string

const (
	PageruleNewParamsActionsCacheByDeviceTypeValueOn  PageruleNewParamsActionsCacheByDeviceTypeValue = "on"
	PageruleNewParamsActionsCacheByDeviceTypeValueOff PageruleNewParamsActionsCacheByDeviceTypeValue = "off"
)

func (r PageruleNewParamsActionsCacheByDeviceTypeValue) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsCacheByDeviceTypeValueOn, PageruleNewParamsActionsCacheByDeviceTypeValueOff:
		return true
	}
	return false
}

type PageruleNewParamsActionsCacheDeceptionArmor struct {
	// Protect from web cache deception attacks while still allowing static assets to
	// be cached. This setting verifies that the URL's extension matches the returned
	// `Content-Type`.
	ID param.Field[PageruleNewParamsActionsCacheDeceptionArmorID] `json:"id"`
	// The status of Cache Deception Armor.
	Value param.Field[PageruleNewParamsActionsCacheDeceptionArmorValue] `json:"value"`
}

func (r PageruleNewParamsActionsCacheDeceptionArmor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsCacheDeceptionArmor) ImplementsPagerulesPageruleNewParamsActionUnion() {
}

// Protect from web cache deception attacks while still allowing static assets to
// be cached. This setting verifies that the URL's extension matches the returned
// `Content-Type`.
type PageruleNewParamsActionsCacheDeceptionArmorID string

const (
	PageruleNewParamsActionsCacheDeceptionArmorIDCacheDeceptionArmor PageruleNewParamsActionsCacheDeceptionArmorID = "cache_deception_armor"
)

func (r PageruleNewParamsActionsCacheDeceptionArmorID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsCacheDeceptionArmorIDCacheDeceptionArmor:
		return true
	}
	return false
}

// The status of Cache Deception Armor.
type PageruleNewParamsActionsCacheDeceptionArmorValue string

const (
	PageruleNewParamsActionsCacheDeceptionArmorValueOn  PageruleNewParamsActionsCacheDeceptionArmorValue = "on"
	PageruleNewParamsActionsCacheDeceptionArmorValueOff PageruleNewParamsActionsCacheDeceptionArmorValue = "off"
)

func (r PageruleNewParamsActionsCacheDeceptionArmorValue) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsCacheDeceptionArmorValueOn, PageruleNewParamsActionsCacheDeceptionArmorValueOff:
		return true
	}
	return false
}

type PageruleNewParamsActionsCacheKey struct {
	// Control specifically what variables to include when deciding which resources to
	// cache. This allows customers to determine what to cache based on something other
	// than just the URL.
	ID    param.Field[PageruleNewParamsActionsCacheKeyID]    `json:"id"`
	Value param.Field[PageruleNewParamsActionsCacheKeyValue] `json:"value"`
}

func (r PageruleNewParamsActionsCacheKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsCacheKey) ImplementsPagerulesPageruleNewParamsActionUnion() {}

// Control specifically what variables to include when deciding which resources to
// cache. This allows customers to determine what to cache based on something other
// than just the URL.
type PageruleNewParamsActionsCacheKeyID string

const (
	PageruleNewParamsActionsCacheKeyIDCacheKey PageruleNewParamsActionsCacheKeyID = "cache_key"
)

func (r PageruleNewParamsActionsCacheKeyID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsCacheKeyIDCacheKey:
		return true
	}
	return false
}

type PageruleNewParamsActionsCacheKeyValue struct {
	// Controls which cookies appear in the Cache Key.
	Cookie param.Field[PageruleNewParamsActionsCacheKeyValueCookie] `json:"cookie"`
	// Controls which headers go into the Cache Key. Exactly one of `include` or
	// `exclude` is expected.
	Header param.Field[PageruleNewParamsActionsCacheKeyValueHeader] `json:"header"`
	// Determines which host header to include in the Cache Key.
	Host param.Field[PageruleNewParamsActionsCacheKeyValueHost] `json:"host"`
	// Controls which URL query string parameters go into the Cache Key. Exactly one of
	// `include` or `exclude` is expected.
	QueryString param.Field[PageruleNewParamsActionsCacheKeyValueQueryString] `json:"query_string"`
	// Feature fields to add features about the end-user (client) into the Cache Key.
	User param.Field[PageruleNewParamsActionsCacheKeyValueUser] `json:"user"`
}

func (r PageruleNewParamsActionsCacheKeyValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which cookies appear in the Cache Key.
type PageruleNewParamsActionsCacheKeyValueCookie struct {
	// A list of cookies to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of cookies to include.
	Include param.Field[[]string] `json:"include"`
}

func (r PageruleNewParamsActionsCacheKeyValueCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which headers go into the Cache Key. Exactly one of `include` or
// `exclude` is expected.
type PageruleNewParamsActionsCacheKeyValueHeader struct {
	// A list of headers to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of headers to ignore.
	Exclude param.Field[[]string] `json:"exclude"`
	// A list of headers to include.
	Include param.Field[[]string] `json:"include"`
}

func (r PageruleNewParamsActionsCacheKeyValueHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which host header to include in the Cache Key.
type PageruleNewParamsActionsCacheKeyValueHost struct {
	// Whether to include the Host header in the HTTP request sent to the origin.
	Resolved param.Field[bool] `json:"resolved"`
}

func (r PageruleNewParamsActionsCacheKeyValueHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which URL query string parameters go into the Cache Key. Exactly one of
// `include` or `exclude` is expected.
type PageruleNewParamsActionsCacheKeyValueQueryString struct {
	// Ignore all query string parameters.
	Exclude param.Field[PageruleNewParamsActionsCacheKeyValueQueryStringExcludeUnion] `json:"exclude"`
	// Include all query string parameters.
	Include param.Field[PageruleNewParamsActionsCacheKeyValueQueryStringIncludeUnion] `json:"include"`
}

func (r PageruleNewParamsActionsCacheKeyValueQueryString) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Ignore all query string parameters.
//
// Satisfied by
// [pagerules.PageruleNewParamsActionsCacheKeyValueQueryStringExcludeString],
// [pagerules.PageruleNewParamsActionsCacheKeyValueQueryStringExcludeArray].
type PageruleNewParamsActionsCacheKeyValueQueryStringExcludeUnion interface {
	implementsPagerulesPageruleNewParamsActionsCacheKeyValueQueryStringExcludeUnion()
}

// Ignore all query string parameters.
type PageruleNewParamsActionsCacheKeyValueQueryStringExcludeString string

const (
	PageruleNewParamsActionsCacheKeyValueQueryStringExcludeStringStar PageruleNewParamsActionsCacheKeyValueQueryStringExcludeString = "*"
)

func (r PageruleNewParamsActionsCacheKeyValueQueryStringExcludeString) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsCacheKeyValueQueryStringExcludeStringStar:
		return true
	}
	return false
}

func (r PageruleNewParamsActionsCacheKeyValueQueryStringExcludeString) implementsPagerulesPageruleNewParamsActionsCacheKeyValueQueryStringExcludeUnion() {
}

type PageruleNewParamsActionsCacheKeyValueQueryStringExcludeArray []string

func (r PageruleNewParamsActionsCacheKeyValueQueryStringExcludeArray) implementsPagerulesPageruleNewParamsActionsCacheKeyValueQueryStringExcludeUnion() {
}

// Include all query string parameters.
//
// Satisfied by
// [pagerules.PageruleNewParamsActionsCacheKeyValueQueryStringIncludeString],
// [pagerules.PageruleNewParamsActionsCacheKeyValueQueryStringIncludeArray].
type PageruleNewParamsActionsCacheKeyValueQueryStringIncludeUnion interface {
	implementsPagerulesPageruleNewParamsActionsCacheKeyValueQueryStringIncludeUnion()
}

// Include all query string parameters.
type PageruleNewParamsActionsCacheKeyValueQueryStringIncludeString string

const (
	PageruleNewParamsActionsCacheKeyValueQueryStringIncludeStringStar PageruleNewParamsActionsCacheKeyValueQueryStringIncludeString = "*"
)

func (r PageruleNewParamsActionsCacheKeyValueQueryStringIncludeString) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsCacheKeyValueQueryStringIncludeStringStar:
		return true
	}
	return false
}

func (r PageruleNewParamsActionsCacheKeyValueQueryStringIncludeString) implementsPagerulesPageruleNewParamsActionsCacheKeyValueQueryStringIncludeUnion() {
}

type PageruleNewParamsActionsCacheKeyValueQueryStringIncludeArray []string

func (r PageruleNewParamsActionsCacheKeyValueQueryStringIncludeArray) implementsPagerulesPageruleNewParamsActionsCacheKeyValueQueryStringIncludeUnion() {
}

// Feature fields to add features about the end-user (client) into the Cache Key.
type PageruleNewParamsActionsCacheKeyValueUser struct {
	// Classifies a request as `mobile`, `desktop`, or `tablet` based on the User
	// Agent.
	DeviceType param.Field[bool] `json:"device_type"`
	// Includes the client's country, derived from the IP address.
	Geo param.Field[bool] `json:"geo"`
	// Includes the first language code contained in the `Accept-Language` header sent
	// by the client.
	Lang param.Field[bool] `json:"lang"`
}

func (r PageruleNewParamsActionsCacheKeyValueUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleNewParamsActionsCacheOnCookie struct {
	// Apply the Cache Everything option (Cache Level setting) based on a regular
	// expression match against a cookie name.
	ID param.Field[PageruleNewParamsActionsCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request.
	Value param.Field[string] `json:"value"`
}

func (r PageruleNewParamsActionsCacheOnCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsCacheOnCookie) ImplementsPagerulesPageruleNewParamsActionUnion() {}

// Apply the Cache Everything option (Cache Level setting) based on a regular
// expression match against a cookie name.
type PageruleNewParamsActionsCacheOnCookieID string

const (
	PageruleNewParamsActionsCacheOnCookieIDCacheOnCookie PageruleNewParamsActionsCacheOnCookieID = "cache_on_cookie"
)

func (r PageruleNewParamsActionsCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsCacheOnCookieIDCacheOnCookie:
		return true
	}
	return false
}

type PageruleNewParamsActionsDisableApps struct {
	// Turn off all active
	// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
	// (deprecated).
	ID param.Field[PageruleNewParamsActionsDisableAppsID] `json:"id"`
}

func (r PageruleNewParamsActionsDisableApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsDisableApps) ImplementsPagerulesPageruleNewParamsActionUnion() {}

// Turn off all active
// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
// (deprecated).
type PageruleNewParamsActionsDisableAppsID string

const (
	PageruleNewParamsActionsDisableAppsIDDisableApps PageruleNewParamsActionsDisableAppsID = "disable_apps"
)

func (r PageruleNewParamsActionsDisableAppsID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsDisableAppsIDDisableApps:
		return true
	}
	return false
}

type PageruleNewParamsActionsDisablePerformance struct {
	// Turn off
	// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
	// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
	// and [Polish](https://developers.cloudflare.com/images/polish/).
	ID param.Field[PageruleNewParamsActionsDisablePerformanceID] `json:"id"`
}

func (r PageruleNewParamsActionsDisablePerformance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsDisablePerformance) ImplementsPagerulesPageruleNewParamsActionUnion() {
}

// Turn off
// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
// and [Polish](https://developers.cloudflare.com/images/polish/).
type PageruleNewParamsActionsDisablePerformanceID string

const (
	PageruleNewParamsActionsDisablePerformanceIDDisablePerformance PageruleNewParamsActionsDisablePerformanceID = "disable_performance"
)

func (r PageruleNewParamsActionsDisablePerformanceID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsDisablePerformanceIDDisablePerformance:
		return true
	}
	return false
}

type PageruleNewParamsActionsDisableSecurity struct {
	// Turn off
	// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
	// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
	// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
	// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
	// and
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	ID param.Field[PageruleNewParamsActionsDisableSecurityID] `json:"id"`
}

func (r PageruleNewParamsActionsDisableSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsDisableSecurity) ImplementsPagerulesPageruleNewParamsActionUnion() {}

// Turn off
// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
// and
// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
type PageruleNewParamsActionsDisableSecurityID string

const (
	PageruleNewParamsActionsDisableSecurityIDDisableSecurity PageruleNewParamsActionsDisableSecurityID = "disable_security"
)

func (r PageruleNewParamsActionsDisableSecurityID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsDisableSecurityIDDisableSecurity:
		return true
	}
	return false
}

type PageruleNewParamsActionsDisableZaraz struct {
	// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
	ID param.Field[PageruleNewParamsActionsDisableZarazID] `json:"id"`
}

func (r PageruleNewParamsActionsDisableZaraz) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsDisableZaraz) ImplementsPagerulesPageruleNewParamsActionUnion() {}

// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
type PageruleNewParamsActionsDisableZarazID string

const (
	PageruleNewParamsActionsDisableZarazIDDisableZaraz PageruleNewParamsActionsDisableZarazID = "disable_zaraz"
)

func (r PageruleNewParamsActionsDisableZarazID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsDisableZarazIDDisableZaraz:
		return true
	}
	return false
}

type PageruleNewParamsActionsEdgeCacheTTL struct {
	// Specify how long to cache a resource in the Cloudflare global network. _Edge
	// Cache TTL_ is not visible in response headers.
	ID    param.Field[PageruleNewParamsActionsEdgeCacheTTLID] `json:"id"`
	Value param.Field[int64]                                  `json:"value"`
}

func (r PageruleNewParamsActionsEdgeCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsEdgeCacheTTL) ImplementsPagerulesPageruleNewParamsActionUnion() {}

// Specify how long to cache a resource in the Cloudflare global network. _Edge
// Cache TTL_ is not visible in response headers.
type PageruleNewParamsActionsEdgeCacheTTLID string

const (
	PageruleNewParamsActionsEdgeCacheTTLIDEdgeCacheTTL PageruleNewParamsActionsEdgeCacheTTLID = "edge_cache_ttl"
)

func (r PageruleNewParamsActionsEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

type PageruleNewParamsActionsExplicitCacheControl struct {
	// Origin Cache Control is enabled by default for Free, Pro, and Business domains
	// and disabled by default for Enterprise domains.
	ID param.Field[PageruleNewParamsActionsExplicitCacheControlID] `json:"id"`
	// The status of Origin Cache Control.
	Value param.Field[PageruleNewParamsActionsExplicitCacheControlValue] `json:"value"`
}

func (r PageruleNewParamsActionsExplicitCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsExplicitCacheControl) ImplementsPagerulesPageruleNewParamsActionUnion() {
}

// Origin Cache Control is enabled by default for Free, Pro, and Business domains
// and disabled by default for Enterprise domains.
type PageruleNewParamsActionsExplicitCacheControlID string

const (
	PageruleNewParamsActionsExplicitCacheControlIDExplicitCacheControl PageruleNewParamsActionsExplicitCacheControlID = "explicit_cache_control"
)

func (r PageruleNewParamsActionsExplicitCacheControlID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsExplicitCacheControlIDExplicitCacheControl:
		return true
	}
	return false
}

// The status of Origin Cache Control.
type PageruleNewParamsActionsExplicitCacheControlValue string

const (
	PageruleNewParamsActionsExplicitCacheControlValueOn  PageruleNewParamsActionsExplicitCacheControlValue = "on"
	PageruleNewParamsActionsExplicitCacheControlValueOff PageruleNewParamsActionsExplicitCacheControlValue = "off"
)

func (r PageruleNewParamsActionsExplicitCacheControlValue) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsExplicitCacheControlValueOn, PageruleNewParamsActionsExplicitCacheControlValueOff:
		return true
	}
	return false
}

type PageruleNewParamsActionsForwardingURL struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    param.Field[PageruleNewParamsActionsForwardingURLID]    `json:"id"`
	Value param.Field[PageruleNewParamsActionsForwardingURLValue] `json:"value"`
}

func (r PageruleNewParamsActionsForwardingURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsForwardingURL) ImplementsPagerulesPageruleNewParamsActionUnion() {}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageruleNewParamsActionsForwardingURLID string

const (
	PageruleNewParamsActionsForwardingURLIDForwardingURL PageruleNewParamsActionsForwardingURLID = "forwarding_url"
)

func (r PageruleNewParamsActionsForwardingURLID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsForwardingURLIDForwardingURL:
		return true
	}
	return false
}

type PageruleNewParamsActionsForwardingURLValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode param.Field[PageruleNewParamsActionsForwardingURLValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageruleNewParamsActionsForwardingURLValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageruleNewParamsActionsForwardingURLValueStatusCode int64

const (
	PageruleNewParamsActionsForwardingURLValueStatusCode301 PageruleNewParamsActionsForwardingURLValueStatusCode = 301
	PageruleNewParamsActionsForwardingURLValueStatusCode302 PageruleNewParamsActionsForwardingURLValueStatusCode = 302
)

func (r PageruleNewParamsActionsForwardingURLValueStatusCode) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsForwardingURLValueStatusCode301, PageruleNewParamsActionsForwardingURLValueStatusCode302:
		return true
	}
	return false
}

type PageruleNewParamsActionsHostHeaderOverride struct {
	// Apply a specific host header.
	ID param.Field[PageruleNewParamsActionsHostHeaderOverrideID] `json:"id"`
	// The hostname to use in the `Host` header
	Value param.Field[string] `json:"value"`
}

func (r PageruleNewParamsActionsHostHeaderOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsHostHeaderOverride) ImplementsPagerulesPageruleNewParamsActionUnion() {
}

// Apply a specific host header.
type PageruleNewParamsActionsHostHeaderOverrideID string

const (
	PageruleNewParamsActionsHostHeaderOverrideIDHostHeaderOverride PageruleNewParamsActionsHostHeaderOverrideID = "host_header_override"
)

func (r PageruleNewParamsActionsHostHeaderOverrideID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsHostHeaderOverrideIDHostHeaderOverride:
		return true
	}
	return false
}

type PageruleNewParamsActionsResolveOverride struct {
	// Change the origin address to the value specified in this setting.
	ID param.Field[PageruleNewParamsActionsResolveOverrideID] `json:"id"`
	// The origin address you want to override with.
	Value param.Field[string] `json:"value"`
}

func (r PageruleNewParamsActionsResolveOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsResolveOverride) ImplementsPagerulesPageruleNewParamsActionUnion() {}

// Change the origin address to the value specified in this setting.
type PageruleNewParamsActionsResolveOverrideID string

const (
	PageruleNewParamsActionsResolveOverrideIDResolveOverride PageruleNewParamsActionsResolveOverrideID = "resolve_override"
)

func (r PageruleNewParamsActionsResolveOverrideID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsResolveOverrideIDResolveOverride:
		return true
	}
	return false
}

type PageruleNewParamsActionsRespectStrongEtag struct {
	// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
	// the origin server.
	ID param.Field[PageruleNewParamsActionsRespectStrongEtagID] `json:"id"`
	// The status of Respect Strong ETags
	Value param.Field[PageruleNewParamsActionsRespectStrongEtagValue] `json:"value"`
}

func (r PageruleNewParamsActionsRespectStrongEtag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleNewParamsActionsRespectStrongEtag) ImplementsPagerulesPageruleNewParamsActionUnion() {
}

// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
// the origin server.
type PageruleNewParamsActionsRespectStrongEtagID string

const (
	PageruleNewParamsActionsRespectStrongEtagIDRespectStrongEtag PageruleNewParamsActionsRespectStrongEtagID = "respect_strong_etag"
)

func (r PageruleNewParamsActionsRespectStrongEtagID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsRespectStrongEtagIDRespectStrongEtag:
		return true
	}
	return false
}

// The status of Respect Strong ETags
type PageruleNewParamsActionsRespectStrongEtagValue string

const (
	PageruleNewParamsActionsRespectStrongEtagValueOn  PageruleNewParamsActionsRespectStrongEtagValue = "on"
	PageruleNewParamsActionsRespectStrongEtagValueOff PageruleNewParamsActionsRespectStrongEtagValue = "off"
)

func (r PageruleNewParamsActionsRespectStrongEtagValue) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsRespectStrongEtagValueOn, PageruleNewParamsActionsRespectStrongEtagValueOff:
		return true
	}
	return false
}

// If enabled, any ` http://“ URL is converted to  `https://` through a 301
// redirect.
type PageruleNewParamsActionsID string

const (
	PageruleNewParamsActionsIDAlwaysUseHTTPS          PageruleNewParamsActionsID = "always_use_https"
	PageruleNewParamsActionsIDAutomaticHTTPSRewrites  PageruleNewParamsActionsID = "automatic_https_rewrites"
	PageruleNewParamsActionsIDBrowserCacheTTL         PageruleNewParamsActionsID = "browser_cache_ttl"
	PageruleNewParamsActionsIDBrowserCheck            PageruleNewParamsActionsID = "browser_check"
	PageruleNewParamsActionsIDBypassCacheOnCookie     PageruleNewParamsActionsID = "bypass_cache_on_cookie"
	PageruleNewParamsActionsIDCacheByDeviceType       PageruleNewParamsActionsID = "cache_by_device_type"
	PageruleNewParamsActionsIDCacheDeceptionArmor     PageruleNewParamsActionsID = "cache_deception_armor"
	PageruleNewParamsActionsIDCacheKey                PageruleNewParamsActionsID = "cache_key"
	PageruleNewParamsActionsIDCacheLevel              PageruleNewParamsActionsID = "cache_level"
	PageruleNewParamsActionsIDCacheOnCookie           PageruleNewParamsActionsID = "cache_on_cookie"
	PageruleNewParamsActionsIDDisableApps             PageruleNewParamsActionsID = "disable_apps"
	PageruleNewParamsActionsIDDisablePerformance      PageruleNewParamsActionsID = "disable_performance"
	PageruleNewParamsActionsIDDisableSecurity         PageruleNewParamsActionsID = "disable_security"
	PageruleNewParamsActionsIDDisableZaraz            PageruleNewParamsActionsID = "disable_zaraz"
	PageruleNewParamsActionsIDEdgeCacheTTL            PageruleNewParamsActionsID = "edge_cache_ttl"
	PageruleNewParamsActionsIDEmailObfuscation        PageruleNewParamsActionsID = "email_obfuscation"
	PageruleNewParamsActionsIDExplicitCacheControl    PageruleNewParamsActionsID = "explicit_cache_control"
	PageruleNewParamsActionsIDForwardingURL           PageruleNewParamsActionsID = "forwarding_url"
	PageruleNewParamsActionsIDHostHeaderOverride      PageruleNewParamsActionsID = "host_header_override"
	PageruleNewParamsActionsIDIPGeolocation           PageruleNewParamsActionsID = "ip_geolocation"
	PageruleNewParamsActionsIDMirage                  PageruleNewParamsActionsID = "mirage"
	PageruleNewParamsActionsIDOpportunisticEncryption PageruleNewParamsActionsID = "opportunistic_encryption"
	PageruleNewParamsActionsIDOriginErrorPagePassThru PageruleNewParamsActionsID = "origin_error_page_pass_thru"
	PageruleNewParamsActionsIDPolish                  PageruleNewParamsActionsID = "polish"
	PageruleNewParamsActionsIDResolveOverride         PageruleNewParamsActionsID = "resolve_override"
	PageruleNewParamsActionsIDRespectStrongEtag       PageruleNewParamsActionsID = "respect_strong_etag"
	PageruleNewParamsActionsIDResponseBuffering       PageruleNewParamsActionsID = "response_buffering"
	PageruleNewParamsActionsIDRocketLoader            PageruleNewParamsActionsID = "rocket_loader"
	PageruleNewParamsActionsIDSecurityLevel           PageruleNewParamsActionsID = "security_level"
	PageruleNewParamsActionsIDSortQueryStringForCache PageruleNewParamsActionsID = "sort_query_string_for_cache"
	PageruleNewParamsActionsIDSSL                     PageruleNewParamsActionsID = "ssl"
	PageruleNewParamsActionsIDTrueClientIPHeader      PageruleNewParamsActionsID = "true_client_ip_header"
	PageruleNewParamsActionsIDWAF                     PageruleNewParamsActionsID = "waf"
)

func (r PageruleNewParamsActionsID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsIDAlwaysUseHTTPS, PageruleNewParamsActionsIDAutomaticHTTPSRewrites, PageruleNewParamsActionsIDBrowserCacheTTL, PageruleNewParamsActionsIDBrowserCheck, PageruleNewParamsActionsIDBypassCacheOnCookie, PageruleNewParamsActionsIDCacheByDeviceType, PageruleNewParamsActionsIDCacheDeceptionArmor, PageruleNewParamsActionsIDCacheKey, PageruleNewParamsActionsIDCacheLevel, PageruleNewParamsActionsIDCacheOnCookie, PageruleNewParamsActionsIDDisableApps, PageruleNewParamsActionsIDDisablePerformance, PageruleNewParamsActionsIDDisableSecurity, PageruleNewParamsActionsIDDisableZaraz, PageruleNewParamsActionsIDEdgeCacheTTL, PageruleNewParamsActionsIDEmailObfuscation, PageruleNewParamsActionsIDExplicitCacheControl, PageruleNewParamsActionsIDForwardingURL, PageruleNewParamsActionsIDHostHeaderOverride, PageruleNewParamsActionsIDIPGeolocation, PageruleNewParamsActionsIDMirage, PageruleNewParamsActionsIDOpportunisticEncryption, PageruleNewParamsActionsIDOriginErrorPagePassThru, PageruleNewParamsActionsIDPolish, PageruleNewParamsActionsIDResolveOverride, PageruleNewParamsActionsIDRespectStrongEtag, PageruleNewParamsActionsIDResponseBuffering, PageruleNewParamsActionsIDRocketLoader, PageruleNewParamsActionsIDSecurityLevel, PageruleNewParamsActionsIDSortQueryStringForCache, PageruleNewParamsActionsIDSSL, PageruleNewParamsActionsIDTrueClientIPHeader, PageruleNewParamsActionsIDWAF:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleNewParamsStatus string

const (
	PageruleNewParamsStatusActive   PageruleNewParamsStatus = "active"
	PageruleNewParamsStatusDisabled PageruleNewParamsStatus = "disabled"
)

func (r PageruleNewParamsStatus) IsKnown() bool {
	switch r {
	case PageruleNewParamsStatusActive, PageruleNewParamsStatusDisabled:
		return true
	}
	return false
}

type PageruleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageruleNewResponseEnvelopeSuccess `json:"success,required"`
	Result  PageRule                           `json:"result"`
	JSON    pageruleNewResponseEnvelopeJSON    `json:"-"`
}

// pageruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleNewResponseEnvelope]
type pageruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleNewResponseEnvelopeSuccess bool

const (
	PageruleNewResponseEnvelopeSuccessTrue PageruleNewResponseEnvelopeSuccess = true
)

func (r PageruleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]PageruleUpdateParamsActionUnion] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargetParam] `json:"targets,required"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageruleUpdateParamsStatus] `json:"status"`
}

func (r PageruleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleUpdateParamsAction struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID    param.Field[PageruleUpdateParamsActionsID] `json:"id"`
	Value param.Field[interface{}]                   `json:"value"`
}

func (r PageruleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsAction) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

// Satisfied by [zones.AlwaysUseHTTPSParam], [zones.AutomaticHTTPSRewritesParam],
// [zones.BrowserCacheTTLParam], [zones.BrowserCheckParam],
// [pagerules.PageruleUpdateParamsActionsBypassCacheOnCookie],
// [pagerules.PageruleUpdateParamsActionsCacheByDeviceType],
// [pagerules.PageruleUpdateParamsActionsCacheDeceptionArmor],
// [pagerules.PageruleUpdateParamsActionsCacheKey], [zones.CacheLevelParam],
// [pagerules.PageruleUpdateParamsActionsCacheOnCookie],
// [pagerules.PageruleUpdateParamsActionsDisableApps],
// [pagerules.PageruleUpdateParamsActionsDisablePerformance],
// [pagerules.PageruleUpdateParamsActionsDisableSecurity],
// [pagerules.PageruleUpdateParamsActionsDisableZaraz],
// [pagerules.PageruleUpdateParamsActionsEdgeCacheTTL],
// [zones.EmailObfuscationParam],
// [pagerules.PageruleUpdateParamsActionsExplicitCacheControl],
// [pagerules.PageruleUpdateParamsActionsForwardingURL],
// [pagerules.PageruleUpdateParamsActionsHostHeaderOverride],
// [zones.IPGeolocationParam], [zones.MirageParam],
// [zones.OpportunisticEncryptionParam], [zones.OriginErrorPagePassThruParam],
// [zones.PolishParam], [pagerules.PageruleUpdateParamsActionsResolveOverride],
// [pagerules.PageruleUpdateParamsActionsRespectStrongEtag],
// [zones.ResponseBufferingParam], [zones.RocketLoaderParam],
// [zones.SecurityLevelParam], [zones.SortQueryStringForCacheParam],
// [zones.SSLParam], [zones.TrueClientIPHeaderParam], [zones.WAFParam],
// [PageruleUpdateParamsAction].
type PageruleUpdateParamsActionUnion interface {
	ImplementsPagerulesPageruleUpdateParamsActionUnion()
}

type PageruleUpdateParamsActionsBypassCacheOnCookie struct {
	// Bypass cache and fetch resources from the origin server if a regular expression
	// matches against a cookie name present in the request.
	ID param.Field[PageruleUpdateParamsActionsBypassCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request. Refer to
	// [Bypass Cache on Cookie setting](https://developers.cloudflare.com/rules/page-rules/reference/additional-reference/#bypass-cache-on-cookie-setting)
	// to learn about limited regular expression support.
	Value param.Field[string] `json:"value"`
}

func (r PageruleUpdateParamsActionsBypassCacheOnCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsBypassCacheOnCookie) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Bypass cache and fetch resources from the origin server if a regular expression
// matches against a cookie name present in the request.
type PageruleUpdateParamsActionsBypassCacheOnCookieID string

const (
	PageruleUpdateParamsActionsBypassCacheOnCookieIDBypassCacheOnCookie PageruleUpdateParamsActionsBypassCacheOnCookieID = "bypass_cache_on_cookie"
)

func (r PageruleUpdateParamsActionsBypassCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsBypassCacheOnCookieIDBypassCacheOnCookie:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsCacheByDeviceType struct {
	// Separate cached content based on the visitor's device type.
	ID param.Field[PageruleUpdateParamsActionsCacheByDeviceTypeID] `json:"id"`
	// The status of Cache By Device Type.
	Value param.Field[PageruleUpdateParamsActionsCacheByDeviceTypeValue] `json:"value"`
}

func (r PageruleUpdateParamsActionsCacheByDeviceType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsCacheByDeviceType) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Separate cached content based on the visitor's device type.
type PageruleUpdateParamsActionsCacheByDeviceTypeID string

const (
	PageruleUpdateParamsActionsCacheByDeviceTypeIDCacheByDeviceType PageruleUpdateParamsActionsCacheByDeviceTypeID = "cache_by_device_type"
)

func (r PageruleUpdateParamsActionsCacheByDeviceTypeID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsCacheByDeviceTypeIDCacheByDeviceType:
		return true
	}
	return false
}

// The status of Cache By Device Type.
type PageruleUpdateParamsActionsCacheByDeviceTypeValue string

const (
	PageruleUpdateParamsActionsCacheByDeviceTypeValueOn  PageruleUpdateParamsActionsCacheByDeviceTypeValue = "on"
	PageruleUpdateParamsActionsCacheByDeviceTypeValueOff PageruleUpdateParamsActionsCacheByDeviceTypeValue = "off"
)

func (r PageruleUpdateParamsActionsCacheByDeviceTypeValue) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsCacheByDeviceTypeValueOn, PageruleUpdateParamsActionsCacheByDeviceTypeValueOff:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsCacheDeceptionArmor struct {
	// Protect from web cache deception attacks while still allowing static assets to
	// be cached. This setting verifies that the URL's extension matches the returned
	// `Content-Type`.
	ID param.Field[PageruleUpdateParamsActionsCacheDeceptionArmorID] `json:"id"`
	// The status of Cache Deception Armor.
	Value param.Field[PageruleUpdateParamsActionsCacheDeceptionArmorValue] `json:"value"`
}

func (r PageruleUpdateParamsActionsCacheDeceptionArmor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsCacheDeceptionArmor) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Protect from web cache deception attacks while still allowing static assets to
// be cached. This setting verifies that the URL's extension matches the returned
// `Content-Type`.
type PageruleUpdateParamsActionsCacheDeceptionArmorID string

const (
	PageruleUpdateParamsActionsCacheDeceptionArmorIDCacheDeceptionArmor PageruleUpdateParamsActionsCacheDeceptionArmorID = "cache_deception_armor"
)

func (r PageruleUpdateParamsActionsCacheDeceptionArmorID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsCacheDeceptionArmorIDCacheDeceptionArmor:
		return true
	}
	return false
}

// The status of Cache Deception Armor.
type PageruleUpdateParamsActionsCacheDeceptionArmorValue string

const (
	PageruleUpdateParamsActionsCacheDeceptionArmorValueOn  PageruleUpdateParamsActionsCacheDeceptionArmorValue = "on"
	PageruleUpdateParamsActionsCacheDeceptionArmorValueOff PageruleUpdateParamsActionsCacheDeceptionArmorValue = "off"
)

func (r PageruleUpdateParamsActionsCacheDeceptionArmorValue) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsCacheDeceptionArmorValueOn, PageruleUpdateParamsActionsCacheDeceptionArmorValueOff:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsCacheKey struct {
	// Control specifically what variables to include when deciding which resources to
	// cache. This allows customers to determine what to cache based on something other
	// than just the URL.
	ID    param.Field[PageruleUpdateParamsActionsCacheKeyID]    `json:"id"`
	Value param.Field[PageruleUpdateParamsActionsCacheKeyValue] `json:"value"`
}

func (r PageruleUpdateParamsActionsCacheKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsCacheKey) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

// Control specifically what variables to include when deciding which resources to
// cache. This allows customers to determine what to cache based on something other
// than just the URL.
type PageruleUpdateParamsActionsCacheKeyID string

const (
	PageruleUpdateParamsActionsCacheKeyIDCacheKey PageruleUpdateParamsActionsCacheKeyID = "cache_key"
)

func (r PageruleUpdateParamsActionsCacheKeyID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsCacheKeyIDCacheKey:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsCacheKeyValue struct {
	// Controls which cookies appear in the Cache Key.
	Cookie param.Field[PageruleUpdateParamsActionsCacheKeyValueCookie] `json:"cookie"`
	// Controls which headers go into the Cache Key. Exactly one of `include` or
	// `exclude` is expected.
	Header param.Field[PageruleUpdateParamsActionsCacheKeyValueHeader] `json:"header"`
	// Determines which host header to include in the Cache Key.
	Host param.Field[PageruleUpdateParamsActionsCacheKeyValueHost] `json:"host"`
	// Controls which URL query string parameters go into the Cache Key. Exactly one of
	// `include` or `exclude` is expected.
	QueryString param.Field[PageruleUpdateParamsActionsCacheKeyValueQueryString] `json:"query_string"`
	// Feature fields to add features about the end-user (client) into the Cache Key.
	User param.Field[PageruleUpdateParamsActionsCacheKeyValueUser] `json:"user"`
}

func (r PageruleUpdateParamsActionsCacheKeyValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which cookies appear in the Cache Key.
type PageruleUpdateParamsActionsCacheKeyValueCookie struct {
	// A list of cookies to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of cookies to include.
	Include param.Field[[]string] `json:"include"`
}

func (r PageruleUpdateParamsActionsCacheKeyValueCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which headers go into the Cache Key. Exactly one of `include` or
// `exclude` is expected.
type PageruleUpdateParamsActionsCacheKeyValueHeader struct {
	// A list of headers to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of headers to ignore.
	Exclude param.Field[[]string] `json:"exclude"`
	// A list of headers to include.
	Include param.Field[[]string] `json:"include"`
}

func (r PageruleUpdateParamsActionsCacheKeyValueHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which host header to include in the Cache Key.
type PageruleUpdateParamsActionsCacheKeyValueHost struct {
	// Whether to include the Host header in the HTTP request sent to the origin.
	Resolved param.Field[bool] `json:"resolved"`
}

func (r PageruleUpdateParamsActionsCacheKeyValueHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which URL query string parameters go into the Cache Key. Exactly one of
// `include` or `exclude` is expected.
type PageruleUpdateParamsActionsCacheKeyValueQueryString struct {
	// Ignore all query string parameters.
	Exclude param.Field[PageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeUnion] `json:"exclude"`
	// Include all query string parameters.
	Include param.Field[PageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeUnion] `json:"include"`
}

func (r PageruleUpdateParamsActionsCacheKeyValueQueryString) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Ignore all query string parameters.
//
// Satisfied by
// [pagerules.PageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeString],
// [pagerules.PageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeArray].
type PageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeUnion interface {
	implementsPagerulesPageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeUnion()
}

// Ignore all query string parameters.
type PageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeString string

const (
	PageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeStringStar PageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeString = "*"
)

func (r PageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeString) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeStringStar:
		return true
	}
	return false
}

func (r PageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeString) implementsPagerulesPageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeUnion() {
}

type PageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeArray []string

func (r PageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeArray) implementsPagerulesPageruleUpdateParamsActionsCacheKeyValueQueryStringExcludeUnion() {
}

// Include all query string parameters.
//
// Satisfied by
// [pagerules.PageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeString],
// [pagerules.PageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeArray].
type PageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeUnion interface {
	implementsPagerulesPageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeUnion()
}

// Include all query string parameters.
type PageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeString string

const (
	PageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeStringStar PageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeString = "*"
)

func (r PageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeString) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeStringStar:
		return true
	}
	return false
}

func (r PageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeString) implementsPagerulesPageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeUnion() {
}

type PageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeArray []string

func (r PageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeArray) implementsPagerulesPageruleUpdateParamsActionsCacheKeyValueQueryStringIncludeUnion() {
}

// Feature fields to add features about the end-user (client) into the Cache Key.
type PageruleUpdateParamsActionsCacheKeyValueUser struct {
	// Classifies a request as `mobile`, `desktop`, or `tablet` based on the User
	// Agent.
	DeviceType param.Field[bool] `json:"device_type"`
	// Includes the client's country, derived from the IP address.
	Geo param.Field[bool] `json:"geo"`
	// Includes the first language code contained in the `Accept-Language` header sent
	// by the client.
	Lang param.Field[bool] `json:"lang"`
}

func (r PageruleUpdateParamsActionsCacheKeyValueUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleUpdateParamsActionsCacheOnCookie struct {
	// Apply the Cache Everything option (Cache Level setting) based on a regular
	// expression match against a cookie name.
	ID param.Field[PageruleUpdateParamsActionsCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request.
	Value param.Field[string] `json:"value"`
}

func (r PageruleUpdateParamsActionsCacheOnCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsCacheOnCookie) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Apply the Cache Everything option (Cache Level setting) based on a regular
// expression match against a cookie name.
type PageruleUpdateParamsActionsCacheOnCookieID string

const (
	PageruleUpdateParamsActionsCacheOnCookieIDCacheOnCookie PageruleUpdateParamsActionsCacheOnCookieID = "cache_on_cookie"
)

func (r PageruleUpdateParamsActionsCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsCacheOnCookieIDCacheOnCookie:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsDisableApps struct {
	// Turn off all active
	// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
	// (deprecated).
	ID param.Field[PageruleUpdateParamsActionsDisableAppsID] `json:"id"`
}

func (r PageruleUpdateParamsActionsDisableApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsDisableApps) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Turn off all active
// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
// (deprecated).
type PageruleUpdateParamsActionsDisableAppsID string

const (
	PageruleUpdateParamsActionsDisableAppsIDDisableApps PageruleUpdateParamsActionsDisableAppsID = "disable_apps"
)

func (r PageruleUpdateParamsActionsDisableAppsID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsDisableAppsIDDisableApps:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsDisablePerformance struct {
	// Turn off
	// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
	// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
	// and [Polish](https://developers.cloudflare.com/images/polish/).
	ID param.Field[PageruleUpdateParamsActionsDisablePerformanceID] `json:"id"`
}

func (r PageruleUpdateParamsActionsDisablePerformance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsDisablePerformance) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Turn off
// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
// and [Polish](https://developers.cloudflare.com/images/polish/).
type PageruleUpdateParamsActionsDisablePerformanceID string

const (
	PageruleUpdateParamsActionsDisablePerformanceIDDisablePerformance PageruleUpdateParamsActionsDisablePerformanceID = "disable_performance"
)

func (r PageruleUpdateParamsActionsDisablePerformanceID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsDisablePerformanceIDDisablePerformance:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsDisableSecurity struct {
	// Turn off
	// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
	// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
	// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
	// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
	// and
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	ID param.Field[PageruleUpdateParamsActionsDisableSecurityID] `json:"id"`
}

func (r PageruleUpdateParamsActionsDisableSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsDisableSecurity) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Turn off
// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
// and
// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
type PageruleUpdateParamsActionsDisableSecurityID string

const (
	PageruleUpdateParamsActionsDisableSecurityIDDisableSecurity PageruleUpdateParamsActionsDisableSecurityID = "disable_security"
)

func (r PageruleUpdateParamsActionsDisableSecurityID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsDisableSecurityIDDisableSecurity:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsDisableZaraz struct {
	// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
	ID param.Field[PageruleUpdateParamsActionsDisableZarazID] `json:"id"`
}

func (r PageruleUpdateParamsActionsDisableZaraz) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsDisableZaraz) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
type PageruleUpdateParamsActionsDisableZarazID string

const (
	PageruleUpdateParamsActionsDisableZarazIDDisableZaraz PageruleUpdateParamsActionsDisableZarazID = "disable_zaraz"
)

func (r PageruleUpdateParamsActionsDisableZarazID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsDisableZarazIDDisableZaraz:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsEdgeCacheTTL struct {
	// Specify how long to cache a resource in the Cloudflare global network. _Edge
	// Cache TTL_ is not visible in response headers.
	ID    param.Field[PageruleUpdateParamsActionsEdgeCacheTTLID] `json:"id"`
	Value param.Field[int64]                                     `json:"value"`
}

func (r PageruleUpdateParamsActionsEdgeCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsEdgeCacheTTL) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Specify how long to cache a resource in the Cloudflare global network. _Edge
// Cache TTL_ is not visible in response headers.
type PageruleUpdateParamsActionsEdgeCacheTTLID string

const (
	PageruleUpdateParamsActionsEdgeCacheTTLIDEdgeCacheTTL PageruleUpdateParamsActionsEdgeCacheTTLID = "edge_cache_ttl"
)

func (r PageruleUpdateParamsActionsEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsExplicitCacheControl struct {
	// Origin Cache Control is enabled by default for Free, Pro, and Business domains
	// and disabled by default for Enterprise domains.
	ID param.Field[PageruleUpdateParamsActionsExplicitCacheControlID] `json:"id"`
	// The status of Origin Cache Control.
	Value param.Field[PageruleUpdateParamsActionsExplicitCacheControlValue] `json:"value"`
}

func (r PageruleUpdateParamsActionsExplicitCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsExplicitCacheControl) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Origin Cache Control is enabled by default for Free, Pro, and Business domains
// and disabled by default for Enterprise domains.
type PageruleUpdateParamsActionsExplicitCacheControlID string

const (
	PageruleUpdateParamsActionsExplicitCacheControlIDExplicitCacheControl PageruleUpdateParamsActionsExplicitCacheControlID = "explicit_cache_control"
)

func (r PageruleUpdateParamsActionsExplicitCacheControlID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsExplicitCacheControlIDExplicitCacheControl:
		return true
	}
	return false
}

// The status of Origin Cache Control.
type PageruleUpdateParamsActionsExplicitCacheControlValue string

const (
	PageruleUpdateParamsActionsExplicitCacheControlValueOn  PageruleUpdateParamsActionsExplicitCacheControlValue = "on"
	PageruleUpdateParamsActionsExplicitCacheControlValueOff PageruleUpdateParamsActionsExplicitCacheControlValue = "off"
)

func (r PageruleUpdateParamsActionsExplicitCacheControlValue) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsExplicitCacheControlValueOn, PageruleUpdateParamsActionsExplicitCacheControlValueOff:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsForwardingURL struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    param.Field[PageruleUpdateParamsActionsForwardingURLID]    `json:"id"`
	Value param.Field[PageruleUpdateParamsActionsForwardingURLValue] `json:"value"`
}

func (r PageruleUpdateParamsActionsForwardingURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsForwardingURL) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageruleUpdateParamsActionsForwardingURLID string

const (
	PageruleUpdateParamsActionsForwardingURLIDForwardingURL PageruleUpdateParamsActionsForwardingURLID = "forwarding_url"
)

func (r PageruleUpdateParamsActionsForwardingURLID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsForwardingURLIDForwardingURL:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsForwardingURLValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode param.Field[PageruleUpdateParamsActionsForwardingURLValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageruleUpdateParamsActionsForwardingURLValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageruleUpdateParamsActionsForwardingURLValueStatusCode int64

const (
	PageruleUpdateParamsActionsForwardingURLValueStatusCode301 PageruleUpdateParamsActionsForwardingURLValueStatusCode = 301
	PageruleUpdateParamsActionsForwardingURLValueStatusCode302 PageruleUpdateParamsActionsForwardingURLValueStatusCode = 302
)

func (r PageruleUpdateParamsActionsForwardingURLValueStatusCode) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsForwardingURLValueStatusCode301, PageruleUpdateParamsActionsForwardingURLValueStatusCode302:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsHostHeaderOverride struct {
	// Apply a specific host header.
	ID param.Field[PageruleUpdateParamsActionsHostHeaderOverrideID] `json:"id"`
	// The hostname to use in the `Host` header
	Value param.Field[string] `json:"value"`
}

func (r PageruleUpdateParamsActionsHostHeaderOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsHostHeaderOverride) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Apply a specific host header.
type PageruleUpdateParamsActionsHostHeaderOverrideID string

const (
	PageruleUpdateParamsActionsHostHeaderOverrideIDHostHeaderOverride PageruleUpdateParamsActionsHostHeaderOverrideID = "host_header_override"
)

func (r PageruleUpdateParamsActionsHostHeaderOverrideID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsHostHeaderOverrideIDHostHeaderOverride:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsResolveOverride struct {
	// Change the origin address to the value specified in this setting.
	ID param.Field[PageruleUpdateParamsActionsResolveOverrideID] `json:"id"`
	// The origin address you want to override with.
	Value param.Field[string] `json:"value"`
}

func (r PageruleUpdateParamsActionsResolveOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsResolveOverride) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Change the origin address to the value specified in this setting.
type PageruleUpdateParamsActionsResolveOverrideID string

const (
	PageruleUpdateParamsActionsResolveOverrideIDResolveOverride PageruleUpdateParamsActionsResolveOverrideID = "resolve_override"
)

func (r PageruleUpdateParamsActionsResolveOverrideID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsResolveOverrideIDResolveOverride:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsRespectStrongEtag struct {
	// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
	// the origin server.
	ID param.Field[PageruleUpdateParamsActionsRespectStrongEtagID] `json:"id"`
	// The status of Respect Strong ETags
	Value param.Field[PageruleUpdateParamsActionsRespectStrongEtagValue] `json:"value"`
}

func (r PageruleUpdateParamsActionsRespectStrongEtag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleUpdateParamsActionsRespectStrongEtag) ImplementsPagerulesPageruleUpdateParamsActionUnion() {
}

// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
// the origin server.
type PageruleUpdateParamsActionsRespectStrongEtagID string

const (
	PageruleUpdateParamsActionsRespectStrongEtagIDRespectStrongEtag PageruleUpdateParamsActionsRespectStrongEtagID = "respect_strong_etag"
)

func (r PageruleUpdateParamsActionsRespectStrongEtagID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsRespectStrongEtagIDRespectStrongEtag:
		return true
	}
	return false
}

// The status of Respect Strong ETags
type PageruleUpdateParamsActionsRespectStrongEtagValue string

const (
	PageruleUpdateParamsActionsRespectStrongEtagValueOn  PageruleUpdateParamsActionsRespectStrongEtagValue = "on"
	PageruleUpdateParamsActionsRespectStrongEtagValueOff PageruleUpdateParamsActionsRespectStrongEtagValue = "off"
)

func (r PageruleUpdateParamsActionsRespectStrongEtagValue) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsRespectStrongEtagValueOn, PageruleUpdateParamsActionsRespectStrongEtagValueOff:
		return true
	}
	return false
}

// If enabled, any ` http://“ URL is converted to  `https://` through a 301
// redirect.
type PageruleUpdateParamsActionsID string

const (
	PageruleUpdateParamsActionsIDAlwaysUseHTTPS          PageruleUpdateParamsActionsID = "always_use_https"
	PageruleUpdateParamsActionsIDAutomaticHTTPSRewrites  PageruleUpdateParamsActionsID = "automatic_https_rewrites"
	PageruleUpdateParamsActionsIDBrowserCacheTTL         PageruleUpdateParamsActionsID = "browser_cache_ttl"
	PageruleUpdateParamsActionsIDBrowserCheck            PageruleUpdateParamsActionsID = "browser_check"
	PageruleUpdateParamsActionsIDBypassCacheOnCookie     PageruleUpdateParamsActionsID = "bypass_cache_on_cookie"
	PageruleUpdateParamsActionsIDCacheByDeviceType       PageruleUpdateParamsActionsID = "cache_by_device_type"
	PageruleUpdateParamsActionsIDCacheDeceptionArmor     PageruleUpdateParamsActionsID = "cache_deception_armor"
	PageruleUpdateParamsActionsIDCacheKey                PageruleUpdateParamsActionsID = "cache_key"
	PageruleUpdateParamsActionsIDCacheLevel              PageruleUpdateParamsActionsID = "cache_level"
	PageruleUpdateParamsActionsIDCacheOnCookie           PageruleUpdateParamsActionsID = "cache_on_cookie"
	PageruleUpdateParamsActionsIDDisableApps             PageruleUpdateParamsActionsID = "disable_apps"
	PageruleUpdateParamsActionsIDDisablePerformance      PageruleUpdateParamsActionsID = "disable_performance"
	PageruleUpdateParamsActionsIDDisableSecurity         PageruleUpdateParamsActionsID = "disable_security"
	PageruleUpdateParamsActionsIDDisableZaraz            PageruleUpdateParamsActionsID = "disable_zaraz"
	PageruleUpdateParamsActionsIDEdgeCacheTTL            PageruleUpdateParamsActionsID = "edge_cache_ttl"
	PageruleUpdateParamsActionsIDEmailObfuscation        PageruleUpdateParamsActionsID = "email_obfuscation"
	PageruleUpdateParamsActionsIDExplicitCacheControl    PageruleUpdateParamsActionsID = "explicit_cache_control"
	PageruleUpdateParamsActionsIDForwardingURL           PageruleUpdateParamsActionsID = "forwarding_url"
	PageruleUpdateParamsActionsIDHostHeaderOverride      PageruleUpdateParamsActionsID = "host_header_override"
	PageruleUpdateParamsActionsIDIPGeolocation           PageruleUpdateParamsActionsID = "ip_geolocation"
	PageruleUpdateParamsActionsIDMirage                  PageruleUpdateParamsActionsID = "mirage"
	PageruleUpdateParamsActionsIDOpportunisticEncryption PageruleUpdateParamsActionsID = "opportunistic_encryption"
	PageruleUpdateParamsActionsIDOriginErrorPagePassThru PageruleUpdateParamsActionsID = "origin_error_page_pass_thru"
	PageruleUpdateParamsActionsIDPolish                  PageruleUpdateParamsActionsID = "polish"
	PageruleUpdateParamsActionsIDResolveOverride         PageruleUpdateParamsActionsID = "resolve_override"
	PageruleUpdateParamsActionsIDRespectStrongEtag       PageruleUpdateParamsActionsID = "respect_strong_etag"
	PageruleUpdateParamsActionsIDResponseBuffering       PageruleUpdateParamsActionsID = "response_buffering"
	PageruleUpdateParamsActionsIDRocketLoader            PageruleUpdateParamsActionsID = "rocket_loader"
	PageruleUpdateParamsActionsIDSecurityLevel           PageruleUpdateParamsActionsID = "security_level"
	PageruleUpdateParamsActionsIDSortQueryStringForCache PageruleUpdateParamsActionsID = "sort_query_string_for_cache"
	PageruleUpdateParamsActionsIDSSL                     PageruleUpdateParamsActionsID = "ssl"
	PageruleUpdateParamsActionsIDTrueClientIPHeader      PageruleUpdateParamsActionsID = "true_client_ip_header"
	PageruleUpdateParamsActionsIDWAF                     PageruleUpdateParamsActionsID = "waf"
)

func (r PageruleUpdateParamsActionsID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsIDAlwaysUseHTTPS, PageruleUpdateParamsActionsIDAutomaticHTTPSRewrites, PageruleUpdateParamsActionsIDBrowserCacheTTL, PageruleUpdateParamsActionsIDBrowserCheck, PageruleUpdateParamsActionsIDBypassCacheOnCookie, PageruleUpdateParamsActionsIDCacheByDeviceType, PageruleUpdateParamsActionsIDCacheDeceptionArmor, PageruleUpdateParamsActionsIDCacheKey, PageruleUpdateParamsActionsIDCacheLevel, PageruleUpdateParamsActionsIDCacheOnCookie, PageruleUpdateParamsActionsIDDisableApps, PageruleUpdateParamsActionsIDDisablePerformance, PageruleUpdateParamsActionsIDDisableSecurity, PageruleUpdateParamsActionsIDDisableZaraz, PageruleUpdateParamsActionsIDEdgeCacheTTL, PageruleUpdateParamsActionsIDEmailObfuscation, PageruleUpdateParamsActionsIDExplicitCacheControl, PageruleUpdateParamsActionsIDForwardingURL, PageruleUpdateParamsActionsIDHostHeaderOverride, PageruleUpdateParamsActionsIDIPGeolocation, PageruleUpdateParamsActionsIDMirage, PageruleUpdateParamsActionsIDOpportunisticEncryption, PageruleUpdateParamsActionsIDOriginErrorPagePassThru, PageruleUpdateParamsActionsIDPolish, PageruleUpdateParamsActionsIDResolveOverride, PageruleUpdateParamsActionsIDRespectStrongEtag, PageruleUpdateParamsActionsIDResponseBuffering, PageruleUpdateParamsActionsIDRocketLoader, PageruleUpdateParamsActionsIDSecurityLevel, PageruleUpdateParamsActionsIDSortQueryStringForCache, PageruleUpdateParamsActionsIDSSL, PageruleUpdateParamsActionsIDTrueClientIPHeader, PageruleUpdateParamsActionsIDWAF:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleUpdateParamsStatus string

const (
	PageruleUpdateParamsStatusActive   PageruleUpdateParamsStatus = "active"
	PageruleUpdateParamsStatusDisabled PageruleUpdateParamsStatus = "disabled"
)

func (r PageruleUpdateParamsStatus) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsStatusActive, PageruleUpdateParamsStatusDisabled:
		return true
	}
	return false
}

type PageruleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageruleUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  PageRule                              `json:"result"`
	JSON    pageruleUpdateResponseEnvelopeJSON    `json:"-"`
}

// pageruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleUpdateResponseEnvelope]
type pageruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleUpdateResponseEnvelopeSuccess bool

const (
	PageruleUpdateResponseEnvelopeSuccessTrue PageruleUpdateResponseEnvelopeSuccess = true
)

func (r PageruleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned Page Rules.
	Direction param.Field[PageruleListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[PageruleListParamsMatch] `query:"match"`
	// The field used to sort returned Page Rules.
	Order param.Field[PageruleListParamsOrder] `query:"order"`
	// The status of the Page Rule.
	Status param.Field[PageruleListParamsStatus] `query:"status"`
}

// URLQuery serializes [PageruleListParams]'s query parameters as `url.Values`.
func (r PageruleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The direction used to sort returned Page Rules.
type PageruleListParamsDirection string

const (
	PageruleListParamsDirectionAsc  PageruleListParamsDirection = "asc"
	PageruleListParamsDirectionDesc PageruleListParamsDirection = "desc"
)

func (r PageruleListParamsDirection) IsKnown() bool {
	switch r {
	case PageruleListParamsDirectionAsc, PageruleListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type PageruleListParamsMatch string

const (
	PageruleListParamsMatchAny PageruleListParamsMatch = "any"
	PageruleListParamsMatchAll PageruleListParamsMatch = "all"
)

func (r PageruleListParamsMatch) IsKnown() bool {
	switch r {
	case PageruleListParamsMatchAny, PageruleListParamsMatchAll:
		return true
	}
	return false
}

// The field used to sort returned Page Rules.
type PageruleListParamsOrder string

const (
	PageruleListParamsOrderStatus   PageruleListParamsOrder = "status"
	PageruleListParamsOrderPriority PageruleListParamsOrder = "priority"
)

func (r PageruleListParamsOrder) IsKnown() bool {
	switch r {
	case PageruleListParamsOrderStatus, PageruleListParamsOrderPriority:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleListParamsStatus string

const (
	PageruleListParamsStatusActive   PageruleListParamsStatus = "active"
	PageruleListParamsStatusDisabled PageruleListParamsStatus = "disabled"
)

func (r PageruleListParamsStatus) IsKnown() bool {
	switch r {
	case PageruleListParamsStatusActive, PageruleListParamsStatusDisabled:
		return true
	}
	return false
}

type PageruleListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageruleListResponseEnvelopeSuccess `json:"success,required"`
	Result  []PageRule                          `json:"result"`
	JSON    pageruleListResponseEnvelopeJSON    `json:"-"`
}

// pageruleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleListResponseEnvelope]
type pageruleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleListResponseEnvelopeSuccess bool

const (
	PageruleListResponseEnvelopeSuccessTrue PageruleListResponseEnvelopeSuccess = true
)

func (r PageruleListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageruleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageruleDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  PageruleDeleteResponse                `json:"result,nullable"`
	JSON    pageruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// pageruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleDeleteResponseEnvelope]
type pageruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleDeleteResponseEnvelopeSuccess bool

const (
	PageruleDeleteResponseEnvelopeSuccessTrue PageruleDeleteResponseEnvelopeSuccess = true
)

func (r PageruleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]PageruleEditParamsActionUnion] `json:"actions"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageruleEditParamsStatus] `json:"status"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargetParam] `json:"targets"`
}

func (r PageruleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleEditParamsAction struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID    param.Field[PageruleEditParamsActionsID] `json:"id"`
	Value param.Field[interface{}]                 `json:"value"`
}

func (r PageruleEditParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsAction) ImplementsPagerulesPageruleEditParamsActionUnion() {}

// Satisfied by [zones.AlwaysUseHTTPSParam], [zones.AutomaticHTTPSRewritesParam],
// [zones.BrowserCacheTTLParam], [zones.BrowserCheckParam],
// [pagerules.PageruleEditParamsActionsBypassCacheOnCookie],
// [pagerules.PageruleEditParamsActionsCacheByDeviceType],
// [pagerules.PageruleEditParamsActionsCacheDeceptionArmor],
// [pagerules.PageruleEditParamsActionsCacheKey], [zones.CacheLevelParam],
// [pagerules.PageruleEditParamsActionsCacheOnCookie],
// [pagerules.PageruleEditParamsActionsDisableApps],
// [pagerules.PageruleEditParamsActionsDisablePerformance],
// [pagerules.PageruleEditParamsActionsDisableSecurity],
// [pagerules.PageruleEditParamsActionsDisableZaraz],
// [pagerules.PageruleEditParamsActionsEdgeCacheTTL],
// [zones.EmailObfuscationParam],
// [pagerules.PageruleEditParamsActionsExplicitCacheControl],
// [pagerules.PageruleEditParamsActionsForwardingURL],
// [pagerules.PageruleEditParamsActionsHostHeaderOverride],
// [zones.IPGeolocationParam], [zones.MirageParam],
// [zones.OpportunisticEncryptionParam], [zones.OriginErrorPagePassThruParam],
// [zones.PolishParam], [pagerules.PageruleEditParamsActionsResolveOverride],
// [pagerules.PageruleEditParamsActionsRespectStrongEtag],
// [zones.ResponseBufferingParam], [zones.RocketLoaderParam],
// [zones.SecurityLevelParam], [zones.SortQueryStringForCacheParam],
// [zones.SSLParam], [zones.TrueClientIPHeaderParam], [zones.WAFParam],
// [PageruleEditParamsAction].
type PageruleEditParamsActionUnion interface {
	ImplementsPagerulesPageruleEditParamsActionUnion()
}

type PageruleEditParamsActionsBypassCacheOnCookie struct {
	// Bypass cache and fetch resources from the origin server if a regular expression
	// matches against a cookie name present in the request.
	ID param.Field[PageruleEditParamsActionsBypassCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request. Refer to
	// [Bypass Cache on Cookie setting](https://developers.cloudflare.com/rules/page-rules/reference/additional-reference/#bypass-cache-on-cookie-setting)
	// to learn about limited regular expression support.
	Value param.Field[string] `json:"value"`
}

func (r PageruleEditParamsActionsBypassCacheOnCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsBypassCacheOnCookie) ImplementsPagerulesPageruleEditParamsActionUnion() {
}

// Bypass cache and fetch resources from the origin server if a regular expression
// matches against a cookie name present in the request.
type PageruleEditParamsActionsBypassCacheOnCookieID string

const (
	PageruleEditParamsActionsBypassCacheOnCookieIDBypassCacheOnCookie PageruleEditParamsActionsBypassCacheOnCookieID = "bypass_cache_on_cookie"
)

func (r PageruleEditParamsActionsBypassCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsBypassCacheOnCookieIDBypassCacheOnCookie:
		return true
	}
	return false
}

type PageruleEditParamsActionsCacheByDeviceType struct {
	// Separate cached content based on the visitor's device type.
	ID param.Field[PageruleEditParamsActionsCacheByDeviceTypeID] `json:"id"`
	// The status of Cache By Device Type.
	Value param.Field[PageruleEditParamsActionsCacheByDeviceTypeValue] `json:"value"`
}

func (r PageruleEditParamsActionsCacheByDeviceType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsCacheByDeviceType) ImplementsPagerulesPageruleEditParamsActionUnion() {
}

// Separate cached content based on the visitor's device type.
type PageruleEditParamsActionsCacheByDeviceTypeID string

const (
	PageruleEditParamsActionsCacheByDeviceTypeIDCacheByDeviceType PageruleEditParamsActionsCacheByDeviceTypeID = "cache_by_device_type"
)

func (r PageruleEditParamsActionsCacheByDeviceTypeID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsCacheByDeviceTypeIDCacheByDeviceType:
		return true
	}
	return false
}

// The status of Cache By Device Type.
type PageruleEditParamsActionsCacheByDeviceTypeValue string

const (
	PageruleEditParamsActionsCacheByDeviceTypeValueOn  PageruleEditParamsActionsCacheByDeviceTypeValue = "on"
	PageruleEditParamsActionsCacheByDeviceTypeValueOff PageruleEditParamsActionsCacheByDeviceTypeValue = "off"
)

func (r PageruleEditParamsActionsCacheByDeviceTypeValue) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsCacheByDeviceTypeValueOn, PageruleEditParamsActionsCacheByDeviceTypeValueOff:
		return true
	}
	return false
}

type PageruleEditParamsActionsCacheDeceptionArmor struct {
	// Protect from web cache deception attacks while still allowing static assets to
	// be cached. This setting verifies that the URL's extension matches the returned
	// `Content-Type`.
	ID param.Field[PageruleEditParamsActionsCacheDeceptionArmorID] `json:"id"`
	// The status of Cache Deception Armor.
	Value param.Field[PageruleEditParamsActionsCacheDeceptionArmorValue] `json:"value"`
}

func (r PageruleEditParamsActionsCacheDeceptionArmor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsCacheDeceptionArmor) ImplementsPagerulesPageruleEditParamsActionUnion() {
}

// Protect from web cache deception attacks while still allowing static assets to
// be cached. This setting verifies that the URL's extension matches the returned
// `Content-Type`.
type PageruleEditParamsActionsCacheDeceptionArmorID string

const (
	PageruleEditParamsActionsCacheDeceptionArmorIDCacheDeceptionArmor PageruleEditParamsActionsCacheDeceptionArmorID = "cache_deception_armor"
)

func (r PageruleEditParamsActionsCacheDeceptionArmorID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsCacheDeceptionArmorIDCacheDeceptionArmor:
		return true
	}
	return false
}

// The status of Cache Deception Armor.
type PageruleEditParamsActionsCacheDeceptionArmorValue string

const (
	PageruleEditParamsActionsCacheDeceptionArmorValueOn  PageruleEditParamsActionsCacheDeceptionArmorValue = "on"
	PageruleEditParamsActionsCacheDeceptionArmorValueOff PageruleEditParamsActionsCacheDeceptionArmorValue = "off"
)

func (r PageruleEditParamsActionsCacheDeceptionArmorValue) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsCacheDeceptionArmorValueOn, PageruleEditParamsActionsCacheDeceptionArmorValueOff:
		return true
	}
	return false
}

type PageruleEditParamsActionsCacheKey struct {
	// Control specifically what variables to include when deciding which resources to
	// cache. This allows customers to determine what to cache based on something other
	// than just the URL.
	ID    param.Field[PageruleEditParamsActionsCacheKeyID]    `json:"id"`
	Value param.Field[PageruleEditParamsActionsCacheKeyValue] `json:"value"`
}

func (r PageruleEditParamsActionsCacheKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsCacheKey) ImplementsPagerulesPageruleEditParamsActionUnion() {}

// Control specifically what variables to include when deciding which resources to
// cache. This allows customers to determine what to cache based on something other
// than just the URL.
type PageruleEditParamsActionsCacheKeyID string

const (
	PageruleEditParamsActionsCacheKeyIDCacheKey PageruleEditParamsActionsCacheKeyID = "cache_key"
)

func (r PageruleEditParamsActionsCacheKeyID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsCacheKeyIDCacheKey:
		return true
	}
	return false
}

type PageruleEditParamsActionsCacheKeyValue struct {
	// Controls which cookies appear in the Cache Key.
	Cookie param.Field[PageruleEditParamsActionsCacheKeyValueCookie] `json:"cookie"`
	// Controls which headers go into the Cache Key. Exactly one of `include` or
	// `exclude` is expected.
	Header param.Field[PageruleEditParamsActionsCacheKeyValueHeader] `json:"header"`
	// Determines which host header to include in the Cache Key.
	Host param.Field[PageruleEditParamsActionsCacheKeyValueHost] `json:"host"`
	// Controls which URL query string parameters go into the Cache Key. Exactly one of
	// `include` or `exclude` is expected.
	QueryString param.Field[PageruleEditParamsActionsCacheKeyValueQueryString] `json:"query_string"`
	// Feature fields to add features about the end-user (client) into the Cache Key.
	User param.Field[PageruleEditParamsActionsCacheKeyValueUser] `json:"user"`
}

func (r PageruleEditParamsActionsCacheKeyValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which cookies appear in the Cache Key.
type PageruleEditParamsActionsCacheKeyValueCookie struct {
	// A list of cookies to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of cookies to include.
	Include param.Field[[]string] `json:"include"`
}

func (r PageruleEditParamsActionsCacheKeyValueCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which headers go into the Cache Key. Exactly one of `include` or
// `exclude` is expected.
type PageruleEditParamsActionsCacheKeyValueHeader struct {
	// A list of headers to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of headers to ignore.
	Exclude param.Field[[]string] `json:"exclude"`
	// A list of headers to include.
	Include param.Field[[]string] `json:"include"`
}

func (r PageruleEditParamsActionsCacheKeyValueHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which host header to include in the Cache Key.
type PageruleEditParamsActionsCacheKeyValueHost struct {
	// Whether to include the Host header in the HTTP request sent to the origin.
	Resolved param.Field[bool] `json:"resolved"`
}

func (r PageruleEditParamsActionsCacheKeyValueHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which URL query string parameters go into the Cache Key. Exactly one of
// `include` or `exclude` is expected.
type PageruleEditParamsActionsCacheKeyValueQueryString struct {
	// Ignore all query string parameters.
	Exclude param.Field[PageruleEditParamsActionsCacheKeyValueQueryStringExcludeUnion] `json:"exclude"`
	// Include all query string parameters.
	Include param.Field[PageruleEditParamsActionsCacheKeyValueQueryStringIncludeUnion] `json:"include"`
}

func (r PageruleEditParamsActionsCacheKeyValueQueryString) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Ignore all query string parameters.
//
// Satisfied by
// [pagerules.PageruleEditParamsActionsCacheKeyValueQueryStringExcludeString],
// [pagerules.PageruleEditParamsActionsCacheKeyValueQueryStringExcludeArray].
type PageruleEditParamsActionsCacheKeyValueQueryStringExcludeUnion interface {
	implementsPagerulesPageruleEditParamsActionsCacheKeyValueQueryStringExcludeUnion()
}

// Ignore all query string parameters.
type PageruleEditParamsActionsCacheKeyValueQueryStringExcludeString string

const (
	PageruleEditParamsActionsCacheKeyValueQueryStringExcludeStringStar PageruleEditParamsActionsCacheKeyValueQueryStringExcludeString = "*"
)

func (r PageruleEditParamsActionsCacheKeyValueQueryStringExcludeString) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsCacheKeyValueQueryStringExcludeStringStar:
		return true
	}
	return false
}

func (r PageruleEditParamsActionsCacheKeyValueQueryStringExcludeString) implementsPagerulesPageruleEditParamsActionsCacheKeyValueQueryStringExcludeUnion() {
}

type PageruleEditParamsActionsCacheKeyValueQueryStringExcludeArray []string

func (r PageruleEditParamsActionsCacheKeyValueQueryStringExcludeArray) implementsPagerulesPageruleEditParamsActionsCacheKeyValueQueryStringExcludeUnion() {
}

// Include all query string parameters.
//
// Satisfied by
// [pagerules.PageruleEditParamsActionsCacheKeyValueQueryStringIncludeString],
// [pagerules.PageruleEditParamsActionsCacheKeyValueQueryStringIncludeArray].
type PageruleEditParamsActionsCacheKeyValueQueryStringIncludeUnion interface {
	implementsPagerulesPageruleEditParamsActionsCacheKeyValueQueryStringIncludeUnion()
}

// Include all query string parameters.
type PageruleEditParamsActionsCacheKeyValueQueryStringIncludeString string

const (
	PageruleEditParamsActionsCacheKeyValueQueryStringIncludeStringStar PageruleEditParamsActionsCacheKeyValueQueryStringIncludeString = "*"
)

func (r PageruleEditParamsActionsCacheKeyValueQueryStringIncludeString) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsCacheKeyValueQueryStringIncludeStringStar:
		return true
	}
	return false
}

func (r PageruleEditParamsActionsCacheKeyValueQueryStringIncludeString) implementsPagerulesPageruleEditParamsActionsCacheKeyValueQueryStringIncludeUnion() {
}

type PageruleEditParamsActionsCacheKeyValueQueryStringIncludeArray []string

func (r PageruleEditParamsActionsCacheKeyValueQueryStringIncludeArray) implementsPagerulesPageruleEditParamsActionsCacheKeyValueQueryStringIncludeUnion() {
}

// Feature fields to add features about the end-user (client) into the Cache Key.
type PageruleEditParamsActionsCacheKeyValueUser struct {
	// Classifies a request as `mobile`, `desktop`, or `tablet` based on the User
	// Agent.
	DeviceType param.Field[bool] `json:"device_type"`
	// Includes the client's country, derived from the IP address.
	Geo param.Field[bool] `json:"geo"`
	// Includes the first language code contained in the `Accept-Language` header sent
	// by the client.
	Lang param.Field[bool] `json:"lang"`
}

func (r PageruleEditParamsActionsCacheKeyValueUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleEditParamsActionsCacheOnCookie struct {
	// Apply the Cache Everything option (Cache Level setting) based on a regular
	// expression match against a cookie name.
	ID param.Field[PageruleEditParamsActionsCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request.
	Value param.Field[string] `json:"value"`
}

func (r PageruleEditParamsActionsCacheOnCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsCacheOnCookie) ImplementsPagerulesPageruleEditParamsActionUnion() {}

// Apply the Cache Everything option (Cache Level setting) based on a regular
// expression match against a cookie name.
type PageruleEditParamsActionsCacheOnCookieID string

const (
	PageruleEditParamsActionsCacheOnCookieIDCacheOnCookie PageruleEditParamsActionsCacheOnCookieID = "cache_on_cookie"
)

func (r PageruleEditParamsActionsCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsCacheOnCookieIDCacheOnCookie:
		return true
	}
	return false
}

type PageruleEditParamsActionsDisableApps struct {
	// Turn off all active
	// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
	// (deprecated).
	ID param.Field[PageruleEditParamsActionsDisableAppsID] `json:"id"`
}

func (r PageruleEditParamsActionsDisableApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsDisableApps) ImplementsPagerulesPageruleEditParamsActionUnion() {}

// Turn off all active
// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
// (deprecated).
type PageruleEditParamsActionsDisableAppsID string

const (
	PageruleEditParamsActionsDisableAppsIDDisableApps PageruleEditParamsActionsDisableAppsID = "disable_apps"
)

func (r PageruleEditParamsActionsDisableAppsID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsDisableAppsIDDisableApps:
		return true
	}
	return false
}

type PageruleEditParamsActionsDisablePerformance struct {
	// Turn off
	// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
	// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
	// and [Polish](https://developers.cloudflare.com/images/polish/).
	ID param.Field[PageruleEditParamsActionsDisablePerformanceID] `json:"id"`
}

func (r PageruleEditParamsActionsDisablePerformance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsDisablePerformance) ImplementsPagerulesPageruleEditParamsActionUnion() {
}

// Turn off
// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
// and [Polish](https://developers.cloudflare.com/images/polish/).
type PageruleEditParamsActionsDisablePerformanceID string

const (
	PageruleEditParamsActionsDisablePerformanceIDDisablePerformance PageruleEditParamsActionsDisablePerformanceID = "disable_performance"
)

func (r PageruleEditParamsActionsDisablePerformanceID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsDisablePerformanceIDDisablePerformance:
		return true
	}
	return false
}

type PageruleEditParamsActionsDisableSecurity struct {
	// Turn off
	// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
	// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
	// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
	// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
	// and
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	ID param.Field[PageruleEditParamsActionsDisableSecurityID] `json:"id"`
}

func (r PageruleEditParamsActionsDisableSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsDisableSecurity) ImplementsPagerulesPageruleEditParamsActionUnion() {
}

// Turn off
// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
// and
// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
type PageruleEditParamsActionsDisableSecurityID string

const (
	PageruleEditParamsActionsDisableSecurityIDDisableSecurity PageruleEditParamsActionsDisableSecurityID = "disable_security"
)

func (r PageruleEditParamsActionsDisableSecurityID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsDisableSecurityIDDisableSecurity:
		return true
	}
	return false
}

type PageruleEditParamsActionsDisableZaraz struct {
	// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
	ID param.Field[PageruleEditParamsActionsDisableZarazID] `json:"id"`
}

func (r PageruleEditParamsActionsDisableZaraz) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsDisableZaraz) ImplementsPagerulesPageruleEditParamsActionUnion() {}

// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
type PageruleEditParamsActionsDisableZarazID string

const (
	PageruleEditParamsActionsDisableZarazIDDisableZaraz PageruleEditParamsActionsDisableZarazID = "disable_zaraz"
)

func (r PageruleEditParamsActionsDisableZarazID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsDisableZarazIDDisableZaraz:
		return true
	}
	return false
}

type PageruleEditParamsActionsEdgeCacheTTL struct {
	// Specify how long to cache a resource in the Cloudflare global network. _Edge
	// Cache TTL_ is not visible in response headers.
	ID    param.Field[PageruleEditParamsActionsEdgeCacheTTLID] `json:"id"`
	Value param.Field[int64]                                   `json:"value"`
}

func (r PageruleEditParamsActionsEdgeCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsEdgeCacheTTL) ImplementsPagerulesPageruleEditParamsActionUnion() {}

// Specify how long to cache a resource in the Cloudflare global network. _Edge
// Cache TTL_ is not visible in response headers.
type PageruleEditParamsActionsEdgeCacheTTLID string

const (
	PageruleEditParamsActionsEdgeCacheTTLIDEdgeCacheTTL PageruleEditParamsActionsEdgeCacheTTLID = "edge_cache_ttl"
)

func (r PageruleEditParamsActionsEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

type PageruleEditParamsActionsExplicitCacheControl struct {
	// Origin Cache Control is enabled by default for Free, Pro, and Business domains
	// and disabled by default for Enterprise domains.
	ID param.Field[PageruleEditParamsActionsExplicitCacheControlID] `json:"id"`
	// The status of Origin Cache Control.
	Value param.Field[PageruleEditParamsActionsExplicitCacheControlValue] `json:"value"`
}

func (r PageruleEditParamsActionsExplicitCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsExplicitCacheControl) ImplementsPagerulesPageruleEditParamsActionUnion() {
}

// Origin Cache Control is enabled by default for Free, Pro, and Business domains
// and disabled by default for Enterprise domains.
type PageruleEditParamsActionsExplicitCacheControlID string

const (
	PageruleEditParamsActionsExplicitCacheControlIDExplicitCacheControl PageruleEditParamsActionsExplicitCacheControlID = "explicit_cache_control"
)

func (r PageruleEditParamsActionsExplicitCacheControlID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsExplicitCacheControlIDExplicitCacheControl:
		return true
	}
	return false
}

// The status of Origin Cache Control.
type PageruleEditParamsActionsExplicitCacheControlValue string

const (
	PageruleEditParamsActionsExplicitCacheControlValueOn  PageruleEditParamsActionsExplicitCacheControlValue = "on"
	PageruleEditParamsActionsExplicitCacheControlValueOff PageruleEditParamsActionsExplicitCacheControlValue = "off"
)

func (r PageruleEditParamsActionsExplicitCacheControlValue) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsExplicitCacheControlValueOn, PageruleEditParamsActionsExplicitCacheControlValueOff:
		return true
	}
	return false
}

type PageruleEditParamsActionsForwardingURL struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    param.Field[PageruleEditParamsActionsForwardingURLID]    `json:"id"`
	Value param.Field[PageruleEditParamsActionsForwardingURLValue] `json:"value"`
}

func (r PageruleEditParamsActionsForwardingURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsForwardingURL) ImplementsPagerulesPageruleEditParamsActionUnion() {}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageruleEditParamsActionsForwardingURLID string

const (
	PageruleEditParamsActionsForwardingURLIDForwardingURL PageruleEditParamsActionsForwardingURLID = "forwarding_url"
)

func (r PageruleEditParamsActionsForwardingURLID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsForwardingURLIDForwardingURL:
		return true
	}
	return false
}

type PageruleEditParamsActionsForwardingURLValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode param.Field[PageruleEditParamsActionsForwardingURLValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageruleEditParamsActionsForwardingURLValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageruleEditParamsActionsForwardingURLValueStatusCode int64

const (
	PageruleEditParamsActionsForwardingURLValueStatusCode301 PageruleEditParamsActionsForwardingURLValueStatusCode = 301
	PageruleEditParamsActionsForwardingURLValueStatusCode302 PageruleEditParamsActionsForwardingURLValueStatusCode = 302
)

func (r PageruleEditParamsActionsForwardingURLValueStatusCode) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsForwardingURLValueStatusCode301, PageruleEditParamsActionsForwardingURLValueStatusCode302:
		return true
	}
	return false
}

type PageruleEditParamsActionsHostHeaderOverride struct {
	// Apply a specific host header.
	ID param.Field[PageruleEditParamsActionsHostHeaderOverrideID] `json:"id"`
	// The hostname to use in the `Host` header
	Value param.Field[string] `json:"value"`
}

func (r PageruleEditParamsActionsHostHeaderOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsHostHeaderOverride) ImplementsPagerulesPageruleEditParamsActionUnion() {
}

// Apply a specific host header.
type PageruleEditParamsActionsHostHeaderOverrideID string

const (
	PageruleEditParamsActionsHostHeaderOverrideIDHostHeaderOverride PageruleEditParamsActionsHostHeaderOverrideID = "host_header_override"
)

func (r PageruleEditParamsActionsHostHeaderOverrideID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsHostHeaderOverrideIDHostHeaderOverride:
		return true
	}
	return false
}

type PageruleEditParamsActionsResolveOverride struct {
	// Change the origin address to the value specified in this setting.
	ID param.Field[PageruleEditParamsActionsResolveOverrideID] `json:"id"`
	// The origin address you want to override with.
	Value param.Field[string] `json:"value"`
}

func (r PageruleEditParamsActionsResolveOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsResolveOverride) ImplementsPagerulesPageruleEditParamsActionUnion() {
}

// Change the origin address to the value specified in this setting.
type PageruleEditParamsActionsResolveOverrideID string

const (
	PageruleEditParamsActionsResolveOverrideIDResolveOverride PageruleEditParamsActionsResolveOverrideID = "resolve_override"
)

func (r PageruleEditParamsActionsResolveOverrideID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsResolveOverrideIDResolveOverride:
		return true
	}
	return false
}

type PageruleEditParamsActionsRespectStrongEtag struct {
	// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
	// the origin server.
	ID param.Field[PageruleEditParamsActionsRespectStrongEtagID] `json:"id"`
	// The status of Respect Strong ETags
	Value param.Field[PageruleEditParamsActionsRespectStrongEtagValue] `json:"value"`
}

func (r PageruleEditParamsActionsRespectStrongEtag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageruleEditParamsActionsRespectStrongEtag) ImplementsPagerulesPageruleEditParamsActionUnion() {
}

// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
// the origin server.
type PageruleEditParamsActionsRespectStrongEtagID string

const (
	PageruleEditParamsActionsRespectStrongEtagIDRespectStrongEtag PageruleEditParamsActionsRespectStrongEtagID = "respect_strong_etag"
)

func (r PageruleEditParamsActionsRespectStrongEtagID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsRespectStrongEtagIDRespectStrongEtag:
		return true
	}
	return false
}

// The status of Respect Strong ETags
type PageruleEditParamsActionsRespectStrongEtagValue string

const (
	PageruleEditParamsActionsRespectStrongEtagValueOn  PageruleEditParamsActionsRespectStrongEtagValue = "on"
	PageruleEditParamsActionsRespectStrongEtagValueOff PageruleEditParamsActionsRespectStrongEtagValue = "off"
)

func (r PageruleEditParamsActionsRespectStrongEtagValue) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsRespectStrongEtagValueOn, PageruleEditParamsActionsRespectStrongEtagValueOff:
		return true
	}
	return false
}

// If enabled, any ` http://“ URL is converted to  `https://` through a 301
// redirect.
type PageruleEditParamsActionsID string

const (
	PageruleEditParamsActionsIDAlwaysUseHTTPS          PageruleEditParamsActionsID = "always_use_https"
	PageruleEditParamsActionsIDAutomaticHTTPSRewrites  PageruleEditParamsActionsID = "automatic_https_rewrites"
	PageruleEditParamsActionsIDBrowserCacheTTL         PageruleEditParamsActionsID = "browser_cache_ttl"
	PageruleEditParamsActionsIDBrowserCheck            PageruleEditParamsActionsID = "browser_check"
	PageruleEditParamsActionsIDBypassCacheOnCookie     PageruleEditParamsActionsID = "bypass_cache_on_cookie"
	PageruleEditParamsActionsIDCacheByDeviceType       PageruleEditParamsActionsID = "cache_by_device_type"
	PageruleEditParamsActionsIDCacheDeceptionArmor     PageruleEditParamsActionsID = "cache_deception_armor"
	PageruleEditParamsActionsIDCacheKey                PageruleEditParamsActionsID = "cache_key"
	PageruleEditParamsActionsIDCacheLevel              PageruleEditParamsActionsID = "cache_level"
	PageruleEditParamsActionsIDCacheOnCookie           PageruleEditParamsActionsID = "cache_on_cookie"
	PageruleEditParamsActionsIDDisableApps             PageruleEditParamsActionsID = "disable_apps"
	PageruleEditParamsActionsIDDisablePerformance      PageruleEditParamsActionsID = "disable_performance"
	PageruleEditParamsActionsIDDisableSecurity         PageruleEditParamsActionsID = "disable_security"
	PageruleEditParamsActionsIDDisableZaraz            PageruleEditParamsActionsID = "disable_zaraz"
	PageruleEditParamsActionsIDEdgeCacheTTL            PageruleEditParamsActionsID = "edge_cache_ttl"
	PageruleEditParamsActionsIDEmailObfuscation        PageruleEditParamsActionsID = "email_obfuscation"
	PageruleEditParamsActionsIDExplicitCacheControl    PageruleEditParamsActionsID = "explicit_cache_control"
	PageruleEditParamsActionsIDForwardingURL           PageruleEditParamsActionsID = "forwarding_url"
	PageruleEditParamsActionsIDHostHeaderOverride      PageruleEditParamsActionsID = "host_header_override"
	PageruleEditParamsActionsIDIPGeolocation           PageruleEditParamsActionsID = "ip_geolocation"
	PageruleEditParamsActionsIDMirage                  PageruleEditParamsActionsID = "mirage"
	PageruleEditParamsActionsIDOpportunisticEncryption PageruleEditParamsActionsID = "opportunistic_encryption"
	PageruleEditParamsActionsIDOriginErrorPagePassThru PageruleEditParamsActionsID = "origin_error_page_pass_thru"
	PageruleEditParamsActionsIDPolish                  PageruleEditParamsActionsID = "polish"
	PageruleEditParamsActionsIDResolveOverride         PageruleEditParamsActionsID = "resolve_override"
	PageruleEditParamsActionsIDRespectStrongEtag       PageruleEditParamsActionsID = "respect_strong_etag"
	PageruleEditParamsActionsIDResponseBuffering       PageruleEditParamsActionsID = "response_buffering"
	PageruleEditParamsActionsIDRocketLoader            PageruleEditParamsActionsID = "rocket_loader"
	PageruleEditParamsActionsIDSecurityLevel           PageruleEditParamsActionsID = "security_level"
	PageruleEditParamsActionsIDSortQueryStringForCache PageruleEditParamsActionsID = "sort_query_string_for_cache"
	PageruleEditParamsActionsIDSSL                     PageruleEditParamsActionsID = "ssl"
	PageruleEditParamsActionsIDTrueClientIPHeader      PageruleEditParamsActionsID = "true_client_ip_header"
	PageruleEditParamsActionsIDWAF                     PageruleEditParamsActionsID = "waf"
)

func (r PageruleEditParamsActionsID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsIDAlwaysUseHTTPS, PageruleEditParamsActionsIDAutomaticHTTPSRewrites, PageruleEditParamsActionsIDBrowserCacheTTL, PageruleEditParamsActionsIDBrowserCheck, PageruleEditParamsActionsIDBypassCacheOnCookie, PageruleEditParamsActionsIDCacheByDeviceType, PageruleEditParamsActionsIDCacheDeceptionArmor, PageruleEditParamsActionsIDCacheKey, PageruleEditParamsActionsIDCacheLevel, PageruleEditParamsActionsIDCacheOnCookie, PageruleEditParamsActionsIDDisableApps, PageruleEditParamsActionsIDDisablePerformance, PageruleEditParamsActionsIDDisableSecurity, PageruleEditParamsActionsIDDisableZaraz, PageruleEditParamsActionsIDEdgeCacheTTL, PageruleEditParamsActionsIDEmailObfuscation, PageruleEditParamsActionsIDExplicitCacheControl, PageruleEditParamsActionsIDForwardingURL, PageruleEditParamsActionsIDHostHeaderOverride, PageruleEditParamsActionsIDIPGeolocation, PageruleEditParamsActionsIDMirage, PageruleEditParamsActionsIDOpportunisticEncryption, PageruleEditParamsActionsIDOriginErrorPagePassThru, PageruleEditParamsActionsIDPolish, PageruleEditParamsActionsIDResolveOverride, PageruleEditParamsActionsIDRespectStrongEtag, PageruleEditParamsActionsIDResponseBuffering, PageruleEditParamsActionsIDRocketLoader, PageruleEditParamsActionsIDSecurityLevel, PageruleEditParamsActionsIDSortQueryStringForCache, PageruleEditParamsActionsIDSSL, PageruleEditParamsActionsIDTrueClientIPHeader, PageruleEditParamsActionsIDWAF:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleEditParamsStatus string

const (
	PageruleEditParamsStatusActive   PageruleEditParamsStatus = "active"
	PageruleEditParamsStatusDisabled PageruleEditParamsStatus = "disabled"
)

func (r PageruleEditParamsStatus) IsKnown() bool {
	switch r {
	case PageruleEditParamsStatusActive, PageruleEditParamsStatusDisabled:
		return true
	}
	return false
}

type PageruleEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageruleEditResponseEnvelopeSuccess `json:"success,required"`
	Result  PageRule                            `json:"result"`
	JSON    pageruleEditResponseEnvelopeJSON    `json:"-"`
}

// pageruleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleEditResponseEnvelope]
type pageruleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleEditResponseEnvelopeSuccess bool

const (
	PageruleEditResponseEnvelopeSuccessTrue PageruleEditResponseEnvelopeSuccess = true
)

func (r PageruleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageruleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageruleGetResponseEnvelopeSuccess `json:"success,required"`
	Result  PageRule                           `json:"result"`
	JSON    pageruleGetResponseEnvelopeJSON    `json:"-"`
}

// pageruleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleGetResponseEnvelope]
type pageruleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleGetResponseEnvelopeSuccess bool

const (
	PageruleGetResponseEnvelopeSuccessTrue PageruleGetResponseEnvelopeSuccess = true
)

func (r PageruleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
