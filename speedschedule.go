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

// SpeedScheduleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSpeedScheduleService] method
// instead.
type SpeedScheduleService struct {
	Options []option.RequestOption
}

// NewSpeedScheduleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpeedScheduleService(opts ...option.RequestOption) (r *SpeedScheduleService) {
	r = &SpeedScheduleService{}
	r.Options = opts
	return
}

// Creates a scheduled test for a page.
func (r *SpeedScheduleService) New(ctx context.Context, url string, params SpeedScheduleNewParams, opts ...option.RequestOption) (res *SpeedScheduleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedScheduleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SpeedScheduleNewResponse struct {
	// The test schedule.
	Schedule ObservatorySchedule          `json:"schedule"`
	Test     ObservatoryPageTest          `json:"test"`
	JSON     speedScheduleNewResponseJSON `json:"-"`
}

// speedScheduleNewResponseJSON contains the JSON metadata for the struct
// [SpeedScheduleNewResponse]
type speedScheduleNewResponseJSON struct {
	Schedule    apijson.Field
	Test        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedScheduleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedScheduleNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[SpeedScheduleNewParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedScheduleNewParams]'s query parameters as `url.Values`.
func (r SpeedScheduleNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedScheduleNewParamsRegion string

const (
	SpeedScheduleNewParamsRegionAsiaEast1           SpeedScheduleNewParamsRegion = "asia-east1"
	SpeedScheduleNewParamsRegionAsiaNortheast1      SpeedScheduleNewParamsRegion = "asia-northeast1"
	SpeedScheduleNewParamsRegionAsiaNortheast2      SpeedScheduleNewParamsRegion = "asia-northeast2"
	SpeedScheduleNewParamsRegionAsiaSouth1          SpeedScheduleNewParamsRegion = "asia-south1"
	SpeedScheduleNewParamsRegionAsiaSoutheast1      SpeedScheduleNewParamsRegion = "asia-southeast1"
	SpeedScheduleNewParamsRegionAustraliaSoutheast1 SpeedScheduleNewParamsRegion = "australia-southeast1"
	SpeedScheduleNewParamsRegionEuropeNorth1        SpeedScheduleNewParamsRegion = "europe-north1"
	SpeedScheduleNewParamsRegionEuropeSouthwest1    SpeedScheduleNewParamsRegion = "europe-southwest1"
	SpeedScheduleNewParamsRegionEuropeWest1         SpeedScheduleNewParamsRegion = "europe-west1"
	SpeedScheduleNewParamsRegionEuropeWest2         SpeedScheduleNewParamsRegion = "europe-west2"
	SpeedScheduleNewParamsRegionEuropeWest3         SpeedScheduleNewParamsRegion = "europe-west3"
	SpeedScheduleNewParamsRegionEuropeWest4         SpeedScheduleNewParamsRegion = "europe-west4"
	SpeedScheduleNewParamsRegionEuropeWest8         SpeedScheduleNewParamsRegion = "europe-west8"
	SpeedScheduleNewParamsRegionEuropeWest9         SpeedScheduleNewParamsRegion = "europe-west9"
	SpeedScheduleNewParamsRegionMeWest1             SpeedScheduleNewParamsRegion = "me-west1"
	SpeedScheduleNewParamsRegionSouthamericaEast1   SpeedScheduleNewParamsRegion = "southamerica-east1"
	SpeedScheduleNewParamsRegionUsCentral1          SpeedScheduleNewParamsRegion = "us-central1"
	SpeedScheduleNewParamsRegionUsEast1             SpeedScheduleNewParamsRegion = "us-east1"
	SpeedScheduleNewParamsRegionUsEast4             SpeedScheduleNewParamsRegion = "us-east4"
	SpeedScheduleNewParamsRegionUsSouth1            SpeedScheduleNewParamsRegion = "us-south1"
	SpeedScheduleNewParamsRegionUsWest1             SpeedScheduleNewParamsRegion = "us-west1"
)

type SpeedScheduleNewResponseEnvelope struct {
	Result SpeedScheduleNewResponse             `json:"result"`
	JSON   speedScheduleNewResponseEnvelopeJSON `json:"-"`
}

// speedScheduleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedScheduleNewResponseEnvelope]
type speedScheduleNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedScheduleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
