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

// ScheduleService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewScheduleService] method instead.
type ScheduleService struct {
	Options []option.RequestOption
}

// NewScheduleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScheduleService(opts ...option.RequestOption) (r *ScheduleService) {
	r = &ScheduleService{}
	r.Options = opts
	return
}

// Creates a scheduled test for a page.
func (r *ScheduleService) New(ctx context.Context, url string, params ScheduleNewParams, opts ...option.RequestOption) (res *ScheduleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScheduleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScheduleNewResponse struct {
	// The test schedule.
	Schedule ScheduleNewResponseSchedule `json:"schedule"`
	Test     ScheduleNewResponseTest     `json:"test"`
	JSON     scheduleNewResponseJSON     `json:"-"`
}

// scheduleNewResponseJSON contains the JSON metadata for the struct
// [ScheduleNewResponse]
type scheduleNewResponseJSON struct {
	Schedule    apijson.Field
	Test        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The test schedule.
type ScheduleNewResponseSchedule struct {
	// The frequency of the test.
	Frequency ScheduleNewResponseScheduleFrequency `json:"frequency"`
	// A test region.
	Region ScheduleNewResponseScheduleRegion `json:"region"`
	// A URL.
	URL  string                          `json:"url"`
	JSON scheduleNewResponseScheduleJSON `json:"-"`
}

// scheduleNewResponseScheduleJSON contains the JSON metadata for the struct
// [ScheduleNewResponseSchedule]
type scheduleNewResponseScheduleJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleNewResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseScheduleJSON) RawJSON() string {
	return r.raw
}

// The frequency of the test.
type ScheduleNewResponseScheduleFrequency string

const (
	ScheduleNewResponseScheduleFrequencyDaily  ScheduleNewResponseScheduleFrequency = "DAILY"
	ScheduleNewResponseScheduleFrequencyWeekly ScheduleNewResponseScheduleFrequency = "WEEKLY"
)

// A test region.
type ScheduleNewResponseScheduleRegion string

const (
	ScheduleNewResponseScheduleRegionAsiaEast1           ScheduleNewResponseScheduleRegion = "asia-east1"
	ScheduleNewResponseScheduleRegionAsiaNortheast1      ScheduleNewResponseScheduleRegion = "asia-northeast1"
	ScheduleNewResponseScheduleRegionAsiaNortheast2      ScheduleNewResponseScheduleRegion = "asia-northeast2"
	ScheduleNewResponseScheduleRegionAsiaSouth1          ScheduleNewResponseScheduleRegion = "asia-south1"
	ScheduleNewResponseScheduleRegionAsiaSoutheast1      ScheduleNewResponseScheduleRegion = "asia-southeast1"
	ScheduleNewResponseScheduleRegionAustraliaSoutheast1 ScheduleNewResponseScheduleRegion = "australia-southeast1"
	ScheduleNewResponseScheduleRegionEuropeNorth1        ScheduleNewResponseScheduleRegion = "europe-north1"
	ScheduleNewResponseScheduleRegionEuropeSouthwest1    ScheduleNewResponseScheduleRegion = "europe-southwest1"
	ScheduleNewResponseScheduleRegionEuropeWest1         ScheduleNewResponseScheduleRegion = "europe-west1"
	ScheduleNewResponseScheduleRegionEuropeWest2         ScheduleNewResponseScheduleRegion = "europe-west2"
	ScheduleNewResponseScheduleRegionEuropeWest3         ScheduleNewResponseScheduleRegion = "europe-west3"
	ScheduleNewResponseScheduleRegionEuropeWest4         ScheduleNewResponseScheduleRegion = "europe-west4"
	ScheduleNewResponseScheduleRegionEuropeWest8         ScheduleNewResponseScheduleRegion = "europe-west8"
	ScheduleNewResponseScheduleRegionEuropeWest9         ScheduleNewResponseScheduleRegion = "europe-west9"
	ScheduleNewResponseScheduleRegionMeWest1             ScheduleNewResponseScheduleRegion = "me-west1"
	ScheduleNewResponseScheduleRegionSouthamericaEast1   ScheduleNewResponseScheduleRegion = "southamerica-east1"
	ScheduleNewResponseScheduleRegionUsCentral1          ScheduleNewResponseScheduleRegion = "us-central1"
	ScheduleNewResponseScheduleRegionUsEast1             ScheduleNewResponseScheduleRegion = "us-east1"
	ScheduleNewResponseScheduleRegionUsEast4             ScheduleNewResponseScheduleRegion = "us-east4"
	ScheduleNewResponseScheduleRegionUsSouth1            ScheduleNewResponseScheduleRegion = "us-south1"
	ScheduleNewResponseScheduleRegionUsWest1             ScheduleNewResponseScheduleRegion = "us-west1"
)

type ScheduleNewResponseTest struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport ScheduleNewResponseTestDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport ScheduleNewResponseTestMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region ScheduleNewResponseTestRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency ScheduleNewResponseTestScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                      `json:"url"`
	JSON scheduleNewResponseTestJSON `json:"-"`
}

// scheduleNewResponseTestJSON contains the JSON metadata for the struct
// [ScheduleNewResponseTest]
type scheduleNewResponseTestJSON struct {
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

func (r *ScheduleNewResponseTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseTestJSON) RawJSON() string {
	return r.raw
}

// The Lighthouse report.
type ScheduleNewResponseTestDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ScheduleNewResponseTestDesktopReportDeviceType `json:"deviceType"`
	Error      ScheduleNewResponseTestDesktopReportError      `json:"error"`
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
	State ScheduleNewResponseTestDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                  `json:"tti"`
	JSON scheduleNewResponseTestDesktopReportJSON `json:"-"`
}

// scheduleNewResponseTestDesktopReportJSON contains the JSON metadata for the
// struct [ScheduleNewResponseTestDesktopReport]
type scheduleNewResponseTestDesktopReportJSON struct {
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

func (r *ScheduleNewResponseTestDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseTestDesktopReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type ScheduleNewResponseTestDesktopReportDeviceType string

const (
	ScheduleNewResponseTestDesktopReportDeviceTypeDesktop ScheduleNewResponseTestDesktopReportDeviceType = "DESKTOP"
	ScheduleNewResponseTestDesktopReportDeviceTypeMobile  ScheduleNewResponseTestDesktopReportDeviceType = "MOBILE"
)

type ScheduleNewResponseTestDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code ScheduleNewResponseTestDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                        `json:"finalDisplayedUrl"`
	JSON              scheduleNewResponseTestDesktopReportErrorJSON `json:"-"`
}

// scheduleNewResponseTestDesktopReportErrorJSON contains the JSON metadata for the
// struct [ScheduleNewResponseTestDesktopReportError]
type scheduleNewResponseTestDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScheduleNewResponseTestDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseTestDesktopReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type ScheduleNewResponseTestDesktopReportErrorCode string

const (
	ScheduleNewResponseTestDesktopReportErrorCodeNotReachable      ScheduleNewResponseTestDesktopReportErrorCode = "NOT_REACHABLE"
	ScheduleNewResponseTestDesktopReportErrorCodeDNSFailure        ScheduleNewResponseTestDesktopReportErrorCode = "DNS_FAILURE"
	ScheduleNewResponseTestDesktopReportErrorCodeNotHTML           ScheduleNewResponseTestDesktopReportErrorCode = "NOT_HTML"
	ScheduleNewResponseTestDesktopReportErrorCodeLighthouseTimeout ScheduleNewResponseTestDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ScheduleNewResponseTestDesktopReportErrorCodeUnknown           ScheduleNewResponseTestDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ScheduleNewResponseTestDesktopReportState string

const (
	ScheduleNewResponseTestDesktopReportStateRunning  ScheduleNewResponseTestDesktopReportState = "RUNNING"
	ScheduleNewResponseTestDesktopReportStateComplete ScheduleNewResponseTestDesktopReportState = "COMPLETE"
	ScheduleNewResponseTestDesktopReportStateFailed   ScheduleNewResponseTestDesktopReportState = "FAILED"
)

// The Lighthouse report.
type ScheduleNewResponseTestMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ScheduleNewResponseTestMobileReportDeviceType `json:"deviceType"`
	Error      ScheduleNewResponseTestMobileReportError      `json:"error"`
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
	State ScheduleNewResponseTestMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                 `json:"tti"`
	JSON scheduleNewResponseTestMobileReportJSON `json:"-"`
}

// scheduleNewResponseTestMobileReportJSON contains the JSON metadata for the
// struct [ScheduleNewResponseTestMobileReport]
type scheduleNewResponseTestMobileReportJSON struct {
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

func (r *ScheduleNewResponseTestMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseTestMobileReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type ScheduleNewResponseTestMobileReportDeviceType string

const (
	ScheduleNewResponseTestMobileReportDeviceTypeDesktop ScheduleNewResponseTestMobileReportDeviceType = "DESKTOP"
	ScheduleNewResponseTestMobileReportDeviceTypeMobile  ScheduleNewResponseTestMobileReportDeviceType = "MOBILE"
)

type ScheduleNewResponseTestMobileReportError struct {
	// The error code of the Lighthouse result.
	Code ScheduleNewResponseTestMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                       `json:"finalDisplayedUrl"`
	JSON              scheduleNewResponseTestMobileReportErrorJSON `json:"-"`
}

// scheduleNewResponseTestMobileReportErrorJSON contains the JSON metadata for the
// struct [ScheduleNewResponseTestMobileReportError]
type scheduleNewResponseTestMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScheduleNewResponseTestMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseTestMobileReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type ScheduleNewResponseTestMobileReportErrorCode string

const (
	ScheduleNewResponseTestMobileReportErrorCodeNotReachable      ScheduleNewResponseTestMobileReportErrorCode = "NOT_REACHABLE"
	ScheduleNewResponseTestMobileReportErrorCodeDNSFailure        ScheduleNewResponseTestMobileReportErrorCode = "DNS_FAILURE"
	ScheduleNewResponseTestMobileReportErrorCodeNotHTML           ScheduleNewResponseTestMobileReportErrorCode = "NOT_HTML"
	ScheduleNewResponseTestMobileReportErrorCodeLighthouseTimeout ScheduleNewResponseTestMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ScheduleNewResponseTestMobileReportErrorCodeUnknown           ScheduleNewResponseTestMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ScheduleNewResponseTestMobileReportState string

const (
	ScheduleNewResponseTestMobileReportStateRunning  ScheduleNewResponseTestMobileReportState = "RUNNING"
	ScheduleNewResponseTestMobileReportStateComplete ScheduleNewResponseTestMobileReportState = "COMPLETE"
	ScheduleNewResponseTestMobileReportStateFailed   ScheduleNewResponseTestMobileReportState = "FAILED"
)

// A test region with a label.
type ScheduleNewResponseTestRegion struct {
	Label string `json:"label"`
	// A test region.
	Value ScheduleNewResponseTestRegionValue `json:"value"`
	JSON  scheduleNewResponseTestRegionJSON  `json:"-"`
}

// scheduleNewResponseTestRegionJSON contains the JSON metadata for the struct
// [ScheduleNewResponseTestRegion]
type scheduleNewResponseTestRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleNewResponseTestRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseTestRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type ScheduleNewResponseTestRegionValue string

const (
	ScheduleNewResponseTestRegionValueAsiaEast1           ScheduleNewResponseTestRegionValue = "asia-east1"
	ScheduleNewResponseTestRegionValueAsiaNortheast1      ScheduleNewResponseTestRegionValue = "asia-northeast1"
	ScheduleNewResponseTestRegionValueAsiaNortheast2      ScheduleNewResponseTestRegionValue = "asia-northeast2"
	ScheduleNewResponseTestRegionValueAsiaSouth1          ScheduleNewResponseTestRegionValue = "asia-south1"
	ScheduleNewResponseTestRegionValueAsiaSoutheast1      ScheduleNewResponseTestRegionValue = "asia-southeast1"
	ScheduleNewResponseTestRegionValueAustraliaSoutheast1 ScheduleNewResponseTestRegionValue = "australia-southeast1"
	ScheduleNewResponseTestRegionValueEuropeNorth1        ScheduleNewResponseTestRegionValue = "europe-north1"
	ScheduleNewResponseTestRegionValueEuropeSouthwest1    ScheduleNewResponseTestRegionValue = "europe-southwest1"
	ScheduleNewResponseTestRegionValueEuropeWest1         ScheduleNewResponseTestRegionValue = "europe-west1"
	ScheduleNewResponseTestRegionValueEuropeWest2         ScheduleNewResponseTestRegionValue = "europe-west2"
	ScheduleNewResponseTestRegionValueEuropeWest3         ScheduleNewResponseTestRegionValue = "europe-west3"
	ScheduleNewResponseTestRegionValueEuropeWest4         ScheduleNewResponseTestRegionValue = "europe-west4"
	ScheduleNewResponseTestRegionValueEuropeWest8         ScheduleNewResponseTestRegionValue = "europe-west8"
	ScheduleNewResponseTestRegionValueEuropeWest9         ScheduleNewResponseTestRegionValue = "europe-west9"
	ScheduleNewResponseTestRegionValueMeWest1             ScheduleNewResponseTestRegionValue = "me-west1"
	ScheduleNewResponseTestRegionValueSouthamericaEast1   ScheduleNewResponseTestRegionValue = "southamerica-east1"
	ScheduleNewResponseTestRegionValueUsCentral1          ScheduleNewResponseTestRegionValue = "us-central1"
	ScheduleNewResponseTestRegionValueUsEast1             ScheduleNewResponseTestRegionValue = "us-east1"
	ScheduleNewResponseTestRegionValueUsEast4             ScheduleNewResponseTestRegionValue = "us-east4"
	ScheduleNewResponseTestRegionValueUsSouth1            ScheduleNewResponseTestRegionValue = "us-south1"
	ScheduleNewResponseTestRegionValueUsWest1             ScheduleNewResponseTestRegionValue = "us-west1"
)

// The frequency of the test.
type ScheduleNewResponseTestScheduleFrequency string

const (
	ScheduleNewResponseTestScheduleFrequencyDaily  ScheduleNewResponseTestScheduleFrequency = "DAILY"
	ScheduleNewResponseTestScheduleFrequencyWeekly ScheduleNewResponseTestScheduleFrequency = "WEEKLY"
)

type ScheduleNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[ScheduleNewParamsRegion] `query:"region"`
}

// URLQuery serializes [ScheduleNewParams]'s query parameters as `url.Values`.
func (r ScheduleNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type ScheduleNewParamsRegion string

const (
	ScheduleNewParamsRegionAsiaEast1           ScheduleNewParamsRegion = "asia-east1"
	ScheduleNewParamsRegionAsiaNortheast1      ScheduleNewParamsRegion = "asia-northeast1"
	ScheduleNewParamsRegionAsiaNortheast2      ScheduleNewParamsRegion = "asia-northeast2"
	ScheduleNewParamsRegionAsiaSouth1          ScheduleNewParamsRegion = "asia-south1"
	ScheduleNewParamsRegionAsiaSoutheast1      ScheduleNewParamsRegion = "asia-southeast1"
	ScheduleNewParamsRegionAustraliaSoutheast1 ScheduleNewParamsRegion = "australia-southeast1"
	ScheduleNewParamsRegionEuropeNorth1        ScheduleNewParamsRegion = "europe-north1"
	ScheduleNewParamsRegionEuropeSouthwest1    ScheduleNewParamsRegion = "europe-southwest1"
	ScheduleNewParamsRegionEuropeWest1         ScheduleNewParamsRegion = "europe-west1"
	ScheduleNewParamsRegionEuropeWest2         ScheduleNewParamsRegion = "europe-west2"
	ScheduleNewParamsRegionEuropeWest3         ScheduleNewParamsRegion = "europe-west3"
	ScheduleNewParamsRegionEuropeWest4         ScheduleNewParamsRegion = "europe-west4"
	ScheduleNewParamsRegionEuropeWest8         ScheduleNewParamsRegion = "europe-west8"
	ScheduleNewParamsRegionEuropeWest9         ScheduleNewParamsRegion = "europe-west9"
	ScheduleNewParamsRegionMeWest1             ScheduleNewParamsRegion = "me-west1"
	ScheduleNewParamsRegionSouthamericaEast1   ScheduleNewParamsRegion = "southamerica-east1"
	ScheduleNewParamsRegionUsCentral1          ScheduleNewParamsRegion = "us-central1"
	ScheduleNewParamsRegionUsEast1             ScheduleNewParamsRegion = "us-east1"
	ScheduleNewParamsRegionUsEast4             ScheduleNewParamsRegion = "us-east4"
	ScheduleNewParamsRegionUsSouth1            ScheduleNewParamsRegion = "us-south1"
	ScheduleNewParamsRegionUsWest1             ScheduleNewParamsRegion = "us-west1"
)

type ScheduleNewResponseEnvelope struct {
	Result ScheduleNewResponse             `json:"result"`
	JSON   scheduleNewResponseEnvelopeJSON `json:"-"`
}

// scheduleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScheduleNewResponseEnvelope]
type scheduleNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
