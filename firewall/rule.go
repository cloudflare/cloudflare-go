// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/filters"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/rate_limits"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/tidwall/gjson"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
type RuleService struct {
	Options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	return
}

// Create one or more firewall rules.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) New(ctx context.Context, params RuleNewParams, opts ...option.RequestOption) (res *[]FirewallRule, err error) {
	var env RuleNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing firewall rule.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) Update(ctx context.Context, ruleID string, params RuleUpdateParams, opts ...option.RequestOption) (res *FirewallRule, err error) {
	var env RuleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", params.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) List(ctx context.Context, params RuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[FirewallRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules", params.ZoneID)
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

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) ListAutoPaging(ctx context.Context, params RuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[FirewallRule] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing firewall rule.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) Delete(ctx context.Context, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) (res *FirewallRule, err error) {
	var env RuleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", body.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes existing firewall rules.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) BulkDelete(ctx context.Context, body RuleBulkDeleteParams, opts ...option.RequestOption) (res *[]FirewallRule, err error) {
	var env RuleBulkDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the priority of existing firewall rules.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) BulkEdit(ctx context.Context, params RuleBulkEditParams, opts ...option.RequestOption) (res *[]FirewallRule, err error) {
	var env RuleBulkEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates one or more existing firewall rules.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) BulkUpdate(ctx context.Context, params RuleBulkUpdateParams, opts ...option.RequestOption) (res *[]FirewallRule, err error) {
	var env RuleBulkUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the priority of an existing firewall rule.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) Edit(ctx context.Context, ruleID string, body RuleEditParams, opts ...option.RequestOption) (res *[]FirewallRule, err error) {
	var env RuleEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", body.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a firewall rule.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) Get(ctx context.Context, ruleID string, params RuleGetParams, opts ...option.RequestOption) (res *FirewallRule, err error) {
	var env RuleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", params.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool              `json:"deleted,required"`
	JSON    deletedFilterJSON `json:"-"`
}

// deletedFilterJSON contains the JSON metadata for the struct [DeletedFilter]
type deletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r DeletedFilter) ImplementsFirewallFirewallRuleFilter() {}

type FirewallRule struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action rate_limits.Action `json:"action"`
	// An informative summary of the firewall rule.
	Description string             `json:"description"`
	Filter      FirewallRuleFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64   `json:"priority"`
	Products []Product `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string           `json:"ref"`
	JSON firewallRuleJSON `json:"-"`
}

// firewallRuleJSON contains the JSON metadata for the struct [FirewallRule]
type firewallRuleJSON struct {
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

func (r *FirewallRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool `json:"deleted"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref   string                 `json:"ref"`
	JSON  firewallRuleFilterJSON `json:"-"`
	union FirewallRuleFilterUnion
}

// firewallRuleFilterJSON contains the JSON metadata for the struct
// [FirewallRuleFilter]
type firewallRuleFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r firewallRuleFilterJSON) RawJSON() string {
	return r.raw
}

func (r *FirewallRuleFilter) UnmarshalJSON(data []byte) (err error) {
	*r = FirewallRuleFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FirewallRuleFilterUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [filters.FirewallFilter],
// [firewall.DeletedFilter].
func (r FirewallRuleFilter) AsUnion() FirewallRuleFilterUnion {
	return r.union
}

// Union satisfied by [filters.FirewallFilter] or [firewall.DeletedFilter].
type FirewallRuleFilterUnion interface {
	ImplementsFirewallFirewallRuleFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallRuleFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(filters.FirewallFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeletedFilter{}),
		},
	)
}

// A list of products to bypass for a request when using the `bypass` action.
type Product string

const (
	ProductZoneLockdown  Product = "zoneLockdown"
	ProductUABlock       Product = "uaBlock"
	ProductBIC           Product = "bic"
	ProductHot           Product = "hot"
	ProductSecurityLevel Product = "securityLevel"
	ProductRateLimit     Product = "rateLimit"
	ProductWAF           Product = "waf"
)

func (r Product) IsKnown() bool {
	switch r {
	case ProductZoneLockdown, ProductUABlock, ProductBIC, ProductHot, ProductSecurityLevel, ProductRateLimit, ProductWAF:
		return true
	}
	return false
}

type RuleNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action param.Field[RuleNewParamsAction]         `json:"action,required"`
	Filter param.Field[filters.FirewallFilterParam] `json:"filter,required"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RuleNewParamsAction struct {
	// The action to perform.
	Mode param.Field[RuleNewParamsActionMode] `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response param.Field[RuleNewParamsActionResponse] `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r RuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform.
type RuleNewParamsActionMode string

const (
	RuleNewParamsActionModeSimulate         RuleNewParamsActionMode = "simulate"
	RuleNewParamsActionModeBan              RuleNewParamsActionMode = "ban"
	RuleNewParamsActionModeChallenge        RuleNewParamsActionMode = "challenge"
	RuleNewParamsActionModeJSChallenge      RuleNewParamsActionMode = "js_challenge"
	RuleNewParamsActionModeManagedChallenge RuleNewParamsActionMode = "managed_challenge"
)

func (r RuleNewParamsActionMode) IsKnown() bool {
	switch r {
	case RuleNewParamsActionModeSimulate, RuleNewParamsActionModeBan, RuleNewParamsActionModeChallenge, RuleNewParamsActionModeJSChallenge, RuleNewParamsActionModeManagedChallenge:
		return true
	}
	return false
}

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type RuleNewParamsActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body param.Field[string] `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType param.Field[string] `json:"content_type"`
}

func (r RuleNewParamsActionResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []FirewallRule        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleNewResponseEnvelopeJSON       `json:"-"`
}

// ruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelope]
type ruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleNewResponseEnvelopeSuccess bool

const (
	RuleNewResponseEnvelopeSuccessTrue RuleNewResponseEnvelopeSuccess = true
)

func (r RuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       ruleNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeResultInfo]
type ruleNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action param.Field[RuleUpdateParamsAction]      `json:"action,required"`
	Filter param.Field[filters.FirewallFilterParam] `json:"filter,required"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RuleUpdateParamsAction struct {
	// The action to perform.
	Mode param.Field[RuleUpdateParamsActionMode] `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response param.Field[RuleUpdateParamsActionResponse] `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r RuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform.
type RuleUpdateParamsActionMode string

const (
	RuleUpdateParamsActionModeSimulate         RuleUpdateParamsActionMode = "simulate"
	RuleUpdateParamsActionModeBan              RuleUpdateParamsActionMode = "ban"
	RuleUpdateParamsActionModeChallenge        RuleUpdateParamsActionMode = "challenge"
	RuleUpdateParamsActionModeJSChallenge      RuleUpdateParamsActionMode = "js_challenge"
	RuleUpdateParamsActionModeManagedChallenge RuleUpdateParamsActionMode = "managed_challenge"
)

func (r RuleUpdateParamsActionMode) IsKnown() bool {
	switch r {
	case RuleUpdateParamsActionModeSimulate, RuleUpdateParamsActionModeBan, RuleUpdateParamsActionModeChallenge, RuleUpdateParamsActionModeJSChallenge, RuleUpdateParamsActionModeManagedChallenge:
		return true
	}
	return false
}

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type RuleUpdateParamsActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body param.Field[string] `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType param.Field[string] `json:"content_type"`
}

func (r RuleUpdateParamsActionResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   FirewallRule          `json:"result,required"`
	// Whether the API call was successful
	Success RuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleUpdateResponseEnvelopeJSON    `json:"-"`
}

// ruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelope]
type ruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleUpdateResponseEnvelopeSuccess bool

const (
	RuleUpdateResponseEnvelopeSuccessTrue RuleUpdateResponseEnvelopeSuccess = true
)

func (r RuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The unique identifier of the firewall rule.
	ID param.Field[string] `query:"id"`
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

// URLQuery serializes [RuleListParams]'s query parameters as `url.Values`.
func (r RuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RuleDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RuleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   FirewallRule          `json:"result,required"`
	// Whether the API call was successful
	Success RuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelope]
type ruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleDeleteResponseEnvelopeSuccess bool

const (
	RuleDeleteResponseEnvelopeSuccessTrue RuleDeleteResponseEnvelopeSuccess = true
)

func (r RuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleBulkDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RuleBulkDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []FirewallRule        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleBulkDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleBulkDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleBulkDeleteResponseEnvelopeJSON       `json:"-"`
}

// ruleBulkDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleBulkDeleteResponseEnvelope]
type ruleBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleBulkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleBulkDeleteResponseEnvelopeSuccess bool

const (
	RuleBulkDeleteResponseEnvelopeSuccessTrue RuleBulkDeleteResponseEnvelopeSuccess = true
)

func (r RuleBulkDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleBulkDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleBulkDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       ruleBulkDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleBulkDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [RuleBulkDeleteResponseEnvelopeResultInfo]
type ruleBulkDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleBulkDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleBulkDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleBulkEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r RuleBulkEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleBulkEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []FirewallRule        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleBulkEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleBulkEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleBulkEditResponseEnvelopeJSON       `json:"-"`
}

// ruleBulkEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleBulkEditResponseEnvelope]
type ruleBulkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleBulkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleBulkEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleBulkEditResponseEnvelopeSuccess bool

const (
	RuleBulkEditResponseEnvelopeSuccessTrue RuleBulkEditResponseEnvelopeSuccess = true
)

func (r RuleBulkEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleBulkEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleBulkEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       ruleBulkEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleBulkEditResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [RuleBulkEditResponseEnvelopeResultInfo]
type ruleBulkEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleBulkEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleBulkEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleBulkUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r RuleBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleBulkUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []FirewallRule        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleBulkUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleBulkUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleBulkUpdateResponseEnvelopeJSON       `json:"-"`
}

// ruleBulkUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleBulkUpdateResponseEnvelope]
type ruleBulkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleBulkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleBulkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleBulkUpdateResponseEnvelopeSuccess bool

const (
	RuleBulkUpdateResponseEnvelopeSuccessTrue RuleBulkUpdateResponseEnvelopeSuccess = true
)

func (r RuleBulkUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleBulkUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleBulkUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       ruleBulkUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleBulkUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [RuleBulkUpdateResponseEnvelopeResultInfo]
type ruleBulkUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleBulkUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleBulkUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []FirewallRule        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleEditResponseEnvelopeJSON       `json:"-"`
}

// ruleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelope]
type ruleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleEditResponseEnvelopeSuccess bool

const (
	RuleEditResponseEnvelopeSuccessTrue RuleEditResponseEnvelopeSuccess = true
)

func (r RuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       ruleEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleEditResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeResultInfo]
type ruleEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The unique identifier of the firewall rule.
	ID param.Field[string] `query:"id"`
}

// URLQuery serializes [RuleGetParams]'s query parameters as `url.Values`.
func (r RuleGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RuleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   FirewallRule          `json:"result,required"`
	// Whether the API call was successful
	Success RuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleGetResponseEnvelopeJSON    `json:"-"`
}

// ruleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleGetResponseEnvelope]
type ruleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleGetResponseEnvelopeSuccess bool

const (
	RuleGetResponseEnvelopeSuccessTrue RuleGetResponseEnvelopeSuccess = true
)

func (r RuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
