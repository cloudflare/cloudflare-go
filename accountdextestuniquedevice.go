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

// AccountDexTestUniqueDeviceService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDexTestUniqueDeviceService] method instead.
type AccountDexTestUniqueDeviceService struct {
	Options []option.RequestOption
}

// NewAccountDexTestUniqueDeviceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDexTestUniqueDeviceService(opts ...option.RequestOption) (r *AccountDexTestUniqueDeviceService) {
	r = &AccountDexTestUniqueDeviceService{}
	r.Options = opts
	return
}

// Returns unique count of devices that have run synthetic application monitoring
// tests in the past 7 days.
func (r *AccountDexTestUniqueDeviceService) List(ctx context.Context, accountIdentifier string, query AccountDexTestUniqueDeviceListParams, opts ...option.RequestOption) (res *AccountDexTestUniqueDeviceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/tests/unique-devices", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountDexTestUniqueDeviceListResponse struct {
	Errors   []AccountDexTestUniqueDeviceListResponseError   `json:"errors"`
	Messages []AccountDexTestUniqueDeviceListResponseMessage `json:"messages"`
	Result   AccountDexTestUniqueDeviceListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDexTestUniqueDeviceListResponseSuccess `json:"success"`
	JSON    accountDexTestUniqueDeviceListResponseJSON    `json:"-"`
}

// accountDexTestUniqueDeviceListResponseJSON contains the JSON metadata for the
// struct [AccountDexTestUniqueDeviceListResponse]
type accountDexTestUniqueDeviceListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestUniqueDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestUniqueDeviceListResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountDexTestUniqueDeviceListResponseErrorJSON `json:"-"`
}

// accountDexTestUniqueDeviceListResponseErrorJSON contains the JSON metadata for
// the struct [AccountDexTestUniqueDeviceListResponseError]
type accountDexTestUniqueDeviceListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestUniqueDeviceListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestUniqueDeviceListResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountDexTestUniqueDeviceListResponseMessageJSON `json:"-"`
}

// accountDexTestUniqueDeviceListResponseMessageJSON contains the JSON metadata for
// the struct [AccountDexTestUniqueDeviceListResponseMessage]
type accountDexTestUniqueDeviceListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestUniqueDeviceListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexTestUniqueDeviceListResponseResult struct {
	// total number of unique devices
	UniqueDevicesTotal int64                                            `json:"uniqueDevicesTotal,required"`
	JSON               accountDexTestUniqueDeviceListResponseResultJSON `json:"-"`
}

// accountDexTestUniqueDeviceListResponseResultJSON contains the JSON metadata for
// the struct [AccountDexTestUniqueDeviceListResponseResult]
type accountDexTestUniqueDeviceListResponseResultJSON struct {
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexTestUniqueDeviceListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDexTestUniqueDeviceListResponseSuccess bool

const (
	AccountDexTestUniqueDeviceListResponseSuccessTrue AccountDexTestUniqueDeviceListResponseSuccess = true
)

type AccountDexTestUniqueDeviceListParams struct {
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
	// Optionally filter results by test name
	TestName param.Field[string] `query:"testName"`
}

// URLQuery serializes [AccountDexTestUniqueDeviceListParams]'s query parameters as
// `url.Values`.
func (r AccountDexTestUniqueDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
