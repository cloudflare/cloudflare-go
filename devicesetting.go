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
func (r *DeviceSettingService) List(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *DeviceSettingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceSettingListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the current device settings for a Zero Trust account.
func (r *DeviceSettingService) Replace(ctx context.Context, identifier interface{}, body DeviceSettingReplaceParams, opts ...option.RequestOption) (res *DeviceSettingReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceSettingReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceSettingListResponse struct {
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled bool `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled bool `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled bool `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP bool                          `json:"use_zt_virtual_ip"`
	JSON           deviceSettingListResponseJSON `json:"-"`
}

// deviceSettingListResponseJSON contains the JSON metadata for the struct
// [DeviceSettingListResponse]
type deviceSettingListResponseJSON struct {
	GatewayProxyEnabled                apijson.Field
	GatewayUdpProxyEnabled             apijson.Field
	RootCertificateInstallationEnabled apijson.Field
	UseZtVirtualIP                     apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *DeviceSettingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingReplaceResponse struct {
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled bool `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled bool `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled bool `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP bool                             `json:"use_zt_virtual_ip"`
	JSON           deviceSettingReplaceResponseJSON `json:"-"`
}

// deviceSettingReplaceResponseJSON contains the JSON metadata for the struct
// [DeviceSettingReplaceResponse]
type deviceSettingReplaceResponseJSON struct {
	GatewayProxyEnabled                apijson.Field
	GatewayUdpProxyEnabled             apijson.Field
	RootCertificateInstallationEnabled apijson.Field
	UseZtVirtualIP                     apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *DeviceSettingReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingListResponseEnvelope struct {
	Errors   []DeviceSettingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceSettingListResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceSettingListResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceSettingListResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceSettingListResponseEnvelopeJSON    `json:"-"`
}

// deviceSettingListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceSettingListResponseEnvelope]
type deviceSettingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingListResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    deviceSettingListResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceSettingListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceSettingListResponseEnvelopeErrors]
type deviceSettingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingListResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    deviceSettingListResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceSettingListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceSettingListResponseEnvelopeMessages]
type deviceSettingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceSettingListResponseEnvelopeSuccess bool

const (
	DeviceSettingListResponseEnvelopeSuccessTrue DeviceSettingListResponseEnvelopeSuccess = true
)

type DeviceSettingReplaceParams struct {
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled param.Field[bool] `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled param.Field[bool] `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled param.Field[bool] `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP param.Field[bool] `json:"use_zt_virtual_ip"`
}

func (r DeviceSettingReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceSettingReplaceResponseEnvelope struct {
	Errors   []DeviceSettingReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceSettingReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceSettingReplaceResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceSettingReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceSettingReplaceResponseEnvelopeJSON    `json:"-"`
}

// deviceSettingReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceSettingReplaceResponseEnvelope]
type deviceSettingReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingReplaceResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    deviceSettingReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceSettingReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DeviceSettingReplaceResponseEnvelopeErrors]
type deviceSettingReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingReplaceResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    deviceSettingReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceSettingReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceSettingReplaceResponseEnvelopeMessages]
type deviceSettingReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceSettingReplaceResponseEnvelopeSuccess bool

const (
	DeviceSettingReplaceResponseEnvelopeSuccessTrue DeviceSettingReplaceResponseEnvelopeSuccess = true
)
