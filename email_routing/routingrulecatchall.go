// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RoutingRuleCatchAllService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRoutingRuleCatchAllService]
// method instead.
type RoutingRuleCatchAllService struct {
	Options []option.RequestOption
}

// NewRoutingRuleCatchAllService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRoutingRuleCatchAllService(opts ...option.RequestOption) (r *RoutingRuleCatchAllService) {
	r = &RoutingRuleCatchAllService{}
	r.Options = opts
	return
}

// Enable or disable catch-all routing rule, or change action to forward to
// specific destination address.
func (r *RoutingRuleCatchAllService) Update(ctx context.Context, zoneIdentifier string, body RoutingRuleCatchAllUpdateParams, opts ...option.RequestOption) (res *EmailCatchAllRule, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingRuleCatchAllUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information on the default catch-all routing rule.
func (r *RoutingRuleCatchAllService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailCatchAllRule, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingRuleCatchAllGetResponseEnvelope
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

func (r emailCatchAllRuleJSON) RawJSON() string {
	return r.raw
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

func (r emailCatchAllRuleActionJSON) RawJSON() string {
	return r.raw
}

// Type of action for catch-all rule.
type EmailCatchAllRuleActionsType string

const (
	EmailCatchAllRuleActionsTypeDrop    EmailCatchAllRuleActionsType = "drop"
	EmailCatchAllRuleActionsTypeForward EmailCatchAllRuleActionsType = "forward"
	EmailCatchAllRuleActionsTypeWorker  EmailCatchAllRuleActionsType = "worker"
)

func (r EmailCatchAllRuleActionsType) IsKnown() bool {
	switch r {
	case EmailCatchAllRuleActionsTypeDrop, EmailCatchAllRuleActionsTypeForward, EmailCatchAllRuleActionsTypeWorker:
		return true
	}
	return false
}

// Routing rule status.
type EmailCatchAllRuleEnabled bool

const (
	EmailCatchAllRuleEnabledTrue  EmailCatchAllRuleEnabled = true
	EmailCatchAllRuleEnabledFalse EmailCatchAllRuleEnabled = false
)

func (r EmailCatchAllRuleEnabled) IsKnown() bool {
	switch r {
	case EmailCatchAllRuleEnabledTrue, EmailCatchAllRuleEnabledFalse:
		return true
	}
	return false
}

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

func (r emailCatchAllRuleMatcherJSON) RawJSON() string {
	return r.raw
}

// Type of matcher. Default is 'all'.
type EmailCatchAllRuleMatchersType string

const (
	EmailCatchAllRuleMatchersTypeAll EmailCatchAllRuleMatchersType = "all"
)

func (r EmailCatchAllRuleMatchersType) IsKnown() bool {
	switch r {
	case EmailCatchAllRuleMatchersTypeAll:
		return true
	}
	return false
}

type RoutingRuleCatchAllUpdateParams struct {
	// List actions for the catch-all routing rule.
	Actions param.Field[[]RoutingRuleCatchAllUpdateParamsAction] `json:"actions,required"`
	// List of matchers for the catch-all routing rule.
	Matchers param.Field[[]RoutingRuleCatchAllUpdateParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[RoutingRuleCatchAllUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
}

func (r RoutingRuleCatchAllUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action for the catch-all routing rule.
type RoutingRuleCatchAllUpdateParamsAction struct {
	// Type of action for catch-all rule.
	Type  param.Field[RoutingRuleCatchAllUpdateParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                   `json:"value"`
}

func (r RoutingRuleCatchAllUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of action for catch-all rule.
type RoutingRuleCatchAllUpdateParamsActionsType string

const (
	RoutingRuleCatchAllUpdateParamsActionsTypeDrop    RoutingRuleCatchAllUpdateParamsActionsType = "drop"
	RoutingRuleCatchAllUpdateParamsActionsTypeForward RoutingRuleCatchAllUpdateParamsActionsType = "forward"
	RoutingRuleCatchAllUpdateParamsActionsTypeWorker  RoutingRuleCatchAllUpdateParamsActionsType = "worker"
)

func (r RoutingRuleCatchAllUpdateParamsActionsType) IsKnown() bool {
	switch r {
	case RoutingRuleCatchAllUpdateParamsActionsTypeDrop, RoutingRuleCatchAllUpdateParamsActionsTypeForward, RoutingRuleCatchAllUpdateParamsActionsTypeWorker:
		return true
	}
	return false
}

// Matcher for catch-all routing rule.
type RoutingRuleCatchAllUpdateParamsMatcher struct {
	// Type of matcher. Default is 'all'.
	Type param.Field[RoutingRuleCatchAllUpdateParamsMatchersType] `json:"type,required"`
}

func (r RoutingRuleCatchAllUpdateParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of matcher. Default is 'all'.
type RoutingRuleCatchAllUpdateParamsMatchersType string

const (
	RoutingRuleCatchAllUpdateParamsMatchersTypeAll RoutingRuleCatchAllUpdateParamsMatchersType = "all"
)

func (r RoutingRuleCatchAllUpdateParamsMatchersType) IsKnown() bool {
	switch r {
	case RoutingRuleCatchAllUpdateParamsMatchersTypeAll:
		return true
	}
	return false
}

// Routing rule status.
type RoutingRuleCatchAllUpdateParamsEnabled bool

const (
	RoutingRuleCatchAllUpdateParamsEnabledTrue  RoutingRuleCatchAllUpdateParamsEnabled = true
	RoutingRuleCatchAllUpdateParamsEnabledFalse RoutingRuleCatchAllUpdateParamsEnabled = false
)

func (r RoutingRuleCatchAllUpdateParamsEnabled) IsKnown() bool {
	switch r {
	case RoutingRuleCatchAllUpdateParamsEnabledTrue, RoutingRuleCatchAllUpdateParamsEnabledFalse:
		return true
	}
	return false
}

type RoutingRuleCatchAllUpdateResponseEnvelope struct {
	Errors   []RoutingRuleCatchAllUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingRuleCatchAllUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailCatchAllRule                                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingRuleCatchAllUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingRuleCatchAllUpdateResponseEnvelopeJSON    `json:"-"`
}

// routingRuleCatchAllUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [RoutingRuleCatchAllUpdateResponseEnvelope]
type routingRuleCatchAllUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleCatchAllUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleCatchAllUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingRuleCatchAllUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    routingRuleCatchAllUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// routingRuleCatchAllUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [RoutingRuleCatchAllUpdateResponseEnvelopeErrors]
type routingRuleCatchAllUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleCatchAllUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleCatchAllUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingRuleCatchAllUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    routingRuleCatchAllUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// routingRuleCatchAllUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RoutingRuleCatchAllUpdateResponseEnvelopeMessages]
type routingRuleCatchAllUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleCatchAllUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleCatchAllUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingRuleCatchAllUpdateResponseEnvelopeSuccess bool

const (
	RoutingRuleCatchAllUpdateResponseEnvelopeSuccessTrue RoutingRuleCatchAllUpdateResponseEnvelopeSuccess = true
)

func (r RoutingRuleCatchAllUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RoutingRuleCatchAllUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RoutingRuleCatchAllGetResponseEnvelope struct {
	Errors   []RoutingRuleCatchAllGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingRuleCatchAllGetResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailCatchAllRule                                `json:"result,required"`
	// Whether the API call was successful
	Success RoutingRuleCatchAllGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingRuleCatchAllGetResponseEnvelopeJSON    `json:"-"`
}

// routingRuleCatchAllGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RoutingRuleCatchAllGetResponseEnvelope]
type routingRuleCatchAllGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleCatchAllGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleCatchAllGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingRuleCatchAllGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    routingRuleCatchAllGetResponseEnvelopeErrorsJSON `json:"-"`
}

// routingRuleCatchAllGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RoutingRuleCatchAllGetResponseEnvelopeErrors]
type routingRuleCatchAllGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleCatchAllGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleCatchAllGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingRuleCatchAllGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    routingRuleCatchAllGetResponseEnvelopeMessagesJSON `json:"-"`
}

// routingRuleCatchAllGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RoutingRuleCatchAllGetResponseEnvelopeMessages]
type routingRuleCatchAllGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleCatchAllGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleCatchAllGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingRuleCatchAllGetResponseEnvelopeSuccess bool

const (
	RoutingRuleCatchAllGetResponseEnvelopeSuccessTrue RoutingRuleCatchAllGetResponseEnvelopeSuccess = true
)

func (r RoutingRuleCatchAllGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RoutingRuleCatchAllGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
