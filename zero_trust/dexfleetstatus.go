// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DEXFleetStatusService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDEXFleetStatusService] method
// instead.
type DEXFleetStatusService struct {
	Options []option.RequestOption
	Devices *DEXFleetStatusDeviceService
}

// NewDEXFleetStatusService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDEXFleetStatusService(opts ...option.RequestOption) (r *DEXFleetStatusService) {
	r = &DEXFleetStatusService{}
	r.Options = opts
	r.Devices = NewDEXFleetStatusDeviceService(opts...)
	return
}

// List details for live (up to 60 minutes) devices using WARP
func (r *DEXFleetStatusService) Live(ctx context.Context, params DEXFleetStatusLiveParams, opts ...option.RequestOption) (res *DEXFleetStatusLiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXFleetStatusLiveResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/live", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List details for devices using WARP, up to 7 days
func (r *DEXFleetStatusService) OverTime(ctx context.Context, params DEXFleetStatusOverTimeParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/over-time", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

type DEXFleetStatusLiveResponse struct {
	DeviceStats DEXFleetStatusLiveResponseDeviceStats `json:"deviceStats"`
	JSON        dexFleetStatusLiveResponseJSON        `json:"-"`
}

// dexFleetStatusLiveResponseJSON contains the JSON metadata for the struct
// [DEXFleetStatusLiveResponse]
type dexFleetStatusLiveResponseJSON struct {
	DeviceStats apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveResponseDeviceStats struct {
	ByColo     []DEXFleetStatusLiveResponseDeviceStatsByColo     `json:"byColo,nullable"`
	ByMode     []DEXFleetStatusLiveResponseDeviceStatsByMode     `json:"byMode,nullable"`
	ByPlatform []DEXFleetStatusLiveResponseDeviceStatsByPlatform `json:"byPlatform,nullable"`
	ByStatus   []DEXFleetStatusLiveResponseDeviceStatsByStatus   `json:"byStatus,nullable"`
	ByVersion  []DEXFleetStatusLiveResponseDeviceStatsByVersion  `json:"byVersion,nullable"`
	// Number of unique devices
	UniqueDevicesTotal float64                                   `json:"uniqueDevicesTotal"`
	JSON               dexFleetStatusLiveResponseDeviceStatsJSON `json:"-"`
}

// dexFleetStatusLiveResponseDeviceStatsJSON contains the JSON metadata for the
// struct [DEXFleetStatusLiveResponseDeviceStats]
type dexFleetStatusLiveResponseDeviceStatsJSON struct {
	ByColo             apijson.Field
	ByMode             apijson.Field
	ByPlatform         apijson.Field
	ByStatus           apijson.Field
	ByVersion          apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseDeviceStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseDeviceStatsJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveResponseDeviceStatsByColo struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                         `json:"uniqueDevicesTotal"`
	Value              string                                          `json:"value"`
	JSON               dexFleetStatusLiveResponseDeviceStatsByColoJSON `json:"-"`
}

// dexFleetStatusLiveResponseDeviceStatsByColoJSON contains the JSON metadata for
// the struct [DEXFleetStatusLiveResponseDeviceStatsByColo]
type dexFleetStatusLiveResponseDeviceStatsByColoJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseDeviceStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseDeviceStatsByColoJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveResponseDeviceStatsByMode struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                         `json:"uniqueDevicesTotal"`
	Value              string                                          `json:"value"`
	JSON               dexFleetStatusLiveResponseDeviceStatsByModeJSON `json:"-"`
}

// dexFleetStatusLiveResponseDeviceStatsByModeJSON contains the JSON metadata for
// the struct [DEXFleetStatusLiveResponseDeviceStatsByMode]
type dexFleetStatusLiveResponseDeviceStatsByModeJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseDeviceStatsByMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseDeviceStatsByModeJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveResponseDeviceStatsByPlatform struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                             `json:"uniqueDevicesTotal"`
	Value              string                                              `json:"value"`
	JSON               dexFleetStatusLiveResponseDeviceStatsByPlatformJSON `json:"-"`
}

// dexFleetStatusLiveResponseDeviceStatsByPlatformJSON contains the JSON metadata
// for the struct [DEXFleetStatusLiveResponseDeviceStatsByPlatform]
type dexFleetStatusLiveResponseDeviceStatsByPlatformJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseDeviceStatsByPlatform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseDeviceStatsByPlatformJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveResponseDeviceStatsByStatus struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                           `json:"uniqueDevicesTotal"`
	Value              string                                            `json:"value"`
	JSON               dexFleetStatusLiveResponseDeviceStatsByStatusJSON `json:"-"`
}

// dexFleetStatusLiveResponseDeviceStatsByStatusJSON contains the JSON metadata for
// the struct [DEXFleetStatusLiveResponseDeviceStatsByStatus]
type dexFleetStatusLiveResponseDeviceStatsByStatusJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseDeviceStatsByStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseDeviceStatsByStatusJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveResponseDeviceStatsByVersion struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                            `json:"uniqueDevicesTotal"`
	Value              string                                             `json:"value"`
	JSON               dexFleetStatusLiveResponseDeviceStatsByVersionJSON `json:"-"`
}

// dexFleetStatusLiveResponseDeviceStatsByVersionJSON contains the JSON metadata
// for the struct [DEXFleetStatusLiveResponseDeviceStatsByVersion]
type dexFleetStatusLiveResponseDeviceStatsByVersionJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseDeviceStatsByVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseDeviceStatsByVersionJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Number of minutes before current time
	SinceMinutes param.Field[float64] `query:"since_minutes,required"`
}

// URLQuery serializes [DEXFleetStatusLiveParams]'s query parameters as
// `url.Values`.
func (r DEXFleetStatusLiveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DEXFleetStatusLiveResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   DEXFleetStatusLiveResponse   `json:"result,required"`
	// Whether the API call was successful
	Success DEXFleetStatusLiveResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexFleetStatusLiveResponseEnvelopeJSON    `json:"-"`
}

// dexFleetStatusLiveResponseEnvelopeJSON contains the JSON metadata for the struct
// [DEXFleetStatusLiveResponseEnvelope]
type dexFleetStatusLiveResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXFleetStatusLiveResponseEnvelopeSuccess bool

const (
	DEXFleetStatusLiveResponseEnvelopeSuccessTrue DEXFleetStatusLiveResponseEnvelopeSuccess = true
)

func (r DEXFleetStatusLiveResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXFleetStatusLiveResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DEXFleetStatusOverTimeParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Timestamp in ISO format
	TimeEnd param.Field[string] `query:"time_end,required"`
	// Timestamp in ISO format
	TimeStart param.Field[string] `query:"time_start,required"`
	// Cloudflare colo
	Colo param.Field[string] `query:"colo"`
	// Device-specific ID, given as UUID v4
	DeviceID param.Field[string] `query:"device_id"`
}

// URLQuery serializes [DEXFleetStatusOverTimeParams]'s query parameters as
// `url.Values`.
func (r DEXFleetStatusOverTimeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
