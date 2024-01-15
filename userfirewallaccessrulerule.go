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

// Updates an IP Access rule defined at the user level. You can only update the
// rule action (`mode` parameter) and notes.
func (r *UserFirewallAccessRuleRuleService) Update(ctx context.Context, identifier string, body UserFirewallAccessRuleRuleUpdateParams, opts ...option.RequestOption) (res *UserFirewallAccessRuleRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an IP Access rule at the user level.
//
// Note: Deleting a user-level rule will affect all zones owned by the user.
func (r *UserFirewallAccessRuleRuleService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserFirewallAccessRuleRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new IP Access rule for all zones owned by the current user.
//
// Note: To create an IP Access rule that applies to a specific zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *UserFirewallAccessRuleRuleService) IPAccessRulesForAUserNewAnIPAccessRule(ctx context.Context, body UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParams, opts ...option.RequestOption) (res *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/firewall/access_rules/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches IP Access rules of the user. You can filter the results using several
// optional parameters.
func (r *UserFirewallAccessRuleRuleService) IPAccessRulesForAUserListIPAccessRules(ctx context.Context, query UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParams, opts ...option.RequestOption) (res *shared.Page[UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponse], err error) {
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

type UserFirewallAccessRuleRuleUpdateResponse struct {
	Errors   []UserFirewallAccessRuleRuleUpdateResponseError   `json:"errors"`
	Messages []UserFirewallAccessRuleRuleUpdateResponseMessage `json:"messages"`
	Result   UserFirewallAccessRuleRuleUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success UserFirewallAccessRuleRuleUpdateResponseSuccess `json:"success"`
	JSON    userFirewallAccessRuleRuleUpdateResponseJSON    `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseJSON contains the JSON metadata for the
// struct [UserFirewallAccessRuleRuleUpdateResponse]
type userFirewallAccessRuleRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleUpdateResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    userFirewallAccessRuleRuleUpdateResponseErrorJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseErrorJSON contains the JSON metadata for
// the struct [UserFirewallAccessRuleRuleUpdateResponseError]
type userFirewallAccessRuleRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleUpdateResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    userFirewallAccessRuleRuleUpdateResponseMessageJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseMessageJSON contains the JSON metadata
// for the struct [UserFirewallAccessRuleRuleUpdateResponseMessage]
type userFirewallAccessRuleRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleUpdateResponseResult struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []UserFirewallAccessRuleRuleUpdateResponseResultAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration UserFirewallAccessRuleRuleUpdateResponseResultConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode UserFirewallAccessRuleRuleUpdateResponseResultMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                             `json:"notes"`
	JSON  userFirewallAccessRuleRuleUpdateResponseResultJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseResultJSON contains the JSON metadata
// for the struct [UserFirewallAccessRuleRuleUpdateResponseResult]
type userFirewallAccessRuleRuleUpdateResponseResultJSON struct {
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

func (r *UserFirewallAccessRuleRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleUpdateResponseResultAllowedMode string

const (
	UserFirewallAccessRuleRuleUpdateResponseResultAllowedModeBlock            UserFirewallAccessRuleRuleUpdateResponseResultAllowedMode = "block"
	UserFirewallAccessRuleRuleUpdateResponseResultAllowedModeChallenge        UserFirewallAccessRuleRuleUpdateResponseResultAllowedMode = "challenge"
	UserFirewallAccessRuleRuleUpdateResponseResultAllowedModeWhitelist        UserFirewallAccessRuleRuleUpdateResponseResultAllowedMode = "whitelist"
	UserFirewallAccessRuleRuleUpdateResponseResultAllowedModeJsChallenge      UserFirewallAccessRuleRuleUpdateResponseResultAllowedMode = "js_challenge"
	UserFirewallAccessRuleRuleUpdateResponseResultAllowedModeManagedChallenge UserFirewallAccessRuleRuleUpdateResponseResultAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfiguration],
// [UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6Configuration],
// [UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfiguration],
// [UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfiguration]
// or
// [UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfiguration].
type UserFirewallAccessRuleRuleUpdateResponseResultConfiguration interface {
	implementsUserFirewallAccessRuleRuleUpdateResponseResultConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*UserFirewallAccessRuleRuleUpdateResponseResultConfiguration)(nil)).Elem(), "")
}

type UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                 `json:"value"`
	JSON  userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTipConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTipConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfiguration]
type userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTipConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfiguration) implementsUserFirewallAccessRuleRuleUpdateResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfigurationTargetIP UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                                   `json:"value"`
	JSON  userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6Configuration]
type userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6Configuration) implementsUserFirewallAccessRuleRuleUpdateResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationTargetIp6 UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                                   `json:"value"`
	JSON  userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfiguration]
type userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfiguration) implementsUserFirewallAccessRuleRuleUpdateResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                                  `json:"value"`
	JSON  userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTasnConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTasnConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfiguration]
type userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTasnConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfiguration) implementsUserFirewallAccessRuleRuleUpdateResponseResultConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfigurationTarget string

const (
	UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfigurationTargetASN UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                                      `json:"value"`
	JSON  userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfiguration]
type userFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfiguration) implementsUserFirewallAccessRuleRuleUpdateResponseResultConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationTargetCountry UserFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleUpdateResponseResultMode string

const (
	UserFirewallAccessRuleRuleUpdateResponseResultModeBlock            UserFirewallAccessRuleRuleUpdateResponseResultMode = "block"
	UserFirewallAccessRuleRuleUpdateResponseResultModeChallenge        UserFirewallAccessRuleRuleUpdateResponseResultMode = "challenge"
	UserFirewallAccessRuleRuleUpdateResponseResultModeWhitelist        UserFirewallAccessRuleRuleUpdateResponseResultMode = "whitelist"
	UserFirewallAccessRuleRuleUpdateResponseResultModeJsChallenge      UserFirewallAccessRuleRuleUpdateResponseResultMode = "js_challenge"
	UserFirewallAccessRuleRuleUpdateResponseResultModeManagedChallenge UserFirewallAccessRuleRuleUpdateResponseResultMode = "managed_challenge"
)

// Whether the API call was successful
type UserFirewallAccessRuleRuleUpdateResponseSuccess bool

const (
	UserFirewallAccessRuleRuleUpdateResponseSuccessTrue UserFirewallAccessRuleRuleUpdateResponseSuccess = true
)

type UserFirewallAccessRuleRuleDeleteResponse struct {
	Errors   []UserFirewallAccessRuleRuleDeleteResponseError   `json:"errors"`
	Messages []UserFirewallAccessRuleRuleDeleteResponseMessage `json:"messages"`
	Result   UserFirewallAccessRuleRuleDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success UserFirewallAccessRuleRuleDeleteResponseSuccess `json:"success"`
	JSON    userFirewallAccessRuleRuleDeleteResponseJSON    `json:"-"`
}

// userFirewallAccessRuleRuleDeleteResponseJSON contains the JSON metadata for the
// struct [UserFirewallAccessRuleRuleDeleteResponse]
type userFirewallAccessRuleRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleDeleteResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    userFirewallAccessRuleRuleDeleteResponseErrorJSON `json:"-"`
}

// userFirewallAccessRuleRuleDeleteResponseErrorJSON contains the JSON metadata for
// the struct [UserFirewallAccessRuleRuleDeleteResponseError]
type userFirewallAccessRuleRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleDeleteResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    userFirewallAccessRuleRuleDeleteResponseMessageJSON `json:"-"`
}

// userFirewallAccessRuleRuleDeleteResponseMessageJSON contains the JSON metadata
// for the struct [UserFirewallAccessRuleRuleDeleteResponseMessage]
type userFirewallAccessRuleRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleDeleteResponseResult struct {
	// The unique identifier of the IP Access rule.
	ID   string                                             `json:"id"`
	JSON userFirewallAccessRuleRuleDeleteResponseResultJSON `json:"-"`
}

// userFirewallAccessRuleRuleDeleteResponseResultJSON contains the JSON metadata
// for the struct [UserFirewallAccessRuleRuleDeleteResponseResult]
type userFirewallAccessRuleRuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserFirewallAccessRuleRuleDeleteResponseSuccess bool

const (
	UserFirewallAccessRuleRuleDeleteResponseSuccessTrue UserFirewallAccessRuleRuleDeleteResponseSuccess = true
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponse struct {
	Errors   []UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseError   `json:"errors"`
	Messages []UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMessage `json:"messages"`
	Result   UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseSuccess `json:"success"`
	JSON    userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseJSON    `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponse]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseError struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseErrorJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseErrorJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseError]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMessage struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMessageJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMessage]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResult struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                                                             `json:"notes"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResult]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultJSON struct {
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

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultAllowedMode string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultAllowedModeBlock            UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultAllowedMode = "block"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultAllowedModeChallenge        UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultAllowedMode = "challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultAllowedModeWhitelist        UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultAllowedMode = "whitelist"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultAllowedModeJsChallenge      UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultAllowedMode = "js_challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultAllowedModeManagedChallenge UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6Configuration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfiguration]
// or
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfiguration].
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfiguration interface {
	implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfiguration)(nil)).Elem(), "")
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                                                 `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTipConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTipConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTipConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfigurationTargetIP UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                                                                   `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6Configuration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6Configuration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationTargetIp6 UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                                                                   `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                                                                  `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTasnConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTasnConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTasnConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfigurationTargetASN UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                                                                      `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationTargetCountry UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultMode string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultModeBlock            UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultMode = "block"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultModeChallenge        UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultMode = "challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultModeWhitelist        UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultMode = "whitelist"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultModeJsChallenge      UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultMode = "js_challenge"
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultModeManagedChallenge UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseResultMode = "managed_challenge"
)

// Whether the API call was successful
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseSuccess bool

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseSuccessTrue UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleResponseSuccess = true
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
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIPConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIpv6Configuration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCidrConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTASNConfiguration]
// or
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCountryConfiguration].
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration interface {
	implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration)(nil)).Elem(), "")
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                                           `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTipConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTipConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIPConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTipConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIPConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIPConfigurationTargetIP UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                                                             `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIpv6Configuration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIpv6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIpv6Configuration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationTargetIp6 UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                                                             `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCidrConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCidrConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                                                            `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTasnConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTasnConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTASNConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTasnConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTASNConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTASNConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTASNConfigurationTargetASN UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                                                                `json:"value"`
	JSON  userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationJSON `json:"-"`
}

// userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationJSON
// contains the JSON metadata for the struct
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCountryConfiguration]
type userFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCountryConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationTargetCountry UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationTarget = "country"
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
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6Configuration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfiguration].
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration interface {
	implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration()
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTargetIP UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6Configuration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6ConfigurationTargetIp6 UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfigurationTargetASN UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfigurationTargetCountry UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfigurationTarget = "country"
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
	UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTargetASN     UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParamsFiltersConfigurationTarget = "asn"
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
