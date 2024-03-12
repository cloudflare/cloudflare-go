// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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

// Updates the current device settings for a Zero Trust account.
func (r *DeviceSettingService) Update(ctx context.Context, params DeviceSettingUpdateParams, opts ...option.RequestOption) (res *TeamsDevicesZeroTrustAccountDeviceSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceSettingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Describes the current device settings for a Zero Trust account.
func (r *DeviceSettingService) List(ctx context.Context, query DeviceSettingListParams, opts ...option.RequestOption) (res *TeamsDevicesZeroTrustAccountDeviceSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceSettingListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/settings", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TeamsDevicesZeroTrustAccountDeviceSettings struct {
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled bool `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled bool `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled bool `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP bool                                           `json:"use_zt_virtual_ip"`
	JSON           teamsDevicesZeroTrustAccountDeviceSettingsJSON `json:"-"`
}

// teamsDevicesZeroTrustAccountDeviceSettingsJSON contains the JSON metadata for
// the struct [TeamsDevicesZeroTrustAccountDeviceSettings]
type teamsDevicesZeroTrustAccountDeviceSettingsJSON struct {
	GatewayProxyEnabled                apijson.Field
	GatewayUdpProxyEnabled             apijson.Field
	RootCertificateInstallationEnabled apijson.Field
	UseZtVirtualIP                     apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *TeamsDevicesZeroTrustAccountDeviceSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesZeroTrustAccountDeviceSettingsJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled param.Field[bool] `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled param.Field[bool] `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled param.Field[bool] `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP param.Field[bool] `json:"use_zt_virtual_ip"`
}

func (r DeviceSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceSettingUpdateResponseEnvelope struct {
	Errors   []DeviceSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamsDevicesZeroTrustAccountDeviceSettings    `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceSettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// deviceSettingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceSettingUpdateResponseEnvelope]
type deviceSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    deviceSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceSettingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceSettingUpdateResponseEnvelopeErrors]
type deviceSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    deviceSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceSettingUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceSettingUpdateResponseEnvelopeMessages]
type deviceSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceSettingUpdateResponseEnvelopeSuccess bool

const (
	DeviceSettingUpdateResponseEnvelopeSuccessTrue DeviceSettingUpdateResponseEnvelopeSuccess = true
)

type DeviceSettingListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type DeviceSettingListResponseEnvelope struct {
	Errors   []DeviceSettingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceSettingListResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamsDevicesZeroTrustAccountDeviceSettings  `json:"result,required,nullable"`
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

func (r deviceSettingListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r deviceSettingListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r deviceSettingListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceSettingListResponseEnvelopeSuccess bool

const (
	DeviceSettingListResponseEnvelopeSuccessTrue DeviceSettingListResponseEnvelopeSuccess = true
)
