// File generated from our OpenAPI spec by Stainless.

package speed

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// TestService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewTestService] method instead.
type TestService struct {
	Options []option.RequestOption
}

// NewTestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTestService(opts ...option.RequestOption) (r *TestService) {
	r = &TestService{}
	r.Options = opts
	return
}

// Starts a test for a specific webpage, in a specific region.
func (r *TestService) New(ctx context.Context, url string, params TestNewParams, opts ...option.RequestOption) (res *TestNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TestNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Test history (list of tests) for a specific webpage.
func (r *TestService) List(ctx context.Context, url string, params TestListParams, opts ...option.RequestOption) (res *TestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Deletes all tests for a specific webpage from a specific region. Deleted tests
// are still counted as part of the quota.
func (r *TestService) Delete(ctx context.Context, url string, params TestDeleteParams, opts ...option.RequestOption) (res *TestDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TestDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the result of a specific test.
func (r *TestService) Get(ctx context.Context, url string, testID string, query TestGetParams, opts ...option.RequestOption) (res *TestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TestGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests/%s", query.ZoneID, url, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TestNewResponse struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport TestNewResponseDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport TestNewResponseMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region TestNewResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency TestNewResponseScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string              `json:"url"`
	JSON testNewResponseJSON `json:"-"`
}

// testNewResponseJSON contains the JSON metadata for the struct [TestNewResponse]
type testNewResponseJSON struct {
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

func (r *TestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testNewResponseJSON) RawJSON() string {
	return r.raw
}

// The Lighthouse report.
type TestNewResponseDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType TestNewResponseDesktopReportDeviceType `json:"deviceType"`
	Error      TestNewResponseDesktopReportError      `json:"error"`
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
	State TestNewResponseDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                          `json:"tti"`
	JSON testNewResponseDesktopReportJSON `json:"-"`
}

// testNewResponseDesktopReportJSON contains the JSON metadata for the struct
// [TestNewResponseDesktopReport]
type testNewResponseDesktopReportJSON struct {
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

func (r *TestNewResponseDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testNewResponseDesktopReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type TestNewResponseDesktopReportDeviceType string

const (
	TestNewResponseDesktopReportDeviceTypeDesktop TestNewResponseDesktopReportDeviceType = "DESKTOP"
	TestNewResponseDesktopReportDeviceTypeMobile  TestNewResponseDesktopReportDeviceType = "MOBILE"
)

type TestNewResponseDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code TestNewResponseDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                `json:"finalDisplayedUrl"`
	JSON              testNewResponseDesktopReportErrorJSON `json:"-"`
}

// testNewResponseDesktopReportErrorJSON contains the JSON metadata for the struct
// [TestNewResponseDesktopReportError]
type testNewResponseDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TestNewResponseDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testNewResponseDesktopReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type TestNewResponseDesktopReportErrorCode string

const (
	TestNewResponseDesktopReportErrorCodeNotReachable      TestNewResponseDesktopReportErrorCode = "NOT_REACHABLE"
	TestNewResponseDesktopReportErrorCodeDNSFailure        TestNewResponseDesktopReportErrorCode = "DNS_FAILURE"
	TestNewResponseDesktopReportErrorCodeNotHTML           TestNewResponseDesktopReportErrorCode = "NOT_HTML"
	TestNewResponseDesktopReportErrorCodeLighthouseTimeout TestNewResponseDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	TestNewResponseDesktopReportErrorCodeUnknown           TestNewResponseDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type TestNewResponseDesktopReportState string

const (
	TestNewResponseDesktopReportStateRunning  TestNewResponseDesktopReportState = "RUNNING"
	TestNewResponseDesktopReportStateComplete TestNewResponseDesktopReportState = "COMPLETE"
	TestNewResponseDesktopReportStateFailed   TestNewResponseDesktopReportState = "FAILED"
)

// The Lighthouse report.
type TestNewResponseMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType TestNewResponseMobileReportDeviceType `json:"deviceType"`
	Error      TestNewResponseMobileReportError      `json:"error"`
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
	State TestNewResponseMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                         `json:"tti"`
	JSON testNewResponseMobileReportJSON `json:"-"`
}

// testNewResponseMobileReportJSON contains the JSON metadata for the struct
// [TestNewResponseMobileReport]
type testNewResponseMobileReportJSON struct {
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

func (r *TestNewResponseMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testNewResponseMobileReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type TestNewResponseMobileReportDeviceType string

const (
	TestNewResponseMobileReportDeviceTypeDesktop TestNewResponseMobileReportDeviceType = "DESKTOP"
	TestNewResponseMobileReportDeviceTypeMobile  TestNewResponseMobileReportDeviceType = "MOBILE"
)

type TestNewResponseMobileReportError struct {
	// The error code of the Lighthouse result.
	Code TestNewResponseMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                               `json:"finalDisplayedUrl"`
	JSON              testNewResponseMobileReportErrorJSON `json:"-"`
}

// testNewResponseMobileReportErrorJSON contains the JSON metadata for the struct
// [TestNewResponseMobileReportError]
type testNewResponseMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TestNewResponseMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testNewResponseMobileReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type TestNewResponseMobileReportErrorCode string

const (
	TestNewResponseMobileReportErrorCodeNotReachable      TestNewResponseMobileReportErrorCode = "NOT_REACHABLE"
	TestNewResponseMobileReportErrorCodeDNSFailure        TestNewResponseMobileReportErrorCode = "DNS_FAILURE"
	TestNewResponseMobileReportErrorCodeNotHTML           TestNewResponseMobileReportErrorCode = "NOT_HTML"
	TestNewResponseMobileReportErrorCodeLighthouseTimeout TestNewResponseMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	TestNewResponseMobileReportErrorCodeUnknown           TestNewResponseMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type TestNewResponseMobileReportState string

const (
	TestNewResponseMobileReportStateRunning  TestNewResponseMobileReportState = "RUNNING"
	TestNewResponseMobileReportStateComplete TestNewResponseMobileReportState = "COMPLETE"
	TestNewResponseMobileReportStateFailed   TestNewResponseMobileReportState = "FAILED"
)

// A test region with a label.
type TestNewResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value TestNewResponseRegionValue `json:"value"`
	JSON  testNewResponseRegionJSON  `json:"-"`
}

// testNewResponseRegionJSON contains the JSON metadata for the struct
// [TestNewResponseRegion]
type testNewResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestNewResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testNewResponseRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type TestNewResponseRegionValue string

const (
	TestNewResponseRegionValueAsiaEast1           TestNewResponseRegionValue = "asia-east1"
	TestNewResponseRegionValueAsiaNortheast1      TestNewResponseRegionValue = "asia-northeast1"
	TestNewResponseRegionValueAsiaNortheast2      TestNewResponseRegionValue = "asia-northeast2"
	TestNewResponseRegionValueAsiaSouth1          TestNewResponseRegionValue = "asia-south1"
	TestNewResponseRegionValueAsiaSoutheast1      TestNewResponseRegionValue = "asia-southeast1"
	TestNewResponseRegionValueAustraliaSoutheast1 TestNewResponseRegionValue = "australia-southeast1"
	TestNewResponseRegionValueEuropeNorth1        TestNewResponseRegionValue = "europe-north1"
	TestNewResponseRegionValueEuropeSouthwest1    TestNewResponseRegionValue = "europe-southwest1"
	TestNewResponseRegionValueEuropeWest1         TestNewResponseRegionValue = "europe-west1"
	TestNewResponseRegionValueEuropeWest2         TestNewResponseRegionValue = "europe-west2"
	TestNewResponseRegionValueEuropeWest3         TestNewResponseRegionValue = "europe-west3"
	TestNewResponseRegionValueEuropeWest4         TestNewResponseRegionValue = "europe-west4"
	TestNewResponseRegionValueEuropeWest8         TestNewResponseRegionValue = "europe-west8"
	TestNewResponseRegionValueEuropeWest9         TestNewResponseRegionValue = "europe-west9"
	TestNewResponseRegionValueMeWest1             TestNewResponseRegionValue = "me-west1"
	TestNewResponseRegionValueSouthamericaEast1   TestNewResponseRegionValue = "southamerica-east1"
	TestNewResponseRegionValueUsCentral1          TestNewResponseRegionValue = "us-central1"
	TestNewResponseRegionValueUsEast1             TestNewResponseRegionValue = "us-east1"
	TestNewResponseRegionValueUsEast4             TestNewResponseRegionValue = "us-east4"
	TestNewResponseRegionValueUsSouth1            TestNewResponseRegionValue = "us-south1"
	TestNewResponseRegionValueUsWest1             TestNewResponseRegionValue = "us-west1"
)

// The frequency of the test.
type TestNewResponseScheduleFrequency string

const (
	TestNewResponseScheduleFrequencyDaily  TestNewResponseScheduleFrequency = "DAILY"
	TestNewResponseScheduleFrequencyWeekly TestNewResponseScheduleFrequency = "WEEKLY"
)

type TestListResponse struct {
	Errors   []TestListResponseError   `json:"errors,required"`
	Messages []TestListResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    bool                       `json:"success,required"`
	ResultInfo TestListResponseResultInfo `json:"result_info"`
	JSON       testListResponseJSON       `json:"-"`
}

// testListResponseJSON contains the JSON metadata for the struct
// [TestListResponse]
type testListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResponseJSON) RawJSON() string {
	return r.raw
}

type TestListResponseError struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    testListResponseErrorJSON `json:"-"`
}

// testListResponseErrorJSON contains the JSON metadata for the struct
// [TestListResponseError]
type testListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type TestListResponseMessage struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    testListResponseMessageJSON `json:"-"`
}

// testListResponseMessageJSON contains the JSON metadata for the struct
// [TestListResponseMessage]
type testListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type TestListResponseResultInfo struct {
	Count      int64                          `json:"count"`
	Page       int64                          `json:"page"`
	PerPage    int64                          `json:"per_page"`
	TotalCount int64                          `json:"total_count"`
	JSON       testListResponseResultInfoJSON `json:"-"`
}

// testListResponseResultInfoJSON contains the JSON metadata for the struct
// [TestListResponseResultInfo]
type testListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type TestDeleteResponse struct {
	// Number of items affected.
	Count float64                `json:"count"`
	JSON  testDeleteResponseJSON `json:"-"`
}

// testDeleteResponseJSON contains the JSON metadata for the struct
// [TestDeleteResponse]
type testDeleteResponseJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TestGetResponse struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport TestGetResponseDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport TestGetResponseMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region TestGetResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency TestGetResponseScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string              `json:"url"`
	JSON testGetResponseJSON `json:"-"`
}

// testGetResponseJSON contains the JSON metadata for the struct [TestGetResponse]
type testGetResponseJSON struct {
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

func (r *TestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testGetResponseJSON) RawJSON() string {
	return r.raw
}

// The Lighthouse report.
type TestGetResponseDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType TestGetResponseDesktopReportDeviceType `json:"deviceType"`
	Error      TestGetResponseDesktopReportError      `json:"error"`
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
	State TestGetResponseDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                          `json:"tti"`
	JSON testGetResponseDesktopReportJSON `json:"-"`
}

// testGetResponseDesktopReportJSON contains the JSON metadata for the struct
// [TestGetResponseDesktopReport]
type testGetResponseDesktopReportJSON struct {
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

func (r *TestGetResponseDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testGetResponseDesktopReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type TestGetResponseDesktopReportDeviceType string

const (
	TestGetResponseDesktopReportDeviceTypeDesktop TestGetResponseDesktopReportDeviceType = "DESKTOP"
	TestGetResponseDesktopReportDeviceTypeMobile  TestGetResponseDesktopReportDeviceType = "MOBILE"
)

type TestGetResponseDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code TestGetResponseDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                `json:"finalDisplayedUrl"`
	JSON              testGetResponseDesktopReportErrorJSON `json:"-"`
}

// testGetResponseDesktopReportErrorJSON contains the JSON metadata for the struct
// [TestGetResponseDesktopReportError]
type testGetResponseDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TestGetResponseDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testGetResponseDesktopReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type TestGetResponseDesktopReportErrorCode string

const (
	TestGetResponseDesktopReportErrorCodeNotReachable      TestGetResponseDesktopReportErrorCode = "NOT_REACHABLE"
	TestGetResponseDesktopReportErrorCodeDNSFailure        TestGetResponseDesktopReportErrorCode = "DNS_FAILURE"
	TestGetResponseDesktopReportErrorCodeNotHTML           TestGetResponseDesktopReportErrorCode = "NOT_HTML"
	TestGetResponseDesktopReportErrorCodeLighthouseTimeout TestGetResponseDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	TestGetResponseDesktopReportErrorCodeUnknown           TestGetResponseDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type TestGetResponseDesktopReportState string

const (
	TestGetResponseDesktopReportStateRunning  TestGetResponseDesktopReportState = "RUNNING"
	TestGetResponseDesktopReportStateComplete TestGetResponseDesktopReportState = "COMPLETE"
	TestGetResponseDesktopReportStateFailed   TestGetResponseDesktopReportState = "FAILED"
)

// The Lighthouse report.
type TestGetResponseMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType TestGetResponseMobileReportDeviceType `json:"deviceType"`
	Error      TestGetResponseMobileReportError      `json:"error"`
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
	State TestGetResponseMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                         `json:"tti"`
	JSON testGetResponseMobileReportJSON `json:"-"`
}

// testGetResponseMobileReportJSON contains the JSON metadata for the struct
// [TestGetResponseMobileReport]
type testGetResponseMobileReportJSON struct {
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

func (r *TestGetResponseMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testGetResponseMobileReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type TestGetResponseMobileReportDeviceType string

const (
	TestGetResponseMobileReportDeviceTypeDesktop TestGetResponseMobileReportDeviceType = "DESKTOP"
	TestGetResponseMobileReportDeviceTypeMobile  TestGetResponseMobileReportDeviceType = "MOBILE"
)

type TestGetResponseMobileReportError struct {
	// The error code of the Lighthouse result.
	Code TestGetResponseMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                               `json:"finalDisplayedUrl"`
	JSON              testGetResponseMobileReportErrorJSON `json:"-"`
}

// testGetResponseMobileReportErrorJSON contains the JSON metadata for the struct
// [TestGetResponseMobileReportError]
type testGetResponseMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TestGetResponseMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testGetResponseMobileReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type TestGetResponseMobileReportErrorCode string

const (
	TestGetResponseMobileReportErrorCodeNotReachable      TestGetResponseMobileReportErrorCode = "NOT_REACHABLE"
	TestGetResponseMobileReportErrorCodeDNSFailure        TestGetResponseMobileReportErrorCode = "DNS_FAILURE"
	TestGetResponseMobileReportErrorCodeNotHTML           TestGetResponseMobileReportErrorCode = "NOT_HTML"
	TestGetResponseMobileReportErrorCodeLighthouseTimeout TestGetResponseMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	TestGetResponseMobileReportErrorCodeUnknown           TestGetResponseMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type TestGetResponseMobileReportState string

const (
	TestGetResponseMobileReportStateRunning  TestGetResponseMobileReportState = "RUNNING"
	TestGetResponseMobileReportStateComplete TestGetResponseMobileReportState = "COMPLETE"
	TestGetResponseMobileReportStateFailed   TestGetResponseMobileReportState = "FAILED"
)

// A test region with a label.
type TestGetResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value TestGetResponseRegionValue `json:"value"`
	JSON  testGetResponseRegionJSON  `json:"-"`
}

// testGetResponseRegionJSON contains the JSON metadata for the struct
// [TestGetResponseRegion]
type testGetResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestGetResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testGetResponseRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type TestGetResponseRegionValue string

const (
	TestGetResponseRegionValueAsiaEast1           TestGetResponseRegionValue = "asia-east1"
	TestGetResponseRegionValueAsiaNortheast1      TestGetResponseRegionValue = "asia-northeast1"
	TestGetResponseRegionValueAsiaNortheast2      TestGetResponseRegionValue = "asia-northeast2"
	TestGetResponseRegionValueAsiaSouth1          TestGetResponseRegionValue = "asia-south1"
	TestGetResponseRegionValueAsiaSoutheast1      TestGetResponseRegionValue = "asia-southeast1"
	TestGetResponseRegionValueAustraliaSoutheast1 TestGetResponseRegionValue = "australia-southeast1"
	TestGetResponseRegionValueEuropeNorth1        TestGetResponseRegionValue = "europe-north1"
	TestGetResponseRegionValueEuropeSouthwest1    TestGetResponseRegionValue = "europe-southwest1"
	TestGetResponseRegionValueEuropeWest1         TestGetResponseRegionValue = "europe-west1"
	TestGetResponseRegionValueEuropeWest2         TestGetResponseRegionValue = "europe-west2"
	TestGetResponseRegionValueEuropeWest3         TestGetResponseRegionValue = "europe-west3"
	TestGetResponseRegionValueEuropeWest4         TestGetResponseRegionValue = "europe-west4"
	TestGetResponseRegionValueEuropeWest8         TestGetResponseRegionValue = "europe-west8"
	TestGetResponseRegionValueEuropeWest9         TestGetResponseRegionValue = "europe-west9"
	TestGetResponseRegionValueMeWest1             TestGetResponseRegionValue = "me-west1"
	TestGetResponseRegionValueSouthamericaEast1   TestGetResponseRegionValue = "southamerica-east1"
	TestGetResponseRegionValueUsCentral1          TestGetResponseRegionValue = "us-central1"
	TestGetResponseRegionValueUsEast1             TestGetResponseRegionValue = "us-east1"
	TestGetResponseRegionValueUsEast4             TestGetResponseRegionValue = "us-east4"
	TestGetResponseRegionValueUsSouth1            TestGetResponseRegionValue = "us-south1"
	TestGetResponseRegionValueUsWest1             TestGetResponseRegionValue = "us-west1"
)

// The frequency of the test.
type TestGetResponseScheduleFrequency string

const (
	TestGetResponseScheduleFrequencyDaily  TestGetResponseScheduleFrequency = "DAILY"
	TestGetResponseScheduleFrequencyWeekly TestGetResponseScheduleFrequency = "WEEKLY"
)

type TestNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[TestNewParamsRegion] `json:"region"`
}

func (r TestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A test region.
type TestNewParamsRegion string

const (
	TestNewParamsRegionAsiaEast1           TestNewParamsRegion = "asia-east1"
	TestNewParamsRegionAsiaNortheast1      TestNewParamsRegion = "asia-northeast1"
	TestNewParamsRegionAsiaNortheast2      TestNewParamsRegion = "asia-northeast2"
	TestNewParamsRegionAsiaSouth1          TestNewParamsRegion = "asia-south1"
	TestNewParamsRegionAsiaSoutheast1      TestNewParamsRegion = "asia-southeast1"
	TestNewParamsRegionAustraliaSoutheast1 TestNewParamsRegion = "australia-southeast1"
	TestNewParamsRegionEuropeNorth1        TestNewParamsRegion = "europe-north1"
	TestNewParamsRegionEuropeSouthwest1    TestNewParamsRegion = "europe-southwest1"
	TestNewParamsRegionEuropeWest1         TestNewParamsRegion = "europe-west1"
	TestNewParamsRegionEuropeWest2         TestNewParamsRegion = "europe-west2"
	TestNewParamsRegionEuropeWest3         TestNewParamsRegion = "europe-west3"
	TestNewParamsRegionEuropeWest4         TestNewParamsRegion = "europe-west4"
	TestNewParamsRegionEuropeWest8         TestNewParamsRegion = "europe-west8"
	TestNewParamsRegionEuropeWest9         TestNewParamsRegion = "europe-west9"
	TestNewParamsRegionMeWest1             TestNewParamsRegion = "me-west1"
	TestNewParamsRegionSouthamericaEast1   TestNewParamsRegion = "southamerica-east1"
	TestNewParamsRegionUsCentral1          TestNewParamsRegion = "us-central1"
	TestNewParamsRegionUsEast1             TestNewParamsRegion = "us-east1"
	TestNewParamsRegionUsEast4             TestNewParamsRegion = "us-east4"
	TestNewParamsRegionUsSouth1            TestNewParamsRegion = "us-south1"
	TestNewParamsRegionUsWest1             TestNewParamsRegion = "us-west1"
)

type TestNewResponseEnvelope struct {
	Result TestNewResponse             `json:"result"`
	JSON   testNewResponseEnvelopeJSON `json:"-"`
}

// testNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TestNewResponseEnvelope]
type testNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TestListParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Page    param.Field[int64]  `query:"page"`
	PerPage param.Field[int64]  `query:"per_page"`
	// A test region.
	Region param.Field[TestListParamsRegion] `query:"region"`
}

// URLQuery serializes [TestListParams]'s query parameters as `url.Values`.
func (r TestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type TestListParamsRegion string

const (
	TestListParamsRegionAsiaEast1           TestListParamsRegion = "asia-east1"
	TestListParamsRegionAsiaNortheast1      TestListParamsRegion = "asia-northeast1"
	TestListParamsRegionAsiaNortheast2      TestListParamsRegion = "asia-northeast2"
	TestListParamsRegionAsiaSouth1          TestListParamsRegion = "asia-south1"
	TestListParamsRegionAsiaSoutheast1      TestListParamsRegion = "asia-southeast1"
	TestListParamsRegionAustraliaSoutheast1 TestListParamsRegion = "australia-southeast1"
	TestListParamsRegionEuropeNorth1        TestListParamsRegion = "europe-north1"
	TestListParamsRegionEuropeSouthwest1    TestListParamsRegion = "europe-southwest1"
	TestListParamsRegionEuropeWest1         TestListParamsRegion = "europe-west1"
	TestListParamsRegionEuropeWest2         TestListParamsRegion = "europe-west2"
	TestListParamsRegionEuropeWest3         TestListParamsRegion = "europe-west3"
	TestListParamsRegionEuropeWest4         TestListParamsRegion = "europe-west4"
	TestListParamsRegionEuropeWest8         TestListParamsRegion = "europe-west8"
	TestListParamsRegionEuropeWest9         TestListParamsRegion = "europe-west9"
	TestListParamsRegionMeWest1             TestListParamsRegion = "me-west1"
	TestListParamsRegionSouthamericaEast1   TestListParamsRegion = "southamerica-east1"
	TestListParamsRegionUsCentral1          TestListParamsRegion = "us-central1"
	TestListParamsRegionUsEast1             TestListParamsRegion = "us-east1"
	TestListParamsRegionUsEast4             TestListParamsRegion = "us-east4"
	TestListParamsRegionUsSouth1            TestListParamsRegion = "us-south1"
	TestListParamsRegionUsWest1             TestListParamsRegion = "us-west1"
)

type TestDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[TestDeleteParamsRegion] `query:"region"`
}

// URLQuery serializes [TestDeleteParams]'s query parameters as `url.Values`.
func (r TestDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type TestDeleteParamsRegion string

const (
	TestDeleteParamsRegionAsiaEast1           TestDeleteParamsRegion = "asia-east1"
	TestDeleteParamsRegionAsiaNortheast1      TestDeleteParamsRegion = "asia-northeast1"
	TestDeleteParamsRegionAsiaNortheast2      TestDeleteParamsRegion = "asia-northeast2"
	TestDeleteParamsRegionAsiaSouth1          TestDeleteParamsRegion = "asia-south1"
	TestDeleteParamsRegionAsiaSoutheast1      TestDeleteParamsRegion = "asia-southeast1"
	TestDeleteParamsRegionAustraliaSoutheast1 TestDeleteParamsRegion = "australia-southeast1"
	TestDeleteParamsRegionEuropeNorth1        TestDeleteParamsRegion = "europe-north1"
	TestDeleteParamsRegionEuropeSouthwest1    TestDeleteParamsRegion = "europe-southwest1"
	TestDeleteParamsRegionEuropeWest1         TestDeleteParamsRegion = "europe-west1"
	TestDeleteParamsRegionEuropeWest2         TestDeleteParamsRegion = "europe-west2"
	TestDeleteParamsRegionEuropeWest3         TestDeleteParamsRegion = "europe-west3"
	TestDeleteParamsRegionEuropeWest4         TestDeleteParamsRegion = "europe-west4"
	TestDeleteParamsRegionEuropeWest8         TestDeleteParamsRegion = "europe-west8"
	TestDeleteParamsRegionEuropeWest9         TestDeleteParamsRegion = "europe-west9"
	TestDeleteParamsRegionMeWest1             TestDeleteParamsRegion = "me-west1"
	TestDeleteParamsRegionSouthamericaEast1   TestDeleteParamsRegion = "southamerica-east1"
	TestDeleteParamsRegionUsCentral1          TestDeleteParamsRegion = "us-central1"
	TestDeleteParamsRegionUsEast1             TestDeleteParamsRegion = "us-east1"
	TestDeleteParamsRegionUsEast4             TestDeleteParamsRegion = "us-east4"
	TestDeleteParamsRegionUsSouth1            TestDeleteParamsRegion = "us-south1"
	TestDeleteParamsRegionUsWest1             TestDeleteParamsRegion = "us-west1"
)

type TestDeleteResponseEnvelope struct {
	Result TestDeleteResponse             `json:"result"`
	JSON   testDeleteResponseEnvelopeJSON `json:"-"`
}

// testDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TestDeleteResponseEnvelope]
type testDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TestGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type TestGetResponseEnvelope struct {
	Result TestGetResponse             `json:"result"`
	JSON   testGetResponseEnvelopeJSON `json:"-"`
}

// testGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TestGetResponseEnvelope]
type testGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
