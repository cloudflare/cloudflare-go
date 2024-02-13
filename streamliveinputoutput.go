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

// Deletes an output and removes it from the associated live input.
func (r *StreamLiveInputOutputService) Delete(ctx context.Context, accountID string, liveInputIdentifier string, outputIdentifier string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs/%s", accountID, liveInputIdentifier, outputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Creates a new output that can be used to simulcast or restream live video to
// other RTMP or SRT destinations. Outputs are always linked to a specific live
// input — one live input can have many outputs.
func (r *StreamLiveInputOutputService) StreamLiveInputsNewANewOutputConnectedToALiveInput(ctx context.Context, accountID string, liveInputIdentifier string, body StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputParams, opts ...option.RequestOption) (res *StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", accountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves all outputs associated with a specified live input.
func (r *StreamLiveInputOutputService) StreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInput(ctx context.Context, accountID string, liveInputIdentifier string, opts ...option.RequestOption) (res *[]StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", accountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponse struct {
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
	URL  string                                                                              `json:"url"`
	JSON streamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseJSON `json:"-"`
}

// streamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseJSON
// contains the JSON metadata for the struct
// [StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponse]
type streamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponse struct {
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
	URL  string                                                                                           `json:"url"`
	JSON streamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseJSON `json:"-"`
}

// streamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseJSON
// contains the JSON metadata for the struct
// [StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponse]
type streamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputParams struct {
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

func (r StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelope struct {
	Errors   []StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeJSON    `json:"-"`
}

// streamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelope]
type streamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeErrors struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    streamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeErrorsJSON `json:"-"`
}

// streamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeErrors]
type streamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeMessages struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    streamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeMessagesJSON `json:"-"`
}

// streamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeMessages]
type streamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeSuccess bool

const (
	StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeSuccessTrue StreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseEnvelopeSuccess = true
)

type StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelope struct {
	Errors   []StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeMessages `json:"messages,required"`
	Result   []StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeJSON    `json:"-"`
}

// streamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelope]
type streamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeErrors struct {
	Code    int64                                                                                                          `json:"code,required"`
	Message string                                                                                                         `json:"message,required"`
	JSON    streamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeErrorsJSON `json:"-"`
}

// streamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeErrors]
type streamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeMessages struct {
	Code    int64                                                                                                            `json:"code,required"`
	Message string                                                                                                           `json:"message,required"`
	JSON    streamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeMessagesJSON `json:"-"`
}

// streamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeMessages]
type streamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeSuccess bool

const (
	StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeSuccessTrue StreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseEnvelopeSuccess = true
)
