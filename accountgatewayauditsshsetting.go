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

// AccountGatewayAuditSSHSettingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountGatewayAuditSSHSettingService] method instead.
type AccountGatewayAuditSSHSettingService struct {
	Options []option.RequestOption
}

// NewAccountGatewayAuditSSHSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountGatewayAuditSSHSettingService(opts ...option.RequestOption) (r *AccountGatewayAuditSSHSettingService) {
	r = &AccountGatewayAuditSSHSettingService{}
	r.Options = opts
	return
}

// Get all Zero Trust Audit SSH settings for an account.
func (r *AccountGatewayAuditSSHSettingService) Get(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountGatewayAuditSSHSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/audit_ssh_settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates Zero Trust Audit SSH settings.
func (r *AccountGatewayAuditSSHSettingService) Update(ctx context.Context, identifier interface{}, body AccountGatewayAuditSSHSettingUpdateParams, opts ...option.RequestOption) (res *AccountGatewayAuditSSHSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/audit_ssh_settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountGatewayAuditSSHSettingGetResponse struct {
	Errors   []AccountGatewayAuditSSHSettingGetResponseError   `json:"errors"`
	Messages []AccountGatewayAuditSSHSettingGetResponseMessage `json:"messages"`
	Result   AccountGatewayAuditSSHSettingGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayAuditSSHSettingGetResponseSuccess `json:"success"`
	JSON    accountGatewayAuditSSHSettingGetResponseJSON    `json:"-"`
}

// accountGatewayAuditSSHSettingGetResponseJSON contains the JSON metadata for the
// struct [AccountGatewayAuditSSHSettingGetResponse]
type accountGatewayAuditSSHSettingGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAuditSSHSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayAuditSSHSettingGetResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountGatewayAuditSSHSettingGetResponseErrorJSON `json:"-"`
}

// accountGatewayAuditSSHSettingGetResponseErrorJSON contains the JSON metadata for
// the struct [AccountGatewayAuditSSHSettingGetResponseError]
type accountGatewayAuditSSHSettingGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAuditSSHSettingGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayAuditSSHSettingGetResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountGatewayAuditSSHSettingGetResponseMessageJSON `json:"-"`
}

// accountGatewayAuditSSHSettingGetResponseMessageJSON contains the JSON metadata
// for the struct [AccountGatewayAuditSSHSettingGetResponseMessage]
type accountGatewayAuditSSHSettingGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAuditSSHSettingGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayAuditSSHSettingGetResponseResult struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// SSH encryption public key
	PublicKey string `json:"public_key"`
	// Seed ID
	SeedID    string                                             `json:"seed_id"`
	UpdatedAt time.Time                                          `json:"updated_at" format:"date-time"`
	JSON      accountGatewayAuditSSHSettingGetResponseResultJSON `json:"-"`
}

// accountGatewayAuditSSHSettingGetResponseResultJSON contains the JSON metadata
// for the struct [AccountGatewayAuditSSHSettingGetResponseResult]
type accountGatewayAuditSSHSettingGetResponseResultJSON struct {
	CreatedAt   apijson.Field
	PublicKey   apijson.Field
	SeedID      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAuditSSHSettingGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayAuditSSHSettingGetResponseSuccess bool

const (
	AccountGatewayAuditSSHSettingGetResponseSuccessTrue AccountGatewayAuditSSHSettingGetResponseSuccess = true
)

type AccountGatewayAuditSSHSettingUpdateResponse struct {
	Errors   []AccountGatewayAuditSSHSettingUpdateResponseError   `json:"errors"`
	Messages []AccountGatewayAuditSSHSettingUpdateResponseMessage `json:"messages"`
	Result   AccountGatewayAuditSSHSettingUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayAuditSSHSettingUpdateResponseSuccess `json:"success"`
	JSON    accountGatewayAuditSSHSettingUpdateResponseJSON    `json:"-"`
}

// accountGatewayAuditSSHSettingUpdateResponseJSON contains the JSON metadata for
// the struct [AccountGatewayAuditSSHSettingUpdateResponse]
type accountGatewayAuditSSHSettingUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAuditSSHSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayAuditSSHSettingUpdateResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountGatewayAuditSSHSettingUpdateResponseErrorJSON `json:"-"`
}

// accountGatewayAuditSSHSettingUpdateResponseErrorJSON contains the JSON metadata
// for the struct [AccountGatewayAuditSSHSettingUpdateResponseError]
type accountGatewayAuditSSHSettingUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAuditSSHSettingUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayAuditSSHSettingUpdateResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountGatewayAuditSSHSettingUpdateResponseMessageJSON `json:"-"`
}

// accountGatewayAuditSSHSettingUpdateResponseMessageJSON contains the JSON
// metadata for the struct [AccountGatewayAuditSSHSettingUpdateResponseMessage]
type accountGatewayAuditSSHSettingUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAuditSSHSettingUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayAuditSSHSettingUpdateResponseResult struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// SSH encryption public key
	PublicKey string `json:"public_key"`
	// Seed ID
	SeedID    string                                                `json:"seed_id"`
	UpdatedAt time.Time                                             `json:"updated_at" format:"date-time"`
	JSON      accountGatewayAuditSSHSettingUpdateResponseResultJSON `json:"-"`
}

// accountGatewayAuditSSHSettingUpdateResponseResultJSON contains the JSON metadata
// for the struct [AccountGatewayAuditSSHSettingUpdateResponseResult]
type accountGatewayAuditSSHSettingUpdateResponseResultJSON struct {
	CreatedAt   apijson.Field
	PublicKey   apijson.Field
	SeedID      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayAuditSSHSettingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayAuditSSHSettingUpdateResponseSuccess bool

const (
	AccountGatewayAuditSSHSettingUpdateResponseSuccessTrue AccountGatewayAuditSSHSettingUpdateResponseSuccess = true
)

type AccountGatewayAuditSSHSettingUpdateParams struct {
	// SSH encryption public key
	PublicKey param.Field[string] `json:"public_key,required"`
	// Seed ID
	SeedID param.Field[string] `json:"seed_id"`
}

func (r AccountGatewayAuditSSHSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
