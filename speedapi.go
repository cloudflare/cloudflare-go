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
	path := fmt.Sprintf("zones/%s/speed_api/availabilities", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all webpages which have been tested.
func (r *SpeedAPIService) PagesList(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SpeedAPIPagesListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a scheduled test for a page.
func (r *SpeedAPIService) ScheduleDelete(ctx context.Context, zoneID string, url string, body SpeedAPIScheduleDeleteParams, opts ...option.RequestOption) (res *SpeedAPIScheduleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Retrieves the test schedule for a page in a specific region.
func (r *SpeedAPIService) ScheduleGet(ctx context.Context, zoneID string, url string, query SpeedAPIScheduleGetParams, opts ...option.RequestOption) (res *SpeedAPIScheduleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Starts a test for a specific webpage, in a specific region.
func (r *SpeedAPIService) TestsNew(ctx context.Context, zoneID string, url string, body SpeedAPITestsNewParams, opts ...option.RequestOption) (res *SpeedAPITestsNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes all tests for a specific webpage from a specific region. Deleted tests
// are still counted as part of the quota.
func (r *SpeedAPIService) TestsDelete(ctx context.Context, zoneID string, url string, body SpeedAPITestsDeleteParams, opts ...option.RequestOption) (res *SpeedAPITestsDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
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
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests/%s", zoneID, url, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists the core web vital metrics trend over time for a specific page.
func (r *SpeedAPIService) TrendsList(ctx context.Context, zoneID string, url string, query SpeedAPITrendsListParams, opts ...option.RequestOption) (res *SpeedAPITrendsListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/trend", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SpeedAPIAvailabilitiesListResponse struct {
	Errors   []SpeedAPIAvailabilitiesListResponseError   `json:"errors"`
	Messages []SpeedAPIAvailabilitiesListResponseMessage `json:"messages"`
	Result   SpeedAPIAvailabilitiesListResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                                   `json:"success"`
	JSON    speedAPIAvailabilitiesListResponseJSON `json:"-"`
}

// speedAPIAvailabilitiesListResponseJSON contains the JSON metadata for the struct
// [SpeedAPIAvailabilitiesListResponse]
type speedAPIAvailabilitiesListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIAvailabilitiesListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIAvailabilitiesListResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    speedAPIAvailabilitiesListResponseErrorJSON `json:"-"`
}

// speedAPIAvailabilitiesListResponseErrorJSON contains the JSON metadata for the
// struct [SpeedAPIAvailabilitiesListResponseError]
type speedAPIAvailabilitiesListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIAvailabilitiesListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIAvailabilitiesListResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    speedAPIAvailabilitiesListResponseMessageJSON `json:"-"`
}

// speedAPIAvailabilitiesListResponseMessageJSON contains the JSON metadata for the
// struct [SpeedAPIAvailabilitiesListResponseMessage]
type speedAPIAvailabilitiesListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIAvailabilitiesListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIAvailabilitiesListResponseResult struct {
	Quota          SpeedAPIAvailabilitiesListResponseResultQuota    `json:"quota"`
	Regions        []SpeedAPIAvailabilitiesListResponseResultRegion `json:"regions"`
	RegionsPerPlan interface{}                                      `json:"regionsPerPlan"`
	JSON           speedAPIAvailabilitiesListResponseResultJSON     `json:"-"`
}

// speedAPIAvailabilitiesListResponseResultJSON contains the JSON metadata for the
// struct [SpeedAPIAvailabilitiesListResponseResult]
type speedAPIAvailabilitiesListResponseResultJSON struct {
	Quota          apijson.Field
	Regions        apijson.Field
	RegionsPerPlan apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SpeedAPIAvailabilitiesListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIAvailabilitiesListResponseResultQuota struct {
	// Cloudflare plan.
	Plan string `json:"plan"`
	// The number of tests available per plan.
	QuotasPerPlan interface{} `json:"quotasPerPlan"`
	// The number of remaining schedules available.
	RemainingSchedules float64 `json:"remainingSchedules"`
	// The number of remaining tests available.
	RemainingTests float64 `json:"remainingTests"`
	// The number of schedules available per plan.
	ScheduleQuotasPerPlan interface{}                                       `json:"scheduleQuotasPerPlan"`
	JSON                  speedAPIAvailabilitiesListResponseResultQuotaJSON `json:"-"`
}

// speedAPIAvailabilitiesListResponseResultQuotaJSON contains the JSON metadata for
// the struct [SpeedAPIAvailabilitiesListResponseResultQuota]
type speedAPIAvailabilitiesListResponseResultQuotaJSON struct {
	Plan                  apijson.Field
	QuotasPerPlan         apijson.Field
	RemainingSchedules    apijson.Field
	RemainingTests        apijson.Field
	ScheduleQuotasPerPlan apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SpeedAPIAvailabilitiesListResponseResultQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region with a label.
type SpeedAPIAvailabilitiesListResponseResultRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPIAvailabilitiesListResponseResultRegionsValue `json:"value"`
	JSON  speedAPIAvailabilitiesListResponseResultRegionJSON   `json:"-"`
}

// speedAPIAvailabilitiesListResponseResultRegionJSON contains the JSON metadata
// for the struct [SpeedAPIAvailabilitiesListResponseResultRegion]
type speedAPIAvailabilitiesListResponseResultRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIAvailabilitiesListResponseResultRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPIAvailabilitiesListResponseResultRegionsValue string

const (
	SpeedAPIAvailabilitiesListResponseResultRegionsValueAsiaEast1           SpeedAPIAvailabilitiesListResponseResultRegionsValue = "asia-east1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueAsiaNortheast1      SpeedAPIAvailabilitiesListResponseResultRegionsValue = "asia-northeast1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueAsiaNortheast2      SpeedAPIAvailabilitiesListResponseResultRegionsValue = "asia-northeast2"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueAsiaSouth1          SpeedAPIAvailabilitiesListResponseResultRegionsValue = "asia-south1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueAsiaSoutheast1      SpeedAPIAvailabilitiesListResponseResultRegionsValue = "asia-southeast1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueAustraliaSoutheast1 SpeedAPIAvailabilitiesListResponseResultRegionsValue = "australia-southeast1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueEuropeNorth1        SpeedAPIAvailabilitiesListResponseResultRegionsValue = "europe-north1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueEuropeSouthwest1    SpeedAPIAvailabilitiesListResponseResultRegionsValue = "europe-southwest1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueEuropeWest1         SpeedAPIAvailabilitiesListResponseResultRegionsValue = "europe-west1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueEuropeWest2         SpeedAPIAvailabilitiesListResponseResultRegionsValue = "europe-west2"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueEuropeWest3         SpeedAPIAvailabilitiesListResponseResultRegionsValue = "europe-west3"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueEuropeWest4         SpeedAPIAvailabilitiesListResponseResultRegionsValue = "europe-west4"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueEuropeWest8         SpeedAPIAvailabilitiesListResponseResultRegionsValue = "europe-west8"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueEuropeWest9         SpeedAPIAvailabilitiesListResponseResultRegionsValue = "europe-west9"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueMeWest1             SpeedAPIAvailabilitiesListResponseResultRegionsValue = "me-west1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueSouthamericaEast1   SpeedAPIAvailabilitiesListResponseResultRegionsValue = "southamerica-east1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueUsCentral1          SpeedAPIAvailabilitiesListResponseResultRegionsValue = "us-central1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueUsEast1             SpeedAPIAvailabilitiesListResponseResultRegionsValue = "us-east1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueUsEast4             SpeedAPIAvailabilitiesListResponseResultRegionsValue = "us-east4"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueUsSouth1            SpeedAPIAvailabilitiesListResponseResultRegionsValue = "us-south1"
	SpeedAPIAvailabilitiesListResponseResultRegionsValueUsWest1             SpeedAPIAvailabilitiesListResponseResultRegionsValue = "us-west1"
)

type SpeedAPIPagesListResponse struct {
	Errors   []SpeedAPIPagesListResponseError   `json:"errors"`
	Messages []SpeedAPIPagesListResponseMessage `json:"messages"`
	Result   []SpeedAPIPagesListResponseResult  `json:"result"`
	// Whether the API call was successful.
	Success bool                          `json:"success"`
	JSON    speedAPIPagesListResponseJSON `json:"-"`
}

// speedAPIPagesListResponseJSON contains the JSON metadata for the struct
// [SpeedAPIPagesListResponse]
type speedAPIPagesListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIPagesListResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    speedAPIPagesListResponseErrorJSON `json:"-"`
}

// speedAPIPagesListResponseErrorJSON contains the JSON metadata for the struct
// [SpeedAPIPagesListResponseError]
type speedAPIPagesListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIPagesListResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    speedAPIPagesListResponseMessageJSON `json:"-"`
}

// speedAPIPagesListResponseMessageJSON contains the JSON metadata for the struct
// [SpeedAPIPagesListResponseMessage]
type speedAPIPagesListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIPagesListResponseResult struct {
	// A test region with a label.
	Region SpeedAPIPagesListResponseResultRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedAPIPagesListResponseResultScheduleFrequency `json:"scheduleFrequency"`
	Tests             []SpeedAPIPagesListResponseResultTest            `json:"tests"`
	// A URL.
	URL  string                              `json:"url"`
	JSON speedAPIPagesListResponseResultJSON `json:"-"`
}

// speedAPIPagesListResponseResultJSON contains the JSON metadata for the struct
// [SpeedAPIPagesListResponseResult]
type speedAPIPagesListResponseResultJSON struct {
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	Tests             apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region with a label.
type SpeedAPIPagesListResponseResultRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPIPagesListResponseResultRegionValue `json:"value"`
	JSON  speedAPIPagesListResponseResultRegionJSON  `json:"-"`
}

// speedAPIPagesListResponseResultRegionJSON contains the JSON metadata for the
// struct [SpeedAPIPagesListResponseResultRegion]
type speedAPIPagesListResponseResultRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponseResultRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPIPagesListResponseResultRegionValue string

const (
	SpeedAPIPagesListResponseResultRegionValueAsiaEast1           SpeedAPIPagesListResponseResultRegionValue = "asia-east1"
	SpeedAPIPagesListResponseResultRegionValueAsiaNortheast1      SpeedAPIPagesListResponseResultRegionValue = "asia-northeast1"
	SpeedAPIPagesListResponseResultRegionValueAsiaNortheast2      SpeedAPIPagesListResponseResultRegionValue = "asia-northeast2"
	SpeedAPIPagesListResponseResultRegionValueAsiaSouth1          SpeedAPIPagesListResponseResultRegionValue = "asia-south1"
	SpeedAPIPagesListResponseResultRegionValueAsiaSoutheast1      SpeedAPIPagesListResponseResultRegionValue = "asia-southeast1"
	SpeedAPIPagesListResponseResultRegionValueAustraliaSoutheast1 SpeedAPIPagesListResponseResultRegionValue = "australia-southeast1"
	SpeedAPIPagesListResponseResultRegionValueEuropeNorth1        SpeedAPIPagesListResponseResultRegionValue = "europe-north1"
	SpeedAPIPagesListResponseResultRegionValueEuropeSouthwest1    SpeedAPIPagesListResponseResultRegionValue = "europe-southwest1"
	SpeedAPIPagesListResponseResultRegionValueEuropeWest1         SpeedAPIPagesListResponseResultRegionValue = "europe-west1"
	SpeedAPIPagesListResponseResultRegionValueEuropeWest2         SpeedAPIPagesListResponseResultRegionValue = "europe-west2"
	SpeedAPIPagesListResponseResultRegionValueEuropeWest3         SpeedAPIPagesListResponseResultRegionValue = "europe-west3"
	SpeedAPIPagesListResponseResultRegionValueEuropeWest4         SpeedAPIPagesListResponseResultRegionValue = "europe-west4"
	SpeedAPIPagesListResponseResultRegionValueEuropeWest8         SpeedAPIPagesListResponseResultRegionValue = "europe-west8"
	SpeedAPIPagesListResponseResultRegionValueEuropeWest9         SpeedAPIPagesListResponseResultRegionValue = "europe-west9"
	SpeedAPIPagesListResponseResultRegionValueMeWest1             SpeedAPIPagesListResponseResultRegionValue = "me-west1"
	SpeedAPIPagesListResponseResultRegionValueSouthamericaEast1   SpeedAPIPagesListResponseResultRegionValue = "southamerica-east1"
	SpeedAPIPagesListResponseResultRegionValueUsCentral1          SpeedAPIPagesListResponseResultRegionValue = "us-central1"
	SpeedAPIPagesListResponseResultRegionValueUsEast1             SpeedAPIPagesListResponseResultRegionValue = "us-east1"
	SpeedAPIPagesListResponseResultRegionValueUsEast4             SpeedAPIPagesListResponseResultRegionValue = "us-east4"
	SpeedAPIPagesListResponseResultRegionValueUsSouth1            SpeedAPIPagesListResponseResultRegionValue = "us-south1"
	SpeedAPIPagesListResponseResultRegionValueUsWest1             SpeedAPIPagesListResponseResultRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedAPIPagesListResponseResultScheduleFrequency string

const (
	SpeedAPIPagesListResponseResultScheduleFrequencyDaily  SpeedAPIPagesListResponseResultScheduleFrequency = "DAILY"
	SpeedAPIPagesListResponseResultScheduleFrequencyWeekly SpeedAPIPagesListResponseResultScheduleFrequency = "WEEKLY"
)

type SpeedAPIPagesListResponseResultTest struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedAPIPagesListResponseResultTestsDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedAPIPagesListResponseResultTestsMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedAPIPagesListResponseResultTestsRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedAPIPagesListResponseResultTestsScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                                  `json:"url"`
	JSON speedAPIPagesListResponseResultTestJSON `json:"-"`
}

// speedAPIPagesListResponseResultTestJSON contains the JSON metadata for the
// struct [SpeedAPIPagesListResponseResultTest]
type speedAPIPagesListResponseResultTestJSON struct {
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

func (r *SpeedAPIPagesListResponseResultTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedAPIPagesListResponseResultTestsDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPIPagesListResponseResultTestsDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedAPIPagesListResponseResultTestsDesktopReportError      `json:"error"`
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
	State SpeedAPIPagesListResponseResultTestsDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                               `json:"tti"`
	JSON speedAPIPagesListResponseResultTestsDesktopReportJSON `json:"-"`
}

// speedAPIPagesListResponseResultTestsDesktopReportJSON contains the JSON metadata
// for the struct [SpeedAPIPagesListResponseResultTestsDesktopReport]
type speedAPIPagesListResponseResultTestsDesktopReportJSON struct {
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

func (r *SpeedAPIPagesListResponseResultTestsDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPIPagesListResponseResultTestsDesktopReportDeviceType string

const (
	SpeedAPIPagesListResponseResultTestsDesktopReportDeviceTypeDesktop SpeedAPIPagesListResponseResultTestsDesktopReportDeviceType = "DESKTOP"
	SpeedAPIPagesListResponseResultTestsDesktopReportDeviceTypeMobile  SpeedAPIPagesListResponseResultTestsDesktopReportDeviceType = "MOBILE"
)

type SpeedAPIPagesListResponseResultTestsDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPIPagesListResponseResultTestsDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                     `json:"finalDisplayedUrl"`
	JSON              speedAPIPagesListResponseResultTestsDesktopReportErrorJSON `json:"-"`
}

// speedAPIPagesListResponseResultTestsDesktopReportErrorJSON contains the JSON
// metadata for the struct [SpeedAPIPagesListResponseResultTestsDesktopReportError]
type speedAPIPagesListResponseResultTestsDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponseResultTestsDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPIPagesListResponseResultTestsDesktopReportErrorCode string

const (
	SpeedAPIPagesListResponseResultTestsDesktopReportErrorCodeNotReachable      SpeedAPIPagesListResponseResultTestsDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedAPIPagesListResponseResultTestsDesktopReportErrorCodeDNSFailure        SpeedAPIPagesListResponseResultTestsDesktopReportErrorCode = "DNS_FAILURE"
	SpeedAPIPagesListResponseResultTestsDesktopReportErrorCodeNotHTML           SpeedAPIPagesListResponseResultTestsDesktopReportErrorCode = "NOT_HTML"
	SpeedAPIPagesListResponseResultTestsDesktopReportErrorCodeLighthouseTimeout SpeedAPIPagesListResponseResultTestsDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPIPagesListResponseResultTestsDesktopReportErrorCodeUnknown           SpeedAPIPagesListResponseResultTestsDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPIPagesListResponseResultTestsDesktopReportState string

const (
	SpeedAPIPagesListResponseResultTestsDesktopReportStateRunning  SpeedAPIPagesListResponseResultTestsDesktopReportState = "RUNNING"
	SpeedAPIPagesListResponseResultTestsDesktopReportStateComplete SpeedAPIPagesListResponseResultTestsDesktopReportState = "COMPLETE"
	SpeedAPIPagesListResponseResultTestsDesktopReportStateFailed   SpeedAPIPagesListResponseResultTestsDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedAPIPagesListResponseResultTestsMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPIPagesListResponseResultTestsMobileReportDeviceType `json:"deviceType"`
	Error      SpeedAPIPagesListResponseResultTestsMobileReportError      `json:"error"`
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
	State SpeedAPIPagesListResponseResultTestsMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                              `json:"tti"`
	JSON speedAPIPagesListResponseResultTestsMobileReportJSON `json:"-"`
}

// speedAPIPagesListResponseResultTestsMobileReportJSON contains the JSON metadata
// for the struct [SpeedAPIPagesListResponseResultTestsMobileReport]
type speedAPIPagesListResponseResultTestsMobileReportJSON struct {
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

func (r *SpeedAPIPagesListResponseResultTestsMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPIPagesListResponseResultTestsMobileReportDeviceType string

const (
	SpeedAPIPagesListResponseResultTestsMobileReportDeviceTypeDesktop SpeedAPIPagesListResponseResultTestsMobileReportDeviceType = "DESKTOP"
	SpeedAPIPagesListResponseResultTestsMobileReportDeviceTypeMobile  SpeedAPIPagesListResponseResultTestsMobileReportDeviceType = "MOBILE"
)

type SpeedAPIPagesListResponseResultTestsMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPIPagesListResponseResultTestsMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                    `json:"finalDisplayedUrl"`
	JSON              speedAPIPagesListResponseResultTestsMobileReportErrorJSON `json:"-"`
}

// speedAPIPagesListResponseResultTestsMobileReportErrorJSON contains the JSON
// metadata for the struct [SpeedAPIPagesListResponseResultTestsMobileReportError]
type speedAPIPagesListResponseResultTestsMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponseResultTestsMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPIPagesListResponseResultTestsMobileReportErrorCode string

const (
	SpeedAPIPagesListResponseResultTestsMobileReportErrorCodeNotReachable      SpeedAPIPagesListResponseResultTestsMobileReportErrorCode = "NOT_REACHABLE"
	SpeedAPIPagesListResponseResultTestsMobileReportErrorCodeDNSFailure        SpeedAPIPagesListResponseResultTestsMobileReportErrorCode = "DNS_FAILURE"
	SpeedAPIPagesListResponseResultTestsMobileReportErrorCodeNotHTML           SpeedAPIPagesListResponseResultTestsMobileReportErrorCode = "NOT_HTML"
	SpeedAPIPagesListResponseResultTestsMobileReportErrorCodeLighthouseTimeout SpeedAPIPagesListResponseResultTestsMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPIPagesListResponseResultTestsMobileReportErrorCodeUnknown           SpeedAPIPagesListResponseResultTestsMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPIPagesListResponseResultTestsMobileReportState string

const (
	SpeedAPIPagesListResponseResultTestsMobileReportStateRunning  SpeedAPIPagesListResponseResultTestsMobileReportState = "RUNNING"
	SpeedAPIPagesListResponseResultTestsMobileReportStateComplete SpeedAPIPagesListResponseResultTestsMobileReportState = "COMPLETE"
	SpeedAPIPagesListResponseResultTestsMobileReportStateFailed   SpeedAPIPagesListResponseResultTestsMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedAPIPagesListResponseResultTestsRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPIPagesListResponseResultTestsRegionValue `json:"value"`
	JSON  speedAPIPagesListResponseResultTestsRegionJSON  `json:"-"`
}

// speedAPIPagesListResponseResultTestsRegionJSON contains the JSON metadata for
// the struct [SpeedAPIPagesListResponseResultTestsRegion]
type speedAPIPagesListResponseResultTestsRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIPagesListResponseResultTestsRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPIPagesListResponseResultTestsRegionValue string

const (
	SpeedAPIPagesListResponseResultTestsRegionValueAsiaEast1           SpeedAPIPagesListResponseResultTestsRegionValue = "asia-east1"
	SpeedAPIPagesListResponseResultTestsRegionValueAsiaNortheast1      SpeedAPIPagesListResponseResultTestsRegionValue = "asia-northeast1"
	SpeedAPIPagesListResponseResultTestsRegionValueAsiaNortheast2      SpeedAPIPagesListResponseResultTestsRegionValue = "asia-northeast2"
	SpeedAPIPagesListResponseResultTestsRegionValueAsiaSouth1          SpeedAPIPagesListResponseResultTestsRegionValue = "asia-south1"
	SpeedAPIPagesListResponseResultTestsRegionValueAsiaSoutheast1      SpeedAPIPagesListResponseResultTestsRegionValue = "asia-southeast1"
	SpeedAPIPagesListResponseResultTestsRegionValueAustraliaSoutheast1 SpeedAPIPagesListResponseResultTestsRegionValue = "australia-southeast1"
	SpeedAPIPagesListResponseResultTestsRegionValueEuropeNorth1        SpeedAPIPagesListResponseResultTestsRegionValue = "europe-north1"
	SpeedAPIPagesListResponseResultTestsRegionValueEuropeSouthwest1    SpeedAPIPagesListResponseResultTestsRegionValue = "europe-southwest1"
	SpeedAPIPagesListResponseResultTestsRegionValueEuropeWest1         SpeedAPIPagesListResponseResultTestsRegionValue = "europe-west1"
	SpeedAPIPagesListResponseResultTestsRegionValueEuropeWest2         SpeedAPIPagesListResponseResultTestsRegionValue = "europe-west2"
	SpeedAPIPagesListResponseResultTestsRegionValueEuropeWest3         SpeedAPIPagesListResponseResultTestsRegionValue = "europe-west3"
	SpeedAPIPagesListResponseResultTestsRegionValueEuropeWest4         SpeedAPIPagesListResponseResultTestsRegionValue = "europe-west4"
	SpeedAPIPagesListResponseResultTestsRegionValueEuropeWest8         SpeedAPIPagesListResponseResultTestsRegionValue = "europe-west8"
	SpeedAPIPagesListResponseResultTestsRegionValueEuropeWest9         SpeedAPIPagesListResponseResultTestsRegionValue = "europe-west9"
	SpeedAPIPagesListResponseResultTestsRegionValueMeWest1             SpeedAPIPagesListResponseResultTestsRegionValue = "me-west1"
	SpeedAPIPagesListResponseResultTestsRegionValueSouthamericaEast1   SpeedAPIPagesListResponseResultTestsRegionValue = "southamerica-east1"
	SpeedAPIPagesListResponseResultTestsRegionValueUsCentral1          SpeedAPIPagesListResponseResultTestsRegionValue = "us-central1"
	SpeedAPIPagesListResponseResultTestsRegionValueUsEast1             SpeedAPIPagesListResponseResultTestsRegionValue = "us-east1"
	SpeedAPIPagesListResponseResultTestsRegionValueUsEast4             SpeedAPIPagesListResponseResultTestsRegionValue = "us-east4"
	SpeedAPIPagesListResponseResultTestsRegionValueUsSouth1            SpeedAPIPagesListResponseResultTestsRegionValue = "us-south1"
	SpeedAPIPagesListResponseResultTestsRegionValueUsWest1             SpeedAPIPagesListResponseResultTestsRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedAPIPagesListResponseResultTestsScheduleFrequency string

const (
	SpeedAPIPagesListResponseResultTestsScheduleFrequencyDaily  SpeedAPIPagesListResponseResultTestsScheduleFrequency = "DAILY"
	SpeedAPIPagesListResponseResultTestsScheduleFrequencyWeekly SpeedAPIPagesListResponseResultTestsScheduleFrequency = "WEEKLY"
)

type SpeedAPIScheduleDeleteResponse struct {
	Errors   []SpeedAPIScheduleDeleteResponseError   `json:"errors"`
	Messages []SpeedAPIScheduleDeleteResponseMessage `json:"messages"`
	Result   SpeedAPIScheduleDeleteResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                               `json:"success"`
	JSON    speedAPIScheduleDeleteResponseJSON `json:"-"`
}

// speedAPIScheduleDeleteResponseJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleDeleteResponse]
type speedAPIScheduleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    speedAPIScheduleDeleteResponseErrorJSON `json:"-"`
}

// speedAPIScheduleDeleteResponseErrorJSON contains the JSON metadata for the
// struct [SpeedAPIScheduleDeleteResponseError]
type speedAPIScheduleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    speedAPIScheduleDeleteResponseMessageJSON `json:"-"`
}

// speedAPIScheduleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [SpeedAPIScheduleDeleteResponseMessage]
type speedAPIScheduleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleDeleteResponseResult struct {
	// Number of items affected.
	Count float64                                  `json:"count"`
	JSON  speedAPIScheduleDeleteResponseResultJSON `json:"-"`
}

// speedAPIScheduleDeleteResponseResultJSON contains the JSON metadata for the
// struct [SpeedAPIScheduleDeleteResponseResult]
type speedAPIScheduleDeleteResponseResultJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleGetResponse struct {
	Errors   []SpeedAPIScheduleGetResponseError   `json:"errors"`
	Messages []SpeedAPIScheduleGetResponseMessage `json:"messages"`
	// The test schedule.
	Result SpeedAPIScheduleGetResponseResult `json:"result"`
	// Whether the API call was successful.
	Success bool                            `json:"success"`
	JSON    speedAPIScheduleGetResponseJSON `json:"-"`
}

// speedAPIScheduleGetResponseJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleGetResponse]
type speedAPIScheduleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleGetResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    speedAPIScheduleGetResponseErrorJSON `json:"-"`
}

// speedAPIScheduleGetResponseErrorJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleGetResponseError]
type speedAPIScheduleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleGetResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    speedAPIScheduleGetResponseMessageJSON `json:"-"`
}

// speedAPIScheduleGetResponseMessageJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleGetResponseMessage]
type speedAPIScheduleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The test schedule.
type SpeedAPIScheduleGetResponseResult struct {
	// The frequency of the test.
	Frequency SpeedAPIScheduleGetResponseResultFrequency `json:"frequency"`
	// A test region.
	Region SpeedAPIScheduleGetResponseResultRegion `json:"region"`
	// A URL.
	URL  string                                `json:"url"`
	JSON speedAPIScheduleGetResponseResultJSON `json:"-"`
}

// speedAPIScheduleGetResponseResultJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleGetResponseResult]
type speedAPIScheduleGetResponseResultJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency of the test.
type SpeedAPIScheduleGetResponseResultFrequency string

const (
	SpeedAPIScheduleGetResponseResultFrequencyDaily  SpeedAPIScheduleGetResponseResultFrequency = "DAILY"
	SpeedAPIScheduleGetResponseResultFrequencyWeekly SpeedAPIScheduleGetResponseResultFrequency = "WEEKLY"
)

// A test region.
type SpeedAPIScheduleGetResponseResultRegion string

const (
	SpeedAPIScheduleGetResponseResultRegionAsiaEast1           SpeedAPIScheduleGetResponseResultRegion = "asia-east1"
	SpeedAPIScheduleGetResponseResultRegionAsiaNortheast1      SpeedAPIScheduleGetResponseResultRegion = "asia-northeast1"
	SpeedAPIScheduleGetResponseResultRegionAsiaNortheast2      SpeedAPIScheduleGetResponseResultRegion = "asia-northeast2"
	SpeedAPIScheduleGetResponseResultRegionAsiaSouth1          SpeedAPIScheduleGetResponseResultRegion = "asia-south1"
	SpeedAPIScheduleGetResponseResultRegionAsiaSoutheast1      SpeedAPIScheduleGetResponseResultRegion = "asia-southeast1"
	SpeedAPIScheduleGetResponseResultRegionAustraliaSoutheast1 SpeedAPIScheduleGetResponseResultRegion = "australia-southeast1"
	SpeedAPIScheduleGetResponseResultRegionEuropeNorth1        SpeedAPIScheduleGetResponseResultRegion = "europe-north1"
	SpeedAPIScheduleGetResponseResultRegionEuropeSouthwest1    SpeedAPIScheduleGetResponseResultRegion = "europe-southwest1"
	SpeedAPIScheduleGetResponseResultRegionEuropeWest1         SpeedAPIScheduleGetResponseResultRegion = "europe-west1"
	SpeedAPIScheduleGetResponseResultRegionEuropeWest2         SpeedAPIScheduleGetResponseResultRegion = "europe-west2"
	SpeedAPIScheduleGetResponseResultRegionEuropeWest3         SpeedAPIScheduleGetResponseResultRegion = "europe-west3"
	SpeedAPIScheduleGetResponseResultRegionEuropeWest4         SpeedAPIScheduleGetResponseResultRegion = "europe-west4"
	SpeedAPIScheduleGetResponseResultRegionEuropeWest8         SpeedAPIScheduleGetResponseResultRegion = "europe-west8"
	SpeedAPIScheduleGetResponseResultRegionEuropeWest9         SpeedAPIScheduleGetResponseResultRegion = "europe-west9"
	SpeedAPIScheduleGetResponseResultRegionMeWest1             SpeedAPIScheduleGetResponseResultRegion = "me-west1"
	SpeedAPIScheduleGetResponseResultRegionSouthamericaEast1   SpeedAPIScheduleGetResponseResultRegion = "southamerica-east1"
	SpeedAPIScheduleGetResponseResultRegionUsCentral1          SpeedAPIScheduleGetResponseResultRegion = "us-central1"
	SpeedAPIScheduleGetResponseResultRegionUsEast1             SpeedAPIScheduleGetResponseResultRegion = "us-east1"
	SpeedAPIScheduleGetResponseResultRegionUsEast4             SpeedAPIScheduleGetResponseResultRegion = "us-east4"
	SpeedAPIScheduleGetResponseResultRegionUsSouth1            SpeedAPIScheduleGetResponseResultRegion = "us-south1"
	SpeedAPIScheduleGetResponseResultRegionUsWest1             SpeedAPIScheduleGetResponseResultRegion = "us-west1"
)

type SpeedAPITestsNewResponse struct {
	Errors   []SpeedAPITestsNewResponseError   `json:"errors"`
	Messages []SpeedAPITestsNewResponseMessage `json:"messages"`
	Result   SpeedAPITestsNewResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                         `json:"success"`
	JSON    speedAPITestsNewResponseJSON `json:"-"`
}

// speedAPITestsNewResponseJSON contains the JSON metadata for the struct
// [SpeedAPITestsNewResponse]
type speedAPITestsNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsNewResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    speedAPITestsNewResponseErrorJSON `json:"-"`
}

// speedAPITestsNewResponseErrorJSON contains the JSON metadata for the struct
// [SpeedAPITestsNewResponseError]
type speedAPITestsNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsNewResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    speedAPITestsNewResponseMessageJSON `json:"-"`
}

// speedAPITestsNewResponseMessageJSON contains the JSON metadata for the struct
// [SpeedAPITestsNewResponseMessage]
type speedAPITestsNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsNewResponseResult struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedAPITestsNewResponseResultDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedAPITestsNewResponseResultMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedAPITestsNewResponseResultRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedAPITestsNewResponseResultScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                             `json:"url"`
	JSON speedAPITestsNewResponseResultJSON `json:"-"`
}

// speedAPITestsNewResponseResultJSON contains the JSON metadata for the struct
// [SpeedAPITestsNewResponseResult]
type speedAPITestsNewResponseResultJSON struct {
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

func (r *SpeedAPITestsNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedAPITestsNewResponseResultDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPITestsNewResponseResultDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedAPITestsNewResponseResultDesktopReportError      `json:"error"`
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
	State SpeedAPITestsNewResponseResultDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                         `json:"tti"`
	JSON speedAPITestsNewResponseResultDesktopReportJSON `json:"-"`
}

// speedAPITestsNewResponseResultDesktopReportJSON contains the JSON metadata for
// the struct [SpeedAPITestsNewResponseResultDesktopReport]
type speedAPITestsNewResponseResultDesktopReportJSON struct {
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

func (r *SpeedAPITestsNewResponseResultDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPITestsNewResponseResultDesktopReportDeviceType string

const (
	SpeedAPITestsNewResponseResultDesktopReportDeviceTypeDesktop SpeedAPITestsNewResponseResultDesktopReportDeviceType = "DESKTOP"
	SpeedAPITestsNewResponseResultDesktopReportDeviceTypeMobile  SpeedAPITestsNewResponseResultDesktopReportDeviceType = "MOBILE"
)

type SpeedAPITestsNewResponseResultDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPITestsNewResponseResultDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                               `json:"finalDisplayedUrl"`
	JSON              speedAPITestsNewResponseResultDesktopReportErrorJSON `json:"-"`
}

// speedAPITestsNewResponseResultDesktopReportErrorJSON contains the JSON metadata
// for the struct [SpeedAPITestsNewResponseResultDesktopReportError]
type speedAPITestsNewResponseResultDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPITestsNewResponseResultDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPITestsNewResponseResultDesktopReportErrorCode string

const (
	SpeedAPITestsNewResponseResultDesktopReportErrorCodeNotReachable      SpeedAPITestsNewResponseResultDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedAPITestsNewResponseResultDesktopReportErrorCodeDNSFailure        SpeedAPITestsNewResponseResultDesktopReportErrorCode = "DNS_FAILURE"
	SpeedAPITestsNewResponseResultDesktopReportErrorCodeNotHTML           SpeedAPITestsNewResponseResultDesktopReportErrorCode = "NOT_HTML"
	SpeedAPITestsNewResponseResultDesktopReportErrorCodeLighthouseTimeout SpeedAPITestsNewResponseResultDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPITestsNewResponseResultDesktopReportErrorCodeUnknown           SpeedAPITestsNewResponseResultDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPITestsNewResponseResultDesktopReportState string

const (
	SpeedAPITestsNewResponseResultDesktopReportStateRunning  SpeedAPITestsNewResponseResultDesktopReportState = "RUNNING"
	SpeedAPITestsNewResponseResultDesktopReportStateComplete SpeedAPITestsNewResponseResultDesktopReportState = "COMPLETE"
	SpeedAPITestsNewResponseResultDesktopReportStateFailed   SpeedAPITestsNewResponseResultDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedAPITestsNewResponseResultMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPITestsNewResponseResultMobileReportDeviceType `json:"deviceType"`
	Error      SpeedAPITestsNewResponseResultMobileReportError      `json:"error"`
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
	State SpeedAPITestsNewResponseResultMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                        `json:"tti"`
	JSON speedAPITestsNewResponseResultMobileReportJSON `json:"-"`
}

// speedAPITestsNewResponseResultMobileReportJSON contains the JSON metadata for
// the struct [SpeedAPITestsNewResponseResultMobileReport]
type speedAPITestsNewResponseResultMobileReportJSON struct {
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

func (r *SpeedAPITestsNewResponseResultMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPITestsNewResponseResultMobileReportDeviceType string

const (
	SpeedAPITestsNewResponseResultMobileReportDeviceTypeDesktop SpeedAPITestsNewResponseResultMobileReportDeviceType = "DESKTOP"
	SpeedAPITestsNewResponseResultMobileReportDeviceTypeMobile  SpeedAPITestsNewResponseResultMobileReportDeviceType = "MOBILE"
)

type SpeedAPITestsNewResponseResultMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPITestsNewResponseResultMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                              `json:"finalDisplayedUrl"`
	JSON              speedAPITestsNewResponseResultMobileReportErrorJSON `json:"-"`
}

// speedAPITestsNewResponseResultMobileReportErrorJSON contains the JSON metadata
// for the struct [SpeedAPITestsNewResponseResultMobileReportError]
type speedAPITestsNewResponseResultMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPITestsNewResponseResultMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPITestsNewResponseResultMobileReportErrorCode string

const (
	SpeedAPITestsNewResponseResultMobileReportErrorCodeNotReachable      SpeedAPITestsNewResponseResultMobileReportErrorCode = "NOT_REACHABLE"
	SpeedAPITestsNewResponseResultMobileReportErrorCodeDNSFailure        SpeedAPITestsNewResponseResultMobileReportErrorCode = "DNS_FAILURE"
	SpeedAPITestsNewResponseResultMobileReportErrorCodeNotHTML           SpeedAPITestsNewResponseResultMobileReportErrorCode = "NOT_HTML"
	SpeedAPITestsNewResponseResultMobileReportErrorCodeLighthouseTimeout SpeedAPITestsNewResponseResultMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPITestsNewResponseResultMobileReportErrorCodeUnknown           SpeedAPITestsNewResponseResultMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPITestsNewResponseResultMobileReportState string

const (
	SpeedAPITestsNewResponseResultMobileReportStateRunning  SpeedAPITestsNewResponseResultMobileReportState = "RUNNING"
	SpeedAPITestsNewResponseResultMobileReportStateComplete SpeedAPITestsNewResponseResultMobileReportState = "COMPLETE"
	SpeedAPITestsNewResponseResultMobileReportStateFailed   SpeedAPITestsNewResponseResultMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedAPITestsNewResponseResultRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPITestsNewResponseResultRegionValue `json:"value"`
	JSON  speedAPITestsNewResponseResultRegionJSON  `json:"-"`
}

// speedAPITestsNewResponseResultRegionJSON contains the JSON metadata for the
// struct [SpeedAPITestsNewResponseResultRegion]
type speedAPITestsNewResponseResultRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsNewResponseResultRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPITestsNewResponseResultRegionValue string

const (
	SpeedAPITestsNewResponseResultRegionValueAsiaEast1           SpeedAPITestsNewResponseResultRegionValue = "asia-east1"
	SpeedAPITestsNewResponseResultRegionValueAsiaNortheast1      SpeedAPITestsNewResponseResultRegionValue = "asia-northeast1"
	SpeedAPITestsNewResponseResultRegionValueAsiaNortheast2      SpeedAPITestsNewResponseResultRegionValue = "asia-northeast2"
	SpeedAPITestsNewResponseResultRegionValueAsiaSouth1          SpeedAPITestsNewResponseResultRegionValue = "asia-south1"
	SpeedAPITestsNewResponseResultRegionValueAsiaSoutheast1      SpeedAPITestsNewResponseResultRegionValue = "asia-southeast1"
	SpeedAPITestsNewResponseResultRegionValueAustraliaSoutheast1 SpeedAPITestsNewResponseResultRegionValue = "australia-southeast1"
	SpeedAPITestsNewResponseResultRegionValueEuropeNorth1        SpeedAPITestsNewResponseResultRegionValue = "europe-north1"
	SpeedAPITestsNewResponseResultRegionValueEuropeSouthwest1    SpeedAPITestsNewResponseResultRegionValue = "europe-southwest1"
	SpeedAPITestsNewResponseResultRegionValueEuropeWest1         SpeedAPITestsNewResponseResultRegionValue = "europe-west1"
	SpeedAPITestsNewResponseResultRegionValueEuropeWest2         SpeedAPITestsNewResponseResultRegionValue = "europe-west2"
	SpeedAPITestsNewResponseResultRegionValueEuropeWest3         SpeedAPITestsNewResponseResultRegionValue = "europe-west3"
	SpeedAPITestsNewResponseResultRegionValueEuropeWest4         SpeedAPITestsNewResponseResultRegionValue = "europe-west4"
	SpeedAPITestsNewResponseResultRegionValueEuropeWest8         SpeedAPITestsNewResponseResultRegionValue = "europe-west8"
	SpeedAPITestsNewResponseResultRegionValueEuropeWest9         SpeedAPITestsNewResponseResultRegionValue = "europe-west9"
	SpeedAPITestsNewResponseResultRegionValueMeWest1             SpeedAPITestsNewResponseResultRegionValue = "me-west1"
	SpeedAPITestsNewResponseResultRegionValueSouthamericaEast1   SpeedAPITestsNewResponseResultRegionValue = "southamerica-east1"
	SpeedAPITestsNewResponseResultRegionValueUsCentral1          SpeedAPITestsNewResponseResultRegionValue = "us-central1"
	SpeedAPITestsNewResponseResultRegionValueUsEast1             SpeedAPITestsNewResponseResultRegionValue = "us-east1"
	SpeedAPITestsNewResponseResultRegionValueUsEast4             SpeedAPITestsNewResponseResultRegionValue = "us-east4"
	SpeedAPITestsNewResponseResultRegionValueUsSouth1            SpeedAPITestsNewResponseResultRegionValue = "us-south1"
	SpeedAPITestsNewResponseResultRegionValueUsWest1             SpeedAPITestsNewResponseResultRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedAPITestsNewResponseResultScheduleFrequency string

const (
	SpeedAPITestsNewResponseResultScheduleFrequencyDaily  SpeedAPITestsNewResponseResultScheduleFrequency = "DAILY"
	SpeedAPITestsNewResponseResultScheduleFrequencyWeekly SpeedAPITestsNewResponseResultScheduleFrequency = "WEEKLY"
)

type SpeedAPITestsDeleteResponse struct {
	Errors   []SpeedAPITestsDeleteResponseError   `json:"errors"`
	Messages []SpeedAPITestsDeleteResponseMessage `json:"messages"`
	Result   SpeedAPITestsDeleteResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                            `json:"success"`
	JSON    speedAPITestsDeleteResponseJSON `json:"-"`
}

// speedAPITestsDeleteResponseJSON contains the JSON metadata for the struct
// [SpeedAPITestsDeleteResponse]
type speedAPITestsDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsDeleteResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    speedAPITestsDeleteResponseErrorJSON `json:"-"`
}

// speedAPITestsDeleteResponseErrorJSON contains the JSON metadata for the struct
// [SpeedAPITestsDeleteResponseError]
type speedAPITestsDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsDeleteResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    speedAPITestsDeleteResponseMessageJSON `json:"-"`
}

// speedAPITestsDeleteResponseMessageJSON contains the JSON metadata for the struct
// [SpeedAPITestsDeleteResponseMessage]
type speedAPITestsDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsDeleteResponseResult struct {
	// Number of items affected.
	Count float64                               `json:"count"`
	JSON  speedAPITestsDeleteResponseResultJSON `json:"-"`
}

// speedAPITestsDeleteResponseResultJSON contains the JSON metadata for the struct
// [SpeedAPITestsDeleteResponseResult]
type speedAPITestsDeleteResponseResultJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsListResponse struct {
	Errors     []SpeedAPITestsListResponseError    `json:"errors"`
	Messages   []SpeedAPITestsListResponseMessage  `json:"messages"`
	Result     []SpeedAPITestsListResponseResult   `json:"result"`
	ResultInfo SpeedAPITestsListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success bool                          `json:"success"`
	JSON    speedAPITestsListResponseJSON `json:"-"`
}

// speedAPITestsListResponseJSON contains the JSON metadata for the struct
// [SpeedAPITestsListResponse]
type speedAPITestsListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
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

type SpeedAPITestsListResponseResult struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedAPITestsListResponseResultDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedAPITestsListResponseResultMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedAPITestsListResponseResultRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedAPITestsListResponseResultScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                              `json:"url"`
	JSON speedAPITestsListResponseResultJSON `json:"-"`
}

// speedAPITestsListResponseResultJSON contains the JSON metadata for the struct
// [SpeedAPITestsListResponseResult]
type speedAPITestsListResponseResultJSON struct {
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

func (r *SpeedAPITestsListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedAPITestsListResponseResultDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPITestsListResponseResultDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedAPITestsListResponseResultDesktopReportError      `json:"error"`
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
	State SpeedAPITestsListResponseResultDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                          `json:"tti"`
	JSON speedAPITestsListResponseResultDesktopReportJSON `json:"-"`
}

// speedAPITestsListResponseResultDesktopReportJSON contains the JSON metadata for
// the struct [SpeedAPITestsListResponseResultDesktopReport]
type speedAPITestsListResponseResultDesktopReportJSON struct {
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

func (r *SpeedAPITestsListResponseResultDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPITestsListResponseResultDesktopReportDeviceType string

const (
	SpeedAPITestsListResponseResultDesktopReportDeviceTypeDesktop SpeedAPITestsListResponseResultDesktopReportDeviceType = "DESKTOP"
	SpeedAPITestsListResponseResultDesktopReportDeviceTypeMobile  SpeedAPITestsListResponseResultDesktopReportDeviceType = "MOBILE"
)

type SpeedAPITestsListResponseResultDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPITestsListResponseResultDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                `json:"finalDisplayedUrl"`
	JSON              speedAPITestsListResponseResultDesktopReportErrorJSON `json:"-"`
}

// speedAPITestsListResponseResultDesktopReportErrorJSON contains the JSON metadata
// for the struct [SpeedAPITestsListResponseResultDesktopReportError]
type speedAPITestsListResponseResultDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPITestsListResponseResultDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPITestsListResponseResultDesktopReportErrorCode string

const (
	SpeedAPITestsListResponseResultDesktopReportErrorCodeNotReachable      SpeedAPITestsListResponseResultDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedAPITestsListResponseResultDesktopReportErrorCodeDNSFailure        SpeedAPITestsListResponseResultDesktopReportErrorCode = "DNS_FAILURE"
	SpeedAPITestsListResponseResultDesktopReportErrorCodeNotHTML           SpeedAPITestsListResponseResultDesktopReportErrorCode = "NOT_HTML"
	SpeedAPITestsListResponseResultDesktopReportErrorCodeLighthouseTimeout SpeedAPITestsListResponseResultDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPITestsListResponseResultDesktopReportErrorCodeUnknown           SpeedAPITestsListResponseResultDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPITestsListResponseResultDesktopReportState string

const (
	SpeedAPITestsListResponseResultDesktopReportStateRunning  SpeedAPITestsListResponseResultDesktopReportState = "RUNNING"
	SpeedAPITestsListResponseResultDesktopReportStateComplete SpeedAPITestsListResponseResultDesktopReportState = "COMPLETE"
	SpeedAPITestsListResponseResultDesktopReportStateFailed   SpeedAPITestsListResponseResultDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedAPITestsListResponseResultMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPITestsListResponseResultMobileReportDeviceType `json:"deviceType"`
	Error      SpeedAPITestsListResponseResultMobileReportError      `json:"error"`
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
	State SpeedAPITestsListResponseResultMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                         `json:"tti"`
	JSON speedAPITestsListResponseResultMobileReportJSON `json:"-"`
}

// speedAPITestsListResponseResultMobileReportJSON contains the JSON metadata for
// the struct [SpeedAPITestsListResponseResultMobileReport]
type speedAPITestsListResponseResultMobileReportJSON struct {
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

func (r *SpeedAPITestsListResponseResultMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPITestsListResponseResultMobileReportDeviceType string

const (
	SpeedAPITestsListResponseResultMobileReportDeviceTypeDesktop SpeedAPITestsListResponseResultMobileReportDeviceType = "DESKTOP"
	SpeedAPITestsListResponseResultMobileReportDeviceTypeMobile  SpeedAPITestsListResponseResultMobileReportDeviceType = "MOBILE"
)

type SpeedAPITestsListResponseResultMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPITestsListResponseResultMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                               `json:"finalDisplayedUrl"`
	JSON              speedAPITestsListResponseResultMobileReportErrorJSON `json:"-"`
}

// speedAPITestsListResponseResultMobileReportErrorJSON contains the JSON metadata
// for the struct [SpeedAPITestsListResponseResultMobileReportError]
type speedAPITestsListResponseResultMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPITestsListResponseResultMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPITestsListResponseResultMobileReportErrorCode string

const (
	SpeedAPITestsListResponseResultMobileReportErrorCodeNotReachable      SpeedAPITestsListResponseResultMobileReportErrorCode = "NOT_REACHABLE"
	SpeedAPITestsListResponseResultMobileReportErrorCodeDNSFailure        SpeedAPITestsListResponseResultMobileReportErrorCode = "DNS_FAILURE"
	SpeedAPITestsListResponseResultMobileReportErrorCodeNotHTML           SpeedAPITestsListResponseResultMobileReportErrorCode = "NOT_HTML"
	SpeedAPITestsListResponseResultMobileReportErrorCodeLighthouseTimeout SpeedAPITestsListResponseResultMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPITestsListResponseResultMobileReportErrorCodeUnknown           SpeedAPITestsListResponseResultMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPITestsListResponseResultMobileReportState string

const (
	SpeedAPITestsListResponseResultMobileReportStateRunning  SpeedAPITestsListResponseResultMobileReportState = "RUNNING"
	SpeedAPITestsListResponseResultMobileReportStateComplete SpeedAPITestsListResponseResultMobileReportState = "COMPLETE"
	SpeedAPITestsListResponseResultMobileReportStateFailed   SpeedAPITestsListResponseResultMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedAPITestsListResponseResultRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPITestsListResponseResultRegionValue `json:"value"`
	JSON  speedAPITestsListResponseResultRegionJSON  `json:"-"`
}

// speedAPITestsListResponseResultRegionJSON contains the JSON metadata for the
// struct [SpeedAPITestsListResponseResultRegion]
type speedAPITestsListResponseResultRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsListResponseResultRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPITestsListResponseResultRegionValue string

const (
	SpeedAPITestsListResponseResultRegionValueAsiaEast1           SpeedAPITestsListResponseResultRegionValue = "asia-east1"
	SpeedAPITestsListResponseResultRegionValueAsiaNortheast1      SpeedAPITestsListResponseResultRegionValue = "asia-northeast1"
	SpeedAPITestsListResponseResultRegionValueAsiaNortheast2      SpeedAPITestsListResponseResultRegionValue = "asia-northeast2"
	SpeedAPITestsListResponseResultRegionValueAsiaSouth1          SpeedAPITestsListResponseResultRegionValue = "asia-south1"
	SpeedAPITestsListResponseResultRegionValueAsiaSoutheast1      SpeedAPITestsListResponseResultRegionValue = "asia-southeast1"
	SpeedAPITestsListResponseResultRegionValueAustraliaSoutheast1 SpeedAPITestsListResponseResultRegionValue = "australia-southeast1"
	SpeedAPITestsListResponseResultRegionValueEuropeNorth1        SpeedAPITestsListResponseResultRegionValue = "europe-north1"
	SpeedAPITestsListResponseResultRegionValueEuropeSouthwest1    SpeedAPITestsListResponseResultRegionValue = "europe-southwest1"
	SpeedAPITestsListResponseResultRegionValueEuropeWest1         SpeedAPITestsListResponseResultRegionValue = "europe-west1"
	SpeedAPITestsListResponseResultRegionValueEuropeWest2         SpeedAPITestsListResponseResultRegionValue = "europe-west2"
	SpeedAPITestsListResponseResultRegionValueEuropeWest3         SpeedAPITestsListResponseResultRegionValue = "europe-west3"
	SpeedAPITestsListResponseResultRegionValueEuropeWest4         SpeedAPITestsListResponseResultRegionValue = "europe-west4"
	SpeedAPITestsListResponseResultRegionValueEuropeWest8         SpeedAPITestsListResponseResultRegionValue = "europe-west8"
	SpeedAPITestsListResponseResultRegionValueEuropeWest9         SpeedAPITestsListResponseResultRegionValue = "europe-west9"
	SpeedAPITestsListResponseResultRegionValueMeWest1             SpeedAPITestsListResponseResultRegionValue = "me-west1"
	SpeedAPITestsListResponseResultRegionValueSouthamericaEast1   SpeedAPITestsListResponseResultRegionValue = "southamerica-east1"
	SpeedAPITestsListResponseResultRegionValueUsCentral1          SpeedAPITestsListResponseResultRegionValue = "us-central1"
	SpeedAPITestsListResponseResultRegionValueUsEast1             SpeedAPITestsListResponseResultRegionValue = "us-east1"
	SpeedAPITestsListResponseResultRegionValueUsEast4             SpeedAPITestsListResponseResultRegionValue = "us-east4"
	SpeedAPITestsListResponseResultRegionValueUsSouth1            SpeedAPITestsListResponseResultRegionValue = "us-south1"
	SpeedAPITestsListResponseResultRegionValueUsWest1             SpeedAPITestsListResponseResultRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedAPITestsListResponseResultScheduleFrequency string

const (
	SpeedAPITestsListResponseResultScheduleFrequencyDaily  SpeedAPITestsListResponseResultScheduleFrequency = "DAILY"
	SpeedAPITestsListResponseResultScheduleFrequencyWeekly SpeedAPITestsListResponseResultScheduleFrequency = "WEEKLY"
)

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
	Errors   []SpeedAPITestsGetResponseError   `json:"errors"`
	Messages []SpeedAPITestsGetResponseMessage `json:"messages"`
	Result   SpeedAPITestsGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                         `json:"success"`
	JSON    speedAPITestsGetResponseJSON `json:"-"`
}

// speedAPITestsGetResponseJSON contains the JSON metadata for the struct
// [SpeedAPITestsGetResponse]
type speedAPITestsGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsGetResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    speedAPITestsGetResponseErrorJSON `json:"-"`
}

// speedAPITestsGetResponseErrorJSON contains the JSON metadata for the struct
// [SpeedAPITestsGetResponseError]
type speedAPITestsGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsGetResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    speedAPITestsGetResponseMessageJSON `json:"-"`
}

// speedAPITestsGetResponseMessageJSON contains the JSON metadata for the struct
// [SpeedAPITestsGetResponseMessage]
type speedAPITestsGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITestsGetResponseResult struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedAPITestsGetResponseResultDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedAPITestsGetResponseResultMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedAPITestsGetResponseResultRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedAPITestsGetResponseResultScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                             `json:"url"`
	JSON speedAPITestsGetResponseResultJSON `json:"-"`
}

// speedAPITestsGetResponseResultJSON contains the JSON metadata for the struct
// [SpeedAPITestsGetResponseResult]
type speedAPITestsGetResponseResultJSON struct {
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

func (r *SpeedAPITestsGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedAPITestsGetResponseResultDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPITestsGetResponseResultDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedAPITestsGetResponseResultDesktopReportError      `json:"error"`
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
	State SpeedAPITestsGetResponseResultDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                         `json:"tti"`
	JSON speedAPITestsGetResponseResultDesktopReportJSON `json:"-"`
}

// speedAPITestsGetResponseResultDesktopReportJSON contains the JSON metadata for
// the struct [SpeedAPITestsGetResponseResultDesktopReport]
type speedAPITestsGetResponseResultDesktopReportJSON struct {
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

func (r *SpeedAPITestsGetResponseResultDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPITestsGetResponseResultDesktopReportDeviceType string

const (
	SpeedAPITestsGetResponseResultDesktopReportDeviceTypeDesktop SpeedAPITestsGetResponseResultDesktopReportDeviceType = "DESKTOP"
	SpeedAPITestsGetResponseResultDesktopReportDeviceTypeMobile  SpeedAPITestsGetResponseResultDesktopReportDeviceType = "MOBILE"
)

type SpeedAPITestsGetResponseResultDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPITestsGetResponseResultDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                               `json:"finalDisplayedUrl"`
	JSON              speedAPITestsGetResponseResultDesktopReportErrorJSON `json:"-"`
}

// speedAPITestsGetResponseResultDesktopReportErrorJSON contains the JSON metadata
// for the struct [SpeedAPITestsGetResponseResultDesktopReportError]
type speedAPITestsGetResponseResultDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPITestsGetResponseResultDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPITestsGetResponseResultDesktopReportErrorCode string

const (
	SpeedAPITestsGetResponseResultDesktopReportErrorCodeNotReachable      SpeedAPITestsGetResponseResultDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedAPITestsGetResponseResultDesktopReportErrorCodeDNSFailure        SpeedAPITestsGetResponseResultDesktopReportErrorCode = "DNS_FAILURE"
	SpeedAPITestsGetResponseResultDesktopReportErrorCodeNotHTML           SpeedAPITestsGetResponseResultDesktopReportErrorCode = "NOT_HTML"
	SpeedAPITestsGetResponseResultDesktopReportErrorCodeLighthouseTimeout SpeedAPITestsGetResponseResultDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPITestsGetResponseResultDesktopReportErrorCodeUnknown           SpeedAPITestsGetResponseResultDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPITestsGetResponseResultDesktopReportState string

const (
	SpeedAPITestsGetResponseResultDesktopReportStateRunning  SpeedAPITestsGetResponseResultDesktopReportState = "RUNNING"
	SpeedAPITestsGetResponseResultDesktopReportStateComplete SpeedAPITestsGetResponseResultDesktopReportState = "COMPLETE"
	SpeedAPITestsGetResponseResultDesktopReportStateFailed   SpeedAPITestsGetResponseResultDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedAPITestsGetResponseResultMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPITestsGetResponseResultMobileReportDeviceType `json:"deviceType"`
	Error      SpeedAPITestsGetResponseResultMobileReportError      `json:"error"`
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
	State SpeedAPITestsGetResponseResultMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                        `json:"tti"`
	JSON speedAPITestsGetResponseResultMobileReportJSON `json:"-"`
}

// speedAPITestsGetResponseResultMobileReportJSON contains the JSON metadata for
// the struct [SpeedAPITestsGetResponseResultMobileReport]
type speedAPITestsGetResponseResultMobileReportJSON struct {
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

func (r *SpeedAPITestsGetResponseResultMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPITestsGetResponseResultMobileReportDeviceType string

const (
	SpeedAPITestsGetResponseResultMobileReportDeviceTypeDesktop SpeedAPITestsGetResponseResultMobileReportDeviceType = "DESKTOP"
	SpeedAPITestsGetResponseResultMobileReportDeviceTypeMobile  SpeedAPITestsGetResponseResultMobileReportDeviceType = "MOBILE"
)

type SpeedAPITestsGetResponseResultMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPITestsGetResponseResultMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                              `json:"finalDisplayedUrl"`
	JSON              speedAPITestsGetResponseResultMobileReportErrorJSON `json:"-"`
}

// speedAPITestsGetResponseResultMobileReportErrorJSON contains the JSON metadata
// for the struct [SpeedAPITestsGetResponseResultMobileReportError]
type speedAPITestsGetResponseResultMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPITestsGetResponseResultMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPITestsGetResponseResultMobileReportErrorCode string

const (
	SpeedAPITestsGetResponseResultMobileReportErrorCodeNotReachable      SpeedAPITestsGetResponseResultMobileReportErrorCode = "NOT_REACHABLE"
	SpeedAPITestsGetResponseResultMobileReportErrorCodeDNSFailure        SpeedAPITestsGetResponseResultMobileReportErrorCode = "DNS_FAILURE"
	SpeedAPITestsGetResponseResultMobileReportErrorCodeNotHTML           SpeedAPITestsGetResponseResultMobileReportErrorCode = "NOT_HTML"
	SpeedAPITestsGetResponseResultMobileReportErrorCodeLighthouseTimeout SpeedAPITestsGetResponseResultMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPITestsGetResponseResultMobileReportErrorCodeUnknown           SpeedAPITestsGetResponseResultMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPITestsGetResponseResultMobileReportState string

const (
	SpeedAPITestsGetResponseResultMobileReportStateRunning  SpeedAPITestsGetResponseResultMobileReportState = "RUNNING"
	SpeedAPITestsGetResponseResultMobileReportStateComplete SpeedAPITestsGetResponseResultMobileReportState = "COMPLETE"
	SpeedAPITestsGetResponseResultMobileReportStateFailed   SpeedAPITestsGetResponseResultMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedAPITestsGetResponseResultRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPITestsGetResponseResultRegionValue `json:"value"`
	JSON  speedAPITestsGetResponseResultRegionJSON  `json:"-"`
}

// speedAPITestsGetResponseResultRegionJSON contains the JSON metadata for the
// struct [SpeedAPITestsGetResponseResultRegion]
type speedAPITestsGetResponseResultRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITestsGetResponseResultRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPITestsGetResponseResultRegionValue string

const (
	SpeedAPITestsGetResponseResultRegionValueAsiaEast1           SpeedAPITestsGetResponseResultRegionValue = "asia-east1"
	SpeedAPITestsGetResponseResultRegionValueAsiaNortheast1      SpeedAPITestsGetResponseResultRegionValue = "asia-northeast1"
	SpeedAPITestsGetResponseResultRegionValueAsiaNortheast2      SpeedAPITestsGetResponseResultRegionValue = "asia-northeast2"
	SpeedAPITestsGetResponseResultRegionValueAsiaSouth1          SpeedAPITestsGetResponseResultRegionValue = "asia-south1"
	SpeedAPITestsGetResponseResultRegionValueAsiaSoutheast1      SpeedAPITestsGetResponseResultRegionValue = "asia-southeast1"
	SpeedAPITestsGetResponseResultRegionValueAustraliaSoutheast1 SpeedAPITestsGetResponseResultRegionValue = "australia-southeast1"
	SpeedAPITestsGetResponseResultRegionValueEuropeNorth1        SpeedAPITestsGetResponseResultRegionValue = "europe-north1"
	SpeedAPITestsGetResponseResultRegionValueEuropeSouthwest1    SpeedAPITestsGetResponseResultRegionValue = "europe-southwest1"
	SpeedAPITestsGetResponseResultRegionValueEuropeWest1         SpeedAPITestsGetResponseResultRegionValue = "europe-west1"
	SpeedAPITestsGetResponseResultRegionValueEuropeWest2         SpeedAPITestsGetResponseResultRegionValue = "europe-west2"
	SpeedAPITestsGetResponseResultRegionValueEuropeWest3         SpeedAPITestsGetResponseResultRegionValue = "europe-west3"
	SpeedAPITestsGetResponseResultRegionValueEuropeWest4         SpeedAPITestsGetResponseResultRegionValue = "europe-west4"
	SpeedAPITestsGetResponseResultRegionValueEuropeWest8         SpeedAPITestsGetResponseResultRegionValue = "europe-west8"
	SpeedAPITestsGetResponseResultRegionValueEuropeWest9         SpeedAPITestsGetResponseResultRegionValue = "europe-west9"
	SpeedAPITestsGetResponseResultRegionValueMeWest1             SpeedAPITestsGetResponseResultRegionValue = "me-west1"
	SpeedAPITestsGetResponseResultRegionValueSouthamericaEast1   SpeedAPITestsGetResponseResultRegionValue = "southamerica-east1"
	SpeedAPITestsGetResponseResultRegionValueUsCentral1          SpeedAPITestsGetResponseResultRegionValue = "us-central1"
	SpeedAPITestsGetResponseResultRegionValueUsEast1             SpeedAPITestsGetResponseResultRegionValue = "us-east1"
	SpeedAPITestsGetResponseResultRegionValueUsEast4             SpeedAPITestsGetResponseResultRegionValue = "us-east4"
	SpeedAPITestsGetResponseResultRegionValueUsSouth1            SpeedAPITestsGetResponseResultRegionValue = "us-south1"
	SpeedAPITestsGetResponseResultRegionValueUsWest1             SpeedAPITestsGetResponseResultRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedAPITestsGetResponseResultScheduleFrequency string

const (
	SpeedAPITestsGetResponseResultScheduleFrequencyDaily  SpeedAPITestsGetResponseResultScheduleFrequency = "DAILY"
	SpeedAPITestsGetResponseResultScheduleFrequencyWeekly SpeedAPITestsGetResponseResultScheduleFrequency = "WEEKLY"
)

type SpeedAPITrendsListResponse struct {
	Errors   []SpeedAPITrendsListResponseError   `json:"errors"`
	Messages []SpeedAPITrendsListResponseMessage `json:"messages"`
	Result   SpeedAPITrendsListResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                           `json:"success"`
	JSON    speedAPITrendsListResponseJSON `json:"-"`
}

// speedAPITrendsListResponseJSON contains the JSON metadata for the struct
// [SpeedAPITrendsListResponse]
type speedAPITrendsListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITrendsListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITrendsListResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    speedAPITrendsListResponseErrorJSON `json:"-"`
}

// speedAPITrendsListResponseErrorJSON contains the JSON metadata for the struct
// [SpeedAPITrendsListResponseError]
type speedAPITrendsListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITrendsListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITrendsListResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    speedAPITrendsListResponseMessageJSON `json:"-"`
}

// speedAPITrendsListResponseMessageJSON contains the JSON metadata for the struct
// [SpeedAPITrendsListResponseMessage]
type speedAPITrendsListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPITrendsListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPITrendsListResponseResult struct {
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
	Tti  []float64                            `json:"tti"`
	JSON speedAPITrendsListResponseResultJSON `json:"-"`
}

// speedAPITrendsListResponseResultJSON contains the JSON metadata for the struct
// [SpeedAPITrendsListResponseResult]
type speedAPITrendsListResponseResultJSON struct {
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

func (r *SpeedAPITrendsListResponseResult) UnmarshalJSON(data []byte) (err error) {
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
