// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

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

// Creates a new IP Access rule for an account or zone. The rule will apply to all
// zones in the account or zone.
//
// Note: To create an IP Access rule that applies to a single zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *FirewallAccessRuleService) New(ctx context.Context, params FirewallAccessRuleNewParams, opts ...option.RequestOption) (res *FirewallAccessRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAccessRuleNewResponseEnvelope
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
func (r *FirewallAccessRuleService) List(ctx context.Context, params FirewallAccessRuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[FirewallAccessRuleListResponse], err error) {
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
func (r *FirewallAccessRuleService) ListAutoPaging(ctx context.Context, params FirewallAccessRuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[FirewallAccessRuleListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing IP Access rule defined.
//
// Note: This operation will affect all zones in the account or zone.
func (r *FirewallAccessRuleService) Delete(ctx context.Context, identifier interface{}, body FirewallAccessRuleDeleteParams, opts ...option.RequestOption) (res *FirewallAccessRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAccessRuleDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules/%v", accountOrZone, accountOrZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an IP Access rule defined at the account level.
//
// Note: This operation will affect all zones in the account.
func (r *FirewallAccessRuleService) Edit(ctx context.Context, identifier interface{}, params FirewallAccessRuleEditParams, opts ...option.RequestOption) (res *FirewallAccessRuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAccessRuleEditResponseEnvelope
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
func (r *FirewallAccessRuleService) Get(ctx context.Context, identifier interface{}, query FirewallAccessRuleGetParams, opts ...option.RequestOption) (res *FirewallAccessRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAccessRuleGetResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [FirewallAccessRuleNewResponseUnknown] or
// [shared.UnionString].
type FirewallAccessRuleNewResponse interface {
	ImplementsFirewallAccessRuleNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallAccessRuleNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallAccessRuleListResponse = interface{}

type FirewallAccessRuleDeleteResponse struct {
	// Identifier
	ID   string                               `json:"id,required"`
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

// Union satisfied by [FirewallAccessRuleEditResponseUnknown] or
// [shared.UnionString].
type FirewallAccessRuleEditResponse interface {
	ImplementsFirewallAccessRuleEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallAccessRuleEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [FirewallAccessRuleGetResponseUnknown] or
// [shared.UnionString].
type FirewallAccessRuleGetResponse interface {
	ImplementsFirewallAccessRuleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallAccessRuleGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallAccessRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[FirewallAccessRuleNewParamsConfiguration] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallAccessRuleNewParamsMode] `json:"mode,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r FirewallAccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
//
// Satisfied by [FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration],
// [FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration],
// [FirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfiguration],
// [FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration],
// [FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration].
type FirewallAccessRuleNewParamsConfiguration interface {
	implementsFirewallAccessRuleNewParamsConfiguration()
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

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration) implementsFirewallAccessRuleNewParamsConfiguration() {
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

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration) implementsFirewallAccessRuleNewParamsConfiguration() {
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

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsCidrConfiguration) implementsFirewallAccessRuleNewParamsConfiguration() {
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

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration) implementsFirewallAccessRuleNewParamsConfiguration() {
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

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration) implementsFirewallAccessRuleNewParamsConfiguration() {
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
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
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

type FirewallAccessRuleDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

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
	// The rule configuration.
	Configuration param.Field[FirewallAccessRuleEditParamsConfiguration] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallAccessRuleEditParamsMode] `json:"mode,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r FirewallAccessRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
//
// Satisfied by
// [FirewallAccessRuleEditParamsConfigurationLegacyJhsIPConfiguration],
// [FirewallAccessRuleEditParamsConfigurationLegacyJhsIPV6Configuration],
// [FirewallAccessRuleEditParamsConfigurationLegacyJhsCidrConfiguration],
// [FirewallAccessRuleEditParamsConfigurationLegacyJhsASNConfiguration],
// [FirewallAccessRuleEditParamsConfigurationLegacyJhsCountryConfiguration].
type FirewallAccessRuleEditParamsConfiguration interface {
	implementsFirewallAccessRuleEditParamsConfiguration()
}

type FirewallAccessRuleEditParamsConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[FirewallAccessRuleEditParamsConfigurationLegacyJhsIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleEditParamsConfigurationLegacyJhsIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleEditParamsConfigurationLegacyJhsIPConfiguration) implementsFirewallAccessRuleEditParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleEditParamsConfigurationLegacyJhsIPConfigurationTarget string

const (
	FirewallAccessRuleEditParamsConfigurationLegacyJhsIPConfigurationTargetIP FirewallAccessRuleEditParamsConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

type FirewallAccessRuleEditParamsConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[FirewallAccessRuleEditParamsConfigurationLegacyJhsIPV6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleEditParamsConfigurationLegacyJhsIPV6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleEditParamsConfigurationLegacyJhsIPV6Configuration) implementsFirewallAccessRuleEditParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type FirewallAccessRuleEditParamsConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	FirewallAccessRuleEditParamsConfigurationLegacyJhsIPV6ConfigurationTargetIp6 FirewallAccessRuleEditParamsConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

type FirewallAccessRuleEditParamsConfigurationLegacyJhsCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[FirewallAccessRuleEditParamsConfigurationLegacyJhsCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleEditParamsConfigurationLegacyJhsCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleEditParamsConfigurationLegacyJhsCidrConfiguration) implementsFirewallAccessRuleEditParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type FirewallAccessRuleEditParamsConfigurationLegacyJhsCidrConfigurationTarget string

const (
	FirewallAccessRuleEditParamsConfigurationLegacyJhsCidrConfigurationTargetIPRange FirewallAccessRuleEditParamsConfigurationLegacyJhsCidrConfigurationTarget = "ip_range"
)

type FirewallAccessRuleEditParamsConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[FirewallAccessRuleEditParamsConfigurationLegacyJhsASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleEditParamsConfigurationLegacyJhsASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleEditParamsConfigurationLegacyJhsASNConfiguration) implementsFirewallAccessRuleEditParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type FirewallAccessRuleEditParamsConfigurationLegacyJhsASNConfigurationTarget string

const (
	FirewallAccessRuleEditParamsConfigurationLegacyJhsASNConfigurationTargetASN FirewallAccessRuleEditParamsConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

type FirewallAccessRuleEditParamsConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[FirewallAccessRuleEditParamsConfigurationLegacyJhsCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleEditParamsConfigurationLegacyJhsCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleEditParamsConfigurationLegacyJhsCountryConfiguration) implementsFirewallAccessRuleEditParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type FirewallAccessRuleEditParamsConfigurationLegacyJhsCountryConfigurationTarget string

const (
	FirewallAccessRuleEditParamsConfigurationLegacyJhsCountryConfigurationTargetCountry FirewallAccessRuleEditParamsConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

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

type FirewallAccessRuleGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type FirewallAccessRuleGetResponseEnvelope struct {
	Errors   []FirewallAccessRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallAccessRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallAccessRuleGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallAccessRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallAccessRuleGetResponseEnvelopeJSON    `json:"-"`
}

// firewallAccessRuleGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallAccessRuleGetResponseEnvelope]
type firewallAccessRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    firewallAccessRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallAccessRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallAccessRuleGetResponseEnvelopeErrors]
type firewallAccessRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    firewallAccessRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallAccessRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallAccessRuleGetResponseEnvelopeMessages]
type firewallAccessRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleGetResponseEnvelopeSuccess bool

const (
	FirewallAccessRuleGetResponseEnvelopeSuccessTrue FirewallAccessRuleGetResponseEnvelopeSuccess = true
)
