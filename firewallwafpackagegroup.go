// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
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

// Fetches the WAF rule groups in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageGroupService) List(ctx context.Context, packageID string, params FirewallWAFPackageGroupListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[FirewallWAFPackageGroupListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups", params.ZoneID, packageID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Fetches the WAF rule groups in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageGroupService) ListAutoPaging(ctx context.Context, packageID string, params FirewallWAFPackageGroupListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[FirewallWAFPackageGroupListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, packageID, params, opts...))
}

// Updates a WAF rule group. You can update the state (`mode` parameter) of a rule
// group.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageGroupService) Edit(ctx context.Context, packageID string, groupID string, params FirewallWAFPackageGroupEditParams, opts ...option.RequestOption) (res *FirewallWAFPackageGroupEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFPackageGroupEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", params.ZoneID, packageID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a WAF rule group.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageGroupService) Get(ctx context.Context, packageID string, groupID string, query FirewallWAFPackageGroupGetParams, opts ...option.RequestOption) (res *FirewallWAFPackageGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFPackageGroupGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", query.ZoneID, packageID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallWAFPackageGroupListResponse struct {
	// The unique identifier of the rule group.
	ID string `json:"id,required"`
	// An informative summary of what the rule group does.
	Description string `json:"description,required,nullable"`
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode FirewallWAFPackageGroupListResponseMode `json:"mode,required"`
	// The name of the rule group.
	Name string `json:"name,required"`
	// The number of rules in the current rule group.
	RulesCount float64 `json:"rules_count,required"`
	// The available states for the rule group.
	AllowedModes []FirewallWAFPackageGroupListResponseAllowedMode `json:"allowed_modes"`
	// The number of rules within the group that have been modified from their default
	// configuration.
	ModifiedRulesCount float64 `json:"modified_rules_count"`
	// The unique identifier of a WAF package.
	PackageID string                                  `json:"package_id"`
	JSON      firewallWAFPackageGroupListResponseJSON `json:"-"`
}

// firewallWAFPackageGroupListResponseJSON contains the JSON metadata for the
// struct [FirewallWAFPackageGroupListResponse]
type firewallWAFPackageGroupListResponseJSON struct {
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

func (r *FirewallWAFPackageGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageGroupListResponseJSON) RawJSON() string {
	return r.raw
}

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type FirewallWAFPackageGroupListResponseMode string

const (
	FirewallWAFPackageGroupListResponseModeOn  FirewallWAFPackageGroupListResponseMode = "on"
	FirewallWAFPackageGroupListResponseModeOff FirewallWAFPackageGroupListResponseMode = "off"
)

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type FirewallWAFPackageGroupListResponseAllowedMode string

const (
	FirewallWAFPackageGroupListResponseAllowedModeOn  FirewallWAFPackageGroupListResponseAllowedMode = "on"
	FirewallWAFPackageGroupListResponseAllowedModeOff FirewallWAFPackageGroupListResponseAllowedMode = "off"
)

// Union satisfied by [FirewallWAFPackageGroupEditResponseUnknown],
// [FirewallWAFPackageGroupEditResponseArray] or [shared.UnionString].
type FirewallWAFPackageGroupEditResponse interface {
	ImplementsFirewallWAFPackageGroupEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallWAFPackageGroupEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallWAFPackageGroupEditResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallWAFPackageGroupEditResponseArray []interface{}

func (r FirewallWAFPackageGroupEditResponseArray) ImplementsFirewallWAFPackageGroupEditResponse() {}

// Union satisfied by [FirewallWAFPackageGroupGetResponseUnknown],
// [FirewallWAFPackageGroupGetResponseArray] or [shared.UnionString].
type FirewallWAFPackageGroupGetResponse interface {
	ImplementsFirewallWAFPackageGroupGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallWAFPackageGroupGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallWAFPackageGroupGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallWAFPackageGroupGetResponseArray []interface{}

func (r FirewallWAFPackageGroupGetResponseArray) ImplementsFirewallWAFPackageGroupGetResponse() {}

type FirewallWAFPackageGroupListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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

type FirewallWAFPackageGroupEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode param.Field[FirewallWAFPackageGroupEditParamsMode] `json:"mode"`
}

func (r FirewallWAFPackageGroupEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type FirewallWAFPackageGroupEditParamsMode string

const (
	FirewallWAFPackageGroupEditParamsModeOn  FirewallWAFPackageGroupEditParamsMode = "on"
	FirewallWAFPackageGroupEditParamsModeOff FirewallWAFPackageGroupEditParamsMode = "off"
)

type FirewallWAFPackageGroupEditResponseEnvelope struct {
	Errors   []FirewallWAFPackageGroupEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFPackageGroupEditResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallWAFPackageGroupEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FirewallWAFPackageGroupEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFPackageGroupEditResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFPackageGroupEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [FirewallWAFPackageGroupEditResponseEnvelope]
type firewallWAFPackageGroupEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageGroupEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallWAFPackageGroupEditResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    firewallWAFPackageGroupEditResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFPackageGroupEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FirewallWAFPackageGroupEditResponseEnvelopeErrors]
type firewallWAFPackageGroupEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageGroupEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallWAFPackageGroupEditResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    firewallWAFPackageGroupEditResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFPackageGroupEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [FirewallWAFPackageGroupEditResponseEnvelopeMessages]
type firewallWAFPackageGroupEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageGroupEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallWAFPackageGroupEditResponseEnvelopeSuccess bool

const (
	FirewallWAFPackageGroupEditResponseEnvelopeSuccessTrue FirewallWAFPackageGroupEditResponseEnvelopeSuccess = true
)

type FirewallWAFPackageGroupGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type FirewallWAFPackageGroupGetResponseEnvelope struct {
	Errors   []FirewallWAFPackageGroupGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFPackageGroupGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallWAFPackageGroupGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FirewallWAFPackageGroupGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFPackageGroupGetResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFPackageGroupGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [FirewallWAFPackageGroupGetResponseEnvelope]
type firewallWAFPackageGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallWAFPackageGroupGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    firewallWAFPackageGroupGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFPackageGroupGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FirewallWAFPackageGroupGetResponseEnvelopeErrors]
type firewallWAFPackageGroupGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageGroupGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallWAFPackageGroupGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    firewallWAFPackageGroupGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFPackageGroupGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [FirewallWAFPackageGroupGetResponseEnvelopeMessages]
type firewallWAFPackageGroupGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageGroupGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallWAFPackageGroupGetResponseEnvelopeSuccess bool

const (
	FirewallWAFPackageGroupGetResponseEnvelopeSuccessTrue FirewallWAFPackageGroupGetResponseEnvelopeSuccess = true
)
