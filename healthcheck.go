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

// HealthcheckService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewHealthcheckService] method
// instead.
type HealthcheckService struct {
	Options  []option.RequestOption
	Previews *HealthcheckPreviewService
}

// NewHealthcheckService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHealthcheckService(opts ...option.RequestOption) (r *HealthcheckService) {
	r = &HealthcheckService{}
	r.Options = opts
	r.Previews = NewHealthcheckPreviewService(opts...)
	return
}

// Update a configured health check.
func (r *HealthcheckService) Update(ctx context.Context, zoneIdentifier string, identifier string, body HealthcheckUpdateParams, opts ...option.RequestOption) (res *HealthcheckUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a health check.
func (r *HealthcheckService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *HealthcheckDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured health check.
func (r *HealthcheckService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *HealthcheckGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a new health check.
func (r *HealthcheckService) HealthChecksNewHealthCheck(ctx context.Context, zoneIdentifier string, body HealthcheckHealthChecksNewHealthCheckParams, opts ...option.RequestOption) (res *HealthcheckHealthChecksNewHealthCheckResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckHealthChecksNewHealthCheckResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured health checks.
func (r *HealthcheckService) HealthChecksListHealthChecks(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]HealthcheckHealthChecksListHealthChecksResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckHealthChecksListHealthChecksResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HealthcheckUpdateResponse struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthcheckUpdateResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig HealthcheckUpdateResponseHTTPConfig `json:"http_config,nullable"`
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
	Status HealthcheckUpdateResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthcheckUpdateResponseTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                        `json:"type"`
	JSON healthcheckUpdateResponseJSON `json:"-"`
}

// healthcheckUpdateResponseJSON contains the JSON metadata for the struct
// [HealthcheckUpdateResponse]
type healthcheckUpdateResponseJSON struct {
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

func (r *HealthcheckUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckUpdateResponseCheckRegion string

const (
	HealthcheckUpdateResponseCheckRegionWnam       HealthcheckUpdateResponseCheckRegion = "WNAM"
	HealthcheckUpdateResponseCheckRegionEnam       HealthcheckUpdateResponseCheckRegion = "ENAM"
	HealthcheckUpdateResponseCheckRegionWeu        HealthcheckUpdateResponseCheckRegion = "WEU"
	HealthcheckUpdateResponseCheckRegionEeu        HealthcheckUpdateResponseCheckRegion = "EEU"
	HealthcheckUpdateResponseCheckRegionNsam       HealthcheckUpdateResponseCheckRegion = "NSAM"
	HealthcheckUpdateResponseCheckRegionSsam       HealthcheckUpdateResponseCheckRegion = "SSAM"
	HealthcheckUpdateResponseCheckRegionOc         HealthcheckUpdateResponseCheckRegion = "OC"
	HealthcheckUpdateResponseCheckRegionMe         HealthcheckUpdateResponseCheckRegion = "ME"
	HealthcheckUpdateResponseCheckRegionNaf        HealthcheckUpdateResponseCheckRegion = "NAF"
	HealthcheckUpdateResponseCheckRegionSaf        HealthcheckUpdateResponseCheckRegion = "SAF"
	HealthcheckUpdateResponseCheckRegionIn         HealthcheckUpdateResponseCheckRegion = "IN"
	HealthcheckUpdateResponseCheckRegionSeas       HealthcheckUpdateResponseCheckRegion = "SEAS"
	HealthcheckUpdateResponseCheckRegionNeas       HealthcheckUpdateResponseCheckRegion = "NEAS"
	HealthcheckUpdateResponseCheckRegionAllRegions HealthcheckUpdateResponseCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckUpdateResponseHTTPConfig struct {
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
	Method HealthcheckUpdateResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                   `json:"port"`
	JSON healthcheckUpdateResponseHTTPConfigJSON `json:"-"`
}

// healthcheckUpdateResponseHTTPConfigJSON contains the JSON metadata for the
// struct [HealthcheckUpdateResponseHTTPConfig]
type healthcheckUpdateResponseHTTPConfigJSON struct {
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

func (r *HealthcheckUpdateResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type HealthcheckUpdateResponseHTTPConfigMethod string

const (
	HealthcheckUpdateResponseHTTPConfigMethodGet  HealthcheckUpdateResponseHTTPConfigMethod = "GET"
	HealthcheckUpdateResponseHTTPConfigMethodHead HealthcheckUpdateResponseHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type HealthcheckUpdateResponseStatus string

const (
	HealthcheckUpdateResponseStatusUnknown   HealthcheckUpdateResponseStatus = "unknown"
	HealthcheckUpdateResponseStatusHealthy   HealthcheckUpdateResponseStatus = "healthy"
	HealthcheckUpdateResponseStatusUnhealthy HealthcheckUpdateResponseStatus = "unhealthy"
	HealthcheckUpdateResponseStatusSuspended HealthcheckUpdateResponseStatus = "suspended"
)

// Parameters specific to TCP health check.
type HealthcheckUpdateResponseTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthcheckUpdateResponseTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                  `json:"port"`
	JSON healthcheckUpdateResponseTcpConfigJSON `json:"-"`
}

// healthcheckUpdateResponseTcpConfigJSON contains the JSON metadata for the struct
// [HealthcheckUpdateResponseTcpConfig]
type healthcheckUpdateResponseTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckUpdateResponseTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type HealthcheckUpdateResponseTcpConfigMethod string

const (
	HealthcheckUpdateResponseTcpConfigMethodConnectionEstablished HealthcheckUpdateResponseTcpConfigMethod = "connection_established"
)

type HealthcheckDeleteResponse struct {
	// Identifier
	ID   string                        `json:"id"`
	JSON healthcheckDeleteResponseJSON `json:"-"`
}

// healthcheckDeleteResponseJSON contains the JSON metadata for the struct
// [HealthcheckDeleteResponse]
type healthcheckDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthcheckGetResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig HealthcheckGetResponseHTTPConfig `json:"http_config,nullable"`
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
	Status HealthcheckGetResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthcheckGetResponseTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                     `json:"type"`
	JSON healthcheckGetResponseJSON `json:"-"`
}

// healthcheckGetResponseJSON contains the JSON metadata for the struct
// [HealthcheckGetResponse]
type healthcheckGetResponseJSON struct {
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

func (r *HealthcheckGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckGetResponseCheckRegion string

const (
	HealthcheckGetResponseCheckRegionWnam       HealthcheckGetResponseCheckRegion = "WNAM"
	HealthcheckGetResponseCheckRegionEnam       HealthcheckGetResponseCheckRegion = "ENAM"
	HealthcheckGetResponseCheckRegionWeu        HealthcheckGetResponseCheckRegion = "WEU"
	HealthcheckGetResponseCheckRegionEeu        HealthcheckGetResponseCheckRegion = "EEU"
	HealthcheckGetResponseCheckRegionNsam       HealthcheckGetResponseCheckRegion = "NSAM"
	HealthcheckGetResponseCheckRegionSsam       HealthcheckGetResponseCheckRegion = "SSAM"
	HealthcheckGetResponseCheckRegionOc         HealthcheckGetResponseCheckRegion = "OC"
	HealthcheckGetResponseCheckRegionMe         HealthcheckGetResponseCheckRegion = "ME"
	HealthcheckGetResponseCheckRegionNaf        HealthcheckGetResponseCheckRegion = "NAF"
	HealthcheckGetResponseCheckRegionSaf        HealthcheckGetResponseCheckRegion = "SAF"
	HealthcheckGetResponseCheckRegionIn         HealthcheckGetResponseCheckRegion = "IN"
	HealthcheckGetResponseCheckRegionSeas       HealthcheckGetResponseCheckRegion = "SEAS"
	HealthcheckGetResponseCheckRegionNeas       HealthcheckGetResponseCheckRegion = "NEAS"
	HealthcheckGetResponseCheckRegionAllRegions HealthcheckGetResponseCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckGetResponseHTTPConfig struct {
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
	Method HealthcheckGetResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                `json:"port"`
	JSON healthcheckGetResponseHTTPConfigJSON `json:"-"`
}

// healthcheckGetResponseHTTPConfigJSON contains the JSON metadata for the struct
// [HealthcheckGetResponseHTTPConfig]
type healthcheckGetResponseHTTPConfigJSON struct {
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

func (r *HealthcheckGetResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type HealthcheckGetResponseHTTPConfigMethod string

const (
	HealthcheckGetResponseHTTPConfigMethodGet  HealthcheckGetResponseHTTPConfigMethod = "GET"
	HealthcheckGetResponseHTTPConfigMethodHead HealthcheckGetResponseHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type HealthcheckGetResponseStatus string

const (
	HealthcheckGetResponseStatusUnknown   HealthcheckGetResponseStatus = "unknown"
	HealthcheckGetResponseStatusHealthy   HealthcheckGetResponseStatus = "healthy"
	HealthcheckGetResponseStatusUnhealthy HealthcheckGetResponseStatus = "unhealthy"
	HealthcheckGetResponseStatusSuspended HealthcheckGetResponseStatus = "suspended"
)

// Parameters specific to TCP health check.
type HealthcheckGetResponseTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthcheckGetResponseTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                               `json:"port"`
	JSON healthcheckGetResponseTcpConfigJSON `json:"-"`
}

// healthcheckGetResponseTcpConfigJSON contains the JSON metadata for the struct
// [HealthcheckGetResponseTcpConfig]
type healthcheckGetResponseTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckGetResponseTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type HealthcheckGetResponseTcpConfigMethod string

const (
	HealthcheckGetResponseTcpConfigMethodConnectionEstablished HealthcheckGetResponseTcpConfigMethod = "connection_established"
)

type HealthcheckHealthChecksNewHealthCheckResponse struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthcheckHealthChecksNewHealthCheckResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig HealthcheckHealthChecksNewHealthCheckResponseHTTPConfig `json:"http_config,nullable"`
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
	Status HealthcheckHealthChecksNewHealthCheckResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthcheckHealthChecksNewHealthCheckResponseTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                                            `json:"type"`
	JSON healthcheckHealthChecksNewHealthCheckResponseJSON `json:"-"`
}

// healthcheckHealthChecksNewHealthCheckResponseJSON contains the JSON metadata for
// the struct [HealthcheckHealthChecksNewHealthCheckResponse]
type healthcheckHealthChecksNewHealthCheckResponseJSON struct {
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

func (r *HealthcheckHealthChecksNewHealthCheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckHealthChecksNewHealthCheckResponseCheckRegion string

const (
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionWnam       HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "WNAM"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionEnam       HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "ENAM"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionWeu        HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "WEU"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionEeu        HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "EEU"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionNsam       HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "NSAM"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionSsam       HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "SSAM"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionOc         HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "OC"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionMe         HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "ME"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionNaf        HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "NAF"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionSaf        HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "SAF"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionIn         HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "IN"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionSeas       HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "SEAS"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionNeas       HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "NEAS"
	HealthcheckHealthChecksNewHealthCheckResponseCheckRegionAllRegions HealthcheckHealthChecksNewHealthCheckResponseCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckHealthChecksNewHealthCheckResponseHTTPConfig struct {
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
	Method HealthcheckHealthChecksNewHealthCheckResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                                       `json:"port"`
	JSON healthcheckHealthChecksNewHealthCheckResponseHTTPConfigJSON `json:"-"`
}

// healthcheckHealthChecksNewHealthCheckResponseHTTPConfigJSON contains the JSON
// metadata for the struct
// [HealthcheckHealthChecksNewHealthCheckResponseHTTPConfig]
type healthcheckHealthChecksNewHealthCheckResponseHTTPConfigJSON struct {
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

func (r *HealthcheckHealthChecksNewHealthCheckResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type HealthcheckHealthChecksNewHealthCheckResponseHTTPConfigMethod string

const (
	HealthcheckHealthChecksNewHealthCheckResponseHTTPConfigMethodGet  HealthcheckHealthChecksNewHealthCheckResponseHTTPConfigMethod = "GET"
	HealthcheckHealthChecksNewHealthCheckResponseHTTPConfigMethodHead HealthcheckHealthChecksNewHealthCheckResponseHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type HealthcheckHealthChecksNewHealthCheckResponseStatus string

const (
	HealthcheckHealthChecksNewHealthCheckResponseStatusUnknown   HealthcheckHealthChecksNewHealthCheckResponseStatus = "unknown"
	HealthcheckHealthChecksNewHealthCheckResponseStatusHealthy   HealthcheckHealthChecksNewHealthCheckResponseStatus = "healthy"
	HealthcheckHealthChecksNewHealthCheckResponseStatusUnhealthy HealthcheckHealthChecksNewHealthCheckResponseStatus = "unhealthy"
	HealthcheckHealthChecksNewHealthCheckResponseStatusSuspended HealthcheckHealthChecksNewHealthCheckResponseStatus = "suspended"
)

// Parameters specific to TCP health check.
type HealthcheckHealthChecksNewHealthCheckResponseTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthcheckHealthChecksNewHealthCheckResponseTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                                      `json:"port"`
	JSON healthcheckHealthChecksNewHealthCheckResponseTcpConfigJSON `json:"-"`
}

// healthcheckHealthChecksNewHealthCheckResponseTcpConfigJSON contains the JSON
// metadata for the struct [HealthcheckHealthChecksNewHealthCheckResponseTcpConfig]
type healthcheckHealthChecksNewHealthCheckResponseTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckHealthChecksNewHealthCheckResponseTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type HealthcheckHealthChecksNewHealthCheckResponseTcpConfigMethod string

const (
	HealthcheckHealthChecksNewHealthCheckResponseTcpConfigMethodConnectionEstablished HealthcheckHealthChecksNewHealthCheckResponseTcpConfigMethod = "connection_established"
)

type HealthcheckHealthChecksListHealthChecksResponse struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthcheckHealthChecksListHealthChecksResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig HealthcheckHealthChecksListHealthChecksResponseHTTPConfig `json:"http_config,nullable"`
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
	Status HealthcheckHealthChecksListHealthChecksResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthcheckHealthChecksListHealthChecksResponseTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                                              `json:"type"`
	JSON healthcheckHealthChecksListHealthChecksResponseJSON `json:"-"`
}

// healthcheckHealthChecksListHealthChecksResponseJSON contains the JSON metadata
// for the struct [HealthcheckHealthChecksListHealthChecksResponse]
type healthcheckHealthChecksListHealthChecksResponseJSON struct {
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

func (r *HealthcheckHealthChecksListHealthChecksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckHealthChecksListHealthChecksResponseCheckRegion string

const (
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionWnam       HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "WNAM"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionEnam       HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "ENAM"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionWeu        HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "WEU"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionEeu        HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "EEU"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionNsam       HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "NSAM"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionSsam       HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "SSAM"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionOc         HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "OC"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionMe         HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "ME"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionNaf        HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "NAF"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionSaf        HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "SAF"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionIn         HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "IN"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionSeas       HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "SEAS"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionNeas       HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "NEAS"
	HealthcheckHealthChecksListHealthChecksResponseCheckRegionAllRegions HealthcheckHealthChecksListHealthChecksResponseCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckHealthChecksListHealthChecksResponseHTTPConfig struct {
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
	Method HealthcheckHealthChecksListHealthChecksResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                                         `json:"port"`
	JSON healthcheckHealthChecksListHealthChecksResponseHTTPConfigJSON `json:"-"`
}

// healthcheckHealthChecksListHealthChecksResponseHTTPConfigJSON contains the JSON
// metadata for the struct
// [HealthcheckHealthChecksListHealthChecksResponseHTTPConfig]
type healthcheckHealthChecksListHealthChecksResponseHTTPConfigJSON struct {
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

func (r *HealthcheckHealthChecksListHealthChecksResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type HealthcheckHealthChecksListHealthChecksResponseHTTPConfigMethod string

const (
	HealthcheckHealthChecksListHealthChecksResponseHTTPConfigMethodGet  HealthcheckHealthChecksListHealthChecksResponseHTTPConfigMethod = "GET"
	HealthcheckHealthChecksListHealthChecksResponseHTTPConfigMethodHead HealthcheckHealthChecksListHealthChecksResponseHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type HealthcheckHealthChecksListHealthChecksResponseStatus string

const (
	HealthcheckHealthChecksListHealthChecksResponseStatusUnknown   HealthcheckHealthChecksListHealthChecksResponseStatus = "unknown"
	HealthcheckHealthChecksListHealthChecksResponseStatusHealthy   HealthcheckHealthChecksListHealthChecksResponseStatus = "healthy"
	HealthcheckHealthChecksListHealthChecksResponseStatusUnhealthy HealthcheckHealthChecksListHealthChecksResponseStatus = "unhealthy"
	HealthcheckHealthChecksListHealthChecksResponseStatusSuspended HealthcheckHealthChecksListHealthChecksResponseStatus = "suspended"
)

// Parameters specific to TCP health check.
type HealthcheckHealthChecksListHealthChecksResponseTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthcheckHealthChecksListHealthChecksResponseTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                                        `json:"port"`
	JSON healthcheckHealthChecksListHealthChecksResponseTcpConfigJSON `json:"-"`
}

// healthcheckHealthChecksListHealthChecksResponseTcpConfigJSON contains the JSON
// metadata for the struct
// [HealthcheckHealthChecksListHealthChecksResponseTcpConfig]
type healthcheckHealthChecksListHealthChecksResponseTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckHealthChecksListHealthChecksResponseTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type HealthcheckHealthChecksListHealthChecksResponseTcpConfigMethod string

const (
	HealthcheckHealthChecksListHealthChecksResponseTcpConfigMethodConnectionEstablished HealthcheckHealthChecksListHealthChecksResponseTcpConfigMethod = "connection_established"
)

type HealthcheckUpdateParams struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]HealthcheckUpdateParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[HealthcheckUpdateParamsHTTPConfig] `json:"http_config"`
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
	TcpConfig param.Field[HealthcheckUpdateParamsTcpConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r HealthcheckUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckUpdateParamsCheckRegion string

const (
	HealthcheckUpdateParamsCheckRegionWnam       HealthcheckUpdateParamsCheckRegion = "WNAM"
	HealthcheckUpdateParamsCheckRegionEnam       HealthcheckUpdateParamsCheckRegion = "ENAM"
	HealthcheckUpdateParamsCheckRegionWeu        HealthcheckUpdateParamsCheckRegion = "WEU"
	HealthcheckUpdateParamsCheckRegionEeu        HealthcheckUpdateParamsCheckRegion = "EEU"
	HealthcheckUpdateParamsCheckRegionNsam       HealthcheckUpdateParamsCheckRegion = "NSAM"
	HealthcheckUpdateParamsCheckRegionSsam       HealthcheckUpdateParamsCheckRegion = "SSAM"
	HealthcheckUpdateParamsCheckRegionOc         HealthcheckUpdateParamsCheckRegion = "OC"
	HealthcheckUpdateParamsCheckRegionMe         HealthcheckUpdateParamsCheckRegion = "ME"
	HealthcheckUpdateParamsCheckRegionNaf        HealthcheckUpdateParamsCheckRegion = "NAF"
	HealthcheckUpdateParamsCheckRegionSaf        HealthcheckUpdateParamsCheckRegion = "SAF"
	HealthcheckUpdateParamsCheckRegionIn         HealthcheckUpdateParamsCheckRegion = "IN"
	HealthcheckUpdateParamsCheckRegionSeas       HealthcheckUpdateParamsCheckRegion = "SEAS"
	HealthcheckUpdateParamsCheckRegionNeas       HealthcheckUpdateParamsCheckRegion = "NEAS"
	HealthcheckUpdateParamsCheckRegionAllRegions HealthcheckUpdateParamsCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckUpdateParamsHTTPConfig struct {
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
	Method param.Field[HealthcheckUpdateParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckUpdateParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type HealthcheckUpdateParamsHTTPConfigMethod string

const (
	HealthcheckUpdateParamsHTTPConfigMethodGet  HealthcheckUpdateParamsHTTPConfigMethod = "GET"
	HealthcheckUpdateParamsHTTPConfigMethodHead HealthcheckUpdateParamsHTTPConfigMethod = "HEAD"
)

// Parameters specific to TCP health check.
type HealthcheckUpdateParamsTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[HealthcheckUpdateParamsTcpConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckUpdateParamsTcpConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type HealthcheckUpdateParamsTcpConfigMethod string

const (
	HealthcheckUpdateParamsTcpConfigMethodConnectionEstablished HealthcheckUpdateParamsTcpConfigMethod = "connection_established"
)

type HealthcheckUpdateResponseEnvelope struct {
	Errors   []HealthcheckUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthcheckUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckUpdateResponseEnvelopeJSON    `json:"-"`
}

// healthcheckUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [HealthcheckUpdateResponseEnvelope]
type healthcheckUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    healthcheckUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HealthcheckUpdateResponseEnvelopeErrors]
type healthcheckUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    healthcheckUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HealthcheckUpdateResponseEnvelopeMessages]
type healthcheckUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckUpdateResponseEnvelopeSuccess bool

const (
	HealthcheckUpdateResponseEnvelopeSuccessTrue HealthcheckUpdateResponseEnvelopeSuccess = true
)

type HealthcheckDeleteResponseEnvelope struct {
	Errors   []HealthcheckDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthcheckDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckDeleteResponseEnvelopeJSON    `json:"-"`
}

// healthcheckDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [HealthcheckDeleteResponseEnvelope]
type healthcheckDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    healthcheckDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HealthcheckDeleteResponseEnvelopeErrors]
type healthcheckDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    healthcheckDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HealthcheckDeleteResponseEnvelopeMessages]
type healthcheckDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckDeleteResponseEnvelopeSuccess bool

const (
	HealthcheckDeleteResponseEnvelopeSuccessTrue HealthcheckDeleteResponseEnvelopeSuccess = true
)

type HealthcheckGetResponseEnvelope struct {
	Errors   []HealthcheckGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckGetResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthcheckGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckGetResponseEnvelopeJSON    `json:"-"`
}

// healthcheckGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HealthcheckGetResponseEnvelope]
type healthcheckGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    healthcheckGetResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HealthcheckGetResponseEnvelopeErrors]
type healthcheckGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    healthcheckGetResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HealthcheckGetResponseEnvelopeMessages]
type healthcheckGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckGetResponseEnvelopeSuccess bool

const (
	HealthcheckGetResponseEnvelopeSuccessTrue HealthcheckGetResponseEnvelopeSuccess = true
)

type HealthcheckHealthChecksNewHealthCheckParams struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]HealthcheckHealthChecksNewHealthCheckParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[HealthcheckHealthChecksNewHealthCheckParamsHTTPConfig] `json:"http_config"`
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
	TcpConfig param.Field[HealthcheckHealthChecksNewHealthCheckParamsTcpConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r HealthcheckHealthChecksNewHealthCheckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckHealthChecksNewHealthCheckParamsCheckRegion string

const (
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionWnam       HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "WNAM"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionEnam       HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "ENAM"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionWeu        HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "WEU"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionEeu        HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "EEU"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionNsam       HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "NSAM"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionSsam       HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "SSAM"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionOc         HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "OC"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionMe         HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "ME"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionNaf        HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "NAF"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionSaf        HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "SAF"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionIn         HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "IN"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionSeas       HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "SEAS"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionNeas       HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "NEAS"
	HealthcheckHealthChecksNewHealthCheckParamsCheckRegionAllRegions HealthcheckHealthChecksNewHealthCheckParamsCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckHealthChecksNewHealthCheckParamsHTTPConfig struct {
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
	Method param.Field[HealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckHealthChecksNewHealthCheckParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type HealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethod string

const (
	HealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethodGet  HealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethod = "GET"
	HealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethodHead HealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethod = "HEAD"
)

// Parameters specific to TCP health check.
type HealthcheckHealthChecksNewHealthCheckParamsTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[HealthcheckHealthChecksNewHealthCheckParamsTcpConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckHealthChecksNewHealthCheckParamsTcpConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type HealthcheckHealthChecksNewHealthCheckParamsTcpConfigMethod string

const (
	HealthcheckHealthChecksNewHealthCheckParamsTcpConfigMethodConnectionEstablished HealthcheckHealthChecksNewHealthCheckParamsTcpConfigMethod = "connection_established"
)

type HealthcheckHealthChecksNewHealthCheckResponseEnvelope struct {
	Errors   []HealthcheckHealthChecksNewHealthCheckResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckHealthChecksNewHealthCheckResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthcheckHealthChecksNewHealthCheckResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckHealthChecksNewHealthCheckResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckHealthChecksNewHealthCheckResponseEnvelopeJSON    `json:"-"`
}

// healthcheckHealthChecksNewHealthCheckResponseEnvelopeJSON contains the JSON
// metadata for the struct [HealthcheckHealthChecksNewHealthCheckResponseEnvelope]
type healthcheckHealthChecksNewHealthCheckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckHealthChecksNewHealthCheckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckHealthChecksNewHealthCheckResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    healthcheckHealthChecksNewHealthCheckResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckHealthChecksNewHealthCheckResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [HealthcheckHealthChecksNewHealthCheckResponseEnvelopeErrors]
type healthcheckHealthChecksNewHealthCheckResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckHealthChecksNewHealthCheckResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckHealthChecksNewHealthCheckResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    healthcheckHealthChecksNewHealthCheckResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckHealthChecksNewHealthCheckResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [HealthcheckHealthChecksNewHealthCheckResponseEnvelopeMessages]
type healthcheckHealthChecksNewHealthCheckResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckHealthChecksNewHealthCheckResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckHealthChecksNewHealthCheckResponseEnvelopeSuccess bool

const (
	HealthcheckHealthChecksNewHealthCheckResponseEnvelopeSuccessTrue HealthcheckHealthChecksNewHealthCheckResponseEnvelopeSuccess = true
)

type HealthcheckHealthChecksListHealthChecksResponseEnvelope struct {
	Errors   []HealthcheckHealthChecksListHealthChecksResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckHealthChecksListHealthChecksResponseEnvelopeMessages `json:"messages,required"`
	Result   []HealthcheckHealthChecksListHealthChecksResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    HealthcheckHealthChecksListHealthChecksResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo HealthcheckHealthChecksListHealthChecksResponseEnvelopeResultInfo `json:"result_info"`
	JSON       healthcheckHealthChecksListHealthChecksResponseEnvelopeJSON       `json:"-"`
}

// healthcheckHealthChecksListHealthChecksResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [HealthcheckHealthChecksListHealthChecksResponseEnvelope]
type healthcheckHealthChecksListHealthChecksResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckHealthChecksListHealthChecksResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckHealthChecksListHealthChecksResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    healthcheckHealthChecksListHealthChecksResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckHealthChecksListHealthChecksResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [HealthcheckHealthChecksListHealthChecksResponseEnvelopeErrors]
type healthcheckHealthChecksListHealthChecksResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckHealthChecksListHealthChecksResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckHealthChecksListHealthChecksResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    healthcheckHealthChecksListHealthChecksResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckHealthChecksListHealthChecksResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [HealthcheckHealthChecksListHealthChecksResponseEnvelopeMessages]
type healthcheckHealthChecksListHealthChecksResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckHealthChecksListHealthChecksResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckHealthChecksListHealthChecksResponseEnvelopeSuccess bool

const (
	HealthcheckHealthChecksListHealthChecksResponseEnvelopeSuccessTrue HealthcheckHealthChecksListHealthChecksResponseEnvelopeSuccess = true
)

type HealthcheckHealthChecksListHealthChecksResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                               `json:"total_count"`
	JSON       healthcheckHealthChecksListHealthChecksResponseEnvelopeResultInfoJSON `json:"-"`
}

// healthcheckHealthChecksListHealthChecksResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [HealthcheckHealthChecksListHealthChecksResponseEnvelopeResultInfo]
type healthcheckHealthChecksListHealthChecksResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckHealthChecksListHealthChecksResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
