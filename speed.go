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

// SpeedService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewSpeedService] method instead.
type SpeedService struct {
	Options        []option.RequestOption
	Tests          *SpeedTestService
	Schedule       *SpeedScheduleService
	Availabilities *SpeedAvailabilityService
	Pages          *SpeedPageService
}

// NewSpeedService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSpeedService(opts ...option.RequestOption) (r *SpeedService) {
	r = &SpeedService{}
	r.Options = opts
	r.Tests = NewSpeedTestService(opts...)
	r.Schedule = NewSpeedScheduleService(opts...)
	r.Availabilities = NewSpeedAvailabilityService(opts...)
	r.Pages = NewSpeedPageService(opts...)
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
func (r *SpeedService) ScheduleGet(ctx context.Context, url string, params SpeedScheduleGetParams, opts ...option.RequestOption) (res *SpeedScheduleGetResponse, err error) {
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
func (r *SpeedService) TrendsList(ctx context.Context, url string, params SpeedTrendsListParams, opts ...option.RequestOption) (res *SpeedTrendsListResponse, err error) {
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

// The test schedule.
type SpeedScheduleGetResponse struct {
	// The frequency of the test.
	Frequency SpeedScheduleGetResponseFrequency `json:"frequency"`
	// A test region.
	Region SpeedScheduleGetResponseRegion `json:"region"`
	// A URL.
	URL  string                       `json:"url"`
	JSON speedScheduleGetResponseJSON `json:"-"`
}

// speedScheduleGetResponseJSON contains the JSON metadata for the struct
// [SpeedScheduleGetResponse]
type speedScheduleGetResponseJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedScheduleGetResponseJSON) RawJSON() string {
	return r.raw
}

// The frequency of the test.
type SpeedScheduleGetResponseFrequency string

const (
	SpeedScheduleGetResponseFrequencyDaily  SpeedScheduleGetResponseFrequency = "DAILY"
	SpeedScheduleGetResponseFrequencyWeekly SpeedScheduleGetResponseFrequency = "WEEKLY"
)

// A test region.
type SpeedScheduleGetResponseRegion string

const (
	SpeedScheduleGetResponseRegionAsiaEast1           SpeedScheduleGetResponseRegion = "asia-east1"
	SpeedScheduleGetResponseRegionAsiaNortheast1      SpeedScheduleGetResponseRegion = "asia-northeast1"
	SpeedScheduleGetResponseRegionAsiaNortheast2      SpeedScheduleGetResponseRegion = "asia-northeast2"
	SpeedScheduleGetResponseRegionAsiaSouth1          SpeedScheduleGetResponseRegion = "asia-south1"
	SpeedScheduleGetResponseRegionAsiaSoutheast1      SpeedScheduleGetResponseRegion = "asia-southeast1"
	SpeedScheduleGetResponseRegionAustraliaSoutheast1 SpeedScheduleGetResponseRegion = "australia-southeast1"
	SpeedScheduleGetResponseRegionEuropeNorth1        SpeedScheduleGetResponseRegion = "europe-north1"
	SpeedScheduleGetResponseRegionEuropeSouthwest1    SpeedScheduleGetResponseRegion = "europe-southwest1"
	SpeedScheduleGetResponseRegionEuropeWest1         SpeedScheduleGetResponseRegion = "europe-west1"
	SpeedScheduleGetResponseRegionEuropeWest2         SpeedScheduleGetResponseRegion = "europe-west2"
	SpeedScheduleGetResponseRegionEuropeWest3         SpeedScheduleGetResponseRegion = "europe-west3"
	SpeedScheduleGetResponseRegionEuropeWest4         SpeedScheduleGetResponseRegion = "europe-west4"
	SpeedScheduleGetResponseRegionEuropeWest8         SpeedScheduleGetResponseRegion = "europe-west8"
	SpeedScheduleGetResponseRegionEuropeWest9         SpeedScheduleGetResponseRegion = "europe-west9"
	SpeedScheduleGetResponseRegionMeWest1             SpeedScheduleGetResponseRegion = "me-west1"
	SpeedScheduleGetResponseRegionSouthamericaEast1   SpeedScheduleGetResponseRegion = "southamerica-east1"
	SpeedScheduleGetResponseRegionUsCentral1          SpeedScheduleGetResponseRegion = "us-central1"
	SpeedScheduleGetResponseRegionUsEast1             SpeedScheduleGetResponseRegion = "us-east1"
	SpeedScheduleGetResponseRegionUsEast4             SpeedScheduleGetResponseRegion = "us-east4"
	SpeedScheduleGetResponseRegionUsSouth1            SpeedScheduleGetResponseRegion = "us-south1"
	SpeedScheduleGetResponseRegionUsWest1             SpeedScheduleGetResponseRegion = "us-west1"
)

type SpeedTrendsListResponse struct {
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
	Tti  []float64                   `json:"tti"`
	JSON speedTrendsListResponseJSON `json:"-"`
}

// speedTrendsListResponseJSON contains the JSON metadata for the struct
// [SpeedTrendsListResponse]
type speedTrendsListResponseJSON struct {
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

func (r *SpeedTrendsListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedTrendsListResponseJSON) RawJSON() string {
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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

type SpeedDeleteResponseEnvelope struct {
	Result SpeedDeleteResponse             `json:"result"`
	JSON   speedDeleteResponseEnvelopeJSON `json:"-"`
}

// speedDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedDeleteResponseEnvelope]
type speedDeleteResponseEnvelopeJSON struct {
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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

type SpeedScheduleGetResponseEnvelope struct {
	// The test schedule.
	Result SpeedScheduleGetResponse             `json:"result"`
	JSON   speedScheduleGetResponseEnvelopeJSON `json:"-"`
}

// speedScheduleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedScheduleGetResponseEnvelope]
type speedScheduleGetResponseEnvelopeJSON struct {
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
	// The timezone of the start and end timestamps.
	Tz param.Field[string] `query:"tz,required"`
}

// URLQuery serializes [SpeedTrendsListParams]'s query parameters as `url.Values`.
func (r SpeedTrendsListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of device.
type SpeedTrendsListParamsDeviceType string

const (
	SpeedTrendsListParamsDeviceTypeDesktop SpeedTrendsListParamsDeviceType = "DESKTOP"
	SpeedTrendsListParamsDeviceTypeMobile  SpeedTrendsListParamsDeviceType = "MOBILE"
)

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

type SpeedTrendsListResponseEnvelope struct {
	Result SpeedTrendsListResponse             `json:"result"`
	JSON   speedTrendsListResponseEnvelopeJSON `json:"-"`
}

// speedTrendsListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedTrendsListResponseEnvelope]
type speedTrendsListResponseEnvelopeJSON struct {
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
