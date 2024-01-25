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

// ZoneSpeedAPIScheduleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSpeedAPIScheduleService]
// method instead.
type ZoneSpeedAPIScheduleService struct {
	Options []option.RequestOption
}

// NewZoneSpeedAPIScheduleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSpeedAPIScheduleService(opts ...option.RequestOption) (r *ZoneSpeedAPIScheduleService) {
	r = &ZoneSpeedAPIScheduleService{}
	r.Options = opts
	return
}

// Creates a scheduled test for a page.
func (r *ZoneSpeedAPIScheduleService) New(ctx context.Context, zoneIdentifier string, url string, body ZoneSpeedAPIScheduleNewParams, opts ...option.RequestOption) (res *ZoneSpeedAPIScheduleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneIdentifier, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the test schedule for a page in a specific region.
func (r *ZoneSpeedAPIScheduleService) Get(ctx context.Context, zoneIdentifier string, url string, query ZoneSpeedAPIScheduleGetParams, opts ...option.RequestOption) (res *ZoneSpeedAPIScheduleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneIdentifier, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a scheduled test for a page.
func (r *ZoneSpeedAPIScheduleService) Delete(ctx context.Context, zoneIdentifier string, url string, body ZoneSpeedAPIScheduleDeleteParams, opts ...option.RequestOption) (res *ZoneSpeedAPIScheduleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneIdentifier, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type ZoneSpeedAPIScheduleNewResponse struct {
	Errors   []ZoneSpeedAPIScheduleNewResponseError   `json:"errors"`
	Messages []ZoneSpeedAPIScheduleNewResponseMessage `json:"messages"`
	Result   ZoneSpeedAPIScheduleNewResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                                `json:"success"`
	JSON    zoneSpeedAPIScheduleNewResponseJSON `json:"-"`
}

// zoneSpeedAPIScheduleNewResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIScheduleNewResponse]
type zoneSpeedAPIScheduleNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIScheduleNewResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSpeedAPIScheduleNewResponseErrorJSON `json:"-"`
}

// zoneSpeedAPIScheduleNewResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIScheduleNewResponseError]
type zoneSpeedAPIScheduleNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIScheduleNewResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSpeedAPIScheduleNewResponseMessageJSON `json:"-"`
}

// zoneSpeedAPIScheduleNewResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIScheduleNewResponseMessage]
type zoneSpeedAPIScheduleNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIScheduleNewResponseResult struct {
	// The test schedule.
	Schedule ZoneSpeedAPIScheduleNewResponseResultSchedule `json:"schedule"`
	Test     ZoneSpeedAPIScheduleNewResponseResultTest     `json:"test"`
	JSON     zoneSpeedAPIScheduleNewResponseResultJSON     `json:"-"`
}

// zoneSpeedAPIScheduleNewResponseResultJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIScheduleNewResponseResult]
type zoneSpeedAPIScheduleNewResponseResultJSON struct {
	Schedule    apijson.Field
	Test        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The test schedule.
type ZoneSpeedAPIScheduleNewResponseResultSchedule struct {
	// The frequency of the test.
	Frequency ZoneSpeedAPIScheduleNewResponseResultScheduleFrequency `json:"frequency"`
	// A test region.
	Region ZoneSpeedAPIScheduleNewResponseResultScheduleRegion `json:"region"`
	// A URL.
	URL  string                                            `json:"url"`
	JSON zoneSpeedAPIScheduleNewResponseResultScheduleJSON `json:"-"`
}

// zoneSpeedAPIScheduleNewResponseResultScheduleJSON contains the JSON metadata for
// the struct [ZoneSpeedAPIScheduleNewResponseResultSchedule]
type zoneSpeedAPIScheduleNewResponseResultScheduleJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleNewResponseResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency of the test.
type ZoneSpeedAPIScheduleNewResponseResultScheduleFrequency string

const (
	ZoneSpeedAPIScheduleNewResponseResultScheduleFrequencyDaily  ZoneSpeedAPIScheduleNewResponseResultScheduleFrequency = "DAILY"
	ZoneSpeedAPIScheduleNewResponseResultScheduleFrequencyWeekly ZoneSpeedAPIScheduleNewResponseResultScheduleFrequency = "WEEKLY"
)

// A test region.
type ZoneSpeedAPIScheduleNewResponseResultScheduleRegion string

const (
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionAsiaEast1           ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "asia-east1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionAsiaNortheast1      ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "asia-northeast1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionAsiaNortheast2      ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "asia-northeast2"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionAsiaSouth1          ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "asia-south1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionAsiaSoutheast1      ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "asia-southeast1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionAustraliaSoutheast1 ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "australia-southeast1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionEuropeNorth1        ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "europe-north1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionEuropeSouthwest1    ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "europe-southwest1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionEuropeWest1         ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "europe-west1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionEuropeWest2         ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "europe-west2"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionEuropeWest3         ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "europe-west3"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionEuropeWest4         ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "europe-west4"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionEuropeWest8         ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "europe-west8"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionEuropeWest9         ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "europe-west9"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionMeWest1             ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "me-west1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionSouthamericaEast1   ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "southamerica-east1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionUsCentral1          ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "us-central1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionUsEast1             ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "us-east1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionUsEast4             ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "us-east4"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionUsSouth1            ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "us-south1"
	ZoneSpeedAPIScheduleNewResponseResultScheduleRegionUsWest1             ZoneSpeedAPIScheduleNewResponseResultScheduleRegion = "us-west1"
)

type ZoneSpeedAPIScheduleNewResponseResultTest struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport ZoneSpeedAPIScheduleNewResponseResultTestDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport ZoneSpeedAPIScheduleNewResponseResultTestMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region ZoneSpeedAPIScheduleNewResponseResultTestRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency ZoneSpeedAPIScheduleNewResponseResultTestScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                                        `json:"url"`
	JSON zoneSpeedAPIScheduleNewResponseResultTestJSON `json:"-"`
}

// zoneSpeedAPIScheduleNewResponseResultTestJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIScheduleNewResponseResultTest]
type zoneSpeedAPIScheduleNewResponseResultTestJSON struct {
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

func (r *ZoneSpeedAPIScheduleNewResponseResultTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type ZoneSpeedAPIScheduleNewResponseResultTestDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportDeviceType `json:"deviceType"`
	Error      ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportError      `json:"error"`
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
	State ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                                    `json:"tti"`
	JSON zoneSpeedAPIScheduleNewResponseResultTestDesktopReportJSON `json:"-"`
}

// zoneSpeedAPIScheduleNewResponseResultTestDesktopReportJSON contains the JSON
// metadata for the struct [ZoneSpeedAPIScheduleNewResponseResultTestDesktopReport]
type zoneSpeedAPIScheduleNewResponseResultTestDesktopReportJSON struct {
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

func (r *ZoneSpeedAPIScheduleNewResponseResultTestDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportDeviceType string

const (
	ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportDeviceTypeDesktop ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportDeviceType = "DESKTOP"
	ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportDeviceTypeMobile  ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportDeviceType = "MOBILE"
)

type ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                          `json:"finalDisplayedUrl"`
	JSON              zoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorJSON `json:"-"`
}

// zoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorJSON contains the
// JSON metadata for the struct
// [ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportError]
type zoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode string

const (
	ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorCodeNotReachable      ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode = "NOT_REACHABLE"
	ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorCodeDNSFailure        ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode = "DNS_FAILURE"
	ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorCodeNotHTML           ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode = "NOT_HTML"
	ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorCodeLighthouseTimeout ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorCodeUnknown           ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportState string

const (
	ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportStateRunning  ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportState = "RUNNING"
	ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportStateComplete ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportState = "COMPLETE"
	ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportStateFailed   ZoneSpeedAPIScheduleNewResponseResultTestDesktopReportState = "FAILED"
)

// The Lighthouse report.
type ZoneSpeedAPIScheduleNewResponseResultTestMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ZoneSpeedAPIScheduleNewResponseResultTestMobileReportDeviceType `json:"deviceType"`
	Error      ZoneSpeedAPIScheduleNewResponseResultTestMobileReportError      `json:"error"`
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
	State ZoneSpeedAPIScheduleNewResponseResultTestMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                                   `json:"tti"`
	JSON zoneSpeedAPIScheduleNewResponseResultTestMobileReportJSON `json:"-"`
}

// zoneSpeedAPIScheduleNewResponseResultTestMobileReportJSON contains the JSON
// metadata for the struct [ZoneSpeedAPIScheduleNewResponseResultTestMobileReport]
type zoneSpeedAPIScheduleNewResponseResultTestMobileReportJSON struct {
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

func (r *ZoneSpeedAPIScheduleNewResponseResultTestMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type ZoneSpeedAPIScheduleNewResponseResultTestMobileReportDeviceType string

const (
	ZoneSpeedAPIScheduleNewResponseResultTestMobileReportDeviceTypeDesktop ZoneSpeedAPIScheduleNewResponseResultTestMobileReportDeviceType = "DESKTOP"
	ZoneSpeedAPIScheduleNewResponseResultTestMobileReportDeviceTypeMobile  ZoneSpeedAPIScheduleNewResponseResultTestMobileReportDeviceType = "MOBILE"
)

type ZoneSpeedAPIScheduleNewResponseResultTestMobileReportError struct {
	// The error code of the Lighthouse result.
	Code ZoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                         `json:"finalDisplayedUrl"`
	JSON              zoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorJSON `json:"-"`
}

// zoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorJSON contains the JSON
// metadata for the struct
// [ZoneSpeedAPIScheduleNewResponseResultTestMobileReportError]
type zoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleNewResponseResultTestMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type ZoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorCode string

const (
	ZoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorCodeNotReachable      ZoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorCode = "NOT_REACHABLE"
	ZoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorCodeDNSFailure        ZoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorCode = "DNS_FAILURE"
	ZoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorCodeNotHTML           ZoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorCode = "NOT_HTML"
	ZoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorCodeLighthouseTimeout ZoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ZoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorCodeUnknown           ZoneSpeedAPIScheduleNewResponseResultTestMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ZoneSpeedAPIScheduleNewResponseResultTestMobileReportState string

const (
	ZoneSpeedAPIScheduleNewResponseResultTestMobileReportStateRunning  ZoneSpeedAPIScheduleNewResponseResultTestMobileReportState = "RUNNING"
	ZoneSpeedAPIScheduleNewResponseResultTestMobileReportStateComplete ZoneSpeedAPIScheduleNewResponseResultTestMobileReportState = "COMPLETE"
	ZoneSpeedAPIScheduleNewResponseResultTestMobileReportStateFailed   ZoneSpeedAPIScheduleNewResponseResultTestMobileReportState = "FAILED"
)

// A test region with a label.
type ZoneSpeedAPIScheduleNewResponseResultTestRegion struct {
	Label string `json:"label"`
	// A test region.
	Value ZoneSpeedAPIScheduleNewResponseResultTestRegionValue `json:"value"`
	JSON  zoneSpeedAPIScheduleNewResponseResultTestRegionJSON  `json:"-"`
}

// zoneSpeedAPIScheduleNewResponseResultTestRegionJSON contains the JSON metadata
// for the struct [ZoneSpeedAPIScheduleNewResponseResultTestRegion]
type zoneSpeedAPIScheduleNewResponseResultTestRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleNewResponseResultTestRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type ZoneSpeedAPIScheduleNewResponseResultTestRegionValue string

const (
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueAsiaEast1           ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "asia-east1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueAsiaNortheast1      ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "asia-northeast1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueAsiaNortheast2      ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "asia-northeast2"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueAsiaSouth1          ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "asia-south1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueAsiaSoutheast1      ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "asia-southeast1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueAustraliaSoutheast1 ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "australia-southeast1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueEuropeNorth1        ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "europe-north1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueEuropeSouthwest1    ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "europe-southwest1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueEuropeWest1         ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "europe-west1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueEuropeWest2         ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "europe-west2"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueEuropeWest3         ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "europe-west3"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueEuropeWest4         ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "europe-west4"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueEuropeWest8         ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "europe-west8"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueEuropeWest9         ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "europe-west9"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueMeWest1             ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "me-west1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueSouthamericaEast1   ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "southamerica-east1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueUsCentral1          ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "us-central1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueUsEast1             ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "us-east1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueUsEast4             ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "us-east4"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueUsSouth1            ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "us-south1"
	ZoneSpeedAPIScheduleNewResponseResultTestRegionValueUsWest1             ZoneSpeedAPIScheduleNewResponseResultTestRegionValue = "us-west1"
)

// The frequency of the test.
type ZoneSpeedAPIScheduleNewResponseResultTestScheduleFrequency string

const (
	ZoneSpeedAPIScheduleNewResponseResultTestScheduleFrequencyDaily  ZoneSpeedAPIScheduleNewResponseResultTestScheduleFrequency = "DAILY"
	ZoneSpeedAPIScheduleNewResponseResultTestScheduleFrequencyWeekly ZoneSpeedAPIScheduleNewResponseResultTestScheduleFrequency = "WEEKLY"
)

type ZoneSpeedAPIScheduleGetResponse struct {
	Errors   []ZoneSpeedAPIScheduleGetResponseError   `json:"errors"`
	Messages []ZoneSpeedAPIScheduleGetResponseMessage `json:"messages"`
	// The test schedule.
	Result ZoneSpeedAPIScheduleGetResponseResult `json:"result"`
	// Whether the API call was successful.
	Success bool                                `json:"success"`
	JSON    zoneSpeedAPIScheduleGetResponseJSON `json:"-"`
}

// zoneSpeedAPIScheduleGetResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIScheduleGetResponse]
type zoneSpeedAPIScheduleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIScheduleGetResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSpeedAPIScheduleGetResponseErrorJSON `json:"-"`
}

// zoneSpeedAPIScheduleGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIScheduleGetResponseError]
type zoneSpeedAPIScheduleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIScheduleGetResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSpeedAPIScheduleGetResponseMessageJSON `json:"-"`
}

// zoneSpeedAPIScheduleGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIScheduleGetResponseMessage]
type zoneSpeedAPIScheduleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The test schedule.
type ZoneSpeedAPIScheduleGetResponseResult struct {
	// The frequency of the test.
	Frequency ZoneSpeedAPIScheduleGetResponseResultFrequency `json:"frequency"`
	// A test region.
	Region ZoneSpeedAPIScheduleGetResponseResultRegion `json:"region"`
	// A URL.
	URL  string                                    `json:"url"`
	JSON zoneSpeedAPIScheduleGetResponseResultJSON `json:"-"`
}

// zoneSpeedAPIScheduleGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIScheduleGetResponseResult]
type zoneSpeedAPIScheduleGetResponseResultJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency of the test.
type ZoneSpeedAPIScheduleGetResponseResultFrequency string

const (
	ZoneSpeedAPIScheduleGetResponseResultFrequencyDaily  ZoneSpeedAPIScheduleGetResponseResultFrequency = "DAILY"
	ZoneSpeedAPIScheduleGetResponseResultFrequencyWeekly ZoneSpeedAPIScheduleGetResponseResultFrequency = "WEEKLY"
)

// A test region.
type ZoneSpeedAPIScheduleGetResponseResultRegion string

const (
	ZoneSpeedAPIScheduleGetResponseResultRegionAsiaEast1           ZoneSpeedAPIScheduleGetResponseResultRegion = "asia-east1"
	ZoneSpeedAPIScheduleGetResponseResultRegionAsiaNortheast1      ZoneSpeedAPIScheduleGetResponseResultRegion = "asia-northeast1"
	ZoneSpeedAPIScheduleGetResponseResultRegionAsiaNortheast2      ZoneSpeedAPIScheduleGetResponseResultRegion = "asia-northeast2"
	ZoneSpeedAPIScheduleGetResponseResultRegionAsiaSouth1          ZoneSpeedAPIScheduleGetResponseResultRegion = "asia-south1"
	ZoneSpeedAPIScheduleGetResponseResultRegionAsiaSoutheast1      ZoneSpeedAPIScheduleGetResponseResultRegion = "asia-southeast1"
	ZoneSpeedAPIScheduleGetResponseResultRegionAustraliaSoutheast1 ZoneSpeedAPIScheduleGetResponseResultRegion = "australia-southeast1"
	ZoneSpeedAPIScheduleGetResponseResultRegionEuropeNorth1        ZoneSpeedAPIScheduleGetResponseResultRegion = "europe-north1"
	ZoneSpeedAPIScheduleGetResponseResultRegionEuropeSouthwest1    ZoneSpeedAPIScheduleGetResponseResultRegion = "europe-southwest1"
	ZoneSpeedAPIScheduleGetResponseResultRegionEuropeWest1         ZoneSpeedAPIScheduleGetResponseResultRegion = "europe-west1"
	ZoneSpeedAPIScheduleGetResponseResultRegionEuropeWest2         ZoneSpeedAPIScheduleGetResponseResultRegion = "europe-west2"
	ZoneSpeedAPIScheduleGetResponseResultRegionEuropeWest3         ZoneSpeedAPIScheduleGetResponseResultRegion = "europe-west3"
	ZoneSpeedAPIScheduleGetResponseResultRegionEuropeWest4         ZoneSpeedAPIScheduleGetResponseResultRegion = "europe-west4"
	ZoneSpeedAPIScheduleGetResponseResultRegionEuropeWest8         ZoneSpeedAPIScheduleGetResponseResultRegion = "europe-west8"
	ZoneSpeedAPIScheduleGetResponseResultRegionEuropeWest9         ZoneSpeedAPIScheduleGetResponseResultRegion = "europe-west9"
	ZoneSpeedAPIScheduleGetResponseResultRegionMeWest1             ZoneSpeedAPIScheduleGetResponseResultRegion = "me-west1"
	ZoneSpeedAPIScheduleGetResponseResultRegionSouthamericaEast1   ZoneSpeedAPIScheduleGetResponseResultRegion = "southamerica-east1"
	ZoneSpeedAPIScheduleGetResponseResultRegionUsCentral1          ZoneSpeedAPIScheduleGetResponseResultRegion = "us-central1"
	ZoneSpeedAPIScheduleGetResponseResultRegionUsEast1             ZoneSpeedAPIScheduleGetResponseResultRegion = "us-east1"
	ZoneSpeedAPIScheduleGetResponseResultRegionUsEast4             ZoneSpeedAPIScheduleGetResponseResultRegion = "us-east4"
	ZoneSpeedAPIScheduleGetResponseResultRegionUsSouth1            ZoneSpeedAPIScheduleGetResponseResultRegion = "us-south1"
	ZoneSpeedAPIScheduleGetResponseResultRegionUsWest1             ZoneSpeedAPIScheduleGetResponseResultRegion = "us-west1"
)

type ZoneSpeedAPIScheduleDeleteResponse struct {
	Errors   []ZoneSpeedAPIScheduleDeleteResponseError   `json:"errors"`
	Messages []ZoneSpeedAPIScheduleDeleteResponseMessage `json:"messages"`
	Result   ZoneSpeedAPIScheduleDeleteResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                                   `json:"success"`
	JSON    zoneSpeedAPIScheduleDeleteResponseJSON `json:"-"`
}

// zoneSpeedAPIScheduleDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIScheduleDeleteResponse]
type zoneSpeedAPIScheduleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIScheduleDeleteResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSpeedAPIScheduleDeleteResponseErrorJSON `json:"-"`
}

// zoneSpeedAPIScheduleDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIScheduleDeleteResponseError]
type zoneSpeedAPIScheduleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIScheduleDeleteResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSpeedAPIScheduleDeleteResponseMessageJSON `json:"-"`
}

// zoneSpeedAPIScheduleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIScheduleDeleteResponseMessage]
type zoneSpeedAPIScheduleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIScheduleDeleteResponseResult struct {
	// Number of items affected.
	Count float64                                      `json:"count"`
	JSON  zoneSpeedAPIScheduleDeleteResponseResultJSON `json:"-"`
}

// zoneSpeedAPIScheduleDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIScheduleDeleteResponseResult]
type zoneSpeedAPIScheduleDeleteResponseResultJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIScheduleNewParams struct {
	// A test region.
	Region param.Field[ZoneSpeedAPIScheduleNewParamsRegion] `query:"region"`
}

// URLQuery serializes [ZoneSpeedAPIScheduleNewParams]'s query parameters as
// `url.Values`.
func (r ZoneSpeedAPIScheduleNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type ZoneSpeedAPIScheduleNewParamsRegion string

const (
	ZoneSpeedAPIScheduleNewParamsRegionAsiaEast1           ZoneSpeedAPIScheduleNewParamsRegion = "asia-east1"
	ZoneSpeedAPIScheduleNewParamsRegionAsiaNortheast1      ZoneSpeedAPIScheduleNewParamsRegion = "asia-northeast1"
	ZoneSpeedAPIScheduleNewParamsRegionAsiaNortheast2      ZoneSpeedAPIScheduleNewParamsRegion = "asia-northeast2"
	ZoneSpeedAPIScheduleNewParamsRegionAsiaSouth1          ZoneSpeedAPIScheduleNewParamsRegion = "asia-south1"
	ZoneSpeedAPIScheduleNewParamsRegionAsiaSoutheast1      ZoneSpeedAPIScheduleNewParamsRegion = "asia-southeast1"
	ZoneSpeedAPIScheduleNewParamsRegionAustraliaSoutheast1 ZoneSpeedAPIScheduleNewParamsRegion = "australia-southeast1"
	ZoneSpeedAPIScheduleNewParamsRegionEuropeNorth1        ZoneSpeedAPIScheduleNewParamsRegion = "europe-north1"
	ZoneSpeedAPIScheduleNewParamsRegionEuropeSouthwest1    ZoneSpeedAPIScheduleNewParamsRegion = "europe-southwest1"
	ZoneSpeedAPIScheduleNewParamsRegionEuropeWest1         ZoneSpeedAPIScheduleNewParamsRegion = "europe-west1"
	ZoneSpeedAPIScheduleNewParamsRegionEuropeWest2         ZoneSpeedAPIScheduleNewParamsRegion = "europe-west2"
	ZoneSpeedAPIScheduleNewParamsRegionEuropeWest3         ZoneSpeedAPIScheduleNewParamsRegion = "europe-west3"
	ZoneSpeedAPIScheduleNewParamsRegionEuropeWest4         ZoneSpeedAPIScheduleNewParamsRegion = "europe-west4"
	ZoneSpeedAPIScheduleNewParamsRegionEuropeWest8         ZoneSpeedAPIScheduleNewParamsRegion = "europe-west8"
	ZoneSpeedAPIScheduleNewParamsRegionEuropeWest9         ZoneSpeedAPIScheduleNewParamsRegion = "europe-west9"
	ZoneSpeedAPIScheduleNewParamsRegionMeWest1             ZoneSpeedAPIScheduleNewParamsRegion = "me-west1"
	ZoneSpeedAPIScheduleNewParamsRegionSouthamericaEast1   ZoneSpeedAPIScheduleNewParamsRegion = "southamerica-east1"
	ZoneSpeedAPIScheduleNewParamsRegionUsCentral1          ZoneSpeedAPIScheduleNewParamsRegion = "us-central1"
	ZoneSpeedAPIScheduleNewParamsRegionUsEast1             ZoneSpeedAPIScheduleNewParamsRegion = "us-east1"
	ZoneSpeedAPIScheduleNewParamsRegionUsEast4             ZoneSpeedAPIScheduleNewParamsRegion = "us-east4"
	ZoneSpeedAPIScheduleNewParamsRegionUsSouth1            ZoneSpeedAPIScheduleNewParamsRegion = "us-south1"
	ZoneSpeedAPIScheduleNewParamsRegionUsWest1             ZoneSpeedAPIScheduleNewParamsRegion = "us-west1"
)

type ZoneSpeedAPIScheduleGetParams struct {
	// A test region.
	Region param.Field[ZoneSpeedAPIScheduleGetParamsRegion] `query:"region"`
}

// URLQuery serializes [ZoneSpeedAPIScheduleGetParams]'s query parameters as
// `url.Values`.
func (r ZoneSpeedAPIScheduleGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type ZoneSpeedAPIScheduleGetParamsRegion string

const (
	ZoneSpeedAPIScheduleGetParamsRegionAsiaEast1           ZoneSpeedAPIScheduleGetParamsRegion = "asia-east1"
	ZoneSpeedAPIScheduleGetParamsRegionAsiaNortheast1      ZoneSpeedAPIScheduleGetParamsRegion = "asia-northeast1"
	ZoneSpeedAPIScheduleGetParamsRegionAsiaNortheast2      ZoneSpeedAPIScheduleGetParamsRegion = "asia-northeast2"
	ZoneSpeedAPIScheduleGetParamsRegionAsiaSouth1          ZoneSpeedAPIScheduleGetParamsRegion = "asia-south1"
	ZoneSpeedAPIScheduleGetParamsRegionAsiaSoutheast1      ZoneSpeedAPIScheduleGetParamsRegion = "asia-southeast1"
	ZoneSpeedAPIScheduleGetParamsRegionAustraliaSoutheast1 ZoneSpeedAPIScheduleGetParamsRegion = "australia-southeast1"
	ZoneSpeedAPIScheduleGetParamsRegionEuropeNorth1        ZoneSpeedAPIScheduleGetParamsRegion = "europe-north1"
	ZoneSpeedAPIScheduleGetParamsRegionEuropeSouthwest1    ZoneSpeedAPIScheduleGetParamsRegion = "europe-southwest1"
	ZoneSpeedAPIScheduleGetParamsRegionEuropeWest1         ZoneSpeedAPIScheduleGetParamsRegion = "europe-west1"
	ZoneSpeedAPIScheduleGetParamsRegionEuropeWest2         ZoneSpeedAPIScheduleGetParamsRegion = "europe-west2"
	ZoneSpeedAPIScheduleGetParamsRegionEuropeWest3         ZoneSpeedAPIScheduleGetParamsRegion = "europe-west3"
	ZoneSpeedAPIScheduleGetParamsRegionEuropeWest4         ZoneSpeedAPIScheduleGetParamsRegion = "europe-west4"
	ZoneSpeedAPIScheduleGetParamsRegionEuropeWest8         ZoneSpeedAPIScheduleGetParamsRegion = "europe-west8"
	ZoneSpeedAPIScheduleGetParamsRegionEuropeWest9         ZoneSpeedAPIScheduleGetParamsRegion = "europe-west9"
	ZoneSpeedAPIScheduleGetParamsRegionMeWest1             ZoneSpeedAPIScheduleGetParamsRegion = "me-west1"
	ZoneSpeedAPIScheduleGetParamsRegionSouthamericaEast1   ZoneSpeedAPIScheduleGetParamsRegion = "southamerica-east1"
	ZoneSpeedAPIScheduleGetParamsRegionUsCentral1          ZoneSpeedAPIScheduleGetParamsRegion = "us-central1"
	ZoneSpeedAPIScheduleGetParamsRegionUsEast1             ZoneSpeedAPIScheduleGetParamsRegion = "us-east1"
	ZoneSpeedAPIScheduleGetParamsRegionUsEast4             ZoneSpeedAPIScheduleGetParamsRegion = "us-east4"
	ZoneSpeedAPIScheduleGetParamsRegionUsSouth1            ZoneSpeedAPIScheduleGetParamsRegion = "us-south1"
	ZoneSpeedAPIScheduleGetParamsRegionUsWest1             ZoneSpeedAPIScheduleGetParamsRegion = "us-west1"
)

type ZoneSpeedAPIScheduleDeleteParams struct {
	// A test region.
	Region param.Field[ZoneSpeedAPIScheduleDeleteParamsRegion] `query:"region"`
}

// URLQuery serializes [ZoneSpeedAPIScheduleDeleteParams]'s query parameters as
// `url.Values`.
func (r ZoneSpeedAPIScheduleDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type ZoneSpeedAPIScheduleDeleteParamsRegion string

const (
	ZoneSpeedAPIScheduleDeleteParamsRegionAsiaEast1           ZoneSpeedAPIScheduleDeleteParamsRegion = "asia-east1"
	ZoneSpeedAPIScheduleDeleteParamsRegionAsiaNortheast1      ZoneSpeedAPIScheduleDeleteParamsRegion = "asia-northeast1"
	ZoneSpeedAPIScheduleDeleteParamsRegionAsiaNortheast2      ZoneSpeedAPIScheduleDeleteParamsRegion = "asia-northeast2"
	ZoneSpeedAPIScheduleDeleteParamsRegionAsiaSouth1          ZoneSpeedAPIScheduleDeleteParamsRegion = "asia-south1"
	ZoneSpeedAPIScheduleDeleteParamsRegionAsiaSoutheast1      ZoneSpeedAPIScheduleDeleteParamsRegion = "asia-southeast1"
	ZoneSpeedAPIScheduleDeleteParamsRegionAustraliaSoutheast1 ZoneSpeedAPIScheduleDeleteParamsRegion = "australia-southeast1"
	ZoneSpeedAPIScheduleDeleteParamsRegionEuropeNorth1        ZoneSpeedAPIScheduleDeleteParamsRegion = "europe-north1"
	ZoneSpeedAPIScheduleDeleteParamsRegionEuropeSouthwest1    ZoneSpeedAPIScheduleDeleteParamsRegion = "europe-southwest1"
	ZoneSpeedAPIScheduleDeleteParamsRegionEuropeWest1         ZoneSpeedAPIScheduleDeleteParamsRegion = "europe-west1"
	ZoneSpeedAPIScheduleDeleteParamsRegionEuropeWest2         ZoneSpeedAPIScheduleDeleteParamsRegion = "europe-west2"
	ZoneSpeedAPIScheduleDeleteParamsRegionEuropeWest3         ZoneSpeedAPIScheduleDeleteParamsRegion = "europe-west3"
	ZoneSpeedAPIScheduleDeleteParamsRegionEuropeWest4         ZoneSpeedAPIScheduleDeleteParamsRegion = "europe-west4"
	ZoneSpeedAPIScheduleDeleteParamsRegionEuropeWest8         ZoneSpeedAPIScheduleDeleteParamsRegion = "europe-west8"
	ZoneSpeedAPIScheduleDeleteParamsRegionEuropeWest9         ZoneSpeedAPIScheduleDeleteParamsRegion = "europe-west9"
	ZoneSpeedAPIScheduleDeleteParamsRegionMeWest1             ZoneSpeedAPIScheduleDeleteParamsRegion = "me-west1"
	ZoneSpeedAPIScheduleDeleteParamsRegionSouthamericaEast1   ZoneSpeedAPIScheduleDeleteParamsRegion = "southamerica-east1"
	ZoneSpeedAPIScheduleDeleteParamsRegionUsCentral1          ZoneSpeedAPIScheduleDeleteParamsRegion = "us-central1"
	ZoneSpeedAPIScheduleDeleteParamsRegionUsEast1             ZoneSpeedAPIScheduleDeleteParamsRegion = "us-east1"
	ZoneSpeedAPIScheduleDeleteParamsRegionUsEast4             ZoneSpeedAPIScheduleDeleteParamsRegion = "us-east4"
	ZoneSpeedAPIScheduleDeleteParamsRegionUsSouth1            ZoneSpeedAPIScheduleDeleteParamsRegion = "us-south1"
	ZoneSpeedAPIScheduleDeleteParamsRegionUsWest1             ZoneSpeedAPIScheduleDeleteParamsRegion = "us-west1"
)
