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

// Updates an IP Access rule defined at the user level. You can only update the
// rule action (`mode` parameter) and notes.
func (r *UserFirewallAccessRuleRuleService) Update(ctx context.Context, identifier string, body UserFirewallAccessRuleRuleUpdateParams, opts ...option.RequestOption) (res *UserFirewallAccessRuleRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserFirewallAccessRuleRuleUpdateResponseEnvelope
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

// Creates a new IP Access rule for all zones owned by the current user.
//
// Note: To create an IP Access rule that applies to a specific zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *UserFirewallAccessRuleRuleService) IPAccessRulesForAUserNewAnIPAccessRule(ctx context.Context, body UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParams, opts ...option.RequestOption) (res *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelope
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
func (r *UserFirewallAccessRuleRuleService) IPAccessRulesForAUserListIPAccessRules(ctx context.Context, query UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParams, opts ...option.RequestOption) (res *[]UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelope
	path := "user/firewall/access_rules/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserFirewallAccessRuleRuleUpdateResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []UserFirewallAccessRuleRuleUpdateResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration UserFirewallAccessRuleRuleUpdateResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode UserFirewallAccessRuleRuleUpdateResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                       `json:"notes"`
	JSON  userFirewallAccessRuleRuleUpdateResponseJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseJSON contains the JSON metadata for the
// struct [UserFirewallAccessRuleRuleUpdateResponse]
type userFirewallAccessRuleRuleUpdateResponseJSON struct {
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

func (r *UserFirewallAccessRuleRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleUpdateResponseAllowedMode string

const (
	UserFirewallAccessRuleRuleUpdateResponseAllowedModeBlock            UserFirewallAccessRuleRuleUpdateResponseAllowedMode = "block"
	UserFirewallAccessRuleRuleUpdateResponseAllowedModeChallenge        UserFirewallAccessRuleRuleUpdateResponseAllowedMode = "challenge"
	UserFirewallAccessRuleRuleUpdateResponseAllowedModeWhitelist        UserFirewallAccessRuleRuleUpdateResponseAllowedMode = "whitelist"
	UserFirewallAccessRuleRuleUpdateResponseAllowedModeJsChallenge      UserFirewallAccessRuleRuleUpdateResponseAllowedMode = "js_challenge"
	UserFirewallAccessRuleRuleUpdateResponseAllowedModeManagedChallenge UserFirewallAccessRuleRuleUpdateResponseAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPConfiguration],
// [UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPV6Configuration],
// [UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCidrConfiguration],
// [UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsAsnConfiguration]
// or
// [UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCountryConfiguration].
type UserFirewallAccessRuleRuleUpdateResponseConfiguration interface {
	implementsUserFirewallAccessRuleRuleUpdateResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*UserFirewallAccessRuleRuleUpdateResponseConfiguration)(nil)).Elem(), "")
}

type UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                            `json:"value"`
	JSON  userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPConfiguration]
type userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleRuleUpdateResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPConfigurationTargetIP UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                              `json:"value"`
	JSON  userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPV6Configuration]
type userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleRuleUpdateResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPV6ConfigurationTargetIp6 UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                              `json:"value"`
	JSON  userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCidrConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCidrConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCidrConfiguration]
type userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleRuleUpdateResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsAsnConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsAsnConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                             `json:"value"`
	JSON  userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsAsnConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsAsnConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsAsnConfiguration]
type userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsAsnConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsAsnConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsAsnConfiguration) implementsUserFirewallAccessRuleRuleUpdateResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsAsnConfigurationTarget string

const (
	UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsAsnConfigurationTargetAsn UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsAsnConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                                 `json:"value"`
	JSON  userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCountryConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCountryConfiguration]
type userFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleRuleUpdateResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCountryConfigurationTargetCountry UserFirewallAccessRuleRuleUpdateResponseConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleUpdateResponseMode string

const (
	UserFirewallAccessRuleRuleUpdateResponseModeBlock            UserFirewallAccessRuleRuleUpdateResponseMode = "block"
	UserFirewallAccessRuleRuleUpdateResponseModeChallenge        UserFirewallAccessRuleRuleUpdateResponseMode = "challenge"
	UserFirewallAccessRuleRuleUpdateResponseModeWhitelist        UserFirewallAccessRuleRuleUpdateResponseMode = "whitelist"
	UserFirewallAccessRuleRuleUpdateResponseModeJsChallenge      UserFirewallAccessRuleRuleUpdateResponseMode = "js_challenge"
	UserFirewallAccessRuleRuleUpdateResponseModeManagedChallenge UserFirewallAccessRuleRuleUpdateResponseMode = "managed_challenge"
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

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                                                       `json:"notes"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponse]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseJSON struct {
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

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseAllowedMode string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseAllowedModeBlock            UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseAllowedMode = "block"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseAllowedModeChallenge        UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseAllowedMode = "challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseAllowedModeWhitelist        UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseAllowedMode = "whitelist"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseAllowedModeJsChallenge      UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseAllowedMode = "js_challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseAllowedModeManagedChallenge UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPV6Configuration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCidrConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsAsnConfiguration]
// or
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCountryConfiguration].
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfiguration interface {
	implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfiguration)(nil)).Elem(), "")
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                                            `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPConfigurationTargetIP UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                                                              `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPV6Configuration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPV6ConfigurationTargetIp6 UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                                                              `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCidrConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCidrConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCidrConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsAsnConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsAsnConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                                                             `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsAsnConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsAsnConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsAsnConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsAsnConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsAsnConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsAsnConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsAsnConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsAsnConfigurationTargetAsn UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsAsnConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                                                                 `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCountryConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCountryConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCountryConfigurationTargetCountry UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMode string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseModeBlock            UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMode = "block"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseModeChallenge        UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMode = "challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseModeWhitelist        UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMode = "whitelist"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseModeJsChallenge      UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMode = "js_challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseModeManagedChallenge UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMode = "managed_challenge"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                                                       `json:"notes"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponse]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseJSON struct {
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

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseAllowedMode string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseAllowedModeBlock            UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseAllowedMode = "block"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseAllowedModeChallenge        UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseAllowedMode = "challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseAllowedModeWhitelist        UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseAllowedMode = "whitelist"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseAllowedModeJsChallenge      UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseAllowedMode = "js_challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseAllowedModeManagedChallenge UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPV6Configuration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCidrConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsAsnConfiguration]
// or
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCountryConfiguration].
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration interface {
	implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration)(nil)).Elem(), "")
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                                            `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPConfigurationTargetIP UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                                                              `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPV6ConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPV6Configuration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPV6ConfigurationTargetIp6 UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                                                              `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCidrConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCidrConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCidrConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsAsnConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsAsnConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                                                             `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsAsnConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsAsnConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsAsnConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsAsnConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsAsnConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsAsnConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsAsnConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsAsnConfigurationTargetAsn UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsAsnConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                                                                 `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCountryConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCountryConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCountryConfigurationTargetCountry UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseMode string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseModeBlock            UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseMode = "block"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseModeChallenge        UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseMode = "challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseModeWhitelist        UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseMode = "whitelist"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseModeJsChallenge      UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseMode = "js_challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseModeManagedChallenge UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseMode = "managed_challenge"
)

type UserFirewallAccessRuleRuleUpdateParams struct {
	// The action to apply to a matched request.
	Mode param.Field[UserFirewallAccessRuleRuleUpdateParamsMode] `json:"mode"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r UserFirewallAccessRuleRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleUpdateParamsMode string

const (
	UserFirewallAccessRuleRuleUpdateParamsModeBlock            UserFirewallAccessRuleRuleUpdateParamsMode = "block"
	UserFirewallAccessRuleRuleUpdateParamsModeChallenge        UserFirewallAccessRuleRuleUpdateParamsMode = "challenge"
	UserFirewallAccessRuleRuleUpdateParamsModeWhitelist        UserFirewallAccessRuleRuleUpdateParamsMode = "whitelist"
	UserFirewallAccessRuleRuleUpdateParamsModeJsChallenge      UserFirewallAccessRuleRuleUpdateParamsMode = "js_challenge"
	UserFirewallAccessRuleRuleUpdateParamsModeManagedChallenge UserFirewallAccessRuleRuleUpdateParamsMode = "managed_challenge"
)

type UserFirewallAccessRuleRuleUpdateResponseEnvelope struct {
	Errors   []UserFirewallAccessRuleRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserFirewallAccessRuleRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   UserFirewallAccessRuleRuleUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success UserFirewallAccessRuleRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    userFirewallAccessRuleRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [UserFirewallAccessRuleRuleUpdateResponseEnvelope]
type userFirewallAccessRuleRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    userFirewallAccessRuleRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserFirewallAccessRuleRuleUpdateResponseEnvelopeErrors]
type userFirewallAccessRuleRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    userFirewallAccessRuleRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [UserFirewallAccessRuleRuleUpdateResponseEnvelopeMessages]
type userFirewallAccessRuleRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserFirewallAccessRuleRuleUpdateResponseEnvelopeSuccess bool

const (
	UserFirewallAccessRuleRuleUpdateResponseEnvelopeSuccessTrue UserFirewallAccessRuleRuleUpdateResponseEnvelopeSuccess = true
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

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParams struct {
	// The rule configuration.
	Configuration param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
//
// Satisfied by
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPV6Configuration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCidrConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsAsnConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCountryConfiguration].
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration interface {
	implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration()
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPConfigurationTargetIP UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPV6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPV6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPV6ConfigurationTargetIp6 UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCidrConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsAsnConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsAsnConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsAsnConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsAsnConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsAsnConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsAsnConfigurationTargetAsn UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsAsnConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCountryConfigurationTargetCountry UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsMode string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsModeBlock            UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsMode = "block"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsModeChallenge        UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsMode = "challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsModeWhitelist        UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsMode = "whitelist"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsModeJsChallenge      UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsMode = "js_challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsModeManagedChallenge UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsMode = "managed_challenge"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelope struct {
	Errors   []UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeMessages `json:"messages,required"`
	Result   UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeSuccess `json:"success,required"`
	JSON    userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeJSON    `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelope]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeErrors struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeErrors]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeMessages struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeMessages]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeSuccess bool

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeSuccessTrue UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseEnvelopeSuccess = true
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParams struct {
	// The direction used to sort returned rules.
	Direction     param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsDirection]     `query:"direction"`
	EgsPagination param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsEgsPagination] `query:"egs-pagination"`
	Filters       param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFilters]       `query:"filters"`
	// The field used to sort returned rules.
	Order param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsOrder] `query:"order"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParams]'s query
// parameters as `url.Values`.
func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsDirection string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsDirectionAsc  UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsDirection = "asc"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsDirectionDesc UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsDirection = "desc"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsEgsPagination struct {
	Json param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsEgsPaginationJson] `query:"json"`
}

// URLQuery serializes
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsEgsPagination]'s
// query parameters as `url.Values`.
func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsEgsPagination) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsEgsPaginationJson struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsEgsPaginationJson]'s
// query parameters as `url.Values`.
func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsEgsPaginationJson) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFilters struct {
	// The target to search in existing rules.
	ConfigurationTarget param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTarget] `query:"configuration.target"`
	// The target value to search for in existing rules: an IP address, an IP address
	// range, or a country code, depending on the provided `configuration.target`.
	// Notes: You can search for a single IPv4 address, an IP address range with a
	// subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	ConfigurationValue param.Field[string] `query:"configuration.value"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMode] `query:"mode"`
	// The string to search for in the notes of existing IP Access rules. Notes: For
	// example, the string 'attack' would match IP Access rules with notes 'Attack
	// 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
}

// URLQuery serializes
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFilters]'s
// query parameters as `url.Values`.
func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFilters) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The target to search in existing rules.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTargetIP      UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTarget = "ip"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTargetIPRange UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTarget = "ip_range"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTargetAsn     UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTarget = "asn"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTargetCountry UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTarget = "country"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMatch string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMatchAny UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMatch = "any"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMatchAll UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMatch = "all"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMode string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersModeBlock            UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMode = "block"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersModeChallenge        UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMode = "challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersModeWhitelist        UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMode = "whitelist"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersModeJsChallenge      UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMode = "js_challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersModeManagedChallenge UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersMode = "managed_challenge"
)

// The field used to sort returned rules.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsOrder string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsOrderConfigurationTarget UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsOrder = "configuration.target"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsOrderConfigurationValue  UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsOrder = "configuration.value"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsOrderMode                UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsOrder = "mode"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelope struct {
	Errors   []UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeJSON       `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelope]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeErrors struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeErrors]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeMessages struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeMessages]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeSuccess bool

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeSuccessTrue UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeSuccess = true
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                        `json:"total_count"`
	JSON       userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeResultInfo]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
