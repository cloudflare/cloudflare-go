// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rules

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
	"github.com/cloudflare/cloudflare-go/v3/zones"
	"github.com/tidwall/gjson"
)

// PageRuleService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageRuleService] method instead.
type PageRuleService struct {
	Options []option.RequestOption
}

// NewPageRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageRuleService(opts ...option.RequestOption) (r *PageRuleService) {
	r = &PageRuleService{}
	r.Options = opts
	return
}

// Creates a new Page Rule.
func (r *PageRuleService) New(ctx context.Context, params PageRuleNewParams, opts ...option.RequestOption) (res *PageRule, err error) {
	var env PageRuleNewResponseEnvelope
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
func (r *PageRuleService) Update(ctx context.Context, pageruleID string, params PageRuleUpdateParams, opts ...option.RequestOption) (res *PageRule, err error) {
	var env PageRuleUpdateResponseEnvelope
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
func (r *PageRuleService) List(ctx context.Context, params PageRuleListParams, opts ...option.RequestOption) (res *[]PageRule, err error) {
	var env PageRuleListResponseEnvelope
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
func (r *PageRuleService) Delete(ctx context.Context, pageruleID string, body PageRuleDeleteParams, opts ...option.RequestOption) (res *PageRuleDeleteResponse, err error) {
	var env PageRuleDeleteResponseEnvelope
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
func (r *PageRuleService) Edit(ctx context.Context, pageruleID string, params PageRuleEditParams, opts ...option.RequestOption) (res *PageRule, err error) {
	var env PageRuleEditResponseEnvelope
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
func (r *PageRuleService) Get(ctx context.Context, pageruleID string, query PageRuleGetParams, opts ...option.RequestOption) (res *PageRule, err error) {
	var env PageRuleGetResponseEnvelope
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
// [page_rules.PageRuleActionsBypassCacheOnCookie],
// [page_rules.PageRuleActionsCacheByDeviceType],
// [page_rules.PageRuleActionsCacheDeceptionArmor],
// [page_rules.PageRuleActionsCacheKey], [zones.CacheLevel],
// [page_rules.PageRuleActionsCacheOnCookie],
// [page_rules.PageRuleActionsDisableApps],
// [page_rules.PageRuleActionsDisablePerformance],
// [page_rules.PageRuleActionsDisableSecurity],
// [page_rules.PageRuleActionsDisableZaraz],
// [page_rules.PageRuleActionsEdgeCacheTTL], [zones.EmailObfuscation],
// [page_rules.PageRuleActionsExplicitCacheControl],
// [page_rules.PageRuleActionsForwardingURL],
// [page_rules.PageRuleActionsHostHeaderOverride], [zones.IPGeolocation],
// [zones.Mirage], [zones.OpportunisticEncryption],
// [zones.OriginErrorPagePassThru], [zones.Polish],
// [page_rules.PageRuleActionsResolveOverride],
// [page_rules.PageRuleActionsRespectStrongEtag], [zones.ResponseBuffering],
// [zones.RocketLoader], [zones.SecurityLevel], [zones.SortQueryStringForCache],
// [zones.SSL], [zones.TrueClientIPHeader], [zones.WAF].
func (r PageRuleAction) AsUnion() PageRuleActionsUnion {
	return r.union
}

// Union satisfied by [zones.AlwaysUseHTTPS], [zones.AutomaticHTTPSRewrites],
// [zones.BrowserCacheTTL], [zones.BrowserCheck],
// [page_rules.PageRuleActionsBypassCacheOnCookie],
// [page_rules.PageRuleActionsCacheByDeviceType],
// [page_rules.PageRuleActionsCacheDeceptionArmor],
// [page_rules.PageRuleActionsCacheKey], [zones.CacheLevel],
// [page_rules.PageRuleActionsCacheOnCookie],
// [page_rules.PageRuleActionsDisableApps],
// [page_rules.PageRuleActionsDisablePerformance],
// [page_rules.PageRuleActionsDisableSecurity],
// [page_rules.PageRuleActionsDisableZaraz],
// [page_rules.PageRuleActionsEdgeCacheTTL], [zones.EmailObfuscation],
// [page_rules.PageRuleActionsExplicitCacheControl],
// [page_rules.PageRuleActionsForwardingURL],
// [page_rules.PageRuleActionsHostHeaderOverride], [zones.IPGeolocation],
// [zones.Mirage], [zones.OpportunisticEncryption],
// [zones.OriginErrorPagePassThru], [zones.Polish],
// [page_rules.PageRuleActionsResolveOverride],
// [page_rules.PageRuleActionsRespectStrongEtag], [zones.ResponseBuffering],
// [zones.RocketLoader], [zones.SecurityLevel], [zones.SortQueryStringForCache],
// [zones.SSL], [zones.TrueClientIPHeader] or [zones.WAF].
type PageRuleActionsUnion interface {
	ImplementsPageRulesPageRuleAction()
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

func (r PageRuleActionsBypassCacheOnCookie) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsCacheByDeviceType) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsCacheDeceptionArmor) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsCacheKey) ImplementsPageRulesPageRuleAction() {}

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
// [page_rules.PageRuleActionsCacheKeyValueQueryStringExcludeString] or
// [page_rules.PageRuleActionsCacheKeyValueQueryStringExcludeArray].
type PageRuleActionsCacheKeyValueQueryStringExcludeUnion interface {
	implementsPageRulesPageRuleActionsCacheKeyValueQueryStringExcludeUnion()
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

func (r PageRuleActionsCacheKeyValueQueryStringExcludeString) implementsPageRulesPageRuleActionsCacheKeyValueQueryStringExcludeUnion() {
}

type PageRuleActionsCacheKeyValueQueryStringExcludeArray []string

func (r PageRuleActionsCacheKeyValueQueryStringExcludeArray) implementsPageRulesPageRuleActionsCacheKeyValueQueryStringExcludeUnion() {
}

// Include all query string parameters.
//
// Union satisfied by
// [page_rules.PageRuleActionsCacheKeyValueQueryStringIncludeString] or
// [page_rules.PageRuleActionsCacheKeyValueQueryStringIncludeArray].
type PageRuleActionsCacheKeyValueQueryStringIncludeUnion interface {
	implementsPageRulesPageRuleActionsCacheKeyValueQueryStringIncludeUnion()
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

func (r PageRuleActionsCacheKeyValueQueryStringIncludeString) implementsPageRulesPageRuleActionsCacheKeyValueQueryStringIncludeUnion() {
}

type PageRuleActionsCacheKeyValueQueryStringIncludeArray []string

func (r PageRuleActionsCacheKeyValueQueryStringIncludeArray) implementsPageRulesPageRuleActionsCacheKeyValueQueryStringIncludeUnion() {
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

func (r PageRuleActionsCacheOnCookie) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsDisableApps) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsDisablePerformance) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsDisableSecurity) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsDisableZaraz) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsEdgeCacheTTL) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsExplicitCacheControl) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsForwardingURL) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsHostHeaderOverride) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsResolveOverride) ImplementsPageRulesPageRuleAction() {}

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

func (r PageRuleActionsRespectStrongEtag) ImplementsPageRulesPageRuleAction() {}

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

type PageRuleDeleteResponse struct {
	// Identifier
	ID   string                     `json:"id,required"`
	JSON pageRuleDeleteResponseJSON `json:"-"`
}

// pageRuleDeleteResponseJSON contains the JSON metadata for the struct
// [PageRuleDeleteResponse]
type pageRuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PageRuleNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]PageRuleNewParamsActionUnion] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargetParam] `json:"targets,required"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageRuleNewParamsStatus] `json:"status"`
}

func (r PageRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageRuleNewParamsAction struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID    param.Field[PageRuleNewParamsActionsID] `json:"id"`
	Value param.Field[interface{}]                `json:"value"`
}

func (r PageRuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsAction) ImplementsPageRulesPageRuleNewParamsActionUnion() {}

// Satisfied by [zones.AlwaysUseHTTPSParam], [zones.AutomaticHTTPSRewritesParam],
// [zones.BrowserCacheTTLParam], [zones.BrowserCheckParam],
// [page_rules.PageRuleNewParamsActionsBypassCacheOnCookie],
// [page_rules.PageRuleNewParamsActionsCacheByDeviceType],
// [page_rules.PageRuleNewParamsActionsCacheDeceptionArmor],
// [page_rules.PageRuleNewParamsActionsCacheKey], [zones.CacheLevelParam],
// [page_rules.PageRuleNewParamsActionsCacheOnCookie],
// [page_rules.PageRuleNewParamsActionsDisableApps],
// [page_rules.PageRuleNewParamsActionsDisablePerformance],
// [page_rules.PageRuleNewParamsActionsDisableSecurity],
// [page_rules.PageRuleNewParamsActionsDisableZaraz],
// [page_rules.PageRuleNewParamsActionsEdgeCacheTTL],
// [zones.EmailObfuscationParam],
// [page_rules.PageRuleNewParamsActionsExplicitCacheControl],
// [page_rules.PageRuleNewParamsActionsForwardingURL],
// [page_rules.PageRuleNewParamsActionsHostHeaderOverride],
// [zones.IPGeolocationParam], [zones.MirageParam],
// [zones.OpportunisticEncryptionParam], [zones.OriginErrorPagePassThruParam],
// [zones.PolishParam], [page_rules.PageRuleNewParamsActionsResolveOverride],
// [page_rules.PageRuleNewParamsActionsRespectStrongEtag],
// [zones.ResponseBufferingParam], [zones.RocketLoaderParam],
// [zones.SecurityLevelParam], [zones.SortQueryStringForCacheParam],
// [zones.SSLParam], [zones.TrueClientIPHeaderParam], [zones.WAFParam],
// [PageRuleNewParamsAction].
type PageRuleNewParamsActionUnion interface {
	ImplementsPageRulesPageRuleNewParamsActionUnion()
}

type PageRuleNewParamsActionsBypassCacheOnCookie struct {
	// Bypass cache and fetch resources from the origin server if a regular expression
	// matches against a cookie name present in the request.
	ID param.Field[PageRuleNewParamsActionsBypassCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request. Refer to
	// [Bypass Cache on Cookie setting](https://developers.cloudflare.com/rules/page-rules/reference/additional-reference/#bypass-cache-on-cookie-setting)
	// to learn about limited regular expression support.
	Value param.Field[string] `json:"value"`
}

func (r PageRuleNewParamsActionsBypassCacheOnCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsBypassCacheOnCookie) ImplementsPageRulesPageRuleNewParamsActionUnion() {
}

// Bypass cache and fetch resources from the origin server if a regular expression
// matches against a cookie name present in the request.
type PageRuleNewParamsActionsBypassCacheOnCookieID string

const (
	PageRuleNewParamsActionsBypassCacheOnCookieIDBypassCacheOnCookie PageRuleNewParamsActionsBypassCacheOnCookieID = "bypass_cache_on_cookie"
)

func (r PageRuleNewParamsActionsBypassCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsBypassCacheOnCookieIDBypassCacheOnCookie:
		return true
	}
	return false
}

type PageRuleNewParamsActionsCacheByDeviceType struct {
	// Separate cached content based on the visitor's device type.
	ID param.Field[PageRuleNewParamsActionsCacheByDeviceTypeID] `json:"id"`
	// The status of Cache By Device Type.
	Value param.Field[PageRuleNewParamsActionsCacheByDeviceTypeValue] `json:"value"`
}

func (r PageRuleNewParamsActionsCacheByDeviceType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsCacheByDeviceType) ImplementsPageRulesPageRuleNewParamsActionUnion() {
}

// Separate cached content based on the visitor's device type.
type PageRuleNewParamsActionsCacheByDeviceTypeID string

const (
	PageRuleNewParamsActionsCacheByDeviceTypeIDCacheByDeviceType PageRuleNewParamsActionsCacheByDeviceTypeID = "cache_by_device_type"
)

func (r PageRuleNewParamsActionsCacheByDeviceTypeID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsCacheByDeviceTypeIDCacheByDeviceType:
		return true
	}
	return false
}

// The status of Cache By Device Type.
type PageRuleNewParamsActionsCacheByDeviceTypeValue string

const (
	PageRuleNewParamsActionsCacheByDeviceTypeValueOn  PageRuleNewParamsActionsCacheByDeviceTypeValue = "on"
	PageRuleNewParamsActionsCacheByDeviceTypeValueOff PageRuleNewParamsActionsCacheByDeviceTypeValue = "off"
)

func (r PageRuleNewParamsActionsCacheByDeviceTypeValue) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsCacheByDeviceTypeValueOn, PageRuleNewParamsActionsCacheByDeviceTypeValueOff:
		return true
	}
	return false
}

type PageRuleNewParamsActionsCacheDeceptionArmor struct {
	// Protect from web cache deception attacks while still allowing static assets to
	// be cached. This setting verifies that the URL's extension matches the returned
	// `Content-Type`.
	ID param.Field[PageRuleNewParamsActionsCacheDeceptionArmorID] `json:"id"`
	// The status of Cache Deception Armor.
	Value param.Field[PageRuleNewParamsActionsCacheDeceptionArmorValue] `json:"value"`
}

func (r PageRuleNewParamsActionsCacheDeceptionArmor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsCacheDeceptionArmor) ImplementsPageRulesPageRuleNewParamsActionUnion() {
}

// Protect from web cache deception attacks while still allowing static assets to
// be cached. This setting verifies that the URL's extension matches the returned
// `Content-Type`.
type PageRuleNewParamsActionsCacheDeceptionArmorID string

const (
	PageRuleNewParamsActionsCacheDeceptionArmorIDCacheDeceptionArmor PageRuleNewParamsActionsCacheDeceptionArmorID = "cache_deception_armor"
)

func (r PageRuleNewParamsActionsCacheDeceptionArmorID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsCacheDeceptionArmorIDCacheDeceptionArmor:
		return true
	}
	return false
}

// The status of Cache Deception Armor.
type PageRuleNewParamsActionsCacheDeceptionArmorValue string

const (
	PageRuleNewParamsActionsCacheDeceptionArmorValueOn  PageRuleNewParamsActionsCacheDeceptionArmorValue = "on"
	PageRuleNewParamsActionsCacheDeceptionArmorValueOff PageRuleNewParamsActionsCacheDeceptionArmorValue = "off"
)

func (r PageRuleNewParamsActionsCacheDeceptionArmorValue) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsCacheDeceptionArmorValueOn, PageRuleNewParamsActionsCacheDeceptionArmorValueOff:
		return true
	}
	return false
}

type PageRuleNewParamsActionsCacheKey struct {
	// Control specifically what variables to include when deciding which resources to
	// cache. This allows customers to determine what to cache based on something other
	// than just the URL.
	ID    param.Field[PageRuleNewParamsActionsCacheKeyID]    `json:"id"`
	Value param.Field[PageRuleNewParamsActionsCacheKeyValue] `json:"value"`
}

func (r PageRuleNewParamsActionsCacheKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsCacheKey) ImplementsPageRulesPageRuleNewParamsActionUnion() {}

// Control specifically what variables to include when deciding which resources to
// cache. This allows customers to determine what to cache based on something other
// than just the URL.
type PageRuleNewParamsActionsCacheKeyID string

const (
	PageRuleNewParamsActionsCacheKeyIDCacheKey PageRuleNewParamsActionsCacheKeyID = "cache_key"
)

func (r PageRuleNewParamsActionsCacheKeyID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsCacheKeyIDCacheKey:
		return true
	}
	return false
}

type PageRuleNewParamsActionsCacheKeyValue struct {
	// Controls which cookies appear in the Cache Key.
	Cookie param.Field[PageRuleNewParamsActionsCacheKeyValueCookie] `json:"cookie"`
	// Controls which headers go into the Cache Key. Exactly one of `include` or
	// `exclude` is expected.
	Header param.Field[PageRuleNewParamsActionsCacheKeyValueHeader] `json:"header"`
	// Determines which host header to include in the Cache Key.
	Host param.Field[PageRuleNewParamsActionsCacheKeyValueHost] `json:"host"`
	// Controls which URL query string parameters go into the Cache Key. Exactly one of
	// `include` or `exclude` is expected.
	QueryString param.Field[PageRuleNewParamsActionsCacheKeyValueQueryString] `json:"query_string"`
	// Feature fields to add features about the end-user (client) into the Cache Key.
	User param.Field[PageRuleNewParamsActionsCacheKeyValueUser] `json:"user"`
}

func (r PageRuleNewParamsActionsCacheKeyValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which cookies appear in the Cache Key.
type PageRuleNewParamsActionsCacheKeyValueCookie struct {
	// A list of cookies to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of cookies to include.
	Include param.Field[[]string] `json:"include"`
}

func (r PageRuleNewParamsActionsCacheKeyValueCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which headers go into the Cache Key. Exactly one of `include` or
// `exclude` is expected.
type PageRuleNewParamsActionsCacheKeyValueHeader struct {
	// A list of headers to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of headers to ignore.
	Exclude param.Field[[]string] `json:"exclude"`
	// A list of headers to include.
	Include param.Field[[]string] `json:"include"`
}

func (r PageRuleNewParamsActionsCacheKeyValueHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which host header to include in the Cache Key.
type PageRuleNewParamsActionsCacheKeyValueHost struct {
	// Whether to include the Host header in the HTTP request sent to the origin.
	Resolved param.Field[bool] `json:"resolved"`
}

func (r PageRuleNewParamsActionsCacheKeyValueHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which URL query string parameters go into the Cache Key. Exactly one of
// `include` or `exclude` is expected.
type PageRuleNewParamsActionsCacheKeyValueQueryString struct {
	// Ignore all query string parameters.
	Exclude param.Field[PageRuleNewParamsActionsCacheKeyValueQueryStringExcludeUnion] `json:"exclude"`
	// Include all query string parameters.
	Include param.Field[PageRuleNewParamsActionsCacheKeyValueQueryStringIncludeUnion] `json:"include"`
}

func (r PageRuleNewParamsActionsCacheKeyValueQueryString) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Ignore all query string parameters.
//
// Satisfied by
// [page_rules.PageRuleNewParamsActionsCacheKeyValueQueryStringExcludeString],
// [page_rules.PageRuleNewParamsActionsCacheKeyValueQueryStringExcludeArray].
type PageRuleNewParamsActionsCacheKeyValueQueryStringExcludeUnion interface {
	implementsPageRulesPageRuleNewParamsActionsCacheKeyValueQueryStringExcludeUnion()
}

// Ignore all query string parameters.
type PageRuleNewParamsActionsCacheKeyValueQueryStringExcludeString string

const (
	PageRuleNewParamsActionsCacheKeyValueQueryStringExcludeStringStar PageRuleNewParamsActionsCacheKeyValueQueryStringExcludeString = "*"
)

func (r PageRuleNewParamsActionsCacheKeyValueQueryStringExcludeString) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsCacheKeyValueQueryStringExcludeStringStar:
		return true
	}
	return false
}

func (r PageRuleNewParamsActionsCacheKeyValueQueryStringExcludeString) implementsPageRulesPageRuleNewParamsActionsCacheKeyValueQueryStringExcludeUnion() {
}

type PageRuleNewParamsActionsCacheKeyValueQueryStringExcludeArray []string

func (r PageRuleNewParamsActionsCacheKeyValueQueryStringExcludeArray) implementsPageRulesPageRuleNewParamsActionsCacheKeyValueQueryStringExcludeUnion() {
}

// Include all query string parameters.
//
// Satisfied by
// [page_rules.PageRuleNewParamsActionsCacheKeyValueQueryStringIncludeString],
// [page_rules.PageRuleNewParamsActionsCacheKeyValueQueryStringIncludeArray].
type PageRuleNewParamsActionsCacheKeyValueQueryStringIncludeUnion interface {
	implementsPageRulesPageRuleNewParamsActionsCacheKeyValueQueryStringIncludeUnion()
}

// Include all query string parameters.
type PageRuleNewParamsActionsCacheKeyValueQueryStringIncludeString string

const (
	PageRuleNewParamsActionsCacheKeyValueQueryStringIncludeStringStar PageRuleNewParamsActionsCacheKeyValueQueryStringIncludeString = "*"
)

func (r PageRuleNewParamsActionsCacheKeyValueQueryStringIncludeString) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsCacheKeyValueQueryStringIncludeStringStar:
		return true
	}
	return false
}

func (r PageRuleNewParamsActionsCacheKeyValueQueryStringIncludeString) implementsPageRulesPageRuleNewParamsActionsCacheKeyValueQueryStringIncludeUnion() {
}

type PageRuleNewParamsActionsCacheKeyValueQueryStringIncludeArray []string

func (r PageRuleNewParamsActionsCacheKeyValueQueryStringIncludeArray) implementsPageRulesPageRuleNewParamsActionsCacheKeyValueQueryStringIncludeUnion() {
}

// Feature fields to add features about the end-user (client) into the Cache Key.
type PageRuleNewParamsActionsCacheKeyValueUser struct {
	// Classifies a request as `mobile`, `desktop`, or `tablet` based on the User
	// Agent.
	DeviceType param.Field[bool] `json:"device_type"`
	// Includes the client's country, derived from the IP address.
	Geo param.Field[bool] `json:"geo"`
	// Includes the first language code contained in the `Accept-Language` header sent
	// by the client.
	Lang param.Field[bool] `json:"lang"`
}

func (r PageRuleNewParamsActionsCacheKeyValueUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageRuleNewParamsActionsCacheOnCookie struct {
	// Apply the Cache Everything option (Cache Level setting) based on a regular
	// expression match against a cookie name.
	ID param.Field[PageRuleNewParamsActionsCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request.
	Value param.Field[string] `json:"value"`
}

func (r PageRuleNewParamsActionsCacheOnCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsCacheOnCookie) ImplementsPageRulesPageRuleNewParamsActionUnion() {}

// Apply the Cache Everything option (Cache Level setting) based on a regular
// expression match against a cookie name.
type PageRuleNewParamsActionsCacheOnCookieID string

const (
	PageRuleNewParamsActionsCacheOnCookieIDCacheOnCookie PageRuleNewParamsActionsCacheOnCookieID = "cache_on_cookie"
)

func (r PageRuleNewParamsActionsCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsCacheOnCookieIDCacheOnCookie:
		return true
	}
	return false
}

type PageRuleNewParamsActionsDisableApps struct {
	// Turn off all active
	// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
	// (deprecated).
	ID param.Field[PageRuleNewParamsActionsDisableAppsID] `json:"id"`
}

func (r PageRuleNewParamsActionsDisableApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsDisableApps) ImplementsPageRulesPageRuleNewParamsActionUnion() {}

// Turn off all active
// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
// (deprecated).
type PageRuleNewParamsActionsDisableAppsID string

const (
	PageRuleNewParamsActionsDisableAppsIDDisableApps PageRuleNewParamsActionsDisableAppsID = "disable_apps"
)

func (r PageRuleNewParamsActionsDisableAppsID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsDisableAppsIDDisableApps:
		return true
	}
	return false
}

type PageRuleNewParamsActionsDisablePerformance struct {
	// Turn off
	// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
	// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
	// and [Polish](https://developers.cloudflare.com/images/polish/).
	ID param.Field[PageRuleNewParamsActionsDisablePerformanceID] `json:"id"`
}

func (r PageRuleNewParamsActionsDisablePerformance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsDisablePerformance) ImplementsPageRulesPageRuleNewParamsActionUnion() {
}

// Turn off
// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
// and [Polish](https://developers.cloudflare.com/images/polish/).
type PageRuleNewParamsActionsDisablePerformanceID string

const (
	PageRuleNewParamsActionsDisablePerformanceIDDisablePerformance PageRuleNewParamsActionsDisablePerformanceID = "disable_performance"
)

func (r PageRuleNewParamsActionsDisablePerformanceID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsDisablePerformanceIDDisablePerformance:
		return true
	}
	return false
}

type PageRuleNewParamsActionsDisableSecurity struct {
	// Turn off
	// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
	// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
	// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
	// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
	// and
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	ID param.Field[PageRuleNewParamsActionsDisableSecurityID] `json:"id"`
}

func (r PageRuleNewParamsActionsDisableSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsDisableSecurity) ImplementsPageRulesPageRuleNewParamsActionUnion() {}

// Turn off
// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
// and
// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
type PageRuleNewParamsActionsDisableSecurityID string

const (
	PageRuleNewParamsActionsDisableSecurityIDDisableSecurity PageRuleNewParamsActionsDisableSecurityID = "disable_security"
)

func (r PageRuleNewParamsActionsDisableSecurityID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsDisableSecurityIDDisableSecurity:
		return true
	}
	return false
}

type PageRuleNewParamsActionsDisableZaraz struct {
	// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
	ID param.Field[PageRuleNewParamsActionsDisableZarazID] `json:"id"`
}

func (r PageRuleNewParamsActionsDisableZaraz) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsDisableZaraz) ImplementsPageRulesPageRuleNewParamsActionUnion() {}

// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
type PageRuleNewParamsActionsDisableZarazID string

const (
	PageRuleNewParamsActionsDisableZarazIDDisableZaraz PageRuleNewParamsActionsDisableZarazID = "disable_zaraz"
)

func (r PageRuleNewParamsActionsDisableZarazID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsDisableZarazIDDisableZaraz:
		return true
	}
	return false
}

type PageRuleNewParamsActionsEdgeCacheTTL struct {
	// Specify how long to cache a resource in the Cloudflare global network. _Edge
	// Cache TTL_ is not visible in response headers.
	ID    param.Field[PageRuleNewParamsActionsEdgeCacheTTLID] `json:"id"`
	Value param.Field[int64]                                  `json:"value"`
}

func (r PageRuleNewParamsActionsEdgeCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsEdgeCacheTTL) ImplementsPageRulesPageRuleNewParamsActionUnion() {}

// Specify how long to cache a resource in the Cloudflare global network. _Edge
// Cache TTL_ is not visible in response headers.
type PageRuleNewParamsActionsEdgeCacheTTLID string

const (
	PageRuleNewParamsActionsEdgeCacheTTLIDEdgeCacheTTL PageRuleNewParamsActionsEdgeCacheTTLID = "edge_cache_ttl"
)

func (r PageRuleNewParamsActionsEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

type PageRuleNewParamsActionsExplicitCacheControl struct {
	// Origin Cache Control is enabled by default for Free, Pro, and Business domains
	// and disabled by default for Enterprise domains.
	ID param.Field[PageRuleNewParamsActionsExplicitCacheControlID] `json:"id"`
	// The status of Origin Cache Control.
	Value param.Field[PageRuleNewParamsActionsExplicitCacheControlValue] `json:"value"`
}

func (r PageRuleNewParamsActionsExplicitCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsExplicitCacheControl) ImplementsPageRulesPageRuleNewParamsActionUnion() {
}

// Origin Cache Control is enabled by default for Free, Pro, and Business domains
// and disabled by default for Enterprise domains.
type PageRuleNewParamsActionsExplicitCacheControlID string

const (
	PageRuleNewParamsActionsExplicitCacheControlIDExplicitCacheControl PageRuleNewParamsActionsExplicitCacheControlID = "explicit_cache_control"
)

func (r PageRuleNewParamsActionsExplicitCacheControlID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsExplicitCacheControlIDExplicitCacheControl:
		return true
	}
	return false
}

// The status of Origin Cache Control.
type PageRuleNewParamsActionsExplicitCacheControlValue string

const (
	PageRuleNewParamsActionsExplicitCacheControlValueOn  PageRuleNewParamsActionsExplicitCacheControlValue = "on"
	PageRuleNewParamsActionsExplicitCacheControlValueOff PageRuleNewParamsActionsExplicitCacheControlValue = "off"
)

func (r PageRuleNewParamsActionsExplicitCacheControlValue) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsExplicitCacheControlValueOn, PageRuleNewParamsActionsExplicitCacheControlValueOff:
		return true
	}
	return false
}

type PageRuleNewParamsActionsForwardingURL struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    param.Field[PageRuleNewParamsActionsForwardingURLID]    `json:"id"`
	Value param.Field[PageRuleNewParamsActionsForwardingURLValue] `json:"value"`
}

func (r PageRuleNewParamsActionsForwardingURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsForwardingURL) ImplementsPageRulesPageRuleNewParamsActionUnion() {}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageRuleNewParamsActionsForwardingURLID string

const (
	PageRuleNewParamsActionsForwardingURLIDForwardingURL PageRuleNewParamsActionsForwardingURLID = "forwarding_url"
)

func (r PageRuleNewParamsActionsForwardingURLID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsForwardingURLIDForwardingURL:
		return true
	}
	return false
}

type PageRuleNewParamsActionsForwardingURLValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode param.Field[PageRuleNewParamsActionsForwardingURLValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageRuleNewParamsActionsForwardingURLValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageRuleNewParamsActionsForwardingURLValueStatusCode int64

const (
	PageRuleNewParamsActionsForwardingURLValueStatusCode301 PageRuleNewParamsActionsForwardingURLValueStatusCode = 301
	PageRuleNewParamsActionsForwardingURLValueStatusCode302 PageRuleNewParamsActionsForwardingURLValueStatusCode = 302
)

func (r PageRuleNewParamsActionsForwardingURLValueStatusCode) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsForwardingURLValueStatusCode301, PageRuleNewParamsActionsForwardingURLValueStatusCode302:
		return true
	}
	return false
}

type PageRuleNewParamsActionsHostHeaderOverride struct {
	// Apply a specific host header.
	ID param.Field[PageRuleNewParamsActionsHostHeaderOverrideID] `json:"id"`
	// The hostname to use in the `Host` header
	Value param.Field[string] `json:"value"`
}

func (r PageRuleNewParamsActionsHostHeaderOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsHostHeaderOverride) ImplementsPageRulesPageRuleNewParamsActionUnion() {
}

// Apply a specific host header.
type PageRuleNewParamsActionsHostHeaderOverrideID string

const (
	PageRuleNewParamsActionsHostHeaderOverrideIDHostHeaderOverride PageRuleNewParamsActionsHostHeaderOverrideID = "host_header_override"
)

func (r PageRuleNewParamsActionsHostHeaderOverrideID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsHostHeaderOverrideIDHostHeaderOverride:
		return true
	}
	return false
}

type PageRuleNewParamsActionsResolveOverride struct {
	// Change the origin address to the value specified in this setting.
	ID param.Field[PageRuleNewParamsActionsResolveOverrideID] `json:"id"`
	// The origin address you want to override with.
	Value param.Field[string] `json:"value"`
}

func (r PageRuleNewParamsActionsResolveOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsResolveOverride) ImplementsPageRulesPageRuleNewParamsActionUnion() {}

// Change the origin address to the value specified in this setting.
type PageRuleNewParamsActionsResolveOverrideID string

const (
	PageRuleNewParamsActionsResolveOverrideIDResolveOverride PageRuleNewParamsActionsResolveOverrideID = "resolve_override"
)

func (r PageRuleNewParamsActionsResolveOverrideID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsResolveOverrideIDResolveOverride:
		return true
	}
	return false
}

type PageRuleNewParamsActionsRespectStrongEtag struct {
	// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
	// the origin server.
	ID param.Field[PageRuleNewParamsActionsRespectStrongEtagID] `json:"id"`
	// The status of Respect Strong ETags
	Value param.Field[PageRuleNewParamsActionsRespectStrongEtagValue] `json:"value"`
}

func (r PageRuleNewParamsActionsRespectStrongEtag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleNewParamsActionsRespectStrongEtag) ImplementsPageRulesPageRuleNewParamsActionUnion() {
}

// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
// the origin server.
type PageRuleNewParamsActionsRespectStrongEtagID string

const (
	PageRuleNewParamsActionsRespectStrongEtagIDRespectStrongEtag PageRuleNewParamsActionsRespectStrongEtagID = "respect_strong_etag"
)

func (r PageRuleNewParamsActionsRespectStrongEtagID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsRespectStrongEtagIDRespectStrongEtag:
		return true
	}
	return false
}

// The status of Respect Strong ETags
type PageRuleNewParamsActionsRespectStrongEtagValue string

const (
	PageRuleNewParamsActionsRespectStrongEtagValueOn  PageRuleNewParamsActionsRespectStrongEtagValue = "on"
	PageRuleNewParamsActionsRespectStrongEtagValueOff PageRuleNewParamsActionsRespectStrongEtagValue = "off"
)

func (r PageRuleNewParamsActionsRespectStrongEtagValue) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsRespectStrongEtagValueOn, PageRuleNewParamsActionsRespectStrongEtagValueOff:
		return true
	}
	return false
}

// If enabled, any ` http://“ URL is converted to  `https://` through a 301
// redirect.
type PageRuleNewParamsActionsID string

const (
	PageRuleNewParamsActionsIDAlwaysUseHTTPS          PageRuleNewParamsActionsID = "always_use_https"
	PageRuleNewParamsActionsIDAutomaticHTTPSRewrites  PageRuleNewParamsActionsID = "automatic_https_rewrites"
	PageRuleNewParamsActionsIDBrowserCacheTTL         PageRuleNewParamsActionsID = "browser_cache_ttl"
	PageRuleNewParamsActionsIDBrowserCheck            PageRuleNewParamsActionsID = "browser_check"
	PageRuleNewParamsActionsIDBypassCacheOnCookie     PageRuleNewParamsActionsID = "bypass_cache_on_cookie"
	PageRuleNewParamsActionsIDCacheByDeviceType       PageRuleNewParamsActionsID = "cache_by_device_type"
	PageRuleNewParamsActionsIDCacheDeceptionArmor     PageRuleNewParamsActionsID = "cache_deception_armor"
	PageRuleNewParamsActionsIDCacheKey                PageRuleNewParamsActionsID = "cache_key"
	PageRuleNewParamsActionsIDCacheLevel              PageRuleNewParamsActionsID = "cache_level"
	PageRuleNewParamsActionsIDCacheOnCookie           PageRuleNewParamsActionsID = "cache_on_cookie"
	PageRuleNewParamsActionsIDDisableApps             PageRuleNewParamsActionsID = "disable_apps"
	PageRuleNewParamsActionsIDDisablePerformance      PageRuleNewParamsActionsID = "disable_performance"
	PageRuleNewParamsActionsIDDisableSecurity         PageRuleNewParamsActionsID = "disable_security"
	PageRuleNewParamsActionsIDDisableZaraz            PageRuleNewParamsActionsID = "disable_zaraz"
	PageRuleNewParamsActionsIDEdgeCacheTTL            PageRuleNewParamsActionsID = "edge_cache_ttl"
	PageRuleNewParamsActionsIDEmailObfuscation        PageRuleNewParamsActionsID = "email_obfuscation"
	PageRuleNewParamsActionsIDExplicitCacheControl    PageRuleNewParamsActionsID = "explicit_cache_control"
	PageRuleNewParamsActionsIDForwardingURL           PageRuleNewParamsActionsID = "forwarding_url"
	PageRuleNewParamsActionsIDHostHeaderOverride      PageRuleNewParamsActionsID = "host_header_override"
	PageRuleNewParamsActionsIDIPGeolocation           PageRuleNewParamsActionsID = "ip_geolocation"
	PageRuleNewParamsActionsIDMirage                  PageRuleNewParamsActionsID = "mirage"
	PageRuleNewParamsActionsIDOpportunisticEncryption PageRuleNewParamsActionsID = "opportunistic_encryption"
	PageRuleNewParamsActionsIDOriginErrorPagePassThru PageRuleNewParamsActionsID = "origin_error_page_pass_thru"
	PageRuleNewParamsActionsIDPolish                  PageRuleNewParamsActionsID = "polish"
	PageRuleNewParamsActionsIDResolveOverride         PageRuleNewParamsActionsID = "resolve_override"
	PageRuleNewParamsActionsIDRespectStrongEtag       PageRuleNewParamsActionsID = "respect_strong_etag"
	PageRuleNewParamsActionsIDResponseBuffering       PageRuleNewParamsActionsID = "response_buffering"
	PageRuleNewParamsActionsIDRocketLoader            PageRuleNewParamsActionsID = "rocket_loader"
	PageRuleNewParamsActionsIDSecurityLevel           PageRuleNewParamsActionsID = "security_level"
	PageRuleNewParamsActionsIDSortQueryStringForCache PageRuleNewParamsActionsID = "sort_query_string_for_cache"
	PageRuleNewParamsActionsIDSSL                     PageRuleNewParamsActionsID = "ssl"
	PageRuleNewParamsActionsIDTrueClientIPHeader      PageRuleNewParamsActionsID = "true_client_ip_header"
	PageRuleNewParamsActionsIDWAF                     PageRuleNewParamsActionsID = "waf"
)

func (r PageRuleNewParamsActionsID) IsKnown() bool {
	switch r {
	case PageRuleNewParamsActionsIDAlwaysUseHTTPS, PageRuleNewParamsActionsIDAutomaticHTTPSRewrites, PageRuleNewParamsActionsIDBrowserCacheTTL, PageRuleNewParamsActionsIDBrowserCheck, PageRuleNewParamsActionsIDBypassCacheOnCookie, PageRuleNewParamsActionsIDCacheByDeviceType, PageRuleNewParamsActionsIDCacheDeceptionArmor, PageRuleNewParamsActionsIDCacheKey, PageRuleNewParamsActionsIDCacheLevel, PageRuleNewParamsActionsIDCacheOnCookie, PageRuleNewParamsActionsIDDisableApps, PageRuleNewParamsActionsIDDisablePerformance, PageRuleNewParamsActionsIDDisableSecurity, PageRuleNewParamsActionsIDDisableZaraz, PageRuleNewParamsActionsIDEdgeCacheTTL, PageRuleNewParamsActionsIDEmailObfuscation, PageRuleNewParamsActionsIDExplicitCacheControl, PageRuleNewParamsActionsIDForwardingURL, PageRuleNewParamsActionsIDHostHeaderOverride, PageRuleNewParamsActionsIDIPGeolocation, PageRuleNewParamsActionsIDMirage, PageRuleNewParamsActionsIDOpportunisticEncryption, PageRuleNewParamsActionsIDOriginErrorPagePassThru, PageRuleNewParamsActionsIDPolish, PageRuleNewParamsActionsIDResolveOverride, PageRuleNewParamsActionsIDRespectStrongEtag, PageRuleNewParamsActionsIDResponseBuffering, PageRuleNewParamsActionsIDRocketLoader, PageRuleNewParamsActionsIDSecurityLevel, PageRuleNewParamsActionsIDSortQueryStringForCache, PageRuleNewParamsActionsIDSSL, PageRuleNewParamsActionsIDTrueClientIPHeader, PageRuleNewParamsActionsIDWAF:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageRuleNewParamsStatus string

const (
	PageRuleNewParamsStatusActive   PageRuleNewParamsStatus = "active"
	PageRuleNewParamsStatusDisabled PageRuleNewParamsStatus = "disabled"
)

func (r PageRuleNewParamsStatus) IsKnown() bool {
	switch r {
	case PageRuleNewParamsStatusActive, PageRuleNewParamsStatusDisabled:
		return true
	}
	return false
}

type PageRuleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageRuleNewResponseEnvelopeSuccess `json:"success,required"`
	Result  PageRule                           `json:"result"`
	JSON    pageRuleNewResponseEnvelopeJSON    `json:"-"`
}

// pageRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageRuleNewResponseEnvelope]
type pageRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageRuleNewResponseEnvelopeSuccess bool

const (
	PageRuleNewResponseEnvelopeSuccessTrue PageRuleNewResponseEnvelopeSuccess = true
)

func (r PageRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageRuleUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]PageRuleUpdateParamsActionUnion] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargetParam] `json:"targets,required"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageRuleUpdateParamsStatus] `json:"status"`
}

func (r PageRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageRuleUpdateParamsAction struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID    param.Field[PageRuleUpdateParamsActionsID] `json:"id"`
	Value param.Field[interface{}]                   `json:"value"`
}

func (r PageRuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsAction) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {}

// Satisfied by [zones.AlwaysUseHTTPSParam], [zones.AutomaticHTTPSRewritesParam],
// [zones.BrowserCacheTTLParam], [zones.BrowserCheckParam],
// [page_rules.PageRuleUpdateParamsActionsBypassCacheOnCookie],
// [page_rules.PageRuleUpdateParamsActionsCacheByDeviceType],
// [page_rules.PageRuleUpdateParamsActionsCacheDeceptionArmor],
// [page_rules.PageRuleUpdateParamsActionsCacheKey], [zones.CacheLevelParam],
// [page_rules.PageRuleUpdateParamsActionsCacheOnCookie],
// [page_rules.PageRuleUpdateParamsActionsDisableApps],
// [page_rules.PageRuleUpdateParamsActionsDisablePerformance],
// [page_rules.PageRuleUpdateParamsActionsDisableSecurity],
// [page_rules.PageRuleUpdateParamsActionsDisableZaraz],
// [page_rules.PageRuleUpdateParamsActionsEdgeCacheTTL],
// [zones.EmailObfuscationParam],
// [page_rules.PageRuleUpdateParamsActionsExplicitCacheControl],
// [page_rules.PageRuleUpdateParamsActionsForwardingURL],
// [page_rules.PageRuleUpdateParamsActionsHostHeaderOverride],
// [zones.IPGeolocationParam], [zones.MirageParam],
// [zones.OpportunisticEncryptionParam], [zones.OriginErrorPagePassThruParam],
// [zones.PolishParam], [page_rules.PageRuleUpdateParamsActionsResolveOverride],
// [page_rules.PageRuleUpdateParamsActionsRespectStrongEtag],
// [zones.ResponseBufferingParam], [zones.RocketLoaderParam],
// [zones.SecurityLevelParam], [zones.SortQueryStringForCacheParam],
// [zones.SSLParam], [zones.TrueClientIPHeaderParam], [zones.WAFParam],
// [PageRuleUpdateParamsAction].
type PageRuleUpdateParamsActionUnion interface {
	ImplementsPageRulesPageRuleUpdateParamsActionUnion()
}

type PageRuleUpdateParamsActionsBypassCacheOnCookie struct {
	// Bypass cache and fetch resources from the origin server if a regular expression
	// matches against a cookie name present in the request.
	ID param.Field[PageRuleUpdateParamsActionsBypassCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request. Refer to
	// [Bypass Cache on Cookie setting](https://developers.cloudflare.com/rules/page-rules/reference/additional-reference/#bypass-cache-on-cookie-setting)
	// to learn about limited regular expression support.
	Value param.Field[string] `json:"value"`
}

func (r PageRuleUpdateParamsActionsBypassCacheOnCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsBypassCacheOnCookie) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Bypass cache and fetch resources from the origin server if a regular expression
// matches against a cookie name present in the request.
type PageRuleUpdateParamsActionsBypassCacheOnCookieID string

const (
	PageRuleUpdateParamsActionsBypassCacheOnCookieIDBypassCacheOnCookie PageRuleUpdateParamsActionsBypassCacheOnCookieID = "bypass_cache_on_cookie"
)

func (r PageRuleUpdateParamsActionsBypassCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsBypassCacheOnCookieIDBypassCacheOnCookie:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsCacheByDeviceType struct {
	// Separate cached content based on the visitor's device type.
	ID param.Field[PageRuleUpdateParamsActionsCacheByDeviceTypeID] `json:"id"`
	// The status of Cache By Device Type.
	Value param.Field[PageRuleUpdateParamsActionsCacheByDeviceTypeValue] `json:"value"`
}

func (r PageRuleUpdateParamsActionsCacheByDeviceType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsCacheByDeviceType) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Separate cached content based on the visitor's device type.
type PageRuleUpdateParamsActionsCacheByDeviceTypeID string

const (
	PageRuleUpdateParamsActionsCacheByDeviceTypeIDCacheByDeviceType PageRuleUpdateParamsActionsCacheByDeviceTypeID = "cache_by_device_type"
)

func (r PageRuleUpdateParamsActionsCacheByDeviceTypeID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsCacheByDeviceTypeIDCacheByDeviceType:
		return true
	}
	return false
}

// The status of Cache By Device Type.
type PageRuleUpdateParamsActionsCacheByDeviceTypeValue string

const (
	PageRuleUpdateParamsActionsCacheByDeviceTypeValueOn  PageRuleUpdateParamsActionsCacheByDeviceTypeValue = "on"
	PageRuleUpdateParamsActionsCacheByDeviceTypeValueOff PageRuleUpdateParamsActionsCacheByDeviceTypeValue = "off"
)

func (r PageRuleUpdateParamsActionsCacheByDeviceTypeValue) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsCacheByDeviceTypeValueOn, PageRuleUpdateParamsActionsCacheByDeviceTypeValueOff:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsCacheDeceptionArmor struct {
	// Protect from web cache deception attacks while still allowing static assets to
	// be cached. This setting verifies that the URL's extension matches the returned
	// `Content-Type`.
	ID param.Field[PageRuleUpdateParamsActionsCacheDeceptionArmorID] `json:"id"`
	// The status of Cache Deception Armor.
	Value param.Field[PageRuleUpdateParamsActionsCacheDeceptionArmorValue] `json:"value"`
}

func (r PageRuleUpdateParamsActionsCacheDeceptionArmor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsCacheDeceptionArmor) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Protect from web cache deception attacks while still allowing static assets to
// be cached. This setting verifies that the URL's extension matches the returned
// `Content-Type`.
type PageRuleUpdateParamsActionsCacheDeceptionArmorID string

const (
	PageRuleUpdateParamsActionsCacheDeceptionArmorIDCacheDeceptionArmor PageRuleUpdateParamsActionsCacheDeceptionArmorID = "cache_deception_armor"
)

func (r PageRuleUpdateParamsActionsCacheDeceptionArmorID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsCacheDeceptionArmorIDCacheDeceptionArmor:
		return true
	}
	return false
}

// The status of Cache Deception Armor.
type PageRuleUpdateParamsActionsCacheDeceptionArmorValue string

const (
	PageRuleUpdateParamsActionsCacheDeceptionArmorValueOn  PageRuleUpdateParamsActionsCacheDeceptionArmorValue = "on"
	PageRuleUpdateParamsActionsCacheDeceptionArmorValueOff PageRuleUpdateParamsActionsCacheDeceptionArmorValue = "off"
)

func (r PageRuleUpdateParamsActionsCacheDeceptionArmorValue) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsCacheDeceptionArmorValueOn, PageRuleUpdateParamsActionsCacheDeceptionArmorValueOff:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsCacheKey struct {
	// Control specifically what variables to include when deciding which resources to
	// cache. This allows customers to determine what to cache based on something other
	// than just the URL.
	ID    param.Field[PageRuleUpdateParamsActionsCacheKeyID]    `json:"id"`
	Value param.Field[PageRuleUpdateParamsActionsCacheKeyValue] `json:"value"`
}

func (r PageRuleUpdateParamsActionsCacheKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsCacheKey) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {}

// Control specifically what variables to include when deciding which resources to
// cache. This allows customers to determine what to cache based on something other
// than just the URL.
type PageRuleUpdateParamsActionsCacheKeyID string

const (
	PageRuleUpdateParamsActionsCacheKeyIDCacheKey PageRuleUpdateParamsActionsCacheKeyID = "cache_key"
)

func (r PageRuleUpdateParamsActionsCacheKeyID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsCacheKeyIDCacheKey:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsCacheKeyValue struct {
	// Controls which cookies appear in the Cache Key.
	Cookie param.Field[PageRuleUpdateParamsActionsCacheKeyValueCookie] `json:"cookie"`
	// Controls which headers go into the Cache Key. Exactly one of `include` or
	// `exclude` is expected.
	Header param.Field[PageRuleUpdateParamsActionsCacheKeyValueHeader] `json:"header"`
	// Determines which host header to include in the Cache Key.
	Host param.Field[PageRuleUpdateParamsActionsCacheKeyValueHost] `json:"host"`
	// Controls which URL query string parameters go into the Cache Key. Exactly one of
	// `include` or `exclude` is expected.
	QueryString param.Field[PageRuleUpdateParamsActionsCacheKeyValueQueryString] `json:"query_string"`
	// Feature fields to add features about the end-user (client) into the Cache Key.
	User param.Field[PageRuleUpdateParamsActionsCacheKeyValueUser] `json:"user"`
}

func (r PageRuleUpdateParamsActionsCacheKeyValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which cookies appear in the Cache Key.
type PageRuleUpdateParamsActionsCacheKeyValueCookie struct {
	// A list of cookies to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of cookies to include.
	Include param.Field[[]string] `json:"include"`
}

func (r PageRuleUpdateParamsActionsCacheKeyValueCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which headers go into the Cache Key. Exactly one of `include` or
// `exclude` is expected.
type PageRuleUpdateParamsActionsCacheKeyValueHeader struct {
	// A list of headers to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of headers to ignore.
	Exclude param.Field[[]string] `json:"exclude"`
	// A list of headers to include.
	Include param.Field[[]string] `json:"include"`
}

func (r PageRuleUpdateParamsActionsCacheKeyValueHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which host header to include in the Cache Key.
type PageRuleUpdateParamsActionsCacheKeyValueHost struct {
	// Whether to include the Host header in the HTTP request sent to the origin.
	Resolved param.Field[bool] `json:"resolved"`
}

func (r PageRuleUpdateParamsActionsCacheKeyValueHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which URL query string parameters go into the Cache Key. Exactly one of
// `include` or `exclude` is expected.
type PageRuleUpdateParamsActionsCacheKeyValueQueryString struct {
	// Ignore all query string parameters.
	Exclude param.Field[PageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeUnion] `json:"exclude"`
	// Include all query string parameters.
	Include param.Field[PageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeUnion] `json:"include"`
}

func (r PageRuleUpdateParamsActionsCacheKeyValueQueryString) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Ignore all query string parameters.
//
// Satisfied by
// [page_rules.PageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeString],
// [page_rules.PageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeArray].
type PageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeUnion interface {
	implementsPageRulesPageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeUnion()
}

// Ignore all query string parameters.
type PageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeString string

const (
	PageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeStringStar PageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeString = "*"
)

func (r PageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeString) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeStringStar:
		return true
	}
	return false
}

func (r PageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeString) implementsPageRulesPageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeUnion() {
}

type PageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeArray []string

func (r PageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeArray) implementsPageRulesPageRuleUpdateParamsActionsCacheKeyValueQueryStringExcludeUnion() {
}

// Include all query string parameters.
//
// Satisfied by
// [page_rules.PageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeString],
// [page_rules.PageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeArray].
type PageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeUnion interface {
	implementsPageRulesPageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeUnion()
}

// Include all query string parameters.
type PageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeString string

const (
	PageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeStringStar PageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeString = "*"
)

func (r PageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeString) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeStringStar:
		return true
	}
	return false
}

func (r PageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeString) implementsPageRulesPageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeUnion() {
}

type PageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeArray []string

func (r PageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeArray) implementsPageRulesPageRuleUpdateParamsActionsCacheKeyValueQueryStringIncludeUnion() {
}

// Feature fields to add features about the end-user (client) into the Cache Key.
type PageRuleUpdateParamsActionsCacheKeyValueUser struct {
	// Classifies a request as `mobile`, `desktop`, or `tablet` based on the User
	// Agent.
	DeviceType param.Field[bool] `json:"device_type"`
	// Includes the client's country, derived from the IP address.
	Geo param.Field[bool] `json:"geo"`
	// Includes the first language code contained in the `Accept-Language` header sent
	// by the client.
	Lang param.Field[bool] `json:"lang"`
}

func (r PageRuleUpdateParamsActionsCacheKeyValueUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageRuleUpdateParamsActionsCacheOnCookie struct {
	// Apply the Cache Everything option (Cache Level setting) based on a regular
	// expression match against a cookie name.
	ID param.Field[PageRuleUpdateParamsActionsCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request.
	Value param.Field[string] `json:"value"`
}

func (r PageRuleUpdateParamsActionsCacheOnCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsCacheOnCookie) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Apply the Cache Everything option (Cache Level setting) based on a regular
// expression match against a cookie name.
type PageRuleUpdateParamsActionsCacheOnCookieID string

const (
	PageRuleUpdateParamsActionsCacheOnCookieIDCacheOnCookie PageRuleUpdateParamsActionsCacheOnCookieID = "cache_on_cookie"
)

func (r PageRuleUpdateParamsActionsCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsCacheOnCookieIDCacheOnCookie:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsDisableApps struct {
	// Turn off all active
	// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
	// (deprecated).
	ID param.Field[PageRuleUpdateParamsActionsDisableAppsID] `json:"id"`
}

func (r PageRuleUpdateParamsActionsDisableApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsDisableApps) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Turn off all active
// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
// (deprecated).
type PageRuleUpdateParamsActionsDisableAppsID string

const (
	PageRuleUpdateParamsActionsDisableAppsIDDisableApps PageRuleUpdateParamsActionsDisableAppsID = "disable_apps"
)

func (r PageRuleUpdateParamsActionsDisableAppsID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsDisableAppsIDDisableApps:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsDisablePerformance struct {
	// Turn off
	// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
	// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
	// and [Polish](https://developers.cloudflare.com/images/polish/).
	ID param.Field[PageRuleUpdateParamsActionsDisablePerformanceID] `json:"id"`
}

func (r PageRuleUpdateParamsActionsDisablePerformance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsDisablePerformance) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Turn off
// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
// and [Polish](https://developers.cloudflare.com/images/polish/).
type PageRuleUpdateParamsActionsDisablePerformanceID string

const (
	PageRuleUpdateParamsActionsDisablePerformanceIDDisablePerformance PageRuleUpdateParamsActionsDisablePerformanceID = "disable_performance"
)

func (r PageRuleUpdateParamsActionsDisablePerformanceID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsDisablePerformanceIDDisablePerformance:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsDisableSecurity struct {
	// Turn off
	// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
	// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
	// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
	// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
	// and
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	ID param.Field[PageRuleUpdateParamsActionsDisableSecurityID] `json:"id"`
}

func (r PageRuleUpdateParamsActionsDisableSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsDisableSecurity) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Turn off
// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
// and
// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
type PageRuleUpdateParamsActionsDisableSecurityID string

const (
	PageRuleUpdateParamsActionsDisableSecurityIDDisableSecurity PageRuleUpdateParamsActionsDisableSecurityID = "disable_security"
)

func (r PageRuleUpdateParamsActionsDisableSecurityID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsDisableSecurityIDDisableSecurity:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsDisableZaraz struct {
	// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
	ID param.Field[PageRuleUpdateParamsActionsDisableZarazID] `json:"id"`
}

func (r PageRuleUpdateParamsActionsDisableZaraz) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsDisableZaraz) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
type PageRuleUpdateParamsActionsDisableZarazID string

const (
	PageRuleUpdateParamsActionsDisableZarazIDDisableZaraz PageRuleUpdateParamsActionsDisableZarazID = "disable_zaraz"
)

func (r PageRuleUpdateParamsActionsDisableZarazID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsDisableZarazIDDisableZaraz:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsEdgeCacheTTL struct {
	// Specify how long to cache a resource in the Cloudflare global network. _Edge
	// Cache TTL_ is not visible in response headers.
	ID    param.Field[PageRuleUpdateParamsActionsEdgeCacheTTLID] `json:"id"`
	Value param.Field[int64]                                     `json:"value"`
}

func (r PageRuleUpdateParamsActionsEdgeCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsEdgeCacheTTL) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Specify how long to cache a resource in the Cloudflare global network. _Edge
// Cache TTL_ is not visible in response headers.
type PageRuleUpdateParamsActionsEdgeCacheTTLID string

const (
	PageRuleUpdateParamsActionsEdgeCacheTTLIDEdgeCacheTTL PageRuleUpdateParamsActionsEdgeCacheTTLID = "edge_cache_ttl"
)

func (r PageRuleUpdateParamsActionsEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsExplicitCacheControl struct {
	// Origin Cache Control is enabled by default for Free, Pro, and Business domains
	// and disabled by default for Enterprise domains.
	ID param.Field[PageRuleUpdateParamsActionsExplicitCacheControlID] `json:"id"`
	// The status of Origin Cache Control.
	Value param.Field[PageRuleUpdateParamsActionsExplicitCacheControlValue] `json:"value"`
}

func (r PageRuleUpdateParamsActionsExplicitCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsExplicitCacheControl) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Origin Cache Control is enabled by default for Free, Pro, and Business domains
// and disabled by default for Enterprise domains.
type PageRuleUpdateParamsActionsExplicitCacheControlID string

const (
	PageRuleUpdateParamsActionsExplicitCacheControlIDExplicitCacheControl PageRuleUpdateParamsActionsExplicitCacheControlID = "explicit_cache_control"
)

func (r PageRuleUpdateParamsActionsExplicitCacheControlID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsExplicitCacheControlIDExplicitCacheControl:
		return true
	}
	return false
}

// The status of Origin Cache Control.
type PageRuleUpdateParamsActionsExplicitCacheControlValue string

const (
	PageRuleUpdateParamsActionsExplicitCacheControlValueOn  PageRuleUpdateParamsActionsExplicitCacheControlValue = "on"
	PageRuleUpdateParamsActionsExplicitCacheControlValueOff PageRuleUpdateParamsActionsExplicitCacheControlValue = "off"
)

func (r PageRuleUpdateParamsActionsExplicitCacheControlValue) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsExplicitCacheControlValueOn, PageRuleUpdateParamsActionsExplicitCacheControlValueOff:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsForwardingURL struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    param.Field[PageRuleUpdateParamsActionsForwardingURLID]    `json:"id"`
	Value param.Field[PageRuleUpdateParamsActionsForwardingURLValue] `json:"value"`
}

func (r PageRuleUpdateParamsActionsForwardingURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsForwardingURL) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageRuleUpdateParamsActionsForwardingURLID string

const (
	PageRuleUpdateParamsActionsForwardingURLIDForwardingURL PageRuleUpdateParamsActionsForwardingURLID = "forwarding_url"
)

func (r PageRuleUpdateParamsActionsForwardingURLID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsForwardingURLIDForwardingURL:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsForwardingURLValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode param.Field[PageRuleUpdateParamsActionsForwardingURLValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageRuleUpdateParamsActionsForwardingURLValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageRuleUpdateParamsActionsForwardingURLValueStatusCode int64

const (
	PageRuleUpdateParamsActionsForwardingURLValueStatusCode301 PageRuleUpdateParamsActionsForwardingURLValueStatusCode = 301
	PageRuleUpdateParamsActionsForwardingURLValueStatusCode302 PageRuleUpdateParamsActionsForwardingURLValueStatusCode = 302
)

func (r PageRuleUpdateParamsActionsForwardingURLValueStatusCode) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsForwardingURLValueStatusCode301, PageRuleUpdateParamsActionsForwardingURLValueStatusCode302:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsHostHeaderOverride struct {
	// Apply a specific host header.
	ID param.Field[PageRuleUpdateParamsActionsHostHeaderOverrideID] `json:"id"`
	// The hostname to use in the `Host` header
	Value param.Field[string] `json:"value"`
}

func (r PageRuleUpdateParamsActionsHostHeaderOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsHostHeaderOverride) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Apply a specific host header.
type PageRuleUpdateParamsActionsHostHeaderOverrideID string

const (
	PageRuleUpdateParamsActionsHostHeaderOverrideIDHostHeaderOverride PageRuleUpdateParamsActionsHostHeaderOverrideID = "host_header_override"
)

func (r PageRuleUpdateParamsActionsHostHeaderOverrideID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsHostHeaderOverrideIDHostHeaderOverride:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsResolveOverride struct {
	// Change the origin address to the value specified in this setting.
	ID param.Field[PageRuleUpdateParamsActionsResolveOverrideID] `json:"id"`
	// The origin address you want to override with.
	Value param.Field[string] `json:"value"`
}

func (r PageRuleUpdateParamsActionsResolveOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsResolveOverride) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Change the origin address to the value specified in this setting.
type PageRuleUpdateParamsActionsResolveOverrideID string

const (
	PageRuleUpdateParamsActionsResolveOverrideIDResolveOverride PageRuleUpdateParamsActionsResolveOverrideID = "resolve_override"
)

func (r PageRuleUpdateParamsActionsResolveOverrideID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsResolveOverrideIDResolveOverride:
		return true
	}
	return false
}

type PageRuleUpdateParamsActionsRespectStrongEtag struct {
	// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
	// the origin server.
	ID param.Field[PageRuleUpdateParamsActionsRespectStrongEtagID] `json:"id"`
	// The status of Respect Strong ETags
	Value param.Field[PageRuleUpdateParamsActionsRespectStrongEtagValue] `json:"value"`
}

func (r PageRuleUpdateParamsActionsRespectStrongEtag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleUpdateParamsActionsRespectStrongEtag) ImplementsPageRulesPageRuleUpdateParamsActionUnion() {
}

// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
// the origin server.
type PageRuleUpdateParamsActionsRespectStrongEtagID string

const (
	PageRuleUpdateParamsActionsRespectStrongEtagIDRespectStrongEtag PageRuleUpdateParamsActionsRespectStrongEtagID = "respect_strong_etag"
)

func (r PageRuleUpdateParamsActionsRespectStrongEtagID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsRespectStrongEtagIDRespectStrongEtag:
		return true
	}
	return false
}

// The status of Respect Strong ETags
type PageRuleUpdateParamsActionsRespectStrongEtagValue string

const (
	PageRuleUpdateParamsActionsRespectStrongEtagValueOn  PageRuleUpdateParamsActionsRespectStrongEtagValue = "on"
	PageRuleUpdateParamsActionsRespectStrongEtagValueOff PageRuleUpdateParamsActionsRespectStrongEtagValue = "off"
)

func (r PageRuleUpdateParamsActionsRespectStrongEtagValue) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsRespectStrongEtagValueOn, PageRuleUpdateParamsActionsRespectStrongEtagValueOff:
		return true
	}
	return false
}

// If enabled, any ` http://“ URL is converted to  `https://` through a 301
// redirect.
type PageRuleUpdateParamsActionsID string

const (
	PageRuleUpdateParamsActionsIDAlwaysUseHTTPS          PageRuleUpdateParamsActionsID = "always_use_https"
	PageRuleUpdateParamsActionsIDAutomaticHTTPSRewrites  PageRuleUpdateParamsActionsID = "automatic_https_rewrites"
	PageRuleUpdateParamsActionsIDBrowserCacheTTL         PageRuleUpdateParamsActionsID = "browser_cache_ttl"
	PageRuleUpdateParamsActionsIDBrowserCheck            PageRuleUpdateParamsActionsID = "browser_check"
	PageRuleUpdateParamsActionsIDBypassCacheOnCookie     PageRuleUpdateParamsActionsID = "bypass_cache_on_cookie"
	PageRuleUpdateParamsActionsIDCacheByDeviceType       PageRuleUpdateParamsActionsID = "cache_by_device_type"
	PageRuleUpdateParamsActionsIDCacheDeceptionArmor     PageRuleUpdateParamsActionsID = "cache_deception_armor"
	PageRuleUpdateParamsActionsIDCacheKey                PageRuleUpdateParamsActionsID = "cache_key"
	PageRuleUpdateParamsActionsIDCacheLevel              PageRuleUpdateParamsActionsID = "cache_level"
	PageRuleUpdateParamsActionsIDCacheOnCookie           PageRuleUpdateParamsActionsID = "cache_on_cookie"
	PageRuleUpdateParamsActionsIDDisableApps             PageRuleUpdateParamsActionsID = "disable_apps"
	PageRuleUpdateParamsActionsIDDisablePerformance      PageRuleUpdateParamsActionsID = "disable_performance"
	PageRuleUpdateParamsActionsIDDisableSecurity         PageRuleUpdateParamsActionsID = "disable_security"
	PageRuleUpdateParamsActionsIDDisableZaraz            PageRuleUpdateParamsActionsID = "disable_zaraz"
	PageRuleUpdateParamsActionsIDEdgeCacheTTL            PageRuleUpdateParamsActionsID = "edge_cache_ttl"
	PageRuleUpdateParamsActionsIDEmailObfuscation        PageRuleUpdateParamsActionsID = "email_obfuscation"
	PageRuleUpdateParamsActionsIDExplicitCacheControl    PageRuleUpdateParamsActionsID = "explicit_cache_control"
	PageRuleUpdateParamsActionsIDForwardingURL           PageRuleUpdateParamsActionsID = "forwarding_url"
	PageRuleUpdateParamsActionsIDHostHeaderOverride      PageRuleUpdateParamsActionsID = "host_header_override"
	PageRuleUpdateParamsActionsIDIPGeolocation           PageRuleUpdateParamsActionsID = "ip_geolocation"
	PageRuleUpdateParamsActionsIDMirage                  PageRuleUpdateParamsActionsID = "mirage"
	PageRuleUpdateParamsActionsIDOpportunisticEncryption PageRuleUpdateParamsActionsID = "opportunistic_encryption"
	PageRuleUpdateParamsActionsIDOriginErrorPagePassThru PageRuleUpdateParamsActionsID = "origin_error_page_pass_thru"
	PageRuleUpdateParamsActionsIDPolish                  PageRuleUpdateParamsActionsID = "polish"
	PageRuleUpdateParamsActionsIDResolveOverride         PageRuleUpdateParamsActionsID = "resolve_override"
	PageRuleUpdateParamsActionsIDRespectStrongEtag       PageRuleUpdateParamsActionsID = "respect_strong_etag"
	PageRuleUpdateParamsActionsIDResponseBuffering       PageRuleUpdateParamsActionsID = "response_buffering"
	PageRuleUpdateParamsActionsIDRocketLoader            PageRuleUpdateParamsActionsID = "rocket_loader"
	PageRuleUpdateParamsActionsIDSecurityLevel           PageRuleUpdateParamsActionsID = "security_level"
	PageRuleUpdateParamsActionsIDSortQueryStringForCache PageRuleUpdateParamsActionsID = "sort_query_string_for_cache"
	PageRuleUpdateParamsActionsIDSSL                     PageRuleUpdateParamsActionsID = "ssl"
	PageRuleUpdateParamsActionsIDTrueClientIPHeader      PageRuleUpdateParamsActionsID = "true_client_ip_header"
	PageRuleUpdateParamsActionsIDWAF                     PageRuleUpdateParamsActionsID = "waf"
)

func (r PageRuleUpdateParamsActionsID) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsActionsIDAlwaysUseHTTPS, PageRuleUpdateParamsActionsIDAutomaticHTTPSRewrites, PageRuleUpdateParamsActionsIDBrowserCacheTTL, PageRuleUpdateParamsActionsIDBrowserCheck, PageRuleUpdateParamsActionsIDBypassCacheOnCookie, PageRuleUpdateParamsActionsIDCacheByDeviceType, PageRuleUpdateParamsActionsIDCacheDeceptionArmor, PageRuleUpdateParamsActionsIDCacheKey, PageRuleUpdateParamsActionsIDCacheLevel, PageRuleUpdateParamsActionsIDCacheOnCookie, PageRuleUpdateParamsActionsIDDisableApps, PageRuleUpdateParamsActionsIDDisablePerformance, PageRuleUpdateParamsActionsIDDisableSecurity, PageRuleUpdateParamsActionsIDDisableZaraz, PageRuleUpdateParamsActionsIDEdgeCacheTTL, PageRuleUpdateParamsActionsIDEmailObfuscation, PageRuleUpdateParamsActionsIDExplicitCacheControl, PageRuleUpdateParamsActionsIDForwardingURL, PageRuleUpdateParamsActionsIDHostHeaderOverride, PageRuleUpdateParamsActionsIDIPGeolocation, PageRuleUpdateParamsActionsIDMirage, PageRuleUpdateParamsActionsIDOpportunisticEncryption, PageRuleUpdateParamsActionsIDOriginErrorPagePassThru, PageRuleUpdateParamsActionsIDPolish, PageRuleUpdateParamsActionsIDResolveOverride, PageRuleUpdateParamsActionsIDRespectStrongEtag, PageRuleUpdateParamsActionsIDResponseBuffering, PageRuleUpdateParamsActionsIDRocketLoader, PageRuleUpdateParamsActionsIDSecurityLevel, PageRuleUpdateParamsActionsIDSortQueryStringForCache, PageRuleUpdateParamsActionsIDSSL, PageRuleUpdateParamsActionsIDTrueClientIPHeader, PageRuleUpdateParamsActionsIDWAF:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageRuleUpdateParamsStatus string

const (
	PageRuleUpdateParamsStatusActive   PageRuleUpdateParamsStatus = "active"
	PageRuleUpdateParamsStatusDisabled PageRuleUpdateParamsStatus = "disabled"
)

func (r PageRuleUpdateParamsStatus) IsKnown() bool {
	switch r {
	case PageRuleUpdateParamsStatusActive, PageRuleUpdateParamsStatusDisabled:
		return true
	}
	return false
}

type PageRuleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  PageRule                              `json:"result"`
	JSON    pageRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// pageRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageRuleUpdateResponseEnvelope]
type pageRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageRuleUpdateResponseEnvelopeSuccess bool

const (
	PageRuleUpdateResponseEnvelopeSuccessTrue PageRuleUpdateResponseEnvelopeSuccess = true
)

func (r PageRuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageRuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageRuleListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned Page Rules.
	Direction param.Field[PageRuleListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[PageRuleListParamsMatch] `query:"match"`
	// The field used to sort returned Page Rules.
	Order param.Field[PageRuleListParamsOrder] `query:"order"`
	// The status of the Page Rule.
	Status param.Field[PageRuleListParamsStatus] `query:"status"`
}

// URLQuery serializes [PageRuleListParams]'s query parameters as `url.Values`.
func (r PageRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The direction used to sort returned Page Rules.
type PageRuleListParamsDirection string

const (
	PageRuleListParamsDirectionAsc  PageRuleListParamsDirection = "asc"
	PageRuleListParamsDirectionDesc PageRuleListParamsDirection = "desc"
)

func (r PageRuleListParamsDirection) IsKnown() bool {
	switch r {
	case PageRuleListParamsDirectionAsc, PageRuleListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type PageRuleListParamsMatch string

const (
	PageRuleListParamsMatchAny PageRuleListParamsMatch = "any"
	PageRuleListParamsMatchAll PageRuleListParamsMatch = "all"
)

func (r PageRuleListParamsMatch) IsKnown() bool {
	switch r {
	case PageRuleListParamsMatchAny, PageRuleListParamsMatchAll:
		return true
	}
	return false
}

// The field used to sort returned Page Rules.
type PageRuleListParamsOrder string

const (
	PageRuleListParamsOrderStatus   PageRuleListParamsOrder = "status"
	PageRuleListParamsOrderPriority PageRuleListParamsOrder = "priority"
)

func (r PageRuleListParamsOrder) IsKnown() bool {
	switch r {
	case PageRuleListParamsOrderStatus, PageRuleListParamsOrderPriority:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageRuleListParamsStatus string

const (
	PageRuleListParamsStatusActive   PageRuleListParamsStatus = "active"
	PageRuleListParamsStatusDisabled PageRuleListParamsStatus = "disabled"
)

func (r PageRuleListParamsStatus) IsKnown() bool {
	switch r {
	case PageRuleListParamsStatusActive, PageRuleListParamsStatusDisabled:
		return true
	}
	return false
}

type PageRuleListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageRuleListResponseEnvelopeSuccess `json:"success,required"`
	Result  []PageRule                          `json:"result"`
	JSON    pageRuleListResponseEnvelopeJSON    `json:"-"`
}

// pageRuleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageRuleListResponseEnvelope]
type pageRuleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageRuleListResponseEnvelopeSuccess bool

const (
	PageRuleListResponseEnvelopeSuccessTrue PageRuleListResponseEnvelopeSuccess = true
)

func (r PageRuleListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageRuleListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageRuleDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageRuleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  PageRuleDeleteResponse                `json:"result,nullable"`
	JSON    pageRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// pageRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageRuleDeleteResponseEnvelope]
type pageRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageRuleDeleteResponseEnvelopeSuccess bool

const (
	PageRuleDeleteResponseEnvelopeSuccessTrue PageRuleDeleteResponseEnvelopeSuccess = true
)

func (r PageRuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageRuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageRuleEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]PageRuleEditParamsActionUnion] `json:"actions"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageRuleEditParamsStatus] `json:"status"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargetParam] `json:"targets"`
}

func (r PageRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageRuleEditParamsAction struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID    param.Field[PageRuleEditParamsActionsID] `json:"id"`
	Value param.Field[interface{}]                 `json:"value"`
}

func (r PageRuleEditParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsAction) ImplementsPageRulesPageRuleEditParamsActionUnion() {}

// Satisfied by [zones.AlwaysUseHTTPSParam], [zones.AutomaticHTTPSRewritesParam],
// [zones.BrowserCacheTTLParam], [zones.BrowserCheckParam],
// [page_rules.PageRuleEditParamsActionsBypassCacheOnCookie],
// [page_rules.PageRuleEditParamsActionsCacheByDeviceType],
// [page_rules.PageRuleEditParamsActionsCacheDeceptionArmor],
// [page_rules.PageRuleEditParamsActionsCacheKey], [zones.CacheLevelParam],
// [page_rules.PageRuleEditParamsActionsCacheOnCookie],
// [page_rules.PageRuleEditParamsActionsDisableApps],
// [page_rules.PageRuleEditParamsActionsDisablePerformance],
// [page_rules.PageRuleEditParamsActionsDisableSecurity],
// [page_rules.PageRuleEditParamsActionsDisableZaraz],
// [page_rules.PageRuleEditParamsActionsEdgeCacheTTL],
// [zones.EmailObfuscationParam],
// [page_rules.PageRuleEditParamsActionsExplicitCacheControl],
// [page_rules.PageRuleEditParamsActionsForwardingURL],
// [page_rules.PageRuleEditParamsActionsHostHeaderOverride],
// [zones.IPGeolocationParam], [zones.MirageParam],
// [zones.OpportunisticEncryptionParam], [zones.OriginErrorPagePassThruParam],
// [zones.PolishParam], [page_rules.PageRuleEditParamsActionsResolveOverride],
// [page_rules.PageRuleEditParamsActionsRespectStrongEtag],
// [zones.ResponseBufferingParam], [zones.RocketLoaderParam],
// [zones.SecurityLevelParam], [zones.SortQueryStringForCacheParam],
// [zones.SSLParam], [zones.TrueClientIPHeaderParam], [zones.WAFParam],
// [PageRuleEditParamsAction].
type PageRuleEditParamsActionUnion interface {
	ImplementsPageRulesPageRuleEditParamsActionUnion()
}

type PageRuleEditParamsActionsBypassCacheOnCookie struct {
	// Bypass cache and fetch resources from the origin server if a regular expression
	// matches against a cookie name present in the request.
	ID param.Field[PageRuleEditParamsActionsBypassCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request. Refer to
	// [Bypass Cache on Cookie setting](https://developers.cloudflare.com/rules/page-rules/reference/additional-reference/#bypass-cache-on-cookie-setting)
	// to learn about limited regular expression support.
	Value param.Field[string] `json:"value"`
}

func (r PageRuleEditParamsActionsBypassCacheOnCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsBypassCacheOnCookie) ImplementsPageRulesPageRuleEditParamsActionUnion() {
}

// Bypass cache and fetch resources from the origin server if a regular expression
// matches against a cookie name present in the request.
type PageRuleEditParamsActionsBypassCacheOnCookieID string

const (
	PageRuleEditParamsActionsBypassCacheOnCookieIDBypassCacheOnCookie PageRuleEditParamsActionsBypassCacheOnCookieID = "bypass_cache_on_cookie"
)

func (r PageRuleEditParamsActionsBypassCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsBypassCacheOnCookieIDBypassCacheOnCookie:
		return true
	}
	return false
}

type PageRuleEditParamsActionsCacheByDeviceType struct {
	// Separate cached content based on the visitor's device type.
	ID param.Field[PageRuleEditParamsActionsCacheByDeviceTypeID] `json:"id"`
	// The status of Cache By Device Type.
	Value param.Field[PageRuleEditParamsActionsCacheByDeviceTypeValue] `json:"value"`
}

func (r PageRuleEditParamsActionsCacheByDeviceType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsCacheByDeviceType) ImplementsPageRulesPageRuleEditParamsActionUnion() {
}

// Separate cached content based on the visitor's device type.
type PageRuleEditParamsActionsCacheByDeviceTypeID string

const (
	PageRuleEditParamsActionsCacheByDeviceTypeIDCacheByDeviceType PageRuleEditParamsActionsCacheByDeviceTypeID = "cache_by_device_type"
)

func (r PageRuleEditParamsActionsCacheByDeviceTypeID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsCacheByDeviceTypeIDCacheByDeviceType:
		return true
	}
	return false
}

// The status of Cache By Device Type.
type PageRuleEditParamsActionsCacheByDeviceTypeValue string

const (
	PageRuleEditParamsActionsCacheByDeviceTypeValueOn  PageRuleEditParamsActionsCacheByDeviceTypeValue = "on"
	PageRuleEditParamsActionsCacheByDeviceTypeValueOff PageRuleEditParamsActionsCacheByDeviceTypeValue = "off"
)

func (r PageRuleEditParamsActionsCacheByDeviceTypeValue) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsCacheByDeviceTypeValueOn, PageRuleEditParamsActionsCacheByDeviceTypeValueOff:
		return true
	}
	return false
}

type PageRuleEditParamsActionsCacheDeceptionArmor struct {
	// Protect from web cache deception attacks while still allowing static assets to
	// be cached. This setting verifies that the URL's extension matches the returned
	// `Content-Type`.
	ID param.Field[PageRuleEditParamsActionsCacheDeceptionArmorID] `json:"id"`
	// The status of Cache Deception Armor.
	Value param.Field[PageRuleEditParamsActionsCacheDeceptionArmorValue] `json:"value"`
}

func (r PageRuleEditParamsActionsCacheDeceptionArmor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsCacheDeceptionArmor) ImplementsPageRulesPageRuleEditParamsActionUnion() {
}

// Protect from web cache deception attacks while still allowing static assets to
// be cached. This setting verifies that the URL's extension matches the returned
// `Content-Type`.
type PageRuleEditParamsActionsCacheDeceptionArmorID string

const (
	PageRuleEditParamsActionsCacheDeceptionArmorIDCacheDeceptionArmor PageRuleEditParamsActionsCacheDeceptionArmorID = "cache_deception_armor"
)

func (r PageRuleEditParamsActionsCacheDeceptionArmorID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsCacheDeceptionArmorIDCacheDeceptionArmor:
		return true
	}
	return false
}

// The status of Cache Deception Armor.
type PageRuleEditParamsActionsCacheDeceptionArmorValue string

const (
	PageRuleEditParamsActionsCacheDeceptionArmorValueOn  PageRuleEditParamsActionsCacheDeceptionArmorValue = "on"
	PageRuleEditParamsActionsCacheDeceptionArmorValueOff PageRuleEditParamsActionsCacheDeceptionArmorValue = "off"
)

func (r PageRuleEditParamsActionsCacheDeceptionArmorValue) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsCacheDeceptionArmorValueOn, PageRuleEditParamsActionsCacheDeceptionArmorValueOff:
		return true
	}
	return false
}

type PageRuleEditParamsActionsCacheKey struct {
	// Control specifically what variables to include when deciding which resources to
	// cache. This allows customers to determine what to cache based on something other
	// than just the URL.
	ID    param.Field[PageRuleEditParamsActionsCacheKeyID]    `json:"id"`
	Value param.Field[PageRuleEditParamsActionsCacheKeyValue] `json:"value"`
}

func (r PageRuleEditParamsActionsCacheKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsCacheKey) ImplementsPageRulesPageRuleEditParamsActionUnion() {}

// Control specifically what variables to include when deciding which resources to
// cache. This allows customers to determine what to cache based on something other
// than just the URL.
type PageRuleEditParamsActionsCacheKeyID string

const (
	PageRuleEditParamsActionsCacheKeyIDCacheKey PageRuleEditParamsActionsCacheKeyID = "cache_key"
)

func (r PageRuleEditParamsActionsCacheKeyID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsCacheKeyIDCacheKey:
		return true
	}
	return false
}

type PageRuleEditParamsActionsCacheKeyValue struct {
	// Controls which cookies appear in the Cache Key.
	Cookie param.Field[PageRuleEditParamsActionsCacheKeyValueCookie] `json:"cookie"`
	// Controls which headers go into the Cache Key. Exactly one of `include` or
	// `exclude` is expected.
	Header param.Field[PageRuleEditParamsActionsCacheKeyValueHeader] `json:"header"`
	// Determines which host header to include in the Cache Key.
	Host param.Field[PageRuleEditParamsActionsCacheKeyValueHost] `json:"host"`
	// Controls which URL query string parameters go into the Cache Key. Exactly one of
	// `include` or `exclude` is expected.
	QueryString param.Field[PageRuleEditParamsActionsCacheKeyValueQueryString] `json:"query_string"`
	// Feature fields to add features about the end-user (client) into the Cache Key.
	User param.Field[PageRuleEditParamsActionsCacheKeyValueUser] `json:"user"`
}

func (r PageRuleEditParamsActionsCacheKeyValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which cookies appear in the Cache Key.
type PageRuleEditParamsActionsCacheKeyValueCookie struct {
	// A list of cookies to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of cookies to include.
	Include param.Field[[]string] `json:"include"`
}

func (r PageRuleEditParamsActionsCacheKeyValueCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which headers go into the Cache Key. Exactly one of `include` or
// `exclude` is expected.
type PageRuleEditParamsActionsCacheKeyValueHeader struct {
	// A list of headers to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of headers to ignore.
	Exclude param.Field[[]string] `json:"exclude"`
	// A list of headers to include.
	Include param.Field[[]string] `json:"include"`
}

func (r PageRuleEditParamsActionsCacheKeyValueHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which host header to include in the Cache Key.
type PageRuleEditParamsActionsCacheKeyValueHost struct {
	// Whether to include the Host header in the HTTP request sent to the origin.
	Resolved param.Field[bool] `json:"resolved"`
}

func (r PageRuleEditParamsActionsCacheKeyValueHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which URL query string parameters go into the Cache Key. Exactly one of
// `include` or `exclude` is expected.
type PageRuleEditParamsActionsCacheKeyValueQueryString struct {
	// Ignore all query string parameters.
	Exclude param.Field[PageRuleEditParamsActionsCacheKeyValueQueryStringExcludeUnion] `json:"exclude"`
	// Include all query string parameters.
	Include param.Field[PageRuleEditParamsActionsCacheKeyValueQueryStringIncludeUnion] `json:"include"`
}

func (r PageRuleEditParamsActionsCacheKeyValueQueryString) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Ignore all query string parameters.
//
// Satisfied by
// [page_rules.PageRuleEditParamsActionsCacheKeyValueQueryStringExcludeString],
// [page_rules.PageRuleEditParamsActionsCacheKeyValueQueryStringExcludeArray].
type PageRuleEditParamsActionsCacheKeyValueQueryStringExcludeUnion interface {
	implementsPageRulesPageRuleEditParamsActionsCacheKeyValueQueryStringExcludeUnion()
}

// Ignore all query string parameters.
type PageRuleEditParamsActionsCacheKeyValueQueryStringExcludeString string

const (
	PageRuleEditParamsActionsCacheKeyValueQueryStringExcludeStringStar PageRuleEditParamsActionsCacheKeyValueQueryStringExcludeString = "*"
)

func (r PageRuleEditParamsActionsCacheKeyValueQueryStringExcludeString) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsCacheKeyValueQueryStringExcludeStringStar:
		return true
	}
	return false
}

func (r PageRuleEditParamsActionsCacheKeyValueQueryStringExcludeString) implementsPageRulesPageRuleEditParamsActionsCacheKeyValueQueryStringExcludeUnion() {
}

type PageRuleEditParamsActionsCacheKeyValueQueryStringExcludeArray []string

func (r PageRuleEditParamsActionsCacheKeyValueQueryStringExcludeArray) implementsPageRulesPageRuleEditParamsActionsCacheKeyValueQueryStringExcludeUnion() {
}

// Include all query string parameters.
//
// Satisfied by
// [page_rules.PageRuleEditParamsActionsCacheKeyValueQueryStringIncludeString],
// [page_rules.PageRuleEditParamsActionsCacheKeyValueQueryStringIncludeArray].
type PageRuleEditParamsActionsCacheKeyValueQueryStringIncludeUnion interface {
	implementsPageRulesPageRuleEditParamsActionsCacheKeyValueQueryStringIncludeUnion()
}

// Include all query string parameters.
type PageRuleEditParamsActionsCacheKeyValueQueryStringIncludeString string

const (
	PageRuleEditParamsActionsCacheKeyValueQueryStringIncludeStringStar PageRuleEditParamsActionsCacheKeyValueQueryStringIncludeString = "*"
)

func (r PageRuleEditParamsActionsCacheKeyValueQueryStringIncludeString) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsCacheKeyValueQueryStringIncludeStringStar:
		return true
	}
	return false
}

func (r PageRuleEditParamsActionsCacheKeyValueQueryStringIncludeString) implementsPageRulesPageRuleEditParamsActionsCacheKeyValueQueryStringIncludeUnion() {
}

type PageRuleEditParamsActionsCacheKeyValueQueryStringIncludeArray []string

func (r PageRuleEditParamsActionsCacheKeyValueQueryStringIncludeArray) implementsPageRulesPageRuleEditParamsActionsCacheKeyValueQueryStringIncludeUnion() {
}

// Feature fields to add features about the end-user (client) into the Cache Key.
type PageRuleEditParamsActionsCacheKeyValueUser struct {
	// Classifies a request as `mobile`, `desktop`, or `tablet` based on the User
	// Agent.
	DeviceType param.Field[bool] `json:"device_type"`
	// Includes the client's country, derived from the IP address.
	Geo param.Field[bool] `json:"geo"`
	// Includes the first language code contained in the `Accept-Language` header sent
	// by the client.
	Lang param.Field[bool] `json:"lang"`
}

func (r PageRuleEditParamsActionsCacheKeyValueUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageRuleEditParamsActionsCacheOnCookie struct {
	// Apply the Cache Everything option (Cache Level setting) based on a regular
	// expression match against a cookie name.
	ID param.Field[PageRuleEditParamsActionsCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request.
	Value param.Field[string] `json:"value"`
}

func (r PageRuleEditParamsActionsCacheOnCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsCacheOnCookie) ImplementsPageRulesPageRuleEditParamsActionUnion() {}

// Apply the Cache Everything option (Cache Level setting) based on a regular
// expression match against a cookie name.
type PageRuleEditParamsActionsCacheOnCookieID string

const (
	PageRuleEditParamsActionsCacheOnCookieIDCacheOnCookie PageRuleEditParamsActionsCacheOnCookieID = "cache_on_cookie"
)

func (r PageRuleEditParamsActionsCacheOnCookieID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsCacheOnCookieIDCacheOnCookie:
		return true
	}
	return false
}

type PageRuleEditParamsActionsDisableApps struct {
	// Turn off all active
	// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
	// (deprecated).
	ID param.Field[PageRuleEditParamsActionsDisableAppsID] `json:"id"`
}

func (r PageRuleEditParamsActionsDisableApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsDisableApps) ImplementsPageRulesPageRuleEditParamsActionUnion() {}

// Turn off all active
// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
// (deprecated).
type PageRuleEditParamsActionsDisableAppsID string

const (
	PageRuleEditParamsActionsDisableAppsIDDisableApps PageRuleEditParamsActionsDisableAppsID = "disable_apps"
)

func (r PageRuleEditParamsActionsDisableAppsID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsDisableAppsIDDisableApps:
		return true
	}
	return false
}

type PageRuleEditParamsActionsDisablePerformance struct {
	// Turn off
	// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
	// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
	// and [Polish](https://developers.cloudflare.com/images/polish/).
	ID param.Field[PageRuleEditParamsActionsDisablePerformanceID] `json:"id"`
}

func (r PageRuleEditParamsActionsDisablePerformance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsDisablePerformance) ImplementsPageRulesPageRuleEditParamsActionUnion() {
}

// Turn off
// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
// and [Polish](https://developers.cloudflare.com/images/polish/).
type PageRuleEditParamsActionsDisablePerformanceID string

const (
	PageRuleEditParamsActionsDisablePerformanceIDDisablePerformance PageRuleEditParamsActionsDisablePerformanceID = "disable_performance"
)

func (r PageRuleEditParamsActionsDisablePerformanceID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsDisablePerformanceIDDisablePerformance:
		return true
	}
	return false
}

type PageRuleEditParamsActionsDisableSecurity struct {
	// Turn off
	// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
	// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
	// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
	// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
	// and
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	ID param.Field[PageRuleEditParamsActionsDisableSecurityID] `json:"id"`
}

func (r PageRuleEditParamsActionsDisableSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsDisableSecurity) ImplementsPageRulesPageRuleEditParamsActionUnion() {
}

// Turn off
// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
// and
// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
type PageRuleEditParamsActionsDisableSecurityID string

const (
	PageRuleEditParamsActionsDisableSecurityIDDisableSecurity PageRuleEditParamsActionsDisableSecurityID = "disable_security"
)

func (r PageRuleEditParamsActionsDisableSecurityID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsDisableSecurityIDDisableSecurity:
		return true
	}
	return false
}

type PageRuleEditParamsActionsDisableZaraz struct {
	// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
	ID param.Field[PageRuleEditParamsActionsDisableZarazID] `json:"id"`
}

func (r PageRuleEditParamsActionsDisableZaraz) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsDisableZaraz) ImplementsPageRulesPageRuleEditParamsActionUnion() {}

// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
type PageRuleEditParamsActionsDisableZarazID string

const (
	PageRuleEditParamsActionsDisableZarazIDDisableZaraz PageRuleEditParamsActionsDisableZarazID = "disable_zaraz"
)

func (r PageRuleEditParamsActionsDisableZarazID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsDisableZarazIDDisableZaraz:
		return true
	}
	return false
}

type PageRuleEditParamsActionsEdgeCacheTTL struct {
	// Specify how long to cache a resource in the Cloudflare global network. _Edge
	// Cache TTL_ is not visible in response headers.
	ID    param.Field[PageRuleEditParamsActionsEdgeCacheTTLID] `json:"id"`
	Value param.Field[int64]                                   `json:"value"`
}

func (r PageRuleEditParamsActionsEdgeCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsEdgeCacheTTL) ImplementsPageRulesPageRuleEditParamsActionUnion() {}

// Specify how long to cache a resource in the Cloudflare global network. _Edge
// Cache TTL_ is not visible in response headers.
type PageRuleEditParamsActionsEdgeCacheTTLID string

const (
	PageRuleEditParamsActionsEdgeCacheTTLIDEdgeCacheTTL PageRuleEditParamsActionsEdgeCacheTTLID = "edge_cache_ttl"
)

func (r PageRuleEditParamsActionsEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

type PageRuleEditParamsActionsExplicitCacheControl struct {
	// Origin Cache Control is enabled by default for Free, Pro, and Business domains
	// and disabled by default for Enterprise domains.
	ID param.Field[PageRuleEditParamsActionsExplicitCacheControlID] `json:"id"`
	// The status of Origin Cache Control.
	Value param.Field[PageRuleEditParamsActionsExplicitCacheControlValue] `json:"value"`
}

func (r PageRuleEditParamsActionsExplicitCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsExplicitCacheControl) ImplementsPageRulesPageRuleEditParamsActionUnion() {
}

// Origin Cache Control is enabled by default for Free, Pro, and Business domains
// and disabled by default for Enterprise domains.
type PageRuleEditParamsActionsExplicitCacheControlID string

const (
	PageRuleEditParamsActionsExplicitCacheControlIDExplicitCacheControl PageRuleEditParamsActionsExplicitCacheControlID = "explicit_cache_control"
)

func (r PageRuleEditParamsActionsExplicitCacheControlID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsExplicitCacheControlIDExplicitCacheControl:
		return true
	}
	return false
}

// The status of Origin Cache Control.
type PageRuleEditParamsActionsExplicitCacheControlValue string

const (
	PageRuleEditParamsActionsExplicitCacheControlValueOn  PageRuleEditParamsActionsExplicitCacheControlValue = "on"
	PageRuleEditParamsActionsExplicitCacheControlValueOff PageRuleEditParamsActionsExplicitCacheControlValue = "off"
)

func (r PageRuleEditParamsActionsExplicitCacheControlValue) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsExplicitCacheControlValueOn, PageRuleEditParamsActionsExplicitCacheControlValueOff:
		return true
	}
	return false
}

type PageRuleEditParamsActionsForwardingURL struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    param.Field[PageRuleEditParamsActionsForwardingURLID]    `json:"id"`
	Value param.Field[PageRuleEditParamsActionsForwardingURLValue] `json:"value"`
}

func (r PageRuleEditParamsActionsForwardingURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsForwardingURL) ImplementsPageRulesPageRuleEditParamsActionUnion() {}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageRuleEditParamsActionsForwardingURLID string

const (
	PageRuleEditParamsActionsForwardingURLIDForwardingURL PageRuleEditParamsActionsForwardingURLID = "forwarding_url"
)

func (r PageRuleEditParamsActionsForwardingURLID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsForwardingURLIDForwardingURL:
		return true
	}
	return false
}

type PageRuleEditParamsActionsForwardingURLValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode param.Field[PageRuleEditParamsActionsForwardingURLValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageRuleEditParamsActionsForwardingURLValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageRuleEditParamsActionsForwardingURLValueStatusCode int64

const (
	PageRuleEditParamsActionsForwardingURLValueStatusCode301 PageRuleEditParamsActionsForwardingURLValueStatusCode = 301
	PageRuleEditParamsActionsForwardingURLValueStatusCode302 PageRuleEditParamsActionsForwardingURLValueStatusCode = 302
)

func (r PageRuleEditParamsActionsForwardingURLValueStatusCode) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsForwardingURLValueStatusCode301, PageRuleEditParamsActionsForwardingURLValueStatusCode302:
		return true
	}
	return false
}

type PageRuleEditParamsActionsHostHeaderOverride struct {
	// Apply a specific host header.
	ID param.Field[PageRuleEditParamsActionsHostHeaderOverrideID] `json:"id"`
	// The hostname to use in the `Host` header
	Value param.Field[string] `json:"value"`
}

func (r PageRuleEditParamsActionsHostHeaderOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsHostHeaderOverride) ImplementsPageRulesPageRuleEditParamsActionUnion() {
}

// Apply a specific host header.
type PageRuleEditParamsActionsHostHeaderOverrideID string

const (
	PageRuleEditParamsActionsHostHeaderOverrideIDHostHeaderOverride PageRuleEditParamsActionsHostHeaderOverrideID = "host_header_override"
)

func (r PageRuleEditParamsActionsHostHeaderOverrideID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsHostHeaderOverrideIDHostHeaderOverride:
		return true
	}
	return false
}

type PageRuleEditParamsActionsResolveOverride struct {
	// Change the origin address to the value specified in this setting.
	ID param.Field[PageRuleEditParamsActionsResolveOverrideID] `json:"id"`
	// The origin address you want to override with.
	Value param.Field[string] `json:"value"`
}

func (r PageRuleEditParamsActionsResolveOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsResolveOverride) ImplementsPageRulesPageRuleEditParamsActionUnion() {
}

// Change the origin address to the value specified in this setting.
type PageRuleEditParamsActionsResolveOverrideID string

const (
	PageRuleEditParamsActionsResolveOverrideIDResolveOverride PageRuleEditParamsActionsResolveOverrideID = "resolve_override"
)

func (r PageRuleEditParamsActionsResolveOverrideID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsResolveOverrideIDResolveOverride:
		return true
	}
	return false
}

type PageRuleEditParamsActionsRespectStrongEtag struct {
	// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
	// the origin server.
	ID param.Field[PageRuleEditParamsActionsRespectStrongEtagID] `json:"id"`
	// The status of Respect Strong ETags
	Value param.Field[PageRuleEditParamsActionsRespectStrongEtagValue] `json:"value"`
}

func (r PageRuleEditParamsActionsRespectStrongEtag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRuleEditParamsActionsRespectStrongEtag) ImplementsPageRulesPageRuleEditParamsActionUnion() {
}

// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
// the origin server.
type PageRuleEditParamsActionsRespectStrongEtagID string

const (
	PageRuleEditParamsActionsRespectStrongEtagIDRespectStrongEtag PageRuleEditParamsActionsRespectStrongEtagID = "respect_strong_etag"
)

func (r PageRuleEditParamsActionsRespectStrongEtagID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsRespectStrongEtagIDRespectStrongEtag:
		return true
	}
	return false
}

// The status of Respect Strong ETags
type PageRuleEditParamsActionsRespectStrongEtagValue string

const (
	PageRuleEditParamsActionsRespectStrongEtagValueOn  PageRuleEditParamsActionsRespectStrongEtagValue = "on"
	PageRuleEditParamsActionsRespectStrongEtagValueOff PageRuleEditParamsActionsRespectStrongEtagValue = "off"
)

func (r PageRuleEditParamsActionsRespectStrongEtagValue) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsRespectStrongEtagValueOn, PageRuleEditParamsActionsRespectStrongEtagValueOff:
		return true
	}
	return false
}

// If enabled, any ` http://“ URL is converted to  `https://` through a 301
// redirect.
type PageRuleEditParamsActionsID string

const (
	PageRuleEditParamsActionsIDAlwaysUseHTTPS          PageRuleEditParamsActionsID = "always_use_https"
	PageRuleEditParamsActionsIDAutomaticHTTPSRewrites  PageRuleEditParamsActionsID = "automatic_https_rewrites"
	PageRuleEditParamsActionsIDBrowserCacheTTL         PageRuleEditParamsActionsID = "browser_cache_ttl"
	PageRuleEditParamsActionsIDBrowserCheck            PageRuleEditParamsActionsID = "browser_check"
	PageRuleEditParamsActionsIDBypassCacheOnCookie     PageRuleEditParamsActionsID = "bypass_cache_on_cookie"
	PageRuleEditParamsActionsIDCacheByDeviceType       PageRuleEditParamsActionsID = "cache_by_device_type"
	PageRuleEditParamsActionsIDCacheDeceptionArmor     PageRuleEditParamsActionsID = "cache_deception_armor"
	PageRuleEditParamsActionsIDCacheKey                PageRuleEditParamsActionsID = "cache_key"
	PageRuleEditParamsActionsIDCacheLevel              PageRuleEditParamsActionsID = "cache_level"
	PageRuleEditParamsActionsIDCacheOnCookie           PageRuleEditParamsActionsID = "cache_on_cookie"
	PageRuleEditParamsActionsIDDisableApps             PageRuleEditParamsActionsID = "disable_apps"
	PageRuleEditParamsActionsIDDisablePerformance      PageRuleEditParamsActionsID = "disable_performance"
	PageRuleEditParamsActionsIDDisableSecurity         PageRuleEditParamsActionsID = "disable_security"
	PageRuleEditParamsActionsIDDisableZaraz            PageRuleEditParamsActionsID = "disable_zaraz"
	PageRuleEditParamsActionsIDEdgeCacheTTL            PageRuleEditParamsActionsID = "edge_cache_ttl"
	PageRuleEditParamsActionsIDEmailObfuscation        PageRuleEditParamsActionsID = "email_obfuscation"
	PageRuleEditParamsActionsIDExplicitCacheControl    PageRuleEditParamsActionsID = "explicit_cache_control"
	PageRuleEditParamsActionsIDForwardingURL           PageRuleEditParamsActionsID = "forwarding_url"
	PageRuleEditParamsActionsIDHostHeaderOverride      PageRuleEditParamsActionsID = "host_header_override"
	PageRuleEditParamsActionsIDIPGeolocation           PageRuleEditParamsActionsID = "ip_geolocation"
	PageRuleEditParamsActionsIDMirage                  PageRuleEditParamsActionsID = "mirage"
	PageRuleEditParamsActionsIDOpportunisticEncryption PageRuleEditParamsActionsID = "opportunistic_encryption"
	PageRuleEditParamsActionsIDOriginErrorPagePassThru PageRuleEditParamsActionsID = "origin_error_page_pass_thru"
	PageRuleEditParamsActionsIDPolish                  PageRuleEditParamsActionsID = "polish"
	PageRuleEditParamsActionsIDResolveOverride         PageRuleEditParamsActionsID = "resolve_override"
	PageRuleEditParamsActionsIDRespectStrongEtag       PageRuleEditParamsActionsID = "respect_strong_etag"
	PageRuleEditParamsActionsIDResponseBuffering       PageRuleEditParamsActionsID = "response_buffering"
	PageRuleEditParamsActionsIDRocketLoader            PageRuleEditParamsActionsID = "rocket_loader"
	PageRuleEditParamsActionsIDSecurityLevel           PageRuleEditParamsActionsID = "security_level"
	PageRuleEditParamsActionsIDSortQueryStringForCache PageRuleEditParamsActionsID = "sort_query_string_for_cache"
	PageRuleEditParamsActionsIDSSL                     PageRuleEditParamsActionsID = "ssl"
	PageRuleEditParamsActionsIDTrueClientIPHeader      PageRuleEditParamsActionsID = "true_client_ip_header"
	PageRuleEditParamsActionsIDWAF                     PageRuleEditParamsActionsID = "waf"
)

func (r PageRuleEditParamsActionsID) IsKnown() bool {
	switch r {
	case PageRuleEditParamsActionsIDAlwaysUseHTTPS, PageRuleEditParamsActionsIDAutomaticHTTPSRewrites, PageRuleEditParamsActionsIDBrowserCacheTTL, PageRuleEditParamsActionsIDBrowserCheck, PageRuleEditParamsActionsIDBypassCacheOnCookie, PageRuleEditParamsActionsIDCacheByDeviceType, PageRuleEditParamsActionsIDCacheDeceptionArmor, PageRuleEditParamsActionsIDCacheKey, PageRuleEditParamsActionsIDCacheLevel, PageRuleEditParamsActionsIDCacheOnCookie, PageRuleEditParamsActionsIDDisableApps, PageRuleEditParamsActionsIDDisablePerformance, PageRuleEditParamsActionsIDDisableSecurity, PageRuleEditParamsActionsIDDisableZaraz, PageRuleEditParamsActionsIDEdgeCacheTTL, PageRuleEditParamsActionsIDEmailObfuscation, PageRuleEditParamsActionsIDExplicitCacheControl, PageRuleEditParamsActionsIDForwardingURL, PageRuleEditParamsActionsIDHostHeaderOverride, PageRuleEditParamsActionsIDIPGeolocation, PageRuleEditParamsActionsIDMirage, PageRuleEditParamsActionsIDOpportunisticEncryption, PageRuleEditParamsActionsIDOriginErrorPagePassThru, PageRuleEditParamsActionsIDPolish, PageRuleEditParamsActionsIDResolveOverride, PageRuleEditParamsActionsIDRespectStrongEtag, PageRuleEditParamsActionsIDResponseBuffering, PageRuleEditParamsActionsIDRocketLoader, PageRuleEditParamsActionsIDSecurityLevel, PageRuleEditParamsActionsIDSortQueryStringForCache, PageRuleEditParamsActionsIDSSL, PageRuleEditParamsActionsIDTrueClientIPHeader, PageRuleEditParamsActionsIDWAF:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageRuleEditParamsStatus string

const (
	PageRuleEditParamsStatusActive   PageRuleEditParamsStatus = "active"
	PageRuleEditParamsStatusDisabled PageRuleEditParamsStatus = "disabled"
)

func (r PageRuleEditParamsStatus) IsKnown() bool {
	switch r {
	case PageRuleEditParamsStatusActive, PageRuleEditParamsStatusDisabled:
		return true
	}
	return false
}

type PageRuleEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageRuleEditResponseEnvelopeSuccess `json:"success,required"`
	Result  PageRule                            `json:"result"`
	JSON    pageRuleEditResponseEnvelopeJSON    `json:"-"`
}

// pageRuleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageRuleEditResponseEnvelope]
type pageRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageRuleEditResponseEnvelopeSuccess bool

const (
	PageRuleEditResponseEnvelopeSuccessTrue PageRuleEditResponseEnvelopeSuccess = true
)

func (r PageRuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageRuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageRuleGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageRuleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageRuleGetResponseEnvelopeSuccess `json:"success,required"`
	Result  PageRule                           `json:"result"`
	JSON    pageRuleGetResponseEnvelopeJSON    `json:"-"`
}

// pageRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageRuleGetResponseEnvelope]
type pageRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageRuleGetResponseEnvelopeSuccess bool

const (
	PageRuleGetResponseEnvelopeSuccessTrue PageRuleGetResponseEnvelopeSuccess = true
)

func (r PageRuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageRuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
