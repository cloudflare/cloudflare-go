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

// SpeedPageService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSpeedPageService] method instead.
type SpeedPageService struct {
	Options []option.RequestOption
}

// NewSpeedPageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpeedPageService(opts ...option.RequestOption) (r *SpeedPageService) {
	r = &SpeedPageService{}
	r.Options = opts
	return
}

// Lists all webpages which have been tested.
func (r *SpeedPageService) List(ctx context.Context, query SpeedPageListParams, opts ...option.RequestOption) (res *[]SpeedPageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedPageListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SpeedPageListResponse struct {
	// A test region with a label.
	Region SpeedPageListResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedPageListResponseScheduleFrequency `json:"scheduleFrequency"`
	Tests             []ObservatoryPageTest                  `json:"tests"`
	// A URL.
	URL  string                    `json:"url"`
	JSON speedPageListResponseJSON `json:"-"`
}

// speedPageListResponseJSON contains the JSON metadata for the struct
// [SpeedPageListResponse]
type speedPageListResponseJSON struct {
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	Tests             apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedPageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region with a label.
type SpeedPageListResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedPageListResponseRegionValue `json:"value"`
	JSON  speedPageListResponseRegionJSON  `json:"-"`
}

// speedPageListResponseRegionJSON contains the JSON metadata for the struct
// [SpeedPageListResponseRegion]
type speedPageListResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedPageListResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedPageListResponseRegionValue string

const (
	SpeedPageListResponseRegionValueAsiaEast1           SpeedPageListResponseRegionValue = "asia-east1"
	SpeedPageListResponseRegionValueAsiaNortheast1      SpeedPageListResponseRegionValue = "asia-northeast1"
	SpeedPageListResponseRegionValueAsiaNortheast2      SpeedPageListResponseRegionValue = "asia-northeast2"
	SpeedPageListResponseRegionValueAsiaSouth1          SpeedPageListResponseRegionValue = "asia-south1"
	SpeedPageListResponseRegionValueAsiaSoutheast1      SpeedPageListResponseRegionValue = "asia-southeast1"
	SpeedPageListResponseRegionValueAustraliaSoutheast1 SpeedPageListResponseRegionValue = "australia-southeast1"
	SpeedPageListResponseRegionValueEuropeNorth1        SpeedPageListResponseRegionValue = "europe-north1"
	SpeedPageListResponseRegionValueEuropeSouthwest1    SpeedPageListResponseRegionValue = "europe-southwest1"
	SpeedPageListResponseRegionValueEuropeWest1         SpeedPageListResponseRegionValue = "europe-west1"
	SpeedPageListResponseRegionValueEuropeWest2         SpeedPageListResponseRegionValue = "europe-west2"
	SpeedPageListResponseRegionValueEuropeWest3         SpeedPageListResponseRegionValue = "europe-west3"
	SpeedPageListResponseRegionValueEuropeWest4         SpeedPageListResponseRegionValue = "europe-west4"
	SpeedPageListResponseRegionValueEuropeWest8         SpeedPageListResponseRegionValue = "europe-west8"
	SpeedPageListResponseRegionValueEuropeWest9         SpeedPageListResponseRegionValue = "europe-west9"
	SpeedPageListResponseRegionValueMeWest1             SpeedPageListResponseRegionValue = "me-west1"
	SpeedPageListResponseRegionValueSouthamericaEast1   SpeedPageListResponseRegionValue = "southamerica-east1"
	SpeedPageListResponseRegionValueUsCentral1          SpeedPageListResponseRegionValue = "us-central1"
	SpeedPageListResponseRegionValueUsEast1             SpeedPageListResponseRegionValue = "us-east1"
	SpeedPageListResponseRegionValueUsEast4             SpeedPageListResponseRegionValue = "us-east4"
	SpeedPageListResponseRegionValueUsSouth1            SpeedPageListResponseRegionValue = "us-south1"
	SpeedPageListResponseRegionValueUsWest1             SpeedPageListResponseRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedPageListResponseScheduleFrequency string

const (
	SpeedPageListResponseScheduleFrequencyDaily  SpeedPageListResponseScheduleFrequency = "DAILY"
	SpeedPageListResponseScheduleFrequencyWeekly SpeedPageListResponseScheduleFrequency = "WEEKLY"
)

type SpeedPageListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SpeedPageListResponseEnvelope struct {
	Result []SpeedPageListResponse           `json:"result"`
	JSON   speedPageListResponseEnvelopeJSON `json:"-"`
}

// speedPageListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedPageListResponseEnvelope]
type speedPageListResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedPageListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
