// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthchecks

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// HealthcheckService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewHealthcheckService] method
// instead.
type HealthcheckService struct {
	Options  []option.RequestOption
	Previews *PreviewService
}

// NewHealthcheckService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHealthcheckService(opts ...option.RequestOption) (r *HealthcheckService) {
	r = &HealthcheckService{}
	r.Options = opts
	r.Previews = NewPreviewService(opts...)
	return
}

// Create a new health check.
func (r *HealthcheckService) New(ctx context.Context, params HealthcheckNewParams, opts ...option.RequestOption) (res *Healthcheck, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a configured health check.
func (r *HealthcheckService) Update(ctx context.Context, healthcheckID string, params HealthcheckUpdateParams, opts ...option.RequestOption) (res *Healthcheck, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/%s", params.ZoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured health checks.
func (r *HealthcheckService) List(ctx context.Context, query HealthcheckListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Healthcheck], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/healthchecks", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List configured health checks.
func (r *HealthcheckService) ListAutoPaging(ctx context.Context, query HealthcheckListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Healthcheck] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a health check.
func (r *HealthcheckService) Delete(ctx context.Context, healthcheckID string, params HealthcheckDeleteParams, opts ...option.RequestOption) (res *HealthcheckDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/%s", params.ZoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patch a configured health check.
func (r *HealthcheckService) Edit(ctx context.Context, healthcheckID string, params HealthcheckEditParams, opts ...option.RequestOption) (res *Healthcheck, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/%s", params.ZoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured health check.
func (r *HealthcheckService) Get(ctx context.Context, healthcheckID string, query HealthcheckGetParams, opts ...option.RequestOption) (res *Healthcheck, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/%s", query.ZoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Healthcheck struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthcheckCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig HTTPConfiguration `json:"http_config,nullable"`
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
	Status HealthcheckStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TCPConfig TCPConfiguration `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string          `json:"type"`
	JSON healthcheckJSON `json:"-"`
}

// healthcheckJSON contains the JSON metadata for the struct [Healthcheck]
type healthcheckJSON struct {
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
	TCPConfig            apijson.Field
	Timeout              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Healthcheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckCheckRegion string

const (
	HealthcheckCheckRegionWnam       HealthcheckCheckRegion = "WNAM"
	HealthcheckCheckRegionEnam       HealthcheckCheckRegion = "ENAM"
	HealthcheckCheckRegionWeu        HealthcheckCheckRegion = "WEU"
	HealthcheckCheckRegionEeu        HealthcheckCheckRegion = "EEU"
	HealthcheckCheckRegionNsam       HealthcheckCheckRegion = "NSAM"
	HealthcheckCheckRegionSsam       HealthcheckCheckRegion = "SSAM"
	HealthcheckCheckRegionOc         HealthcheckCheckRegion = "OC"
	HealthcheckCheckRegionMe         HealthcheckCheckRegion = "ME"
	HealthcheckCheckRegionNaf        HealthcheckCheckRegion = "NAF"
	HealthcheckCheckRegionSaf        HealthcheckCheckRegion = "SAF"
	HealthcheckCheckRegionIn         HealthcheckCheckRegion = "IN"
	HealthcheckCheckRegionSeas       HealthcheckCheckRegion = "SEAS"
	HealthcheckCheckRegionNeas       HealthcheckCheckRegion = "NEAS"
	HealthcheckCheckRegionAllRegions HealthcheckCheckRegion = "ALL_REGIONS"
)

func (r HealthcheckCheckRegion) IsKnown() bool {
	switch r {
	case HealthcheckCheckRegionWnam, HealthcheckCheckRegionEnam, HealthcheckCheckRegionWeu, HealthcheckCheckRegionEeu, HealthcheckCheckRegionNsam, HealthcheckCheckRegionSsam, HealthcheckCheckRegionOc, HealthcheckCheckRegionMe, HealthcheckCheckRegionNaf, HealthcheckCheckRegionSaf, HealthcheckCheckRegionIn, HealthcheckCheckRegionSeas, HealthcheckCheckRegionNeas, HealthcheckCheckRegionAllRegions:
		return true
	}
	return false
}

// The current status of the origin server according to the health check.
type HealthcheckStatus string

const (
	HealthcheckStatusUnknown   HealthcheckStatus = "unknown"
	HealthcheckStatusHealthy   HealthcheckStatus = "healthy"
	HealthcheckStatusUnhealthy HealthcheckStatus = "unhealthy"
	HealthcheckStatusSuspended HealthcheckStatus = "suspended"
)

func (r HealthcheckStatus) IsKnown() bool {
	switch r {
	case HealthcheckStatusUnknown, HealthcheckStatusHealthy, HealthcheckStatusUnhealthy, HealthcheckStatusSuspended:
		return true
	}
	return false
}

// Parameters specific to an HTTP or HTTPS health check.
type HTTPConfiguration struct {
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
	Method HTTPConfigurationMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                 `json:"port"`
	JSON httpConfigurationJSON `json:"-"`
}

// httpConfigurationJSON contains the JSON metadata for the struct
// [HTTPConfiguration]
type httpConfigurationJSON struct {
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

func (r *HTTPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpConfigurationJSON) RawJSON() string {
	return r.raw
}

// The HTTP method to use for the health check.
type HTTPConfigurationMethod string

const (
	HTTPConfigurationMethodGet  HTTPConfigurationMethod = "GET"
	HTTPConfigurationMethodHead HTTPConfigurationMethod = "HEAD"
)

func (r HTTPConfigurationMethod) IsKnown() bool {
	switch r {
	case HTTPConfigurationMethodGet, HTTPConfigurationMethodHead:
		return true
	}
	return false
}

// Parameters specific to an HTTP or HTTPS health check.
type HTTPConfigurationParam struct {
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
	Method param.Field[HTTPConfigurationMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r HTTPConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Parameters specific to TCP health check.
type TCPConfiguration struct {
	// The TCP connection method to use for the health check.
	Method TCPConfigurationMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                `json:"port"`
	JSON tcpConfigurationJSON `json:"-"`
}

// tcpConfigurationJSON contains the JSON metadata for the struct
// [TCPConfiguration]
type tcpConfigurationJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpConfigurationJSON) RawJSON() string {
	return r.raw
}

// The TCP connection method to use for the health check.
type TCPConfigurationMethod string

const (
	TCPConfigurationMethodConnectionEstablished TCPConfigurationMethod = "connection_established"
)

func (r TCPConfigurationMethod) IsKnown() bool {
	switch r {
	case TCPConfigurationMethodConnectionEstablished:
		return true
	}
	return false
}

// Parameters specific to TCP health check.
type TCPConfigurationParam struct {
	// The TCP connection method to use for the health check.
	Method param.Field[TCPConfigurationMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r TCPConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

func (r healthcheckDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type HealthcheckNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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
	HTTPConfig param.Field[HTTPConfigurationParam] `json:"http_config"`
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
	TCPConfig param.Field[TCPConfigurationParam] `json:"tcp_config"`
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

func (r HealthcheckNewParamsCheckRegion) IsKnown() bool {
	switch r {
	case HealthcheckNewParamsCheckRegionWnam, HealthcheckNewParamsCheckRegionEnam, HealthcheckNewParamsCheckRegionWeu, HealthcheckNewParamsCheckRegionEeu, HealthcheckNewParamsCheckRegionNsam, HealthcheckNewParamsCheckRegionSsam, HealthcheckNewParamsCheckRegionOc, HealthcheckNewParamsCheckRegionMe, HealthcheckNewParamsCheckRegionNaf, HealthcheckNewParamsCheckRegionSaf, HealthcheckNewParamsCheckRegionIn, HealthcheckNewParamsCheckRegionSeas, HealthcheckNewParamsCheckRegionNeas, HealthcheckNewParamsCheckRegionAllRegions:
		return true
	}
	return false
}

type HealthcheckNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Healthcheck                                               `json:"result,required"`
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

func (r healthcheckNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HealthcheckNewResponseEnvelopeSuccess bool

const (
	HealthcheckNewResponseEnvelopeSuccessTrue HealthcheckNewResponseEnvelopeSuccess = true
)

func (r HealthcheckNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HealthcheckNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HealthcheckUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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
	HTTPConfig param.Field[HTTPConfigurationParam] `json:"http_config"`
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
	TCPConfig param.Field[TCPConfigurationParam] `json:"tcp_config"`
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

func (r HealthcheckUpdateParamsCheckRegion) IsKnown() bool {
	switch r {
	case HealthcheckUpdateParamsCheckRegionWnam, HealthcheckUpdateParamsCheckRegionEnam, HealthcheckUpdateParamsCheckRegionWeu, HealthcheckUpdateParamsCheckRegionEeu, HealthcheckUpdateParamsCheckRegionNsam, HealthcheckUpdateParamsCheckRegionSsam, HealthcheckUpdateParamsCheckRegionOc, HealthcheckUpdateParamsCheckRegionMe, HealthcheckUpdateParamsCheckRegionNaf, HealthcheckUpdateParamsCheckRegionSaf, HealthcheckUpdateParamsCheckRegionIn, HealthcheckUpdateParamsCheckRegionSeas, HealthcheckUpdateParamsCheckRegionNeas, HealthcheckUpdateParamsCheckRegionAllRegions:
		return true
	}
	return false
}

type HealthcheckUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Healthcheck                                               `json:"result,required"`
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

func (r healthcheckUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HealthcheckUpdateResponseEnvelopeSuccess bool

const (
	HealthcheckUpdateResponseEnvelopeSuccessTrue HealthcheckUpdateResponseEnvelopeSuccess = true
)

func (r HealthcheckUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HealthcheckUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HealthcheckListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HealthcheckDeleteParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r HealthcheckDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type HealthcheckDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   HealthcheckDeleteResponse                                 `json:"result,required"`
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

func (r healthcheckDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HealthcheckDeleteResponseEnvelopeSuccess bool

const (
	HealthcheckDeleteResponseEnvelopeSuccessTrue HealthcheckDeleteResponseEnvelopeSuccess = true
)

func (r HealthcheckDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HealthcheckDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HealthcheckEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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
	HTTPConfig param.Field[HTTPConfigurationParam] `json:"http_config"`
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
	TCPConfig param.Field[TCPConfigurationParam] `json:"tcp_config"`
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

func (r HealthcheckEditParamsCheckRegion) IsKnown() bool {
	switch r {
	case HealthcheckEditParamsCheckRegionWnam, HealthcheckEditParamsCheckRegionEnam, HealthcheckEditParamsCheckRegionWeu, HealthcheckEditParamsCheckRegionEeu, HealthcheckEditParamsCheckRegionNsam, HealthcheckEditParamsCheckRegionSsam, HealthcheckEditParamsCheckRegionOc, HealthcheckEditParamsCheckRegionMe, HealthcheckEditParamsCheckRegionNaf, HealthcheckEditParamsCheckRegionSaf, HealthcheckEditParamsCheckRegionIn, HealthcheckEditParamsCheckRegionSeas, HealthcheckEditParamsCheckRegionNeas, HealthcheckEditParamsCheckRegionAllRegions:
		return true
	}
	return false
}

type HealthcheckEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Healthcheck                                               `json:"result,required"`
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

func (r healthcheckEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HealthcheckEditResponseEnvelopeSuccess bool

const (
	HealthcheckEditResponseEnvelopeSuccessTrue HealthcheckEditResponseEnvelopeSuccess = true
)

func (r HealthcheckEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HealthcheckEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HealthcheckGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HealthcheckGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Healthcheck                                               `json:"result,required"`
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

func (r healthcheckGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HealthcheckGetResponseEnvelopeSuccess bool

const (
	HealthcheckGetResponseEnvelopeSuccessTrue HealthcheckGetResponseEnvelopeSuccess = true
)

func (r HealthcheckGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HealthcheckGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
