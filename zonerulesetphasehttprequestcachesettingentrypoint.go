// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointService contains methods and
// other services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRulesetPhaseHTTPRequestCacheSettingEntrypointService] method
// instead.
type ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointService struct {
	Options []option.RequestOption
}

// NewZoneRulesetPhaseHTTPRequestCacheSettingEntrypointService generates a new
// service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewZoneRulesetPhaseHTTPRequestCacheSettingEntrypointService(opts ...option.RequestOption) (r *ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointService) {
	r = &ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointService{}
	r.Options = opts
	return
}

// Fetches all Cache Rules in a zone.
func (r *ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointService) CacheRulesListCacheRules(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CacheRulesRuleset, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/http_request_cache_settings/entrypoint", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Cache Rules of a zone.
func (r *ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointService) CacheRulesUpdateCacheRules(ctx context.Context, zoneID string, body ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/http_request_cache_settings/entrypoint", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CacheRulesRuleset struct {
	ID          interface{} `json:"id"`
	Description interface{} `json:"description"`
	Kind        interface{} `json:"kind"`
	Name        interface{} `json:"name"`
	Phase       interface{} `json:"phase"`
	// The rules in the ruleset.
	Rules []CacheRulesRulesetRule `json:"rules"`
	JSON  cacheRulesRulesetJSON   `json:"-"`
}

// cacheRulesRulesetJSON contains the JSON metadata for the struct
// [CacheRulesRuleset]
type cacheRulesRulesetJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRulesRulesetRule struct {
	ID     interface{} `json:"id"`
	Action interface{} `json:"action"`
	// The parameters configuring the action.
	ActionParameters CacheRulesRulesetRulesActionParameters `json:"action_parameters"`
	Description      interface{}                            `json:"description"`
	Expression       interface{}                            `json:"expression"`
	Version          interface{}                            `json:"version"`
	JSON             cacheRulesRulesetRuleJSON              `json:"-"`
}

// cacheRulesRulesetRuleJSON contains the JSON metadata for the struct
// [CacheRulesRulesetRule]
type cacheRulesRulesetRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Description      apijson.Field
	Expression       apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CacheRulesRulesetRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The parameters configuring the action.
type CacheRulesRulesetRulesActionParameters struct {
	// Set the Browser TTL.
	BrowserTtl CacheRulesRulesetRulesActionParametersBrowserTtl `json:"browser_ttl"`
	// Set the Cache TTL.
	CacheKey CacheRulesRulesetRulesActionParametersCacheKey `json:"cache_key"`
	// Set the Cache TTL.
	EdgeTtl                 CacheRulesRulesetRulesActionParametersEdgeTtl    `json:"edge_ttl"`
	OriginErrorPagePassthru bool                                             `json:"origin_error_page_passthru"`
	RespectStrongEtags      bool                                             `json:"respect_strong_etags"`
	ServeStale              CacheRulesRulesetRulesActionParametersServeStale `json:"serve_stale"`
	JSON                    cacheRulesRulesetRulesActionParametersJSON       `json:"-"`
}

// cacheRulesRulesetRulesActionParametersJSON contains the JSON metadata for the
// struct [CacheRulesRulesetRulesActionParameters]
type cacheRulesRulesetRulesActionParametersJSON struct {
	BrowserTtl              apijson.Field
	CacheKey                apijson.Field
	EdgeTtl                 apijson.Field
	OriginErrorPagePassthru apijson.Field
	RespectStrongEtags      apijson.Field
	ServeStale              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CacheRulesRulesetRulesActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Set the Browser TTL.
type CacheRulesRulesetRulesActionParametersBrowserTtl struct {
	Default int64                                                `json:"default"`
	Mode    string                                               `json:"mode"`
	JSON    cacheRulesRulesetRulesActionParametersBrowserTtlJSON `json:"-"`
}

// cacheRulesRulesetRulesActionParametersBrowserTtlJSON contains the JSON metadata
// for the struct [CacheRulesRulesetRulesActionParametersBrowserTtl]
type cacheRulesRulesetRulesActionParametersBrowserTtlJSON struct {
	Default     apijson.Field
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesRulesetRulesActionParametersBrowserTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Set the Cache TTL.
type CacheRulesRulesetRulesActionParametersCacheKey struct {
	CacheDeceptionArmor     bool                                                    `json:"cache_deception_armor"`
	CustomKey               CacheRulesRulesetRulesActionParametersCacheKeyCustomKey `json:"custom_key"`
	IgnoreQueryStringsOrder bool                                                    `json:"ignore_query_strings_order"`
	JSON                    cacheRulesRulesetRulesActionParametersCacheKeyJSON      `json:"-"`
}

// cacheRulesRulesetRulesActionParametersCacheKeyJSON contains the JSON metadata
// for the struct [CacheRulesRulesetRulesActionParametersCacheKey]
type cacheRulesRulesetRulesActionParametersCacheKeyJSON struct {
	CacheDeceptionArmor     apijson.Field
	CustomKey               apijson.Field
	IgnoreQueryStringsOrder apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CacheRulesRulesetRulesActionParametersCacheKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRulesRulesetRulesActionParametersCacheKeyCustomKey struct {
	Cookie      CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyCookie      `json:"cookie"`
	Header      CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHeader      `json:"header"`
	Host        CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHost        `json:"host"`
	QueryString CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyQueryString `json:"query_string"`
	User        CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyUser        `json:"user"`
	JSON        cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyJSON        `json:"-"`
}

// cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyJSON contains the JSON
// metadata for the struct
// [CacheRulesRulesetRulesActionParametersCacheKeyCustomKey]
type cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesRulesetRulesActionParametersCacheKeyCustomKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyCookie struct {
	CheckPresence []interface{}                                                     `json:"check_presence"`
	Include       []interface{}                                                     `json:"include"`
	JSON          cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyCookieJSON `json:"-"`
}

// cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyCookieJSON contains the
// JSON metadata for the struct
// [CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyCookie]
type cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHeader struct {
	CheckPresence []interface{}                                                     `json:"check_presence"`
	Include       []interface{}                                                     `json:"include"`
	JSON          cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHeaderJSON `json:"-"`
}

// cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHeaderJSON contains the
// JSON metadata for the struct
// [CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHeader]
type cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHeaderJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHost struct {
	Resolved bool                                                            `json:"resolved"`
	JSON     cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHostJSON `json:"-"`
}

// cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHostJSON contains the
// JSON metadata for the struct
// [CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHost]
type cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyQueryString struct {
	Include string                                                                 `json:"include"`
	JSON    cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyQueryStringJSON `json:"-"`
}

// cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyQueryStringJSON contains
// the JSON metadata for the struct
// [CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyQueryString]
type cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyQueryStringJSON struct {
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyQueryString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyUser struct {
	DeviceType bool                                                            `json:"device_type"`
	Geo        bool                                                            `json:"geo"`
	Lang       bool                                                            `json:"lang"`
	JSON       cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyUserJSON `json:"-"`
}

// cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyUserJSON contains the
// JSON metadata for the struct
// [CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyUser]
type cacheRulesRulesetRulesActionParametersCacheKeyCustomKeyUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesRulesetRulesActionParametersCacheKeyCustomKeyUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Set the Cache TTL.
type CacheRulesRulesetRulesActionParametersEdgeTtl struct {
	Mode          string                                                     `json:"mode"`
	StatusCodeTtl CacheRulesRulesetRulesActionParametersEdgeTtlStatusCodeTtl `json:"status_code_ttl"`
	JSON          cacheRulesRulesetRulesActionParametersEdgeTtlJSON          `json:"-"`
}

// cacheRulesRulesetRulesActionParametersEdgeTtlJSON contains the JSON metadata for
// the struct [CacheRulesRulesetRulesActionParametersEdgeTtl]
type cacheRulesRulesetRulesActionParametersEdgeTtlJSON struct {
	Mode          apijson.Field
	StatusCodeTtl apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CacheRulesRulesetRulesActionParametersEdgeTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRulesRulesetRulesActionParametersEdgeTtlStatusCodeTtl struct {
	StatusCode int64                                                          `json:"status_code"`
	Value      int64                                                          `json:"value"`
	JSON       cacheRulesRulesetRulesActionParametersEdgeTtlStatusCodeTtlJSON `json:"-"`
}

// cacheRulesRulesetRulesActionParametersEdgeTtlStatusCodeTtlJSON contains the JSON
// metadata for the struct
// [CacheRulesRulesetRulesActionParametersEdgeTtlStatusCodeTtl]
type cacheRulesRulesetRulesActionParametersEdgeTtlStatusCodeTtlJSON struct {
	StatusCode  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesRulesetRulesActionParametersEdgeTtlStatusCodeTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRulesRulesetRulesActionParametersServeStale struct {
	DisableStaleWhileUpdating bool                                                 `json:"disable_stale_while_updating"`
	JSON                      cacheRulesRulesetRulesActionParametersServeStaleJSON `json:"-"`
}

// cacheRulesRulesetRulesActionParametersServeStaleJSON contains the JSON metadata
// for the struct [CacheRulesRulesetRulesActionParametersServeStale]
type cacheRulesRulesetRulesActionParametersServeStaleJSON struct {
	DisableStaleWhileUpdating apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CacheRulesRulesetRulesActionParametersServeStale) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParams struct {
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object.
//
// Satisfied by
// [ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParamsRulesCreateUpdateRule],
// [ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParamsRulesObject].
type ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParamsRule interface {
	implementsZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParamsRule()
}

// A rule object.
type ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParamsRulesCreateUpdateRule struct {
	// The action to perform when the rule matches.
	Action param.Field[string] `json:"action,required"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression,required"`
	// The parameters configuring the rule action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParamsRulesCreateUpdateRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParamsRulesCreateUpdateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParamsRulesCreateUpdateRule) implementsZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParamsRule() {
}

// An object configuring the rule's logging behavior.
type ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParamsRulesCreateUpdateRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneRulesetPhaseHTTPRequestCacheSettingEntrypointCacheRulesUpdateCacheRulesParamsRulesCreateUpdateRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
