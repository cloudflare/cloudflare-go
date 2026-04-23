// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package google_tag_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// ConfigService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	return
}

// Updates the Google Tag Gateway configuration for a zone.
func (r *ConfigService) Update(ctx context.Context, params ConfigUpdateParams, opts ...option.RequestOption) (res *Config, err error) {
	var env ConfigUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/google-tag-gateway/config", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets the Google Tag Gateway configuration for a zone.
func (r *ConfigService) Get(ctx context.Context, query ConfigGetParams, opts ...option.RequestOption) (res *Config, err error) {
	var env ConfigGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/google-tag-gateway/config", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Google Tag Gateway configuration for a zone.
type Config struct {
	// Enables or disables Google Tag Gateway for this zone.
	Enabled bool `json:"enabled" api:"required"`
	// Specifies the endpoint path for proxying Google Tag Manager requests. Use an
	// absolute path starting with '/', with no nested paths and alphanumeric
	// characters only (e.g. /metrics).
	Endpoint string `json:"endpoint" api:"required"`
	// Hides the original client IP address from Google when enabled.
	HideOriginalIP bool `json:"hideOriginalIp" api:"required"`
	// Specify the Google Tag Manager container or measurement ID (e.g. GTM-XXXXXXX or
	// G-XXXXXXXXXX).
	MeasurementID string `json:"measurementId" api:"required"`
	// Set up the associated Google Tag on the zone automatically when enabled.
	SetUpTag bool       `json:"setUpTag" api:"nullable"`
	JSON     configJSON `json:"-"`
}

// configJSON contains the JSON metadata for the struct [Config]
type configJSON struct {
	Enabled        apijson.Field
	Endpoint       apijson.Field
	HideOriginalIP apijson.Field
	MeasurementID  apijson.Field
	SetUpTag       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Config) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configJSON) RawJSON() string {
	return r.raw
}

// Google Tag Gateway configuration for a zone.
type ConfigParam struct {
	// Enables or disables Google Tag Gateway for this zone.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Specifies the endpoint path for proxying Google Tag Manager requests. Use an
	// absolute path starting with '/', with no nested paths and alphanumeric
	// characters only (e.g. /metrics).
	Endpoint param.Field[string] `json:"endpoint" api:"required"`
	// Hides the original client IP address from Google when enabled.
	HideOriginalIP param.Field[bool] `json:"hideOriginalIp" api:"required"`
	// Specify the Google Tag Manager container or measurement ID (e.g. GTM-XXXXXXX or
	// G-XXXXXXXXXX).
	MeasurementID param.Field[string] `json:"measurementId" api:"required"`
	// Set up the associated Google Tag on the zone automatically when enabled.
	SetUpTag param.Field[bool] `json:"setUpTag"`
}

func (r ConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Google Tag Gateway configuration for a zone.
	Config ConfigParam `json:"config" api:"required"`
}

func (r ConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Config)
}

type ConfigUpdateResponseEnvelope struct {
	Errors   []ConfigUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ConfigUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success ConfigUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	// Google Tag Gateway configuration for a zone.
	Result Config                           `json:"result"`
	JSON   configUpdateResponseEnvelopeJSON `json:"-"`
}

// configUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseEnvelope]
type configUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ConfigUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             configUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// configUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseEnvelopeErrors]
type configUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    configUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// configUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ConfigUpdateResponseEnvelopeErrorsSource]
type configUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ConfigUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             configUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// configUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConfigUpdateResponseEnvelopeMessages]
type configUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    configUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// configUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ConfigUpdateResponseEnvelopeMessagesSource]
type configUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConfigUpdateResponseEnvelopeSuccess bool

const (
	ConfigUpdateResponseEnvelopeSuccessTrue ConfigUpdateResponseEnvelopeSuccess = true
)

func (r ConfigUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type ConfigGetResponseEnvelope struct {
	Errors   []ConfigGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ConfigGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success ConfigGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// Google Tag Gateway configuration for a zone.
	Result Config                        `json:"result"`
	JSON   configGetResponseEnvelopeJSON `json:"-"`
}

// configGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigGetResponseEnvelope]
type configGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code" api:"required"`
	Message          string                                `json:"message" api:"required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           ConfigGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             configGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// configGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConfigGetResponseEnvelopeErrors]
type configGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    configGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// configGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ConfigGetResponseEnvelopeErrorsSource]
type configGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code" api:"required"`
	Message          string                                  `json:"message" api:"required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           ConfigGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             configGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// configGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ConfigGetResponseEnvelopeMessages]
type configGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    configGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// configGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [ConfigGetResponseEnvelopeMessagesSource]
type configGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConfigGetResponseEnvelopeSuccess bool

const (
	ConfigGetResponseEnvelopeSuccessTrue ConfigGetResponseEnvelopeSuccess = true
)

func (r ConfigGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
