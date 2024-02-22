// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserFirewallAccessRuleRuleService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewUserFirewallAccessRuleRuleService] method instead.
type UserFirewallAccessRuleRuleService struct {
	Options []option.RequestOption
}

// NewUserFirewallAccessRuleRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserFirewallAccessRuleRuleService(opts ...option.RequestOption) (r *UserFirewallAccessRuleRuleService) {
	r = &UserFirewallAccessRuleRuleService{}
	r.Options = opts
	return
}

// Creates a new IP Access rule for all zones owned by the current user.
//
// Note: To create an IP Access rule that applies to a specific zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *UserFirewallAccessRuleRuleService) New(ctx context.Context, body UserFirewallAccessRuleRuleNewParams, opts ...option.RequestOption) (res *UserFirewallAccessRuleRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserFirewallAccessRuleRuleNewResponseEnvelope
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
func (r *UserFirewallAccessRuleRuleService) List(ctx context.Context, query UserFirewallAccessRuleRuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[UserFirewallAccessRuleRuleListResponse], err error) {
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
func (r *UserFirewallAccessRuleRuleService) ListAutoPaging(ctx context.Context, query UserFirewallAccessRuleRuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[UserFirewallAccessRuleRuleListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Deletes an IP Access rule at the user level.
//
// Note: Deleting a user-level rule will affect all zones owned by the user.
func (r *UserFirewallAccessRuleRuleService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserFirewallAccessRuleRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserFirewallAccessRuleRuleDeleteResponseEnvelope
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
func (r *UserFirewallAccessRuleRuleService) Edit(ctx context.Context, identifier string, body UserFirewallAccessRuleRuleEditParams, opts ...option.RequestOption) (res *UserFirewallAccessRuleRuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserFirewallAccessRuleRuleEditResponseEnvelope
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserFirewallAccessRuleRuleNewResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []UserFirewallAccessRuleRuleNewResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration UserFirewallAccessRuleRuleNewResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode UserFirewallAccessRuleRuleNewResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                    `json:"notes"`
	JSON  userFirewallAccessRuleRuleNewResponseJSON `json:"-"`
}

// userFirewallAccessRuleRuleNewResponseJSON contains the JSON metadata for the
// struct [UserFirewallAccessRuleRuleNewResponse]
type userFirewallAccessRuleRuleNewResponseJSON struct {
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

func (r *UserFirewallAccessRuleRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleNewResponseAllowedMode string

const (
	UserFirewallAccessRuleRuleNewResponseAllowedModeBlock            UserFirewallAccessRuleRuleNewResponseAllowedMode = "block"
	UserFirewallAccessRuleRuleNewResponseAllowedModeChallenge        UserFirewallAccessRuleRuleNewResponseAllowedMode = "challenge"
	UserFirewallAccessRuleRuleNewResponseAllowedModeWhitelist        UserFirewallAccessRuleRuleNewResponseAllowedMode = "whitelist"
	UserFirewallAccessRuleRuleNewResponseAllowedModeJsChallenge      UserFirewallAccessRuleRuleNewResponseAllowedMode = "js_challenge"
	UserFirewallAccessRuleRuleNewResponseAllowedModeManagedChallenge UserFirewallAccessRuleRuleNewResponseAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPConfiguration],
// [UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPV6Configuration],
// [UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCidrConfiguration],
// [UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsASNConfiguration] or
// [UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCountryConfiguration].
type UserFirewallAccessRuleRuleNewResponseConfiguration interface {
	implementsUserFirewallAccessRuleRuleNewResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*UserFirewallAccessRuleRuleNewResponseConfiguration)(nil)).Elem(), "")
}

type UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                         `json:"value"`
	JSON  userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPConfiguration]
type userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPConfigurationTargetIP UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                           `json:"value"`
	JSON  userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPV6Configuration]
type userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationTargetIp6 UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                           `json:"value"`
	JSON  userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCidrConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCidrConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCidrConfiguration]
type userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                          `json:"value"`
	JSON  userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsASNConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsASNConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsASNConfiguration]
type userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsASNConfigurationTarget string

const (
	UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsASNConfigurationTargetASN UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                              `json:"value"`
	JSON  userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCountryConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCountryConfiguration]
type userFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleRuleNewResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCountryConfigurationTargetCountry UserFirewallAccessRuleRuleNewResponseConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleNewResponseMode string

const (
	UserFirewallAccessRuleRuleNewResponseModeBlock            UserFirewallAccessRuleRuleNewResponseMode = "block"
	UserFirewallAccessRuleRuleNewResponseModeChallenge        UserFirewallAccessRuleRuleNewResponseMode = "challenge"
	UserFirewallAccessRuleRuleNewResponseModeWhitelist        UserFirewallAccessRuleRuleNewResponseMode = "whitelist"
	UserFirewallAccessRuleRuleNewResponseModeJsChallenge      UserFirewallAccessRuleRuleNewResponseMode = "js_challenge"
	UserFirewallAccessRuleRuleNewResponseModeManagedChallenge UserFirewallAccessRuleRuleNewResponseMode = "managed_challenge"
)

type UserFirewallAccessRuleRuleListResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []UserFirewallAccessRuleRuleListResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration UserFirewallAccessRuleRuleListResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode UserFirewallAccessRuleRuleListResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                     `json:"notes"`
	JSON  userFirewallAccessRuleRuleListResponseJSON `json:"-"`
}

// userFirewallAccessRuleRuleListResponseJSON contains the JSON metadata for the
// struct [UserFirewallAccessRuleRuleListResponse]
type userFirewallAccessRuleRuleListResponseJSON struct {
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

func (r *UserFirewallAccessRuleRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleListResponseAllowedMode string

const (
	UserFirewallAccessRuleRuleListResponseAllowedModeBlock            UserFirewallAccessRuleRuleListResponseAllowedMode = "block"
	UserFirewallAccessRuleRuleListResponseAllowedModeChallenge        UserFirewallAccessRuleRuleListResponseAllowedMode = "challenge"
	UserFirewallAccessRuleRuleListResponseAllowedModeWhitelist        UserFirewallAccessRuleRuleListResponseAllowedMode = "whitelist"
	UserFirewallAccessRuleRuleListResponseAllowedModeJsChallenge      UserFirewallAccessRuleRuleListResponseAllowedMode = "js_challenge"
	UserFirewallAccessRuleRuleListResponseAllowedModeManagedChallenge UserFirewallAccessRuleRuleListResponseAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPConfiguration],
// [UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPV6Configuration],
// [UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCidrConfiguration],
// [UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsASNConfiguration]
// or
// [UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCountryConfiguration].
type UserFirewallAccessRuleRuleListResponseConfiguration interface {
	implementsUserFirewallAccessRuleRuleListResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*UserFirewallAccessRuleRuleListResponseConfiguration)(nil)).Elem(), "")
}

type UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                          `json:"value"`
	JSON  userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPConfiguration]
type userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPConfigurationTargetIP UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                            `json:"value"`
	JSON  userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPV6Configuration]
type userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPV6ConfigurationTargetIp6 UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                            `json:"value"`
	JSON  userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCidrConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCidrConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCidrConfiguration]
type userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                           `json:"value"`
	JSON  userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsASNConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsASNConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsASNConfiguration]
type userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsASNConfigurationTarget string

const (
	UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsASNConfigurationTargetASN UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                               `json:"value"`
	JSON  userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCountryConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCountryConfiguration]
type userFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleRuleListResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCountryConfigurationTargetCountry UserFirewallAccessRuleRuleListResponseConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleListResponseMode string

const (
	UserFirewallAccessRuleRuleListResponseModeBlock            UserFirewallAccessRuleRuleListResponseMode = "block"
	UserFirewallAccessRuleRuleListResponseModeChallenge        UserFirewallAccessRuleRuleListResponseMode = "challenge"
	UserFirewallAccessRuleRuleListResponseModeWhitelist        UserFirewallAccessRuleRuleListResponseMode = "whitelist"
	UserFirewallAccessRuleRuleListResponseModeJsChallenge      UserFirewallAccessRuleRuleListResponseMode = "js_challenge"
	UserFirewallAccessRuleRuleListResponseModeManagedChallenge UserFirewallAccessRuleRuleListResponseMode = "managed_challenge"
)

type UserFirewallAccessRuleRuleDeleteResponse struct {
	// The unique identifier of the IP Access rule.
	ID   string                                       `json:"id"`
	JSON userFirewallAccessRuleRuleDeleteResponseJSON `json:"-"`
}

// userFirewallAccessRuleRuleDeleteResponseJSON contains the JSON metadata for the
// struct [UserFirewallAccessRuleRuleDeleteResponse]
type userFirewallAccessRuleRuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleEditResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []UserFirewallAccessRuleRuleEditResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration UserFirewallAccessRuleRuleEditResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode UserFirewallAccessRuleRuleEditResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                     `json:"notes"`
	JSON  userFirewallAccessRuleRuleEditResponseJSON `json:"-"`
}

// userFirewallAccessRuleRuleEditResponseJSON contains the JSON metadata for the
// struct [UserFirewallAccessRuleRuleEditResponse]
type userFirewallAccessRuleRuleEditResponseJSON struct {
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

func (r *UserFirewallAccessRuleRuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleEditResponseAllowedMode string

const (
	UserFirewallAccessRuleRuleEditResponseAllowedModeBlock            UserFirewallAccessRuleRuleEditResponseAllowedMode = "block"
	UserFirewallAccessRuleRuleEditResponseAllowedModeChallenge        UserFirewallAccessRuleRuleEditResponseAllowedMode = "challenge"
	UserFirewallAccessRuleRuleEditResponseAllowedModeWhitelist        UserFirewallAccessRuleRuleEditResponseAllowedMode = "whitelist"
	UserFirewallAccessRuleRuleEditResponseAllowedModeJsChallenge      UserFirewallAccessRuleRuleEditResponseAllowedMode = "js_challenge"
	UserFirewallAccessRuleRuleEditResponseAllowedModeManagedChallenge UserFirewallAccessRuleRuleEditResponseAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPConfiguration],
// [UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPV6Configuration],
// [UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCidrConfiguration],
// [UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsASNConfiguration]
// or
// [UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCountryConfiguration].
type UserFirewallAccessRuleRuleEditResponseConfiguration interface {
	implementsUserFirewallAccessRuleRuleEditResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*UserFirewallAccessRuleRuleEditResponseConfiguration)(nil)).Elem(), "")
}

type UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                          `json:"value"`
	JSON  userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPConfiguration]
type userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPConfigurationTargetIP UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                            `json:"value"`
	JSON  userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPV6Configuration]
type userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationTargetIp6 UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                            `json:"value"`
	JSON  userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCidrConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCidrConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCidrConfiguration]
type userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                           `json:"value"`
	JSON  userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsASNConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsASNConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsASNConfiguration]
type userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsASNConfigurationTarget string

const (
	UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsASNConfigurationTargetASN UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                               `json:"value"`
	JSON  userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCountryConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCountryConfiguration]
type userFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleRuleEditResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCountryConfigurationTargetCountry UserFirewallAccessRuleRuleEditResponseConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleEditResponseMode string

const (
	UserFirewallAccessRuleRuleEditResponseModeBlock            UserFirewallAccessRuleRuleEditResponseMode = "block"
	UserFirewallAccessRuleRuleEditResponseModeChallenge        UserFirewallAccessRuleRuleEditResponseMode = "challenge"
	UserFirewallAccessRuleRuleEditResponseModeWhitelist        UserFirewallAccessRuleRuleEditResponseMode = "whitelist"
	UserFirewallAccessRuleRuleEditResponseModeJsChallenge      UserFirewallAccessRuleRuleEditResponseMode = "js_challenge"
	UserFirewallAccessRuleRuleEditResponseModeManagedChallenge UserFirewallAccessRuleRuleEditResponseMode = "managed_challenge"
)

type UserFirewallAccessRuleRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[UserFirewallAccessRuleRuleNewParamsConfiguration] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[UserFirewallAccessRuleRuleNewParamsMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r UserFirewallAccessRuleRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
//
// Satisfied by
// [UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPConfiguration],
// [UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPV6Configuration],
// [UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCidrConfiguration],
// [UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsASNConfiguration],
// [UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCountryConfiguration].
type UserFirewallAccessRuleRuleNewParamsConfiguration interface {
	implementsUserFirewallAccessRuleRuleNewParamsConfiguration()
}

type UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPConfigurationTargetIP UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPV6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTargetIp6 UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget string

const (
	UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsASNConfigurationTargetASN UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCountryConfigurationTargetCountry UserFirewallAccessRuleRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleNewParamsMode string

const (
	UserFirewallAccessRuleRuleNewParamsModeBlock            UserFirewallAccessRuleRuleNewParamsMode = "block"
	UserFirewallAccessRuleRuleNewParamsModeChallenge        UserFirewallAccessRuleRuleNewParamsMode = "challenge"
	UserFirewallAccessRuleRuleNewParamsModeWhitelist        UserFirewallAccessRuleRuleNewParamsMode = "whitelist"
	UserFirewallAccessRuleRuleNewParamsModeJsChallenge      UserFirewallAccessRuleRuleNewParamsMode = "js_challenge"
	UserFirewallAccessRuleRuleNewParamsModeManagedChallenge UserFirewallAccessRuleRuleNewParamsMode = "managed_challenge"
)

type UserFirewallAccessRuleRuleNewResponseEnvelope struct {
	Errors   []UserFirewallAccessRuleRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserFirewallAccessRuleRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   UserFirewallAccessRuleRuleNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success UserFirewallAccessRuleRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    userFirewallAccessRuleRuleNewResponseEnvelopeJSON    `json:"-"`
}

// userFirewallAccessRuleRuleNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserFirewallAccessRuleRuleNewResponseEnvelope]
type userFirewallAccessRuleRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleNewResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    userFirewallAccessRuleRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// userFirewallAccessRuleRuleNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserFirewallAccessRuleRuleNewResponseEnvelopeErrors]
type userFirewallAccessRuleRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleNewResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    userFirewallAccessRuleRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// userFirewallAccessRuleRuleNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserFirewallAccessRuleRuleNewResponseEnvelopeMessages]
type userFirewallAccessRuleRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserFirewallAccessRuleRuleNewResponseEnvelopeSuccess bool

const (
	UserFirewallAccessRuleRuleNewResponseEnvelopeSuccessTrue UserFirewallAccessRuleRuleNewResponseEnvelopeSuccess = true
)

type UserFirewallAccessRuleRuleListParams struct {
	// The direction used to sort returned rules.
	Direction     param.Field[UserFirewallAccessRuleRuleListParamsDirection]     `query:"direction"`
	EgsPagination param.Field[UserFirewallAccessRuleRuleListParamsEgsPagination] `query:"egs-pagination"`
	Filters       param.Field[UserFirewallAccessRuleRuleListParamsFilters]       `query:"filters"`
	// The field used to sort returned rules.
	Order param.Field[UserFirewallAccessRuleRuleListParamsOrder] `query:"order"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserFirewallAccessRuleRuleListParams]'s query parameters as
// `url.Values`.
func (r UserFirewallAccessRuleRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type UserFirewallAccessRuleRuleListParamsDirection string

const (
	UserFirewallAccessRuleRuleListParamsDirectionAsc  UserFirewallAccessRuleRuleListParamsDirection = "asc"
	UserFirewallAccessRuleRuleListParamsDirectionDesc UserFirewallAccessRuleRuleListParamsDirection = "desc"
)

type UserFirewallAccessRuleRuleListParamsEgsPagination struct {
	Json param.Field[UserFirewallAccessRuleRuleListParamsEgsPaginationJson] `query:"json"`
}

// URLQuery serializes [UserFirewallAccessRuleRuleListParamsEgsPagination]'s query
// parameters as `url.Values`.
func (r UserFirewallAccessRuleRuleListParamsEgsPagination) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserFirewallAccessRuleRuleListParamsEgsPaginationJson struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserFirewallAccessRuleRuleListParamsEgsPaginationJson]'s
// query parameters as `url.Values`.
func (r UserFirewallAccessRuleRuleListParamsEgsPaginationJson) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserFirewallAccessRuleRuleListParamsFilters struct {
	// The target to search in existing rules.
	ConfigurationTarget param.Field[UserFirewallAccessRuleRuleListParamsFiltersConfigurationTarget] `query:"configuration.target"`
	// The target value to search for in existing rules: an IP address, an IP address
	// range, or a country code, depending on the provided `configuration.target`.
	// Notes: You can search for a single IPv4 address, an IP address range with a
	// subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	ConfigurationValue param.Field[string] `query:"configuration.value"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[UserFirewallAccessRuleRuleListParamsFiltersMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[UserFirewallAccessRuleRuleListParamsFiltersMode] `query:"mode"`
	// The string to search for in the notes of existing IP Access rules. Notes: For
	// example, the string 'attack' would match IP Access rules with notes 'Attack
	// 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
}

// URLQuery serializes [UserFirewallAccessRuleRuleListParamsFilters]'s query
// parameters as `url.Values`.
func (r UserFirewallAccessRuleRuleListParamsFilters) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The target to search in existing rules.
type UserFirewallAccessRuleRuleListParamsFiltersConfigurationTarget string

const (
	UserFirewallAccessRuleRuleListParamsFiltersConfigurationTargetIP      UserFirewallAccessRuleRuleListParamsFiltersConfigurationTarget = "ip"
	UserFirewallAccessRuleRuleListParamsFiltersConfigurationTargetIPRange UserFirewallAccessRuleRuleListParamsFiltersConfigurationTarget = "ip_range"
	UserFirewallAccessRuleRuleListParamsFiltersConfigurationTargetASN     UserFirewallAccessRuleRuleListParamsFiltersConfigurationTarget = "asn"
	UserFirewallAccessRuleRuleListParamsFiltersConfigurationTargetCountry UserFirewallAccessRuleRuleListParamsFiltersConfigurationTarget = "country"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type UserFirewallAccessRuleRuleListParamsFiltersMatch string

const (
	UserFirewallAccessRuleRuleListParamsFiltersMatchAny UserFirewallAccessRuleRuleListParamsFiltersMatch = "any"
	UserFirewallAccessRuleRuleListParamsFiltersMatchAll UserFirewallAccessRuleRuleListParamsFiltersMatch = "all"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleListParamsFiltersMode string

const (
	UserFirewallAccessRuleRuleListParamsFiltersModeBlock            UserFirewallAccessRuleRuleListParamsFiltersMode = "block"
	UserFirewallAccessRuleRuleListParamsFiltersModeChallenge        UserFirewallAccessRuleRuleListParamsFiltersMode = "challenge"
	UserFirewallAccessRuleRuleListParamsFiltersModeWhitelist        UserFirewallAccessRuleRuleListParamsFiltersMode = "whitelist"
	UserFirewallAccessRuleRuleListParamsFiltersModeJsChallenge      UserFirewallAccessRuleRuleListParamsFiltersMode = "js_challenge"
	UserFirewallAccessRuleRuleListParamsFiltersModeManagedChallenge UserFirewallAccessRuleRuleListParamsFiltersMode = "managed_challenge"
)

// The field used to sort returned rules.
type UserFirewallAccessRuleRuleListParamsOrder string

const (
	UserFirewallAccessRuleRuleListParamsOrderConfigurationTarget UserFirewallAccessRuleRuleListParamsOrder = "configuration.target"
	UserFirewallAccessRuleRuleListParamsOrderConfigurationValue  UserFirewallAccessRuleRuleListParamsOrder = "configuration.value"
	UserFirewallAccessRuleRuleListParamsOrderMode                UserFirewallAccessRuleRuleListParamsOrder = "mode"
)

type UserFirewallAccessRuleRuleDeleteResponseEnvelope struct {
	Errors   []UserFirewallAccessRuleRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserFirewallAccessRuleRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   UserFirewallAccessRuleRuleDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success UserFirewallAccessRuleRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    userFirewallAccessRuleRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// userFirewallAccessRuleRuleDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [UserFirewallAccessRuleRuleDeleteResponseEnvelope]
type userFirewallAccessRuleRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    userFirewallAccessRuleRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// userFirewallAccessRuleRuleDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserFirewallAccessRuleRuleDeleteResponseEnvelopeErrors]
type userFirewallAccessRuleRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    userFirewallAccessRuleRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// userFirewallAccessRuleRuleDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [UserFirewallAccessRuleRuleDeleteResponseEnvelopeMessages]
type userFirewallAccessRuleRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserFirewallAccessRuleRuleDeleteResponseEnvelopeSuccess bool

const (
	UserFirewallAccessRuleRuleDeleteResponseEnvelopeSuccessTrue UserFirewallAccessRuleRuleDeleteResponseEnvelopeSuccess = true
)

type UserFirewallAccessRuleRuleEditParams struct {
	// The action to apply to a matched request.
	Mode param.Field[UserFirewallAccessRuleRuleEditParamsMode] `json:"mode"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r UserFirewallAccessRuleRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleEditParamsMode string

const (
	UserFirewallAccessRuleRuleEditParamsModeBlock            UserFirewallAccessRuleRuleEditParamsMode = "block"
	UserFirewallAccessRuleRuleEditParamsModeChallenge        UserFirewallAccessRuleRuleEditParamsMode = "challenge"
	UserFirewallAccessRuleRuleEditParamsModeWhitelist        UserFirewallAccessRuleRuleEditParamsMode = "whitelist"
	UserFirewallAccessRuleRuleEditParamsModeJsChallenge      UserFirewallAccessRuleRuleEditParamsMode = "js_challenge"
	UserFirewallAccessRuleRuleEditParamsModeManagedChallenge UserFirewallAccessRuleRuleEditParamsMode = "managed_challenge"
)

type UserFirewallAccessRuleRuleEditResponseEnvelope struct {
	Errors   []UserFirewallAccessRuleRuleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserFirewallAccessRuleRuleEditResponseEnvelopeMessages `json:"messages,required"`
	Result   UserFirewallAccessRuleRuleEditResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success UserFirewallAccessRuleRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    userFirewallAccessRuleRuleEditResponseEnvelopeJSON    `json:"-"`
}

// userFirewallAccessRuleRuleEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [UserFirewallAccessRuleRuleEditResponseEnvelope]
type userFirewallAccessRuleRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleEditResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    userFirewallAccessRuleRuleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// userFirewallAccessRuleRuleEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserFirewallAccessRuleRuleEditResponseEnvelopeErrors]
type userFirewallAccessRuleRuleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleEditResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    userFirewallAccessRuleRuleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// userFirewallAccessRuleRuleEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserFirewallAccessRuleRuleEditResponseEnvelopeMessages]
type userFirewallAccessRuleRuleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserFirewallAccessRuleRuleEditResponseEnvelopeSuccess bool

const (
	UserFirewallAccessRuleRuleEditResponseEnvelopeSuccessTrue UserFirewallAccessRuleRuleEditResponseEnvelopeSuccess = true
)
