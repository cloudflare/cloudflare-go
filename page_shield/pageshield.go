// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// PageShieldService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPageShieldService] method instead.
type PageShieldService struct {
	Options     []option.RequestOption
	Policies    *PolicyService
	Connections *ConnectionService
	Scripts     *ScriptService
}

// NewPageShieldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageShieldService(opts ...option.RequestOption) (r *PageShieldService) {
	r = &PageShieldService{}
	r.Options = opts
	r.Policies = NewPolicyService(opts...)
	r.Connections = NewConnectionService(opts...)
	r.Scripts = NewScriptService(opts...)
	return
}

// Updates Page Shield settings.
func (r *PageShieldService) Update(ctx context.Context, params PageShieldUpdateParams, opts ...option.RequestOption) (res *PageShieldUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageShieldUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/page_shield", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the Page Shield settings.
func (r *PageShieldService) Get(ctx context.Context, query PageShieldGetParams, opts ...option.RequestOption) (res *PageShieldSetting, err error) {
	opts = append(r.Options[:], opts...)
	var env PageShieldGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/page_shield", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageShieldSetting struct {
	// When true, indicates that Page Shield is enabled.
	Enabled bool `json:"enabled"`
	// The timestamp of when Page Shield was last updated.
	UpdatedAt string `json:"updated_at"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint bool `json:"use_cloudflare_reporting_endpoint"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath bool                  `json:"use_connection_url_path"`
	JSON                 pageShieldSettingJSON `json:"-"`
}

// pageShieldSettingJSON contains the JSON metadata for the struct
// [PageShieldSetting]
type pageShieldSettingJSON struct {
	Enabled                        apijson.Field
	UpdatedAt                      apijson.Field
	UseCloudflareReportingEndpoint apijson.Field
	UseConnectionURLPath           apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *PageShieldSetting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldSettingJSON) RawJSON() string {
	return r.raw
}

type PageShieldUpdateResponse struct {
	// When true, indicates that Page Shield is enabled.
	Enabled bool `json:"enabled"`
	// The timestamp of when Page Shield was last updated.
	UpdatedAt string `json:"updated_at"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint bool `json:"use_cloudflare_reporting_endpoint"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath bool                         `json:"use_connection_url_path"`
	JSON                 pageShieldUpdateResponseJSON `json:"-"`
}

// pageShieldUpdateResponseJSON contains the JSON metadata for the struct
// [PageShieldUpdateResponse]
type pageShieldUpdateResponseJSON struct {
	Enabled                        apijson.Field
	UpdatedAt                      apijson.Field
	UseCloudflareReportingEndpoint apijson.Field
	UseConnectionURLPath           apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *PageShieldUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PageShieldUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// When true, indicates that Page Shield is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint param.Field[bool] `json:"use_cloudflare_reporting_endpoint"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath param.Field[bool] `json:"use_connection_url_path"`
}

func (r PageShieldUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageShieldUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   PageShieldUpdateResponse     `json:"result,required"`
	// Whether the API call was successful
	Success PageShieldUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageShieldUpdateResponseEnvelopeJSON    `json:"-"`
}

// pageShieldUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageShieldUpdateResponseEnvelope]
type pageShieldUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageShieldUpdateResponseEnvelopeSuccess bool

const (
	PageShieldUpdateResponseEnvelopeSuccessTrue PageShieldUpdateResponseEnvelopeSuccess = true
)

func (r PageShieldUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageShieldUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageShieldGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageShieldGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   PageShieldSetting            `json:"result,required"`
	// Whether the API call was successful
	Success PageShieldGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageShieldGetResponseEnvelopeJSON    `json:"-"`
}

// pageShieldGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageShieldGetResponseEnvelope]
type pageShieldGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageShieldGetResponseEnvelopeSuccess bool

const (
	PageShieldGetResponseEnvelopeSuccessTrue PageShieldGetResponseEnvelopeSuccess = true
)

func (r PageShieldGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageShieldGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
