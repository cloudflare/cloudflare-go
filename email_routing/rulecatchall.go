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

// Action for the catch-all routing rule.
type Action struct {
	// Type of action for catch-all rule.
	Type  ActionType `json:"type,required"`
	Value []string   `json:"value"`
	JSON  actionJSON `json:"-"`
}

// actionJSON contains the JSON metadata for the struct [Action]
type actionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Action) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r actionJSON) RawJSON() string {
	return r.raw
}

// Type of action for catch-all rule.
type ActionType string

const (
	ActionTypeDrop    ActionType = "drop"
	ActionTypeForward ActionType = "forward"
	ActionTypeWorker  ActionType = "worker"
)

func (r ActionType) IsKnown() bool {
	switch r {
	case ActionTypeDrop, ActionTypeForward, ActionTypeWorker:
		return true
	}
	return false
}

// Action for the catch-all routing rule.
type ActionParam struct {
	// Type of action for catch-all rule.
	Type  param.Field[ActionType] `json:"type,required"`
	Value param.Field[[]string]   `json:"value"`
}

func (r ActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmailCatchAllRule struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []Action `json:"actions"`
	// Routing rule status.
	Enabled EmailCatchAllRuleEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []Matcher `json:"matchers"`
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
type Matcher struct {
	// Type of matcher. Default is 'all'.
	Type MatcherType `json:"type,required"`
	JSON matcherJSON `json:"-"`
}

// matcherJSON contains the JSON metadata for the struct [Matcher]
type matcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Matcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matcherJSON) RawJSON() string {
	return r.raw
}

// Type of matcher. Default is 'all'.
type MatcherType string

const (
	MatcherTypeAll MatcherType = "all"
)

func (r MatcherType) IsKnown() bool {
	switch r {
	case MatcherTypeAll:
		return true
	}
	return false
}

// Matcher for catch-all routing rule.
type MatcherParam struct {
	// Type of matcher. Default is 'all'.
	Type param.Field[MatcherType] `json:"type,required"`
}

func (r MatcherParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleCatchAllUpdateParams struct {
	// List actions for the catch-all routing rule.
	Actions param.Field[[]ActionParam] `json:"actions,required"`
	// List of matchers for the catch-all routing rule.
	Matchers param.Field[[]MatcherParam] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[RuleCatchAllUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
}

func (r RuleCatchAllUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
