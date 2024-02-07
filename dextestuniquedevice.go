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

// DexTestUniqueDeviceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDexTestUniqueDeviceService]
// method instead.
type DexTestUniqueDeviceService struct {
	Options []option.RequestOption
}

// NewDexTestUniqueDeviceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDexTestUniqueDeviceService(opts ...option.RequestOption) (r *DexTestUniqueDeviceService) {
	r = &DexTestUniqueDeviceService{}
	r.Options = opts
	return
}

// Returns unique count of devices that have run synthetic application monitoring
// tests in the past 7 days.
func (r *DexTestUniqueDeviceService) List(ctx context.Context, accountID string, query DexTestUniqueDeviceListParams, opts ...option.RequestOption) (res *DexTestUniqueDeviceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/tests/unique-devices", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DexTestUniqueDeviceListResponse struct {
	Errors   []DexTestUniqueDeviceListResponseError   `json:"errors"`
	Messages []DexTestUniqueDeviceListResponseMessage `json:"messages"`
	Result   DexTestUniqueDeviceListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DexTestUniqueDeviceListResponseSuccess `json:"success"`
	JSON    dexTestUniqueDeviceListResponseJSON    `json:"-"`
}

// dexTestUniqueDeviceListResponseJSON contains the JSON metadata for the struct
// [DexTestUniqueDeviceListResponse]
type dexTestUniqueDeviceListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestUniqueDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestUniqueDeviceListResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dexTestUniqueDeviceListResponseErrorJSON `json:"-"`
}

// dexTestUniqueDeviceListResponseErrorJSON contains the JSON metadata for the
// struct [DexTestUniqueDeviceListResponseError]
type dexTestUniqueDeviceListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestUniqueDeviceListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestUniqueDeviceListResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dexTestUniqueDeviceListResponseMessageJSON `json:"-"`
}

// dexTestUniqueDeviceListResponseMessageJSON contains the JSON metadata for the
// struct [DexTestUniqueDeviceListResponseMessage]
type dexTestUniqueDeviceListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexTestUniqueDeviceListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexTestUniqueDeviceListResponseResult struct {
	// total number of unique devices
	UniqueDevicesTotal int64                                     `json:"uniqueDevicesTotal,required"`
	JSON               dexTestUniqueDeviceListResponseResultJSON `json:"-"`
}

// dexTestUniqueDeviceListResponseResultJSON contains the JSON metadata for the
// struct [DexTestUniqueDeviceListResponseResult]
type dexTestUniqueDeviceListResponseResultJSON struct {
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexTestUniqueDeviceListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexTestUniqueDeviceListResponseSuccess bool

const (
	DexTestUniqueDeviceListResponseSuccessTrue DexTestUniqueDeviceListResponseSuccess = true
)

type DexTestUniqueDeviceListParams struct {
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
	// Optionally filter results by test name
	TestName param.Field[string] `query:"testName"`
}

// URLQuery serializes [DexTestUniqueDeviceListParams]'s query parameters as
// `url.Values`.
func (r DexTestUniqueDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
