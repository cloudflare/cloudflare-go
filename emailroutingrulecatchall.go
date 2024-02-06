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
	var env EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable or disable catch-all routing rule, or change action to forward to
// specific destination address.
func (r *EmailRoutingRuleCatchAllService) EmailRoutingRoutingRulesUpdateCatchAllRule(ctx context.Context, zoneIdentifier string, body EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleParams, opts ...option.RequestOption) (res *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                                                      `json:"tag"`
	JSON emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseJSON `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Action for the catch-all routing rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseAction struct {
	// Type of action for catch-all rule.
	Type  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseActionsType `json:"type,required"`
	Value []string                                                                           `json:"value"`
	JSON  emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseActionJSON  `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseActionJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseAction]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of action for catch-all rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseActionsType string

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseActionsTypeDrop    EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseActionsType = "drop"
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseActionsTypeForward EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseActionsType = "forward"
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseActionsTypeWorker  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnabled bool

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnabledTrue  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnabled = true
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnabledFalse EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnabled = false
)

// Matcher for catch-all routing rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMatcher struct {
	// Type of matcher. Default is 'all'.
	Type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMatchersType `json:"type,required"`
	JSON emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMatcherJSON  `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMatcherJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMatcher]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of matcher. Default is 'all'.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMatchersType string

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMatchersTypeAll EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseMatchersType = "all"
)

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                                                         `json:"tag"`
	JSON emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseJSON `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Action for the catch-all routing rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseAction struct {
	// Type of action for catch-all rule.
	Type  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseActionsType `json:"type,required"`
	Value []string                                                                              `json:"value"`
	JSON  emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseActionJSON  `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseActionJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseAction]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of action for catch-all rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseActionsType string

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseActionsTypeDrop    EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseActionsType = "drop"
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseActionsTypeForward EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseActionsType = "forward"
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseActionsTypeWorker  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnabled bool

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnabledTrue  EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnabled = true
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnabledFalse EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnabled = false
)

// Matcher for catch-all routing rule.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMatcher struct {
	// Type of matcher. Default is 'all'.
	Type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMatchersType `json:"type,required"`
	JSON emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMatcherJSON  `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMatcherJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMatcher]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of matcher. Default is 'all'.
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMatchersType string

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMatchersTypeAll EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseMatchersType = "all"
)

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelope struct {
	Errors   []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeErrors   `json:"errors"`
	Messages []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeMessages `json:"messages"`
	Result   EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponse                   `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeSuccess `json:"success"`
	JSON    emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelope]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeErrors struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeErrors]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeMessages struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeMessages]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeSuccess bool

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeSuccessTrue EmailRoutingRuleCatchAllEmailRoutingRoutingRulesGetCatchAllRuleResponseEnvelopeSuccess = true
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

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelope struct {
	Errors   []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeErrors   `json:"errors"`
	Messages []EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeMessages `json:"messages"`
	Result   EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponse                   `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeSuccess `json:"success"`
	JSON    emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelope]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeErrors struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeErrors]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeMessages struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeMessages]
type emailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeSuccess bool

const (
	EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeSuccessTrue EmailRoutingRuleCatchAllEmailRoutingRoutingRulesUpdateCatchAllRuleResponseEnvelopeSuccess = true
)
