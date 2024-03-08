// File generated from our OpenAPI spec by Stainless.

package healthchecks

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

// PreviewService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPreviewService] method instead.
type PreviewService struct {
	Options []option.RequestOption
}

// NewPreviewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPreviewService(opts ...option.RequestOption) (r *PreviewService) {
	r = &PreviewService{}
	r.Options = opts
	return
}

// Create a new preview health check.
func (r *PreviewService) New(ctx context.Context, zoneIdentifier string, body PreviewNewParams, opts ...option.RequestOption) (res *PreviewNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PreviewNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/preview", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a health check.
func (r *PreviewService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *PreviewDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PreviewDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/preview/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured health check preview.
func (r *PreviewService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *PreviewGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PreviewGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/preview/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PreviewNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []PreviewNewResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig PreviewNewResponseHTTPConfig `json:"http_config,nullable"`
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
	Status PreviewNewResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig PreviewNewResponseTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                 `json:"type"`
	JSON previewNewResponseJSON `json:"-"`
}

// previewNewResponseJSON contains the JSON metadata for the struct
// [PreviewNewResponse]
type previewNewResponseJSON struct {
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

func (r *PreviewNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewNewResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type PreviewNewResponseCheckRegion string

const (
	PreviewNewResponseCheckRegionWnam       PreviewNewResponseCheckRegion = "WNAM"
	PreviewNewResponseCheckRegionEnam       PreviewNewResponseCheckRegion = "ENAM"
	PreviewNewResponseCheckRegionWeu        PreviewNewResponseCheckRegion = "WEU"
	PreviewNewResponseCheckRegionEeu        PreviewNewResponseCheckRegion = "EEU"
	PreviewNewResponseCheckRegionNsam       PreviewNewResponseCheckRegion = "NSAM"
	PreviewNewResponseCheckRegionSsam       PreviewNewResponseCheckRegion = "SSAM"
	PreviewNewResponseCheckRegionOc         PreviewNewResponseCheckRegion = "OC"
	PreviewNewResponseCheckRegionMe         PreviewNewResponseCheckRegion = "ME"
	PreviewNewResponseCheckRegionNaf        PreviewNewResponseCheckRegion = "NAF"
	PreviewNewResponseCheckRegionSaf        PreviewNewResponseCheckRegion = "SAF"
	PreviewNewResponseCheckRegionIn         PreviewNewResponseCheckRegion = "IN"
	PreviewNewResponseCheckRegionSeas       PreviewNewResponseCheckRegion = "SEAS"
	PreviewNewResponseCheckRegionNeas       PreviewNewResponseCheckRegion = "NEAS"
	PreviewNewResponseCheckRegionAllRegions PreviewNewResponseCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type PreviewNewResponseHTTPConfig struct {
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
	Method PreviewNewResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                            `json:"port"`
	JSON previewNewResponseHTTPConfigJSON `json:"-"`
}

// previewNewResponseHTTPConfigJSON contains the JSON metadata for the struct
// [PreviewNewResponseHTTPConfig]
type previewNewResponseHTTPConfigJSON struct {
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

func (r *PreviewNewResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewNewResponseHTTPConfigJSON) RawJSON() string {
	return r.raw
}

// The HTTP method to use for the health check.
type PreviewNewResponseHTTPConfigMethod string

const (
	PreviewNewResponseHTTPConfigMethodGet  PreviewNewResponseHTTPConfigMethod = "GET"
	PreviewNewResponseHTTPConfigMethodHead PreviewNewResponseHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type PreviewNewResponseStatus string

const (
	PreviewNewResponseStatusUnknown   PreviewNewResponseStatus = "unknown"
	PreviewNewResponseStatusHealthy   PreviewNewResponseStatus = "healthy"
	PreviewNewResponseStatusUnhealthy PreviewNewResponseStatus = "unhealthy"
	PreviewNewResponseStatusSuspended PreviewNewResponseStatus = "suspended"
)

// Parameters specific to TCP health check.
type PreviewNewResponseTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method PreviewNewResponseTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                           `json:"port"`
	JSON previewNewResponseTcpConfigJSON `json:"-"`
}

// previewNewResponseTcpConfigJSON contains the JSON metadata for the struct
// [PreviewNewResponseTcpConfig]
type previewNewResponseTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewNewResponseTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewNewResponseTcpConfigJSON) RawJSON() string {
	return r.raw
}

// The TCP connection method to use for the health check.
type PreviewNewResponseTcpConfigMethod string

const (
	PreviewNewResponseTcpConfigMethodConnectionEstablished PreviewNewResponseTcpConfigMethod = "connection_established"
)

type PreviewDeleteResponse struct {
	// Identifier
	ID   string                    `json:"id"`
	JSON previewDeleteResponseJSON `json:"-"`
}

// previewDeleteResponseJSON contains the JSON metadata for the struct
// [PreviewDeleteResponse]
type previewDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PreviewGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []PreviewGetResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig PreviewGetResponseHTTPConfig `json:"http_config,nullable"`
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
	Status PreviewGetResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig PreviewGetResponseTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                 `json:"type"`
	JSON previewGetResponseJSON `json:"-"`
}

// previewGetResponseJSON contains the JSON metadata for the struct
// [PreviewGetResponse]
type previewGetResponseJSON struct {
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

func (r *PreviewGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewGetResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type PreviewGetResponseCheckRegion string

const (
	PreviewGetResponseCheckRegionWnam       PreviewGetResponseCheckRegion = "WNAM"
	PreviewGetResponseCheckRegionEnam       PreviewGetResponseCheckRegion = "ENAM"
	PreviewGetResponseCheckRegionWeu        PreviewGetResponseCheckRegion = "WEU"
	PreviewGetResponseCheckRegionEeu        PreviewGetResponseCheckRegion = "EEU"
	PreviewGetResponseCheckRegionNsam       PreviewGetResponseCheckRegion = "NSAM"
	PreviewGetResponseCheckRegionSsam       PreviewGetResponseCheckRegion = "SSAM"
	PreviewGetResponseCheckRegionOc         PreviewGetResponseCheckRegion = "OC"
	PreviewGetResponseCheckRegionMe         PreviewGetResponseCheckRegion = "ME"
	PreviewGetResponseCheckRegionNaf        PreviewGetResponseCheckRegion = "NAF"
	PreviewGetResponseCheckRegionSaf        PreviewGetResponseCheckRegion = "SAF"
	PreviewGetResponseCheckRegionIn         PreviewGetResponseCheckRegion = "IN"
	PreviewGetResponseCheckRegionSeas       PreviewGetResponseCheckRegion = "SEAS"
	PreviewGetResponseCheckRegionNeas       PreviewGetResponseCheckRegion = "NEAS"
	PreviewGetResponseCheckRegionAllRegions PreviewGetResponseCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type PreviewGetResponseHTTPConfig struct {
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
	Method PreviewGetResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                            `json:"port"`
	JSON previewGetResponseHTTPConfigJSON `json:"-"`
}

// previewGetResponseHTTPConfigJSON contains the JSON metadata for the struct
// [PreviewGetResponseHTTPConfig]
type previewGetResponseHTTPConfigJSON struct {
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

func (r *PreviewGetResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewGetResponseHTTPConfigJSON) RawJSON() string {
	return r.raw
}

// The HTTP method to use for the health check.
type PreviewGetResponseHTTPConfigMethod string

const (
	PreviewGetResponseHTTPConfigMethodGet  PreviewGetResponseHTTPConfigMethod = "GET"
	PreviewGetResponseHTTPConfigMethodHead PreviewGetResponseHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type PreviewGetResponseStatus string

const (
	PreviewGetResponseStatusUnknown   PreviewGetResponseStatus = "unknown"
	PreviewGetResponseStatusHealthy   PreviewGetResponseStatus = "healthy"
	PreviewGetResponseStatusUnhealthy PreviewGetResponseStatus = "unhealthy"
	PreviewGetResponseStatusSuspended PreviewGetResponseStatus = "suspended"
)

// Parameters specific to TCP health check.
type PreviewGetResponseTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method PreviewGetResponseTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                           `json:"port"`
	JSON previewGetResponseTcpConfigJSON `json:"-"`
}

// previewGetResponseTcpConfigJSON contains the JSON metadata for the struct
// [PreviewGetResponseTcpConfig]
type previewGetResponseTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewGetResponseTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewGetResponseTcpConfigJSON) RawJSON() string {
	return r.raw
}

// The TCP connection method to use for the health check.
type PreviewGetResponseTcpConfigMethod string

const (
	PreviewGetResponseTcpConfigMethodConnectionEstablished PreviewGetResponseTcpConfigMethod = "connection_established"
)

type PreviewNewParams struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]PreviewNewParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[PreviewNewParamsHTTPConfig] `json:"http_config"`
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
	TcpConfig param.Field[PreviewNewParamsTcpConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r PreviewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type PreviewNewParamsCheckRegion string

const (
	PreviewNewParamsCheckRegionWnam       PreviewNewParamsCheckRegion = "WNAM"
	PreviewNewParamsCheckRegionEnam       PreviewNewParamsCheckRegion = "ENAM"
	PreviewNewParamsCheckRegionWeu        PreviewNewParamsCheckRegion = "WEU"
	PreviewNewParamsCheckRegionEeu        PreviewNewParamsCheckRegion = "EEU"
	PreviewNewParamsCheckRegionNsam       PreviewNewParamsCheckRegion = "NSAM"
	PreviewNewParamsCheckRegionSsam       PreviewNewParamsCheckRegion = "SSAM"
	PreviewNewParamsCheckRegionOc         PreviewNewParamsCheckRegion = "OC"
	PreviewNewParamsCheckRegionMe         PreviewNewParamsCheckRegion = "ME"
	PreviewNewParamsCheckRegionNaf        PreviewNewParamsCheckRegion = "NAF"
	PreviewNewParamsCheckRegionSaf        PreviewNewParamsCheckRegion = "SAF"
	PreviewNewParamsCheckRegionIn         PreviewNewParamsCheckRegion = "IN"
	PreviewNewParamsCheckRegionSeas       PreviewNewParamsCheckRegion = "SEAS"
	PreviewNewParamsCheckRegionNeas       PreviewNewParamsCheckRegion = "NEAS"
	PreviewNewParamsCheckRegionAllRegions PreviewNewParamsCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type PreviewNewParamsHTTPConfig struct {
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
	Method param.Field[PreviewNewParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r PreviewNewParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type PreviewNewParamsHTTPConfigMethod string

const (
	PreviewNewParamsHTTPConfigMethodGet  PreviewNewParamsHTTPConfigMethod = "GET"
	PreviewNewParamsHTTPConfigMethodHead PreviewNewParamsHTTPConfigMethod = "HEAD"
)

// Parameters specific to TCP health check.
type PreviewNewParamsTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[PreviewNewParamsTcpConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r PreviewNewParamsTcpConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type PreviewNewParamsTcpConfigMethod string

const (
	PreviewNewParamsTcpConfigMethodConnectionEstablished PreviewNewParamsTcpConfigMethod = "connection_established"
)

type PreviewNewResponseEnvelope struct {
	Errors   []PreviewNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PreviewNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PreviewNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PreviewNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    previewNewResponseEnvelopeJSON    `json:"-"`
}

// previewNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PreviewNewResponseEnvelope]
type previewNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PreviewNewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    previewNewResponseEnvelopeErrorsJSON `json:"-"`
}

// previewNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PreviewNewResponseEnvelopeErrors]
type previewNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PreviewNewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    previewNewResponseEnvelopeMessagesJSON `json:"-"`
}

// previewNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PreviewNewResponseEnvelopeMessages]
type previewNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PreviewNewResponseEnvelopeSuccess bool

const (
	PreviewNewResponseEnvelopeSuccessTrue PreviewNewResponseEnvelopeSuccess = true
)

type PreviewDeleteResponseEnvelope struct {
	Errors   []PreviewDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PreviewDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   PreviewDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PreviewDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    previewDeleteResponseEnvelopeJSON    `json:"-"`
}

// previewDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PreviewDeleteResponseEnvelope]
type previewDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PreviewDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    previewDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// previewDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PreviewDeleteResponseEnvelopeErrors]
type previewDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PreviewDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    previewDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// previewDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PreviewDeleteResponseEnvelopeMessages]
type previewDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PreviewDeleteResponseEnvelopeSuccess bool

const (
	PreviewDeleteResponseEnvelopeSuccessTrue PreviewDeleteResponseEnvelopeSuccess = true
)

type PreviewGetResponseEnvelope struct {
	Errors   []PreviewGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PreviewGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PreviewGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PreviewGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    previewGetResponseEnvelopeJSON    `json:"-"`
}

// previewGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PreviewGetResponseEnvelope]
type previewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PreviewGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    previewGetResponseEnvelopeErrorsJSON `json:"-"`
}

// previewGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PreviewGetResponseEnvelopeErrors]
type previewGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PreviewGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    previewGetResponseEnvelopeMessagesJSON `json:"-"`
}

// previewGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PreviewGetResponseEnvelopeMessages]
type previewGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PreviewGetResponseEnvelopeSuccess bool

const (
	PreviewGetResponseEnvelopeSuccessTrue PreviewGetResponseEnvelopeSuccess = true
)
