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

// ZoneFirewallWafPackageGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneFirewallWafPackageGroupService] method instead.
type ZoneFirewallWafPackageGroupService struct {
	Options []option.RequestOption
}

// NewZoneFirewallWafPackageGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneFirewallWafPackageGroupService(opts ...option.RequestOption) (r *ZoneFirewallWafPackageGroupService) {
	r = &ZoneFirewallWafPackageGroupService{}
	r.Options = opts
	return
}

// Fetches the details of a WAF rule group.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageGroupService) Get(ctx context.Context, zoneIdentifier string, packageIdentifier string, identifier string, opts ...option.RequestOption) (res *RuleGroupResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", zoneIdentifier, packageIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a WAF rule group. You can update the state (`mode` parameter) of a rule
// group.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageGroupService) Update(ctx context.Context, zoneIdentifier string, packageIdentifier string, identifier string, body ZoneFirewallWafPackageGroupUpdateParams, opts ...option.RequestOption) (res *RuleGroupResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", zoneIdentifier, packageIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches the WAF rule groups in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageGroupService) WafRuleGroupsListWafRuleGroups(ctx context.Context, zoneIdentifier string, packageIdentifier string, query ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParams, opts ...option.RequestOption) (res *RuleGroupResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups", zoneIdentifier, packageIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RuleGroupResponseCollection struct {
	Errors     []RuleGroupResponseCollectionError    `json:"errors"`
	Messages   []RuleGroupResponseCollectionMessage  `json:"messages"`
	Result     []RuleGroupResponseCollectionResult   `json:"result"`
	ResultInfo RuleGroupResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success RuleGroupResponseCollectionSuccess `json:"success"`
	JSON    ruleGroupResponseCollectionJSON    `json:"-"`
}

// ruleGroupResponseCollectionJSON contains the JSON metadata for the struct
// [RuleGroupResponseCollection]
type ruleGroupResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGroupResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleGroupResponseCollectionError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    ruleGroupResponseCollectionErrorJSON `json:"-"`
}

// ruleGroupResponseCollectionErrorJSON contains the JSON metadata for the struct
// [RuleGroupResponseCollectionError]
type ruleGroupResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGroupResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleGroupResponseCollectionMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    ruleGroupResponseCollectionMessageJSON `json:"-"`
}

// ruleGroupResponseCollectionMessageJSON contains the JSON metadata for the struct
// [RuleGroupResponseCollectionMessage]
type ruleGroupResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGroupResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleGroupResponseCollectionResult struct {
	// The unique identifier of the rule group.
	ID string `json:"id,required"`
	// An informative summary of what the rule group does.
	Description string `json:"description,required,nullable"`
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode RuleGroupResponseCollectionResultMode `json:"mode,required"`
	// The name of the rule group.
	Name string `json:"name,required"`
	// The number of rules in the current rule group.
	RulesCount float64 `json:"rules_count,required"`
	// The available states for the rule group.
	AllowedModes []RuleGroupResponseCollectionResultAllowedMode `json:"allowed_modes"`
	// The number of rules within the group that have been modified from their default
	// configuration.
	ModifiedRulesCount float64 `json:"modified_rules_count"`
	// The unique identifier of a WAF package.
	PackageID string                                `json:"package_id"`
	JSON      ruleGroupResponseCollectionResultJSON `json:"-"`
}

// ruleGroupResponseCollectionResultJSON contains the JSON metadata for the struct
// [RuleGroupResponseCollectionResult]
type ruleGroupResponseCollectionResultJSON struct {
	ID                 apijson.Field
	Description        apijson.Field
	Mode               apijson.Field
	Name               apijson.Field
	RulesCount         apijson.Field
	AllowedModes       apijson.Field
	ModifiedRulesCount apijson.Field
	PackageID          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RuleGroupResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type RuleGroupResponseCollectionResultMode string

const (
	RuleGroupResponseCollectionResultModeOn  RuleGroupResponseCollectionResultMode = "on"
	RuleGroupResponseCollectionResultModeOff RuleGroupResponseCollectionResultMode = "off"
)

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type RuleGroupResponseCollectionResultAllowedMode string

const (
	RuleGroupResponseCollectionResultAllowedModeOn  RuleGroupResponseCollectionResultAllowedMode = "on"
	RuleGroupResponseCollectionResultAllowedModeOff RuleGroupResponseCollectionResultAllowedMode = "off"
)

type RuleGroupResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       ruleGroupResponseCollectionResultInfoJSON `json:"-"`
}

// ruleGroupResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [RuleGroupResponseCollectionResultInfo]
type ruleGroupResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGroupResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleGroupResponseCollectionSuccess bool

const (
	RuleGroupResponseCollectionSuccessTrue RuleGroupResponseCollectionSuccess = true
)

type RuleGroupResponseSingle struct {
	Errors   []RuleGroupResponseSingleError   `json:"errors"`
	Messages []RuleGroupResponseSingleMessage `json:"messages"`
	Result   interface{}                      `json:"result"`
	// Whether the API call was successful
	Success RuleGroupResponseSingleSuccess `json:"success"`
	JSON    ruleGroupResponseSingleJSON    `json:"-"`
}

// ruleGroupResponseSingleJSON contains the JSON metadata for the struct
// [RuleGroupResponseSingle]
type ruleGroupResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGroupResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleGroupResponseSingleError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    ruleGroupResponseSingleErrorJSON `json:"-"`
}

// ruleGroupResponseSingleErrorJSON contains the JSON metadata for the struct
// [RuleGroupResponseSingleError]
type ruleGroupResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGroupResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleGroupResponseSingleMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    ruleGroupResponseSingleMessageJSON `json:"-"`
}

// ruleGroupResponseSingleMessageJSON contains the JSON metadata for the struct
// [RuleGroupResponseSingleMessage]
type ruleGroupResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGroupResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleGroupResponseSingleSuccess bool

const (
	RuleGroupResponseSingleSuccessTrue RuleGroupResponseSingleSuccess = true
)

type ZoneFirewallWafPackageGroupUpdateParams struct {
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode param.Field[ZoneFirewallWafPackageGroupUpdateParamsMode] `json:"mode"`
}

func (r ZoneFirewallWafPackageGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type ZoneFirewallWafPackageGroupUpdateParamsMode string

const (
	ZoneFirewallWafPackageGroupUpdateParamsModeOn  ZoneFirewallWafPackageGroupUpdateParamsMode = "on"
	ZoneFirewallWafPackageGroupUpdateParamsModeOff ZoneFirewallWafPackageGroupUpdateParamsMode = "off"
)

type ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParams struct {
	// The direction used to sort returned rule groups.
	Direction param.Field[ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsMatch] `query:"match"`
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode param.Field[ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsMode] `query:"mode"`
	// The field used to sort returned rule groups.
	Order param.Field[ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of rule groups per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParams]'s query
// parameters as `url.Values`.
func (r ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rule groups.
type ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsDirection string

const (
	ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsDirectionAsc  ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsDirection = "asc"
	ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsDirectionDesc ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsDirection = "desc"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsMatch string

const (
	ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsMatchAny ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsMatch = "any"
	ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsMatchAll ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsMatch = "all"
)

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsMode string

const (
	ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsModeOn  ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsMode = "on"
	ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsModeOff ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsMode = "off"
)

// The field used to sort returned rule groups.
type ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsOrder string

const (
	ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsOrderMode       ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsOrder = "mode"
	ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsOrderRulesCount ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParamsOrder = "rules_count"
)
