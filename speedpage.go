// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SpeedPageService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSpeedPageService] method instead.
type SpeedPageService struct {
	Options []option.RequestOption
}

// NewSpeedPageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpeedPageService(opts ...option.RequestOption) (r *SpeedPageService) {
	r = &SpeedPageService{}
	r.Options = opts
	return
}

// Lists all webpages which have been tested.
func (r *SpeedPageService) List(ctx context.Context, query SpeedPageListParams, opts ...option.RequestOption) (res *[]SpeedPageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedPageListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SpeedPageListResponse struct {
	// A test region with a label.
	Region SpeedPageListResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedPageListResponseScheduleFrequency `json:"scheduleFrequency"`
	Tests             []SpeedPageListResponseTest            `json:"tests"`
	// A URL.
	URL  string                    `json:"url"`
	JSON speedPageListResponseJSON `json:"-"`
}

// speedPageListResponseJSON contains the JSON metadata for the struct
// [SpeedPageListResponse]
type speedPageListResponseJSON struct {
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	Tests             apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedPageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedPageListResponseJSON) RawJSON() string {
	return r.raw
}

// A test region with a label.
type SpeedPageListResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedPageListResponseRegionValue `json:"value"`
	JSON  speedPageListResponseRegionJSON  `json:"-"`
}

// speedPageListResponseRegionJSON contains the JSON metadata for the struct
// [SpeedPageListResponseRegion]
type speedPageListResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedPageListResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedPageListResponseRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type SpeedPageListResponseRegionValue string

const (
	SpeedPageListResponseRegionValueAsiaEast1           SpeedPageListResponseRegionValue = "asia-east1"
	SpeedPageListResponseRegionValueAsiaNortheast1      SpeedPageListResponseRegionValue = "asia-northeast1"
	SpeedPageListResponseRegionValueAsiaNortheast2      SpeedPageListResponseRegionValue = "asia-northeast2"
	SpeedPageListResponseRegionValueAsiaSouth1          SpeedPageListResponseRegionValue = "asia-south1"
	SpeedPageListResponseRegionValueAsiaSoutheast1      SpeedPageListResponseRegionValue = "asia-southeast1"
	SpeedPageListResponseRegionValueAustraliaSoutheast1 SpeedPageListResponseRegionValue = "australia-southeast1"
	SpeedPageListResponseRegionValueEuropeNorth1        SpeedPageListResponseRegionValue = "europe-north1"
	SpeedPageListResponseRegionValueEuropeSouthwest1    SpeedPageListResponseRegionValue = "europe-southwest1"
	SpeedPageListResponseRegionValueEuropeWest1         SpeedPageListResponseRegionValue = "europe-west1"
	SpeedPageListResponseRegionValueEuropeWest2         SpeedPageListResponseRegionValue = "europe-west2"
	SpeedPageListResponseRegionValueEuropeWest3         SpeedPageListResponseRegionValue = "europe-west3"
	SpeedPageListResponseRegionValueEuropeWest4         SpeedPageListResponseRegionValue = "europe-west4"
	SpeedPageListResponseRegionValueEuropeWest8         SpeedPageListResponseRegionValue = "europe-west8"
	SpeedPageListResponseRegionValueEuropeWest9         SpeedPageListResponseRegionValue = "europe-west9"
	SpeedPageListResponseRegionValueMeWest1             SpeedPageListResponseRegionValue = "me-west1"
	SpeedPageListResponseRegionValueSouthamericaEast1   SpeedPageListResponseRegionValue = "southamerica-east1"
	SpeedPageListResponseRegionValueUsCentral1          SpeedPageListResponseRegionValue = "us-central1"
	SpeedPageListResponseRegionValueUsEast1             SpeedPageListResponseRegionValue = "us-east1"
	SpeedPageListResponseRegionValueUsEast4             SpeedPageListResponseRegionValue = "us-east4"
	SpeedPageListResponseRegionValueUsSouth1            SpeedPageListResponseRegionValue = "us-south1"
	SpeedPageListResponseRegionValueUsWest1             SpeedPageListResponseRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedPageListResponseScheduleFrequency string

const (
	SpeedPageListResponseScheduleFrequencyDaily  SpeedPageListResponseScheduleFrequency = "DAILY"
	SpeedPageListResponseScheduleFrequencyWeekly SpeedPageListResponseScheduleFrequency = "WEEKLY"
)

type SpeedPageListResponseTest struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedPageListResponseTestsDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedPageListResponseTestsMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedPageListResponseTestsRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedPageListResponseTestsScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                        `json:"url"`
	JSON speedPageListResponseTestJSON `json:"-"`
}

// speedPageListResponseTestJSON contains the JSON metadata for the struct
// [SpeedPageListResponseTest]
type speedPageListResponseTestJSON struct {
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

func (r *SpeedPageListResponseTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedPageListResponseTestJSON) RawJSON() string {
	return r.raw
}

// The Lighthouse report.
type SpeedPageListResponseTestsDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedPageListResponseTestsDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedPageListResponseTestsDesktopReportError      `json:"error"`
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
	State SpeedPageListResponseTestsDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                     `json:"tti"`
	JSON speedPageListResponseTestsDesktopReportJSON `json:"-"`
}

// speedPageListResponseTestsDesktopReportJSON contains the JSON metadata for the
// struct [SpeedPageListResponseTestsDesktopReport]
type speedPageListResponseTestsDesktopReportJSON struct {
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

func (r *SpeedPageListResponseTestsDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedPageListResponseTestsDesktopReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type SpeedPageListResponseTestsDesktopReportDeviceType string

const (
	SpeedPageListResponseTestsDesktopReportDeviceTypeDesktop SpeedPageListResponseTestsDesktopReportDeviceType = "DESKTOP"
	SpeedPageListResponseTestsDesktopReportDeviceTypeMobile  SpeedPageListResponseTestsDesktopReportDeviceType = "MOBILE"
)

type SpeedPageListResponseTestsDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedPageListResponseTestsDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                           `json:"finalDisplayedUrl"`
	JSON              speedPageListResponseTestsDesktopReportErrorJSON `json:"-"`
}

// speedPageListResponseTestsDesktopReportErrorJSON contains the JSON metadata for
// the struct [SpeedPageListResponseTestsDesktopReportError]
type speedPageListResponseTestsDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedPageListResponseTestsDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedPageListResponseTestsDesktopReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type SpeedPageListResponseTestsDesktopReportErrorCode string

const (
	SpeedPageListResponseTestsDesktopReportErrorCodeNotReachable      SpeedPageListResponseTestsDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedPageListResponseTestsDesktopReportErrorCodeDNSFailure        SpeedPageListResponseTestsDesktopReportErrorCode = "DNS_FAILURE"
	SpeedPageListResponseTestsDesktopReportErrorCodeNotHTML           SpeedPageListResponseTestsDesktopReportErrorCode = "NOT_HTML"
	SpeedPageListResponseTestsDesktopReportErrorCodeLighthouseTimeout SpeedPageListResponseTestsDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedPageListResponseTestsDesktopReportErrorCodeUnknown           SpeedPageListResponseTestsDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedPageListResponseTestsDesktopReportState string

const (
	SpeedPageListResponseTestsDesktopReportStateRunning  SpeedPageListResponseTestsDesktopReportState = "RUNNING"
	SpeedPageListResponseTestsDesktopReportStateComplete SpeedPageListResponseTestsDesktopReportState = "COMPLETE"
	SpeedPageListResponseTestsDesktopReportStateFailed   SpeedPageListResponseTestsDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedPageListResponseTestsMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedPageListResponseTestsMobileReportDeviceType `json:"deviceType"`
	Error      SpeedPageListResponseTestsMobileReportError      `json:"error"`
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
	State SpeedPageListResponseTestsMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                    `json:"tti"`
	JSON speedPageListResponseTestsMobileReportJSON `json:"-"`
}

// speedPageListResponseTestsMobileReportJSON contains the JSON metadata for the
// struct [SpeedPageListResponseTestsMobileReport]
type speedPageListResponseTestsMobileReportJSON struct {
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

func (r *SpeedPageListResponseTestsMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedPageListResponseTestsMobileReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type SpeedPageListResponseTestsMobileReportDeviceType string

const (
	SpeedPageListResponseTestsMobileReportDeviceTypeDesktop SpeedPageListResponseTestsMobileReportDeviceType = "DESKTOP"
	SpeedPageListResponseTestsMobileReportDeviceTypeMobile  SpeedPageListResponseTestsMobileReportDeviceType = "MOBILE"
)

type SpeedPageListResponseTestsMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedPageListResponseTestsMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                          `json:"finalDisplayedUrl"`
	JSON              speedPageListResponseTestsMobileReportErrorJSON `json:"-"`
}

// speedPageListResponseTestsMobileReportErrorJSON contains the JSON metadata for
// the struct [SpeedPageListResponseTestsMobileReportError]
type speedPageListResponseTestsMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedPageListResponseTestsMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedPageListResponseTestsMobileReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type SpeedPageListResponseTestsMobileReportErrorCode string

const (
	SpeedPageListResponseTestsMobileReportErrorCodeNotReachable      SpeedPageListResponseTestsMobileReportErrorCode = "NOT_REACHABLE"
	SpeedPageListResponseTestsMobileReportErrorCodeDNSFailure        SpeedPageListResponseTestsMobileReportErrorCode = "DNS_FAILURE"
	SpeedPageListResponseTestsMobileReportErrorCodeNotHTML           SpeedPageListResponseTestsMobileReportErrorCode = "NOT_HTML"
	SpeedPageListResponseTestsMobileReportErrorCodeLighthouseTimeout SpeedPageListResponseTestsMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedPageListResponseTestsMobileReportErrorCodeUnknown           SpeedPageListResponseTestsMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedPageListResponseTestsMobileReportState string

const (
	SpeedPageListResponseTestsMobileReportStateRunning  SpeedPageListResponseTestsMobileReportState = "RUNNING"
	SpeedPageListResponseTestsMobileReportStateComplete SpeedPageListResponseTestsMobileReportState = "COMPLETE"
	SpeedPageListResponseTestsMobileReportStateFailed   SpeedPageListResponseTestsMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedPageListResponseTestsRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedPageListResponseTestsRegionValue `json:"value"`
	JSON  speedPageListResponseTestsRegionJSON  `json:"-"`
}

// speedPageListResponseTestsRegionJSON contains the JSON metadata for the struct
// [SpeedPageListResponseTestsRegion]
type speedPageListResponseTestsRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedPageListResponseTestsRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedPageListResponseTestsRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type SpeedPageListResponseTestsRegionValue string

const (
	SpeedPageListResponseTestsRegionValueAsiaEast1           SpeedPageListResponseTestsRegionValue = "asia-east1"
	SpeedPageListResponseTestsRegionValueAsiaNortheast1      SpeedPageListResponseTestsRegionValue = "asia-northeast1"
	SpeedPageListResponseTestsRegionValueAsiaNortheast2      SpeedPageListResponseTestsRegionValue = "asia-northeast2"
	SpeedPageListResponseTestsRegionValueAsiaSouth1          SpeedPageListResponseTestsRegionValue = "asia-south1"
	SpeedPageListResponseTestsRegionValueAsiaSoutheast1      SpeedPageListResponseTestsRegionValue = "asia-southeast1"
	SpeedPageListResponseTestsRegionValueAustraliaSoutheast1 SpeedPageListResponseTestsRegionValue = "australia-southeast1"
	SpeedPageListResponseTestsRegionValueEuropeNorth1        SpeedPageListResponseTestsRegionValue = "europe-north1"
	SpeedPageListResponseTestsRegionValueEuropeSouthwest1    SpeedPageListResponseTestsRegionValue = "europe-southwest1"
	SpeedPageListResponseTestsRegionValueEuropeWest1         SpeedPageListResponseTestsRegionValue = "europe-west1"
	SpeedPageListResponseTestsRegionValueEuropeWest2         SpeedPageListResponseTestsRegionValue = "europe-west2"
	SpeedPageListResponseTestsRegionValueEuropeWest3         SpeedPageListResponseTestsRegionValue = "europe-west3"
	SpeedPageListResponseTestsRegionValueEuropeWest4         SpeedPageListResponseTestsRegionValue = "europe-west4"
	SpeedPageListResponseTestsRegionValueEuropeWest8         SpeedPageListResponseTestsRegionValue = "europe-west8"
	SpeedPageListResponseTestsRegionValueEuropeWest9         SpeedPageListResponseTestsRegionValue = "europe-west9"
	SpeedPageListResponseTestsRegionValueMeWest1             SpeedPageListResponseTestsRegionValue = "me-west1"
	SpeedPageListResponseTestsRegionValueSouthamericaEast1   SpeedPageListResponseTestsRegionValue = "southamerica-east1"
	SpeedPageListResponseTestsRegionValueUsCentral1          SpeedPageListResponseTestsRegionValue = "us-central1"
	SpeedPageListResponseTestsRegionValueUsEast1             SpeedPageListResponseTestsRegionValue = "us-east1"
	SpeedPageListResponseTestsRegionValueUsEast4             SpeedPageListResponseTestsRegionValue = "us-east4"
	SpeedPageListResponseTestsRegionValueUsSouth1            SpeedPageListResponseTestsRegionValue = "us-south1"
	SpeedPageListResponseTestsRegionValueUsWest1             SpeedPageListResponseTestsRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedPageListResponseTestsScheduleFrequency string

const (
	SpeedPageListResponseTestsScheduleFrequencyDaily  SpeedPageListResponseTestsScheduleFrequency = "DAILY"
	SpeedPageListResponseTestsScheduleFrequencyWeekly SpeedPageListResponseTestsScheduleFrequency = "WEEKLY"
)

type SpeedPageListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SpeedPageListResponseEnvelope struct {
	Result []SpeedPageListResponse           `json:"result"`
	JSON   speedPageListResponseEnvelopeJSON `json:"-"`
}

// speedPageListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedPageListResponseEnvelope]
type speedPageListResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedPageListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedPageListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
