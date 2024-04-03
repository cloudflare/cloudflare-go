// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package speed

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AvailabilityService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAvailabilityService] method
// instead.
type AvailabilityService struct {
	Options []option.RequestOption
}

// NewAvailabilityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAvailabilityService(opts ...option.RequestOption) (r *AvailabilityService) {
	r = &AvailabilityService{}
	r.Options = opts
	return
}

// Retrieves quota for all plans, as well as the current zone quota.
func (r *AvailabilityService) List(ctx context.Context, query AvailabilityListParams, opts ...option.RequestOption) (res *ObservatoryAvailabilities, err error) {
	opts = append(r.Options[:], opts...)
	var env AvailabilityListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/availabilities", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ObservatoryAvailabilities struct {
	Quota          ObservatoryAvailabilitiesQuota    `json:"quota"`
	Regions        []ObservatoryAvailabilitiesRegion `json:"regions"`
	RegionsPerPlan interface{}                       `json:"regionsPerPlan"`
	JSON           observatoryAvailabilitiesJSON     `json:"-"`
}

// observatoryAvailabilitiesJSON contains the JSON metadata for the struct
// [ObservatoryAvailabilities]
type observatoryAvailabilitiesJSON struct {
	Quota          apijson.Field
	Regions        apijson.Field
	RegionsPerPlan apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservatoryAvailabilities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryAvailabilitiesJSON) RawJSON() string {
	return r.raw
}

type ObservatoryAvailabilitiesQuota struct {
	// Cloudflare plan.
	Plan string `json:"plan"`
	// The number of tests available per plan.
	QuotasPerPlan map[string]float64 `json:"quotasPerPlan"`
	// The number of remaining schedules available.
	RemainingSchedules float64 `json:"remainingSchedules"`
	// The number of remaining tests available.
	RemainingTests float64 `json:"remainingTests"`
	// The number of schedules available per plan.
	ScheduleQuotasPerPlan map[string]float64                 `json:"scheduleQuotasPerPlan"`
	JSON                  observatoryAvailabilitiesQuotaJSON `json:"-"`
}

// observatoryAvailabilitiesQuotaJSON contains the JSON metadata for the struct
// [ObservatoryAvailabilitiesQuota]
type observatoryAvailabilitiesQuotaJSON struct {
	Plan                  apijson.Field
	QuotasPerPlan         apijson.Field
	RemainingSchedules    apijson.Field
	RemainingTests        apijson.Field
	ScheduleQuotasPerPlan apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ObservatoryAvailabilitiesQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryAvailabilitiesQuotaJSON) RawJSON() string {
	return r.raw
}

// A test region with a label.
type ObservatoryAvailabilitiesRegion struct {
	Label string `json:"label"`
	// A test region.
	Value ObservatoryAvailabilitiesRegionsValue `json:"value"`
	JSON  observatoryAvailabilitiesRegionJSON   `json:"-"`
}

// observatoryAvailabilitiesRegionJSON contains the JSON metadata for the struct
// [ObservatoryAvailabilitiesRegion]
type observatoryAvailabilitiesRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservatoryAvailabilitiesRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryAvailabilitiesRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type ObservatoryAvailabilitiesRegionsValue string

const (
	ObservatoryAvailabilitiesRegionsValueAsiaEast1           ObservatoryAvailabilitiesRegionsValue = "asia-east1"
	ObservatoryAvailabilitiesRegionsValueAsiaNortheast1      ObservatoryAvailabilitiesRegionsValue = "asia-northeast1"
	ObservatoryAvailabilitiesRegionsValueAsiaNortheast2      ObservatoryAvailabilitiesRegionsValue = "asia-northeast2"
	ObservatoryAvailabilitiesRegionsValueAsiaSouth1          ObservatoryAvailabilitiesRegionsValue = "asia-south1"
	ObservatoryAvailabilitiesRegionsValueAsiaSoutheast1      ObservatoryAvailabilitiesRegionsValue = "asia-southeast1"
	ObservatoryAvailabilitiesRegionsValueAustraliaSoutheast1 ObservatoryAvailabilitiesRegionsValue = "australia-southeast1"
	ObservatoryAvailabilitiesRegionsValueEuropeNorth1        ObservatoryAvailabilitiesRegionsValue = "europe-north1"
	ObservatoryAvailabilitiesRegionsValueEuropeSouthwest1    ObservatoryAvailabilitiesRegionsValue = "europe-southwest1"
	ObservatoryAvailabilitiesRegionsValueEuropeWest1         ObservatoryAvailabilitiesRegionsValue = "europe-west1"
	ObservatoryAvailabilitiesRegionsValueEuropeWest2         ObservatoryAvailabilitiesRegionsValue = "europe-west2"
	ObservatoryAvailabilitiesRegionsValueEuropeWest3         ObservatoryAvailabilitiesRegionsValue = "europe-west3"
	ObservatoryAvailabilitiesRegionsValueEuropeWest4         ObservatoryAvailabilitiesRegionsValue = "europe-west4"
	ObservatoryAvailabilitiesRegionsValueEuropeWest8         ObservatoryAvailabilitiesRegionsValue = "europe-west8"
	ObservatoryAvailabilitiesRegionsValueEuropeWest9         ObservatoryAvailabilitiesRegionsValue = "europe-west9"
	ObservatoryAvailabilitiesRegionsValueMeWest1             ObservatoryAvailabilitiesRegionsValue = "me-west1"
	ObservatoryAvailabilitiesRegionsValueSouthamericaEast1   ObservatoryAvailabilitiesRegionsValue = "southamerica-east1"
	ObservatoryAvailabilitiesRegionsValueUsCentral1          ObservatoryAvailabilitiesRegionsValue = "us-central1"
	ObservatoryAvailabilitiesRegionsValueUsEast1             ObservatoryAvailabilitiesRegionsValue = "us-east1"
	ObservatoryAvailabilitiesRegionsValueUsEast4             ObservatoryAvailabilitiesRegionsValue = "us-east4"
	ObservatoryAvailabilitiesRegionsValueUsSouth1            ObservatoryAvailabilitiesRegionsValue = "us-south1"
	ObservatoryAvailabilitiesRegionsValueUsWest1             ObservatoryAvailabilitiesRegionsValue = "us-west1"
)

func (r ObservatoryAvailabilitiesRegionsValue) IsKnown() bool {
	switch r {
	case ObservatoryAvailabilitiesRegionsValueAsiaEast1, ObservatoryAvailabilitiesRegionsValueAsiaNortheast1, ObservatoryAvailabilitiesRegionsValueAsiaNortheast2, ObservatoryAvailabilitiesRegionsValueAsiaSouth1, ObservatoryAvailabilitiesRegionsValueAsiaSoutheast1, ObservatoryAvailabilitiesRegionsValueAustraliaSoutheast1, ObservatoryAvailabilitiesRegionsValueEuropeNorth1, ObservatoryAvailabilitiesRegionsValueEuropeSouthwest1, ObservatoryAvailabilitiesRegionsValueEuropeWest1, ObservatoryAvailabilitiesRegionsValueEuropeWest2, ObservatoryAvailabilitiesRegionsValueEuropeWest3, ObservatoryAvailabilitiesRegionsValueEuropeWest4, ObservatoryAvailabilitiesRegionsValueEuropeWest8, ObservatoryAvailabilitiesRegionsValueEuropeWest9, ObservatoryAvailabilitiesRegionsValueMeWest1, ObservatoryAvailabilitiesRegionsValueSouthamericaEast1, ObservatoryAvailabilitiesRegionsValueUsCentral1, ObservatoryAvailabilitiesRegionsValueUsEast1, ObservatoryAvailabilitiesRegionsValueUsEast4, ObservatoryAvailabilitiesRegionsValueUsSouth1, ObservatoryAvailabilitiesRegionsValueUsWest1:
		return true
	}
	return false
}

type AvailabilityListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type AvailabilityListResponseEnvelope struct {
	Errors   interface{}                          `json:"errors,required"`
	Messages interface{}                          `json:"messages,required"`
	Success  interface{}                          `json:"success,required"`
	Result   ObservatoryAvailabilities            `json:"result"`
	JSON     availabilityListResponseEnvelopeJSON `json:"-"`
}

// availabilityListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AvailabilityListResponseEnvelope]
type availabilityListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
