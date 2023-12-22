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

// Get the device settings policy by ID.
func (r *AccountDevicePolicyService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a device settings policy.
func (r *AccountDevicePolicyService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountDevicePolicyUpdateParams, opts ...option.RequestOption) (res *DeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete a device settings policy. Returns the remaining policies for the account.
func (r *AccountDevicePolicyService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DeviceSettingsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a device settings policy to be applied to certain devices matching the
// criteria.
func (r *AccountDevicePolicyService) DevicesNewDeviceSettingsPolicy(ctx context.Context, identifier interface{}, body AccountDevicePolicyDevicesNewDeviceSettingsPolicyParams, opts ...option.RequestOption) (res *DeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the default device settings policy for an account.
func (r *AccountDevicePolicyService) DevicesGetDefaultDeviceSettingsPolicy(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *DefaultDeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists the device settings policies for an account.
func (r *AccountDevicePolicyService) DevicesListDeviceSettingsPolicies(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *DeviceSettingsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policies", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the default device settings policy for an account.
func (r *AccountDevicePolicyService) DevicesUpdateDefaultDeviceSettingsPolicy(ctx context.Context, identifier interface{}, body AccountDevicePolicyDevicesUpdateDefaultDeviceSettingsPolicyParams, opts ...option.RequestOption) (res *DefaultDeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type DefaultDeviceSettingsResponse struct {
	Errors     []DefaultDeviceSettingsResponseError    `json:"errors"`
	Messages   []DefaultDeviceSettingsResponseMessage  `json:"messages"`
	Result     DefaultDeviceSettingsResponseResult     `json:"result"`
	ResultInfo DefaultDeviceSettingsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DefaultDeviceSettingsResponseSuccess `json:"success"`
	JSON    defaultDeviceSettingsResponseJSON    `json:"-"`
}

// defaultDeviceSettingsResponseJSON contains the JSON metadata for the struct
// [DefaultDeviceSettingsResponse]
type defaultDeviceSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultDeviceSettingsResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    defaultDeviceSettingsResponseErrorJSON `json:"-"`
}

// defaultDeviceSettingsResponseErrorJSON contains the JSON metadata for the struct
// [DefaultDeviceSettingsResponseError]
type defaultDeviceSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultDeviceSettingsResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    defaultDeviceSettingsResponseMessageJSON `json:"-"`
}

// defaultDeviceSettingsResponseMessageJSON contains the JSON metadata for the
// struct [DefaultDeviceSettingsResponseMessage]
type defaultDeviceSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultDeviceSettingsResponseResult struct {
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
	// If the dns_server field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers, unless this
	// policy option is set.
	DisableAutoFallback bool `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled bool                                         `json:"enabled"`
	Exclude []DefaultDeviceSettingsResponseResultExclude `json:"exclude"`
	// Whether to add Microsoft IPs to split tunnel exclusions.
	ExcludeOfficeIPs bool                                                `json:"exclude_office_ips"`
	FallbackDomains  []DefaultDeviceSettingsResponseResultFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                              `json:"gateway_unique_id"`
	Include          []DefaultDeviceSettingsResponseResultInclude        `json:"include"`
	ServiceModeV2    DefaultDeviceSettingsResponseResultServiceModeV2    `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                    `json:"switch_locked"`
	JSON         defaultDeviceSettingsResponseResultJSON `json:"-"`
}

// defaultDeviceSettingsResponseResultJSON contains the JSON metadata for the
// struct [DefaultDeviceSettingsResponseResult]
type defaultDeviceSettingsResponseResultJSON struct {
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

func (r *DefaultDeviceSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultDeviceSettingsResponseResultExclude struct {
	// The address in CIDR format to exclude from the tunnel. If address is present,
	// host must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If host is present, address must not
	// be present.
	Host string                                         `json:"host"`
	JSON defaultDeviceSettingsResponseResultExcludeJSON `json:"-"`
}

// defaultDeviceSettingsResponseResultExcludeJSON contains the JSON metadata for
// the struct [DefaultDeviceSettingsResponseResultExclude]
type defaultDeviceSettingsResponseResultExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponseResultExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultDeviceSettingsResponseResultFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                         `json:"dns_server"`
	JSON      defaultDeviceSettingsResponseResultFallbackDomainJSON `json:"-"`
}

// defaultDeviceSettingsResponseResultFallbackDomainJSON contains the JSON metadata
// for the struct [DefaultDeviceSettingsResponseResultFallbackDomain]
type defaultDeviceSettingsResponseResultFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponseResultFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultDeviceSettingsResponseResultInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                         `json:"host"`
	JSON defaultDeviceSettingsResponseResultIncludeJSON `json:"-"`
}

// defaultDeviceSettingsResponseResultIncludeJSON contains the JSON metadata for
// the struct [DefaultDeviceSettingsResponseResultInclude]
type defaultDeviceSettingsResponseResultIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponseResultInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultDeviceSettingsResponseResultServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                              `json:"port"`
	JSON defaultDeviceSettingsResponseResultServiceModeV2JSON `json:"-"`
}

// defaultDeviceSettingsResponseResultServiceModeV2JSON contains the JSON metadata
// for the struct [DefaultDeviceSettingsResponseResultServiceModeV2]
type defaultDeviceSettingsResponseResultServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponseResultServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultDeviceSettingsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       defaultDeviceSettingsResponseResultInfoJSON `json:"-"`
}

// defaultDeviceSettingsResponseResultInfoJSON contains the JSON metadata for the
// struct [DefaultDeviceSettingsResponseResultInfo]
type defaultDeviceSettingsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DefaultDeviceSettingsResponseSuccess bool

const (
	DefaultDeviceSettingsResponseSuccessTrue DefaultDeviceSettingsResponseSuccess = true
)

type DeviceSettingsResponse struct {
	Errors     []DeviceSettingsResponseError    `json:"errors"`
	Messages   []DeviceSettingsResponseMessage  `json:"messages"`
	Result     DeviceSettingsResponseResult     `json:"result"`
	ResultInfo DeviceSettingsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DeviceSettingsResponseSuccess `json:"success"`
	JSON    deviceSettingsResponseJSON    `json:"-"`
}

// deviceSettingsResponseJSON contains the JSON metadata for the struct
// [DeviceSettingsResponse]
type deviceSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    deviceSettingsResponseErrorJSON `json:"-"`
}

// deviceSettingsResponseErrorJSON contains the JSON metadata for the struct
// [DeviceSettingsResponseError]
type deviceSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    deviceSettingsResponseMessageJSON `json:"-"`
}

// deviceSettingsResponseMessageJSON contains the JSON metadata for the struct
// [DeviceSettingsResponseMessage]
type deviceSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseResult struct {
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
	// If the dns_server field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers, unless this
	// policy option is set.
	DisableAutoFallback bool `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled bool                                  `json:"enabled"`
	Exclude []DeviceSettingsResponseResultExclude `json:"exclude"`
	// Whether to add Microsoft IPs to split tunnel exclusions.
	ExcludeOfficeIPs bool                                         `json:"exclude_office_ips"`
	FallbackDomains  []DeviceSettingsResponseResultFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                       `json:"gateway_unique_id"`
	Include          []DeviceSettingsResponseResultInclude        `json:"include"`
	// The wirefilter expression to match devices.
	Match string `json:"match"`
	// The name of the device settings policy.
	Name string `json:"name"`
	// Device ID.
	PolicyID string `json:"policy_id"`
	// The precedence of the policy. Lower values indicate higher precedence. Policies
	// will be evaluated in ascending order of this field.
	Precedence    float64                                   `json:"precedence"`
	ServiceModeV2 DeviceSettingsResponseResultServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                             `json:"switch_locked"`
	JSON         deviceSettingsResponseResultJSON `json:"-"`
}

// deviceSettingsResponseResultJSON contains the JSON metadata for the struct
// [DeviceSettingsResponseResult]
type deviceSettingsResponseResultJSON struct {
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

func (r *DeviceSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseResultExclude struct {
	// The address in CIDR format to exclude from the tunnel. If address is present,
	// host must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If host is present, address must not
	// be present.
	Host string                                  `json:"host"`
	JSON deviceSettingsResponseResultExcludeJSON `json:"-"`
}

// deviceSettingsResponseResultExcludeJSON contains the JSON metadata for the
// struct [DeviceSettingsResponseResultExclude]
type deviceSettingsResponseResultExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseResultExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseResultFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                  `json:"dns_server"`
	JSON      deviceSettingsResponseResultFallbackDomainJSON `json:"-"`
}

// deviceSettingsResponseResultFallbackDomainJSON contains the JSON metadata for
// the struct [DeviceSettingsResponseResultFallbackDomain]
type deviceSettingsResponseResultFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseResultFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseResultInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                  `json:"host"`
	JSON deviceSettingsResponseResultIncludeJSON `json:"-"`
}

// deviceSettingsResponseResultIncludeJSON contains the JSON metadata for the
// struct [DeviceSettingsResponseResultInclude]
type deviceSettingsResponseResultIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseResultInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseResultServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                       `json:"port"`
	JSON deviceSettingsResponseResultServiceModeV2JSON `json:"-"`
}

// deviceSettingsResponseResultServiceModeV2JSON contains the JSON metadata for the
// struct [DeviceSettingsResponseResultServiceModeV2]
type deviceSettingsResponseResultServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseResultServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       deviceSettingsResponseResultInfoJSON `json:"-"`
}

// deviceSettingsResponseResultInfoJSON contains the JSON metadata for the struct
// [DeviceSettingsResponseResultInfo]
type deviceSettingsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DeviceSettingsResponseSuccess bool

const (
	DeviceSettingsResponseSuccessTrue DeviceSettingsResponseSuccess = true
)

type DeviceSettingsResponseCollection struct {
	Errors     []DeviceSettingsResponseCollectionError    `json:"errors"`
	Messages   []DeviceSettingsResponseCollectionMessage  `json:"messages"`
	Result     []DeviceSettingsResponseCollectionResult   `json:"result"`
	ResultInfo DeviceSettingsResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DeviceSettingsResponseCollectionSuccess `json:"success"`
	JSON    deviceSettingsResponseCollectionJSON    `json:"-"`
}

// deviceSettingsResponseCollectionJSON contains the JSON metadata for the struct
// [DeviceSettingsResponseCollection]
type deviceSettingsResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseCollectionError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    deviceSettingsResponseCollectionErrorJSON `json:"-"`
}

// deviceSettingsResponseCollectionErrorJSON contains the JSON metadata for the
// struct [DeviceSettingsResponseCollectionError]
type deviceSettingsResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseCollectionMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    deviceSettingsResponseCollectionMessageJSON `json:"-"`
}

// deviceSettingsResponseCollectionMessageJSON contains the JSON metadata for the
// struct [DeviceSettingsResponseCollectionMessage]
type deviceSettingsResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseCollectionResult struct {
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
	// If the dns_server field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers, unless this
	// policy option is set.
	DisableAutoFallback bool `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled bool                                            `json:"enabled"`
	Exclude []DeviceSettingsResponseCollectionResultExclude `json:"exclude"`
	// Whether to add Microsoft IPs to split tunnel exclusions.
	ExcludeOfficeIPs bool                                                   `json:"exclude_office_ips"`
	FallbackDomains  []DeviceSettingsResponseCollectionResultFallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string                                                 `json:"gateway_unique_id"`
	Include          []DeviceSettingsResponseCollectionResultInclude        `json:"include"`
	// The wirefilter expression to match devices.
	Match string `json:"match"`
	// The name of the device settings policy.
	Name string `json:"name"`
	// Device ID.
	PolicyID string `json:"policy_id"`
	// The precedence of the policy. Lower values indicate higher precedence. Policies
	// will be evaluated in ascending order of this field.
	Precedence    float64                                             `json:"precedence"`
	ServiceModeV2 DeviceSettingsResponseCollectionResultServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                                       `json:"switch_locked"`
	JSON         deviceSettingsResponseCollectionResultJSON `json:"-"`
}

// deviceSettingsResponseCollectionResultJSON contains the JSON metadata for the
// struct [DeviceSettingsResponseCollectionResult]
type deviceSettingsResponseCollectionResultJSON struct {
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

func (r *DeviceSettingsResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseCollectionResultExclude struct {
	// The address in CIDR format to exclude from the tunnel. If address is present,
	// host must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If host is present, address must not
	// be present.
	Host string                                            `json:"host"`
	JSON deviceSettingsResponseCollectionResultExcludeJSON `json:"-"`
}

// deviceSettingsResponseCollectionResultExcludeJSON contains the JSON metadata for
// the struct [DeviceSettingsResponseCollectionResultExclude]
type deviceSettingsResponseCollectionResultExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollectionResultExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseCollectionResultFallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                            `json:"dns_server"`
	JSON      deviceSettingsResponseCollectionResultFallbackDomainJSON `json:"-"`
}

// deviceSettingsResponseCollectionResultFallbackDomainJSON contains the JSON
// metadata for the struct [DeviceSettingsResponseCollectionResultFallbackDomain]
type deviceSettingsResponseCollectionResultFallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollectionResultFallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseCollectionResultInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                            `json:"host"`
	JSON deviceSettingsResponseCollectionResultIncludeJSON `json:"-"`
}

// deviceSettingsResponseCollectionResultIncludeJSON contains the JSON metadata for
// the struct [DeviceSettingsResponseCollectionResultInclude]
type deviceSettingsResponseCollectionResultIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollectionResultInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseCollectionResultServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                                 `json:"port"`
	JSON deviceSettingsResponseCollectionResultServiceModeV2JSON `json:"-"`
}

// deviceSettingsResponseCollectionResultServiceModeV2JSON contains the JSON
// metadata for the struct [DeviceSettingsResponseCollectionResultServiceModeV2]
type deviceSettingsResponseCollectionResultServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollectionResultServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSettingsResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       deviceSettingsResponseCollectionResultInfoJSON `json:"-"`
}

// deviceSettingsResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [DeviceSettingsResponseCollectionResultInfo]
type deviceSettingsResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DeviceSettingsResponseCollectionSuccess bool

const (
	DeviceSettingsResponseCollectionSuccessTrue DeviceSettingsResponseCollectionSuccess = true
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
	// If the dns_server field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers, unless this
	// policy option is set.
	DisableAutoFallback param.Field[bool] `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled param.Field[bool] `json:"enabled"`
	// Whether to add Microsoft IPs to split tunnel exclusions.
	ExcludeOfficeIPs param.Field[bool] `json:"exclude_office_ips"`
	// The wirefilter expression to match devices.
	Match param.Field[string] `json:"match"`
	// The name of the device settings policy.
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
	// The name of the device settings policy.
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
	// If the dns_server field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers, unless this
	// policy option is set.
	DisableAutoFallback param.Field[bool] `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled param.Field[bool] `json:"enabled"`
	// Whether to add Microsoft IPs to split tunnel exclusions.
	ExcludeOfficeIPs param.Field[bool]                                                                 `json:"exclude_office_ips"`
	ServiceModeV2    param.Field[AccountDevicePolicyDevicesNewDeviceSettingsPolicyParamsServiceModeV2] `json:"service_mode_v2"`
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
	// If the dns_server field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers, unless this
	// policy option is set.
	DisableAutoFallback param.Field[bool] `json:"disable_auto_fallback"`
	// Whether to add Microsoft IPs to split tunnel exclusions.
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
