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
func (r *LogpushEdgeService) Update(ctx context.Context, zoneID string, body LogpushEdgeUpdateParams, opts ...option.RequestOption) (res *LogpushEdgeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushEdgeUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/logpush/edge", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Instant Logs jobs for a zone.
func (r *LogpushEdgeService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *[]LogpushEdgeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushEdgeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/logpush/edge", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogpushEdgeUpdateResponse struct {
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
	SessionID string                        `json:"session_id"`
	JSON      logpushEdgeUpdateResponseJSON `json:"-"`
}

// logpushEdgeUpdateResponseJSON contains the JSON metadata for the struct
// [LogpushEdgeUpdateResponse]
type logpushEdgeUpdateResponseJSON struct {
	DestinationConf apijson.Field
	Fields          apijson.Field
	Filter          apijson.Field
	Sample          apijson.Field
	SessionID       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LogpushEdgeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type LogpushEdgeUpdateParams struct {
	// Comma-separated list of fields.
	Fields param.Field[string] `json:"fields"`
	// Filters to drill down into specific events.
	Filter param.Field[string] `json:"filter"`
	// The sample parameter is the sample rate of the records set by the client:
	// "sample": 1 is 100% of records "sample": 10 is 10% and so on.
	Sample param.Field[int64] `json:"sample"`
}

func (r LogpushEdgeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogpushEdgeUpdateResponseEnvelope struct {
	Errors   []LogpushEdgeUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushEdgeUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushEdgeUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushEdgeUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushEdgeUpdateResponseEnvelopeJSON    `json:"-"`
}

// logpushEdgeUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogpushEdgeUpdateResponseEnvelope]
type logpushEdgeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushEdgeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushEdgeUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    logpushEdgeUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushEdgeUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LogpushEdgeUpdateResponseEnvelopeErrors]
type logpushEdgeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushEdgeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushEdgeUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    logpushEdgeUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushEdgeUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LogpushEdgeUpdateResponseEnvelopeMessages]
type logpushEdgeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushEdgeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushEdgeUpdateResponseEnvelopeSuccess bool

const (
	LogpushEdgeUpdateResponseEnvelopeSuccessTrue LogpushEdgeUpdateResponseEnvelopeSuccess = true
)

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

// Whether the API call was successful
type LogpushEdgeGetResponseEnvelopeSuccess bool

const (
	LogpushEdgeGetResponseEnvelopeSuccessTrue LogpushEdgeGetResponseEnvelopeSuccess = true
)
