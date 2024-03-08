// File generated from our OpenAPI spec by Stainless.

package user

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *FirewallAccessRuleService) List(ctx context.Context, query FirewallAccessRuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[FirewallAccessRuleListResponse], err error) {
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
func (r *FirewallAccessRuleService) ListAutoPaging(ctx context.Context, query FirewallAccessRuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[FirewallAccessRuleListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Deletes an IP Access rule at the user level.
//
// Note: Deleting a user-level rule will affect all zones owned by the user.
func (r *FirewallAccessRuleService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *FirewallAccessRuleDeleteResponse, err error) {
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

// The rule configuration.
//
// Union satisfied by
// [user.FirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfiguration],
// [user.FirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6Configuration],
// [user.FirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfiguration],
// [user.FirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfiguration] or
// [user.FirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfiguration].
type FirewallAccessRuleNewResponseConfiguration interface {
	implementsUserFirewallAccessRuleNewResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallAccessRuleNewResponseConfiguration)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfiguration{}),
		},
	)
}

type FirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target FirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                 `json:"value"`
	JSON  firewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// firewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfiguration]
type firewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationTarget string

const (
	FirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationTargetIP FirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type FirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target FirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                   `json:"value"`
	JSON  firewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// firewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6Configuration]
type firewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type FirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	FirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationTargetIp6 FirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type FirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target FirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                   `json:"value"`
	JSON  firewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationJSON `json:"-"`
}

// firewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfiguration]
type firewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type FirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationTarget string

const (
	FirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationTargetIPRange FirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type FirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target FirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                  `json:"value"`
	JSON  firewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationJSON `json:"-"`
}

// firewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfiguration]
type firewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type FirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationTarget string

const (
	FirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationTargetASN FirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type FirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target FirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                      `json:"value"`
	JSON  firewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// firewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfiguration]
type firewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type FirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationTarget string

const (
	FirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationTargetCountry FirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type FirewallAccessRuleNewResponseMode string

const (
	FirewallAccessRuleNewResponseModeBlock            FirewallAccessRuleNewResponseMode = "block"
	FirewallAccessRuleNewResponseModeChallenge        FirewallAccessRuleNewResponseMode = "challenge"
	FirewallAccessRuleNewResponseModeWhitelist        FirewallAccessRuleNewResponseMode = "whitelist"
	FirewallAccessRuleNewResponseModeJsChallenge      FirewallAccessRuleNewResponseMode = "js_challenge"
	FirewallAccessRuleNewResponseModeManagedChallenge FirewallAccessRuleNewResponseMode = "managed_challenge"
)

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

// The rule configuration.
//
// Union satisfied by
// [user.FirewallAccessRuleListResponseConfigurationLegacyJhsIPConfiguration],
// [user.FirewallAccessRuleListResponseConfigurationLegacyJhsIPV6Configuration],
// [user.FirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfiguration],
// [user.FirewallAccessRuleListResponseConfigurationLegacyJhsASNConfiguration] or
// [user.FirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfiguration].
type FirewallAccessRuleListResponseConfiguration interface {
	implementsUserFirewallAccessRuleListResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallAccessRuleListResponseConfiguration)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleListResponseConfigurationLegacyJhsIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleListResponseConfigurationLegacyJhsIPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleListResponseConfigurationLegacyJhsASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfiguration{}),
		},
	)
}

type FirewallAccessRuleListResponseConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target FirewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                  `json:"value"`
	JSON  firewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// firewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallAccessRuleListResponseConfigurationLegacyJhsIPConfiguration]
type firewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleListResponseConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleListResponseConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationTarget string

const (
	FirewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationTargetIP FirewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type FirewallAccessRuleListResponseConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target FirewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                    `json:"value"`
	JSON  firewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// firewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleListResponseConfigurationLegacyJhsIPV6Configuration]
type firewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleListResponseConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleListResponseConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type FirewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	FirewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationTargetIp6 FirewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type FirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target FirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                    `json:"value"`
	JSON  firewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationJSON `json:"-"`
}

// firewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfiguration]
type firewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type FirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationTarget string

const (
	FirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationTargetIPRange FirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type FirewallAccessRuleListResponseConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target FirewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                   `json:"value"`
	JSON  firewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationJSON `json:"-"`
}

// firewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleListResponseConfigurationLegacyJhsASNConfiguration]
type firewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleListResponseConfigurationLegacyJhsASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleListResponseConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type FirewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationTarget string

const (
	FirewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationTargetASN FirewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type FirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target FirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                       `json:"value"`
	JSON  firewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// firewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfiguration]
type firewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type FirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationTarget string

const (
	FirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationTargetCountry FirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type FirewallAccessRuleListResponseMode string

const (
	FirewallAccessRuleListResponseModeBlock            FirewallAccessRuleListResponseMode = "block"
	FirewallAccessRuleListResponseModeChallenge        FirewallAccessRuleListResponseMode = "challenge"
	FirewallAccessRuleListResponseModeWhitelist        FirewallAccessRuleListResponseMode = "whitelist"
	FirewallAccessRuleListResponseModeJsChallenge      FirewallAccessRuleListResponseMode = "js_challenge"
	FirewallAccessRuleListResponseModeManagedChallenge FirewallAccessRuleListResponseMode = "managed_challenge"
)

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

// The rule configuration.
//
// Union satisfied by
// [user.FirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfiguration],
// [user.FirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6Configuration],
// [user.FirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfiguration],
// [user.FirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfiguration] or
// [user.FirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfiguration].
type FirewallAccessRuleEditResponseConfiguration interface {
	implementsUserFirewallAccessRuleEditResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallAccessRuleEditResponseConfiguration)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfiguration{}),
		},
	)
}

type FirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target FirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                  `json:"value"`
	JSON  firewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// firewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfiguration]
type firewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationTarget string

const (
	FirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationTargetIP FirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type FirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target FirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                    `json:"value"`
	JSON  firewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// firewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6Configuration]
type firewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type FirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	FirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationTargetIp6 FirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type FirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target FirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                    `json:"value"`
	JSON  firewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationJSON `json:"-"`
}

// firewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfiguration]
type firewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type FirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationTarget string

const (
	FirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationTargetIPRange FirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type FirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target FirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                   `json:"value"`
	JSON  firewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationJSON `json:"-"`
}

// firewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfiguration]
type firewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type FirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationTarget string

const (
	FirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationTargetASN FirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type FirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target FirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                       `json:"value"`
	JSON  firewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// firewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfiguration]
type firewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type FirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationTarget string

const (
	FirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationTargetCountry FirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type FirewallAccessRuleEditResponseMode string

const (
	FirewallAccessRuleEditResponseModeBlock            FirewallAccessRuleEditResponseMode = "block"
	FirewallAccessRuleEditResponseModeChallenge        FirewallAccessRuleEditResponseMode = "challenge"
	FirewallAccessRuleEditResponseModeWhitelist        FirewallAccessRuleEditResponseMode = "whitelist"
	FirewallAccessRuleEditResponseModeJsChallenge      FirewallAccessRuleEditResponseMode = "js_challenge"
	FirewallAccessRuleEditResponseModeManagedChallenge FirewallAccessRuleEditResponseMode = "managed_challenge"
)

type FirewallAccessRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[FirewallAccessRuleNewParamsConfiguration] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallAccessRuleNewParamsMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r FirewallAccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
//
// Satisfied by
// [user.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration],
// [user.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration],
// [user.FirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfiguration],
// [user.FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration],
// [user.FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration].
type FirewallAccessRuleNewParamsConfiguration interface {
	implementsUserFirewallAccessRuleNewParamsConfiguration()
}

type FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTargetIP FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTargetIp6 FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type FirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type FirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfigurationTargetIPRange FirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTargetASN FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTargetCountry FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type FirewallAccessRuleNewParamsMode string

const (
	FirewallAccessRuleNewParamsModeBlock            FirewallAccessRuleNewParamsMode = "block"
	FirewallAccessRuleNewParamsModeChallenge        FirewallAccessRuleNewParamsMode = "challenge"
	FirewallAccessRuleNewParamsModeWhitelist        FirewallAccessRuleNewParamsMode = "whitelist"
	FirewallAccessRuleNewParamsModeJsChallenge      FirewallAccessRuleNewParamsMode = "js_challenge"
	FirewallAccessRuleNewParamsModeManagedChallenge FirewallAccessRuleNewParamsMode = "managed_challenge"
)

type FirewallAccessRuleNewResponseEnvelope struct {
	Errors   []FirewallAccessRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallAccessRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallAccessRuleNewResponse                   `json:"result,required,nullable"`
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

type FirewallAccessRuleNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    firewallAccessRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallAccessRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallAccessRuleNewResponseEnvelopeErrors]
type firewallAccessRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    firewallAccessRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallAccessRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallAccessRuleNewResponseEnvelopeMessages]
type firewallAccessRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleNewResponseEnvelopeSuccess bool

const (
	FirewallAccessRuleNewResponseEnvelopeSuccessTrue FirewallAccessRuleNewResponseEnvelopeSuccess = true
)

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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type FirewallAccessRuleListParamsDirection string

const (
	FirewallAccessRuleListParamsDirectionAsc  FirewallAccessRuleListParamsDirection = "asc"
	FirewallAccessRuleListParamsDirectionDesc FirewallAccessRuleListParamsDirection = "desc"
)

type FirewallAccessRuleListParamsEgsPagination struct {
	Json param.Field[FirewallAccessRuleListParamsEgsPaginationJson] `query:"json"`
}

// URLQuery serializes [FirewallAccessRuleListParamsEgsPagination]'s query
// parameters as `url.Values`.
func (r FirewallAccessRuleListParamsEgsPagination) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type FirewallAccessRuleListParamsFiltersMatch string

const (
	FirewallAccessRuleListParamsFiltersMatchAny FirewallAccessRuleListParamsFiltersMatch = "any"
	FirewallAccessRuleListParamsFiltersMatchAll FirewallAccessRuleListParamsFiltersMatch = "all"
)

// The action to apply to a matched request.
type FirewallAccessRuleListParamsFiltersMode string

const (
	FirewallAccessRuleListParamsFiltersModeBlock            FirewallAccessRuleListParamsFiltersMode = "block"
	FirewallAccessRuleListParamsFiltersModeChallenge        FirewallAccessRuleListParamsFiltersMode = "challenge"
	FirewallAccessRuleListParamsFiltersModeWhitelist        FirewallAccessRuleListParamsFiltersMode = "whitelist"
	FirewallAccessRuleListParamsFiltersModeJsChallenge      FirewallAccessRuleListParamsFiltersMode = "js_challenge"
	FirewallAccessRuleListParamsFiltersModeManagedChallenge FirewallAccessRuleListParamsFiltersMode = "managed_challenge"
)

// The field used to sort returned rules.
type FirewallAccessRuleListParamsOrder string

const (
	FirewallAccessRuleListParamsOrderConfigurationTarget FirewallAccessRuleListParamsOrder = "configuration.target"
	FirewallAccessRuleListParamsOrderConfigurationValue  FirewallAccessRuleListParamsOrder = "configuration.value"
	FirewallAccessRuleListParamsOrderMode                FirewallAccessRuleListParamsOrder = "mode"
)

type FirewallAccessRuleDeleteResponseEnvelope struct {
	Errors   []FirewallAccessRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallAccessRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallAccessRuleDeleteResponse                   `json:"result,required,nullable"`
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

type FirewallAccessRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    firewallAccessRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallAccessRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FirewallAccessRuleDeleteResponseEnvelopeErrors]
type firewallAccessRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    firewallAccessRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallAccessRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallAccessRuleDeleteResponseEnvelopeMessages]
type firewallAccessRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleDeleteResponseEnvelopeSuccess bool

const (
	FirewallAccessRuleDeleteResponseEnvelopeSuccessTrue FirewallAccessRuleDeleteResponseEnvelopeSuccess = true
)

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

type FirewallAccessRuleEditResponseEnvelope struct {
	Errors   []FirewallAccessRuleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallAccessRuleEditResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallAccessRuleEditResponse                   `json:"result,required,nullable"`
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

type FirewallAccessRuleEditResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    firewallAccessRuleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallAccessRuleEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallAccessRuleEditResponseEnvelopeErrors]
type firewallAccessRuleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleEditResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    firewallAccessRuleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallAccessRuleEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallAccessRuleEditResponseEnvelopeMessages]
type firewallAccessRuleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleEditResponseEnvelopeSuccess bool

const (
	FirewallAccessRuleEditResponseEnvelopeSuccessTrue FirewallAccessRuleEditResponseEnvelopeSuccess = true
)
