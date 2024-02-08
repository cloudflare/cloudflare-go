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

// SpeedAPIService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSpeedAPIService] method instead.
type SpeedAPIService struct {
	Options  []option.RequestOption
	Schedule *SpeedAPIScheduleService
}

// NewSpeedAPIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpeedAPIService(opts ...option.RequestOption) (r *SpeedAPIService) {
	r = &SpeedAPIService{}
	r.Options = opts
	r.Schedule = NewSpeedAPIScheduleService(opts...)
	return
}

// Retrieves quota for all plans, as well as the current zone quota.
func (r *SpeedAPIService) AvailabilitiesList(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SpeedAPIAvailabilitiesListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedAPIAvailabilitiesListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/availabilities", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all webpages which have been tested.
func (r *SpeedAPIService) PagesList(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *[]SpeedAPIPagesListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedAPIPagesListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a scheduled test for a page.
func (r *SpeedAPIService) ScheduleDelete(ctx context.Context, zoneID string, url string, body SpeedAPIScheduleDeleteParams, opts ...option.RequestOption) (res *SpeedAPIScheduleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedAPIScheduleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the test schedule for a page in a specific region.
func (r *SpeedAPIService) ScheduleGet(ctx context.Context, zoneID string, url string, query SpeedAPIScheduleGetParams, opts ...option.RequestOption) (res *SpeedAPIScheduleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedAPIScheduleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Starts a test for a specific webpage, in a specific region.
func (r *SpeedAPIService) TestsNew(ctx context.Context, zoneID string, url string, body SpeedAPITestsNewParams, opts ...option.RequestOption) (res *SpeedAPITestsNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedAPITestsNewResponseEnvelope
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
func (r *SpeedAPIService) TestsDelete(ctx context.Context, zoneID string, url string, body SpeedAPITestsDeleteParams, opts ...option.RequestOption) (res *SpeedAPITestsDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedAPITestsDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Test history (list of tests) for a specific webpage.
func (r *SpeedAPIService) TestsList(ctx context.Context, zoneID string, url string, query SpeedAPITestsListParams, opts ...option.RequestOption) (res *SpeedAPITestsListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the result of a specific test.
func (r *SpeedAPIService) TestsGet(ctx context.Context, zoneID string, url string, testID string, opts ...option.RequestOption) (res *SpeedAPITestsGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedAPITestsGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests/%s", zoneID, url, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the core web vital metrics trend over time for a specific page.
func (r *SpeedAPIService) TrendsList(ctx context.Context, zoneID string, url string, query SpeedAPITrendsListParams, opts ...option.RequestOption) (res *SpeedAPITrendsListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedAPITrendsListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/trend", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SpeedAPIAvailabilitiesListResponse struct {
	Quota          SpeedAPIAvailabilitiesListResponseQuota    `json:"quota"`
	Regions        []SpeedAPIAvailabilitiesListResponseRegion `json:"regions"`
	RegionsPerPlan interface{}                                `json:"regionsPerPlan"`
	JSON           speedAPIAvailabilitiesListResponseJSON     `json:"-"`
}

// speedAPIAvailabilitiesListResponseJSON contains the JSON metadata for the struct
// [SpeedAPIAvailabilitiesListResponse]
type speedAPIAvailabilitiesListResponseJSON struct {
	Quota          apijson.Field
	Regions        apijson.Field
	RegionsPerPlan apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SpeedAPIAvailabilitiesListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIAvailabilitiesListResponseQuota struct {
	// Cloudflare plan.
	Plan string `json:"plan"`
	// The number of tests available per plan.
	QuotasPerPlan map[string]float64 `json:"quotasPerPlan"`
	// The number of remaining schedules available.
	RemainingSchedules float64 `json:"remainingSchedules"`
	// The number of remaining tests available.
	RemainingTests float64 `json:"remainingTests"`
	// The number of schedules available per plan.
	ScheduleQuotasPerPlan map[string]float64                          `json:"scheduleQuotasPerPlan"`
	JSON                  speedAPIAvailabilitiesListResponseQuotaJSON `json:"-"`
}

// speedAPIAvailabilitiesListResponseQuotaJSON contains the JSON metadata for the
// struct [SpeedAPIAvailabilitiesListResponseQuota]
type speedAPIAvailabilitiesListResponseQuotaJSON struct {
	Plan                  apijson.Field
	QuotasPerPlan         apijson.Field
	RemainingSchedules    apijson.Field
	RemainingTests        apijson.Field
	ScheduleQuotasPerPlan apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SpeedAPIAvailabilitiesListResponseQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region with a label.
type SpeedAPIAvailabilitiesListResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPIAvailabilitiesListResponseRegionsValue `json:"value"`
	JSON  speedAPIAvailabilitiesListResponseRegionJSON   `json:"-"`
}

// speedAPIAvailabilitiesListResponseRegionJSON contains the JSON metadata for the
// struct [SpeedAPIAvailabilitiesListResponseRegion]
type speedAPIAvailabilitiesListResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIAvailabilitiesListResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPIAvailabilitiesListResponseRegionsValue string

const (
	SpeedAPIAvailabilitiesListResponseRegionsValueAsiaEast1           SpeedAPIAvailabilitiesListResponseRegionsValue = "asia-east1"
	SpeedAPIAvailabilitiesListResponseRegionsValueAsiaNortheast1      SpeedAPIAvailabilitiesListResponseRegionsValue = "asia-northeast1"
	SpeedAPIAvailabilitiesListResponseRegionsValueAsiaNortheast2      SpeedAPIAvailabilitiesListResponseRegionsValue = "asia-northeast2"
	SpeedAPIAvailabilitiesListResponseRegionsValueAsiaSouth1          SpeedAPIAvailabilitiesListResponseRegionsValue = "asia-south1"
	SpeedAPIAvailabilitiesListResponseRegionsValueAsiaSoutheast1      SpeedAPIAvailabilitiesListResponseRegionsValue = "asia-southeast1"
	SpeedAPIAvailabilitiesListResponseRegionsValueAustraliaSoutheast1 SpeedAPIAvailabilitiesListResponseRegionsValue = "australia-southeast1"
	SpeedAPIAvailabilitiesListResponseRegionsValueEuropeNorth1        SpeedAPIAvailabilitiesListResponseRegionsValue = "europe-north1"
	SpeedAPIAvailabilitiesListResponseRegionsValueEuropeSouthwest1    SpeedAPIAvailabilitiesListResponseRegionsValue = "europe-southwest1"
	SpeedAPIAvailabilitiesListResponseRegionsValueEuropeWest1         SpeedAPIAvailabilitiesListResponseRegionsValue = "europe-west1"
	SpeedAPIAvailabilitiesListResponseRegionsValueEuropeWest2         SpeedAPIAvailabilitiesListResponseRegionsValue = "europe-west2"
	SpeedAPIAvailabilitiesListResponseRegionsValueEuropeWest3         SpeedAPIAvailabilitiesListResponseRegionsValue = "europe-west3"
	SpeedAPIAvailabilitiesListResponseRegionsValueEuropeWest4         SpeedAPIAvailabilitiesListResponseRegionsValue = "europe-west4"
	SpeedAPIAvailabilitiesListResponseRegionsValueEuropeWest8         SpeedAPIAvailabilitiesListResponseRegionsValue = "europe-west8"
	SpeedAPIAvailabilitiesListResponseRegionsValueEuropeWest9         SpeedAPIAvailabilitiesListResponseRegionsValue = "europe-west9"
	SpeedAPIAvailabilitiesListResponseRegionsValueMeWest1             SpeedAPIAvailabilitiesListResponseRegionsValue = "me-west1"
	SpeedAPIAvailabilitiesListResponseRegionsValueSouthamericaEast1   SpeedAPIAvailabilitiesListResponseRegionsValue = "southamerica-east1"
	SpeedAPIAvailabilitiesListResponseRegionsValueUsCentral1          SpeedAPIAvailabilitiesListResponseRegionsValue = "us-central1"
	SpeedAPIAvailabilitiesListResponseRegionsValueUsEast1             SpeedAPIAvailabilitiesListResponseRegionsValue = "us-east1"
	SpeedAPIAvailabilitiesListResponseRegionsValueUsEast4             SpeedAPIAvailabilitiesListResponseRegionsValue = "us-east4"
	SpeedAPIAvailabilitiesListResponseRegionsValueUsSouth1            SpeedAPIAvailabilitiesListResponseRegionsValue = "us-south1"
	SpeedAPIAvailabilitiesListResponseRegionsValueUsWest1             SpeedAPIAvailabilitiesListResponseRegionsValue = "us-west1"
)

type SpeedAPIPagesListResponse struct {
	// A test region with a label.
	Region SpeedAPIPagesListResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedAPIPagesListResponseScheduleFrequency `json:"scheduleFrequency"`
	Tests             []SpeedAPIPagesListResponseTest            `json:"tests"`
	// A URL.
	URL  string                        `json:"url"`
	JSON speedAPIPagesListResponseJSON `json:"-"`
}

// speedAPIPagesListResponseJSON contains the JSON metadata for the struct
// [SpeedAPIPagesListResponse]
type speedAPIPagesListResponseJSON struct {
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	Tests             apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region with a label.
type SpeedAPIPagesListResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPIPagesListResponseRegionValue `json:"value"`
	JSON  speedAPIPagesListResponseRegionJSON  `json:"-"`
}

// speedAPIPagesListResponseRegionJSON contains the JSON metadata for the struct
// [SpeedAPIPagesListResponseRegion]
type speedAPIPagesListResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPIPagesListResponseRegionValue string

const (
	SpeedAPIPagesListResponseRegionValueAsiaEast1           SpeedAPIPagesListResponseRegionValue = "asia-east1"
	SpeedAPIPagesListResponseRegionValueAsiaNortheast1      SpeedAPIPagesListResponseRegionValue = "asia-northeast1"
	SpeedAPIPagesListResponseRegionValueAsiaNortheast2      SpeedAPIPagesListResponseRegionValue = "asia-northeast2"
	SpeedAPIPagesListResponseRegionValueAsiaSouth1          SpeedAPIPagesListResponseRegionValue = "asia-south1"
	SpeedAPIPagesListResponseRegionValueAsiaSoutheast1      SpeedAPIPagesListResponseRegionValue = "asia-southeast1"
	SpeedAPIPagesListResponseRegionValueAustraliaSoutheast1 SpeedAPIPagesListResponseRegionValue = "australia-southeast1"
	SpeedAPIPagesListResponseRegionValueEuropeNorth1        SpeedAPIPagesListResponseRegionValue = "europe-north1"
	SpeedAPIPagesListResponseRegionValueEuropeSouthwest1    SpeedAPIPagesListResponseRegionValue = "europe-southwest1"
	SpeedAPIPagesListResponseRegionValueEuropeWest1         SpeedAPIPagesListResponseRegionValue = "europe-west1"
	SpeedAPIPagesListResponseRegionValueEuropeWest2         SpeedAPIPagesListResponseRegionValue = "europe-west2"
	SpeedAPIPagesListResponseRegionValueEuropeWest3         SpeedAPIPagesListResponseRegionValue = "europe-west3"
	SpeedAPIPagesListResponseRegionValueEuropeWest4         SpeedAPIPagesListResponseRegionValue = "europe-west4"
	SpeedAPIPagesListResponseRegionValueEuropeWest8         SpeedAPIPagesListResponseRegionValue = "europe-west8"
	SpeedAPIPagesListResponseRegionValueEuropeWest9         SpeedAPIPagesListResponseRegionValue = "europe-west9"
	SpeedAPIPagesListResponseRegionValueMeWest1             SpeedAPIPagesListResponseRegionValue = "me-west1"
	SpeedAPIPagesListResponseRegionValueSouthamericaEast1   SpeedAPIPagesListResponseRegionValue = "southamerica-east1"
	SpeedAPIPagesListResponseRegionValueUsCentral1          SpeedAPIPagesListResponseRegionValue = "us-central1"
	SpeedAPIPagesListResponseRegionValueUsEast1             SpeedAPIPagesListResponseRegionValue = "us-east1"
	SpeedAPIPagesListResponseRegionValueUsEast4             SpeedAPIPagesListResponseRegionValue = "us-east4"
	SpeedAPIPagesListResponseRegionValueUsSouth1            SpeedAPIPagesListResponseRegionValue = "us-south1"
	SpeedAPIPagesListResponseRegionValueUsWest1             SpeedAPIPagesListResponseRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedAPIPagesListResponseScheduleFrequency string

const (
	SpeedAPIPagesListResponseScheduleFrequencyDaily  SpeedAPIPagesListResponseScheduleFrequency = "DAILY"
	SpeedAPIPagesListResponseScheduleFrequencyWeekly SpeedAPIPagesListResponseScheduleFrequency = "WEEKLY"
)

type SpeedAPIPagesListResponseTest struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedAPIPagesListResponseTestsDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedAPIPagesListResponseTestsMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedAPIPagesListResponseTestsRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedAPIPagesListResponseTestsScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                            `json:"url"`
	JSON speedAPIPagesListResponseTestJSON `json:"-"`
}

// speedAPIPagesListResponseTestJSON contains the JSON metadata for the struct
// [SpeedAPIPagesListResponseTest]
type speedAPIPagesListResponseTestJSON struct {
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

func (r *SpeedAPIPagesListResponseTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedAPIPagesListResponseTestsDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPIPagesListResponseTestsDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedAPIPagesListResponseTestsDesktopReportError      `json:"error"`
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
	State SpeedAPIPagesListResponseTestsDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                         `json:"tti"`
	JSON speedAPIPagesListResponseTestsDesktopReportJSON `json:"-"`
}

// speedAPIPagesListResponseTestsDesktopReportJSON contains the JSON metadata for
// the struct [SpeedAPIPagesListResponseTestsDesktopReport]
type speedAPIPagesListResponseTestsDesktopReportJSON struct {
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

func (r *SpeedAPIPagesListResponseTestsDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPIPagesListResponseTestsDesktopReportDeviceType string

const (
	SpeedAPIPagesListResponseTestsDesktopReportDeviceTypeDesktop SpeedAPIPagesListResponseTestsDesktopReportDeviceType = "DESKTOP"
	SpeedAPIPagesListResponseTestsDesktopReportDeviceTypeMobile  SpeedAPIPagesListResponseTestsDesktopReportDeviceType = "MOBILE"
)

type SpeedAPIPagesListResponseTestsDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPIPagesListResponseTestsDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                               `json:"finalDisplayedUrl"`
	JSON              speedAPIPagesListResponseTestsDesktopReportErrorJSON `json:"-"`
}

// speedAPIPagesListResponseTestsDesktopReportErrorJSON contains the JSON metadata
// for the struct [SpeedAPIPagesListResponseTestsDesktopReportError]
type speedAPIPagesListResponseTestsDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponseTestsDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPIPagesListResponseTestsDesktopReportErrorCode string

const (
	SpeedAPIPagesListResponseTestsDesktopReportErrorCodeNotReachable      SpeedAPIPagesListResponseTestsDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedAPIPagesListResponseTestsDesktopReportErrorCodeDNSFailure        SpeedAPIPagesListResponseTestsDesktopReportErrorCode = "DNS_FAILURE"
	SpeedAPIPagesListResponseTestsDesktopReportErrorCodeNotHTML           SpeedAPIPagesListResponseTestsDesktopReportErrorCode = "NOT_HTML"
	SpeedAPIPagesListResponseTestsDesktopReportErrorCodeLighthouseTimeout SpeedAPIPagesListResponseTestsDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPIPagesListResponseTestsDesktopReportErrorCodeUnknown           SpeedAPIPagesListResponseTestsDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPIPagesListResponseTestsDesktopReportState string

const (
	SpeedAPIPagesListResponseTestsDesktopReportStateRunning  SpeedAPIPagesListResponseTestsDesktopReportState = "RUNNING"
	SpeedAPIPagesListResponseTestsDesktopReportStateComplete SpeedAPIPagesListResponseTestsDesktopReportState = "COMPLETE"
	SpeedAPIPagesListResponseTestsDesktopReportStateFailed   SpeedAPIPagesListResponseTestsDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedAPIPagesListResponseTestsMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPIPagesListResponseTestsMobileReportDeviceType `json:"deviceType"`
	Error      SpeedAPIPagesListResponseTestsMobileReportError      `json:"error"`
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
	State SpeedAPIPagesListResponseTestsMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                        `json:"tti"`
	JSON speedAPIPagesListResponseTestsMobileReportJSON `json:"-"`
}

// speedAPIPagesListResponseTestsMobileReportJSON contains the JSON metadata for
// the struct [SpeedAPIPagesListResponseTestsMobileReport]
type speedAPIPagesListResponseTestsMobileReportJSON struct {
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

func (r *SpeedAPIPagesListResponseTestsMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPIPagesListResponseTestsMobileReportDeviceType string

const (
	SpeedAPIPagesListResponseTestsMobileReportDeviceTypeDesktop SpeedAPIPagesListResponseTestsMobileReportDeviceType = "DESKTOP"
	SpeedAPIPagesListResponseTestsMobileReportDeviceTypeMobile  SpeedAPIPagesListResponseTestsMobileReportDeviceType = "MOBILE"
)

type SpeedAPIPagesListResponseTestsMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPIPagesListResponseTestsMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                              `json:"finalDisplayedUrl"`
	JSON              speedAPIPagesListResponseTestsMobileReportErrorJSON `json:"-"`
}

// speedAPIPagesListResponseTestsMobileReportErrorJSON contains the JSON metadata
// for the struct [SpeedAPIPagesListResponseTestsMobileReportError]
type speedAPIPagesListResponseTestsMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponseTestsMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPIPagesListResponseTestsMobileReportErrorCode string

const (
	SpeedAPIPagesListResponseTestsMobileReportErrorCodeNotReachable      SpeedAPIPagesListResponseTestsMobileReportErrorCode = "NOT_REACHABLE"
	SpeedAPIPagesListResponseTestsMobileReportErrorCodeDNSFailure        SpeedAPIPagesListResponseTestsMobileReportErrorCode = "DNS_FAILURE"
	SpeedAPIPagesListResponseTestsMobileReportErrorCodeNotHTML           SpeedAPIPagesListResponseTestsMobileReportErrorCode = "NOT_HTML"
	SpeedAPIPagesListResponseTestsMobileReportErrorCodeLighthouseTimeout SpeedAPIPagesListResponseTestsMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPIPagesListResponseTestsMobileReportErrorCodeUnknown           SpeedAPIPagesListResponseTestsMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPIPagesListResponseTestsMobileReportState string

const (
	SpeedAPIPagesListResponseTestsMobileReportStateRunning  SpeedAPIPagesListResponseTestsMobileReportState = "RUNNING"
	SpeedAPIPagesListResponseTestsMobileReportStateComplete SpeedAPIPagesListResponseTestsMobileReportState = "COMPLETE"
	SpeedAPIPagesListResponseTestsMobileReportStateFailed   SpeedAPIPagesListResponseTestsMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedAPIPagesListResponseTestsRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPIPagesListResponseTestsRegionValue `json:"value"`
	JSON  speedAPIPagesListResponseTestsRegionJSON  `json:"-"`
}

// speedAPIPagesListResponseTestsRegionJSON contains the JSON metadata for the
// struct [SpeedAPIPagesListResponseTestsRegion]
type speedAPIPagesListResponseTestsRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponseTestsRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPIPagesListResponseTestsRegionValue string

const (
	SpeedAPIPagesListResponseTestsRegionValueAsiaEast1           SpeedAPIPagesListResponseTestsRegionValue = "asia-east1"
	SpeedAPIPagesListResponseTestsRegionValueAsiaNortheast1      SpeedAPIPagesListResponseTestsRegionValue = "asia-northeast1"
	SpeedAPIPagesListResponseTestsRegionValueAsiaNortheast2      SpeedAPIPagesListResponseTestsRegionValue = "asia-northeast2"
	SpeedAPIPagesListResponseTestsRegionValueAsiaSouth1          SpeedAPIPagesListResponseTestsRegionValue = "asia-south1"
	SpeedAPIPagesListResponseTestsRegionValueAsiaSoutheast1      SpeedAPIPagesListResponseTestsRegionValue = "asia-southeast1"
	SpeedAPIPagesListResponseTestsRegionValueAustraliaSoutheast1 SpeedAPIPagesListResponseTestsRegionValue = "australia-southeast1"
	SpeedAPIPagesListResponseTestsRegionValueEuropeNorth1        SpeedAPIPagesListResponseTestsRegionValue = "europe-north1"
	SpeedAPIPagesListResponseTestsRegionValueEuropeSouthwest1    SpeedAPIPagesListResponseTestsRegionValue = "europe-southwest1"
	SpeedAPIPagesListResponseTestsRegionValueEuropeWest1         SpeedAPIPagesListResponseTestsRegionValue = "europe-west1"
	SpeedAPIPagesListResponseTestsRegionValueEuropeWest2         SpeedAPIPagesListResponseTestsRegionValue = "europe-west2"
	SpeedAPIPagesListResponseTestsRegionValueEuropeWest3         SpeedAPIPagesListResponseTestsRegionValue = "europe-west3"
	SpeedAPIPagesListResponseTestsRegionValueEuropeWest4         SpeedAPIPagesListResponseTestsRegionValue = "europe-west4"
	SpeedAPIPagesListResponseTestsRegionValueEuropeWest8         SpeedAPIPagesListResponseTestsRegionValue = "europe-west8"
	SpeedAPIPagesListResponseTestsRegionValueEuropeWest9         SpeedAPIPagesListResponseTestsRegionValue = "europe-west9"
	SpeedAPIPagesListResponseTestsRegionValueMeWest1             SpeedAPIPagesListResponseTestsRegionValue = "me-west1"
	SpeedAPIPagesListResponseTestsRegionValueSouthamericaEast1   SpeedAPIPagesListResponseTestsRegionValue = "southamerica-east1"
	SpeedAPIPagesListResponseTestsRegionValueUsCentral1          SpeedAPIPagesListResponseTestsRegionValue = "us-central1"
	SpeedAPIPagesListResponseTestsRegionValueUsEast1             SpeedAPIPagesListResponseTestsRegionValue = "us-east1"
	SpeedAPIPagesListResponseTestsRegionValueUsEast4             SpeedAPIPagesListResponseTestsRegionValue = "us-east4"
	SpeedAPIPagesListResponseTestsRegionValueUsSouth1            SpeedAPIPagesListResponseTestsRegionValue = "us-south1"
	SpeedAPIPagesListResponseTestsRegionValueUsWest1             SpeedAPIPagesListResponseTestsRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedAPIPagesListResponseTestsScheduleFrequency string

const (
	SpeedAPIPagesListResponseTestsScheduleFrequencyDaily  SpeedAPIPagesListResponseTestsScheduleFrequency = "DAILY"
	SpeedAPIPagesListResponseTestsScheduleFrequencyWeekly SpeedAPIPagesListResponseTestsScheduleFrequency = "WEEKLY"
)

type SpeedAPIScheduleDeleteResponse struct {
	// Number of items affected.
	Count float64                            `json:"count"`
	JSON  speedAPIScheduleDeleteResponseJSON `json:"-"`
}

// speedAPIScheduleDeleteResponseJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleDeleteResponse]
type speedAPIScheduleDeleteResponseJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The test schedule.
type SpeedAPIScheduleGetResponse struct {
	// The frequency of the test.
	Frequency SpeedAPIScheduleGetResponseFrequency `json:"frequency"`
	// A test region.
	Region SpeedAPIScheduleGetResponseRegion `json:"region"`
	// A URL.
	URL  string                          `json:"url"`
	JSON speedAPIScheduleGetResponseJSON `json:"-"`
}

// speedAPIScheduleGetResponseJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleGetResponse]
type speedAPIScheduleGetResponseJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency of the test.
type SpeedAPIScheduleGetResponseFrequency string

const (
	SpeedAPIScheduleGetResponseFrequencyDaily  SpeedAPIScheduleGetResponseFrequency = "DAILY"
	SpeedAPIScheduleGetResponseFrequencyWeekly SpeedAPIScheduleGetResponseFrequency = "WEEKLY"
)

// A test region.
type SpeedAPIScheduleGetResponseRegion string

const (
	SpeedAPIScheduleGetResponseRegionAsiaEast1           SpeedAPIScheduleGetResponseRegion = "asia-east1"
	SpeedAPIScheduleGetResponseRegionAsiaNortheast1      SpeedAPIScheduleGetResponseRegion = "asia-northeast1"
	SpeedAPIScheduleGetResponseRegionAsiaNortheast2      SpeedAPIScheduleGetResponseRegion = "asia-northeast2"
	SpeedAPIScheduleGetResponseRegionAsiaSouth1          SpeedAPIScheduleGetResponseRegion = "asia-south1"
	SpeedAPIScheduleGetResponseRegionAsiaSoutheast1      SpeedAPIScheduleGetResponseRegion = "asia-southeast1"
	SpeedAPIScheduleGetResponseRegionAustraliaSoutheast1 SpeedAPIScheduleGetResponseRegion = "australia-southeast1"
	SpeedAPIScheduleGetResponseRegionEuropeNorth1        SpeedAPIScheduleGetResponseRegion = "europe-north1"
	SpeedAPIScheduleGetResponseRegionEuropeSouthwest1    SpeedAPIScheduleGetResponseRegion = "europe-southwest1"
	SpeedAPIScheduleGetResponseRegionEuropeWest1         SpeedAPIScheduleGetResponseRegion = "europe-west1"
	SpeedAPIScheduleGetResponseRegionEuropeWest2         SpeedAPIScheduleGetResponseRegion = "europe-west2"
	SpeedAPIScheduleGetResponseRegionEuropeWest3         SpeedAPIScheduleGetResponseRegion = "europe-west3"
	SpeedAPIScheduleGetResponseRegionEuropeWest4         SpeedAPIScheduleGetResponseRegion = "europe-west4"
	SpeedAPIScheduleGetResponseRegionEuropeWest8         SpeedAPIScheduleGetResponseRegion = "europe-west8"
	SpeedAPIScheduleGetResponseRegionEuropeWest9         SpeedAPIScheduleGetResponseRegion = "europe-west9"
	SpeedAPIScheduleGetResponseRegionMeWest1             SpeedAPIScheduleGetResponseRegion = "me-west1"
	SpeedAPIScheduleGetResponseRegionSouthamericaEast1   SpeedAPIScheduleGetResponseRegion = "southamerica-east1"
	SpeedAPIScheduleGetResponseRegionUsCentral1          SpeedAPIScheduleGetResponseRegion = "us-central1"
	SpeedAPIScheduleGetResponseRegionUsEast1             SpeedAPIScheduleGetResponseRegion = "us-east1"
	SpeedAPIScheduleGetResponseRegionUsEast4             SpeedAPIScheduleGetResponseRegion = "us-east4"
	SpeedAPIScheduleGetResponseRegionUsSouth1            SpeedAPIScheduleGetResponseRegion = "us-south1"
	SpeedAPIScheduleGetResponseRegionUsWest1             SpeedAPIScheduleGetResponseRegion = "us-west1"
)

type SpeedAPITestsNewResponse struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedAPITestsNewResponseDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedAPITestsNewResponseMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedAPITestsNewResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedAPITestsNewResponseScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                       `json:"url"`
	JSON speedAPITestsNewResponseJSON `json:"-"`
}

// speedAPITestsNewResponseJSON contains the JSON metadata for the struct
// [SpeedAPITestsNewResponse]
type speedAPITestsNewResponseJSON struct {
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

func (r *SpeedAPITestsNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedAPITestsNewResponseDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPITestsNewResponseDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedAPITestsNewResponseDesktopReportError      `json:"error"`
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
	State SpeedAPITestsNewResponseDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                   `json:"tti"`
	JSON speedAPITestsNewResponseDesktopReportJSON `json:"-"`
}

// speedAPITestsNewResponseDesktopReportJSON contains the JSON metadata for the
// struct [SpeedAPITestsNewResponseDesktopReport]
type speedAPITestsNewResponseDesktopReportJSON struct {
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

func (r *SpeedAPITestsNewResponseDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPITestsNewResponseDesktopReportDeviceType string

const (
	SpeedAPITestsNewResponseDesktopReportDeviceTypeDesktop SpeedAPITestsNewResponseDesktopReportDeviceType = "DESKTOP"
	SpeedAPITestsNewResponseDesktopReportDeviceTypeMobile  SpeedAPITestsNewResponseDesktopReportDeviceType = "MOBILE"
)

type SpeedAPITestsNewResponseDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPITestsNewResponseDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                         `json:"finalDisplayedUrl"`
	JSON              speedAPITestsNewResponseDesktopReportErrorJSON `json:"-"`
}

// speedAPITestsNewResponseDesktopReportErrorJSON contains the JSON metadata for
// the struct [SpeedAPITestsNewResponseDesktopReportError]
type speedAPITestsNewResponseDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPITestsNewResponseDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPITestsNewResponseDesktopReportErrorCode string

const (
	SpeedAPITestsNewResponseDesktopReportErrorCodeNotReachable      SpeedAPITestsNewResponseDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedAPITestsNewResponseDesktopReportErrorCodeDNSFailure        SpeedAPITestsNewResponseDesktopReportErrorCode = "DNS_FAILURE"
	SpeedAPITestsNewResponseDesktopReportErrorCodeNotHTML           SpeedAPITestsNewResponseDesktopReportErrorCode = "NOT_HTML"
	SpeedAPITestsNewResponseDesktopReportErrorCodeLighthouseTimeout SpeedAPITestsNewResponseDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPITestsNewResponseDesktopReportErrorCodeUnknown           SpeedAPITestsNewResponseDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPITestsNewResponseDesktopReportState string

const (
	SpeedAPITestsNewResponseDesktopReportStateRunning  SpeedAPITestsNewResponseDesktopReportState = "RUNNING"
	SpeedAPITestsNewResponseDesktopReportStateComplete SpeedAPITestsNewResponseDesktopReportState = "COMPLETE"
	SpeedAPITestsNewResponseDesktopReportStateFailed   SpeedAPITestsNewResponseDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedAPITestsNewResponseMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPITestsNewResponseMobileReportDeviceType `json:"deviceType"`
	Error      SpeedAPITestsNewResponseMobileReportError      `json:"error"`
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
	State SpeedAPITestsNewResponseMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                  `json:"tti"`
	JSON speedAPITestsNewResponseMobileReportJSON `json:"-"`
}

// speedAPITestsNewResponseMobileReportJSON contains the JSON metadata for the
// struct [SpeedAPITestsNewResponseMobileReport]
type speedAPITestsNewResponseMobileReportJSON struct {
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

func (r *SpeedAPITestsNewResponseMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPITestsNewResponseMobileReportDeviceType string

const (
	SpeedAPITestsNewResponseMobileReportDeviceTypeDesktop SpeedAPITestsNewResponseMobileReportDeviceType = "DESKTOP"
	SpeedAPITestsNewResponseMobileReportDeviceTypeMobile  SpeedAPITestsNewResponseMobileReportDeviceType = "MOBILE"
)

type SpeedAPITestsNewResponseMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPITestsNewResponseMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                        `json:"finalDisplayedUrl"`
	JSON              speedAPITestsNewResponseMobileReportErrorJSON `json:"-"`
}

// speedAPITestsNewResponseMobileReportErrorJSON contains the JSON metadata for the
// struct [SpeedAPITestsNewResponseMobileReportError]
type speedAPITestsNewResponseMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPITestsNewResponseMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPITestsNewResponseMobileReportErrorCode string

const (
	SpeedAPITestsNewResponseMobileReportErrorCodeNotReachable      SpeedAPITestsNewResponseMobileReportErrorCode = "NOT_REACHABLE"
	SpeedAPITestsNewResponseMobileReportErrorCodeDNSFailure        SpeedAPITestsNewResponseMobileReportErrorCode = "DNS_FAILURE"
	SpeedAPITestsNewResponseMobileReportErrorCodeNotHTML           SpeedAPITestsNewResponseMobileReportErrorCode = "NOT_HTML"
	SpeedAPITestsNewResponseMobileReportErrorCodeLighthouseTimeout SpeedAPITestsNewResponseMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPITestsNewResponseMobileReportErrorCodeUnknown           SpeedAPITestsNewResponseMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPITestsNewResponseMobileReportState string

const (
	SpeedAPITestsNewResponseMobileReportStateRunning  SpeedAPITestsNewResponseMobileReportState = "RUNNING"
	SpeedAPITestsNewResponseMobileReportStateComplete SpeedAPITestsNewResponseMobileReportState = "COMPLETE"
	SpeedAPITestsNewResponseMobileReportStateFailed   SpeedAPITestsNewResponseMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedAPITestsNewResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPITestsNewResponseRegionValue `json:"value"`
	JSON  speedAPITestsNewResponseRegionJSON  `json:"-"`
}

// speedAPITestsNewResponseRegionJSON contains the JSON metadata for the struct
// [SpeedAPITestsNewResponseRegion]
type speedAPITestsNewResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsNewResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPITestsNewResponseRegionValue string

const (
	SpeedAPITestsNewResponseRegionValueAsiaEast1           SpeedAPITestsNewResponseRegionValue = "asia-east1"
	SpeedAPITestsNewResponseRegionValueAsiaNortheast1      SpeedAPITestsNewResponseRegionValue = "asia-northeast1"
	SpeedAPITestsNewResponseRegionValueAsiaNortheast2      SpeedAPITestsNewResponseRegionValue = "asia-northeast2"
	SpeedAPITestsNewResponseRegionValueAsiaSouth1          SpeedAPITestsNewResponseRegionValue = "asia-south1"
	SpeedAPITestsNewResponseRegionValueAsiaSoutheast1      SpeedAPITestsNewResponseRegionValue = "asia-southeast1"
	SpeedAPITestsNewResponseRegionValueAustraliaSoutheast1 SpeedAPITestsNewResponseRegionValue = "australia-southeast1"
	SpeedAPITestsNewResponseRegionValueEuropeNorth1        SpeedAPITestsNewResponseRegionValue = "europe-north1"
	SpeedAPITestsNewResponseRegionValueEuropeSouthwest1    SpeedAPITestsNewResponseRegionValue = "europe-southwest1"
	SpeedAPITestsNewResponseRegionValueEuropeWest1         SpeedAPITestsNewResponseRegionValue = "europe-west1"
	SpeedAPITestsNewResponseRegionValueEuropeWest2         SpeedAPITestsNewResponseRegionValue = "europe-west2"
	SpeedAPITestsNewResponseRegionValueEuropeWest3         SpeedAPITestsNewResponseRegionValue = "europe-west3"
	SpeedAPITestsNewResponseRegionValueEuropeWest4         SpeedAPITestsNewResponseRegionValue = "europe-west4"
	SpeedAPITestsNewResponseRegionValueEuropeWest8         SpeedAPITestsNewResponseRegionValue = "europe-west8"
	SpeedAPITestsNewResponseRegionValueEuropeWest9         SpeedAPITestsNewResponseRegionValue = "europe-west9"
	SpeedAPITestsNewResponseRegionValueMeWest1             SpeedAPITestsNewResponseRegionValue = "me-west1"
	SpeedAPITestsNewResponseRegionValueSouthamericaEast1   SpeedAPITestsNewResponseRegionValue = "southamerica-east1"
	SpeedAPITestsNewResponseRegionValueUsCentral1          SpeedAPITestsNewResponseRegionValue = "us-central1"
	SpeedAPITestsNewResponseRegionValueUsEast1             SpeedAPITestsNewResponseRegionValue = "us-east1"
	SpeedAPITestsNewResponseRegionValueUsEast4             SpeedAPITestsNewResponseRegionValue = "us-east4"
	SpeedAPITestsNewResponseRegionValueUsSouth1            SpeedAPITestsNewResponseRegionValue = "us-south1"
	SpeedAPITestsNewResponseRegionValueUsWest1             SpeedAPITestsNewResponseRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedAPITestsNewResponseScheduleFrequency string

const (
	SpeedAPITestsNewResponseScheduleFrequencyDaily  SpeedAPITestsNewResponseScheduleFrequency = "DAILY"
	SpeedAPITestsNewResponseScheduleFrequencyWeekly SpeedAPITestsNewResponseScheduleFrequency = "WEEKLY"
)

type SpeedAPITestsDeleteResponse struct {
	// Number of items affected.
	Count float64                         `json:"count"`
	JSON  speedAPITestsDeleteResponseJSON `json:"-"`
}

// speedAPITestsDeleteResponseJSON contains the JSON metadata for the struct
// [SpeedAPITestsDeleteResponse]
type speedAPITestsDeleteResponseJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsListResponse struct {
	Errors   []SpeedAPITestsListResponseError   `json:"errors,required"`
	Messages []SpeedAPITestsListResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    bool                                `json:"success,required"`
	ResultInfo SpeedAPITestsListResponseResultInfo `json:"result_info"`
	JSON       speedAPITestsListResponseJSON       `json:"-"`
}

// speedAPITestsListResponseJSON contains the JSON metadata for the struct
// [SpeedAPITestsListResponse]
type speedAPITestsListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsListResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    speedAPITestsListResponseErrorJSON `json:"-"`
}

// speedAPITestsListResponseErrorJSON contains the JSON metadata for the struct
// [SpeedAPITestsListResponseError]
type speedAPITestsListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsListResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    speedAPITestsListResponseMessageJSON `json:"-"`
}

// speedAPITestsListResponseMessageJSON contains the JSON metadata for the struct
// [SpeedAPITestsListResponseMessage]
type speedAPITestsListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsListResponseResultInfo struct {
	Count      int64                                   `json:"count"`
	Page       int64                                   `json:"page"`
	PerPage    int64                                   `json:"per_page"`
	TotalCount int64                                   `json:"total_count"`
	JSON       speedAPITestsListResponseResultInfoJSON `json:"-"`
}

// speedAPITestsListResponseResultInfoJSON contains the JSON metadata for the
// struct [SpeedAPITestsListResponseResultInfo]
type speedAPITestsListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsGetResponse struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedAPITestsGetResponseDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedAPITestsGetResponseMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedAPITestsGetResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedAPITestsGetResponseScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                       `json:"url"`
	JSON speedAPITestsGetResponseJSON `json:"-"`
}

// speedAPITestsGetResponseJSON contains the JSON metadata for the struct
// [SpeedAPITestsGetResponse]
type speedAPITestsGetResponseJSON struct {
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

func (r *SpeedAPITestsGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedAPITestsGetResponseDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPITestsGetResponseDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedAPITestsGetResponseDesktopReportError      `json:"error"`
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
	State SpeedAPITestsGetResponseDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                   `json:"tti"`
	JSON speedAPITestsGetResponseDesktopReportJSON `json:"-"`
}

// speedAPITestsGetResponseDesktopReportJSON contains the JSON metadata for the
// struct [SpeedAPITestsGetResponseDesktopReport]
type speedAPITestsGetResponseDesktopReportJSON struct {
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

func (r *SpeedAPITestsGetResponseDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPITestsGetResponseDesktopReportDeviceType string

const (
	SpeedAPITestsGetResponseDesktopReportDeviceTypeDesktop SpeedAPITestsGetResponseDesktopReportDeviceType = "DESKTOP"
	SpeedAPITestsGetResponseDesktopReportDeviceTypeMobile  SpeedAPITestsGetResponseDesktopReportDeviceType = "MOBILE"
)

type SpeedAPITestsGetResponseDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPITestsGetResponseDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                         `json:"finalDisplayedUrl"`
	JSON              speedAPITestsGetResponseDesktopReportErrorJSON `json:"-"`
}

// speedAPITestsGetResponseDesktopReportErrorJSON contains the JSON metadata for
// the struct [SpeedAPITestsGetResponseDesktopReportError]
type speedAPITestsGetResponseDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPITestsGetResponseDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPITestsGetResponseDesktopReportErrorCode string

const (
	SpeedAPITestsGetResponseDesktopReportErrorCodeNotReachable      SpeedAPITestsGetResponseDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedAPITestsGetResponseDesktopReportErrorCodeDNSFailure        SpeedAPITestsGetResponseDesktopReportErrorCode = "DNS_FAILURE"
	SpeedAPITestsGetResponseDesktopReportErrorCodeNotHTML           SpeedAPITestsGetResponseDesktopReportErrorCode = "NOT_HTML"
	SpeedAPITestsGetResponseDesktopReportErrorCodeLighthouseTimeout SpeedAPITestsGetResponseDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPITestsGetResponseDesktopReportErrorCodeUnknown           SpeedAPITestsGetResponseDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPITestsGetResponseDesktopReportState string

const (
	SpeedAPITestsGetResponseDesktopReportStateRunning  SpeedAPITestsGetResponseDesktopReportState = "RUNNING"
	SpeedAPITestsGetResponseDesktopReportStateComplete SpeedAPITestsGetResponseDesktopReportState = "COMPLETE"
	SpeedAPITestsGetResponseDesktopReportStateFailed   SpeedAPITestsGetResponseDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedAPITestsGetResponseMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPITestsGetResponseMobileReportDeviceType `json:"deviceType"`
	Error      SpeedAPITestsGetResponseMobileReportError      `json:"error"`
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
	State SpeedAPITestsGetResponseMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                  `json:"tti"`
	JSON speedAPITestsGetResponseMobileReportJSON `json:"-"`
}

// speedAPITestsGetResponseMobileReportJSON contains the JSON metadata for the
// struct [SpeedAPITestsGetResponseMobileReport]
type speedAPITestsGetResponseMobileReportJSON struct {
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

func (r *SpeedAPITestsGetResponseMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPITestsGetResponseMobileReportDeviceType string

const (
	SpeedAPITestsGetResponseMobileReportDeviceTypeDesktop SpeedAPITestsGetResponseMobileReportDeviceType = "DESKTOP"
	SpeedAPITestsGetResponseMobileReportDeviceTypeMobile  SpeedAPITestsGetResponseMobileReportDeviceType = "MOBILE"
)

type SpeedAPITestsGetResponseMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPITestsGetResponseMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                        `json:"finalDisplayedUrl"`
	JSON              speedAPITestsGetResponseMobileReportErrorJSON `json:"-"`
}

// speedAPITestsGetResponseMobileReportErrorJSON contains the JSON metadata for the
// struct [SpeedAPITestsGetResponseMobileReportError]
type speedAPITestsGetResponseMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPITestsGetResponseMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPITestsGetResponseMobileReportErrorCode string

const (
	SpeedAPITestsGetResponseMobileReportErrorCodeNotReachable      SpeedAPITestsGetResponseMobileReportErrorCode = "NOT_REACHABLE"
	SpeedAPITestsGetResponseMobileReportErrorCodeDNSFailure        SpeedAPITestsGetResponseMobileReportErrorCode = "DNS_FAILURE"
	SpeedAPITestsGetResponseMobileReportErrorCodeNotHTML           SpeedAPITestsGetResponseMobileReportErrorCode = "NOT_HTML"
	SpeedAPITestsGetResponseMobileReportErrorCodeLighthouseTimeout SpeedAPITestsGetResponseMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPITestsGetResponseMobileReportErrorCodeUnknown           SpeedAPITestsGetResponseMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPITestsGetResponseMobileReportState string

const (
	SpeedAPITestsGetResponseMobileReportStateRunning  SpeedAPITestsGetResponseMobileReportState = "RUNNING"
	SpeedAPITestsGetResponseMobileReportStateComplete SpeedAPITestsGetResponseMobileReportState = "COMPLETE"
	SpeedAPITestsGetResponseMobileReportStateFailed   SpeedAPITestsGetResponseMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedAPITestsGetResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPITestsGetResponseRegionValue `json:"value"`
	JSON  speedAPITestsGetResponseRegionJSON  `json:"-"`
}

// speedAPITestsGetResponseRegionJSON contains the JSON metadata for the struct
// [SpeedAPITestsGetResponseRegion]
type speedAPITestsGetResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsGetResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPITestsGetResponseRegionValue string

const (
	SpeedAPITestsGetResponseRegionValueAsiaEast1           SpeedAPITestsGetResponseRegionValue = "asia-east1"
	SpeedAPITestsGetResponseRegionValueAsiaNortheast1      SpeedAPITestsGetResponseRegionValue = "asia-northeast1"
	SpeedAPITestsGetResponseRegionValueAsiaNortheast2      SpeedAPITestsGetResponseRegionValue = "asia-northeast2"
	SpeedAPITestsGetResponseRegionValueAsiaSouth1          SpeedAPITestsGetResponseRegionValue = "asia-south1"
	SpeedAPITestsGetResponseRegionValueAsiaSoutheast1      SpeedAPITestsGetResponseRegionValue = "asia-southeast1"
	SpeedAPITestsGetResponseRegionValueAustraliaSoutheast1 SpeedAPITestsGetResponseRegionValue = "australia-southeast1"
	SpeedAPITestsGetResponseRegionValueEuropeNorth1        SpeedAPITestsGetResponseRegionValue = "europe-north1"
	SpeedAPITestsGetResponseRegionValueEuropeSouthwest1    SpeedAPITestsGetResponseRegionValue = "europe-southwest1"
	SpeedAPITestsGetResponseRegionValueEuropeWest1         SpeedAPITestsGetResponseRegionValue = "europe-west1"
	SpeedAPITestsGetResponseRegionValueEuropeWest2         SpeedAPITestsGetResponseRegionValue = "europe-west2"
	SpeedAPITestsGetResponseRegionValueEuropeWest3         SpeedAPITestsGetResponseRegionValue = "europe-west3"
	SpeedAPITestsGetResponseRegionValueEuropeWest4         SpeedAPITestsGetResponseRegionValue = "europe-west4"
	SpeedAPITestsGetResponseRegionValueEuropeWest8         SpeedAPITestsGetResponseRegionValue = "europe-west8"
	SpeedAPITestsGetResponseRegionValueEuropeWest9         SpeedAPITestsGetResponseRegionValue = "europe-west9"
	SpeedAPITestsGetResponseRegionValueMeWest1             SpeedAPITestsGetResponseRegionValue = "me-west1"
	SpeedAPITestsGetResponseRegionValueSouthamericaEast1   SpeedAPITestsGetResponseRegionValue = "southamerica-east1"
	SpeedAPITestsGetResponseRegionValueUsCentral1          SpeedAPITestsGetResponseRegionValue = "us-central1"
	SpeedAPITestsGetResponseRegionValueUsEast1             SpeedAPITestsGetResponseRegionValue = "us-east1"
	SpeedAPITestsGetResponseRegionValueUsEast4             SpeedAPITestsGetResponseRegionValue = "us-east4"
	SpeedAPITestsGetResponseRegionValueUsSouth1            SpeedAPITestsGetResponseRegionValue = "us-south1"
	SpeedAPITestsGetResponseRegionValueUsWest1             SpeedAPITestsGetResponseRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedAPITestsGetResponseScheduleFrequency string

const (
	SpeedAPITestsGetResponseScheduleFrequencyDaily  SpeedAPITestsGetResponseScheduleFrequency = "DAILY"
	SpeedAPITestsGetResponseScheduleFrequencyWeekly SpeedAPITestsGetResponseScheduleFrequency = "WEEKLY"
)

type SpeedAPITrendsListResponse struct {
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
	Tti  []float64                      `json:"tti"`
	JSON speedAPITrendsListResponseJSON `json:"-"`
}

// speedAPITrendsListResponseJSON contains the JSON metadata for the struct
// [SpeedAPITrendsListResponse]
type speedAPITrendsListResponseJSON struct {
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

func (r *SpeedAPITrendsListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIAvailabilitiesListResponseEnvelope struct {
	Result SpeedAPIAvailabilitiesListResponse             `json:"result"`
	JSON   speedAPIAvailabilitiesListResponseEnvelopeJSON `json:"-"`
}

// speedAPIAvailabilitiesListResponseEnvelopeJSON contains the JSON metadata for
// the struct [SpeedAPIAvailabilitiesListResponseEnvelope]
type speedAPIAvailabilitiesListResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIAvailabilitiesListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIPagesListResponseEnvelope struct {
	Result []SpeedAPIPagesListResponse           `json:"result"`
	JSON   speedAPIPagesListResponseEnvelopeJSON `json:"-"`
}

// speedAPIPagesListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedAPIPagesListResponseEnvelope]
type speedAPIPagesListResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleDeleteParams struct {
	// A test region.
	Region param.Field[SpeedAPIScheduleDeleteParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedAPIScheduleDeleteParams]'s query parameters as
// `url.Values`.
func (r SpeedAPIScheduleDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedAPIScheduleDeleteParamsRegion string

const (
	SpeedAPIScheduleDeleteParamsRegionAsiaEast1           SpeedAPIScheduleDeleteParamsRegion = "asia-east1"
	SpeedAPIScheduleDeleteParamsRegionAsiaNortheast1      SpeedAPIScheduleDeleteParamsRegion = "asia-northeast1"
	SpeedAPIScheduleDeleteParamsRegionAsiaNortheast2      SpeedAPIScheduleDeleteParamsRegion = "asia-northeast2"
	SpeedAPIScheduleDeleteParamsRegionAsiaSouth1          SpeedAPIScheduleDeleteParamsRegion = "asia-south1"
	SpeedAPIScheduleDeleteParamsRegionAsiaSoutheast1      SpeedAPIScheduleDeleteParamsRegion = "asia-southeast1"
	SpeedAPIScheduleDeleteParamsRegionAustraliaSoutheast1 SpeedAPIScheduleDeleteParamsRegion = "australia-southeast1"
	SpeedAPIScheduleDeleteParamsRegionEuropeNorth1        SpeedAPIScheduleDeleteParamsRegion = "europe-north1"
	SpeedAPIScheduleDeleteParamsRegionEuropeSouthwest1    SpeedAPIScheduleDeleteParamsRegion = "europe-southwest1"
	SpeedAPIScheduleDeleteParamsRegionEuropeWest1         SpeedAPIScheduleDeleteParamsRegion = "europe-west1"
	SpeedAPIScheduleDeleteParamsRegionEuropeWest2         SpeedAPIScheduleDeleteParamsRegion = "europe-west2"
	SpeedAPIScheduleDeleteParamsRegionEuropeWest3         SpeedAPIScheduleDeleteParamsRegion = "europe-west3"
	SpeedAPIScheduleDeleteParamsRegionEuropeWest4         SpeedAPIScheduleDeleteParamsRegion = "europe-west4"
	SpeedAPIScheduleDeleteParamsRegionEuropeWest8         SpeedAPIScheduleDeleteParamsRegion = "europe-west8"
	SpeedAPIScheduleDeleteParamsRegionEuropeWest9         SpeedAPIScheduleDeleteParamsRegion = "europe-west9"
	SpeedAPIScheduleDeleteParamsRegionMeWest1             SpeedAPIScheduleDeleteParamsRegion = "me-west1"
	SpeedAPIScheduleDeleteParamsRegionSouthamericaEast1   SpeedAPIScheduleDeleteParamsRegion = "southamerica-east1"
	SpeedAPIScheduleDeleteParamsRegionUsCentral1          SpeedAPIScheduleDeleteParamsRegion = "us-central1"
	SpeedAPIScheduleDeleteParamsRegionUsEast1             SpeedAPIScheduleDeleteParamsRegion = "us-east1"
	SpeedAPIScheduleDeleteParamsRegionUsEast4             SpeedAPIScheduleDeleteParamsRegion = "us-east4"
	SpeedAPIScheduleDeleteParamsRegionUsSouth1            SpeedAPIScheduleDeleteParamsRegion = "us-south1"
	SpeedAPIScheduleDeleteParamsRegionUsWest1             SpeedAPIScheduleDeleteParamsRegion = "us-west1"
)

type SpeedAPIScheduleDeleteResponseEnvelope struct {
	Result SpeedAPIScheduleDeleteResponse             `json:"result"`
	JSON   speedAPIScheduleDeleteResponseEnvelopeJSON `json:"-"`
}

// speedAPIScheduleDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SpeedAPIScheduleDeleteResponseEnvelope]
type speedAPIScheduleDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleGetParams struct {
	// A test region.
	Region param.Field[SpeedAPIScheduleGetParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedAPIScheduleGetParams]'s query parameters as
// `url.Values`.
func (r SpeedAPIScheduleGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedAPIScheduleGetParamsRegion string

const (
	SpeedAPIScheduleGetParamsRegionAsiaEast1           SpeedAPIScheduleGetParamsRegion = "asia-east1"
	SpeedAPIScheduleGetParamsRegionAsiaNortheast1      SpeedAPIScheduleGetParamsRegion = "asia-northeast1"
	SpeedAPIScheduleGetParamsRegionAsiaNortheast2      SpeedAPIScheduleGetParamsRegion = "asia-northeast2"
	SpeedAPIScheduleGetParamsRegionAsiaSouth1          SpeedAPIScheduleGetParamsRegion = "asia-south1"
	SpeedAPIScheduleGetParamsRegionAsiaSoutheast1      SpeedAPIScheduleGetParamsRegion = "asia-southeast1"
	SpeedAPIScheduleGetParamsRegionAustraliaSoutheast1 SpeedAPIScheduleGetParamsRegion = "australia-southeast1"
	SpeedAPIScheduleGetParamsRegionEuropeNorth1        SpeedAPIScheduleGetParamsRegion = "europe-north1"
	SpeedAPIScheduleGetParamsRegionEuropeSouthwest1    SpeedAPIScheduleGetParamsRegion = "europe-southwest1"
	SpeedAPIScheduleGetParamsRegionEuropeWest1         SpeedAPIScheduleGetParamsRegion = "europe-west1"
	SpeedAPIScheduleGetParamsRegionEuropeWest2         SpeedAPIScheduleGetParamsRegion = "europe-west2"
	SpeedAPIScheduleGetParamsRegionEuropeWest3         SpeedAPIScheduleGetParamsRegion = "europe-west3"
	SpeedAPIScheduleGetParamsRegionEuropeWest4         SpeedAPIScheduleGetParamsRegion = "europe-west4"
	SpeedAPIScheduleGetParamsRegionEuropeWest8         SpeedAPIScheduleGetParamsRegion = "europe-west8"
	SpeedAPIScheduleGetParamsRegionEuropeWest9         SpeedAPIScheduleGetParamsRegion = "europe-west9"
	SpeedAPIScheduleGetParamsRegionMeWest1             SpeedAPIScheduleGetParamsRegion = "me-west1"
	SpeedAPIScheduleGetParamsRegionSouthamericaEast1   SpeedAPIScheduleGetParamsRegion = "southamerica-east1"
	SpeedAPIScheduleGetParamsRegionUsCentral1          SpeedAPIScheduleGetParamsRegion = "us-central1"
	SpeedAPIScheduleGetParamsRegionUsEast1             SpeedAPIScheduleGetParamsRegion = "us-east1"
	SpeedAPIScheduleGetParamsRegionUsEast4             SpeedAPIScheduleGetParamsRegion = "us-east4"
	SpeedAPIScheduleGetParamsRegionUsSouth1            SpeedAPIScheduleGetParamsRegion = "us-south1"
	SpeedAPIScheduleGetParamsRegionUsWest1             SpeedAPIScheduleGetParamsRegion = "us-west1"
)

type SpeedAPIScheduleGetResponseEnvelope struct {
	// The test schedule.
	Result SpeedAPIScheduleGetResponse             `json:"result"`
	JSON   speedAPIScheduleGetResponseEnvelopeJSON `json:"-"`
}

// speedAPIScheduleGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SpeedAPIScheduleGetResponseEnvelope]
type speedAPIScheduleGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsNewParams struct {
	// A test region.
	Region param.Field[SpeedAPITestsNewParamsRegion] `json:"region"`
}

func (r SpeedAPITestsNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A test region.
type SpeedAPITestsNewParamsRegion string

const (
	SpeedAPITestsNewParamsRegionAsiaEast1           SpeedAPITestsNewParamsRegion = "asia-east1"
	SpeedAPITestsNewParamsRegionAsiaNortheast1      SpeedAPITestsNewParamsRegion = "asia-northeast1"
	SpeedAPITestsNewParamsRegionAsiaNortheast2      SpeedAPITestsNewParamsRegion = "asia-northeast2"
	SpeedAPITestsNewParamsRegionAsiaSouth1          SpeedAPITestsNewParamsRegion = "asia-south1"
	SpeedAPITestsNewParamsRegionAsiaSoutheast1      SpeedAPITestsNewParamsRegion = "asia-southeast1"
	SpeedAPITestsNewParamsRegionAustraliaSoutheast1 SpeedAPITestsNewParamsRegion = "australia-southeast1"
	SpeedAPITestsNewParamsRegionEuropeNorth1        SpeedAPITestsNewParamsRegion = "europe-north1"
	SpeedAPITestsNewParamsRegionEuropeSouthwest1    SpeedAPITestsNewParamsRegion = "europe-southwest1"
	SpeedAPITestsNewParamsRegionEuropeWest1         SpeedAPITestsNewParamsRegion = "europe-west1"
	SpeedAPITestsNewParamsRegionEuropeWest2         SpeedAPITestsNewParamsRegion = "europe-west2"
	SpeedAPITestsNewParamsRegionEuropeWest3         SpeedAPITestsNewParamsRegion = "europe-west3"
	SpeedAPITestsNewParamsRegionEuropeWest4         SpeedAPITestsNewParamsRegion = "europe-west4"
	SpeedAPITestsNewParamsRegionEuropeWest8         SpeedAPITestsNewParamsRegion = "europe-west8"
	SpeedAPITestsNewParamsRegionEuropeWest9         SpeedAPITestsNewParamsRegion = "europe-west9"
	SpeedAPITestsNewParamsRegionMeWest1             SpeedAPITestsNewParamsRegion = "me-west1"
	SpeedAPITestsNewParamsRegionSouthamericaEast1   SpeedAPITestsNewParamsRegion = "southamerica-east1"
	SpeedAPITestsNewParamsRegionUsCentral1          SpeedAPITestsNewParamsRegion = "us-central1"
	SpeedAPITestsNewParamsRegionUsEast1             SpeedAPITestsNewParamsRegion = "us-east1"
	SpeedAPITestsNewParamsRegionUsEast4             SpeedAPITestsNewParamsRegion = "us-east4"
	SpeedAPITestsNewParamsRegionUsSouth1            SpeedAPITestsNewParamsRegion = "us-south1"
	SpeedAPITestsNewParamsRegionUsWest1             SpeedAPITestsNewParamsRegion = "us-west1"
)

type SpeedAPITestsNewResponseEnvelope struct {
	Result SpeedAPITestsNewResponse             `json:"result"`
	JSON   speedAPITestsNewResponseEnvelopeJSON `json:"-"`
}

// speedAPITestsNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedAPITestsNewResponseEnvelope]
type speedAPITestsNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsDeleteParams struct {
	// A test region.
	Region param.Field[SpeedAPITestsDeleteParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedAPITestsDeleteParams]'s query parameters as
// `url.Values`.
func (r SpeedAPITestsDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedAPITestsDeleteParamsRegion string

const (
	SpeedAPITestsDeleteParamsRegionAsiaEast1           SpeedAPITestsDeleteParamsRegion = "asia-east1"
	SpeedAPITestsDeleteParamsRegionAsiaNortheast1      SpeedAPITestsDeleteParamsRegion = "asia-northeast1"
	SpeedAPITestsDeleteParamsRegionAsiaNortheast2      SpeedAPITestsDeleteParamsRegion = "asia-northeast2"
	SpeedAPITestsDeleteParamsRegionAsiaSouth1          SpeedAPITestsDeleteParamsRegion = "asia-south1"
	SpeedAPITestsDeleteParamsRegionAsiaSoutheast1      SpeedAPITestsDeleteParamsRegion = "asia-southeast1"
	SpeedAPITestsDeleteParamsRegionAustraliaSoutheast1 SpeedAPITestsDeleteParamsRegion = "australia-southeast1"
	SpeedAPITestsDeleteParamsRegionEuropeNorth1        SpeedAPITestsDeleteParamsRegion = "europe-north1"
	SpeedAPITestsDeleteParamsRegionEuropeSouthwest1    SpeedAPITestsDeleteParamsRegion = "europe-southwest1"
	SpeedAPITestsDeleteParamsRegionEuropeWest1         SpeedAPITestsDeleteParamsRegion = "europe-west1"
	SpeedAPITestsDeleteParamsRegionEuropeWest2         SpeedAPITestsDeleteParamsRegion = "europe-west2"
	SpeedAPITestsDeleteParamsRegionEuropeWest3         SpeedAPITestsDeleteParamsRegion = "europe-west3"
	SpeedAPITestsDeleteParamsRegionEuropeWest4         SpeedAPITestsDeleteParamsRegion = "europe-west4"
	SpeedAPITestsDeleteParamsRegionEuropeWest8         SpeedAPITestsDeleteParamsRegion = "europe-west8"
	SpeedAPITestsDeleteParamsRegionEuropeWest9         SpeedAPITestsDeleteParamsRegion = "europe-west9"
	SpeedAPITestsDeleteParamsRegionMeWest1             SpeedAPITestsDeleteParamsRegion = "me-west1"
	SpeedAPITestsDeleteParamsRegionSouthamericaEast1   SpeedAPITestsDeleteParamsRegion = "southamerica-east1"
	SpeedAPITestsDeleteParamsRegionUsCentral1          SpeedAPITestsDeleteParamsRegion = "us-central1"
	SpeedAPITestsDeleteParamsRegionUsEast1             SpeedAPITestsDeleteParamsRegion = "us-east1"
	SpeedAPITestsDeleteParamsRegionUsEast4             SpeedAPITestsDeleteParamsRegion = "us-east4"
	SpeedAPITestsDeleteParamsRegionUsSouth1            SpeedAPITestsDeleteParamsRegion = "us-south1"
	SpeedAPITestsDeleteParamsRegionUsWest1             SpeedAPITestsDeleteParamsRegion = "us-west1"
)

type SpeedAPITestsDeleteResponseEnvelope struct {
	Result SpeedAPITestsDeleteResponse             `json:"result"`
	JSON   speedAPITestsDeleteResponseEnvelopeJSON `json:"-"`
}

// speedAPITestsDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SpeedAPITestsDeleteResponseEnvelope]
type speedAPITestsDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsListParams struct {
	Page    param.Field[int64] `query:"page"`
	PerPage param.Field[int64] `query:"per_page"`
	// A test region.
	Region param.Field[SpeedAPITestsListParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedAPITestsListParams]'s query parameters as
// `url.Values`.
func (r SpeedAPITestsListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedAPITestsListParamsRegion string

const (
	SpeedAPITestsListParamsRegionAsiaEast1           SpeedAPITestsListParamsRegion = "asia-east1"
	SpeedAPITestsListParamsRegionAsiaNortheast1      SpeedAPITestsListParamsRegion = "asia-northeast1"
	SpeedAPITestsListParamsRegionAsiaNortheast2      SpeedAPITestsListParamsRegion = "asia-northeast2"
	SpeedAPITestsListParamsRegionAsiaSouth1          SpeedAPITestsListParamsRegion = "asia-south1"
	SpeedAPITestsListParamsRegionAsiaSoutheast1      SpeedAPITestsListParamsRegion = "asia-southeast1"
	SpeedAPITestsListParamsRegionAustraliaSoutheast1 SpeedAPITestsListParamsRegion = "australia-southeast1"
	SpeedAPITestsListParamsRegionEuropeNorth1        SpeedAPITestsListParamsRegion = "europe-north1"
	SpeedAPITestsListParamsRegionEuropeSouthwest1    SpeedAPITestsListParamsRegion = "europe-southwest1"
	SpeedAPITestsListParamsRegionEuropeWest1         SpeedAPITestsListParamsRegion = "europe-west1"
	SpeedAPITestsListParamsRegionEuropeWest2         SpeedAPITestsListParamsRegion = "europe-west2"
	SpeedAPITestsListParamsRegionEuropeWest3         SpeedAPITestsListParamsRegion = "europe-west3"
	SpeedAPITestsListParamsRegionEuropeWest4         SpeedAPITestsListParamsRegion = "europe-west4"
	SpeedAPITestsListParamsRegionEuropeWest8         SpeedAPITestsListParamsRegion = "europe-west8"
	SpeedAPITestsListParamsRegionEuropeWest9         SpeedAPITestsListParamsRegion = "europe-west9"
	SpeedAPITestsListParamsRegionMeWest1             SpeedAPITestsListParamsRegion = "me-west1"
	SpeedAPITestsListParamsRegionSouthamericaEast1   SpeedAPITestsListParamsRegion = "southamerica-east1"
	SpeedAPITestsListParamsRegionUsCentral1          SpeedAPITestsListParamsRegion = "us-central1"
	SpeedAPITestsListParamsRegionUsEast1             SpeedAPITestsListParamsRegion = "us-east1"
	SpeedAPITestsListParamsRegionUsEast4             SpeedAPITestsListParamsRegion = "us-east4"
	SpeedAPITestsListParamsRegionUsSouth1            SpeedAPITestsListParamsRegion = "us-south1"
	SpeedAPITestsListParamsRegionUsWest1             SpeedAPITestsListParamsRegion = "us-west1"
)

type SpeedAPITestsGetResponseEnvelope struct {
	Result SpeedAPITestsGetResponse             `json:"result"`
	JSON   speedAPITestsGetResponseEnvelopeJSON `json:"-"`
}

// speedAPITestsGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedAPITestsGetResponseEnvelope]
type speedAPITestsGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITrendsListParams struct {
	// The type of device.
	DeviceType param.Field[SpeedAPITrendsListParamsDeviceType] `query:"deviceType,required"`
	// A comma-separated list of metrics to include in the results.
	Metrics param.Field[string] `query:"metrics,required"`
	// A test region.
	Region param.Field[SpeedAPITrendsListParamsRegion] `query:"region,required"`
	// The timezone of the start and end timestamps.
	Tz param.Field[string] `query:"tz,required"`
}

// URLQuery serializes [SpeedAPITrendsListParams]'s query parameters as
// `url.Values`.
func (r SpeedAPITrendsListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of device.
type SpeedAPITrendsListParamsDeviceType string

const (
	SpeedAPITrendsListParamsDeviceTypeDesktop SpeedAPITrendsListParamsDeviceType = "DESKTOP"
	SpeedAPITrendsListParamsDeviceTypeMobile  SpeedAPITrendsListParamsDeviceType = "MOBILE"
)

// A test region.
type SpeedAPITrendsListParamsRegion string

const (
	SpeedAPITrendsListParamsRegionAsiaEast1           SpeedAPITrendsListParamsRegion = "asia-east1"
	SpeedAPITrendsListParamsRegionAsiaNortheast1      SpeedAPITrendsListParamsRegion = "asia-northeast1"
	SpeedAPITrendsListParamsRegionAsiaNortheast2      SpeedAPITrendsListParamsRegion = "asia-northeast2"
	SpeedAPITrendsListParamsRegionAsiaSouth1          SpeedAPITrendsListParamsRegion = "asia-south1"
	SpeedAPITrendsListParamsRegionAsiaSoutheast1      SpeedAPITrendsListParamsRegion = "asia-southeast1"
	SpeedAPITrendsListParamsRegionAustraliaSoutheast1 SpeedAPITrendsListParamsRegion = "australia-southeast1"
	SpeedAPITrendsListParamsRegionEuropeNorth1        SpeedAPITrendsListParamsRegion = "europe-north1"
	SpeedAPITrendsListParamsRegionEuropeSouthwest1    SpeedAPITrendsListParamsRegion = "europe-southwest1"
	SpeedAPITrendsListParamsRegionEuropeWest1         SpeedAPITrendsListParamsRegion = "europe-west1"
	SpeedAPITrendsListParamsRegionEuropeWest2         SpeedAPITrendsListParamsRegion = "europe-west2"
	SpeedAPITrendsListParamsRegionEuropeWest3         SpeedAPITrendsListParamsRegion = "europe-west3"
	SpeedAPITrendsListParamsRegionEuropeWest4         SpeedAPITrendsListParamsRegion = "europe-west4"
	SpeedAPITrendsListParamsRegionEuropeWest8         SpeedAPITrendsListParamsRegion = "europe-west8"
	SpeedAPITrendsListParamsRegionEuropeWest9         SpeedAPITrendsListParamsRegion = "europe-west9"
	SpeedAPITrendsListParamsRegionMeWest1             SpeedAPITrendsListParamsRegion = "me-west1"
	SpeedAPITrendsListParamsRegionSouthamericaEast1   SpeedAPITrendsListParamsRegion = "southamerica-east1"
	SpeedAPITrendsListParamsRegionUsCentral1          SpeedAPITrendsListParamsRegion = "us-central1"
	SpeedAPITrendsListParamsRegionUsEast1             SpeedAPITrendsListParamsRegion = "us-east1"
	SpeedAPITrendsListParamsRegionUsEast4             SpeedAPITrendsListParamsRegion = "us-east4"
	SpeedAPITrendsListParamsRegionUsSouth1            SpeedAPITrendsListParamsRegion = "us-south1"
	SpeedAPITrendsListParamsRegionUsWest1             SpeedAPITrendsListParamsRegion = "us-west1"
)

type SpeedAPITrendsListResponseEnvelope struct {
	Result SpeedAPITrendsListResponse             `json:"result"`
	JSON   speedAPITrendsListResponseEnvelopeJSON `json:"-"`
}

// speedAPITrendsListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedAPITrendsListResponseEnvelope]
type speedAPITrendsListResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITrendsListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
