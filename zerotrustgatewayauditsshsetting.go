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

// ZeroTrustGatewayAuditSSHSettingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustGatewayAuditSSHSettingService] method instead.
type ZeroTrustGatewayAuditSSHSettingService struct {
	Options []option.RequestOption
}

// NewZeroTrustGatewayAuditSSHSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustGatewayAuditSSHSettingService(opts ...option.RequestOption) (r *ZeroTrustGatewayAuditSSHSettingService) {
	r = &ZeroTrustGatewayAuditSSHSettingService{}
	r.Options = opts
	return
}

// Updates Zero Trust Audit SSH settings.
func (r *ZeroTrustGatewayAuditSSHSettingService) Update(ctx context.Context, params ZeroTrustGatewayAuditSSHSettingUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGatewayAuditSSHSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/audit_ssh_settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get all Zero Trust Audit SSH settings for an account.
func (r *ZeroTrustGatewayAuditSSHSettingService) Get(ctx context.Context, query ZeroTrustGatewayAuditSSHSettingGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewayAuditSSHSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayAuditSSHSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/audit_ssh_settings", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGatewayAuditSSHSettingUpdateResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// SSH encryption public key
	PublicKey string `json:"public_key"`
	// Seed ID
	SeedID    string                                            `json:"seed_id"`
	UpdatedAt time.Time                                         `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayAuditSSHSettingUpdateResponseJSON `json:"-"`
}

// zeroTrustGatewayAuditSSHSettingUpdateResponseJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayAuditSSHSettingUpdateResponse]
type zeroTrustGatewayAuditSSHSettingUpdateResponseJSON struct {
	CreatedAt   apijson.Field
	PublicKey   apijson.Field
	SeedID      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAuditSSHSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayAuditSSHSettingGetResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// SSH encryption public key
	PublicKey string `json:"public_key"`
	// Seed ID
	SeedID    string                                         `json:"seed_id"`
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayAuditSSHSettingGetResponseJSON `json:"-"`
}

// zeroTrustGatewayAuditSSHSettingGetResponseJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayAuditSSHSettingGetResponse]
type zeroTrustGatewayAuditSSHSettingGetResponseJSON struct {
	CreatedAt   apijson.Field
	PublicKey   apijson.Field
	SeedID      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAuditSSHSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayAuditSSHSettingUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// SSH encryption public key
	PublicKey param.Field[string] `json:"public_key,required"`
	// Seed ID
	SeedID param.Field[string] `json:"seed_id"`
}

func (r ZeroTrustGatewayAuditSSHSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelope struct {
	Errors   []ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayAuditSSHSettingUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelope]
type zeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeErrors]
type zeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeMessages]
type zeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeSuccessTrue ZeroTrustGatewayAuditSSHSettingUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayAuditSSHSettingGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayAuditSSHSettingGetResponseEnvelope struct {
	Errors   []ZeroTrustGatewayAuditSSHSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayAuditSSHSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayAuditSSHSettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayAuditSSHSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayAuditSSHSettingGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayAuditSSHSettingGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayAuditSSHSettingGetResponseEnvelope]
type zeroTrustGatewayAuditSSHSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAuditSSHSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayAuditSSHSettingGetResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustGatewayAuditSSHSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayAuditSSHSettingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayAuditSSHSettingGetResponseEnvelopeErrors]
type zeroTrustGatewayAuditSSHSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAuditSSHSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayAuditSSHSettingGetResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustGatewayAuditSSHSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayAuditSSHSettingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayAuditSSHSettingGetResponseEnvelopeMessages]
type zeroTrustGatewayAuditSSHSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAuditSSHSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayAuditSSHSettingGetResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayAuditSSHSettingGetResponseEnvelopeSuccessTrue ZeroTrustGatewayAuditSSHSettingGetResponseEnvelopeSuccess = true
)
