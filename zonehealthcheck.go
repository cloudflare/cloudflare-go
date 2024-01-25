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

// ZoneHealthcheckService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneHealthcheckService] method
// instead.
type ZoneHealthcheckService struct {
	Options  []option.RequestOption
	Previews *ZoneHealthcheckPreviewService
}

// NewZoneHealthcheckService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneHealthcheckService(opts ...option.RequestOption) (r *ZoneHealthcheckService) {
	r = &ZoneHealthcheckService{}
	r.Options = opts
	r.Previews = NewZoneHealthcheckPreviewService(opts...)
	return
}

// Fetch a single configured health check.
func (r *ZoneHealthcheckService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *HealthchecksSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a configured health check.
func (r *ZoneHealthcheckService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneHealthcheckUpdateParams, opts ...option.RequestOption) (res *HealthchecksSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a health check.
func (r *ZoneHealthcheckService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneHealthcheckDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new health check.
func (r *ZoneHealthcheckService) HealthChecksNewHealthCheck(ctx context.Context, zoneIdentifier string, body ZoneHealthcheckHealthChecksNewHealthCheckParams, opts ...option.RequestOption) (res *HealthchecksSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/healthchecks", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured health checks.
func (r *ZoneHealthcheckService) HealthChecksListHealthChecks(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneHealthcheckHealthChecksListHealthChecksResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/healthchecks", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HealthchecksSingleResponse struct {
	Errors   []HealthchecksSingleResponseError   `json:"errors"`
	Messages []HealthchecksSingleResponseMessage `json:"messages"`
	Result   HealthchecksSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success HealthchecksSingleResponseSuccess `json:"success"`
	JSON    healthchecksSingleResponseJSON    `json:"-"`
}

// healthchecksSingleResponseJSON contains the JSON metadata for the struct
// [HealthchecksSingleResponse]
type healthchecksSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthchecksSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthchecksSingleResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    healthchecksSingleResponseErrorJSON `json:"-"`
}

// healthchecksSingleResponseErrorJSON contains the JSON metadata for the struct
// [HealthchecksSingleResponseError]
type healthchecksSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthchecksSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthchecksSingleResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    healthchecksSingleResponseMessageJSON `json:"-"`
}

// healthchecksSingleResponseMessageJSON contains the JSON metadata for the struct
// [HealthchecksSingleResponseMessage]
type healthchecksSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthchecksSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthchecksSingleResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthchecksSingleResponseResultCheckRegion `json:"check_regions,nullable"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails int64 `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses int64     `json:"consecutive_successes"`
	CreatedOn            time.Time `json:"created_on" format:"date-time"`
	// A human-readable description of the health check.
	Description string `json:"description"`
	// The current failure reason if status is unhealthy.
	FailureReason string `json:"failure_reason"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig HealthchecksSingleResponseResultHTTPConfig `json:"http_config,nullable"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval   int64     `json:"interval"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name string `json:"name"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The current status of the origin server according to the health check.
	Status HealthchecksSingleResponseResultStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthchecksSingleResponseResultTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                               `json:"type"`
	JSON healthchecksSingleResponseResultJSON `json:"-"`
}

// healthchecksSingleResponseResultJSON contains the JSON metadata for the struct
// [HealthchecksSingleResponseResult]
type healthchecksSingleResponseResultJSON struct {
	ID                   apijson.Field
	Address              apijson.Field
	CheckRegions         apijson.Field
	ConsecutiveFails     apijson.Field
	ConsecutiveSuccesses apijson.Field
	CreatedOn            apijson.Field
	Description          apijson.Field
	FailureReason        apijson.Field
	HTTPConfig           apijson.Field
	Interval             apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	Retries              apijson.Field
	Status               apijson.Field
	Suspended            apijson.Field
	TcpConfig            apijson.Field
	Timeout              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HealthchecksSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthchecksSingleResponseResultCheckRegion string

const (
	HealthchecksSingleResponseResultCheckRegionWnam       HealthchecksSingleResponseResultCheckRegion = "WNAM"
	HealthchecksSingleResponseResultCheckRegionEnam       HealthchecksSingleResponseResultCheckRegion = "ENAM"
	HealthchecksSingleResponseResultCheckRegionWeu        HealthchecksSingleResponseResultCheckRegion = "WEU"
	HealthchecksSingleResponseResultCheckRegionEeu        HealthchecksSingleResponseResultCheckRegion = "EEU"
	HealthchecksSingleResponseResultCheckRegionNsam       HealthchecksSingleResponseResultCheckRegion = "NSAM"
	HealthchecksSingleResponseResultCheckRegionSsam       HealthchecksSingleResponseResultCheckRegion = "SSAM"
	HealthchecksSingleResponseResultCheckRegionOc         HealthchecksSingleResponseResultCheckRegion = "OC"
	HealthchecksSingleResponseResultCheckRegionMe         HealthchecksSingleResponseResultCheckRegion = "ME"
	HealthchecksSingleResponseResultCheckRegionNaf        HealthchecksSingleResponseResultCheckRegion = "NAF"
	HealthchecksSingleResponseResultCheckRegionSaf        HealthchecksSingleResponseResultCheckRegion = "SAF"
	HealthchecksSingleResponseResultCheckRegionIn         HealthchecksSingleResponseResultCheckRegion = "IN"
	HealthchecksSingleResponseResultCheckRegionSeas       HealthchecksSingleResponseResultCheckRegion = "SEAS"
	HealthchecksSingleResponseResultCheckRegionNeas       HealthchecksSingleResponseResultCheckRegion = "NEAS"
	HealthchecksSingleResponseResultCheckRegionAllRegions HealthchecksSingleResponseResultCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthchecksSingleResponseResultHTTPConfig struct {
	// Do not validate the certificate when the health check uses HTTPS.
	AllowInsecure bool `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy.
	ExpectedBody string `json:"expected_body"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all
	// codes starting with 2) of the health check.
	ExpectedCodes []string `json:"expected_codes,nullable"`
	// Follow redirects if the origin returns a 3xx status code.
	FollowRedirects bool `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden.
	Header interface{} `json:"header,nullable"`
	// The HTTP method to use for the health check.
	Method HealthchecksSingleResponseResultHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                          `json:"port"`
	JSON healthchecksSingleResponseResultHTTPConfigJSON `json:"-"`
}

// healthchecksSingleResponseResultHTTPConfigJSON contains the JSON metadata for
// the struct [HealthchecksSingleResponseResultHTTPConfig]
type healthchecksSingleResponseResultHTTPConfigJSON struct {
	AllowInsecure   apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Method          apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *HealthchecksSingleResponseResultHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type HealthchecksSingleResponseResultHTTPConfigMethod string

const (
	HealthchecksSingleResponseResultHTTPConfigMethodGet  HealthchecksSingleResponseResultHTTPConfigMethod = "GET"
	HealthchecksSingleResponseResultHTTPConfigMethodHead HealthchecksSingleResponseResultHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type HealthchecksSingleResponseResultStatus string

const (
	HealthchecksSingleResponseResultStatusUnknown   HealthchecksSingleResponseResultStatus = "unknown"
	HealthchecksSingleResponseResultStatusHealthy   HealthchecksSingleResponseResultStatus = "healthy"
	HealthchecksSingleResponseResultStatusUnhealthy HealthchecksSingleResponseResultStatus = "unhealthy"
	HealthchecksSingleResponseResultStatusSuspended HealthchecksSingleResponseResultStatus = "suspended"
)

// Parameters specific to TCP health check.
type HealthchecksSingleResponseResultTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthchecksSingleResponseResultTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                         `json:"port"`
	JSON healthchecksSingleResponseResultTcpConfigJSON `json:"-"`
}

// healthchecksSingleResponseResultTcpConfigJSON contains the JSON metadata for the
// struct [HealthchecksSingleResponseResultTcpConfig]
type healthchecksSingleResponseResultTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthchecksSingleResponseResultTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type HealthchecksSingleResponseResultTcpConfigMethod string

const (
	HealthchecksSingleResponseResultTcpConfigMethodConnectionEstablished HealthchecksSingleResponseResultTcpConfigMethod = "connection_established"
)

// Whether the API call was successful
type HealthchecksSingleResponseSuccess bool

const (
	HealthchecksSingleResponseSuccessTrue HealthchecksSingleResponseSuccess = true
)

type ZoneHealthcheckDeleteResponse struct {
	Errors   []ZoneHealthcheckDeleteResponseError   `json:"errors"`
	Messages []ZoneHealthcheckDeleteResponseMessage `json:"messages"`
	Result   ZoneHealthcheckDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneHealthcheckDeleteResponseSuccess `json:"success"`
	JSON    zoneHealthcheckDeleteResponseJSON    `json:"-"`
}

// zoneHealthcheckDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneHealthcheckDeleteResponse]
type zoneHealthcheckDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHealthcheckDeleteResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneHealthcheckDeleteResponseErrorJSON `json:"-"`
}

// zoneHealthcheckDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ZoneHealthcheckDeleteResponseError]
type zoneHealthcheckDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHealthcheckDeleteResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneHealthcheckDeleteResponseMessageJSON `json:"-"`
}

// zoneHealthcheckDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneHealthcheckDeleteResponseMessage]
type zoneHealthcheckDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHealthcheckDeleteResponseResult struct {
	// Identifier
	ID   string                                  `json:"id"`
	JSON zoneHealthcheckDeleteResponseResultJSON `json:"-"`
}

// zoneHealthcheckDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneHealthcheckDeleteResponseResult]
type zoneHealthcheckDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneHealthcheckDeleteResponseSuccess bool

const (
	ZoneHealthcheckDeleteResponseSuccessTrue ZoneHealthcheckDeleteResponseSuccess = true
)

type ZoneHealthcheckHealthChecksListHealthChecksResponse struct {
	Errors     []ZoneHealthcheckHealthChecksListHealthChecksResponseError    `json:"errors"`
	Messages   []ZoneHealthcheckHealthChecksListHealthChecksResponseMessage  `json:"messages"`
	Result     []ZoneHealthcheckHealthChecksListHealthChecksResponseResult   `json:"result"`
	ResultInfo ZoneHealthcheckHealthChecksListHealthChecksResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneHealthcheckHealthChecksListHealthChecksResponseSuccess `json:"success"`
	JSON    zoneHealthcheckHealthChecksListHealthChecksResponseJSON    `json:"-"`
}

// zoneHealthcheckHealthChecksListHealthChecksResponseJSON contains the JSON
// metadata for the struct [ZoneHealthcheckHealthChecksListHealthChecksResponse]
type zoneHealthcheckHealthChecksListHealthChecksResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckHealthChecksListHealthChecksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHealthcheckHealthChecksListHealthChecksResponseError struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zoneHealthcheckHealthChecksListHealthChecksResponseErrorJSON `json:"-"`
}

// zoneHealthcheckHealthChecksListHealthChecksResponseErrorJSON contains the JSON
// metadata for the struct
// [ZoneHealthcheckHealthChecksListHealthChecksResponseError]
type zoneHealthcheckHealthChecksListHealthChecksResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckHealthChecksListHealthChecksResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHealthcheckHealthChecksListHealthChecksResponseMessage struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zoneHealthcheckHealthChecksListHealthChecksResponseMessageJSON `json:"-"`
}

// zoneHealthcheckHealthChecksListHealthChecksResponseMessageJSON contains the JSON
// metadata for the struct
// [ZoneHealthcheckHealthChecksListHealthChecksResponseMessage]
type zoneHealthcheckHealthChecksListHealthChecksResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckHealthChecksListHealthChecksResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHealthcheckHealthChecksListHealthChecksResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion `json:"check_regions,nullable"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails int64 `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses int64     `json:"consecutive_successes"`
	CreatedOn            time.Time `json:"created_on" format:"date-time"`
	// A human-readable description of the health check.
	Description string `json:"description"`
	// The current failure reason if status is unhealthy.
	FailureReason string `json:"failure_reason"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig ZoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfig `json:"http_config,nullable"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval   int64     `json:"interval"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name string `json:"name"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The current status of the origin server according to the health check.
	Status ZoneHealthcheckHealthChecksListHealthChecksResponseResultStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig ZoneHealthcheckHealthChecksListHealthChecksResponseResultTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                                                        `json:"type"`
	JSON zoneHealthcheckHealthChecksListHealthChecksResponseResultJSON `json:"-"`
}

// zoneHealthcheckHealthChecksListHealthChecksResponseResultJSON contains the JSON
// metadata for the struct
// [ZoneHealthcheckHealthChecksListHealthChecksResponseResult]
type zoneHealthcheckHealthChecksListHealthChecksResponseResultJSON struct {
	ID                   apijson.Field
	Address              apijson.Field
	CheckRegions         apijson.Field
	ConsecutiveFails     apijson.Field
	ConsecutiveSuccesses apijson.Field
	CreatedOn            apijson.Field
	Description          apijson.Field
	FailureReason        apijson.Field
	HTTPConfig           apijson.Field
	Interval             apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	Retries              apijson.Field
	Status               apijson.Field
	Suspended            apijson.Field
	TcpConfig            apijson.Field
	Timeout              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneHealthcheckHealthChecksListHealthChecksResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion string

const (
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionWnam       ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "WNAM"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionEnam       ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "ENAM"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionWeu        ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "WEU"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionEeu        ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "EEU"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionNsam       ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "NSAM"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionSsam       ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "SSAM"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionOc         ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "OC"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionMe         ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "ME"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionNaf        ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "NAF"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionSaf        ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "SAF"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionIn         ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "IN"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionSeas       ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "SEAS"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionNeas       ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "NEAS"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegionAllRegions ZoneHealthcheckHealthChecksListHealthChecksResponseResultCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type ZoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfig struct {
	// Do not validate the certificate when the health check uses HTTPS.
	AllowInsecure bool `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy.
	ExpectedBody string `json:"expected_body"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all
	// codes starting with 2) of the health check.
	ExpectedCodes []string `json:"expected_codes,nullable"`
	// Follow redirects if the origin returns a 3xx status code.
	FollowRedirects bool `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden.
	Header interface{} `json:"header,nullable"`
	// The HTTP method to use for the health check.
	Method ZoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                                                   `json:"port"`
	JSON zoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfigJSON `json:"-"`
}

// zoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfigJSON contains
// the JSON metadata for the struct
// [ZoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfig]
type zoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfigJSON struct {
	AllowInsecure   apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Method          apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type ZoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfigMethod string

const (
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfigMethodGet  ZoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfigMethod = "GET"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfigMethodHead ZoneHealthcheckHealthChecksListHealthChecksResponseResultHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type ZoneHealthcheckHealthChecksListHealthChecksResponseResultStatus string

const (
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultStatusUnknown   ZoneHealthcheckHealthChecksListHealthChecksResponseResultStatus = "unknown"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultStatusHealthy   ZoneHealthcheckHealthChecksListHealthChecksResponseResultStatus = "healthy"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultStatusUnhealthy ZoneHealthcheckHealthChecksListHealthChecksResponseResultStatus = "unhealthy"
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultStatusSuspended ZoneHealthcheckHealthChecksListHealthChecksResponseResultStatus = "suspended"
)

// Parameters specific to TCP health check.
type ZoneHealthcheckHealthChecksListHealthChecksResponseResultTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method ZoneHealthcheckHealthChecksListHealthChecksResponseResultTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                                                  `json:"port"`
	JSON zoneHealthcheckHealthChecksListHealthChecksResponseResultTcpConfigJSON `json:"-"`
}

// zoneHealthcheckHealthChecksListHealthChecksResponseResultTcpConfigJSON contains
// the JSON metadata for the struct
// [ZoneHealthcheckHealthChecksListHealthChecksResponseResultTcpConfig]
type zoneHealthcheckHealthChecksListHealthChecksResponseResultTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckHealthChecksListHealthChecksResponseResultTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type ZoneHealthcheckHealthChecksListHealthChecksResponseResultTcpConfigMethod string

const (
	ZoneHealthcheckHealthChecksListHealthChecksResponseResultTcpConfigMethodConnectionEstablished ZoneHealthcheckHealthChecksListHealthChecksResponseResultTcpConfigMethod = "connection_established"
)

type ZoneHealthcheckHealthChecksListHealthChecksResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                           `json:"total_count"`
	JSON       zoneHealthcheckHealthChecksListHealthChecksResponseResultInfoJSON `json:"-"`
}

// zoneHealthcheckHealthChecksListHealthChecksResponseResultInfoJSON contains the
// JSON metadata for the struct
// [ZoneHealthcheckHealthChecksListHealthChecksResponseResultInfo]
type zoneHealthcheckHealthChecksListHealthChecksResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckHealthChecksListHealthChecksResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneHealthcheckHealthChecksListHealthChecksResponseSuccess bool

const (
	ZoneHealthcheckHealthChecksListHealthChecksResponseSuccessTrue ZoneHealthcheckHealthChecksListHealthChecksResponseSuccess = true
)

type ZoneHealthcheckUpdateParams struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]ZoneHealthcheckUpdateParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[ZoneHealthcheckUpdateParamsHTTPConfig] `json:"http_config"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// If suspended, no health checks are sent to the origin.
	Suspended param.Field[bool] `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig param.Field[ZoneHealthcheckUpdateParamsTcpConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r ZoneHealthcheckUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type ZoneHealthcheckUpdateParamsCheckRegion string

const (
	ZoneHealthcheckUpdateParamsCheckRegionWnam       ZoneHealthcheckUpdateParamsCheckRegion = "WNAM"
	ZoneHealthcheckUpdateParamsCheckRegionEnam       ZoneHealthcheckUpdateParamsCheckRegion = "ENAM"
	ZoneHealthcheckUpdateParamsCheckRegionWeu        ZoneHealthcheckUpdateParamsCheckRegion = "WEU"
	ZoneHealthcheckUpdateParamsCheckRegionEeu        ZoneHealthcheckUpdateParamsCheckRegion = "EEU"
	ZoneHealthcheckUpdateParamsCheckRegionNsam       ZoneHealthcheckUpdateParamsCheckRegion = "NSAM"
	ZoneHealthcheckUpdateParamsCheckRegionSsam       ZoneHealthcheckUpdateParamsCheckRegion = "SSAM"
	ZoneHealthcheckUpdateParamsCheckRegionOc         ZoneHealthcheckUpdateParamsCheckRegion = "OC"
	ZoneHealthcheckUpdateParamsCheckRegionMe         ZoneHealthcheckUpdateParamsCheckRegion = "ME"
	ZoneHealthcheckUpdateParamsCheckRegionNaf        ZoneHealthcheckUpdateParamsCheckRegion = "NAF"
	ZoneHealthcheckUpdateParamsCheckRegionSaf        ZoneHealthcheckUpdateParamsCheckRegion = "SAF"
	ZoneHealthcheckUpdateParamsCheckRegionIn         ZoneHealthcheckUpdateParamsCheckRegion = "IN"
	ZoneHealthcheckUpdateParamsCheckRegionSeas       ZoneHealthcheckUpdateParamsCheckRegion = "SEAS"
	ZoneHealthcheckUpdateParamsCheckRegionNeas       ZoneHealthcheckUpdateParamsCheckRegion = "NEAS"
	ZoneHealthcheckUpdateParamsCheckRegionAllRegions ZoneHealthcheckUpdateParamsCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type ZoneHealthcheckUpdateParamsHTTPConfig struct {
	// Do not validate the certificate when the health check uses HTTPS.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all
	// codes starting with 2) of the health check.
	ExpectedCodes param.Field[[]string] `json:"expected_codes"`
	// Follow redirects if the origin returns a 3xx status code.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden.
	Header param.Field[interface{}] `json:"header"`
	// The HTTP method to use for the health check.
	Method param.Field[ZoneHealthcheckUpdateParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r ZoneHealthcheckUpdateParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type ZoneHealthcheckUpdateParamsHTTPConfigMethod string

const (
	ZoneHealthcheckUpdateParamsHTTPConfigMethodGet  ZoneHealthcheckUpdateParamsHTTPConfigMethod = "GET"
	ZoneHealthcheckUpdateParamsHTTPConfigMethodHead ZoneHealthcheckUpdateParamsHTTPConfigMethod = "HEAD"
)

// Parameters specific to TCP health check.
type ZoneHealthcheckUpdateParamsTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[ZoneHealthcheckUpdateParamsTcpConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r ZoneHealthcheckUpdateParamsTcpConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type ZoneHealthcheckUpdateParamsTcpConfigMethod string

const (
	ZoneHealthcheckUpdateParamsTcpConfigMethodConnectionEstablished ZoneHealthcheckUpdateParamsTcpConfigMethod = "connection_established"
)

type ZoneHealthcheckHealthChecksNewHealthCheckParams struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[ZoneHealthcheckHealthChecksNewHealthCheckParamsHTTPConfig] `json:"http_config"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// If suspended, no health checks are sent to the origin.
	Suspended param.Field[bool] `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig param.Field[ZoneHealthcheckHealthChecksNewHealthCheckParamsTcpConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r ZoneHealthcheckHealthChecksNewHealthCheckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion string

const (
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionWnam       ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "WNAM"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionEnam       ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "ENAM"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionWeu        ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "WEU"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionEeu        ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "EEU"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionNsam       ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "NSAM"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionSsam       ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "SSAM"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionOc         ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "OC"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionMe         ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "ME"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionNaf        ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "NAF"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionSaf        ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "SAF"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionIn         ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "IN"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionSeas       ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "SEAS"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionNeas       ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "NEAS"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegionAllRegions ZoneHealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type ZoneHealthcheckHealthChecksNewHealthCheckParamsHTTPConfig struct {
	// Do not validate the certificate when the health check uses HTTPS.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all
	// codes starting with 2) of the health check.
	ExpectedCodes param.Field[[]string] `json:"expected_codes"`
	// Follow redirects if the origin returns a 3xx status code.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden.
	Header param.Field[interface{}] `json:"header"`
	// The HTTP method to use for the health check.
	Method param.Field[ZoneHealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r ZoneHealthcheckHealthChecksNewHealthCheckParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type ZoneHealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethod string

const (
	ZoneHealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethodGet  ZoneHealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethod = "GET"
	ZoneHealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethodHead ZoneHealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethod = "HEAD"
)

// Parameters specific to TCP health check.
type ZoneHealthcheckHealthChecksNewHealthCheckParamsTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[ZoneHealthcheckHealthChecksNewHealthCheckParamsTcpConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r ZoneHealthcheckHealthChecksNewHealthCheckParamsTcpConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type ZoneHealthcheckHealthChecksNewHealthCheckParamsTcpConfigMethod string

const (
	ZoneHealthcheckHealthChecksNewHealthCheckParamsTcpConfigMethodConnectionEstablished ZoneHealthcheckHealthChecksNewHealthCheckParamsTcpConfigMethod = "connection_established"
)
