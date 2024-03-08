// File generated from our OpenAPI spec by Stainless.

package logpush

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// EdgeService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewEdgeService] method instead.
type EdgeService struct {
	Options []option.RequestOption
}

// NewEdgeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEdgeService(opts ...option.RequestOption) (r *EdgeService) {
	r = &EdgeService{}
	r.Options = opts
	return
}

// Creates a new Instant Logs job for a zone.
func (r *EdgeService) New(ctx context.Context, params EdgeNewParams, opts ...option.RequestOption) (res *EdgeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EdgeNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/logpush/edge", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Instant Logs jobs for a zone.
func (r *EdgeService) Get(ctx context.Context, query EdgeGetParams, opts ...option.RequestOption) (res *[]EdgeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EdgeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/logpush/edge", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EdgeNewResponse struct {
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
	SessionID string              `json:"session_id"`
	JSON      edgeNewResponseJSON `json:"-"`
}

// edgeNewResponseJSON contains the JSON metadata for the struct [EdgeNewResponse]
type edgeNewResponseJSON struct {
	DestinationConf apijson.Field
	Fields          apijson.Field
	Filter          apijson.Field
	Sample          apijson.Field
	SessionID       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EdgeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeNewResponseJSON) RawJSON() string {
	return r.raw
}

type EdgeGetResponse struct {
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
	SessionID string              `json:"session_id"`
	JSON      edgeGetResponseJSON `json:"-"`
}

// edgeGetResponseJSON contains the JSON metadata for the struct [EdgeGetResponse]
type edgeGetResponseJSON struct {
	DestinationConf apijson.Field
	Fields          apijson.Field
	Filter          apijson.Field
	Sample          apijson.Field
	SessionID       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EdgeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeGetResponseJSON) RawJSON() string {
	return r.raw
}

type EdgeNewParams struct {
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

func (r EdgeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EdgeNewResponseEnvelope struct {
	Errors   []EdgeNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EdgeNewResponseEnvelopeMessages `json:"messages,required"`
	Result   EdgeNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success EdgeNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    edgeNewResponseEnvelopeJSON    `json:"-"`
}

// edgeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [EdgeNewResponseEnvelope]
type edgeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EdgeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EdgeNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    edgeNewResponseEnvelopeErrorsJSON `json:"-"`
}

// edgeNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [EdgeNewResponseEnvelopeErrors]
type edgeNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EdgeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EdgeNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    edgeNewResponseEnvelopeMessagesJSON `json:"-"`
}

// edgeNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [EdgeNewResponseEnvelopeMessages]
type edgeNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EdgeNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EdgeNewResponseEnvelopeSuccess bool

const (
	EdgeNewResponseEnvelopeSuccessTrue EdgeNewResponseEnvelopeSuccess = true
)

type EdgeGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type EdgeGetResponseEnvelope struct {
	Errors   []EdgeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EdgeGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []EdgeGetResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success EdgeGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    edgeGetResponseEnvelopeJSON    `json:"-"`
}

// edgeGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EdgeGetResponseEnvelope]
type edgeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EdgeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EdgeGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    edgeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// edgeGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [EdgeGetResponseEnvelopeErrors]
type edgeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EdgeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EdgeGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    edgeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// edgeGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [EdgeGetResponseEnvelopeMessages]
type edgeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EdgeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EdgeGetResponseEnvelopeSuccess bool

const (
	EdgeGetResponseEnvelopeSuccessTrue EdgeGetResponseEnvelopeSuccess = true
)
