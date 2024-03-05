// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SpeedTestService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSpeedTestService] method instead.
type SpeedTestService struct {
	Options []option.RequestOption
}

// NewSpeedTestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpeedTestService(opts ...option.RequestOption) (r *SpeedTestService) {
	r = &SpeedTestService{}
	r.Options = opts
	return
}

// Starts a test for a specific webpage, in a specific region.
func (r *SpeedTestService) New(ctx context.Context, url string, params SpeedTestNewParams, opts ...option.RequestOption) (res *ObservatoryPageTest, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedTestNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Test history (list of tests) for a specific webpage.
func (r *SpeedTestService) List(ctx context.Context, url string, params SpeedTestListParams, opts ...option.RequestOption) (res *SpeedTestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Deletes all tests for a specific webpage from a specific region. Deleted tests
// are still counted as part of the quota.
func (r *SpeedTestService) Delete(ctx context.Context, url string, params SpeedTestDeleteParams, opts ...option.RequestOption) (res *SpeedTestDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedTestDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the result of a specific test.
func (r *SpeedTestService) Get(ctx context.Context, url string, testID string, query SpeedTestGetParams, opts ...option.RequestOption) (res *ObservatoryPageTest, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedTestGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests/%s", query.ZoneID, url, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ObservatoryPageTest struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport ObservatoryPageTestDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport ObservatoryPageTestMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region ObservatoryPageTestRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency ObservatoryPageTestScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                  `json:"url"`
	JSON observatoryPageTestJSON `json:"-"`
}

// observatoryPageTestJSON contains the JSON metadata for the struct
// [ObservatoryPageTest]
type observatoryPageTestJSON struct {
	ID                apijson.Field
	Date              apijson.Field
	DesktopReport     apijson.Field
	MobileReport      apijson.Field
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservatoryPageTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type ObservatoryPageTestDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ObservatoryPageTestDesktopReportDeviceType `json:"deviceType"`
	Error      ObservatoryPageTestDesktopReportError      `json:"error"`
	// First Contentful Paint.
	Fcp float64 `json:"fcp"`
	// The URL to the full Lighthouse JSON report.
	JsonReportURL string `json:"jsonReportUrl"`
	// Largest Contentful Paint.
	Lcp float64 `json:"lcp"`
	// The Lighthouse performance score.
	PerformanceScore float64 `json:"performanceScore"`
	// Speed Index.
	Si float64 `json:"si"`
	// The state of the Lighthouse report.
	State ObservatoryPageTestDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                              `json:"tti"`
	JSON observatoryPageTestDesktopReportJSON `json:"-"`
}

// observatoryPageTestDesktopReportJSON contains the JSON metadata for the struct
// [ObservatoryPageTestDesktopReport]
type observatoryPageTestDesktopReportJSON struct {
	Cls              apijson.Field
	DeviceType       apijson.Field
	Error            apijson.Field
	Fcp              apijson.Field
	JsonReportURL    apijson.Field
	Lcp              apijson.Field
	PerformanceScore apijson.Field
	Si               apijson.Field
	State            apijson.Field
	Tbt              apijson.Field
	Ttfb             apijson.Field
	Tti              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ObservatoryPageTestDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type ObservatoryPageTestDesktopReportDeviceType string

const (
	ObservatoryPageTestDesktopReportDeviceTypeDesktop ObservatoryPageTestDesktopReportDeviceType = "DESKTOP"
	ObservatoryPageTestDesktopReportDeviceTypeMobile  ObservatoryPageTestDesktopReportDeviceType = "MOBILE"
)

type ObservatoryPageTestDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code ObservatoryPageTestDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                    `json:"finalDisplayedUrl"`
	JSON              observatoryPageTestDesktopReportErrorJSON `json:"-"`
}

// observatoryPageTestDesktopReportErrorJSON contains the JSON metadata for the
// struct [ObservatoryPageTestDesktopReportError]
type observatoryPageTestDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservatoryPageTestDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type ObservatoryPageTestDesktopReportErrorCode string

const (
	ObservatoryPageTestDesktopReportErrorCodeNotReachable      ObservatoryPageTestDesktopReportErrorCode = "NOT_REACHABLE"
	ObservatoryPageTestDesktopReportErrorCodeDNSFailure        ObservatoryPageTestDesktopReportErrorCode = "DNS_FAILURE"
	ObservatoryPageTestDesktopReportErrorCodeNotHTML           ObservatoryPageTestDesktopReportErrorCode = "NOT_HTML"
	ObservatoryPageTestDesktopReportErrorCodeLighthouseTimeout ObservatoryPageTestDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ObservatoryPageTestDesktopReportErrorCodeUnknown           ObservatoryPageTestDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ObservatoryPageTestDesktopReportState string

const (
	ObservatoryPageTestDesktopReportStateRunning  ObservatoryPageTestDesktopReportState = "RUNNING"
	ObservatoryPageTestDesktopReportStateComplete ObservatoryPageTestDesktopReportState = "COMPLETE"
	ObservatoryPageTestDesktopReportStateFailed   ObservatoryPageTestDesktopReportState = "FAILED"
)

// The Lighthouse report.
type ObservatoryPageTestMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ObservatoryPageTestMobileReportDeviceType `json:"deviceType"`
	Error      ObservatoryPageTestMobileReportError      `json:"error"`
	// First Contentful Paint.
	Fcp float64 `json:"fcp"`
	// The URL to the full Lighthouse JSON report.
	JsonReportURL string `json:"jsonReportUrl"`
	// Largest Contentful Paint.
	Lcp float64 `json:"lcp"`
	// The Lighthouse performance score.
	PerformanceScore float64 `json:"performanceScore"`
	// Speed Index.
	Si float64 `json:"si"`
	// The state of the Lighthouse report.
	State ObservatoryPageTestMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                             `json:"tti"`
	JSON observatoryPageTestMobileReportJSON `json:"-"`
}

// observatoryPageTestMobileReportJSON contains the JSON metadata for the struct
// [ObservatoryPageTestMobileReport]
type observatoryPageTestMobileReportJSON struct {
	Cls              apijson.Field
	DeviceType       apijson.Field
	Error            apijson.Field
	Fcp              apijson.Field
	JsonReportURL    apijson.Field
	Lcp              apijson.Field
	PerformanceScore apijson.Field
	Si               apijson.Field
	State            apijson.Field
	Tbt              apijson.Field
	Ttfb             apijson.Field
	Tti              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ObservatoryPageTestMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type ObservatoryPageTestMobileReportDeviceType string

const (
	ObservatoryPageTestMobileReportDeviceTypeDesktop ObservatoryPageTestMobileReportDeviceType = "DESKTOP"
	ObservatoryPageTestMobileReportDeviceTypeMobile  ObservatoryPageTestMobileReportDeviceType = "MOBILE"
)

type ObservatoryPageTestMobileReportError struct {
	// The error code of the Lighthouse result.
	Code ObservatoryPageTestMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                   `json:"finalDisplayedUrl"`
	JSON              observatoryPageTestMobileReportErrorJSON `json:"-"`
}

// observatoryPageTestMobileReportErrorJSON contains the JSON metadata for the
// struct [ObservatoryPageTestMobileReportError]
type observatoryPageTestMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservatoryPageTestMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type ObservatoryPageTestMobileReportErrorCode string

const (
	ObservatoryPageTestMobileReportErrorCodeNotReachable      ObservatoryPageTestMobileReportErrorCode = "NOT_REACHABLE"
	ObservatoryPageTestMobileReportErrorCodeDNSFailure        ObservatoryPageTestMobileReportErrorCode = "DNS_FAILURE"
	ObservatoryPageTestMobileReportErrorCodeNotHTML           ObservatoryPageTestMobileReportErrorCode = "NOT_HTML"
	ObservatoryPageTestMobileReportErrorCodeLighthouseTimeout ObservatoryPageTestMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ObservatoryPageTestMobileReportErrorCodeUnknown           ObservatoryPageTestMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ObservatoryPageTestMobileReportState string

const (
	ObservatoryPageTestMobileReportStateRunning  ObservatoryPageTestMobileReportState = "RUNNING"
	ObservatoryPageTestMobileReportStateComplete ObservatoryPageTestMobileReportState = "COMPLETE"
	ObservatoryPageTestMobileReportStateFailed   ObservatoryPageTestMobileReportState = "FAILED"
)

// A test region with a label.
type ObservatoryPageTestRegion struct {
	Label string `json:"label"`
	// A test region.
	Value ObservatoryPageTestRegionValue `json:"value"`
	JSON  observatoryPageTestRegionJSON  `json:"-"`
}

// observatoryPageTestRegionJSON contains the JSON metadata for the struct
// [ObservatoryPageTestRegion]
type observatoryPageTestRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservatoryPageTestRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type ObservatoryPageTestRegionValue string

const (
	ObservatoryPageTestRegionValueAsiaEast1           ObservatoryPageTestRegionValue = "asia-east1"
	ObservatoryPageTestRegionValueAsiaNortheast1      ObservatoryPageTestRegionValue = "asia-northeast1"
	ObservatoryPageTestRegionValueAsiaNortheast2      ObservatoryPageTestRegionValue = "asia-northeast2"
	ObservatoryPageTestRegionValueAsiaSouth1          ObservatoryPageTestRegionValue = "asia-south1"
	ObservatoryPageTestRegionValueAsiaSoutheast1      ObservatoryPageTestRegionValue = "asia-southeast1"
	ObservatoryPageTestRegionValueAustraliaSoutheast1 ObservatoryPageTestRegionValue = "australia-southeast1"
	ObservatoryPageTestRegionValueEuropeNorth1        ObservatoryPageTestRegionValue = "europe-north1"
	ObservatoryPageTestRegionValueEuropeSouthwest1    ObservatoryPageTestRegionValue = "europe-southwest1"
	ObservatoryPageTestRegionValueEuropeWest1         ObservatoryPageTestRegionValue = "europe-west1"
	ObservatoryPageTestRegionValueEuropeWest2         ObservatoryPageTestRegionValue = "europe-west2"
	ObservatoryPageTestRegionValueEuropeWest3         ObservatoryPageTestRegionValue = "europe-west3"
	ObservatoryPageTestRegionValueEuropeWest4         ObservatoryPageTestRegionValue = "europe-west4"
	ObservatoryPageTestRegionValueEuropeWest8         ObservatoryPageTestRegionValue = "europe-west8"
	ObservatoryPageTestRegionValueEuropeWest9         ObservatoryPageTestRegionValue = "europe-west9"
	ObservatoryPageTestRegionValueMeWest1             ObservatoryPageTestRegionValue = "me-west1"
	ObservatoryPageTestRegionValueSouthamericaEast1   ObservatoryPageTestRegionValue = "southamerica-east1"
	ObservatoryPageTestRegionValueUsCentral1          ObservatoryPageTestRegionValue = "us-central1"
	ObservatoryPageTestRegionValueUsEast1             ObservatoryPageTestRegionValue = "us-east1"
	ObservatoryPageTestRegionValueUsEast4             ObservatoryPageTestRegionValue = "us-east4"
	ObservatoryPageTestRegionValueUsSouth1            ObservatoryPageTestRegionValue = "us-south1"
	ObservatoryPageTestRegionValueUsWest1             ObservatoryPageTestRegionValue = "us-west1"
)

// The frequency of the test.
type ObservatoryPageTestScheduleFrequency string

const (
	ObservatoryPageTestScheduleFrequencyDaily  ObservatoryPageTestScheduleFrequency = "DAILY"
	ObservatoryPageTestScheduleFrequencyWeekly ObservatoryPageTestScheduleFrequency = "WEEKLY"
)

type SpeedTestListResponse struct {
	Errors   []SpeedTestListResponseError   `json:"errors,required"`
	Messages []SpeedTestListResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    bool                            `json:"success,required"`
	ResultInfo SpeedTestListResponseResultInfo `json:"result_info"`
	JSON       speedTestListResponseJSON       `json:"-"`
}

// speedTestListResponseJSON contains the JSON metadata for the struct
// [SpeedTestListResponse]
type speedTestListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedTestListResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    speedTestListResponseErrorJSON `json:"-"`
}

// speedTestListResponseErrorJSON contains the JSON metadata for the struct
// [SpeedTestListResponseError]
type speedTestListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedTestListResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    speedTestListResponseMessageJSON `json:"-"`
}

// speedTestListResponseMessageJSON contains the JSON metadata for the struct
// [SpeedTestListResponseMessage]
type speedTestListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedTestListResponseResultInfo struct {
	Count      int64                               `json:"count"`
	Page       int64                               `json:"page"`
	PerPage    int64                               `json:"per_page"`
	TotalCount int64                               `json:"total_count"`
	JSON       speedTestListResponseResultInfoJSON `json:"-"`
}

// speedTestListResponseResultInfoJSON contains the JSON metadata for the struct
// [SpeedTestListResponseResultInfo]
type speedTestListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedTestDeleteResponse struct {
	// Number of items affected.
	Count float64                     `json:"count"`
	JSON  speedTestDeleteResponseJSON `json:"-"`
}

// speedTestDeleteResponseJSON contains the JSON metadata for the struct
// [SpeedTestDeleteResponse]
type speedTestDeleteResponseJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedTestNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[SpeedTestNewParamsRegion] `json:"region"`
}

func (r SpeedTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A test region.
type SpeedTestNewParamsRegion string

const (
	SpeedTestNewParamsRegionAsiaEast1           SpeedTestNewParamsRegion = "asia-east1"
	SpeedTestNewParamsRegionAsiaNortheast1      SpeedTestNewParamsRegion = "asia-northeast1"
	SpeedTestNewParamsRegionAsiaNortheast2      SpeedTestNewParamsRegion = "asia-northeast2"
	SpeedTestNewParamsRegionAsiaSouth1          SpeedTestNewParamsRegion = "asia-south1"
	SpeedTestNewParamsRegionAsiaSoutheast1      SpeedTestNewParamsRegion = "asia-southeast1"
	SpeedTestNewParamsRegionAustraliaSoutheast1 SpeedTestNewParamsRegion = "australia-southeast1"
	SpeedTestNewParamsRegionEuropeNorth1        SpeedTestNewParamsRegion = "europe-north1"
	SpeedTestNewParamsRegionEuropeSouthwest1    SpeedTestNewParamsRegion = "europe-southwest1"
	SpeedTestNewParamsRegionEuropeWest1         SpeedTestNewParamsRegion = "europe-west1"
	SpeedTestNewParamsRegionEuropeWest2         SpeedTestNewParamsRegion = "europe-west2"
	SpeedTestNewParamsRegionEuropeWest3         SpeedTestNewParamsRegion = "europe-west3"
	SpeedTestNewParamsRegionEuropeWest4         SpeedTestNewParamsRegion = "europe-west4"
	SpeedTestNewParamsRegionEuropeWest8         SpeedTestNewParamsRegion = "europe-west8"
	SpeedTestNewParamsRegionEuropeWest9         SpeedTestNewParamsRegion = "europe-west9"
	SpeedTestNewParamsRegionMeWest1             SpeedTestNewParamsRegion = "me-west1"
	SpeedTestNewParamsRegionSouthamericaEast1   SpeedTestNewParamsRegion = "southamerica-east1"
	SpeedTestNewParamsRegionUsCentral1          SpeedTestNewParamsRegion = "us-central1"
	SpeedTestNewParamsRegionUsEast1             SpeedTestNewParamsRegion = "us-east1"
	SpeedTestNewParamsRegionUsEast4             SpeedTestNewParamsRegion = "us-east4"
	SpeedTestNewParamsRegionUsSouth1            SpeedTestNewParamsRegion = "us-south1"
	SpeedTestNewParamsRegionUsWest1             SpeedTestNewParamsRegion = "us-west1"
)

type SpeedTestNewResponseEnvelope struct {
	Result ObservatoryPageTest              `json:"result"`
	JSON   speedTestNewResponseEnvelopeJSON `json:"-"`
}

// speedTestNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedTestNewResponseEnvelope]
type speedTestNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedTestListParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Page    param.Field[int64]  `query:"page"`
	PerPage param.Field[int64]  `query:"per_page"`
	// A test region.
	Region param.Field[SpeedTestListParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedTestListParams]'s query parameters as `url.Values`.
func (r SpeedTestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedTestListParamsRegion string

const (
	SpeedTestListParamsRegionAsiaEast1           SpeedTestListParamsRegion = "asia-east1"
	SpeedTestListParamsRegionAsiaNortheast1      SpeedTestListParamsRegion = "asia-northeast1"
	SpeedTestListParamsRegionAsiaNortheast2      SpeedTestListParamsRegion = "asia-northeast2"
	SpeedTestListParamsRegionAsiaSouth1          SpeedTestListParamsRegion = "asia-south1"
	SpeedTestListParamsRegionAsiaSoutheast1      SpeedTestListParamsRegion = "asia-southeast1"
	SpeedTestListParamsRegionAustraliaSoutheast1 SpeedTestListParamsRegion = "australia-southeast1"
	SpeedTestListParamsRegionEuropeNorth1        SpeedTestListParamsRegion = "europe-north1"
	SpeedTestListParamsRegionEuropeSouthwest1    SpeedTestListParamsRegion = "europe-southwest1"
	SpeedTestListParamsRegionEuropeWest1         SpeedTestListParamsRegion = "europe-west1"
	SpeedTestListParamsRegionEuropeWest2         SpeedTestListParamsRegion = "europe-west2"
	SpeedTestListParamsRegionEuropeWest3         SpeedTestListParamsRegion = "europe-west3"
	SpeedTestListParamsRegionEuropeWest4         SpeedTestListParamsRegion = "europe-west4"
	SpeedTestListParamsRegionEuropeWest8         SpeedTestListParamsRegion = "europe-west8"
	SpeedTestListParamsRegionEuropeWest9         SpeedTestListParamsRegion = "europe-west9"
	SpeedTestListParamsRegionMeWest1             SpeedTestListParamsRegion = "me-west1"
	SpeedTestListParamsRegionSouthamericaEast1   SpeedTestListParamsRegion = "southamerica-east1"
	SpeedTestListParamsRegionUsCentral1          SpeedTestListParamsRegion = "us-central1"
	SpeedTestListParamsRegionUsEast1             SpeedTestListParamsRegion = "us-east1"
	SpeedTestListParamsRegionUsEast4             SpeedTestListParamsRegion = "us-east4"
	SpeedTestListParamsRegionUsSouth1            SpeedTestListParamsRegion = "us-south1"
	SpeedTestListParamsRegionUsWest1             SpeedTestListParamsRegion = "us-west1"
)

type SpeedTestDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[SpeedTestDeleteParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedTestDeleteParams]'s query parameters as `url.Values`.
func (r SpeedTestDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedTestDeleteParamsRegion string

const (
	SpeedTestDeleteParamsRegionAsiaEast1           SpeedTestDeleteParamsRegion = "asia-east1"
	SpeedTestDeleteParamsRegionAsiaNortheast1      SpeedTestDeleteParamsRegion = "asia-northeast1"
	SpeedTestDeleteParamsRegionAsiaNortheast2      SpeedTestDeleteParamsRegion = "asia-northeast2"
	SpeedTestDeleteParamsRegionAsiaSouth1          SpeedTestDeleteParamsRegion = "asia-south1"
	SpeedTestDeleteParamsRegionAsiaSoutheast1      SpeedTestDeleteParamsRegion = "asia-southeast1"
	SpeedTestDeleteParamsRegionAustraliaSoutheast1 SpeedTestDeleteParamsRegion = "australia-southeast1"
	SpeedTestDeleteParamsRegionEuropeNorth1        SpeedTestDeleteParamsRegion = "europe-north1"
	SpeedTestDeleteParamsRegionEuropeSouthwest1    SpeedTestDeleteParamsRegion = "europe-southwest1"
	SpeedTestDeleteParamsRegionEuropeWest1         SpeedTestDeleteParamsRegion = "europe-west1"
	SpeedTestDeleteParamsRegionEuropeWest2         SpeedTestDeleteParamsRegion = "europe-west2"
	SpeedTestDeleteParamsRegionEuropeWest3         SpeedTestDeleteParamsRegion = "europe-west3"
	SpeedTestDeleteParamsRegionEuropeWest4         SpeedTestDeleteParamsRegion = "europe-west4"
	SpeedTestDeleteParamsRegionEuropeWest8         SpeedTestDeleteParamsRegion = "europe-west8"
	SpeedTestDeleteParamsRegionEuropeWest9         SpeedTestDeleteParamsRegion = "europe-west9"
	SpeedTestDeleteParamsRegionMeWest1             SpeedTestDeleteParamsRegion = "me-west1"
	SpeedTestDeleteParamsRegionSouthamericaEast1   SpeedTestDeleteParamsRegion = "southamerica-east1"
	SpeedTestDeleteParamsRegionUsCentral1          SpeedTestDeleteParamsRegion = "us-central1"
	SpeedTestDeleteParamsRegionUsEast1             SpeedTestDeleteParamsRegion = "us-east1"
	SpeedTestDeleteParamsRegionUsEast4             SpeedTestDeleteParamsRegion = "us-east4"
	SpeedTestDeleteParamsRegionUsSouth1            SpeedTestDeleteParamsRegion = "us-south1"
	SpeedTestDeleteParamsRegionUsWest1             SpeedTestDeleteParamsRegion = "us-west1"
)

type SpeedTestDeleteResponseEnvelope struct {
	Result SpeedTestDeleteResponse             `json:"result"`
	JSON   speedTestDeleteResponseEnvelopeJSON `json:"-"`
}

// speedTestDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedTestDeleteResponseEnvelope]
type speedTestDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedTestGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SpeedTestGetResponseEnvelope struct {
	Result ObservatoryPageTest              `json:"result"`
	JSON   speedTestGetResponseEnvelopeJSON `json:"-"`
}

// speedTestGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedTestGetResponseEnvelope]
type speedTestGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
