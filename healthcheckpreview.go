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
func (r *HealthcheckPreviewService) New(ctx context.Context, zoneIdentifier string, body HealthcheckPreviewNewParams, opts ...option.RequestOption) (res *HealthchecksHealthchecks, err error) {
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
func (r *HealthcheckPreviewService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *HealthchecksHealthchecks, err error) {
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
	Result   HealthchecksHealthchecks                        `json:"result,required"`
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

// Whether the API call was successful
type HealthcheckPreviewDeleteResponseEnvelopeSuccess bool

const (
	HealthcheckPreviewDeleteResponseEnvelopeSuccessTrue HealthcheckPreviewDeleteResponseEnvelopeSuccess = true
)

type HealthcheckPreviewGetResponseEnvelope struct {
	Errors   []HealthcheckPreviewGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckPreviewGetResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthchecksHealthchecks                        `json:"result,required"`
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

// Whether the API call was successful
type HealthcheckPreviewGetResponseEnvelopeSuccess bool

const (
	HealthcheckPreviewGetResponseEnvelopeSuccessTrue HealthcheckPreviewGetResponseEnvelopeSuccess = true
)
