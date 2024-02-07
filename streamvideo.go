// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// StreamVideoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamVideoService] method
// instead.
type StreamVideoService struct {
	Options []option.RequestOption
}

// NewStreamVideoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamVideoService(opts ...option.RequestOption) (r *StreamVideoService) {
	r = &StreamVideoService{}
	r.Options = opts
	return
}

// Returns information about an account's storage use.
func (r *StreamVideoService) StorageUsage(ctx context.Context, accountID string, query StreamVideoStorageUsageParams, opts ...option.RequestOption) (res *StreamVideoStorageUsageResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/storage-usage", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type StreamVideoStorageUsageResponse struct {
	Errors   []StreamVideoStorageUsageResponseError   `json:"errors"`
	Messages []StreamVideoStorageUsageResponseMessage `json:"messages"`
	Result   StreamVideoStorageUsageResponseResult    `json:"result"`
	// Whether the API call was successful
	Success StreamVideoStorageUsageResponseSuccess `json:"success"`
	JSON    streamVideoStorageUsageResponseJSON    `json:"-"`
}

// streamVideoStorageUsageResponseJSON contains the JSON metadata for the struct
// [StreamVideoStorageUsageResponse]
type streamVideoStorageUsageResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamVideoStorageUsageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamVideoStorageUsageResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    streamVideoStorageUsageResponseErrorJSON `json:"-"`
}

// streamVideoStorageUsageResponseErrorJSON contains the JSON metadata for the
// struct [StreamVideoStorageUsageResponseError]
type streamVideoStorageUsageResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamVideoStorageUsageResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamVideoStorageUsageResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    streamVideoStorageUsageResponseMessageJSON `json:"-"`
}

// streamVideoStorageUsageResponseMessageJSON contains the JSON metadata for the
// struct [StreamVideoStorageUsageResponseMessage]
type streamVideoStorageUsageResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamVideoStorageUsageResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamVideoStorageUsageResponseResult struct {
	// A user-defined identifier for the media creator.
	Creator string `json:"creator"`
	// The total minutes of video content stored in the account.
	TotalStorageMinutes int64 `json:"totalStorageMinutes"`
	// The storage capacity alloted for the account.
	TotalStorageMinutesLimit int64 `json:"totalStorageMinutesLimit"`
	// The total count of videos associated with the account.
	VideoCount int64                                     `json:"videoCount"`
	JSON       streamVideoStorageUsageResponseResultJSON `json:"-"`
}

// streamVideoStorageUsageResponseResultJSON contains the JSON metadata for the
// struct [StreamVideoStorageUsageResponseResult]
type streamVideoStorageUsageResponseResultJSON struct {
	Creator                  apijson.Field
	TotalStorageMinutes      apijson.Field
	TotalStorageMinutesLimit apijson.Field
	VideoCount               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *StreamVideoStorageUsageResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamVideoStorageUsageResponseSuccess bool

const (
	StreamVideoStorageUsageResponseSuccessTrue StreamVideoStorageUsageResponseSuccess = true
)

type StreamVideoStorageUsageParams struct {
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `query:"creator"`
}

// URLQuery serializes [StreamVideoStorageUsageParams]'s query parameters as
// `url.Values`.
func (r StreamVideoStorageUsageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
