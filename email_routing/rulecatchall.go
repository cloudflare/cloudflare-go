// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RuleCatchAllService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRuleCatchAllService] method
// instead.
type RuleCatchAllService struct {
	Options []option.RequestOption
}

// NewRuleCatchAllService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRuleCatchAllService(opts ...option.RequestOption) (r *RuleCatchAllService) {
	r = &RuleCatchAllService{}
	r.Options = opts
	return
}

// Enable or disable catch-all routing rule, or change action to forward to
// specific destination address.
func (r *RuleCatchAllService) Update(ctx context.Context, zoneIdentifier string, body RuleCatchAllUpdateParams, opts ...option.RequestOption) (res *EmailCatchAllRule, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleCatchAllUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information on the default catch-all routing rule.
func (r *RuleCatchAllService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailCatchAllRule, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleCatchAllGetResponseEnvelope
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

type RuleCatchAllUpdateParams struct {
	// List actions for the catch-all routing rule.
	Actions param.Field[[]RuleCatchAllUpdateParamsAction] `json:"actions,required"`
	// List of matchers for the catch-all routing rule.
	Matchers param.Field[[]RuleCatchAllUpdateParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[RuleCatchAllUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
}

func (r RuleCatchAllUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action for the catch-all routing rule.
type RuleCatchAllUpdateParamsAction struct {
	// Type of action for catch-all rule.
	Type  param.Field[RuleCatchAllUpdateParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                            `json:"value"`
}

func (r RuleCatchAllUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of action for catch-all rule.
type RuleCatchAllUpdateParamsActionsType string

const (
	RuleCatchAllUpdateParamsActionsTypeDrop    RuleCatchAllUpdateParamsActionsType = "drop"
	RuleCatchAllUpdateParamsActionsTypeForward RuleCatchAllUpdateParamsActionsType = "forward"
	RuleCatchAllUpdateParamsActionsTypeWorker  RuleCatchAllUpdateParamsActionsType = "worker"
)

func (r RuleCatchAllUpdateParamsActionsType) IsKnown() bool {
	switch r {
	case RuleCatchAllUpdateParamsActionsTypeDrop, RuleCatchAllUpdateParamsActionsTypeForward, RuleCatchAllUpdateParamsActionsTypeWorker:
		return true
	}
	return false
}

// Matcher for catch-all routing rule.
type RuleCatchAllUpdateParamsMatcher struct {
	// Type of matcher. Default is 'all'.
	Type param.Field[RuleCatchAllUpdateParamsMatchersType] `json:"type,required"`
}

func (r RuleCatchAllUpdateParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of matcher. Default is 'all'.
type RuleCatchAllUpdateParamsMatchersType string

const (
	RuleCatchAllUpdateParamsMatchersTypeAll RuleCatchAllUpdateParamsMatchersType = "all"
)

func (r RuleCatchAllUpdateParamsMatchersType) IsKnown() bool {
	switch r {
	case RuleCatchAllUpdateParamsMatchersTypeAll:
		return true
	}
	return false
}

// Routing rule status.
type RuleCatchAllUpdateParamsEnabled bool

const (
	RuleCatchAllUpdateParamsEnabledTrue  RuleCatchAllUpdateParamsEnabled = true
	RuleCatchAllUpdateParamsEnabledFalse RuleCatchAllUpdateParamsEnabled = false
)

func (r RuleCatchAllUpdateParamsEnabled) IsKnown() bool {
	switch r {
	case RuleCatchAllUpdateParamsEnabledTrue, RuleCatchAllUpdateParamsEnabledFalse:
		return true
	}
	return false
}

type RuleCatchAllUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   EmailCatchAllRule                                         `json:"result,required"`
	// Whether the API call was successful
	Success RuleCatchAllUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleCatchAllUpdateResponseEnvelopeJSON    `json:"-"`
}

// ruleCatchAllUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleCatchAllUpdateResponseEnvelope]
type ruleCatchAllUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCatchAllUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleCatchAllUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleCatchAllUpdateResponseEnvelopeSuccess bool

const (
	RuleCatchAllUpdateResponseEnvelopeSuccessTrue RuleCatchAllUpdateResponseEnvelopeSuccess = true
)

func (r RuleCatchAllUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleCatchAllUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleCatchAllGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   EmailCatchAllRule                                         `json:"result,required"`
	// Whether the API call was successful
	Success RuleCatchAllGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleCatchAllGetResponseEnvelopeJSON    `json:"-"`
}

// ruleCatchAllGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleCatchAllGetResponseEnvelope]
type ruleCatchAllGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCatchAllGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleCatchAllGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleCatchAllGetResponseEnvelopeSuccess bool

const (
	RuleCatchAllGetResponseEnvelopeSuccessTrue RuleCatchAllGetResponseEnvelopeSuccess = true
)

func (r RuleCatchAllGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleCatchAllGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
