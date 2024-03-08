// File generated from our OpenAPI spec by Stainless.

package cloudflare

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

// UserFirewallAccessRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserFirewallAccessRuleService]
// method instead.
type UserFirewallAccessRuleService struct {
	Options []option.RequestOption
}

// NewUserFirewallAccessRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserFirewallAccessRuleService(opts ...option.RequestOption) (r *UserFirewallAccessRuleService) {
	r = &UserFirewallAccessRuleService{}
	r.Options = opts
	return
}

// Creates a new IP Access rule for all zones owned by the current user.
//
// Note: To create an IP Access rule that applies to a specific zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *UserFirewallAccessRuleService) New(ctx context.Context, body UserFirewallAccessRuleNewParams, opts ...option.RequestOption) (res *UserFirewallAccessRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserFirewallAccessRuleNewResponseEnvelope
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
func (r *UserFirewallAccessRuleService) List(ctx context.Context, query UserFirewallAccessRuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[UserFirewallAccessRuleListResponse], err error) {
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
func (r *UserFirewallAccessRuleService) ListAutoPaging(ctx context.Context, query UserFirewallAccessRuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[UserFirewallAccessRuleListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Deletes an IP Access rule at the user level.
//
// Note: Deleting a user-level rule will affect all zones owned by the user.
func (r *UserFirewallAccessRuleService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserFirewallAccessRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserFirewallAccessRuleDeleteResponseEnvelope
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
func (r *UserFirewallAccessRuleService) Edit(ctx context.Context, identifier string, body UserFirewallAccessRuleEditParams, opts ...option.RequestOption) (res *UserFirewallAccessRuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserFirewallAccessRuleEditResponseEnvelope
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserFirewallAccessRuleNewResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []UserFirewallAccessRuleNewResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration UserFirewallAccessRuleNewResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode UserFirewallAccessRuleNewResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                `json:"notes"`
	JSON  userFirewallAccessRuleNewResponseJSON `json:"-"`
}

// userFirewallAccessRuleNewResponseJSON contains the JSON metadata for the struct
// [UserFirewallAccessRuleNewResponse]
type userFirewallAccessRuleNewResponseJSON struct {
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

func (r *UserFirewallAccessRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type UserFirewallAccessRuleNewResponseAllowedMode string

const (
	UserFirewallAccessRuleNewResponseAllowedModeBlock            UserFirewallAccessRuleNewResponseAllowedMode = "block"
	UserFirewallAccessRuleNewResponseAllowedModeChallenge        UserFirewallAccessRuleNewResponseAllowedMode = "challenge"
	UserFirewallAccessRuleNewResponseAllowedModeWhitelist        UserFirewallAccessRuleNewResponseAllowedMode = "whitelist"
	UserFirewallAccessRuleNewResponseAllowedModeJsChallenge      UserFirewallAccessRuleNewResponseAllowedMode = "js_challenge"
	UserFirewallAccessRuleNewResponseAllowedModeManagedChallenge UserFirewallAccessRuleNewResponseAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfiguration],
// [UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6Configuration],
// [UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfiguration],
// [UserFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfiguration] or
// [UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfiguration].
type UserFirewallAccessRuleNewResponseConfiguration interface {
	implementsUserFirewallAccessRuleNewResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserFirewallAccessRuleNewResponseConfiguration)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfiguration{}),
		},
	)
}

type UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                     `json:"value"`
	JSON  userFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfiguration]
type userFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationTarget string

const (
	UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationTargetIP UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                       `json:"value"`
	JSON  userFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6Configuration]
type userFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationTargetIp6 UserFirewallAccessRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                       `json:"value"`
	JSON  userFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfiguration]
type userFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationTarget string

const (
	UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationTargetIPRange UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target UserFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                      `json:"value"`
	JSON  userFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfiguration]
type userFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationTarget string

const (
	UserFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationTargetASN UserFirewallAccessRuleNewResponseConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                          `json:"value"`
	JSON  userFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfiguration]
type userFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationTarget string

const (
	UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationTargetCountry UserFirewallAccessRuleNewResponseConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleNewResponseMode string

const (
	UserFirewallAccessRuleNewResponseModeBlock            UserFirewallAccessRuleNewResponseMode = "block"
	UserFirewallAccessRuleNewResponseModeChallenge        UserFirewallAccessRuleNewResponseMode = "challenge"
	UserFirewallAccessRuleNewResponseModeWhitelist        UserFirewallAccessRuleNewResponseMode = "whitelist"
	UserFirewallAccessRuleNewResponseModeJsChallenge      UserFirewallAccessRuleNewResponseMode = "js_challenge"
	UserFirewallAccessRuleNewResponseModeManagedChallenge UserFirewallAccessRuleNewResponseMode = "managed_challenge"
)

type UserFirewallAccessRuleListResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []UserFirewallAccessRuleListResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration UserFirewallAccessRuleListResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode UserFirewallAccessRuleListResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                 `json:"notes"`
	JSON  userFirewallAccessRuleListResponseJSON `json:"-"`
}

// userFirewallAccessRuleListResponseJSON contains the JSON metadata for the struct
// [UserFirewallAccessRuleListResponse]
type userFirewallAccessRuleListResponseJSON struct {
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

func (r *UserFirewallAccessRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type UserFirewallAccessRuleListResponseAllowedMode string

const (
	UserFirewallAccessRuleListResponseAllowedModeBlock            UserFirewallAccessRuleListResponseAllowedMode = "block"
	UserFirewallAccessRuleListResponseAllowedModeChallenge        UserFirewallAccessRuleListResponseAllowedMode = "challenge"
	UserFirewallAccessRuleListResponseAllowedModeWhitelist        UserFirewallAccessRuleListResponseAllowedMode = "whitelist"
	UserFirewallAccessRuleListResponseAllowedModeJsChallenge      UserFirewallAccessRuleListResponseAllowedMode = "js_challenge"
	UserFirewallAccessRuleListResponseAllowedModeManagedChallenge UserFirewallAccessRuleListResponseAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfiguration],
// [UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6Configuration],
// [UserFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfiguration],
// [UserFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfiguration] or
// [UserFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfiguration].
type UserFirewallAccessRuleListResponseConfiguration interface {
	implementsUserFirewallAccessRuleListResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserFirewallAccessRuleListResponseConfiguration)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfiguration{}),
		},
	)
}

type UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                      `json:"value"`
	JSON  userFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfiguration]
type userFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationTarget string

const (
	UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationTargetIP UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                        `json:"value"`
	JSON  userFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6Configuration]
type userFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationTargetIp6 UserFirewallAccessRuleListResponseConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target UserFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                        `json:"value"`
	JSON  userFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfiguration]
type userFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationTarget string

const (
	UserFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationTargetIPRange UserFirewallAccessRuleListResponseConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target UserFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                       `json:"value"`
	JSON  userFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfiguration]
type userFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationTarget string

const (
	UserFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationTargetASN UserFirewallAccessRuleListResponseConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target UserFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                           `json:"value"`
	JSON  userFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfiguration]
type userFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationTarget string

const (
	UserFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationTargetCountry UserFirewallAccessRuleListResponseConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleListResponseMode string

const (
	UserFirewallAccessRuleListResponseModeBlock            UserFirewallAccessRuleListResponseMode = "block"
	UserFirewallAccessRuleListResponseModeChallenge        UserFirewallAccessRuleListResponseMode = "challenge"
	UserFirewallAccessRuleListResponseModeWhitelist        UserFirewallAccessRuleListResponseMode = "whitelist"
	UserFirewallAccessRuleListResponseModeJsChallenge      UserFirewallAccessRuleListResponseMode = "js_challenge"
	UserFirewallAccessRuleListResponseModeManagedChallenge UserFirewallAccessRuleListResponseMode = "managed_challenge"
)

type UserFirewallAccessRuleDeleteResponse struct {
	// The unique identifier of the IP Access rule.
	ID   string                                   `json:"id"`
	JSON userFirewallAccessRuleDeleteResponseJSON `json:"-"`
}

// userFirewallAccessRuleDeleteResponseJSON contains the JSON metadata for the
// struct [UserFirewallAccessRuleDeleteResponse]
type userFirewallAccessRuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type UserFirewallAccessRuleEditResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []UserFirewallAccessRuleEditResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration UserFirewallAccessRuleEditResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode UserFirewallAccessRuleEditResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                 `json:"notes"`
	JSON  userFirewallAccessRuleEditResponseJSON `json:"-"`
}

// userFirewallAccessRuleEditResponseJSON contains the JSON metadata for the struct
// [UserFirewallAccessRuleEditResponse]
type userFirewallAccessRuleEditResponseJSON struct {
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

func (r *UserFirewallAccessRuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleEditResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type UserFirewallAccessRuleEditResponseAllowedMode string

const (
	UserFirewallAccessRuleEditResponseAllowedModeBlock            UserFirewallAccessRuleEditResponseAllowedMode = "block"
	UserFirewallAccessRuleEditResponseAllowedModeChallenge        UserFirewallAccessRuleEditResponseAllowedMode = "challenge"
	UserFirewallAccessRuleEditResponseAllowedModeWhitelist        UserFirewallAccessRuleEditResponseAllowedMode = "whitelist"
	UserFirewallAccessRuleEditResponseAllowedModeJsChallenge      UserFirewallAccessRuleEditResponseAllowedMode = "js_challenge"
	UserFirewallAccessRuleEditResponseAllowedModeManagedChallenge UserFirewallAccessRuleEditResponseAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfiguration],
// [UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6Configuration],
// [UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfiguration],
// [UserFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfiguration] or
// [UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfiguration].
type UserFirewallAccessRuleEditResponseConfiguration interface {
	implementsUserFirewallAccessRuleEditResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserFirewallAccessRuleEditResponseConfiguration)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfiguration{}),
		},
	)
}

type UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                      `json:"value"`
	JSON  userFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfiguration]
type userFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationTarget string

const (
	UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationTargetIP UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                        `json:"value"`
	JSON  userFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6Configuration]
type userFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationTargetIp6 UserFirewallAccessRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                        `json:"value"`
	JSON  userFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfiguration]
type userFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationTarget string

const (
	UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationTargetIPRange UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target UserFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                       `json:"value"`
	JSON  userFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfiguration]
type userFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationTarget string

const (
	UserFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationTargetASN UserFirewallAccessRuleEditResponseConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                           `json:"value"`
	JSON  userFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfiguration]
type userFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationTarget string

const (
	UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationTargetCountry UserFirewallAccessRuleEditResponseConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleEditResponseMode string

const (
	UserFirewallAccessRuleEditResponseModeBlock            UserFirewallAccessRuleEditResponseMode = "block"
	UserFirewallAccessRuleEditResponseModeChallenge        UserFirewallAccessRuleEditResponseMode = "challenge"
	UserFirewallAccessRuleEditResponseModeWhitelist        UserFirewallAccessRuleEditResponseMode = "whitelist"
	UserFirewallAccessRuleEditResponseModeJsChallenge      UserFirewallAccessRuleEditResponseMode = "js_challenge"
	UserFirewallAccessRuleEditResponseModeManagedChallenge UserFirewallAccessRuleEditResponseMode = "managed_challenge"
)

type UserFirewallAccessRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[UserFirewallAccessRuleNewParamsConfiguration] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[UserFirewallAccessRuleNewParamsMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r UserFirewallAccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
//
// Satisfied by
// [UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration],
// [UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration],
// [UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfiguration],
// [UserFirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration],
// [UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration].
type UserFirewallAccessRuleNewParamsConfiguration interface {
	implementsUserFirewallAccessRuleNewParamsConfiguration()
}

type UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget string

const (
	UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTargetIP UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTargetIp6 UserFirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfigurationTarget string

const (
	UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfigurationTargetIPRange UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[UserFirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget string

const (
	UserFirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTargetASN UserFirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget string

const (
	UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTargetCountry UserFirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleNewParamsMode string

const (
	UserFirewallAccessRuleNewParamsModeBlock            UserFirewallAccessRuleNewParamsMode = "block"
	UserFirewallAccessRuleNewParamsModeChallenge        UserFirewallAccessRuleNewParamsMode = "challenge"
	UserFirewallAccessRuleNewParamsModeWhitelist        UserFirewallAccessRuleNewParamsMode = "whitelist"
	UserFirewallAccessRuleNewParamsModeJsChallenge      UserFirewallAccessRuleNewParamsMode = "js_challenge"
	UserFirewallAccessRuleNewParamsModeManagedChallenge UserFirewallAccessRuleNewParamsMode = "managed_challenge"
)

type UserFirewallAccessRuleNewResponseEnvelope struct {
	Errors   []UserFirewallAccessRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserFirewallAccessRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   UserFirewallAccessRuleNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success UserFirewallAccessRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    userFirewallAccessRuleNewResponseEnvelopeJSON    `json:"-"`
}

// userFirewallAccessRuleNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserFirewallAccessRuleNewResponseEnvelope]
type userFirewallAccessRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserFirewallAccessRuleNewResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    userFirewallAccessRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// userFirewallAccessRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserFirewallAccessRuleNewResponseEnvelopeErrors]
type userFirewallAccessRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserFirewallAccessRuleNewResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    userFirewallAccessRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// userFirewallAccessRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [UserFirewallAccessRuleNewResponseEnvelopeMessages]
type userFirewallAccessRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserFirewallAccessRuleNewResponseEnvelopeSuccess bool

const (
	UserFirewallAccessRuleNewResponseEnvelopeSuccessTrue UserFirewallAccessRuleNewResponseEnvelopeSuccess = true
)

type UserFirewallAccessRuleListParams struct {
	// The direction used to sort returned rules.
	Direction     param.Field[UserFirewallAccessRuleListParamsDirection]     `query:"direction"`
	EgsPagination param.Field[UserFirewallAccessRuleListParamsEgsPagination] `query:"egs-pagination"`
	Filters       param.Field[UserFirewallAccessRuleListParamsFilters]       `query:"filters"`
	// The field used to sort returned rules.
	Order param.Field[UserFirewallAccessRuleListParamsOrder] `query:"order"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserFirewallAccessRuleListParams]'s query parameters as
// `url.Values`.
func (r UserFirewallAccessRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type UserFirewallAccessRuleListParamsDirection string

const (
	UserFirewallAccessRuleListParamsDirectionAsc  UserFirewallAccessRuleListParamsDirection = "asc"
	UserFirewallAccessRuleListParamsDirectionDesc UserFirewallAccessRuleListParamsDirection = "desc"
)

type UserFirewallAccessRuleListParamsEgsPagination struct {
	Json param.Field[UserFirewallAccessRuleListParamsEgsPaginationJson] `query:"json"`
}

// URLQuery serializes [UserFirewallAccessRuleListParamsEgsPagination]'s query
// parameters as `url.Values`.
func (r UserFirewallAccessRuleListParamsEgsPagination) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserFirewallAccessRuleListParamsEgsPaginationJson struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserFirewallAccessRuleListParamsEgsPaginationJson]'s query
// parameters as `url.Values`.
func (r UserFirewallAccessRuleListParamsEgsPaginationJson) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserFirewallAccessRuleListParamsFilters struct {
	// The target to search in existing rules.
	ConfigurationTarget param.Field[UserFirewallAccessRuleListParamsFiltersConfigurationTarget] `query:"configuration.target"`
	// The target value to search for in existing rules: an IP address, an IP address
	// range, or a country code, depending on the provided `configuration.target`.
	// Notes: You can search for a single IPv4 address, an IP address range with a
	// subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	ConfigurationValue param.Field[string] `query:"configuration.value"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[UserFirewallAccessRuleListParamsFiltersMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[UserFirewallAccessRuleListParamsFiltersMode] `query:"mode"`
	// The string to search for in the notes of existing IP Access rules. Notes: For
	// example, the string 'attack' would match IP Access rules with notes 'Attack
	// 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
}

// URLQuery serializes [UserFirewallAccessRuleListParamsFilters]'s query parameters
// as `url.Values`.
func (r UserFirewallAccessRuleListParamsFilters) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The target to search in existing rules.
type UserFirewallAccessRuleListParamsFiltersConfigurationTarget string

const (
	UserFirewallAccessRuleListParamsFiltersConfigurationTargetIP      UserFirewallAccessRuleListParamsFiltersConfigurationTarget = "ip"
	UserFirewallAccessRuleListParamsFiltersConfigurationTargetIPRange UserFirewallAccessRuleListParamsFiltersConfigurationTarget = "ip_range"
	UserFirewallAccessRuleListParamsFiltersConfigurationTargetASN     UserFirewallAccessRuleListParamsFiltersConfigurationTarget = "asn"
	UserFirewallAccessRuleListParamsFiltersConfigurationTargetCountry UserFirewallAccessRuleListParamsFiltersConfigurationTarget = "country"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type UserFirewallAccessRuleListParamsFiltersMatch string

const (
	UserFirewallAccessRuleListParamsFiltersMatchAny UserFirewallAccessRuleListParamsFiltersMatch = "any"
	UserFirewallAccessRuleListParamsFiltersMatchAll UserFirewallAccessRuleListParamsFiltersMatch = "all"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleListParamsFiltersMode string

const (
	UserFirewallAccessRuleListParamsFiltersModeBlock            UserFirewallAccessRuleListParamsFiltersMode = "block"
	UserFirewallAccessRuleListParamsFiltersModeChallenge        UserFirewallAccessRuleListParamsFiltersMode = "challenge"
	UserFirewallAccessRuleListParamsFiltersModeWhitelist        UserFirewallAccessRuleListParamsFiltersMode = "whitelist"
	UserFirewallAccessRuleListParamsFiltersModeJsChallenge      UserFirewallAccessRuleListParamsFiltersMode = "js_challenge"
	UserFirewallAccessRuleListParamsFiltersModeManagedChallenge UserFirewallAccessRuleListParamsFiltersMode = "managed_challenge"
)

// The field used to sort returned rules.
type UserFirewallAccessRuleListParamsOrder string

const (
	UserFirewallAccessRuleListParamsOrderConfigurationTarget UserFirewallAccessRuleListParamsOrder = "configuration.target"
	UserFirewallAccessRuleListParamsOrderConfigurationValue  UserFirewallAccessRuleListParamsOrder = "configuration.value"
	UserFirewallAccessRuleListParamsOrderMode                UserFirewallAccessRuleListParamsOrder = "mode"
)

type UserFirewallAccessRuleDeleteResponseEnvelope struct {
	Errors   []UserFirewallAccessRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserFirewallAccessRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   UserFirewallAccessRuleDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success UserFirewallAccessRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    userFirewallAccessRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// userFirewallAccessRuleDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserFirewallAccessRuleDeleteResponseEnvelope]
type userFirewallAccessRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserFirewallAccessRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    userFirewallAccessRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// userFirewallAccessRuleDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserFirewallAccessRuleDeleteResponseEnvelopeErrors]
type userFirewallAccessRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserFirewallAccessRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    userFirewallAccessRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// userFirewallAccessRuleDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserFirewallAccessRuleDeleteResponseEnvelopeMessages]
type userFirewallAccessRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserFirewallAccessRuleDeleteResponseEnvelopeSuccess bool

const (
	UserFirewallAccessRuleDeleteResponseEnvelopeSuccessTrue UserFirewallAccessRuleDeleteResponseEnvelopeSuccess = true
)

type UserFirewallAccessRuleEditParams struct {
	// The action to apply to a matched request.
	Mode param.Field[UserFirewallAccessRuleEditParamsMode] `json:"mode"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r UserFirewallAccessRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to apply to a matched request.
type UserFirewallAccessRuleEditParamsMode string

const (
	UserFirewallAccessRuleEditParamsModeBlock            UserFirewallAccessRuleEditParamsMode = "block"
	UserFirewallAccessRuleEditParamsModeChallenge        UserFirewallAccessRuleEditParamsMode = "challenge"
	UserFirewallAccessRuleEditParamsModeWhitelist        UserFirewallAccessRuleEditParamsMode = "whitelist"
	UserFirewallAccessRuleEditParamsModeJsChallenge      UserFirewallAccessRuleEditParamsMode = "js_challenge"
	UserFirewallAccessRuleEditParamsModeManagedChallenge UserFirewallAccessRuleEditParamsMode = "managed_challenge"
)

type UserFirewallAccessRuleEditResponseEnvelope struct {
	Errors   []UserFirewallAccessRuleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserFirewallAccessRuleEditResponseEnvelopeMessages `json:"messages,required"`
	Result   UserFirewallAccessRuleEditResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success UserFirewallAccessRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    userFirewallAccessRuleEditResponseEnvelopeJSON    `json:"-"`
}

// userFirewallAccessRuleEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserFirewallAccessRuleEditResponseEnvelope]
type userFirewallAccessRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserFirewallAccessRuleEditResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    userFirewallAccessRuleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// userFirewallAccessRuleEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserFirewallAccessRuleEditResponseEnvelopeErrors]
type userFirewallAccessRuleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserFirewallAccessRuleEditResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    userFirewallAccessRuleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// userFirewallAccessRuleEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserFirewallAccessRuleEditResponseEnvelopeMessages]
type userFirewallAccessRuleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFirewallAccessRuleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserFirewallAccessRuleEditResponseEnvelopeSuccess bool

const (
	UserFirewallAccessRuleEditResponseEnvelopeSuccessTrue UserFirewallAccessRuleEditResponseEnvelopeSuccess = true
)
