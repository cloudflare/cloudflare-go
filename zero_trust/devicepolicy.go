// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// DevicePolicyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDevicePolicyService] method
// instead.
type DevicePolicyService struct {
	Options         []option.RequestOption
	DefaultPolicy   *DevicePolicyDefaultPolicyService
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
	r.DefaultPolicy = NewDevicePolicyDefaultPolicyService(opts...)
	r.Excludes = NewDevicePolicyExcludeService(opts...)
	r.FallbackDomains = NewDevicePolicyFallbackDomainService(opts...)
	r.Includes = NewDevicePolicyIncludeService(opts...)
	return
}

// Creates a device settings profile to be applied to certain devices matching the
// criteria.
func (r *DevicePolicyService) New(ctx context.Context, params DevicePolicyNewParams, opts ...option.RequestOption) (res *TeamsDevicesDeviceSettingsPolicy, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list of the device settings profiles for an account.
func (r *DevicePolicyService) List(ctx context.Context, query DevicePolicyListParams, opts ...option.RequestOption) (res *[]TeamsDevicesDeviceSettingsPolicy, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policies", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a device settings profile and fetches a list of the remaining profiles
// for an account.
func (r *DevicePolicyService) Delete(ctx context.Context, policyID string, body DevicePolicyDeleteParams, opts ...option.RequestOption) (res *[]TeamsDevicesDeviceSettingsPolicy, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", body.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured device settings profile.
func (r *DevicePolicyService) Edit(ctx context.Context, policyID string, params DevicePolicyEditParams, opts ...option.RequestOption) (res *TeamsDevicesDeviceSettingsPolicy, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a device settings profile by ID.
func (r *DevicePolicyService) Get(ctx context.Context, policyID string, query DevicePolicyGetParams, opts ...option.RequestOption) (res *TeamsDevicesDeviceSettingsPolicy, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TeamsDevicesDeviceSettingsPolicy struct {
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
	Enabled bool                      `json:"enabled"`
	Exclude []TeamsDevicesSplitTunnel `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                             `json:"exclude_office_ips"`
	FallbackDomains  []TeamsDevicesFallbackDomain     `json:"fallback_domains"`
	GatewayUniqueID  string                           `json:"gateway_unique_id"`
	Include          []TeamsDevicesSplitTunnelInclude `json:"include"`
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
	Precedence    float64                                       `json:"precedence"`
	ServiceModeV2 TeamsDevicesDeviceSettingsPolicyServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                 `json:"switch_locked"`
	JSON         teamsDevicesDeviceSettingsPolicyJSON `json:"-"`
}

// teamsDevicesDeviceSettingsPolicyJSON contains the JSON metadata for the struct
// [TeamsDevicesDeviceSettingsPolicy]
type teamsDevicesDeviceSettingsPolicyJSON struct {
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

func (r *TeamsDevicesDeviceSettingsPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDeviceSettingsPolicyJSON) RawJSON() string {
	return r.raw
}

type TeamsDevicesDeviceSettingsPolicyServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                           `json:"port"`
	JSON teamsDevicesDeviceSettingsPolicyServiceModeV2JSON `json:"-"`
}

// teamsDevicesDeviceSettingsPolicyServiceModeV2JSON contains the JSON metadata for
// the struct [TeamsDevicesDeviceSettingsPolicyServiceModeV2]
type teamsDevicesDeviceSettingsPolicyServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamsDevicesDeviceSettingsPolicyServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamsDevicesDeviceSettingsPolicyServiceModeV2JSON) RawJSON() string {
	return r.raw
}

type DevicePolicyNewResponse = interface{}

type DevicePolicyEditResponse = interface{}

type DevicePolicyGetResponse = interface{}

type DevicePolicyNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
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
	LanAllowSubnetSize param.Field[float64]                            `json:"lan_allow_subnet_size"`
	ServiceModeV2      param.Field[DevicePolicyNewParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
}

func (r DevicePolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyNewParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r DevicePolicyNewParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyNewResponseEnvelope struct {
	Errors   []DevicePolicyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamsDevicesDeviceSettingsPolicy          `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyNewResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePolicyNewResponseEnvelope]
type devicePolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    devicePolicyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePolicyNewResponseEnvelopeErrors]
type devicePolicyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    devicePolicyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DevicePolicyNewResponseEnvelopeMessages]
type devicePolicyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyNewResponseEnvelopeSuccess bool

const (
	DevicePolicyNewResponseEnvelopeSuccessTrue DevicePolicyNewResponseEnvelopeSuccess = true
)

type DevicePolicyNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       devicePolicyNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DevicePolicyNewResponseEnvelopeResultInfo]
type devicePolicyNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type DevicePolicyListResponseEnvelope struct {
	Errors   []DevicePolicyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []TeamsDevicesDeviceSettingsPolicy         `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyListResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePolicyListResponseEnvelope]
type devicePolicyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyListResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    devicePolicyListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePolicyListResponseEnvelopeErrors]
type devicePolicyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyListResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    devicePolicyListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DevicePolicyListResponseEnvelopeMessages]
type devicePolicyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyListResponseEnvelopeSuccess bool

const (
	DevicePolicyListResponseEnvelopeSuccessTrue DevicePolicyListResponseEnvelopeSuccess = true
)

type DevicePolicyListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       devicePolicyListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DevicePolicyListResponseEnvelopeResultInfo]
type devicePolicyListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type DevicePolicyDeleteResponseEnvelope struct {
	Errors   []DevicePolicyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   []TeamsDevicesDeviceSettingsPolicy           `json:"result,required,nullable"`
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

func (r devicePolicyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r devicePolicyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r devicePolicyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r devicePolicyDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyEditParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
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
	Precedence    param.Field[float64]                             `json:"precedence"`
	ServiceModeV2 param.Field[DevicePolicyEditParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
}

func (r DevicePolicyEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyEditParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r DevicePolicyEditParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyEditResponseEnvelope struct {
	Errors   []DevicePolicyEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyEditResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamsDevicesDeviceSettingsPolicy           `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyEditResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePolicyEditResponseEnvelope]
type devicePolicyEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyEditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    devicePolicyEditResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DevicePolicyEditResponseEnvelopeErrors]
type devicePolicyEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyEditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    devicePolicyEditResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DevicePolicyEditResponseEnvelopeMessages]
type devicePolicyEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyEditResponseEnvelopeSuccess bool

const (
	DevicePolicyEditResponseEnvelopeSuccessTrue DevicePolicyEditResponseEnvelopeSuccess = true
)

type DevicePolicyEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       devicePolicyEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyEditResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DevicePolicyEditResponseEnvelopeResultInfo]
type devicePolicyEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type DevicePolicyGetResponseEnvelope struct {
	Errors   []DevicePolicyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamsDevicesDeviceSettingsPolicy          `json:"result,required,nullable"`
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

func (r devicePolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r devicePolicyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r devicePolicyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r devicePolicyGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
