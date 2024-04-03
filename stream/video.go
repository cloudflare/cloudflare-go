// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// VideoService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewVideoService] method instead.
type VideoService struct {
	Options []option.RequestOption
}

// NewVideoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVideoService(opts ...option.RequestOption) (r *VideoService) {
	r = &VideoService{}
	r.Options = opts
	return
}

// Returns information about an account's storage use.
func (r *VideoService) StorageUsage(ctx context.Context, params VideoStorageUsageParams, opts ...option.RequestOption) (res *VideoStorageUsageResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VideoStorageUsageResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/storage-usage", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type VideoStorageUsageResponse struct {
	// A user-defined identifier for the media creator.
	Creator string `json:"creator"`
	// The total minutes of video content stored in the account.
	TotalStorageMinutes int64 `json:"totalStorageMinutes"`
	// The storage capacity alloted for the account.
	TotalStorageMinutesLimit int64 `json:"totalStorageMinutesLimit"`
	// The total count of videos associated with the account.
	VideoCount int64                         `json:"videoCount"`
	JSON       videoStorageUsageResponseJSON `json:"-"`
}

// videoStorageUsageResponseJSON contains the JSON metadata for the struct
// [VideoStorageUsageResponse]
type videoStorageUsageResponseJSON struct {
	Creator                  apijson.Field
	TotalStorageMinutes      apijson.Field
	TotalStorageMinutesLimit apijson.Field
	VideoCount               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *VideoStorageUsageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoStorageUsageResponseJSON) RawJSON() string {
	return r.raw
}

type VideoStorageUsageParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `query:"creator"`
}

// URLQuery serializes [VideoStorageUsageParams]'s query parameters as
// `url.Values`.
func (r VideoStorageUsageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VideoStorageUsageResponseEnvelope struct {
	Errors   []VideoStorageUsageResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VideoStorageUsageResponseEnvelopeMessages `json:"messages,required"`
	Result   VideoStorageUsageResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success VideoStorageUsageResponseEnvelopeSuccess `json:"success,required"`
	JSON    videoStorageUsageResponseEnvelopeJSON    `json:"-"`
}

// videoStorageUsageResponseEnvelopeJSON contains the JSON metadata for the struct
// [VideoStorageUsageResponseEnvelope]
type videoStorageUsageResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoStorageUsageResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoStorageUsageResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type VideoStorageUsageResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    videoStorageUsageResponseEnvelopeErrorsJSON `json:"-"`
}

// videoStorageUsageResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VideoStorageUsageResponseEnvelopeErrors]
type videoStorageUsageResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoStorageUsageResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoStorageUsageResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type VideoStorageUsageResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    videoStorageUsageResponseEnvelopeMessagesJSON `json:"-"`
}

// videoStorageUsageResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VideoStorageUsageResponseEnvelopeMessages]
type videoStorageUsageResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoStorageUsageResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoStorageUsageResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type VideoStorageUsageResponseEnvelopeSuccess bool

const (
	VideoStorageUsageResponseEnvelopeSuccessTrue VideoStorageUsageResponseEnvelopeSuccess = true
)

func (r VideoStorageUsageResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VideoStorageUsageResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
