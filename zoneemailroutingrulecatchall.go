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
func (r *ZoneEmailRoutingRuleCatchAllService) EmailRoutingRoutingRulesGetCatchAllRule(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable catch-all routing rule, or change action to forward to
// specific destination address.
func (r *ZoneEmailRoutingRuleCatchAllService) EmailRoutingRoutingRulesUpdateCatchAllRule(ctx context.Context, zoneIdentifier string, body ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParams, opts ...option.RequestOption) (res *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse struct {
	Errors   []ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseError   `json:"errors"`
	Messages []ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessage `json:"messages"`
	Result   ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseSuccess `json:"success"`
	JSON    zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseJSON    `json:"-"`
}

// zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse]
type zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseError struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseErrorJSON `json:"-"`
}

// zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseError]
type zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessage struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessageJSON `json:"-"`
}

// zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessage]
type zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResult struct {
	// List actions for the catch-all routing rule.
	Actions []ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule identifier.
	Tag  string                                                                                `json:"tag"`
	JSON zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultJSON `json:"-"`
}

// zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResult]
type zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultJSON struct {
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Action for the catch-all routing rule.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultAction struct {
	// Type of action for catch-all rule.
	Type  ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsType `json:"type,required"`
	Value []string                                                                                     `json:"value"`
	JSON  zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionJSON  `json:"-"`
}

// zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultAction]
type zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of action for catch-all rule.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsType string

const (
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsTypeDrop    ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsType = "drop"
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsTypeForward ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsType = "forward"
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsTypeWorker  ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsType = "worker"
)

// Routing rule status.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultEnabled bool

const (
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultEnabledTrue  ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultEnabled = true
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultEnabledFalse ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultEnabled = false
)

// Matcher for catch-all routing rule.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcher struct {
	// Type of matcher. Default is 'all'.
	Type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatchersType `json:"type,required"`
	JSON zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcherJSON  `json:"-"`
}

// zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcherJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcher]
type zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of matcher. Default is 'all'.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatchersType string

const (
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatchersTypeAll ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatchersType = "all"
)

// Whether the API call was successful
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseSuccess bool

const (
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseSuccessTrue ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseSuccess = true
)

type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse struct {
	Errors   []ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseError   `json:"errors"`
	Messages []ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessage `json:"messages"`
	Result   ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseSuccess `json:"success"`
	JSON    zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseJSON    `json:"-"`
}

// zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse]
type zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseError struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseErrorJSON `json:"-"`
}

// zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseError]
type zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessage struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessageJSON `json:"-"`
}

// zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessage]
type zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResult struct {
	// List actions for the catch-all routing rule.
	Actions []ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule identifier.
	Tag  string                                                                                   `json:"tag"`
	JSON zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultJSON `json:"-"`
}

// zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResult]
type zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultJSON struct {
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Action for the catch-all routing rule.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultAction struct {
	// Type of action for catch-all rule.
	Type  ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsType `json:"type,required"`
	Value []string                                                                                        `json:"value"`
	JSON  zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionJSON  `json:"-"`
}

// zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultAction]
type zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of action for catch-all rule.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsType string

const (
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsTypeDrop    ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsType = "drop"
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsTypeForward ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsType = "forward"
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsTypeWorker  ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsType = "worker"
)

// Routing rule status.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultEnabled bool

const (
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultEnabledTrue  ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultEnabled = true
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultEnabledFalse ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultEnabled = false
)

// Matcher for catch-all routing rule.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcher struct {
	// Type of matcher. Default is 'all'.
	Type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatchersType `json:"type,required"`
	JSON zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcherJSON  `json:"-"`
}

// zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcherJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcher]
type zoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of matcher. Default is 'all'.
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatchersType string

const (
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatchersTypeAll ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatchersType = "all"
)

// Whether the API call was successful
type ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseSuccess bool

const (
	ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseSuccessTrue ZoneEmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseSuccess = true
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
