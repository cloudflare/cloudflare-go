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

// DexFleetStatusDeviceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDexFleetStatusDeviceService]
// method instead.
type DexFleetStatusDeviceService struct {
	Options []option.RequestOption
}

// NewDexFleetStatusDeviceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDexFleetStatusDeviceService(opts ...option.RequestOption) (r *DexFleetStatusDeviceService) {
	r = &DexFleetStatusDeviceService{}
	r.Options = opts
	return
}

// List details for devices using WARP
func (r *DexFleetStatusDeviceService) List(ctx context.Context, accountID string, query DexFleetStatusDeviceListParams, opts ...option.RequestOption) (res *DexFleetStatusDeviceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/devices", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DexFleetStatusDeviceListResponse struct {
	Errors     []DexFleetStatusDeviceListResponseError    `json:"errors"`
	Messages   []DexFleetStatusDeviceListResponseMessage  `json:"messages"`
	Result     []DexFleetStatusDeviceListResponseResult   `json:"result"`
	ResultInfo DexFleetStatusDeviceListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DexFleetStatusDeviceListResponseSuccess `json:"success"`
	JSON    dexFleetStatusDeviceListResponseJSON    `json:"-"`
}

// dexFleetStatusDeviceListResponseJSON contains the JSON metadata for the struct
// [DexFleetStatusDeviceListResponse]
type dexFleetStatusDeviceListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusDeviceListResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dexFleetStatusDeviceListResponseErrorJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseErrorJSON contains the JSON metadata for the
// struct [DexFleetStatusDeviceListResponseError]
type dexFleetStatusDeviceListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusDeviceListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusDeviceListResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    dexFleetStatusDeviceListResponseMessageJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseMessageJSON contains the JSON metadata for the
// struct [DexFleetStatusDeviceListResponseMessage]
type dexFleetStatusDeviceListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusDeviceListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusDeviceListResponseResult struct {
	// Cloudflare colo
	Colo string `json:"colo,required"`
	// Device identifier (UUID v4)
	DeviceID string `json:"deviceId,required"`
	// Operating system
	Platform string `json:"platform,required"`
	// Network status
	Status string `json:"status,required"`
	// WARP client version
	Version string `json:"version,required"`
	// Device identifier (human readable)
	DeviceName string `json:"deviceName"`
	// User contact email address
	PersonEmail string                                     `json:"personEmail"`
	JSON        dexFleetStatusDeviceListResponseResultJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseResultJSON contains the JSON metadata for the
// struct [DexFleetStatusDeviceListResponseResult]
type dexFleetStatusDeviceListResponseResultJSON struct {
	Colo        apijson.Field
	DeviceID    apijson.Field
	Platform    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	DeviceName  apijson.Field
	PersonEmail apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusDeviceListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusDeviceListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       dexFleetStatusDeviceListResponseResultInfoJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseResultInfoJSON contains the JSON metadata for
// the struct [DexFleetStatusDeviceListResponseResultInfo]
type dexFleetStatusDeviceListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusDeviceListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexFleetStatusDeviceListResponseSuccess bool

const (
	DexFleetStatusDeviceListResponseSuccessTrue DexFleetStatusDeviceListResponseSuccess = true
)

type DexFleetStatusDeviceListParams struct {
	// Page number of paginated results
	Page param.Field[float64] `query:"page,required"`
	// Number of items per page
	PerPage param.Field[float64] `query:"per_page,required"`
	// Timestamp in ISO format
	TimeEnd param.Field[string] `query:"time_end,required"`
	// Timestamp in ISO format
	TimeStart param.Field[string] `query:"time_start,required"`
	// Cloudflare colo
	Colo param.Field[string] `query:"colo"`
	// Device-specific ID, given as UUID v4
	DeviceID param.Field[string] `query:"device_id"`
	// The mode under which the WARP client is run
	Mode param.Field[string] `query:"mode"`
	// Operating system
	Platform param.Field[string] `query:"platform"`
	// Dimension to sort results by
	SortBy param.Field[DexFleetStatusDeviceListParamsSortBy] `query:"sort_by"`
	// Network status
	Status param.Field[string] `query:"status"`
	// WARP client version
	Version param.Field[string] `query:"version"`
}

// URLQuery serializes [DexFleetStatusDeviceListParams]'s query parameters as
// `url.Values`.
func (r DexFleetStatusDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Dimension to sort results by
type DexFleetStatusDeviceListParamsSortBy string

const (
	DexFleetStatusDeviceListParamsSortByColo      DexFleetStatusDeviceListParamsSortBy = "colo"
	DexFleetStatusDeviceListParamsSortByDeviceID  DexFleetStatusDeviceListParamsSortBy = "device_id"
	DexFleetStatusDeviceListParamsSortByMode      DexFleetStatusDeviceListParamsSortBy = "mode"
	DexFleetStatusDeviceListParamsSortByPlatform  DexFleetStatusDeviceListParamsSortBy = "platform"
	DexFleetStatusDeviceListParamsSortByStatus    DexFleetStatusDeviceListParamsSortBy = "status"
	DexFleetStatusDeviceListParamsSortByTimestamp DexFleetStatusDeviceListParamsSortBy = "timestamp"
	DexFleetStatusDeviceListParamsSortByVersion   DexFleetStatusDeviceListParamsSortBy = "version"
)
