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

// HealthcheckPreviewService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewHealthcheckPreviewService] method
// instead.
type HealthcheckPreviewService struct {
	Options []option.RequestOption
}

// NewHealthcheckPreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHealthcheckPreviewService(opts ...option.RequestOption) (r *HealthcheckPreviewService) {
	r = &HealthcheckPreviewService{}
	r.Options = opts
	return
}

// Create a new preview health check.
func (r *HealthcheckPreviewService) New(ctx context.Context, zoneIdentifier string, body HealthcheckPreviewNewParams, opts ...option.RequestOption) (res *HealthcheckPreviewNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckPreviewNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/preview", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a health check.
func (r *HealthcheckPreviewService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *HealthcheckPreviewDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckPreviewDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/preview/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured health check preview.
func (r *HealthcheckPreviewService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *HealthcheckPreviewGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckPreviewGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/preview/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HealthcheckPreviewNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthcheckPreviewNewResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig HealthcheckPreviewNewResponseHTTPConfig `json:"http_config,nullable"`
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
	Status HealthcheckPreviewNewResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthcheckPreviewNewResponseTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                            `json:"type"`
	JSON healthcheckPreviewNewResponseJSON `json:"-"`
}

// healthcheckPreviewNewResponseJSON contains the JSON metadata for the struct
// [HealthcheckPreviewNewResponse]
type healthcheckPreviewNewResponseJSON struct {
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

func (r *HealthcheckPreviewNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewNewResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckPreviewNewResponseCheckRegion string

const (
	HealthcheckPreviewNewResponseCheckRegionWnam       HealthcheckPreviewNewResponseCheckRegion = "WNAM"
	HealthcheckPreviewNewResponseCheckRegionEnam       HealthcheckPreviewNewResponseCheckRegion = "ENAM"
	HealthcheckPreviewNewResponseCheckRegionWeu        HealthcheckPreviewNewResponseCheckRegion = "WEU"
	HealthcheckPreviewNewResponseCheckRegionEeu        HealthcheckPreviewNewResponseCheckRegion = "EEU"
	HealthcheckPreviewNewResponseCheckRegionNsam       HealthcheckPreviewNewResponseCheckRegion = "NSAM"
	HealthcheckPreviewNewResponseCheckRegionSsam       HealthcheckPreviewNewResponseCheckRegion = "SSAM"
	HealthcheckPreviewNewResponseCheckRegionOc         HealthcheckPreviewNewResponseCheckRegion = "OC"
	HealthcheckPreviewNewResponseCheckRegionMe         HealthcheckPreviewNewResponseCheckRegion = "ME"
	HealthcheckPreviewNewResponseCheckRegionNaf        HealthcheckPreviewNewResponseCheckRegion = "NAF"
	HealthcheckPreviewNewResponseCheckRegionSaf        HealthcheckPreviewNewResponseCheckRegion = "SAF"
	HealthcheckPreviewNewResponseCheckRegionIn         HealthcheckPreviewNewResponseCheckRegion = "IN"
	HealthcheckPreviewNewResponseCheckRegionSeas       HealthcheckPreviewNewResponseCheckRegion = "SEAS"
	HealthcheckPreviewNewResponseCheckRegionNeas       HealthcheckPreviewNewResponseCheckRegion = "NEAS"
	HealthcheckPreviewNewResponseCheckRegionAllRegions HealthcheckPreviewNewResponseCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckPreviewNewResponseHTTPConfig struct {
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
	Method HealthcheckPreviewNewResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                       `json:"port"`
	JSON healthcheckPreviewNewResponseHTTPConfigJSON `json:"-"`
}

// healthcheckPreviewNewResponseHTTPConfigJSON contains the JSON metadata for the
// struct [HealthcheckPreviewNewResponseHTTPConfig]
type healthcheckPreviewNewResponseHTTPConfigJSON struct {
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

func (r *HealthcheckPreviewNewResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewNewResponseHTTPConfigJSON) RawJSON() string {
	return r.raw
}

// The HTTP method to use for the health check.
type HealthcheckPreviewNewResponseHTTPConfigMethod string

const (
	HealthcheckPreviewNewResponseHTTPConfigMethodGet  HealthcheckPreviewNewResponseHTTPConfigMethod = "GET"
	HealthcheckPreviewNewResponseHTTPConfigMethodHead HealthcheckPreviewNewResponseHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type HealthcheckPreviewNewResponseStatus string

const (
	HealthcheckPreviewNewResponseStatusUnknown   HealthcheckPreviewNewResponseStatus = "unknown"
	HealthcheckPreviewNewResponseStatusHealthy   HealthcheckPreviewNewResponseStatus = "healthy"
	HealthcheckPreviewNewResponseStatusUnhealthy HealthcheckPreviewNewResponseStatus = "unhealthy"
	HealthcheckPreviewNewResponseStatusSuspended HealthcheckPreviewNewResponseStatus = "suspended"
)

// Parameters specific to TCP health check.
type HealthcheckPreviewNewResponseTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthcheckPreviewNewResponseTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                      `json:"port"`
	JSON healthcheckPreviewNewResponseTcpConfigJSON `json:"-"`
}

// healthcheckPreviewNewResponseTcpConfigJSON contains the JSON metadata for the
// struct [HealthcheckPreviewNewResponseTcpConfig]
type healthcheckPreviewNewResponseTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckPreviewNewResponseTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewNewResponseTcpConfigJSON) RawJSON() string {
	return r.raw
}

// The TCP connection method to use for the health check.
type HealthcheckPreviewNewResponseTcpConfigMethod string

const (
	HealthcheckPreviewNewResponseTcpConfigMethodConnectionEstablished HealthcheckPreviewNewResponseTcpConfigMethod = "connection_established"
)

type HealthcheckPreviewDeleteResponse struct {
	// Identifier
	ID   string                               `json:"id"`
	JSON healthcheckPreviewDeleteResponseJSON `json:"-"`
}

// healthcheckPreviewDeleteResponseJSON contains the JSON metadata for the struct
// [HealthcheckPreviewDeleteResponse]
type healthcheckPreviewDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckPreviewDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type HealthcheckPreviewGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthcheckPreviewGetResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig HealthcheckPreviewGetResponseHTTPConfig `json:"http_config,nullable"`
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
	Status HealthcheckPreviewGetResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthcheckPreviewGetResponseTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                            `json:"type"`
	JSON healthcheckPreviewGetResponseJSON `json:"-"`
}

// healthcheckPreviewGetResponseJSON contains the JSON metadata for the struct
// [HealthcheckPreviewGetResponse]
type healthcheckPreviewGetResponseJSON struct {
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

func (r *HealthcheckPreviewGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewGetResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckPreviewGetResponseCheckRegion string

const (
	HealthcheckPreviewGetResponseCheckRegionWnam       HealthcheckPreviewGetResponseCheckRegion = "WNAM"
	HealthcheckPreviewGetResponseCheckRegionEnam       HealthcheckPreviewGetResponseCheckRegion = "ENAM"
	HealthcheckPreviewGetResponseCheckRegionWeu        HealthcheckPreviewGetResponseCheckRegion = "WEU"
	HealthcheckPreviewGetResponseCheckRegionEeu        HealthcheckPreviewGetResponseCheckRegion = "EEU"
	HealthcheckPreviewGetResponseCheckRegionNsam       HealthcheckPreviewGetResponseCheckRegion = "NSAM"
	HealthcheckPreviewGetResponseCheckRegionSsam       HealthcheckPreviewGetResponseCheckRegion = "SSAM"
	HealthcheckPreviewGetResponseCheckRegionOc         HealthcheckPreviewGetResponseCheckRegion = "OC"
	HealthcheckPreviewGetResponseCheckRegionMe         HealthcheckPreviewGetResponseCheckRegion = "ME"
	HealthcheckPreviewGetResponseCheckRegionNaf        HealthcheckPreviewGetResponseCheckRegion = "NAF"
	HealthcheckPreviewGetResponseCheckRegionSaf        HealthcheckPreviewGetResponseCheckRegion = "SAF"
	HealthcheckPreviewGetResponseCheckRegionIn         HealthcheckPreviewGetResponseCheckRegion = "IN"
	HealthcheckPreviewGetResponseCheckRegionSeas       HealthcheckPreviewGetResponseCheckRegion = "SEAS"
	HealthcheckPreviewGetResponseCheckRegionNeas       HealthcheckPreviewGetResponseCheckRegion = "NEAS"
	HealthcheckPreviewGetResponseCheckRegionAllRegions HealthcheckPreviewGetResponseCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckPreviewGetResponseHTTPConfig struct {
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
	Method HealthcheckPreviewGetResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                       `json:"port"`
	JSON healthcheckPreviewGetResponseHTTPConfigJSON `json:"-"`
}

// healthcheckPreviewGetResponseHTTPConfigJSON contains the JSON metadata for the
// struct [HealthcheckPreviewGetResponseHTTPConfig]
type healthcheckPreviewGetResponseHTTPConfigJSON struct {
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

func (r *HealthcheckPreviewGetResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewGetResponseHTTPConfigJSON) RawJSON() string {
	return r.raw
}

// The HTTP method to use for the health check.
type HealthcheckPreviewGetResponseHTTPConfigMethod string

const (
	HealthcheckPreviewGetResponseHTTPConfigMethodGet  HealthcheckPreviewGetResponseHTTPConfigMethod = "GET"
	HealthcheckPreviewGetResponseHTTPConfigMethodHead HealthcheckPreviewGetResponseHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type HealthcheckPreviewGetResponseStatus string

const (
	HealthcheckPreviewGetResponseStatusUnknown   HealthcheckPreviewGetResponseStatus = "unknown"
	HealthcheckPreviewGetResponseStatusHealthy   HealthcheckPreviewGetResponseStatus = "healthy"
	HealthcheckPreviewGetResponseStatusUnhealthy HealthcheckPreviewGetResponseStatus = "unhealthy"
	HealthcheckPreviewGetResponseStatusSuspended HealthcheckPreviewGetResponseStatus = "suspended"
)

// Parameters specific to TCP health check.
type HealthcheckPreviewGetResponseTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthcheckPreviewGetResponseTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                      `json:"port"`
	JSON healthcheckPreviewGetResponseTcpConfigJSON `json:"-"`
}

// healthcheckPreviewGetResponseTcpConfigJSON contains the JSON metadata for the
// struct [HealthcheckPreviewGetResponseTcpConfig]
type healthcheckPreviewGetResponseTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckPreviewGetResponseTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewGetResponseTcpConfigJSON) RawJSON() string {
	return r.raw
}

// The TCP connection method to use for the health check.
type HealthcheckPreviewGetResponseTcpConfigMethod string

const (
	HealthcheckPreviewGetResponseTcpConfigMethodConnectionEstablished HealthcheckPreviewGetResponseTcpConfigMethod = "connection_established"
)

type HealthcheckPreviewNewParams struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]HealthcheckPreviewNewParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[HealthcheckPreviewNewParamsHTTPConfig] `json:"http_config"`
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
	TcpConfig param.Field[HealthcheckPreviewNewParamsTcpConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r HealthcheckPreviewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckPreviewNewParamsCheckRegion string

const (
	HealthcheckPreviewNewParamsCheckRegionWnam       HealthcheckPreviewNewParamsCheckRegion = "WNAM"
	HealthcheckPreviewNewParamsCheckRegionEnam       HealthcheckPreviewNewParamsCheckRegion = "ENAM"
	HealthcheckPreviewNewParamsCheckRegionWeu        HealthcheckPreviewNewParamsCheckRegion = "WEU"
	HealthcheckPreviewNewParamsCheckRegionEeu        HealthcheckPreviewNewParamsCheckRegion = "EEU"
	HealthcheckPreviewNewParamsCheckRegionNsam       HealthcheckPreviewNewParamsCheckRegion = "NSAM"
	HealthcheckPreviewNewParamsCheckRegionSsam       HealthcheckPreviewNewParamsCheckRegion = "SSAM"
	HealthcheckPreviewNewParamsCheckRegionOc         HealthcheckPreviewNewParamsCheckRegion = "OC"
	HealthcheckPreviewNewParamsCheckRegionMe         HealthcheckPreviewNewParamsCheckRegion = "ME"
	HealthcheckPreviewNewParamsCheckRegionNaf        HealthcheckPreviewNewParamsCheckRegion = "NAF"
	HealthcheckPreviewNewParamsCheckRegionSaf        HealthcheckPreviewNewParamsCheckRegion = "SAF"
	HealthcheckPreviewNewParamsCheckRegionIn         HealthcheckPreviewNewParamsCheckRegion = "IN"
	HealthcheckPreviewNewParamsCheckRegionSeas       HealthcheckPreviewNewParamsCheckRegion = "SEAS"
	HealthcheckPreviewNewParamsCheckRegionNeas       HealthcheckPreviewNewParamsCheckRegion = "NEAS"
	HealthcheckPreviewNewParamsCheckRegionAllRegions HealthcheckPreviewNewParamsCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckPreviewNewParamsHTTPConfig struct {
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
	Method param.Field[HealthcheckPreviewNewParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckPreviewNewParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type HealthcheckPreviewNewParamsHTTPConfigMethod string

const (
	HealthcheckPreviewNewParamsHTTPConfigMethodGet  HealthcheckPreviewNewParamsHTTPConfigMethod = "GET"
	HealthcheckPreviewNewParamsHTTPConfigMethodHead HealthcheckPreviewNewParamsHTTPConfigMethod = "HEAD"
)

// Parameters specific to TCP health check.
type HealthcheckPreviewNewParamsTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[HealthcheckPreviewNewParamsTcpConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckPreviewNewParamsTcpConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type HealthcheckPreviewNewParamsTcpConfigMethod string

const (
	HealthcheckPreviewNewParamsTcpConfigMethodConnectionEstablished HealthcheckPreviewNewParamsTcpConfigMethod = "connection_established"
)

type HealthcheckPreviewNewResponseEnvelope struct {
	Errors   []HealthcheckPreviewNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckPreviewNewResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthcheckPreviewNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckPreviewNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckPreviewNewResponseEnvelopeJSON    `json:"-"`
}

// healthcheckPreviewNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [HealthcheckPreviewNewResponseEnvelope]
type healthcheckPreviewNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckPreviewNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HealthcheckPreviewNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    healthcheckPreviewNewResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckPreviewNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HealthcheckPreviewNewResponseEnvelopeErrors]
type healthcheckPreviewNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckPreviewNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HealthcheckPreviewNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    healthcheckPreviewNewResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckPreviewNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [HealthcheckPreviewNewResponseEnvelopeMessages]
type healthcheckPreviewNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckPreviewNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HealthcheckPreviewNewResponseEnvelopeSuccess bool

const (
	HealthcheckPreviewNewResponseEnvelopeSuccessTrue HealthcheckPreviewNewResponseEnvelopeSuccess = true
)

type HealthcheckPreviewDeleteResponseEnvelope struct {
	Errors   []HealthcheckPreviewDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckPreviewDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthcheckPreviewDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckPreviewDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckPreviewDeleteResponseEnvelopeJSON    `json:"-"`
}

// healthcheckPreviewDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [HealthcheckPreviewDeleteResponseEnvelope]
type healthcheckPreviewDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckPreviewDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HealthcheckPreviewDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    healthcheckPreviewDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckPreviewDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [HealthcheckPreviewDeleteResponseEnvelopeErrors]
type healthcheckPreviewDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckPreviewDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HealthcheckPreviewDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    healthcheckPreviewDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckPreviewDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HealthcheckPreviewDeleteResponseEnvelopeMessages]
type healthcheckPreviewDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckPreviewDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HealthcheckPreviewDeleteResponseEnvelopeSuccess bool

const (
	HealthcheckPreviewDeleteResponseEnvelopeSuccessTrue HealthcheckPreviewDeleteResponseEnvelopeSuccess = true
)

type HealthcheckPreviewGetResponseEnvelope struct {
	Errors   []HealthcheckPreviewGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckPreviewGetResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthcheckPreviewGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckPreviewGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckPreviewGetResponseEnvelopeJSON    `json:"-"`
}

// healthcheckPreviewGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HealthcheckPreviewGetResponseEnvelope]
type healthcheckPreviewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckPreviewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HealthcheckPreviewGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    healthcheckPreviewGetResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckPreviewGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HealthcheckPreviewGetResponseEnvelopeErrors]
type healthcheckPreviewGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckPreviewGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HealthcheckPreviewGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    healthcheckPreviewGetResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckPreviewGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [HealthcheckPreviewGetResponseEnvelopeMessages]
type healthcheckPreviewGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckPreviewGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckPreviewGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HealthcheckPreviewGetResponseEnvelopeSuccess bool

const (
	HealthcheckPreviewGetResponseEnvelopeSuccessTrue HealthcheckPreviewGetResponseEnvelopeSuccess = true
)
