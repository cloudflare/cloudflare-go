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

// EmailRoutingRoutingRuleCatchAllService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewEmailRoutingRoutingRuleCatchAllService] method instead.
type EmailRoutingRoutingRuleCatchAllService struct {
	Options []option.RequestOption
}

// NewEmailRoutingRoutingRuleCatchAllService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailRoutingRoutingRuleCatchAllService(opts ...option.RequestOption) (r *EmailRoutingRoutingRuleCatchAllService) {
	r = &EmailRoutingRoutingRuleCatchAllService{}
	r.Options = opts
	return
}

// Enable or disable catch-all routing rule, or change action to forward to
// specific destination address.
func (r *EmailRoutingRoutingRuleCatchAllService) Update(ctx context.Context, zoneIdentifier string, body EmailRoutingRoutingRuleCatchAllUpdateParams, opts ...option.RequestOption) (res *EmailCatchAllRule, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information on the default catch-all routing rule.
func (r *EmailRoutingRoutingRuleCatchAllService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailCatchAllRule, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingRuleCatchAllGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailCatchAllRule struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []EmailCatchAllRuleAction `json:"actions"`
	// Routing rule status.
	Enabled EmailCatchAllRuleEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []EmailCatchAllRuleMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                `json:"tag"`
	JSON emailCatchAllRuleJSON `json:"-"`
}

// emailCatchAllRuleJSON contains the JSON metadata for the struct
// [EmailCatchAllRule]
type emailCatchAllRuleJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailCatchAllRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Action for the catch-all routing rule.
type EmailCatchAllRuleAction struct {
	// Type of action for catch-all rule.
	Type  EmailCatchAllRuleActionsType `json:"type,required"`
	Value []string                     `json:"value"`
	JSON  emailCatchAllRuleActionJSON  `json:"-"`
}

// emailCatchAllRuleActionJSON contains the JSON metadata for the struct
// [EmailCatchAllRuleAction]
type emailCatchAllRuleActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailCatchAllRuleAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of action for catch-all rule.
type EmailCatchAllRuleActionsType string

const (
	EmailCatchAllRuleActionsTypeDrop    EmailCatchAllRuleActionsType = "drop"
	EmailCatchAllRuleActionsTypeForward EmailCatchAllRuleActionsType = "forward"
	EmailCatchAllRuleActionsTypeWorker  EmailCatchAllRuleActionsType = "worker"
)

// Routing rule status.
type EmailCatchAllRuleEnabled bool

const (
	EmailCatchAllRuleEnabledTrue  EmailCatchAllRuleEnabled = true
	EmailCatchAllRuleEnabledFalse EmailCatchAllRuleEnabled = false
)

// Matcher for catch-all routing rule.
type EmailCatchAllRuleMatcher struct {
	// Type of matcher. Default is 'all'.
	Type EmailCatchAllRuleMatchersType `json:"type,required"`
	JSON emailCatchAllRuleMatcherJSON  `json:"-"`
}

// emailCatchAllRuleMatcherJSON contains the JSON metadata for the struct
// [EmailCatchAllRuleMatcher]
type emailCatchAllRuleMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailCatchAllRuleMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of matcher. Default is 'all'.
type EmailCatchAllRuleMatchersType string

const (
	EmailCatchAllRuleMatchersTypeAll EmailCatchAllRuleMatchersType = "all"
)

type EmailRoutingRoutingRuleCatchAllUpdateParams struct {
	// List actions for the catch-all routing rule.
	Actions param.Field[[]EmailRoutingRoutingRuleCatchAllUpdateParamsAction] `json:"actions,required"`
	// List of matchers for the catch-all routing rule.
	Matchers param.Field[[]EmailRoutingRoutingRuleCatchAllUpdateParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRoutingRoutingRuleCatchAllUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
}

func (r EmailRoutingRoutingRuleCatchAllUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action for the catch-all routing rule.
type EmailRoutingRoutingRuleCatchAllUpdateParamsAction struct {
	// Type of action for catch-all rule.
	Type  param.Field[EmailRoutingRoutingRuleCatchAllUpdateParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                               `json:"value"`
}

func (r EmailRoutingRoutingRuleCatchAllUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of action for catch-all rule.
type EmailRoutingRoutingRuleCatchAllUpdateParamsActionsType string

const (
	EmailRoutingRoutingRuleCatchAllUpdateParamsActionsTypeDrop    EmailRoutingRoutingRuleCatchAllUpdateParamsActionsType = "drop"
	EmailRoutingRoutingRuleCatchAllUpdateParamsActionsTypeForward EmailRoutingRoutingRuleCatchAllUpdateParamsActionsType = "forward"
	EmailRoutingRoutingRuleCatchAllUpdateParamsActionsTypeWorker  EmailRoutingRoutingRuleCatchAllUpdateParamsActionsType = "worker"
)

// Matcher for catch-all routing rule.
type EmailRoutingRoutingRuleCatchAllUpdateParamsMatcher struct {
	// Type of matcher. Default is 'all'.
	Type param.Field[EmailRoutingRoutingRuleCatchAllUpdateParamsMatchersType] `json:"type,required"`
}

func (r EmailRoutingRoutingRuleCatchAllUpdateParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of matcher. Default is 'all'.
type EmailRoutingRoutingRuleCatchAllUpdateParamsMatchersType string

const (
	EmailRoutingRoutingRuleCatchAllUpdateParamsMatchersTypeAll EmailRoutingRoutingRuleCatchAllUpdateParamsMatchersType = "all"
)

// Routing rule status.
type EmailRoutingRoutingRuleCatchAllUpdateParamsEnabled bool

const (
	EmailRoutingRoutingRuleCatchAllUpdateParamsEnabledTrue  EmailRoutingRoutingRuleCatchAllUpdateParamsEnabled = true
	EmailRoutingRoutingRuleCatchAllUpdateParamsEnabledFalse EmailRoutingRoutingRuleCatchAllUpdateParamsEnabled = false
)

type EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelope struct {
	Errors   []EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailCatchAllRule                                               `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelope]
type emailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    emailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeErrors]
type emailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    emailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeMessages]
type emailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeSuccessTrue EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeSuccess = true
)

type EmailRoutingRoutingRuleCatchAllGetResponseEnvelope struct {
	Errors   []EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailCatchAllRule                                            `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingRuleCatchAllGetResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingRuleCatchAllGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingRuleCatchAllGetResponseEnvelope]
type emailRoutingRoutingRuleCatchAllGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleCatchAllGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    emailRoutingRoutingRuleCatchAllGetResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingRuleCatchAllGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeErrors]
type emailRoutingRoutingRuleCatchAllGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    emailRoutingRoutingRuleCatchAllGetResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingRuleCatchAllGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeMessages]
type emailRoutingRoutingRuleCatchAllGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeSuccessTrue EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeSuccess = true
)
