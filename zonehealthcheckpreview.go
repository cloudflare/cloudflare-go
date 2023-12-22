// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneHealthcheckPreviewService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneHealthcheckPreviewService]
// method instead.
type ZoneHealthcheckPreviewService struct {
	Options []option.RequestOption
}

// NewZoneHealthcheckPreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneHealthcheckPreviewService(opts ...option.RequestOption) (r *ZoneHealthcheckPreviewService) {
	r = &ZoneHealthcheckPreviewService{}
	r.Options = opts
	return
}

// Fetch a single configured health check preview.
func (r *ZoneHealthcheckPreviewService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *SingleResponseEwxfoMtB, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/healthchecks/preview/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a health check.
func (r *ZoneHealthcheckPreviewService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneHealthcheckPreviewDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/healthchecks/preview/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new preview health check.
func (r *ZoneHealthcheckPreviewService) HealthChecksNewPreviewHealthCheck(ctx context.Context, zoneIdentifier string, body ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParams, opts ...option.RequestOption) (res *SingleResponseEwxfoMtB, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/healthchecks/preview", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneHealthcheckPreviewDeleteResponse struct {
	Errors   []ZoneHealthcheckPreviewDeleteResponseError   `json:"errors"`
	Messages []ZoneHealthcheckPreviewDeleteResponseMessage `json:"messages"`
	Result   ZoneHealthcheckPreviewDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneHealthcheckPreviewDeleteResponseSuccess `json:"success"`
	JSON    zoneHealthcheckPreviewDeleteResponseJSON    `json:"-"`
}

// zoneHealthcheckPreviewDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneHealthcheckPreviewDeleteResponse]
type zoneHealthcheckPreviewDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckPreviewDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHealthcheckPreviewDeleteResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneHealthcheckPreviewDeleteResponseErrorJSON `json:"-"`
}

// zoneHealthcheckPreviewDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneHealthcheckPreviewDeleteResponseError]
type zoneHealthcheckPreviewDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckPreviewDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHealthcheckPreviewDeleteResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneHealthcheckPreviewDeleteResponseMessageJSON `json:"-"`
}

// zoneHealthcheckPreviewDeleteResponseMessageJSON contains the JSON metadata for
// the struct [ZoneHealthcheckPreviewDeleteResponseMessage]
type zoneHealthcheckPreviewDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckPreviewDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneHealthcheckPreviewDeleteResponseResult struct {
	ID   interface{}                                    `json:"id"`
	JSON zoneHealthcheckPreviewDeleteResponseResultJSON `json:"-"`
}

// zoneHealthcheckPreviewDeleteResponseResultJSON contains the JSON metadata for
// the struct [ZoneHealthcheckPreviewDeleteResponseResult]
type zoneHealthcheckPreviewDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckPreviewDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneHealthcheckPreviewDeleteResponseSuccess bool

const (
	ZoneHealthcheckPreviewDeleteResponseSuccessTrue ZoneHealthcheckPreviewDeleteResponseSuccess = true
)

type ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParams struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsHTTPConfig] `json:"http_config"`
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
	TcpConfig param.Field[ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsTcpConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion string

const (
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionWnam       ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "WNAM"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionEnam       ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "ENAM"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionWeu        ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "WEU"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionEeu        ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "EEU"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionNsam       ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "NSAM"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionSsam       ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "SSAM"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionOc         ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "OC"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionMe         ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "ME"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionNaf        ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "NAF"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionSaf        ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "SAF"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionIn         ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "IN"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionSeas       ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "SEAS"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionNeas       ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "NEAS"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegionAllRegions ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsHTTPConfig struct {
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
	Method param.Field[ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsHTTPConfigMethod string

const (
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsHTTPConfigMethodGet  ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsHTTPConfigMethod = "GET"
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsHTTPConfigMethodHead ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsHTTPConfigMethod = "HEAD"
)

// Parameters specific to TCP health check.
type ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsTcpConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsTcpConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsTcpConfigMethod string

const (
	ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsTcpConfigMethodConnectionEstablished ZoneHealthcheckPreviewHealthChecksNewPreviewHealthCheckParamsTcpConfigMethod = "connection_established"
)
