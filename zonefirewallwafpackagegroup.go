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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
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
func (r *ZoneFirewallWafPackageGroupService) Get(ctx context.Context, zoneIdentifier string, packageIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneFirewallWafPackageGroupGetResponse, err error) {
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
func (r *ZoneFirewallWafPackageGroupService) Update(ctx context.Context, zoneIdentifier string, packageIdentifier string, identifier string, body ZoneFirewallWafPackageGroupUpdateParams, opts ...option.RequestOption) (res *ZoneFirewallWafPackageGroupUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", zoneIdentifier, packageIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches the WAF rule groups in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageGroupService) WafRuleGroupsListWafRuleGroups(ctx context.Context, zoneIdentifier string, packageIdentifier string, query ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsParams, opts ...option.RequestOption) (res *shared.Page[ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups", zoneIdentifier, packageIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

type ZoneFirewallWafPackageGroupGetResponse struct {
	Errors   []ZoneFirewallWafPackageGroupGetResponseError   `json:"errors"`
	Messages []ZoneFirewallWafPackageGroupGetResponseMessage `json:"messages"`
	Result   interface{}                                     `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallWafPackageGroupGetResponseSuccess `json:"success"`
	JSON    zoneFirewallWafPackageGroupGetResponseJSON    `json:"-"`
}

// zoneFirewallWafPackageGroupGetResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafPackageGroupGetResponse]
type zoneFirewallWafPackageGroupGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafPackageGroupGetResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneFirewallWafPackageGroupGetResponseErrorJSON `json:"-"`
}

// zoneFirewallWafPackageGroupGetResponseErrorJSON contains the JSON metadata for
// the struct [ZoneFirewallWafPackageGroupGetResponseError]
type zoneFirewallWafPackageGroupGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageGroupGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafPackageGroupGetResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneFirewallWafPackageGroupGetResponseMessageJSON `json:"-"`
}

// zoneFirewallWafPackageGroupGetResponseMessageJSON contains the JSON metadata for
// the struct [ZoneFirewallWafPackageGroupGetResponseMessage]
type zoneFirewallWafPackageGroupGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageGroupGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFirewallWafPackageGroupGetResponseSuccess bool

const (
	ZoneFirewallWafPackageGroupGetResponseSuccessTrue ZoneFirewallWafPackageGroupGetResponseSuccess = true
)

type ZoneFirewallWafPackageGroupUpdateResponse struct {
	Errors   []ZoneFirewallWafPackageGroupUpdateResponseError   `json:"errors"`
	Messages []ZoneFirewallWafPackageGroupUpdateResponseMessage `json:"messages"`
	Result   interface{}                                        `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallWafPackageGroupUpdateResponseSuccess `json:"success"`
	JSON    zoneFirewallWafPackageGroupUpdateResponseJSON    `json:"-"`
}

// zoneFirewallWafPackageGroupUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafPackageGroupUpdateResponse]
type zoneFirewallWafPackageGroupUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafPackageGroupUpdateResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneFirewallWafPackageGroupUpdateResponseErrorJSON `json:"-"`
}

// zoneFirewallWafPackageGroupUpdateResponseErrorJSON contains the JSON metadata
// for the struct [ZoneFirewallWafPackageGroupUpdateResponseError]
type zoneFirewallWafPackageGroupUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageGroupUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafPackageGroupUpdateResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneFirewallWafPackageGroupUpdateResponseMessageJSON `json:"-"`
}

// zoneFirewallWafPackageGroupUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneFirewallWafPackageGroupUpdateResponseMessage]
type zoneFirewallWafPackageGroupUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageGroupUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFirewallWafPackageGroupUpdateResponseSuccess bool

const (
	ZoneFirewallWafPackageGroupUpdateResponseSuccessTrue ZoneFirewallWafPackageGroupUpdateResponseSuccess = true
)

type ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponse struct {
	// The unique identifier of the rule group.
	ID string `json:"id,required"`
	// An informative summary of what the rule group does.
	Description string `json:"description,required,nullable"`
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseMode `json:"mode,required"`
	// The name of the rule group.
	Name string `json:"name,required"`
	// The number of rules in the current rule group.
	RulesCount float64 `json:"rules_count,required"`
	// The available states for the rule group.
	AllowedModes []ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseAllowedMode `json:"allowed_modes"`
	// The number of rules within the group that have been modified from their default
	// configuration.
	ModifiedRulesCount float64 `json:"modified_rules_count"`
	// The unique identifier of a WAF package.
	PackageID string                                                                `json:"package_id"`
	JSON      zoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseJSON `json:"-"`
}

// zoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseJSON contains
// the JSON metadata for the struct
// [ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponse]
type zoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseJSON struct {
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

func (r *ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseMode string

const (
	ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseModeOn  ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseMode = "on"
	ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseModeOff ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseMode = "off"
)

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseAllowedMode string

const (
	ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseAllowedModeOn  ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseAllowedMode = "on"
	ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseAllowedModeOff ZoneFirewallWafPackageGroupWafRuleGroupsListWafRuleGroupsResponseAllowedMode = "off"
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
