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

// Create a new health check.
func (r *HealthcheckService) New(ctx context.Context, zoneIdentifier string, body HealthcheckNewParams, opts ...option.RequestOption) (res *HealthchecksHealthchecks, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a configured health check.
func (r *HealthcheckService) Update(ctx context.Context, zoneIdentifier string, identifier string, body HealthcheckUpdateParams, opts ...option.RequestOption) (res *HealthchecksHealthchecks, err error) {
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

// List configured health checks.
func (r *HealthcheckService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]HealthchecksHealthchecks, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckListResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

// Patch a configured health check.
func (r *HealthcheckService) Edit(ctx context.Context, zoneIdentifier string, identifier string, body HealthcheckEditParams, opts ...option.RequestOption) (res *HealthchecksHealthchecks, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured health check.
func (r *HealthcheckService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *HealthchecksHealthchecks, err error) {
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

type HealthchecksHealthchecks struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthchecksHealthchecksCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig HealthchecksHealthchecksHTTPConfig `json:"http_config,nullable"`
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
	Status HealthchecksHealthchecksStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthchecksHealthchecksTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                       `json:"type"`
	JSON healthchecksHealthchecksJSON `json:"-"`
}

// healthchecksHealthchecksJSON contains the JSON metadata for the struct
// [HealthchecksHealthchecks]
type healthchecksHealthchecksJSON struct {
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

func (r *HealthchecksHealthchecks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthchecksHealthchecksCheckRegion string

const (
	HealthchecksHealthchecksCheckRegionWnam       HealthchecksHealthchecksCheckRegion = "WNAM"
	HealthchecksHealthchecksCheckRegionEnam       HealthchecksHealthchecksCheckRegion = "ENAM"
	HealthchecksHealthchecksCheckRegionWeu        HealthchecksHealthchecksCheckRegion = "WEU"
	HealthchecksHealthchecksCheckRegionEeu        HealthchecksHealthchecksCheckRegion = "EEU"
	HealthchecksHealthchecksCheckRegionNsam       HealthchecksHealthchecksCheckRegion = "NSAM"
	HealthchecksHealthchecksCheckRegionSsam       HealthchecksHealthchecksCheckRegion = "SSAM"
	HealthchecksHealthchecksCheckRegionOc         HealthchecksHealthchecksCheckRegion = "OC"
	HealthchecksHealthchecksCheckRegionMe         HealthchecksHealthchecksCheckRegion = "ME"
	HealthchecksHealthchecksCheckRegionNaf        HealthchecksHealthchecksCheckRegion = "NAF"
	HealthchecksHealthchecksCheckRegionSaf        HealthchecksHealthchecksCheckRegion = "SAF"
	HealthchecksHealthchecksCheckRegionIn         HealthchecksHealthchecksCheckRegion = "IN"
	HealthchecksHealthchecksCheckRegionSeas       HealthchecksHealthchecksCheckRegion = "SEAS"
	HealthchecksHealthchecksCheckRegionNeas       HealthchecksHealthchecksCheckRegion = "NEAS"
	HealthchecksHealthchecksCheckRegionAllRegions HealthchecksHealthchecksCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthchecksHealthchecksHTTPConfig struct {
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
	Method HealthchecksHealthchecksHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                  `json:"port"`
	JSON healthchecksHealthchecksHTTPConfigJSON `json:"-"`
}

// healthchecksHealthchecksHTTPConfigJSON contains the JSON metadata for the struct
// [HealthchecksHealthchecksHTTPConfig]
type healthchecksHealthchecksHTTPConfigJSON struct {
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

func (r *HealthchecksHealthchecksHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type HealthchecksHealthchecksHTTPConfigMethod string

const (
	HealthchecksHealthchecksHTTPConfigMethodGet  HealthchecksHealthchecksHTTPConfigMethod = "GET"
	HealthchecksHealthchecksHTTPConfigMethodHead HealthchecksHealthchecksHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type HealthchecksHealthchecksStatus string

const (
	HealthchecksHealthchecksStatusUnknown   HealthchecksHealthchecksStatus = "unknown"
	HealthchecksHealthchecksStatusHealthy   HealthchecksHealthchecksStatus = "healthy"
	HealthchecksHealthchecksStatusUnhealthy HealthchecksHealthchecksStatus = "unhealthy"
	HealthchecksHealthchecksStatusSuspended HealthchecksHealthchecksStatus = "suspended"
)

// Parameters specific to TCP health check.
type HealthchecksHealthchecksTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthchecksHealthchecksTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                 `json:"port"`
	JSON healthchecksHealthchecksTcpConfigJSON `json:"-"`
}

// healthchecksHealthchecksTcpConfigJSON contains the JSON metadata for the struct
// [HealthchecksHealthchecksTcpConfig]
type healthchecksHealthchecksTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthchecksHealthchecksTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type HealthchecksHealthchecksTcpConfigMethod string

const (
	HealthchecksHealthchecksTcpConfigMethodConnectionEstablished HealthchecksHealthchecksTcpConfigMethod = "connection_established"
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

type HealthcheckNewParams struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]HealthcheckNewParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[HealthcheckNewParamsHTTPConfig] `json:"http_config"`
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
	TcpConfig param.Field[HealthcheckNewParamsTcpConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r HealthcheckNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckNewParamsCheckRegion string

const (
	HealthcheckNewParamsCheckRegionWnam       HealthcheckNewParamsCheckRegion = "WNAM"
	HealthcheckNewParamsCheckRegionEnam       HealthcheckNewParamsCheckRegion = "ENAM"
	HealthcheckNewParamsCheckRegionWeu        HealthcheckNewParamsCheckRegion = "WEU"
	HealthcheckNewParamsCheckRegionEeu        HealthcheckNewParamsCheckRegion = "EEU"
	HealthcheckNewParamsCheckRegionNsam       HealthcheckNewParamsCheckRegion = "NSAM"
	HealthcheckNewParamsCheckRegionSsam       HealthcheckNewParamsCheckRegion = "SSAM"
	HealthcheckNewParamsCheckRegionOc         HealthcheckNewParamsCheckRegion = "OC"
	HealthcheckNewParamsCheckRegionMe         HealthcheckNewParamsCheckRegion = "ME"
	HealthcheckNewParamsCheckRegionNaf        HealthcheckNewParamsCheckRegion = "NAF"
	HealthcheckNewParamsCheckRegionSaf        HealthcheckNewParamsCheckRegion = "SAF"
	HealthcheckNewParamsCheckRegionIn         HealthcheckNewParamsCheckRegion = "IN"
	HealthcheckNewParamsCheckRegionSeas       HealthcheckNewParamsCheckRegion = "SEAS"
	HealthcheckNewParamsCheckRegionNeas       HealthcheckNewParamsCheckRegion = "NEAS"
	HealthcheckNewParamsCheckRegionAllRegions HealthcheckNewParamsCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckNewParamsHTTPConfig struct {
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
	Method param.Field[HealthcheckNewParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckNewParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type HealthcheckNewParamsHTTPConfigMethod string

const (
	HealthcheckNewParamsHTTPConfigMethodGet  HealthcheckNewParamsHTTPConfigMethod = "GET"
	HealthcheckNewParamsHTTPConfigMethodHead HealthcheckNewParamsHTTPConfigMethod = "HEAD"
)

// Parameters specific to TCP health check.
type HealthcheckNewParamsTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[HealthcheckNewParamsTcpConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckNewParamsTcpConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type HealthcheckNewParamsTcpConfigMethod string

const (
	HealthcheckNewParamsTcpConfigMethodConnectionEstablished HealthcheckNewParamsTcpConfigMethod = "connection_established"
)

type HealthcheckNewResponseEnvelope struct {
	Errors   []HealthcheckNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckNewResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthchecksHealthchecks                 `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckNewResponseEnvelopeJSON    `json:"-"`
}

// healthcheckNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [HealthcheckNewResponseEnvelope]
type healthcheckNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    healthcheckNewResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HealthcheckNewResponseEnvelopeErrors]
type healthcheckNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    healthcheckNewResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HealthcheckNewResponseEnvelopeMessages]
type healthcheckNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckNewResponseEnvelopeSuccess bool

const (
	HealthcheckNewResponseEnvelopeSuccessTrue HealthcheckNewResponseEnvelopeSuccess = true
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
	Result   HealthchecksHealthchecks                    `json:"result,required"`
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

type HealthcheckListResponseEnvelope struct {
	Errors   []HealthcheckListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckListResponseEnvelopeMessages `json:"messages,required"`
	Result   []HealthchecksHealthchecks                `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    HealthcheckListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo HealthcheckListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       healthcheckListResponseEnvelopeJSON       `json:"-"`
}

// healthcheckListResponseEnvelopeJSON contains the JSON metadata for the struct
// [HealthcheckListResponseEnvelope]
type healthcheckListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    healthcheckListResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HealthcheckListResponseEnvelopeErrors]
type healthcheckListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    healthcheckListResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HealthcheckListResponseEnvelopeMessages]
type healthcheckListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckListResponseEnvelopeSuccess bool

const (
	HealthcheckListResponseEnvelopeSuccessTrue HealthcheckListResponseEnvelopeSuccess = true
)

type HealthcheckListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       healthcheckListResponseEnvelopeResultInfoJSON `json:"-"`
}

// healthcheckListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [HealthcheckListResponseEnvelopeResultInfo]
type healthcheckListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type HealthcheckEditParams struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]HealthcheckEditParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[HealthcheckEditParamsHTTPConfig] `json:"http_config"`
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
	TcpConfig param.Field[HealthcheckEditParamsTcpConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r HealthcheckEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckEditParamsCheckRegion string

const (
	HealthcheckEditParamsCheckRegionWnam       HealthcheckEditParamsCheckRegion = "WNAM"
	HealthcheckEditParamsCheckRegionEnam       HealthcheckEditParamsCheckRegion = "ENAM"
	HealthcheckEditParamsCheckRegionWeu        HealthcheckEditParamsCheckRegion = "WEU"
	HealthcheckEditParamsCheckRegionEeu        HealthcheckEditParamsCheckRegion = "EEU"
	HealthcheckEditParamsCheckRegionNsam       HealthcheckEditParamsCheckRegion = "NSAM"
	HealthcheckEditParamsCheckRegionSsam       HealthcheckEditParamsCheckRegion = "SSAM"
	HealthcheckEditParamsCheckRegionOc         HealthcheckEditParamsCheckRegion = "OC"
	HealthcheckEditParamsCheckRegionMe         HealthcheckEditParamsCheckRegion = "ME"
	HealthcheckEditParamsCheckRegionNaf        HealthcheckEditParamsCheckRegion = "NAF"
	HealthcheckEditParamsCheckRegionSaf        HealthcheckEditParamsCheckRegion = "SAF"
	HealthcheckEditParamsCheckRegionIn         HealthcheckEditParamsCheckRegion = "IN"
	HealthcheckEditParamsCheckRegionSeas       HealthcheckEditParamsCheckRegion = "SEAS"
	HealthcheckEditParamsCheckRegionNeas       HealthcheckEditParamsCheckRegion = "NEAS"
	HealthcheckEditParamsCheckRegionAllRegions HealthcheckEditParamsCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckEditParamsHTTPConfig struct {
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
	Method param.Field[HealthcheckEditParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckEditParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type HealthcheckEditParamsHTTPConfigMethod string

const (
	HealthcheckEditParamsHTTPConfigMethodGet  HealthcheckEditParamsHTTPConfigMethod = "GET"
	HealthcheckEditParamsHTTPConfigMethodHead HealthcheckEditParamsHTTPConfigMethod = "HEAD"
)

// Parameters specific to TCP health check.
type HealthcheckEditParamsTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[HealthcheckEditParamsTcpConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckEditParamsTcpConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type HealthcheckEditParamsTcpConfigMethod string

const (
	HealthcheckEditParamsTcpConfigMethodConnectionEstablished HealthcheckEditParamsTcpConfigMethod = "connection_established"
)

type HealthcheckEditResponseEnvelope struct {
	Errors   []HealthcheckEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckEditResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthchecksHealthchecks                  `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckEditResponseEnvelopeJSON    `json:"-"`
}

// healthcheckEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [HealthcheckEditResponseEnvelope]
type healthcheckEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckEditResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    healthcheckEditResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HealthcheckEditResponseEnvelopeErrors]
type healthcheckEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckEditResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    healthcheckEditResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HealthcheckEditResponseEnvelopeMessages]
type healthcheckEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckEditResponseEnvelopeSuccess bool

const (
	HealthcheckEditResponseEnvelopeSuccessTrue HealthcheckEditResponseEnvelopeSuccess = true
)

type HealthcheckGetResponseEnvelope struct {
	Errors   []HealthcheckGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckGetResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthchecksHealthchecks                 `json:"result,required"`
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
