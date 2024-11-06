// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// ConfigService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
	Full    *ConfigFullService
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	r.Full = NewConfigFullService(opts...)
	return
}

// Create a new network monitoring configuration.
func (r *ConfigService) New(ctx context.Context, params ConfigNewParams, opts ...option.RequestOption) (res *ConfigNewResponse, err error) {
	var env ConfigNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing network monitoring configuration, requires the entire
// configuration to be updated at once.
func (r *ConfigService) Update(ctx context.Context, params ConfigUpdateParams, opts ...option.RequestOption) (res *ConfigUpdateResponse, err error) {
	var env ConfigUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an existing network monitoring configuration.
func (r *ConfigService) Delete(ctx context.Context, body ConfigDeleteParams, opts ...option.RequestOption) (res *ConfigDeleteResponse, err error) {
	var env ConfigDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update fields in an existing network monitoring configuration.
func (r *ConfigService) Edit(ctx context.Context, params ConfigEditParams, opts ...option.RequestOption) (res *ConfigEditResponse, err error) {
	var env ConfigEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists default sampling, router IPs and warp devices for account.
func (r *ConfigService) Get(ctx context.Context, query ConfigGetParams, opts ...option.RequestOption) (res *ConfigGetResponse, err error) {
	var env ConfigGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConfigNewResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name        string                        `json:"name,required"`
	RouterIPs   []string                      `json:"router_ips,required"`
	WARPDevices []ConfigNewResponseWARPDevice `json:"warp_devices,required"`
	JSON        configNewResponseJSON         `json:"-"`
}

// configNewResponseJSON contains the JSON metadata for the struct
// [ConfigNewResponse]
type configNewResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	WARPDevices     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseJSON) RawJSON() string {
	return r.raw
}

// Object representing a warp device with an ID and name.
type ConfigNewResponseWARPDevice struct {
	// Unique identifier for the warp device.
	ID string `json:"id,required"`
	// Name of the warp device.
	Name string `json:"name,required"`
	// IPv4 CIDR of the router sourcing flow data associated with this warp device.
	// Only /32 addresses are currently supported.
	RouterIP string                          `json:"router_ip,required"`
	JSON     configNewResponseWARPDeviceJSON `json:"-"`
}

// configNewResponseWARPDeviceJSON contains the JSON metadata for the struct
// [ConfigNewResponseWARPDevice]
type configNewResponseWARPDeviceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	RouterIP    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponseWARPDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseWARPDeviceJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name        string                           `json:"name,required"`
	RouterIPs   []string                         `json:"router_ips,required"`
	WARPDevices []ConfigUpdateResponseWARPDevice `json:"warp_devices,required"`
	JSON        configUpdateResponseJSON         `json:"-"`
}

// configUpdateResponseJSON contains the JSON metadata for the struct
// [ConfigUpdateResponse]
type configUpdateResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	WARPDevices     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Object representing a warp device with an ID and name.
type ConfigUpdateResponseWARPDevice struct {
	// Unique identifier for the warp device.
	ID string `json:"id,required"`
	// Name of the warp device.
	Name string `json:"name,required"`
	// IPv4 CIDR of the router sourcing flow data associated with this warp device.
	// Only /32 addresses are currently supported.
	RouterIP string                             `json:"router_ip,required"`
	JSON     configUpdateResponseWARPDeviceJSON `json:"-"`
}

// configUpdateResponseWARPDeviceJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseWARPDevice]
type configUpdateResponseWARPDeviceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	RouterIP    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseWARPDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseWARPDeviceJSON) RawJSON() string {
	return r.raw
}

type ConfigDeleteResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name        string                           `json:"name,required"`
	RouterIPs   []string                         `json:"router_ips,required"`
	WARPDevices []ConfigDeleteResponseWARPDevice `json:"warp_devices,required"`
	JSON        configDeleteResponseJSON         `json:"-"`
}

// configDeleteResponseJSON contains the JSON metadata for the struct
// [ConfigDeleteResponse]
type configDeleteResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	WARPDevices     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConfigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Object representing a warp device with an ID and name.
type ConfigDeleteResponseWARPDevice struct {
	// Unique identifier for the warp device.
	ID string `json:"id,required"`
	// Name of the warp device.
	Name string `json:"name,required"`
	// IPv4 CIDR of the router sourcing flow data associated with this warp device.
	// Only /32 addresses are currently supported.
	RouterIP string                             `json:"router_ip,required"`
	JSON     configDeleteResponseWARPDeviceJSON `json:"-"`
}

// configDeleteResponseWARPDeviceJSON contains the JSON metadata for the struct
// [ConfigDeleteResponseWARPDevice]
type configDeleteResponseWARPDeviceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	RouterIP    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigDeleteResponseWARPDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configDeleteResponseWARPDeviceJSON) RawJSON() string {
	return r.raw
}

type ConfigEditResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name        string                         `json:"name,required"`
	RouterIPs   []string                       `json:"router_ips,required"`
	WARPDevices []ConfigEditResponseWARPDevice `json:"warp_devices,required"`
	JSON        configEditResponseJSON         `json:"-"`
}

// configEditResponseJSON contains the JSON metadata for the struct
// [ConfigEditResponse]
type configEditResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	WARPDevices     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConfigEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseJSON) RawJSON() string {
	return r.raw
}

// Object representing a warp device with an ID and name.
type ConfigEditResponseWARPDevice struct {
	// Unique identifier for the warp device.
	ID string `json:"id,required"`
	// Name of the warp device.
	Name string `json:"name,required"`
	// IPv4 CIDR of the router sourcing flow data associated with this warp device.
	// Only /32 addresses are currently supported.
	RouterIP string                           `json:"router_ip,required"`
	JSON     configEditResponseWARPDeviceJSON `json:"-"`
}

// configEditResponseWARPDeviceJSON contains the JSON metadata for the struct
// [ConfigEditResponseWARPDevice]
type configEditResponseWARPDeviceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	RouterIP    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigEditResponseWARPDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseWARPDeviceJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name        string                        `json:"name,required"`
	RouterIPs   []string                      `json:"router_ips,required"`
	WARPDevices []ConfigGetResponseWARPDevice `json:"warp_devices,required"`
	JSON        configGetResponseJSON         `json:"-"`
}

// configGetResponseJSON contains the JSON metadata for the struct
// [ConfigGetResponse]
type configGetResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	WARPDevices     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseJSON) RawJSON() string {
	return r.raw
}

// Object representing a warp device with an ID and name.
type ConfigGetResponseWARPDevice struct {
	// Unique identifier for the warp device.
	ID string `json:"id,required"`
	// Name of the warp device.
	Name string `json:"name,required"`
	// IPv4 CIDR of the router sourcing flow data associated with this warp device.
	// Only /32 addresses are currently supported.
	RouterIP string                          `json:"router_ip,required"`
	JSON     configGetResponseWARPDeviceJSON `json:"-"`
}

// configGetResponseWARPDeviceJSON contains the JSON metadata for the struct
// [ConfigGetResponseWARPDevice]
type configGetResponseWARPDeviceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	RouterIP    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseWARPDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseWARPDeviceJSON) RawJSON() string {
	return r.raw
}

type ConfigNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling param.Field[float64] `json:"default_sampling,required"`
	// The account name.
	Name        param.Field[string]                      `json:"name,required"`
	RouterIPs   param.Field[[]string]                    `json:"router_ips"`
	WARPDevices param.Field[[]ConfigNewParamsWARPDevice] `json:"warp_devices"`
}

func (r ConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object representing a warp device with an ID and name.
type ConfigNewParamsWARPDevice struct {
	// Unique identifier for the warp device.
	ID param.Field[string] `json:"id,required"`
	// Name of the warp device.
	Name param.Field[string] `json:"name,required"`
	// IPv4 CIDR of the router sourcing flow data associated with this warp device.
	// Only /32 addresses are currently supported.
	RouterIP param.Field[string] `json:"router_ip,required"`
}

func (r ConfigNewParamsWARPDevice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ConfigNewResponse     `json:"result,required"`
	// Whether the API call was successful
	Success ConfigNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    configNewResponseEnvelopeJSON    `json:"-"`
}

// configNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigNewResponseEnvelope]
type configNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigNewResponseEnvelopeSuccess bool

const (
	ConfigNewResponseEnvelopeSuccessTrue ConfigNewResponseEnvelopeSuccess = true
)

func (r ConfigNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling param.Field[float64] `json:"default_sampling,required"`
	// The account name.
	Name        param.Field[string]                         `json:"name,required"`
	RouterIPs   param.Field[[]string]                       `json:"router_ips"`
	WARPDevices param.Field[[]ConfigUpdateParamsWARPDevice] `json:"warp_devices"`
}

func (r ConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object representing a warp device with an ID and name.
type ConfigUpdateParamsWARPDevice struct {
	// Unique identifier for the warp device.
	ID param.Field[string] `json:"id,required"`
	// Name of the warp device.
	Name param.Field[string] `json:"name,required"`
	// IPv4 CIDR of the router sourcing flow data associated with this warp device.
	// Only /32 addresses are currently supported.
	RouterIP param.Field[string] `json:"router_ip,required"`
}

func (r ConfigUpdateParamsWARPDevice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ConfigUpdateResponse  `json:"result,required"`
	// Whether the API call was successful
	Success ConfigUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    configUpdateResponseEnvelopeJSON    `json:"-"`
}

// configUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseEnvelope]
type configUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigUpdateResponseEnvelopeSuccess bool

const (
	ConfigUpdateResponseEnvelopeSuccessTrue ConfigUpdateResponseEnvelopeSuccess = true
)

func (r ConfigUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ConfigDeleteResponse  `json:"result,required"`
	// Whether the API call was successful
	Success ConfigDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    configDeleteResponseEnvelopeJSON    `json:"-"`
}

// configDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigDeleteResponseEnvelope]
type configDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigDeleteResponseEnvelopeSuccess bool

const (
	ConfigDeleteResponseEnvelopeSuccessTrue ConfigDeleteResponseEnvelopeSuccess = true
)

func (r ConfigDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling param.Field[float64] `json:"default_sampling"`
	// The account name.
	Name        param.Field[string]                       `json:"name"`
	RouterIPs   param.Field[[]string]                     `json:"router_ips"`
	WARPDevices param.Field[[]ConfigEditParamsWARPDevice] `json:"warp_devices"`
}

func (r ConfigEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object representing a warp device with an ID and name.
type ConfigEditParamsWARPDevice struct {
	// Unique identifier for the warp device.
	ID param.Field[string] `json:"id,required"`
	// Name of the warp device.
	Name param.Field[string] `json:"name,required"`
	// IPv4 CIDR of the router sourcing flow data associated with this warp device.
	// Only /32 addresses are currently supported.
	RouterIP param.Field[string] `json:"router_ip,required"`
}

func (r ConfigEditParamsWARPDevice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ConfigEditResponse    `json:"result,required"`
	// Whether the API call was successful
	Success ConfigEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    configEditResponseEnvelopeJSON    `json:"-"`
}

// configEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigEditResponseEnvelope]
type configEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigEditResponseEnvelopeSuccess bool

const (
	ConfigEditResponseEnvelopeSuccessTrue ConfigEditResponseEnvelopeSuccess = true
)

func (r ConfigEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ConfigGetResponse     `json:"result,required"`
	// Whether the API call was successful
	Success ConfigGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    configGetResponseEnvelopeJSON    `json:"-"`
}

// configGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigGetResponseEnvelope]
type configGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigGetResponseEnvelopeSuccess bool

const (
	ConfigGetResponseEnvelopeSuccessTrue ConfigGetResponseEnvelopeSuccess = true
)

func (r ConfigGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
