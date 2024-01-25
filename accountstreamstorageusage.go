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

// AccountStreamStorageUsageService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountStreamStorageUsageService] method instead.
type AccountStreamStorageUsageService struct {
	Options []option.RequestOption
}

// NewAccountStreamStorageUsageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStreamStorageUsageService(opts ...option.RequestOption) (r *AccountStreamStorageUsageService) {
	r = &AccountStreamStorageUsageService{}
	r.Options = opts
	return
}

// Returns information about an account's storage use.
func (r *AccountStreamStorageUsageService) Get(ctx context.Context, accountIdentifier string, query AccountStreamStorageUsageGetParams, opts ...option.RequestOption) (res *AccountStreamStorageUsageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/storage-usage", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountStreamStorageUsageGetResponse struct {
	Errors   []AccountStreamStorageUsageGetResponseError   `json:"errors"`
	Messages []AccountStreamStorageUsageGetResponseMessage `json:"messages"`
	Result   AccountStreamStorageUsageGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamStorageUsageGetResponseSuccess `json:"success"`
	JSON    accountStreamStorageUsageGetResponseJSON    `json:"-"`
}

// accountStreamStorageUsageGetResponseJSON contains the JSON metadata for the
// struct [AccountStreamStorageUsageGetResponse]
type accountStreamStorageUsageGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamStorageUsageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamStorageUsageGetResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountStreamStorageUsageGetResponseErrorJSON `json:"-"`
}

// accountStreamStorageUsageGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamStorageUsageGetResponseError]
type accountStreamStorageUsageGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamStorageUsageGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamStorageUsageGetResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountStreamStorageUsageGetResponseMessageJSON `json:"-"`
}

// accountStreamStorageUsageGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountStreamStorageUsageGetResponseMessage]
type accountStreamStorageUsageGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamStorageUsageGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamStorageUsageGetResponseResult struct {
	// A user-defined identifier for the media creator.
	Creator string `json:"creator"`
	// The total minutes of video content stored in the account.
	TotalStorageMinutes int64 `json:"totalStorageMinutes"`
	// The storage capacity alloted for the account.
	TotalStorageMinutesLimit int64 `json:"totalStorageMinutesLimit"`
	// The total count of videos associated with the account.
	VideoCount int64                                          `json:"videoCount"`
	JSON       accountStreamStorageUsageGetResponseResultJSON `json:"-"`
}

// accountStreamStorageUsageGetResponseResultJSON contains the JSON metadata for
// the struct [AccountStreamStorageUsageGetResponseResult]
type accountStreamStorageUsageGetResponseResultJSON struct {
	Creator                  apijson.Field
	TotalStorageMinutes      apijson.Field
	TotalStorageMinutesLimit apijson.Field
	VideoCount               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccountStreamStorageUsageGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamStorageUsageGetResponseSuccess bool

const (
	AccountStreamStorageUsageGetResponseSuccessTrue AccountStreamStorageUsageGetResponseSuccess = true
)

type AccountStreamStorageUsageGetParams struct {
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `query:"creator"`
}

// URLQuery serializes [AccountStreamStorageUsageGetParams]'s query parameters as
// `url.Values`.
func (r AccountStreamStorageUsageGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
