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

// PageShieldService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPageShieldService] method instead.
type PageShieldService struct {
	Options     []option.RequestOption
	Connections *PageShieldConnectionService
	Scripts     *PageShieldScriptService
}

// NewPageShieldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageShieldService(opts ...option.RequestOption) (r *PageShieldService) {
	r = &PageShieldService{}
	r.Options = opts
	r.Connections = NewPageShieldConnectionService(opts...)
	r.Scripts = NewPageShieldScriptService(opts...)
	return
}

// Fetches the Page Shield settings.
func (r *PageShieldService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *PageShieldListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageShieldListResponseEnvelope
	path := fmt.Sprintf("zones/%s/page_shield", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates Page Shield settings.
func (r *PageShieldService) PageShieldUpdatePageShieldSettings(ctx context.Context, zoneID string, body PageShieldPageShieldUpdatePageShieldSettingsParams, opts ...option.RequestOption) (res *PageShieldPageShieldUpdatePageShieldSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelope
	path := fmt.Sprintf("zones/%s/page_shield", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageShieldListResponse struct {
	// When true, indicates that Page Shield is enabled.
	Enabled bool `json:"enabled"`
	// The timestamp of when Page Shield was last updated.
	UpdatedAt string `json:"updated_at"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint bool `json:"use_cloudflare_reporting_endpoint"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath bool                       `json:"use_connection_url_path"`
	JSON                 pageShieldListResponseJSON `json:"-"`
}

// pageShieldListResponseJSON contains the JSON metadata for the struct
// [PageShieldListResponse]
type pageShieldListResponseJSON struct {
	Enabled                        apijson.Field
	UpdatedAt                      apijson.Field
	UseCloudflareReportingEndpoint apijson.Field
	UseConnectionURLPath           apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *PageShieldListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldPageShieldUpdatePageShieldSettingsResponse struct {
	// When true, indicates that Page Shield is enabled.
	Enabled bool `json:"enabled"`
	// The timestamp of when Page Shield was last updated.
	UpdatedAt string `json:"updated_at"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint bool `json:"use_cloudflare_reporting_endpoint"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath bool                                                     `json:"use_connection_url_path"`
	JSON                 pageShieldPageShieldUpdatePageShieldSettingsResponseJSON `json:"-"`
}

// pageShieldPageShieldUpdatePageShieldSettingsResponseJSON contains the JSON
// metadata for the struct [PageShieldPageShieldUpdatePageShieldSettingsResponse]
type pageShieldPageShieldUpdatePageShieldSettingsResponseJSON struct {
	Enabled                        apijson.Field
	UpdatedAt                      apijson.Field
	UseCloudflareReportingEndpoint apijson.Field
	UseConnectionURLPath           apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *PageShieldPageShieldUpdatePageShieldSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldListResponseEnvelope struct {
	Errors   []PageShieldListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageShieldListResponseEnvelopeMessages `json:"messages,required"`
	Result   PageShieldListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PageShieldListResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageShieldListResponseEnvelopeJSON    `json:"-"`
}

// pageShieldListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageShieldListResponseEnvelope]
type pageShieldListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldListResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    pageShieldListResponseEnvelopeErrorsJSON `json:"-"`
}

// pageShieldListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PageShieldListResponseEnvelopeErrors]
type pageShieldListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldListResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    pageShieldListResponseEnvelopeMessagesJSON `json:"-"`
}

// pageShieldListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageShieldListResponseEnvelopeMessages]
type pageShieldListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageShieldListResponseEnvelopeSuccess bool

const (
	PageShieldListResponseEnvelopeSuccessTrue PageShieldListResponseEnvelopeSuccess = true
)

type PageShieldPageShieldUpdatePageShieldSettingsParams struct {
	// When true, indicates that Page Shield is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint param.Field[bool] `json:"use_cloudflare_reporting_endpoint"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath param.Field[bool] `json:"use_connection_url_path"`
}

func (r PageShieldPageShieldUpdatePageShieldSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelope struct {
	Errors   []PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeMessages `json:"messages,required"`
	Result   PageShieldPageShieldUpdatePageShieldSettingsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeJSON    `json:"-"`
}

// pageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelope]
type pageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    pageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeErrorsJSON `json:"-"`
}

// pageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeErrors]
type pageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    pageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeMessagesJSON `json:"-"`
}

// pageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeMessages]
type pageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeSuccess bool

const (
	PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeSuccessTrue PageShieldPageShieldUpdatePageShieldSettingsResponseEnvelopeSuccess = true
)
