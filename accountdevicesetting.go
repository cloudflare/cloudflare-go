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

// Describes the current device settings for a Zero Trust account.
func (r *AccountDeviceSettingService) ZeroTrustAccountsGetDeviceSettingsForZeroTrustAccount(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the current device settings for a Zero Trust account.
func (r *AccountDeviceSettingService) ZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccount(ctx context.Context, identifier interface{}, body AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountParams, opts ...option.RequestOption) (res *AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponse struct {
	Errors   []AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseError   `json:"errors"`
	Messages []AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseMessage `json:"messages"`
	Result   AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseSuccess `json:"success"`
	JSON    accountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseJSON    `json:"-"`
}

// accountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseJSON
// contains the JSON metadata for the struct
// [AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponse]
type accountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseError struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    accountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseErrorJSON `json:"-"`
}

// accountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseError]
type accountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseMessage struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    accountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseMessageJSON `json:"-"`
}

// accountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseMessage]
type accountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseResult struct {
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled bool `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled bool `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled bool `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP bool                                                                                        `json:"use_zt_virtual_ip"`
	JSON           accountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseResultJSON `json:"-"`
}

// accountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseResult]
type accountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseResultJSON struct {
	GatewayProxyEnabled                apijson.Field
	GatewayUdpProxyEnabled             apijson.Field
	RootCertificateInstallationEnabled apijson.Field
	UseZtVirtualIP                     apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseSuccess bool

const (
	AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseSuccessTrue AccountDeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseSuccess = true
)

type AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponse struct {
	Errors   []AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseError   `json:"errors"`
	Messages []AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseMessage `json:"messages"`
	Result   AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseSuccess `json:"success"`
	JSON    accountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseJSON    `json:"-"`
}

// accountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseJSON
// contains the JSON metadata for the struct
// [AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponse]
type accountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseError struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    accountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseErrorJSON `json:"-"`
}

// accountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseError]
type accountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseMessage struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    accountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseMessageJSON `json:"-"`
}

// accountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseMessage]
type accountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseResult struct {
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled bool `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled bool `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled bool `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP bool                                                                                              `json:"use_zt_virtual_ip"`
	JSON           accountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseResultJSON `json:"-"`
}

// accountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseResult]
type accountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseResultJSON struct {
	GatewayProxyEnabled                apijson.Field
	GatewayUdpProxyEnabled             apijson.Field
	RootCertificateInstallationEnabled apijson.Field
	UseZtVirtualIP                     apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseSuccess bool

const (
	AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseSuccessTrue AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseSuccess = true
)

type AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountParams struct {
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled param.Field[bool] `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled param.Field[bool] `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled param.Field[bool] `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP param.Field[bool] `json:"use_zt_virtual_ip"`
}

func (r AccountDeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
