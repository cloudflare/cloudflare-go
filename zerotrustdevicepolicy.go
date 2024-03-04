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

// ZeroTrustDevicePolicyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDevicePolicyService]
// method instead.
type ZeroTrustDevicePolicyService struct {
	Options         []option.RequestOption
	DefaultPolicy   *ZeroTrustDevicePolicyDefaultPolicyService
	Excludes        *ZeroTrustDevicePolicyExcludeService
	FallbackDomains *ZeroTrustDevicePolicyFallbackDomainService
	Includes        *ZeroTrustDevicePolicyIncludeService
}

// NewZeroTrustDevicePolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDevicePolicyService(opts ...option.RequestOption) (r *ZeroTrustDevicePolicyService) {
	r = &ZeroTrustDevicePolicyService{}
	r.Options = opts
	r.DefaultPolicy = NewZeroTrustDevicePolicyDefaultPolicyService(opts...)
	r.Excludes = NewZeroTrustDevicePolicyExcludeService(opts...)
	r.FallbackDomains = NewZeroTrustDevicePolicyFallbackDomainService(opts...)
	r.Includes = NewZeroTrustDevicePolicyIncludeService(opts...)
	return
}

// Creates a device settings profile to be applied to certain devices matching the
// criteria.
func (r *ZeroTrustDevicePolicyService) New(ctx context.Context, params ZeroTrustDevicePolicyNewParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list of the device settings profiles for an account.
func (r *ZeroTrustDevicePolicyService) List(ctx context.Context, query ZeroTrustDevicePolicyListParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyListResponseEnvelope
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
func (r *ZeroTrustDevicePolicyService) Delete(ctx context.Context, policyID string, body ZeroTrustDevicePolicyDeleteParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", body.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured device settings profile.
func (r *ZeroTrustDevicePolicyService) Edit(ctx context.Context, policyID string, params ZeroTrustDevicePolicyEditParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a device settings profile by ID.
func (r *ZeroTrustDevicePolicyService) Get(ctx context.Context, policyID string, query ZeroTrustDevicePolicyGetParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDevicePolicyNewResponse = interface{}

type ZeroTrustDevicePolicyListResponse struct {
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
	Enabled bool                                       `json:"enabled"`
	Exclude []ZeroTrustDevicePolicyListResponseExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                              `json:"exclude_office_ips"`
	FallbackDomains  []ZeroTrustDevicePolicyListResponseFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                            `json:"gateway_unique_id"`
	Include          []ZeroTrustDevicePolicyListResponseInclude        `json:"include"`
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
	Precedence    float64                                        `json:"precedence"`
	ServiceModeV2 ZeroTrustDevicePolicyListResponseServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                  `json:"switch_locked"`
	JSON         zeroTrustDevicePolicyListResponseJSON `json:"-"`
}

// zeroTrustDevicePolicyListResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDevicePolicyListResponse]
type zeroTrustDevicePolicyListResponseJSON struct {
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

func (r *ZeroTrustDevicePolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyListResponseExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                       `json:"host"`
	JSON zeroTrustDevicePolicyListResponseExcludeJSON `json:"-"`
}

// zeroTrustDevicePolicyListResponseExcludeJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePolicyListResponseExclude]
type zeroTrustDevicePolicyListResponseExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyListResponseExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyListResponseFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                       `json:"dns_server"`
	JSON      zeroTrustDevicePolicyListResponseFallbackDomainJSON `json:"-"`
}

// zeroTrustDevicePolicyListResponseFallbackDomainJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyListResponseFallbackDomain]
type zeroTrustDevicePolicyListResponseFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyListResponseFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyListResponseInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                       `json:"host"`
	JSON zeroTrustDevicePolicyListResponseIncludeJSON `json:"-"`
}

// zeroTrustDevicePolicyListResponseIncludeJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePolicyListResponseInclude]
type zeroTrustDevicePolicyListResponseIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyListResponseInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyListResponseServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                            `json:"port"`
	JSON zeroTrustDevicePolicyListResponseServiceModeV2JSON `json:"-"`
}

// zeroTrustDevicePolicyListResponseServiceModeV2JSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyListResponseServiceModeV2]
type zeroTrustDevicePolicyListResponseServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyListResponseServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyDeleteResponse struct {
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
	Enabled bool                                         `json:"enabled"`
	Exclude []ZeroTrustDevicePolicyDeleteResponseExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                                `json:"exclude_office_ips"`
	FallbackDomains  []ZeroTrustDevicePolicyDeleteResponseFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                              `json:"gateway_unique_id"`
	Include          []ZeroTrustDevicePolicyDeleteResponseInclude        `json:"include"`
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
	Precedence    float64                                          `json:"precedence"`
	ServiceModeV2 ZeroTrustDevicePolicyDeleteResponseServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                    `json:"switch_locked"`
	JSON         zeroTrustDevicePolicyDeleteResponseJSON `json:"-"`
}

// zeroTrustDevicePolicyDeleteResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePolicyDeleteResponse]
type zeroTrustDevicePolicyDeleteResponseJSON struct {
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

func (r *ZeroTrustDevicePolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyDeleteResponseExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                         `json:"host"`
	JSON zeroTrustDevicePolicyDeleteResponseExcludeJSON `json:"-"`
}

// zeroTrustDevicePolicyDeleteResponseExcludeJSON contains the JSON metadata for
// the struct [ZeroTrustDevicePolicyDeleteResponseExclude]
type zeroTrustDevicePolicyDeleteResponseExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyDeleteResponseExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyDeleteResponseFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                         `json:"dns_server"`
	JSON      zeroTrustDevicePolicyDeleteResponseFallbackDomainJSON `json:"-"`
}

// zeroTrustDevicePolicyDeleteResponseFallbackDomainJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyDeleteResponseFallbackDomain]
type zeroTrustDevicePolicyDeleteResponseFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyDeleteResponseFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyDeleteResponseInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                         `json:"host"`
	JSON zeroTrustDevicePolicyDeleteResponseIncludeJSON `json:"-"`
}

// zeroTrustDevicePolicyDeleteResponseIncludeJSON contains the JSON metadata for
// the struct [ZeroTrustDevicePolicyDeleteResponseInclude]
type zeroTrustDevicePolicyDeleteResponseIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyDeleteResponseInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyDeleteResponseServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                              `json:"port"`
	JSON zeroTrustDevicePolicyDeleteResponseServiceModeV2JSON `json:"-"`
}

// zeroTrustDevicePolicyDeleteResponseServiceModeV2JSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyDeleteResponseServiceModeV2]
type zeroTrustDevicePolicyDeleteResponseServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyDeleteResponseServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyEditResponse = interface{}

type ZeroTrustDevicePolicyGetResponse = interface{}

type ZeroTrustDevicePolicyNewParams struct {
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
	LanAllowSubnetSize param.Field[float64]                                     `json:"lan_allow_subnet_size"`
	ServiceModeV2      param.Field[ZeroTrustDevicePolicyNewParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
}

func (r ZeroTrustDevicePolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDevicePolicyNewParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r ZeroTrustDevicePolicyNewParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDevicePolicyNewResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyNewResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyNewResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePolicyNewResponseEnvelope]
type zeroTrustDevicePolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustDevicePolicyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyNewResponseEnvelopeErrors]
type zeroTrustDevicePolicyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustDevicePolicyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyNewResponseEnvelopeMessages]
type zeroTrustDevicePolicyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyNewResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyNewResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyNewResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       zeroTrustDevicePolicyNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyNewResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyNewResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePolicyListResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePolicyListResponseEnvelope]
type zeroTrustDevicePolicyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyListResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDevicePolicyListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyListResponseEnvelopeErrors]
type zeroTrustDevicePolicyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyListResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustDevicePolicyListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyListResponseEnvelopeMessages]
type zeroTrustDevicePolicyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyListResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyListResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyListResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       zeroTrustDevicePolicyListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyListResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePolicyDeleteResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyDeleteResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyDeleteResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDevicePolicyDeleteResponseEnvelope]
type zeroTrustDevicePolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyDeleteResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustDevicePolicyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyDeleteResponseEnvelopeErrors]
type zeroTrustDevicePolicyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyDeleteResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustDevicePolicyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyDeleteResponseEnvelopeMessages]
type zeroTrustDevicePolicyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyDeleteResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       zeroTrustDevicePolicyDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyDeleteResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyEditParams struct {
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
	Precedence    param.Field[float64]                                      `json:"precedence"`
	ServiceModeV2 param.Field[ZeroTrustDevicePolicyEditParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
}

func (r ZeroTrustDevicePolicyEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDevicePolicyEditParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r ZeroTrustDevicePolicyEditParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDevicePolicyEditResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyEditResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyEditResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyEditResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePolicyEditResponseEnvelope]
type zeroTrustDevicePolicyEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDevicePolicyEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyEditResponseEnvelopeErrors]
type zeroTrustDevicePolicyEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustDevicePolicyEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyEditResponseEnvelopeMessages]
type zeroTrustDevicePolicyEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyEditResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyEditResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyEditResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       zeroTrustDevicePolicyEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyEditResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyEditResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePolicyGetResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyGetResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDevicePolicyGetResponseEnvelope]
type zeroTrustDevicePolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustDevicePolicyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyGetResponseEnvelopeErrors]
type zeroTrustDevicePolicyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustDevicePolicyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDevicePolicyGetResponseEnvelopeMessages]
type zeroTrustDevicePolicyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyGetResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyGetResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       zeroTrustDevicePolicyGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyGetResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
