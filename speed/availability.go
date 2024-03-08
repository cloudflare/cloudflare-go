// File generated from our OpenAPI spec by Stainless.

package speed

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *AvailabilityService) List(ctx context.Context, query AvailabilityListParams, opts ...option.RequestOption) (res *AvailabilityListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AvailabilityListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/availabilities", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AvailabilityListResponse struct {
	Quota          AvailabilityListResponseQuota    `json:"quota"`
	Regions        []AvailabilityListResponseRegion `json:"regions"`
	RegionsPerPlan interface{}                      `json:"regionsPerPlan"`
	JSON           availabilityListResponseJSON     `json:"-"`
}

// availabilityListResponseJSON contains the JSON metadata for the struct
// [AvailabilityListResponse]
type availabilityListResponseJSON struct {
	Quota          apijson.Field
	Regions        apijson.Field
	RegionsPerPlan apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AvailabilityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityListResponseJSON) RawJSON() string {
	return r.raw
}

type AvailabilityListResponseQuota struct {
	// Cloudflare plan.
	Plan string `json:"plan"`
	// The number of tests available per plan.
	QuotasPerPlan map[string]float64 `json:"quotasPerPlan"`
	// The number of remaining schedules available.
	RemainingSchedules float64 `json:"remainingSchedules"`
	// The number of remaining tests available.
	RemainingTests float64 `json:"remainingTests"`
	// The number of schedules available per plan.
	ScheduleQuotasPerPlan map[string]float64                `json:"scheduleQuotasPerPlan"`
	JSON                  availabilityListResponseQuotaJSON `json:"-"`
}

// availabilityListResponseQuotaJSON contains the JSON metadata for the struct
// [AvailabilityListResponseQuota]
type availabilityListResponseQuotaJSON struct {
	Plan                  apijson.Field
	QuotasPerPlan         apijson.Field
	RemainingSchedules    apijson.Field
	RemainingTests        apijson.Field
	ScheduleQuotasPerPlan apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AvailabilityListResponseQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityListResponseQuotaJSON) RawJSON() string {
	return r.raw
}

// A test region with a label.
type AvailabilityListResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value AvailabilityListResponseRegionsValue `json:"value"`
	JSON  availabilityListResponseRegionJSON   `json:"-"`
}

// availabilityListResponseRegionJSON contains the JSON metadata for the struct
// [AvailabilityListResponseRegion]
type availabilityListResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityListResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityListResponseRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type AvailabilityListResponseRegionsValue string

const (
	AvailabilityListResponseRegionsValueAsiaEast1           AvailabilityListResponseRegionsValue = "asia-east1"
	AvailabilityListResponseRegionsValueAsiaNortheast1      AvailabilityListResponseRegionsValue = "asia-northeast1"
	AvailabilityListResponseRegionsValueAsiaNortheast2      AvailabilityListResponseRegionsValue = "asia-northeast2"
	AvailabilityListResponseRegionsValueAsiaSouth1          AvailabilityListResponseRegionsValue = "asia-south1"
	AvailabilityListResponseRegionsValueAsiaSoutheast1      AvailabilityListResponseRegionsValue = "asia-southeast1"
	AvailabilityListResponseRegionsValueAustraliaSoutheast1 AvailabilityListResponseRegionsValue = "australia-southeast1"
	AvailabilityListResponseRegionsValueEuropeNorth1        AvailabilityListResponseRegionsValue = "europe-north1"
	AvailabilityListResponseRegionsValueEuropeSouthwest1    AvailabilityListResponseRegionsValue = "europe-southwest1"
	AvailabilityListResponseRegionsValueEuropeWest1         AvailabilityListResponseRegionsValue = "europe-west1"
	AvailabilityListResponseRegionsValueEuropeWest2         AvailabilityListResponseRegionsValue = "europe-west2"
	AvailabilityListResponseRegionsValueEuropeWest3         AvailabilityListResponseRegionsValue = "europe-west3"
	AvailabilityListResponseRegionsValueEuropeWest4         AvailabilityListResponseRegionsValue = "europe-west4"
	AvailabilityListResponseRegionsValueEuropeWest8         AvailabilityListResponseRegionsValue = "europe-west8"
	AvailabilityListResponseRegionsValueEuropeWest9         AvailabilityListResponseRegionsValue = "europe-west9"
	AvailabilityListResponseRegionsValueMeWest1             AvailabilityListResponseRegionsValue = "me-west1"
	AvailabilityListResponseRegionsValueSouthamericaEast1   AvailabilityListResponseRegionsValue = "southamerica-east1"
	AvailabilityListResponseRegionsValueUsCentral1          AvailabilityListResponseRegionsValue = "us-central1"
	AvailabilityListResponseRegionsValueUsEast1             AvailabilityListResponseRegionsValue = "us-east1"
	AvailabilityListResponseRegionsValueUsEast4             AvailabilityListResponseRegionsValue = "us-east4"
	AvailabilityListResponseRegionsValueUsSouth1            AvailabilityListResponseRegionsValue = "us-south1"
	AvailabilityListResponseRegionsValueUsWest1             AvailabilityListResponseRegionsValue = "us-west1"
)

type AvailabilityListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type AvailabilityListResponseEnvelope struct {
	Result AvailabilityListResponse             `json:"result"`
	JSON   availabilityListResponseEnvelopeJSON `json:"-"`
}

// availabilityListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AvailabilityListResponseEnvelope]
type availabilityListResponseEnvelopeJSON struct {
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
