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
	var env SpeedAPIScheduleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SpeedAPIScheduleNewResponse struct {
	// The test schedule.
	Schedule SpeedAPIScheduleNewResponseSchedule `json:"schedule"`
	Test     SpeedAPIScheduleNewResponseTest     `json:"test"`
	JSON     speedAPIScheduleNewResponseJSON     `json:"-"`
}

// speedAPIScheduleNewResponseJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleNewResponse]
type speedAPIScheduleNewResponseJSON struct {
	Schedule    apijson.Field
	Test        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The test schedule.
type SpeedAPIScheduleNewResponseSchedule struct {
	// The frequency of the test.
	Frequency SpeedAPIScheduleNewResponseScheduleFrequency `json:"frequency"`
	// A test region.
	Region SpeedAPIScheduleNewResponseScheduleRegion `json:"region"`
	// A URL.
	URL  string                                  `json:"url"`
	JSON speedAPIScheduleNewResponseScheduleJSON `json:"-"`
}

// speedAPIScheduleNewResponseScheduleJSON contains the JSON metadata for the
// struct [SpeedAPIScheduleNewResponseSchedule]
type speedAPIScheduleNewResponseScheduleJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency of the test.
type SpeedAPIScheduleNewResponseScheduleFrequency string

const (
	SpeedAPIScheduleNewResponseScheduleFrequencyDaily  SpeedAPIScheduleNewResponseScheduleFrequency = "DAILY"
	SpeedAPIScheduleNewResponseScheduleFrequencyWeekly SpeedAPIScheduleNewResponseScheduleFrequency = "WEEKLY"
)

// A test region.
type SpeedAPIScheduleNewResponseScheduleRegion string

const (
	SpeedAPIScheduleNewResponseScheduleRegionAsiaEast1           SpeedAPIScheduleNewResponseScheduleRegion = "asia-east1"
	SpeedAPIScheduleNewResponseScheduleRegionAsiaNortheast1      SpeedAPIScheduleNewResponseScheduleRegion = "asia-northeast1"
	SpeedAPIScheduleNewResponseScheduleRegionAsiaNortheast2      SpeedAPIScheduleNewResponseScheduleRegion = "asia-northeast2"
	SpeedAPIScheduleNewResponseScheduleRegionAsiaSouth1          SpeedAPIScheduleNewResponseScheduleRegion = "asia-south1"
	SpeedAPIScheduleNewResponseScheduleRegionAsiaSoutheast1      SpeedAPIScheduleNewResponseScheduleRegion = "asia-southeast1"
	SpeedAPIScheduleNewResponseScheduleRegionAustraliaSoutheast1 SpeedAPIScheduleNewResponseScheduleRegion = "australia-southeast1"
	SpeedAPIScheduleNewResponseScheduleRegionEuropeNorth1        SpeedAPIScheduleNewResponseScheduleRegion = "europe-north1"
	SpeedAPIScheduleNewResponseScheduleRegionEuropeSouthwest1    SpeedAPIScheduleNewResponseScheduleRegion = "europe-southwest1"
	SpeedAPIScheduleNewResponseScheduleRegionEuropeWest1         SpeedAPIScheduleNewResponseScheduleRegion = "europe-west1"
	SpeedAPIScheduleNewResponseScheduleRegionEuropeWest2         SpeedAPIScheduleNewResponseScheduleRegion = "europe-west2"
	SpeedAPIScheduleNewResponseScheduleRegionEuropeWest3         SpeedAPIScheduleNewResponseScheduleRegion = "europe-west3"
	SpeedAPIScheduleNewResponseScheduleRegionEuropeWest4         SpeedAPIScheduleNewResponseScheduleRegion = "europe-west4"
	SpeedAPIScheduleNewResponseScheduleRegionEuropeWest8         SpeedAPIScheduleNewResponseScheduleRegion = "europe-west8"
	SpeedAPIScheduleNewResponseScheduleRegionEuropeWest9         SpeedAPIScheduleNewResponseScheduleRegion = "europe-west9"
	SpeedAPIScheduleNewResponseScheduleRegionMeWest1             SpeedAPIScheduleNewResponseScheduleRegion = "me-west1"
	SpeedAPIScheduleNewResponseScheduleRegionSouthamericaEast1   SpeedAPIScheduleNewResponseScheduleRegion = "southamerica-east1"
	SpeedAPIScheduleNewResponseScheduleRegionUsCentral1          SpeedAPIScheduleNewResponseScheduleRegion = "us-central1"
	SpeedAPIScheduleNewResponseScheduleRegionUsEast1             SpeedAPIScheduleNewResponseScheduleRegion = "us-east1"
	SpeedAPIScheduleNewResponseScheduleRegionUsEast4             SpeedAPIScheduleNewResponseScheduleRegion = "us-east4"
	SpeedAPIScheduleNewResponseScheduleRegionUsSouth1            SpeedAPIScheduleNewResponseScheduleRegion = "us-south1"
	SpeedAPIScheduleNewResponseScheduleRegionUsWest1             SpeedAPIScheduleNewResponseScheduleRegion = "us-west1"
)

type SpeedAPIScheduleNewResponseTest struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport SpeedAPIScheduleNewResponseTestDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport SpeedAPIScheduleNewResponseTestMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region SpeedAPIScheduleNewResponseTestRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency SpeedAPIScheduleNewResponseTestScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                              `json:"url"`
	JSON speedAPIScheduleNewResponseTestJSON `json:"-"`
}

// speedAPIScheduleNewResponseTestJSON contains the JSON metadata for the struct
// [SpeedAPIScheduleNewResponseTest]
type speedAPIScheduleNewResponseTestJSON struct {
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

func (r *SpeedAPIScheduleNewResponseTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type SpeedAPIScheduleNewResponseTestDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPIScheduleNewResponseTestDesktopReportDeviceType `json:"deviceType"`
	Error      SpeedAPIScheduleNewResponseTestDesktopReportError      `json:"error"`
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
	State SpeedAPIScheduleNewResponseTestDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                          `json:"tti"`
	JSON speedAPIScheduleNewResponseTestDesktopReportJSON `json:"-"`
}

// speedAPIScheduleNewResponseTestDesktopReportJSON contains the JSON metadata for
// the struct [SpeedAPIScheduleNewResponseTestDesktopReport]
type speedAPIScheduleNewResponseTestDesktopReportJSON struct {
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

func (r *SpeedAPIScheduleNewResponseTestDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPIScheduleNewResponseTestDesktopReportDeviceType string

const (
	SpeedAPIScheduleNewResponseTestDesktopReportDeviceTypeDesktop SpeedAPIScheduleNewResponseTestDesktopReportDeviceType = "DESKTOP"
	SpeedAPIScheduleNewResponseTestDesktopReportDeviceTypeMobile  SpeedAPIScheduleNewResponseTestDesktopReportDeviceType = "MOBILE"
)

type SpeedAPIScheduleNewResponseTestDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPIScheduleNewResponseTestDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                `json:"finalDisplayedUrl"`
	JSON              speedAPIScheduleNewResponseTestDesktopReportErrorJSON `json:"-"`
}

// speedAPIScheduleNewResponseTestDesktopReportErrorJSON contains the JSON metadata
// for the struct [SpeedAPIScheduleNewResponseTestDesktopReportError]
type speedAPIScheduleNewResponseTestDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseTestDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPIScheduleNewResponseTestDesktopReportErrorCode string

const (
	SpeedAPIScheduleNewResponseTestDesktopReportErrorCodeNotReachable      SpeedAPIScheduleNewResponseTestDesktopReportErrorCode = "NOT_REACHABLE"
	SpeedAPIScheduleNewResponseTestDesktopReportErrorCodeDNSFailure        SpeedAPIScheduleNewResponseTestDesktopReportErrorCode = "DNS_FAILURE"
	SpeedAPIScheduleNewResponseTestDesktopReportErrorCodeNotHTML           SpeedAPIScheduleNewResponseTestDesktopReportErrorCode = "NOT_HTML"
	SpeedAPIScheduleNewResponseTestDesktopReportErrorCodeLighthouseTimeout SpeedAPIScheduleNewResponseTestDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPIScheduleNewResponseTestDesktopReportErrorCodeUnknown           SpeedAPIScheduleNewResponseTestDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPIScheduleNewResponseTestDesktopReportState string

const (
	SpeedAPIScheduleNewResponseTestDesktopReportStateRunning  SpeedAPIScheduleNewResponseTestDesktopReportState = "RUNNING"
	SpeedAPIScheduleNewResponseTestDesktopReportStateComplete SpeedAPIScheduleNewResponseTestDesktopReportState = "COMPLETE"
	SpeedAPIScheduleNewResponseTestDesktopReportStateFailed   SpeedAPIScheduleNewResponseTestDesktopReportState = "FAILED"
)

// The Lighthouse report.
type SpeedAPIScheduleNewResponseTestMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType SpeedAPIScheduleNewResponseTestMobileReportDeviceType `json:"deviceType"`
	Error      SpeedAPIScheduleNewResponseTestMobileReportError      `json:"error"`
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
	State SpeedAPIScheduleNewResponseTestMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                         `json:"tti"`
	JSON speedAPIScheduleNewResponseTestMobileReportJSON `json:"-"`
}

// speedAPIScheduleNewResponseTestMobileReportJSON contains the JSON metadata for
// the struct [SpeedAPIScheduleNewResponseTestMobileReport]
type speedAPIScheduleNewResponseTestMobileReportJSON struct {
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

func (r *SpeedAPIScheduleNewResponseTestMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type SpeedAPIScheduleNewResponseTestMobileReportDeviceType string

const (
	SpeedAPIScheduleNewResponseTestMobileReportDeviceTypeDesktop SpeedAPIScheduleNewResponseTestMobileReportDeviceType = "DESKTOP"
	SpeedAPIScheduleNewResponseTestMobileReportDeviceTypeMobile  SpeedAPIScheduleNewResponseTestMobileReportDeviceType = "MOBILE"
)

type SpeedAPIScheduleNewResponseTestMobileReportError struct {
	// The error code of the Lighthouse result.
	Code SpeedAPIScheduleNewResponseTestMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                               `json:"finalDisplayedUrl"`
	JSON              speedAPIScheduleNewResponseTestMobileReportErrorJSON `json:"-"`
}

// speedAPIScheduleNewResponseTestMobileReportErrorJSON contains the JSON metadata
// for the struct [SpeedAPIScheduleNewResponseTestMobileReportError]
type speedAPIScheduleNewResponseTestMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseTestMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type SpeedAPIScheduleNewResponseTestMobileReportErrorCode string

const (
	SpeedAPIScheduleNewResponseTestMobileReportErrorCodeNotReachable      SpeedAPIScheduleNewResponseTestMobileReportErrorCode = "NOT_REACHABLE"
	SpeedAPIScheduleNewResponseTestMobileReportErrorCodeDNSFailure        SpeedAPIScheduleNewResponseTestMobileReportErrorCode = "DNS_FAILURE"
	SpeedAPIScheduleNewResponseTestMobileReportErrorCodeNotHTML           SpeedAPIScheduleNewResponseTestMobileReportErrorCode = "NOT_HTML"
	SpeedAPIScheduleNewResponseTestMobileReportErrorCodeLighthouseTimeout SpeedAPIScheduleNewResponseTestMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	SpeedAPIScheduleNewResponseTestMobileReportErrorCodeUnknown           SpeedAPIScheduleNewResponseTestMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type SpeedAPIScheduleNewResponseTestMobileReportState string

const (
	SpeedAPIScheduleNewResponseTestMobileReportStateRunning  SpeedAPIScheduleNewResponseTestMobileReportState = "RUNNING"
	SpeedAPIScheduleNewResponseTestMobileReportStateComplete SpeedAPIScheduleNewResponseTestMobileReportState = "COMPLETE"
	SpeedAPIScheduleNewResponseTestMobileReportStateFailed   SpeedAPIScheduleNewResponseTestMobileReportState = "FAILED"
)

// A test region with a label.
type SpeedAPIScheduleNewResponseTestRegion struct {
	Label string `json:"label"`
	// A test region.
	Value SpeedAPIScheduleNewResponseTestRegionValue `json:"value"`
	JSON  speedAPIScheduleNewResponseTestRegionJSON  `json:"-"`
}

// speedAPIScheduleNewResponseTestRegionJSON contains the JSON metadata for the
// struct [SpeedAPIScheduleNewResponseTestRegion]
type speedAPIScheduleNewResponseTestRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseTestRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type SpeedAPIScheduleNewResponseTestRegionValue string

const (
	SpeedAPIScheduleNewResponseTestRegionValueAsiaEast1           SpeedAPIScheduleNewResponseTestRegionValue = "asia-east1"
	SpeedAPIScheduleNewResponseTestRegionValueAsiaNortheast1      SpeedAPIScheduleNewResponseTestRegionValue = "asia-northeast1"
	SpeedAPIScheduleNewResponseTestRegionValueAsiaNortheast2      SpeedAPIScheduleNewResponseTestRegionValue = "asia-northeast2"
	SpeedAPIScheduleNewResponseTestRegionValueAsiaSouth1          SpeedAPIScheduleNewResponseTestRegionValue = "asia-south1"
	SpeedAPIScheduleNewResponseTestRegionValueAsiaSoutheast1      SpeedAPIScheduleNewResponseTestRegionValue = "asia-southeast1"
	SpeedAPIScheduleNewResponseTestRegionValueAustraliaSoutheast1 SpeedAPIScheduleNewResponseTestRegionValue = "australia-southeast1"
	SpeedAPIScheduleNewResponseTestRegionValueEuropeNorth1        SpeedAPIScheduleNewResponseTestRegionValue = "europe-north1"
	SpeedAPIScheduleNewResponseTestRegionValueEuropeSouthwest1    SpeedAPIScheduleNewResponseTestRegionValue = "europe-southwest1"
	SpeedAPIScheduleNewResponseTestRegionValueEuropeWest1         SpeedAPIScheduleNewResponseTestRegionValue = "europe-west1"
	SpeedAPIScheduleNewResponseTestRegionValueEuropeWest2         SpeedAPIScheduleNewResponseTestRegionValue = "europe-west2"
	SpeedAPIScheduleNewResponseTestRegionValueEuropeWest3         SpeedAPIScheduleNewResponseTestRegionValue = "europe-west3"
	SpeedAPIScheduleNewResponseTestRegionValueEuropeWest4         SpeedAPIScheduleNewResponseTestRegionValue = "europe-west4"
	SpeedAPIScheduleNewResponseTestRegionValueEuropeWest8         SpeedAPIScheduleNewResponseTestRegionValue = "europe-west8"
	SpeedAPIScheduleNewResponseTestRegionValueEuropeWest9         SpeedAPIScheduleNewResponseTestRegionValue = "europe-west9"
	SpeedAPIScheduleNewResponseTestRegionValueMeWest1             SpeedAPIScheduleNewResponseTestRegionValue = "me-west1"
	SpeedAPIScheduleNewResponseTestRegionValueSouthamericaEast1   SpeedAPIScheduleNewResponseTestRegionValue = "southamerica-east1"
	SpeedAPIScheduleNewResponseTestRegionValueUsCentral1          SpeedAPIScheduleNewResponseTestRegionValue = "us-central1"
	SpeedAPIScheduleNewResponseTestRegionValueUsEast1             SpeedAPIScheduleNewResponseTestRegionValue = "us-east1"
	SpeedAPIScheduleNewResponseTestRegionValueUsEast4             SpeedAPIScheduleNewResponseTestRegionValue = "us-east4"
	SpeedAPIScheduleNewResponseTestRegionValueUsSouth1            SpeedAPIScheduleNewResponseTestRegionValue = "us-south1"
	SpeedAPIScheduleNewResponseTestRegionValueUsWest1             SpeedAPIScheduleNewResponseTestRegionValue = "us-west1"
)

// The frequency of the test.
type SpeedAPIScheduleNewResponseTestScheduleFrequency string

const (
	SpeedAPIScheduleNewResponseTestScheduleFrequencyDaily  SpeedAPIScheduleNewResponseTestScheduleFrequency = "DAILY"
	SpeedAPIScheduleNewResponseTestScheduleFrequencyWeekly SpeedAPIScheduleNewResponseTestScheduleFrequency = "WEEKLY"
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

type SpeedAPIScheduleNewResponseEnvelope struct {
	Errors   []SpeedAPIScheduleNewResponseEnvelopeErrors   `json:"errors"`
	Messages []SpeedAPIScheduleNewResponseEnvelopeMessages `json:"messages"`
	Result   SpeedAPIScheduleNewResponse                   `json:"result"`
	// Whether the API call was successful.
	Success bool                                    `json:"success"`
	JSON    speedAPIScheduleNewResponseEnvelopeJSON `json:"-"`
}

// speedAPIScheduleNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SpeedAPIScheduleNewResponseEnvelope]
type speedAPIScheduleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    speedAPIScheduleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// speedAPIScheduleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SpeedAPIScheduleNewResponseEnvelopeErrors]
type speedAPIScheduleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedAPIScheduleNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    speedAPIScheduleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// speedAPIScheduleNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SpeedAPIScheduleNewResponseEnvelopeMessages]
type speedAPIScheduleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedAPIScheduleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
