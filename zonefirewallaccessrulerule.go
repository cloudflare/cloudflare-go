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

// ZoneFirewallAccessRuleRuleService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneFirewallAccessRuleRuleService] method instead.
type ZoneFirewallAccessRuleRuleService struct {
	Options []option.RequestOption
}

// NewZoneFirewallAccessRuleRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneFirewallAccessRuleRuleService(opts ...option.RequestOption) (r *ZoneFirewallAccessRuleRuleService) {
	r = &ZoneFirewallAccessRuleRuleService{}
	r.Options = opts
	return
}

// Updates an IP Access rule defined at the zone level. You can only update the
// rule action (`mode` parameter) and notes.
func (r *ZoneFirewallAccessRuleRuleService) Update(ctx context.Context, zoneID string, identifier string, body ZoneFirewallAccessRuleRuleUpdateParams, opts ...option.RequestOption) (res *RuleSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/access_rules/rules/%s", zoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an IP Access rule defined at the zone level.
//
// Optionally, you can use the `cascade` property to specify that you wish to
// delete similar rules in other zones managed by the same zone owner.
func (r *ZoneFirewallAccessRuleRuleService) Delete(ctx context.Context, zoneID string, identifier string, body ZoneFirewallAccessRuleRuleDeleteParams, opts ...option.RequestOption) (res *RuleSingleIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/access_rules/rules/%s", zoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Creates a new IP Access rule for a zone.
//
// Note: To create an IP Access rule that applies to multiple zones, refer to
// [IP Access rules for a user](#ip-access-rules-for-a-user) or
// [IP Access rules for an account](#ip-access-rules-for-an-account) as
// appropriate.
func (r *ZoneFirewallAccessRuleRuleService) IPAccessRulesForAZoneNewAnIPAccessRule(ctx context.Context, zoneID string, body ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParams, opts ...option.RequestOption) (res *RuleSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/access_rules/rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches IP Access rules of a zone. You can filter the results using several
// optional parameters.
func (r *ZoneFirewallAccessRuleRuleService) IPAccessRulesForAZoneListIPAccessRules(ctx context.Context, zoneID string, query ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParams, opts ...option.RequestOption) (res *RuleCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/access_rules/rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneFirewallAccessRuleRuleUpdateParams struct {
	// The action to apply to a matched request.
	Mode param.Field[ZoneFirewallAccessRuleRuleUpdateParamsMode] `json:"mode"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r ZoneFirewallAccessRuleRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to apply to a matched request.
type ZoneFirewallAccessRuleRuleUpdateParamsMode string

const (
	ZoneFirewallAccessRuleRuleUpdateParamsModeBlock            ZoneFirewallAccessRuleRuleUpdateParamsMode = "block"
	ZoneFirewallAccessRuleRuleUpdateParamsModeChallenge        ZoneFirewallAccessRuleRuleUpdateParamsMode = "challenge"
	ZoneFirewallAccessRuleRuleUpdateParamsModeWhitelist        ZoneFirewallAccessRuleRuleUpdateParamsMode = "whitelist"
	ZoneFirewallAccessRuleRuleUpdateParamsModeJsChallenge      ZoneFirewallAccessRuleRuleUpdateParamsMode = "js_challenge"
	ZoneFirewallAccessRuleRuleUpdateParamsModeManagedChallenge ZoneFirewallAccessRuleRuleUpdateParamsMode = "managed_challenge"
)

type ZoneFirewallAccessRuleRuleDeleteParams struct {
	// The level to attempt to delete similar rules defined for other zones with the
	// same owner. The default value is `none`, which will only delete the current
	// rule. Using `basic` will delete rules that match the same action (mode) and
	// configuration, while using `aggressive` will delete rules that match the same
	// configuration.
	Cascade param.Field[ZoneFirewallAccessRuleRuleDeleteParamsCascade] `json:"cascade"`
}

func (r ZoneFirewallAccessRuleRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The level to attempt to delete similar rules defined for other zones with the
// same owner. The default value is `none`, which will only delete the current
// rule. Using `basic` will delete rules that match the same action (mode) and
// configuration, while using `aggressive` will delete rules that match the same
// configuration.
type ZoneFirewallAccessRuleRuleDeleteParamsCascade string

const (
	ZoneFirewallAccessRuleRuleDeleteParamsCascadeNone       ZoneFirewallAccessRuleRuleDeleteParamsCascade = "none"
	ZoneFirewallAccessRuleRuleDeleteParamsCascadeBasic      ZoneFirewallAccessRuleRuleDeleteParamsCascade = "basic"
	ZoneFirewallAccessRuleRuleDeleteParamsCascadeAggressive ZoneFirewallAccessRuleRuleDeleteParamsCascade = "aggressive"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParams struct {
	// The rule configuration.
	Configuration param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes,required"`
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
//
// Satisfied by
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIPConfiguration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIpv6Configuration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCidrConfiguration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationASNConfiguration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCountryConfiguration].
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration interface {
	implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration()
}

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIPConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIPConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIPConfigurationTargetIP ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIPConfigurationTarget = "ip"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIpv6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIpv6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIpv6Configuration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIpv6ConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIpv6ConfigurationTargetIp6 ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationIpv6ConfigurationTarget = "ip6"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCidrConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCidrConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCidrConfigurationTargetIPRange ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCidrConfigurationTarget = "ip_range"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationASNConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationASNConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationASNConfigurationTargetASN ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationASNConfigurationTarget = "asn"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCountryConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCountryConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCountryConfigurationTargetCountry ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsMode string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsModeBlock            ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsMode = "block"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsModeChallenge        ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsMode = "challenge"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsModeWhitelist        ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsMode = "whitelist"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsModeJsChallenge      ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsMode = "js_challenge"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsModeManagedChallenge ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsMode = "managed_challenge"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParams struct {
	// The direction used to sort returned rules.
	Direction     param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsDirection]     `query:"direction"`
	EgsPagination param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsEgsPagination] `query:"egs-pagination"`
	Filters       param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFilters]       `query:"filters"`
	// The field used to sort returned rules.
	Order param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsOrder] `query:"order"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParams]'s query
// parameters as `url.Values`.
func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsDirection string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsDirectionAsc  ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsDirection = "asc"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsDirectionDesc ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsDirection = "desc"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsEgsPagination struct {
	Json param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsEgsPaginationJson] `query:"json"`
}

// URLQuery serializes
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsEgsPagination]'s
// query parameters as `url.Values`.
func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsEgsPagination) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsEgsPaginationJson struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsEgsPaginationJson]'s
// query parameters as `url.Values`.
func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsEgsPaginationJson) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFilters struct {
	// The target to search in existing rules.
	ConfigurationTarget param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersConfigurationTarget] `query:"configuration.target"`
	// The target value to search for in existing rules: an IP address, an IP address
	// range, or a country code, depending on the provided `configuration.target`.
	// Notes: You can search for a single IPv4 address, an IP address range with a
	// subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	ConfigurationValue param.Field[string] `query:"configuration.value"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMode] `query:"mode"`
	// The string to search for in the notes of existing IP Access rules. Notes: For
	// example, the string 'attack' would match IP Access rules with notes 'Attack
	// 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
}

// URLQuery serializes
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFilters]'s
// query parameters as `url.Values`.
func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFilters) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The target to search in existing rules.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersConfigurationTargetIP      ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersConfigurationTarget = "ip"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersConfigurationTargetIPRange ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersConfigurationTarget = "ip_range"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersConfigurationTargetASN     ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersConfigurationTarget = "asn"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersConfigurationTargetCountry ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersConfigurationTarget = "country"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMatch string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMatchAny ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMatch = "any"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMatchAll ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMatch = "all"
)

// The action to apply to a matched request.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMode string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersModeBlock            ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMode = "block"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersModeChallenge        ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMode = "challenge"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersModeWhitelist        ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMode = "whitelist"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersModeJsChallenge      ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMode = "js_challenge"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersModeManagedChallenge ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsFiltersMode = "managed_challenge"
)

// The field used to sort returned rules.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsOrder string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsOrderConfigurationTarget ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsOrder = "configuration.target"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsOrderConfigurationValue  ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsOrder = "configuration.value"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsOrderMode                ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParamsOrder = "mode"
)
