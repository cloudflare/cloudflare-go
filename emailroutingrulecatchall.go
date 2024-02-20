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

// Enable or disable catch-all routing rule, or change action to forward to
// specific destination address.
func (r *EmailRoutingRuleCatchAllService) Replace(ctx context.Context, zoneIdentifier string, body EmailRoutingRuleCatchAllReplaceParams, opts ...option.RequestOption) (res *EmailRoutingRuleCatchAllReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRuleCatchAllReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

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

type EmailRoutingRuleCatchAllReplaceResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []EmailRoutingRuleCatchAllReplaceResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleCatchAllReplaceResponseEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []EmailRoutingRuleCatchAllReplaceResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                      `json:"tag"`
	JSON emailRoutingRuleCatchAllReplaceResponseJSON `json:"-"`
}

// emailRoutingRuleCatchAllReplaceResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRuleCatchAllReplaceResponse]
type emailRoutingRuleCatchAllReplaceResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Action for the catch-all routing rule.
type EmailRoutingRuleCatchAllReplaceResponseAction struct {
	// Type of action for catch-all rule.
	Type  EmailRoutingRuleCatchAllReplaceResponseActionsType `json:"type,required"`
	Value []string                                           `json:"value"`
	JSON  emailRoutingRuleCatchAllReplaceResponseActionJSON  `json:"-"`
}

// emailRoutingRuleCatchAllReplaceResponseActionJSON contains the JSON metadata for
// the struct [EmailRoutingRuleCatchAllReplaceResponseAction]
type emailRoutingRuleCatchAllReplaceResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllReplaceResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of action for catch-all rule.
type EmailRoutingRuleCatchAllReplaceResponseActionsType string

const (
	EmailRoutingRuleCatchAllReplaceResponseActionsTypeDrop    EmailRoutingRuleCatchAllReplaceResponseActionsType = "drop"
	EmailRoutingRuleCatchAllReplaceResponseActionsTypeForward EmailRoutingRuleCatchAllReplaceResponseActionsType = "forward"
	EmailRoutingRuleCatchAllReplaceResponseActionsTypeWorker  EmailRoutingRuleCatchAllReplaceResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleCatchAllReplaceResponseEnabled bool

const (
	EmailRoutingRuleCatchAllReplaceResponseEnabledTrue  EmailRoutingRuleCatchAllReplaceResponseEnabled = true
	EmailRoutingRuleCatchAllReplaceResponseEnabledFalse EmailRoutingRuleCatchAllReplaceResponseEnabled = false
)

// Matcher for catch-all routing rule.
type EmailRoutingRuleCatchAllReplaceResponseMatcher struct {
	// Type of matcher. Default is 'all'.
	Type EmailRoutingRuleCatchAllReplaceResponseMatchersType `json:"type,required"`
	JSON emailRoutingRuleCatchAllReplaceResponseMatcherJSON  `json:"-"`
}

// emailRoutingRuleCatchAllReplaceResponseMatcherJSON contains the JSON metadata
// for the struct [EmailRoutingRuleCatchAllReplaceResponseMatcher]
type emailRoutingRuleCatchAllReplaceResponseMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllReplaceResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of matcher. Default is 'all'.
type EmailRoutingRuleCatchAllReplaceResponseMatchersType string

const (
	EmailRoutingRuleCatchAllReplaceResponseMatchersTypeAll EmailRoutingRuleCatchAllReplaceResponseMatchersType = "all"
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

type EmailRoutingRuleCatchAllReplaceParams struct {
	// List actions for the catch-all routing rule.
	Actions param.Field[[]EmailRoutingRuleCatchAllReplaceParamsAction] `json:"actions,required"`
	// List of matchers for the catch-all routing rule.
	Matchers param.Field[[]EmailRoutingRuleCatchAllReplaceParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRoutingRuleCatchAllReplaceParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
}

func (r EmailRoutingRuleCatchAllReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action for the catch-all routing rule.
type EmailRoutingRuleCatchAllReplaceParamsAction struct {
	// Type of action for catch-all rule.
	Type  param.Field[EmailRoutingRuleCatchAllReplaceParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                         `json:"value"`
}

func (r EmailRoutingRuleCatchAllReplaceParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of action for catch-all rule.
type EmailRoutingRuleCatchAllReplaceParamsActionsType string

const (
	EmailRoutingRuleCatchAllReplaceParamsActionsTypeDrop    EmailRoutingRuleCatchAllReplaceParamsActionsType = "drop"
	EmailRoutingRuleCatchAllReplaceParamsActionsTypeForward EmailRoutingRuleCatchAllReplaceParamsActionsType = "forward"
	EmailRoutingRuleCatchAllReplaceParamsActionsTypeWorker  EmailRoutingRuleCatchAllReplaceParamsActionsType = "worker"
)

// Matcher for catch-all routing rule.
type EmailRoutingRuleCatchAllReplaceParamsMatcher struct {
	// Type of matcher. Default is 'all'.
	Type param.Field[EmailRoutingRuleCatchAllReplaceParamsMatchersType] `json:"type,required"`
}

func (r EmailRoutingRuleCatchAllReplaceParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of matcher. Default is 'all'.
type EmailRoutingRuleCatchAllReplaceParamsMatchersType string

const (
	EmailRoutingRuleCatchAllReplaceParamsMatchersTypeAll EmailRoutingRuleCatchAllReplaceParamsMatchersType = "all"
)

// Routing rule status.
type EmailRoutingRuleCatchAllReplaceParamsEnabled bool

const (
	EmailRoutingRuleCatchAllReplaceParamsEnabledTrue  EmailRoutingRuleCatchAllReplaceParamsEnabled = true
	EmailRoutingRuleCatchAllReplaceParamsEnabledFalse EmailRoutingRuleCatchAllReplaceParamsEnabled = false
)

type EmailRoutingRuleCatchAllReplaceResponseEnvelope struct {
	Errors   []EmailRoutingRuleCatchAllReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRuleCatchAllReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRuleCatchAllReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRuleCatchAllReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRuleCatchAllReplaceResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRuleCatchAllReplaceResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingRuleCatchAllReplaceResponseEnvelope]
type emailRoutingRuleCatchAllReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllReplaceResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    emailRoutingRuleCatchAllReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRuleCatchAllReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [EmailRoutingRuleCatchAllReplaceResponseEnvelopeErrors]
type emailRoutingRuleCatchAllReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleCatchAllReplaceResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    emailRoutingRuleCatchAllReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRuleCatchAllReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [EmailRoutingRuleCatchAllReplaceResponseEnvelopeMessages]
type emailRoutingRuleCatchAllReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleCatchAllReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleCatchAllReplaceResponseEnvelopeSuccess bool

const (
	EmailRoutingRuleCatchAllReplaceResponseEnvelopeSuccessTrue EmailRoutingRuleCatchAllReplaceResponseEnvelopeSuccess = true
)
