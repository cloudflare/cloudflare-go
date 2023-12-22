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
	Options        []option.RequestOption
	Categories     *AccountGatewayCategoryService
	AppTypes       *AccountGatewayAppTypeService
	Configurations *AccountGatewayConfigurationService
	Lists          *AccountGatewayListService
	Locations      *AccountGatewayLocationService
	Loggings       *AccountGatewayLoggingService
	ProxyEndpoints *AccountGatewayProxyEndpointService
	Rules          *AccountGatewayRuleService
}

// NewAccountGatewayService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountGatewayService(opts ...option.RequestOption) (r *AccountGatewayService) {
	r = &AccountGatewayService{}
	r.Options = opts
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

// Create Zero Trust account with existing cloudflare account.
func (r *AccountGatewayService) ZeroTrustAccountsNewZeroTrustAccount(ctx context.Context, identifier interface{}, body AccountGatewayZeroTrustAccountsNewZeroTrustAccountParams, opts ...option.RequestOption) (res *GatewayAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Zero Trust account information.
func (r *AccountGatewayService) ZeroTrustAccountsGetZeroTrustAccountInformation(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *GatewayAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GatewayAccount struct {
	Errors   []GatewayAccountError   `json:"errors"`
	Messages []GatewayAccountMessage `json:"messages"`
	Result   GatewayAccountResult    `json:"result"`
	// Whether the API call was successful
	Success GatewayAccountSuccess `json:"success"`
	JSON    gatewayAccountJSON    `json:"-"`
}

// gatewayAccountJSON contains the JSON metadata for the struct [GatewayAccount]
type gatewayAccountJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAccountError struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    gatewayAccountErrorJSON `json:"-"`
}

// gatewayAccountErrorJSON contains the JSON metadata for the struct
// [GatewayAccountError]
type gatewayAccountErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAccountMessage struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    gatewayAccountMessageJSON `json:"-"`
}

// gatewayAccountMessageJSON contains the JSON metadata for the struct
// [GatewayAccountMessage]
type gatewayAccountMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAccountResult struct {
	// Cloudflare account ID
	ID string `json:"id"`
	// Gateway internal id.
	GatewayTag string `json:"gateway_tag"`
	// The name of provider. Usually cloudflare.
	ProviderName string                   `json:"provider_name"`
	JSON         gatewayAccountResultJSON `json:"-"`
}

// gatewayAccountResultJSON contains the JSON metadata for the struct
// [GatewayAccountResult]
type gatewayAccountResultJSON struct {
	ID           apijson.Field
	GatewayTag   apijson.Field
	ProviderName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayAccountResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayAccountSuccess bool

const (
	GatewayAccountSuccessTrue GatewayAccountSuccess = true
)

type AccountGatewayZeroTrustAccountsNewZeroTrustAccountParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `json:"account_id,required"`
}

func (r AccountGatewayZeroTrustAccountsNewZeroTrustAccountParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
