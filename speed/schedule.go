// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package speed

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

// ScheduleService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewScheduleService] method instead.
type ScheduleService struct {
	Options []option.RequestOption
}

// NewScheduleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScheduleService(opts ...option.RequestOption) (r *ScheduleService) {
	r = &ScheduleService{}
	r.Options = opts
	return
}

// Creates a scheduled test for a page.
func (r *ScheduleService) New(ctx context.Context, url string, params ScheduleNewParams, opts ...option.RequestOption) (res *ScheduleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScheduleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The test schedule.
type Schedule struct {
	// The frequency of the test.
	Frequency ScheduleFrequency `json:"frequency"`
	// A test region.
	Region ScheduleRegion `json:"region"`
	// A URL.
	URL  string       `json:"url"`
	JSON scheduleJSON `json:"-"`
}

// scheduleJSON contains the JSON metadata for the struct [Schedule]
type scheduleJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Schedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleJSON) RawJSON() string {
	return r.raw
}

// The frequency of the test.
type ScheduleFrequency string

const (
	ScheduleFrequencyDaily  ScheduleFrequency = "DAILY"
	ScheduleFrequencyWeekly ScheduleFrequency = "WEEKLY"
)

func (r ScheduleFrequency) IsKnown() bool {
	switch r {
	case ScheduleFrequencyDaily, ScheduleFrequencyWeekly:
		return true
	}
	return false
}

// A test region.
type ScheduleRegion string

const (
	ScheduleRegionAsiaEast1           ScheduleRegion = "asia-east1"
	ScheduleRegionAsiaNortheast1      ScheduleRegion = "asia-northeast1"
	ScheduleRegionAsiaNortheast2      ScheduleRegion = "asia-northeast2"
	ScheduleRegionAsiaSouth1          ScheduleRegion = "asia-south1"
	ScheduleRegionAsiaSoutheast1      ScheduleRegion = "asia-southeast1"
	ScheduleRegionAustraliaSoutheast1 ScheduleRegion = "australia-southeast1"
	ScheduleRegionEuropeNorth1        ScheduleRegion = "europe-north1"
	ScheduleRegionEuropeSouthwest1    ScheduleRegion = "europe-southwest1"
	ScheduleRegionEuropeWest1         ScheduleRegion = "europe-west1"
	ScheduleRegionEuropeWest2         ScheduleRegion = "europe-west2"
	ScheduleRegionEuropeWest3         ScheduleRegion = "europe-west3"
	ScheduleRegionEuropeWest4         ScheduleRegion = "europe-west4"
	ScheduleRegionEuropeWest8         ScheduleRegion = "europe-west8"
	ScheduleRegionEuropeWest9         ScheduleRegion = "europe-west9"
	ScheduleRegionMeWest1             ScheduleRegion = "me-west1"
	ScheduleRegionSouthamericaEast1   ScheduleRegion = "southamerica-east1"
	ScheduleRegionUsCentral1          ScheduleRegion = "us-central1"
	ScheduleRegionUsEast1             ScheduleRegion = "us-east1"
	ScheduleRegionUsEast4             ScheduleRegion = "us-east4"
	ScheduleRegionUsSouth1            ScheduleRegion = "us-south1"
	ScheduleRegionUsWest1             ScheduleRegion = "us-west1"
)

func (r ScheduleRegion) IsKnown() bool {
	switch r {
	case ScheduleRegionAsiaEast1, ScheduleRegionAsiaNortheast1, ScheduleRegionAsiaNortheast2, ScheduleRegionAsiaSouth1, ScheduleRegionAsiaSoutheast1, ScheduleRegionAustraliaSoutheast1, ScheduleRegionEuropeNorth1, ScheduleRegionEuropeSouthwest1, ScheduleRegionEuropeWest1, ScheduleRegionEuropeWest2, ScheduleRegionEuropeWest3, ScheduleRegionEuropeWest4, ScheduleRegionEuropeWest8, ScheduleRegionEuropeWest9, ScheduleRegionMeWest1, ScheduleRegionSouthamericaEast1, ScheduleRegionUsCentral1, ScheduleRegionUsEast1, ScheduleRegionUsEast4, ScheduleRegionUsSouth1, ScheduleRegionUsWest1:
		return true
	}
	return false
}

type ScheduleNewResponse struct {
	// The test schedule.
	Schedule Schedule                `json:"schedule"`
	Test     Test                    `json:"test"`
	JSON     scheduleNewResponseJSON `json:"-"`
}

// scheduleNewResponseJSON contains the JSON metadata for the struct
// [ScheduleNewResponse]
type scheduleNewResponseJSON struct {
	Schedule    apijson.Field
	Test        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScheduleNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[ScheduleNewParamsRegion] `query:"region"`
}

// URLQuery serializes [ScheduleNewParams]'s query parameters as `url.Values`.
func (r ScheduleNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type ScheduleNewParamsRegion string

const (
	ScheduleNewParamsRegionAsiaEast1           ScheduleNewParamsRegion = "asia-east1"
	ScheduleNewParamsRegionAsiaNortheast1      ScheduleNewParamsRegion = "asia-northeast1"
	ScheduleNewParamsRegionAsiaNortheast2      ScheduleNewParamsRegion = "asia-northeast2"
	ScheduleNewParamsRegionAsiaSouth1          ScheduleNewParamsRegion = "asia-south1"
	ScheduleNewParamsRegionAsiaSoutheast1      ScheduleNewParamsRegion = "asia-southeast1"
	ScheduleNewParamsRegionAustraliaSoutheast1 ScheduleNewParamsRegion = "australia-southeast1"
	ScheduleNewParamsRegionEuropeNorth1        ScheduleNewParamsRegion = "europe-north1"
	ScheduleNewParamsRegionEuropeSouthwest1    ScheduleNewParamsRegion = "europe-southwest1"
	ScheduleNewParamsRegionEuropeWest1         ScheduleNewParamsRegion = "europe-west1"
	ScheduleNewParamsRegionEuropeWest2         ScheduleNewParamsRegion = "europe-west2"
	ScheduleNewParamsRegionEuropeWest3         ScheduleNewParamsRegion = "europe-west3"
	ScheduleNewParamsRegionEuropeWest4         ScheduleNewParamsRegion = "europe-west4"
	ScheduleNewParamsRegionEuropeWest8         ScheduleNewParamsRegion = "europe-west8"
	ScheduleNewParamsRegionEuropeWest9         ScheduleNewParamsRegion = "europe-west9"
	ScheduleNewParamsRegionMeWest1             ScheduleNewParamsRegion = "me-west1"
	ScheduleNewParamsRegionSouthamericaEast1   ScheduleNewParamsRegion = "southamerica-east1"
	ScheduleNewParamsRegionUsCentral1          ScheduleNewParamsRegion = "us-central1"
	ScheduleNewParamsRegionUsEast1             ScheduleNewParamsRegion = "us-east1"
	ScheduleNewParamsRegionUsEast4             ScheduleNewParamsRegion = "us-east4"
	ScheduleNewParamsRegionUsSouth1            ScheduleNewParamsRegion = "us-south1"
	ScheduleNewParamsRegionUsWest1             ScheduleNewParamsRegion = "us-west1"
)

func (r ScheduleNewParamsRegion) IsKnown() bool {
	switch r {
	case ScheduleNewParamsRegionAsiaEast1, ScheduleNewParamsRegionAsiaNortheast1, ScheduleNewParamsRegionAsiaNortheast2, ScheduleNewParamsRegionAsiaSouth1, ScheduleNewParamsRegionAsiaSoutheast1, ScheduleNewParamsRegionAustraliaSoutheast1, ScheduleNewParamsRegionEuropeNorth1, ScheduleNewParamsRegionEuropeSouthwest1, ScheduleNewParamsRegionEuropeWest1, ScheduleNewParamsRegionEuropeWest2, ScheduleNewParamsRegionEuropeWest3, ScheduleNewParamsRegionEuropeWest4, ScheduleNewParamsRegionEuropeWest8, ScheduleNewParamsRegionEuropeWest9, ScheduleNewParamsRegionMeWest1, ScheduleNewParamsRegionSouthamericaEast1, ScheduleNewParamsRegionUsCentral1, ScheduleNewParamsRegionUsEast1, ScheduleNewParamsRegionUsEast4, ScheduleNewParamsRegionUsSouth1, ScheduleNewParamsRegionUsWest1:
		return true
	}
	return false
}

type ScheduleNewResponseEnvelope struct {
	Errors   interface{}                     `json:"errors,required"`
	Messages interface{}                     `json:"messages,required"`
	Success  interface{}                     `json:"success,required"`
	Result   ScheduleNewResponse             `json:"result"`
	JSON     scheduleNewResponseEnvelopeJSON `json:"-"`
}

// scheduleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScheduleNewResponseEnvelope]
type scheduleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
