// File generated from our OpenAPI spec by Stainless.

package stream

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// LiveInputOutputService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLiveInputOutputService] method
// instead.
type LiveInputOutputService struct {
	Options []option.RequestOption
}

// NewLiveInputOutputService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLiveInputOutputService(opts ...option.RequestOption) (r *LiveInputOutputService) {
	r = &LiveInputOutputService{}
	r.Options = opts
	return
}

// Creates a new output that can be used to simulcast or restream live video to
// other RTMP or SRT destinations. Outputs are always linked to a specific live
// input — one live input can have many outputs.
func (r *LiveInputOutputService) New(ctx context.Context, liveInputIdentifier string, params LiveInputOutputNewParams, opts ...option.RequestOption) (res *LiveInputOutputNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LiveInputOutputNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", params.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the state of an output.
func (r *LiveInputOutputService) Update(ctx context.Context, liveInputIdentifier string, outputIdentifier string, params LiveInputOutputUpdateParams, opts ...option.RequestOption) (res *LiveInputOutputUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LiveInputOutputUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs/%s", params.AccountID, liveInputIdentifier, outputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves all outputs associated with a specified live input.
func (r *LiveInputOutputService) List(ctx context.Context, liveInputIdentifier string, query LiveInputOutputListParams, opts ...option.RequestOption) (res *[]LiveInputOutputListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LiveInputOutputListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", query.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an output and removes it from the associated live input.
func (r *LiveInputOutputService) Delete(ctx context.Context, liveInputIdentifier string, outputIdentifier string, body LiveInputOutputDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs/%s", body.AccountID, liveInputIdentifier, outputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type LiveInputOutputNewResponse struct {
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
	URL  string                         `json:"url"`
	JSON liveInputOutputNewResponseJSON `json:"-"`
}

// liveInputOutputNewResponseJSON contains the JSON metadata for the struct
// [LiveInputOutputNewResponse]
type liveInputOutputNewResponseJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputNewResponseJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputUpdateResponse struct {
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
	URL  string                            `json:"url"`
	JSON liveInputOutputUpdateResponseJSON `json:"-"`
}

// liveInputOutputUpdateResponseJSON contains the JSON metadata for the struct
// [LiveInputOutputUpdateResponse]
type liveInputOutputUpdateResponseJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputListResponse struct {
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
	URL  string                          `json:"url"`
	JSON liveInputOutputListResponseJSON `json:"-"`
}

// liveInputOutputListResponseJSON contains the JSON metadata for the struct
// [LiveInputOutputListResponse]
type liveInputOutputListResponseJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputListResponseJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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

func (r LiveInputOutputNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LiveInputOutputNewResponseEnvelope struct {
	Errors   []LiveInputOutputNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputOutputNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LiveInputOutputNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LiveInputOutputNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    liveInputOutputNewResponseEnvelopeJSON    `json:"-"`
}

// liveInputOutputNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LiveInputOutputNewResponseEnvelope]
type liveInputOutputNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputNewResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    liveInputOutputNewResponseEnvelopeErrorsJSON `json:"-"`
}

// liveInputOutputNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LiveInputOutputNewResponseEnvelopeErrors]
type liveInputOutputNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputNewResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    liveInputOutputNewResponseEnvelopeMessagesJSON `json:"-"`
}

// liveInputOutputNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LiveInputOutputNewResponseEnvelopeMessages]
type liveInputOutputNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LiveInputOutputNewResponseEnvelopeSuccess bool

const (
	LiveInputOutputNewResponseEnvelopeSuccessTrue LiveInputOutputNewResponseEnvelopeSuccess = true
)

type LiveInputOutputUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r LiveInputOutputUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LiveInputOutputUpdateResponseEnvelope struct {
	Errors   []LiveInputOutputUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputOutputUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   LiveInputOutputUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LiveInputOutputUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    liveInputOutputUpdateResponseEnvelopeJSON    `json:"-"`
}

// liveInputOutputUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [LiveInputOutputUpdateResponseEnvelope]
type liveInputOutputUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    liveInputOutputUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// liveInputOutputUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LiveInputOutputUpdateResponseEnvelopeErrors]
type liveInputOutputUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    liveInputOutputUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// liveInputOutputUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LiveInputOutputUpdateResponseEnvelopeMessages]
type liveInputOutputUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LiveInputOutputUpdateResponseEnvelopeSuccess bool

const (
	LiveInputOutputUpdateResponseEnvelopeSuccessTrue LiveInputOutputUpdateResponseEnvelopeSuccess = true
)

type LiveInputOutputListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type LiveInputOutputListResponseEnvelope struct {
	Errors   []LiveInputOutputListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputOutputListResponseEnvelopeMessages `json:"messages,required"`
	Result   []LiveInputOutputListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success LiveInputOutputListResponseEnvelopeSuccess `json:"success,required"`
	JSON    liveInputOutputListResponseEnvelopeJSON    `json:"-"`
}

// liveInputOutputListResponseEnvelopeJSON contains the JSON metadata for the
// struct [LiveInputOutputListResponseEnvelope]
type liveInputOutputListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    liveInputOutputListResponseEnvelopeErrorsJSON `json:"-"`
}

// liveInputOutputListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LiveInputOutputListResponseEnvelopeErrors]
type liveInputOutputListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    liveInputOutputListResponseEnvelopeMessagesJSON `json:"-"`
}

// liveInputOutputListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LiveInputOutputListResponseEnvelopeMessages]
type liveInputOutputListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LiveInputOutputListResponseEnvelopeSuccess bool

const (
	LiveInputOutputListResponseEnvelopeSuccessTrue LiveInputOutputListResponseEnvelopeSuccess = true
)

type LiveInputOutputDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
