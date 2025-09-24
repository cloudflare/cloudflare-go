// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package smart_shield

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// SmartShieldService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSmartShieldService] method instead.
type SmartShieldService struct {
	Options []option.RequestOption
}

// NewSmartShieldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSmartShieldService(opts ...option.RequestOption) (r *SmartShieldService) {
	r = &SmartShieldService{}
	r.Options = opts
	return
}

// Set Smart Shield Settings.
func (r *SmartShieldService) Update(ctx context.Context, params SmartShieldUpdateParams, opts ...option.RequestOption) (res *SmartShieldUpdateResponse, err error) {
	var env SmartShieldUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/smart_shield", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a new health check.
func (r *SmartShieldService) NewHealthcheck(ctx context.Context, params SmartShieldNewHealthcheckParams, opts ...option.RequestOption) (res *SmartShieldNewHealthcheckResponse, err error) {
	var env SmartShieldNewHealthcheckResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/smart_shield/healthchecks", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a health check.
func (r *SmartShieldService) DeleteHealthcheck(ctx context.Context, healthcheckID string, body SmartShieldDeleteHealthcheckParams, opts ...option.RequestOption) (res *SmartShieldDeleteHealthcheckResponse, err error) {
	var env SmartShieldDeleteHealthcheckResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if healthcheckID == "" {
		err = errors.New("missing required healthcheck_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/smart_shield/healthchecks/%s", body.ZoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patch a configured health check.
func (r *SmartShieldService) EditHealthcheck(ctx context.Context, healthcheckID string, params SmartShieldEditHealthcheckParams, opts ...option.RequestOption) (res *SmartShieldEditHealthcheckResponse, err error) {
	var env SmartShieldEditHealthcheckResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if healthcheckID == "" {
		err = errors.New("missing required healthcheck_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/smart_shield/healthchecks/%s", params.ZoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve Smart Shield Settings.
func (r *SmartShieldService) Get(ctx context.Context, query SmartShieldGetParams, opts ...option.RequestOption) (res *SmartShieldGetResponse, err error) {
	var env SmartShieldGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/smart_shield", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured health check.
func (r *SmartShieldService) GetHealthcheck(ctx context.Context, healthcheckID string, query SmartShieldGetHealthcheckParams, opts ...option.RequestOption) (res *SmartShieldGetHealthcheckResponse, err error) {
	var env SmartShieldGetHealthcheckResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if healthcheckID == "" {
		err = errors.New("missing required healthcheck_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/smart_shield/healthchecks/%s", query.ZoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured health checks.
func (r *SmartShieldService) ListHealthchecks(ctx context.Context, params SmartShieldListHealthchecksParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SmartShieldListHealthchecksResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/smart_shield/healthchecks", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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
func (r *SmartShieldService) ListHealthchecksAutoPaging(ctx context.Context, params SmartShieldListHealthchecksParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SmartShieldListHealthchecksResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.ListHealthchecks(ctx, params, opts...))
}

// Update a configured health check.
func (r *SmartShieldService) UpdateHealthcheck(ctx context.Context, healthcheckID string, params SmartShieldUpdateHealthcheckParams, opts ...option.RequestOption) (res *SmartShieldUpdateHealthcheckResponse, err error) {
	var env SmartShieldUpdateHealthcheckResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if healthcheckID == "" {
		err = errors.New("missing required healthcheck_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/smart_shield/healthchecks/%s", params.ZoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A consolidated object containing settings from multiple APIs for partial
// updates.
type SmartShieldUpdateResponse struct {
	SmartTieredCache SmartShieldUpdateResponseSmartTieredCache `json:"smart_tiered_cache"`
	JSON             smartShieldUpdateResponseJSON             `json:"-"`
}

// smartShieldUpdateResponseJSON contains the JSON metadata for the struct
// [SmartShieldUpdateResponse]
type smartShieldUpdateResponseJSON struct {
	SmartTieredCache apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SmartShieldUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SmartShieldUpdateResponseSmartTieredCache struct {
	// The id of the Smart Tiered Cache setting.
	ID string `json:"id"`
	// Whether the setting is editable.
	Editable bool `json:"editable"`
	// The last time the setting was modified.
	ModifiedOn string `json:"modified_on"`
	// Specifies the enablement value of Tiered Cache.
	Value SmartShieldUpdateResponseSmartTieredCacheValue `json:"value"`
	JSON  smartShieldUpdateResponseSmartTieredCacheJSON  `json:"-"`
}

// smartShieldUpdateResponseSmartTieredCacheJSON contains the JSON metadata for the
// struct [SmartShieldUpdateResponseSmartTieredCache]
type smartShieldUpdateResponseSmartTieredCacheJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldUpdateResponseSmartTieredCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldUpdateResponseSmartTieredCacheJSON) RawJSON() string {
	return r.raw
}

// Specifies the enablement value of Tiered Cache.
type SmartShieldUpdateResponseSmartTieredCacheValue string

const (
	SmartShieldUpdateResponseSmartTieredCacheValueOn  SmartShieldUpdateResponseSmartTieredCacheValue = "on"
	SmartShieldUpdateResponseSmartTieredCacheValueOff SmartShieldUpdateResponseSmartTieredCacheValue = "off"
)

func (r SmartShieldUpdateResponseSmartTieredCacheValue) IsKnown() bool {
	switch r {
	case SmartShieldUpdateResponseSmartTieredCacheValueOn, SmartShieldUpdateResponseSmartTieredCacheValueOff:
		return true
	}
	return false
}

type SmartShieldNewHealthcheckResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []SmartShieldNewHealthcheckResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig SmartShieldNewHealthcheckResponseHTTPConfig `json:"http_config,nullable"`
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
	Status SmartShieldNewHealthcheckResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TCPConfig SmartShieldNewHealthcheckResponseTCPConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                                `json:"type"`
	JSON smartShieldNewHealthcheckResponseJSON `json:"-"`
}

// smartShieldNewHealthcheckResponseJSON contains the JSON metadata for the struct
// [SmartShieldNewHealthcheckResponse]
type smartShieldNewHealthcheckResponseJSON struct {
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

func (r *SmartShieldNewHealthcheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldNewHealthcheckResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type SmartShieldNewHealthcheckResponseCheckRegion string

const (
	SmartShieldNewHealthcheckResponseCheckRegionWnam       SmartShieldNewHealthcheckResponseCheckRegion = "WNAM"
	SmartShieldNewHealthcheckResponseCheckRegionEnam       SmartShieldNewHealthcheckResponseCheckRegion = "ENAM"
	SmartShieldNewHealthcheckResponseCheckRegionWeu        SmartShieldNewHealthcheckResponseCheckRegion = "WEU"
	SmartShieldNewHealthcheckResponseCheckRegionEeu        SmartShieldNewHealthcheckResponseCheckRegion = "EEU"
	SmartShieldNewHealthcheckResponseCheckRegionNsam       SmartShieldNewHealthcheckResponseCheckRegion = "NSAM"
	SmartShieldNewHealthcheckResponseCheckRegionSsam       SmartShieldNewHealthcheckResponseCheckRegion = "SSAM"
	SmartShieldNewHealthcheckResponseCheckRegionOc         SmartShieldNewHealthcheckResponseCheckRegion = "OC"
	SmartShieldNewHealthcheckResponseCheckRegionMe         SmartShieldNewHealthcheckResponseCheckRegion = "ME"
	SmartShieldNewHealthcheckResponseCheckRegionNaf        SmartShieldNewHealthcheckResponseCheckRegion = "NAF"
	SmartShieldNewHealthcheckResponseCheckRegionSaf        SmartShieldNewHealthcheckResponseCheckRegion = "SAF"
	SmartShieldNewHealthcheckResponseCheckRegionIn         SmartShieldNewHealthcheckResponseCheckRegion = "IN"
	SmartShieldNewHealthcheckResponseCheckRegionSeas       SmartShieldNewHealthcheckResponseCheckRegion = "SEAS"
	SmartShieldNewHealthcheckResponseCheckRegionNeas       SmartShieldNewHealthcheckResponseCheckRegion = "NEAS"
	SmartShieldNewHealthcheckResponseCheckRegionAllRegions SmartShieldNewHealthcheckResponseCheckRegion = "ALL_REGIONS"
)

func (r SmartShieldNewHealthcheckResponseCheckRegion) IsKnown() bool {
	switch r {
	case SmartShieldNewHealthcheckResponseCheckRegionWnam, SmartShieldNewHealthcheckResponseCheckRegionEnam, SmartShieldNewHealthcheckResponseCheckRegionWeu, SmartShieldNewHealthcheckResponseCheckRegionEeu, SmartShieldNewHealthcheckResponseCheckRegionNsam, SmartShieldNewHealthcheckResponseCheckRegionSsam, SmartShieldNewHealthcheckResponseCheckRegionOc, SmartShieldNewHealthcheckResponseCheckRegionMe, SmartShieldNewHealthcheckResponseCheckRegionNaf, SmartShieldNewHealthcheckResponseCheckRegionSaf, SmartShieldNewHealthcheckResponseCheckRegionIn, SmartShieldNewHealthcheckResponseCheckRegionSeas, SmartShieldNewHealthcheckResponseCheckRegionNeas, SmartShieldNewHealthcheckResponseCheckRegionAllRegions:
		return true
	}
	return false
}

// Parameters specific to an HTTP or HTTPS health check.
type SmartShieldNewHealthcheckResponseHTTPConfig struct {
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
	Header map[string][]string `json:"header,nullable"`
	// The HTTP method to use for the health check.
	Method SmartShieldNewHealthcheckResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                           `json:"port"`
	JSON smartShieldNewHealthcheckResponseHTTPConfigJSON `json:"-"`
}

// smartShieldNewHealthcheckResponseHTTPConfigJSON contains the JSON metadata for
// the struct [SmartShieldNewHealthcheckResponseHTTPConfig]
type smartShieldNewHealthcheckResponseHTTPConfigJSON struct {
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

func (r *SmartShieldNewHealthcheckResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldNewHealthcheckResponseHTTPConfigJSON) RawJSON() string {
	return r.raw
}

// The HTTP method to use for the health check.
type SmartShieldNewHealthcheckResponseHTTPConfigMethod string

const (
	SmartShieldNewHealthcheckResponseHTTPConfigMethodGet  SmartShieldNewHealthcheckResponseHTTPConfigMethod = "GET"
	SmartShieldNewHealthcheckResponseHTTPConfigMethodHead SmartShieldNewHealthcheckResponseHTTPConfigMethod = "HEAD"
)

func (r SmartShieldNewHealthcheckResponseHTTPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldNewHealthcheckResponseHTTPConfigMethodGet, SmartShieldNewHealthcheckResponseHTTPConfigMethodHead:
		return true
	}
	return false
}

// The current status of the origin server according to the health check.
type SmartShieldNewHealthcheckResponseStatus string

const (
	SmartShieldNewHealthcheckResponseStatusUnknown   SmartShieldNewHealthcheckResponseStatus = "unknown"
	SmartShieldNewHealthcheckResponseStatusHealthy   SmartShieldNewHealthcheckResponseStatus = "healthy"
	SmartShieldNewHealthcheckResponseStatusUnhealthy SmartShieldNewHealthcheckResponseStatus = "unhealthy"
	SmartShieldNewHealthcheckResponseStatusSuspended SmartShieldNewHealthcheckResponseStatus = "suspended"
)

func (r SmartShieldNewHealthcheckResponseStatus) IsKnown() bool {
	switch r {
	case SmartShieldNewHealthcheckResponseStatusUnknown, SmartShieldNewHealthcheckResponseStatusHealthy, SmartShieldNewHealthcheckResponseStatusUnhealthy, SmartShieldNewHealthcheckResponseStatusSuspended:
		return true
	}
	return false
}

// Parameters specific to TCP health check.
type SmartShieldNewHealthcheckResponseTCPConfig struct {
	// The TCP connection method to use for the health check.
	Method SmartShieldNewHealthcheckResponseTCPConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                          `json:"port"`
	JSON smartShieldNewHealthcheckResponseTCPConfigJSON `json:"-"`
}

// smartShieldNewHealthcheckResponseTCPConfigJSON contains the JSON metadata for
// the struct [SmartShieldNewHealthcheckResponseTCPConfig]
type smartShieldNewHealthcheckResponseTCPConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldNewHealthcheckResponseTCPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldNewHealthcheckResponseTCPConfigJSON) RawJSON() string {
	return r.raw
}

// The TCP connection method to use for the health check.
type SmartShieldNewHealthcheckResponseTCPConfigMethod string

const (
	SmartShieldNewHealthcheckResponseTCPConfigMethodConnectionEstablished SmartShieldNewHealthcheckResponseTCPConfigMethod = "connection_established"
)

func (r SmartShieldNewHealthcheckResponseTCPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldNewHealthcheckResponseTCPConfigMethodConnectionEstablished:
		return true
	}
	return false
}

type SmartShieldDeleteHealthcheckResponse struct {
	// Identifier.
	ID   string                                   `json:"id"`
	JSON smartShieldDeleteHealthcheckResponseJSON `json:"-"`
}

// smartShieldDeleteHealthcheckResponseJSON contains the JSON metadata for the
// struct [SmartShieldDeleteHealthcheckResponse]
type smartShieldDeleteHealthcheckResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldDeleteHealthcheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldDeleteHealthcheckResponseJSON) RawJSON() string {
	return r.raw
}

type SmartShieldEditHealthcheckResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []SmartShieldEditHealthcheckResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig SmartShieldEditHealthcheckResponseHTTPConfig `json:"http_config,nullable"`
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
	Status SmartShieldEditHealthcheckResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TCPConfig SmartShieldEditHealthcheckResponseTCPConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                                 `json:"type"`
	JSON smartShieldEditHealthcheckResponseJSON `json:"-"`
}

// smartShieldEditHealthcheckResponseJSON contains the JSON metadata for the struct
// [SmartShieldEditHealthcheckResponse]
type smartShieldEditHealthcheckResponseJSON struct {
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

func (r *SmartShieldEditHealthcheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldEditHealthcheckResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type SmartShieldEditHealthcheckResponseCheckRegion string

const (
	SmartShieldEditHealthcheckResponseCheckRegionWnam       SmartShieldEditHealthcheckResponseCheckRegion = "WNAM"
	SmartShieldEditHealthcheckResponseCheckRegionEnam       SmartShieldEditHealthcheckResponseCheckRegion = "ENAM"
	SmartShieldEditHealthcheckResponseCheckRegionWeu        SmartShieldEditHealthcheckResponseCheckRegion = "WEU"
	SmartShieldEditHealthcheckResponseCheckRegionEeu        SmartShieldEditHealthcheckResponseCheckRegion = "EEU"
	SmartShieldEditHealthcheckResponseCheckRegionNsam       SmartShieldEditHealthcheckResponseCheckRegion = "NSAM"
	SmartShieldEditHealthcheckResponseCheckRegionSsam       SmartShieldEditHealthcheckResponseCheckRegion = "SSAM"
	SmartShieldEditHealthcheckResponseCheckRegionOc         SmartShieldEditHealthcheckResponseCheckRegion = "OC"
	SmartShieldEditHealthcheckResponseCheckRegionMe         SmartShieldEditHealthcheckResponseCheckRegion = "ME"
	SmartShieldEditHealthcheckResponseCheckRegionNaf        SmartShieldEditHealthcheckResponseCheckRegion = "NAF"
	SmartShieldEditHealthcheckResponseCheckRegionSaf        SmartShieldEditHealthcheckResponseCheckRegion = "SAF"
	SmartShieldEditHealthcheckResponseCheckRegionIn         SmartShieldEditHealthcheckResponseCheckRegion = "IN"
	SmartShieldEditHealthcheckResponseCheckRegionSeas       SmartShieldEditHealthcheckResponseCheckRegion = "SEAS"
	SmartShieldEditHealthcheckResponseCheckRegionNeas       SmartShieldEditHealthcheckResponseCheckRegion = "NEAS"
	SmartShieldEditHealthcheckResponseCheckRegionAllRegions SmartShieldEditHealthcheckResponseCheckRegion = "ALL_REGIONS"
)

func (r SmartShieldEditHealthcheckResponseCheckRegion) IsKnown() bool {
	switch r {
	case SmartShieldEditHealthcheckResponseCheckRegionWnam, SmartShieldEditHealthcheckResponseCheckRegionEnam, SmartShieldEditHealthcheckResponseCheckRegionWeu, SmartShieldEditHealthcheckResponseCheckRegionEeu, SmartShieldEditHealthcheckResponseCheckRegionNsam, SmartShieldEditHealthcheckResponseCheckRegionSsam, SmartShieldEditHealthcheckResponseCheckRegionOc, SmartShieldEditHealthcheckResponseCheckRegionMe, SmartShieldEditHealthcheckResponseCheckRegionNaf, SmartShieldEditHealthcheckResponseCheckRegionSaf, SmartShieldEditHealthcheckResponseCheckRegionIn, SmartShieldEditHealthcheckResponseCheckRegionSeas, SmartShieldEditHealthcheckResponseCheckRegionNeas, SmartShieldEditHealthcheckResponseCheckRegionAllRegions:
		return true
	}
	return false
}

// Parameters specific to an HTTP or HTTPS health check.
type SmartShieldEditHealthcheckResponseHTTPConfig struct {
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
	Header map[string][]string `json:"header,nullable"`
	// The HTTP method to use for the health check.
	Method SmartShieldEditHealthcheckResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                            `json:"port"`
	JSON smartShieldEditHealthcheckResponseHTTPConfigJSON `json:"-"`
}

// smartShieldEditHealthcheckResponseHTTPConfigJSON contains the JSON metadata for
// the struct [SmartShieldEditHealthcheckResponseHTTPConfig]
type smartShieldEditHealthcheckResponseHTTPConfigJSON struct {
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

func (r *SmartShieldEditHealthcheckResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldEditHealthcheckResponseHTTPConfigJSON) RawJSON() string {
	return r.raw
}

// The HTTP method to use for the health check.
type SmartShieldEditHealthcheckResponseHTTPConfigMethod string

const (
	SmartShieldEditHealthcheckResponseHTTPConfigMethodGet  SmartShieldEditHealthcheckResponseHTTPConfigMethod = "GET"
	SmartShieldEditHealthcheckResponseHTTPConfigMethodHead SmartShieldEditHealthcheckResponseHTTPConfigMethod = "HEAD"
)

func (r SmartShieldEditHealthcheckResponseHTTPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldEditHealthcheckResponseHTTPConfigMethodGet, SmartShieldEditHealthcheckResponseHTTPConfigMethodHead:
		return true
	}
	return false
}

// The current status of the origin server according to the health check.
type SmartShieldEditHealthcheckResponseStatus string

const (
	SmartShieldEditHealthcheckResponseStatusUnknown   SmartShieldEditHealthcheckResponseStatus = "unknown"
	SmartShieldEditHealthcheckResponseStatusHealthy   SmartShieldEditHealthcheckResponseStatus = "healthy"
	SmartShieldEditHealthcheckResponseStatusUnhealthy SmartShieldEditHealthcheckResponseStatus = "unhealthy"
	SmartShieldEditHealthcheckResponseStatusSuspended SmartShieldEditHealthcheckResponseStatus = "suspended"
)

func (r SmartShieldEditHealthcheckResponseStatus) IsKnown() bool {
	switch r {
	case SmartShieldEditHealthcheckResponseStatusUnknown, SmartShieldEditHealthcheckResponseStatusHealthy, SmartShieldEditHealthcheckResponseStatusUnhealthy, SmartShieldEditHealthcheckResponseStatusSuspended:
		return true
	}
	return false
}

// Parameters specific to TCP health check.
type SmartShieldEditHealthcheckResponseTCPConfig struct {
	// The TCP connection method to use for the health check.
	Method SmartShieldEditHealthcheckResponseTCPConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                           `json:"port"`
	JSON smartShieldEditHealthcheckResponseTCPConfigJSON `json:"-"`
}

// smartShieldEditHealthcheckResponseTCPConfigJSON contains the JSON metadata for
// the struct [SmartShieldEditHealthcheckResponseTCPConfig]
type smartShieldEditHealthcheckResponseTCPConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldEditHealthcheckResponseTCPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldEditHealthcheckResponseTCPConfigJSON) RawJSON() string {
	return r.raw
}

// The TCP connection method to use for the health check.
type SmartShieldEditHealthcheckResponseTCPConfigMethod string

const (
	SmartShieldEditHealthcheckResponseTCPConfigMethodConnectionEstablished SmartShieldEditHealthcheckResponseTCPConfigMethod = "connection_established"
)

func (r SmartShieldEditHealthcheckResponseTCPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldEditHealthcheckResponseTCPConfigMethodConnectionEstablished:
		return true
	}
	return false
}

// A consolidated object containing settings from multiple APIs for partial
// updates.
type SmartShieldGetResponse struct {
	CacheReserve SmartShieldGetResponseCacheReserve `json:"cache_reserve"`
	// The total number of health checks associated with the zone.
	HealthchecksCount   int64                                     `json:"healthchecks_count"`
	RegionalTieredCache SmartShieldGetResponseRegionalTieredCache `json:"regional_tiered_cache"`
	SmartRouting        SmartShieldGetResponseSmartRouting        `json:"smart_routing"`
	SmartTieredCache    SmartShieldGetResponseSmartTieredCache    `json:"smart_tiered_cache"`
	JSON                smartShieldGetResponseJSON                `json:"-"`
}

// smartShieldGetResponseJSON contains the JSON metadata for the struct
// [SmartShieldGetResponse]
type smartShieldGetResponseJSON struct {
	CacheReserve        apijson.Field
	HealthchecksCount   apijson.Field
	RegionalTieredCache apijson.Field
	SmartRouting        apijson.Field
	SmartTieredCache    apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SmartShieldGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldGetResponseJSON) RawJSON() string {
	return r.raw
}

type SmartShieldGetResponseCacheReserve struct {
	// The id of the Cache Reserve setting.
	ID string `json:"id"`
	// Whether the setting is editable.
	Editable bool `json:"editable"`
	// Specifies the enablement value of Cache Reserve.
	Value SmartShieldGetResponseCacheReserveValue `json:"value"`
	JSON  smartShieldGetResponseCacheReserveJSON  `json:"-"`
}

// smartShieldGetResponseCacheReserveJSON contains the JSON metadata for the struct
// [SmartShieldGetResponseCacheReserve]
type smartShieldGetResponseCacheReserveJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldGetResponseCacheReserve) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldGetResponseCacheReserveJSON) RawJSON() string {
	return r.raw
}

// Specifies the enablement value of Cache Reserve.
type SmartShieldGetResponseCacheReserveValue string

const (
	SmartShieldGetResponseCacheReserveValueOn  SmartShieldGetResponseCacheReserveValue = "on"
	SmartShieldGetResponseCacheReserveValueOff SmartShieldGetResponseCacheReserveValue = "off"
)

func (r SmartShieldGetResponseCacheReserveValue) IsKnown() bool {
	switch r {
	case SmartShieldGetResponseCacheReserveValueOn, SmartShieldGetResponseCacheReserveValueOff:
		return true
	}
	return false
}

type SmartShieldGetResponseRegionalTieredCache struct {
	// The id of the Regional Tiered Cache setting.
	ID string `json:"id"`
	// Whether the setting is editable.
	Editable bool `json:"editable"`
	// Specifies the enablement value of Cache Reserve.
	Value SmartShieldGetResponseRegionalTieredCacheValue `json:"value"`
	JSON  smartShieldGetResponseRegionalTieredCacheJSON  `json:"-"`
}

// smartShieldGetResponseRegionalTieredCacheJSON contains the JSON metadata for the
// struct [SmartShieldGetResponseRegionalTieredCache]
type smartShieldGetResponseRegionalTieredCacheJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldGetResponseRegionalTieredCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldGetResponseRegionalTieredCacheJSON) RawJSON() string {
	return r.raw
}

// Specifies the enablement value of Cache Reserve.
type SmartShieldGetResponseRegionalTieredCacheValue string

const (
	SmartShieldGetResponseRegionalTieredCacheValueOn  SmartShieldGetResponseRegionalTieredCacheValue = "on"
	SmartShieldGetResponseRegionalTieredCacheValueOff SmartShieldGetResponseRegionalTieredCacheValue = "off"
)

func (r SmartShieldGetResponseRegionalTieredCacheValue) IsKnown() bool {
	switch r {
	case SmartShieldGetResponseRegionalTieredCacheValueOn, SmartShieldGetResponseRegionalTieredCacheValueOff:
		return true
	}
	return false
}

type SmartShieldGetResponseSmartRouting struct {
	// The id of the Smart Routing setting.
	ID string `json:"id"`
	// Whether the setting is editable.
	Editable bool `json:"editable"`
	// Specifies the enablement value of Argo Smart Routing.
	Value SmartShieldGetResponseSmartRoutingValue `json:"value"`
	JSON  smartShieldGetResponseSmartRoutingJSON  `json:"-"`
}

// smartShieldGetResponseSmartRoutingJSON contains the JSON metadata for the struct
// [SmartShieldGetResponseSmartRouting]
type smartShieldGetResponseSmartRoutingJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldGetResponseSmartRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldGetResponseSmartRoutingJSON) RawJSON() string {
	return r.raw
}

// Specifies the enablement value of Argo Smart Routing.
type SmartShieldGetResponseSmartRoutingValue string

const (
	SmartShieldGetResponseSmartRoutingValueOn  SmartShieldGetResponseSmartRoutingValue = "on"
	SmartShieldGetResponseSmartRoutingValueOff SmartShieldGetResponseSmartRoutingValue = "off"
)

func (r SmartShieldGetResponseSmartRoutingValue) IsKnown() bool {
	switch r {
	case SmartShieldGetResponseSmartRoutingValueOn, SmartShieldGetResponseSmartRoutingValueOff:
		return true
	}
	return false
}

type SmartShieldGetResponseSmartTieredCache struct {
	// The id of the Smart Tiered Cache setting.
	ID string `json:"id"`
	// Whether the setting is editable.
	Editable bool `json:"editable"`
	// The last time the setting was modified.
	ModifiedOn string `json:"modified_on"`
	// Specifies the enablement value of Tiered Cache.
	Value SmartShieldGetResponseSmartTieredCacheValue `json:"value"`
	JSON  smartShieldGetResponseSmartTieredCacheJSON  `json:"-"`
}

// smartShieldGetResponseSmartTieredCacheJSON contains the JSON metadata for the
// struct [SmartShieldGetResponseSmartTieredCache]
type smartShieldGetResponseSmartTieredCacheJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldGetResponseSmartTieredCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldGetResponseSmartTieredCacheJSON) RawJSON() string {
	return r.raw
}

// Specifies the enablement value of Tiered Cache.
type SmartShieldGetResponseSmartTieredCacheValue string

const (
	SmartShieldGetResponseSmartTieredCacheValueOn  SmartShieldGetResponseSmartTieredCacheValue = "on"
	SmartShieldGetResponseSmartTieredCacheValueOff SmartShieldGetResponseSmartTieredCacheValue = "off"
)

func (r SmartShieldGetResponseSmartTieredCacheValue) IsKnown() bool {
	switch r {
	case SmartShieldGetResponseSmartTieredCacheValueOn, SmartShieldGetResponseSmartTieredCacheValueOff:
		return true
	}
	return false
}

type SmartShieldGetHealthcheckResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []SmartShieldGetHealthcheckResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig SmartShieldGetHealthcheckResponseHTTPConfig `json:"http_config,nullable"`
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
	Status SmartShieldGetHealthcheckResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TCPConfig SmartShieldGetHealthcheckResponseTCPConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                                `json:"type"`
	JSON smartShieldGetHealthcheckResponseJSON `json:"-"`
}

// smartShieldGetHealthcheckResponseJSON contains the JSON metadata for the struct
// [SmartShieldGetHealthcheckResponse]
type smartShieldGetHealthcheckResponseJSON struct {
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

func (r *SmartShieldGetHealthcheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldGetHealthcheckResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type SmartShieldGetHealthcheckResponseCheckRegion string

const (
	SmartShieldGetHealthcheckResponseCheckRegionWnam       SmartShieldGetHealthcheckResponseCheckRegion = "WNAM"
	SmartShieldGetHealthcheckResponseCheckRegionEnam       SmartShieldGetHealthcheckResponseCheckRegion = "ENAM"
	SmartShieldGetHealthcheckResponseCheckRegionWeu        SmartShieldGetHealthcheckResponseCheckRegion = "WEU"
	SmartShieldGetHealthcheckResponseCheckRegionEeu        SmartShieldGetHealthcheckResponseCheckRegion = "EEU"
	SmartShieldGetHealthcheckResponseCheckRegionNsam       SmartShieldGetHealthcheckResponseCheckRegion = "NSAM"
	SmartShieldGetHealthcheckResponseCheckRegionSsam       SmartShieldGetHealthcheckResponseCheckRegion = "SSAM"
	SmartShieldGetHealthcheckResponseCheckRegionOc         SmartShieldGetHealthcheckResponseCheckRegion = "OC"
	SmartShieldGetHealthcheckResponseCheckRegionMe         SmartShieldGetHealthcheckResponseCheckRegion = "ME"
	SmartShieldGetHealthcheckResponseCheckRegionNaf        SmartShieldGetHealthcheckResponseCheckRegion = "NAF"
	SmartShieldGetHealthcheckResponseCheckRegionSaf        SmartShieldGetHealthcheckResponseCheckRegion = "SAF"
	SmartShieldGetHealthcheckResponseCheckRegionIn         SmartShieldGetHealthcheckResponseCheckRegion = "IN"
	SmartShieldGetHealthcheckResponseCheckRegionSeas       SmartShieldGetHealthcheckResponseCheckRegion = "SEAS"
	SmartShieldGetHealthcheckResponseCheckRegionNeas       SmartShieldGetHealthcheckResponseCheckRegion = "NEAS"
	SmartShieldGetHealthcheckResponseCheckRegionAllRegions SmartShieldGetHealthcheckResponseCheckRegion = "ALL_REGIONS"
)

func (r SmartShieldGetHealthcheckResponseCheckRegion) IsKnown() bool {
	switch r {
	case SmartShieldGetHealthcheckResponseCheckRegionWnam, SmartShieldGetHealthcheckResponseCheckRegionEnam, SmartShieldGetHealthcheckResponseCheckRegionWeu, SmartShieldGetHealthcheckResponseCheckRegionEeu, SmartShieldGetHealthcheckResponseCheckRegionNsam, SmartShieldGetHealthcheckResponseCheckRegionSsam, SmartShieldGetHealthcheckResponseCheckRegionOc, SmartShieldGetHealthcheckResponseCheckRegionMe, SmartShieldGetHealthcheckResponseCheckRegionNaf, SmartShieldGetHealthcheckResponseCheckRegionSaf, SmartShieldGetHealthcheckResponseCheckRegionIn, SmartShieldGetHealthcheckResponseCheckRegionSeas, SmartShieldGetHealthcheckResponseCheckRegionNeas, SmartShieldGetHealthcheckResponseCheckRegionAllRegions:
		return true
	}
	return false
}

// Parameters specific to an HTTP or HTTPS health check.
type SmartShieldGetHealthcheckResponseHTTPConfig struct {
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
	Header map[string][]string `json:"header,nullable"`
	// The HTTP method to use for the health check.
	Method SmartShieldGetHealthcheckResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                           `json:"port"`
	JSON smartShieldGetHealthcheckResponseHTTPConfigJSON `json:"-"`
}

// smartShieldGetHealthcheckResponseHTTPConfigJSON contains the JSON metadata for
// the struct [SmartShieldGetHealthcheckResponseHTTPConfig]
type smartShieldGetHealthcheckResponseHTTPConfigJSON struct {
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

func (r *SmartShieldGetHealthcheckResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldGetHealthcheckResponseHTTPConfigJSON) RawJSON() string {
	return r.raw
}

// The HTTP method to use for the health check.
type SmartShieldGetHealthcheckResponseHTTPConfigMethod string

const (
	SmartShieldGetHealthcheckResponseHTTPConfigMethodGet  SmartShieldGetHealthcheckResponseHTTPConfigMethod = "GET"
	SmartShieldGetHealthcheckResponseHTTPConfigMethodHead SmartShieldGetHealthcheckResponseHTTPConfigMethod = "HEAD"
)

func (r SmartShieldGetHealthcheckResponseHTTPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldGetHealthcheckResponseHTTPConfigMethodGet, SmartShieldGetHealthcheckResponseHTTPConfigMethodHead:
		return true
	}
	return false
}

// The current status of the origin server according to the health check.
type SmartShieldGetHealthcheckResponseStatus string

const (
	SmartShieldGetHealthcheckResponseStatusUnknown   SmartShieldGetHealthcheckResponseStatus = "unknown"
	SmartShieldGetHealthcheckResponseStatusHealthy   SmartShieldGetHealthcheckResponseStatus = "healthy"
	SmartShieldGetHealthcheckResponseStatusUnhealthy SmartShieldGetHealthcheckResponseStatus = "unhealthy"
	SmartShieldGetHealthcheckResponseStatusSuspended SmartShieldGetHealthcheckResponseStatus = "suspended"
)

func (r SmartShieldGetHealthcheckResponseStatus) IsKnown() bool {
	switch r {
	case SmartShieldGetHealthcheckResponseStatusUnknown, SmartShieldGetHealthcheckResponseStatusHealthy, SmartShieldGetHealthcheckResponseStatusUnhealthy, SmartShieldGetHealthcheckResponseStatusSuspended:
		return true
	}
	return false
}

// Parameters specific to TCP health check.
type SmartShieldGetHealthcheckResponseTCPConfig struct {
	// The TCP connection method to use for the health check.
	Method SmartShieldGetHealthcheckResponseTCPConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                          `json:"port"`
	JSON smartShieldGetHealthcheckResponseTCPConfigJSON `json:"-"`
}

// smartShieldGetHealthcheckResponseTCPConfigJSON contains the JSON metadata for
// the struct [SmartShieldGetHealthcheckResponseTCPConfig]
type smartShieldGetHealthcheckResponseTCPConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldGetHealthcheckResponseTCPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldGetHealthcheckResponseTCPConfigJSON) RawJSON() string {
	return r.raw
}

// The TCP connection method to use for the health check.
type SmartShieldGetHealthcheckResponseTCPConfigMethod string

const (
	SmartShieldGetHealthcheckResponseTCPConfigMethodConnectionEstablished SmartShieldGetHealthcheckResponseTCPConfigMethod = "connection_established"
)

func (r SmartShieldGetHealthcheckResponseTCPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldGetHealthcheckResponseTCPConfigMethodConnectionEstablished:
		return true
	}
	return false
}

type SmartShieldListHealthchecksResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []SmartShieldListHealthchecksResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig SmartShieldListHealthchecksResponseHTTPConfig `json:"http_config,nullable"`
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
	Status SmartShieldListHealthchecksResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TCPConfig SmartShieldListHealthchecksResponseTCPConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                                  `json:"type"`
	JSON smartShieldListHealthchecksResponseJSON `json:"-"`
}

// smartShieldListHealthchecksResponseJSON contains the JSON metadata for the
// struct [SmartShieldListHealthchecksResponse]
type smartShieldListHealthchecksResponseJSON struct {
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

func (r *SmartShieldListHealthchecksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldListHealthchecksResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type SmartShieldListHealthchecksResponseCheckRegion string

const (
	SmartShieldListHealthchecksResponseCheckRegionWnam       SmartShieldListHealthchecksResponseCheckRegion = "WNAM"
	SmartShieldListHealthchecksResponseCheckRegionEnam       SmartShieldListHealthchecksResponseCheckRegion = "ENAM"
	SmartShieldListHealthchecksResponseCheckRegionWeu        SmartShieldListHealthchecksResponseCheckRegion = "WEU"
	SmartShieldListHealthchecksResponseCheckRegionEeu        SmartShieldListHealthchecksResponseCheckRegion = "EEU"
	SmartShieldListHealthchecksResponseCheckRegionNsam       SmartShieldListHealthchecksResponseCheckRegion = "NSAM"
	SmartShieldListHealthchecksResponseCheckRegionSsam       SmartShieldListHealthchecksResponseCheckRegion = "SSAM"
	SmartShieldListHealthchecksResponseCheckRegionOc         SmartShieldListHealthchecksResponseCheckRegion = "OC"
	SmartShieldListHealthchecksResponseCheckRegionMe         SmartShieldListHealthchecksResponseCheckRegion = "ME"
	SmartShieldListHealthchecksResponseCheckRegionNaf        SmartShieldListHealthchecksResponseCheckRegion = "NAF"
	SmartShieldListHealthchecksResponseCheckRegionSaf        SmartShieldListHealthchecksResponseCheckRegion = "SAF"
	SmartShieldListHealthchecksResponseCheckRegionIn         SmartShieldListHealthchecksResponseCheckRegion = "IN"
	SmartShieldListHealthchecksResponseCheckRegionSeas       SmartShieldListHealthchecksResponseCheckRegion = "SEAS"
	SmartShieldListHealthchecksResponseCheckRegionNeas       SmartShieldListHealthchecksResponseCheckRegion = "NEAS"
	SmartShieldListHealthchecksResponseCheckRegionAllRegions SmartShieldListHealthchecksResponseCheckRegion = "ALL_REGIONS"
)

func (r SmartShieldListHealthchecksResponseCheckRegion) IsKnown() bool {
	switch r {
	case SmartShieldListHealthchecksResponseCheckRegionWnam, SmartShieldListHealthchecksResponseCheckRegionEnam, SmartShieldListHealthchecksResponseCheckRegionWeu, SmartShieldListHealthchecksResponseCheckRegionEeu, SmartShieldListHealthchecksResponseCheckRegionNsam, SmartShieldListHealthchecksResponseCheckRegionSsam, SmartShieldListHealthchecksResponseCheckRegionOc, SmartShieldListHealthchecksResponseCheckRegionMe, SmartShieldListHealthchecksResponseCheckRegionNaf, SmartShieldListHealthchecksResponseCheckRegionSaf, SmartShieldListHealthchecksResponseCheckRegionIn, SmartShieldListHealthchecksResponseCheckRegionSeas, SmartShieldListHealthchecksResponseCheckRegionNeas, SmartShieldListHealthchecksResponseCheckRegionAllRegions:
		return true
	}
	return false
}

// Parameters specific to an HTTP or HTTPS health check.
type SmartShieldListHealthchecksResponseHTTPConfig struct {
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
	Header map[string][]string `json:"header,nullable"`
	// The HTTP method to use for the health check.
	Method SmartShieldListHealthchecksResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                             `json:"port"`
	JSON smartShieldListHealthchecksResponseHTTPConfigJSON `json:"-"`
}

// smartShieldListHealthchecksResponseHTTPConfigJSON contains the JSON metadata for
// the struct [SmartShieldListHealthchecksResponseHTTPConfig]
type smartShieldListHealthchecksResponseHTTPConfigJSON struct {
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

func (r *SmartShieldListHealthchecksResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldListHealthchecksResponseHTTPConfigJSON) RawJSON() string {
	return r.raw
}

// The HTTP method to use for the health check.
type SmartShieldListHealthchecksResponseHTTPConfigMethod string

const (
	SmartShieldListHealthchecksResponseHTTPConfigMethodGet  SmartShieldListHealthchecksResponseHTTPConfigMethod = "GET"
	SmartShieldListHealthchecksResponseHTTPConfigMethodHead SmartShieldListHealthchecksResponseHTTPConfigMethod = "HEAD"
)

func (r SmartShieldListHealthchecksResponseHTTPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldListHealthchecksResponseHTTPConfigMethodGet, SmartShieldListHealthchecksResponseHTTPConfigMethodHead:
		return true
	}
	return false
}

// The current status of the origin server according to the health check.
type SmartShieldListHealthchecksResponseStatus string

const (
	SmartShieldListHealthchecksResponseStatusUnknown   SmartShieldListHealthchecksResponseStatus = "unknown"
	SmartShieldListHealthchecksResponseStatusHealthy   SmartShieldListHealthchecksResponseStatus = "healthy"
	SmartShieldListHealthchecksResponseStatusUnhealthy SmartShieldListHealthchecksResponseStatus = "unhealthy"
	SmartShieldListHealthchecksResponseStatusSuspended SmartShieldListHealthchecksResponseStatus = "suspended"
)

func (r SmartShieldListHealthchecksResponseStatus) IsKnown() bool {
	switch r {
	case SmartShieldListHealthchecksResponseStatusUnknown, SmartShieldListHealthchecksResponseStatusHealthy, SmartShieldListHealthchecksResponseStatusUnhealthy, SmartShieldListHealthchecksResponseStatusSuspended:
		return true
	}
	return false
}

// Parameters specific to TCP health check.
type SmartShieldListHealthchecksResponseTCPConfig struct {
	// The TCP connection method to use for the health check.
	Method SmartShieldListHealthchecksResponseTCPConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                            `json:"port"`
	JSON smartShieldListHealthchecksResponseTCPConfigJSON `json:"-"`
}

// smartShieldListHealthchecksResponseTCPConfigJSON contains the JSON metadata for
// the struct [SmartShieldListHealthchecksResponseTCPConfig]
type smartShieldListHealthchecksResponseTCPConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldListHealthchecksResponseTCPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldListHealthchecksResponseTCPConfigJSON) RawJSON() string {
	return r.raw
}

// The TCP connection method to use for the health check.
type SmartShieldListHealthchecksResponseTCPConfigMethod string

const (
	SmartShieldListHealthchecksResponseTCPConfigMethodConnectionEstablished SmartShieldListHealthchecksResponseTCPConfigMethod = "connection_established"
)

func (r SmartShieldListHealthchecksResponseTCPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldListHealthchecksResponseTCPConfigMethodConnectionEstablished:
		return true
	}
	return false
}

type SmartShieldUpdateHealthcheckResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []SmartShieldUpdateHealthcheckResponseCheckRegion `json:"check_regions,nullable"`
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
	HTTPConfig SmartShieldUpdateHealthcheckResponseHTTPConfig `json:"http_config,nullable"`
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
	Status SmartShieldUpdateHealthcheckResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TCPConfig SmartShieldUpdateHealthcheckResponseTCPConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                                   `json:"type"`
	JSON smartShieldUpdateHealthcheckResponseJSON `json:"-"`
}

// smartShieldUpdateHealthcheckResponseJSON contains the JSON metadata for the
// struct [SmartShieldUpdateHealthcheckResponse]
type smartShieldUpdateHealthcheckResponseJSON struct {
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

func (r *SmartShieldUpdateHealthcheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldUpdateHealthcheckResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type SmartShieldUpdateHealthcheckResponseCheckRegion string

const (
	SmartShieldUpdateHealthcheckResponseCheckRegionWnam       SmartShieldUpdateHealthcheckResponseCheckRegion = "WNAM"
	SmartShieldUpdateHealthcheckResponseCheckRegionEnam       SmartShieldUpdateHealthcheckResponseCheckRegion = "ENAM"
	SmartShieldUpdateHealthcheckResponseCheckRegionWeu        SmartShieldUpdateHealthcheckResponseCheckRegion = "WEU"
	SmartShieldUpdateHealthcheckResponseCheckRegionEeu        SmartShieldUpdateHealthcheckResponseCheckRegion = "EEU"
	SmartShieldUpdateHealthcheckResponseCheckRegionNsam       SmartShieldUpdateHealthcheckResponseCheckRegion = "NSAM"
	SmartShieldUpdateHealthcheckResponseCheckRegionSsam       SmartShieldUpdateHealthcheckResponseCheckRegion = "SSAM"
	SmartShieldUpdateHealthcheckResponseCheckRegionOc         SmartShieldUpdateHealthcheckResponseCheckRegion = "OC"
	SmartShieldUpdateHealthcheckResponseCheckRegionMe         SmartShieldUpdateHealthcheckResponseCheckRegion = "ME"
	SmartShieldUpdateHealthcheckResponseCheckRegionNaf        SmartShieldUpdateHealthcheckResponseCheckRegion = "NAF"
	SmartShieldUpdateHealthcheckResponseCheckRegionSaf        SmartShieldUpdateHealthcheckResponseCheckRegion = "SAF"
	SmartShieldUpdateHealthcheckResponseCheckRegionIn         SmartShieldUpdateHealthcheckResponseCheckRegion = "IN"
	SmartShieldUpdateHealthcheckResponseCheckRegionSeas       SmartShieldUpdateHealthcheckResponseCheckRegion = "SEAS"
	SmartShieldUpdateHealthcheckResponseCheckRegionNeas       SmartShieldUpdateHealthcheckResponseCheckRegion = "NEAS"
	SmartShieldUpdateHealthcheckResponseCheckRegionAllRegions SmartShieldUpdateHealthcheckResponseCheckRegion = "ALL_REGIONS"
)

func (r SmartShieldUpdateHealthcheckResponseCheckRegion) IsKnown() bool {
	switch r {
	case SmartShieldUpdateHealthcheckResponseCheckRegionWnam, SmartShieldUpdateHealthcheckResponseCheckRegionEnam, SmartShieldUpdateHealthcheckResponseCheckRegionWeu, SmartShieldUpdateHealthcheckResponseCheckRegionEeu, SmartShieldUpdateHealthcheckResponseCheckRegionNsam, SmartShieldUpdateHealthcheckResponseCheckRegionSsam, SmartShieldUpdateHealthcheckResponseCheckRegionOc, SmartShieldUpdateHealthcheckResponseCheckRegionMe, SmartShieldUpdateHealthcheckResponseCheckRegionNaf, SmartShieldUpdateHealthcheckResponseCheckRegionSaf, SmartShieldUpdateHealthcheckResponseCheckRegionIn, SmartShieldUpdateHealthcheckResponseCheckRegionSeas, SmartShieldUpdateHealthcheckResponseCheckRegionNeas, SmartShieldUpdateHealthcheckResponseCheckRegionAllRegions:
		return true
	}
	return false
}

// Parameters specific to an HTTP or HTTPS health check.
type SmartShieldUpdateHealthcheckResponseHTTPConfig struct {
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
	Header map[string][]string `json:"header,nullable"`
	// The HTTP method to use for the health check.
	Method SmartShieldUpdateHealthcheckResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                              `json:"port"`
	JSON smartShieldUpdateHealthcheckResponseHTTPConfigJSON `json:"-"`
}

// smartShieldUpdateHealthcheckResponseHTTPConfigJSON contains the JSON metadata
// for the struct [SmartShieldUpdateHealthcheckResponseHTTPConfig]
type smartShieldUpdateHealthcheckResponseHTTPConfigJSON struct {
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

func (r *SmartShieldUpdateHealthcheckResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldUpdateHealthcheckResponseHTTPConfigJSON) RawJSON() string {
	return r.raw
}

// The HTTP method to use for the health check.
type SmartShieldUpdateHealthcheckResponseHTTPConfigMethod string

const (
	SmartShieldUpdateHealthcheckResponseHTTPConfigMethodGet  SmartShieldUpdateHealthcheckResponseHTTPConfigMethod = "GET"
	SmartShieldUpdateHealthcheckResponseHTTPConfigMethodHead SmartShieldUpdateHealthcheckResponseHTTPConfigMethod = "HEAD"
)

func (r SmartShieldUpdateHealthcheckResponseHTTPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldUpdateHealthcheckResponseHTTPConfigMethodGet, SmartShieldUpdateHealthcheckResponseHTTPConfigMethodHead:
		return true
	}
	return false
}

// The current status of the origin server according to the health check.
type SmartShieldUpdateHealthcheckResponseStatus string

const (
	SmartShieldUpdateHealthcheckResponseStatusUnknown   SmartShieldUpdateHealthcheckResponseStatus = "unknown"
	SmartShieldUpdateHealthcheckResponseStatusHealthy   SmartShieldUpdateHealthcheckResponseStatus = "healthy"
	SmartShieldUpdateHealthcheckResponseStatusUnhealthy SmartShieldUpdateHealthcheckResponseStatus = "unhealthy"
	SmartShieldUpdateHealthcheckResponseStatusSuspended SmartShieldUpdateHealthcheckResponseStatus = "suspended"
)

func (r SmartShieldUpdateHealthcheckResponseStatus) IsKnown() bool {
	switch r {
	case SmartShieldUpdateHealthcheckResponseStatusUnknown, SmartShieldUpdateHealthcheckResponseStatusHealthy, SmartShieldUpdateHealthcheckResponseStatusUnhealthy, SmartShieldUpdateHealthcheckResponseStatusSuspended:
		return true
	}
	return false
}

// Parameters specific to TCP health check.
type SmartShieldUpdateHealthcheckResponseTCPConfig struct {
	// The TCP connection method to use for the health check.
	Method SmartShieldUpdateHealthcheckResponseTCPConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                             `json:"port"`
	JSON smartShieldUpdateHealthcheckResponseTCPConfigJSON `json:"-"`
}

// smartShieldUpdateHealthcheckResponseTCPConfigJSON contains the JSON metadata for
// the struct [SmartShieldUpdateHealthcheckResponseTCPConfig]
type smartShieldUpdateHealthcheckResponseTCPConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldUpdateHealthcheckResponseTCPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldUpdateHealthcheckResponseTCPConfigJSON) RawJSON() string {
	return r.raw
}

// The TCP connection method to use for the health check.
type SmartShieldUpdateHealthcheckResponseTCPConfigMethod string

const (
	SmartShieldUpdateHealthcheckResponseTCPConfigMethodConnectionEstablished SmartShieldUpdateHealthcheckResponseTCPConfigMethod = "connection_established"
)

func (r SmartShieldUpdateHealthcheckResponseTCPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldUpdateHealthcheckResponseTCPConfigMethodConnectionEstablished:
		return true
	}
	return false
}

type SmartShieldUpdateParams struct {
	// Identifier.
	ZoneID              param.Field[string]                                     `path:"zone_id,required"`
	CacheReserve        param.Field[SmartShieldUpdateParamsCacheReserve]        `json:"cache_reserve"`
	RegionalTieredCache param.Field[SmartShieldUpdateParamsRegionalTieredCache] `json:"regional_tiered_cache"`
	SmartRouting        param.Field[SmartShieldUpdateParamsSmartRouting]        `json:"smart_routing"`
	SmartTieredCache    param.Field[SmartShieldUpdateParamsSmartTieredCache]    `json:"smart_tiered_cache"`
}

func (r SmartShieldUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SmartShieldUpdateParamsCacheReserve struct {
	// Specifies the enablement value of Cache Reserve.
	Value param.Field[SmartShieldUpdateParamsCacheReserveValue] `json:"value"`
}

func (r SmartShieldUpdateParamsCacheReserve) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the enablement value of Cache Reserve.
type SmartShieldUpdateParamsCacheReserveValue string

const (
	SmartShieldUpdateParamsCacheReserveValueOn  SmartShieldUpdateParamsCacheReserveValue = "on"
	SmartShieldUpdateParamsCacheReserveValueOff SmartShieldUpdateParamsCacheReserveValue = "off"
)

func (r SmartShieldUpdateParamsCacheReserveValue) IsKnown() bool {
	switch r {
	case SmartShieldUpdateParamsCacheReserveValueOn, SmartShieldUpdateParamsCacheReserveValueOff:
		return true
	}
	return false
}

type SmartShieldUpdateParamsRegionalTieredCache struct {
	// Specifies the enablement value of Regional Tiered Cache.
	Value param.Field[SmartShieldUpdateParamsRegionalTieredCacheValue] `json:"value"`
}

func (r SmartShieldUpdateParamsRegionalTieredCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the enablement value of Regional Tiered Cache.
type SmartShieldUpdateParamsRegionalTieredCacheValue string

const (
	SmartShieldUpdateParamsRegionalTieredCacheValueOn  SmartShieldUpdateParamsRegionalTieredCacheValue = "on"
	SmartShieldUpdateParamsRegionalTieredCacheValueOff SmartShieldUpdateParamsRegionalTieredCacheValue = "off"
)

func (r SmartShieldUpdateParamsRegionalTieredCacheValue) IsKnown() bool {
	switch r {
	case SmartShieldUpdateParamsRegionalTieredCacheValueOn, SmartShieldUpdateParamsRegionalTieredCacheValueOff:
		return true
	}
	return false
}

type SmartShieldUpdateParamsSmartRouting struct {
	// Specifies the enablement value of Smart Routing.
	Value param.Field[SmartShieldUpdateParamsSmartRoutingValue] `json:"value"`
}

func (r SmartShieldUpdateParamsSmartRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the enablement value of Smart Routing.
type SmartShieldUpdateParamsSmartRoutingValue string

const (
	SmartShieldUpdateParamsSmartRoutingValueOn  SmartShieldUpdateParamsSmartRoutingValue = "on"
	SmartShieldUpdateParamsSmartRoutingValueOff SmartShieldUpdateParamsSmartRoutingValue = "off"
)

func (r SmartShieldUpdateParamsSmartRoutingValue) IsKnown() bool {
	switch r {
	case SmartShieldUpdateParamsSmartRoutingValueOn, SmartShieldUpdateParamsSmartRoutingValueOff:
		return true
	}
	return false
}

type SmartShieldUpdateParamsSmartTieredCache struct {
	// Specifies the enablement value of Smart Tiered Cache.
	Value param.Field[SmartShieldUpdateParamsSmartTieredCacheValue] `json:"value"`
}

func (r SmartShieldUpdateParamsSmartTieredCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the enablement value of Smart Tiered Cache.
type SmartShieldUpdateParamsSmartTieredCacheValue string

const (
	SmartShieldUpdateParamsSmartTieredCacheValueOn  SmartShieldUpdateParamsSmartTieredCacheValue = "on"
	SmartShieldUpdateParamsSmartTieredCacheValueOff SmartShieldUpdateParamsSmartTieredCacheValue = "off"
)

func (r SmartShieldUpdateParamsSmartTieredCacheValue) IsKnown() bool {
	switch r {
	case SmartShieldUpdateParamsSmartTieredCacheValueOn, SmartShieldUpdateParamsSmartTieredCacheValueOff:
		return true
	}
	return false
}

type SmartShieldUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A consolidated object containing settings from multiple APIs for partial
	// updates.
	Result SmartShieldUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SmartShieldUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartShieldUpdateResponseEnvelopeJSON    `json:"-"`
}

// smartShieldUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SmartShieldUpdateResponseEnvelope]
type smartShieldUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SmartShieldUpdateResponseEnvelopeSuccess bool

const (
	SmartShieldUpdateResponseEnvelopeSuccessTrue SmartShieldUpdateResponseEnvelopeSuccess = true
)

func (r SmartShieldUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartShieldUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SmartShieldNewHealthcheckParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]SmartShieldNewHealthcheckParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[SmartShieldNewHealthcheckParamsHTTPConfig] `json:"http_config"`
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
	TCPConfig param.Field[SmartShieldNewHealthcheckParamsTCPConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r SmartShieldNewHealthcheckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type SmartShieldNewHealthcheckParamsCheckRegion string

const (
	SmartShieldNewHealthcheckParamsCheckRegionWnam       SmartShieldNewHealthcheckParamsCheckRegion = "WNAM"
	SmartShieldNewHealthcheckParamsCheckRegionEnam       SmartShieldNewHealthcheckParamsCheckRegion = "ENAM"
	SmartShieldNewHealthcheckParamsCheckRegionWeu        SmartShieldNewHealthcheckParamsCheckRegion = "WEU"
	SmartShieldNewHealthcheckParamsCheckRegionEeu        SmartShieldNewHealthcheckParamsCheckRegion = "EEU"
	SmartShieldNewHealthcheckParamsCheckRegionNsam       SmartShieldNewHealthcheckParamsCheckRegion = "NSAM"
	SmartShieldNewHealthcheckParamsCheckRegionSsam       SmartShieldNewHealthcheckParamsCheckRegion = "SSAM"
	SmartShieldNewHealthcheckParamsCheckRegionOc         SmartShieldNewHealthcheckParamsCheckRegion = "OC"
	SmartShieldNewHealthcheckParamsCheckRegionMe         SmartShieldNewHealthcheckParamsCheckRegion = "ME"
	SmartShieldNewHealthcheckParamsCheckRegionNaf        SmartShieldNewHealthcheckParamsCheckRegion = "NAF"
	SmartShieldNewHealthcheckParamsCheckRegionSaf        SmartShieldNewHealthcheckParamsCheckRegion = "SAF"
	SmartShieldNewHealthcheckParamsCheckRegionIn         SmartShieldNewHealthcheckParamsCheckRegion = "IN"
	SmartShieldNewHealthcheckParamsCheckRegionSeas       SmartShieldNewHealthcheckParamsCheckRegion = "SEAS"
	SmartShieldNewHealthcheckParamsCheckRegionNeas       SmartShieldNewHealthcheckParamsCheckRegion = "NEAS"
	SmartShieldNewHealthcheckParamsCheckRegionAllRegions SmartShieldNewHealthcheckParamsCheckRegion = "ALL_REGIONS"
)

func (r SmartShieldNewHealthcheckParamsCheckRegion) IsKnown() bool {
	switch r {
	case SmartShieldNewHealthcheckParamsCheckRegionWnam, SmartShieldNewHealthcheckParamsCheckRegionEnam, SmartShieldNewHealthcheckParamsCheckRegionWeu, SmartShieldNewHealthcheckParamsCheckRegionEeu, SmartShieldNewHealthcheckParamsCheckRegionNsam, SmartShieldNewHealthcheckParamsCheckRegionSsam, SmartShieldNewHealthcheckParamsCheckRegionOc, SmartShieldNewHealthcheckParamsCheckRegionMe, SmartShieldNewHealthcheckParamsCheckRegionNaf, SmartShieldNewHealthcheckParamsCheckRegionSaf, SmartShieldNewHealthcheckParamsCheckRegionIn, SmartShieldNewHealthcheckParamsCheckRegionSeas, SmartShieldNewHealthcheckParamsCheckRegionNeas, SmartShieldNewHealthcheckParamsCheckRegionAllRegions:
		return true
	}
	return false
}

// Parameters specific to an HTTP or HTTPS health check.
type SmartShieldNewHealthcheckParamsHTTPConfig struct {
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
	Header param.Field[map[string][]string] `json:"header"`
	// The HTTP method to use for the health check.
	Method param.Field[SmartShieldNewHealthcheckParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r SmartShieldNewHealthcheckParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type SmartShieldNewHealthcheckParamsHTTPConfigMethod string

const (
	SmartShieldNewHealthcheckParamsHTTPConfigMethodGet  SmartShieldNewHealthcheckParamsHTTPConfigMethod = "GET"
	SmartShieldNewHealthcheckParamsHTTPConfigMethodHead SmartShieldNewHealthcheckParamsHTTPConfigMethod = "HEAD"
)

func (r SmartShieldNewHealthcheckParamsHTTPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldNewHealthcheckParamsHTTPConfigMethodGet, SmartShieldNewHealthcheckParamsHTTPConfigMethodHead:
		return true
	}
	return false
}

// Parameters specific to TCP health check.
type SmartShieldNewHealthcheckParamsTCPConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[SmartShieldNewHealthcheckParamsTCPConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r SmartShieldNewHealthcheckParamsTCPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type SmartShieldNewHealthcheckParamsTCPConfigMethod string

const (
	SmartShieldNewHealthcheckParamsTCPConfigMethodConnectionEstablished SmartShieldNewHealthcheckParamsTCPConfigMethod = "connection_established"
)

func (r SmartShieldNewHealthcheckParamsTCPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldNewHealthcheckParamsTCPConfigMethodConnectionEstablished:
		return true
	}
	return false
}

type SmartShieldNewHealthcheckResponseEnvelope struct {
	Errors   []shared.ResponseInfo             `json:"errors,required"`
	Messages []shared.ResponseInfo             `json:"messages,required"`
	Result   SmartShieldNewHealthcheckResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SmartShieldNewHealthcheckResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartShieldNewHealthcheckResponseEnvelopeJSON    `json:"-"`
}

// smartShieldNewHealthcheckResponseEnvelopeJSON contains the JSON metadata for the
// struct [SmartShieldNewHealthcheckResponseEnvelope]
type smartShieldNewHealthcheckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldNewHealthcheckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldNewHealthcheckResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SmartShieldNewHealthcheckResponseEnvelopeSuccess bool

const (
	SmartShieldNewHealthcheckResponseEnvelopeSuccessTrue SmartShieldNewHealthcheckResponseEnvelopeSuccess = true
)

func (r SmartShieldNewHealthcheckResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartShieldNewHealthcheckResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SmartShieldDeleteHealthcheckParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SmartShieldDeleteHealthcheckResponseEnvelope struct {
	Errors   []shared.ResponseInfo                `json:"errors,required"`
	Messages []shared.ResponseInfo                `json:"messages,required"`
	Result   SmartShieldDeleteHealthcheckResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SmartShieldDeleteHealthcheckResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartShieldDeleteHealthcheckResponseEnvelopeJSON    `json:"-"`
}

// smartShieldDeleteHealthcheckResponseEnvelopeJSON contains the JSON metadata for
// the struct [SmartShieldDeleteHealthcheckResponseEnvelope]
type smartShieldDeleteHealthcheckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldDeleteHealthcheckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldDeleteHealthcheckResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SmartShieldDeleteHealthcheckResponseEnvelopeSuccess bool

const (
	SmartShieldDeleteHealthcheckResponseEnvelopeSuccessTrue SmartShieldDeleteHealthcheckResponseEnvelopeSuccess = true
)

func (r SmartShieldDeleteHealthcheckResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartShieldDeleteHealthcheckResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SmartShieldEditHealthcheckParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]SmartShieldEditHealthcheckParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[SmartShieldEditHealthcheckParamsHTTPConfig] `json:"http_config"`
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
	TCPConfig param.Field[SmartShieldEditHealthcheckParamsTCPConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r SmartShieldEditHealthcheckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type SmartShieldEditHealthcheckParamsCheckRegion string

const (
	SmartShieldEditHealthcheckParamsCheckRegionWnam       SmartShieldEditHealthcheckParamsCheckRegion = "WNAM"
	SmartShieldEditHealthcheckParamsCheckRegionEnam       SmartShieldEditHealthcheckParamsCheckRegion = "ENAM"
	SmartShieldEditHealthcheckParamsCheckRegionWeu        SmartShieldEditHealthcheckParamsCheckRegion = "WEU"
	SmartShieldEditHealthcheckParamsCheckRegionEeu        SmartShieldEditHealthcheckParamsCheckRegion = "EEU"
	SmartShieldEditHealthcheckParamsCheckRegionNsam       SmartShieldEditHealthcheckParamsCheckRegion = "NSAM"
	SmartShieldEditHealthcheckParamsCheckRegionSsam       SmartShieldEditHealthcheckParamsCheckRegion = "SSAM"
	SmartShieldEditHealthcheckParamsCheckRegionOc         SmartShieldEditHealthcheckParamsCheckRegion = "OC"
	SmartShieldEditHealthcheckParamsCheckRegionMe         SmartShieldEditHealthcheckParamsCheckRegion = "ME"
	SmartShieldEditHealthcheckParamsCheckRegionNaf        SmartShieldEditHealthcheckParamsCheckRegion = "NAF"
	SmartShieldEditHealthcheckParamsCheckRegionSaf        SmartShieldEditHealthcheckParamsCheckRegion = "SAF"
	SmartShieldEditHealthcheckParamsCheckRegionIn         SmartShieldEditHealthcheckParamsCheckRegion = "IN"
	SmartShieldEditHealthcheckParamsCheckRegionSeas       SmartShieldEditHealthcheckParamsCheckRegion = "SEAS"
	SmartShieldEditHealthcheckParamsCheckRegionNeas       SmartShieldEditHealthcheckParamsCheckRegion = "NEAS"
	SmartShieldEditHealthcheckParamsCheckRegionAllRegions SmartShieldEditHealthcheckParamsCheckRegion = "ALL_REGIONS"
)

func (r SmartShieldEditHealthcheckParamsCheckRegion) IsKnown() bool {
	switch r {
	case SmartShieldEditHealthcheckParamsCheckRegionWnam, SmartShieldEditHealthcheckParamsCheckRegionEnam, SmartShieldEditHealthcheckParamsCheckRegionWeu, SmartShieldEditHealthcheckParamsCheckRegionEeu, SmartShieldEditHealthcheckParamsCheckRegionNsam, SmartShieldEditHealthcheckParamsCheckRegionSsam, SmartShieldEditHealthcheckParamsCheckRegionOc, SmartShieldEditHealthcheckParamsCheckRegionMe, SmartShieldEditHealthcheckParamsCheckRegionNaf, SmartShieldEditHealthcheckParamsCheckRegionSaf, SmartShieldEditHealthcheckParamsCheckRegionIn, SmartShieldEditHealthcheckParamsCheckRegionSeas, SmartShieldEditHealthcheckParamsCheckRegionNeas, SmartShieldEditHealthcheckParamsCheckRegionAllRegions:
		return true
	}
	return false
}

// Parameters specific to an HTTP or HTTPS health check.
type SmartShieldEditHealthcheckParamsHTTPConfig struct {
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
	Header param.Field[map[string][]string] `json:"header"`
	// The HTTP method to use for the health check.
	Method param.Field[SmartShieldEditHealthcheckParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r SmartShieldEditHealthcheckParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type SmartShieldEditHealthcheckParamsHTTPConfigMethod string

const (
	SmartShieldEditHealthcheckParamsHTTPConfigMethodGet  SmartShieldEditHealthcheckParamsHTTPConfigMethod = "GET"
	SmartShieldEditHealthcheckParamsHTTPConfigMethodHead SmartShieldEditHealthcheckParamsHTTPConfigMethod = "HEAD"
)

func (r SmartShieldEditHealthcheckParamsHTTPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldEditHealthcheckParamsHTTPConfigMethodGet, SmartShieldEditHealthcheckParamsHTTPConfigMethodHead:
		return true
	}
	return false
}

// Parameters specific to TCP health check.
type SmartShieldEditHealthcheckParamsTCPConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[SmartShieldEditHealthcheckParamsTCPConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r SmartShieldEditHealthcheckParamsTCPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type SmartShieldEditHealthcheckParamsTCPConfigMethod string

const (
	SmartShieldEditHealthcheckParamsTCPConfigMethodConnectionEstablished SmartShieldEditHealthcheckParamsTCPConfigMethod = "connection_established"
)

func (r SmartShieldEditHealthcheckParamsTCPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldEditHealthcheckParamsTCPConfigMethodConnectionEstablished:
		return true
	}
	return false
}

type SmartShieldEditHealthcheckResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors,required"`
	Messages []shared.ResponseInfo              `json:"messages,required"`
	Result   SmartShieldEditHealthcheckResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SmartShieldEditHealthcheckResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartShieldEditHealthcheckResponseEnvelopeJSON    `json:"-"`
}

// smartShieldEditHealthcheckResponseEnvelopeJSON contains the JSON metadata for
// the struct [SmartShieldEditHealthcheckResponseEnvelope]
type smartShieldEditHealthcheckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldEditHealthcheckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldEditHealthcheckResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SmartShieldEditHealthcheckResponseEnvelopeSuccess bool

const (
	SmartShieldEditHealthcheckResponseEnvelopeSuccessTrue SmartShieldEditHealthcheckResponseEnvelopeSuccess = true
)

func (r SmartShieldEditHealthcheckResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartShieldEditHealthcheckResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SmartShieldGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SmartShieldGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A consolidated object containing settings from multiple APIs for partial
	// updates.
	Result SmartShieldGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SmartShieldGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartShieldGetResponseEnvelopeJSON    `json:"-"`
}

// smartShieldGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SmartShieldGetResponseEnvelope]
type smartShieldGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SmartShieldGetResponseEnvelopeSuccess bool

const (
	SmartShieldGetResponseEnvelopeSuccessTrue SmartShieldGetResponseEnvelopeSuccess = true
)

func (r SmartShieldGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartShieldGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SmartShieldGetHealthcheckParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SmartShieldGetHealthcheckResponseEnvelope struct {
	Errors   []shared.ResponseInfo             `json:"errors,required"`
	Messages []shared.ResponseInfo             `json:"messages,required"`
	Result   SmartShieldGetHealthcheckResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SmartShieldGetHealthcheckResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartShieldGetHealthcheckResponseEnvelopeJSON    `json:"-"`
}

// smartShieldGetHealthcheckResponseEnvelopeJSON contains the JSON metadata for the
// struct [SmartShieldGetHealthcheckResponseEnvelope]
type smartShieldGetHealthcheckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldGetHealthcheckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldGetHealthcheckResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SmartShieldGetHealthcheckResponseEnvelopeSuccess bool

const (
	SmartShieldGetHealthcheckResponseEnvelopeSuccessTrue SmartShieldGetHealthcheckResponseEnvelopeSuccess = true
)

func (r SmartShieldGetHealthcheckResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartShieldGetHealthcheckResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SmartShieldListHealthchecksParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page. Must be a multiple of 5.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [SmartShieldListHealthchecksParams]'s query parameters as
// `url.Values`.
func (r SmartShieldListHealthchecksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SmartShieldUpdateHealthcheckParams struct {
	// Identifier.
	ZoneID   param.Field[string]                                   `path:"zone_id,required"`
	Errors   param.Field[[]shared.ResponseInfoParam]               `json:"errors,required"`
	Messages param.Field[[]shared.ResponseInfoParam]               `json:"messages,required"`
	Result   param.Field[SmartShieldUpdateHealthcheckParamsResult] `json:"result,required"`
	// Whether the API call was successful.
	Success param.Field[SmartShieldUpdateHealthcheckParamsSuccess] `json:"success,required"`
}

func (r SmartShieldUpdateHealthcheckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SmartShieldUpdateHealthcheckParamsResult struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]SmartShieldUpdateHealthcheckParamsResultCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[SmartShieldUpdateHealthcheckParamsResultHTTPConfig] `json:"http_config"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// If suspended, no health checks are sent to the origin.
	Suspended param.Field[bool] `json:"suspended"`
	// Parameters specific to TCP health check.
	TCPConfig param.Field[SmartShieldUpdateHealthcheckParamsResultTCPConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r SmartShieldUpdateHealthcheckParamsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type SmartShieldUpdateHealthcheckParamsResultCheckRegion string

const (
	SmartShieldUpdateHealthcheckParamsResultCheckRegionWnam       SmartShieldUpdateHealthcheckParamsResultCheckRegion = "WNAM"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionEnam       SmartShieldUpdateHealthcheckParamsResultCheckRegion = "ENAM"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionWeu        SmartShieldUpdateHealthcheckParamsResultCheckRegion = "WEU"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionEeu        SmartShieldUpdateHealthcheckParamsResultCheckRegion = "EEU"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionNsam       SmartShieldUpdateHealthcheckParamsResultCheckRegion = "NSAM"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionSsam       SmartShieldUpdateHealthcheckParamsResultCheckRegion = "SSAM"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionOc         SmartShieldUpdateHealthcheckParamsResultCheckRegion = "OC"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionMe         SmartShieldUpdateHealthcheckParamsResultCheckRegion = "ME"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionNaf        SmartShieldUpdateHealthcheckParamsResultCheckRegion = "NAF"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionSaf        SmartShieldUpdateHealthcheckParamsResultCheckRegion = "SAF"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionIn         SmartShieldUpdateHealthcheckParamsResultCheckRegion = "IN"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionSeas       SmartShieldUpdateHealthcheckParamsResultCheckRegion = "SEAS"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionNeas       SmartShieldUpdateHealthcheckParamsResultCheckRegion = "NEAS"
	SmartShieldUpdateHealthcheckParamsResultCheckRegionAllRegions SmartShieldUpdateHealthcheckParamsResultCheckRegion = "ALL_REGIONS"
)

func (r SmartShieldUpdateHealthcheckParamsResultCheckRegion) IsKnown() bool {
	switch r {
	case SmartShieldUpdateHealthcheckParamsResultCheckRegionWnam, SmartShieldUpdateHealthcheckParamsResultCheckRegionEnam, SmartShieldUpdateHealthcheckParamsResultCheckRegionWeu, SmartShieldUpdateHealthcheckParamsResultCheckRegionEeu, SmartShieldUpdateHealthcheckParamsResultCheckRegionNsam, SmartShieldUpdateHealthcheckParamsResultCheckRegionSsam, SmartShieldUpdateHealthcheckParamsResultCheckRegionOc, SmartShieldUpdateHealthcheckParamsResultCheckRegionMe, SmartShieldUpdateHealthcheckParamsResultCheckRegionNaf, SmartShieldUpdateHealthcheckParamsResultCheckRegionSaf, SmartShieldUpdateHealthcheckParamsResultCheckRegionIn, SmartShieldUpdateHealthcheckParamsResultCheckRegionSeas, SmartShieldUpdateHealthcheckParamsResultCheckRegionNeas, SmartShieldUpdateHealthcheckParamsResultCheckRegionAllRegions:
		return true
	}
	return false
}

// Parameters specific to an HTTP or HTTPS health check.
type SmartShieldUpdateHealthcheckParamsResultHTTPConfig struct {
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
	Header param.Field[map[string][]string] `json:"header"`
	// The HTTP method to use for the health check.
	Method param.Field[SmartShieldUpdateHealthcheckParamsResultHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r SmartShieldUpdateHealthcheckParamsResultHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type SmartShieldUpdateHealthcheckParamsResultHTTPConfigMethod string

const (
	SmartShieldUpdateHealthcheckParamsResultHTTPConfigMethodGet  SmartShieldUpdateHealthcheckParamsResultHTTPConfigMethod = "GET"
	SmartShieldUpdateHealthcheckParamsResultHTTPConfigMethodHead SmartShieldUpdateHealthcheckParamsResultHTTPConfigMethod = "HEAD"
)

func (r SmartShieldUpdateHealthcheckParamsResultHTTPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldUpdateHealthcheckParamsResultHTTPConfigMethodGet, SmartShieldUpdateHealthcheckParamsResultHTTPConfigMethodHead:
		return true
	}
	return false
}

// The current status of the origin server according to the health check.
type SmartShieldUpdateHealthcheckParamsResultStatus string

const (
	SmartShieldUpdateHealthcheckParamsResultStatusUnknown   SmartShieldUpdateHealthcheckParamsResultStatus = "unknown"
	SmartShieldUpdateHealthcheckParamsResultStatusHealthy   SmartShieldUpdateHealthcheckParamsResultStatus = "healthy"
	SmartShieldUpdateHealthcheckParamsResultStatusUnhealthy SmartShieldUpdateHealthcheckParamsResultStatus = "unhealthy"
	SmartShieldUpdateHealthcheckParamsResultStatusSuspended SmartShieldUpdateHealthcheckParamsResultStatus = "suspended"
)

func (r SmartShieldUpdateHealthcheckParamsResultStatus) IsKnown() bool {
	switch r {
	case SmartShieldUpdateHealthcheckParamsResultStatusUnknown, SmartShieldUpdateHealthcheckParamsResultStatusHealthy, SmartShieldUpdateHealthcheckParamsResultStatusUnhealthy, SmartShieldUpdateHealthcheckParamsResultStatusSuspended:
		return true
	}
	return false
}

// Parameters specific to TCP health check.
type SmartShieldUpdateHealthcheckParamsResultTCPConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[SmartShieldUpdateHealthcheckParamsResultTCPConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r SmartShieldUpdateHealthcheckParamsResultTCPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type SmartShieldUpdateHealthcheckParamsResultTCPConfigMethod string

const (
	SmartShieldUpdateHealthcheckParamsResultTCPConfigMethodConnectionEstablished SmartShieldUpdateHealthcheckParamsResultTCPConfigMethod = "connection_established"
)

func (r SmartShieldUpdateHealthcheckParamsResultTCPConfigMethod) IsKnown() bool {
	switch r {
	case SmartShieldUpdateHealthcheckParamsResultTCPConfigMethodConnectionEstablished:
		return true
	}
	return false
}

// Whether the API call was successful.
type SmartShieldUpdateHealthcheckParamsSuccess bool

const (
	SmartShieldUpdateHealthcheckParamsSuccessTrue SmartShieldUpdateHealthcheckParamsSuccess = true
)

func (r SmartShieldUpdateHealthcheckParamsSuccess) IsKnown() bool {
	switch r {
	case SmartShieldUpdateHealthcheckParamsSuccessTrue:
		return true
	}
	return false
}

type SmartShieldUpdateHealthcheckResponseEnvelope struct {
	Errors   []shared.ResponseInfo                `json:"errors,required"`
	Messages []shared.ResponseInfo                `json:"messages,required"`
	Result   SmartShieldUpdateHealthcheckResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SmartShieldUpdateHealthcheckResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartShieldUpdateHealthcheckResponseEnvelopeJSON    `json:"-"`
}

// smartShieldUpdateHealthcheckResponseEnvelopeJSON contains the JSON metadata for
// the struct [SmartShieldUpdateHealthcheckResponseEnvelope]
type smartShieldUpdateHealthcheckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartShieldUpdateHealthcheckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartShieldUpdateHealthcheckResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SmartShieldUpdateHealthcheckResponseEnvelopeSuccess bool

const (
	SmartShieldUpdateHealthcheckResponseEnvelopeSuccessTrue SmartShieldUpdateHealthcheckResponseEnvelopeSuccess = true
)

func (r SmartShieldUpdateHealthcheckResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartShieldUpdateHealthcheckResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
