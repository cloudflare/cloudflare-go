// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// StreamCaptionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamCaptionService] method
// instead.
type StreamCaptionService struct {
	Options []option.RequestOption
}

// NewStreamCaptionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamCaptionService(opts ...option.RequestOption) (r *StreamCaptionService) {
	r = &StreamCaptionService{}
	r.Options = opts
	return
}

// Uploads the caption or subtitle file to the endpoint for a specific BCP47
// language. One caption or subtitle file per language is allowed.
func (r *StreamCaptionService) Update(ctx context.Context, identifier string, language string, params StreamCaptionUpdateParams, opts ...option.RequestOption) (res *StreamCaptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamCaptionUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", params.AccountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Removes the captions or subtitles from a video.
func (r *StreamCaptionService) Delete(ctx context.Context, identifier string, language string, body StreamCaptionDeleteParams, opts ...option.RequestOption) (res *StreamCaptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamCaptionDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", body.AccountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the available captions or subtitles for a specific video.
func (r *StreamCaptionService) Get(ctx context.Context, identifier string, query StreamCaptionGetParams, opts ...option.RequestOption) (res *[]StreamCaptions, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamCaptionGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/captions", query.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamCaptions struct {
	// The language label displayed in the native language to users.
	Label string `json:"label"`
	// The language tag in BCP 47 format.
	Language string             `json:"language"`
	JSON     streamCaptionsJSON `json:"-"`
}

// streamCaptionsJSON contains the JSON metadata for the struct [StreamCaptions]
type streamCaptionsJSON struct {
	Label       apijson.Field
	Language    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamCaptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [StreamCaptionUpdateResponseUnknown] or [shared.UnionString].
type StreamCaptionUpdateResponse interface {
	ImplementsStreamCaptionUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamCaptionUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StreamCaptionDeleteResponseUnknown],
// [StreamCaptionDeleteResponseArray] or [shared.UnionString].
type StreamCaptionDeleteResponse interface {
	ImplementsStreamCaptionDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamCaptionDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StreamCaptionDeleteResponseArray []interface{}

func (r StreamCaptionDeleteResponseArray) ImplementsStreamCaptionDeleteResponse() {}

type StreamCaptionUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The WebVTT file containing the caption or subtitle content.
	File param.Field[string] `json:"file,required"`
}

func (r StreamCaptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamCaptionUpdateResponseEnvelope struct {
	Errors   []StreamCaptionUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamCaptionUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamCaptionUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamCaptionUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamCaptionUpdateResponseEnvelopeJSON    `json:"-"`
}

// streamCaptionUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamCaptionUpdateResponseEnvelope]
type streamCaptionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamCaptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamCaptionUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    streamCaptionUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// streamCaptionUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamCaptionUpdateResponseEnvelopeErrors]
type streamCaptionUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamCaptionUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamCaptionUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    streamCaptionUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// streamCaptionUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamCaptionUpdateResponseEnvelopeMessages]
type streamCaptionUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamCaptionUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamCaptionUpdateResponseEnvelopeSuccess bool

const (
	StreamCaptionUpdateResponseEnvelopeSuccessTrue StreamCaptionUpdateResponseEnvelopeSuccess = true
)

type StreamCaptionDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamCaptionDeleteResponseEnvelope struct {
	Errors   []StreamCaptionDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamCaptionDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamCaptionDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamCaptionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamCaptionDeleteResponseEnvelopeJSON    `json:"-"`
}

// streamCaptionDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamCaptionDeleteResponseEnvelope]
type streamCaptionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamCaptionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamCaptionDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    streamCaptionDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// streamCaptionDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamCaptionDeleteResponseEnvelopeErrors]
type streamCaptionDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamCaptionDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamCaptionDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    streamCaptionDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// streamCaptionDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamCaptionDeleteResponseEnvelopeMessages]
type streamCaptionDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamCaptionDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamCaptionDeleteResponseEnvelopeSuccess bool

const (
	StreamCaptionDeleteResponseEnvelopeSuccessTrue StreamCaptionDeleteResponseEnvelopeSuccess = true
)

type StreamCaptionGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamCaptionGetResponseEnvelope struct {
	Errors   []StreamCaptionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamCaptionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []StreamCaptions                           `json:"result,required"`
	// Whether the API call was successful
	Success StreamCaptionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamCaptionGetResponseEnvelopeJSON    `json:"-"`
}

// streamCaptionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamCaptionGetResponseEnvelope]
type streamCaptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamCaptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamCaptionGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    streamCaptionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// streamCaptionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamCaptionGetResponseEnvelopeErrors]
type streamCaptionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamCaptionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamCaptionGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    streamCaptionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// streamCaptionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamCaptionGetResponseEnvelopeMessages]
type streamCaptionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamCaptionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamCaptionGetResponseEnvelopeSuccess bool

const (
	StreamCaptionGetResponseEnvelopeSuccessTrue StreamCaptionGetResponseEnvelopeSuccess = true
)
