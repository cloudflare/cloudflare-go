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

// ZeroTrustDeviceSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDeviceSettingService]
// method instead.
type ZeroTrustDeviceSettingService struct {
	Options []option.RequestOption
}

// NewZeroTrustDeviceSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDeviceSettingService(opts ...option.RequestOption) (r *ZeroTrustDeviceSettingService) {
	r = &ZeroTrustDeviceSettingService{}
	r.Options = opts
	return
}

// Updates the current device settings for a Zero Trust account.
func (r *ZeroTrustDeviceSettingService) Update(ctx context.Context, params ZeroTrustDeviceSettingUpdateParams, opts ...option.RequestOption) (res *TeamsDevicesZeroTrustAccountDeviceSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceSettingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Describes the current device settings for a Zero Trust account.
func (r *ZeroTrustDeviceSettingService) List(ctx context.Context, query ZeroTrustDeviceSettingListParams, opts ...option.RequestOption) (res *TeamsDevicesZeroTrustAccountDeviceSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceSettingListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/settings", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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

type ZeroTrustDeviceSettingUpdateParams struct {
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

func (r ZeroTrustDeviceSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDeviceSettingUpdateResponseEnvelope struct {
	Errors   []ZeroTrustDeviceSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamsDevicesZeroTrustAccountDeviceSettings             `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceSettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceSettingUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceSettingUpdateResponseEnvelope]
type zeroTrustDeviceSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDeviceSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDeviceSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceSettingUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceSettingUpdateResponseEnvelopeErrors]
type zeroTrustDeviceSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDeviceSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustDeviceSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceSettingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceSettingUpdateResponseEnvelopeMessages]
type zeroTrustDeviceSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDeviceSettingUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceSettingUpdateResponseEnvelopeSuccessTrue ZeroTrustDeviceSettingUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustDeviceSettingListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDeviceSettingListResponseEnvelope struct {
	Errors   []ZeroTrustDeviceSettingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceSettingListResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamsDevicesZeroTrustAccountDeviceSettings           `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceSettingListResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceSettingListResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceSettingListResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceSettingListResponseEnvelope]
type zeroTrustDeviceSettingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceSettingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDeviceSettingListResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustDeviceSettingListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceSettingListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceSettingListResponseEnvelopeErrors]
type zeroTrustDeviceSettingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceSettingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDeviceSettingListResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDeviceSettingListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceSettingListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceSettingListResponseEnvelopeMessages]
type zeroTrustDeviceSettingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceSettingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDeviceSettingListResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceSettingListResponseEnvelopeSuccessTrue ZeroTrustDeviceSettingListResponseEnvelopeSuccess = true
)
