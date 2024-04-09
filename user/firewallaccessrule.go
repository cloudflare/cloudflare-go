// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
  "context"
  "fmt"
  "net/http"
  "net/url"
  "reflect"
  "time"

  "github.com/cloudflare/cloudflare-go/v2/firewall"
  "github.com/cloudflare/cloudflare-go/v2/internal/apijson"
  "github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
  "github.com/cloudflare/cloudflare-go/v2/internal/pagination"
  "github.com/cloudflare/cloudflare-go/v2/internal/param"
  "github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
  "github.com/cloudflare/cloudflare-go/v2/internal/shared"
  "github.com/cloudflare/cloudflare-go/v2/option"
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

// Creates a new IP Access rule for all zones owned by the current user.
//
// Note: To create an IP Access rule that applies to a specific zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *FirewallAccessRuleService) New(ctx context.Context, body FirewallAccessRuleNewParams, opts ...option.RequestOption) (res *FirewallRule, err error) {
  opts = append(r.Options[:], opts...)
  var env FirewallAccessRuleNewResponseEnvelope
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
func (r *FirewallAccessRuleService) List(ctx context.Context, query FirewallAccessRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[FirewallRule], err error) {
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
func (r *FirewallAccessRuleService) ListAutoPaging(ctx context.Context, query FirewallAccessRuleListParams, opts ...option.RequestOption) (*pagination.V4PagePaginationArrayAutoPager[FirewallRule]) {
  return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Deletes an IP Access rule at the user level.
//
// Note: Deleting a user-level rule will affect all zones owned by the user.
func (r *FirewallAccessRuleService) Delete(ctx context.Context, identifier string, body FirewallAccessRuleDeleteParams, opts ...option.RequestOption) (res *FirewallAccessRuleDeleteResponse, err error) {
  opts = append(r.Options[:], opts...)
  var env FirewallAccessRuleDeleteResponseEnvelope
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
func (r *FirewallAccessRuleService) Edit(ctx context.Context, identifier string, body FirewallAccessRuleEditParams, opts ...option.RequestOption) (res *FirewallRule, err error) {
  opts = append(r.Options[:], opts...)
  var env FirewallAccessRuleEditResponseEnvelope
  path := fmt.Sprintf("user/firewall/access_rules/rules/%s", identifier)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

type FirewallRule struct {
// The unique identifier of the IP Access rule.
ID string `json:"id,required"`
// The available actions that a rule can apply to a matched request.
AllowedModes []FirewallRuleAllowedMode `json:"allowed_modes,required"`
// The rule configuration.
Configuration FirewallRuleConfiguration `json:"configuration,required"`
// The action to apply to a matched request.
Mode FirewallRuleMode `json:"mode,required"`
// The timestamp of when the rule was created.
CreatedOn time.Time `json:"created_on" format:"date-time"`
// The timestamp of when the rule was last modified.
ModifiedOn time.Time `json:"modified_on" format:"date-time"`
// An informative summary of the rule, typically used as a reminder or explanation.
Notes string `json:"notes"`
JSON firewallRuleJSON `json:"-"`
}

// firewallRuleJSON contains the JSON metadata for the struct [FirewallRule]
type firewallRuleJSON struct {
ID apijson.Field
AllowedModes apijson.Field
Configuration apijson.Field
Mode apijson.Field
CreatedOn apijson.Field
ModifiedOn apijson.Field
Notes apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *FirewallRule) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleJSON) RawJSON() (string) {
  return r.raw
}

// The action to apply to a matched request.
type FirewallRuleAllowedMode string

const (
  FirewallRuleAllowedModeBlock FirewallRuleAllowedMode = "block"
  FirewallRuleAllowedModeChallenge FirewallRuleAllowedMode = "challenge"
  FirewallRuleAllowedModeWhitelist FirewallRuleAllowedMode = "whitelist"
  FirewallRuleAllowedModeJsChallenge FirewallRuleAllowedMode = "js_challenge"
  FirewallRuleAllowedModeManagedChallenge FirewallRuleAllowedMode = "managed_challenge"
)

func (r FirewallRuleAllowedMode) IsKnown() (bool) {
  switch r {
  case FirewallRuleAllowedModeBlock, FirewallRuleAllowedModeChallenge, FirewallRuleAllowedModeWhitelist, FirewallRuleAllowedModeJsChallenge, FirewallRuleAllowedModeManagedChallenge:
      return true
  }
  return false
}

// The rule configuration.
type FirewallRuleConfiguration struct {
// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
Target firewall.FirewallRuleConfigurationTarget `json:"target"`
// The IP address to match. This address will be compared to the IP address of
// incoming requests.
Value string `json:"value"`
JSON firewallRuleConfigurationJSON `json:"-"`
union FirewallRuleConfigurationUnion
}

// firewallRuleConfigurationJSON contains the JSON metadata for the struct
// [FirewallRuleConfiguration]
type firewallRuleConfigurationJSON struct {
Target apijson.Field
Value apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r firewallRuleConfigurationJSON) RawJSON() (string) {
  return r.raw
}

func (r *FirewallRuleConfiguration) UnmarshalJSON(data []byte) (err error) {
  err = apijson.UnmarshalRoot(data, &r.union)
  if err != nil {
    return err
  }
  return apijson.Port(r.union, &r)
}

func (r FirewallRuleConfiguration) AsUnion() (FirewallRuleConfigurationUnion) {
  return r.union
}

// The rule configuration.
//
// Union satisfied by [firewall.IPConfiguration], [firewall.IPV6Configuration],
// [firewall.CIDRConfiguration], [firewall.ASNConfiguration] or
// [firewall.CountryConfiguration].
type FirewallRuleConfigurationUnion interface {
  implementsUserFirewallRuleConfiguration()
}

func init() {
  apijson.RegisterUnion(
    reflect.TypeOf((*FirewallRuleConfigurationUnion)(nil)).Elem(),
    "",
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(firewall.IPConfiguration{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(firewall.IPV6Configuration{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(firewall.CIDRConfiguration{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(firewall.ASNConfiguration{}),
    },
    apijson.UnionVariant{
      TypeFilter: gjson.JSON,
      Type: reflect.TypeOf(firewall.CountryConfiguration{}),
    },
  )
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type firewall.FirewallRuleConfigurationTarget string

const (
  firewall.FirewallRuleConfigurationTargetIP firewall.FirewallRuleConfigurationTarget = "ip"
  firewall.FirewallRuleConfigurationTargetIp6 firewall.FirewallRuleConfigurationTarget = "ip6"
  firewall.FirewallRuleConfigurationTargetIPRange firewall.FirewallRuleConfigurationTarget = "ip_range"
  firewall.FirewallRuleConfigurationTargetASN firewall.FirewallRuleConfigurationTarget = "asn"
  firewall.FirewallRuleConfigurationTargetCountry firewall.FirewallRuleConfigurationTarget = "country"
)

func (r firewall.FirewallRuleConfigurationTarget) IsKnown() (bool) {
  switch r {
  case firewall.FirewallRuleConfigurationTargetIP, firewall.FirewallRuleConfigurationTargetIp6, firewall.FirewallRuleConfigurationTargetIPRange, firewall.FirewallRuleConfigurationTargetASN, firewall.FirewallRuleConfigurationTargetCountry:
      return true
  }
  return false
}

// The action to apply to a matched request.
type FirewallRuleMode string

const (
  FirewallRuleModeBlock FirewallRuleMode = "block"
  FirewallRuleModeChallenge FirewallRuleMode = "challenge"
  FirewallRuleModeWhitelist FirewallRuleMode = "whitelist"
  FirewallRuleModeJsChallenge FirewallRuleMode = "js_challenge"
  FirewallRuleModeManagedChallenge FirewallRuleMode = "managed_challenge"
)

func (r FirewallRuleMode) IsKnown() (bool) {
  switch r {
  case FirewallRuleModeBlock, FirewallRuleModeChallenge, FirewallRuleModeWhitelist, FirewallRuleModeJsChallenge, FirewallRuleModeManagedChallenge:
      return true
  }
  return false
}

type FirewallAccessRuleDeleteResponse struct {
// The unique identifier of the IP Access rule.
ID string `json:"id"`
JSON firewallAccessRuleDeleteResponseJSON `json:"-"`
}

// firewallAccessRuleDeleteResponseJSON contains the JSON metadata for the struct
// [FirewallAccessRuleDeleteResponse]
type firewallAccessRuleDeleteResponseJSON struct {
ID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleDeleteResponseJSON) RawJSON() (string) {
  return r.raw
}

type FirewallAccessRuleNewParams struct {
// The rule configuration.
Configuration param.Field[FirewallAccessRuleNewParamsConfigurationUnion] `json:"configuration,required"`
// The action to apply to a matched request.
Mode param.Field[FirewallAccessRuleNewParamsMode] `json:"mode,required"`
// An informative summary of the rule, typically used as a reminder or explanation.
Notes param.Field[string] `json:"notes"`
}

func (r FirewallAccessRuleNewParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// The rule configuration.
type FirewallAccessRuleNewParamsConfiguration struct {
// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
Target param.Field[FirewallAccessRuleNewParamsConfigurationTarget] `json:"target"`
// The IP address to match. This address will be compared to the IP address of
// incoming requests.
Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfiguration) implementsUserFirewallAccessRuleNewParamsConfigurationUnion() {}

// The rule configuration.
//
// Satisfied by [firewall.IPConfigurationParam], [firewall.IPV6ConfigurationParam],
// [firewall.CIDRConfigurationParam], [firewall.ASNConfigurationParam],
// [firewall.CountryConfigurationParam],
// [FirewallAccessRuleNewParamsConfiguration].
type FirewallAccessRuleNewParamsConfigurationUnion interface {
  implementsUserFirewallAccessRuleNewParamsConfigurationUnion()
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleNewParamsConfigurationTarget string

const (
  FirewallAccessRuleNewParamsConfigurationTargetIP FirewallAccessRuleNewParamsConfigurationTarget = "ip"
  FirewallAccessRuleNewParamsConfigurationTargetIp6 FirewallAccessRuleNewParamsConfigurationTarget = "ip6"
  FirewallAccessRuleNewParamsConfigurationTargetIPRange FirewallAccessRuleNewParamsConfigurationTarget = "ip_range"
  FirewallAccessRuleNewParamsConfigurationTargetASN FirewallAccessRuleNewParamsConfigurationTarget = "asn"
  FirewallAccessRuleNewParamsConfigurationTargetCountry FirewallAccessRuleNewParamsConfigurationTarget = "country"
)

func (r FirewallAccessRuleNewParamsConfigurationTarget) IsKnown() (bool) {
  switch r {
  case FirewallAccessRuleNewParamsConfigurationTargetIP, FirewallAccessRuleNewParamsConfigurationTargetIp6, FirewallAccessRuleNewParamsConfigurationTargetIPRange, FirewallAccessRuleNewParamsConfigurationTargetASN, FirewallAccessRuleNewParamsConfigurationTargetCountry:
      return true
  }
  return false
}

// The action to apply to a matched request.
type FirewallAccessRuleNewParamsMode string

const (
  FirewallAccessRuleNewParamsModeBlock FirewallAccessRuleNewParamsMode = "block"
  FirewallAccessRuleNewParamsModeChallenge FirewallAccessRuleNewParamsMode = "challenge"
  FirewallAccessRuleNewParamsModeWhitelist FirewallAccessRuleNewParamsMode = "whitelist"
  FirewallAccessRuleNewParamsModeJsChallenge FirewallAccessRuleNewParamsMode = "js_challenge"
  FirewallAccessRuleNewParamsModeManagedChallenge FirewallAccessRuleNewParamsMode = "managed_challenge"
)

func (r FirewallAccessRuleNewParamsMode) IsKnown() (bool) {
  switch r {
  case FirewallAccessRuleNewParamsModeBlock, FirewallAccessRuleNewParamsModeChallenge, FirewallAccessRuleNewParamsModeWhitelist, FirewallAccessRuleNewParamsModeJsChallenge, FirewallAccessRuleNewParamsModeManagedChallenge:
      return true
  }
  return false
}

type FirewallAccessRuleNewResponseEnvelope struct {
Errors []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
Result FirewallRule `json:"result,required"`
// Whether the API call was successful
Success FirewallAccessRuleNewResponseEnvelopeSuccess `json:"success,required"`
JSON firewallAccessRuleNewResponseEnvelopeJSON `json:"-"`
}

// firewallAccessRuleNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallAccessRuleNewResponseEnvelope]
type firewallAccessRuleNewResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleNewResponseEnvelopeSuccess bool

const (
  FirewallAccessRuleNewResponseEnvelopeSuccessTrue FirewallAccessRuleNewResponseEnvelopeSuccess = true
)

func (r FirewallAccessRuleNewResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case FirewallAccessRuleNewResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}

type FirewallAccessRuleListParams struct {
// The direction used to sort returned rules.
Direction param.Field[FirewallAccessRuleListParamsDirection] `query:"direction"`
EgsPagination param.Field[FirewallAccessRuleListParamsEgsPagination] `query:"egs-pagination"`
Filters param.Field[FirewallAccessRuleListParamsFilters] `query:"filters"`
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
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

// The direction used to sort returned rules.
type FirewallAccessRuleListParamsDirection string

const (
  FirewallAccessRuleListParamsDirectionAsc FirewallAccessRuleListParamsDirection = "asc"
  FirewallAccessRuleListParamsDirectionDesc FirewallAccessRuleListParamsDirection = "desc"
)

func (r FirewallAccessRuleListParamsDirection) IsKnown() (bool) {
  switch r {
  case FirewallAccessRuleListParamsDirectionAsc, FirewallAccessRuleListParamsDirectionDesc:
      return true
  }
  return false
}

type FirewallAccessRuleListParamsEgsPagination struct {
Json param.Field[FirewallAccessRuleListParamsEgsPaginationJson] `query:"json"`
}

// URLQuery serializes [FirewallAccessRuleListParamsEgsPagination]'s query
// parameters as `url.Values`.
func (r FirewallAccessRuleListParamsEgsPagination) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
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
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
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
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

// The target to search in existing rules.
type FirewallAccessRuleListParamsFiltersConfigurationTarget string

const (
  FirewallAccessRuleListParamsFiltersConfigurationTargetIP FirewallAccessRuleListParamsFiltersConfigurationTarget = "ip"
  FirewallAccessRuleListParamsFiltersConfigurationTargetIPRange FirewallAccessRuleListParamsFiltersConfigurationTarget = "ip_range"
  FirewallAccessRuleListParamsFiltersConfigurationTargetASN FirewallAccessRuleListParamsFiltersConfigurationTarget = "asn"
  FirewallAccessRuleListParamsFiltersConfigurationTargetCountry FirewallAccessRuleListParamsFiltersConfigurationTarget = "country"
)

func (r FirewallAccessRuleListParamsFiltersConfigurationTarget) IsKnown() (bool) {
  switch r {
  case FirewallAccessRuleListParamsFiltersConfigurationTargetIP, FirewallAccessRuleListParamsFiltersConfigurationTargetIPRange, FirewallAccessRuleListParamsFiltersConfigurationTargetASN, FirewallAccessRuleListParamsFiltersConfigurationTargetCountry:
      return true
  }
  return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type FirewallAccessRuleListParamsFiltersMatch string

const (
  FirewallAccessRuleListParamsFiltersMatchAny FirewallAccessRuleListParamsFiltersMatch = "any"
  FirewallAccessRuleListParamsFiltersMatchAll FirewallAccessRuleListParamsFiltersMatch = "all"
)

func (r FirewallAccessRuleListParamsFiltersMatch) IsKnown() (bool) {
  switch r {
  case FirewallAccessRuleListParamsFiltersMatchAny, FirewallAccessRuleListParamsFiltersMatchAll:
      return true
  }
  return false
}

// The action to apply to a matched request.
type FirewallAccessRuleListParamsFiltersMode string

const (
  FirewallAccessRuleListParamsFiltersModeBlock FirewallAccessRuleListParamsFiltersMode = "block"
  FirewallAccessRuleListParamsFiltersModeChallenge FirewallAccessRuleListParamsFiltersMode = "challenge"
  FirewallAccessRuleListParamsFiltersModeWhitelist FirewallAccessRuleListParamsFiltersMode = "whitelist"
  FirewallAccessRuleListParamsFiltersModeJsChallenge FirewallAccessRuleListParamsFiltersMode = "js_challenge"
  FirewallAccessRuleListParamsFiltersModeManagedChallenge FirewallAccessRuleListParamsFiltersMode = "managed_challenge"
)

func (r FirewallAccessRuleListParamsFiltersMode) IsKnown() (bool) {
  switch r {
  case FirewallAccessRuleListParamsFiltersModeBlock, FirewallAccessRuleListParamsFiltersModeChallenge, FirewallAccessRuleListParamsFiltersModeWhitelist, FirewallAccessRuleListParamsFiltersModeJsChallenge, FirewallAccessRuleListParamsFiltersModeManagedChallenge:
      return true
  }
  return false
}

// The field used to sort returned rules.
type FirewallAccessRuleListParamsOrder string

const (
  FirewallAccessRuleListParamsOrderConfigurationTarget FirewallAccessRuleListParamsOrder = "configuration.target"
  FirewallAccessRuleListParamsOrderConfigurationValue FirewallAccessRuleListParamsOrder = "configuration.value"
  FirewallAccessRuleListParamsOrderMode FirewallAccessRuleListParamsOrder = "mode"
)

func (r FirewallAccessRuleListParamsOrder) IsKnown() (bool) {
  switch r {
  case FirewallAccessRuleListParamsOrderConfigurationTarget, FirewallAccessRuleListParamsOrderConfigurationValue, FirewallAccessRuleListParamsOrderMode:
      return true
  }
  return false
}

type FirewallAccessRuleDeleteParams struct {
Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallAccessRuleDeleteParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r.Body)
}

type FirewallAccessRuleDeleteResponseEnvelope struct {
Errors []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
Result FirewallAccessRuleDeleteResponse `json:"result,required"`
// Whether the API call was successful
Success FirewallAccessRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
JSON firewallAccessRuleDeleteResponseEnvelopeJSON `json:"-"`
}

// firewallAccessRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallAccessRuleDeleteResponseEnvelope]
type firewallAccessRuleDeleteResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleDeleteResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleDeleteResponseEnvelopeSuccess bool

const (
  FirewallAccessRuleDeleteResponseEnvelopeSuccessTrue FirewallAccessRuleDeleteResponseEnvelopeSuccess = true
)

func (r FirewallAccessRuleDeleteResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case FirewallAccessRuleDeleteResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}

type FirewallAccessRuleEditParams struct {
// The action to apply to a matched request.
Mode param.Field[FirewallAccessRuleEditParamsMode] `json:"mode"`
// An informative summary of the rule, typically used as a reminder or explanation.
Notes param.Field[string] `json:"notes"`
}

func (r FirewallAccessRuleEditParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// The action to apply to a matched request.
type FirewallAccessRuleEditParamsMode string

const (
  FirewallAccessRuleEditParamsModeBlock FirewallAccessRuleEditParamsMode = "block"
  FirewallAccessRuleEditParamsModeChallenge FirewallAccessRuleEditParamsMode = "challenge"
  FirewallAccessRuleEditParamsModeWhitelist FirewallAccessRuleEditParamsMode = "whitelist"
  FirewallAccessRuleEditParamsModeJsChallenge FirewallAccessRuleEditParamsMode = "js_challenge"
  FirewallAccessRuleEditParamsModeManagedChallenge FirewallAccessRuleEditParamsMode = "managed_challenge"
)

func (r FirewallAccessRuleEditParamsMode) IsKnown() (bool) {
  switch r {
  case FirewallAccessRuleEditParamsModeBlock, FirewallAccessRuleEditParamsModeChallenge, FirewallAccessRuleEditParamsModeWhitelist, FirewallAccessRuleEditParamsModeJsChallenge, FirewallAccessRuleEditParamsModeManagedChallenge:
      return true
  }
  return false
}

type FirewallAccessRuleEditResponseEnvelope struct {
Errors []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
Result FirewallRule `json:"result,required"`
// Whether the API call was successful
Success FirewallAccessRuleEditResponseEnvelopeSuccess `json:"success,required"`
JSON firewallAccessRuleEditResponseEnvelopeJSON `json:"-"`
}

// firewallAccessRuleEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallAccessRuleEditResponseEnvelope]
type firewallAccessRuleEditResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleEditResponseEnvelopeSuccess bool

const (
  FirewallAccessRuleEditResponseEnvelopeSuccessTrue FirewallAccessRuleEditResponseEnvelopeSuccess = true
)

func (r FirewallAccessRuleEditResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case FirewallAccessRuleEditResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}
