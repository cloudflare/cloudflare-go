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

// ZeroTrustDEXFleetStatusLiveService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustDEXFleetStatusLiveService] method instead.
type ZeroTrustDEXFleetStatusLiveService struct {
	Options []option.RequestOption
}

// NewZeroTrustDEXFleetStatusLiveService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDEXFleetStatusLiveService(opts ...option.RequestOption) (r *ZeroTrustDEXFleetStatusLiveService) {
	r = &ZeroTrustDEXFleetStatusLiveService{}
	r.Options = opts
	return
}

// List details for live (up to 60 minutes) devices using WARP
func (r *ZeroTrustDEXFleetStatusLiveService) List(ctx context.Context, params ZeroTrustDEXFleetStatusLiveListParams, opts ...option.RequestOption) (res *ZeroTrustDEXFleetStatusLiveListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDEXFleetStatusLiveListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/live", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDEXFleetStatusLiveListResponse struct {
	DeviceStats ZeroTrustDEXFleetStatusLiveListResponseDeviceStats `json:"deviceStats"`
	JSON        zeroTrustDEXFleetStatusLiveListResponseJSON        `json:"-"`
}

// zeroTrustDEXFleetStatusLiveListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDEXFleetStatusLiveListResponse]
type zeroTrustDEXFleetStatusLiveListResponseJSON struct {
	DeviceStats apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXFleetStatusLiveListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXFleetStatusLiveListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXFleetStatusLiveListResponseDeviceStats struct {
	ByColo     []ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByColo     `json:"byColo,nullable"`
	ByMode     []ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByMode     `json:"byMode,nullable"`
	ByPlatform []ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByPlatform `json:"byPlatform,nullable"`
	ByStatus   []ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByStatus   `json:"byStatus,nullable"`
	ByVersion  []ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByVersion  `json:"byVersion,nullable"`
	// Number of unique devices
	UniqueDevicesTotal float64                                                `json:"uniqueDevicesTotal"`
	JSON               zeroTrustDEXFleetStatusLiveListResponseDeviceStatsJSON `json:"-"`
}

// zeroTrustDEXFleetStatusLiveListResponseDeviceStatsJSON contains the JSON
// metadata for the struct [ZeroTrustDEXFleetStatusLiveListResponseDeviceStats]
type zeroTrustDEXFleetStatusLiveListResponseDeviceStatsJSON struct {
	ByColo             apijson.Field
	ByMode             apijson.Field
	ByPlatform         apijson.Field
	ByStatus           apijson.Field
	ByVersion          apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustDEXFleetStatusLiveListResponseDeviceStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXFleetStatusLiveListResponseDeviceStatsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByColo struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                      `json:"uniqueDevicesTotal"`
	Value              string                                                       `json:"value"`
	JSON               zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByColoJSON `json:"-"`
}

// zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByColoJSON contains the JSON
// metadata for the struct
// [ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByColo]
type zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByColoJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByColoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByMode struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                      `json:"uniqueDevicesTotal"`
	Value              string                                                       `json:"value"`
	JSON               zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByModeJSON `json:"-"`
}

// zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByModeJSON contains the JSON
// metadata for the struct
// [ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByMode]
type zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByModeJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByModeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByPlatform struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                          `json:"uniqueDevicesTotal"`
	Value              string                                                           `json:"value"`
	JSON               zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByPlatformJSON `json:"-"`
}

// zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByPlatformJSON contains the
// JSON metadata for the struct
// [ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByPlatform]
type zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByPlatformJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByPlatform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByPlatformJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByStatus struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                        `json:"uniqueDevicesTotal"`
	Value              string                                                         `json:"value"`
	JSON               zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByStatusJSON `json:"-"`
}

// zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByStatusJSON contains the JSON
// metadata for the struct
// [ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByStatus]
type zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByStatusJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByStatusJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByVersion struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                         `json:"uniqueDevicesTotal"`
	Value              string                                                          `json:"value"`
	JSON               zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByVersionJSON `json:"-"`
}

// zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByVersionJSON contains the
// JSON metadata for the struct
// [ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByVersion]
type zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByVersionJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustDEXFleetStatusLiveListResponseDeviceStatsByVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXFleetStatusLiveListResponseDeviceStatsByVersionJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXFleetStatusLiveListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Number of minutes before current time
	SinceMinutes param.Field[float64] `query:"since_minutes,required"`
}

// URLQuery serializes [ZeroTrustDEXFleetStatusLiveListParams]'s query parameters
// as `url.Values`.
func (r ZeroTrustDEXFleetStatusLiveListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZeroTrustDEXFleetStatusLiveListResponseEnvelope struct {
	Errors   []ZeroTrustDEXFleetStatusLiveListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDEXFleetStatusLiveListResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDEXFleetStatusLiveListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDEXFleetStatusLiveListResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDEXFleetStatusLiveListResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDEXFleetStatusLiveListResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustDEXFleetStatusLiveListResponseEnvelope]
type zeroTrustDEXFleetStatusLiveListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXFleetStatusLiveListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXFleetStatusLiveListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXFleetStatusLiveListResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustDEXFleetStatusLiveListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDEXFleetStatusLiveListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDEXFleetStatusLiveListResponseEnvelopeErrors]
type zeroTrustDEXFleetStatusLiveListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXFleetStatusLiveListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXFleetStatusLiveListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXFleetStatusLiveListResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustDEXFleetStatusLiveListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDEXFleetStatusLiveListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustDEXFleetStatusLiveListResponseEnvelopeMessages]
type zeroTrustDEXFleetStatusLiveListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDEXFleetStatusLiveListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXFleetStatusLiveListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustDEXFleetStatusLiveListResponseEnvelopeSuccess bool

const (
	ZeroTrustDEXFleetStatusLiveListResponseEnvelopeSuccessTrue ZeroTrustDEXFleetStatusLiveListResponseEnvelopeSuccess = true
)
