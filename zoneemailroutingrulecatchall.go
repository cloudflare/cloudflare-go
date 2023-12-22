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

// ZoneEmailRoutingRuleCatchAllService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneEmailRoutingRuleCatchAllService] method instead.
type ZoneEmailRoutingRuleCatchAllService struct {
	Options []option.RequestOption
}

// NewZoneEmailRoutingRuleCatchAllService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneEmailRoutingRuleCatchAllService(opts ...option.RequestOption) (r *ZoneEmailRoutingRuleCatchAllService) {
	r = &ZoneEmailRoutingRuleCatchAllService{}
	r.Options = opts
	return
}

// Get information on the default catch-all routing rule.
func (r *ZoneEmailRoutingRuleCatchAllService) EmailRoutingRoutingRulesGetCatchAllRule(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *CatchAllRuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable catch-all routing rule, or change action to forward to
// specific destination address.
func (r *ZoneEmailRoutingRuleCatchAllService) EmailRoutingRoutingRulesUpdateCatchAllRule(ctx context.Context, zoneIdentifier string, body ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParams, opts ...option.RequestOption) (res *CatchAllRuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CatchAllRuleResponseSingle struct {
	Errors   []CatchAllRuleResponseSingleError   `json:"errors"`
	Messages []CatchAllRuleResponseSingleMessage `json:"messages"`
	Result   CatchAllRuleResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success CatchAllRuleResponseSingleSuccess `json:"success"`
	JSON    catchAllRuleResponseSingleJSON    `json:"-"`
}

// catchAllRuleResponseSingleJSON contains the JSON metadata for the struct
// [CatchAllRuleResponseSingle]
type catchAllRuleResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatchAllRuleResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CatchAllRuleResponseSingleError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    catchAllRuleResponseSingleErrorJSON `json:"-"`
}

// catchAllRuleResponseSingleErrorJSON contains the JSON metadata for the struct
// [CatchAllRuleResponseSingleError]
type catchAllRuleResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatchAllRuleResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CatchAllRuleResponseSingleMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    catchAllRuleResponseSingleMessageJSON `json:"-"`
}

// catchAllRuleResponseSingleMessageJSON contains the JSON metadata for the struct
// [CatchAllRuleResponseSingleMessage]
type catchAllRuleResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatchAllRuleResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CatchAllRuleResponseSingleResult struct {
	// List actions for the catch-all routing rule.
	Actions []CatchAllRuleResponseSingleResultAction `json:"actions"`
	// Routing rule status.
	Enabled CatchAllRuleResponseSingleResultEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []CatchAllRuleResponseSingleResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule identifier.
	Tag  string                               `json:"tag"`
	JSON catchAllRuleResponseSingleResultJSON `json:"-"`
}

// catchAllRuleResponseSingleResultJSON contains the JSON metadata for the struct
// [CatchAllRuleResponseSingleResult]
type catchAllRuleResponseSingleResultJSON struct {
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatchAllRuleResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Action for the catch-all routing rule.
type CatchAllRuleResponseSingleResultAction struct {
	// Type of action for catch-all rule.
	Type  CatchAllRuleResponseSingleResultActionsType `json:"type,required"`
	Value []string                                    `json:"value"`
	JSON  catchAllRuleResponseSingleResultActionJSON  `json:"-"`
}

// catchAllRuleResponseSingleResultActionJSON contains the JSON metadata for the
// struct [CatchAllRuleResponseSingleResultAction]
type catchAllRuleResponseSingleResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatchAllRuleResponseSingleResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of action for catch-all rule.
type CatchAllRuleResponseSingleResultActionsType string

const (
	CatchAllRuleResponseSingleResultActionsTypeDrop    CatchAllRuleResponseSingleResultActionsType = "drop"
	CatchAllRuleResponseSingleResultActionsTypeForward CatchAllRuleResponseSingleResultActionsType = "forward"
	CatchAllRuleResponseSingleResultActionsTypeWorker  CatchAllRuleResponseSingleResultActionsType = "worker"
)

// Routing rule status.
type CatchAllRuleResponseSingleResultEnabled bool

const (
	CatchAllRuleResponseSingleResultEnabledTrue  CatchAllRuleResponseSingleResultEnabled = true
	CatchAllRuleResponseSingleResultEnabledFalse CatchAllRuleResponseSingleResultEnabled = false
)

// Matcher for catch-all routing rule.
type CatchAllRuleResponseSingleResultMatcher struct {
	// Type of matcher. Default is 'all'.
	Type CatchAllRuleResponseSingleResultMatchersType `json:"type,required"`
	JSON catchAllRuleResponseSingleResultMatcherJSON  `json:"-"`
}

// catchAllRuleResponseSingleResultMatcherJSON contains the JSON metadata for the
// struct [CatchAllRuleResponseSingleResultMatcher]
type catchAllRuleResponseSingleResultMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatchAllRuleResponseSingleResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of matcher. Default is 'all'.
type CatchAllRuleResponseSingleResultMatchersType string

const (
	CatchAllRuleResponseSingleResultMatchersTypeAll CatchAllRuleResponseSingleResultMatchersType = "all"
)

// Whether the API call was successful
type CatchAllRuleResponseSingleSuccess bool

const (
	CatchAllRuleResponseSingleSuccessTrue CatchAllRuleResponseSingleSuccess = true
)

type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParams struct {
	// List actions for the catch-all routing rule.
	Actions param.Field[[]ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsAction] `json:"actions,required"`
	// List of matchers for the catch-all routing rule.
	Matchers param.Field[[]ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
}

func (r ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action for the catch-all routing rule.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsAction struct {
	// Type of action for catch-all rule.
	Type  param.Field[ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                                                                `json:"value"`
}

func (r ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of action for catch-all rule.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsType string

const (
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsTypeDrop    ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsType = "drop"
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsTypeForward ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsType = "forward"
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsTypeWorker  ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsType = "worker"
)

// Matcher for catch-all routing rule.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatcher struct {
	// Type of matcher. Default is 'all'.
	Type param.Field[ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersType] `json:"type,required"`
}

func (r ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of matcher. Default is 'all'.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersType string

const (
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersTypeAll ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersType = "all"
)

// Routing rule status.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabled bool

const (
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabledTrue  ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabled = true
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabledFalse ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabled = false
)
