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

// ZoneEmailRoutingRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneEmailRoutingRuleService]
// method instead.
type ZoneEmailRoutingRuleService struct {
	Options   []option.RequestOption
	CatchAlls *ZoneEmailRoutingRuleCatchAllService
}

// NewZoneEmailRoutingRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneEmailRoutingRuleService(opts ...option.RequestOption) (r *ZoneEmailRoutingRuleService) {
	r = &ZoneEmailRoutingRuleService{}
	r.Options = opts
	r.CatchAlls = NewZoneEmailRoutingRuleCatchAllService(opts...)
	return
}

// Get information for a specific routing rule already created.
func (r *ZoneEmailRoutingRuleService) Get(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *RuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update actions and matches, or enable/disable specific routing rules.
func (r *ZoneEmailRoutingRuleService) Update(ctx context.Context, zoneIdentifier string, ruleIdentifier string, body ZoneEmailRoutingRuleUpdateParams, opts ...option.RequestOption) (res *RuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a specific routing rule.
func (r *ZoneEmailRoutingRuleService) Delete(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *RuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Rules consist of a set of criteria for matching emails (such as an email being
// sent to a specific custom email address) plus a set of actions to take on the
// email (like forwarding it to a specific destination address).
func (r *ZoneEmailRoutingRuleService) EmailRoutingRoutingRulesNewRoutingRule(ctx context.Context, zoneIdentifier string, body ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams, opts ...option.RequestOption) (res *RuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists existing routing rules.
func (r *ZoneEmailRoutingRuleService) EmailRoutingRoutingRulesListRoutingRules(ctx context.Context, zoneIdentifier string, query ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams, opts ...option.RequestOption) (res *RulesResponseCollection7vDQjAdp, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RuleResponseSingle struct {
	Errors   []RuleResponseSingleError   `json:"errors"`
	Messages []RuleResponseSingleMessage `json:"messages"`
	Result   RuleResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success RuleResponseSingleSuccess `json:"success"`
	JSON    ruleResponseSingleJSON    `json:"-"`
}

// ruleResponseSingleJSON contains the JSON metadata for the struct
// [RuleResponseSingle]
type ruleResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleResponseSingleError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    ruleResponseSingleErrorJSON `json:"-"`
}

// ruleResponseSingleErrorJSON contains the JSON metadata for the struct
// [RuleResponseSingleError]
type ruleResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleResponseSingleMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    ruleResponseSingleMessageJSON `json:"-"`
}

// ruleResponseSingleMessageJSON contains the JSON metadata for the struct
// [RuleResponseSingleMessage]
type ruleResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleResponseSingleResult struct {
	// List actions patterns.
	Actions []RuleResponseSingleResultAction `json:"actions"`
	// Routing rule status.
	Enabled RuleResponseSingleResultEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []RuleResponseSingleResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule identifier.
	Tag  string                       `json:"tag"`
	JSON ruleResponseSingleResultJSON `json:"-"`
}

// ruleResponseSingleResultJSON contains the JSON metadata for the struct
// [RuleResponseSingleResult]
type ruleResponseSingleResultJSON struct {
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type RuleResponseSingleResultAction struct {
	// Type of supported action.
	Type  RuleResponseSingleResultActionsType `json:"type,required"`
	Value []string                            `json:"value,required"`
	JSON  ruleResponseSingleResultActionJSON  `json:"-"`
}

// ruleResponseSingleResultActionJSON contains the JSON metadata for the struct
// [RuleResponseSingleResultAction]
type ruleResponseSingleResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseSingleResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type RuleResponseSingleResultActionsType string

const (
	RuleResponseSingleResultActionsTypeForward RuleResponseSingleResultActionsType = "forward"
	RuleResponseSingleResultActionsTypeWorker  RuleResponseSingleResultActionsType = "worker"
)

// Routing rule status.
type RuleResponseSingleResultEnabled bool

const (
	RuleResponseSingleResultEnabledTrue  RuleResponseSingleResultEnabled = true
	RuleResponseSingleResultEnabledFalse RuleResponseSingleResultEnabled = false
)

// Matching pattern to forward your actions.
type RuleResponseSingleResultMatcher struct {
	// Field for type matcher.
	Field RuleResponseSingleResultMatchersField `json:"field,required"`
	// Type of matcher.
	Type RuleResponseSingleResultMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                              `json:"value,required"`
	JSON  ruleResponseSingleResultMatcherJSON `json:"-"`
}

// ruleResponseSingleResultMatcherJSON contains the JSON metadata for the struct
// [RuleResponseSingleResultMatcher]
type ruleResponseSingleResultMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseSingleResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type RuleResponseSingleResultMatchersField string

const (
	RuleResponseSingleResultMatchersFieldTo RuleResponseSingleResultMatchersField = "to"
)

// Type of matcher.
type RuleResponseSingleResultMatchersType string

const (
	RuleResponseSingleResultMatchersTypeLiteral RuleResponseSingleResultMatchersType = "literal"
)

// Whether the API call was successful
type RuleResponseSingleSuccess bool

const (
	RuleResponseSingleSuccessTrue RuleResponseSingleSuccess = true
)

type RulesResponseCollection7vDQjAdp struct {
	Errors     []RulesResponseCollection7vDQjAdpError    `json:"errors"`
	Messages   []RulesResponseCollection7vDQjAdpMessage  `json:"messages"`
	Result     []RulesResponseCollection7vDQjAdpResult   `json:"result"`
	ResultInfo RulesResponseCollection7vDQjAdpResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success RulesResponseCollection7vDQjAdpSuccess `json:"success"`
	JSON    rulesResponseCollection7vDQjAdpJSON    `json:"-"`
}

// rulesResponseCollection7vDQjAdpJSON contains the JSON metadata for the struct
// [RulesResponseCollection7vDQjAdp]
type rulesResponseCollection7vDQjAdpJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesResponseCollection7vDQjAdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesResponseCollection7vDQjAdpError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    rulesResponseCollection7vDQjAdpErrorJSON `json:"-"`
}

// rulesResponseCollection7vDQjAdpErrorJSON contains the JSON metadata for the
// struct [RulesResponseCollection7vDQjAdpError]
type rulesResponseCollection7vDQjAdpErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesResponseCollection7vDQjAdpError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesResponseCollection7vDQjAdpMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    rulesResponseCollection7vDQjAdpMessageJSON `json:"-"`
}

// rulesResponseCollection7vDQjAdpMessageJSON contains the JSON metadata for the
// struct [RulesResponseCollection7vDQjAdpMessage]
type rulesResponseCollection7vDQjAdpMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesResponseCollection7vDQjAdpMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesResponseCollection7vDQjAdpResult struct {
	// List actions patterns.
	Actions []RulesResponseCollection7vDQjAdpResultAction `json:"actions"`
	// Routing rule status.
	Enabled RulesResponseCollection7vDQjAdpResultEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []RulesResponseCollection7vDQjAdpResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule identifier.
	Tag  string                                    `json:"tag"`
	JSON rulesResponseCollection7vDQjAdpResultJSON `json:"-"`
}

// rulesResponseCollection7vDQjAdpResultJSON contains the JSON metadata for the
// struct [RulesResponseCollection7vDQjAdpResult]
type rulesResponseCollection7vDQjAdpResultJSON struct {
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesResponseCollection7vDQjAdpResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type RulesResponseCollection7vDQjAdpResultAction struct {
	// Type of supported action.
	Type  RulesResponseCollection7vDQjAdpResultActionsType `json:"type,required"`
	Value []string                                         `json:"value,required"`
	JSON  rulesResponseCollection7vDQjAdpResultActionJSON  `json:"-"`
}

// rulesResponseCollection7vDQjAdpResultActionJSON contains the JSON metadata for
// the struct [RulesResponseCollection7vDQjAdpResultAction]
type rulesResponseCollection7vDQjAdpResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesResponseCollection7vDQjAdpResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type RulesResponseCollection7vDQjAdpResultActionsType string

const (
	RulesResponseCollection7vDQjAdpResultActionsTypeForward RulesResponseCollection7vDQjAdpResultActionsType = "forward"
	RulesResponseCollection7vDQjAdpResultActionsTypeWorker  RulesResponseCollection7vDQjAdpResultActionsType = "worker"
)

// Routing rule status.
type RulesResponseCollection7vDQjAdpResultEnabled bool

const (
	RulesResponseCollection7vDQjAdpResultEnabledTrue  RulesResponseCollection7vDQjAdpResultEnabled = true
	RulesResponseCollection7vDQjAdpResultEnabledFalse RulesResponseCollection7vDQjAdpResultEnabled = false
)

// Matching pattern to forward your actions.
type RulesResponseCollection7vDQjAdpResultMatcher struct {
	// Field for type matcher.
	Field RulesResponseCollection7vDQjAdpResultMatchersField `json:"field,required"`
	// Type of matcher.
	Type RulesResponseCollection7vDQjAdpResultMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                           `json:"value,required"`
	JSON  rulesResponseCollection7vDQjAdpResultMatcherJSON `json:"-"`
}

// rulesResponseCollection7vDQjAdpResultMatcherJSON contains the JSON metadata for
// the struct [RulesResponseCollection7vDQjAdpResultMatcher]
type rulesResponseCollection7vDQjAdpResultMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesResponseCollection7vDQjAdpResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type RulesResponseCollection7vDQjAdpResultMatchersField string

const (
	RulesResponseCollection7vDQjAdpResultMatchersFieldTo RulesResponseCollection7vDQjAdpResultMatchersField = "to"
)

// Type of matcher.
type RulesResponseCollection7vDQjAdpResultMatchersType string

const (
	RulesResponseCollection7vDQjAdpResultMatchersTypeLiteral RulesResponseCollection7vDQjAdpResultMatchersType = "literal"
)

type RulesResponseCollection7vDQjAdpResultInfo struct {
	Count      interface{}                                   `json:"count"`
	Page       interface{}                                   `json:"page"`
	PerPage    interface{}                                   `json:"per_page"`
	TotalCount interface{}                                   `json:"total_count"`
	JSON       rulesResponseCollection7vDQjAdpResultInfoJSON `json:"-"`
}

// rulesResponseCollection7vDQjAdpResultInfoJSON contains the JSON metadata for the
// struct [RulesResponseCollection7vDQjAdpResultInfo]
type rulesResponseCollection7vDQjAdpResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesResponseCollection7vDQjAdpResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RulesResponseCollection7vDQjAdpSuccess bool

const (
	RulesResponseCollection7vDQjAdpSuccessTrue RulesResponseCollection7vDQjAdpSuccess = true
)

type ZoneEmailRoutingRuleUpdateParams struct {
	// List actions patterns.
	Actions param.Field[[]ZoneEmailRoutingRuleUpdateParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]ZoneEmailRoutingRuleUpdateParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[ZoneEmailRoutingRuleUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r ZoneEmailRoutingRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type ZoneEmailRoutingRuleUpdateParamsAction struct {
	// Type of supported action.
	Type  param.Field[ZoneEmailRoutingRuleUpdateParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                    `json:"value,required"`
}

func (r ZoneEmailRoutingRuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type ZoneEmailRoutingRuleUpdateParamsActionsType string

const (
	ZoneEmailRoutingRuleUpdateParamsActionsTypeForward ZoneEmailRoutingRuleUpdateParamsActionsType = "forward"
	ZoneEmailRoutingRuleUpdateParamsActionsTypeWorker  ZoneEmailRoutingRuleUpdateParamsActionsType = "worker"
)

// Matching pattern to forward your actions.
type ZoneEmailRoutingRuleUpdateParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[ZoneEmailRoutingRuleUpdateParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[ZoneEmailRoutingRuleUpdateParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r ZoneEmailRoutingRuleUpdateParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type ZoneEmailRoutingRuleUpdateParamsMatchersField string

const (
	ZoneEmailRoutingRuleUpdateParamsMatchersFieldTo ZoneEmailRoutingRuleUpdateParamsMatchersField = "to"
)

// Type of matcher.
type ZoneEmailRoutingRuleUpdateParamsMatchersType string

const (
	ZoneEmailRoutingRuleUpdateParamsMatchersTypeLiteral ZoneEmailRoutingRuleUpdateParamsMatchersType = "literal"
)

// Routing rule status.
type ZoneEmailRoutingRuleUpdateParamsEnabled bool

const (
	ZoneEmailRoutingRuleUpdateParamsEnabledTrue  ZoneEmailRoutingRuleUpdateParamsEnabled = true
	ZoneEmailRoutingRuleUpdateParamsEnabledFalse ZoneEmailRoutingRuleUpdateParamsEnabled = false
)

type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams struct {
	// List actions patterns.
	Actions param.Field[[]ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsAction struct {
	// Type of supported action.
	Type  param.Field[ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                                                    `json:"value,required"`
}

func (r ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType string

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsTypeForward ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType = "forward"
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsTypeWorker  ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType = "worker"
)

// Matching pattern to forward your actions.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersField string

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersFieldTo ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersField = "to"
)

// Type of matcher.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersType string

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersTypeLiteral ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersType = "literal"
)

// Routing rule status.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabled bool

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabledTrue  ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabled = true
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabledFalse ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabled = false
)

type ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams struct {
	// Filter by enabled routing rules.
	Enabled param.Field[ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabled] `query:"enabled"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams]'s query
// parameters as `url.Values`.
func (r ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by enabled routing rules.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabled bool

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabledTrue  ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabled = true
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabledFalse ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabled = false
)
