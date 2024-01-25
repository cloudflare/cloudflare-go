// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSpeedAPIService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSpeedAPIService] method
// instead.
type ZoneSpeedAPIService struct {
	Options  []option.RequestOption
	Pages    *ZoneSpeedAPIPageService
	Schedule *ZoneSpeedAPIScheduleService
}

// NewZoneSpeedAPIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSpeedAPIService(opts ...option.RequestOption) (r *ZoneSpeedAPIService) {
	r = &ZoneSpeedAPIService{}
	r.Options = opts
	r.Pages = NewZoneSpeedAPIPageService(opts...)
	r.Schedule = NewZoneSpeedAPIScheduleService(opts...)
	return
}

// Retrieves quota for all plans, as well as the current zone quota.
func (r *ZoneSpeedAPIService) Availabilities(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSpeedAPIAvailabilitiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/availabilities", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSpeedAPIAvailabilitiesResponse struct {
	Errors   []ZoneSpeedAPIAvailabilitiesResponseError   `json:"errors"`
	Messages []ZoneSpeedAPIAvailabilitiesResponseMessage `json:"messages"`
	Result   ZoneSpeedAPIAvailabilitiesResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                                   `json:"success"`
	JSON    zoneSpeedAPIAvailabilitiesResponseJSON `json:"-"`
}

// zoneSpeedAPIAvailabilitiesResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIAvailabilitiesResponse]
type zoneSpeedAPIAvailabilitiesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIAvailabilitiesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIAvailabilitiesResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSpeedAPIAvailabilitiesResponseErrorJSON `json:"-"`
}

// zoneSpeedAPIAvailabilitiesResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIAvailabilitiesResponseError]
type zoneSpeedAPIAvailabilitiesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIAvailabilitiesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIAvailabilitiesResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSpeedAPIAvailabilitiesResponseMessageJSON `json:"-"`
}

// zoneSpeedAPIAvailabilitiesResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIAvailabilitiesResponseMessage]
type zoneSpeedAPIAvailabilitiesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIAvailabilitiesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIAvailabilitiesResponseResult struct {
	Quota          ZoneSpeedAPIAvailabilitiesResponseResultQuota    `json:"quota"`
	Regions        []ZoneSpeedAPIAvailabilitiesResponseResultRegion `json:"regions"`
	RegionsPerPlan interface{}                                      `json:"regionsPerPlan"`
	JSON           zoneSpeedAPIAvailabilitiesResponseResultJSON     `json:"-"`
}

// zoneSpeedAPIAvailabilitiesResponseResultJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIAvailabilitiesResponseResult]
type zoneSpeedAPIAvailabilitiesResponseResultJSON struct {
	Quota          apijson.Field
	Regions        apijson.Field
	RegionsPerPlan apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneSpeedAPIAvailabilitiesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIAvailabilitiesResponseResultQuota struct {
	// Cloudflare plan.
	Plan string `json:"plan"`
	// The number of tests available per plan.
	QuotasPerPlan interface{} `json:"quotasPerPlan"`
	// The number of remaining schedules available.
	RemainingSchedules float64 `json:"remainingSchedules"`
	// The number of remaining tests available.
	RemainingTests float64 `json:"remainingTests"`
	// The number of schedules available per plan.
	ScheduleQuotasPerPlan interface{}                                       `json:"scheduleQuotasPerPlan"`
	JSON                  zoneSpeedAPIAvailabilitiesResponseResultQuotaJSON `json:"-"`
}

// zoneSpeedAPIAvailabilitiesResponseResultQuotaJSON contains the JSON metadata for
// the struct [ZoneSpeedAPIAvailabilitiesResponseResultQuota]
type zoneSpeedAPIAvailabilitiesResponseResultQuotaJSON struct {
	Plan                  apijson.Field
	QuotasPerPlan         apijson.Field
	RemainingSchedules    apijson.Field
	RemainingTests        apijson.Field
	ScheduleQuotasPerPlan apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZoneSpeedAPIAvailabilitiesResponseResultQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region with a label.
type ZoneSpeedAPIAvailabilitiesResponseResultRegion struct {
	Label string `json:"label"`
	// A test region.
	Value ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue `json:"value"`
	JSON  zoneSpeedAPIAvailabilitiesResponseResultRegionJSON   `json:"-"`
}

// zoneSpeedAPIAvailabilitiesResponseResultRegionJSON contains the JSON metadata
// for the struct [ZoneSpeedAPIAvailabilitiesResponseResultRegion]
type zoneSpeedAPIAvailabilitiesResponseResultRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIAvailabilitiesResponseResultRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue string

const (
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueAsiaEast1           ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "asia-east1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueAsiaNortheast1      ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "asia-northeast1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueAsiaNortheast2      ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "asia-northeast2"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueAsiaSouth1          ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "asia-south1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueAsiaSoutheast1      ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "asia-southeast1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueAustraliaSoutheast1 ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "australia-southeast1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueEuropeNorth1        ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "europe-north1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueEuropeSouthwest1    ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "europe-southwest1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueEuropeWest1         ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "europe-west1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueEuropeWest2         ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "europe-west2"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueEuropeWest3         ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "europe-west3"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueEuropeWest4         ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "europe-west4"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueEuropeWest8         ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "europe-west8"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueEuropeWest9         ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "europe-west9"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueMeWest1             ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "me-west1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueSouthamericaEast1   ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "southamerica-east1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueUsCentral1          ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "us-central1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueUsEast1             ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "us-east1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueUsEast4             ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "us-east4"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueUsSouth1            ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "us-south1"
	ZoneSpeedAPIAvailabilitiesResponseResultRegionsValueUsWest1             ZoneSpeedAPIAvailabilitiesResponseResultRegionsValue = "us-west1"
)
