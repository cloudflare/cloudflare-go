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
	var env DexFleetStatusLiveListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/live", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DexFleetStatusLiveListResponse struct {
	DeviceStats DexFleetStatusLiveListResponseDeviceStats `json:"deviceStats"`
	JSON        dexFleetStatusLiveListResponseJSON        `json:"-"`
}

// dexFleetStatusLiveListResponseJSON contains the JSON metadata for the struct
// [DexFleetStatusLiveListResponse]
type dexFleetStatusLiveListResponseJSON struct {
	DeviceStats apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseDeviceStats struct {
	ByColo     []DexFleetStatusLiveListResponseDeviceStatsByColo     `json:"byColo,nullable"`
	ByMode     []DexFleetStatusLiveListResponseDeviceStatsByMode     `json:"byMode,nullable"`
	ByPlatform []DexFleetStatusLiveListResponseDeviceStatsByPlatform `json:"byPlatform,nullable"`
	ByStatus   []DexFleetStatusLiveListResponseDeviceStatsByStatus   `json:"byStatus,nullable"`
	ByVersion  []DexFleetStatusLiveListResponseDeviceStatsByVersion  `json:"byVersion,nullable"`
	// Number of unique devices
	UniqueDevicesTotal float64                                       `json:"uniqueDevicesTotal"`
	JSON               dexFleetStatusLiveListResponseDeviceStatsJSON `json:"-"`
}

// dexFleetStatusLiveListResponseDeviceStatsJSON contains the JSON metadata for the
// struct [DexFleetStatusLiveListResponseDeviceStats]
type dexFleetStatusLiveListResponseDeviceStatsJSON struct {
	ByColo             apijson.Field
	ByMode             apijson.Field
	ByPlatform         apijson.Field
	ByStatus           apijson.Field
	ByVersion          apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseDeviceStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseDeviceStatsByColo struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                             `json:"uniqueDevicesTotal"`
	Value              string                                              `json:"value"`
	JSON               dexFleetStatusLiveListResponseDeviceStatsByColoJSON `json:"-"`
}

// dexFleetStatusLiveListResponseDeviceStatsByColoJSON contains the JSON metadata
// for the struct [DexFleetStatusLiveListResponseDeviceStatsByColo]
type dexFleetStatusLiveListResponseDeviceStatsByColoJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseDeviceStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseDeviceStatsByMode struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                             `json:"uniqueDevicesTotal"`
	Value              string                                              `json:"value"`
	JSON               dexFleetStatusLiveListResponseDeviceStatsByModeJSON `json:"-"`
}

// dexFleetStatusLiveListResponseDeviceStatsByModeJSON contains the JSON metadata
// for the struct [DexFleetStatusLiveListResponseDeviceStatsByMode]
type dexFleetStatusLiveListResponseDeviceStatsByModeJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseDeviceStatsByMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseDeviceStatsByPlatform struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                 `json:"uniqueDevicesTotal"`
	Value              string                                                  `json:"value"`
	JSON               dexFleetStatusLiveListResponseDeviceStatsByPlatformJSON `json:"-"`
}

// dexFleetStatusLiveListResponseDeviceStatsByPlatformJSON contains the JSON
// metadata for the struct [DexFleetStatusLiveListResponseDeviceStatsByPlatform]
type dexFleetStatusLiveListResponseDeviceStatsByPlatformJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseDeviceStatsByPlatform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseDeviceStatsByStatus struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                               `json:"uniqueDevicesTotal"`
	Value              string                                                `json:"value"`
	JSON               dexFleetStatusLiveListResponseDeviceStatsByStatusJSON `json:"-"`
}

// dexFleetStatusLiveListResponseDeviceStatsByStatusJSON contains the JSON metadata
// for the struct [DexFleetStatusLiveListResponseDeviceStatsByStatus]
type dexFleetStatusLiveListResponseDeviceStatsByStatusJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseDeviceStatsByStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseDeviceStatsByVersion struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                `json:"uniqueDevicesTotal"`
	Value              string                                                 `json:"value"`
	JSON               dexFleetStatusLiveListResponseDeviceStatsByVersionJSON `json:"-"`
}

// dexFleetStatusLiveListResponseDeviceStatsByVersionJSON contains the JSON
// metadata for the struct [DexFleetStatusLiveListResponseDeviceStatsByVersion]
type dexFleetStatusLiveListResponseDeviceStatsByVersionJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseDeviceStatsByVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type DexFleetStatusLiveListResponseEnvelope struct {
	Errors   []DexFleetStatusLiveListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DexFleetStatusLiveListResponseEnvelopeMessages `json:"messages,required"`
	Result   DexFleetStatusLiveListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DexFleetStatusLiveListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexFleetStatusLiveListResponseEnvelopeJSON    `json:"-"`
}

// dexFleetStatusLiveListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DexFleetStatusLiveListResponseEnvelope]
type dexFleetStatusLiveListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dexFleetStatusLiveListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexFleetStatusLiveListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DexFleetStatusLiveListResponseEnvelopeErrors]
type dexFleetStatusLiveListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DexFleetStatusLiveListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    dexFleetStatusLiveListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexFleetStatusLiveListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DexFleetStatusLiveListResponseEnvelopeMessages]
type dexFleetStatusLiveListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexFleetStatusLiveListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DexFleetStatusLiveListResponseEnvelopeSuccess bool

const (
	DexFleetStatusLiveListResponseEnvelopeSuccessTrue DexFleetStatusLiveListResponseEnvelopeSuccess = true
)
