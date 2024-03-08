// File generated from our OpenAPI spec by Stainless.

package speed

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// PageService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPageService] method instead.
type PageService struct {
	Options []option.RequestOption
}

// NewPageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPageService(opts ...option.RequestOption) (r *PageService) {
	r = &PageService{}
	r.Options = opts
	return
}

// Lists all webpages which have been tested.
func (r *PageService) List(ctx context.Context, query PageListParams, opts ...option.RequestOption) (res *[]PageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageListResponse struct {
	// A test region with a label.
	Region PageListResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency PageListResponseScheduleFrequency `json:"scheduleFrequency"`
	Tests             []PageListResponseTest            `json:"tests"`
	// A URL.
	URL  string               `json:"url"`
	JSON pageListResponseJSON `json:"-"`
}

// pageListResponseJSON contains the JSON metadata for the struct
// [PageListResponse]
type pageListResponseJSON struct {
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	Tests             apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseJSON) RawJSON() string {
	return r.raw
}

// A test region with a label.
type PageListResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value PageListResponseRegionValue `json:"value"`
	JSON  pageListResponseRegionJSON  `json:"-"`
}

// pageListResponseRegionJSON contains the JSON metadata for the struct
// [PageListResponseRegion]
type pageListResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageListResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type PageListResponseRegionValue string

const (
	PageListResponseRegionValueAsiaEast1           PageListResponseRegionValue = "asia-east1"
	PageListResponseRegionValueAsiaNortheast1      PageListResponseRegionValue = "asia-northeast1"
	PageListResponseRegionValueAsiaNortheast2      PageListResponseRegionValue = "asia-northeast2"
	PageListResponseRegionValueAsiaSouth1          PageListResponseRegionValue = "asia-south1"
	PageListResponseRegionValueAsiaSoutheast1      PageListResponseRegionValue = "asia-southeast1"
	PageListResponseRegionValueAustraliaSoutheast1 PageListResponseRegionValue = "australia-southeast1"
	PageListResponseRegionValueEuropeNorth1        PageListResponseRegionValue = "europe-north1"
	PageListResponseRegionValueEuropeSouthwest1    PageListResponseRegionValue = "europe-southwest1"
	PageListResponseRegionValueEuropeWest1         PageListResponseRegionValue = "europe-west1"
	PageListResponseRegionValueEuropeWest2         PageListResponseRegionValue = "europe-west2"
	PageListResponseRegionValueEuropeWest3         PageListResponseRegionValue = "europe-west3"
	PageListResponseRegionValueEuropeWest4         PageListResponseRegionValue = "europe-west4"
	PageListResponseRegionValueEuropeWest8         PageListResponseRegionValue = "europe-west8"
	PageListResponseRegionValueEuropeWest9         PageListResponseRegionValue = "europe-west9"
	PageListResponseRegionValueMeWest1             PageListResponseRegionValue = "me-west1"
	PageListResponseRegionValueSouthamericaEast1   PageListResponseRegionValue = "southamerica-east1"
	PageListResponseRegionValueUsCentral1          PageListResponseRegionValue = "us-central1"
	PageListResponseRegionValueUsEast1             PageListResponseRegionValue = "us-east1"
	PageListResponseRegionValueUsEast4             PageListResponseRegionValue = "us-east4"
	PageListResponseRegionValueUsSouth1            PageListResponseRegionValue = "us-south1"
	PageListResponseRegionValueUsWest1             PageListResponseRegionValue = "us-west1"
)

// The frequency of the test.
type PageListResponseScheduleFrequency string

const (
	PageListResponseScheduleFrequencyDaily  PageListResponseScheduleFrequency = "DAILY"
	PageListResponseScheduleFrequencyWeekly PageListResponseScheduleFrequency = "WEEKLY"
)

type PageListResponseTest struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport PageListResponseTestsDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport PageListResponseTestsMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region PageListResponseTestsRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency PageListResponseTestsScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                   `json:"url"`
	JSON pageListResponseTestJSON `json:"-"`
}

// pageListResponseTestJSON contains the JSON metadata for the struct
// [PageListResponseTest]
type pageListResponseTestJSON struct {
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

func (r *PageListResponseTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseTestJSON) RawJSON() string {
	return r.raw
}

// The Lighthouse report.
type PageListResponseTestsDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType PageListResponseTestsDesktopReportDeviceType `json:"deviceType"`
	Error      PageListResponseTestsDesktopReportError      `json:"error"`
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
	State PageListResponseTestsDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                `json:"tti"`
	JSON pageListResponseTestsDesktopReportJSON `json:"-"`
}

// pageListResponseTestsDesktopReportJSON contains the JSON metadata for the struct
// [PageListResponseTestsDesktopReport]
type pageListResponseTestsDesktopReportJSON struct {
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

func (r *PageListResponseTestsDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseTestsDesktopReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type PageListResponseTestsDesktopReportDeviceType string

const (
	PageListResponseTestsDesktopReportDeviceTypeDesktop PageListResponseTestsDesktopReportDeviceType = "DESKTOP"
	PageListResponseTestsDesktopReportDeviceTypeMobile  PageListResponseTestsDesktopReportDeviceType = "MOBILE"
)

type PageListResponseTestsDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code PageListResponseTestsDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                      `json:"finalDisplayedUrl"`
	JSON              pageListResponseTestsDesktopReportErrorJSON `json:"-"`
}

// pageListResponseTestsDesktopReportErrorJSON contains the JSON metadata for the
// struct [PageListResponseTestsDesktopReportError]
type pageListResponseTestsDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PageListResponseTestsDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseTestsDesktopReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type PageListResponseTestsDesktopReportErrorCode string

const (
	PageListResponseTestsDesktopReportErrorCodeNotReachable      PageListResponseTestsDesktopReportErrorCode = "NOT_REACHABLE"
	PageListResponseTestsDesktopReportErrorCodeDNSFailure        PageListResponseTestsDesktopReportErrorCode = "DNS_FAILURE"
	PageListResponseTestsDesktopReportErrorCodeNotHTML           PageListResponseTestsDesktopReportErrorCode = "NOT_HTML"
	PageListResponseTestsDesktopReportErrorCodeLighthouseTimeout PageListResponseTestsDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	PageListResponseTestsDesktopReportErrorCodeUnknown           PageListResponseTestsDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type PageListResponseTestsDesktopReportState string

const (
	PageListResponseTestsDesktopReportStateRunning  PageListResponseTestsDesktopReportState = "RUNNING"
	PageListResponseTestsDesktopReportStateComplete PageListResponseTestsDesktopReportState = "COMPLETE"
	PageListResponseTestsDesktopReportStateFailed   PageListResponseTestsDesktopReportState = "FAILED"
)

// The Lighthouse report.
type PageListResponseTestsMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType PageListResponseTestsMobileReportDeviceType `json:"deviceType"`
	Error      PageListResponseTestsMobileReportError      `json:"error"`
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
	State PageListResponseTestsMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                               `json:"tti"`
	JSON pageListResponseTestsMobileReportJSON `json:"-"`
}

// pageListResponseTestsMobileReportJSON contains the JSON metadata for the struct
// [PageListResponseTestsMobileReport]
type pageListResponseTestsMobileReportJSON struct {
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

func (r *PageListResponseTestsMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseTestsMobileReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type PageListResponseTestsMobileReportDeviceType string

const (
	PageListResponseTestsMobileReportDeviceTypeDesktop PageListResponseTestsMobileReportDeviceType = "DESKTOP"
	PageListResponseTestsMobileReportDeviceTypeMobile  PageListResponseTestsMobileReportDeviceType = "MOBILE"
)

type PageListResponseTestsMobileReportError struct {
	// The error code of the Lighthouse result.
	Code PageListResponseTestsMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                     `json:"finalDisplayedUrl"`
	JSON              pageListResponseTestsMobileReportErrorJSON `json:"-"`
}

// pageListResponseTestsMobileReportErrorJSON contains the JSON metadata for the
// struct [PageListResponseTestsMobileReportError]
type pageListResponseTestsMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PageListResponseTestsMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseTestsMobileReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type PageListResponseTestsMobileReportErrorCode string

const (
	PageListResponseTestsMobileReportErrorCodeNotReachable      PageListResponseTestsMobileReportErrorCode = "NOT_REACHABLE"
	PageListResponseTestsMobileReportErrorCodeDNSFailure        PageListResponseTestsMobileReportErrorCode = "DNS_FAILURE"
	PageListResponseTestsMobileReportErrorCodeNotHTML           PageListResponseTestsMobileReportErrorCode = "NOT_HTML"
	PageListResponseTestsMobileReportErrorCodeLighthouseTimeout PageListResponseTestsMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	PageListResponseTestsMobileReportErrorCodeUnknown           PageListResponseTestsMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type PageListResponseTestsMobileReportState string

const (
	PageListResponseTestsMobileReportStateRunning  PageListResponseTestsMobileReportState = "RUNNING"
	PageListResponseTestsMobileReportStateComplete PageListResponseTestsMobileReportState = "COMPLETE"
	PageListResponseTestsMobileReportStateFailed   PageListResponseTestsMobileReportState = "FAILED"
)

// A test region with a label.
type PageListResponseTestsRegion struct {
	Label string `json:"label"`
	// A test region.
	Value PageListResponseTestsRegionValue `json:"value"`
	JSON  pageListResponseTestsRegionJSON  `json:"-"`
}

// pageListResponseTestsRegionJSON contains the JSON metadata for the struct
// [PageListResponseTestsRegion]
type pageListResponseTestsRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageListResponseTestsRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseTestsRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type PageListResponseTestsRegionValue string

const (
	PageListResponseTestsRegionValueAsiaEast1           PageListResponseTestsRegionValue = "asia-east1"
	PageListResponseTestsRegionValueAsiaNortheast1      PageListResponseTestsRegionValue = "asia-northeast1"
	PageListResponseTestsRegionValueAsiaNortheast2      PageListResponseTestsRegionValue = "asia-northeast2"
	PageListResponseTestsRegionValueAsiaSouth1          PageListResponseTestsRegionValue = "asia-south1"
	PageListResponseTestsRegionValueAsiaSoutheast1      PageListResponseTestsRegionValue = "asia-southeast1"
	PageListResponseTestsRegionValueAustraliaSoutheast1 PageListResponseTestsRegionValue = "australia-southeast1"
	PageListResponseTestsRegionValueEuropeNorth1        PageListResponseTestsRegionValue = "europe-north1"
	PageListResponseTestsRegionValueEuropeSouthwest1    PageListResponseTestsRegionValue = "europe-southwest1"
	PageListResponseTestsRegionValueEuropeWest1         PageListResponseTestsRegionValue = "europe-west1"
	PageListResponseTestsRegionValueEuropeWest2         PageListResponseTestsRegionValue = "europe-west2"
	PageListResponseTestsRegionValueEuropeWest3         PageListResponseTestsRegionValue = "europe-west3"
	PageListResponseTestsRegionValueEuropeWest4         PageListResponseTestsRegionValue = "europe-west4"
	PageListResponseTestsRegionValueEuropeWest8         PageListResponseTestsRegionValue = "europe-west8"
	PageListResponseTestsRegionValueEuropeWest9         PageListResponseTestsRegionValue = "europe-west9"
	PageListResponseTestsRegionValueMeWest1             PageListResponseTestsRegionValue = "me-west1"
	PageListResponseTestsRegionValueSouthamericaEast1   PageListResponseTestsRegionValue = "southamerica-east1"
	PageListResponseTestsRegionValueUsCentral1          PageListResponseTestsRegionValue = "us-central1"
	PageListResponseTestsRegionValueUsEast1             PageListResponseTestsRegionValue = "us-east1"
	PageListResponseTestsRegionValueUsEast4             PageListResponseTestsRegionValue = "us-east4"
	PageListResponseTestsRegionValueUsSouth1            PageListResponseTestsRegionValue = "us-south1"
	PageListResponseTestsRegionValueUsWest1             PageListResponseTestsRegionValue = "us-west1"
)

// The frequency of the test.
type PageListResponseTestsScheduleFrequency string

const (
	PageListResponseTestsScheduleFrequencyDaily  PageListResponseTestsScheduleFrequency = "DAILY"
	PageListResponseTestsScheduleFrequencyWeekly PageListResponseTestsScheduleFrequency = "WEEKLY"
)

type PageListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageListResponseEnvelope struct {
	Result []PageListResponse           `json:"result"`
	JSON   pageListResponseEnvelopeJSON `json:"-"`
}

// pageListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageListResponseEnvelope]
type pageListResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
