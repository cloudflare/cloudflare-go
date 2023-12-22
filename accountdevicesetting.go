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

// AccountDeviceSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDeviceSettingService]
// method instead.
type AccountDeviceSettingService struct {
	Options []option.RequestOption
}

// NewAccountDeviceSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDeviceSettingService(opts ...option.RequestOption) (r *AccountDeviceSettingService) {
	r = &AccountDeviceSettingService{}
	r.Options = opts
	return
}

// Describes the current device settings Zero Trust account.
func (r *AccountDeviceSettingService) ZeroTrustAccountsGetDeviceSettingsForZeroTrustAccount(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *GatewayAccountDeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the current device settings for Zero Trust account.
func (r *AccountDeviceSettingService) ZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccount(ctx context.Context, identifier interface{}, body AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountParams, opts ...option.RequestOption) (res *GatewayAccountDeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type GatewayAccountDeviceSettingsResponse struct {
	Errors   []GatewayAccountDeviceSettingsResponseError   `json:"errors"`
	Messages []GatewayAccountDeviceSettingsResponseMessage `json:"messages"`
	Result   GatewayAccountDeviceSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success GatewayAccountDeviceSettingsResponseSuccess `json:"success"`
	JSON    gatewayAccountDeviceSettingsResponseJSON    `json:"-"`
}

// gatewayAccountDeviceSettingsResponseJSON contains the JSON metadata for the
// struct [GatewayAccountDeviceSettingsResponse]
type gatewayAccountDeviceSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountDeviceSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAccountDeviceSettingsResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    gatewayAccountDeviceSettingsResponseErrorJSON `json:"-"`
}

// gatewayAccountDeviceSettingsResponseErrorJSON contains the JSON metadata for the
// struct [GatewayAccountDeviceSettingsResponseError]
type gatewayAccountDeviceSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountDeviceSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAccountDeviceSettingsResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    gatewayAccountDeviceSettingsResponseMessageJSON `json:"-"`
}

// gatewayAccountDeviceSettingsResponseMessageJSON contains the JSON metadata for
// the struct [GatewayAccountDeviceSettingsResponseMessage]
type gatewayAccountDeviceSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountDeviceSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAccountDeviceSettingsResponseResult struct {
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled bool `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled bool                                           `json:"gateway_udp_proxy_enabled"`
	JSON                   gatewayAccountDeviceSettingsResponseResultJSON `json:"-"`
}

// gatewayAccountDeviceSettingsResponseResultJSON contains the JSON metadata for
// the struct [GatewayAccountDeviceSettingsResponseResult]
type gatewayAccountDeviceSettingsResponseResultJSON struct {
	GatewayProxyEnabled    apijson.Field
	GatewayUdpProxyEnabled apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *GatewayAccountDeviceSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayAccountDeviceSettingsResponseSuccess bool

const (
	GatewayAccountDeviceSettingsResponseSuccessTrue GatewayAccountDeviceSettingsResponseSuccess = true
)

type AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountParams struct {
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled param.Field[bool] `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled param.Field[bool] `json:"gateway_udp_proxy_enabled"`
}

func (r AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
