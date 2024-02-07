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

// FirewallWAFPackageGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewFirewallWAFPackageGroupService] method instead.
type FirewallWAFPackageGroupService struct {
	Options []option.RequestOption
}

// NewFirewallWAFPackageGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallWAFPackageGroupService(opts ...option.RequestOption) (r *FirewallWAFPackageGroupService) {
	r = &FirewallWAFPackageGroupService{}
	r.Options = opts
	return
}

// Fetches the details of a WAF rule group.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageGroupService) Get(ctx context.Context, zoneID string, packageID string, groupID string, opts ...option.RequestOption) (res *FirewallWAFPackageGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", zoneID, packageID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a WAF rule group. You can update the state (`mode` parameter) of a rule
// group.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageGroupService) Update(ctx context.Context, zoneID string, packageID string, groupID string, body FirewallWAFPackageGroupUpdateParams, opts ...option.RequestOption) (res *FirewallWAFPackageGroupUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", zoneID, packageID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches the WAF rule groups in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageGroupService) List(ctx context.Context, zoneID string, packageID string, query FirewallWAFPackageGroupListParams, opts ...option.RequestOption) (res *FirewallWAFPackageGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups", zoneID, packageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FirewallWAFPackageGroupGetResponse struct {
	Errors   []FirewallWAFPackageGroupGetResponseError   `json:"errors"`
	Messages []FirewallWAFPackageGroupGetResponseMessage `json:"messages"`
	Result   interface{}                                 `json:"result"`
	// Whether the API call was successful
	Success FirewallWAFPackageGroupGetResponseSuccess `json:"success"`
	JSON    firewallWAFPackageGroupGetResponseJSON    `json:"-"`
}

// firewallWAFPackageGroupGetResponseJSON contains the JSON metadata for the struct
// [FirewallWAFPackageGroupGetResponse]
type firewallWAFPackageGroupGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupGetResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    firewallWAFPackageGroupGetResponseErrorJSON `json:"-"`
}

// firewallWAFPackageGroupGetResponseErrorJSON contains the JSON metadata for the
// struct [FirewallWAFPackageGroupGetResponseError]
type firewallWAFPackageGroupGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupGetResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    firewallWAFPackageGroupGetResponseMessageJSON `json:"-"`
}

// firewallWAFPackageGroupGetResponseMessageJSON contains the JSON metadata for the
// struct [FirewallWAFPackageGroupGetResponseMessage]
type firewallWAFPackageGroupGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFPackageGroupGetResponseSuccess bool

const (
	FirewallWAFPackageGroupGetResponseSuccessTrue FirewallWAFPackageGroupGetResponseSuccess = true
)

type FirewallWAFPackageGroupUpdateResponse struct {
	Errors   []FirewallWAFPackageGroupUpdateResponseError   `json:"errors"`
	Messages []FirewallWAFPackageGroupUpdateResponseMessage `json:"messages"`
	Result   interface{}                                    `json:"result"`
	// Whether the API call was successful
	Success FirewallWAFPackageGroupUpdateResponseSuccess `json:"success"`
	JSON    firewallWAFPackageGroupUpdateResponseJSON    `json:"-"`
}

// firewallWAFPackageGroupUpdateResponseJSON contains the JSON metadata for the
// struct [FirewallWAFPackageGroupUpdateResponse]
type firewallWAFPackageGroupUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupUpdateResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    firewallWAFPackageGroupUpdateResponseErrorJSON `json:"-"`
}

// firewallWAFPackageGroupUpdateResponseErrorJSON contains the JSON metadata for
// the struct [FirewallWAFPackageGroupUpdateResponseError]
type firewallWAFPackageGroupUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupUpdateResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    firewallWAFPackageGroupUpdateResponseMessageJSON `json:"-"`
}

// firewallWAFPackageGroupUpdateResponseMessageJSON contains the JSON metadata for
// the struct [FirewallWAFPackageGroupUpdateResponseMessage]
type firewallWAFPackageGroupUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFPackageGroupUpdateResponseSuccess bool

const (
	FirewallWAFPackageGroupUpdateResponseSuccessTrue FirewallWAFPackageGroupUpdateResponseSuccess = true
)

type FirewallWAFPackageGroupListResponse struct {
	Errors     []FirewallWAFPackageGroupListResponseError    `json:"errors"`
	Messages   []FirewallWAFPackageGroupListResponseMessage  `json:"messages"`
	Result     []FirewallWAFPackageGroupListResponseResult   `json:"result"`
	ResultInfo FirewallWAFPackageGroupListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success FirewallWAFPackageGroupListResponseSuccess `json:"success"`
	JSON    firewallWAFPackageGroupListResponseJSON    `json:"-"`
}

// firewallWAFPackageGroupListResponseJSON contains the JSON metadata for the
// struct [FirewallWAFPackageGroupListResponse]
type firewallWAFPackageGroupListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    firewallWAFPackageGroupListResponseErrorJSON `json:"-"`
}

// firewallWAFPackageGroupListResponseErrorJSON contains the JSON metadata for the
// struct [FirewallWAFPackageGroupListResponseError]
type firewallWAFPackageGroupListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    firewallWAFPackageGroupListResponseMessageJSON `json:"-"`
}

// firewallWAFPackageGroupListResponseMessageJSON contains the JSON metadata for
// the struct [FirewallWAFPackageGroupListResponseMessage]
type firewallWAFPackageGroupListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupListResponseResult struct {
	// The unique identifier of the rule group.
	ID string `json:"id,required"`
	// An informative summary of what the rule group does.
	Description string `json:"description,required,nullable"`
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode FirewallWAFPackageGroupListResponseResultMode `json:"mode,required"`
	// The name of the rule group.
	Name string `json:"name,required"`
	// The number of rules in the current rule group.
	RulesCount float64 `json:"rules_count,required"`
	// The available states for the rule group.
	AllowedModes []FirewallWAFPackageGroupListResponseResultAllowedMode `json:"allowed_modes"`
	// The number of rules within the group that have been modified from their default
	// configuration.
	ModifiedRulesCount float64 `json:"modified_rules_count"`
	// The unique identifier of a WAF package.
	PackageID string                                        `json:"package_id"`
	JSON      firewallWAFPackageGroupListResponseResultJSON `json:"-"`
}

// firewallWAFPackageGroupListResponseResultJSON contains the JSON metadata for the
// struct [FirewallWAFPackageGroupListResponseResult]
type firewallWAFPackageGroupListResponseResultJSON struct {
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

func (r *FirewallWAFPackageGroupListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type FirewallWAFPackageGroupListResponseResultMode string

const (
	FirewallWAFPackageGroupListResponseResultModeOn  FirewallWAFPackageGroupListResponseResultMode = "on"
	FirewallWAFPackageGroupListResponseResultModeOff FirewallWAFPackageGroupListResponseResultMode = "off"
)

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type FirewallWAFPackageGroupListResponseResultAllowedMode string

const (
	FirewallWAFPackageGroupListResponseResultAllowedModeOn  FirewallWAFPackageGroupListResponseResultAllowedMode = "on"
	FirewallWAFPackageGroupListResponseResultAllowedModeOff FirewallWAFPackageGroupListResponseResultAllowedMode = "off"
)

type FirewallWAFPackageGroupListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       firewallWAFPackageGroupListResponseResultInfoJSON `json:"-"`
}

// firewallWAFPackageGroupListResponseResultInfoJSON contains the JSON metadata for
// the struct [FirewallWAFPackageGroupListResponseResultInfo]
type firewallWAFPackageGroupListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFPackageGroupListResponseSuccess bool

const (
	FirewallWAFPackageGroupListResponseSuccessTrue FirewallWAFPackageGroupListResponseSuccess = true
)

type FirewallWAFPackageGroupUpdateParams struct {
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode param.Field[FirewallWAFPackageGroupUpdateParamsMode] `json:"mode"`
}

func (r FirewallWAFPackageGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type FirewallWAFPackageGroupUpdateParamsMode string

const (
	FirewallWAFPackageGroupUpdateParamsModeOn  FirewallWAFPackageGroupUpdateParamsMode = "on"
	FirewallWAFPackageGroupUpdateParamsModeOff FirewallWAFPackageGroupUpdateParamsMode = "off"
)

type FirewallWAFPackageGroupListParams struct {
	// The direction used to sort returned rule groups.
	Direction param.Field[FirewallWAFPackageGroupListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[FirewallWAFPackageGroupListParamsMatch] `query:"match"`
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode param.Field[FirewallWAFPackageGroupListParamsMode] `query:"mode"`
	// The field used to sort returned rule groups.
	Order param.Field[FirewallWAFPackageGroupListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of rule groups per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [FirewallWAFPackageGroupListParams]'s query parameters as
// `url.Values`.
func (r FirewallWAFPackageGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rule groups.
type FirewallWAFPackageGroupListParamsDirection string

const (
	FirewallWAFPackageGroupListParamsDirectionAsc  FirewallWAFPackageGroupListParamsDirection = "asc"
	FirewallWAFPackageGroupListParamsDirectionDesc FirewallWAFPackageGroupListParamsDirection = "desc"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type FirewallWAFPackageGroupListParamsMatch string

const (
	FirewallWAFPackageGroupListParamsMatchAny FirewallWAFPackageGroupListParamsMatch = "any"
	FirewallWAFPackageGroupListParamsMatchAll FirewallWAFPackageGroupListParamsMatch = "all"
)

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type FirewallWAFPackageGroupListParamsMode string

const (
	FirewallWAFPackageGroupListParamsModeOn  FirewallWAFPackageGroupListParamsMode = "on"
	FirewallWAFPackageGroupListParamsModeOff FirewallWAFPackageGroupListParamsMode = "off"
)

// The field used to sort returned rule groups.
type FirewallWAFPackageGroupListParamsOrder string

const (
	FirewallWAFPackageGroupListParamsOrderMode       FirewallWAFPackageGroupListParamsOrder = "mode"
	FirewallWAFPackageGroupListParamsOrderRulesCount FirewallWAFPackageGroupListParamsOrder = "rules_count"
)
