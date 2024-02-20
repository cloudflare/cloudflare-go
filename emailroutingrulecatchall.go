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

// Enable or disable catch-all routing rule, or change action to forward to
// specific destination address.
func (r *EmailRoutingRuleCatchAllService) Update(ctx context.Context, zoneIdentifier string, body EmailRoutingRuleCatchAllUpdateParams, opts ...option.RequestOption) (res *EmailRoutingRuleCatchAllUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRuleCatchAllUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information on the default catch-all routing rule.
func (r *EmailRoutingRuleCatchAllService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRuleCatchAllGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRuleCatchAllGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingRuleCatchAllUpdateResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []EmailRoutingRuleCatchAllUpdateResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleCatchAllUpdateResponseEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []EmailRoutingRuleCatchAllUpdateResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                     `json:"tag"`
	JSON emailRoutingRuleCatchAllUpdateResponseJSON `json:"-"`
}

// emailRoutingRuleCatchAllUpdateResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRuleCatchAllUpdateResponse]
type emailRoutingRuleCatchAllUpdateResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Action for the catch-all routing rule.
type EmailRoutingRuleCatchAllUpdateResponseAction struct {
	// Type of action for catch-all rule.
	Type  EmailRoutingRuleCatchAllUpdateResponseActionsType `json:"type,required"`
	Value []string                                          `json:"value"`
	JSON  emailRoutingRuleCatchAllUpdateResponseActionJSON  `json:"-"`
}

// emailRoutingRuleCatchAllUpdateResponseActionJSON contains the JSON metadata for
// the struct [EmailRoutingRuleCatchAllUpdateResponseAction]
type emailRoutingRuleCatchAllUpdateResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllUpdateResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of action for catch-all rule.
type EmailRoutingRuleCatchAllUpdateResponseActionsType string

const (
	EmailRoutingRuleCatchAllUpdateResponseActionsTypeDrop    EmailRoutingRuleCatchAllUpdateResponseActionsType = "drop"
	EmailRoutingRuleCatchAllUpdateResponseActionsTypeForward EmailRoutingRuleCatchAllUpdateResponseActionsType = "forward"
	EmailRoutingRuleCatchAllUpdateResponseActionsTypeWorker  EmailRoutingRuleCatchAllUpdateResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleCatchAllUpdateResponseEnabled bool

const (
	EmailRoutingRuleCatchAllUpdateResponseEnabledTrue  EmailRoutingRuleCatchAllUpdateResponseEnabled = true
	EmailRoutingRuleCatchAllUpdateResponseEnabledFalse EmailRoutingRuleCatchAllUpdateResponseEnabled = false
)

// Matcher for catch-all routing rule.
type EmailRoutingRuleCatchAllUpdateResponseMatcher struct {
	// Type of matcher. Default is 'all'.
	Type EmailRoutingRuleCatchAllUpdateResponseMatchersType `json:"type,required"`
	JSON emailRoutingRuleCatchAllUpdateResponseMatcherJSON  `json:"-"`
}

// emailRoutingRuleCatchAllUpdateResponseMatcherJSON contains the JSON metadata for
// the struct [EmailRoutingRuleCatchAllUpdateResponseMatcher]
type emailRoutingRuleCatchAllUpdateResponseMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllUpdateResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of matcher. Default is 'all'.
type EmailRoutingRuleCatchAllUpdateResponseMatchersType string

const (
	EmailRoutingRuleCatchAllUpdateResponseMatchersTypeAll EmailRoutingRuleCatchAllUpdateResponseMatchersType = "all"
)

type EmailRoutingRuleCatchAllGetResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []EmailRoutingRuleCatchAllGetResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleCatchAllGetResponseEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []EmailRoutingRuleCatchAllGetResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                  `json:"tag"`
	JSON emailRoutingRuleCatchAllGetResponseJSON `json:"-"`
}

// emailRoutingRuleCatchAllGetResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRuleCatchAllGetResponse]
type emailRoutingRuleCatchAllGetResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Action for the catch-all routing rule.
type EmailRoutingRuleCatchAllGetResponseAction struct {
	// Type of action for catch-all rule.
	Type  EmailRoutingRuleCatchAllGetResponseActionsType `json:"type,required"`
	Value []string                                       `json:"value"`
	JSON  emailRoutingRuleCatchAllGetResponseActionJSON  `json:"-"`
}

// emailRoutingRuleCatchAllGetResponseActionJSON contains the JSON metadata for the
// struct [EmailRoutingRuleCatchAllGetResponseAction]
type emailRoutingRuleCatchAllGetResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllGetResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of action for catch-all rule.
type EmailRoutingRuleCatchAllGetResponseActionsType string

const (
	EmailRoutingRuleCatchAllGetResponseActionsTypeDrop    EmailRoutingRuleCatchAllGetResponseActionsType = "drop"
	EmailRoutingRuleCatchAllGetResponseActionsTypeForward EmailRoutingRuleCatchAllGetResponseActionsType = "forward"
	EmailRoutingRuleCatchAllGetResponseActionsTypeWorker  EmailRoutingRuleCatchAllGetResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleCatchAllGetResponseEnabled bool

const (
	EmailRoutingRuleCatchAllGetResponseEnabledTrue  EmailRoutingRuleCatchAllGetResponseEnabled = true
	EmailRoutingRuleCatchAllGetResponseEnabledFalse EmailRoutingRuleCatchAllGetResponseEnabled = false
)

// Matcher for catch-all routing rule.
type EmailRoutingRuleCatchAllGetResponseMatcher struct {
	// Type of matcher. Default is 'all'.
	Type EmailRoutingRuleCatchAllGetResponseMatchersType `json:"type,required"`
	JSON emailRoutingRuleCatchAllGetResponseMatcherJSON  `json:"-"`
}

// emailRoutingRuleCatchAllGetResponseMatcherJSON contains the JSON metadata for
// the struct [EmailRoutingRuleCatchAllGetResponseMatcher]
type emailRoutingRuleCatchAllGetResponseMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllGetResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of matcher. Default is 'all'.
type EmailRoutingRuleCatchAllGetResponseMatchersType string

const (
	EmailRoutingRuleCatchAllGetResponseMatchersTypeAll EmailRoutingRuleCatchAllGetResponseMatchersType = "all"
)

type EmailRoutingRuleCatchAllUpdateParams struct {
	// List actions for the catch-all routing rule.
	Actions param.Field[[]EmailRoutingRuleCatchAllUpdateParamsAction] `json:"actions,required"`
	// List of matchers for the catch-all routing rule.
	Matchers param.Field[[]EmailRoutingRuleCatchAllUpdateParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRoutingRuleCatchAllUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
}

func (r EmailRoutingRuleCatchAllUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action for the catch-all routing rule.
type EmailRoutingRuleCatchAllUpdateParamsAction struct {
	// Type of action for catch-all rule.
	Type  param.Field[EmailRoutingRuleCatchAllUpdateParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                        `json:"value"`
}

func (r EmailRoutingRuleCatchAllUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of action for catch-all rule.
type EmailRoutingRuleCatchAllUpdateParamsActionsType string

const (
	EmailRoutingRuleCatchAllUpdateParamsActionsTypeDrop    EmailRoutingRuleCatchAllUpdateParamsActionsType = "drop"
	EmailRoutingRuleCatchAllUpdateParamsActionsTypeForward EmailRoutingRuleCatchAllUpdateParamsActionsType = "forward"
	EmailRoutingRuleCatchAllUpdateParamsActionsTypeWorker  EmailRoutingRuleCatchAllUpdateParamsActionsType = "worker"
)

// Matcher for catch-all routing rule.
type EmailRoutingRuleCatchAllUpdateParamsMatcher struct {
	// Type of matcher. Default is 'all'.
	Type param.Field[EmailRoutingRuleCatchAllUpdateParamsMatchersType] `json:"type,required"`
}

func (r EmailRoutingRuleCatchAllUpdateParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of matcher. Default is 'all'.
type EmailRoutingRuleCatchAllUpdateParamsMatchersType string

const (
	EmailRoutingRuleCatchAllUpdateParamsMatchersTypeAll EmailRoutingRuleCatchAllUpdateParamsMatchersType = "all"
)

// Routing rule status.
type EmailRoutingRuleCatchAllUpdateParamsEnabled bool

const (
	EmailRoutingRuleCatchAllUpdateParamsEnabledTrue  EmailRoutingRuleCatchAllUpdateParamsEnabled = true
	EmailRoutingRuleCatchAllUpdateParamsEnabledFalse EmailRoutingRuleCatchAllUpdateParamsEnabled = false
)

type EmailRoutingRuleCatchAllUpdateResponseEnvelope struct {
	Errors   []EmailRoutingRuleCatchAllUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRuleCatchAllUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRuleCatchAllUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRuleCatchAllUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRuleCatchAllUpdateResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRuleCatchAllUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingRuleCatchAllUpdateResponseEnvelope]
type emailRoutingRuleCatchAllUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllUpdateResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    emailRoutingRuleCatchAllUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRuleCatchAllUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [EmailRoutingRuleCatchAllUpdateResponseEnvelopeErrors]
type emailRoutingRuleCatchAllUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllUpdateResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    emailRoutingRuleCatchAllUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRuleCatchAllUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [EmailRoutingRuleCatchAllUpdateResponseEnvelopeMessages]
type emailRoutingRuleCatchAllUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleCatchAllUpdateResponseEnvelopeSuccess bool

const (
	EmailRoutingRuleCatchAllUpdateResponseEnvelopeSuccessTrue EmailRoutingRuleCatchAllUpdateResponseEnvelopeSuccess = true
)

type EmailRoutingRuleCatchAllGetResponseEnvelope struct {
	Errors   []EmailRoutingRuleCatchAllGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRuleCatchAllGetResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRuleCatchAllGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRuleCatchAllGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRuleCatchAllGetResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRuleCatchAllGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingRuleCatchAllGetResponseEnvelope]
type emailRoutingRuleCatchAllGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    emailRoutingRuleCatchAllGetResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRuleCatchAllGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [EmailRoutingRuleCatchAllGetResponseEnvelopeErrors]
type emailRoutingRuleCatchAllGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    emailRoutingRuleCatchAllGetResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRuleCatchAllGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [EmailRoutingRuleCatchAllGetResponseEnvelopeMessages]
type emailRoutingRuleCatchAllGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleCatchAllGetResponseEnvelopeSuccess bool

const (
	EmailRoutingRuleCatchAllGetResponseEnvelopeSuccessTrue EmailRoutingRuleCatchAllGetResponseEnvelopeSuccess = true
)
