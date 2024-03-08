// File generated from our OpenAPI spec by Stainless.

package email_routing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *RoutingRuleCatchAllService) Update(ctx context.Context, zoneIdentifier string, body RoutingRuleCatchAllUpdateParams, opts ...option.RequestOption) (res *RoutingRuleCatchAllUpdateResponse, err error) {
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
func (r *RoutingRuleCatchAllService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *RoutingRuleCatchAllGetResponse, err error) {
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

type RoutingRuleCatchAllUpdateResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []RoutingRuleCatchAllUpdateResponseAction `json:"actions"`
	// Routing rule status.
	Enabled RoutingRuleCatchAllUpdateResponseEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []RoutingRuleCatchAllUpdateResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                `json:"tag"`
	JSON routingRuleCatchAllUpdateResponseJSON `json:"-"`
}

// routingRuleCatchAllUpdateResponseJSON contains the JSON metadata for the struct
// [RoutingRuleCatchAllUpdateResponse]
type routingRuleCatchAllUpdateResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleCatchAllUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleCatchAllUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Action for the catch-all routing rule.
type RoutingRuleCatchAllUpdateResponseAction struct {
	// Type of action for catch-all rule.
	Type  RoutingRuleCatchAllUpdateResponseActionsType `json:"type,required"`
	Value []string                                     `json:"value"`
	JSON  routingRuleCatchAllUpdateResponseActionJSON  `json:"-"`
}

// routingRuleCatchAllUpdateResponseActionJSON contains the JSON metadata for the
// struct [RoutingRuleCatchAllUpdateResponseAction]
type routingRuleCatchAllUpdateResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleCatchAllUpdateResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleCatchAllUpdateResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of action for catch-all rule.
type RoutingRuleCatchAllUpdateResponseActionsType string

const (
	RoutingRuleCatchAllUpdateResponseActionsTypeDrop    RoutingRuleCatchAllUpdateResponseActionsType = "drop"
	RoutingRuleCatchAllUpdateResponseActionsTypeForward RoutingRuleCatchAllUpdateResponseActionsType = "forward"
	RoutingRuleCatchAllUpdateResponseActionsTypeWorker  RoutingRuleCatchAllUpdateResponseActionsType = "worker"
)

// Routing rule status.
type RoutingRuleCatchAllUpdateResponseEnabled bool

const (
	RoutingRuleCatchAllUpdateResponseEnabledTrue  RoutingRuleCatchAllUpdateResponseEnabled = true
	RoutingRuleCatchAllUpdateResponseEnabledFalse RoutingRuleCatchAllUpdateResponseEnabled = false
)

// Matcher for catch-all routing rule.
type RoutingRuleCatchAllUpdateResponseMatcher struct {
	// Type of matcher. Default is 'all'.
	Type RoutingRuleCatchAllUpdateResponseMatchersType `json:"type,required"`
	JSON routingRuleCatchAllUpdateResponseMatcherJSON  `json:"-"`
}

// routingRuleCatchAllUpdateResponseMatcherJSON contains the JSON metadata for the
// struct [RoutingRuleCatchAllUpdateResponseMatcher]
type routingRuleCatchAllUpdateResponseMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleCatchAllUpdateResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleCatchAllUpdateResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Type of matcher. Default is 'all'.
type RoutingRuleCatchAllUpdateResponseMatchersType string

const (
	RoutingRuleCatchAllUpdateResponseMatchersTypeAll RoutingRuleCatchAllUpdateResponseMatchersType = "all"
)

type RoutingRuleCatchAllGetResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []RoutingRuleCatchAllGetResponseAction `json:"actions"`
	// Routing rule status.
	Enabled RoutingRuleCatchAllGetResponseEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []RoutingRuleCatchAllGetResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                             `json:"tag"`
	JSON routingRuleCatchAllGetResponseJSON `json:"-"`
}

// routingRuleCatchAllGetResponseJSON contains the JSON metadata for the struct
// [RoutingRuleCatchAllGetResponse]
type routingRuleCatchAllGetResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleCatchAllGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleCatchAllGetResponseJSON) RawJSON() string {
	return r.raw
}

// Action for the catch-all routing rule.
type RoutingRuleCatchAllGetResponseAction struct {
	// Type of action for catch-all rule.
	Type  RoutingRuleCatchAllGetResponseActionsType `json:"type,required"`
	Value []string                                  `json:"value"`
	JSON  routingRuleCatchAllGetResponseActionJSON  `json:"-"`
}

// routingRuleCatchAllGetResponseActionJSON contains the JSON metadata for the
// struct [RoutingRuleCatchAllGetResponseAction]
type routingRuleCatchAllGetResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleCatchAllGetResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleCatchAllGetResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of action for catch-all rule.
type RoutingRuleCatchAllGetResponseActionsType string

const (
	RoutingRuleCatchAllGetResponseActionsTypeDrop    RoutingRuleCatchAllGetResponseActionsType = "drop"
	RoutingRuleCatchAllGetResponseActionsTypeForward RoutingRuleCatchAllGetResponseActionsType = "forward"
	RoutingRuleCatchAllGetResponseActionsTypeWorker  RoutingRuleCatchAllGetResponseActionsType = "worker"
)

// Routing rule status.
type RoutingRuleCatchAllGetResponseEnabled bool

const (
	RoutingRuleCatchAllGetResponseEnabledTrue  RoutingRuleCatchAllGetResponseEnabled = true
	RoutingRuleCatchAllGetResponseEnabledFalse RoutingRuleCatchAllGetResponseEnabled = false
)

// Matcher for catch-all routing rule.
type RoutingRuleCatchAllGetResponseMatcher struct {
	// Type of matcher. Default is 'all'.
	Type RoutingRuleCatchAllGetResponseMatchersType `json:"type,required"`
	JSON routingRuleCatchAllGetResponseMatcherJSON  `json:"-"`
}

// routingRuleCatchAllGetResponseMatcherJSON contains the JSON metadata for the
// struct [RoutingRuleCatchAllGetResponseMatcher]
type routingRuleCatchAllGetResponseMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleCatchAllGetResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleCatchAllGetResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Type of matcher. Default is 'all'.
type RoutingRuleCatchAllGetResponseMatchersType string

const (
	RoutingRuleCatchAllGetResponseMatchersTypeAll RoutingRuleCatchAllGetResponseMatchersType = "all"
)

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

// Routing rule status.
type RoutingRuleCatchAllUpdateParamsEnabled bool

const (
	RoutingRuleCatchAllUpdateParamsEnabledTrue  RoutingRuleCatchAllUpdateParamsEnabled = true
	RoutingRuleCatchAllUpdateParamsEnabledFalse RoutingRuleCatchAllUpdateParamsEnabled = false
)

type RoutingRuleCatchAllUpdateResponseEnvelope struct {
	Errors   []RoutingRuleCatchAllUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingRuleCatchAllUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingRuleCatchAllUpdateResponse                   `json:"result,required"`
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

type RoutingRuleCatchAllGetResponseEnvelope struct {
	Errors   []RoutingRuleCatchAllGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingRuleCatchAllGetResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingRuleCatchAllGetResponse                   `json:"result,required"`
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
