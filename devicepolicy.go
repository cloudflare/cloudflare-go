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

// DevicePolicyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDevicePolicyService] method
// instead.
type DevicePolicyService struct {
	Options         []option.RequestOption
	Excludes        *DevicePolicyExcludeService
	FallbackDomains *DevicePolicyFallbackDomainService
	Includes        *DevicePolicyIncludeService
}

// NewDevicePolicyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDevicePolicyService(opts ...option.RequestOption) (r *DevicePolicyService) {
	r = &DevicePolicyService{}
	r.Options = opts
	r.Excludes = NewDevicePolicyExcludeService(opts...)
	r.FallbackDomains = NewDevicePolicyFallbackDomainService(opts...)
	r.Includes = NewDevicePolicyIncludeService(opts...)
	return
}

// Updates a configured device settings profile.
func (r *DevicePolicyService) Update(ctx context.Context, identifier interface{}, uuid string, body DevicePolicyUpdateParams, opts ...option.RequestOption) (res *[]DevicePolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a device settings profile and fetches a list of the remaining profiles
// for an account.
func (r *DevicePolicyService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *[]DevicePolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a device settings profile to be applied to certain devices matching the
// criteria.
func (r *DevicePolicyService) DevicesNewDeviceSettingsPolicy(ctx context.Context, identifier interface{}, body DevicePolicyDevicesNewDeviceSettingsPolicyParams, opts ...option.RequestOption) (res *[]DevicePolicyDevicesNewDeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the default device settings profile for an account.
func (r *DevicePolicyService) DevicesGetDefaultDeviceSettingsPolicy(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list of the device settings profiles for an account.
func (r *DevicePolicyService) DevicesListDeviceSettingsPolicies(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DevicePolicyDevicesListDeviceSettingsPoliciesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policies", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the default device settings profile for an account.
func (r *DevicePolicyService) DevicesUpdateDefaultDeviceSettingsPolicy(ctx context.Context, identifier interface{}, body DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParams, opts ...option.RequestOption) (res *[]DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a device settings profile by ID.
func (r *DevicePolicyService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *[]DevicePolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyUpdateResponse = interface{}

type DevicePolicyDeleteResponse struct {
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch bool `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates bool `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave bool `json:"allowed_to_leave"`
	// The amount of time in minutes to reconnect after having been disabled.
	AutoConnect float64 `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal float64 `json:"captive_portal"`
	// Whether the policy is the default policy for an account.
	Default bool `json:"default"`
	// A description of the policy.
	Description string `json:"description"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback bool `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled bool                                `json:"enabled"`
	Exclude []DevicePolicyDeleteResponseExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                       `json:"exclude_office_ips"`
	FallbackDomains  []DevicePolicyDeleteResponseFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                     `json:"gateway_unique_id"`
	Include          []DevicePolicyDeleteResponseInclude        `json:"include"`
	// The amount of time in minutes a user is allowed access to their LAN. A value of
	// 0 will allow LAN access until the next WARP reconnection, such as a reboot or a
	// laptop waking from sleep. Note that this field is omitted from the response if
	// null or unset.
	LanAllowMinutes float64 `json:"lan_allow_minutes"`
	// The size of the subnet for the local access network. Note that this field is
	// omitted from the response if null or unset.
	LanAllowSubnetSize float64 `json:"lan_allow_subnet_size"`
	// The wirefilter expression to match devices.
	Match string `json:"match"`
	// The name of the device settings profile.
	Name string `json:"name"`
	// Device ID.
	PolicyID string `json:"policy_id"`
	// The precedence of the policy. Lower values indicate higher precedence. Policies
	// will be evaluated in ascending order of this field.
	Precedence    float64                                 `json:"precedence"`
	ServiceModeV2 DevicePolicyDeleteResponseServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                           `json:"switch_locked"`
	JSON         devicePolicyDeleteResponseJSON `json:"-"`
}

// devicePolicyDeleteResponseJSON contains the JSON metadata for the struct
// [DevicePolicyDeleteResponse]
type devicePolicyDeleteResponseJSON struct {
	AllowModeSwitch     apijson.Field
	AllowUpdates        apijson.Field
	AllowedToLeave      apijson.Field
	AutoConnect         apijson.Field
	CaptivePortal       apijson.Field
	Default             apijson.Field
	Description         apijson.Field
	DisableAutoFallback apijson.Field
	Enabled             apijson.Field
	Exclude             apijson.Field
	ExcludeOfficeIPs    apijson.Field
	FallbackDomains     apijson.Field
	GatewayUniqueID     apijson.Field
	Include             apijson.Field
	LanAllowMinutes     apijson.Field
	LanAllowSubnetSize  apijson.Field
	Match               apijson.Field
	Name                apijson.Field
	PolicyID            apijson.Field
	Precedence          apijson.Field
	ServiceModeV2       apijson.Field
	SupportURL          apijson.Field
	SwitchLocked        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DevicePolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDeleteResponseExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                `json:"host"`
	JSON devicePolicyDeleteResponseExcludeJSON `json:"-"`
}

// devicePolicyDeleteResponseExcludeJSON contains the JSON metadata for the struct
// [DevicePolicyDeleteResponseExclude]
type devicePolicyDeleteResponseExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDeleteResponseExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDeleteResponseFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                `json:"dns_server"`
	JSON      devicePolicyDeleteResponseFallbackDomainJSON `json:"-"`
}

// devicePolicyDeleteResponseFallbackDomainJSON contains the JSON metadata for the
// struct [DevicePolicyDeleteResponseFallbackDomain]
type devicePolicyDeleteResponseFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDeleteResponseFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDeleteResponseInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                `json:"host"`
	JSON devicePolicyDeleteResponseIncludeJSON `json:"-"`
}

// devicePolicyDeleteResponseIncludeJSON contains the JSON metadata for the struct
// [DevicePolicyDeleteResponseInclude]
type devicePolicyDeleteResponseIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDeleteResponseInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDeleteResponseServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                     `json:"port"`
	JSON devicePolicyDeleteResponseServiceModeV2JSON `json:"-"`
}

// devicePolicyDeleteResponseServiceModeV2JSON contains the JSON metadata for the
// struct [DevicePolicyDeleteResponseServiceModeV2]
type devicePolicyDeleteResponseServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDeleteResponseServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesNewDeviceSettingsPolicyResponse = interface{}

type DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponse = interface{}

type DevicePolicyDevicesListDeviceSettingsPoliciesResponse struct {
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch bool `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates bool `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave bool `json:"allowed_to_leave"`
	// The amount of time in minutes to reconnect after having been disabled.
	AutoConnect float64 `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal float64 `json:"captive_portal"`
	// Whether the policy is the default policy for an account.
	Default bool `json:"default"`
	// A description of the policy.
	Description string `json:"description"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback bool `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled bool                                                           `json:"enabled"`
	Exclude []DevicePolicyDevicesListDeviceSettingsPoliciesResponseExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                                                  `json:"exclude_office_ips"`
	FallbackDomains  []DevicePolicyDevicesListDeviceSettingsPoliciesResponseFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                                                `json:"gateway_unique_id"`
	Include          []DevicePolicyDevicesListDeviceSettingsPoliciesResponseInclude        `json:"include"`
	// The amount of time in minutes a user is allowed access to their LAN. A value of
	// 0 will allow LAN access until the next WARP reconnection, such as a reboot or a
	// laptop waking from sleep. Note that this field is omitted from the response if
	// null or unset.
	LanAllowMinutes float64 `json:"lan_allow_minutes"`
	// The size of the subnet for the local access network. Note that this field is
	// omitted from the response if null or unset.
	LanAllowSubnetSize float64 `json:"lan_allow_subnet_size"`
	// The wirefilter expression to match devices.
	Match string `json:"match"`
	// The name of the device settings profile.
	Name string `json:"name"`
	// Device ID.
	PolicyID string `json:"policy_id"`
	// The precedence of the policy. Lower values indicate higher precedence. Policies
	// will be evaluated in ascending order of this field.
	Precedence    float64                                                            `json:"precedence"`
	ServiceModeV2 DevicePolicyDevicesListDeviceSettingsPoliciesResponseServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                                      `json:"switch_locked"`
	JSON         devicePolicyDevicesListDeviceSettingsPoliciesResponseJSON `json:"-"`
}

// devicePolicyDevicesListDeviceSettingsPoliciesResponseJSON contains the JSON
// metadata for the struct [DevicePolicyDevicesListDeviceSettingsPoliciesResponse]
type devicePolicyDevicesListDeviceSettingsPoliciesResponseJSON struct {
	AllowModeSwitch     apijson.Field
	AllowUpdates        apijson.Field
	AllowedToLeave      apijson.Field
	AutoConnect         apijson.Field
	CaptivePortal       apijson.Field
	Default             apijson.Field
	Description         apijson.Field
	DisableAutoFallback apijson.Field
	Enabled             apijson.Field
	Exclude             apijson.Field
	ExcludeOfficeIPs    apijson.Field
	FallbackDomains     apijson.Field
	GatewayUniqueID     apijson.Field
	Include             apijson.Field
	LanAllowMinutes     apijson.Field
	LanAllowSubnetSize  apijson.Field
	Match               apijson.Field
	Name                apijson.Field
	PolicyID            apijson.Field
	Precedence          apijson.Field
	ServiceModeV2       apijson.Field
	SupportURL          apijson.Field
	SwitchLocked        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DevicePolicyDevicesListDeviceSettingsPoliciesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesListDeviceSettingsPoliciesResponseExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                           `json:"host"`
	JSON devicePolicyDevicesListDeviceSettingsPoliciesResponseExcludeJSON `json:"-"`
}

// devicePolicyDevicesListDeviceSettingsPoliciesResponseExcludeJSON contains the
// JSON metadata for the struct
// [DevicePolicyDevicesListDeviceSettingsPoliciesResponseExclude]
type devicePolicyDevicesListDeviceSettingsPoliciesResponseExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesListDeviceSettingsPoliciesResponseExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesListDeviceSettingsPoliciesResponseFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                           `json:"dns_server"`
	JSON      devicePolicyDevicesListDeviceSettingsPoliciesResponseFallbackDomainJSON `json:"-"`
}

// devicePolicyDevicesListDeviceSettingsPoliciesResponseFallbackDomainJSON contains
// the JSON metadata for the struct
// [DevicePolicyDevicesListDeviceSettingsPoliciesResponseFallbackDomain]
type devicePolicyDevicesListDeviceSettingsPoliciesResponseFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesListDeviceSettingsPoliciesResponseFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesListDeviceSettingsPoliciesResponseInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                           `json:"host"`
	JSON devicePolicyDevicesListDeviceSettingsPoliciesResponseIncludeJSON `json:"-"`
}

// devicePolicyDevicesListDeviceSettingsPoliciesResponseIncludeJSON contains the
// JSON metadata for the struct
// [DevicePolicyDevicesListDeviceSettingsPoliciesResponseInclude]
type devicePolicyDevicesListDeviceSettingsPoliciesResponseIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesListDeviceSettingsPoliciesResponseInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesListDeviceSettingsPoliciesResponseServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                                                `json:"port"`
	JSON devicePolicyDevicesListDeviceSettingsPoliciesResponseServiceModeV2JSON `json:"-"`
}

// devicePolicyDevicesListDeviceSettingsPoliciesResponseServiceModeV2JSON contains
// the JSON metadata for the struct
// [DevicePolicyDevicesListDeviceSettingsPoliciesResponseServiceModeV2]
type devicePolicyDevicesListDeviceSettingsPoliciesResponseServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesListDeviceSettingsPoliciesResponseServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponse = interface{}

type DevicePolicyGetResponse = interface{}

type DevicePolicyUpdateParams struct {
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch param.Field[bool] `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates param.Field[bool] `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave param.Field[bool] `json:"allowed_to_leave"`
	// The amount of time in minutes to reconnect after having been disabled.
	AutoConnect param.Field[float64] `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal param.Field[float64] `json:"captive_portal"`
	// A description of the policy.
	Description param.Field[string] `json:"description"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback param.Field[bool] `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled param.Field[bool] `json:"enabled"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs param.Field[bool] `json:"exclude_office_ips"`
	// The wirefilter expression to match devices.
	Match param.Field[string] `json:"match"`
	// The name of the device settings profile.
	Name param.Field[string] `json:"name"`
	// The precedence of the policy. Lower values indicate higher precedence. Policies
	// will be evaluated in ascending order of this field.
	Precedence    param.Field[float64]                               `json:"precedence"`
	ServiceModeV2 param.Field[DevicePolicyUpdateParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
}

func (r DevicePolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyUpdateParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r DevicePolicyUpdateParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyUpdateResponseEnvelope struct {
	Errors   []DevicePolicyUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePolicyUpdateResponseEnvelope]
type devicePolicyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    devicePolicyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePolicyUpdateResponseEnvelopeErrors]
type devicePolicyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    devicePolicyUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DevicePolicyUpdateResponseEnvelopeMessages]
type devicePolicyUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyUpdateResponseEnvelopeSuccessTrue DevicePolicyUpdateResponseEnvelopeSuccess = true
)

type DevicePolicyUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       devicePolicyUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DevicePolicyUpdateResponseEnvelopeResultInfo]
type devicePolicyUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDeleteResponseEnvelope struct {
	Errors   []DevicePolicyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyDeleteResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyDeleteResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePolicyDeleteResponseEnvelope]
type devicePolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDeleteResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    devicePolicyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePolicyDeleteResponseEnvelopeErrors]
type devicePolicyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDeleteResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    devicePolicyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DevicePolicyDeleteResponseEnvelopeMessages]
type devicePolicyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyDeleteResponseEnvelopeSuccess bool

const (
	DevicePolicyDeleteResponseEnvelopeSuccessTrue DevicePolicyDeleteResponseEnvelopeSuccess = true
)

type DevicePolicyDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       devicePolicyDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DevicePolicyDeleteResponseEnvelopeResultInfo]
type devicePolicyDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesNewDeviceSettingsPolicyParams struct {
	// The wirefilter expression to match devices.
	Match param.Field[string] `json:"match,required"`
	// The name of the device settings profile.
	Name param.Field[string] `json:"name,required"`
	// The precedence of the policy. Lower values indicate higher precedence. Policies
	// will be evaluated in ascending order of this field.
	Precedence param.Field[float64] `json:"precedence,required"`
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch param.Field[bool] `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates param.Field[bool] `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave param.Field[bool] `json:"allowed_to_leave"`
	// The amount of time in minutes to reconnect after having been disabled.
	AutoConnect param.Field[float64] `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal param.Field[float64] `json:"captive_portal"`
	// A description of the policy.
	Description param.Field[string] `json:"description"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback param.Field[bool] `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled param.Field[bool] `json:"enabled"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs param.Field[bool] `json:"exclude_office_ips"`
	// The amount of time in minutes a user is allowed access to their LAN. A value of
	// 0 will allow LAN access until the next WARP reconnection, such as a reboot or a
	// laptop waking from sleep. Note that this field is omitted from the response if
	// null or unset.
	LanAllowMinutes param.Field[float64] `json:"lan_allow_minutes"`
	// The size of the subnet for the local access network. Note that this field is
	// omitted from the response if null or unset.
	LanAllowSubnetSize param.Field[float64]                                                       `json:"lan_allow_subnet_size"`
	ServiceModeV2      param.Field[DevicePolicyDevicesNewDeviceSettingsPolicyParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
}

func (r DevicePolicyDevicesNewDeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyDevicesNewDeviceSettingsPolicyParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r DevicePolicyDevicesNewDeviceSettingsPolicyParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelope struct {
	Errors   []DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyDevicesNewDeviceSettingsPolicyResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelope]
type devicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeErrors struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    devicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeErrors]
type devicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeMessages struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    devicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeMessages]
type devicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeSuccess bool

const (
	DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeSuccessTrue DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeSuccess = true
)

type DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                  `json:"total_count"`
	JSON       devicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeResultInfo]
type devicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesNewDeviceSettingsPolicyResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelope struct {
	Errors   []DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelope]
type devicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    devicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeErrors]
type devicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    devicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeMessages]
type devicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeSuccess bool

const (
	DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeSuccessTrue DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeSuccess = true
)

type DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                         `json:"total_count"`
	JSON       devicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeResultInfo]
type devicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelope struct {
	Errors   []DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyDevicesListDeviceSettingsPoliciesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelope]
type devicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeErrors struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    devicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeErrors]
type devicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeMessages struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    devicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeMessages]
type devicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeSuccess bool

const (
	DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeSuccessTrue DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeSuccess = true
)

type DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                     `json:"total_count"`
	JSON       devicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeResultInfo]
type devicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesListDeviceSettingsPoliciesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParams struct {
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch param.Field[bool] `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates param.Field[bool] `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave param.Field[bool] `json:"allowed_to_leave"`
	// The amount of time in minutes to reconnect after having been disabled.
	AutoConnect param.Field[float64] `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal param.Field[float64] `json:"captive_portal"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback param.Field[bool] `json:"disable_auto_fallback"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs param.Field[bool]                                                                    `json:"exclude_office_ips"`
	ServiceModeV2    param.Field[DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
}

func (r DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelope struct {
	Errors   []DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelope]
type devicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeErrors struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    devicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeErrors]
type devicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeMessages struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    devicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeMessages]
type devicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeSuccess bool

const (
	DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeSuccessTrue DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeSuccess = true
)

type DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                            `json:"total_count"`
	JSON       devicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeResultInfo]
type devicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyGetResponseEnvelope struct {
	Errors   []DevicePolicyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePolicyGetResponseEnvelope]
type devicePolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    devicePolicyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePolicyGetResponseEnvelopeErrors]
type devicePolicyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    devicePolicyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DevicePolicyGetResponseEnvelopeMessages]
type devicePolicyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyGetResponseEnvelopeSuccess bool

const (
	DevicePolicyGetResponseEnvelopeSuccessTrue DevicePolicyGetResponseEnvelopeSuccess = true
)

type DevicePolicyGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       devicePolicyGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DevicePolicyGetResponseEnvelopeResultInfo]
type devicePolicyGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
