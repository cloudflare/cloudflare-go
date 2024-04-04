// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AccessRuleService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccessRuleService] method instead.
type AccessRuleService struct {
	Options []option.RequestOption
}

// NewAccessRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessRuleService(opts ...option.RequestOption) (r *AccessRuleService) {
	r = &AccessRuleService{}
	r.Options = opts
	return
}

// Creates a new IP Access rule for an account or zone. The rule will apply to all
// zones in the account or zone.
//
// Note: To create an IP Access rule that applies to a single zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *AccessRuleService) New(ctx context.Context, params AccessRuleNewParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef173Union, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessRuleNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches IP Access rules of an account or zone. These rules apply to all the
// zones in the account or zone. You can filter the results using several optional
// parameters.
func (r *AccessRuleService) List(ctx context.Context, params AccessRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessRuleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Fetches IP Access rules of an account or zone. These rules apply to all the
// zones in the account or zone. You can filter the results using several optional
// parameters.
func (r *AccessRuleService) ListAutoPaging(ctx context.Context, params AccessRuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessRuleListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing IP Access rule defined.
//
// Note: This operation will affect all zones in the account or zone.
func (r *AccessRuleService) Delete(ctx context.Context, identifier interface{}, params AccessRuleDeleteParams, opts ...option.RequestOption) (res *AccessRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessRuleDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules/%v", accountOrZone, accountOrZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an IP Access rule defined.
//
// Note: This operation will affect all zones in the account or zone.
func (r *AccessRuleService) Edit(ctx context.Context, identifier interface{}, params AccessRuleEditParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef173Union, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessRuleEditResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules/%v", accountOrZone, accountOrZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of an IP Access rule defined.
func (r *AccessRuleService) Get(ctx context.Context, identifier interface{}, query AccessRuleGetParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef173Union, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessRuleGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules/%v", accountOrZone, accountOrZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessRuleListResponse = interface{}

type AccessRuleDeleteResponse struct {
	// Identifier
	ID   string                       `json:"id,required"`
	JSON accessRuleDeleteResponseJSON `json:"-"`
}

// accessRuleDeleteResponseJSON contains the JSON metadata for the struct
// [AccessRuleDeleteResponse]
type accessRuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[AccessRuleNewParamsConfigurationUnion] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[AccessRuleNewParamsMode] `json:"mode,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r AccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
type AccessRuleNewParamsConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccessRuleNewParamsConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleNewParamsConfiguration) implementsFirewallAccessRuleNewParamsConfigurationUnion() {}

// The rule configuration.
//
// Satisfied by
// [firewall.AccessRuleNewParamsConfigurationLegacyJhsIPConfiguration],
// [firewall.AccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration],
// [firewall.AccessRuleNewParamsConfigurationLegacyJhsCIDRConfiguration],
// [firewall.AccessRuleNewParamsConfigurationLegacyJhsASNConfiguration],
// [firewall.AccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration],
// [AccessRuleNewParamsConfiguration].
type AccessRuleNewParamsConfigurationUnion interface {
	implementsFirewallAccessRuleNewParamsConfigurationUnion()
}

type AccessRuleNewParamsConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleNewParamsConfigurationLegacyJhsIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleNewParamsConfigurationLegacyJhsIPConfiguration) implementsFirewallAccessRuleNewParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget string

const (
	AccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTargetIP AccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

func (r AccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTargetIP:
		return true
	}
	return false
}

type AccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[AccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration) implementsFirewallAccessRuleNewParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type AccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	AccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTargetIp6 AccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

func (r AccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTargetIp6:
		return true
	}
	return false
}

type AccessRuleNewParamsConfigurationLegacyJhsCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[AccessRuleNewParamsConfigurationLegacyJhsCIDRConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleNewParamsConfigurationLegacyJhsCIDRConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleNewParamsConfigurationLegacyJhsCIDRConfiguration) implementsFirewallAccessRuleNewParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type AccessRuleNewParamsConfigurationLegacyJhsCIDRConfigurationTarget string

const (
	AccessRuleNewParamsConfigurationLegacyJhsCIDRConfigurationTargetIPRange AccessRuleNewParamsConfigurationLegacyJhsCIDRConfigurationTarget = "ip_range"
)

func (r AccessRuleNewParamsConfigurationLegacyJhsCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleNewParamsConfigurationLegacyJhsCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

type AccessRuleNewParamsConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[AccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleNewParamsConfigurationLegacyJhsASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleNewParamsConfigurationLegacyJhsASNConfiguration) implementsFirewallAccessRuleNewParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type AccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget string

const (
	AccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTargetASN AccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

func (r AccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTargetASN:
		return true
	}
	return false
}

type AccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[AccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration) implementsFirewallAccessRuleNewParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type AccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget string

const (
	AccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTargetCountry AccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

func (r AccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTargetCountry:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleNewParamsConfigurationTarget string

const (
	AccessRuleNewParamsConfigurationTargetIP      AccessRuleNewParamsConfigurationTarget = "ip"
	AccessRuleNewParamsConfigurationTargetIp6     AccessRuleNewParamsConfigurationTarget = "ip6"
	AccessRuleNewParamsConfigurationTargetIPRange AccessRuleNewParamsConfigurationTarget = "ip_range"
	AccessRuleNewParamsConfigurationTargetASN     AccessRuleNewParamsConfigurationTarget = "asn"
	AccessRuleNewParamsConfigurationTargetCountry AccessRuleNewParamsConfigurationTarget = "country"
)

func (r AccessRuleNewParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleNewParamsConfigurationTargetIP, AccessRuleNewParamsConfigurationTargetIp6, AccessRuleNewParamsConfigurationTargetIPRange, AccessRuleNewParamsConfigurationTargetASN, AccessRuleNewParamsConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type AccessRuleNewParamsMode string

const (
	AccessRuleNewParamsModeBlock            AccessRuleNewParamsMode = "block"
	AccessRuleNewParamsModeChallenge        AccessRuleNewParamsMode = "challenge"
	AccessRuleNewParamsModeWhitelist        AccessRuleNewParamsMode = "whitelist"
	AccessRuleNewParamsModeJsChallenge      AccessRuleNewParamsMode = "js_challenge"
	AccessRuleNewParamsModeManagedChallenge AccessRuleNewParamsMode = "managed_challenge"
)

func (r AccessRuleNewParamsMode) IsKnown() bool {
	switch r {
	case AccessRuleNewParamsModeBlock, AccessRuleNewParamsModeChallenge, AccessRuleNewParamsModeWhitelist, AccessRuleNewParamsModeJsChallenge, AccessRuleNewParamsModeManagedChallenge:
		return true
	}
	return false
}

type AccessRuleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   shared.UnnamedSchemaRef173Union `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccessRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessRuleNewResponseEnvelopeJSON    `json:"-"`
}

// accessRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessRuleNewResponseEnvelope]
type accessRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessRuleNewResponseEnvelopeSuccess bool

const (
	AccessRuleNewResponseEnvelopeSuccessTrue AccessRuleNewResponseEnvelopeSuccess = true
)

func (r AccessRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessRuleListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The direction used to sort returned rules.
	Direction     param.Field[AccessRuleListParamsDirection]     `query:"direction"`
	EgsPagination param.Field[AccessRuleListParamsEgsPagination] `query:"egs-pagination"`
	Filters       param.Field[AccessRuleListParamsFilters]       `query:"filters"`
	// The field used to sort returned rules.
	Order param.Field[AccessRuleListParamsOrder] `query:"order"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccessRuleListParams]'s query parameters as `url.Values`.
func (r AccessRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type AccessRuleListParamsDirection string

const (
	AccessRuleListParamsDirectionAsc  AccessRuleListParamsDirection = "asc"
	AccessRuleListParamsDirectionDesc AccessRuleListParamsDirection = "desc"
)

func (r AccessRuleListParamsDirection) IsKnown() bool {
	switch r {
	case AccessRuleListParamsDirectionAsc, AccessRuleListParamsDirectionDesc:
		return true
	}
	return false
}

type AccessRuleListParamsEgsPagination struct {
	Json param.Field[AccessRuleListParamsEgsPaginationJson] `query:"json"`
}

// URLQuery serializes [AccessRuleListParamsEgsPagination]'s query parameters as
// `url.Values`.
func (r AccessRuleListParamsEgsPagination) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccessRuleListParamsEgsPaginationJson struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccessRuleListParamsEgsPaginationJson]'s query parameters
// as `url.Values`.
func (r AccessRuleListParamsEgsPaginationJson) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccessRuleListParamsFilters struct {
	// The target to search in existing rules.
	ConfigurationTarget param.Field[AccessRuleListParamsFiltersConfigurationTarget] `query:"configuration.target"`
	// The target value to search for in existing rules: an IP address, an IP address
	// range, or a country code, depending on the provided `configuration.target`.
	// Notes: You can search for a single IPv4 address, an IP address range with a
	// subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	ConfigurationValue param.Field[string] `query:"configuration.value"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[AccessRuleListParamsFiltersMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[AccessRuleListParamsFiltersMode] `query:"mode"`
	// The string to search for in the notes of existing IP Access rules. Notes: For
	// example, the string 'attack' would match IP Access rules with notes 'Attack
	// 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
}

// URLQuery serializes [AccessRuleListParamsFilters]'s query parameters as
// `url.Values`.
func (r AccessRuleListParamsFilters) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The target to search in existing rules.
type AccessRuleListParamsFiltersConfigurationTarget string

const (
	AccessRuleListParamsFiltersConfigurationTargetIP      AccessRuleListParamsFiltersConfigurationTarget = "ip"
	AccessRuleListParamsFiltersConfigurationTargetIPRange AccessRuleListParamsFiltersConfigurationTarget = "ip_range"
	AccessRuleListParamsFiltersConfigurationTargetASN     AccessRuleListParamsFiltersConfigurationTarget = "asn"
	AccessRuleListParamsFiltersConfigurationTargetCountry AccessRuleListParamsFiltersConfigurationTarget = "country"
)

func (r AccessRuleListParamsFiltersConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleListParamsFiltersConfigurationTargetIP, AccessRuleListParamsFiltersConfigurationTargetIPRange, AccessRuleListParamsFiltersConfigurationTargetASN, AccessRuleListParamsFiltersConfigurationTargetCountry:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type AccessRuleListParamsFiltersMatch string

const (
	AccessRuleListParamsFiltersMatchAny AccessRuleListParamsFiltersMatch = "any"
	AccessRuleListParamsFiltersMatchAll AccessRuleListParamsFiltersMatch = "all"
)

func (r AccessRuleListParamsFiltersMatch) IsKnown() bool {
	switch r {
	case AccessRuleListParamsFiltersMatchAny, AccessRuleListParamsFiltersMatchAll:
		return true
	}
	return false
}

// The action to apply to a matched request.
type AccessRuleListParamsFiltersMode string

const (
	AccessRuleListParamsFiltersModeBlock            AccessRuleListParamsFiltersMode = "block"
	AccessRuleListParamsFiltersModeChallenge        AccessRuleListParamsFiltersMode = "challenge"
	AccessRuleListParamsFiltersModeWhitelist        AccessRuleListParamsFiltersMode = "whitelist"
	AccessRuleListParamsFiltersModeJsChallenge      AccessRuleListParamsFiltersMode = "js_challenge"
	AccessRuleListParamsFiltersModeManagedChallenge AccessRuleListParamsFiltersMode = "managed_challenge"
)

func (r AccessRuleListParamsFiltersMode) IsKnown() bool {
	switch r {
	case AccessRuleListParamsFiltersModeBlock, AccessRuleListParamsFiltersModeChallenge, AccessRuleListParamsFiltersModeWhitelist, AccessRuleListParamsFiltersModeJsChallenge, AccessRuleListParamsFiltersModeManagedChallenge:
		return true
	}
	return false
}

// The field used to sort returned rules.
type AccessRuleListParamsOrder string

const (
	AccessRuleListParamsOrderConfigurationTarget AccessRuleListParamsOrder = "configuration.target"
	AccessRuleListParamsOrderConfigurationValue  AccessRuleListParamsOrder = "configuration.value"
	AccessRuleListParamsOrderMode                AccessRuleListParamsOrder = "mode"
)

func (r AccessRuleListParamsOrder) IsKnown() bool {
	switch r {
	case AccessRuleListParamsOrderConfigurationTarget, AccessRuleListParamsOrderConfigurationValue, AccessRuleListParamsOrderMode:
		return true
	}
	return false
}

type AccessRuleDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r AccessRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccessRuleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   AccessRuleDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccessRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessRuleDeleteResponseEnvelope]
type accessRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessRuleDeleteResponseEnvelopeSuccess bool

const (
	AccessRuleDeleteResponseEnvelopeSuccessTrue AccessRuleDeleteResponseEnvelopeSuccess = true
)

func (r AccessRuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessRuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessRuleEditParams struct {
	// The rule configuration.
	Configuration param.Field[AccessRuleEditParamsConfigurationUnion] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[AccessRuleEditParamsMode] `json:"mode,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r AccessRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
type AccessRuleEditParamsConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccessRuleEditParamsConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleEditParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleEditParamsConfiguration) implementsFirewallAccessRuleEditParamsConfigurationUnion() {
}

// The rule configuration.
//
// Satisfied by
// [firewall.AccessRuleEditParamsConfigurationLegacyJhsIPConfiguration],
// [firewall.AccessRuleEditParamsConfigurationLegacyJhsIPV6Configuration],
// [firewall.AccessRuleEditParamsConfigurationLegacyJhsCIDRConfiguration],
// [firewall.AccessRuleEditParamsConfigurationLegacyJhsASNConfiguration],
// [firewall.AccessRuleEditParamsConfigurationLegacyJhsCountryConfiguration],
// [AccessRuleEditParamsConfiguration].
type AccessRuleEditParamsConfigurationUnion interface {
	implementsFirewallAccessRuleEditParamsConfigurationUnion()
}

type AccessRuleEditParamsConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccessRuleEditParamsConfigurationLegacyJhsIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleEditParamsConfigurationLegacyJhsIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleEditParamsConfigurationLegacyJhsIPConfiguration) implementsFirewallAccessRuleEditParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleEditParamsConfigurationLegacyJhsIPConfigurationTarget string

const (
	AccessRuleEditParamsConfigurationLegacyJhsIPConfigurationTargetIP AccessRuleEditParamsConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

func (r AccessRuleEditParamsConfigurationLegacyJhsIPConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleEditParamsConfigurationLegacyJhsIPConfigurationTargetIP:
		return true
	}
	return false
}

type AccessRuleEditParamsConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[AccessRuleEditParamsConfigurationLegacyJhsIPV6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleEditParamsConfigurationLegacyJhsIPV6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleEditParamsConfigurationLegacyJhsIPV6Configuration) implementsFirewallAccessRuleEditParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type AccessRuleEditParamsConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	AccessRuleEditParamsConfigurationLegacyJhsIPV6ConfigurationTargetIp6 AccessRuleEditParamsConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

func (r AccessRuleEditParamsConfigurationLegacyJhsIPV6ConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleEditParamsConfigurationLegacyJhsIPV6ConfigurationTargetIp6:
		return true
	}
	return false
}

type AccessRuleEditParamsConfigurationLegacyJhsCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[AccessRuleEditParamsConfigurationLegacyJhsCIDRConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleEditParamsConfigurationLegacyJhsCIDRConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleEditParamsConfigurationLegacyJhsCIDRConfiguration) implementsFirewallAccessRuleEditParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type AccessRuleEditParamsConfigurationLegacyJhsCIDRConfigurationTarget string

const (
	AccessRuleEditParamsConfigurationLegacyJhsCIDRConfigurationTargetIPRange AccessRuleEditParamsConfigurationLegacyJhsCIDRConfigurationTarget = "ip_range"
)

func (r AccessRuleEditParamsConfigurationLegacyJhsCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleEditParamsConfigurationLegacyJhsCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

type AccessRuleEditParamsConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[AccessRuleEditParamsConfigurationLegacyJhsASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleEditParamsConfigurationLegacyJhsASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleEditParamsConfigurationLegacyJhsASNConfiguration) implementsFirewallAccessRuleEditParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type AccessRuleEditParamsConfigurationLegacyJhsASNConfigurationTarget string

const (
	AccessRuleEditParamsConfigurationLegacyJhsASNConfigurationTargetASN AccessRuleEditParamsConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

func (r AccessRuleEditParamsConfigurationLegacyJhsASNConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleEditParamsConfigurationLegacyJhsASNConfigurationTargetASN:
		return true
	}
	return false
}

type AccessRuleEditParamsConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[AccessRuleEditParamsConfigurationLegacyJhsCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleEditParamsConfigurationLegacyJhsCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleEditParamsConfigurationLegacyJhsCountryConfiguration) implementsFirewallAccessRuleEditParamsConfigurationUnion() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type AccessRuleEditParamsConfigurationLegacyJhsCountryConfigurationTarget string

const (
	AccessRuleEditParamsConfigurationLegacyJhsCountryConfigurationTargetCountry AccessRuleEditParamsConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

func (r AccessRuleEditParamsConfigurationLegacyJhsCountryConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleEditParamsConfigurationLegacyJhsCountryConfigurationTargetCountry:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleEditParamsConfigurationTarget string

const (
	AccessRuleEditParamsConfigurationTargetIP      AccessRuleEditParamsConfigurationTarget = "ip"
	AccessRuleEditParamsConfigurationTargetIp6     AccessRuleEditParamsConfigurationTarget = "ip6"
	AccessRuleEditParamsConfigurationTargetIPRange AccessRuleEditParamsConfigurationTarget = "ip_range"
	AccessRuleEditParamsConfigurationTargetASN     AccessRuleEditParamsConfigurationTarget = "asn"
	AccessRuleEditParamsConfigurationTargetCountry AccessRuleEditParamsConfigurationTarget = "country"
)

func (r AccessRuleEditParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleEditParamsConfigurationTargetIP, AccessRuleEditParamsConfigurationTargetIp6, AccessRuleEditParamsConfigurationTargetIPRange, AccessRuleEditParamsConfigurationTargetASN, AccessRuleEditParamsConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type AccessRuleEditParamsMode string

const (
	AccessRuleEditParamsModeBlock            AccessRuleEditParamsMode = "block"
	AccessRuleEditParamsModeChallenge        AccessRuleEditParamsMode = "challenge"
	AccessRuleEditParamsModeWhitelist        AccessRuleEditParamsMode = "whitelist"
	AccessRuleEditParamsModeJsChallenge      AccessRuleEditParamsMode = "js_challenge"
	AccessRuleEditParamsModeManagedChallenge AccessRuleEditParamsMode = "managed_challenge"
)

func (r AccessRuleEditParamsMode) IsKnown() bool {
	switch r {
	case AccessRuleEditParamsModeBlock, AccessRuleEditParamsModeChallenge, AccessRuleEditParamsModeWhitelist, AccessRuleEditParamsModeJsChallenge, AccessRuleEditParamsModeManagedChallenge:
		return true
	}
	return false
}

type AccessRuleEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   shared.UnnamedSchemaRef173Union `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccessRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessRuleEditResponseEnvelopeJSON    `json:"-"`
}

// accessRuleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessRuleEditResponseEnvelope]
type accessRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessRuleEditResponseEnvelopeSuccess bool

const (
	AccessRuleEditResponseEnvelopeSuccessTrue AccessRuleEditResponseEnvelopeSuccess = true
)

func (r AccessRuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessRuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessRuleGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessRuleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   shared.UnnamedSchemaRef173Union `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccessRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessRuleGetResponseEnvelopeJSON    `json:"-"`
}

// accessRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessRuleGetResponseEnvelope]
type accessRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessRuleGetResponseEnvelopeSuccess bool

const (
	AccessRuleGetResponseEnvelopeSuccessTrue AccessRuleGetResponseEnvelopeSuccess = true
)

func (r AccessRuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessRuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
