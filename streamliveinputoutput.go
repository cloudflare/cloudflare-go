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

// StreamLiveInputOutputService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamLiveInputOutputService]
// method instead.
type StreamLiveInputOutputService struct {
	Options []option.RequestOption
}

// NewStreamLiveInputOutputService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStreamLiveInputOutputService(opts ...option.RequestOption) (r *StreamLiveInputOutputService) {
	r = &StreamLiveInputOutputService{}
	r.Options = opts
	return
}

// Creates a new output that can be used to simulcast or restream live video to
// other RTMP or SRT destinations. Outputs are always linked to a specific live
// input — one live input can have many outputs.
func (r *StreamLiveInputOutputService) New(ctx context.Context, accountID string, liveInputIdentifier string, body StreamLiveInputOutputNewParams, opts ...option.RequestOption) (res *StreamLiveInputOutputNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputOutputNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", accountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the state of an output.
func (r *StreamLiveInputOutputService) Update(ctx context.Context, accountID string, liveInputIdentifier string, outputIdentifier string, body StreamLiveInputOutputUpdateParams, opts ...option.RequestOption) (res *StreamLiveInputOutputUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputOutputUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs/%s", accountID, liveInputIdentifier, outputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves all outputs associated with a specified live input.
func (r *StreamLiveInputOutputService) List(ctx context.Context, accountID string, liveInputIdentifier string, opts ...option.RequestOption) (res *[]StreamLiveInputOutputListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputOutputListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", accountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an output and removes it from the associated live input.
func (r *StreamLiveInputOutputService) Delete(ctx context.Context, accountID string, liveInputIdentifier string, outputIdentifier string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs/%s", accountID, liveInputIdentifier, outputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type StreamLiveInputOutputNewResponse struct {
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled bool `json:"enabled"`
	// The streamKey used to authenticate against an output's target.
	StreamKey string `json:"streamKey"`
	// A unique identifier for the output.
	Uid string `json:"uid"`
	// The URL an output uses to restream.
	URL  string                               `json:"url"`
	JSON streamLiveInputOutputNewResponseJSON `json:"-"`
}

// streamLiveInputOutputNewResponseJSON contains the JSON metadata for the struct
// [StreamLiveInputOutputNewResponse]
type streamLiveInputOutputNewResponseJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputUpdateResponse struct {
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled bool `json:"enabled"`
	// The streamKey used to authenticate against an output's target.
	StreamKey string `json:"streamKey"`
	// A unique identifier for the output.
	Uid string `json:"uid"`
	// The URL an output uses to restream.
	URL  string                                  `json:"url"`
	JSON streamLiveInputOutputUpdateResponseJSON `json:"-"`
}

// streamLiveInputOutputUpdateResponseJSON contains the JSON metadata for the
// struct [StreamLiveInputOutputUpdateResponse]
type streamLiveInputOutputUpdateResponseJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputListResponse struct {
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled bool `json:"enabled"`
	// The streamKey used to authenticate against an output's target.
	StreamKey string `json:"streamKey"`
	// A unique identifier for the output.
	Uid string `json:"uid"`
	// The URL an output uses to restream.
	URL  string                                `json:"url"`
	JSON streamLiveInputOutputListResponseJSON `json:"-"`
}

// streamLiveInputOutputListResponseJSON contains the JSON metadata for the struct
// [StreamLiveInputOutputListResponse]
type streamLiveInputOutputListResponseJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputNewParams struct {
	// The streamKey used to authenticate against an output's target.
	StreamKey param.Field[string] `json:"streamKey,required"`
	// The URL an output uses to restream.
	URL param.Field[string] `json:"url,required"`
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r StreamLiveInputOutputNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamLiveInputOutputNewResponseEnvelope struct {
	Errors   []StreamLiveInputOutputNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamLiveInputOutputNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamLiveInputOutputNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamLiveInputOutputNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamLiveInputOutputNewResponseEnvelopeJSON    `json:"-"`
}

// streamLiveInputOutputNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamLiveInputOutputNewResponseEnvelope]
type streamLiveInputOutputNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    streamLiveInputOutputNewResponseEnvelopeErrorsJSON `json:"-"`
}

// streamLiveInputOutputNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [StreamLiveInputOutputNewResponseEnvelopeErrors]
type streamLiveInputOutputNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    streamLiveInputOutputNewResponseEnvelopeMessagesJSON `json:"-"`
}

// streamLiveInputOutputNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [StreamLiveInputOutputNewResponseEnvelopeMessages]
type streamLiveInputOutputNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamLiveInputOutputNewResponseEnvelopeSuccess bool

const (
	StreamLiveInputOutputNewResponseEnvelopeSuccessTrue StreamLiveInputOutputNewResponseEnvelopeSuccess = true
)

type StreamLiveInputOutputUpdateParams struct {
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r StreamLiveInputOutputUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamLiveInputOutputUpdateResponseEnvelope struct {
	Errors   []StreamLiveInputOutputUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamLiveInputOutputUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamLiveInputOutputUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamLiveInputOutputUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamLiveInputOutputUpdateResponseEnvelopeJSON    `json:"-"`
}

// streamLiveInputOutputUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [StreamLiveInputOutputUpdateResponseEnvelope]
type streamLiveInputOutputUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputUpdateResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    streamLiveInputOutputUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// streamLiveInputOutputUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [StreamLiveInputOutputUpdateResponseEnvelopeErrors]
type streamLiveInputOutputUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputUpdateResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    streamLiveInputOutputUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// streamLiveInputOutputUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [StreamLiveInputOutputUpdateResponseEnvelopeMessages]
type streamLiveInputOutputUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamLiveInputOutputUpdateResponseEnvelopeSuccess bool

const (
	StreamLiveInputOutputUpdateResponseEnvelopeSuccessTrue StreamLiveInputOutputUpdateResponseEnvelopeSuccess = true
)

type StreamLiveInputOutputListResponseEnvelope struct {
	Errors   []StreamLiveInputOutputListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamLiveInputOutputListResponseEnvelopeMessages `json:"messages,required"`
	Result   []StreamLiveInputOutputListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success StreamLiveInputOutputListResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamLiveInputOutputListResponseEnvelopeJSON    `json:"-"`
}

// streamLiveInputOutputListResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamLiveInputOutputListResponseEnvelope]
type streamLiveInputOutputListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputListResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    streamLiveInputOutputListResponseEnvelopeErrorsJSON `json:"-"`
}

// streamLiveInputOutputListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [StreamLiveInputOutputListResponseEnvelopeErrors]
type streamLiveInputOutputListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputListResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    streamLiveInputOutputListResponseEnvelopeMessagesJSON `json:"-"`
}

// streamLiveInputOutputListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [StreamLiveInputOutputListResponseEnvelopeMessages]
type streamLiveInputOutputListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamLiveInputOutputListResponseEnvelopeSuccess bool

const (
	StreamLiveInputOutputListResponseEnvelopeSuccessTrue StreamLiveInputOutputListResponseEnvelopeSuccess = true
)
