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

// DEXFleetStatusDeviceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDEXFleetStatusDeviceService]
// method instead.
type DEXFleetStatusDeviceService struct {
	Options []option.RequestOption
}

// NewDEXFleetStatusDeviceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXFleetStatusDeviceService(opts ...option.RequestOption) (r *DEXFleetStatusDeviceService) {
	r = &DEXFleetStatusDeviceService{}
	r.Options = opts
	return
}

// List details for devices using WARP
func (r *DEXFleetStatusDeviceService) List(ctx context.Context, accountID string, query DEXFleetStatusDeviceListParams, opts ...option.RequestOption) (res *[]DEXFleetStatusDeviceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXFleetStatusDeviceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/devices", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DEXFleetStatusDeviceListResponse struct {
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
	PersonEmail string                               `json:"personEmail"`
	JSON        dexFleetStatusDeviceListResponseJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseJSON contains the JSON metadata for the struct
// [DEXFleetStatusDeviceListResponse]
type dexFleetStatusDeviceListResponseJSON struct {
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

func (r *DEXFleetStatusDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXFleetStatusDeviceListParams struct {
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
	SortBy param.Field[DEXFleetStatusDeviceListParamsSortBy] `query:"sort_by"`
	// Network status
	Status param.Field[string] `query:"status"`
	// WARP client version
	Version param.Field[string] `query:"version"`
}

// URLQuery serializes [DEXFleetStatusDeviceListParams]'s query parameters as
// `url.Values`.
func (r DEXFleetStatusDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Dimension to sort results by
type DEXFleetStatusDeviceListParamsSortBy string

const (
	DEXFleetStatusDeviceListParamsSortByColo      DEXFleetStatusDeviceListParamsSortBy = "colo"
	DEXFleetStatusDeviceListParamsSortByDeviceID  DEXFleetStatusDeviceListParamsSortBy = "device_id"
	DEXFleetStatusDeviceListParamsSortByMode      DEXFleetStatusDeviceListParamsSortBy = "mode"
	DEXFleetStatusDeviceListParamsSortByPlatform  DEXFleetStatusDeviceListParamsSortBy = "platform"
	DEXFleetStatusDeviceListParamsSortByStatus    DEXFleetStatusDeviceListParamsSortBy = "status"
	DEXFleetStatusDeviceListParamsSortByTimestamp DEXFleetStatusDeviceListParamsSortBy = "timestamp"
	DEXFleetStatusDeviceListParamsSortByVersion   DEXFleetStatusDeviceListParamsSortBy = "version"
)

type DEXFleetStatusDeviceListResponseEnvelope struct {
	Errors   []DEXFleetStatusDeviceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXFleetStatusDeviceListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DEXFleetStatusDeviceListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DEXFleetStatusDeviceListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DEXFleetStatusDeviceListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dexFleetStatusDeviceListResponseEnvelopeJSON       `json:"-"`
}

// dexFleetStatusDeviceListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DEXFleetStatusDeviceListResponseEnvelope]
type dexFleetStatusDeviceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXFleetStatusDeviceListResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    dexFleetStatusDeviceListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DEXFleetStatusDeviceListResponseEnvelopeErrors]
type dexFleetStatusDeviceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXFleetStatusDeviceListResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    dexFleetStatusDeviceListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DEXFleetStatusDeviceListResponseEnvelopeMessages]
type dexFleetStatusDeviceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DEXFleetStatusDeviceListResponseEnvelopeSuccess bool

const (
	DEXFleetStatusDeviceListResponseEnvelopeSuccessTrue DEXFleetStatusDeviceListResponseEnvelopeSuccess = true
)

type DEXFleetStatusDeviceListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       dexFleetStatusDeviceListResponseEnvelopeResultInfoJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DEXFleetStatusDeviceListResponseEnvelopeResultInfo]
type dexFleetStatusDeviceListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
