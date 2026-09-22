// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ObservabilityTracingRuleService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityTracingRuleService] method instead.
type ObservabilityTracingRuleService struct {
	Options []option.RequestOption
}

// NewObservabilityTracingRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservabilityTracingRuleService(opts ...option.RequestOption) (r *ObservabilityTracingRuleService) {
	r = &ObservabilityTracingRuleService{}
	r.Options = opts
	return
}

// Replace all sampling overrides in a zone's managed Cloudflare Traces ruleset.
// Rules are evaluated in the supplied order.
func (r *ObservabilityTracingRuleService) Update(ctx context.Context, params ObservabilityTracingRuleUpdateParams, opts ...option.RequestOption) (res *ObservabilityTracingRuleUpdateResponse, err error) {
	var env ObservabilityTracingRuleUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/observability/tracing/rules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Delete every sampling override from a zone's managed Cloudflare Traces ruleset.
func (r *ObservabilityTracingRuleService) Delete(ctx context.Context, body ObservabilityTracingRuleDeleteParams, opts ...option.RequestOption) (res *ObservabilityTracingRuleDeleteResponse, err error) {
	var env ObservabilityTracingRuleDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/observability/tracing/rules", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve the ordered sampling overrides for a zone's managed Cloudflare Traces
// ruleset.
func (r *ObservabilityTracingRuleService) Get(ctx context.Context, query ObservabilityTracingRuleGetParams, opts ...option.RequestOption) (res *ObservabilityTracingRuleGetResponse, err error) {
	var env ObservabilityTracingRuleGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/observability/tracing/rules", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ObservabilityTracingRuleUpdateResponse struct {
	// Trace rules in evaluation order.
	Rules []ObservabilityTracingRuleUpdateResponseRule `json:"rules" api:"required"`
	JSON  observabilityTracingRuleUpdateResponseJSON   `json:"-"`
}

// observabilityTracingRuleUpdateResponseJSON contains the JSON metadata for the
// struct [ObservabilityTracingRuleUpdateResponse]
type observabilityTracingRuleUpdateResponseJSON struct {
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleUpdateResponseRule struct {
	Action           ObservabilityTracingRuleUpdateResponseRulesAction           `json:"action" api:"required"`
	ActionParameters ObservabilityTracingRuleUpdateResponseRulesActionParameters `json:"action_parameters" api:"required"`
	Description      string                                                      `json:"description" api:"required"`
	Enabled          bool                                                        `json:"enabled" api:"required"`
	// A Rules language expression that selects requests.
	Expression string                                         `json:"expression" api:"required"`
	JSON       observabilityTracingRuleUpdateResponseRuleJSON `json:"-"`
}

// observabilityTracingRuleUpdateResponseRuleJSON contains the JSON metadata for
// the struct [ObservabilityTracingRuleUpdateResponseRule]
type observabilityTracingRuleUpdateResponseRuleJSON struct {
	Action           apijson.Field
	ActionParameters apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ObservabilityTracingRuleUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleUpdateResponseRulesAction string

const (
	ObservabilityTracingRuleUpdateResponseRulesActionSetTraceSettings ObservabilityTracingRuleUpdateResponseRulesAction = "set_trace_settings"
)

func (r ObservabilityTracingRuleUpdateResponseRulesAction) IsKnown() bool {
	switch r {
	case ObservabilityTracingRuleUpdateResponseRulesActionSetTraceSettings:
		return true
	}
	return false
}

type ObservabilityTracingRuleUpdateResponseRulesActionParameters struct {
	// The ratio of requests sampled for tracing, from 0 to 1.
	SamplingRatio float64                                                         `json:"sampling_ratio" api:"required"`
	JSON          observabilityTracingRuleUpdateResponseRulesActionParametersJSON `json:"-"`
}

// observabilityTracingRuleUpdateResponseRulesActionParametersJSON contains the
// JSON metadata for the struct
// [ObservabilityTracingRuleUpdateResponseRulesActionParameters]
type observabilityTracingRuleUpdateResponseRulesActionParametersJSON struct {
	SamplingRatio apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityTracingRuleUpdateResponseRulesActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleUpdateResponseRulesActionParametersJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleDeleteResponse struct {
	// Trace rules in evaluation order.
	Rules []ObservabilityTracingRuleDeleteResponseRule `json:"rules" api:"required"`
	JSON  observabilityTracingRuleDeleteResponseJSON   `json:"-"`
}

// observabilityTracingRuleDeleteResponseJSON contains the JSON metadata for the
// struct [ObservabilityTracingRuleDeleteResponse]
type observabilityTracingRuleDeleteResponseJSON struct {
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleDeleteResponseRule struct {
	Action           ObservabilityTracingRuleDeleteResponseRulesAction           `json:"action" api:"required"`
	ActionParameters ObservabilityTracingRuleDeleteResponseRulesActionParameters `json:"action_parameters" api:"required"`
	Description      string                                                      `json:"description" api:"required"`
	Enabled          bool                                                        `json:"enabled" api:"required"`
	// A Rules language expression that selects requests.
	Expression string                                         `json:"expression" api:"required"`
	JSON       observabilityTracingRuleDeleteResponseRuleJSON `json:"-"`
}

// observabilityTracingRuleDeleteResponseRuleJSON contains the JSON metadata for
// the struct [ObservabilityTracingRuleDeleteResponseRule]
type observabilityTracingRuleDeleteResponseRuleJSON struct {
	Action           apijson.Field
	ActionParameters apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ObservabilityTracingRuleDeleteResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleDeleteResponseRuleJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleDeleteResponseRulesAction string

const (
	ObservabilityTracingRuleDeleteResponseRulesActionSetTraceSettings ObservabilityTracingRuleDeleteResponseRulesAction = "set_trace_settings"
)

func (r ObservabilityTracingRuleDeleteResponseRulesAction) IsKnown() bool {
	switch r {
	case ObservabilityTracingRuleDeleteResponseRulesActionSetTraceSettings:
		return true
	}
	return false
}

type ObservabilityTracingRuleDeleteResponseRulesActionParameters struct {
	// The ratio of requests sampled for tracing, from 0 to 1.
	SamplingRatio float64                                                         `json:"sampling_ratio" api:"required"`
	JSON          observabilityTracingRuleDeleteResponseRulesActionParametersJSON `json:"-"`
}

// observabilityTracingRuleDeleteResponseRulesActionParametersJSON contains the
// JSON metadata for the struct
// [ObservabilityTracingRuleDeleteResponseRulesActionParameters]
type observabilityTracingRuleDeleteResponseRulesActionParametersJSON struct {
	SamplingRatio apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityTracingRuleDeleteResponseRulesActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleDeleteResponseRulesActionParametersJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleGetResponse struct {
	// Trace rules in evaluation order.
	Rules []ObservabilityTracingRuleGetResponseRule `json:"rules" api:"required"`
	JSON  observabilityTracingRuleGetResponseJSON   `json:"-"`
}

// observabilityTracingRuleGetResponseJSON contains the JSON metadata for the
// struct [ObservabilityTracingRuleGetResponse]
type observabilityTracingRuleGetResponseJSON struct {
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleGetResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleGetResponseRule struct {
	Action           ObservabilityTracingRuleGetResponseRulesAction           `json:"action" api:"required"`
	ActionParameters ObservabilityTracingRuleGetResponseRulesActionParameters `json:"action_parameters" api:"required"`
	Description      string                                                   `json:"description" api:"required"`
	Enabled          bool                                                     `json:"enabled" api:"required"`
	// A Rules language expression that selects requests.
	Expression string                                      `json:"expression" api:"required"`
	JSON       observabilityTracingRuleGetResponseRuleJSON `json:"-"`
}

// observabilityTracingRuleGetResponseRuleJSON contains the JSON metadata for the
// struct [ObservabilityTracingRuleGetResponseRule]
type observabilityTracingRuleGetResponseRuleJSON struct {
	Action           apijson.Field
	ActionParameters apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ObservabilityTracingRuleGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleGetResponseRulesAction string

const (
	ObservabilityTracingRuleGetResponseRulesActionSetTraceSettings ObservabilityTracingRuleGetResponseRulesAction = "set_trace_settings"
)

func (r ObservabilityTracingRuleGetResponseRulesAction) IsKnown() bool {
	switch r {
	case ObservabilityTracingRuleGetResponseRulesActionSetTraceSettings:
		return true
	}
	return false
}

type ObservabilityTracingRuleGetResponseRulesActionParameters struct {
	// The ratio of requests sampled for tracing, from 0 to 1.
	SamplingRatio float64                                                      `json:"sampling_ratio" api:"required"`
	JSON          observabilityTracingRuleGetResponseRulesActionParametersJSON `json:"-"`
}

// observabilityTracingRuleGetResponseRulesActionParametersJSON contains the JSON
// metadata for the struct
// [ObservabilityTracingRuleGetResponseRulesActionParameters]
type observabilityTracingRuleGetResponseRulesActionParametersJSON struct {
	SamplingRatio apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityTracingRuleGetResponseRulesActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleGetResponseRulesActionParametersJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleUpdateParams struct {
	// Specify the zone ID.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Trace rules in evaluation order.
	Rules param.Field[[]ObservabilityTracingRuleUpdateParamsRule] `json:"rules" api:"required"`
}

func (r ObservabilityTracingRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTracingRuleUpdateParamsRule struct {
	Action           param.Field[ObservabilityTracingRuleUpdateParamsRulesAction]           `json:"action" api:"required"`
	ActionParameters param.Field[ObservabilityTracingRuleUpdateParamsRulesActionParameters] `json:"action_parameters" api:"required"`
	Description      param.Field[string]                                                    `json:"description" api:"required"`
	Enabled          param.Field[bool]                                                      `json:"enabled" api:"required"`
	// A Rules language expression that selects requests.
	Expression param.Field[string] `json:"expression" api:"required"`
}

func (r ObservabilityTracingRuleUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTracingRuleUpdateParamsRulesAction string

const (
	ObservabilityTracingRuleUpdateParamsRulesActionSetTraceSettings ObservabilityTracingRuleUpdateParamsRulesAction = "set_trace_settings"
)

func (r ObservabilityTracingRuleUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case ObservabilityTracingRuleUpdateParamsRulesActionSetTraceSettings:
		return true
	}
	return false
}

type ObservabilityTracingRuleUpdateParamsRulesActionParameters struct {
	// The ratio of requests sampled for tracing, from 0 to 1.
	SamplingRatio param.Field[float64] `json:"sampling_ratio" api:"required"`
}

func (r ObservabilityTracingRuleUpdateParamsRulesActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTracingRuleUpdateResponseEnvelope struct {
	Errors   []ObservabilityTracingRuleUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityTracingRuleUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityTracingRuleUpdateResponse                   `json:"result" api:"required"`
	Success  ObservabilityTracingRuleUpdateResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityTracingRuleUpdateResponseEnvelopeJSON       `json:"-"`
}

// observabilityTracingRuleUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ObservabilityTracingRuleUpdateResponseEnvelope]
type observabilityTracingRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleUpdateResponseEnvelopeErrors struct {
	Message string                                                   `json:"message" api:"required"`
	JSON    observabilityTracingRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityTracingRuleUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ObservabilityTracingRuleUpdateResponseEnvelopeErrors]
type observabilityTracingRuleUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleUpdateResponseEnvelopeMessages struct {
	Message ObservabilityTracingRuleUpdateResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityTracingRuleUpdateResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityTracingRuleUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityTracingRuleUpdateResponseEnvelopeMessages]
type observabilityTracingRuleUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleUpdateResponseEnvelopeMessagesMessage string

const (
	ObservabilityTracingRuleUpdateResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityTracingRuleUpdateResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityTracingRuleUpdateResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityTracingRuleUpdateResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityTracingRuleUpdateResponseEnvelopeSuccess bool

const (
	ObservabilityTracingRuleUpdateResponseEnvelopeSuccessTrue ObservabilityTracingRuleUpdateResponseEnvelopeSuccess = true
)

func (r ObservabilityTracingRuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityTracingRuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityTracingRuleDeleteParams struct {
	// Specify the zone ID.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type ObservabilityTracingRuleDeleteResponseEnvelope struct {
	Errors   []ObservabilityTracingRuleDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityTracingRuleDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityTracingRuleDeleteResponse                   `json:"result" api:"required"`
	Success  ObservabilityTracingRuleDeleteResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityTracingRuleDeleteResponseEnvelopeJSON       `json:"-"`
}

// observabilityTracingRuleDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ObservabilityTracingRuleDeleteResponseEnvelope]
type observabilityTracingRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleDeleteResponseEnvelopeErrors struct {
	Message string                                                   `json:"message" api:"required"`
	JSON    observabilityTracingRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityTracingRuleDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ObservabilityTracingRuleDeleteResponseEnvelopeErrors]
type observabilityTracingRuleDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleDeleteResponseEnvelopeMessages struct {
	Message ObservabilityTracingRuleDeleteResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityTracingRuleDeleteResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityTracingRuleDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityTracingRuleDeleteResponseEnvelopeMessages]
type observabilityTracingRuleDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleDeleteResponseEnvelopeMessagesMessage string

const (
	ObservabilityTracingRuleDeleteResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityTracingRuleDeleteResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityTracingRuleDeleteResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityTracingRuleDeleteResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityTracingRuleDeleteResponseEnvelopeSuccess bool

const (
	ObservabilityTracingRuleDeleteResponseEnvelopeSuccessTrue ObservabilityTracingRuleDeleteResponseEnvelopeSuccess = true
)

func (r ObservabilityTracingRuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityTracingRuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityTracingRuleGetParams struct {
	// Specify the zone ID.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type ObservabilityTracingRuleGetResponseEnvelope struct {
	Errors   []ObservabilityTracingRuleGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityTracingRuleGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityTracingRuleGetResponse                   `json:"result" api:"required"`
	Success  ObservabilityTracingRuleGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityTracingRuleGetResponseEnvelopeJSON       `json:"-"`
}

// observabilityTracingRuleGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ObservabilityTracingRuleGetResponseEnvelope]
type observabilityTracingRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleGetResponseEnvelopeErrors struct {
	Message string                                                `json:"message" api:"required"`
	JSON    observabilityTracingRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityTracingRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ObservabilityTracingRuleGetResponseEnvelopeErrors]
type observabilityTracingRuleGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleGetResponseEnvelopeMessages struct {
	Message ObservabilityTracingRuleGetResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityTracingRuleGetResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityTracingRuleGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityTracingRuleGetResponseEnvelopeMessages]
type observabilityTracingRuleGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingRuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingRuleGetResponseEnvelopeMessagesMessage string

const (
	ObservabilityTracingRuleGetResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityTracingRuleGetResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityTracingRuleGetResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityTracingRuleGetResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityTracingRuleGetResponseEnvelopeSuccess bool

const (
	ObservabilityTracingRuleGetResponseEnvelopeSuccessTrue ObservabilityTracingRuleGetResponseEnvelopeSuccess = true
)

func (r ObservabilityTracingRuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityTracingRuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
