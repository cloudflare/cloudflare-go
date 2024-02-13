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

// DeviceSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDeviceSettingService] method
// instead.
type DeviceSettingService struct {
	Options []option.RequestOption
}

// NewDeviceSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceSettingService(opts ...option.RequestOption) (r *DeviceSettingService) {
	r = &DeviceSettingService{}
	r.Options = opts
	return
}

// Describes the current device settings for a Zero Trust account.
func (r *DeviceSettingService) ZeroTrustAccountsGetDeviceSettingsForZeroTrustAccount(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the current device settings for a Zero Trust account.
func (r *DeviceSettingService) ZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccount(ctx context.Context, identifier interface{}, body DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountParams, opts ...option.RequestOption) (res *DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponse struct {
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled bool `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled bool `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled bool `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP bool                                                                           `json:"use_zt_virtual_ip"`
	JSON           deviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseJSON `json:"-"`
}

// deviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseJSON
// contains the JSON metadata for the struct
// [DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponse]
type deviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseJSON struct {
	GatewayProxyEnabled                apijson.Field
	GatewayUdpProxyEnabled             apijson.Field
	RootCertificateInstallationEnabled apijson.Field
	UseZtVirtualIP                     apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponse struct {
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled bool `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled bool `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled bool `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP bool                                                                                 `json:"use_zt_virtual_ip"`
	JSON           deviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseJSON `json:"-"`
}

// deviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseJSON
// contains the JSON metadata for the struct
// [DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponse]
type deviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseJSON struct {
	GatewayProxyEnabled                apijson.Field
	GatewayUdpProxyEnabled             apijson.Field
	RootCertificateInstallationEnabled apijson.Field
	UseZtVirtualIP                     apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelope struct {
	Errors   []DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeJSON    `json:"-"`
}

// deviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelope]
type deviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeErrors struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    deviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeErrors]
type deviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeMessages struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    deviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeMessages]
type deviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeSuccess bool

const (
	DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeSuccessTrue DeviceSettingZeroTrustAccountsGetDeviceSettingsForZeroTrustAccountResponseEnvelopeSuccess = true
)

type DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountParams struct {
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled param.Field[bool] `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled param.Field[bool] `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled param.Field[bool] `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP param.Field[bool] `json:"use_zt_virtual_ip"`
}

func (r DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelope struct {
	Errors   []DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeJSON    `json:"-"`
}

// deviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelope]
type deviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeErrors struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    deviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeErrors]
type deviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeMessages struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    deviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeMessages]
type deviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeSuccess bool

const (
	DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeSuccessTrue DeviceSettingZeroTrustAccountsUpdateDeviceSettingsForTheZeroTrustAccountResponseEnvelopeSuccess = true
)
