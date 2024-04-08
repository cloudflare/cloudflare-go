// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// FirewallAccessRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallAccessRuleService] method
// instead.
type FirewallAccessRuleService struct {
	Options []option.RequestOption
}

// NewFirewallAccessRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallAccessRuleService(opts ...option.RequestOption) (r *FirewallAccessRuleService) {
	r = &FirewallAccessRuleService{}
	r.Options = opts
	return
}

// Creates a new IP Access rule for all zones owned by the current user.
//
// Note: To create an IP Access rule that applies to a specific zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *FirewallAccessRuleService) New(ctx context.Context, body FirewallAccessRuleNewParams, opts ...option.RequestOption) (res *FirewallAccessRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAccessRuleNewResponseEnvelope
	path := "user/firewall/access_rules/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches IP Access rules of the user. You can filter the results using several
// optional parameters.
func (r *FirewallAccessRuleService) List(ctx context.Context, query FirewallAccessRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[FirewallAccessRuleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/firewall/access_rules/rules"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches IP Access rules of the user. You can filter the results using several
// optional parameters.
func (r *FirewallAccessRuleService) ListAutoPaging(ctx context.Context, query FirewallAccessRuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[FirewallAccessRuleListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Deletes an IP Access rule at the user level.
//
// Note: Deleting a user-level rule will affect all zones owned by the user.
func (r *FirewallAccessRuleService) Delete(ctx context.Context, identifier string, body FirewallAccessRuleDeleteParams, opts ...option.RequestOption) (res *FirewallAccessRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAccessRuleDeleteResponseEnvelope
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an IP Access rule defined at the user level. You can only update the
// rule action (`mode` parameter) and notes.
func (r *FirewallAccessRuleService) Edit(ctx context.Context, identifier string, body FirewallAccessRuleEditParams, opts ...option.RequestOption) (res *FirewallAccessRuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAccessRuleEditResponseEnvelope
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallAccessRuleNewResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []FirewallAccessRuleNewResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration FirewallAccessRuleNewResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode FirewallAccessRuleNewResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                            `json:"notes"`
	JSON  firewallAccessRuleNewResponseJSON `json:"-"`
}

// firewallAccessRuleNewResponseJSON contains the JSON metadata for the struct
// [FirewallAccessRuleNewResponse]
type firewallAccessRuleNewResponseJSON struct {
	ID            apijson.Field
	AllowedModes  apijson.Field
	Configuration apijson.Field
	Mode          apijson.Field
	CreatedOn     apijson.Field
	ModifiedOn    apijson.Field
	Notes         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type FirewallAccessRuleNewResponseAllowedMode string

const (
	FirewallAccessRuleNewResponseAllowedModeBlock            FirewallAccessRuleNewResponseAllowedMode = "block"
	FirewallAccessRuleNewResponseAllowedModeChallenge        FirewallAccessRuleNewResponseAllowedMode = "challenge"
	FirewallAccessRuleNewResponseAllowedModeWhitelist        FirewallAccessRuleNewResponseAllowedMode = "whitelist"
	FirewallAccessRuleNewResponseAllowedModeJsChallenge      FirewallAccessRuleNewResponseAllowedMode = "js_challenge"
	FirewallAccessRuleNewResponseAllowedModeManagedChallenge FirewallAccessRuleNewResponseAllowedMode = "managed_challenge"
)

func (r FirewallAccessRuleNewResponseAllowedMode) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewResponseAllowedModeBlock, FirewallAccessRuleNewResponseAllowedModeChallenge, FirewallAccessRuleNewResponseAllowedModeWhitelist, FirewallAccessRuleNewResponseAllowedModeJsChallenge, FirewallAccessRuleNewResponseAllowedModeManagedChallenge:
		return true
	}
	return false
}

// The rule configuration.
type FirewallAccessRuleNewResponseConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target FirewallAccessRuleNewResponseConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                         `json:"value"`
	JSON  firewallAccessRuleNewResponseConfigurationJSON `json:"-"`
	union FirewallAccessRuleNewResponseConfigurationUnion
}

// firewallAccessRuleNewResponseConfigurationJSON contains the JSON metadata for
// the struct [FirewallAccessRuleNewResponseConfiguration]
type firewallAccessRuleNewResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r firewallAccessRuleNewResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r *FirewallAccessRuleNewResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r FirewallAccessRuleNewResponseConfiguration) AsUnion() FirewallAccessRuleNewResponseConfigurationUnion {
	return r.union
}

// The rule configuration.
//
// Union satisfied by
// [user.FirewallAccessRuleNewResponseConfigurationFirewallIPConfiguration],
// [user.FirewallAccessRuleNewResponseConfigurationFirewallIPV6Configuration],
// [user.FirewallAccessRuleNewResponseConfigurationFirewallCIDRConfiguration],
// [user.FirewallAccessRuleNewResponseConfigurationFirewallASNConfiguration] or
// [user.FirewallAccessRuleNewResponseConfigurationFirewallCountryConfiguration].
type FirewallAccessRuleNewResponseConfigurationUnion interface {
	implementsUserFirewallAccessRuleNewResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallAccessRuleNewResponseConfigurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleNewResponseConfigurationFirewallIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleNewResponseConfigurationFirewallIPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleNewResponseConfigurationFirewallCIDRConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleNewResponseConfigurationFirewallASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleNewResponseConfigurationFirewallCountryConfiguration{}),
		},
	)
}

type FirewallAccessRuleNewResponseConfigurationFirewallIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target FirewallAccessRuleNewResponseConfigurationFirewallIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                `json:"value"`
	JSON  firewallAccessRuleNewResponseConfigurationFirewallIPConfigurationJSON `json:"-"`
}

// firewallAccessRuleNewResponseConfigurationFirewallIPConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallAccessRuleNewResponseConfigurationFirewallIPConfiguration]
type firewallAccessRuleNewResponseConfigurationFirewallIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseConfigurationFirewallIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseConfigurationFirewallIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleNewResponseConfigurationFirewallIPConfiguration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleNewResponseConfigurationFirewallIPConfigurationTarget string

const (
	FirewallAccessRuleNewResponseConfigurationFirewallIPConfigurationTargetIP FirewallAccessRuleNewResponseConfigurationFirewallIPConfigurationTarget = "ip"
)

func (r FirewallAccessRuleNewResponseConfigurationFirewallIPConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewResponseConfigurationFirewallIPConfigurationTargetIP:
		return true
	}
	return false
}

type FirewallAccessRuleNewResponseConfigurationFirewallIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target FirewallAccessRuleNewResponseConfigurationFirewallIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                  `json:"value"`
	JSON  firewallAccessRuleNewResponseConfigurationFirewallIPV6ConfigurationJSON `json:"-"`
}

// firewallAccessRuleNewResponseConfigurationFirewallIPV6ConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallAccessRuleNewResponseConfigurationFirewallIPV6Configuration]
type firewallAccessRuleNewResponseConfigurationFirewallIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseConfigurationFirewallIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseConfigurationFirewallIPV6ConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleNewResponseConfigurationFirewallIPV6Configuration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type FirewallAccessRuleNewResponseConfigurationFirewallIPV6ConfigurationTarget string

const (
	FirewallAccessRuleNewResponseConfigurationFirewallIPV6ConfigurationTargetIp6 FirewallAccessRuleNewResponseConfigurationFirewallIPV6ConfigurationTarget = "ip6"
)

func (r FirewallAccessRuleNewResponseConfigurationFirewallIPV6ConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewResponseConfigurationFirewallIPV6ConfigurationTargetIp6:
		return true
	}
	return false
}

type FirewallAccessRuleNewResponseConfigurationFirewallCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target FirewallAccessRuleNewResponseConfigurationFirewallCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                  `json:"value"`
	JSON  firewallAccessRuleNewResponseConfigurationFirewallCIDRConfigurationJSON `json:"-"`
}

// firewallAccessRuleNewResponseConfigurationFirewallCIDRConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallAccessRuleNewResponseConfigurationFirewallCIDRConfiguration]
type firewallAccessRuleNewResponseConfigurationFirewallCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseConfigurationFirewallCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseConfigurationFirewallCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleNewResponseConfigurationFirewallCIDRConfiguration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type FirewallAccessRuleNewResponseConfigurationFirewallCIDRConfigurationTarget string

const (
	FirewallAccessRuleNewResponseConfigurationFirewallCIDRConfigurationTargetIPRange FirewallAccessRuleNewResponseConfigurationFirewallCIDRConfigurationTarget = "ip_range"
)

func (r FirewallAccessRuleNewResponseConfigurationFirewallCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewResponseConfigurationFirewallCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

type FirewallAccessRuleNewResponseConfigurationFirewallASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target FirewallAccessRuleNewResponseConfigurationFirewallASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                 `json:"value"`
	JSON  firewallAccessRuleNewResponseConfigurationFirewallASNConfigurationJSON `json:"-"`
}

// firewallAccessRuleNewResponseConfigurationFirewallASNConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallAccessRuleNewResponseConfigurationFirewallASNConfiguration]
type firewallAccessRuleNewResponseConfigurationFirewallASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseConfigurationFirewallASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseConfigurationFirewallASNConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleNewResponseConfigurationFirewallASNConfiguration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type FirewallAccessRuleNewResponseConfigurationFirewallASNConfigurationTarget string

const (
	FirewallAccessRuleNewResponseConfigurationFirewallASNConfigurationTargetASN FirewallAccessRuleNewResponseConfigurationFirewallASNConfigurationTarget = "asn"
)

func (r FirewallAccessRuleNewResponseConfigurationFirewallASNConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewResponseConfigurationFirewallASNConfigurationTargetASN:
		return true
	}
	return false
}

type FirewallAccessRuleNewResponseConfigurationFirewallCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target FirewallAccessRuleNewResponseConfigurationFirewallCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                     `json:"value"`
	JSON  firewallAccessRuleNewResponseConfigurationFirewallCountryConfigurationJSON `json:"-"`
}

// firewallAccessRuleNewResponseConfigurationFirewallCountryConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleNewResponseConfigurationFirewallCountryConfiguration]
type firewallAccessRuleNewResponseConfigurationFirewallCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseConfigurationFirewallCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseConfigurationFirewallCountryConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleNewResponseConfigurationFirewallCountryConfiguration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type FirewallAccessRuleNewResponseConfigurationFirewallCountryConfigurationTarget string

const (
	FirewallAccessRuleNewResponseConfigurationFirewallCountryConfigurationTargetCountry FirewallAccessRuleNewResponseConfigurationFirewallCountryConfigurationTarget = "country"
)

func (r FirewallAccessRuleNewResponseConfigurationFirewallCountryConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewResponseConfigurationFirewallCountryConfigurationTargetCountry:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleNewResponseConfigurationTarget string

const (
	FirewallAccessRuleNewResponseConfigurationTargetIP      FirewallAccessRuleNewResponseConfigurationTarget = "ip"
	FirewallAccessRuleNewResponseConfigurationTargetIp6     FirewallAccessRuleNewResponseConfigurationTarget = "ip6"
	FirewallAccessRuleNewResponseConfigurationTargetIPRange FirewallAccessRuleNewResponseConfigurationTarget = "ip_range"
	FirewallAccessRuleNewResponseConfigurationTargetASN     FirewallAccessRuleNewResponseConfigurationTarget = "asn"
	FirewallAccessRuleNewResponseConfigurationTargetCountry FirewallAccessRuleNewResponseConfigurationTarget = "country"
)

func (r FirewallAccessRuleNewResponseConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewResponseConfigurationTargetIP, FirewallAccessRuleNewResponseConfigurationTargetIp6, FirewallAccessRuleNewResponseConfigurationTargetIPRange, FirewallAccessRuleNewResponseConfigurationTargetASN, FirewallAccessRuleNewResponseConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type FirewallAccessRuleNewResponseMode string

const (
	FirewallAccessRuleNewResponseModeBlock            FirewallAccessRuleNewResponseMode = "block"
	FirewallAccessRuleNewResponseModeChallenge        FirewallAccessRuleNewResponseMode = "challenge"
	FirewallAccessRuleNewResponseModeWhitelist        FirewallAccessRuleNewResponseMode = "whitelist"
	FirewallAccessRuleNewResponseModeJsChallenge      FirewallAccessRuleNewResponseMode = "js_challenge"
	FirewallAccessRuleNewResponseModeManagedChallenge FirewallAccessRuleNewResponseMode = "managed_challenge"
)

func (r FirewallAccessRuleNewResponseMode) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewResponseModeBlock, FirewallAccessRuleNewResponseModeChallenge, FirewallAccessRuleNewResponseModeWhitelist, FirewallAccessRuleNewResponseModeJsChallenge, FirewallAccessRuleNewResponseModeManagedChallenge:
		return true
	}
	return false
}

type FirewallAccessRuleListResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []FirewallAccessRuleListResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration FirewallAccessRuleListResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode FirewallAccessRuleListResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                             `json:"notes"`
	JSON  firewallAccessRuleListResponseJSON `json:"-"`
}

// firewallAccessRuleListResponseJSON contains the JSON metadata for the struct
// [FirewallAccessRuleListResponse]
type firewallAccessRuleListResponseJSON struct {
	ID            apijson.Field
	AllowedModes  apijson.Field
	Configuration apijson.Field
	Mode          apijson.Field
	CreatedOn     apijson.Field
	ModifiedOn    apijson.Field
	Notes         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallAccessRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type FirewallAccessRuleListResponseAllowedMode string

const (
	FirewallAccessRuleListResponseAllowedModeBlock            FirewallAccessRuleListResponseAllowedMode = "block"
	FirewallAccessRuleListResponseAllowedModeChallenge        FirewallAccessRuleListResponseAllowedMode = "challenge"
	FirewallAccessRuleListResponseAllowedModeWhitelist        FirewallAccessRuleListResponseAllowedMode = "whitelist"
	FirewallAccessRuleListResponseAllowedModeJsChallenge      FirewallAccessRuleListResponseAllowedMode = "js_challenge"
	FirewallAccessRuleListResponseAllowedModeManagedChallenge FirewallAccessRuleListResponseAllowedMode = "managed_challenge"
)

func (r FirewallAccessRuleListResponseAllowedMode) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListResponseAllowedModeBlock, FirewallAccessRuleListResponseAllowedModeChallenge, FirewallAccessRuleListResponseAllowedModeWhitelist, FirewallAccessRuleListResponseAllowedModeJsChallenge, FirewallAccessRuleListResponseAllowedModeManagedChallenge:
		return true
	}
	return false
}

// The rule configuration.
type FirewallAccessRuleListResponseConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target FirewallAccessRuleListResponseConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                          `json:"value"`
	JSON  firewallAccessRuleListResponseConfigurationJSON `json:"-"`
	union FirewallAccessRuleListResponseConfigurationUnion
}

// firewallAccessRuleListResponseConfigurationJSON contains the JSON metadata for
// the struct [FirewallAccessRuleListResponseConfiguration]
type firewallAccessRuleListResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r firewallAccessRuleListResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r *FirewallAccessRuleListResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r FirewallAccessRuleListResponseConfiguration) AsUnion() FirewallAccessRuleListResponseConfigurationUnion {
	return r.union
}

// The rule configuration.
//
// Union satisfied by
// [user.FirewallAccessRuleListResponseConfigurationFirewallIPConfiguration],
// [user.FirewallAccessRuleListResponseConfigurationFirewallIPV6Configuration],
// [user.FirewallAccessRuleListResponseConfigurationFirewallCIDRConfiguration],
// [user.FirewallAccessRuleListResponseConfigurationFirewallASNConfiguration] or
// [user.FirewallAccessRuleListResponseConfigurationFirewallCountryConfiguration].
type FirewallAccessRuleListResponseConfigurationUnion interface {
	implementsUserFirewallAccessRuleListResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallAccessRuleListResponseConfigurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleListResponseConfigurationFirewallIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleListResponseConfigurationFirewallIPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleListResponseConfigurationFirewallCIDRConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleListResponseConfigurationFirewallASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleListResponseConfigurationFirewallCountryConfiguration{}),
		},
	)
}

type FirewallAccessRuleListResponseConfigurationFirewallIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target FirewallAccessRuleListResponseConfigurationFirewallIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                 `json:"value"`
	JSON  firewallAccessRuleListResponseConfigurationFirewallIPConfigurationJSON `json:"-"`
}

// firewallAccessRuleListResponseConfigurationFirewallIPConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallAccessRuleListResponseConfigurationFirewallIPConfiguration]
type firewallAccessRuleListResponseConfigurationFirewallIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleListResponseConfigurationFirewallIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleListResponseConfigurationFirewallIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleListResponseConfigurationFirewallIPConfiguration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleListResponseConfigurationFirewallIPConfigurationTarget string

const (
	FirewallAccessRuleListResponseConfigurationFirewallIPConfigurationTargetIP FirewallAccessRuleListResponseConfigurationFirewallIPConfigurationTarget = "ip"
)

func (r FirewallAccessRuleListResponseConfigurationFirewallIPConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListResponseConfigurationFirewallIPConfigurationTargetIP:
		return true
	}
	return false
}

type FirewallAccessRuleListResponseConfigurationFirewallIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target FirewallAccessRuleListResponseConfigurationFirewallIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                   `json:"value"`
	JSON  firewallAccessRuleListResponseConfigurationFirewallIPV6ConfigurationJSON `json:"-"`
}

// firewallAccessRuleListResponseConfigurationFirewallIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleListResponseConfigurationFirewallIPV6Configuration]
type firewallAccessRuleListResponseConfigurationFirewallIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleListResponseConfigurationFirewallIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleListResponseConfigurationFirewallIPV6ConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleListResponseConfigurationFirewallIPV6Configuration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type FirewallAccessRuleListResponseConfigurationFirewallIPV6ConfigurationTarget string

const (
	FirewallAccessRuleListResponseConfigurationFirewallIPV6ConfigurationTargetIp6 FirewallAccessRuleListResponseConfigurationFirewallIPV6ConfigurationTarget = "ip6"
)

func (r FirewallAccessRuleListResponseConfigurationFirewallIPV6ConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListResponseConfigurationFirewallIPV6ConfigurationTargetIp6:
		return true
	}
	return false
}

type FirewallAccessRuleListResponseConfigurationFirewallCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target FirewallAccessRuleListResponseConfigurationFirewallCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                   `json:"value"`
	JSON  firewallAccessRuleListResponseConfigurationFirewallCIDRConfigurationJSON `json:"-"`
}

// firewallAccessRuleListResponseConfigurationFirewallCIDRConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleListResponseConfigurationFirewallCIDRConfiguration]
type firewallAccessRuleListResponseConfigurationFirewallCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleListResponseConfigurationFirewallCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleListResponseConfigurationFirewallCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleListResponseConfigurationFirewallCIDRConfiguration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type FirewallAccessRuleListResponseConfigurationFirewallCIDRConfigurationTarget string

const (
	FirewallAccessRuleListResponseConfigurationFirewallCIDRConfigurationTargetIPRange FirewallAccessRuleListResponseConfigurationFirewallCIDRConfigurationTarget = "ip_range"
)

func (r FirewallAccessRuleListResponseConfigurationFirewallCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListResponseConfigurationFirewallCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

type FirewallAccessRuleListResponseConfigurationFirewallASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target FirewallAccessRuleListResponseConfigurationFirewallASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                  `json:"value"`
	JSON  firewallAccessRuleListResponseConfigurationFirewallASNConfigurationJSON `json:"-"`
}

// firewallAccessRuleListResponseConfigurationFirewallASNConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallAccessRuleListResponseConfigurationFirewallASNConfiguration]
type firewallAccessRuleListResponseConfigurationFirewallASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleListResponseConfigurationFirewallASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleListResponseConfigurationFirewallASNConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleListResponseConfigurationFirewallASNConfiguration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type FirewallAccessRuleListResponseConfigurationFirewallASNConfigurationTarget string

const (
	FirewallAccessRuleListResponseConfigurationFirewallASNConfigurationTargetASN FirewallAccessRuleListResponseConfigurationFirewallASNConfigurationTarget = "asn"
)

func (r FirewallAccessRuleListResponseConfigurationFirewallASNConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListResponseConfigurationFirewallASNConfigurationTargetASN:
		return true
	}
	return false
}

type FirewallAccessRuleListResponseConfigurationFirewallCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target FirewallAccessRuleListResponseConfigurationFirewallCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                      `json:"value"`
	JSON  firewallAccessRuleListResponseConfigurationFirewallCountryConfigurationJSON `json:"-"`
}

// firewallAccessRuleListResponseConfigurationFirewallCountryConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleListResponseConfigurationFirewallCountryConfiguration]
type firewallAccessRuleListResponseConfigurationFirewallCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleListResponseConfigurationFirewallCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleListResponseConfigurationFirewallCountryConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleListResponseConfigurationFirewallCountryConfiguration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type FirewallAccessRuleListResponseConfigurationFirewallCountryConfigurationTarget string

const (
	FirewallAccessRuleListResponseConfigurationFirewallCountryConfigurationTargetCountry FirewallAccessRuleListResponseConfigurationFirewallCountryConfigurationTarget = "country"
)

func (r FirewallAccessRuleListResponseConfigurationFirewallCountryConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListResponseConfigurationFirewallCountryConfigurationTargetCountry:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleListResponseConfigurationTarget string

const (
	FirewallAccessRuleListResponseConfigurationTargetIP      FirewallAccessRuleListResponseConfigurationTarget = "ip"
	FirewallAccessRuleListResponseConfigurationTargetIp6     FirewallAccessRuleListResponseConfigurationTarget = "ip6"
	FirewallAccessRuleListResponseConfigurationTargetIPRange FirewallAccessRuleListResponseConfigurationTarget = "ip_range"
	FirewallAccessRuleListResponseConfigurationTargetASN     FirewallAccessRuleListResponseConfigurationTarget = "asn"
	FirewallAccessRuleListResponseConfigurationTargetCountry FirewallAccessRuleListResponseConfigurationTarget = "country"
)

func (r FirewallAccessRuleListResponseConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListResponseConfigurationTargetIP, FirewallAccessRuleListResponseConfigurationTargetIp6, FirewallAccessRuleListResponseConfigurationTargetIPRange, FirewallAccessRuleListResponseConfigurationTargetASN, FirewallAccessRuleListResponseConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type FirewallAccessRuleListResponseMode string

const (
	FirewallAccessRuleListResponseModeBlock            FirewallAccessRuleListResponseMode = "block"
	FirewallAccessRuleListResponseModeChallenge        FirewallAccessRuleListResponseMode = "challenge"
	FirewallAccessRuleListResponseModeWhitelist        FirewallAccessRuleListResponseMode = "whitelist"
	FirewallAccessRuleListResponseModeJsChallenge      FirewallAccessRuleListResponseMode = "js_challenge"
	FirewallAccessRuleListResponseModeManagedChallenge FirewallAccessRuleListResponseMode = "managed_challenge"
)

func (r FirewallAccessRuleListResponseMode) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListResponseModeBlock, FirewallAccessRuleListResponseModeChallenge, FirewallAccessRuleListResponseModeWhitelist, FirewallAccessRuleListResponseModeJsChallenge, FirewallAccessRuleListResponseModeManagedChallenge:
		return true
	}
	return false
}

type FirewallAccessRuleDeleteResponse struct {
	// The unique identifier of the IP Access rule.
	ID   string                               `json:"id"`
	JSON firewallAccessRuleDeleteResponseJSON `json:"-"`
}

// firewallAccessRuleDeleteResponseJSON contains the JSON metadata for the struct
// [FirewallAccessRuleDeleteResponse]
type firewallAccessRuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleEditResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []FirewallAccessRuleEditResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration FirewallAccessRuleEditResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode FirewallAccessRuleEditResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                             `json:"notes"`
	JSON  firewallAccessRuleEditResponseJSON `json:"-"`
}

// firewallAccessRuleEditResponseJSON contains the JSON metadata for the struct
// [FirewallAccessRuleEditResponse]
type firewallAccessRuleEditResponseJSON struct {
	ID            apijson.Field
	AllowedModes  apijson.Field
	Configuration apijson.Field
	Mode          apijson.Field
	CreatedOn     apijson.Field
	ModifiedOn    apijson.Field
	Notes         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type FirewallAccessRuleEditResponseAllowedMode string

const (
	FirewallAccessRuleEditResponseAllowedModeBlock            FirewallAccessRuleEditResponseAllowedMode = "block"
	FirewallAccessRuleEditResponseAllowedModeChallenge        FirewallAccessRuleEditResponseAllowedMode = "challenge"
	FirewallAccessRuleEditResponseAllowedModeWhitelist        FirewallAccessRuleEditResponseAllowedMode = "whitelist"
	FirewallAccessRuleEditResponseAllowedModeJsChallenge      FirewallAccessRuleEditResponseAllowedMode = "js_challenge"
	FirewallAccessRuleEditResponseAllowedModeManagedChallenge FirewallAccessRuleEditResponseAllowedMode = "managed_challenge"
)

func (r FirewallAccessRuleEditResponseAllowedMode) IsKnown() bool {
	switch r {
	case FirewallAccessRuleEditResponseAllowedModeBlock, FirewallAccessRuleEditResponseAllowedModeChallenge, FirewallAccessRuleEditResponseAllowedModeWhitelist, FirewallAccessRuleEditResponseAllowedModeJsChallenge, FirewallAccessRuleEditResponseAllowedModeManagedChallenge:
		return true
	}
	return false
}

// The rule configuration.
type FirewallAccessRuleEditResponseConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target FirewallAccessRuleEditResponseConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                          `json:"value"`
	JSON  firewallAccessRuleEditResponseConfigurationJSON `json:"-"`
	union FirewallAccessRuleEditResponseConfigurationUnion
}

// firewallAccessRuleEditResponseConfigurationJSON contains the JSON metadata for
// the struct [FirewallAccessRuleEditResponseConfiguration]
type firewallAccessRuleEditResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r firewallAccessRuleEditResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r *FirewallAccessRuleEditResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r FirewallAccessRuleEditResponseConfiguration) AsUnion() FirewallAccessRuleEditResponseConfigurationUnion {
	return r.union
}

// The rule configuration.
//
// Union satisfied by
// [user.FirewallAccessRuleEditResponseConfigurationFirewallIPConfiguration],
// [user.FirewallAccessRuleEditResponseConfigurationFirewallIPV6Configuration],
// [user.FirewallAccessRuleEditResponseConfigurationFirewallCIDRConfiguration],
// [user.FirewallAccessRuleEditResponseConfigurationFirewallASNConfiguration] or
// [user.FirewallAccessRuleEditResponseConfigurationFirewallCountryConfiguration].
type FirewallAccessRuleEditResponseConfigurationUnion interface {
	implementsUserFirewallAccessRuleEditResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallAccessRuleEditResponseConfigurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleEditResponseConfigurationFirewallIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleEditResponseConfigurationFirewallIPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleEditResponseConfigurationFirewallCIDRConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleEditResponseConfigurationFirewallASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleEditResponseConfigurationFirewallCountryConfiguration{}),
		},
	)
}

type FirewallAccessRuleEditResponseConfigurationFirewallIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target FirewallAccessRuleEditResponseConfigurationFirewallIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                 `json:"value"`
	JSON  firewallAccessRuleEditResponseConfigurationFirewallIPConfigurationJSON `json:"-"`
}

// firewallAccessRuleEditResponseConfigurationFirewallIPConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallAccessRuleEditResponseConfigurationFirewallIPConfiguration]
type firewallAccessRuleEditResponseConfigurationFirewallIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseConfigurationFirewallIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseConfigurationFirewallIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleEditResponseConfigurationFirewallIPConfiguration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleEditResponseConfigurationFirewallIPConfigurationTarget string

const (
	FirewallAccessRuleEditResponseConfigurationFirewallIPConfigurationTargetIP FirewallAccessRuleEditResponseConfigurationFirewallIPConfigurationTarget = "ip"
)

func (r FirewallAccessRuleEditResponseConfigurationFirewallIPConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleEditResponseConfigurationFirewallIPConfigurationTargetIP:
		return true
	}
	return false
}

type FirewallAccessRuleEditResponseConfigurationFirewallIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target FirewallAccessRuleEditResponseConfigurationFirewallIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                   `json:"value"`
	JSON  firewallAccessRuleEditResponseConfigurationFirewallIPV6ConfigurationJSON `json:"-"`
}

// firewallAccessRuleEditResponseConfigurationFirewallIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleEditResponseConfigurationFirewallIPV6Configuration]
type firewallAccessRuleEditResponseConfigurationFirewallIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseConfigurationFirewallIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseConfigurationFirewallIPV6ConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleEditResponseConfigurationFirewallIPV6Configuration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type FirewallAccessRuleEditResponseConfigurationFirewallIPV6ConfigurationTarget string

const (
	FirewallAccessRuleEditResponseConfigurationFirewallIPV6ConfigurationTargetIp6 FirewallAccessRuleEditResponseConfigurationFirewallIPV6ConfigurationTarget = "ip6"
)

func (r FirewallAccessRuleEditResponseConfigurationFirewallIPV6ConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleEditResponseConfigurationFirewallIPV6ConfigurationTargetIp6:
		return true
	}
	return false
}

type FirewallAccessRuleEditResponseConfigurationFirewallCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target FirewallAccessRuleEditResponseConfigurationFirewallCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                   `json:"value"`
	JSON  firewallAccessRuleEditResponseConfigurationFirewallCIDRConfigurationJSON `json:"-"`
}

// firewallAccessRuleEditResponseConfigurationFirewallCIDRConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleEditResponseConfigurationFirewallCIDRConfiguration]
type firewallAccessRuleEditResponseConfigurationFirewallCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseConfigurationFirewallCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseConfigurationFirewallCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleEditResponseConfigurationFirewallCIDRConfiguration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type FirewallAccessRuleEditResponseConfigurationFirewallCIDRConfigurationTarget string

const (
	FirewallAccessRuleEditResponseConfigurationFirewallCIDRConfigurationTargetIPRange FirewallAccessRuleEditResponseConfigurationFirewallCIDRConfigurationTarget = "ip_range"
)

func (r FirewallAccessRuleEditResponseConfigurationFirewallCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleEditResponseConfigurationFirewallCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

type FirewallAccessRuleEditResponseConfigurationFirewallASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target FirewallAccessRuleEditResponseConfigurationFirewallASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                  `json:"value"`
	JSON  firewallAccessRuleEditResponseConfigurationFirewallASNConfigurationJSON `json:"-"`
}

// firewallAccessRuleEditResponseConfigurationFirewallASNConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallAccessRuleEditResponseConfigurationFirewallASNConfiguration]
type firewallAccessRuleEditResponseConfigurationFirewallASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseConfigurationFirewallASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseConfigurationFirewallASNConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleEditResponseConfigurationFirewallASNConfiguration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type FirewallAccessRuleEditResponseConfigurationFirewallASNConfigurationTarget string

const (
	FirewallAccessRuleEditResponseConfigurationFirewallASNConfigurationTargetASN FirewallAccessRuleEditResponseConfigurationFirewallASNConfigurationTarget = "asn"
)

func (r FirewallAccessRuleEditResponseConfigurationFirewallASNConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleEditResponseConfigurationFirewallASNConfigurationTargetASN:
		return true
	}
	return false
}

type FirewallAccessRuleEditResponseConfigurationFirewallCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target FirewallAccessRuleEditResponseConfigurationFirewallCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                      `json:"value"`
	JSON  firewallAccessRuleEditResponseConfigurationFirewallCountryConfigurationJSON `json:"-"`
}

// firewallAccessRuleEditResponseConfigurationFirewallCountryConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleEditResponseConfigurationFirewallCountryConfiguration]
type firewallAccessRuleEditResponseConfigurationFirewallCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseConfigurationFirewallCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseConfigurationFirewallCountryConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleEditResponseConfigurationFirewallCountryConfiguration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type FirewallAccessRuleEditResponseConfigurationFirewallCountryConfigurationTarget string

const (
	FirewallAccessRuleEditResponseConfigurationFirewallCountryConfigurationTargetCountry FirewallAccessRuleEditResponseConfigurationFirewallCountryConfigurationTarget = "country"
)

func (r FirewallAccessRuleEditResponseConfigurationFirewallCountryConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleEditResponseConfigurationFirewallCountryConfigurationTargetCountry:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleEditResponseConfigurationTarget string

const (
	FirewallAccessRuleEditResponseConfigurationTargetIP      FirewallAccessRuleEditResponseConfigurationTarget = "ip"
	FirewallAccessRuleEditResponseConfigurationTargetIp6     FirewallAccessRuleEditResponseConfigurationTarget = "ip6"
	FirewallAccessRuleEditResponseConfigurationTargetIPRange FirewallAccessRuleEditResponseConfigurationTarget = "ip_range"
	FirewallAccessRuleEditResponseConfigurationTargetASN     FirewallAccessRuleEditResponseConfigurationTarget = "asn"
	FirewallAccessRuleEditResponseConfigurationTargetCountry FirewallAccessRuleEditResponseConfigurationTarget = "country"
)

func (r FirewallAccessRuleEditResponseConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleEditResponseConfigurationTargetIP, FirewallAccessRuleEditResponseConfigurationTargetIp6, FirewallAccessRuleEditResponseConfigurationTargetIPRange, FirewallAccessRuleEditResponseConfigurationTargetASN, FirewallAccessRuleEditResponseConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type FirewallAccessRuleEditResponseMode string

const (
	FirewallAccessRuleEditResponseModeBlock            FirewallAccessRuleEditResponseMode = "block"
	FirewallAccessRuleEditResponseModeChallenge        FirewallAccessRuleEditResponseMode = "challenge"
	FirewallAccessRuleEditResponseModeWhitelist        FirewallAccessRuleEditResponseMode = "whitelist"
	FirewallAccessRuleEditResponseModeJsChallenge      FirewallAccessRuleEditResponseMode = "js_challenge"
	FirewallAccessRuleEditResponseModeManagedChallenge FirewallAccessRuleEditResponseMode = "managed_challenge"
)

func (r FirewallAccessRuleEditResponseMode) IsKnown() bool {
	switch r {
	case FirewallAccessRuleEditResponseModeBlock, FirewallAccessRuleEditResponseModeChallenge, FirewallAccessRuleEditResponseModeWhitelist, FirewallAccessRuleEditResponseModeJsChallenge, FirewallAccessRuleEditResponseModeManagedChallenge:
		return true
	}
	return false
}

type FirewallAccessRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[FirewallAccessRuleNewParamsConfigurationUnion] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallAccessRuleNewParamsMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r FirewallAccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
type FirewallAccessRuleNewParamsConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfiguration) implementsUserFirewallAccessRuleNewParamsConfigurationUnion() {
}

// The rule configuration.
//
// Satisfied by
// [user.FirewallAccessRuleNewParamsConfigurationFirewallIPConfiguration],
// [user.FirewallAccessRuleNewParamsConfigurationFirewallIPV6Configuration],
// [user.FirewallAccessRuleNewParamsConfigurationFirewallCIDRConfiguration],
// [user.FirewallAccessRuleNewParamsConfigurationFirewallASNConfiguration],
// [user.FirewallAccessRuleNewParamsConfigurationFirewallCountryConfiguration],
// [FirewallAccessRuleNewParamsConfiguration].
type FirewallAccessRuleNewParamsConfigurationUnion interface {
	implementsUserFirewallAccessRuleNewParamsConfigurationUnion()
}

type FirewallAccessRuleNewParamsConfigurationFirewallIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationFirewallIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationFirewallIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationFirewallIPConfiguration) implementsUserFirewallAccessRuleNewParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleNewParamsConfigurationFirewallIPConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationFirewallIPConfigurationTargetIP FirewallAccessRuleNewParamsConfigurationFirewallIPConfigurationTarget = "ip"
)

func (r FirewallAccessRuleNewParamsConfigurationFirewallIPConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsConfigurationFirewallIPConfigurationTargetIP:
		return true
	}
	return false
}

type FirewallAccessRuleNewParamsConfigurationFirewallIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationFirewallIPV6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationFirewallIPV6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationFirewallIPV6Configuration) implementsUserFirewallAccessRuleNewParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type FirewallAccessRuleNewParamsConfigurationFirewallIPV6ConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationFirewallIPV6ConfigurationTargetIp6 FirewallAccessRuleNewParamsConfigurationFirewallIPV6ConfigurationTarget = "ip6"
)

func (r FirewallAccessRuleNewParamsConfigurationFirewallIPV6ConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsConfigurationFirewallIPV6ConfigurationTargetIp6:
		return true
	}
	return false
}

type FirewallAccessRuleNewParamsConfigurationFirewallCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationFirewallCIDRConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationFirewallCIDRConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationFirewallCIDRConfiguration) implementsUserFirewallAccessRuleNewParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type FirewallAccessRuleNewParamsConfigurationFirewallCIDRConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationFirewallCIDRConfigurationTargetIPRange FirewallAccessRuleNewParamsConfigurationFirewallCIDRConfigurationTarget = "ip_range"
)

func (r FirewallAccessRuleNewParamsConfigurationFirewallCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsConfigurationFirewallCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

type FirewallAccessRuleNewParamsConfigurationFirewallASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationFirewallASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationFirewallASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationFirewallASNConfiguration) implementsUserFirewallAccessRuleNewParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type FirewallAccessRuleNewParamsConfigurationFirewallASNConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationFirewallASNConfigurationTargetASN FirewallAccessRuleNewParamsConfigurationFirewallASNConfigurationTarget = "asn"
)

func (r FirewallAccessRuleNewParamsConfigurationFirewallASNConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsConfigurationFirewallASNConfigurationTargetASN:
		return true
	}
	return false
}

type FirewallAccessRuleNewParamsConfigurationFirewallCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationFirewallCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationFirewallCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationFirewallCountryConfiguration) implementsUserFirewallAccessRuleNewParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type FirewallAccessRuleNewParamsConfigurationFirewallCountryConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationFirewallCountryConfigurationTargetCountry FirewallAccessRuleNewParamsConfigurationFirewallCountryConfigurationTarget = "country"
)

func (r FirewallAccessRuleNewParamsConfigurationFirewallCountryConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsConfigurationFirewallCountryConfigurationTargetCountry:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleNewParamsConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationTargetIP      FirewallAccessRuleNewParamsConfigurationTarget = "ip"
	FirewallAccessRuleNewParamsConfigurationTargetIp6     FirewallAccessRuleNewParamsConfigurationTarget = "ip6"
	FirewallAccessRuleNewParamsConfigurationTargetIPRange FirewallAccessRuleNewParamsConfigurationTarget = "ip_range"
	FirewallAccessRuleNewParamsConfigurationTargetASN     FirewallAccessRuleNewParamsConfigurationTarget = "asn"
	FirewallAccessRuleNewParamsConfigurationTargetCountry FirewallAccessRuleNewParamsConfigurationTarget = "country"
)

func (r FirewallAccessRuleNewParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsConfigurationTargetIP, FirewallAccessRuleNewParamsConfigurationTargetIp6, FirewallAccessRuleNewParamsConfigurationTargetIPRange, FirewallAccessRuleNewParamsConfigurationTargetASN, FirewallAccessRuleNewParamsConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type FirewallAccessRuleNewParamsMode string

const (
	FirewallAccessRuleNewParamsModeBlock            FirewallAccessRuleNewParamsMode = "block"
	FirewallAccessRuleNewParamsModeChallenge        FirewallAccessRuleNewParamsMode = "challenge"
	FirewallAccessRuleNewParamsModeWhitelist        FirewallAccessRuleNewParamsMode = "whitelist"
	FirewallAccessRuleNewParamsModeJsChallenge      FirewallAccessRuleNewParamsMode = "js_challenge"
	FirewallAccessRuleNewParamsModeManagedChallenge FirewallAccessRuleNewParamsMode = "managed_challenge"
)

func (r FirewallAccessRuleNewParamsMode) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsModeBlock, FirewallAccessRuleNewParamsModeChallenge, FirewallAccessRuleNewParamsModeWhitelist, FirewallAccessRuleNewParamsModeJsChallenge, FirewallAccessRuleNewParamsModeManagedChallenge:
		return true
	}
	return false
}

type FirewallAccessRuleNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   FirewallAccessRuleNewResponse                             `json:"result,required"`
	// Whether the API call was successful
	Success FirewallAccessRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallAccessRuleNewResponseEnvelopeJSON    `json:"-"`
}

// firewallAccessRuleNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallAccessRuleNewResponseEnvelope]
type firewallAccessRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleNewResponseEnvelopeSuccess bool

const (
	FirewallAccessRuleNewResponseEnvelopeSuccessTrue FirewallAccessRuleNewResponseEnvelopeSuccess = true
)

func (r FirewallAccessRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FirewallAccessRuleListParams struct {
	// The direction used to sort returned rules.
	Direction     param.Field[FirewallAccessRuleListParamsDirection]     `query:"direction"`
	EgsPagination param.Field[FirewallAccessRuleListParamsEgsPagination] `query:"egs-pagination"`
	Filters       param.Field[FirewallAccessRuleListParamsFilters]       `query:"filters"`
	// The field used to sort returned rules.
	Order param.Field[FirewallAccessRuleListParamsOrder] `query:"order"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [FirewallAccessRuleListParams]'s query parameters as
// `url.Values`.
func (r FirewallAccessRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type FirewallAccessRuleListParamsDirection string

const (
	FirewallAccessRuleListParamsDirectionAsc  FirewallAccessRuleListParamsDirection = "asc"
	FirewallAccessRuleListParamsDirectionDesc FirewallAccessRuleListParamsDirection = "desc"
)

func (r FirewallAccessRuleListParamsDirection) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListParamsDirectionAsc, FirewallAccessRuleListParamsDirectionDesc:
		return true
	}
	return false
}

type FirewallAccessRuleListParamsEgsPagination struct {
	Json param.Field[FirewallAccessRuleListParamsEgsPaginationJson] `query:"json"`
}

// URLQuery serializes [FirewallAccessRuleListParamsEgsPagination]'s query
// parameters as `url.Values`.
func (r FirewallAccessRuleListParamsEgsPagination) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallAccessRuleListParamsEgsPaginationJson struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [FirewallAccessRuleListParamsEgsPaginationJson]'s query
// parameters as `url.Values`.
func (r FirewallAccessRuleListParamsEgsPaginationJson) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallAccessRuleListParamsFilters struct {
	// The target to search in existing rules.
	ConfigurationTarget param.Field[FirewallAccessRuleListParamsFiltersConfigurationTarget] `query:"configuration.target"`
	// The target value to search for in existing rules: an IP address, an IP address
	// range, or a country code, depending on the provided `configuration.target`.
	// Notes: You can search for a single IPv4 address, an IP address range with a
	// subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	ConfigurationValue param.Field[string] `query:"configuration.value"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[FirewallAccessRuleListParamsFiltersMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallAccessRuleListParamsFiltersMode] `query:"mode"`
	// The string to search for in the notes of existing IP Access rules. Notes: For
	// example, the string 'attack' would match IP Access rules with notes 'Attack
	// 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
}

// URLQuery serializes [FirewallAccessRuleListParamsFilters]'s query parameters as
// `url.Values`.
func (r FirewallAccessRuleListParamsFilters) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The target to search in existing rules.
type FirewallAccessRuleListParamsFiltersConfigurationTarget string

const (
	FirewallAccessRuleListParamsFiltersConfigurationTargetIP      FirewallAccessRuleListParamsFiltersConfigurationTarget = "ip"
	FirewallAccessRuleListParamsFiltersConfigurationTargetIPRange FirewallAccessRuleListParamsFiltersConfigurationTarget = "ip_range"
	FirewallAccessRuleListParamsFiltersConfigurationTargetASN     FirewallAccessRuleListParamsFiltersConfigurationTarget = "asn"
	FirewallAccessRuleListParamsFiltersConfigurationTargetCountry FirewallAccessRuleListParamsFiltersConfigurationTarget = "country"
)

func (r FirewallAccessRuleListParamsFiltersConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListParamsFiltersConfigurationTargetIP, FirewallAccessRuleListParamsFiltersConfigurationTargetIPRange, FirewallAccessRuleListParamsFiltersConfigurationTargetASN, FirewallAccessRuleListParamsFiltersConfigurationTargetCountry:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type FirewallAccessRuleListParamsFiltersMatch string

const (
	FirewallAccessRuleListParamsFiltersMatchAny FirewallAccessRuleListParamsFiltersMatch = "any"
	FirewallAccessRuleListParamsFiltersMatchAll FirewallAccessRuleListParamsFiltersMatch = "all"
)

func (r FirewallAccessRuleListParamsFiltersMatch) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListParamsFiltersMatchAny, FirewallAccessRuleListParamsFiltersMatchAll:
		return true
	}
	return false
}

// The action to apply to a matched request.
type FirewallAccessRuleListParamsFiltersMode string

const (
	FirewallAccessRuleListParamsFiltersModeBlock            FirewallAccessRuleListParamsFiltersMode = "block"
	FirewallAccessRuleListParamsFiltersModeChallenge        FirewallAccessRuleListParamsFiltersMode = "challenge"
	FirewallAccessRuleListParamsFiltersModeWhitelist        FirewallAccessRuleListParamsFiltersMode = "whitelist"
	FirewallAccessRuleListParamsFiltersModeJsChallenge      FirewallAccessRuleListParamsFiltersMode = "js_challenge"
	FirewallAccessRuleListParamsFiltersModeManagedChallenge FirewallAccessRuleListParamsFiltersMode = "managed_challenge"
)

func (r FirewallAccessRuleListParamsFiltersMode) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListParamsFiltersModeBlock, FirewallAccessRuleListParamsFiltersModeChallenge, FirewallAccessRuleListParamsFiltersModeWhitelist, FirewallAccessRuleListParamsFiltersModeJsChallenge, FirewallAccessRuleListParamsFiltersModeManagedChallenge:
		return true
	}
	return false
}

// The field used to sort returned rules.
type FirewallAccessRuleListParamsOrder string

const (
	FirewallAccessRuleListParamsOrderConfigurationTarget FirewallAccessRuleListParamsOrder = "configuration.target"
	FirewallAccessRuleListParamsOrderConfigurationValue  FirewallAccessRuleListParamsOrder = "configuration.value"
	FirewallAccessRuleListParamsOrderMode                FirewallAccessRuleListParamsOrder = "mode"
)

func (r FirewallAccessRuleListParamsOrder) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListParamsOrderConfigurationTarget, FirewallAccessRuleListParamsOrderConfigurationValue, FirewallAccessRuleListParamsOrderMode:
		return true
	}
	return false
}

type FirewallAccessRuleDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallAccessRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallAccessRuleDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   FirewallAccessRuleDeleteResponse                          `json:"result,required"`
	// Whether the API call was successful
	Success FirewallAccessRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallAccessRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// firewallAccessRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallAccessRuleDeleteResponseEnvelope]
type firewallAccessRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleDeleteResponseEnvelopeSuccess bool

const (
	FirewallAccessRuleDeleteResponseEnvelopeSuccessTrue FirewallAccessRuleDeleteResponseEnvelopeSuccess = true
)

func (r FirewallAccessRuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallAccessRuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FirewallAccessRuleEditParams struct {
	// The action to apply to a matched request.
	Mode param.Field[FirewallAccessRuleEditParamsMode] `json:"mode"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r FirewallAccessRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to apply to a matched request.
type FirewallAccessRuleEditParamsMode string

const (
	FirewallAccessRuleEditParamsModeBlock            FirewallAccessRuleEditParamsMode = "block"
	FirewallAccessRuleEditParamsModeChallenge        FirewallAccessRuleEditParamsMode = "challenge"
	FirewallAccessRuleEditParamsModeWhitelist        FirewallAccessRuleEditParamsMode = "whitelist"
	FirewallAccessRuleEditParamsModeJsChallenge      FirewallAccessRuleEditParamsMode = "js_challenge"
	FirewallAccessRuleEditParamsModeManagedChallenge FirewallAccessRuleEditParamsMode = "managed_challenge"
)

func (r FirewallAccessRuleEditParamsMode) IsKnown() bool {
	switch r {
	case FirewallAccessRuleEditParamsModeBlock, FirewallAccessRuleEditParamsModeChallenge, FirewallAccessRuleEditParamsModeWhitelist, FirewallAccessRuleEditParamsModeJsChallenge, FirewallAccessRuleEditParamsModeManagedChallenge:
		return true
	}
	return false
}

type FirewallAccessRuleEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   FirewallAccessRuleEditResponse                            `json:"result,required"`
	// Whether the API call was successful
	Success FirewallAccessRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallAccessRuleEditResponseEnvelopeJSON    `json:"-"`
}

// firewallAccessRuleEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallAccessRuleEditResponseEnvelope]
type firewallAccessRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleEditResponseEnvelopeSuccess bool

const (
	FirewallAccessRuleEditResponseEnvelopeSuccessTrue FirewallAccessRuleEditResponseEnvelopeSuccess = true
)

func (r FirewallAccessRuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallAccessRuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
