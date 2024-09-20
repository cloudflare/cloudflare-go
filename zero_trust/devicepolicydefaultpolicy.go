// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DevicePolicyDefaultPolicyService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyDefaultPolicyService] method instead.
type DevicePolicyDefaultPolicyService struct {
	Options []option.RequestOption
}

// NewDevicePolicyDefaultPolicyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyDefaultPolicyService(opts ...option.RequestOption) (r *DevicePolicyDefaultPolicyService) {
	r = &DevicePolicyDefaultPolicyService{}
	r.Options = opts
	return
}

// Fetches the default device settings profile for an account.
func (r *DevicePolicyDefaultPolicyService) Get(ctx context.Context, query DevicePolicyDefaultPolicyGetParams, opts ...option.RequestOption) (res *DevicePolicyDefaultPolicyGetResponse, err error) {
	var env DevicePolicyDefaultPolicyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyDefaultPolicyGetResponse struct {
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
	Enabled bool                 `json:"enabled"`
	Exclude []SplitTunnelExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool                                              `json:"exclude_office_ips"`
	FallbackDomains  []FallbackDomain                                  `json:"fallback_domains"`
	GatewayUniqueID  string                                            `json:"gateway_unique_id"`
	Include          []SplitTunnelInclude                              `json:"include"`
	ServiceModeV2    DevicePolicyDefaultPolicyGetResponseServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool `json:"switch_locked"`
	// Determines which tunnel protocol to use.
	TunnelProtocol string                                   `json:"tunnel_protocol"`
	JSON           devicePolicyDefaultPolicyGetResponseJSON `json:"-"`
}

// devicePolicyDefaultPolicyGetResponseJSON contains the JSON metadata for the
// struct [DevicePolicyDefaultPolicyGetResponse]
type devicePolicyDefaultPolicyGetResponseJSON struct {
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
	TunnelProtocol      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DevicePolicyDefaultPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyDefaultPolicyGetResponseServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                               `json:"port"`
	JSON devicePolicyDefaultPolicyGetResponseServiceModeV2JSON `json:"-"`
}

// devicePolicyDefaultPolicyGetResponseServiceModeV2JSON contains the JSON metadata
// for the struct [DevicePolicyDefaultPolicyGetResponseServiceModeV2]
type devicePolicyDefaultPolicyGetResponseServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultPolicyGetResponseServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultPolicyGetResponseServiceModeV2JSON) RawJSON() string {
	return r.raw
}

type DevicePolicyDefaultPolicyGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyDefaultPolicyGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                `json:"errors,required"`
	Messages []shared.ResponseInfo                `json:"messages,required"`
	Result   DevicePolicyDefaultPolicyGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePolicyDefaultPolicyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePolicyDefaultPolicyGetResponseEnvelopeJSON    `json:"-"`
}

// devicePolicyDefaultPolicyGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePolicyDefaultPolicyGetResponseEnvelope]
type devicePolicyDefaultPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultPolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyDefaultPolicyGetResponseEnvelopeSuccess bool

const (
	DevicePolicyDefaultPolicyGetResponseEnvelopeSuccessTrue DevicePolicyDefaultPolicyGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyDefaultPolicyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyDefaultPolicyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
