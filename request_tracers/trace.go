// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package request_tracers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// TraceService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewTraceService] method instead.
type TraceService struct {
	Options []option.RequestOption
}

// NewTraceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTraceService(opts ...option.RequestOption) (r *TraceService) {
	r = &TraceService{}
	r.Options = opts
	return
}

// Request Trace
func (r *TraceService) New(ctx context.Context, accountIdentifier string, body TraceNewParams, opts ...option.RequestOption) (res *TraceNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TraceNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/request-tracer/trace", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RequestTrace []RequestTraceItem

// Trace result with an origin status code
type TraceNewResponse struct {
	// HTTP Status code of zone response
	StatusCode int64                `json:"status_code"`
	Trace      RequestTrace         `json:"trace"`
	JSON       traceNewResponseJSON `json:"-"`
}

// traceNewResponseJSON contains the JSON metadata for the struct
// [TraceNewResponse]
type traceNewResponseJSON struct {
	StatusCode  apijson.Field
	Trace       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TraceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r traceNewResponseJSON) RawJSON() string {
	return r.raw
}

type TraceNewParams struct {
	// HTTP Method of tracing request
	Method param.Field[string] `json:"method,required"`
	// URL to which perform tracing request
	URL  param.Field[string]             `json:"url,required"`
	Body param.Field[TraceNewParamsBody] `json:"body"`
	// Additional request parameters
	Context param.Field[TraceNewParamsContext] `json:"context"`
	// Cookies added to tracing request
	Cookies param.Field[map[string]string] `json:"cookies"`
	// Headers added to tracing request
	Headers param.Field[map[string]string] `json:"headers"`
	// HTTP Protocol of tracing request
	Protocol param.Field[string] `json:"protocol"`
	// Skip sending the request to the Origin server after all rules evaluation
	SkipResponse param.Field[bool] `json:"skip_response"`
}

func (r TraceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TraceNewParamsBody struct {
	// Base64 encoded request body
	Base64 param.Field[string] `json:"base64"`
	// Arbitrary json as request body
	Json param.Field[interface{}] `json:"json"`
	// Request body as plain text
	PlainText param.Field[string] `json:"plain_text"`
}

func (r TraceNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional request parameters
type TraceNewParamsContext struct {
	// Bot score used for evaluating tracing request processing
	BotScore param.Field[int64] `json:"bot_score"`
	// Geodata for tracing request
	Geoloc param.Field[TraceNewParamsContextGeoloc] `json:"geoloc"`
	// Whether to skip any challenges for tracing request (e.g.: captcha)
	SkipChallenge param.Field[bool] `json:"skip_challenge"`
	// Threat score used for evaluating tracing request processing
	ThreatScore param.Field[int64] `json:"threat_score"`
}

func (r TraceNewParamsContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Geodata for tracing request
type TraceNewParamsContextGeoloc struct {
	City                param.Field[string]  `json:"city"`
	Continent           param.Field[string]  `json:"continent"`
	IsEuCountry         param.Field[bool]    `json:"is_eu_country"`
	ISOCode             param.Field[string]  `json:"iso_code"`
	Latitude            param.Field[float64] `json:"latitude"`
	Longitude           param.Field[float64] `json:"longitude"`
	PostalCode          param.Field[string]  `json:"postal_code"`
	RegionCode          param.Field[string]  `json:"region_code"`
	Subdivision2ISOCode param.Field[string]  `json:"subdivision_2_iso_code"`
	Timezone            param.Field[string]  `json:"timezone"`
}

func (r TraceNewParamsContextGeoloc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TraceNewResponseEnvelope struct {
	Errors   []TraceNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TraceNewResponseEnvelopeMessages `json:"messages,required"`
	// Trace result with an origin status code
	Result TraceNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success TraceNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    traceNewResponseEnvelopeJSON    `json:"-"`
}

// traceNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TraceNewResponseEnvelope]
type traceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TraceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r traceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TraceNewResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    traceNewResponseEnvelopeErrorsJSON `json:"-"`
}

// traceNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TraceNewResponseEnvelopeErrors]
type traceNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TraceNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r traceNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TraceNewResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    traceNewResponseEnvelopeMessagesJSON `json:"-"`
}

// traceNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TraceNewResponseEnvelopeMessages]
type traceNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TraceNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r traceNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TraceNewResponseEnvelopeSuccess bool

const (
	TraceNewResponseEnvelopeSuccessTrue TraceNewResponseEnvelopeSuccess = true
)

func (r TraceNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TraceNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
