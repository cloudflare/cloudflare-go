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
	path := fmt.Sprintf("accounts/%v/gateway/audit_ssh_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates Zero Trust Audit SSH settings.
func (r *GatewayAuditSSHSettingService) Update(ctx context.Context, accountID interface{}, body GatewayAuditSSHSettingUpdateParams, opts ...option.RequestOption) (res *GatewayAuditSSHSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/audit_ssh_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type GatewayAuditSSHSettingGetResponse struct {
	Errors   []GatewayAuditSSHSettingGetResponseError   `json:"errors"`
	Messages []GatewayAuditSSHSettingGetResponseMessage `json:"messages"`
	Result   GatewayAuditSSHSettingGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success GatewayAuditSSHSettingGetResponseSuccess `json:"success"`
	JSON    gatewayAuditSSHSettingGetResponseJSON    `json:"-"`
}

// gatewayAuditSSHSettingGetResponseJSON contains the JSON metadata for the struct
// [GatewayAuditSSHSettingGetResponse]
type gatewayAuditSSHSettingGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAuditSSHSettingGetResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    gatewayAuditSSHSettingGetResponseErrorJSON `json:"-"`
}

// gatewayAuditSSHSettingGetResponseErrorJSON contains the JSON metadata for the
// struct [GatewayAuditSSHSettingGetResponseError]
type gatewayAuditSSHSettingGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAuditSSHSettingGetResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    gatewayAuditSSHSettingGetResponseMessageJSON `json:"-"`
}

// gatewayAuditSSHSettingGetResponseMessageJSON contains the JSON metadata for the
// struct [GatewayAuditSSHSettingGetResponseMessage]
type gatewayAuditSSHSettingGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAuditSSHSettingGetResponseResult struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// SSH encryption public key
	PublicKey string `json:"public_key"`
	// Seed ID
	SeedID    string                                      `json:"seed_id"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      gatewayAuditSSHSettingGetResponseResultJSON `json:"-"`
}

// gatewayAuditSSHSettingGetResponseResultJSON contains the JSON metadata for the
// struct [GatewayAuditSSHSettingGetResponseResult]
type gatewayAuditSSHSettingGetResponseResultJSON struct {
	CreatedAt   apijson.Field
	PublicKey   apijson.Field
	SeedID      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayAuditSSHSettingGetResponseSuccess bool

const (
	GatewayAuditSSHSettingGetResponseSuccessTrue GatewayAuditSSHSettingGetResponseSuccess = true
)

type GatewayAuditSSHSettingUpdateResponse struct {
	Errors   []GatewayAuditSSHSettingUpdateResponseError   `json:"errors"`
	Messages []GatewayAuditSSHSettingUpdateResponseMessage `json:"messages"`
	Result   GatewayAuditSSHSettingUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success GatewayAuditSSHSettingUpdateResponseSuccess `json:"success"`
	JSON    gatewayAuditSSHSettingUpdateResponseJSON    `json:"-"`
}

// gatewayAuditSSHSettingUpdateResponseJSON contains the JSON metadata for the
// struct [GatewayAuditSSHSettingUpdateResponse]
type gatewayAuditSSHSettingUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAuditSSHSettingUpdateResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    gatewayAuditSSHSettingUpdateResponseErrorJSON `json:"-"`
}

// gatewayAuditSSHSettingUpdateResponseErrorJSON contains the JSON metadata for the
// struct [GatewayAuditSSHSettingUpdateResponseError]
type gatewayAuditSSHSettingUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAuditSSHSettingUpdateResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    gatewayAuditSSHSettingUpdateResponseMessageJSON `json:"-"`
}

// gatewayAuditSSHSettingUpdateResponseMessageJSON contains the JSON metadata for
// the struct [GatewayAuditSSHSettingUpdateResponseMessage]
type gatewayAuditSSHSettingUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAuditSSHSettingUpdateResponseResult struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// SSH encryption public key
	PublicKey string `json:"public_key"`
	// Seed ID
	SeedID    string                                         `json:"seed_id"`
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      gatewayAuditSSHSettingUpdateResponseResultJSON `json:"-"`
}

// gatewayAuditSSHSettingUpdateResponseResultJSON contains the JSON metadata for
// the struct [GatewayAuditSSHSettingUpdateResponseResult]
type gatewayAuditSSHSettingUpdateResponseResultJSON struct {
	CreatedAt   apijson.Field
	PublicKey   apijson.Field
	SeedID      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAuditSSHSettingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayAuditSSHSettingUpdateResponseSuccess bool

const (
	GatewayAuditSSHSettingUpdateResponseSuccessTrue GatewayAuditSSHSettingUpdateResponseSuccess = true
)

type GatewayAuditSSHSettingUpdateParams struct {
	// SSH encryption public key
	PublicKey param.Field[string] `json:"public_key,required"`
	// Seed ID
	SeedID param.Field[string] `json:"seed_id"`
}

func (r GatewayAuditSSHSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
