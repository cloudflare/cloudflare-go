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

// RequestTracerTraceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRequestTracerTraceService] method
// instead.
type RequestTracerTraceService struct {
	Options []option.RequestOption
}

// NewRequestTracerTraceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRequestTracerTraceService(opts ...option.RequestOption) (r *RequestTracerTraceService) {
	r = &RequestTracerTraceService{}
	r.Options = opts
	return
}

// Request Trace
func (r *RequestTracerTraceService) New(ctx context.Context, accountIdentifier string, body RequestTracerTraceNewParams, opts ...option.RequestOption) (res *RequestTracerTraceNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestTracerTraceNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/request-tracer/trace", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type THy7ZvEaTrace []THy7ZvEaTrace

// Trace result with an origin status code
type RequestTracerTraceNewResponse struct {
	// HTTP Status code of zone response
	StatusCode int64                             `json:"status_code"`
	Trace      THy7ZvEaTrace                     `json:"trace"`
	JSON       requestTracerTraceNewResponseJSON `json:"-"`
}

// requestTracerTraceNewResponseJSON contains the JSON metadata for the struct
// [RequestTracerTraceNewResponse]
type requestTracerTraceNewResponseJSON struct {
	StatusCode  apijson.Field
	Trace       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestTracerTraceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RequestTracerTraceNewParams struct {
	// HTTP Method of tracing request
	Method param.Field[string] `json:"method,required"`
	// URL to which perform tracing request
	URL  param.Field[string]                          `json:"url,required"`
	Body param.Field[RequestTracerTraceNewParamsBody] `json:"body"`
	// Additional request parameters
	Context param.Field[RequestTracerTraceNewParamsContext] `json:"context"`
	// Cookies added to tracing request
	Cookies param.Field[map[string]string] `json:"cookies"`
	// Headers added to tracing request
	Headers param.Field[map[string]string] `json:"headers"`
	// HTTP Protocol of tracing request
	Protocol param.Field[string] `json:"protocol"`
	// Skip sending the request to the Origin server after all rules evaluation
	SkipResponse param.Field[bool] `json:"skip_response"`
}

func (r RequestTracerTraceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RequestTracerTraceNewParamsBody struct {
	// Base64 encoded request body
	Base64 param.Field[string] `json:"base64"`
	// Arbitrary json as request body
	Json param.Field[interface{}] `json:"json"`
	// Request body as plain text
	PlainText param.Field[string] `json:"plain_text"`
}

func (r RequestTracerTraceNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional request parameters
type RequestTracerTraceNewParamsContext struct {
	// Bot score used for evaluating tracing request processing
	BotScore param.Field[int64] `json:"bot_score"`
	// Geodata for tracing request
	Geoloc param.Field[RequestTracerTraceNewParamsContextGeoloc] `json:"geoloc"`
	// Whether to skip any challenges for tracing request (e.g.: captcha)
	SkipChallenge param.Field[bool] `json:"skip_challenge"`
	// Threat score used for evaluating tracing request processing
	ThreatScore param.Field[int64] `json:"threat_score"`
}

func (r RequestTracerTraceNewParamsContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Geodata for tracing request
type RequestTracerTraceNewParamsContextGeoloc struct {
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

func (r RequestTracerTraceNewParamsContextGeoloc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RequestTracerTraceNewResponseEnvelope struct {
	Errors   []RequestTracerTraceNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestTracerTraceNewResponseEnvelopeMessages `json:"messages,required"`
	// Trace result with an origin status code
	Result RequestTracerTraceNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success RequestTracerTraceNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestTracerTraceNewResponseEnvelopeJSON    `json:"-"`
}

// requestTracerTraceNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestTracerTraceNewResponseEnvelope]
type requestTracerTraceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestTracerTraceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RequestTracerTraceNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    requestTracerTraceNewResponseEnvelopeErrorsJSON `json:"-"`
}

// requestTracerTraceNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RequestTracerTraceNewResponseEnvelopeErrors]
type requestTracerTraceNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestTracerTraceNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RequestTracerTraceNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    requestTracerTraceNewResponseEnvelopeMessagesJSON `json:"-"`
}

// requestTracerTraceNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestTracerTraceNewResponseEnvelopeMessages]
type requestTracerTraceNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestTracerTraceNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RequestTracerTraceNewResponseEnvelopeSuccess bool

const (
	RequestTracerTraceNewResponseEnvelopeSuccessTrue RequestTracerTraceNewResponseEnvelopeSuccess = true
)
