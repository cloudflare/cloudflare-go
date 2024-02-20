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

// SpeedService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewSpeedService] method instead.
type SpeedService struct {
	Options        []option.RequestOption
	Schedule       *SpeedScheduleService
	Availabilities *SpeedAvailabilityService
	Pages          *SpeedPageService
	Tests          *SpeedTestService
}

// NewSpeedService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSpeedService(opts ...option.RequestOption) (r *SpeedService) {
	r = &SpeedService{}
	r.Options = opts
	r.Schedule = NewSpeedScheduleService(opts...)
	r.Availabilities = NewSpeedAvailabilityService(opts...)
	r.Pages = NewSpeedPageService(opts...)
	r.Tests = NewSpeedTestService(opts...)
	return
}

// Starts a test for a specific webpage, in a specific region.
func (r *SpeedService) New(ctx context.Context, zoneID string, url string, body SpeedNewParams, opts ...option.RequestOption) (res *SpeedNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes all tests for a specific webpage from a specific region. Deleted tests
// are still counted as part of the quota.
func (r *SpeedService) Delete(ctx context.Context, zoneID string, url string, body SpeedDeleteParams, opts ...option.RequestOption) (res *SpeedDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the test schedule for a page in a specific region.
func (r *SpeedService) ScheduleGet(ctx context.Context, zoneID string, url string, query SpeedScheduleGetParams, opts ...option.RequestOption) (res *SpeedScheduleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedScheduleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the result of a specific test.
func (r *SpeedService) TestsGet(ctx context.Context, zoneID string, url string, testID string, opts ...option.RequestOption) (res *SpeedTestsGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedTestsGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests/%s", zoneID, url, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the core web vital metrics trend over time for a specific page.
func (r *SpeedService) TrendsList(ctx context.Context, zoneID string, url string, query SpeedTrendsListParams, opts ...option.RequestOption) (res *SpeedTrendsListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedTrendsListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/trend", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SpeedNewResponse struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedNewResponseDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedNewResponseMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedNewResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedNewResponseScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string               `json:"url"`
	JSON speedNewResponseJSON `json:"-"`
}

// speedNewResponseJSON contains the JSON metadata for the struct
// [SpeedNewResponse]
type speedNewResponseJSON struct {
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

func (r *SpeedNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedNewResponseDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedNewResponseDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedNewResponseDesktopReportError      `json:"error"`
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
	State SpeedNewResponseDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                           `json:"tti"`
	JSON speedNewResponseDesktopReportJSON `json:"-"`
}

// speedNewResponseDesktopReportJSON contains the JSON metadata for the struct
// [SpeedNewResponseDesktopReport]
type speedNewResponseDesktopReportJSON struct {
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

func (r *SpeedNewResponseDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedNewResponseDesktopReportDeviceType string

const (
	SpeedNewResponseDesktopReportDeviceTypeDesktop SpeedNewResponseDesktopReportDeviceType = "DESKTOP"
	SpeedNewResponseDesktopReportDeviceTypeMobile  SpeedNewResponseDesktopReportDeviceType = "MOBILE"
)

type SpeedNewResponseDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedNewResponseDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                 `json:"finalDisplayedUrl"`
	JSON              speedNewResponseDesktopReportErrorJSON `json:"-"`
}

// speedNewResponseDesktopReportErrorJSON contains the JSON metadata for the struct
// [SpeedNewResponseDesktopReportError]
type speedNewResponseDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedNewResponseDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedNewResponseDesktopReportErrorCode string

const (
	SpeedNewResponseDesktopReportErrorCodeNotReachable      SpeedNewResponseDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedNewResponseDesktopReportErrorCodeDNSFailure        SpeedNewResponseDesktopReportErrorCode = "DNS_FAILURE"
	SpeedNewResponseDesktopReportErrorCodeNotHTML           SpeedNewResponseDesktopReportErrorCode = "NOT_HTML"
	SpeedNewResponseDesktopReportErrorCodeLighthouseTimeout SpeedNewResponseDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedNewResponseDesktopReportErrorCodeUnknown           SpeedNewResponseDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedNewResponseDesktopReportState string

const (
	SpeedNewResponseDesktopReportStateRunning  SpeedNewResponseDesktopReportState = "RUNNING"
	SpeedNewResponseDesktopReportStateComplete SpeedNewResponseDesktopReportState = "COMPLETE"
	SpeedNewResponseDesktopReportStateFailed   SpeedNewResponseDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedNewResponseMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedNewResponseMobileReportDeviceType `json:"deviceType"`
	Error      SpeedNewResponseMobileReportError      `json:"error"`
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
	State SpeedNewResponseMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                          `json:"tti"`
	JSON speedNewResponseMobileReportJSON `json:"-"`
}

// speedNewResponseMobileReportJSON contains the JSON metadata for the struct
// [SpeedNewResponseMobileReport]
type speedNewResponseMobileReportJSON struct {
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

func (r *SpeedNewResponseMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedNewResponseMobileReportDeviceType string

const (
	SpeedNewResponseMobileReportDeviceTypeDesktop SpeedNewResponseMobileReportDeviceType = "DESKTOP"
	SpeedNewResponseMobileReportDeviceTypeMobile  SpeedNewResponseMobileReportDeviceType = "MOBILE"
)

type SpeedNewResponseMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedNewResponseMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                `json:"finalDisplayedUrl"`
	JSON              speedNewResponseMobileReportErrorJSON `json:"-"`
}

// speedNewResponseMobileReportErrorJSON contains the JSON metadata for the struct
// [SpeedNewResponseMobileReportError]
type speedNewResponseMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedNewResponseMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedNewResponseMobileReportErrorCode string

const (
	SpeedNewResponseMobileReportErrorCodeNotReachable      SpeedNewResponseMobileReportErrorCode = "NOT_REACHABLE"
	SpeedNewResponseMobileReportErrorCodeDNSFailure        SpeedNewResponseMobileReportErrorCode = "DNS_FAILURE"
	SpeedNewResponseMobileReportErrorCodeNotHTML           SpeedNewResponseMobileReportErrorCode = "NOT_HTML"
	SpeedNewResponseMobileReportErrorCodeLighthouseTimeout SpeedNewResponseMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedNewResponseMobileReportErrorCodeUnknown           SpeedNewResponseMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedNewResponseMobileReportState string

const (
	SpeedNewResponseMobileReportStateRunning  SpeedNewResponseMobileReportState = "RUNNING"
	SpeedNewResponseMobileReportStateComplete SpeedNewResponseMobileReportState = "COMPLETE"
	SpeedNewResponseMobileReportStateFailed   SpeedNewResponseMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedNewResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedNewResponseRegionValue `json:"value"`
	JSON  speedNewResponseRegionJSON  `json:"-"`
}

// speedNewResponseRegionJSON contains the JSON metadata for the struct
// [SpeedNewResponseRegion]
type speedNewResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedNewResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedNewResponseRegionValue string

const (
	SpeedNewResponseRegionValueAsiaEast1           SpeedNewResponseRegionValue = "asia-east1"
	SpeedNewResponseRegionValueAsiaNortheast1      SpeedNewResponseRegionValue = "asia-northeast1"
	SpeedNewResponseRegionValueAsiaNortheast2      SpeedNewResponseRegionValue = "asia-northeast2"
	SpeedNewResponseRegionValueAsiaSouth1          SpeedNewResponseRegionValue = "asia-south1"
	SpeedNewResponseRegionValueAsiaSoutheast1      SpeedNewResponseRegionValue = "asia-southeast1"
	SpeedNewResponseRegionValueAustraliaSoutheast1 SpeedNewResponseRegionValue = "australia-southeast1"
	SpeedNewResponseRegionValueEuropeNorth1        SpeedNewResponseRegionValue = "europe-north1"
	SpeedNewResponseRegionValueEuropeSouthwest1    SpeedNewResponseRegionValue = "europe-southwest1"
	SpeedNewResponseRegionValueEuropeWest1         SpeedNewResponseRegionValue = "europe-west1"
	SpeedNewResponseRegionValueEuropeWest2         SpeedNewResponseRegionValue = "europe-west2"
	SpeedNewResponseRegionValueEuropeWest3         SpeedNewResponseRegionValue = "europe-west3"
	SpeedNewResponseRegionValueEuropeWest4         SpeedNewResponseRegionValue = "europe-west4"
	SpeedNewResponseRegionValueEuropeWest8         SpeedNewResponseRegionValue = "europe-west8"
	SpeedNewResponseRegionValueEuropeWest9         SpeedNewResponseRegionValue = "europe-west9"
	SpeedNewResponseRegionValueMeWest1             SpeedNewResponseRegionValue = "me-west1"
	SpeedNewResponseRegionValueSouthamericaEast1   SpeedNewResponseRegionValue = "southamerica-east1"
	SpeedNewResponseRegionValueUsCentral1          SpeedNewResponseRegionValue = "us-central1"
	SpeedNewResponseRegionValueUsEast1             SpeedNewResponseRegionValue = "us-east1"
	SpeedNewResponseRegionValueUsEast4             SpeedNewResponseRegionValue = "us-east4"
	SpeedNewResponseRegionValueUsSouth1            SpeedNewResponseRegionValue = "us-south1"
	SpeedNewResponseRegionValueUsWest1             SpeedNewResponseRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedNewResponseScheduleFrequency string

const (
	SpeedNewResponseScheduleFrequencyDaily  SpeedNewResponseScheduleFrequency = "DAILY"
	SpeedNewResponseScheduleFrequencyWeekly SpeedNewResponseScheduleFrequency = "WEEKLY"
)

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

type SpeedTestsGetResponse struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedTestsGetResponseDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedTestsGetResponseMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedTestsGetResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedTestsGetResponseScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                    `json:"url"`
	JSON speedTestsGetResponseJSON `json:"-"`
}

// speedTestsGetResponseJSON contains the JSON metadata for the struct
// [SpeedTestsGetResponse]
type speedTestsGetResponseJSON struct {
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

func (r *SpeedTestsGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedTestsGetResponseDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedTestsGetResponseDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedTestsGetResponseDesktopReportError      `json:"error"`
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
	State SpeedTestsGetResponseDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                `json:"tti"`
	JSON speedTestsGetResponseDesktopReportJSON `json:"-"`
}

// speedTestsGetResponseDesktopReportJSON contains the JSON metadata for the struct
// [SpeedTestsGetResponseDesktopReport]
type speedTestsGetResponseDesktopReportJSON struct {
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

func (r *SpeedTestsGetResponseDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedTestsGetResponseDesktopReportDeviceType string

const (
	SpeedTestsGetResponseDesktopReportDeviceTypeDesktop SpeedTestsGetResponseDesktopReportDeviceType = "DESKTOP"
	SpeedTestsGetResponseDesktopReportDeviceTypeMobile  SpeedTestsGetResponseDesktopReportDeviceType = "MOBILE"
)

type SpeedTestsGetResponseDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedTestsGetResponseDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                      `json:"finalDisplayedUrl"`
	JSON              speedTestsGetResponseDesktopReportErrorJSON `json:"-"`
}

// speedTestsGetResponseDesktopReportErrorJSON contains the JSON metadata for the
// struct [SpeedTestsGetResponseDesktopReportError]
type speedTestsGetResponseDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedTestsGetResponseDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedTestsGetResponseDesktopReportErrorCode string

const (
	SpeedTestsGetResponseDesktopReportErrorCodeNotReachable      SpeedTestsGetResponseDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedTestsGetResponseDesktopReportErrorCodeDNSFailure        SpeedTestsGetResponseDesktopReportErrorCode = "DNS_FAILURE"
	SpeedTestsGetResponseDesktopReportErrorCodeNotHTML           SpeedTestsGetResponseDesktopReportErrorCode = "NOT_HTML"
	SpeedTestsGetResponseDesktopReportErrorCodeLighthouseTimeout SpeedTestsGetResponseDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedTestsGetResponseDesktopReportErrorCodeUnknown           SpeedTestsGetResponseDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedTestsGetResponseDesktopReportState string

const (
	SpeedTestsGetResponseDesktopReportStateRunning  SpeedTestsGetResponseDesktopReportState = "RUNNING"
	SpeedTestsGetResponseDesktopReportStateComplete SpeedTestsGetResponseDesktopReportState = "COMPLETE"
	SpeedTestsGetResponseDesktopReportStateFailed   SpeedTestsGetResponseDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedTestsGetResponseMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedTestsGetResponseMobileReportDeviceType `json:"deviceType"`
	Error      SpeedTestsGetResponseMobileReportError      `json:"error"`
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
	State SpeedTestsGetResponseMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                               `json:"tti"`
	JSON speedTestsGetResponseMobileReportJSON `json:"-"`
}

// speedTestsGetResponseMobileReportJSON contains the JSON metadata for the struct
// [SpeedTestsGetResponseMobileReport]
type speedTestsGetResponseMobileReportJSON struct {
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

func (r *SpeedTestsGetResponseMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedTestsGetResponseMobileReportDeviceType string

const (
	SpeedTestsGetResponseMobileReportDeviceTypeDesktop SpeedTestsGetResponseMobileReportDeviceType = "DESKTOP"
	SpeedTestsGetResponseMobileReportDeviceTypeMobile  SpeedTestsGetResponseMobileReportDeviceType = "MOBILE"
)

type SpeedTestsGetResponseMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedTestsGetResponseMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                     `json:"finalDisplayedUrl"`
	JSON              speedTestsGetResponseMobileReportErrorJSON `json:"-"`
}

// speedTestsGetResponseMobileReportErrorJSON contains the JSON metadata for the
// struct [SpeedTestsGetResponseMobileReportError]
type speedTestsGetResponseMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedTestsGetResponseMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedTestsGetResponseMobileReportErrorCode string

const (
	SpeedTestsGetResponseMobileReportErrorCodeNotReachable      SpeedTestsGetResponseMobileReportErrorCode = "NOT_REACHABLE"
	SpeedTestsGetResponseMobileReportErrorCodeDNSFailure        SpeedTestsGetResponseMobileReportErrorCode = "DNS_FAILURE"
	SpeedTestsGetResponseMobileReportErrorCodeNotHTML           SpeedTestsGetResponseMobileReportErrorCode = "NOT_HTML"
	SpeedTestsGetResponseMobileReportErrorCodeLighthouseTimeout SpeedTestsGetResponseMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedTestsGetResponseMobileReportErrorCodeUnknown           SpeedTestsGetResponseMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedTestsGetResponseMobileReportState string

const (
	SpeedTestsGetResponseMobileReportStateRunning  SpeedTestsGetResponseMobileReportState = "RUNNING"
	SpeedTestsGetResponseMobileReportStateComplete SpeedTestsGetResponseMobileReportState = "COMPLETE"
	SpeedTestsGetResponseMobileReportStateFailed   SpeedTestsGetResponseMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedTestsGetResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedTestsGetResponseRegionValue `json:"value"`
	JSON  speedTestsGetResponseRegionJSON  `json:"-"`
}

// speedTestsGetResponseRegionJSON contains the JSON metadata for the struct
// [SpeedTestsGetResponseRegion]
type speedTestsGetResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestsGetResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedTestsGetResponseRegionValue string

const (
	SpeedTestsGetResponseRegionValueAsiaEast1           SpeedTestsGetResponseRegionValue = "asia-east1"
	SpeedTestsGetResponseRegionValueAsiaNortheast1      SpeedTestsGetResponseRegionValue = "asia-northeast1"
	SpeedTestsGetResponseRegionValueAsiaNortheast2      SpeedTestsGetResponseRegionValue = "asia-northeast2"
	SpeedTestsGetResponseRegionValueAsiaSouth1          SpeedTestsGetResponseRegionValue = "asia-south1"
	SpeedTestsGetResponseRegionValueAsiaSoutheast1      SpeedTestsGetResponseRegionValue = "asia-southeast1"
	SpeedTestsGetResponseRegionValueAustraliaSoutheast1 SpeedTestsGetResponseRegionValue = "australia-southeast1"
	SpeedTestsGetResponseRegionValueEuropeNorth1        SpeedTestsGetResponseRegionValue = "europe-north1"
	SpeedTestsGetResponseRegionValueEuropeSouthwest1    SpeedTestsGetResponseRegionValue = "europe-southwest1"
	SpeedTestsGetResponseRegionValueEuropeWest1         SpeedTestsGetResponseRegionValue = "europe-west1"
	SpeedTestsGetResponseRegionValueEuropeWest2         SpeedTestsGetResponseRegionValue = "europe-west2"
	SpeedTestsGetResponseRegionValueEuropeWest3         SpeedTestsGetResponseRegionValue = "europe-west3"
	SpeedTestsGetResponseRegionValueEuropeWest4         SpeedTestsGetResponseRegionValue = "europe-west4"
	SpeedTestsGetResponseRegionValueEuropeWest8         SpeedTestsGetResponseRegionValue = "europe-west8"
	SpeedTestsGetResponseRegionValueEuropeWest9         SpeedTestsGetResponseRegionValue = "europe-west9"
	SpeedTestsGetResponseRegionValueMeWest1             SpeedTestsGetResponseRegionValue = "me-west1"
	SpeedTestsGetResponseRegionValueSouthamericaEast1   SpeedTestsGetResponseRegionValue = "southamerica-east1"
	SpeedTestsGetResponseRegionValueUsCentral1          SpeedTestsGetResponseRegionValue = "us-central1"
	SpeedTestsGetResponseRegionValueUsEast1             SpeedTestsGetResponseRegionValue = "us-east1"
	SpeedTestsGetResponseRegionValueUsEast4             SpeedTestsGetResponseRegionValue = "us-east4"
	SpeedTestsGetResponseRegionValueUsSouth1            SpeedTestsGetResponseRegionValue = "us-south1"
	SpeedTestsGetResponseRegionValueUsWest1             SpeedTestsGetResponseRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedTestsGetResponseScheduleFrequency string

const (
	SpeedTestsGetResponseScheduleFrequencyDaily  SpeedTestsGetResponseScheduleFrequency = "DAILY"
	SpeedTestsGetResponseScheduleFrequencyWeekly SpeedTestsGetResponseScheduleFrequency = "WEEKLY"
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

type SpeedNewParams struct {
	// A test region.
	Region param.Field[SpeedNewParamsRegion] `json:"region"`
}

func (r SpeedNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A test region.
type SpeedNewParamsRegion string

const (
	SpeedNewParamsRegionAsiaEast1           SpeedNewParamsRegion = "asia-east1"
	SpeedNewParamsRegionAsiaNortheast1      SpeedNewParamsRegion = "asia-northeast1"
	SpeedNewParamsRegionAsiaNortheast2      SpeedNewParamsRegion = "asia-northeast2"
	SpeedNewParamsRegionAsiaSouth1          SpeedNewParamsRegion = "asia-south1"
	SpeedNewParamsRegionAsiaSoutheast1      SpeedNewParamsRegion = "asia-southeast1"
	SpeedNewParamsRegionAustraliaSoutheast1 SpeedNewParamsRegion = "australia-southeast1"
	SpeedNewParamsRegionEuropeNorth1        SpeedNewParamsRegion = "europe-north1"
	SpeedNewParamsRegionEuropeSouthwest1    SpeedNewParamsRegion = "europe-southwest1"
	SpeedNewParamsRegionEuropeWest1         SpeedNewParamsRegion = "europe-west1"
	SpeedNewParamsRegionEuropeWest2         SpeedNewParamsRegion = "europe-west2"
	SpeedNewParamsRegionEuropeWest3         SpeedNewParamsRegion = "europe-west3"
	SpeedNewParamsRegionEuropeWest4         SpeedNewParamsRegion = "europe-west4"
	SpeedNewParamsRegionEuropeWest8         SpeedNewParamsRegion = "europe-west8"
	SpeedNewParamsRegionEuropeWest9         SpeedNewParamsRegion = "europe-west9"
	SpeedNewParamsRegionMeWest1             SpeedNewParamsRegion = "me-west1"
	SpeedNewParamsRegionSouthamericaEast1   SpeedNewParamsRegion = "southamerica-east1"
	SpeedNewParamsRegionUsCentral1          SpeedNewParamsRegion = "us-central1"
	SpeedNewParamsRegionUsEast1             SpeedNewParamsRegion = "us-east1"
	SpeedNewParamsRegionUsEast4             SpeedNewParamsRegion = "us-east4"
	SpeedNewParamsRegionUsSouth1            SpeedNewParamsRegion = "us-south1"
	SpeedNewParamsRegionUsWest1             SpeedNewParamsRegion = "us-west1"
)

type SpeedNewResponseEnvelope struct {
	Result SpeedNewResponse             `json:"result"`
	JSON   speedNewResponseEnvelopeJSON `json:"-"`
}

// speedNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedNewResponseEnvelope]
type speedNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedDeleteParams struct {
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

type SpeedScheduleGetParams struct {
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

type SpeedTestsGetResponseEnvelope struct {
	Result SpeedTestsGetResponse             `json:"result"`
	JSON   speedTestsGetResponseEnvelopeJSON `json:"-"`
}

// speedTestsGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedTestsGetResponseEnvelope]
type speedTestsGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestsGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedTrendsListParams struct {
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
