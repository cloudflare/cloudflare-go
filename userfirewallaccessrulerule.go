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
func (r *UserFirewallAccessRuleRuleService) Update(ctx context.Context, identifier string, body UserFirewallAccessRuleRuleUpdateParams, opts ...option.RequestOption) (res *RuleSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an IP Access rule at the user level.
//
// Note: Deleting a user-level rule will affect all zones owned by the user.
func (r *UserFirewallAccessRuleRuleService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *RuleSingleIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new IP Access rule for all zones owned by the current user.
//
// Note: To create an IP Access rule that applies to a specific zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *UserFirewallAccessRuleRuleService) IPAccessRulesForAUserNewAnIPAccessRule(ctx context.Context, body UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParams, opts ...option.RequestOption) (res *RuleSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/firewall/access_rules/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches IP Access rules of the user. You can filter the results using several
// optional parameters.
func (r *UserFirewallAccessRuleRuleService) IPAccessRulesForAUserListIPAccessRules(ctx context.Context, query UserFirewallAccessRuleRuleIPAccessRulesForAUserListIPAccessRulesParams, opts ...option.RequestOption) (res *RuleCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/firewall/access_rules/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RuleCollectionResponse struct {
	Errors     []RuleCollectionResponseError    `json:"errors"`
	Messages   []RuleCollectionResponseMessage  `json:"messages"`
	Result     []RuleCollectionResponseResult   `json:"result"`
	ResultInfo RuleCollectionResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success RuleCollectionResponseSuccess `json:"success"`
	JSON    ruleCollectionResponseJSON    `json:"-"`
}

// ruleCollectionResponseJSON contains the JSON metadata for the struct
// [RuleCollectionResponse]
type ruleCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleCollectionResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    ruleCollectionResponseErrorJSON `json:"-"`
}

// ruleCollectionResponseErrorJSON contains the JSON metadata for the struct
// [RuleCollectionResponseError]
type ruleCollectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCollectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleCollectionResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    ruleCollectionResponseMessageJSON `json:"-"`
}

// ruleCollectionResponseMessageJSON contains the JSON metadata for the struct
// [RuleCollectionResponseMessage]
type ruleCollectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCollectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleCollectionResponseResult struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []RuleCollectionResponseResultAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration RuleCollectionResponseResultConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode RuleCollectionResponseResultMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                           `json:"notes"`
	JSON  ruleCollectionResponseResultJSON `json:"-"`
}

// ruleCollectionResponseResultJSON contains the JSON metadata for the struct
// [RuleCollectionResponseResult]
type ruleCollectionResponseResultJSON struct {
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

func (r *RuleCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type RuleCollectionResponseResultAllowedMode string

const (
	RuleCollectionResponseResultAllowedModeBlock            RuleCollectionResponseResultAllowedMode = "block"
	RuleCollectionResponseResultAllowedModeChallenge        RuleCollectionResponseResultAllowedMode = "challenge"
	RuleCollectionResponseResultAllowedModeWhitelist        RuleCollectionResponseResultAllowedMode = "whitelist"
	RuleCollectionResponseResultAllowedModeJsChallenge      RuleCollectionResponseResultAllowedMode = "js_challenge"
	RuleCollectionResponseResultAllowedModeManagedChallenge RuleCollectionResponseResultAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by [RuleCollectionResponseResultConfigurationIPConfiguration],
// [RuleCollectionResponseResultConfigurationIpv6Configuration],
// [RuleCollectionResponseResultConfigurationCidrConfiguration],
// [RuleCollectionResponseResultConfigurationASNConfiguration] or
// [RuleCollectionResponseResultConfigurationCountryConfiguration].
type RuleCollectionResponseResultConfiguration interface {
	implementsRuleCollectionResponseResultConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*RuleCollectionResponseResultConfiguration)(nil)).Elem(), "")
}

type RuleCollectionResponseResultConfigurationIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target RuleCollectionResponseResultConfigurationIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                       `json:"value"`
	JSON  ruleCollectionResponseResultConfigurationIPConfigurationJSON `json:"-"`
}

// ruleCollectionResponseResultConfigurationIPConfigurationJSON contains the JSON
// metadata for the struct
// [RuleCollectionResponseResultConfigurationIPConfiguration]
type ruleCollectionResponseResultConfigurationIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCollectionResponseResultConfigurationIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleCollectionResponseResultConfigurationIPConfiguration) implementsRuleCollectionResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type RuleCollectionResponseResultConfigurationIPConfigurationTarget string

const (
	RuleCollectionResponseResultConfigurationIPConfigurationTargetIP RuleCollectionResponseResultConfigurationIPConfigurationTarget = "ip"
)

type RuleCollectionResponseResultConfigurationIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target RuleCollectionResponseResultConfigurationIpv6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                         `json:"value"`
	JSON  ruleCollectionResponseResultConfigurationIpv6ConfigurationJSON `json:"-"`
}

// ruleCollectionResponseResultConfigurationIpv6ConfigurationJSON contains the JSON
// metadata for the struct
// [RuleCollectionResponseResultConfigurationIpv6Configuration]
type ruleCollectionResponseResultConfigurationIpv6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCollectionResponseResultConfigurationIpv6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleCollectionResponseResultConfigurationIpv6Configuration) implementsRuleCollectionResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type RuleCollectionResponseResultConfigurationIpv6ConfigurationTarget string

const (
	RuleCollectionResponseResultConfigurationIpv6ConfigurationTargetIp6 RuleCollectionResponseResultConfigurationIpv6ConfigurationTarget = "ip6"
)

type RuleCollectionResponseResultConfigurationCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target RuleCollectionResponseResultConfigurationCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                         `json:"value"`
	JSON  ruleCollectionResponseResultConfigurationCidrConfigurationJSON `json:"-"`
}

// ruleCollectionResponseResultConfigurationCidrConfigurationJSON contains the JSON
// metadata for the struct
// [RuleCollectionResponseResultConfigurationCidrConfiguration]
type ruleCollectionResponseResultConfigurationCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCollectionResponseResultConfigurationCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleCollectionResponseResultConfigurationCidrConfiguration) implementsRuleCollectionResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type RuleCollectionResponseResultConfigurationCidrConfigurationTarget string

const (
	RuleCollectionResponseResultConfigurationCidrConfigurationTargetIPRange RuleCollectionResponseResultConfigurationCidrConfigurationTarget = "ip_range"
)

type RuleCollectionResponseResultConfigurationASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target RuleCollectionResponseResultConfigurationASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                        `json:"value"`
	JSON  ruleCollectionResponseResultConfigurationASNConfigurationJSON `json:"-"`
}

// ruleCollectionResponseResultConfigurationASNConfigurationJSON contains the JSON
// metadata for the struct
// [RuleCollectionResponseResultConfigurationASNConfiguration]
type ruleCollectionResponseResultConfigurationASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCollectionResponseResultConfigurationASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleCollectionResponseResultConfigurationASNConfiguration) implementsRuleCollectionResponseResultConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type RuleCollectionResponseResultConfigurationASNConfigurationTarget string

const (
	RuleCollectionResponseResultConfigurationASNConfigurationTargetASN RuleCollectionResponseResultConfigurationASNConfigurationTarget = "asn"
)

type RuleCollectionResponseResultConfigurationCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target RuleCollectionResponseResultConfigurationCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                            `json:"value"`
	JSON  ruleCollectionResponseResultConfigurationCountryConfigurationJSON `json:"-"`
}

// ruleCollectionResponseResultConfigurationCountryConfigurationJSON contains the
// JSON metadata for the struct
// [RuleCollectionResponseResultConfigurationCountryConfiguration]
type ruleCollectionResponseResultConfigurationCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCollectionResponseResultConfigurationCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleCollectionResponseResultConfigurationCountryConfiguration) implementsRuleCollectionResponseResultConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type RuleCollectionResponseResultConfigurationCountryConfigurationTarget string

const (
	RuleCollectionResponseResultConfigurationCountryConfigurationTargetCountry RuleCollectionResponseResultConfigurationCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type RuleCollectionResponseResultMode string

const (
	RuleCollectionResponseResultModeBlock            RuleCollectionResponseResultMode = "block"
	RuleCollectionResponseResultModeChallenge        RuleCollectionResponseResultMode = "challenge"
	RuleCollectionResponseResultModeWhitelist        RuleCollectionResponseResultMode = "whitelist"
	RuleCollectionResponseResultModeJsChallenge      RuleCollectionResponseResultMode = "js_challenge"
	RuleCollectionResponseResultModeManagedChallenge RuleCollectionResponseResultMode = "managed_challenge"
)

type RuleCollectionResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       ruleCollectionResponseResultInfoJSON `json:"-"`
}

// ruleCollectionResponseResultInfoJSON contains the JSON metadata for the struct
// [RuleCollectionResponseResultInfo]
type ruleCollectionResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCollectionResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleCollectionResponseSuccess bool

const (
	RuleCollectionResponseSuccessTrue RuleCollectionResponseSuccess = true
)

type RuleSingleIDResponse struct {
	Errors   []RuleSingleIDResponseError   `json:"errors"`
	Messages []RuleSingleIDResponseMessage `json:"messages"`
	Result   RuleSingleIDResponseResult    `json:"result"`
	// Whether the API call was successful
	Success RuleSingleIDResponseSuccess `json:"success"`
	JSON    ruleSingleIDResponseJSON    `json:"-"`
}

// ruleSingleIDResponseJSON contains the JSON metadata for the struct
// [RuleSingleIDResponse]
type ruleSingleIDResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSingleIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleSingleIDResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    ruleSingleIDResponseErrorJSON `json:"-"`
}

// ruleSingleIDResponseErrorJSON contains the JSON metadata for the struct
// [RuleSingleIDResponseError]
type ruleSingleIDResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSingleIDResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleSingleIDResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    ruleSingleIDResponseMessageJSON `json:"-"`
}

// ruleSingleIDResponseMessageJSON contains the JSON metadata for the struct
// [RuleSingleIDResponseMessage]
type ruleSingleIDResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSingleIDResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleSingleIDResponseResult struct {
	// The unique identifier of the IP Access rule.
	ID   string                         `json:"id"`
	JSON ruleSingleIDResponseResultJSON `json:"-"`
}

// ruleSingleIDResponseResultJSON contains the JSON metadata for the struct
// [RuleSingleIDResponseResult]
type ruleSingleIDResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSingleIDResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleSingleIDResponseSuccess bool

const (
	RuleSingleIDResponseSuccessTrue RuleSingleIDResponseSuccess = true
)

type RuleSingleResponse struct {
	Errors   []RuleSingleResponseError   `json:"errors"`
	Messages []RuleSingleResponseMessage `json:"messages"`
	Result   RuleSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success RuleSingleResponseSuccess `json:"success"`
	JSON    ruleSingleResponseJSON    `json:"-"`
}

// ruleSingleResponseJSON contains the JSON metadata for the struct
// [RuleSingleResponse]
type ruleSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleSingleResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    ruleSingleResponseErrorJSON `json:"-"`
}

// ruleSingleResponseErrorJSON contains the JSON metadata for the struct
// [RuleSingleResponseError]
type ruleSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleSingleResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    ruleSingleResponseMessageJSON `json:"-"`
}

// ruleSingleResponseMessageJSON contains the JSON metadata for the struct
// [RuleSingleResponseMessage]
type ruleSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleSingleResponseResult struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []RuleSingleResponseResultAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration RuleSingleResponseResultConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode RuleSingleResponseResultMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                       `json:"notes"`
	JSON  ruleSingleResponseResultJSON `json:"-"`
}

// ruleSingleResponseResultJSON contains the JSON metadata for the struct
// [RuleSingleResponseResult]
type ruleSingleResponseResultJSON struct {
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

func (r *RuleSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type RuleSingleResponseResultAllowedMode string

const (
	RuleSingleResponseResultAllowedModeBlock            RuleSingleResponseResultAllowedMode = "block"
	RuleSingleResponseResultAllowedModeChallenge        RuleSingleResponseResultAllowedMode = "challenge"
	RuleSingleResponseResultAllowedModeWhitelist        RuleSingleResponseResultAllowedMode = "whitelist"
	RuleSingleResponseResultAllowedModeJsChallenge      RuleSingleResponseResultAllowedMode = "js_challenge"
	RuleSingleResponseResultAllowedModeManagedChallenge RuleSingleResponseResultAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by [RuleSingleResponseResultConfigurationIPConfiguration],
// [RuleSingleResponseResultConfigurationIpv6Configuration],
// [RuleSingleResponseResultConfigurationCidrConfiguration],
// [RuleSingleResponseResultConfigurationASNConfiguration] or
// [RuleSingleResponseResultConfigurationCountryConfiguration].
type RuleSingleResponseResultConfiguration interface {
	implementsRuleSingleResponseResultConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*RuleSingleResponseResultConfiguration)(nil)).Elem(), "")
}

type RuleSingleResponseResultConfigurationIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target RuleSingleResponseResultConfigurationIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                   `json:"value"`
	JSON  ruleSingleResponseResultConfigurationIPConfigurationJSON `json:"-"`
}

// ruleSingleResponseResultConfigurationIPConfigurationJSON contains the JSON
// metadata for the struct [RuleSingleResponseResultConfigurationIPConfiguration]
type ruleSingleResponseResultConfigurationIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSingleResponseResultConfigurationIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleSingleResponseResultConfigurationIPConfiguration) implementsRuleSingleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type RuleSingleResponseResultConfigurationIPConfigurationTarget string

const (
	RuleSingleResponseResultConfigurationIPConfigurationTargetIP RuleSingleResponseResultConfigurationIPConfigurationTarget = "ip"
)

type RuleSingleResponseResultConfigurationIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target RuleSingleResponseResultConfigurationIpv6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                     `json:"value"`
	JSON  ruleSingleResponseResultConfigurationIpv6ConfigurationJSON `json:"-"`
}

// ruleSingleResponseResultConfigurationIpv6ConfigurationJSON contains the JSON
// metadata for the struct [RuleSingleResponseResultConfigurationIpv6Configuration]
type ruleSingleResponseResultConfigurationIpv6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSingleResponseResultConfigurationIpv6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleSingleResponseResultConfigurationIpv6Configuration) implementsRuleSingleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type RuleSingleResponseResultConfigurationIpv6ConfigurationTarget string

const (
	RuleSingleResponseResultConfigurationIpv6ConfigurationTargetIp6 RuleSingleResponseResultConfigurationIpv6ConfigurationTarget = "ip6"
)

type RuleSingleResponseResultConfigurationCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target RuleSingleResponseResultConfigurationCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                     `json:"value"`
	JSON  ruleSingleResponseResultConfigurationCidrConfigurationJSON `json:"-"`
}

// ruleSingleResponseResultConfigurationCidrConfigurationJSON contains the JSON
// metadata for the struct [RuleSingleResponseResultConfigurationCidrConfiguration]
type ruleSingleResponseResultConfigurationCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSingleResponseResultConfigurationCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleSingleResponseResultConfigurationCidrConfiguration) implementsRuleSingleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type RuleSingleResponseResultConfigurationCidrConfigurationTarget string

const (
	RuleSingleResponseResultConfigurationCidrConfigurationTargetIPRange RuleSingleResponseResultConfigurationCidrConfigurationTarget = "ip_range"
)

type RuleSingleResponseResultConfigurationASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target RuleSingleResponseResultConfigurationASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                    `json:"value"`
	JSON  ruleSingleResponseResultConfigurationASNConfigurationJSON `json:"-"`
}

// ruleSingleResponseResultConfigurationASNConfigurationJSON contains the JSON
// metadata for the struct [RuleSingleResponseResultConfigurationASNConfiguration]
type ruleSingleResponseResultConfigurationASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSingleResponseResultConfigurationASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleSingleResponseResultConfigurationASNConfiguration) implementsRuleSingleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type RuleSingleResponseResultConfigurationASNConfigurationTarget string

const (
	RuleSingleResponseResultConfigurationASNConfigurationTargetASN RuleSingleResponseResultConfigurationASNConfigurationTarget = "asn"
)

type RuleSingleResponseResultConfigurationCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target RuleSingleResponseResultConfigurationCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                        `json:"value"`
	JSON  ruleSingleResponseResultConfigurationCountryConfigurationJSON `json:"-"`
}

// ruleSingleResponseResultConfigurationCountryConfigurationJSON contains the JSON
// metadata for the struct
// [RuleSingleResponseResultConfigurationCountryConfiguration]
type ruleSingleResponseResultConfigurationCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSingleResponseResultConfigurationCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleSingleResponseResultConfigurationCountryConfiguration) implementsRuleSingleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type RuleSingleResponseResultConfigurationCountryConfigurationTarget string

const (
	RuleSingleResponseResultConfigurationCountryConfigurationTargetCountry RuleSingleResponseResultConfigurationCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type RuleSingleResponseResultMode string

const (
	RuleSingleResponseResultModeBlock            RuleSingleResponseResultMode = "block"
	RuleSingleResponseResultModeChallenge        RuleSingleResponseResultMode = "challenge"
	RuleSingleResponseResultModeWhitelist        RuleSingleResponseResultMode = "whitelist"
	RuleSingleResponseResultModeJsChallenge      RuleSingleResponseResultMode = "js_challenge"
	RuleSingleResponseResultModeManagedChallenge RuleSingleResponseResultMode = "managed_challenge"
)

// Whether the API call was successful
type RuleSingleResponseSuccess bool

const (
	RuleSingleResponseSuccessTrue RuleSingleResponseSuccess = true
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
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIPConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIpv6Configuration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCidrConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationASNConfiguration],
// [UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCountryConfiguration].
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration interface {
	implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration()
}

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIPConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIPConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIPConfigurationTargetIP UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIPConfigurationTarget = "ip"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIpv6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIpv6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIpv6Configuration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIpv6ConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIpv6ConfigurationTargetIp6 UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationIpv6ConfigurationTarget = "ip6"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCidrConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCidrConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCidrConfigurationTargetIPRange UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCidrConfigurationTarget = "ip_range"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationASNConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationASNConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationASNConfigurationTargetASN UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationASNConfigurationTarget = "asn"
)

type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCountryConfiguration) implementsUserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCountryConfigurationTarget string

const (
	UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCountryConfigurationTargetCountry UserFirewallAccessRuleRuleIPAccessRulesForAUserNewAnIPAccessRuleParamsConfigurationCountryConfigurationTarget = "country"
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
