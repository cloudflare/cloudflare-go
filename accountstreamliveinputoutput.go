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
func (r *AccountStreamLiveInputOutputService) Update(ctx context.Context, accountIdentifier string, liveInputIdentifier string, outputIdentifier string, body AccountStreamLiveInputOutputUpdateParams, opts ...option.RequestOption) (res *OutputResponseSingle, err error) {
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
func (r *AccountStreamLiveInputOutputService) StreamLiveInputsNewANewOutputConnectedToALiveInput(ctx context.Context, accountIdentifier string, liveInputIdentifier string, body AccountStreamLiveInputOutputStreamLiveInputsNewANewOutputConnectedToALiveInputParams, opts ...option.RequestOption) (res *OutputResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", accountIdentifier, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves all outputs associated with a specified live input.
func (r *AccountStreamLiveInputOutputService) StreamLiveInputsListAllOutputsAssociatedWithASpecifiedLiveInput(ctx context.Context, accountIdentifier string, liveInputIdentifier string, opts ...option.RequestOption) (res *OutputResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", accountIdentifier, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OutputResponseCollection struct {
	Errors     []OutputResponseCollectionError    `json:"errors"`
	Messages   []OutputResponseCollectionMessage  `json:"messages"`
	Result     []OutputResponseCollectionResult   `json:"result"`
	ResultInfo OutputResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success OutputResponseCollectionSuccess `json:"success"`
	JSON    outputResponseCollectionJSON    `json:"-"`
}

// outputResponseCollectionJSON contains the JSON metadata for the struct
// [OutputResponseCollection]
type outputResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OutputResponseCollectionError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    outputResponseCollectionErrorJSON `json:"-"`
}

// outputResponseCollectionErrorJSON contains the JSON metadata for the struct
// [OutputResponseCollectionError]
type outputResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OutputResponseCollectionMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    outputResponseCollectionMessageJSON `json:"-"`
}

// outputResponseCollectionMessageJSON contains the JSON metadata for the struct
// [OutputResponseCollectionMessage]
type outputResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OutputResponseCollectionResult struct {
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
	URL  string                             `json:"url"`
	JSON outputResponseCollectionResultJSON `json:"-"`
}

// outputResponseCollectionResultJSON contains the JSON metadata for the struct
// [OutputResponseCollectionResult]
type outputResponseCollectionResultJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OutputResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       outputResponseCollectionResultInfoJSON `json:"-"`
}

// outputResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [OutputResponseCollectionResultInfo]
type outputResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OutputResponseCollectionSuccess bool

const (
	OutputResponseCollectionSuccessTrue OutputResponseCollectionSuccess = true
)

type OutputResponseSingle struct {
	Errors   []OutputResponseSingleError   `json:"errors"`
	Messages []OutputResponseSingleMessage `json:"messages"`
	Result   OutputResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success OutputResponseSingleSuccess `json:"success"`
	JSON    outputResponseSingleJSON    `json:"-"`
}

// outputResponseSingleJSON contains the JSON metadata for the struct
// [OutputResponseSingle]
type outputResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OutputResponseSingleError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    outputResponseSingleErrorJSON `json:"-"`
}

// outputResponseSingleErrorJSON contains the JSON metadata for the struct
// [OutputResponseSingleError]
type outputResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OutputResponseSingleMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    outputResponseSingleMessageJSON `json:"-"`
}

// outputResponseSingleMessageJSON contains the JSON metadata for the struct
// [OutputResponseSingleMessage]
type outputResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OutputResponseSingleResult struct {
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
	JSON outputResponseSingleResultJSON `json:"-"`
}

// outputResponseSingleResultJSON contains the JSON metadata for the struct
// [OutputResponseSingleResult]
type outputResponseSingleResultJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OutputResponseSingleSuccess bool

const (
	OutputResponseSingleSuccessTrue OutputResponseSingleSuccess = true
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
