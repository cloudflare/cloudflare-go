// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// GatewayConfigurationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayConfigurationService] method instead.
type GatewayConfigurationService struct {
	Options           []option.RequestOption
	CustomCertificate *GatewayConfigurationCustomCertificateService
}

// NewGatewayConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGatewayConfigurationService(opts ...option.RequestOption) (r *GatewayConfigurationService) {
	r = &GatewayConfigurationService{}
	r.Options = opts
	r.CustomCertificate = NewGatewayConfigurationCustomCertificateService(opts...)
	return
}

// Updates the current Zero Trust account configuration.
func (r *GatewayConfigurationService) Update(ctx context.Context, params GatewayConfigurationUpdateParams, opts ...option.RequestOption) (res *GatewayConfigurationUpdateResponse, err error) {
	var env GatewayConfigurationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patches the current Zero Trust account configuration. This endpoint can update a
// single subcollection of settings such as `antivirus`, `tls_decrypt`,
// `activity_log`, `block_page`, `browser_isolation`, `fips`, `body_scanning`, or
// `certificate`, without updating the entire configuration object. Returns an
// error if any collection of settings is not properly configured.
func (r *GatewayConfigurationService) Edit(ctx context.Context, params GatewayConfigurationEditParams, opts ...option.RequestOption) (res *GatewayConfigurationEditResponse, err error) {
	var env GatewayConfigurationEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the current Zero Trust account configuration.
func (r *GatewayConfigurationService) Get(ctx context.Context, query GatewayConfigurationGetParams, opts ...option.RequestOption) (res *GatewayConfigurationGetResponse, err error) {
	var env GatewayConfigurationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Custom certificate settings for BYO-PKI. (deprecated and replaced by
// `certificate`)
//
// Deprecated: deprecated
type CustomCertificateSettings struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                        `json:"binding_status"`
	UpdatedAt     time.Time                     `json:"updated_at" format:"date-time"`
	JSON          customCertificateSettingsJSON `json:"-"`
}

// customCertificateSettingsJSON contains the JSON metadata for the struct
// [CustomCertificateSettings]
type customCertificateSettingsJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomCertificateSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCertificateSettingsJSON) RawJSON() string {
	return r.raw
}

// Account settings
type GatewayConfigurationUpdateResponse struct {
	CreatedAt time.Time                              `json:"created_at" format:"date-time"`
	Settings  interface{}                            `json:"settings"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationUpdateResponseJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseJSON contains the JSON metadata for the struct
// [GatewayConfigurationUpdateResponse]
type gatewayConfigurationUpdateResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Account settings
type GatewayConfigurationEditResponse struct {
	CreatedAt time.Time                            `json:"created_at" format:"date-time"`
	Settings  interface{}                          `json:"settings"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationEditResponseJSON `json:"-"`
}

// gatewayConfigurationEditResponseJSON contains the JSON metadata for the struct
// [GatewayConfigurationEditResponse]
type gatewayConfigurationEditResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseJSON) RawJSON() string {
	return r.raw
}

// Account settings
type GatewayConfigurationGetResponse struct {
	CreatedAt time.Time                           `json:"created_at" format:"date-time"`
	Settings  interface{}                         `json:"settings"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationGetResponseJSON `json:"-"`
}

// gatewayConfigurationGetResponseJSON contains the JSON metadata for the struct
// [GatewayConfigurationGetResponse]
type gatewayConfigurationGetResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
}

type GatewayConfigurationUpdateParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Settings  param.Field[interface{}] `json:"settings"`
}

func (r GatewayConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationUpdateResponseEnvelope struct {
	Errors   []interface{} `json:"errors,required"`
	Messages []interface{} `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayConfigurationUpdateResponseEnvelopeSuccess `json:"success,required"`
	// Account settings
	Result GatewayConfigurationUpdateResponse             `json:"result"`
	JSON   gatewayConfigurationUpdateResponseEnvelopeJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayConfigurationUpdateResponseEnvelope]
type gatewayConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayConfigurationUpdateResponseEnvelopeSuccess bool

const (
	GatewayConfigurationUpdateResponseEnvelopeSuccessTrue GatewayConfigurationUpdateResponseEnvelopeSuccess = true
)

func (r GatewayConfigurationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayConfigurationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayConfigurationEditParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Settings  param.Field[interface{}] `json:"settings"`
}

func (r GatewayConfigurationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationEditResponseEnvelope struct {
	Errors   []interface{} `json:"errors,required"`
	Messages []interface{} `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayConfigurationEditResponseEnvelopeSuccess `json:"success,required"`
	// Account settings
	Result GatewayConfigurationEditResponse             `json:"result"`
	JSON   gatewayConfigurationEditResponseEnvelopeJSON `json:"-"`
}

// gatewayConfigurationEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayConfigurationEditResponseEnvelope]
type gatewayConfigurationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayConfigurationEditResponseEnvelopeSuccess bool

const (
	GatewayConfigurationEditResponseEnvelopeSuccessTrue GatewayConfigurationEditResponseEnvelopeSuccess = true
)

func (r GatewayConfigurationEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayConfigurationEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayConfigurationGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayConfigurationGetResponseEnvelope struct {
	Errors   []interface{} `json:"errors,required"`
	Messages []interface{} `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	// Account settings
	Result GatewayConfigurationGetResponse             `json:"result"`
	JSON   gatewayConfigurationGetResponseEnvelopeJSON `json:"-"`
}

// gatewayConfigurationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayConfigurationGetResponseEnvelope]
type gatewayConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayConfigurationGetResponseEnvelopeSuccess bool

const (
	GatewayConfigurationGetResponseEnvelopeSuccessTrue GatewayConfigurationGetResponseEnvelopeSuccess = true
)

func (r GatewayConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
