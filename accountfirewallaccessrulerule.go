// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountFirewallAccessRuleRuleService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountFirewallAccessRuleRuleService] method instead.
type AccountFirewallAccessRuleRuleService struct {
	Options []option.RequestOption
}

// NewAccountFirewallAccessRuleRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountFirewallAccessRuleRuleService(opts ...option.RequestOption) (r *AccountFirewallAccessRuleRuleService) {
	r = &AccountFirewallAccessRuleRuleService{}
	r.Options = opts
	return
}

// Fetches the details of an IP Access rule defined at the account level.
func (r *AccountFirewallAccessRuleRuleService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *ResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/firewall/access_rules/rules/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an IP Access rule defined at the account level.
//
// Note: This operation will affect all zones in the account.
func (r *AccountFirewallAccessRuleRuleService) Update(ctx context.Context, accountIdentifier interface{}, identifier interface{}, body AccountFirewallAccessRuleRuleUpdateParams, opts ...option.RequestOption) (res *ResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/firewall/access_rules/rules/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an existing IP Access rule defined at the account level.
//
// Note: This operation will affect all zones in the account.
func (r *AccountFirewallAccessRuleRuleService) Delete(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *APIResponseSingleIDKYar7dC1, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/firewall/access_rules/rules/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new IP Access rule for an account. The rule will apply to all zones in
// the account.
//
// Note: To create an IP Access rule that applies to a single zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *AccountFirewallAccessRuleRuleService) IPAccessRulesForAnAccountNewAnIPAccessRule(ctx context.Context, accountIdentifier interface{}, body AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParams, opts ...option.RequestOption) (res *ResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/firewall/access_rules/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches IP Access rules of an account. These rules apply to all the zones in the
// account. You can filter the results using several optional parameters.
func (r *AccountFirewallAccessRuleRuleService) IPAccessRulesForAnAccountListIPAccessRules(ctx context.Context, accountIdentifier interface{}, query AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParams, opts ...option.RequestOption) (res *ResponseCollectionYgt6DzoC, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/firewall/access_rules/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type APIResponseSingleIDKYar7dC1 struct {
	Errors   []APIResponseSingleIDKYar7dC1Error   `json:"errors"`
	Messages []APIResponseSingleIDKYar7dC1Message `json:"messages"`
	Result   APIResponseSingleIDKYar7dC1Result    `json:"result,nullable"`
	// Whether the API call was successful
	Success APIResponseSingleIDKYar7dC1Success `json:"success"`
	JSON    apiResponseSingleIdkYar7dC1JSON    `json:"-"`
}

// apiResponseSingleIdkYar7dC1JSON contains the JSON metadata for the struct
// [APIResponseSingleIDKYar7dC1]
type apiResponseSingleIdkYar7dC1JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDKYar7dC1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIResponseSingleIDKYar7dC1Error struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    apiResponseSingleIdkYar7dC1ErrorJSON `json:"-"`
}

// apiResponseSingleIdkYar7dC1ErrorJSON contains the JSON metadata for the struct
// [APIResponseSingleIDKYar7dC1Error]
type apiResponseSingleIdkYar7dC1ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDKYar7dC1Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIResponseSingleIDKYar7dC1Message struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    apiResponseSingleIdkYar7dC1MessageJSON `json:"-"`
}

// apiResponseSingleIdkYar7dC1MessageJSON contains the JSON metadata for the struct
// [APIResponseSingleIDKYar7dC1Message]
type apiResponseSingleIdkYar7dC1MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDKYar7dC1Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIResponseSingleIDKYar7dC1Result struct {
	// Identifier
	ID   string                                `json:"id,required"`
	JSON apiResponseSingleIdkYar7dC1ResultJSON `json:"-"`
}

// apiResponseSingleIdkYar7dC1ResultJSON contains the JSON metadata for the struct
// [APIResponseSingleIDKYar7dC1Result]
type apiResponseSingleIdkYar7dC1ResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDKYar7dC1Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type APIResponseSingleIDKYar7dC1Success bool

const (
	APIResponseSingleIDKYar7dC1SuccessTrue APIResponseSingleIDKYar7dC1Success = true
)

type AccountFirewallAccessRuleRuleUpdateParams struct {
	// The rule configuration.
	Configuration param.Field[AccountFirewallAccessRuleRuleUpdateParamsConfiguration] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[AccountFirewallAccessRuleRuleUpdateParamsMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r AccountFirewallAccessRuleRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
//
// Satisfied by
// [AccountFirewallAccessRuleRuleUpdateParamsConfigurationIPConfiguration],
// [AccountFirewallAccessRuleRuleUpdateParamsConfigurationIpv6Configuration],
// [AccountFirewallAccessRuleRuleUpdateParamsConfigurationCidrConfiguration],
// [AccountFirewallAccessRuleRuleUpdateParamsConfigurationASNConfiguration],
// [AccountFirewallAccessRuleRuleUpdateParamsConfigurationCountryConfiguration].
type AccountFirewallAccessRuleRuleUpdateParamsConfiguration interface {
	implementsAccountFirewallAccessRuleRuleUpdateParamsConfiguration()
}

type AccountFirewallAccessRuleRuleUpdateParamsConfigurationIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleUpdateParamsConfigurationIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationIPConfiguration) implementsAccountFirewallAccessRuleRuleUpdateParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccountFirewallAccessRuleRuleUpdateParamsConfigurationIPConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleUpdateParamsConfigurationIPConfigurationTargetIP AccountFirewallAccessRuleRuleUpdateParamsConfigurationIPConfigurationTarget = "ip"
)

type AccountFirewallAccessRuleRuleUpdateParamsConfigurationIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleUpdateParamsConfigurationIpv6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationIpv6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationIpv6Configuration) implementsAccountFirewallAccessRuleRuleUpdateParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type AccountFirewallAccessRuleRuleUpdateParamsConfigurationIpv6ConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleUpdateParamsConfigurationIpv6ConfigurationTargetIp6 AccountFirewallAccessRuleRuleUpdateParamsConfigurationIpv6ConfigurationTarget = "ip6"
)

type AccountFirewallAccessRuleRuleUpdateParamsConfigurationCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleUpdateParamsConfigurationCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationCidrConfiguration) implementsAccountFirewallAccessRuleRuleUpdateParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type AccountFirewallAccessRuleRuleUpdateParamsConfigurationCidrConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleUpdateParamsConfigurationCidrConfigurationTargetIPRange AccountFirewallAccessRuleRuleUpdateParamsConfigurationCidrConfigurationTarget = "ip_range"
)

type AccountFirewallAccessRuleRuleUpdateParamsConfigurationASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleUpdateParamsConfigurationASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationASNConfiguration) implementsAccountFirewallAccessRuleRuleUpdateParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type AccountFirewallAccessRuleRuleUpdateParamsConfigurationASNConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleUpdateParamsConfigurationASNConfigurationTargetASN AccountFirewallAccessRuleRuleUpdateParamsConfigurationASNConfigurationTarget = "asn"
)

type AccountFirewallAccessRuleRuleUpdateParamsConfigurationCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleUpdateParamsConfigurationCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationCountryConfiguration) implementsAccountFirewallAccessRuleRuleUpdateParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type AccountFirewallAccessRuleRuleUpdateParamsConfigurationCountryConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleUpdateParamsConfigurationCountryConfigurationTargetCountry AccountFirewallAccessRuleRuleUpdateParamsConfigurationCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type AccountFirewallAccessRuleRuleUpdateParamsMode string

const (
	AccountFirewallAccessRuleRuleUpdateParamsModeBlock            AccountFirewallAccessRuleRuleUpdateParamsMode = "block"
	AccountFirewallAccessRuleRuleUpdateParamsModeChallenge        AccountFirewallAccessRuleRuleUpdateParamsMode = "challenge"
	AccountFirewallAccessRuleRuleUpdateParamsModeWhitelist        AccountFirewallAccessRuleRuleUpdateParamsMode = "whitelist"
	AccountFirewallAccessRuleRuleUpdateParamsModeJsChallenge      AccountFirewallAccessRuleRuleUpdateParamsMode = "js_challenge"
	AccountFirewallAccessRuleRuleUpdateParamsModeManagedChallenge AccountFirewallAccessRuleRuleUpdateParamsMode = "managed_challenge"
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParams struct {
	// The rule configuration.
	Configuration param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
//
// Satisfied by
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIPConfiguration],
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIpv6Configuration],
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCidrConfiguration],
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationASNConfiguration],
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCountryConfiguration].
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration interface {
	implementsAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration()
}

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIPConfiguration) implementsAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIPConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIPConfigurationTargetIP AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIPConfigurationTarget = "ip"
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIpv6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIpv6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIpv6Configuration) implementsAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIpv6ConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIpv6ConfigurationTargetIp6 AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationIpv6ConfigurationTarget = "ip6"
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCidrConfiguration) implementsAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCidrConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCidrConfigurationTargetIPRange AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCidrConfigurationTarget = "ip_range"
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationASNConfiguration) implementsAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationASNConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationASNConfigurationTargetASN AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationASNConfigurationTarget = "asn"
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCountryConfiguration) implementsAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCountryConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCountryConfigurationTargetCountry AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsMode string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsModeBlock            AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsMode = "block"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsModeChallenge        AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsMode = "challenge"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsModeWhitelist        AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsMode = "whitelist"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsModeJsChallenge      AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsMode = "js_challenge"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsModeManagedChallenge AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsMode = "managed_challenge"
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParams struct {
	// The direction used to sort returned rules.
	Direction     param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsDirection]     `query:"direction"`
	EgsPagination param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsEgsPagination] `query:"egs-pagination"`
	Filters       param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFilters]       `query:"filters"`
	// The field used to sort returned rules.
	Order param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsOrder] `query:"order"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParams]'s
// query parameters as `url.Values`.
func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsDirection string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsDirectionAsc  AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsDirection = "asc"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsDirectionDesc AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsDirection = "desc"
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsEgsPagination struct {
	Json param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsEgsPaginationJson] `query:"json"`
}

// URLQuery serializes
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsEgsPagination]'s
// query parameters as `url.Values`.
func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsEgsPagination) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsEgsPaginationJson struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsEgsPaginationJson]'s
// query parameters as `url.Values`.
func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsEgsPaginationJson) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFilters struct {
	// The target to search in existing rules.
	ConfigurationTarget param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersConfigurationTarget] `query:"configuration.target"`
	// The target value to search for in existing rules: an IP address, an IP address
	// range, or a country code, depending on the provided `configuration.target`.
	// Notes: You can search for a single IPv4 address, an IP address range with a
	// subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	ConfigurationValue param.Field[string] `query:"configuration.value"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMode] `query:"mode"`
	// The string to search for in the notes of existing IP Access rules. Notes: For
	// example, the string 'attack' would match IP Access rules with notes 'Attack
	// 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
}

// URLQuery serializes
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFilters]'s
// query parameters as `url.Values`.
func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFilters) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The target to search in existing rules.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersConfigurationTargetIP      AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersConfigurationTarget = "ip"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersConfigurationTargetIPRange AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersConfigurationTarget = "ip_range"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersConfigurationTargetASN     AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersConfigurationTarget = "asn"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersConfigurationTargetCountry AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersConfigurationTarget = "country"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMatch string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMatchAny AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMatch = "any"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMatchAll AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMatch = "all"
)

// The action to apply to a matched request.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMode string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersModeBlock            AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMode = "block"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersModeChallenge        AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMode = "challenge"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersModeWhitelist        AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMode = "whitelist"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersModeJsChallenge      AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMode = "js_challenge"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersModeManagedChallenge AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsFiltersMode = "managed_challenge"
)

// The field used to sort returned rules.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsOrder string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsOrderConfigurationTarget AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsOrder = "configuration.target"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsOrderConfigurationValue  AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsOrder = "configuration.value"
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsOrderMode                AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParamsOrder = "mode"
)
