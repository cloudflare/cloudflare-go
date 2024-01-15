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

// AccountDevicePolicyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDevicePolicyService]
// method instead.
type AccountDevicePolicyService struct {
	Options         []option.RequestOption
	Excludes        *AccountDevicePolicyExcludeService
	FallbackDomains *AccountDevicePolicyFallbackDomainService
	Includes        *AccountDevicePolicyIncludeService
}

// NewAccountDevicePolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDevicePolicyService(opts ...option.RequestOption) (r *AccountDevicePolicyService) {
	r = &AccountDevicePolicyService{}
	r.Options = opts
	r.Excludes = NewAccountDevicePolicyExcludeService(opts...)
	r.FallbackDomains = NewAccountDevicePolicyFallbackDomainService(opts...)
	r.Includes = NewAccountDevicePolicyIncludeService(opts...)
	return
}

// Fetches a device settings profile by ID.
func (r *AccountDevicePolicyService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDevicePolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured device settings profile.
func (r *AccountDevicePolicyService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountDevicePolicyUpdateParams, opts ...option.RequestOption) (res *AccountDevicePolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes a device settings profile and fetches a list of the remaining profiles
// for an account.
func (r *AccountDevicePolicyService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDevicePolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a device settings profile to be applied to certain devices matching the
// criteria.
func (r *AccountDevicePolicyService) DevicesNewDeviceSettingsPolicy(ctx context.Context, identifier interface{}, body AccountDevicePolicyDevicesNewDeviceSettingsPolicyParams, opts ...option.RequestOption) (res *AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the default device settings profile for an account.
func (r *AccountDevicePolicyService) DevicesGetDefaultDeviceSettingsPolicy(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a list of the device settings profiles for an account.
func (r *AccountDevicePolicyService) DevicesListDeviceSettingsPolicies(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policies", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the default device settings profile for an account.
func (r *AccountDevicePolicyService) DevicesUpdateDefaultDeviceSettingsPolicy(ctx context.Context, identifier interface{}, body AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParams, opts ...option.RequestOption) (res *AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccountDevicePolicyGetResponse struct {
	Errors     []AccountDevicePolicyGetResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyGetResponseMessage  `json:"messages"`
	Result     AccountDevicePolicyGetResponseResult     `json:"result"`
	ResultInfo AccountDevicePolicyGetResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyGetResponseSuccess `json:"success"`
	JSON    accountDevicePolicyGetResponseJSON    `json:"-"`
}

// accountDevicePolicyGetResponseJSON contains the JSON metadata for the struct
// [AccountDevicePolicyGetResponse]
type accountDevicePolicyGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyGetResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountDevicePolicyGetResponseErrorJSON `json:"-"`
}

// accountDevicePolicyGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountDevicePolicyGetResponseError]
type accountDevicePolicyGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyGetResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountDevicePolicyGetResponseMessageJSON `json:"-"`
}

// accountDevicePolicyGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountDevicePolicyGetResponseMessage]
type accountDevicePolicyGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyGetResponseResult struct {
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
	Enabled bool                                          `json:"enabled"`
	Exclude []AccountDevicePolicyGetResponseResultExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                                 `json:"exclude_office_ips"`
	FallbackDomains  []AccountDevicePolicyGetResponseResultFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                               `json:"gateway_unique_id"`
	Include          []AccountDevicePolicyGetResponseResultInclude        `json:"include"`
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
	Precedence    float64                                           `json:"precedence"`
	ServiceModeV2 AccountDevicePolicyGetResponseResultServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                     `json:"switch_locked"`
	JSON         accountDevicePolicyGetResponseResultJSON `json:"-"`
}

// accountDevicePolicyGetResponseResultJSON contains the JSON metadata for the
// struct [AccountDevicePolicyGetResponseResult]
type accountDevicePolicyGetResponseResultJSON struct {
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

func (r *AccountDevicePolicyGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyGetResponseResultExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                          `json:"host"`
	JSON accountDevicePolicyGetResponseResultExcludeJSON `json:"-"`
}

// accountDevicePolicyGetResponseResultExcludeJSON contains the JSON metadata for
// the struct [AccountDevicePolicyGetResponseResultExclude]
type accountDevicePolicyGetResponseResultExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyGetResponseResultExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyGetResponseResultFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                          `json:"dns_server"`
	JSON      accountDevicePolicyGetResponseResultFallbackDomainJSON `json:"-"`
}

// accountDevicePolicyGetResponseResultFallbackDomainJSON contains the JSON
// metadata for the struct [AccountDevicePolicyGetResponseResultFallbackDomain]
type accountDevicePolicyGetResponseResultFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyGetResponseResultFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyGetResponseResultInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                          `json:"host"`
	JSON accountDevicePolicyGetResponseResultIncludeJSON `json:"-"`
}

// accountDevicePolicyGetResponseResultIncludeJSON contains the JSON metadata for
// the struct [AccountDevicePolicyGetResponseResultInclude]
type accountDevicePolicyGetResponseResultIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyGetResponseResultInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyGetResponseResultServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                               `json:"port"`
	JSON accountDevicePolicyGetResponseResultServiceModeV2JSON `json:"-"`
}

// accountDevicePolicyGetResponseResultServiceModeV2JSON contains the JSON metadata
// for the struct [AccountDevicePolicyGetResponseResultServiceModeV2]
type accountDevicePolicyGetResponseResultServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyGetResponseResultServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyGetResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       accountDevicePolicyGetResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyGetResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountDevicePolicyGetResponseResultInfo]
type accountDevicePolicyGetResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyGetResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyGetResponseSuccess bool

const (
	AccountDevicePolicyGetResponseSuccessTrue AccountDevicePolicyGetResponseSuccess = true
)

type AccountDevicePolicyUpdateResponse struct {
	Errors     []AccountDevicePolicyUpdateResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyUpdateResponseMessage  `json:"messages"`
	Result     AccountDevicePolicyUpdateResponseResult     `json:"result"`
	ResultInfo AccountDevicePolicyUpdateResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyUpdateResponseSuccess `json:"success"`
	JSON    accountDevicePolicyUpdateResponseJSON    `json:"-"`
}

// accountDevicePolicyUpdateResponseJSON contains the JSON metadata for the struct
// [AccountDevicePolicyUpdateResponse]
type accountDevicePolicyUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyUpdateResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountDevicePolicyUpdateResponseErrorJSON `json:"-"`
}

// accountDevicePolicyUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountDevicePolicyUpdateResponseError]
type accountDevicePolicyUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyUpdateResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountDevicePolicyUpdateResponseMessageJSON `json:"-"`
}

// accountDevicePolicyUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountDevicePolicyUpdateResponseMessage]
type accountDevicePolicyUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyUpdateResponseResult struct {
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
	Enabled bool                                             `json:"enabled"`
	Exclude []AccountDevicePolicyUpdateResponseResultExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                                    `json:"exclude_office_ips"`
	FallbackDomains  []AccountDevicePolicyUpdateResponseResultFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                                  `json:"gateway_unique_id"`
	Include          []AccountDevicePolicyUpdateResponseResultInclude        `json:"include"`
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
	Precedence    float64                                              `json:"precedence"`
	ServiceModeV2 AccountDevicePolicyUpdateResponseResultServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                        `json:"switch_locked"`
	JSON         accountDevicePolicyUpdateResponseResultJSON `json:"-"`
}

// accountDevicePolicyUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountDevicePolicyUpdateResponseResult]
type accountDevicePolicyUpdateResponseResultJSON struct {
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

func (r *AccountDevicePolicyUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyUpdateResponseResultExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                             `json:"host"`
	JSON accountDevicePolicyUpdateResponseResultExcludeJSON `json:"-"`
}

// accountDevicePolicyUpdateResponseResultExcludeJSON contains the JSON metadata
// for the struct [AccountDevicePolicyUpdateResponseResultExclude]
type accountDevicePolicyUpdateResponseResultExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyUpdateResponseResultExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyUpdateResponseResultFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                             `json:"dns_server"`
	JSON      accountDevicePolicyUpdateResponseResultFallbackDomainJSON `json:"-"`
}

// accountDevicePolicyUpdateResponseResultFallbackDomainJSON contains the JSON
// metadata for the struct [AccountDevicePolicyUpdateResponseResultFallbackDomain]
type accountDevicePolicyUpdateResponseResultFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyUpdateResponseResultFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyUpdateResponseResultInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                             `json:"host"`
	JSON accountDevicePolicyUpdateResponseResultIncludeJSON `json:"-"`
}

// accountDevicePolicyUpdateResponseResultIncludeJSON contains the JSON metadata
// for the struct [AccountDevicePolicyUpdateResponseResultInclude]
type accountDevicePolicyUpdateResponseResultIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyUpdateResponseResultInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyUpdateResponseResultServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                                  `json:"port"`
	JSON accountDevicePolicyUpdateResponseResultServiceModeV2JSON `json:"-"`
}

// accountDevicePolicyUpdateResponseResultServiceModeV2JSON contains the JSON
// metadata for the struct [AccountDevicePolicyUpdateResponseResultServiceModeV2]
type accountDevicePolicyUpdateResponseResultServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyUpdateResponseResultServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       accountDevicePolicyUpdateResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyUpdateResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountDevicePolicyUpdateResponseResultInfo]
type accountDevicePolicyUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyUpdateResponseSuccess bool

const (
	AccountDevicePolicyUpdateResponseSuccessTrue AccountDevicePolicyUpdateResponseSuccess = true
)

type AccountDevicePolicyDeleteResponse struct {
	Errors     []AccountDevicePolicyDeleteResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyDeleteResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyDeleteResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyDeleteResponseSuccess `json:"success"`
	JSON    accountDevicePolicyDeleteResponseJSON    `json:"-"`
}

// accountDevicePolicyDeleteResponseJSON contains the JSON metadata for the struct
// [AccountDevicePolicyDeleteResponse]
type accountDevicePolicyDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDeleteResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountDevicePolicyDeleteResponseErrorJSON `json:"-"`
}

// accountDevicePolicyDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountDevicePolicyDeleteResponseError]
type accountDevicePolicyDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDeleteResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountDevicePolicyDeleteResponseMessageJSON `json:"-"`
}

// accountDevicePolicyDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountDevicePolicyDeleteResponseMessage]
type accountDevicePolicyDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDeleteResponseResult struct {
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
	Enabled bool                                             `json:"enabled"`
	Exclude []AccountDevicePolicyDeleteResponseResultExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                                    `json:"exclude_office_ips"`
	FallbackDomains  []AccountDevicePolicyDeleteResponseResultFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                                  `json:"gateway_unique_id"`
	Include          []AccountDevicePolicyDeleteResponseResultInclude        `json:"include"`
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
	Precedence    float64                                              `json:"precedence"`
	ServiceModeV2 AccountDevicePolicyDeleteResponseResultServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                        `json:"switch_locked"`
	JSON         accountDevicePolicyDeleteResponseResultJSON `json:"-"`
}

// accountDevicePolicyDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountDevicePolicyDeleteResponseResult]
type accountDevicePolicyDeleteResponseResultJSON struct {
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

func (r *AccountDevicePolicyDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDeleteResponseResultExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                             `json:"host"`
	JSON accountDevicePolicyDeleteResponseResultExcludeJSON `json:"-"`
}

// accountDevicePolicyDeleteResponseResultExcludeJSON contains the JSON metadata
// for the struct [AccountDevicePolicyDeleteResponseResultExclude]
type accountDevicePolicyDeleteResponseResultExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDeleteResponseResultExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDeleteResponseResultFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                             `json:"dns_server"`
	JSON      accountDevicePolicyDeleteResponseResultFallbackDomainJSON `json:"-"`
}

// accountDevicePolicyDeleteResponseResultFallbackDomainJSON contains the JSON
// metadata for the struct [AccountDevicePolicyDeleteResponseResultFallbackDomain]
type accountDevicePolicyDeleteResponseResultFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDeleteResponseResultFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDeleteResponseResultInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                             `json:"host"`
	JSON accountDevicePolicyDeleteResponseResultIncludeJSON `json:"-"`
}

// accountDevicePolicyDeleteResponseResultIncludeJSON contains the JSON metadata
// for the struct [AccountDevicePolicyDeleteResponseResultInclude]
type accountDevicePolicyDeleteResponseResultIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDeleteResponseResultInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDeleteResponseResultServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                                  `json:"port"`
	JSON accountDevicePolicyDeleteResponseResultServiceModeV2JSON `json:"-"`
}

// accountDevicePolicyDeleteResponseResultServiceModeV2JSON contains the JSON
// metadata for the struct [AccountDevicePolicyDeleteResponseResultServiceModeV2]
type accountDevicePolicyDeleteResponseResultServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDeleteResponseResultServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       accountDevicePolicyDeleteResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyDeleteResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountDevicePolicyDeleteResponseResultInfo]
type accountDevicePolicyDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyDeleteResponseSuccess bool

const (
	AccountDevicePolicyDeleteResponseSuccessTrue AccountDevicePolicyDeleteResponseSuccess = true
)

type AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponse struct {
	Errors     []AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseMessage  `json:"messages"`
	Result     AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResult     `json:"result"`
	ResultInfo AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseSuccess `json:"success"`
	JSON    accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseJSON    `json:"-"`
}

// accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseJSON contains the JSON
// metadata for the struct
// [AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponse]
type accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseError struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseErrorJSON `json:"-"`
}

// accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseError]
type accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseMessage struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseMessageJSON `json:"-"`
}

// accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseMessage]
type accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResult struct {
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
	Enabled bool                                                                     `json:"enabled"`
	Exclude []AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                                                            `json:"exclude_office_ips"`
	FallbackDomains  []AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                                                          `json:"gateway_unique_id"`
	Include          []AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultInclude        `json:"include"`
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
	Precedence    float64                                                                      `json:"precedence"`
	ServiceModeV2 AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                                                `json:"switch_locked"`
	JSON         accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultJSON `json:"-"`
}

// accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultJSON contains the
// JSON metadata for the struct
// [AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResult]
type accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultJSON struct {
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

func (r *AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                                     `json:"host"`
	JSON accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultExcludeJSON `json:"-"`
}

// accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultExcludeJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultExclude]
type accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                                     `json:"dns_server"`
	JSON      accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultFallbackDomainJSON `json:"-"`
}

// accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultFallbackDomainJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultFallbackDomain]
type accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                                     `json:"host"`
	JSON accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultIncludeJSON `json:"-"`
}

// accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultIncludeJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultInclude]
type accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                                                          `json:"port"`
	JSON accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultServiceModeV2JSON `json:"-"`
}

// accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultServiceModeV2JSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultServiceModeV2]
type accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                 `json:"total_count"`
	JSON       accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultInfoJSON contains
// the JSON metadata for the struct
// [AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultInfo]
type accountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseSuccess bool

const (
	AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseSuccessTrue AccountDevicePolicyDevicesNewDeviceSettingsPolicyResponseSuccess = true
)

type AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponse struct {
	Errors     []AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseMessage  `json:"messages"`
	Result     AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResult     `json:"result"`
	ResultInfo AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseSuccess `json:"success"`
	JSON    accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseJSON    `json:"-"`
}

// accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseJSON contains
// the JSON metadata for the struct
// [AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponse]
type accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseErrorJSON `json:"-"`
}

// accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseError]
type accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseMessageJSON `json:"-"`
}

// accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseMessage]
type accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResult struct {
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
	// Whether the policy will be applied to matching devices.
	Default bool `json:"default"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback bool `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled bool                                                                            `json:"enabled"`
	Exclude []AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                                                                   `json:"exclude_office_ips"`
	FallbackDomains  []AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                                                                 `json:"gateway_unique_id"`
	Include          []AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultInclude        `json:"include"`
	ServiceModeV2    AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultServiceModeV2    `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                                                       `json:"switch_locked"`
	JSON         accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultJSON `json:"-"`
}

// accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResult]
type accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultJSON struct {
	AllowModeSwitch     apijson.Field
	AllowUpdates        apijson.Field
	AllowedToLeave      apijson.Field
	AutoConnect         apijson.Field
	CaptivePortal       apijson.Field
	Default             apijson.Field
	DisableAutoFallback apijson.Field
	Enabled             apijson.Field
	Exclude             apijson.Field
	ExcludeOfficeIPs    apijson.Field
	FallbackDomains     apijson.Field
	GatewayUniqueID     apijson.Field
	Include             apijson.Field
	ServiceModeV2       apijson.Field
	SupportURL          apijson.Field
	SwitchLocked        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                                            `json:"host"`
	JSON accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultExcludeJSON `json:"-"`
}

// accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultExcludeJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultExclude]
type accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                                            `json:"dns_server"`
	JSON      accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultFallbackDomainJSON `json:"-"`
}

// accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultFallbackDomainJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultFallbackDomain]
type accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                                            `json:"host"`
	JSON accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultIncludeJSON `json:"-"`
}

// accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultIncludeJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultInclude]
type accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                                                                 `json:"port"`
	JSON accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultServiceModeV2JSON `json:"-"`
}

// accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultServiceModeV2JSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultServiceModeV2]
type accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                        `json:"total_count"`
	JSON       accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultInfo]
type accountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseSuccess bool

const (
	AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseSuccessTrue AccountDevicePolicyDevicesGetDefaultDeviceSettingsPolicyResponseSuccess = true
)

type AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponse struct {
	Errors     []AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseMessage  `json:"messages"`
	Result     []AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResult   `json:"result"`
	ResultInfo AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseSuccess `json:"success"`
	JSON    accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseJSON    `json:"-"`
}

// accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseJSON contains the
// JSON metadata for the struct
// [AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponse]
type accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseErrorJSON `json:"-"`
}

// accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseError]
type accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseMessageJSON `json:"-"`
}

// accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseMessage]
type accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResult struct {
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
	Enabled bool                                                                        `json:"enabled"`
	Exclude []AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                                                               `json:"exclude_office_ips"`
	FallbackDomains  []AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                                                             `json:"gateway_unique_id"`
	Include          []AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultInclude        `json:"include"`
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
	Precedence    float64                                                                         `json:"precedence"`
	ServiceModeV2 AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                                                   `json:"switch_locked"`
	JSON         accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultJSON `json:"-"`
}

// accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultJSON contains
// the JSON metadata for the struct
// [AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResult]
type accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultJSON struct {
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

func (r *AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                                        `json:"host"`
	JSON accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultExcludeJSON `json:"-"`
}

// accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultExcludeJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultExclude]
type accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                                        `json:"dns_server"`
	JSON      accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultFallbackDomainJSON `json:"-"`
}

// accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultFallbackDomainJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultFallbackDomain]
type accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                                        `json:"host"`
	JSON accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultIncludeJSON `json:"-"`
}

// accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultIncludeJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultInclude]
type accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                                                             `json:"port"`
	JSON accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultServiceModeV2JSON `json:"-"`
}

// accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultServiceModeV2JSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultServiceModeV2]
type accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                    `json:"total_count"`
	JSON       accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultInfo]
type accountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseSuccess bool

const (
	AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseSuccessTrue AccountDevicePolicyDevicesListDeviceSettingsPoliciesResponseSuccess = true
)

type AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponse struct {
	Errors     []AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseError    `json:"errors"`
	Messages   []AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseMessage  `json:"messages"`
	Result     AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResult     `json:"result"`
	ResultInfo AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseSuccess `json:"success"`
	JSON    accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseJSON    `json:"-"`
}

// accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseJSON contains
// the JSON metadata for the struct
// [AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponse]
type accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseError struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseErrorJSON `json:"-"`
}

// accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseError]
type accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseMessage struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseMessageJSON `json:"-"`
}

// accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseMessage]
type accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResult struct {
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
	// Whether the policy will be applied to matching devices.
	Default bool `json:"default"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback bool `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled bool                                                                               `json:"enabled"`
	Exclude []AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                                                                      `json:"exclude_office_ips"`
	FallbackDomains  []AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                                                                    `json:"gateway_unique_id"`
	Include          []AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultInclude        `json:"include"`
	ServiceModeV2    AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultServiceModeV2    `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                                                          `json:"switch_locked"`
	JSON         accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultJSON `json:"-"`
}

// accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResult]
type accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultJSON struct {
	AllowModeSwitch     apijson.Field
	AllowUpdates        apijson.Field
	AllowedToLeave      apijson.Field
	AutoConnect         apijson.Field
	CaptivePortal       apijson.Field
	Default             apijson.Field
	DisableAutoFallback apijson.Field
	Enabled             apijson.Field
	Exclude             apijson.Field
	ExcludeOfficeIPs    apijson.Field
	FallbackDomains     apijson.Field
	GatewayUniqueID     apijson.Field
	Include             apijson.Field
	ServiceModeV2       apijson.Field
	SupportURL          apijson.Field
	SwitchLocked        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                                                                               `json:"host"`
	JSON accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultExcludeJSON `json:"-"`
}

// accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultExcludeJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultExclude]
type accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                                                               `json:"dns_server"`
	JSON      accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultFallbackDomainJSON `json:"-"`
}

// accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultFallbackDomainJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultFallbackDomain]
type accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                                                               `json:"host"`
	JSON accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultIncludeJSON `json:"-"`
}

// accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultIncludeJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultInclude]
type accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                                                                    `json:"port"`
	JSON accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultServiceModeV2JSON `json:"-"`
}

// accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultServiceModeV2JSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultServiceModeV2]
type accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                           `json:"total_count"`
	JSON       accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultInfoJSON `json:"-"`
}

// accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultInfo]
type accountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseSuccess bool

const (
	AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseSuccessTrue AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyResponseSuccess = true
)

type AccountDevicePolicyUpdateParams struct {
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
	ServiceModeV2 param.Field[AccountDevicePolicyUpdateParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
}

func (r AccountDevicePolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyUpdateParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r AccountDevicePolicyUpdateParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyDevicesNewDeviceSettingsPolicyParams struct {
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
	LanAllowSubnetSize param.Field[float64]                                                              `json:"lan_allow_subnet_size"`
	ServiceModeV2      param.Field[AccountDevicePolicyDevicesNewDeviceSettingsPolicyParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
}

func (r AccountDevicePolicyDevicesNewDeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyDevicesNewDeviceSettingsPolicyParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r AccountDevicePolicyDevicesNewDeviceSettingsPolicyParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParams struct {
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
	ExcludeOfficeIPs param.Field[bool]                                                                           `json:"exclude_office_ips"`
	ServiceModeV2    param.Field[AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
}

func (r AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
