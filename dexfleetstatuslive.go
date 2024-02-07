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

// DexFleetStatusLiveService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDexFleetStatusLiveService] method
// instead.
type DexFleetStatusLiveService struct {
	Options []option.RequestOption
}

// NewDexFleetStatusLiveService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDexFleetStatusLiveService(opts ...option.RequestOption) (r *DexFleetStatusLiveService) {
	r = &DexFleetStatusLiveService{}
	r.Options = opts
	return
}

// List details for live (up to 60 minutes) devices using WARP
func (r *DexFleetStatusLiveService) List(ctx context.Context, accountID string, query DexFleetStatusLiveListParams, opts ...option.RequestOption) (res *DexFleetStatusLiveListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/live", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DexFleetStatusLiveListResponse struct {
	Errors   []DexFleetStatusLiveListResponseError   `json:"errors"`
	Messages []DexFleetStatusLiveListResponseMessage `json:"messages"`
	Result   DexFleetStatusLiveListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DexFleetStatusLiveListResponseSuccess `json:"success"`
	JSON    dexFleetStatusLiveListResponseJSON    `json:"-"`
}

// dexFleetStatusLiveListResponseJSON contains the JSON metadata for the struct
// [DexFleetStatusLiveListResponse]
type dexFleetStatusLiveListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dexFleetStatusLiveListResponseErrorJSON `json:"-"`
}

// dexFleetStatusLiveListResponseErrorJSON contains the JSON metadata for the
// struct [DexFleetStatusLiveListResponseError]
type dexFleetStatusLiveListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dexFleetStatusLiveListResponseMessageJSON `json:"-"`
}

// dexFleetStatusLiveListResponseMessageJSON contains the JSON metadata for the
// struct [DexFleetStatusLiveListResponseMessage]
type dexFleetStatusLiveListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseResult struct {
	DeviceStats DexFleetStatusLiveListResponseResultDeviceStats `json:"deviceStats"`
	JSON        dexFleetStatusLiveListResponseResultJSON        `json:"-"`
}

// dexFleetStatusLiveListResponseResultJSON contains the JSON metadata for the
// struct [DexFleetStatusLiveListResponseResult]
type dexFleetStatusLiveListResponseResultJSON struct {
	DeviceStats apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseResultDeviceStats struct {
	ByColo     []DexFleetStatusLiveListResponseResultDeviceStatsByColo     `json:"byColo,nullable"`
	ByMode     []DexFleetStatusLiveListResponseResultDeviceStatsByMode     `json:"byMode,nullable"`
	ByPlatform []DexFleetStatusLiveListResponseResultDeviceStatsByPlatform `json:"byPlatform,nullable"`
	ByStatus   []DexFleetStatusLiveListResponseResultDeviceStatsByStatus   `json:"byStatus,nullable"`
	ByVersion  []DexFleetStatusLiveListResponseResultDeviceStatsByVersion  `json:"byVersion,nullable"`
	// Number of unique devices
	UniqueDevicesTotal float64                                             `json:"uniqueDevicesTotal"`
	JSON               dexFleetStatusLiveListResponseResultDeviceStatsJSON `json:"-"`
}

// dexFleetStatusLiveListResponseResultDeviceStatsJSON contains the JSON metadata
// for the struct [DexFleetStatusLiveListResponseResultDeviceStats]
type dexFleetStatusLiveListResponseResultDeviceStatsJSON struct {
	ByColo             apijson.Field
	ByMode             apijson.Field
	ByPlatform         apijson.Field
	ByStatus           apijson.Field
	ByVersion          apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseResultDeviceStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseResultDeviceStatsByColo struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                   `json:"uniqueDevicesTotal"`
	Value              string                                                    `json:"value"`
	JSON               dexFleetStatusLiveListResponseResultDeviceStatsByColoJSON `json:"-"`
}

// dexFleetStatusLiveListResponseResultDeviceStatsByColoJSON contains the JSON
// metadata for the struct [DexFleetStatusLiveListResponseResultDeviceStatsByColo]
type dexFleetStatusLiveListResponseResultDeviceStatsByColoJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseResultDeviceStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseResultDeviceStatsByMode struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                   `json:"uniqueDevicesTotal"`
	Value              string                                                    `json:"value"`
	JSON               dexFleetStatusLiveListResponseResultDeviceStatsByModeJSON `json:"-"`
}

// dexFleetStatusLiveListResponseResultDeviceStatsByModeJSON contains the JSON
// metadata for the struct [DexFleetStatusLiveListResponseResultDeviceStatsByMode]
type dexFleetStatusLiveListResponseResultDeviceStatsByModeJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseResultDeviceStatsByMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseResultDeviceStatsByPlatform struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                       `json:"uniqueDevicesTotal"`
	Value              string                                                        `json:"value"`
	JSON               dexFleetStatusLiveListResponseResultDeviceStatsByPlatformJSON `json:"-"`
}

// dexFleetStatusLiveListResponseResultDeviceStatsByPlatformJSON contains the JSON
// metadata for the struct
// [DexFleetStatusLiveListResponseResultDeviceStatsByPlatform]
type dexFleetStatusLiveListResponseResultDeviceStatsByPlatformJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseResultDeviceStatsByPlatform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseResultDeviceStatsByStatus struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                     `json:"uniqueDevicesTotal"`
	Value              string                                                      `json:"value"`
	JSON               dexFleetStatusLiveListResponseResultDeviceStatsByStatusJSON `json:"-"`
}

// dexFleetStatusLiveListResponseResultDeviceStatsByStatusJSON contains the JSON
// metadata for the struct
// [DexFleetStatusLiveListResponseResultDeviceStatsByStatus]
type dexFleetStatusLiveListResponseResultDeviceStatsByStatusJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseResultDeviceStatsByStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseResultDeviceStatsByVersion struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                      `json:"uniqueDevicesTotal"`
	Value              string                                                       `json:"value"`
	JSON               dexFleetStatusLiveListResponseResultDeviceStatsByVersionJSON `json:"-"`
}

// dexFleetStatusLiveListResponseResultDeviceStatsByVersionJSON contains the JSON
// metadata for the struct
// [DexFleetStatusLiveListResponseResultDeviceStatsByVersion]
type dexFleetStatusLiveListResponseResultDeviceStatsByVersionJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseResultDeviceStatsByVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexFleetStatusLiveListResponseSuccess bool

const (
	DexFleetStatusLiveListResponseSuccessTrue DexFleetStatusLiveListResponseSuccess = true
)

type DexFleetStatusLiveListParams struct {
	// Number of minutes before current time
	SinceMinutes param.Field[float64] `query:"since_minutes,required"`
}

// URLQuery serializes [DexFleetStatusLiveListParams]'s query parameters as
// `url.Values`.
func (r DexFleetStatusLiveListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
