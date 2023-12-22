// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneFirewallWafOverrideService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneFirewallWafOverrideService] method instead.
type ZoneFirewallWafOverrideService struct {
	Options []option.RequestOption
}

// NewZoneFirewallWafOverrideService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneFirewallWafOverrideService(opts ...option.RequestOption) (r *ZoneFirewallWafOverrideService) {
	r = &ZoneFirewallWafOverrideService{}
	r.Options = opts
	return
}

// Fetches the details of a URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *OverrideResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) Update(ctx context.Context, zoneIdentifier string, id string, body ZoneFirewallWafOverrideUpdateParams, opts ...option.RequestOption) (res *OverrideResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneFirewallWafOverrideDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a URI-based WAF override for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) WafOverridesNewAWafOverride(ctx context.Context, zoneIdentifier string, body ZoneFirewallWafOverrideWafOverridesNewAWafOverrideParams, opts ...option.RequestOption) (res *OverrideResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the URI-based WAF overrides in a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) WafOverridesListWafOverrides(ctx context.Context, zoneIdentifier string, query ZoneFirewallWafOverrideWafOverridesListWafOverridesParams, opts ...option.RequestOption) (res *OverrideResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OverrideResponseCollection struct {
	Errors     []OverrideResponseCollectionError    `json:"errors"`
	Messages   []OverrideResponseCollectionMessage  `json:"messages"`
	Result     []OverrideResponseCollectionResult   `json:"result"`
	ResultInfo OverrideResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success OverrideResponseCollectionSuccess `json:"success"`
	JSON    overrideResponseCollectionJSON    `json:"-"`
}

// overrideResponseCollectionJSON contains the JSON metadata for the struct
// [OverrideResponseCollection]
type overrideResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideResponseCollectionError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    overrideResponseCollectionErrorJSON `json:"-"`
}

// overrideResponseCollectionErrorJSON contains the JSON metadata for the struct
// [OverrideResponseCollectionError]
type overrideResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideResponseCollectionMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    overrideResponseCollectionMessageJSON `json:"-"`
}

// overrideResponseCollectionMessageJSON contains the JSON metadata for the struct
// [OverrideResponseCollectionMessage]
type overrideResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideResponseCollectionResult struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction OverrideResponseCollectionResultRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules interface{} `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string                             `json:"urls"`
	JSON overrideResponseCollectionResultJSON `json:"-"`
}

// overrideResponseCollectionResultJSON contains the JSON metadata for the struct
// [OverrideResponseCollectionResult]
type overrideResponseCollectionResultJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	Groups        apijson.Field
	Paused        apijson.Field
	Priority      apijson.Field
	RewriteAction apijson.Field
	Rules         apijson.Field
	URLs          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OverrideResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type OverrideResponseCollectionResultRewriteAction struct {
	// The WAF rule action to apply.
	Block     OverrideResponseCollectionResultRewriteActionBlock `json:"block"`
	Challenge interface{}                                        `json:"challenge"`
	Default   interface{}                                        `json:"default"`
	// The WAF rule action to apply.
	Disable  OverrideResponseCollectionResultRewriteActionDisable `json:"disable"`
	Simulate interface{}                                          `json:"simulate"`
	JSON     overrideResponseCollectionResultRewriteActionJSON    `json:"-"`
}

// overrideResponseCollectionResultRewriteActionJSON contains the JSON metadata for
// the struct [OverrideResponseCollectionResultRewriteAction]
type overrideResponseCollectionResultRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideResponseCollectionResultRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF rule action to apply.
type OverrideResponseCollectionResultRewriteActionBlock string

const (
	OverrideResponseCollectionResultRewriteActionBlockChallenge OverrideResponseCollectionResultRewriteActionBlock = "challenge"
	OverrideResponseCollectionResultRewriteActionBlockBlock     OverrideResponseCollectionResultRewriteActionBlock = "block"
	OverrideResponseCollectionResultRewriteActionBlockSimulate  OverrideResponseCollectionResultRewriteActionBlock = "simulate"
	OverrideResponseCollectionResultRewriteActionBlockDisable   OverrideResponseCollectionResultRewriteActionBlock = "disable"
	OverrideResponseCollectionResultRewriteActionBlockDefault   OverrideResponseCollectionResultRewriteActionBlock = "default"
)

// The WAF rule action to apply.
type OverrideResponseCollectionResultRewriteActionDisable string

const (
	OverrideResponseCollectionResultRewriteActionDisableChallenge OverrideResponseCollectionResultRewriteActionDisable = "challenge"
	OverrideResponseCollectionResultRewriteActionDisableBlock     OverrideResponseCollectionResultRewriteActionDisable = "block"
	OverrideResponseCollectionResultRewriteActionDisableSimulate  OverrideResponseCollectionResultRewriteActionDisable = "simulate"
	OverrideResponseCollectionResultRewriteActionDisableDisable   OverrideResponseCollectionResultRewriteActionDisable = "disable"
	OverrideResponseCollectionResultRewriteActionDisableDefault   OverrideResponseCollectionResultRewriteActionDisable = "default"
)

type OverrideResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       overrideResponseCollectionResultInfoJSON `json:"-"`
}

// overrideResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [OverrideResponseCollectionResultInfo]
type overrideResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OverrideResponseCollectionSuccess bool

const (
	OverrideResponseCollectionSuccessTrue OverrideResponseCollectionSuccess = true
)

type OverrideResponseSingle struct {
	Errors   []OverrideResponseSingleError   `json:"errors"`
	Messages []OverrideResponseSingleMessage `json:"messages"`
	Result   OverrideResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success OverrideResponseSingleSuccess `json:"success"`
	JSON    overrideResponseSingleJSON    `json:"-"`
}

// overrideResponseSingleJSON contains the JSON metadata for the struct
// [OverrideResponseSingle]
type overrideResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideResponseSingleError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    overrideResponseSingleErrorJSON `json:"-"`
}

// overrideResponseSingleErrorJSON contains the JSON metadata for the struct
// [OverrideResponseSingleError]
type overrideResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideResponseSingleMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    overrideResponseSingleMessageJSON `json:"-"`
}

// overrideResponseSingleMessageJSON contains the JSON metadata for the struct
// [OverrideResponseSingleMessage]
type overrideResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideResponseSingleResult struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction OverrideResponseSingleResultRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules interface{} `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string                         `json:"urls"`
	JSON overrideResponseSingleResultJSON `json:"-"`
}

// overrideResponseSingleResultJSON contains the JSON metadata for the struct
// [OverrideResponseSingleResult]
type overrideResponseSingleResultJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	Groups        apijson.Field
	Paused        apijson.Field
	Priority      apijson.Field
	RewriteAction apijson.Field
	Rules         apijson.Field
	URLs          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OverrideResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type OverrideResponseSingleResultRewriteAction struct {
	// The WAF rule action to apply.
	Block     OverrideResponseSingleResultRewriteActionBlock `json:"block"`
	Challenge interface{}                                    `json:"challenge"`
	Default   interface{}                                    `json:"default"`
	// The WAF rule action to apply.
	Disable  OverrideResponseSingleResultRewriteActionDisable `json:"disable"`
	Simulate interface{}                                      `json:"simulate"`
	JSON     overrideResponseSingleResultRewriteActionJSON    `json:"-"`
}

// overrideResponseSingleResultRewriteActionJSON contains the JSON metadata for the
// struct [OverrideResponseSingleResultRewriteAction]
type overrideResponseSingleResultRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideResponseSingleResultRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF rule action to apply.
type OverrideResponseSingleResultRewriteActionBlock string

const (
	OverrideResponseSingleResultRewriteActionBlockChallenge OverrideResponseSingleResultRewriteActionBlock = "challenge"
	OverrideResponseSingleResultRewriteActionBlockBlock     OverrideResponseSingleResultRewriteActionBlock = "block"
	OverrideResponseSingleResultRewriteActionBlockSimulate  OverrideResponseSingleResultRewriteActionBlock = "simulate"
	OverrideResponseSingleResultRewriteActionBlockDisable   OverrideResponseSingleResultRewriteActionBlock = "disable"
	OverrideResponseSingleResultRewriteActionBlockDefault   OverrideResponseSingleResultRewriteActionBlock = "default"
)

// The WAF rule action to apply.
type OverrideResponseSingleResultRewriteActionDisable string

const (
	OverrideResponseSingleResultRewriteActionDisableChallenge OverrideResponseSingleResultRewriteActionDisable = "challenge"
	OverrideResponseSingleResultRewriteActionDisableBlock     OverrideResponseSingleResultRewriteActionDisable = "block"
	OverrideResponseSingleResultRewriteActionDisableSimulate  OverrideResponseSingleResultRewriteActionDisable = "simulate"
	OverrideResponseSingleResultRewriteActionDisableDisable   OverrideResponseSingleResultRewriteActionDisable = "disable"
	OverrideResponseSingleResultRewriteActionDisableDefault   OverrideResponseSingleResultRewriteActionDisable = "default"
)

// Whether the API call was successful
type OverrideResponseSingleSuccess bool

const (
	OverrideResponseSingleSuccessTrue OverrideResponseSingleSuccess = true
)

type ZoneFirewallWafOverrideDeleteResponse struct {
	Result ZoneFirewallWafOverrideDeleteResponseResult `json:"result"`
	JSON   zoneFirewallWafOverrideDeleteResponseJSON   `json:"-"`
}

// zoneFirewallWafOverrideDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafOverrideDeleteResponse]
type zoneFirewallWafOverrideDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideDeleteResponseResult struct {
	// The unique identifier of the WAF override.
	ID   string                                          `json:"id"`
	JSON zoneFirewallWafOverrideDeleteResponseResultJSON `json:"-"`
}

// zoneFirewallWafOverrideDeleteResponseResultJSON contains the JSON metadata for
// the struct [ZoneFirewallWafOverrideDeleteResponseResult]
type zoneFirewallWafOverrideDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallWafOverrideUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallWafOverrideWafOverridesNewAWafOverrideParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallWafOverrideWafOverridesNewAWafOverrideParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallWafOverrideWafOverridesListWafOverridesParams struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of WAF overrides per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [ZoneFirewallWafOverrideWafOverridesListWafOverridesParams]'s query parameters
// as `url.Values`.
func (r ZoneFirewallWafOverrideWafOverridesListWafOverridesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
