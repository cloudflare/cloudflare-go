// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// PageShieldService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageShieldService] method instead.
type PageShieldService struct {
	Options     []option.RequestOption
	Policies    *PolicyService
	Connections *ConnectionService
	Scripts     *ScriptService
	Cookies     *CookieService
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
	r.Cookies = NewCookieService(opts...)
	return
}

// Updates Page Shield settings.
func (r *PageShieldService) Update(ctx context.Context, params PageShieldUpdateParams, opts ...option.RequestOption) (res *PageShieldUpdateResponse, err error) {
	var env PageShieldUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the Page Shield settings.
func (r *PageShieldService) Get(ctx context.Context, query PageShieldGetParams, opts ...option.RequestOption) (res *Setting, err error) {
	var env PageShieldGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Setting struct {
	// When true, indicates that Page Shield is enabled.
	Enabled bool `json:"enabled,required"`
	// The timestamp of when Page Shield was last updated.
	UpdatedAt string `json:"updated_at,required"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint bool `json:"use_cloudflare_reporting_endpoint,required"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath bool        `json:"use_connection_url_path,required"`
	JSON                 settingJSON `json:"-"`
}

// settingJSON contains the JSON metadata for the struct [Setting]
type settingJSON struct {
	Enabled                        apijson.Field
	UpdatedAt                      apijson.Field
	UseCloudflareReportingEndpoint apijson.Field
	UseConnectionURLPath           apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *Setting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingJSON) RawJSON() string {
	return r.raw
}

type PageShieldUpdateResponse struct {
	// When true, indicates that Page Shield is enabled.
	Enabled bool `json:"enabled,required"`
	// The timestamp of when Page Shield was last updated.
	UpdatedAt string `json:"updated_at,required"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint bool `json:"use_cloudflare_reporting_endpoint,required"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath bool                         `json:"use_connection_url_path,required"`
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
	// Whether the API call was successful
	Success  PageShieldUpdateResponseEnvelopeSuccess    `json:"success,required"`
	Errors   []PageShieldUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []PageShieldUpdateResponseEnvelopeMessages `json:"messages"`
	Result   PageShieldUpdateResponse                   `json:"result"`
	JSON     pageShieldUpdateResponseEnvelopeJSON       `json:"-"`
}

// pageShieldUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageShieldUpdateResponseEnvelope]
type pageShieldUpdateResponseEnvelopeJSON struct {
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
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

type PageShieldUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    pageShieldUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// pageShieldUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PageShieldUpdateResponseEnvelopeErrors]
type pageShieldUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageShieldUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    pageShieldUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// pageShieldUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageShieldUpdateResponseEnvelopeMessages]
type pageShieldUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type PageShieldGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageShieldGetResponseEnvelope struct {
	// Whether the API call was successful
	Success  PageShieldGetResponseEnvelopeSuccess    `json:"success,required"`
	Errors   []PageShieldGetResponseEnvelopeErrors   `json:"errors"`
	Messages []PageShieldGetResponseEnvelopeMessages `json:"messages"`
	Result   Setting                                 `json:"result,nullable"`
	JSON     pageShieldGetResponseEnvelopeJSON       `json:"-"`
}

// pageShieldGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageShieldGetResponseEnvelope]
type pageShieldGetResponseEnvelopeJSON struct {
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
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

type PageShieldGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    pageShieldGetResponseEnvelopeErrorsJSON `json:"-"`
}

// pageShieldGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PageShieldGetResponseEnvelopeErrors]
type pageShieldGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageShieldGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    pageShieldGetResponseEnvelopeMessagesJSON `json:"-"`
}

// pageShieldGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageShieldGetResponseEnvelopeMessages]
type pageShieldGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
