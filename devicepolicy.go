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
func (r *DevicePolicyService) New(ctx context.Context, identifier interface{}, body DevicePolicyNewParams, opts ...option.RequestOption) (res *[]DevicePolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

// Fetches a list of the device settings profiles for an account.
func (r *DevicePolicyService) List(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DevicePolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policies", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

type DevicePolicyNewResponse = interface{}

type DevicePolicyUpdateResponse = interface{}

type DevicePolicyListResponse struct {
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
	Enabled bool                              `json:"enabled"`
	Exclude []DevicePolicyListResponseExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                     `json:"exclude_office_ips"`
	FallbackDomains  []DevicePolicyListResponseFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                   `json:"gateway_unique_id"`
	Include          []DevicePolicyListResponseInclude        `json:"include"`
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
	Precedence    float64                               `json:"precedence"`
	ServiceModeV2 DevicePolicyListResponseServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                         `json:"switch_locked"`
	JSON         devicePolicyListResponseJSON `json:"-"`
}

// devicePolicyListResponseJSON contains the JSON metadata for the struct
// [DevicePolicyListResponse]
type devicePolicyListResponseJSON struct {
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

func (r *DevicePolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyListResponseExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                              `json:"host"`
	JSON devicePolicyListResponseExcludeJSON `json:"-"`
}

// devicePolicyListResponseExcludeJSON contains the JSON metadata for the struct
// [DevicePolicyListResponseExclude]
type devicePolicyListResponseExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyListResponseExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyListResponseFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                              `json:"dns_server"`
	JSON      devicePolicyListResponseFallbackDomainJSON `json:"-"`
}

// devicePolicyListResponseFallbackDomainJSON contains the JSON metadata for the
// struct [DevicePolicyListResponseFallbackDomain]
type devicePolicyListResponseFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyListResponseFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyListResponseInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                              `json:"host"`
	JSON devicePolicyListResponseIncludeJSON `json:"-"`
}

// devicePolicyListResponseIncludeJSON contains the JSON metadata for the struct
// [DevicePolicyListResponseInclude]
type devicePolicyListResponseIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyListResponseInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyListResponseServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                   `json:"port"`
	JSON devicePolicyListResponseServiceModeV2JSON `json:"-"`
}

// devicePolicyListResponseServiceModeV2JSON contains the JSON metadata for the
// struct [DevicePolicyListResponseServiceModeV2]
type devicePolicyListResponseServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyListResponseServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type DevicePolicyGetResponse = interface{}

type DevicePolicyNewParams struct {
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
	Result   []DevicePolicyNewResponse                 `json:"result,required,nullable"`
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

type DevicePolicyListResponseEnvelope struct {
	Errors   []DevicePolicyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyListResponse                 `json:"result,required,nullable"`
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
