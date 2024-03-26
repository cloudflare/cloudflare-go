// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// GatewayAuditSSHSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayAuditSSHSettingService]
// method instead.
type GatewayAuditSSHSettingService struct {
	Options []option.RequestOption
}

// NewGatewayAuditSSHSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGatewayAuditSSHSettingService(opts ...option.RequestOption) (r *GatewayAuditSSHSettingService) {
	r = &GatewayAuditSSHSettingService{}
	r.Options = opts
	return
}

// Updates Zero Trust Audit SSH settings.
func (r *GatewayAuditSSHSettingService) Update(ctx context.Context, params GatewayAuditSSHSettingUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGatewaySettings, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayAuditSSHSettingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/audit_ssh_settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get all Zero Trust Audit SSH settings for an account.
func (r *GatewayAuditSSHSettingService) Get(ctx context.Context, query GatewayAuditSSHSettingGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewaySettings, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayAuditSSHSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/audit_ssh_settings", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGatewaySettings struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// SSH encryption public key
	PublicKey string `json:"public_key"`
	// Seed ID
	SeedID    string                       `json:"seed_id"`
	UpdatedAt time.Time                    `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewaySettingsJSON `json:"-"`
}

// zeroTrustGatewaySettingsJSON contains the JSON metadata for the struct
// [ZeroTrustGatewaySettings]
type zeroTrustGatewaySettingsJSON struct {
	CreatedAt   apijson.Field
	PublicKey   apijson.Field
	SeedID      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewaySettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewaySettingsJSON) RawJSON() string {
	return r.raw
}

type GatewayAuditSSHSettingUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// SSH encryption public key
	PublicKey param.Field[string] `json:"public_key,required"`
	// Seed ID
	SeedID param.Field[string] `json:"seed_id"`
}

func (r GatewayAuditSSHSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayAuditSSHSettingUpdateResponseEnvelope struct {
	Errors   []GatewayAuditSSHSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayAuditSSHSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewaySettings                               `json:"result,required"`
	// Whether the API call was successful
	Success GatewayAuditSSHSettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayAuditSSHSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayAuditSSHSettingUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayAuditSSHSettingUpdateResponseEnvelope]
type gatewayAuditSSHSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAuditSSHSettingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayAuditSSHSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    gatewayAuditSSHSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayAuditSSHSettingUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [GatewayAuditSSHSettingUpdateResponseEnvelopeErrors]
type gatewayAuditSSHSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAuditSSHSettingUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayAuditSSHSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    gatewayAuditSSHSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayAuditSSHSettingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [GatewayAuditSSHSettingUpdateResponseEnvelopeMessages]
type gatewayAuditSSHSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAuditSSHSettingUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayAuditSSHSettingUpdateResponseEnvelopeSuccess bool

const (
	GatewayAuditSSHSettingUpdateResponseEnvelopeSuccessTrue GatewayAuditSSHSettingUpdateResponseEnvelopeSuccess = true
)

func (r GatewayAuditSSHSettingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayAuditSSHSettingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayAuditSSHSettingGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayAuditSSHSettingGetResponseEnvelope struct {
	Errors   []GatewayAuditSSHSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayAuditSSHSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewaySettings                            `json:"result,required"`
	// Whether the API call was successful
	Success GatewayAuditSSHSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayAuditSSHSettingGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayAuditSSHSettingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayAuditSSHSettingGetResponseEnvelope]
type gatewayAuditSSHSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAuditSSHSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayAuditSSHSettingGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    gatewayAuditSSHSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayAuditSSHSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [GatewayAuditSSHSettingGetResponseEnvelopeErrors]
type gatewayAuditSSHSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAuditSSHSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayAuditSSHSettingGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    gatewayAuditSSHSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayAuditSSHSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [GatewayAuditSSHSettingGetResponseEnvelopeMessages]
type gatewayAuditSSHSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAuditSSHSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayAuditSSHSettingGetResponseEnvelopeSuccess bool

const (
	GatewayAuditSSHSettingGetResponseEnvelopeSuccessTrue GatewayAuditSSHSettingGetResponseEnvelopeSuccess = true
)

func (r GatewayAuditSSHSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayAuditSSHSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
