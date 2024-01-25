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

// ZoneSpeedAPIPageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSpeedAPIPageService] method
// instead.
type ZoneSpeedAPIPageService struct {
	Options []option.RequestOption
	Tests   *ZoneSpeedAPIPageTestService
}

// NewZoneSpeedAPIPageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSpeedAPIPageService(opts ...option.RequestOption) (r *ZoneSpeedAPIPageService) {
	r = &ZoneSpeedAPIPageService{}
	r.Options = opts
	r.Tests = NewZoneSpeedAPIPageTestService(opts...)
	return
}

// Lists all webpages which have been tested.
func (r *ZoneSpeedAPIPageService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSpeedAPIPageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists the core web vital metrics trend over time for a specific page.
func (r *ZoneSpeedAPIPageService) Trend(ctx context.Context, zoneIdentifier string, url string, query ZoneSpeedAPIPageTrendParams, opts ...option.RequestOption) (res *ZoneSpeedAPIPageTrendResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/trend", zoneIdentifier, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneSpeedAPIPageListResponse struct {
	Errors   []ZoneSpeedAPIPageListResponseError   `json:"errors"`
	Messages []ZoneSpeedAPIPageListResponseMessage `json:"messages"`
	Result   []ZoneSpeedAPIPageListResponseResult  `json:"result"`
	// Whether the API call was successful.
	Success bool                             `json:"success"`
	JSON    zoneSpeedAPIPageListResponseJSON `json:"-"`
}

// zoneSpeedAPIPageListResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIPageListResponse]
type zoneSpeedAPIPageListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageListResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSpeedAPIPageListResponseErrorJSON `json:"-"`
}

// zoneSpeedAPIPageListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIPageListResponseError]
type zoneSpeedAPIPageListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageListResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSpeedAPIPageListResponseMessageJSON `json:"-"`
}

// zoneSpeedAPIPageListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageListResponseMessage]
type zoneSpeedAPIPageListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageListResponseResult struct {
	// A test region with a label.
	Region ZoneSpeedAPIPageListResponseResultRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency ZoneSpeedAPIPageListResponseResultScheduleFrequency `json:"scheduleFrequency"`
	Tests             []ZoneSpeedAPIPageListResponseResultTest            `json:"tests"`
	// A URL.
	URL  string                                 `json:"url"`
	JSON zoneSpeedAPIPageListResponseResultJSON `json:"-"`
}

// zoneSpeedAPIPageListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIPageListResponseResult]
type zoneSpeedAPIPageListResponseResultJSON struct {
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	Tests             apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region with a label.
type ZoneSpeedAPIPageListResponseResultRegion struct {
	Label string `json:"label"`
	// A test region.
	Value ZoneSpeedAPIPageListResponseResultRegionValue `json:"value"`
	JSON  zoneSpeedAPIPageListResponseResultRegionJSON  `json:"-"`
}

// zoneSpeedAPIPageListResponseResultRegionJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageListResponseResultRegion]
type zoneSpeedAPIPageListResponseResultRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageListResponseResultRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type ZoneSpeedAPIPageListResponseResultRegionValue string

const (
	ZoneSpeedAPIPageListResponseResultRegionValueAsiaEast1           ZoneSpeedAPIPageListResponseResultRegionValue = "asia-east1"
	ZoneSpeedAPIPageListResponseResultRegionValueAsiaNortheast1      ZoneSpeedAPIPageListResponseResultRegionValue = "asia-northeast1"
	ZoneSpeedAPIPageListResponseResultRegionValueAsiaNortheast2      ZoneSpeedAPIPageListResponseResultRegionValue = "asia-northeast2"
	ZoneSpeedAPIPageListResponseResultRegionValueAsiaSouth1          ZoneSpeedAPIPageListResponseResultRegionValue = "asia-south1"
	ZoneSpeedAPIPageListResponseResultRegionValueAsiaSoutheast1      ZoneSpeedAPIPageListResponseResultRegionValue = "asia-southeast1"
	ZoneSpeedAPIPageListResponseResultRegionValueAustraliaSoutheast1 ZoneSpeedAPIPageListResponseResultRegionValue = "australia-southeast1"
	ZoneSpeedAPIPageListResponseResultRegionValueEuropeNorth1        ZoneSpeedAPIPageListResponseResultRegionValue = "europe-north1"
	ZoneSpeedAPIPageListResponseResultRegionValueEuropeSouthwest1    ZoneSpeedAPIPageListResponseResultRegionValue = "europe-southwest1"
	ZoneSpeedAPIPageListResponseResultRegionValueEuropeWest1         ZoneSpeedAPIPageListResponseResultRegionValue = "europe-west1"
	ZoneSpeedAPIPageListResponseResultRegionValueEuropeWest2         ZoneSpeedAPIPageListResponseResultRegionValue = "europe-west2"
	ZoneSpeedAPIPageListResponseResultRegionValueEuropeWest3         ZoneSpeedAPIPageListResponseResultRegionValue = "europe-west3"
	ZoneSpeedAPIPageListResponseResultRegionValueEuropeWest4         ZoneSpeedAPIPageListResponseResultRegionValue = "europe-west4"
	ZoneSpeedAPIPageListResponseResultRegionValueEuropeWest8         ZoneSpeedAPIPageListResponseResultRegionValue = "europe-west8"
	ZoneSpeedAPIPageListResponseResultRegionValueEuropeWest9         ZoneSpeedAPIPageListResponseResultRegionValue = "europe-west9"
	ZoneSpeedAPIPageListResponseResultRegionValueMeWest1             ZoneSpeedAPIPageListResponseResultRegionValue = "me-west1"
	ZoneSpeedAPIPageListResponseResultRegionValueSouthamericaEast1   ZoneSpeedAPIPageListResponseResultRegionValue = "southamerica-east1"
	ZoneSpeedAPIPageListResponseResultRegionValueUsCentral1          ZoneSpeedAPIPageListResponseResultRegionValue = "us-central1"
	ZoneSpeedAPIPageListResponseResultRegionValueUsEast1             ZoneSpeedAPIPageListResponseResultRegionValue = "us-east1"
	ZoneSpeedAPIPageListResponseResultRegionValueUsEast4             ZoneSpeedAPIPageListResponseResultRegionValue = "us-east4"
	ZoneSpeedAPIPageListResponseResultRegionValueUsSouth1            ZoneSpeedAPIPageListResponseResultRegionValue = "us-south1"
	ZoneSpeedAPIPageListResponseResultRegionValueUsWest1             ZoneSpeedAPIPageListResponseResultRegionValue = "us-west1"
)

// The frequency of the test.
type ZoneSpeedAPIPageListResponseResultScheduleFrequency string

const (
	ZoneSpeedAPIPageListResponseResultScheduleFrequencyDaily  ZoneSpeedAPIPageListResponseResultScheduleFrequency = "DAILY"
	ZoneSpeedAPIPageListResponseResultScheduleFrequencyWeekly ZoneSpeedAPIPageListResponseResultScheduleFrequency = "WEEKLY"
)

type ZoneSpeedAPIPageListResponseResultTest struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport ZoneSpeedAPIPageListResponseResultTestsDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport ZoneSpeedAPIPageListResponseResultTestsMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region ZoneSpeedAPIPageListResponseResultTestsRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency ZoneSpeedAPIPageListResponseResultTestsScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                                     `json:"url"`
	JSON zoneSpeedAPIPageListResponseResultTestJSON `json:"-"`
}

// zoneSpeedAPIPageListResponseResultTestJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageListResponseResultTest]
type zoneSpeedAPIPageListResponseResultTestJSON struct {
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

func (r *ZoneSpeedAPIPageListResponseResultTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type ZoneSpeedAPIPageListResponseResultTestsDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ZoneSpeedAPIPageListResponseResultTestsDesktopReportDeviceType `json:"deviceType"`
	Error      ZoneSpeedAPIPageListResponseResultTestsDesktopReportError      `json:"error"`
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
	State ZoneSpeedAPIPageListResponseResultTestsDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                                  `json:"tti"`
	JSON zoneSpeedAPIPageListResponseResultTestsDesktopReportJSON `json:"-"`
}

// zoneSpeedAPIPageListResponseResultTestsDesktopReportJSON contains the JSON
// metadata for the struct [ZoneSpeedAPIPageListResponseResultTestsDesktopReport]
type zoneSpeedAPIPageListResponseResultTestsDesktopReportJSON struct {
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

func (r *ZoneSpeedAPIPageListResponseResultTestsDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type ZoneSpeedAPIPageListResponseResultTestsDesktopReportDeviceType string

const (
	ZoneSpeedAPIPageListResponseResultTestsDesktopReportDeviceTypeDesktop ZoneSpeedAPIPageListResponseResultTestsDesktopReportDeviceType = "DESKTOP"
	ZoneSpeedAPIPageListResponseResultTestsDesktopReportDeviceTypeMobile  ZoneSpeedAPIPageListResponseResultTestsDesktopReportDeviceType = "MOBILE"
)

type ZoneSpeedAPIPageListResponseResultTestsDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code ZoneSpeedAPIPageListResponseResultTestsDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                        `json:"finalDisplayedUrl"`
	JSON              zoneSpeedAPIPageListResponseResultTestsDesktopReportErrorJSON `json:"-"`
}

// zoneSpeedAPIPageListResponseResultTestsDesktopReportErrorJSON contains the JSON
// metadata for the struct
// [ZoneSpeedAPIPageListResponseResultTestsDesktopReportError]
type zoneSpeedAPIPageListResponseResultTestsDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageListResponseResultTestsDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type ZoneSpeedAPIPageListResponseResultTestsDesktopReportErrorCode string

const (
	ZoneSpeedAPIPageListResponseResultTestsDesktopReportErrorCodeNotReachable      ZoneSpeedAPIPageListResponseResultTestsDesktopReportErrorCode = "NOT_REACHABLE"
	ZoneSpeedAPIPageListResponseResultTestsDesktopReportErrorCodeDNSFailure        ZoneSpeedAPIPageListResponseResultTestsDesktopReportErrorCode = "DNS_FAILURE"
	ZoneSpeedAPIPageListResponseResultTestsDesktopReportErrorCodeNotHTML           ZoneSpeedAPIPageListResponseResultTestsDesktopReportErrorCode = "NOT_HTML"
	ZoneSpeedAPIPageListResponseResultTestsDesktopReportErrorCodeLighthouseTimeout ZoneSpeedAPIPageListResponseResultTestsDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ZoneSpeedAPIPageListResponseResultTestsDesktopReportErrorCodeUnknown           ZoneSpeedAPIPageListResponseResultTestsDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ZoneSpeedAPIPageListResponseResultTestsDesktopReportState string

const (
	ZoneSpeedAPIPageListResponseResultTestsDesktopReportStateRunning  ZoneSpeedAPIPageListResponseResultTestsDesktopReportState = "RUNNING"
	ZoneSpeedAPIPageListResponseResultTestsDesktopReportStateComplete ZoneSpeedAPIPageListResponseResultTestsDesktopReportState = "COMPLETE"
	ZoneSpeedAPIPageListResponseResultTestsDesktopReportStateFailed   ZoneSpeedAPIPageListResponseResultTestsDesktopReportState = "FAILED"
)

// The Lighthouse report.
type ZoneSpeedAPIPageListResponseResultTestsMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ZoneSpeedAPIPageListResponseResultTestsMobileReportDeviceType `json:"deviceType"`
	Error      ZoneSpeedAPIPageListResponseResultTestsMobileReportError      `json:"error"`
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
	State ZoneSpeedAPIPageListResponseResultTestsMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                                 `json:"tti"`
	JSON zoneSpeedAPIPageListResponseResultTestsMobileReportJSON `json:"-"`
}

// zoneSpeedAPIPageListResponseResultTestsMobileReportJSON contains the JSON
// metadata for the struct [ZoneSpeedAPIPageListResponseResultTestsMobileReport]
type zoneSpeedAPIPageListResponseResultTestsMobileReportJSON struct {
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

func (r *ZoneSpeedAPIPageListResponseResultTestsMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type ZoneSpeedAPIPageListResponseResultTestsMobileReportDeviceType string

const (
	ZoneSpeedAPIPageListResponseResultTestsMobileReportDeviceTypeDesktop ZoneSpeedAPIPageListResponseResultTestsMobileReportDeviceType = "DESKTOP"
	ZoneSpeedAPIPageListResponseResultTestsMobileReportDeviceTypeMobile  ZoneSpeedAPIPageListResponseResultTestsMobileReportDeviceType = "MOBILE"
)

type ZoneSpeedAPIPageListResponseResultTestsMobileReportError struct {
	// The error code of the Lighthouse result.
	Code ZoneSpeedAPIPageListResponseResultTestsMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                       `json:"finalDisplayedUrl"`
	JSON              zoneSpeedAPIPageListResponseResultTestsMobileReportErrorJSON `json:"-"`
}

// zoneSpeedAPIPageListResponseResultTestsMobileReportErrorJSON contains the JSON
// metadata for the struct
// [ZoneSpeedAPIPageListResponseResultTestsMobileReportError]
type zoneSpeedAPIPageListResponseResultTestsMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageListResponseResultTestsMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type ZoneSpeedAPIPageListResponseResultTestsMobileReportErrorCode string

const (
	ZoneSpeedAPIPageListResponseResultTestsMobileReportErrorCodeNotReachable      ZoneSpeedAPIPageListResponseResultTestsMobileReportErrorCode = "NOT_REACHABLE"
	ZoneSpeedAPIPageListResponseResultTestsMobileReportErrorCodeDNSFailure        ZoneSpeedAPIPageListResponseResultTestsMobileReportErrorCode = "DNS_FAILURE"
	ZoneSpeedAPIPageListResponseResultTestsMobileReportErrorCodeNotHTML           ZoneSpeedAPIPageListResponseResultTestsMobileReportErrorCode = "NOT_HTML"
	ZoneSpeedAPIPageListResponseResultTestsMobileReportErrorCodeLighthouseTimeout ZoneSpeedAPIPageListResponseResultTestsMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ZoneSpeedAPIPageListResponseResultTestsMobileReportErrorCodeUnknown           ZoneSpeedAPIPageListResponseResultTestsMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ZoneSpeedAPIPageListResponseResultTestsMobileReportState string

const (
	ZoneSpeedAPIPageListResponseResultTestsMobileReportStateRunning  ZoneSpeedAPIPageListResponseResultTestsMobileReportState = "RUNNING"
	ZoneSpeedAPIPageListResponseResultTestsMobileReportStateComplete ZoneSpeedAPIPageListResponseResultTestsMobileReportState = "COMPLETE"
	ZoneSpeedAPIPageListResponseResultTestsMobileReportStateFailed   ZoneSpeedAPIPageListResponseResultTestsMobileReportState = "FAILED"
)

// A test region with a label.
type ZoneSpeedAPIPageListResponseResultTestsRegion struct {
	Label string `json:"label"`
	// A test region.
	Value ZoneSpeedAPIPageListResponseResultTestsRegionValue `json:"value"`
	JSON  zoneSpeedAPIPageListResponseResultTestsRegionJSON  `json:"-"`
}

// zoneSpeedAPIPageListResponseResultTestsRegionJSON contains the JSON metadata for
// the struct [ZoneSpeedAPIPageListResponseResultTestsRegion]
type zoneSpeedAPIPageListResponseResultTestsRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageListResponseResultTestsRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type ZoneSpeedAPIPageListResponseResultTestsRegionValue string

const (
	ZoneSpeedAPIPageListResponseResultTestsRegionValueAsiaEast1           ZoneSpeedAPIPageListResponseResultTestsRegionValue = "asia-east1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueAsiaNortheast1      ZoneSpeedAPIPageListResponseResultTestsRegionValue = "asia-northeast1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueAsiaNortheast2      ZoneSpeedAPIPageListResponseResultTestsRegionValue = "asia-northeast2"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueAsiaSouth1          ZoneSpeedAPIPageListResponseResultTestsRegionValue = "asia-south1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueAsiaSoutheast1      ZoneSpeedAPIPageListResponseResultTestsRegionValue = "asia-southeast1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueAustraliaSoutheast1 ZoneSpeedAPIPageListResponseResultTestsRegionValue = "australia-southeast1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueEuropeNorth1        ZoneSpeedAPIPageListResponseResultTestsRegionValue = "europe-north1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueEuropeSouthwest1    ZoneSpeedAPIPageListResponseResultTestsRegionValue = "europe-southwest1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueEuropeWest1         ZoneSpeedAPIPageListResponseResultTestsRegionValue = "europe-west1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueEuropeWest2         ZoneSpeedAPIPageListResponseResultTestsRegionValue = "europe-west2"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueEuropeWest3         ZoneSpeedAPIPageListResponseResultTestsRegionValue = "europe-west3"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueEuropeWest4         ZoneSpeedAPIPageListResponseResultTestsRegionValue = "europe-west4"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueEuropeWest8         ZoneSpeedAPIPageListResponseResultTestsRegionValue = "europe-west8"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueEuropeWest9         ZoneSpeedAPIPageListResponseResultTestsRegionValue = "europe-west9"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueMeWest1             ZoneSpeedAPIPageListResponseResultTestsRegionValue = "me-west1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueSouthamericaEast1   ZoneSpeedAPIPageListResponseResultTestsRegionValue = "southamerica-east1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueUsCentral1          ZoneSpeedAPIPageListResponseResultTestsRegionValue = "us-central1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueUsEast1             ZoneSpeedAPIPageListResponseResultTestsRegionValue = "us-east1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueUsEast4             ZoneSpeedAPIPageListResponseResultTestsRegionValue = "us-east4"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueUsSouth1            ZoneSpeedAPIPageListResponseResultTestsRegionValue = "us-south1"
	ZoneSpeedAPIPageListResponseResultTestsRegionValueUsWest1             ZoneSpeedAPIPageListResponseResultTestsRegionValue = "us-west1"
)

// The frequency of the test.
type ZoneSpeedAPIPageListResponseResultTestsScheduleFrequency string

const (
	ZoneSpeedAPIPageListResponseResultTestsScheduleFrequencyDaily  ZoneSpeedAPIPageListResponseResultTestsScheduleFrequency = "DAILY"
	ZoneSpeedAPIPageListResponseResultTestsScheduleFrequencyWeekly ZoneSpeedAPIPageListResponseResultTestsScheduleFrequency = "WEEKLY"
)

type ZoneSpeedAPIPageTrendResponse struct {
	Errors   []ZoneSpeedAPIPageTrendResponseError   `json:"errors"`
	Messages []ZoneSpeedAPIPageTrendResponseMessage `json:"messages"`
	Result   ZoneSpeedAPIPageTrendResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                              `json:"success"`
	JSON    zoneSpeedAPIPageTrendResponseJSON `json:"-"`
}

// zoneSpeedAPIPageTrendResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIPageTrendResponse]
type zoneSpeedAPIPageTrendResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTrendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTrendResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSpeedAPIPageTrendResponseErrorJSON `json:"-"`
}

// zoneSpeedAPIPageTrendResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIPageTrendResponseError]
type zoneSpeedAPIPageTrendResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTrendResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTrendResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSpeedAPIPageTrendResponseMessageJSON `json:"-"`
}

// zoneSpeedAPIPageTrendResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTrendResponseMessage]
type zoneSpeedAPIPageTrendResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTrendResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTrendResponseResult struct {
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
	Tti  []float64                               `json:"tti"`
	JSON zoneSpeedAPIPageTrendResponseResultJSON `json:"-"`
}

// zoneSpeedAPIPageTrendResponseResultJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTrendResponseResult]
type zoneSpeedAPIPageTrendResponseResultJSON struct {
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

func (r *ZoneSpeedAPIPageTrendResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTrendParams struct {
	// The type of device.
	DeviceType param.Field[ZoneSpeedAPIPageTrendParamsDeviceType] `query:"deviceType,required"`
	// A comma-separated list of metrics to include in the results.
	Metrics param.Field[string] `query:"metrics,required"`
	// A test region.
	Region param.Field[ZoneSpeedAPIPageTrendParamsRegion] `query:"region,required"`
	// The timezone of the start and end timestamps.
	Tz param.Field[string] `query:"tz,required"`
}

// URLQuery serializes [ZoneSpeedAPIPageTrendParams]'s query parameters as
// `url.Values`.
func (r ZoneSpeedAPIPageTrendParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of device.
type ZoneSpeedAPIPageTrendParamsDeviceType string

const (
	ZoneSpeedAPIPageTrendParamsDeviceTypeDesktop ZoneSpeedAPIPageTrendParamsDeviceType = "DESKTOP"
	ZoneSpeedAPIPageTrendParamsDeviceTypeMobile  ZoneSpeedAPIPageTrendParamsDeviceType = "MOBILE"
)

// A test region.
type ZoneSpeedAPIPageTrendParamsRegion string

const (
	ZoneSpeedAPIPageTrendParamsRegionAsiaEast1           ZoneSpeedAPIPageTrendParamsRegion = "asia-east1"
	ZoneSpeedAPIPageTrendParamsRegionAsiaNortheast1      ZoneSpeedAPIPageTrendParamsRegion = "asia-northeast1"
	ZoneSpeedAPIPageTrendParamsRegionAsiaNortheast2      ZoneSpeedAPIPageTrendParamsRegion = "asia-northeast2"
	ZoneSpeedAPIPageTrendParamsRegionAsiaSouth1          ZoneSpeedAPIPageTrendParamsRegion = "asia-south1"
	ZoneSpeedAPIPageTrendParamsRegionAsiaSoutheast1      ZoneSpeedAPIPageTrendParamsRegion = "asia-southeast1"
	ZoneSpeedAPIPageTrendParamsRegionAustraliaSoutheast1 ZoneSpeedAPIPageTrendParamsRegion = "australia-southeast1"
	ZoneSpeedAPIPageTrendParamsRegionEuropeNorth1        ZoneSpeedAPIPageTrendParamsRegion = "europe-north1"
	ZoneSpeedAPIPageTrendParamsRegionEuropeSouthwest1    ZoneSpeedAPIPageTrendParamsRegion = "europe-southwest1"
	ZoneSpeedAPIPageTrendParamsRegionEuropeWest1         ZoneSpeedAPIPageTrendParamsRegion = "europe-west1"
	ZoneSpeedAPIPageTrendParamsRegionEuropeWest2         ZoneSpeedAPIPageTrendParamsRegion = "europe-west2"
	ZoneSpeedAPIPageTrendParamsRegionEuropeWest3         ZoneSpeedAPIPageTrendParamsRegion = "europe-west3"
	ZoneSpeedAPIPageTrendParamsRegionEuropeWest4         ZoneSpeedAPIPageTrendParamsRegion = "europe-west4"
	ZoneSpeedAPIPageTrendParamsRegionEuropeWest8         ZoneSpeedAPIPageTrendParamsRegion = "europe-west8"
	ZoneSpeedAPIPageTrendParamsRegionEuropeWest9         ZoneSpeedAPIPageTrendParamsRegion = "europe-west9"
	ZoneSpeedAPIPageTrendParamsRegionMeWest1             ZoneSpeedAPIPageTrendParamsRegion = "me-west1"
	ZoneSpeedAPIPageTrendParamsRegionSouthamericaEast1   ZoneSpeedAPIPageTrendParamsRegion = "southamerica-east1"
	ZoneSpeedAPIPageTrendParamsRegionUsCentral1          ZoneSpeedAPIPageTrendParamsRegion = "us-central1"
	ZoneSpeedAPIPageTrendParamsRegionUsEast1             ZoneSpeedAPIPageTrendParamsRegion = "us-east1"
	ZoneSpeedAPIPageTrendParamsRegionUsEast4             ZoneSpeedAPIPageTrendParamsRegion = "us-east4"
	ZoneSpeedAPIPageTrendParamsRegionUsSouth1            ZoneSpeedAPIPageTrendParamsRegion = "us-south1"
	ZoneSpeedAPIPageTrendParamsRegionUsWest1             ZoneSpeedAPIPageTrendParamsRegion = "us-west1"
)
