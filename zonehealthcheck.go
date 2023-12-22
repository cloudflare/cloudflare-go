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
func (r *ZoneHealthcheckService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *SingleResponseEwxfoMtB, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a configured health check.
func (r *ZoneHealthcheckService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneHealthcheckUpdateParams, opts ...option.RequestOption) (res *SingleResponseEwxfoMtB, err error) {
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
func (r *ZoneHealthcheckService) HealthChecksNewHealthCheck(ctx context.Context, zoneIdentifier string, body ZoneHealthcheckHealthChecksNewHealthCheckParams, opts ...option.RequestOption) (res *SingleResponseEwxfoMtB, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/healthchecks", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured health checks.
func (r *ZoneHealthcheckService) HealthChecksListHealthChecks(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ResponseCollectionVfFXo19J, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/healthchecks", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ResponseCollectionVfFXo19J struct {
	Errors     []ResponseCollectionVfFXo19JError    `json:"errors"`
	Messages   []ResponseCollectionVfFXo19JMessage  `json:"messages"`
	Result     []ResponseCollectionVfFXo19JResult   `json:"result"`
	ResultInfo ResponseCollectionVfFXo19JResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ResponseCollectionVfFXo19JSuccess `json:"success"`
	JSON    responseCollectionVfFXo19JJSON    `json:"-"`
}

// responseCollectionVfFXo19JJSON contains the JSON metadata for the struct
// [ResponseCollectionVfFXo19J]
type responseCollectionVfFXo19JJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVfFXo19J) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionVfFXo19JError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    responseCollectionVfFXo19JErrorJSON `json:"-"`
}

// responseCollectionVfFXo19JErrorJSON contains the JSON metadata for the struct
// [ResponseCollectionVfFXo19JError]
type responseCollectionVfFXo19JErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVfFXo19JError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionVfFXo19JMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    responseCollectionVfFXo19JMessageJSON `json:"-"`
}

// responseCollectionVfFXo19JMessageJSON contains the JSON metadata for the struct
// [ResponseCollectionVfFXo19JMessage]
type responseCollectionVfFXo19JMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVfFXo19JMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionVfFXo19JResult struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []ResponseCollectionVfFXo19JResultCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig ResponseCollectionVfFXo19JResultHTTPConfig `json:"http_config,nullable"`
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
	Status ResponseCollectionVfFXo19JResultStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig ResponseCollectionVfFXo19JResultTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                               `json:"type"`
	JSON responseCollectionVfFXo19JResultJSON `json:"-"`
}

// responseCollectionVfFXo19JResultJSON contains the JSON metadata for the struct
// [ResponseCollectionVfFXo19JResult]
type responseCollectionVfFXo19JResultJSON struct {
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

func (r *ResponseCollectionVfFXo19JResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type ResponseCollectionVfFXo19JResultCheckRegion string

const (
	ResponseCollectionVfFXo19JResultCheckRegionWnam       ResponseCollectionVfFXo19JResultCheckRegion = "WNAM"
	ResponseCollectionVfFXo19JResultCheckRegionEnam       ResponseCollectionVfFXo19JResultCheckRegion = "ENAM"
	ResponseCollectionVfFXo19JResultCheckRegionWeu        ResponseCollectionVfFXo19JResultCheckRegion = "WEU"
	ResponseCollectionVfFXo19JResultCheckRegionEeu        ResponseCollectionVfFXo19JResultCheckRegion = "EEU"
	ResponseCollectionVfFXo19JResultCheckRegionNsam       ResponseCollectionVfFXo19JResultCheckRegion = "NSAM"
	ResponseCollectionVfFXo19JResultCheckRegionSsam       ResponseCollectionVfFXo19JResultCheckRegion = "SSAM"
	ResponseCollectionVfFXo19JResultCheckRegionOc         ResponseCollectionVfFXo19JResultCheckRegion = "OC"
	ResponseCollectionVfFXo19JResultCheckRegionMe         ResponseCollectionVfFXo19JResultCheckRegion = "ME"
	ResponseCollectionVfFXo19JResultCheckRegionNaf        ResponseCollectionVfFXo19JResultCheckRegion = "NAF"
	ResponseCollectionVfFXo19JResultCheckRegionSaf        ResponseCollectionVfFXo19JResultCheckRegion = "SAF"
	ResponseCollectionVfFXo19JResultCheckRegionIn         ResponseCollectionVfFXo19JResultCheckRegion = "IN"
	ResponseCollectionVfFXo19JResultCheckRegionSeas       ResponseCollectionVfFXo19JResultCheckRegion = "SEAS"
	ResponseCollectionVfFXo19JResultCheckRegionNeas       ResponseCollectionVfFXo19JResultCheckRegion = "NEAS"
	ResponseCollectionVfFXo19JResultCheckRegionAllRegions ResponseCollectionVfFXo19JResultCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type ResponseCollectionVfFXo19JResultHTTPConfig struct {
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
	Method ResponseCollectionVfFXo19JResultHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                          `json:"port"`
	JSON responseCollectionVfFXo19JResultHTTPConfigJSON `json:"-"`
}

// responseCollectionVfFXo19JResultHTTPConfigJSON contains the JSON metadata for
// the struct [ResponseCollectionVfFXo19JResultHTTPConfig]
type responseCollectionVfFXo19JResultHTTPConfigJSON struct {
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

func (r *ResponseCollectionVfFXo19JResultHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type ResponseCollectionVfFXo19JResultHTTPConfigMethod string

const (
	ResponseCollectionVfFXo19JResultHTTPConfigMethodGet  ResponseCollectionVfFXo19JResultHTTPConfigMethod = "GET"
	ResponseCollectionVfFXo19JResultHTTPConfigMethodHead ResponseCollectionVfFXo19JResultHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type ResponseCollectionVfFXo19JResultStatus string

const (
	ResponseCollectionVfFXo19JResultStatusUnknown   ResponseCollectionVfFXo19JResultStatus = "unknown"
	ResponseCollectionVfFXo19JResultStatusHealthy   ResponseCollectionVfFXo19JResultStatus = "healthy"
	ResponseCollectionVfFXo19JResultStatusUnhealthy ResponseCollectionVfFXo19JResultStatus = "unhealthy"
	ResponseCollectionVfFXo19JResultStatusSuspended ResponseCollectionVfFXo19JResultStatus = "suspended"
)

// Parameters specific to TCP health check.
type ResponseCollectionVfFXo19JResultTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method ResponseCollectionVfFXo19JResultTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                         `json:"port"`
	JSON responseCollectionVfFXo19JResultTcpConfigJSON `json:"-"`
}

// responseCollectionVfFXo19JResultTcpConfigJSON contains the JSON metadata for the
// struct [ResponseCollectionVfFXo19JResultTcpConfig]
type responseCollectionVfFXo19JResultTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVfFXo19JResultTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type ResponseCollectionVfFXo19JResultTcpConfigMethod string

const (
	ResponseCollectionVfFXo19JResultTcpConfigMethodConnectionEstablished ResponseCollectionVfFXo19JResultTcpConfigMethod = "connection_established"
)

type ResponseCollectionVfFXo19JResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       responseCollectionVfFXo19JResultInfoJSON `json:"-"`
}

// responseCollectionVfFXo19JResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionVfFXo19JResultInfo]
type responseCollectionVfFXo19JResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVfFXo19JResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseCollectionVfFXo19JSuccess bool

const (
	ResponseCollectionVfFXo19JSuccessTrue ResponseCollectionVfFXo19JSuccess = true
)

type SingleResponseEwxfoMtB struct {
	Errors   []SingleResponseEwxfoMtBError   `json:"errors"`
	Messages []SingleResponseEwxfoMtBMessage `json:"messages"`
	Result   SingleResponseEwxfoMtBResult    `json:"result"`
	// Whether the API call was successful
	Success SingleResponseEwxfoMtBSuccess `json:"success"`
	JSON    singleResponseEwxfoMtBJSON    `json:"-"`
}

// singleResponseEwxfoMtBJSON contains the JSON metadata for the struct
// [SingleResponseEwxfoMtB]
type singleResponseEwxfoMtBJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseEwxfoMtB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseEwxfoMtBError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleResponseEwxfoMtBErrorJSON `json:"-"`
}

// singleResponseEwxfoMtBErrorJSON contains the JSON metadata for the struct
// [SingleResponseEwxfoMtBError]
type singleResponseEwxfoMtBErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseEwxfoMtBError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseEwxfoMtBMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleResponseEwxfoMtBMessageJSON `json:"-"`
}

// singleResponseEwxfoMtBMessageJSON contains the JSON metadata for the struct
// [SingleResponseEwxfoMtBMessage]
type singleResponseEwxfoMtBMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseEwxfoMtBMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseEwxfoMtBResult struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []SingleResponseEwxfoMtBResultCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig SingleResponseEwxfoMtBResultHTTPConfig `json:"http_config,nullable"`
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
	Status SingleResponseEwxfoMtBResultStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig SingleResponseEwxfoMtBResultTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                           `json:"type"`
	JSON singleResponseEwxfoMtBResultJSON `json:"-"`
}

// singleResponseEwxfoMtBResultJSON contains the JSON metadata for the struct
// [SingleResponseEwxfoMtBResult]
type singleResponseEwxfoMtBResultJSON struct {
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

func (r *SingleResponseEwxfoMtBResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type SingleResponseEwxfoMtBResultCheckRegion string

const (
	SingleResponseEwxfoMtBResultCheckRegionWnam       SingleResponseEwxfoMtBResultCheckRegion = "WNAM"
	SingleResponseEwxfoMtBResultCheckRegionEnam       SingleResponseEwxfoMtBResultCheckRegion = "ENAM"
	SingleResponseEwxfoMtBResultCheckRegionWeu        SingleResponseEwxfoMtBResultCheckRegion = "WEU"
	SingleResponseEwxfoMtBResultCheckRegionEeu        SingleResponseEwxfoMtBResultCheckRegion = "EEU"
	SingleResponseEwxfoMtBResultCheckRegionNsam       SingleResponseEwxfoMtBResultCheckRegion = "NSAM"
	SingleResponseEwxfoMtBResultCheckRegionSsam       SingleResponseEwxfoMtBResultCheckRegion = "SSAM"
	SingleResponseEwxfoMtBResultCheckRegionOc         SingleResponseEwxfoMtBResultCheckRegion = "OC"
	SingleResponseEwxfoMtBResultCheckRegionMe         SingleResponseEwxfoMtBResultCheckRegion = "ME"
	SingleResponseEwxfoMtBResultCheckRegionNaf        SingleResponseEwxfoMtBResultCheckRegion = "NAF"
	SingleResponseEwxfoMtBResultCheckRegionSaf        SingleResponseEwxfoMtBResultCheckRegion = "SAF"
	SingleResponseEwxfoMtBResultCheckRegionIn         SingleResponseEwxfoMtBResultCheckRegion = "IN"
	SingleResponseEwxfoMtBResultCheckRegionSeas       SingleResponseEwxfoMtBResultCheckRegion = "SEAS"
	SingleResponseEwxfoMtBResultCheckRegionNeas       SingleResponseEwxfoMtBResultCheckRegion = "NEAS"
	SingleResponseEwxfoMtBResultCheckRegionAllRegions SingleResponseEwxfoMtBResultCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type SingleResponseEwxfoMtBResultHTTPConfig struct {
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
	Method SingleResponseEwxfoMtBResultHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                      `json:"port"`
	JSON singleResponseEwxfoMtBResultHTTPConfigJSON `json:"-"`
}

// singleResponseEwxfoMtBResultHTTPConfigJSON contains the JSON metadata for the
// struct [SingleResponseEwxfoMtBResultHTTPConfig]
type singleResponseEwxfoMtBResultHTTPConfigJSON struct {
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

func (r *SingleResponseEwxfoMtBResultHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type SingleResponseEwxfoMtBResultHTTPConfigMethod string

const (
	SingleResponseEwxfoMtBResultHTTPConfigMethodGet  SingleResponseEwxfoMtBResultHTTPConfigMethod = "GET"
	SingleResponseEwxfoMtBResultHTTPConfigMethodHead SingleResponseEwxfoMtBResultHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type SingleResponseEwxfoMtBResultStatus string

const (
	SingleResponseEwxfoMtBResultStatusUnknown   SingleResponseEwxfoMtBResultStatus = "unknown"
	SingleResponseEwxfoMtBResultStatusHealthy   SingleResponseEwxfoMtBResultStatus = "healthy"
	SingleResponseEwxfoMtBResultStatusUnhealthy SingleResponseEwxfoMtBResultStatus = "unhealthy"
	SingleResponseEwxfoMtBResultStatusSuspended SingleResponseEwxfoMtBResultStatus = "suspended"
)

// Parameters specific to TCP health check.
type SingleResponseEwxfoMtBResultTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method SingleResponseEwxfoMtBResultTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                     `json:"port"`
	JSON singleResponseEwxfoMtBResultTcpConfigJSON `json:"-"`
}

// singleResponseEwxfoMtBResultTcpConfigJSON contains the JSON metadata for the
// struct [SingleResponseEwxfoMtBResultTcpConfig]
type singleResponseEwxfoMtBResultTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseEwxfoMtBResultTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type SingleResponseEwxfoMtBResultTcpConfigMethod string

const (
	SingleResponseEwxfoMtBResultTcpConfigMethodConnectionEstablished SingleResponseEwxfoMtBResultTcpConfigMethod = "connection_established"
)

// Whether the API call was successful
type SingleResponseEwxfoMtBSuccess bool

const (
	SingleResponseEwxfoMtBSuccessTrue SingleResponseEwxfoMtBSuccess = true
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
	ID   interface{}                             `json:"id"`
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
