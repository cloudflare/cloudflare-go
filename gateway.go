// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// GatewayService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewGatewayService] method instead.
type GatewayService struct {
	Options        []option.RequestOption
	Categories     *GatewayCategoryService
	AppTypes       *GatewayAppTypeService
	Configurations *GatewayConfigurationService
	Lists          *GatewayListService
	Locations      *GatewayLocationService
	Loggings       *GatewayLoggingService
	ProxyEndpoints *GatewayProxyEndpointService
	Rules          *GatewayRuleService
}

// NewGatewayService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGatewayService(opts ...option.RequestOption) (r *GatewayService) {
	r = &GatewayService{}
	r.Options = opts
	r.Categories = NewGatewayCategoryService(opts...)
	r.AppTypes = NewGatewayAppTypeService(opts...)
	r.Configurations = NewGatewayConfigurationService(opts...)
	r.Lists = NewGatewayListService(opts...)
	r.Locations = NewGatewayLocationService(opts...)
	r.Loggings = NewGatewayLoggingService(opts...)
	r.ProxyEndpoints = NewGatewayProxyEndpointService(opts...)
	r.Rules = NewGatewayRuleService(opts...)
	return
}

// Creates a Zero Trust account with an existing Cloudflare account.
func (r *GatewayService) ZeroTrustAccountsNewZeroTrustAccount(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *GatewayZeroTrustAccountsNewZeroTrustAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets information about the current Zero Trust account.
func (r *GatewayService) ZeroTrustAccountsGetZeroTrustAccountInformation(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayZeroTrustAccountsNewZeroTrustAccountResponse struct {
	// Cloudflare account ID.
	ID string `json:"id"`
	// Gateway internal ID.
	GatewayTag string `json:"gateway_tag"`
	// The name of the provider. Usually Cloudflare.
	ProviderName string                                                  `json:"provider_name"`
	JSON         gatewayZeroTrustAccountsNewZeroTrustAccountResponseJSON `json:"-"`
}

// gatewayZeroTrustAccountsNewZeroTrustAccountResponseJSON contains the JSON
// metadata for the struct [GatewayZeroTrustAccountsNewZeroTrustAccountResponse]
type gatewayZeroTrustAccountsNewZeroTrustAccountResponseJSON struct {
	ID           apijson.Field
	GatewayTag   apijson.Field
	ProviderName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayZeroTrustAccountsNewZeroTrustAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponse struct {
	// Cloudflare account ID.
	ID string `json:"id"`
	// Gateway internal ID.
	GatewayTag string `json:"gateway_tag"`
	// The name of the provider. Usually Cloudflare.
	ProviderName string                                                             `json:"provider_name"`
	JSON         gatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseJSON `json:"-"`
}

// gatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseJSON contains the
// JSON metadata for the struct
// [GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponse]
type gatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseJSON struct {
	ID           apijson.Field
	GatewayTag   apijson.Field
	ProviderName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelope struct {
	Errors   []GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayZeroTrustAccountsNewZeroTrustAccountResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeJSON    `json:"-"`
}

// gatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelope]
type gatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeErrors struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    gatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeErrors]
type gatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeMessages struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    gatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeMessages]
type gatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeSuccess bool

const (
	GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeSuccessTrue GatewayZeroTrustAccountsNewZeroTrustAccountResponseEnvelopeSuccess = true
)

type GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelope struct {
	Errors   []GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeJSON    `json:"-"`
}

// gatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelope]
type gatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeErrors struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    gatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeErrors]
type gatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeMessages struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    gatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeMessages]
type gatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeSuccess bool

const (
	GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeSuccessTrue GatewayZeroTrustAccountsGetZeroTrustAccountInformationResponseEnvelopeSuccess = true
)
