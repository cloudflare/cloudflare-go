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

// AccountDexFleetStatusLiveService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDexFleetStatusLiveService] method instead.
type AccountDexFleetStatusLiveService struct {
	Options []option.RequestOption
}

// NewAccountDexFleetStatusLiveService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDexFleetStatusLiveService(opts ...option.RequestOption) (r *AccountDexFleetStatusLiveService) {
	r = &AccountDexFleetStatusLiveService{}
	r.Options = opts
	return
}

// List details for live (up to 60 minutes) devices using WARP
func (r *AccountDexFleetStatusLiveService) List(ctx context.Context, accountIdentifier string, query AccountDexFleetStatusLiveListParams, opts ...option.RequestOption) (res *AccountDexFleetStatusLiveListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/live", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountDexFleetStatusLiveListResponse struct {
	Errors   []AccountDexFleetStatusLiveListResponseError   `json:"errors"`
	Messages []AccountDexFleetStatusLiveListResponseMessage `json:"messages"`
	Result   AccountDexFleetStatusLiveListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDexFleetStatusLiveListResponseSuccess `json:"success"`
	JSON    accountDexFleetStatusLiveListResponseJSON    `json:"-"`
}

// accountDexFleetStatusLiveListResponseJSON contains the JSON metadata for the
// struct [AccountDexFleetStatusLiveListResponse]
type accountDexFleetStatusLiveListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexFleetStatusLiveListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexFleetStatusLiveListResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountDexFleetStatusLiveListResponseErrorJSON `json:"-"`
}

// accountDexFleetStatusLiveListResponseErrorJSON contains the JSON metadata for
// the struct [AccountDexFleetStatusLiveListResponseError]
type accountDexFleetStatusLiveListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexFleetStatusLiveListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexFleetStatusLiveListResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountDexFleetStatusLiveListResponseMessageJSON `json:"-"`
}

// accountDexFleetStatusLiveListResponseMessageJSON contains the JSON metadata for
// the struct [AccountDexFleetStatusLiveListResponseMessage]
type accountDexFleetStatusLiveListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexFleetStatusLiveListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexFleetStatusLiveListResponseResult struct {
	DeviceStats AccountDexFleetStatusLiveListResponseResultDeviceStats `json:"deviceStats"`
	JSON        accountDexFleetStatusLiveListResponseResultJSON        `json:"-"`
}

// accountDexFleetStatusLiveListResponseResultJSON contains the JSON metadata for
// the struct [AccountDexFleetStatusLiveListResponseResult]
type accountDexFleetStatusLiveListResponseResultJSON struct {
	DeviceStats apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexFleetStatusLiveListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexFleetStatusLiveListResponseResultDeviceStats struct {
	ByColo     []AccountDexFleetStatusLiveListResponseResultDeviceStatsByColo     `json:"byColo,nullable"`
	ByMode     []AccountDexFleetStatusLiveListResponseResultDeviceStatsByMode     `json:"byMode,nullable"`
	ByPlatform []AccountDexFleetStatusLiveListResponseResultDeviceStatsByPlatform `json:"byPlatform,nullable"`
	ByStatus   []AccountDexFleetStatusLiveListResponseResultDeviceStatsByStatus   `json:"byStatus,nullable"`
	ByVersion  []AccountDexFleetStatusLiveListResponseResultDeviceStatsByVersion  `json:"byVersion,nullable"`
	// Number of unique devices
	UniqueDevicesTotal float64                                                    `json:"uniqueDevicesTotal"`
	JSON               accountDexFleetStatusLiveListResponseResultDeviceStatsJSON `json:"-"`
}

// accountDexFleetStatusLiveListResponseResultDeviceStatsJSON contains the JSON
// metadata for the struct [AccountDexFleetStatusLiveListResponseResultDeviceStats]
type accountDexFleetStatusLiveListResponseResultDeviceStatsJSON struct {
	ByColo             apijson.Field
	ByMode             apijson.Field
	ByPlatform         apijson.Field
	ByStatus           apijson.Field
	ByVersion          apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexFleetStatusLiveListResponseResultDeviceStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexFleetStatusLiveListResponseResultDeviceStatsByColo struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                          `json:"uniqueDevicesTotal"`
	Value              string                                                           `json:"value"`
	JSON               accountDexFleetStatusLiveListResponseResultDeviceStatsByColoJSON `json:"-"`
}

// accountDexFleetStatusLiveListResponseResultDeviceStatsByColoJSON contains the
// JSON metadata for the struct
// [AccountDexFleetStatusLiveListResponseResultDeviceStatsByColo]
type accountDexFleetStatusLiveListResponseResultDeviceStatsByColoJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexFleetStatusLiveListResponseResultDeviceStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexFleetStatusLiveListResponseResultDeviceStatsByMode struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                          `json:"uniqueDevicesTotal"`
	Value              string                                                           `json:"value"`
	JSON               accountDexFleetStatusLiveListResponseResultDeviceStatsByModeJSON `json:"-"`
}

// accountDexFleetStatusLiveListResponseResultDeviceStatsByModeJSON contains the
// JSON metadata for the struct
// [AccountDexFleetStatusLiveListResponseResultDeviceStatsByMode]
type accountDexFleetStatusLiveListResponseResultDeviceStatsByModeJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexFleetStatusLiveListResponseResultDeviceStatsByMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexFleetStatusLiveListResponseResultDeviceStatsByPlatform struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                              `json:"uniqueDevicesTotal"`
	Value              string                                                               `json:"value"`
	JSON               accountDexFleetStatusLiveListResponseResultDeviceStatsByPlatformJSON `json:"-"`
}

// accountDexFleetStatusLiveListResponseResultDeviceStatsByPlatformJSON contains
// the JSON metadata for the struct
// [AccountDexFleetStatusLiveListResponseResultDeviceStatsByPlatform]
type accountDexFleetStatusLiveListResponseResultDeviceStatsByPlatformJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexFleetStatusLiveListResponseResultDeviceStatsByPlatform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexFleetStatusLiveListResponseResultDeviceStatsByStatus struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                            `json:"uniqueDevicesTotal"`
	Value              string                                                             `json:"value"`
	JSON               accountDexFleetStatusLiveListResponseResultDeviceStatsByStatusJSON `json:"-"`
}

// accountDexFleetStatusLiveListResponseResultDeviceStatsByStatusJSON contains the
// JSON metadata for the struct
// [AccountDexFleetStatusLiveListResponseResultDeviceStatsByStatus]
type accountDexFleetStatusLiveListResponseResultDeviceStatsByStatusJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexFleetStatusLiveListResponseResultDeviceStatsByStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexFleetStatusLiveListResponseResultDeviceStatsByVersion struct {
	// Number of unique devices
	UniqueDevicesTotal float64                                                             `json:"uniqueDevicesTotal"`
	Value              string                                                              `json:"value"`
	JSON               accountDexFleetStatusLiveListResponseResultDeviceStatsByVersionJSON `json:"-"`
}

// accountDexFleetStatusLiveListResponseResultDeviceStatsByVersionJSON contains the
// JSON metadata for the struct
// [AccountDexFleetStatusLiveListResponseResultDeviceStatsByVersion]
type accountDexFleetStatusLiveListResponseResultDeviceStatsByVersionJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexFleetStatusLiveListResponseResultDeviceStatsByVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDexFleetStatusLiveListResponseSuccess bool

const (
	AccountDexFleetStatusLiveListResponseSuccessTrue AccountDexFleetStatusLiveListResponseSuccess = true
)

type AccountDexFleetStatusLiveListParams struct {
	// Number of minutes before current time
	SinceMinutes param.Field[float64] `query:"since_minutes,required"`
}

// URLQuery serializes [AccountDexFleetStatusLiveListParams]'s query parameters as
// `url.Values`.
func (r AccountDexFleetStatusLiveListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
