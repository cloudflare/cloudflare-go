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

// LogpushEdgeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLogpushEdgeService] method
// instead.
type LogpushEdgeService struct {
	Options []option.RequestOption
}

// NewLogpushEdgeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLogpushEdgeService(opts ...option.RequestOption) (r *LogpushEdgeService) {
	r = &LogpushEdgeService{}
	r.Options = opts
	return
}

// Creates a new Instant Logs job for a zone.
func (r *LogpushEdgeService) New(ctx context.Context, params LogpushEdgeNewParams, opts ...option.RequestOption) (res *LogpushEdgeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushEdgeNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/logpush/edge", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Instant Logs jobs for a zone.
func (r *LogpushEdgeService) Get(ctx context.Context, query LogpushEdgeGetParams, opts ...option.RequestOption) (res *[]LogpushEdgeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushEdgeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/logpush/edge", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogpushEdgeNewResponse struct {
	// Unique WebSocket address that will receive messages from Cloudflare’s edge.
	DestinationConf string `json:"destination_conf" format:"uri"`
	// Comma-separated list of fields.
	Fields string `json:"fields"`
	// Filters to drill down into specific events.
	Filter string `json:"filter"`
	// The sample parameter is the sample rate of the records set by the client:
	// "sample": 1 is 100% of records "sample": 10 is 10% and so on.
	Sample int64 `json:"sample"`
	// Unique session id of the job.
	SessionID string                     `json:"session_id"`
	JSON      logpushEdgeNewResponseJSON `json:"-"`
}

// logpushEdgeNewResponseJSON contains the JSON metadata for the struct
// [LogpushEdgeNewResponse]
type logpushEdgeNewResponseJSON struct {
	DestinationConf apijson.Field
	Fields          apijson.Field
	Filter          apijson.Field
	Sample          apijson.Field
	SessionID       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LogpushEdgeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logpushEdgeNewResponseJSON) RawJSON() string {
	return r.raw
}

type LogpushEdgeGetResponse struct {
	// Unique WebSocket address that will receive messages from Cloudflare’s edge.
	DestinationConf string `json:"destination_conf" format:"uri"`
	// Comma-separated list of fields.
	Fields string `json:"fields"`
	// Filters to drill down into specific events.
	Filter string `json:"filter"`
	// The sample parameter is the sample rate of the records set by the client:
	// "sample": 1 is 100% of records "sample": 10 is 10% and so on.
	Sample int64 `json:"sample"`
	// Unique session id of the job.
	SessionID string                     `json:"session_id"`
	JSON      logpushEdgeGetResponseJSON `json:"-"`
}

// logpushEdgeGetResponseJSON contains the JSON metadata for the struct
// [LogpushEdgeGetResponse]
type logpushEdgeGetResponseJSON struct {
	DestinationConf apijson.Field
	Fields          apijson.Field
	Filter          apijson.Field
	Sample          apijson.Field
	SessionID       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LogpushEdgeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logpushEdgeGetResponseJSON) RawJSON() string {
	return r.raw
}

type LogpushEdgeNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Comma-separated list of fields.
	Fields param.Field[string] `json:"fields"`
	// Filters to drill down into specific events.
	Filter param.Field[string] `json:"filter"`
	// The sample parameter is the sample rate of the records set by the client:
	// "sample": 1 is 100% of records "sample": 10 is 10% and so on.
	Sample param.Field[int64] `json:"sample"`
}

func (r LogpushEdgeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogpushEdgeNewResponseEnvelope struct {
	Errors   []LogpushEdgeNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushEdgeNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushEdgeNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushEdgeNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushEdgeNewResponseEnvelopeJSON    `json:"-"`
}

// logpushEdgeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogpushEdgeNewResponseEnvelope]
type logpushEdgeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushEdgeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logpushEdgeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LogpushEdgeNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    logpushEdgeNewResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushEdgeNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LogpushEdgeNewResponseEnvelopeErrors]
type logpushEdgeNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushEdgeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logpushEdgeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LogpushEdgeNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    logpushEdgeNewResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushEdgeNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LogpushEdgeNewResponseEnvelopeMessages]
type logpushEdgeNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushEdgeNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logpushEdgeNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LogpushEdgeNewResponseEnvelopeSuccess bool

const (
	LogpushEdgeNewResponseEnvelopeSuccessTrue LogpushEdgeNewResponseEnvelopeSuccess = true
)

type LogpushEdgeGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LogpushEdgeGetResponseEnvelope struct {
	Errors   []LogpushEdgeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushEdgeGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []LogpushEdgeGetResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success LogpushEdgeGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushEdgeGetResponseEnvelopeJSON    `json:"-"`
}

// logpushEdgeGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogpushEdgeGetResponseEnvelope]
type logpushEdgeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushEdgeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logpushEdgeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LogpushEdgeGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    logpushEdgeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushEdgeGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LogpushEdgeGetResponseEnvelopeErrors]
type logpushEdgeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushEdgeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logpushEdgeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LogpushEdgeGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    logpushEdgeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushEdgeGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LogpushEdgeGetResponseEnvelopeMessages]
type logpushEdgeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushEdgeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logpushEdgeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LogpushEdgeGetResponseEnvelopeSuccess bool

const (
	LogpushEdgeGetResponseEnvelopeSuccessTrue LogpushEdgeGetResponseEnvelopeSuccess = true
)
