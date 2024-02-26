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
func (r *SpeedTestService) New(ctx context.Context, url string, params SpeedTestNewParams, opts ...option.RequestOption) (res *SpeedTestNewResponse, err error) {
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
func (r *SpeedTestService) Get(ctx context.Context, url string, testID string, query SpeedTestGetParams, opts ...option.RequestOption) (res *SpeedTestGetResponse, err error) {
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

type SpeedTestNewResponse struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedTestNewResponseDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedTestNewResponseMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedTestNewResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedTestNewResponseScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                   `json:"url"`
	JSON speedTestNewResponseJSON `json:"-"`
}

// speedTestNewResponseJSON contains the JSON metadata for the struct
// [SpeedTestNewResponse]
type speedTestNewResponseJSON struct {
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

func (r *SpeedTestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedTestNewResponseDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedTestNewResponseDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedTestNewResponseDesktopReportError      `json:"error"`
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
	State SpeedTestNewResponseDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                               `json:"tti"`
	JSON speedTestNewResponseDesktopReportJSON `json:"-"`
}

// speedTestNewResponseDesktopReportJSON contains the JSON metadata for the struct
// [SpeedTestNewResponseDesktopReport]
type speedTestNewResponseDesktopReportJSON struct {
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

func (r *SpeedTestNewResponseDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedTestNewResponseDesktopReportDeviceType string

const (
	SpeedTestNewResponseDesktopReportDeviceTypeDesktop SpeedTestNewResponseDesktopReportDeviceType = "DESKTOP"
	SpeedTestNewResponseDesktopReportDeviceTypeMobile  SpeedTestNewResponseDesktopReportDeviceType = "MOBILE"
)

type SpeedTestNewResponseDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedTestNewResponseDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                     `json:"finalDisplayedUrl"`
	JSON              speedTestNewResponseDesktopReportErrorJSON `json:"-"`
}

// speedTestNewResponseDesktopReportErrorJSON contains the JSON metadata for the
// struct [SpeedTestNewResponseDesktopReportError]
type speedTestNewResponseDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedTestNewResponseDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedTestNewResponseDesktopReportErrorCode string

const (
	SpeedTestNewResponseDesktopReportErrorCodeNotReachable      SpeedTestNewResponseDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedTestNewResponseDesktopReportErrorCodeDNSFailure        SpeedTestNewResponseDesktopReportErrorCode = "DNS_FAILURE"
	SpeedTestNewResponseDesktopReportErrorCodeNotHTML           SpeedTestNewResponseDesktopReportErrorCode = "NOT_HTML"
	SpeedTestNewResponseDesktopReportErrorCodeLighthouseTimeout SpeedTestNewResponseDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedTestNewResponseDesktopReportErrorCodeUnknown           SpeedTestNewResponseDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedTestNewResponseDesktopReportState string

const (
	SpeedTestNewResponseDesktopReportStateRunning  SpeedTestNewResponseDesktopReportState = "RUNNING"
	SpeedTestNewResponseDesktopReportStateComplete SpeedTestNewResponseDesktopReportState = "COMPLETE"
	SpeedTestNewResponseDesktopReportStateFailed   SpeedTestNewResponseDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedTestNewResponseMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedTestNewResponseMobileReportDeviceType `json:"deviceType"`
	Error      SpeedTestNewResponseMobileReportError      `json:"error"`
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
	State SpeedTestNewResponseMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                              `json:"tti"`
	JSON speedTestNewResponseMobileReportJSON `json:"-"`
}

// speedTestNewResponseMobileReportJSON contains the JSON metadata for the struct
// [SpeedTestNewResponseMobileReport]
type speedTestNewResponseMobileReportJSON struct {
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

func (r *SpeedTestNewResponseMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedTestNewResponseMobileReportDeviceType string

const (
	SpeedTestNewResponseMobileReportDeviceTypeDesktop SpeedTestNewResponseMobileReportDeviceType = "DESKTOP"
	SpeedTestNewResponseMobileReportDeviceTypeMobile  SpeedTestNewResponseMobileReportDeviceType = "MOBILE"
)

type SpeedTestNewResponseMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedTestNewResponseMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                    `json:"finalDisplayedUrl"`
	JSON              speedTestNewResponseMobileReportErrorJSON `json:"-"`
}

// speedTestNewResponseMobileReportErrorJSON contains the JSON metadata for the
// struct [SpeedTestNewResponseMobileReportError]
type speedTestNewResponseMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedTestNewResponseMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedTestNewResponseMobileReportErrorCode string

const (
	SpeedTestNewResponseMobileReportErrorCodeNotReachable      SpeedTestNewResponseMobileReportErrorCode = "NOT_REACHABLE"
	SpeedTestNewResponseMobileReportErrorCodeDNSFailure        SpeedTestNewResponseMobileReportErrorCode = "DNS_FAILURE"
	SpeedTestNewResponseMobileReportErrorCodeNotHTML           SpeedTestNewResponseMobileReportErrorCode = "NOT_HTML"
	SpeedTestNewResponseMobileReportErrorCodeLighthouseTimeout SpeedTestNewResponseMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedTestNewResponseMobileReportErrorCodeUnknown           SpeedTestNewResponseMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedTestNewResponseMobileReportState string

const (
	SpeedTestNewResponseMobileReportStateRunning  SpeedTestNewResponseMobileReportState = "RUNNING"
	SpeedTestNewResponseMobileReportStateComplete SpeedTestNewResponseMobileReportState = "COMPLETE"
	SpeedTestNewResponseMobileReportStateFailed   SpeedTestNewResponseMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedTestNewResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedTestNewResponseRegionValue `json:"value"`
	JSON  speedTestNewResponseRegionJSON  `json:"-"`
}

// speedTestNewResponseRegionJSON contains the JSON metadata for the struct
// [SpeedTestNewResponseRegion]
type speedTestNewResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestNewResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedTestNewResponseRegionValue string

const (
	SpeedTestNewResponseRegionValueAsiaEast1           SpeedTestNewResponseRegionValue = "asia-east1"
	SpeedTestNewResponseRegionValueAsiaNortheast1      SpeedTestNewResponseRegionValue = "asia-northeast1"
	SpeedTestNewResponseRegionValueAsiaNortheast2      SpeedTestNewResponseRegionValue = "asia-northeast2"
	SpeedTestNewResponseRegionValueAsiaSouth1          SpeedTestNewResponseRegionValue = "asia-south1"
	SpeedTestNewResponseRegionValueAsiaSoutheast1      SpeedTestNewResponseRegionValue = "asia-southeast1"
	SpeedTestNewResponseRegionValueAustraliaSoutheast1 SpeedTestNewResponseRegionValue = "australia-southeast1"
	SpeedTestNewResponseRegionValueEuropeNorth1        SpeedTestNewResponseRegionValue = "europe-north1"
	SpeedTestNewResponseRegionValueEuropeSouthwest1    SpeedTestNewResponseRegionValue = "europe-southwest1"
	SpeedTestNewResponseRegionValueEuropeWest1         SpeedTestNewResponseRegionValue = "europe-west1"
	SpeedTestNewResponseRegionValueEuropeWest2         SpeedTestNewResponseRegionValue = "europe-west2"
	SpeedTestNewResponseRegionValueEuropeWest3         SpeedTestNewResponseRegionValue = "europe-west3"
	SpeedTestNewResponseRegionValueEuropeWest4         SpeedTestNewResponseRegionValue = "europe-west4"
	SpeedTestNewResponseRegionValueEuropeWest8         SpeedTestNewResponseRegionValue = "europe-west8"
	SpeedTestNewResponseRegionValueEuropeWest9         SpeedTestNewResponseRegionValue = "europe-west9"
	SpeedTestNewResponseRegionValueMeWest1             SpeedTestNewResponseRegionValue = "me-west1"
	SpeedTestNewResponseRegionValueSouthamericaEast1   SpeedTestNewResponseRegionValue = "southamerica-east1"
	SpeedTestNewResponseRegionValueUsCentral1          SpeedTestNewResponseRegionValue = "us-central1"
	SpeedTestNewResponseRegionValueUsEast1             SpeedTestNewResponseRegionValue = "us-east1"
	SpeedTestNewResponseRegionValueUsEast4             SpeedTestNewResponseRegionValue = "us-east4"
	SpeedTestNewResponseRegionValueUsSouth1            SpeedTestNewResponseRegionValue = "us-south1"
	SpeedTestNewResponseRegionValueUsWest1             SpeedTestNewResponseRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedTestNewResponseScheduleFrequency string

const (
	SpeedTestNewResponseScheduleFrequencyDaily  SpeedTestNewResponseScheduleFrequency = "DAILY"
	SpeedTestNewResponseScheduleFrequencyWeekly SpeedTestNewResponseScheduleFrequency = "WEEKLY"
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

type SpeedTestGetResponse struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedTestGetResponseDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedTestGetResponseMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedTestGetResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedTestGetResponseScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                   `json:"url"`
	JSON speedTestGetResponseJSON `json:"-"`
}

// speedTestGetResponseJSON contains the JSON metadata for the struct
// [SpeedTestGetResponse]
type speedTestGetResponseJSON struct {
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

func (r *SpeedTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedTestGetResponseDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedTestGetResponseDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedTestGetResponseDesktopReportError      `json:"error"`
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
	State SpeedTestGetResponseDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                               `json:"tti"`
	JSON speedTestGetResponseDesktopReportJSON `json:"-"`
}

// speedTestGetResponseDesktopReportJSON contains the JSON metadata for the struct
// [SpeedTestGetResponseDesktopReport]
type speedTestGetResponseDesktopReportJSON struct {
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

func (r *SpeedTestGetResponseDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedTestGetResponseDesktopReportDeviceType string

const (
	SpeedTestGetResponseDesktopReportDeviceTypeDesktop SpeedTestGetResponseDesktopReportDeviceType = "DESKTOP"
	SpeedTestGetResponseDesktopReportDeviceTypeMobile  SpeedTestGetResponseDesktopReportDeviceType = "MOBILE"
)

type SpeedTestGetResponseDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedTestGetResponseDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                     `json:"finalDisplayedUrl"`
	JSON              speedTestGetResponseDesktopReportErrorJSON `json:"-"`
}

// speedTestGetResponseDesktopReportErrorJSON contains the JSON metadata for the
// struct [SpeedTestGetResponseDesktopReportError]
type speedTestGetResponseDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedTestGetResponseDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedTestGetResponseDesktopReportErrorCode string

const (
	SpeedTestGetResponseDesktopReportErrorCodeNotReachable      SpeedTestGetResponseDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedTestGetResponseDesktopReportErrorCodeDNSFailure        SpeedTestGetResponseDesktopReportErrorCode = "DNS_FAILURE"
	SpeedTestGetResponseDesktopReportErrorCodeNotHTML           SpeedTestGetResponseDesktopReportErrorCode = "NOT_HTML"
	SpeedTestGetResponseDesktopReportErrorCodeLighthouseTimeout SpeedTestGetResponseDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedTestGetResponseDesktopReportErrorCodeUnknown           SpeedTestGetResponseDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedTestGetResponseDesktopReportState string

const (
	SpeedTestGetResponseDesktopReportStateRunning  SpeedTestGetResponseDesktopReportState = "RUNNING"
	SpeedTestGetResponseDesktopReportStateComplete SpeedTestGetResponseDesktopReportState = "COMPLETE"
	SpeedTestGetResponseDesktopReportStateFailed   SpeedTestGetResponseDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedTestGetResponseMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedTestGetResponseMobileReportDeviceType `json:"deviceType"`
	Error      SpeedTestGetResponseMobileReportError      `json:"error"`
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
	State SpeedTestGetResponseMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                              `json:"tti"`
	JSON speedTestGetResponseMobileReportJSON `json:"-"`
}

// speedTestGetResponseMobileReportJSON contains the JSON metadata for the struct
// [SpeedTestGetResponseMobileReport]
type speedTestGetResponseMobileReportJSON struct {
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

func (r *SpeedTestGetResponseMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedTestGetResponseMobileReportDeviceType string

const (
	SpeedTestGetResponseMobileReportDeviceTypeDesktop SpeedTestGetResponseMobileReportDeviceType = "DESKTOP"
	SpeedTestGetResponseMobileReportDeviceTypeMobile  SpeedTestGetResponseMobileReportDeviceType = "MOBILE"
)

type SpeedTestGetResponseMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedTestGetResponseMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                    `json:"finalDisplayedUrl"`
	JSON              speedTestGetResponseMobileReportErrorJSON `json:"-"`
}

// speedTestGetResponseMobileReportErrorJSON contains the JSON metadata for the
// struct [SpeedTestGetResponseMobileReportError]
type speedTestGetResponseMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedTestGetResponseMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedTestGetResponseMobileReportErrorCode string

const (
	SpeedTestGetResponseMobileReportErrorCodeNotReachable      SpeedTestGetResponseMobileReportErrorCode = "NOT_REACHABLE"
	SpeedTestGetResponseMobileReportErrorCodeDNSFailure        SpeedTestGetResponseMobileReportErrorCode = "DNS_FAILURE"
	SpeedTestGetResponseMobileReportErrorCodeNotHTML           SpeedTestGetResponseMobileReportErrorCode = "NOT_HTML"
	SpeedTestGetResponseMobileReportErrorCodeLighthouseTimeout SpeedTestGetResponseMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedTestGetResponseMobileReportErrorCodeUnknown           SpeedTestGetResponseMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedTestGetResponseMobileReportState string

const (
	SpeedTestGetResponseMobileReportStateRunning  SpeedTestGetResponseMobileReportState = "RUNNING"
	SpeedTestGetResponseMobileReportStateComplete SpeedTestGetResponseMobileReportState = "COMPLETE"
	SpeedTestGetResponseMobileReportStateFailed   SpeedTestGetResponseMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedTestGetResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedTestGetResponseRegionValue `json:"value"`
	JSON  speedTestGetResponseRegionJSON  `json:"-"`
}

// speedTestGetResponseRegionJSON contains the JSON metadata for the struct
// [SpeedTestGetResponseRegion]
type speedTestGetResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestGetResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedTestGetResponseRegionValue string

const (
	SpeedTestGetResponseRegionValueAsiaEast1           SpeedTestGetResponseRegionValue = "asia-east1"
	SpeedTestGetResponseRegionValueAsiaNortheast1      SpeedTestGetResponseRegionValue = "asia-northeast1"
	SpeedTestGetResponseRegionValueAsiaNortheast2      SpeedTestGetResponseRegionValue = "asia-northeast2"
	SpeedTestGetResponseRegionValueAsiaSouth1          SpeedTestGetResponseRegionValue = "asia-south1"
	SpeedTestGetResponseRegionValueAsiaSoutheast1      SpeedTestGetResponseRegionValue = "asia-southeast1"
	SpeedTestGetResponseRegionValueAustraliaSoutheast1 SpeedTestGetResponseRegionValue = "australia-southeast1"
	SpeedTestGetResponseRegionValueEuropeNorth1        SpeedTestGetResponseRegionValue = "europe-north1"
	SpeedTestGetResponseRegionValueEuropeSouthwest1    SpeedTestGetResponseRegionValue = "europe-southwest1"
	SpeedTestGetResponseRegionValueEuropeWest1         SpeedTestGetResponseRegionValue = "europe-west1"
	SpeedTestGetResponseRegionValueEuropeWest2         SpeedTestGetResponseRegionValue = "europe-west2"
	SpeedTestGetResponseRegionValueEuropeWest3         SpeedTestGetResponseRegionValue = "europe-west3"
	SpeedTestGetResponseRegionValueEuropeWest4         SpeedTestGetResponseRegionValue = "europe-west4"
	SpeedTestGetResponseRegionValueEuropeWest8         SpeedTestGetResponseRegionValue = "europe-west8"
	SpeedTestGetResponseRegionValueEuropeWest9         SpeedTestGetResponseRegionValue = "europe-west9"
	SpeedTestGetResponseRegionValueMeWest1             SpeedTestGetResponseRegionValue = "me-west1"
	SpeedTestGetResponseRegionValueSouthamericaEast1   SpeedTestGetResponseRegionValue = "southamerica-east1"
	SpeedTestGetResponseRegionValueUsCentral1          SpeedTestGetResponseRegionValue = "us-central1"
	SpeedTestGetResponseRegionValueUsEast1             SpeedTestGetResponseRegionValue = "us-east1"
	SpeedTestGetResponseRegionValueUsEast4             SpeedTestGetResponseRegionValue = "us-east4"
	SpeedTestGetResponseRegionValueUsSouth1            SpeedTestGetResponseRegionValue = "us-south1"
	SpeedTestGetResponseRegionValueUsWest1             SpeedTestGetResponseRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedTestGetResponseScheduleFrequency string

const (
	SpeedTestGetResponseScheduleFrequencyDaily  SpeedTestGetResponseScheduleFrequency = "DAILY"
	SpeedTestGetResponseScheduleFrequencyWeekly SpeedTestGetResponseScheduleFrequency = "WEEKLY"
)

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
	Result SpeedTestNewResponse             `json:"result"`
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
	Result SpeedTestGetResponse             `json:"result"`
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
