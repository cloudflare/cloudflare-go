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

// SpeedAPIScheduleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSpeedAPIScheduleService] method
// instead.
type SpeedAPIScheduleService struct {
	Options []option.RequestOption
}

// NewSpeedAPIScheduleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSpeedAPIScheduleService(opts ...option.RequestOption) (r *SpeedAPIScheduleService) {
	r = &SpeedAPIScheduleService{}
	r.Options = opts
	return
}

// Creates a scheduled test for a page.
func (r *SpeedAPIScheduleService) New(ctx context.Context, zoneID string, url string, body SpeedAPIScheduleNewParams, opts ...option.RequestOption) (res *SpeedAPIScheduleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SpeedAPIScheduleNewResponse struct {
	Errors   []SpeedAPIScheduleNewResponseError   `json:"errors"`
	Messages []SpeedAPIScheduleNewResponseMessage `json:"messages"`
	Result   SpeedAPIScheduleNewResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                            `json:"success"`
	JSON    speedAPIScheduleNewResponseJSON `json:"-"`
}

// speedAPIScheduleNewResponseJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleNewResponse]
type speedAPIScheduleNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleNewResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    speedAPIScheduleNewResponseErrorJSON `json:"-"`
}

// speedAPIScheduleNewResponseErrorJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleNewResponseError]
type speedAPIScheduleNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleNewResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    speedAPIScheduleNewResponseMessageJSON `json:"-"`
}

// speedAPIScheduleNewResponseMessageJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleNewResponseMessage]
type speedAPIScheduleNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleNewResponseResult struct {
	// The test schedule.
	Schedule SpeedAPIScheduleNewResponseResultSchedule `json:"schedule"`
	Test     SpeedAPIScheduleNewResponseResultTest     `json:"test"`
	JSON     speedAPIScheduleNewResponseResultJSON     `json:"-"`
}

// speedAPIScheduleNewResponseResultJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleNewResponseResult]
type speedAPIScheduleNewResponseResultJSON struct {
	Schedule    apijson.Field
	Test        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The test schedule.
type SpeedAPIScheduleNewResponseResultSchedule struct {
	// The frequency of the test.
	Frequency SpeedAPIScheduleNewResponseResultScheduleFrequency `json:"frequency"`
	// A test region.
	Region SpeedAPIScheduleNewResponseResultScheduleRegion `json:"region"`
	// A URL.
	URL  string                                        `json:"url"`
	JSON speedAPIScheduleNewResponseResultScheduleJSON `json:"-"`
}

// speedAPIScheduleNewResponseResultScheduleJSON contains the JSON metadata for the
// struct [SpeedAPIScheduleNewResponseResultSchedule]
type speedAPIScheduleNewResponseResultScheduleJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency of the test.
type SpeedAPIScheduleNewResponseResultScheduleFrequency string

const (
	SpeedAPIScheduleNewResponseResultScheduleFrequencyDaily  SpeedAPIScheduleNewResponseResultScheduleFrequency = "DAILY"
	SpeedAPIScheduleNewResponseResultScheduleFrequencyWeekly SpeedAPIScheduleNewResponseResultScheduleFrequency = "WEEKLY"
)

// A test region.
type SpeedAPIScheduleNewResponseResultScheduleRegion string

const (
	SpeedAPIScheduleNewResponseResultScheduleRegionAsiaEast1           SpeedAPIScheduleNewResponseResultScheduleRegion = "asia-east1"
	SpeedAPIScheduleNewResponseResultScheduleRegionAsiaNortheast1      SpeedAPIScheduleNewResponseResultScheduleRegion = "asia-northeast1"
	SpeedAPIScheduleNewResponseResultScheduleRegionAsiaNortheast2      SpeedAPIScheduleNewResponseResultScheduleRegion = "asia-northeast2"
	SpeedAPIScheduleNewResponseResultScheduleRegionAsiaSouth1          SpeedAPIScheduleNewResponseResultScheduleRegion = "asia-south1"
	SpeedAPIScheduleNewResponseResultScheduleRegionAsiaSoutheast1      SpeedAPIScheduleNewResponseResultScheduleRegion = "asia-southeast1"
	SpeedAPIScheduleNewResponseResultScheduleRegionAustraliaSoutheast1 SpeedAPIScheduleNewResponseResultScheduleRegion = "australia-southeast1"
	SpeedAPIScheduleNewResponseResultScheduleRegionEuropeNorth1        SpeedAPIScheduleNewResponseResultScheduleRegion = "europe-north1"
	SpeedAPIScheduleNewResponseResultScheduleRegionEuropeSouthwest1    SpeedAPIScheduleNewResponseResultScheduleRegion = "europe-southwest1"
	SpeedAPIScheduleNewResponseResultScheduleRegionEuropeWest1         SpeedAPIScheduleNewResponseResultScheduleRegion = "europe-west1"
	SpeedAPIScheduleNewResponseResultScheduleRegionEuropeWest2         SpeedAPIScheduleNewResponseResultScheduleRegion = "europe-west2"
	SpeedAPIScheduleNewResponseResultScheduleRegionEuropeWest3         SpeedAPIScheduleNewResponseResultScheduleRegion = "europe-west3"
	SpeedAPIScheduleNewResponseResultScheduleRegionEuropeWest4         SpeedAPIScheduleNewResponseResultScheduleRegion = "europe-west4"
	SpeedAPIScheduleNewResponseResultScheduleRegionEuropeWest8         SpeedAPIScheduleNewResponseResultScheduleRegion = "europe-west8"
	SpeedAPIScheduleNewResponseResultScheduleRegionEuropeWest9         SpeedAPIScheduleNewResponseResultScheduleRegion = "europe-west9"
	SpeedAPIScheduleNewResponseResultScheduleRegionMeWest1             SpeedAPIScheduleNewResponseResultScheduleRegion = "me-west1"
	SpeedAPIScheduleNewResponseResultScheduleRegionSouthamericaEast1   SpeedAPIScheduleNewResponseResultScheduleRegion = "southamerica-east1"
	SpeedAPIScheduleNewResponseResultScheduleRegionUsCentral1          SpeedAPIScheduleNewResponseResultScheduleRegion = "us-central1"
	SpeedAPIScheduleNewResponseResultScheduleRegionUsEast1             SpeedAPIScheduleNewResponseResultScheduleRegion = "us-east1"
	SpeedAPIScheduleNewResponseResultScheduleRegionUsEast4             SpeedAPIScheduleNewResponseResultScheduleRegion = "us-east4"
	SpeedAPIScheduleNewResponseResultScheduleRegionUsSouth1            SpeedAPIScheduleNewResponseResultScheduleRegion = "us-south1"
	SpeedAPIScheduleNewResponseResultScheduleRegionUsWest1             SpeedAPIScheduleNewResponseResultScheduleRegion = "us-west1"
)

type SpeedAPIScheduleNewResponseResultTest struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedAPIScheduleNewResponseResultTestDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedAPIScheduleNewResponseResultTestMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedAPIScheduleNewResponseResultTestRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedAPIScheduleNewResponseResultTestScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                                    `json:"url"`
	JSON speedAPIScheduleNewResponseResultTestJSON `json:"-"`
}

// speedAPIScheduleNewResponseResultTestJSON contains the JSON metadata for the
// struct [SpeedAPIScheduleNewResponseResultTest]
type speedAPIScheduleNewResponseResultTestJSON struct {
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

func (r *SpeedAPIScheduleNewResponseResultTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedAPIScheduleNewResponseResultTestDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPIScheduleNewResponseResultTestDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedAPIScheduleNewResponseResultTestDesktopReportError      `json:"error"`
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
	State SpeedAPIScheduleNewResponseResultTestDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                                `json:"tti"`
	JSON speedAPIScheduleNewResponseResultTestDesktopReportJSON `json:"-"`
}

// speedAPIScheduleNewResponseResultTestDesktopReportJSON contains the JSON
// metadata for the struct [SpeedAPIScheduleNewResponseResultTestDesktopReport]
type speedAPIScheduleNewResponseResultTestDesktopReportJSON struct {
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

func (r *SpeedAPIScheduleNewResponseResultTestDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPIScheduleNewResponseResultTestDesktopReportDeviceType string

const (
	SpeedAPIScheduleNewResponseResultTestDesktopReportDeviceTypeDesktop SpeedAPIScheduleNewResponseResultTestDesktopReportDeviceType = "DESKTOP"
	SpeedAPIScheduleNewResponseResultTestDesktopReportDeviceTypeMobile  SpeedAPIScheduleNewResponseResultTestDesktopReportDeviceType = "MOBILE"
)

type SpeedAPIScheduleNewResponseResultTestDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                      `json:"finalDisplayedUrl"`
	JSON              speedAPIScheduleNewResponseResultTestDesktopReportErrorJSON `json:"-"`
}

// speedAPIScheduleNewResponseResultTestDesktopReportErrorJSON contains the JSON
// metadata for the struct
// [SpeedAPIScheduleNewResponseResultTestDesktopReportError]
type speedAPIScheduleNewResponseResultTestDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseResultTestDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode string

const (
	SpeedAPIScheduleNewResponseResultTestDesktopReportErrorCodeNotReachable      SpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedAPIScheduleNewResponseResultTestDesktopReportErrorCodeDNSFailure        SpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode = "DNS_FAILURE"
	SpeedAPIScheduleNewResponseResultTestDesktopReportErrorCodeNotHTML           SpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode = "NOT_HTML"
	SpeedAPIScheduleNewResponseResultTestDesktopReportErrorCodeLighthouseTimeout SpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPIScheduleNewResponseResultTestDesktopReportErrorCodeUnknown           SpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPIScheduleNewResponseResultTestDesktopReportState string

const (
	SpeedAPIScheduleNewResponseResultTestDesktopReportStateRunning  SpeedAPIScheduleNewResponseResultTestDesktopReportState = "RUNNING"
	SpeedAPIScheduleNewResponseResultTestDesktopReportStateComplete SpeedAPIScheduleNewResponseResultTestDesktopReportState = "COMPLETE"
	SpeedAPIScheduleNewResponseResultTestDesktopReportStateFailed   SpeedAPIScheduleNewResponseResultTestDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedAPIScheduleNewResponseResultTestMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPIScheduleNewResponseResultTestMobileReportDeviceType `json:"deviceType"`
	Error      SpeedAPIScheduleNewResponseResultTestMobileReportError      `json:"error"`
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
	State SpeedAPIScheduleNewResponseResultTestMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                               `json:"tti"`
	JSON speedAPIScheduleNewResponseResultTestMobileReportJSON `json:"-"`
}

// speedAPIScheduleNewResponseResultTestMobileReportJSON contains the JSON metadata
// for the struct [SpeedAPIScheduleNewResponseResultTestMobileReport]
type speedAPIScheduleNewResponseResultTestMobileReportJSON struct {
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

func (r *SpeedAPIScheduleNewResponseResultTestMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPIScheduleNewResponseResultTestMobileReportDeviceType string

const (
	SpeedAPIScheduleNewResponseResultTestMobileReportDeviceTypeDesktop SpeedAPIScheduleNewResponseResultTestMobileReportDeviceType = "DESKTOP"
	SpeedAPIScheduleNewResponseResultTestMobileReportDeviceTypeMobile  SpeedAPIScheduleNewResponseResultTestMobileReportDeviceType = "MOBILE"
)

type SpeedAPIScheduleNewResponseResultTestMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPIScheduleNewResponseResultTestMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                     `json:"finalDisplayedUrl"`
	JSON              speedAPIScheduleNewResponseResultTestMobileReportErrorJSON `json:"-"`
}

// speedAPIScheduleNewResponseResultTestMobileReportErrorJSON contains the JSON
// metadata for the struct [SpeedAPIScheduleNewResponseResultTestMobileReportError]
type speedAPIScheduleNewResponseResultTestMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseResultTestMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPIScheduleNewResponseResultTestMobileReportErrorCode string

const (
	SpeedAPIScheduleNewResponseResultTestMobileReportErrorCodeNotReachable      SpeedAPIScheduleNewResponseResultTestMobileReportErrorCode = "NOT_REACHABLE"
	SpeedAPIScheduleNewResponseResultTestMobileReportErrorCodeDNSFailure        SpeedAPIScheduleNewResponseResultTestMobileReportErrorCode = "DNS_FAILURE"
	SpeedAPIScheduleNewResponseResultTestMobileReportErrorCodeNotHTML           SpeedAPIScheduleNewResponseResultTestMobileReportErrorCode = "NOT_HTML"
	SpeedAPIScheduleNewResponseResultTestMobileReportErrorCodeLighthouseTimeout SpeedAPIScheduleNewResponseResultTestMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPIScheduleNewResponseResultTestMobileReportErrorCodeUnknown           SpeedAPIScheduleNewResponseResultTestMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPIScheduleNewResponseResultTestMobileReportState string

const (
	SpeedAPIScheduleNewResponseResultTestMobileReportStateRunning  SpeedAPIScheduleNewResponseResultTestMobileReportState = "RUNNING"
	SpeedAPIScheduleNewResponseResultTestMobileReportStateComplete SpeedAPIScheduleNewResponseResultTestMobileReportState = "COMPLETE"
	SpeedAPIScheduleNewResponseResultTestMobileReportStateFailed   SpeedAPIScheduleNewResponseResultTestMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedAPIScheduleNewResponseResultTestRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPIScheduleNewResponseResultTestRegionValue `json:"value"`
	JSON  speedAPIScheduleNewResponseResultTestRegionJSON  `json:"-"`
}

// speedAPIScheduleNewResponseResultTestRegionJSON contains the JSON metadata for
// the struct [SpeedAPIScheduleNewResponseResultTestRegion]
type speedAPIScheduleNewResponseResultTestRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseResultTestRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPIScheduleNewResponseResultTestRegionValue string

const (
	SpeedAPIScheduleNewResponseResultTestRegionValueAsiaEast1           SpeedAPIScheduleNewResponseResultTestRegionValue = "asia-east1"
	SpeedAPIScheduleNewResponseResultTestRegionValueAsiaNortheast1      SpeedAPIScheduleNewResponseResultTestRegionValue = "asia-northeast1"
	SpeedAPIScheduleNewResponseResultTestRegionValueAsiaNortheast2      SpeedAPIScheduleNewResponseResultTestRegionValue = "asia-northeast2"
	SpeedAPIScheduleNewResponseResultTestRegionValueAsiaSouth1          SpeedAPIScheduleNewResponseResultTestRegionValue = "asia-south1"
	SpeedAPIScheduleNewResponseResultTestRegionValueAsiaSoutheast1      SpeedAPIScheduleNewResponseResultTestRegionValue = "asia-southeast1"
	SpeedAPIScheduleNewResponseResultTestRegionValueAustraliaSoutheast1 SpeedAPIScheduleNewResponseResultTestRegionValue = "australia-southeast1"
	SpeedAPIScheduleNewResponseResultTestRegionValueEuropeNorth1        SpeedAPIScheduleNewResponseResultTestRegionValue = "europe-north1"
	SpeedAPIScheduleNewResponseResultTestRegionValueEuropeSouthwest1    SpeedAPIScheduleNewResponseResultTestRegionValue = "europe-southwest1"
	SpeedAPIScheduleNewResponseResultTestRegionValueEuropeWest1         SpeedAPIScheduleNewResponseResultTestRegionValue = "europe-west1"
	SpeedAPIScheduleNewResponseResultTestRegionValueEuropeWest2         SpeedAPIScheduleNewResponseResultTestRegionValue = "europe-west2"
	SpeedAPIScheduleNewResponseResultTestRegionValueEuropeWest3         SpeedAPIScheduleNewResponseResultTestRegionValue = "europe-west3"
	SpeedAPIScheduleNewResponseResultTestRegionValueEuropeWest4         SpeedAPIScheduleNewResponseResultTestRegionValue = "europe-west4"
	SpeedAPIScheduleNewResponseResultTestRegionValueEuropeWest8         SpeedAPIScheduleNewResponseResultTestRegionValue = "europe-west8"
	SpeedAPIScheduleNewResponseResultTestRegionValueEuropeWest9         SpeedAPIScheduleNewResponseResultTestRegionValue = "europe-west9"
	SpeedAPIScheduleNewResponseResultTestRegionValueMeWest1             SpeedAPIScheduleNewResponseResultTestRegionValue = "me-west1"
	SpeedAPIScheduleNewResponseResultTestRegionValueSouthamericaEast1   SpeedAPIScheduleNewResponseResultTestRegionValue = "southamerica-east1"
	SpeedAPIScheduleNewResponseResultTestRegionValueUsCentral1          SpeedAPIScheduleNewResponseResultTestRegionValue = "us-central1"
	SpeedAPIScheduleNewResponseResultTestRegionValueUsEast1             SpeedAPIScheduleNewResponseResultTestRegionValue = "us-east1"
	SpeedAPIScheduleNewResponseResultTestRegionValueUsEast4             SpeedAPIScheduleNewResponseResultTestRegionValue = "us-east4"
	SpeedAPIScheduleNewResponseResultTestRegionValueUsSouth1            SpeedAPIScheduleNewResponseResultTestRegionValue = "us-south1"
	SpeedAPIScheduleNewResponseResultTestRegionValueUsWest1             SpeedAPIScheduleNewResponseResultTestRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedAPIScheduleNewResponseResultTestScheduleFrequency string

const (
	SpeedAPIScheduleNewResponseResultTestScheduleFrequencyDaily  SpeedAPIScheduleNewResponseResultTestScheduleFrequency = "DAILY"
	SpeedAPIScheduleNewResponseResultTestScheduleFrequencyWeekly SpeedAPIScheduleNewResponseResultTestScheduleFrequency = "WEEKLY"
)

type SpeedAPIScheduleNewParams struct {
	// A test region.
	Region param.Field[SpeedAPIScheduleNewParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedAPIScheduleNewParams]'s query parameters as
// `url.Values`.
func (r SpeedAPIScheduleNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedAPIScheduleNewParamsRegion string

const (
	SpeedAPIScheduleNewParamsRegionAsiaEast1           SpeedAPIScheduleNewParamsRegion = "asia-east1"
	SpeedAPIScheduleNewParamsRegionAsiaNortheast1      SpeedAPIScheduleNewParamsRegion = "asia-northeast1"
	SpeedAPIScheduleNewParamsRegionAsiaNortheast2      SpeedAPIScheduleNewParamsRegion = "asia-northeast2"
	SpeedAPIScheduleNewParamsRegionAsiaSouth1          SpeedAPIScheduleNewParamsRegion = "asia-south1"
	SpeedAPIScheduleNewParamsRegionAsiaSoutheast1      SpeedAPIScheduleNewParamsRegion = "asia-southeast1"
	SpeedAPIScheduleNewParamsRegionAustraliaSoutheast1 SpeedAPIScheduleNewParamsRegion = "australia-southeast1"
	SpeedAPIScheduleNewParamsRegionEuropeNorth1        SpeedAPIScheduleNewParamsRegion = "europe-north1"
	SpeedAPIScheduleNewParamsRegionEuropeSouthwest1    SpeedAPIScheduleNewParamsRegion = "europe-southwest1"
	SpeedAPIScheduleNewParamsRegionEuropeWest1         SpeedAPIScheduleNewParamsRegion = "europe-west1"
	SpeedAPIScheduleNewParamsRegionEuropeWest2         SpeedAPIScheduleNewParamsRegion = "europe-west2"
	SpeedAPIScheduleNewParamsRegionEuropeWest3         SpeedAPIScheduleNewParamsRegion = "europe-west3"
	SpeedAPIScheduleNewParamsRegionEuropeWest4         SpeedAPIScheduleNewParamsRegion = "europe-west4"
	SpeedAPIScheduleNewParamsRegionEuropeWest8         SpeedAPIScheduleNewParamsRegion = "europe-west8"
	SpeedAPIScheduleNewParamsRegionEuropeWest9         SpeedAPIScheduleNewParamsRegion = "europe-west9"
	SpeedAPIScheduleNewParamsRegionMeWest1             SpeedAPIScheduleNewParamsRegion = "me-west1"
	SpeedAPIScheduleNewParamsRegionSouthamericaEast1   SpeedAPIScheduleNewParamsRegion = "southamerica-east1"
	SpeedAPIScheduleNewParamsRegionUsCentral1          SpeedAPIScheduleNewParamsRegion = "us-central1"
	SpeedAPIScheduleNewParamsRegionUsEast1             SpeedAPIScheduleNewParamsRegion = "us-east1"
	SpeedAPIScheduleNewParamsRegionUsEast4             SpeedAPIScheduleNewParamsRegion = "us-east4"
	SpeedAPIScheduleNewParamsRegionUsSouth1            SpeedAPIScheduleNewParamsRegion = "us-south1"
	SpeedAPIScheduleNewParamsRegionUsWest1             SpeedAPIScheduleNewParamsRegion = "us-west1"
)
