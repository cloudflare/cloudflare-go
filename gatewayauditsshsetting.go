// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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

// Get all Zero Trust Audit SSH settings for an account.
func (r *GatewayAuditSSHSettingService) Get(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *GatewayAuditSSHSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayAuditSSHSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/audit_ssh_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates Zero Trust Audit SSH settings.
func (r *GatewayAuditSSHSettingService) Replace(ctx context.Context, accountID interface{}, body GatewayAuditSSHSettingReplaceParams, opts ...option.RequestOption) (res *GatewayAuditSSHSettingReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayAuditSSHSettingReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/audit_ssh_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayAuditSSHSettingGetResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// SSH encryption public key
	PublicKey string `json:"public_key"`
	// Seed ID
	SeedID    string                                `json:"seed_id"`
	UpdatedAt time.Time                             `json:"updated_at" format:"date-time"`
	JSON      gatewayAuditSSHSettingGetResponseJSON `json:"-"`
}

// gatewayAuditSSHSettingGetResponseJSON contains the JSON metadata for the struct
// [GatewayAuditSSHSettingGetResponse]
type gatewayAuditSSHSettingGetResponseJSON struct {
	CreatedAt   apijson.Field
	PublicKey   apijson.Field
	SeedID      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAuditSSHSettingReplaceResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// SSH encryption public key
	PublicKey string `json:"public_key"`
	// Seed ID
	SeedID    string                                    `json:"seed_id"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      gatewayAuditSSHSettingReplaceResponseJSON `json:"-"`
}

// gatewayAuditSSHSettingReplaceResponseJSON contains the JSON metadata for the
// struct [GatewayAuditSSHSettingReplaceResponse]
type gatewayAuditSSHSettingReplaceResponseJSON struct {
	CreatedAt   apijson.Field
	PublicKey   apijson.Field
	SeedID      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAuditSSHSettingGetResponseEnvelope struct {
	Errors   []GatewayAuditSSHSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayAuditSSHSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayAuditSSHSettingGetResponse                   `json:"result,required"`
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

// Whether the API call was successful
type GatewayAuditSSHSettingGetResponseEnvelopeSuccess bool

const (
	GatewayAuditSSHSettingGetResponseEnvelopeSuccessTrue GatewayAuditSSHSettingGetResponseEnvelopeSuccess = true
)

type GatewayAuditSSHSettingReplaceParams struct {
	// SSH encryption public key
	PublicKey param.Field[string] `json:"public_key,required"`
	// Seed ID
	SeedID param.Field[string] `json:"seed_id"`
}

func (r GatewayAuditSSHSettingReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayAuditSSHSettingReplaceResponseEnvelope struct {
	Errors   []GatewayAuditSSHSettingReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayAuditSSHSettingReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayAuditSSHSettingReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayAuditSSHSettingReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayAuditSSHSettingReplaceResponseEnvelopeJSON    `json:"-"`
}

// gatewayAuditSSHSettingReplaceResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayAuditSSHSettingReplaceResponseEnvelope]
type gatewayAuditSSHSettingReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAuditSSHSettingReplaceResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    gatewayAuditSSHSettingReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayAuditSSHSettingReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [GatewayAuditSSHSettingReplaceResponseEnvelopeErrors]
type gatewayAuditSSHSettingReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAuditSSHSettingReplaceResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    gatewayAuditSSHSettingReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayAuditSSHSettingReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [GatewayAuditSSHSettingReplaceResponseEnvelopeMessages]
type gatewayAuditSSHSettingReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayAuditSSHSettingReplaceResponseEnvelopeSuccess bool

const (
	GatewayAuditSSHSettingReplaceResponseEnvelopeSuccessTrue GatewayAuditSSHSettingReplaceResponseEnvelopeSuccess = true
)
