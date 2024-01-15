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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
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
func (r *AccountFirewallAccessRuleRuleService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *AccountFirewallAccessRuleRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/firewall/access_rules/rules/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an IP Access rule defined at the account level.
//
// Note: This operation will affect all zones in the account.
func (r *AccountFirewallAccessRuleRuleService) Update(ctx context.Context, accountIdentifier interface{}, identifier interface{}, body AccountFirewallAccessRuleRuleUpdateParams, opts ...option.RequestOption) (res *AccountFirewallAccessRuleRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/firewall/access_rules/rules/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an existing IP Access rule defined at the account level.
//
// Note: This operation will affect all zones in the account.
func (r *AccountFirewallAccessRuleRuleService) Delete(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *AccountFirewallAccessRuleRuleDeleteResponse, err error) {
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
func (r *AccountFirewallAccessRuleRuleService) IPAccessRulesForAnAccountNewAnIPAccessRule(ctx context.Context, accountIdentifier interface{}, body AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParams, opts ...option.RequestOption) (res *AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/firewall/access_rules/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches IP Access rules of an account. These rules apply to all the zones in the
// account. You can filter the results using several optional parameters.
func (r *AccountFirewallAccessRuleRuleService) IPAccessRulesForAnAccountListIPAccessRules(ctx context.Context, accountIdentifier interface{}, query AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesParams, opts ...option.RequestOption) (res *shared.Page[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%v/firewall/access_rules/rules", accountIdentifier)
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

type AccountFirewallAccessRuleRuleGetResponse struct {
	Errors   []AccountFirewallAccessRuleRuleGetResponseError   `json:"errors"`
	Messages []AccountFirewallAccessRuleRuleGetResponseMessage `json:"messages"`
	Result   interface{}                                       `json:"result"`
	// Whether the API call was successful
	Success AccountFirewallAccessRuleRuleGetResponseSuccess `json:"success"`
	JSON    accountFirewallAccessRuleRuleGetResponseJSON    `json:"-"`
}

// accountFirewallAccessRuleRuleGetResponseJSON contains the JSON metadata for the
// struct [AccountFirewallAccessRuleRuleGetResponse]
type accountFirewallAccessRuleRuleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountFirewallAccessRuleRuleGetResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountFirewallAccessRuleRuleGetResponseErrorJSON `json:"-"`
}

// accountFirewallAccessRuleRuleGetResponseErrorJSON contains the JSON metadata for
// the struct [AccountFirewallAccessRuleRuleGetResponseError]
type accountFirewallAccessRuleRuleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountFirewallAccessRuleRuleGetResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountFirewallAccessRuleRuleGetResponseMessageJSON `json:"-"`
}

// accountFirewallAccessRuleRuleGetResponseMessageJSON contains the JSON metadata
// for the struct [AccountFirewallAccessRuleRuleGetResponseMessage]
type accountFirewallAccessRuleRuleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountFirewallAccessRuleRuleGetResponseSuccess bool

const (
	AccountFirewallAccessRuleRuleGetResponseSuccessTrue AccountFirewallAccessRuleRuleGetResponseSuccess = true
)

type AccountFirewallAccessRuleRuleUpdateResponse struct {
	Errors   []AccountFirewallAccessRuleRuleUpdateResponseError   `json:"errors"`
	Messages []AccountFirewallAccessRuleRuleUpdateResponseMessage `json:"messages"`
	Result   interface{}                                          `json:"result"`
	// Whether the API call was successful
	Success AccountFirewallAccessRuleRuleUpdateResponseSuccess `json:"success"`
	JSON    accountFirewallAccessRuleRuleUpdateResponseJSON    `json:"-"`
}

// accountFirewallAccessRuleRuleUpdateResponseJSON contains the JSON metadata for
// the struct [AccountFirewallAccessRuleRuleUpdateResponse]
type accountFirewallAccessRuleRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountFirewallAccessRuleRuleUpdateResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountFirewallAccessRuleRuleUpdateResponseErrorJSON `json:"-"`
}

// accountFirewallAccessRuleRuleUpdateResponseErrorJSON contains the JSON metadata
// for the struct [AccountFirewallAccessRuleRuleUpdateResponseError]
type accountFirewallAccessRuleRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountFirewallAccessRuleRuleUpdateResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountFirewallAccessRuleRuleUpdateResponseMessageJSON `json:"-"`
}

// accountFirewallAccessRuleRuleUpdateResponseMessageJSON contains the JSON
// metadata for the struct [AccountFirewallAccessRuleRuleUpdateResponseMessage]
type accountFirewallAccessRuleRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountFirewallAccessRuleRuleUpdateResponseSuccess bool

const (
	AccountFirewallAccessRuleRuleUpdateResponseSuccessTrue AccountFirewallAccessRuleRuleUpdateResponseSuccess = true
)

type AccountFirewallAccessRuleRuleDeleteResponse struct {
	Errors   []AccountFirewallAccessRuleRuleDeleteResponseError   `json:"errors"`
	Messages []AccountFirewallAccessRuleRuleDeleteResponseMessage `json:"messages"`
	Result   AccountFirewallAccessRuleRuleDeleteResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountFirewallAccessRuleRuleDeleteResponseSuccess `json:"success"`
	JSON    accountFirewallAccessRuleRuleDeleteResponseJSON    `json:"-"`
}

// accountFirewallAccessRuleRuleDeleteResponseJSON contains the JSON metadata for
// the struct [AccountFirewallAccessRuleRuleDeleteResponse]
type accountFirewallAccessRuleRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountFirewallAccessRuleRuleDeleteResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountFirewallAccessRuleRuleDeleteResponseErrorJSON `json:"-"`
}

// accountFirewallAccessRuleRuleDeleteResponseErrorJSON contains the JSON metadata
// for the struct [AccountFirewallAccessRuleRuleDeleteResponseError]
type accountFirewallAccessRuleRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountFirewallAccessRuleRuleDeleteResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountFirewallAccessRuleRuleDeleteResponseMessageJSON `json:"-"`
}

// accountFirewallAccessRuleRuleDeleteResponseMessageJSON contains the JSON
// metadata for the struct [AccountFirewallAccessRuleRuleDeleteResponseMessage]
type accountFirewallAccessRuleRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountFirewallAccessRuleRuleDeleteResponseResult struct {
	// Identifier
	ID   string                                                `json:"id,required"`
	JSON accountFirewallAccessRuleRuleDeleteResponseResultJSON `json:"-"`
}

// accountFirewallAccessRuleRuleDeleteResponseResultJSON contains the JSON metadata
// for the struct [AccountFirewallAccessRuleRuleDeleteResponseResult]
type accountFirewallAccessRuleRuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountFirewallAccessRuleRuleDeleteResponseSuccess bool

const (
	AccountFirewallAccessRuleRuleDeleteResponseSuccessTrue AccountFirewallAccessRuleRuleDeleteResponseSuccess = true
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponse struct {
	Errors   []AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseError   `json:"errors"`
	Messages []AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseMessage `json:"messages"`
	Result   interface{}                                                                              `json:"result"`
	// Whether the API call was successful
	Success AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseSuccess `json:"success"`
	JSON    accountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseJSON    `json:"-"`
}

// accountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseJSON
// contains the JSON metadata for the struct
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponse]
type accountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseError struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseErrorJSON `json:"-"`
}

// accountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseError]
type accountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseMessage struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    accountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseMessageJSON `json:"-"`
}

// accountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseMessage]
type accountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseSuccess bool

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseSuccessTrue AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleResponseSuccess = true
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountListIPAccessRulesResponse = interface{}

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
// [AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIPConfiguration],
// [AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIpv6Configuration],
// [AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCidrConfiguration],
// [AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTASNConfiguration],
// [AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCountryConfiguration].
type AccountFirewallAccessRuleRuleUpdateParamsConfiguration interface {
	implementsAccountFirewallAccessRuleRuleUpdateParamsConfiguration()
}

type AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIPConfiguration) implementsAccountFirewallAccessRuleRuleUpdateParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIPConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIPConfigurationTargetIP AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIPConfigurationTarget = "ip"
)

type AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIpv6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIpv6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIpv6Configuration) implementsAccountFirewallAccessRuleRuleUpdateParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIpv6ConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIpv6ConfigurationTargetIp6 AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTIpv6ConfigurationTarget = "ip6"
)

type AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCidrConfiguration) implementsAccountFirewallAccessRuleRuleUpdateParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCidrConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCidrConfigurationTargetIPRange AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCidrConfigurationTarget = "ip_range"
)

type AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTASNConfiguration) implementsAccountFirewallAccessRuleRuleUpdateParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTASNConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTASNConfigurationTargetASN AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTASNConfigurationTarget = "asn"
)

type AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCountryConfiguration) implementsAccountFirewallAccessRuleRuleUpdateParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCountryConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCountryConfigurationTargetCountry AccountFirewallAccessRuleRuleUpdateParamsConfigurationDZw70ubTCountryConfigurationTarget = "country"
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
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration],
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6Configuration],
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfiguration],
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfiguration],
// [AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfiguration].
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration interface {
	implementsAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration()
}

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration) implementsAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTargetIP AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTarget = "ip"
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6Configuration) implementsAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6ConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6ConfigurationTargetIp6 AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6ConfigurationTarget = "ip6"
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfiguration) implementsAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfigurationTargetIPRange AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfigurationTarget = "ip_range"
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfiguration) implementsAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfigurationTargetASN AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfigurationTarget = "asn"
)

type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfiguration) implementsAccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfigurationTargetCountry AccountFirewallAccessRuleRuleIPAccessRulesForAnAccountNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfigurationTarget = "country"
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
