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

// AccountGatewayService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountGatewayService] method
// instead.
type AccountGatewayService struct {
	Options          []option.RequestOption
	AuditSSHSettings *AccountGatewayAuditSSHSettingService
	Categories       *AccountGatewayCategoryService
	AppTypes         *AccountGatewayAppTypeService
	Configurations   *AccountGatewayConfigurationService
	Lists            *AccountGatewayListService
	Locations        *AccountGatewayLocationService
	Loggings         *AccountGatewayLoggingService
	ProxyEndpoints   *AccountGatewayProxyEndpointService
	Rules            *AccountGatewayRuleService
}

// NewAccountGatewayService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountGatewayService(opts ...option.RequestOption) (r *AccountGatewayService) {
	r = &AccountGatewayService{}
	r.Options = opts
	r.AuditSSHSettings = NewAccountGatewayAuditSSHSettingService(opts...)
	r.Categories = NewAccountGatewayCategoryService(opts...)
	r.AppTypes = NewAccountGatewayAppTypeService(opts...)
	r.Configurations = NewAccountGatewayConfigurationService(opts...)
	r.Lists = NewAccountGatewayListService(opts...)
	r.Locations = NewAccountGatewayLocationService(opts...)
	r.Loggings = NewAccountGatewayLoggingService(opts...)
	r.ProxyEndpoints = NewAccountGatewayProxyEndpointService(opts...)
	r.Rules = NewAccountGatewayRuleService(opts...)
	return
}

// Creates a Zero Trust account with an existing Cloudflare account.
func (r *AccountGatewayService) ZeroTrustAccountsNewZeroTrustAccount(ctx context.Context, identifier interface{}, body AccountGatewayZeroTrustAccountsNewZeroTrustAccountParams, opts ...option.RequestOption) (res *AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets information about the current Zero Trust account.
func (r *AccountGatewayService) ZeroTrustAccountsGetZeroTrustAccountInformation(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponse struct {
	Errors   []AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseError   `json:"errors"`
	Messages []AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseMessage `json:"messages"`
	Result   AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseSuccess `json:"success"`
	JSON    accountGatewayZeroTrustAccountsNewZeroTrustAccountResponseJSON    `json:"-"`
}

// accountGatewayZeroTrustAccountsNewZeroTrustAccountResponseJSON contains the JSON
// metadata for the struct
// [AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponse]
type accountGatewayZeroTrustAccountsNewZeroTrustAccountResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accountGatewayZeroTrustAccountsNewZeroTrustAccountResponseErrorJSON `json:"-"`
}

// accountGatewayZeroTrustAccountsNewZeroTrustAccountResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseError]
type accountGatewayZeroTrustAccountsNewZeroTrustAccountResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountGatewayZeroTrustAccountsNewZeroTrustAccountResponseMessageJSON `json:"-"`
}

// accountGatewayZeroTrustAccountsNewZeroTrustAccountResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseMessage]
type accountGatewayZeroTrustAccountsNewZeroTrustAccountResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseResult struct {
	// Cloudflare account ID.
	ID string `json:"id"`
	// Gateway internal ID.
	GatewayTag string `json:"gateway_tag"`
	// The name of the provider. Usually Cloudflare.
	ProviderName string                                                               `json:"provider_name"`
	JSON         accountGatewayZeroTrustAccountsNewZeroTrustAccountResponseResultJSON `json:"-"`
}

// accountGatewayZeroTrustAccountsNewZeroTrustAccountResponseResultJSON contains
// the JSON metadata for the struct
// [AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseResult]
type accountGatewayZeroTrustAccountsNewZeroTrustAccountResponseResultJSON struct {
	ID           apijson.Field
	GatewayTag   apijson.Field
	ProviderName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseSuccess bool

const (
	AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseSuccessTrue AccountGatewayZeroTrustAccountsNewZeroTrustAccountResponseSuccess = true
)

type AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponse struct {
	Errors   []AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseError   `json:"errors"`
	Messages []AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseMessage `json:"messages"`
	Result   AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseSuccess `json:"success"`
	JSON    accountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseJSON    `json:"-"`
}

// accountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponse]
type accountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseError struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    accountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseErrorJSON `json:"-"`
}

// accountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseError]
type accountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseMessage struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    accountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseMessageJSON `json:"-"`
}

// accountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseMessage]
type accountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseResult struct {
	// Cloudflare account ID.
	ID string `json:"id"`
	// Gateway internal ID.
	GatewayTag string `json:"gateway_tag"`
	// The name of the provider. Usually Cloudflare.
	ProviderName string                                                                          `json:"provider_name"`
	JSON         accountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseResultJSON `json:"-"`
}

// accountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseResult]
type accountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseResultJSON struct {
	ID           apijson.Field
	GatewayTag   apijson.Field
	ProviderName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseSuccess bool

const (
	AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseSuccessTrue AccountGatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseSuccess = true
)

type AccountGatewayZeroTrustAccountsNewZeroTrustAccountParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `json:"account_id,required"`
}

func (r AccountGatewayZeroTrustAccountsNewZeroTrustAccountParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
