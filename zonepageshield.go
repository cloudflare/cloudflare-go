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

// ZonePageShieldService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZonePageShieldService] method
// instead.
type ZonePageShieldService struct {
	Options     []option.RequestOption
	Connections *ZonePageShieldConnectionService
	Scripts     *ZonePageShieldScriptService
	Policies    *ZonePageShieldPolicyService
}

// NewZonePageShieldService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZonePageShieldService(opts ...option.RequestOption) (r *ZonePageShieldService) {
	r = &ZonePageShieldService{}
	r.Options = opts
	r.Connections = NewZonePageShieldConnectionService(opts...)
	r.Scripts = NewZonePageShieldScriptService(opts...)
	r.Policies = NewZonePageShieldPolicyService(opts...)
	return
}

// Fetches the Page Shield settings.
func (r *ZonePageShieldService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZonePageShieldListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates Page Shield settings.
func (r *ZonePageShieldService) PageShieldUpdatePageShieldSettings(ctx context.Context, zoneID string, body ZonePageShieldPageShieldUpdatePageShieldSettingsParams, opts ...option.RequestOption) (res *ZonePageShieldPageShieldUpdatePageShieldSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZonePageShieldListResponse struct {
	Errors   []ZonePageShieldListResponseError   `json:"errors"`
	Messages []ZonePageShieldListResponseMessage `json:"messages"`
	Result   ZonePageShieldListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZonePageShieldListResponseSuccess `json:"success"`
	JSON    zonePageShieldListResponseJSON    `json:"-"`
}

// zonePageShieldListResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldListResponse]
type zonePageShieldListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldListResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zonePageShieldListResponseErrorJSON `json:"-"`
}

// zonePageShieldListResponseErrorJSON contains the JSON metadata for the struct
// [ZonePageShieldListResponseError]
type zonePageShieldListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldListResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zonePageShieldListResponseMessageJSON `json:"-"`
}

// zonePageShieldListResponseMessageJSON contains the JSON metadata for the struct
// [ZonePageShieldListResponseMessage]
type zonePageShieldListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldListResponseResult struct {
	// When true, indicates that Page Shield is enabled.
	Enabled bool `json:"enabled"`
	// The timestamp of when Page Shield was last updated.
	UpdatedAt string `json:"updated_at"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint bool `json:"use_cloudflare_reporting_endpoint"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath bool                                 `json:"use_connection_url_path"`
	JSON                 zonePageShieldListResponseResultJSON `json:"-"`
}

// zonePageShieldListResponseResultJSON contains the JSON metadata for the struct
// [ZonePageShieldListResponseResult]
type zonePageShieldListResponseResultJSON struct {
	Enabled                        apijson.Field
	UpdatedAt                      apijson.Field
	UseCloudflareReportingEndpoint apijson.Field
	UseConnectionURLPath           apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *ZonePageShieldListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZonePageShieldListResponseSuccess bool

const (
	ZonePageShieldListResponseSuccessTrue ZonePageShieldListResponseSuccess = true
)

type ZonePageShieldPageShieldUpdatePageShieldSettingsResponse struct {
	Errors   []ZonePageShieldPageShieldUpdatePageShieldSettingsResponseError   `json:"errors"`
	Messages []ZonePageShieldPageShieldUpdatePageShieldSettingsResponseMessage `json:"messages"`
	Result   ZonePageShieldPageShieldUpdatePageShieldSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZonePageShieldPageShieldUpdatePageShieldSettingsResponseSuccess `json:"success"`
	JSON    zonePageShieldPageShieldUpdatePageShieldSettingsResponseJSON    `json:"-"`
}

// zonePageShieldPageShieldUpdatePageShieldSettingsResponseJSON contains the JSON
// metadata for the struct
// [ZonePageShieldPageShieldUpdatePageShieldSettingsResponse]
type zonePageShieldPageShieldUpdatePageShieldSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPageShieldUpdatePageShieldSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldPageShieldUpdatePageShieldSettingsResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zonePageShieldPageShieldUpdatePageShieldSettingsResponseErrorJSON `json:"-"`
}

// zonePageShieldPageShieldUpdatePageShieldSettingsResponseErrorJSON contains the
// JSON metadata for the struct
// [ZonePageShieldPageShieldUpdatePageShieldSettingsResponseError]
type zonePageShieldPageShieldUpdatePageShieldSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPageShieldUpdatePageShieldSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldPageShieldUpdatePageShieldSettingsResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zonePageShieldPageShieldUpdatePageShieldSettingsResponseMessageJSON `json:"-"`
}

// zonePageShieldPageShieldUpdatePageShieldSettingsResponseMessageJSON contains the
// JSON metadata for the struct
// [ZonePageShieldPageShieldUpdatePageShieldSettingsResponseMessage]
type zonePageShieldPageShieldUpdatePageShieldSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPageShieldUpdatePageShieldSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldPageShieldUpdatePageShieldSettingsResponseResult struct {
	// When true, indicates that Page Shield is enabled.
	Enabled bool `json:"enabled"`
	// The timestamp of when Page Shield was last updated.
	UpdatedAt string `json:"updated_at"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint bool `json:"use_cloudflare_reporting_endpoint"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath bool                                                               `json:"use_connection_url_path"`
	JSON                 zonePageShieldPageShieldUpdatePageShieldSettingsResponseResultJSON `json:"-"`
}

// zonePageShieldPageShieldUpdatePageShieldSettingsResponseResultJSON contains the
// JSON metadata for the struct
// [ZonePageShieldPageShieldUpdatePageShieldSettingsResponseResult]
type zonePageShieldPageShieldUpdatePageShieldSettingsResponseResultJSON struct {
	Enabled                        apijson.Field
	UpdatedAt                      apijson.Field
	UseCloudflareReportingEndpoint apijson.Field
	UseConnectionURLPath           apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *ZonePageShieldPageShieldUpdatePageShieldSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZonePageShieldPageShieldUpdatePageShieldSettingsResponseSuccess bool

const (
	ZonePageShieldPageShieldUpdatePageShieldSettingsResponseSuccessTrue ZonePageShieldPageShieldUpdatePageShieldSettingsResponseSuccess = true
)

type ZonePageShieldPageShieldUpdatePageShieldSettingsParams struct {
	// When true, indicates that Page Shield is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint param.Field[bool] `json:"use_cloudflare_reporting_endpoint"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath param.Field[bool] `json:"use_connection_url_path"`
}

func (r ZonePageShieldPageShieldUpdatePageShieldSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
