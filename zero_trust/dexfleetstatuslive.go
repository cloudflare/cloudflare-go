// File generated from our OpenAPI spec by Stainless.

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
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DEXFleetStatusLiveService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDEXFleetStatusLiveService] method
// instead.
type DEXFleetStatusLiveService struct {
	Options []option.RequestOption
}

// NewDEXFleetStatusLiveService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXFleetStatusLiveService(opts ...option.RequestOption) (r *DEXFleetStatusLiveService) {
	r = &DEXFleetStatusLiveService{}
	r.Options = opts
	return
}

// List details for live (up to 60 minutes) devices using WARP
func (r *DEXFleetStatusLiveService) List(ctx context.Context, params DEXFleetStatusLiveListParams, opts ...option.RequestOption) (res *DEXFleetStatusLiveListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXFleetStatusLiveListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/live", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DEXFleetStatusLiveListResponse struct {
	DeviceStats DEXFleetStatusLiveListResponseDeviceStats `json:"deviceStats"`
	JSON        dexFleetStatusLiveListResponseJSON        `json:"-"`
}

// dexFleetStatusLiveListResponseJSON contains the JSON metadata for the struct
// [DEXFleetStatusLiveListResponse]
type dexFleetStatusLiveListResponseJSON struct {
	DeviceStats apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusLiveListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveListResponseJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveListResponseDeviceStats struct {
	ByColo     []DEXFleetStatusLiveListResponseDeviceStatsByColo     `json:"byColo,nullable"`
	ByMode     []DEXFleetStatusLiveListResponseDeviceStatsByMode     `json:"byMode,nullable"`
	ByPlatform []DEXFleetStatusLiveListResponseDeviceStatsByPlatform `json:"byPlatform,nullable"`
	ByStatus   []DEXFleetStatusLiveListResponseDeviceStatsByStatus   `json:"byStatus,nullable"`
	ByVersion  []DEXFleetStatusLiveListResponseDeviceStatsByVersion  `json:"byVersion,nullable"`
	// Number of unique devices
	UniqueDevicesTotal float64                                       `json:"uniqueDevicesTotal"`
	JSON               dexFleetStatusLiveListResponseDeviceStatsJSON `json:"-"`
}

// dexFleetStatusLiveListResponseDeviceStatsJSON contains the JSON metadata for the
// struct [DEXFleetStatusLiveListResponseDeviceStats]
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

func (r *DEXFleetStatusLiveListResponseDeviceStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveListResponseDeviceStatsJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveListResponseDeviceStatsByColo struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                             `json:"uniqueDevicesTotal"`
	Value              string                                              `json:"value"`
	JSON               dexFleetStatusLiveListResponseDeviceStatsByColoJSON `json:"-"`
}

// dexFleetStatusLiveListResponseDeviceStatsByColoJSON contains the JSON metadata
// for the struct [DEXFleetStatusLiveListResponseDeviceStatsByColo]
type dexFleetStatusLiveListResponseDeviceStatsByColoJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusLiveListResponseDeviceStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveListResponseDeviceStatsByColoJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveListResponseDeviceStatsByMode struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                             `json:"uniqueDevicesTotal"`
	Value              string                                              `json:"value"`
	JSON               dexFleetStatusLiveListResponseDeviceStatsByModeJSON `json:"-"`
}

// dexFleetStatusLiveListResponseDeviceStatsByModeJSON contains the JSON metadata
// for the struct [DEXFleetStatusLiveListResponseDeviceStatsByMode]
type dexFleetStatusLiveListResponseDeviceStatsByModeJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusLiveListResponseDeviceStatsByMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveListResponseDeviceStatsByModeJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveListResponseDeviceStatsByPlatform struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                 `json:"uniqueDevicesTotal"`
	Value              string                                                  `json:"value"`
	JSON               dexFleetStatusLiveListResponseDeviceStatsByPlatformJSON `json:"-"`
}

// dexFleetStatusLiveListResponseDeviceStatsByPlatformJSON contains the JSON
// metadata for the struct [DEXFleetStatusLiveListResponseDeviceStatsByPlatform]
type dexFleetStatusLiveListResponseDeviceStatsByPlatformJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusLiveListResponseDeviceStatsByPlatform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveListResponseDeviceStatsByPlatformJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveListResponseDeviceStatsByStatus struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                               `json:"uniqueDevicesTotal"`
	Value              string                                                `json:"value"`
	JSON               dexFleetStatusLiveListResponseDeviceStatsByStatusJSON `json:"-"`
}

// dexFleetStatusLiveListResponseDeviceStatsByStatusJSON contains the JSON metadata
// for the struct [DEXFleetStatusLiveListResponseDeviceStatsByStatus]
type dexFleetStatusLiveListResponseDeviceStatsByStatusJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusLiveListResponseDeviceStatsByStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveListResponseDeviceStatsByStatusJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveListResponseDeviceStatsByVersion struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                `json:"uniqueDevicesTotal"`
	Value              string                                                 `json:"value"`
	JSON               dexFleetStatusLiveListResponseDeviceStatsByVersionJSON `json:"-"`
}

// dexFleetStatusLiveListResponseDeviceStatsByVersionJSON contains the JSON
// metadata for the struct [DEXFleetStatusLiveListResponseDeviceStatsByVersion]
type dexFleetStatusLiveListResponseDeviceStatsByVersionJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusLiveListResponseDeviceStatsByVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveListResponseDeviceStatsByVersionJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Number of minutes before current time
	SinceMinutes param.Field[float64] `query:"since_minutes,required"`
}

// URLQuery serializes [DEXFleetStatusLiveListParams]'s query parameters as
// `url.Values`.
func (r DEXFleetStatusLiveListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DEXFleetStatusLiveListResponseEnvelope struct {
	Errors   []DEXFleetStatusLiveListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXFleetStatusLiveListResponseEnvelopeMessages `json:"messages,required"`
	Result   DEXFleetStatusLiveListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DEXFleetStatusLiveListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexFleetStatusLiveListResponseEnvelopeJSON    `json:"-"`
}

// dexFleetStatusLiveListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DEXFleetStatusLiveListResponseEnvelope]
type dexFleetStatusLiveListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusLiveListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dexFleetStatusLiveListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexFleetStatusLiveListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DEXFleetStatusLiveListResponseEnvelopeErrors]
type dexFleetStatusLiveListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusLiveListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    dexFleetStatusLiveListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexFleetStatusLiveListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DEXFleetStatusLiveListResponseEnvelopeMessages]
type dexFleetStatusLiveListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusLiveListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXFleetStatusLiveListResponseEnvelopeSuccess bool

const (
	DEXFleetStatusLiveListResponseEnvelopeSuccessTrue DEXFleetStatusLiveListResponseEnvelopeSuccess = true
)
