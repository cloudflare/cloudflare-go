// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *EmailRoutingRoutingRuleCatchAllService) Update(ctx context.Context, zoneIdentifier string, body EmailRoutingRoutingRuleCatchAllUpdateParams, opts ...option.RequestOption) (res *EmailRoutingRoutingRuleCatchAllUpdateResponse, err error) {
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
func (r *EmailRoutingRoutingRuleCatchAllService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRoutingRuleCatchAllGetResponse, err error) {
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

type EmailRoutingRoutingRuleCatchAllUpdateResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []EmailRoutingRoutingRuleCatchAllUpdateResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRoutingRuleCatchAllUpdateResponseEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []EmailRoutingRoutingRuleCatchAllUpdateResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                            `json:"tag"`
	JSON emailRoutingRoutingRuleCatchAllUpdateResponseJSON `json:"-"`
}

// emailRoutingRoutingRuleCatchAllUpdateResponseJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingRuleCatchAllUpdateResponse]
type emailRoutingRoutingRuleCatchAllUpdateResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleCatchAllUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleCatchAllUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Action for the catch-all routing rule.
type EmailRoutingRoutingRuleCatchAllUpdateResponseAction struct {
	// Type of action for catch-all rule.
	Type  EmailRoutingRoutingRuleCatchAllUpdateResponseActionsType `json:"type,required"`
	Value []string                                                 `json:"value"`
	JSON  emailRoutingRoutingRuleCatchAllUpdateResponseActionJSON  `json:"-"`
}

// emailRoutingRoutingRuleCatchAllUpdateResponseActionJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingRuleCatchAllUpdateResponseAction]
type emailRoutingRoutingRuleCatchAllUpdateResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleCatchAllUpdateResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleCatchAllUpdateResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of action for catch-all rule.
type EmailRoutingRoutingRuleCatchAllUpdateResponseActionsType string

const (
	EmailRoutingRoutingRuleCatchAllUpdateResponseActionsTypeDrop    EmailRoutingRoutingRuleCatchAllUpdateResponseActionsType = "drop"
	EmailRoutingRoutingRuleCatchAllUpdateResponseActionsTypeForward EmailRoutingRoutingRuleCatchAllUpdateResponseActionsType = "forward"
	EmailRoutingRoutingRuleCatchAllUpdateResponseActionsTypeWorker  EmailRoutingRoutingRuleCatchAllUpdateResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRoutingRuleCatchAllUpdateResponseEnabled bool

const (
	EmailRoutingRoutingRuleCatchAllUpdateResponseEnabledTrue  EmailRoutingRoutingRuleCatchAllUpdateResponseEnabled = true
	EmailRoutingRoutingRuleCatchAllUpdateResponseEnabledFalse EmailRoutingRoutingRuleCatchAllUpdateResponseEnabled = false
)

// Matcher for catch-all routing rule.
type EmailRoutingRoutingRuleCatchAllUpdateResponseMatcher struct {
	// Type of matcher. Default is 'all'.
	Type EmailRoutingRoutingRuleCatchAllUpdateResponseMatchersType `json:"type,required"`
	JSON emailRoutingRoutingRuleCatchAllUpdateResponseMatcherJSON  `json:"-"`
}

// emailRoutingRoutingRuleCatchAllUpdateResponseMatcherJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingRuleCatchAllUpdateResponseMatcher]
type emailRoutingRoutingRuleCatchAllUpdateResponseMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleCatchAllUpdateResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleCatchAllUpdateResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Type of matcher. Default is 'all'.
type EmailRoutingRoutingRuleCatchAllUpdateResponseMatchersType string

const (
	EmailRoutingRoutingRuleCatchAllUpdateResponseMatchersTypeAll EmailRoutingRoutingRuleCatchAllUpdateResponseMatchersType = "all"
)

type EmailRoutingRoutingRuleCatchAllGetResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []EmailRoutingRoutingRuleCatchAllGetResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRoutingRuleCatchAllGetResponseEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []EmailRoutingRoutingRuleCatchAllGetResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                         `json:"tag"`
	JSON emailRoutingRoutingRuleCatchAllGetResponseJSON `json:"-"`
}

// emailRoutingRoutingRuleCatchAllGetResponseJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingRuleCatchAllGetResponse]
type emailRoutingRoutingRuleCatchAllGetResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleCatchAllGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleCatchAllGetResponseJSON) RawJSON() string {
	return r.raw
}

// Action for the catch-all routing rule.
type EmailRoutingRoutingRuleCatchAllGetResponseAction struct {
	// Type of action for catch-all rule.
	Type  EmailRoutingRoutingRuleCatchAllGetResponseActionsType `json:"type,required"`
	Value []string                                              `json:"value"`
	JSON  emailRoutingRoutingRuleCatchAllGetResponseActionJSON  `json:"-"`
}

// emailRoutingRoutingRuleCatchAllGetResponseActionJSON contains the JSON metadata
// for the struct [EmailRoutingRoutingRuleCatchAllGetResponseAction]
type emailRoutingRoutingRuleCatchAllGetResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleCatchAllGetResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleCatchAllGetResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of action for catch-all rule.
type EmailRoutingRoutingRuleCatchAllGetResponseActionsType string

const (
	EmailRoutingRoutingRuleCatchAllGetResponseActionsTypeDrop    EmailRoutingRoutingRuleCatchAllGetResponseActionsType = "drop"
	EmailRoutingRoutingRuleCatchAllGetResponseActionsTypeForward EmailRoutingRoutingRuleCatchAllGetResponseActionsType = "forward"
	EmailRoutingRoutingRuleCatchAllGetResponseActionsTypeWorker  EmailRoutingRoutingRuleCatchAllGetResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRoutingRuleCatchAllGetResponseEnabled bool

const (
	EmailRoutingRoutingRuleCatchAllGetResponseEnabledTrue  EmailRoutingRoutingRuleCatchAllGetResponseEnabled = true
	EmailRoutingRoutingRuleCatchAllGetResponseEnabledFalse EmailRoutingRoutingRuleCatchAllGetResponseEnabled = false
)

// Matcher for catch-all routing rule.
type EmailRoutingRoutingRuleCatchAllGetResponseMatcher struct {
	// Type of matcher. Default is 'all'.
	Type EmailRoutingRoutingRuleCatchAllGetResponseMatchersType `json:"type,required"`
	JSON emailRoutingRoutingRuleCatchAllGetResponseMatcherJSON  `json:"-"`
}

// emailRoutingRoutingRuleCatchAllGetResponseMatcherJSON contains the JSON metadata
// for the struct [EmailRoutingRoutingRuleCatchAllGetResponseMatcher]
type emailRoutingRoutingRuleCatchAllGetResponseMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleCatchAllGetResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleCatchAllGetResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Type of matcher. Default is 'all'.
type EmailRoutingRoutingRuleCatchAllGetResponseMatchersType string

const (
	EmailRoutingRoutingRuleCatchAllGetResponseMatchersTypeAll EmailRoutingRoutingRuleCatchAllGetResponseMatchersType = "all"
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
	Result   EmailRoutingRoutingRuleCatchAllUpdateResponse                   `json:"result,required"`
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

func (r emailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r emailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r emailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeSuccessTrue EmailRoutingRoutingRuleCatchAllUpdateResponseEnvelopeSuccess = true
)

type EmailRoutingRoutingRuleCatchAllGetResponseEnvelope struct {
	Errors   []EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingRuleCatchAllGetResponse                   `json:"result,required"`
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

func (r emailRoutingRoutingRuleCatchAllGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r emailRoutingRoutingRuleCatchAllGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r emailRoutingRoutingRuleCatchAllGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeSuccessTrue EmailRoutingRoutingRuleCatchAllGetResponseEnvelopeSuccess = true
)
