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

// EmailRoutingRuleCatchAllService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewEmailRoutingRuleCatchAllService] method instead.
type EmailRoutingRuleCatchAllService struct {
	Options []option.RequestOption
}

// NewEmailRoutingRuleCatchAllService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailRoutingRuleCatchAllService(opts ...option.RequestOption) (r *EmailRoutingRuleCatchAllService) {
	r = &EmailRoutingRuleCatchAllService{}
	r.Options = opts
	return
}

// Get information on the default catch-all routing rule.
func (r *EmailRoutingRuleCatchAllService) EmailRoutingRoutingRulesGetCatchAllRule(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable catch-all routing rule, or change action to forward to
// specific destination address.
func (r *EmailRoutingRuleCatchAllService) EmailRoutingRoutingRulesUpdateCatchAllRule(ctx context.Context, zoneIdentifier string, body EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParams, opts ...option.RequestOption) (res *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse struct {
	Errors   []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseError   `json:"errors"`
	Messages []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessage `json:"messages"`
	Result   EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseSuccess `json:"success"`
	JSON    emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseJSON    `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseError struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseErrorJSON `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseErrorJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseError]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessage struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessageJSON `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessage]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResult struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                                                            `json:"tag"`
	JSON emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultJSON `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResult]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Action for the catch-all routing rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultAction struct {
	// Type of action for catch-all rule.
	Type  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsType `json:"type,required"`
	Value []string                                                                                 `json:"value"`
	JSON  emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionJSON  `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultAction]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of action for catch-all rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsType string

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsTypeDrop    EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsType = "drop"
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsTypeForward EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsType = "forward"
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsTypeWorker  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultEnabled bool

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultEnabledTrue  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultEnabled = true
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultEnabledFalse EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultEnabled = false
)

// Matcher for catch-all routing rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcher struct {
	// Type of matcher. Default is 'all'.
	Type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatchersType `json:"type,required"`
	JSON emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcherJSON  `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcherJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcher]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of matcher. Default is 'all'.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatchersType string

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatchersTypeAll EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseResultMatchersType = "all"
)

// Whether the API call was successful
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseSuccess bool

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseSuccessTrue EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseSuccess = true
)

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse struct {
	Errors   []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseError   `json:"errors"`
	Messages []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessage `json:"messages"`
	Result   EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseSuccess `json:"success"`
	JSON    emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseJSON    `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseError struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseErrorJSON `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseErrorJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseError]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessage struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessageJSON `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessage]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResult struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                                                               `json:"tag"`
	JSON emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultJSON `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResult]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Action for the catch-all routing rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultAction struct {
	// Type of action for catch-all rule.
	Type  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsType `json:"type,required"`
	Value []string                                                                                    `json:"value"`
	JSON  emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionJSON  `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultAction]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of action for catch-all rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsType string

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsTypeDrop    EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsType = "drop"
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsTypeForward EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsType = "forward"
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsTypeWorker  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultEnabled bool

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultEnabledTrue  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultEnabled = true
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultEnabledFalse EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultEnabled = false
)

// Matcher for catch-all routing rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcher struct {
	// Type of matcher. Default is 'all'.
	Type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatchersType `json:"type,required"`
	JSON emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcherJSON  `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcherJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcher]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of matcher. Default is 'all'.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatchersType string

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatchersTypeAll EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseResultMatchersType = "all"
)

// Whether the API call was successful
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseSuccess bool

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseSuccessTrue EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseSuccess = true
)

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParams struct {
	// List actions for the catch-all routing rule.
	Actions param.Field[[]EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsAction] `json:"actions,required"`
	// List of matchers for the catch-all routing rule.
	Matchers param.Field[[]EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
}

func (r EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action for the catch-all routing rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsAction struct {
	// Type of action for catch-all rule.
	Type  param.Field[EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                                                            `json:"value"`
}

func (r EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of action for catch-all rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsType string

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsTypeDrop    EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsType = "drop"
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsTypeForward EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsType = "forward"
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsTypeWorker  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsActionsType = "worker"
)

// Matcher for catch-all routing rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatcher struct {
	// Type of matcher. Default is 'all'.
	Type param.Field[EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersType] `json:"type,required"`
}

func (r EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of matcher. Default is 'all'.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersType string

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersTypeAll EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsMatchersType = "all"
)

// Routing rule status.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabled bool

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabledTrue  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabled = true
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabledFalse EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParamsEnabled = false
)
