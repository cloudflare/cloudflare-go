// File generated from our OpenAPI spec by Stainless.

package cloudflare

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

// SpeedScheduleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSpeedScheduleService] method
// instead.
type SpeedScheduleService struct {
	Options []option.RequestOption
}

// NewSpeedScheduleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpeedScheduleService(opts ...option.RequestOption) (r *SpeedScheduleService) {
	r = &SpeedScheduleService{}
	r.Options = opts
	return
}

// Creates a scheduled test for a page.
func (r *SpeedScheduleService) New(ctx context.Context, url string, params SpeedScheduleNewParams, opts ...option.RequestOption) (res *SpeedScheduleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedScheduleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SpeedScheduleNewResponse struct {
	// The test schedule.
	Schedule SpeedScheduleNewResponseSchedule `json:"schedule"`
	Test     SpeedScheduleNewResponseTest     `json:"test"`
	JSON     speedScheduleNewResponseJSON     `json:"-"`
}

// speedScheduleNewResponseJSON contains the JSON metadata for the struct
// [SpeedScheduleNewResponse]
type speedScheduleNewResponseJSON struct {
	Schedule    apijson.Field
	Test        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedScheduleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedScheduleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The test schedule.
type SpeedScheduleNewResponseSchedule struct {
	// The frequency of the test.
	Frequency SpeedScheduleNewResponseScheduleFrequency `json:"frequency"`
	// A test region.
	Region SpeedScheduleNewResponseScheduleRegion `json:"region"`
	// A URL.
	URL  string                               `json:"url"`
	JSON speedScheduleNewResponseScheduleJSON `json:"-"`
}

// speedScheduleNewResponseScheduleJSON contains the JSON metadata for the struct
// [SpeedScheduleNewResponseSchedule]
type speedScheduleNewResponseScheduleJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedScheduleNewResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedScheduleNewResponseScheduleJSON) RawJSON() string {
	return r.raw
}

// The frequency of the test.
type SpeedScheduleNewResponseScheduleFrequency string

const (
	SpeedScheduleNewResponseScheduleFrequencyDaily  SpeedScheduleNewResponseScheduleFrequency = "DAILY"
	SpeedScheduleNewResponseScheduleFrequencyWeekly SpeedScheduleNewResponseScheduleFrequency = "WEEKLY"
)

// A test region.
type SpeedScheduleNewResponseScheduleRegion string

const (
	SpeedScheduleNewResponseScheduleRegionAsiaEast1           SpeedScheduleNewResponseScheduleRegion = "asia-east1"
	SpeedScheduleNewResponseScheduleRegionAsiaNortheast1      SpeedScheduleNewResponseScheduleRegion = "asia-northeast1"
	SpeedScheduleNewResponseScheduleRegionAsiaNortheast2      SpeedScheduleNewResponseScheduleRegion = "asia-northeast2"
	SpeedScheduleNewResponseScheduleRegionAsiaSouth1          SpeedScheduleNewResponseScheduleRegion = "asia-south1"
	SpeedScheduleNewResponseScheduleRegionAsiaSoutheast1      SpeedScheduleNewResponseScheduleRegion = "asia-southeast1"
	SpeedScheduleNewResponseScheduleRegionAustraliaSoutheast1 SpeedScheduleNewResponseScheduleRegion = "australia-southeast1"
	SpeedScheduleNewResponseScheduleRegionEuropeNorth1        SpeedScheduleNewResponseScheduleRegion = "europe-north1"
	SpeedScheduleNewResponseScheduleRegionEuropeSouthwest1    SpeedScheduleNewResponseScheduleRegion = "europe-southwest1"
	SpeedScheduleNewResponseScheduleRegionEuropeWest1         SpeedScheduleNewResponseScheduleRegion = "europe-west1"
	SpeedScheduleNewResponseScheduleRegionEuropeWest2         SpeedScheduleNewResponseScheduleRegion = "europe-west2"
	SpeedScheduleNewResponseScheduleRegionEuropeWest3         SpeedScheduleNewResponseScheduleRegion = "europe-west3"
	SpeedScheduleNewResponseScheduleRegionEuropeWest4         SpeedScheduleNewResponseScheduleRegion = "europe-west4"
	SpeedScheduleNewResponseScheduleRegionEuropeWest8         SpeedScheduleNewResponseScheduleRegion = "europe-west8"
	SpeedScheduleNewResponseScheduleRegionEuropeWest9         SpeedScheduleNewResponseScheduleRegion = "europe-west9"
	SpeedScheduleNewResponseScheduleRegionMeWest1             SpeedScheduleNewResponseScheduleRegion = "me-west1"
	SpeedScheduleNewResponseScheduleRegionSouthamericaEast1   SpeedScheduleNewResponseScheduleRegion = "southamerica-east1"
	SpeedScheduleNewResponseScheduleRegionUsCentral1          SpeedScheduleNewResponseScheduleRegion = "us-central1"
	SpeedScheduleNewResponseScheduleRegionUsEast1             SpeedScheduleNewResponseScheduleRegion = "us-east1"
	SpeedScheduleNewResponseScheduleRegionUsEast4             SpeedScheduleNewResponseScheduleRegion = "us-east4"
	SpeedScheduleNewResponseScheduleRegionUsSouth1            SpeedScheduleNewResponseScheduleRegion = "us-south1"
	SpeedScheduleNewResponseScheduleRegionUsWest1             SpeedScheduleNewResponseScheduleRegion = "us-west1"
)

type SpeedScheduleNewResponseTest struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedScheduleNewResponseTestDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedScheduleNewResponseTestMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedScheduleNewResponseTestRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedScheduleNewResponseTestScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                           `json:"url"`
	JSON speedScheduleNewResponseTestJSON `json:"-"`
}

// speedScheduleNewResponseTestJSON contains the JSON metadata for the struct
// [SpeedScheduleNewResponseTest]
type speedScheduleNewResponseTestJSON struct {
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

func (r *SpeedScheduleNewResponseTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedScheduleNewResponseTestJSON) RawJSON() string {
	return r.raw
}

// The Lighthouse report.
type SpeedScheduleNewResponseTestDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedScheduleNewResponseTestDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedScheduleNewResponseTestDesktopReportError      `json:"error"`
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
	State SpeedScheduleNewResponseTestDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                       `json:"tti"`
	JSON speedScheduleNewResponseTestDesktopReportJSON `json:"-"`
}

// speedScheduleNewResponseTestDesktopReportJSON contains the JSON metadata for the
// struct [SpeedScheduleNewResponseTestDesktopReport]
type speedScheduleNewResponseTestDesktopReportJSON struct {
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

func (r *SpeedScheduleNewResponseTestDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedScheduleNewResponseTestDesktopReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type SpeedScheduleNewResponseTestDesktopReportDeviceType string

const (
	SpeedScheduleNewResponseTestDesktopReportDeviceTypeDesktop SpeedScheduleNewResponseTestDesktopReportDeviceType = "DESKTOP"
	SpeedScheduleNewResponseTestDesktopReportDeviceTypeMobile  SpeedScheduleNewResponseTestDesktopReportDeviceType = "MOBILE"
)

type SpeedScheduleNewResponseTestDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedScheduleNewResponseTestDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                             `json:"finalDisplayedUrl"`
	JSON              speedScheduleNewResponseTestDesktopReportErrorJSON `json:"-"`
}

// speedScheduleNewResponseTestDesktopReportErrorJSON contains the JSON metadata
// for the struct [SpeedScheduleNewResponseTestDesktopReportError]
type speedScheduleNewResponseTestDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedScheduleNewResponseTestDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedScheduleNewResponseTestDesktopReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type SpeedScheduleNewResponseTestDesktopReportErrorCode string

const (
	SpeedScheduleNewResponseTestDesktopReportErrorCodeNotReachable      SpeedScheduleNewResponseTestDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedScheduleNewResponseTestDesktopReportErrorCodeDNSFailure        SpeedScheduleNewResponseTestDesktopReportErrorCode = "DNS_FAILURE"
	SpeedScheduleNewResponseTestDesktopReportErrorCodeNotHTML           SpeedScheduleNewResponseTestDesktopReportErrorCode = "NOT_HTML"
	SpeedScheduleNewResponseTestDesktopReportErrorCodeLighthouseTimeout SpeedScheduleNewResponseTestDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedScheduleNewResponseTestDesktopReportErrorCodeUnknown           SpeedScheduleNewResponseTestDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedScheduleNewResponseTestDesktopReportState string

const (
	SpeedScheduleNewResponseTestDesktopReportStateRunning  SpeedScheduleNewResponseTestDesktopReportState = "RUNNING"
	SpeedScheduleNewResponseTestDesktopReportStateComplete SpeedScheduleNewResponseTestDesktopReportState = "COMPLETE"
	SpeedScheduleNewResponseTestDesktopReportStateFailed   SpeedScheduleNewResponseTestDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedScheduleNewResponseTestMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedScheduleNewResponseTestMobileReportDeviceType `json:"deviceType"`
	Error      SpeedScheduleNewResponseTestMobileReportError      `json:"error"`
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
	State SpeedScheduleNewResponseTestMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                      `json:"tti"`
	JSON speedScheduleNewResponseTestMobileReportJSON `json:"-"`
}

// speedScheduleNewResponseTestMobileReportJSON contains the JSON metadata for the
// struct [SpeedScheduleNewResponseTestMobileReport]
type speedScheduleNewResponseTestMobileReportJSON struct {
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

func (r *SpeedScheduleNewResponseTestMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedScheduleNewResponseTestMobileReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type SpeedScheduleNewResponseTestMobileReportDeviceType string

const (
	SpeedScheduleNewResponseTestMobileReportDeviceTypeDesktop SpeedScheduleNewResponseTestMobileReportDeviceType = "DESKTOP"
	SpeedScheduleNewResponseTestMobileReportDeviceTypeMobile  SpeedScheduleNewResponseTestMobileReportDeviceType = "MOBILE"
)

type SpeedScheduleNewResponseTestMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedScheduleNewResponseTestMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                            `json:"finalDisplayedUrl"`
	JSON              speedScheduleNewResponseTestMobileReportErrorJSON `json:"-"`
}

// speedScheduleNewResponseTestMobileReportErrorJSON contains the JSON metadata for
// the struct [SpeedScheduleNewResponseTestMobileReportError]
type speedScheduleNewResponseTestMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedScheduleNewResponseTestMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedScheduleNewResponseTestMobileReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type SpeedScheduleNewResponseTestMobileReportErrorCode string

const (
	SpeedScheduleNewResponseTestMobileReportErrorCodeNotReachable      SpeedScheduleNewResponseTestMobileReportErrorCode = "NOT_REACHABLE"
	SpeedScheduleNewResponseTestMobileReportErrorCodeDNSFailure        SpeedScheduleNewResponseTestMobileReportErrorCode = "DNS_FAILURE"
	SpeedScheduleNewResponseTestMobileReportErrorCodeNotHTML           SpeedScheduleNewResponseTestMobileReportErrorCode = "NOT_HTML"
	SpeedScheduleNewResponseTestMobileReportErrorCodeLighthouseTimeout SpeedScheduleNewResponseTestMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedScheduleNewResponseTestMobileReportErrorCodeUnknown           SpeedScheduleNewResponseTestMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedScheduleNewResponseTestMobileReportState string

const (
	SpeedScheduleNewResponseTestMobileReportStateRunning  SpeedScheduleNewResponseTestMobileReportState = "RUNNING"
	SpeedScheduleNewResponseTestMobileReportStateComplete SpeedScheduleNewResponseTestMobileReportState = "COMPLETE"
	SpeedScheduleNewResponseTestMobileReportStateFailed   SpeedScheduleNewResponseTestMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedScheduleNewResponseTestRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedScheduleNewResponseTestRegionValue `json:"value"`
	JSON  speedScheduleNewResponseTestRegionJSON  `json:"-"`
}

// speedScheduleNewResponseTestRegionJSON contains the JSON metadata for the struct
// [SpeedScheduleNewResponseTestRegion]
type speedScheduleNewResponseTestRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedScheduleNewResponseTestRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedScheduleNewResponseTestRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type SpeedScheduleNewResponseTestRegionValue string

const (
	SpeedScheduleNewResponseTestRegionValueAsiaEast1           SpeedScheduleNewResponseTestRegionValue = "asia-east1"
	SpeedScheduleNewResponseTestRegionValueAsiaNortheast1      SpeedScheduleNewResponseTestRegionValue = "asia-northeast1"
	SpeedScheduleNewResponseTestRegionValueAsiaNortheast2      SpeedScheduleNewResponseTestRegionValue = "asia-northeast2"
	SpeedScheduleNewResponseTestRegionValueAsiaSouth1          SpeedScheduleNewResponseTestRegionValue = "asia-south1"
	SpeedScheduleNewResponseTestRegionValueAsiaSoutheast1      SpeedScheduleNewResponseTestRegionValue = "asia-southeast1"
	SpeedScheduleNewResponseTestRegionValueAustraliaSoutheast1 SpeedScheduleNewResponseTestRegionValue = "australia-southeast1"
	SpeedScheduleNewResponseTestRegionValueEuropeNorth1        SpeedScheduleNewResponseTestRegionValue = "europe-north1"
	SpeedScheduleNewResponseTestRegionValueEuropeSouthwest1    SpeedScheduleNewResponseTestRegionValue = "europe-southwest1"
	SpeedScheduleNewResponseTestRegionValueEuropeWest1         SpeedScheduleNewResponseTestRegionValue = "europe-west1"
	SpeedScheduleNewResponseTestRegionValueEuropeWest2         SpeedScheduleNewResponseTestRegionValue = "europe-west2"
	SpeedScheduleNewResponseTestRegionValueEuropeWest3         SpeedScheduleNewResponseTestRegionValue = "europe-west3"
	SpeedScheduleNewResponseTestRegionValueEuropeWest4         SpeedScheduleNewResponseTestRegionValue = "europe-west4"
	SpeedScheduleNewResponseTestRegionValueEuropeWest8         SpeedScheduleNewResponseTestRegionValue = "europe-west8"
	SpeedScheduleNewResponseTestRegionValueEuropeWest9         SpeedScheduleNewResponseTestRegionValue = "europe-west9"
	SpeedScheduleNewResponseTestRegionValueMeWest1             SpeedScheduleNewResponseTestRegionValue = "me-west1"
	SpeedScheduleNewResponseTestRegionValueSouthamericaEast1   SpeedScheduleNewResponseTestRegionValue = "southamerica-east1"
	SpeedScheduleNewResponseTestRegionValueUsCentral1          SpeedScheduleNewResponseTestRegionValue = "us-central1"
	SpeedScheduleNewResponseTestRegionValueUsEast1             SpeedScheduleNewResponseTestRegionValue = "us-east1"
	SpeedScheduleNewResponseTestRegionValueUsEast4             SpeedScheduleNewResponseTestRegionValue = "us-east4"
	SpeedScheduleNewResponseTestRegionValueUsSouth1            SpeedScheduleNewResponseTestRegionValue = "us-south1"
	SpeedScheduleNewResponseTestRegionValueUsWest1             SpeedScheduleNewResponseTestRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedScheduleNewResponseTestScheduleFrequency string

const (
	SpeedScheduleNewResponseTestScheduleFrequencyDaily  SpeedScheduleNewResponseTestScheduleFrequency = "DAILY"
	SpeedScheduleNewResponseTestScheduleFrequencyWeekly SpeedScheduleNewResponseTestScheduleFrequency = "WEEKLY"
)

type SpeedScheduleNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[SpeedScheduleNewParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedScheduleNewParams]'s query parameters as `url.Values`.
func (r SpeedScheduleNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedScheduleNewParamsRegion string

const (
	SpeedScheduleNewParamsRegionAsiaEast1           SpeedScheduleNewParamsRegion = "asia-east1"
	SpeedScheduleNewParamsRegionAsiaNortheast1      SpeedScheduleNewParamsRegion = "asia-northeast1"
	SpeedScheduleNewParamsRegionAsiaNortheast2      SpeedScheduleNewParamsRegion = "asia-northeast2"
	SpeedScheduleNewParamsRegionAsiaSouth1          SpeedScheduleNewParamsRegion = "asia-south1"
	SpeedScheduleNewParamsRegionAsiaSoutheast1      SpeedScheduleNewParamsRegion = "asia-southeast1"
	SpeedScheduleNewParamsRegionAustraliaSoutheast1 SpeedScheduleNewParamsRegion = "australia-southeast1"
	SpeedScheduleNewParamsRegionEuropeNorth1        SpeedScheduleNewParamsRegion = "europe-north1"
	SpeedScheduleNewParamsRegionEuropeSouthwest1    SpeedScheduleNewParamsRegion = "europe-southwest1"
	SpeedScheduleNewParamsRegionEuropeWest1         SpeedScheduleNewParamsRegion = "europe-west1"
	SpeedScheduleNewParamsRegionEuropeWest2         SpeedScheduleNewParamsRegion = "europe-west2"
	SpeedScheduleNewParamsRegionEuropeWest3         SpeedScheduleNewParamsRegion = "europe-west3"
	SpeedScheduleNewParamsRegionEuropeWest4         SpeedScheduleNewParamsRegion = "europe-west4"
	SpeedScheduleNewParamsRegionEuropeWest8         SpeedScheduleNewParamsRegion = "europe-west8"
	SpeedScheduleNewParamsRegionEuropeWest9         SpeedScheduleNewParamsRegion = "europe-west9"
	SpeedScheduleNewParamsRegionMeWest1             SpeedScheduleNewParamsRegion = "me-west1"
	SpeedScheduleNewParamsRegionSouthamericaEast1   SpeedScheduleNewParamsRegion = "southamerica-east1"
	SpeedScheduleNewParamsRegionUsCentral1          SpeedScheduleNewParamsRegion = "us-central1"
	SpeedScheduleNewParamsRegionUsEast1             SpeedScheduleNewParamsRegion = "us-east1"
	SpeedScheduleNewParamsRegionUsEast4             SpeedScheduleNewParamsRegion = "us-east4"
	SpeedScheduleNewParamsRegionUsSouth1            SpeedScheduleNewParamsRegion = "us-south1"
	SpeedScheduleNewParamsRegionUsWest1             SpeedScheduleNewParamsRegion = "us-west1"
)

type SpeedScheduleNewResponseEnvelope struct {
	Result SpeedScheduleNewResponse             `json:"result"`
	JSON   speedScheduleNewResponseEnvelopeJSON `json:"-"`
}

// speedScheduleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedScheduleNewResponseEnvelope]
type speedScheduleNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedScheduleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedScheduleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
