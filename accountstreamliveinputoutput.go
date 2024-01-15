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

// AccountStreamLiveInputOutputService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountStreamLiveInputOutputService] method instead.
type AccountStreamLiveInputOutputService struct {
	Options []option.RequestOption
}

// NewAccountStreamLiveInputOutputService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStreamLiveInputOutputService(opts ...option.RequestOption) (r *AccountStreamLiveInputOutputService) {
	r = &AccountStreamLiveInputOutputService{}
	r.Options = opts
	return
}

// Updates the state of an output.
func (r *AccountStreamLiveInputOutputService) Update(ctx context.Context, accountIdentifier string, liveInputIdentifier string, outputIdentifier string, body AccountStreamLiveInputOutputUpdateParams, opts ...option.RequestOption) (res *AccountStreamLiveInputOutputUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs/%s", accountIdentifier, liveInputIdentifier, outputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an output and removes it from the associated live input.
func (r *AccountStreamLiveInputOutputService) Delete(ctx context.Context, accountIdentifier string, liveInputIdentifier string, outputIdentifier string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs/%s", accountIdentifier, liveInputIdentifier, outputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Creates a new output that can be used to simulcast or restream live video to
// other RTMP or SRT destinations. Outputs are always linked to a specific live
// input — one live input can have many outputs.
func (r *AccountStreamLiveInputOutputService) StreamLiveInputsNewANewOutputConnectedToALiveInput(ctx context.Context, accountIdentifier string, liveInputIdentifier string, body AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputParams, opts ...option.RequestOption) (res *AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", accountIdentifier, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves all outputs associated with a specified live input.
func (r *AccountStreamLiveInputOutputService) StreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInput(ctx context.Context, accountIdentifier string, liveInputIdentifier string, opts ...option.RequestOption) (res *AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", accountIdentifier, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountStreamLiveInputOutputUpdateResponse struct {
	Errors   []AccountStreamLiveInputOutputUpdateResponseError   `json:"errors"`
	Messages []AccountStreamLiveInputOutputUpdateResponseMessage `json:"messages"`
	Result   AccountStreamLiveInputOutputUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamLiveInputOutputUpdateResponseSuccess `json:"success"`
	JSON    accountStreamLiveInputOutputUpdateResponseJSON    `json:"-"`
}

// accountStreamLiveInputOutputUpdateResponseJSON contains the JSON metadata for
// the struct [AccountStreamLiveInputOutputUpdateResponse]
type accountStreamLiveInputOutputUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputOutputUpdateResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountStreamLiveInputOutputUpdateResponseErrorJSON `json:"-"`
}

// accountStreamLiveInputOutputUpdateResponseErrorJSON contains the JSON metadata
// for the struct [AccountStreamLiveInputOutputUpdateResponseError]
type accountStreamLiveInputOutputUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputOutputUpdateResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountStreamLiveInputOutputUpdateResponseMessageJSON `json:"-"`
}

// accountStreamLiveInputOutputUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AccountStreamLiveInputOutputUpdateResponseMessage]
type accountStreamLiveInputOutputUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputOutputUpdateResponseResult struct {
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
	URL  string                                               `json:"url"`
	JSON accountStreamLiveInputOutputUpdateResponseResultJSON `json:"-"`
}

// accountStreamLiveInputOutputUpdateResponseResultJSON contains the JSON metadata
// for the struct [AccountStreamLiveInputOutputUpdateResponseResult]
type accountStreamLiveInputOutputUpdateResponseResultJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamLiveInputOutputUpdateResponseSuccess bool

const (
	AccountStreamLiveInputOutputUpdateResponseSuccessTrue AccountStreamLiveInputOutputUpdateResponseSuccess = true
)

type AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponse struct {
	Errors   []AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseError   `json:"errors"`
	Messages []AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseMessage `json:"messages"`
	Result   AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseSuccess `json:"success"`
	JSON    accountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseJSON    `json:"-"`
}

// accountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponse]
type accountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseError struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    accountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseErrorJSON `json:"-"`
}

// accountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseError]
type accountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseMessage struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    accountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseMessageJSON `json:"-"`
}

// accountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseMessage]
type accountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseResult struct {
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
	JSON accountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseResultJSON `json:"-"`
}

// accountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseResultJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseResult]
type accountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseResultJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseSuccess bool

const (
	AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseSuccessTrue AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputResponseSuccess = true
)

type AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponse struct {
	Errors   []AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseError   `json:"errors"`
	Messages []AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseMessage `json:"messages"`
	Result   []AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseSuccess `json:"success"`
	JSON    accountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseJSON    `json:"-"`
}

// accountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponse]
type accountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseError struct {
	Code    int64                                                                                                        `json:"code,required"`
	Message string                                                                                                       `json:"message,required"`
	JSON    accountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseErrorJSON `json:"-"`
}

// accountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseError]
type accountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseMessage struct {
	Code    int64                                                                                                          `json:"code,required"`
	Message string                                                                                                         `json:"message,required"`
	JSON    accountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseMessageJSON `json:"-"`
}

// accountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseMessage]
type accountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseResult struct {
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
	URL  string                                                                                                        `json:"url"`
	JSON accountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseResultJSON `json:"-"`
}

// accountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseResultJSON
// contains the JSON metadata for the struct
// [AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseResult]
type accountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseResultJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseSuccess bool

const (
	AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseSuccessTrue AccountStreamLiveInputOutputStreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInputResponseSuccess = true
)

type AccountStreamLiveInputOutputUpdateParams struct {
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountStreamLiveInputOutputUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputParams struct {
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

func (r AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
