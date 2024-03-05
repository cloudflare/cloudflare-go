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
)

// FirewallRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallRuleService] method
// instead.
type FirewallRuleService struct {
	Options []option.RequestOption
}

// NewFirewallRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFirewallRuleService(opts ...option.RequestOption) (r *FirewallRuleService) {
	r = &FirewallRuleService{}
	r.Options = opts
	return
}

// Create one or more firewall rules.
func (r *FirewallRuleService) New(ctx context.Context, zoneIdentifier string, body FirewallRuleNewParams, opts ...option.RequestOption) (res *[]LegacyJhsFilterRule, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing firewall rule.
func (r *FirewallRuleService) Update(ctx context.Context, zoneIdentifier string, id string, body FirewallRuleUpdateParams, opts ...option.RequestOption) (res *LegacyJhsFilterRule, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
func (r *FirewallRuleService) List(ctx context.Context, zoneIdentifier string, query FirewallRuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[LegacyJhsFilterRule], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
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

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
func (r *FirewallRuleService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query FirewallRuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[LegacyJhsFilterRule] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing firewall rule.
func (r *FirewallRuleService) Delete(ctx context.Context, zoneIdentifier string, id string, body FirewallRuleDeleteParams, opts ...option.RequestOption) (res *LegacyJhsFilterRule, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the priority of an existing firewall rule.
func (r *FirewallRuleService) Edit(ctx context.Context, zoneIdentifier string, id string, body FirewallRuleEditParams, opts ...option.RequestOption) (res *[]LegacyJhsFilterRule, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a firewall rule.
func (r *FirewallRuleService) Get(ctx context.Context, zoneIdentifier string, id string, query FirewallRuleGetParams, opts ...option.RequestOption) (res *LegacyJhsFilterRule, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LegacyJhsFilterRule struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action LegacyJhsFilterRuleAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                    `json:"description"`
	Filter      LegacyJhsFilterRuleFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                      `json:"priority"`
	Products []LegacyJhsFilterRuleProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                  `json:"ref"`
	JSON legacyJhsFilterRuleJSON `json:"-"`
}

// legacyJhsFilterRuleJSON contains the JSON metadata for the struct
// [LegacyJhsFilterRule]
type legacyJhsFilterRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegacyJhsFilterRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type LegacyJhsFilterRuleAction string

const (
	LegacyJhsFilterRuleActionBlock            LegacyJhsFilterRuleAction = "block"
	LegacyJhsFilterRuleActionChallenge        LegacyJhsFilterRuleAction = "challenge"
	LegacyJhsFilterRuleActionJsChallenge      LegacyJhsFilterRuleAction = "js_challenge"
	LegacyJhsFilterRuleActionManagedChallenge LegacyJhsFilterRuleAction = "managed_challenge"
	LegacyJhsFilterRuleActionAllow            LegacyJhsFilterRuleAction = "allow"
	LegacyJhsFilterRuleActionLog              LegacyJhsFilterRuleAction = "log"
	LegacyJhsFilterRuleActionBypass           LegacyJhsFilterRuleAction = "bypass"
)

// Union satisfied by [LegacyJhsFilter] or
// [LegacyJhsFilterRuleFilterLegacyJhsDeletedFilter].
type LegacyJhsFilterRuleFilter interface {
	implementsLegacyJhsFilterRuleFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*LegacyJhsFilterRuleFilter)(nil)).Elem(), "")
}

type LegacyJhsFilterRuleFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                `json:"deleted,required"`
	JSON    legacyJhsFilterRuleFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// legacyJhsFilterRuleFilterLegacyJhsDeletedFilterJSON contains the JSON metadata
// for the struct [LegacyJhsFilterRuleFilterLegacyJhsDeletedFilter]
type legacyJhsFilterRuleFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegacyJhsFilterRuleFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r LegacyJhsFilterRuleFilterLegacyJhsDeletedFilter) implementsLegacyJhsFilterRuleFilter() {}

// A list of products to bypass for a request when using the `bypass` action.
type LegacyJhsFilterRuleProduct string

const (
	LegacyJhsFilterRuleProductZoneLockdown  LegacyJhsFilterRuleProduct = "zoneLockdown"
	LegacyJhsFilterRuleProductUABlock       LegacyJhsFilterRuleProduct = "uaBlock"
	LegacyJhsFilterRuleProductBic           LegacyJhsFilterRuleProduct = "bic"
	LegacyJhsFilterRuleProductHot           LegacyJhsFilterRuleProduct = "hot"
	LegacyJhsFilterRuleProductSecurityLevel LegacyJhsFilterRuleProduct = "securityLevel"
	LegacyJhsFilterRuleProductRateLimit     LegacyJhsFilterRuleProduct = "rateLimit"
	LegacyJhsFilterRuleProductWAF           LegacyJhsFilterRuleProduct = "waf"
)

type FirewallRuleNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallRuleNewResponseEnvelope struct {
	Errors   []FirewallRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   []LegacyJhsFilterRule                     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallRuleNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallRuleNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallRuleNewResponseEnvelopeJSON       `json:"-"`
}

// firewallRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallRuleNewResponseEnvelope]
type firewallRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    firewallRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallRuleNewResponseEnvelopeErrors]
type firewallRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    firewallRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallRuleNewResponseEnvelopeMessages]
type firewallRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallRuleNewResponseEnvelopeSuccess bool

const (
	FirewallRuleNewResponseEnvelopeSuccessTrue FirewallRuleNewResponseEnvelopeSuccess = true
)

type FirewallRuleNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       firewallRuleNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallRuleNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [FirewallRuleNewResponseEnvelopeResultInfo]
type firewallRuleNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallRuleUpdateResponseEnvelope struct {
	Errors   []FirewallRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsFilterRule                          `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// firewallRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallRuleUpdateResponseEnvelope]
type firewallRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    firewallRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallRuleUpdateResponseEnvelopeErrors]
type firewallRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    firewallRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallRuleUpdateResponseEnvelopeMessages]
type firewallRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallRuleUpdateResponseEnvelopeSuccess bool

const (
	FirewallRuleUpdateResponseEnvelopeSuccessTrue FirewallRuleUpdateResponseEnvelopeSuccess = true
)

type FirewallRuleListParams struct {
	// The action to search for. Must be an exact match.
	Action param.Field[string] `query:"action"`
	// A case-insensitive string to find in the description.
	Description param.Field[string] `query:"description"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// When true, indicates that the firewall rule is currently paused.
	Paused param.Field[bool] `query:"paused"`
	// Number of firewall rules per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [FirewallRuleListParams]'s query parameters as `url.Values`.
func (r FirewallRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallRuleDeleteParams struct {
	// When true, indicates that Cloudflare should also delete the associated filter if
	// there are no other firewall rules referencing the filter.
	DeleteFilterIfUnused param.Field[bool] `json:"delete_filter_if_unused"`
}

func (r FirewallRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallRuleDeleteResponseEnvelope struct {
	Errors   []FirewallRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsFilterRule                          `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// firewallRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallRuleDeleteResponseEnvelope]
type firewallRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    firewallRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallRuleDeleteResponseEnvelopeErrors]
type firewallRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    firewallRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallRuleDeleteResponseEnvelopeMessages]
type firewallRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallRuleDeleteResponseEnvelopeSuccess bool

const (
	FirewallRuleDeleteResponseEnvelopeSuccessTrue FirewallRuleDeleteResponseEnvelopeSuccess = true
)

type FirewallRuleEditParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallRuleEditResponseEnvelope struct {
	Errors   []FirewallRuleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleEditResponseEnvelopeMessages `json:"messages,required"`
	Result   []LegacyJhsFilterRule                      `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallRuleEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallRuleEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallRuleEditResponseEnvelopeJSON       `json:"-"`
}

// firewallRuleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallRuleEditResponseEnvelope]
type firewallRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleEditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    firewallRuleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallRuleEditResponseEnvelopeErrors]
type firewallRuleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleEditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    firewallRuleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallRuleEditResponseEnvelopeMessages]
type firewallRuleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallRuleEditResponseEnvelopeSuccess bool

const (
	FirewallRuleEditResponseEnvelopeSuccessTrue FirewallRuleEditResponseEnvelopeSuccess = true
)

type FirewallRuleEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       firewallRuleEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallRuleEditResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [FirewallRuleEditResponseEnvelopeResultInfo]
type firewallRuleEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleGetParams struct {
}

type FirewallRuleGetResponseEnvelope struct {
	Errors   []FirewallRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsFilterRule                       `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallRuleGetResponseEnvelopeJSON    `json:"-"`
}

// firewallRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallRuleGetResponseEnvelope]
type firewallRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    firewallRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallRuleGetResponseEnvelopeErrors]
type firewallRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    firewallRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallRuleGetResponseEnvelopeMessages]
type firewallRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallRuleGetResponseEnvelopeSuccess bool

const (
	FirewallRuleGetResponseEnvelopeSuccessTrue FirewallRuleGetResponseEnvelopeSuccess = true
)
