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

// ZoneSpeedAPIPageTestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSpeedAPIPageTestService]
// method instead.
type ZoneSpeedAPIPageTestService struct {
	Options []option.RequestOption
}

// NewZoneSpeedAPIPageTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSpeedAPIPageTestService(opts ...option.RequestOption) (r *ZoneSpeedAPIPageTestService) {
	r = &ZoneSpeedAPIPageTestService{}
	r.Options = opts
	return
}

// Starts a test for a specific webpage, in a specific region.
func (r *ZoneSpeedAPIPageTestService) New(ctx context.Context, zoneIdentifier string, url string, body ZoneSpeedAPIPageTestNewParams, opts ...option.RequestOption) (res *ZoneSpeedAPIPageTestNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneIdentifier, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the result of a specific test.
func (r *ZoneSpeedAPIPageTestService) Get(ctx context.Context, zoneIdentifier string, url string, testIdentifier string, opts ...option.RequestOption) (res *ZoneSpeedAPIPageTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests/%s", zoneIdentifier, url, testIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Test history (list of tests) for a specific webpage.
func (r *ZoneSpeedAPIPageTestService) List(ctx context.Context, zoneIdentifier string, url string, query ZoneSpeedAPIPageTestListParams, opts ...option.RequestOption) (res *ZoneSpeedAPIPageTestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneIdentifier, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes all tests for a specific webpage from a specific region. Deleted tests
// are still counted as part of the quota.
func (r *ZoneSpeedAPIPageTestService) Delete(ctx context.Context, zoneIdentifier string, url string, body ZoneSpeedAPIPageTestDeleteParams, opts ...option.RequestOption) (res *ZoneSpeedAPIPageTestDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneIdentifier, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type ZoneSpeedAPIPageTestNewResponse struct {
	Errors   []ZoneSpeedAPIPageTestNewResponseError   `json:"errors"`
	Messages []ZoneSpeedAPIPageTestNewResponseMessage `json:"messages"`
	Result   ZoneSpeedAPIPageTestNewResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                                `json:"success"`
	JSON    zoneSpeedAPIPageTestNewResponseJSON `json:"-"`
}

// zoneSpeedAPIPageTestNewResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIPageTestNewResponse]
type zoneSpeedAPIPageTestNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestNewResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSpeedAPIPageTestNewResponseErrorJSON `json:"-"`
}

// zoneSpeedAPIPageTestNewResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestNewResponseError]
type zoneSpeedAPIPageTestNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestNewResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSpeedAPIPageTestNewResponseMessageJSON `json:"-"`
}

// zoneSpeedAPIPageTestNewResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestNewResponseMessage]
type zoneSpeedAPIPageTestNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestNewResponseResult struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport ZoneSpeedAPIPageTestNewResponseResultDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport ZoneSpeedAPIPageTestNewResponseResultMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region ZoneSpeedAPIPageTestNewResponseResultRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency ZoneSpeedAPIPageTestNewResponseResultScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                                    `json:"url"`
	JSON zoneSpeedAPIPageTestNewResponseResultJSON `json:"-"`
}

// zoneSpeedAPIPageTestNewResponseResultJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestNewResponseResult]
type zoneSpeedAPIPageTestNewResponseResultJSON struct {
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

func (r *ZoneSpeedAPIPageTestNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type ZoneSpeedAPIPageTestNewResponseResultDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ZoneSpeedAPIPageTestNewResponseResultDesktopReportDeviceType `json:"deviceType"`
	Error      ZoneSpeedAPIPageTestNewResponseResultDesktopReportError      `json:"error"`
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
	State ZoneSpeedAPIPageTestNewResponseResultDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                                `json:"tti"`
	JSON zoneSpeedAPIPageTestNewResponseResultDesktopReportJSON `json:"-"`
}

// zoneSpeedAPIPageTestNewResponseResultDesktopReportJSON contains the JSON
// metadata for the struct [ZoneSpeedAPIPageTestNewResponseResultDesktopReport]
type zoneSpeedAPIPageTestNewResponseResultDesktopReportJSON struct {
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

func (r *ZoneSpeedAPIPageTestNewResponseResultDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type ZoneSpeedAPIPageTestNewResponseResultDesktopReportDeviceType string

const (
	ZoneSpeedAPIPageTestNewResponseResultDesktopReportDeviceTypeDesktop ZoneSpeedAPIPageTestNewResponseResultDesktopReportDeviceType = "DESKTOP"
	ZoneSpeedAPIPageTestNewResponseResultDesktopReportDeviceTypeMobile  ZoneSpeedAPIPageTestNewResponseResultDesktopReportDeviceType = "MOBILE"
)

type ZoneSpeedAPIPageTestNewResponseResultDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code ZoneSpeedAPIPageTestNewResponseResultDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                      `json:"finalDisplayedUrl"`
	JSON              zoneSpeedAPIPageTestNewResponseResultDesktopReportErrorJSON `json:"-"`
}

// zoneSpeedAPIPageTestNewResponseResultDesktopReportErrorJSON contains the JSON
// metadata for the struct
// [ZoneSpeedAPIPageTestNewResponseResultDesktopReportError]
type zoneSpeedAPIPageTestNewResponseResultDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestNewResponseResultDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type ZoneSpeedAPIPageTestNewResponseResultDesktopReportErrorCode string

const (
	ZoneSpeedAPIPageTestNewResponseResultDesktopReportErrorCodeNotReachable      ZoneSpeedAPIPageTestNewResponseResultDesktopReportErrorCode = "NOT_REACHABLE"
	ZoneSpeedAPIPageTestNewResponseResultDesktopReportErrorCodeDNSFailure        ZoneSpeedAPIPageTestNewResponseResultDesktopReportErrorCode = "DNS_FAILURE"
	ZoneSpeedAPIPageTestNewResponseResultDesktopReportErrorCodeNotHTML           ZoneSpeedAPIPageTestNewResponseResultDesktopReportErrorCode = "NOT_HTML"
	ZoneSpeedAPIPageTestNewResponseResultDesktopReportErrorCodeLighthouseTimeout ZoneSpeedAPIPageTestNewResponseResultDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ZoneSpeedAPIPageTestNewResponseResultDesktopReportErrorCodeUnknown           ZoneSpeedAPIPageTestNewResponseResultDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ZoneSpeedAPIPageTestNewResponseResultDesktopReportState string

const (
	ZoneSpeedAPIPageTestNewResponseResultDesktopReportStateRunning  ZoneSpeedAPIPageTestNewResponseResultDesktopReportState = "RUNNING"
	ZoneSpeedAPIPageTestNewResponseResultDesktopReportStateComplete ZoneSpeedAPIPageTestNewResponseResultDesktopReportState = "COMPLETE"
	ZoneSpeedAPIPageTestNewResponseResultDesktopReportStateFailed   ZoneSpeedAPIPageTestNewResponseResultDesktopReportState = "FAILED"
)

// The Lighthouse report.
type ZoneSpeedAPIPageTestNewResponseResultMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ZoneSpeedAPIPageTestNewResponseResultMobileReportDeviceType `json:"deviceType"`
	Error      ZoneSpeedAPIPageTestNewResponseResultMobileReportError      `json:"error"`
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
	State ZoneSpeedAPIPageTestNewResponseResultMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                               `json:"tti"`
	JSON zoneSpeedAPIPageTestNewResponseResultMobileReportJSON `json:"-"`
}

// zoneSpeedAPIPageTestNewResponseResultMobileReportJSON contains the JSON metadata
// for the struct [ZoneSpeedAPIPageTestNewResponseResultMobileReport]
type zoneSpeedAPIPageTestNewResponseResultMobileReportJSON struct {
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

func (r *ZoneSpeedAPIPageTestNewResponseResultMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type ZoneSpeedAPIPageTestNewResponseResultMobileReportDeviceType string

const (
	ZoneSpeedAPIPageTestNewResponseResultMobileReportDeviceTypeDesktop ZoneSpeedAPIPageTestNewResponseResultMobileReportDeviceType = "DESKTOP"
	ZoneSpeedAPIPageTestNewResponseResultMobileReportDeviceTypeMobile  ZoneSpeedAPIPageTestNewResponseResultMobileReportDeviceType = "MOBILE"
)

type ZoneSpeedAPIPageTestNewResponseResultMobileReportError struct {
	// The error code of the Lighthouse result.
	Code ZoneSpeedAPIPageTestNewResponseResultMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                     `json:"finalDisplayedUrl"`
	JSON              zoneSpeedAPIPageTestNewResponseResultMobileReportErrorJSON `json:"-"`
}

// zoneSpeedAPIPageTestNewResponseResultMobileReportErrorJSON contains the JSON
// metadata for the struct [ZoneSpeedAPIPageTestNewResponseResultMobileReportError]
type zoneSpeedAPIPageTestNewResponseResultMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestNewResponseResultMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type ZoneSpeedAPIPageTestNewResponseResultMobileReportErrorCode string

const (
	ZoneSpeedAPIPageTestNewResponseResultMobileReportErrorCodeNotReachable      ZoneSpeedAPIPageTestNewResponseResultMobileReportErrorCode = "NOT_REACHABLE"
	ZoneSpeedAPIPageTestNewResponseResultMobileReportErrorCodeDNSFailure        ZoneSpeedAPIPageTestNewResponseResultMobileReportErrorCode = "DNS_FAILURE"
	ZoneSpeedAPIPageTestNewResponseResultMobileReportErrorCodeNotHTML           ZoneSpeedAPIPageTestNewResponseResultMobileReportErrorCode = "NOT_HTML"
	ZoneSpeedAPIPageTestNewResponseResultMobileReportErrorCodeLighthouseTimeout ZoneSpeedAPIPageTestNewResponseResultMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ZoneSpeedAPIPageTestNewResponseResultMobileReportErrorCodeUnknown           ZoneSpeedAPIPageTestNewResponseResultMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ZoneSpeedAPIPageTestNewResponseResultMobileReportState string

const (
	ZoneSpeedAPIPageTestNewResponseResultMobileReportStateRunning  ZoneSpeedAPIPageTestNewResponseResultMobileReportState = "RUNNING"
	ZoneSpeedAPIPageTestNewResponseResultMobileReportStateComplete ZoneSpeedAPIPageTestNewResponseResultMobileReportState = "COMPLETE"
	ZoneSpeedAPIPageTestNewResponseResultMobileReportStateFailed   ZoneSpeedAPIPageTestNewResponseResultMobileReportState = "FAILED"
)

// A test region with a label.
type ZoneSpeedAPIPageTestNewResponseResultRegion struct {
	Label string `json:"label"`
	// A test region.
	Value ZoneSpeedAPIPageTestNewResponseResultRegionValue `json:"value"`
	JSON  zoneSpeedAPIPageTestNewResponseResultRegionJSON  `json:"-"`
}

// zoneSpeedAPIPageTestNewResponseResultRegionJSON contains the JSON metadata for
// the struct [ZoneSpeedAPIPageTestNewResponseResultRegion]
type zoneSpeedAPIPageTestNewResponseResultRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestNewResponseResultRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type ZoneSpeedAPIPageTestNewResponseResultRegionValue string

const (
	ZoneSpeedAPIPageTestNewResponseResultRegionValueAsiaEast1           ZoneSpeedAPIPageTestNewResponseResultRegionValue = "asia-east1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueAsiaNortheast1      ZoneSpeedAPIPageTestNewResponseResultRegionValue = "asia-northeast1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueAsiaNortheast2      ZoneSpeedAPIPageTestNewResponseResultRegionValue = "asia-northeast2"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueAsiaSouth1          ZoneSpeedAPIPageTestNewResponseResultRegionValue = "asia-south1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueAsiaSoutheast1      ZoneSpeedAPIPageTestNewResponseResultRegionValue = "asia-southeast1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueAustraliaSoutheast1 ZoneSpeedAPIPageTestNewResponseResultRegionValue = "australia-southeast1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueEuropeNorth1        ZoneSpeedAPIPageTestNewResponseResultRegionValue = "europe-north1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueEuropeSouthwest1    ZoneSpeedAPIPageTestNewResponseResultRegionValue = "europe-southwest1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueEuropeWest1         ZoneSpeedAPIPageTestNewResponseResultRegionValue = "europe-west1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueEuropeWest2         ZoneSpeedAPIPageTestNewResponseResultRegionValue = "europe-west2"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueEuropeWest3         ZoneSpeedAPIPageTestNewResponseResultRegionValue = "europe-west3"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueEuropeWest4         ZoneSpeedAPIPageTestNewResponseResultRegionValue = "europe-west4"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueEuropeWest8         ZoneSpeedAPIPageTestNewResponseResultRegionValue = "europe-west8"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueEuropeWest9         ZoneSpeedAPIPageTestNewResponseResultRegionValue = "europe-west9"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueMeWest1             ZoneSpeedAPIPageTestNewResponseResultRegionValue = "me-west1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueSouthamericaEast1   ZoneSpeedAPIPageTestNewResponseResultRegionValue = "southamerica-east1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueUsCentral1          ZoneSpeedAPIPageTestNewResponseResultRegionValue = "us-central1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueUsEast1             ZoneSpeedAPIPageTestNewResponseResultRegionValue = "us-east1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueUsEast4             ZoneSpeedAPIPageTestNewResponseResultRegionValue = "us-east4"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueUsSouth1            ZoneSpeedAPIPageTestNewResponseResultRegionValue = "us-south1"
	ZoneSpeedAPIPageTestNewResponseResultRegionValueUsWest1             ZoneSpeedAPIPageTestNewResponseResultRegionValue = "us-west1"
)

// The frequency of the test.
type ZoneSpeedAPIPageTestNewResponseResultScheduleFrequency string

const (
	ZoneSpeedAPIPageTestNewResponseResultScheduleFrequencyDaily  ZoneSpeedAPIPageTestNewResponseResultScheduleFrequency = "DAILY"
	ZoneSpeedAPIPageTestNewResponseResultScheduleFrequencyWeekly ZoneSpeedAPIPageTestNewResponseResultScheduleFrequency = "WEEKLY"
)

type ZoneSpeedAPIPageTestGetResponse struct {
	Errors   []ZoneSpeedAPIPageTestGetResponseError   `json:"errors"`
	Messages []ZoneSpeedAPIPageTestGetResponseMessage `json:"messages"`
	Result   ZoneSpeedAPIPageTestGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                                `json:"success"`
	JSON    zoneSpeedAPIPageTestGetResponseJSON `json:"-"`
}

// zoneSpeedAPIPageTestGetResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIPageTestGetResponse]
type zoneSpeedAPIPageTestGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestGetResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSpeedAPIPageTestGetResponseErrorJSON `json:"-"`
}

// zoneSpeedAPIPageTestGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestGetResponseError]
type zoneSpeedAPIPageTestGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestGetResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSpeedAPIPageTestGetResponseMessageJSON `json:"-"`
}

// zoneSpeedAPIPageTestGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestGetResponseMessage]
type zoneSpeedAPIPageTestGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestGetResponseResult struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport ZoneSpeedAPIPageTestGetResponseResultDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport ZoneSpeedAPIPageTestGetResponseResultMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region ZoneSpeedAPIPageTestGetResponseResultRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency ZoneSpeedAPIPageTestGetResponseResultScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                                    `json:"url"`
	JSON zoneSpeedAPIPageTestGetResponseResultJSON `json:"-"`
}

// zoneSpeedAPIPageTestGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestGetResponseResult]
type zoneSpeedAPIPageTestGetResponseResultJSON struct {
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

func (r *ZoneSpeedAPIPageTestGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type ZoneSpeedAPIPageTestGetResponseResultDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ZoneSpeedAPIPageTestGetResponseResultDesktopReportDeviceType `json:"deviceType"`
	Error      ZoneSpeedAPIPageTestGetResponseResultDesktopReportError      `json:"error"`
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
	State ZoneSpeedAPIPageTestGetResponseResultDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                                `json:"tti"`
	JSON zoneSpeedAPIPageTestGetResponseResultDesktopReportJSON `json:"-"`
}

// zoneSpeedAPIPageTestGetResponseResultDesktopReportJSON contains the JSON
// metadata for the struct [ZoneSpeedAPIPageTestGetResponseResultDesktopReport]
type zoneSpeedAPIPageTestGetResponseResultDesktopReportJSON struct {
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

func (r *ZoneSpeedAPIPageTestGetResponseResultDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type ZoneSpeedAPIPageTestGetResponseResultDesktopReportDeviceType string

const (
	ZoneSpeedAPIPageTestGetResponseResultDesktopReportDeviceTypeDesktop ZoneSpeedAPIPageTestGetResponseResultDesktopReportDeviceType = "DESKTOP"
	ZoneSpeedAPIPageTestGetResponseResultDesktopReportDeviceTypeMobile  ZoneSpeedAPIPageTestGetResponseResultDesktopReportDeviceType = "MOBILE"
)

type ZoneSpeedAPIPageTestGetResponseResultDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code ZoneSpeedAPIPageTestGetResponseResultDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                      `json:"finalDisplayedUrl"`
	JSON              zoneSpeedAPIPageTestGetResponseResultDesktopReportErrorJSON `json:"-"`
}

// zoneSpeedAPIPageTestGetResponseResultDesktopReportErrorJSON contains the JSON
// metadata for the struct
// [ZoneSpeedAPIPageTestGetResponseResultDesktopReportError]
type zoneSpeedAPIPageTestGetResponseResultDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestGetResponseResultDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type ZoneSpeedAPIPageTestGetResponseResultDesktopReportErrorCode string

const (
	ZoneSpeedAPIPageTestGetResponseResultDesktopReportErrorCodeNotReachable      ZoneSpeedAPIPageTestGetResponseResultDesktopReportErrorCode = "NOT_REACHABLE"
	ZoneSpeedAPIPageTestGetResponseResultDesktopReportErrorCodeDNSFailure        ZoneSpeedAPIPageTestGetResponseResultDesktopReportErrorCode = "DNS_FAILURE"
	ZoneSpeedAPIPageTestGetResponseResultDesktopReportErrorCodeNotHTML           ZoneSpeedAPIPageTestGetResponseResultDesktopReportErrorCode = "NOT_HTML"
	ZoneSpeedAPIPageTestGetResponseResultDesktopReportErrorCodeLighthouseTimeout ZoneSpeedAPIPageTestGetResponseResultDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ZoneSpeedAPIPageTestGetResponseResultDesktopReportErrorCodeUnknown           ZoneSpeedAPIPageTestGetResponseResultDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ZoneSpeedAPIPageTestGetResponseResultDesktopReportState string

const (
	ZoneSpeedAPIPageTestGetResponseResultDesktopReportStateRunning  ZoneSpeedAPIPageTestGetResponseResultDesktopReportState = "RUNNING"
	ZoneSpeedAPIPageTestGetResponseResultDesktopReportStateComplete ZoneSpeedAPIPageTestGetResponseResultDesktopReportState = "COMPLETE"
	ZoneSpeedAPIPageTestGetResponseResultDesktopReportStateFailed   ZoneSpeedAPIPageTestGetResponseResultDesktopReportState = "FAILED"
)

// The Lighthouse report.
type ZoneSpeedAPIPageTestGetResponseResultMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ZoneSpeedAPIPageTestGetResponseResultMobileReportDeviceType `json:"deviceType"`
	Error      ZoneSpeedAPIPageTestGetResponseResultMobileReportError      `json:"error"`
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
	State ZoneSpeedAPIPageTestGetResponseResultMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                               `json:"tti"`
	JSON zoneSpeedAPIPageTestGetResponseResultMobileReportJSON `json:"-"`
}

// zoneSpeedAPIPageTestGetResponseResultMobileReportJSON contains the JSON metadata
// for the struct [ZoneSpeedAPIPageTestGetResponseResultMobileReport]
type zoneSpeedAPIPageTestGetResponseResultMobileReportJSON struct {
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

func (r *ZoneSpeedAPIPageTestGetResponseResultMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type ZoneSpeedAPIPageTestGetResponseResultMobileReportDeviceType string

const (
	ZoneSpeedAPIPageTestGetResponseResultMobileReportDeviceTypeDesktop ZoneSpeedAPIPageTestGetResponseResultMobileReportDeviceType = "DESKTOP"
	ZoneSpeedAPIPageTestGetResponseResultMobileReportDeviceTypeMobile  ZoneSpeedAPIPageTestGetResponseResultMobileReportDeviceType = "MOBILE"
)

type ZoneSpeedAPIPageTestGetResponseResultMobileReportError struct {
	// The error code of the Lighthouse result.
	Code ZoneSpeedAPIPageTestGetResponseResultMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                     `json:"finalDisplayedUrl"`
	JSON              zoneSpeedAPIPageTestGetResponseResultMobileReportErrorJSON `json:"-"`
}

// zoneSpeedAPIPageTestGetResponseResultMobileReportErrorJSON contains the JSON
// metadata for the struct [ZoneSpeedAPIPageTestGetResponseResultMobileReportError]
type zoneSpeedAPIPageTestGetResponseResultMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestGetResponseResultMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type ZoneSpeedAPIPageTestGetResponseResultMobileReportErrorCode string

const (
	ZoneSpeedAPIPageTestGetResponseResultMobileReportErrorCodeNotReachable      ZoneSpeedAPIPageTestGetResponseResultMobileReportErrorCode = "NOT_REACHABLE"
	ZoneSpeedAPIPageTestGetResponseResultMobileReportErrorCodeDNSFailure        ZoneSpeedAPIPageTestGetResponseResultMobileReportErrorCode = "DNS_FAILURE"
	ZoneSpeedAPIPageTestGetResponseResultMobileReportErrorCodeNotHTML           ZoneSpeedAPIPageTestGetResponseResultMobileReportErrorCode = "NOT_HTML"
	ZoneSpeedAPIPageTestGetResponseResultMobileReportErrorCodeLighthouseTimeout ZoneSpeedAPIPageTestGetResponseResultMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ZoneSpeedAPIPageTestGetResponseResultMobileReportErrorCodeUnknown           ZoneSpeedAPIPageTestGetResponseResultMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ZoneSpeedAPIPageTestGetResponseResultMobileReportState string

const (
	ZoneSpeedAPIPageTestGetResponseResultMobileReportStateRunning  ZoneSpeedAPIPageTestGetResponseResultMobileReportState = "RUNNING"
	ZoneSpeedAPIPageTestGetResponseResultMobileReportStateComplete ZoneSpeedAPIPageTestGetResponseResultMobileReportState = "COMPLETE"
	ZoneSpeedAPIPageTestGetResponseResultMobileReportStateFailed   ZoneSpeedAPIPageTestGetResponseResultMobileReportState = "FAILED"
)

// A test region with a label.
type ZoneSpeedAPIPageTestGetResponseResultRegion struct {
	Label string `json:"label"`
	// A test region.
	Value ZoneSpeedAPIPageTestGetResponseResultRegionValue `json:"value"`
	JSON  zoneSpeedAPIPageTestGetResponseResultRegionJSON  `json:"-"`
}

// zoneSpeedAPIPageTestGetResponseResultRegionJSON contains the JSON metadata for
// the struct [ZoneSpeedAPIPageTestGetResponseResultRegion]
type zoneSpeedAPIPageTestGetResponseResultRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestGetResponseResultRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type ZoneSpeedAPIPageTestGetResponseResultRegionValue string

const (
	ZoneSpeedAPIPageTestGetResponseResultRegionValueAsiaEast1           ZoneSpeedAPIPageTestGetResponseResultRegionValue = "asia-east1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueAsiaNortheast1      ZoneSpeedAPIPageTestGetResponseResultRegionValue = "asia-northeast1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueAsiaNortheast2      ZoneSpeedAPIPageTestGetResponseResultRegionValue = "asia-northeast2"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueAsiaSouth1          ZoneSpeedAPIPageTestGetResponseResultRegionValue = "asia-south1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueAsiaSoutheast1      ZoneSpeedAPIPageTestGetResponseResultRegionValue = "asia-southeast1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueAustraliaSoutheast1 ZoneSpeedAPIPageTestGetResponseResultRegionValue = "australia-southeast1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueEuropeNorth1        ZoneSpeedAPIPageTestGetResponseResultRegionValue = "europe-north1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueEuropeSouthwest1    ZoneSpeedAPIPageTestGetResponseResultRegionValue = "europe-southwest1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueEuropeWest1         ZoneSpeedAPIPageTestGetResponseResultRegionValue = "europe-west1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueEuropeWest2         ZoneSpeedAPIPageTestGetResponseResultRegionValue = "europe-west2"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueEuropeWest3         ZoneSpeedAPIPageTestGetResponseResultRegionValue = "europe-west3"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueEuropeWest4         ZoneSpeedAPIPageTestGetResponseResultRegionValue = "europe-west4"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueEuropeWest8         ZoneSpeedAPIPageTestGetResponseResultRegionValue = "europe-west8"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueEuropeWest9         ZoneSpeedAPIPageTestGetResponseResultRegionValue = "europe-west9"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueMeWest1             ZoneSpeedAPIPageTestGetResponseResultRegionValue = "me-west1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueSouthamericaEast1   ZoneSpeedAPIPageTestGetResponseResultRegionValue = "southamerica-east1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueUsCentral1          ZoneSpeedAPIPageTestGetResponseResultRegionValue = "us-central1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueUsEast1             ZoneSpeedAPIPageTestGetResponseResultRegionValue = "us-east1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueUsEast4             ZoneSpeedAPIPageTestGetResponseResultRegionValue = "us-east4"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueUsSouth1            ZoneSpeedAPIPageTestGetResponseResultRegionValue = "us-south1"
	ZoneSpeedAPIPageTestGetResponseResultRegionValueUsWest1             ZoneSpeedAPIPageTestGetResponseResultRegionValue = "us-west1"
)

// The frequency of the test.
type ZoneSpeedAPIPageTestGetResponseResultScheduleFrequency string

const (
	ZoneSpeedAPIPageTestGetResponseResultScheduleFrequencyDaily  ZoneSpeedAPIPageTestGetResponseResultScheduleFrequency = "DAILY"
	ZoneSpeedAPIPageTestGetResponseResultScheduleFrequencyWeekly ZoneSpeedAPIPageTestGetResponseResultScheduleFrequency = "WEEKLY"
)

type ZoneSpeedAPIPageTestListResponse struct {
	Errors     []ZoneSpeedAPIPageTestListResponseError    `json:"errors"`
	Messages   []ZoneSpeedAPIPageTestListResponseMessage  `json:"messages"`
	Result     []ZoneSpeedAPIPageTestListResponseResult   `json:"result"`
	ResultInfo ZoneSpeedAPIPageTestListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success bool                                 `json:"success"`
	JSON    zoneSpeedAPIPageTestListResponseJSON `json:"-"`
}

// zoneSpeedAPIPageTestListResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIPageTestListResponse]
type zoneSpeedAPIPageTestListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestListResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneSpeedAPIPageTestListResponseErrorJSON `json:"-"`
}

// zoneSpeedAPIPageTestListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestListResponseError]
type zoneSpeedAPIPageTestListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestListResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSpeedAPIPageTestListResponseMessageJSON `json:"-"`
}

// zoneSpeedAPIPageTestListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestListResponseMessage]
type zoneSpeedAPIPageTestListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestListResponseResult struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport ZoneSpeedAPIPageTestListResponseResultDesktopReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport ZoneSpeedAPIPageTestListResponseResultMobileReport `json:"mobileReport"`
	// A test region with a label.
	Region ZoneSpeedAPIPageTestListResponseResultRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency ZoneSpeedAPIPageTestListResponseResultScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                                     `json:"url"`
	JSON zoneSpeedAPIPageTestListResponseResultJSON `json:"-"`
}

// zoneSpeedAPIPageTestListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestListResponseResult]
type zoneSpeedAPIPageTestListResponseResultJSON struct {
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

func (r *ZoneSpeedAPIPageTestListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Lighthouse report.
type ZoneSpeedAPIPageTestListResponseResultDesktopReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ZoneSpeedAPIPageTestListResponseResultDesktopReportDeviceType `json:"deviceType"`
	Error      ZoneSpeedAPIPageTestListResponseResultDesktopReportError      `json:"error"`
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
	State ZoneSpeedAPIPageTestListResponseResultDesktopReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                                 `json:"tti"`
	JSON zoneSpeedAPIPageTestListResponseResultDesktopReportJSON `json:"-"`
}

// zoneSpeedAPIPageTestListResponseResultDesktopReportJSON contains the JSON
// metadata for the struct [ZoneSpeedAPIPageTestListResponseResultDesktopReport]
type zoneSpeedAPIPageTestListResponseResultDesktopReportJSON struct {
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

func (r *ZoneSpeedAPIPageTestListResponseResultDesktopReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type ZoneSpeedAPIPageTestListResponseResultDesktopReportDeviceType string

const (
	ZoneSpeedAPIPageTestListResponseResultDesktopReportDeviceTypeDesktop ZoneSpeedAPIPageTestListResponseResultDesktopReportDeviceType = "DESKTOP"
	ZoneSpeedAPIPageTestListResponseResultDesktopReportDeviceTypeMobile  ZoneSpeedAPIPageTestListResponseResultDesktopReportDeviceType = "MOBILE"
)

type ZoneSpeedAPIPageTestListResponseResultDesktopReportError struct {
	// The error code of the Lighthouse result.
	Code ZoneSpeedAPIPageTestListResponseResultDesktopReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                       `json:"finalDisplayedUrl"`
	JSON              zoneSpeedAPIPageTestListResponseResultDesktopReportErrorJSON `json:"-"`
}

// zoneSpeedAPIPageTestListResponseResultDesktopReportErrorJSON contains the JSON
// metadata for the struct
// [ZoneSpeedAPIPageTestListResponseResultDesktopReportError]
type zoneSpeedAPIPageTestListResponseResultDesktopReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestListResponseResultDesktopReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type ZoneSpeedAPIPageTestListResponseResultDesktopReportErrorCode string

const (
	ZoneSpeedAPIPageTestListResponseResultDesktopReportErrorCodeNotReachable      ZoneSpeedAPIPageTestListResponseResultDesktopReportErrorCode = "NOT_REACHABLE"
	ZoneSpeedAPIPageTestListResponseResultDesktopReportErrorCodeDNSFailure        ZoneSpeedAPIPageTestListResponseResultDesktopReportErrorCode = "DNS_FAILURE"
	ZoneSpeedAPIPageTestListResponseResultDesktopReportErrorCodeNotHTML           ZoneSpeedAPIPageTestListResponseResultDesktopReportErrorCode = "NOT_HTML"
	ZoneSpeedAPIPageTestListResponseResultDesktopReportErrorCodeLighthouseTimeout ZoneSpeedAPIPageTestListResponseResultDesktopReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ZoneSpeedAPIPageTestListResponseResultDesktopReportErrorCodeUnknown           ZoneSpeedAPIPageTestListResponseResultDesktopReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ZoneSpeedAPIPageTestListResponseResultDesktopReportState string

const (
	ZoneSpeedAPIPageTestListResponseResultDesktopReportStateRunning  ZoneSpeedAPIPageTestListResponseResultDesktopReportState = "RUNNING"
	ZoneSpeedAPIPageTestListResponseResultDesktopReportStateComplete ZoneSpeedAPIPageTestListResponseResultDesktopReportState = "COMPLETE"
	ZoneSpeedAPIPageTestListResponseResultDesktopReportStateFailed   ZoneSpeedAPIPageTestListResponseResultDesktopReportState = "FAILED"
)

// The Lighthouse report.
type ZoneSpeedAPIPageTestListResponseResultMobileReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ZoneSpeedAPIPageTestListResponseResultMobileReportDeviceType `json:"deviceType"`
	Error      ZoneSpeedAPIPageTestListResponseResultMobileReportError      `json:"error"`
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
	State ZoneSpeedAPIPageTestListResponseResultMobileReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                                                `json:"tti"`
	JSON zoneSpeedAPIPageTestListResponseResultMobileReportJSON `json:"-"`
}

// zoneSpeedAPIPageTestListResponseResultMobileReportJSON contains the JSON
// metadata for the struct [ZoneSpeedAPIPageTestListResponseResultMobileReport]
type zoneSpeedAPIPageTestListResponseResultMobileReportJSON struct {
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

func (r *ZoneSpeedAPIPageTestListResponseResultMobileReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device.
type ZoneSpeedAPIPageTestListResponseResultMobileReportDeviceType string

const (
	ZoneSpeedAPIPageTestListResponseResultMobileReportDeviceTypeDesktop ZoneSpeedAPIPageTestListResponseResultMobileReportDeviceType = "DESKTOP"
	ZoneSpeedAPIPageTestListResponseResultMobileReportDeviceTypeMobile  ZoneSpeedAPIPageTestListResponseResultMobileReportDeviceType = "MOBILE"
)

type ZoneSpeedAPIPageTestListResponseResultMobileReportError struct {
	// The error code of the Lighthouse result.
	Code ZoneSpeedAPIPageTestListResponseResultMobileReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                                                      `json:"finalDisplayedUrl"`
	JSON              zoneSpeedAPIPageTestListResponseResultMobileReportErrorJSON `json:"-"`
}

// zoneSpeedAPIPageTestListResponseResultMobileReportErrorJSON contains the JSON
// metadata for the struct
// [ZoneSpeedAPIPageTestListResponseResultMobileReportError]
type zoneSpeedAPIPageTestListResponseResultMobileReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestListResponseResultMobileReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The error code of the Lighthouse result.
type ZoneSpeedAPIPageTestListResponseResultMobileReportErrorCode string

const (
	ZoneSpeedAPIPageTestListResponseResultMobileReportErrorCodeNotReachable      ZoneSpeedAPIPageTestListResponseResultMobileReportErrorCode = "NOT_REACHABLE"
	ZoneSpeedAPIPageTestListResponseResultMobileReportErrorCodeDNSFailure        ZoneSpeedAPIPageTestListResponseResultMobileReportErrorCode = "DNS_FAILURE"
	ZoneSpeedAPIPageTestListResponseResultMobileReportErrorCodeNotHTML           ZoneSpeedAPIPageTestListResponseResultMobileReportErrorCode = "NOT_HTML"
	ZoneSpeedAPIPageTestListResponseResultMobileReportErrorCodeLighthouseTimeout ZoneSpeedAPIPageTestListResponseResultMobileReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ZoneSpeedAPIPageTestListResponseResultMobileReportErrorCodeUnknown           ZoneSpeedAPIPageTestListResponseResultMobileReportErrorCode = "UNKNOWN"
)

// The state of the Lighthouse report.
type ZoneSpeedAPIPageTestListResponseResultMobileReportState string

const (
	ZoneSpeedAPIPageTestListResponseResultMobileReportStateRunning  ZoneSpeedAPIPageTestListResponseResultMobileReportState = "RUNNING"
	ZoneSpeedAPIPageTestListResponseResultMobileReportStateComplete ZoneSpeedAPIPageTestListResponseResultMobileReportState = "COMPLETE"
	ZoneSpeedAPIPageTestListResponseResultMobileReportStateFailed   ZoneSpeedAPIPageTestListResponseResultMobileReportState = "FAILED"
)

// A test region with a label.
type ZoneSpeedAPIPageTestListResponseResultRegion struct {
	Label string `json:"label"`
	// A test region.
	Value ZoneSpeedAPIPageTestListResponseResultRegionValue `json:"value"`
	JSON  zoneSpeedAPIPageTestListResponseResultRegionJSON  `json:"-"`
}

// zoneSpeedAPIPageTestListResponseResultRegionJSON contains the JSON metadata for
// the struct [ZoneSpeedAPIPageTestListResponseResultRegion]
type zoneSpeedAPIPageTestListResponseResultRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestListResponseResultRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A test region.
type ZoneSpeedAPIPageTestListResponseResultRegionValue string

const (
	ZoneSpeedAPIPageTestListResponseResultRegionValueAsiaEast1           ZoneSpeedAPIPageTestListResponseResultRegionValue = "asia-east1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueAsiaNortheast1      ZoneSpeedAPIPageTestListResponseResultRegionValue = "asia-northeast1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueAsiaNortheast2      ZoneSpeedAPIPageTestListResponseResultRegionValue = "asia-northeast2"
	ZoneSpeedAPIPageTestListResponseResultRegionValueAsiaSouth1          ZoneSpeedAPIPageTestListResponseResultRegionValue = "asia-south1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueAsiaSoutheast1      ZoneSpeedAPIPageTestListResponseResultRegionValue = "asia-southeast1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueAustraliaSoutheast1 ZoneSpeedAPIPageTestListResponseResultRegionValue = "australia-southeast1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueEuropeNorth1        ZoneSpeedAPIPageTestListResponseResultRegionValue = "europe-north1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueEuropeSouthwest1    ZoneSpeedAPIPageTestListResponseResultRegionValue = "europe-southwest1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueEuropeWest1         ZoneSpeedAPIPageTestListResponseResultRegionValue = "europe-west1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueEuropeWest2         ZoneSpeedAPIPageTestListResponseResultRegionValue = "europe-west2"
	ZoneSpeedAPIPageTestListResponseResultRegionValueEuropeWest3         ZoneSpeedAPIPageTestListResponseResultRegionValue = "europe-west3"
	ZoneSpeedAPIPageTestListResponseResultRegionValueEuropeWest4         ZoneSpeedAPIPageTestListResponseResultRegionValue = "europe-west4"
	ZoneSpeedAPIPageTestListResponseResultRegionValueEuropeWest8         ZoneSpeedAPIPageTestListResponseResultRegionValue = "europe-west8"
	ZoneSpeedAPIPageTestListResponseResultRegionValueEuropeWest9         ZoneSpeedAPIPageTestListResponseResultRegionValue = "europe-west9"
	ZoneSpeedAPIPageTestListResponseResultRegionValueMeWest1             ZoneSpeedAPIPageTestListResponseResultRegionValue = "me-west1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueSouthamericaEast1   ZoneSpeedAPIPageTestListResponseResultRegionValue = "southamerica-east1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueUsCentral1          ZoneSpeedAPIPageTestListResponseResultRegionValue = "us-central1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueUsEast1             ZoneSpeedAPIPageTestListResponseResultRegionValue = "us-east1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueUsEast4             ZoneSpeedAPIPageTestListResponseResultRegionValue = "us-east4"
	ZoneSpeedAPIPageTestListResponseResultRegionValueUsSouth1            ZoneSpeedAPIPageTestListResponseResultRegionValue = "us-south1"
	ZoneSpeedAPIPageTestListResponseResultRegionValueUsWest1             ZoneSpeedAPIPageTestListResponseResultRegionValue = "us-west1"
)

// The frequency of the test.
type ZoneSpeedAPIPageTestListResponseResultScheduleFrequency string

const (
	ZoneSpeedAPIPageTestListResponseResultScheduleFrequencyDaily  ZoneSpeedAPIPageTestListResponseResultScheduleFrequency = "DAILY"
	ZoneSpeedAPIPageTestListResponseResultScheduleFrequencyWeekly ZoneSpeedAPIPageTestListResponseResultScheduleFrequency = "WEEKLY"
)

type ZoneSpeedAPIPageTestListResponseResultInfo struct {
	Count      int64                                          `json:"count"`
	Page       int64                                          `json:"page"`
	PerPage    int64                                          `json:"per_page"`
	TotalCount int64                                          `json:"total_count"`
	JSON       zoneSpeedAPIPageTestListResponseResultInfoJSON `json:"-"`
}

// zoneSpeedAPIPageTestListResponseResultInfoJSON contains the JSON metadata for
// the struct [ZoneSpeedAPIPageTestListResponseResultInfo]
type zoneSpeedAPIPageTestListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestDeleteResponse struct {
	Errors   []ZoneSpeedAPIPageTestDeleteResponseError   `json:"errors"`
	Messages []ZoneSpeedAPIPageTestDeleteResponseMessage `json:"messages"`
	Result   ZoneSpeedAPIPageTestDeleteResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                                   `json:"success"`
	JSON    zoneSpeedAPIPageTestDeleteResponseJSON `json:"-"`
}

// zoneSpeedAPIPageTestDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIPageTestDeleteResponse]
type zoneSpeedAPIPageTestDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestDeleteResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSpeedAPIPageTestDeleteResponseErrorJSON `json:"-"`
}

// zoneSpeedAPIPageTestDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestDeleteResponseError]
type zoneSpeedAPIPageTestDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestDeleteResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSpeedAPIPageTestDeleteResponseMessageJSON `json:"-"`
}

// zoneSpeedAPIPageTestDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestDeleteResponseMessage]
type zoneSpeedAPIPageTestDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestDeleteResponseResult struct {
	// Number of items affected.
	Count float64                                      `json:"count"`
	JSON  zoneSpeedAPIPageTestDeleteResponseResultJSON `json:"-"`
}

// zoneSpeedAPIPageTestDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestDeleteResponseResult]
type zoneSpeedAPIPageTestDeleteResponseResultJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSpeedAPIPageTestNewParams struct {
	// A test region.
	Region param.Field[ZoneSpeedAPIPageTestNewParamsRegion] `json:"region"`
}

func (r ZoneSpeedAPIPageTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A test region.
type ZoneSpeedAPIPageTestNewParamsRegion string

const (
	ZoneSpeedAPIPageTestNewParamsRegionAsiaEast1           ZoneSpeedAPIPageTestNewParamsRegion = "asia-east1"
	ZoneSpeedAPIPageTestNewParamsRegionAsiaNortheast1      ZoneSpeedAPIPageTestNewParamsRegion = "asia-northeast1"
	ZoneSpeedAPIPageTestNewParamsRegionAsiaNortheast2      ZoneSpeedAPIPageTestNewParamsRegion = "asia-northeast2"
	ZoneSpeedAPIPageTestNewParamsRegionAsiaSouth1          ZoneSpeedAPIPageTestNewParamsRegion = "asia-south1"
	ZoneSpeedAPIPageTestNewParamsRegionAsiaSoutheast1      ZoneSpeedAPIPageTestNewParamsRegion = "asia-southeast1"
	ZoneSpeedAPIPageTestNewParamsRegionAustraliaSoutheast1 ZoneSpeedAPIPageTestNewParamsRegion = "australia-southeast1"
	ZoneSpeedAPIPageTestNewParamsRegionEuropeNorth1        ZoneSpeedAPIPageTestNewParamsRegion = "europe-north1"
	ZoneSpeedAPIPageTestNewParamsRegionEuropeSouthwest1    ZoneSpeedAPIPageTestNewParamsRegion = "europe-southwest1"
	ZoneSpeedAPIPageTestNewParamsRegionEuropeWest1         ZoneSpeedAPIPageTestNewParamsRegion = "europe-west1"
	ZoneSpeedAPIPageTestNewParamsRegionEuropeWest2         ZoneSpeedAPIPageTestNewParamsRegion = "europe-west2"
	ZoneSpeedAPIPageTestNewParamsRegionEuropeWest3         ZoneSpeedAPIPageTestNewParamsRegion = "europe-west3"
	ZoneSpeedAPIPageTestNewParamsRegionEuropeWest4         ZoneSpeedAPIPageTestNewParamsRegion = "europe-west4"
	ZoneSpeedAPIPageTestNewParamsRegionEuropeWest8         ZoneSpeedAPIPageTestNewParamsRegion = "europe-west8"
	ZoneSpeedAPIPageTestNewParamsRegionEuropeWest9         ZoneSpeedAPIPageTestNewParamsRegion = "europe-west9"
	ZoneSpeedAPIPageTestNewParamsRegionMeWest1             ZoneSpeedAPIPageTestNewParamsRegion = "me-west1"
	ZoneSpeedAPIPageTestNewParamsRegionSouthamericaEast1   ZoneSpeedAPIPageTestNewParamsRegion = "southamerica-east1"
	ZoneSpeedAPIPageTestNewParamsRegionUsCentral1          ZoneSpeedAPIPageTestNewParamsRegion = "us-central1"
	ZoneSpeedAPIPageTestNewParamsRegionUsEast1             ZoneSpeedAPIPageTestNewParamsRegion = "us-east1"
	ZoneSpeedAPIPageTestNewParamsRegionUsEast4             ZoneSpeedAPIPageTestNewParamsRegion = "us-east4"
	ZoneSpeedAPIPageTestNewParamsRegionUsSouth1            ZoneSpeedAPIPageTestNewParamsRegion = "us-south1"
	ZoneSpeedAPIPageTestNewParamsRegionUsWest1             ZoneSpeedAPIPageTestNewParamsRegion = "us-west1"
)

type ZoneSpeedAPIPageTestListParams struct {
	Page    param.Field[int64] `query:"page"`
	PerPage param.Field[int64] `query:"per_page"`
	// A test region.
	Region param.Field[ZoneSpeedAPIPageTestListParamsRegion] `query:"region"`
}

// URLQuery serializes [ZoneSpeedAPIPageTestListParams]'s query parameters as
// `url.Values`.
func (r ZoneSpeedAPIPageTestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type ZoneSpeedAPIPageTestListParamsRegion string

const (
	ZoneSpeedAPIPageTestListParamsRegionAsiaEast1           ZoneSpeedAPIPageTestListParamsRegion = "asia-east1"
	ZoneSpeedAPIPageTestListParamsRegionAsiaNortheast1      ZoneSpeedAPIPageTestListParamsRegion = "asia-northeast1"
	ZoneSpeedAPIPageTestListParamsRegionAsiaNortheast2      ZoneSpeedAPIPageTestListParamsRegion = "asia-northeast2"
	ZoneSpeedAPIPageTestListParamsRegionAsiaSouth1          ZoneSpeedAPIPageTestListParamsRegion = "asia-south1"
	ZoneSpeedAPIPageTestListParamsRegionAsiaSoutheast1      ZoneSpeedAPIPageTestListParamsRegion = "asia-southeast1"
	ZoneSpeedAPIPageTestListParamsRegionAustraliaSoutheast1 ZoneSpeedAPIPageTestListParamsRegion = "australia-southeast1"
	ZoneSpeedAPIPageTestListParamsRegionEuropeNorth1        ZoneSpeedAPIPageTestListParamsRegion = "europe-north1"
	ZoneSpeedAPIPageTestListParamsRegionEuropeSouthwest1    ZoneSpeedAPIPageTestListParamsRegion = "europe-southwest1"
	ZoneSpeedAPIPageTestListParamsRegionEuropeWest1         ZoneSpeedAPIPageTestListParamsRegion = "europe-west1"
	ZoneSpeedAPIPageTestListParamsRegionEuropeWest2         ZoneSpeedAPIPageTestListParamsRegion = "europe-west2"
	ZoneSpeedAPIPageTestListParamsRegionEuropeWest3         ZoneSpeedAPIPageTestListParamsRegion = "europe-west3"
	ZoneSpeedAPIPageTestListParamsRegionEuropeWest4         ZoneSpeedAPIPageTestListParamsRegion = "europe-west4"
	ZoneSpeedAPIPageTestListParamsRegionEuropeWest8         ZoneSpeedAPIPageTestListParamsRegion = "europe-west8"
	ZoneSpeedAPIPageTestListParamsRegionEuropeWest9         ZoneSpeedAPIPageTestListParamsRegion = "europe-west9"
	ZoneSpeedAPIPageTestListParamsRegionMeWest1             ZoneSpeedAPIPageTestListParamsRegion = "me-west1"
	ZoneSpeedAPIPageTestListParamsRegionSouthamericaEast1   ZoneSpeedAPIPageTestListParamsRegion = "southamerica-east1"
	ZoneSpeedAPIPageTestListParamsRegionUsCentral1          ZoneSpeedAPIPageTestListParamsRegion = "us-central1"
	ZoneSpeedAPIPageTestListParamsRegionUsEast1             ZoneSpeedAPIPageTestListParamsRegion = "us-east1"
	ZoneSpeedAPIPageTestListParamsRegionUsEast4             ZoneSpeedAPIPageTestListParamsRegion = "us-east4"
	ZoneSpeedAPIPageTestListParamsRegionUsSouth1            ZoneSpeedAPIPageTestListParamsRegion = "us-south1"
	ZoneSpeedAPIPageTestListParamsRegionUsWest1             ZoneSpeedAPIPageTestListParamsRegion = "us-west1"
)

type ZoneSpeedAPIPageTestDeleteParams struct {
	// A test region.
	Region param.Field[ZoneSpeedAPIPageTestDeleteParamsRegion] `query:"region"`
}

// URLQuery serializes [ZoneSpeedAPIPageTestDeleteParams]'s query parameters as
// `url.Values`.
func (r ZoneSpeedAPIPageTestDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type ZoneSpeedAPIPageTestDeleteParamsRegion string

const (
	ZoneSpeedAPIPageTestDeleteParamsRegionAsiaEast1           ZoneSpeedAPIPageTestDeleteParamsRegion = "asia-east1"
	ZoneSpeedAPIPageTestDeleteParamsRegionAsiaNortheast1      ZoneSpeedAPIPageTestDeleteParamsRegion = "asia-northeast1"
	ZoneSpeedAPIPageTestDeleteParamsRegionAsiaNortheast2      ZoneSpeedAPIPageTestDeleteParamsRegion = "asia-northeast2"
	ZoneSpeedAPIPageTestDeleteParamsRegionAsiaSouth1          ZoneSpeedAPIPageTestDeleteParamsRegion = "asia-south1"
	ZoneSpeedAPIPageTestDeleteParamsRegionAsiaSoutheast1      ZoneSpeedAPIPageTestDeleteParamsRegion = "asia-southeast1"
	ZoneSpeedAPIPageTestDeleteParamsRegionAustraliaSoutheast1 ZoneSpeedAPIPageTestDeleteParamsRegion = "australia-southeast1"
	ZoneSpeedAPIPageTestDeleteParamsRegionEuropeNorth1        ZoneSpeedAPIPageTestDeleteParamsRegion = "europe-north1"
	ZoneSpeedAPIPageTestDeleteParamsRegionEuropeSouthwest1    ZoneSpeedAPIPageTestDeleteParamsRegion = "europe-southwest1"
	ZoneSpeedAPIPageTestDeleteParamsRegionEuropeWest1         ZoneSpeedAPIPageTestDeleteParamsRegion = "europe-west1"
	ZoneSpeedAPIPageTestDeleteParamsRegionEuropeWest2         ZoneSpeedAPIPageTestDeleteParamsRegion = "europe-west2"
	ZoneSpeedAPIPageTestDeleteParamsRegionEuropeWest3         ZoneSpeedAPIPageTestDeleteParamsRegion = "europe-west3"
	ZoneSpeedAPIPageTestDeleteParamsRegionEuropeWest4         ZoneSpeedAPIPageTestDeleteParamsRegion = "europe-west4"
	ZoneSpeedAPIPageTestDeleteParamsRegionEuropeWest8         ZoneSpeedAPIPageTestDeleteParamsRegion = "europe-west8"
	ZoneSpeedAPIPageTestDeleteParamsRegionEuropeWest9         ZoneSpeedAPIPageTestDeleteParamsRegion = "europe-west9"
	ZoneSpeedAPIPageTestDeleteParamsRegionMeWest1             ZoneSpeedAPIPageTestDeleteParamsRegion = "me-west1"
	ZoneSpeedAPIPageTestDeleteParamsRegionSouthamericaEast1   ZoneSpeedAPIPageTestDeleteParamsRegion = "southamerica-east1"
	ZoneSpeedAPIPageTestDeleteParamsRegionUsCentral1          ZoneSpeedAPIPageTestDeleteParamsRegion = "us-central1"
	ZoneSpeedAPIPageTestDeleteParamsRegionUsEast1             ZoneSpeedAPIPageTestDeleteParamsRegion = "us-east1"
	ZoneSpeedAPIPageTestDeleteParamsRegionUsEast4             ZoneSpeedAPIPageTestDeleteParamsRegion = "us-east4"
	ZoneSpeedAPIPageTestDeleteParamsRegionUsSouth1            ZoneSpeedAPIPageTestDeleteParamsRegion = "us-south1"
	ZoneSpeedAPIPageTestDeleteParamsRegionUsWest1             ZoneSpeedAPIPageTestDeleteParamsRegion = "us-west1"
)
