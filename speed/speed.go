// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package speed

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SpeedService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewSpeedService] method instead.
type SpeedService struct {
	Options        []option.RequestOption
	Tests          *TestService
	Schedule       *ScheduleService
	Availabilities *AvailabilityService
	Pages          *PageService
}

// NewSpeedService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSpeedService(opts ...option.RequestOption) (r *SpeedService) {
	r = &SpeedService{}
	r.Options = opts
	r.Tests = NewTestService(opts...)
	r.Schedule = NewScheduleService(opts...)
	r.Availabilities = NewAvailabilityService(opts...)
	r.Pages = NewPageService(opts...)
	return
}

// Deletes a scheduled test for a page.
func (r *SpeedService) Delete(ctx context.Context, url string, params SpeedDeleteParams, opts ...option.RequestOption) (res *SpeedDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the test schedule for a page in a specific region.
func (r *SpeedService) ScheduleGet(ctx context.Context, url string, params SpeedScheduleGetParams, opts ...option.RequestOption) (res *ObservatorySchedule, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedScheduleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the core web vital metrics trend over time for a specific page.
func (r *SpeedService) TrendsList(ctx context.Context, url string, params SpeedTrendsListParams, opts ...option.RequestOption) (res *ObservatoryTrend, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedTrendsListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/trend", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The test schedule.
type ObservatorySchedule struct {
	// The frequency of the test.
	Frequency ObservatoryScheduleFrequency `json:"frequency"`
	// A test region.
	Region ObservatoryScheduleRegion `json:"region"`
	// A URL.
	URL  string                  `json:"url"`
	JSON observatoryScheduleJSON `json:"-"`
}

// observatoryScheduleJSON contains the JSON metadata for the struct
// [ObservatorySchedule]
type observatoryScheduleJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservatorySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryScheduleJSON) RawJSON() string {
	return r.raw
}

// The frequency of the test.
type ObservatoryScheduleFrequency string

const (
	ObservatoryScheduleFrequencyDaily  ObservatoryScheduleFrequency = "DAILY"
	ObservatoryScheduleFrequencyWeekly ObservatoryScheduleFrequency = "WEEKLY"
)

func (r ObservatoryScheduleFrequency) IsKnown() bool {
	switch r {
	case ObservatoryScheduleFrequencyDaily, ObservatoryScheduleFrequencyWeekly:
		return true
	}
	return false
}

// A test region.
type ObservatoryScheduleRegion string

const (
	ObservatoryScheduleRegionAsiaEast1           ObservatoryScheduleRegion = "asia-east1"
	ObservatoryScheduleRegionAsiaNortheast1      ObservatoryScheduleRegion = "asia-northeast1"
	ObservatoryScheduleRegionAsiaNortheast2      ObservatoryScheduleRegion = "asia-northeast2"
	ObservatoryScheduleRegionAsiaSouth1          ObservatoryScheduleRegion = "asia-south1"
	ObservatoryScheduleRegionAsiaSoutheast1      ObservatoryScheduleRegion = "asia-southeast1"
	ObservatoryScheduleRegionAustraliaSoutheast1 ObservatoryScheduleRegion = "australia-southeast1"
	ObservatoryScheduleRegionEuropeNorth1        ObservatoryScheduleRegion = "europe-north1"
	ObservatoryScheduleRegionEuropeSouthwest1    ObservatoryScheduleRegion = "europe-southwest1"
	ObservatoryScheduleRegionEuropeWest1         ObservatoryScheduleRegion = "europe-west1"
	ObservatoryScheduleRegionEuropeWest2         ObservatoryScheduleRegion = "europe-west2"
	ObservatoryScheduleRegionEuropeWest3         ObservatoryScheduleRegion = "europe-west3"
	ObservatoryScheduleRegionEuropeWest4         ObservatoryScheduleRegion = "europe-west4"
	ObservatoryScheduleRegionEuropeWest8         ObservatoryScheduleRegion = "europe-west8"
	ObservatoryScheduleRegionEuropeWest9         ObservatoryScheduleRegion = "europe-west9"
	ObservatoryScheduleRegionMeWest1             ObservatoryScheduleRegion = "me-west1"
	ObservatoryScheduleRegionSouthamericaEast1   ObservatoryScheduleRegion = "southamerica-east1"
	ObservatoryScheduleRegionUsCentral1          ObservatoryScheduleRegion = "us-central1"
	ObservatoryScheduleRegionUsEast1             ObservatoryScheduleRegion = "us-east1"
	ObservatoryScheduleRegionUsEast4             ObservatoryScheduleRegion = "us-east4"
	ObservatoryScheduleRegionUsSouth1            ObservatoryScheduleRegion = "us-south1"
	ObservatoryScheduleRegionUsWest1             ObservatoryScheduleRegion = "us-west1"
)

func (r ObservatoryScheduleRegion) IsKnown() bool {
	switch r {
	case ObservatoryScheduleRegionAsiaEast1, ObservatoryScheduleRegionAsiaNortheast1, ObservatoryScheduleRegionAsiaNortheast2, ObservatoryScheduleRegionAsiaSouth1, ObservatoryScheduleRegionAsiaSoutheast1, ObservatoryScheduleRegionAustraliaSoutheast1, ObservatoryScheduleRegionEuropeNorth1, ObservatoryScheduleRegionEuropeSouthwest1, ObservatoryScheduleRegionEuropeWest1, ObservatoryScheduleRegionEuropeWest2, ObservatoryScheduleRegionEuropeWest3, ObservatoryScheduleRegionEuropeWest4, ObservatoryScheduleRegionEuropeWest8, ObservatoryScheduleRegionEuropeWest9, ObservatoryScheduleRegionMeWest1, ObservatoryScheduleRegionSouthamericaEast1, ObservatoryScheduleRegionUsCentral1, ObservatoryScheduleRegionUsEast1, ObservatoryScheduleRegionUsEast4, ObservatoryScheduleRegionUsSouth1, ObservatoryScheduleRegionUsWest1:
		return true
	}
	return false
}

type ObservatoryTrend struct {
	// Cumulative Layout Shift trend.
	Cls []float64 `json:"cls"`
	// First Contentful Paint trend.
	Fcp []float64 `json:"fcp"`
	// Largest Contentful Paint trend.
	Lcp []float64 `json:"lcp"`
	// The Lighthouse score trend.
	PerformanceScore []float64 `json:"performanceScore"`
	// Speed Index trend.
	Si []float64 `json:"si"`
	// Total Blocking Time trend.
	Tbt []float64 `json:"tbt"`
	// Time To First Byte trend.
	Ttfb []float64 `json:"ttfb"`
	// Time To Interactive trend.
	Tti  []float64            `json:"tti"`
	JSON observatoryTrendJSON `json:"-"`
}

// observatoryTrendJSON contains the JSON metadata for the struct
// [ObservatoryTrend]
type observatoryTrendJSON struct {
	Cls              apijson.Field
	Fcp              apijson.Field
	Lcp              apijson.Field
	PerformanceScore apijson.Field
	Si               apijson.Field
	Tbt              apijson.Field
	Ttfb             apijson.Field
	Tti              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ObservatoryTrend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryTrendJSON) RawJSON() string {
	return r.raw
}

type SpeedDeleteResponse struct {
	// Number of items affected.
	Count float64                 `json:"count"`
	JSON  speedDeleteResponseJSON `json:"-"`
}

// speedDeleteResponseJSON contains the JSON metadata for the struct
// [SpeedDeleteResponse]
type speedDeleteResponseJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SpeedDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[SpeedDeleteParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedDeleteParams]'s query parameters as `url.Values`.
func (r SpeedDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedDeleteParamsRegion string

const (
	SpeedDeleteParamsRegionAsiaEast1           SpeedDeleteParamsRegion = "asia-east1"
	SpeedDeleteParamsRegionAsiaNortheast1      SpeedDeleteParamsRegion = "asia-northeast1"
	SpeedDeleteParamsRegionAsiaNortheast2      SpeedDeleteParamsRegion = "asia-northeast2"
	SpeedDeleteParamsRegionAsiaSouth1          SpeedDeleteParamsRegion = "asia-south1"
	SpeedDeleteParamsRegionAsiaSoutheast1      SpeedDeleteParamsRegion = "asia-southeast1"
	SpeedDeleteParamsRegionAustraliaSoutheast1 SpeedDeleteParamsRegion = "australia-southeast1"
	SpeedDeleteParamsRegionEuropeNorth1        SpeedDeleteParamsRegion = "europe-north1"
	SpeedDeleteParamsRegionEuropeSouthwest1    SpeedDeleteParamsRegion = "europe-southwest1"
	SpeedDeleteParamsRegionEuropeWest1         SpeedDeleteParamsRegion = "europe-west1"
	SpeedDeleteParamsRegionEuropeWest2         SpeedDeleteParamsRegion = "europe-west2"
	SpeedDeleteParamsRegionEuropeWest3         SpeedDeleteParamsRegion = "europe-west3"
	SpeedDeleteParamsRegionEuropeWest4         SpeedDeleteParamsRegion = "europe-west4"
	SpeedDeleteParamsRegionEuropeWest8         SpeedDeleteParamsRegion = "europe-west8"
	SpeedDeleteParamsRegionEuropeWest9         SpeedDeleteParamsRegion = "europe-west9"
	SpeedDeleteParamsRegionMeWest1             SpeedDeleteParamsRegion = "me-west1"
	SpeedDeleteParamsRegionSouthamericaEast1   SpeedDeleteParamsRegion = "southamerica-east1"
	SpeedDeleteParamsRegionUsCentral1          SpeedDeleteParamsRegion = "us-central1"
	SpeedDeleteParamsRegionUsEast1             SpeedDeleteParamsRegion = "us-east1"
	SpeedDeleteParamsRegionUsEast4             SpeedDeleteParamsRegion = "us-east4"
	SpeedDeleteParamsRegionUsSouth1            SpeedDeleteParamsRegion = "us-south1"
	SpeedDeleteParamsRegionUsWest1             SpeedDeleteParamsRegion = "us-west1"
)

func (r SpeedDeleteParamsRegion) IsKnown() bool {
	switch r {
	case SpeedDeleteParamsRegionAsiaEast1, SpeedDeleteParamsRegionAsiaNortheast1, SpeedDeleteParamsRegionAsiaNortheast2, SpeedDeleteParamsRegionAsiaSouth1, SpeedDeleteParamsRegionAsiaSoutheast1, SpeedDeleteParamsRegionAustraliaSoutheast1, SpeedDeleteParamsRegionEuropeNorth1, SpeedDeleteParamsRegionEuropeSouthwest1, SpeedDeleteParamsRegionEuropeWest1, SpeedDeleteParamsRegionEuropeWest2, SpeedDeleteParamsRegionEuropeWest3, SpeedDeleteParamsRegionEuropeWest4, SpeedDeleteParamsRegionEuropeWest8, SpeedDeleteParamsRegionEuropeWest9, SpeedDeleteParamsRegionMeWest1, SpeedDeleteParamsRegionSouthamericaEast1, SpeedDeleteParamsRegionUsCentral1, SpeedDeleteParamsRegionUsEast1, SpeedDeleteParamsRegionUsEast4, SpeedDeleteParamsRegionUsSouth1, SpeedDeleteParamsRegionUsWest1:
		return true
	}
	return false
}

type SpeedDeleteResponseEnvelope struct {
	Errors   interface{}                     `json:"errors,required"`
	Messages interface{}                     `json:"messages,required"`
	Success  interface{}                     `json:"success,required"`
	Result   SpeedDeleteResponse             `json:"result"`
	JSON     speedDeleteResponseEnvelopeJSON `json:"-"`
}

// speedDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedDeleteResponseEnvelope]
type speedDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SpeedScheduleGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[SpeedScheduleGetParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedScheduleGetParams]'s query parameters as `url.Values`.
func (r SpeedScheduleGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedScheduleGetParamsRegion string

const (
	SpeedScheduleGetParamsRegionAsiaEast1           SpeedScheduleGetParamsRegion = "asia-east1"
	SpeedScheduleGetParamsRegionAsiaNortheast1      SpeedScheduleGetParamsRegion = "asia-northeast1"
	SpeedScheduleGetParamsRegionAsiaNortheast2      SpeedScheduleGetParamsRegion = "asia-northeast2"
	SpeedScheduleGetParamsRegionAsiaSouth1          SpeedScheduleGetParamsRegion = "asia-south1"
	SpeedScheduleGetParamsRegionAsiaSoutheast1      SpeedScheduleGetParamsRegion = "asia-southeast1"
	SpeedScheduleGetParamsRegionAustraliaSoutheast1 SpeedScheduleGetParamsRegion = "australia-southeast1"
	SpeedScheduleGetParamsRegionEuropeNorth1        SpeedScheduleGetParamsRegion = "europe-north1"
	SpeedScheduleGetParamsRegionEuropeSouthwest1    SpeedScheduleGetParamsRegion = "europe-southwest1"
	SpeedScheduleGetParamsRegionEuropeWest1         SpeedScheduleGetParamsRegion = "europe-west1"
	SpeedScheduleGetParamsRegionEuropeWest2         SpeedScheduleGetParamsRegion = "europe-west2"
	SpeedScheduleGetParamsRegionEuropeWest3         SpeedScheduleGetParamsRegion = "europe-west3"
	SpeedScheduleGetParamsRegionEuropeWest4         SpeedScheduleGetParamsRegion = "europe-west4"
	SpeedScheduleGetParamsRegionEuropeWest8         SpeedScheduleGetParamsRegion = "europe-west8"
	SpeedScheduleGetParamsRegionEuropeWest9         SpeedScheduleGetParamsRegion = "europe-west9"
	SpeedScheduleGetParamsRegionMeWest1             SpeedScheduleGetParamsRegion = "me-west1"
	SpeedScheduleGetParamsRegionSouthamericaEast1   SpeedScheduleGetParamsRegion = "southamerica-east1"
	SpeedScheduleGetParamsRegionUsCentral1          SpeedScheduleGetParamsRegion = "us-central1"
	SpeedScheduleGetParamsRegionUsEast1             SpeedScheduleGetParamsRegion = "us-east1"
	SpeedScheduleGetParamsRegionUsEast4             SpeedScheduleGetParamsRegion = "us-east4"
	SpeedScheduleGetParamsRegionUsSouth1            SpeedScheduleGetParamsRegion = "us-south1"
	SpeedScheduleGetParamsRegionUsWest1             SpeedScheduleGetParamsRegion = "us-west1"
)

func (r SpeedScheduleGetParamsRegion) IsKnown() bool {
	switch r {
	case SpeedScheduleGetParamsRegionAsiaEast1, SpeedScheduleGetParamsRegionAsiaNortheast1, SpeedScheduleGetParamsRegionAsiaNortheast2, SpeedScheduleGetParamsRegionAsiaSouth1, SpeedScheduleGetParamsRegionAsiaSoutheast1, SpeedScheduleGetParamsRegionAustraliaSoutheast1, SpeedScheduleGetParamsRegionEuropeNorth1, SpeedScheduleGetParamsRegionEuropeSouthwest1, SpeedScheduleGetParamsRegionEuropeWest1, SpeedScheduleGetParamsRegionEuropeWest2, SpeedScheduleGetParamsRegionEuropeWest3, SpeedScheduleGetParamsRegionEuropeWest4, SpeedScheduleGetParamsRegionEuropeWest8, SpeedScheduleGetParamsRegionEuropeWest9, SpeedScheduleGetParamsRegionMeWest1, SpeedScheduleGetParamsRegionSouthamericaEast1, SpeedScheduleGetParamsRegionUsCentral1, SpeedScheduleGetParamsRegionUsEast1, SpeedScheduleGetParamsRegionUsEast4, SpeedScheduleGetParamsRegionUsSouth1, SpeedScheduleGetParamsRegionUsWest1:
		return true
	}
	return false
}

type SpeedScheduleGetResponseEnvelope struct {
	Errors   interface{} `json:"errors,required"`
	Messages interface{} `json:"messages,required"`
	Success  interface{} `json:"success,required"`
	// The test schedule.
	Result ObservatorySchedule                  `json:"result"`
	JSON   speedScheduleGetResponseEnvelopeJSON `json:"-"`
}

// speedScheduleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedScheduleGetResponseEnvelope]
type speedScheduleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedScheduleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedScheduleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SpeedTrendsListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The type of device.
	DeviceType param.Field[SpeedTrendsListParamsDeviceType] `query:"deviceType,required"`
	// A comma-separated list of metrics to include in the results.
	Metrics param.Field[string] `query:"metrics,required"`
	// A test region.
	Region param.Field[SpeedTrendsListParamsRegion] `query:"region,required"`
	Start  param.Field[time.Time]                   `query:"start,required" format:"date-time"`
	// The timezone of the start and end timestamps.
	Tz  param.Field[string]    `query:"tz,required"`
	End param.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes [SpeedTrendsListParams]'s query parameters as `url.Values`.
func (r SpeedTrendsListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of device.
type SpeedTrendsListParamsDeviceType string

const (
	SpeedTrendsListParamsDeviceTypeDesktop SpeedTrendsListParamsDeviceType = "DESKTOP"
	SpeedTrendsListParamsDeviceTypeMobile  SpeedTrendsListParamsDeviceType = "MOBILE"
)

func (r SpeedTrendsListParamsDeviceType) IsKnown() bool {
	switch r {
	case SpeedTrendsListParamsDeviceTypeDesktop, SpeedTrendsListParamsDeviceTypeMobile:
		return true
	}
	return false
}

// A test region.
type SpeedTrendsListParamsRegion string

const (
	SpeedTrendsListParamsRegionAsiaEast1           SpeedTrendsListParamsRegion = "asia-east1"
	SpeedTrendsListParamsRegionAsiaNortheast1      SpeedTrendsListParamsRegion = "asia-northeast1"
	SpeedTrendsListParamsRegionAsiaNortheast2      SpeedTrendsListParamsRegion = "asia-northeast2"
	SpeedTrendsListParamsRegionAsiaSouth1          SpeedTrendsListParamsRegion = "asia-south1"
	SpeedTrendsListParamsRegionAsiaSoutheast1      SpeedTrendsListParamsRegion = "asia-southeast1"
	SpeedTrendsListParamsRegionAustraliaSoutheast1 SpeedTrendsListParamsRegion = "australia-southeast1"
	SpeedTrendsListParamsRegionEuropeNorth1        SpeedTrendsListParamsRegion = "europe-north1"
	SpeedTrendsListParamsRegionEuropeSouthwest1    SpeedTrendsListParamsRegion = "europe-southwest1"
	SpeedTrendsListParamsRegionEuropeWest1         SpeedTrendsListParamsRegion = "europe-west1"
	SpeedTrendsListParamsRegionEuropeWest2         SpeedTrendsListParamsRegion = "europe-west2"
	SpeedTrendsListParamsRegionEuropeWest3         SpeedTrendsListParamsRegion = "europe-west3"
	SpeedTrendsListParamsRegionEuropeWest4         SpeedTrendsListParamsRegion = "europe-west4"
	SpeedTrendsListParamsRegionEuropeWest8         SpeedTrendsListParamsRegion = "europe-west8"
	SpeedTrendsListParamsRegionEuropeWest9         SpeedTrendsListParamsRegion = "europe-west9"
	SpeedTrendsListParamsRegionMeWest1             SpeedTrendsListParamsRegion = "me-west1"
	SpeedTrendsListParamsRegionSouthamericaEast1   SpeedTrendsListParamsRegion = "southamerica-east1"
	SpeedTrendsListParamsRegionUsCentral1          SpeedTrendsListParamsRegion = "us-central1"
	SpeedTrendsListParamsRegionUsEast1             SpeedTrendsListParamsRegion = "us-east1"
	SpeedTrendsListParamsRegionUsEast4             SpeedTrendsListParamsRegion = "us-east4"
	SpeedTrendsListParamsRegionUsSouth1            SpeedTrendsListParamsRegion = "us-south1"
	SpeedTrendsListParamsRegionUsWest1             SpeedTrendsListParamsRegion = "us-west1"
)

func (r SpeedTrendsListParamsRegion) IsKnown() bool {
	switch r {
	case SpeedTrendsListParamsRegionAsiaEast1, SpeedTrendsListParamsRegionAsiaNortheast1, SpeedTrendsListParamsRegionAsiaNortheast2, SpeedTrendsListParamsRegionAsiaSouth1, SpeedTrendsListParamsRegionAsiaSoutheast1, SpeedTrendsListParamsRegionAustraliaSoutheast1, SpeedTrendsListParamsRegionEuropeNorth1, SpeedTrendsListParamsRegionEuropeSouthwest1, SpeedTrendsListParamsRegionEuropeWest1, SpeedTrendsListParamsRegionEuropeWest2, SpeedTrendsListParamsRegionEuropeWest3, SpeedTrendsListParamsRegionEuropeWest4, SpeedTrendsListParamsRegionEuropeWest8, SpeedTrendsListParamsRegionEuropeWest9, SpeedTrendsListParamsRegionMeWest1, SpeedTrendsListParamsRegionSouthamericaEast1, SpeedTrendsListParamsRegionUsCentral1, SpeedTrendsListParamsRegionUsEast1, SpeedTrendsListParamsRegionUsEast4, SpeedTrendsListParamsRegionUsSouth1, SpeedTrendsListParamsRegionUsWest1:
		return true
	}
	return false
}

type SpeedTrendsListResponseEnvelope struct {
	Errors   interface{}                         `json:"errors,required"`
	Messages interface{}                         `json:"messages,required"`
	Success  interface{}                         `json:"success,required"`
	Result   ObservatoryTrend                    `json:"result"`
	JSON     speedTrendsListResponseEnvelopeJSON `json:"-"`
}

// speedTrendsListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedTrendsListResponseEnvelope]
type speedTrendsListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTrendsListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedTrendsListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
