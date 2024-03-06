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
func (r *StreamVideoService) StorageUsage(ctx context.Context, params StreamVideoStorageUsageParams, opts ...option.RequestOption) (res *StreamVideoStorageUsageResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamVideoStorageUsageResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/storage-usage", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamVideoStorageUsageResponse struct {
	// A user-defined identifier for the media creator.
	Creator string `json:"creator"`
	// The total minutes of video content stored in the account.
	TotalStorageMinutes int64 `json:"totalStorageMinutes"`
	// The storage capacity alloted for the account.
	TotalStorageMinutesLimit int64 `json:"totalStorageMinutesLimit"`
	// The total count of videos associated with the account.
	VideoCount int64                               `json:"videoCount"`
	JSON       streamVideoStorageUsageResponseJSON `json:"-"`
}

// streamVideoStorageUsageResponseJSON contains the JSON metadata for the struct
// [StreamVideoStorageUsageResponse]
type streamVideoStorageUsageResponseJSON struct {
	Creator                  apijson.Field
	TotalStorageMinutes      apijson.Field
	TotalStorageMinutesLimit apijson.Field
	VideoCount               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *StreamVideoStorageUsageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamVideoStorageUsageResponseJSON) RawJSON() string {
	return r.raw
}

type StreamVideoStorageUsageParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
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

type StreamVideoStorageUsageResponseEnvelope struct {
	Errors   []StreamVideoStorageUsageResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamVideoStorageUsageResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamVideoStorageUsageResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamVideoStorageUsageResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamVideoStorageUsageResponseEnvelopeJSON    `json:"-"`
}

// streamVideoStorageUsageResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamVideoStorageUsageResponseEnvelope]
type streamVideoStorageUsageResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamVideoStorageUsageResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamVideoStorageUsageResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamVideoStorageUsageResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    streamVideoStorageUsageResponseEnvelopeErrorsJSON `json:"-"`
}

// streamVideoStorageUsageResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamVideoStorageUsageResponseEnvelopeErrors]
type streamVideoStorageUsageResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamVideoStorageUsageResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamVideoStorageUsageResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamVideoStorageUsageResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    streamVideoStorageUsageResponseEnvelopeMessagesJSON `json:"-"`
}

// streamVideoStorageUsageResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [StreamVideoStorageUsageResponseEnvelopeMessages]
type streamVideoStorageUsageResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamVideoStorageUsageResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamVideoStorageUsageResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamVideoStorageUsageResponseEnvelopeSuccess bool

const (
	StreamVideoStorageUsageResponseEnvelopeSuccessTrue StreamVideoStorageUsageResponseEnvelopeSuccess = true
)
