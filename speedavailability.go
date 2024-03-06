// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SpeedAvailabilityService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSpeedAvailabilityService] method
// instead.
type SpeedAvailabilityService struct {
	Options []option.RequestOption
}

// NewSpeedAvailabilityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSpeedAvailabilityService(opts ...option.RequestOption) (r *SpeedAvailabilityService) {
	r = &SpeedAvailabilityService{}
	r.Options = opts
	return
}

// Retrieves quota for all plans, as well as the current zone quota.
func (r *SpeedAvailabilityService) List(ctx context.Context, query SpeedAvailabilityListParams, opts ...option.RequestOption) (res *SpeedAvailabilityListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedAvailabilityListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/availabilities", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SpeedAvailabilityListResponse struct {
	Quota          SpeedAvailabilityListResponseQuota    `json:"quota"`
	Regions        []SpeedAvailabilityListResponseRegion `json:"regions"`
	RegionsPerPlan interface{}                           `json:"regionsPerPlan"`
	JSON           speedAvailabilityListResponseJSON     `json:"-"`
}

// speedAvailabilityListResponseJSON contains the JSON metadata for the struct
// [SpeedAvailabilityListResponse]
type speedAvailabilityListResponseJSON struct {
	Quota          apijson.Field
	Regions        apijson.Field
	RegionsPerPlan apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SpeedAvailabilityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedAvailabilityListResponseJSON) RawJSON() string {
	return r.raw
}

type SpeedAvailabilityListResponseQuota struct {
	// Cloudflare plan.
	Plan string `json:"plan"`
	// The number of tests available per plan.
	QuotasPerPlan map[string]float64 `json:"quotasPerPlan"`
	// The number of remaining schedules available.
	RemainingSchedules float64 `json:"remainingSchedules"`
	// The number of remaining tests available.
	RemainingTests float64 `json:"remainingTests"`
	// The number of schedules available per plan.
	ScheduleQuotasPerPlan map[string]float64                     `json:"scheduleQuotasPerPlan"`
	JSON                  speedAvailabilityListResponseQuotaJSON `json:"-"`
}

// speedAvailabilityListResponseQuotaJSON contains the JSON metadata for the struct
// [SpeedAvailabilityListResponseQuota]
type speedAvailabilityListResponseQuotaJSON struct {
	Plan                  apijson.Field
	QuotasPerPlan         apijson.Field
	RemainingSchedules    apijson.Field
	RemainingTests        apijson.Field
	ScheduleQuotasPerPlan apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SpeedAvailabilityListResponseQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedAvailabilityListResponseQuotaJSON) RawJSON() string {
	return r.raw
}

// A test region with a label.
type SpeedAvailabilityListResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAvailabilityListResponseRegionsValue `json:"value"`
	JSON  speedAvailabilityListResponseRegionJSON   `json:"-"`
}

// speedAvailabilityListResponseRegionJSON contains the JSON metadata for the
// struct [SpeedAvailabilityListResponseRegion]
type speedAvailabilityListResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAvailabilityListResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedAvailabilityListResponseRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type SpeedAvailabilityListResponseRegionsValue string

const (
	SpeedAvailabilityListResponseRegionsValueAsiaEast1           SpeedAvailabilityListResponseRegionsValue = "asia-east1"
	SpeedAvailabilityListResponseRegionsValueAsiaNortheast1      SpeedAvailabilityListResponseRegionsValue = "asia-northeast1"
	SpeedAvailabilityListResponseRegionsValueAsiaNortheast2      SpeedAvailabilityListResponseRegionsValue = "asia-northeast2"
	SpeedAvailabilityListResponseRegionsValueAsiaSouth1          SpeedAvailabilityListResponseRegionsValue = "asia-south1"
	SpeedAvailabilityListResponseRegionsValueAsiaSoutheast1      SpeedAvailabilityListResponseRegionsValue = "asia-southeast1"
	SpeedAvailabilityListResponseRegionsValueAustraliaSoutheast1 SpeedAvailabilityListResponseRegionsValue = "australia-southeast1"
	SpeedAvailabilityListResponseRegionsValueEuropeNorth1        SpeedAvailabilityListResponseRegionsValue = "europe-north1"
	SpeedAvailabilityListResponseRegionsValueEuropeSouthwest1    SpeedAvailabilityListResponseRegionsValue = "europe-southwest1"
	SpeedAvailabilityListResponseRegionsValueEuropeWest1         SpeedAvailabilityListResponseRegionsValue = "europe-west1"
	SpeedAvailabilityListResponseRegionsValueEuropeWest2         SpeedAvailabilityListResponseRegionsValue = "europe-west2"
	SpeedAvailabilityListResponseRegionsValueEuropeWest3         SpeedAvailabilityListResponseRegionsValue = "europe-west3"
	SpeedAvailabilityListResponseRegionsValueEuropeWest4         SpeedAvailabilityListResponseRegionsValue = "europe-west4"
	SpeedAvailabilityListResponseRegionsValueEuropeWest8         SpeedAvailabilityListResponseRegionsValue = "europe-west8"
	SpeedAvailabilityListResponseRegionsValueEuropeWest9         SpeedAvailabilityListResponseRegionsValue = "europe-west9"
	SpeedAvailabilityListResponseRegionsValueMeWest1             SpeedAvailabilityListResponseRegionsValue = "me-west1"
	SpeedAvailabilityListResponseRegionsValueSouthamericaEast1   SpeedAvailabilityListResponseRegionsValue = "southamerica-east1"
	SpeedAvailabilityListResponseRegionsValueUsCentral1          SpeedAvailabilityListResponseRegionsValue = "us-central1"
	SpeedAvailabilityListResponseRegionsValueUsEast1             SpeedAvailabilityListResponseRegionsValue = "us-east1"
	SpeedAvailabilityListResponseRegionsValueUsEast4             SpeedAvailabilityListResponseRegionsValue = "us-east4"
	SpeedAvailabilityListResponseRegionsValueUsSouth1            SpeedAvailabilityListResponseRegionsValue = "us-south1"
	SpeedAvailabilityListResponseRegionsValueUsWest1             SpeedAvailabilityListResponseRegionsValue = "us-west1"
)

type SpeedAvailabilityListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SpeedAvailabilityListResponseEnvelope struct {
	Result SpeedAvailabilityListResponse             `json:"result"`
	JSON   speedAvailabilityListResponseEnvelopeJSON `json:"-"`
}

// speedAvailabilityListResponseEnvelopeJSON contains the JSON metadata for the
// struct [SpeedAvailabilityListResponseEnvelope]
type speedAvailabilityListResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAvailabilityListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedAvailabilityListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
