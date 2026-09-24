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

// ObservabilityTracingSettingService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityTracingSettingService] method instead.
type ObservabilityTracingSettingService struct {
	Options []option.RequestOption
}

// NewObservabilityTracingSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservabilityTracingSettingService(opts ...option.RequestOption) (r *ObservabilityTracingSettingService) {
	r = &ObservabilityTracingSettingService{}
	r.Options = opts
	return
}

// Update the zone-level Cloudflare Traces settings.
func (r *ObservabilityTracingSettingService) Update(ctx context.Context, params ObservabilityTracingSettingUpdateParams, opts ...option.RequestOption) (res *ObservabilityTracingSettingUpdateResponse, err error) {
	var env ObservabilityTracingSettingUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/observability/tracing/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Reset the zone-level Cloudflare Traces settings to their defaults while
// preserving the sampling rules.
func (r *ObservabilityTracingSettingService) Delete(ctx context.Context, body ObservabilityTracingSettingDeleteParams, opts ...option.RequestOption) (res *ObservabilityTracingSettingDeleteResponse, err error) {
	var env ObservabilityTracingSettingDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/observability/tracing/settings", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve the zone-level Cloudflare Traces settings.
func (r *ObservabilityTracingSettingService) Get(ctx context.Context, query ObservabilityTracingSettingGetParams, opts ...option.RequestOption) (res *ObservabilityTracingSettingGetResponse, err error) {
	var env ObservabilityTracingSettingGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/observability/tracing/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ObservabilityTracingSettingUpdateResponse struct {
	// Up to 100 OpenTelemetry destination identifiers that receive traces.
	Destinations []string `json:"destinations" api:"required"`
	// Whether Cloudflare Traces is enabled for the zone.
	Enabled bool `json:"enabled" api:"required"`
	// Whether trace context is sent externally or across a zone boundary.
	ForwardContext bool `json:"forward_context" api:"required"`
	// Whether traces are persisted in Cloudflare.
	Persist bool `json:"persist" api:"required"`
	// When inbound trace context may be continued. Authenticated propagation is not
	// supported yet.
	PropagationPolicy ObservabilityTracingSettingUpdateResponsePropagationPolicy `json:"propagation_policy" api:"required"`
	// The ratio of requests sampled for tracing, from 0 to 1.
	SamplingRatio float64                                       `json:"sampling_ratio" api:"required"`
	JSON          observabilityTracingSettingUpdateResponseJSON `json:"-"`
}

// observabilityTracingSettingUpdateResponseJSON contains the JSON metadata for the
// struct [ObservabilityTracingSettingUpdateResponse]
type observabilityTracingSettingUpdateResponseJSON struct {
	Destinations      apijson.Field
	Enabled           apijson.Field
	ForwardContext    apijson.Field
	Persist           apijson.Field
	PropagationPolicy apijson.Field
	SamplingRatio     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityTracingSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingSettingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// When inbound trace context may be continued. Authenticated propagation is not
// supported yet.
type ObservabilityTracingSettingUpdateResponsePropagationPolicy string

const (
	ObservabilityTracingSettingUpdateResponsePropagationPolicyAccept        ObservabilityTracingSettingUpdateResponsePropagationPolicy = "accept"
	ObservabilityTracingSettingUpdateResponsePropagationPolicyAuthenticated ObservabilityTracingSettingUpdateResponsePropagationPolicy = "authenticated"
	ObservabilityTracingSettingUpdateResponsePropagationPolicyReject        ObservabilityTracingSettingUpdateResponsePropagationPolicy = "reject"
)

func (r ObservabilityTracingSettingUpdateResponsePropagationPolicy) IsKnown() bool {
	switch r {
	case ObservabilityTracingSettingUpdateResponsePropagationPolicyAccept, ObservabilityTracingSettingUpdateResponsePropagationPolicyAuthenticated, ObservabilityTracingSettingUpdateResponsePropagationPolicyReject:
		return true
	}
	return false
}

type ObservabilityTracingSettingDeleteResponse struct {
	// Up to 100 OpenTelemetry destination identifiers that receive traces.
	Destinations []string `json:"destinations" api:"required"`
	// Whether Cloudflare Traces is enabled for the zone.
	Enabled bool `json:"enabled" api:"required"`
	// Whether trace context is sent externally or across a zone boundary.
	ForwardContext bool `json:"forward_context" api:"required"`
	// Whether traces are persisted in Cloudflare.
	Persist bool `json:"persist" api:"required"`
	// When inbound trace context may be continued. Authenticated propagation is not
	// supported yet.
	PropagationPolicy ObservabilityTracingSettingDeleteResponsePropagationPolicy `json:"propagation_policy" api:"required"`
	// The ratio of requests sampled for tracing, from 0 to 1.
	SamplingRatio float64                                       `json:"sampling_ratio" api:"required"`
	JSON          observabilityTracingSettingDeleteResponseJSON `json:"-"`
}

// observabilityTracingSettingDeleteResponseJSON contains the JSON metadata for the
// struct [ObservabilityTracingSettingDeleteResponse]
type observabilityTracingSettingDeleteResponseJSON struct {
	Destinations      apijson.Field
	Enabled           apijson.Field
	ForwardContext    apijson.Field
	Persist           apijson.Field
	PropagationPolicy apijson.Field
	SamplingRatio     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityTracingSettingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingSettingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// When inbound trace context may be continued. Authenticated propagation is not
// supported yet.
type ObservabilityTracingSettingDeleteResponsePropagationPolicy string

const (
	ObservabilityTracingSettingDeleteResponsePropagationPolicyAccept        ObservabilityTracingSettingDeleteResponsePropagationPolicy = "accept"
	ObservabilityTracingSettingDeleteResponsePropagationPolicyAuthenticated ObservabilityTracingSettingDeleteResponsePropagationPolicy = "authenticated"
	ObservabilityTracingSettingDeleteResponsePropagationPolicyReject        ObservabilityTracingSettingDeleteResponsePropagationPolicy = "reject"
)

func (r ObservabilityTracingSettingDeleteResponsePropagationPolicy) IsKnown() bool {
	switch r {
	case ObservabilityTracingSettingDeleteResponsePropagationPolicyAccept, ObservabilityTracingSettingDeleteResponsePropagationPolicyAuthenticated, ObservabilityTracingSettingDeleteResponsePropagationPolicyReject:
		return true
	}
	return false
}

type ObservabilityTracingSettingGetResponse struct {
	// Up to 100 OpenTelemetry destination identifiers that receive traces.
	Destinations []string `json:"destinations" api:"required"`
	// Whether Cloudflare Traces is enabled for the zone.
	Enabled bool `json:"enabled" api:"required"`
	// Whether trace context is sent externally or across a zone boundary.
	ForwardContext bool `json:"forward_context" api:"required"`
	// Whether traces are persisted in Cloudflare.
	Persist bool `json:"persist" api:"required"`
	// When inbound trace context may be continued. Authenticated propagation is not
	// supported yet.
	PropagationPolicy ObservabilityTracingSettingGetResponsePropagationPolicy `json:"propagation_policy" api:"required"`
	// The ratio of requests sampled for tracing, from 0 to 1.
	SamplingRatio float64                                    `json:"sampling_ratio" api:"required"`
	JSON          observabilityTracingSettingGetResponseJSON `json:"-"`
}

// observabilityTracingSettingGetResponseJSON contains the JSON metadata for the
// struct [ObservabilityTracingSettingGetResponse]
type observabilityTracingSettingGetResponseJSON struct {
	Destinations      apijson.Field
	Enabled           apijson.Field
	ForwardContext    apijson.Field
	Persist           apijson.Field
	PropagationPolicy apijson.Field
	SamplingRatio     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityTracingSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// When inbound trace context may be continued. Authenticated propagation is not
// supported yet.
type ObservabilityTracingSettingGetResponsePropagationPolicy string

const (
	ObservabilityTracingSettingGetResponsePropagationPolicyAccept        ObservabilityTracingSettingGetResponsePropagationPolicy = "accept"
	ObservabilityTracingSettingGetResponsePropagationPolicyAuthenticated ObservabilityTracingSettingGetResponsePropagationPolicy = "authenticated"
	ObservabilityTracingSettingGetResponsePropagationPolicyReject        ObservabilityTracingSettingGetResponsePropagationPolicy = "reject"
)

func (r ObservabilityTracingSettingGetResponsePropagationPolicy) IsKnown() bool {
	switch r {
	case ObservabilityTracingSettingGetResponsePropagationPolicyAccept, ObservabilityTracingSettingGetResponsePropagationPolicyAuthenticated, ObservabilityTracingSettingGetResponsePropagationPolicyReject:
		return true
	}
	return false
}

type ObservabilityTracingSettingUpdateParams struct {
	// Specify the zone ID.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Up to 100 OpenTelemetry destination identifiers that receive traces.
	Destinations param.Field[[]string] `json:"destinations"`
	// Whether Cloudflare Traces is enabled for the zone.
	Enabled param.Field[bool] `json:"enabled"`
	// Whether trace context is sent externally or across a zone boundary.
	ForwardContext param.Field[bool] `json:"forward_context"`
	// Whether traces are persisted in Cloudflare.
	Persist param.Field[bool] `json:"persist"`
	// When inbound trace context may be continued. Authenticated propagation is not
	// supported yet.
	PropagationPolicy param.Field[ObservabilityTracingSettingUpdateParamsPropagationPolicy] `json:"propagation_policy"`
	// The ratio of requests sampled for tracing, from 0 to 1.
	SamplingRatio param.Field[float64] `json:"sampling_ratio"`
}

func (r ObservabilityTracingSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When inbound trace context may be continued. Authenticated propagation is not
// supported yet.
type ObservabilityTracingSettingUpdateParamsPropagationPolicy string

const (
	ObservabilityTracingSettingUpdateParamsPropagationPolicyAccept        ObservabilityTracingSettingUpdateParamsPropagationPolicy = "accept"
	ObservabilityTracingSettingUpdateParamsPropagationPolicyAuthenticated ObservabilityTracingSettingUpdateParamsPropagationPolicy = "authenticated"
	ObservabilityTracingSettingUpdateParamsPropagationPolicyReject        ObservabilityTracingSettingUpdateParamsPropagationPolicy = "reject"
)

func (r ObservabilityTracingSettingUpdateParamsPropagationPolicy) IsKnown() bool {
	switch r {
	case ObservabilityTracingSettingUpdateParamsPropagationPolicyAccept, ObservabilityTracingSettingUpdateParamsPropagationPolicyAuthenticated, ObservabilityTracingSettingUpdateParamsPropagationPolicyReject:
		return true
	}
	return false
}

type ObservabilityTracingSettingUpdateResponseEnvelope struct {
	Errors   []ObservabilityTracingSettingUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityTracingSettingUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityTracingSettingUpdateResponse                   `json:"result" api:"required"`
	Success  ObservabilityTracingSettingUpdateResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityTracingSettingUpdateResponseEnvelopeJSON       `json:"-"`
}

// observabilityTracingSettingUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ObservabilityTracingSettingUpdateResponseEnvelope]
type observabilityTracingSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingSettingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingSettingUpdateResponseEnvelopeErrors struct {
	Message string                                                      `json:"message" api:"required"`
	JSON    observabilityTracingSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityTracingSettingUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ObservabilityTracingSettingUpdateResponseEnvelopeErrors]
type observabilityTracingSettingUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingSettingUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingSettingUpdateResponseEnvelopeMessages struct {
	Message ObservabilityTracingSettingUpdateResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityTracingSettingUpdateResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityTracingSettingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ObservabilityTracingSettingUpdateResponseEnvelopeMessages]
type observabilityTracingSettingUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingSettingUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingSettingUpdateResponseEnvelopeMessagesMessage string

const (
	ObservabilityTracingSettingUpdateResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityTracingSettingUpdateResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityTracingSettingUpdateResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityTracingSettingUpdateResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityTracingSettingUpdateResponseEnvelopeSuccess bool

const (
	ObservabilityTracingSettingUpdateResponseEnvelopeSuccessTrue ObservabilityTracingSettingUpdateResponseEnvelopeSuccess = true
)

func (r ObservabilityTracingSettingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityTracingSettingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityTracingSettingDeleteParams struct {
	// Specify the zone ID.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type ObservabilityTracingSettingDeleteResponseEnvelope struct {
	Errors   []ObservabilityTracingSettingDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityTracingSettingDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityTracingSettingDeleteResponse                   `json:"result" api:"required"`
	Success  ObservabilityTracingSettingDeleteResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityTracingSettingDeleteResponseEnvelopeJSON       `json:"-"`
}

// observabilityTracingSettingDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ObservabilityTracingSettingDeleteResponseEnvelope]
type observabilityTracingSettingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingSettingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingSettingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingSettingDeleteResponseEnvelopeErrors struct {
	Message string                                                      `json:"message" api:"required"`
	JSON    observabilityTracingSettingDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityTracingSettingDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ObservabilityTracingSettingDeleteResponseEnvelopeErrors]
type observabilityTracingSettingDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingSettingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingSettingDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingSettingDeleteResponseEnvelopeMessages struct {
	Message ObservabilityTracingSettingDeleteResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityTracingSettingDeleteResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityTracingSettingDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ObservabilityTracingSettingDeleteResponseEnvelopeMessages]
type observabilityTracingSettingDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingSettingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingSettingDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingSettingDeleteResponseEnvelopeMessagesMessage string

const (
	ObservabilityTracingSettingDeleteResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityTracingSettingDeleteResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityTracingSettingDeleteResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityTracingSettingDeleteResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityTracingSettingDeleteResponseEnvelopeSuccess bool

const (
	ObservabilityTracingSettingDeleteResponseEnvelopeSuccessTrue ObservabilityTracingSettingDeleteResponseEnvelopeSuccess = true
)

func (r ObservabilityTracingSettingDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityTracingSettingDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityTracingSettingGetParams struct {
	// Specify the zone ID.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type ObservabilityTracingSettingGetResponseEnvelope struct {
	Errors   []ObservabilityTracingSettingGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityTracingSettingGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityTracingSettingGetResponse                   `json:"result" api:"required"`
	Success  ObservabilityTracingSettingGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityTracingSettingGetResponseEnvelopeJSON       `json:"-"`
}

// observabilityTracingSettingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ObservabilityTracingSettingGetResponseEnvelope]
type observabilityTracingSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingSettingGetResponseEnvelopeErrors struct {
	Message string                                                   `json:"message" api:"required"`
	JSON    observabilityTracingSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityTracingSettingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ObservabilityTracingSettingGetResponseEnvelopeErrors]
type observabilityTracingSettingGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingSettingGetResponseEnvelopeMessages struct {
	Message ObservabilityTracingSettingGetResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityTracingSettingGetResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityTracingSettingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityTracingSettingGetResponseEnvelopeMessages]
type observabilityTracingSettingGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTracingSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTracingSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTracingSettingGetResponseEnvelopeMessagesMessage string

const (
	ObservabilityTracingSettingGetResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityTracingSettingGetResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityTracingSettingGetResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityTracingSettingGetResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityTracingSettingGetResponseEnvelopeSuccess bool

const (
	ObservabilityTracingSettingGetResponseEnvelopeSuccessTrue ObservabilityTracingSettingGetResponseEnvelopeSuccess = true
)

func (r ObservabilityTracingSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityTracingSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
