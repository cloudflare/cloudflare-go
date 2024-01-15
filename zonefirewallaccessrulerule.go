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
func (r *ZoneFirewallAccessRuleRuleService) Update(ctx context.Context, zoneID string, identifier string, body ZoneFirewallAccessRuleRuleUpdateParams, opts ...option.RequestOption) (res *ZoneFirewallAccessRuleRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/access_rules/rules/%s", zoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an IP Access rule defined at the zone level.
//
// Optionally, you can use the `cascade` property to specify that you wish to
// delete similar rules in other zones managed by the same zone owner.
func (r *ZoneFirewallAccessRuleRuleService) Delete(ctx context.Context, zoneID string, identifier string, body ZoneFirewallAccessRuleRuleDeleteParams, opts ...option.RequestOption) (res *ZoneFirewallAccessRuleRuleDeleteResponse, err error) {
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
func (r *ZoneFirewallAccessRuleRuleService) IPAccessRulesForAZoneNewAnIPAccessRule(ctx context.Context, zoneID string, body ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParams, opts ...option.RequestOption) (res *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/access_rules/rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches IP Access rules of a zone. You can filter the results using several
// optional parameters.
func (r *ZoneFirewallAccessRuleRuleService) IPAccessRulesForAZoneListIPAccessRules(ctx context.Context, zoneID string, query ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesParams, opts ...option.RequestOption) (res *shared.Page[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/access_rules/rules", zoneID)
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

type ZoneFirewallAccessRuleRuleUpdateResponse struct {
	Errors   []ZoneFirewallAccessRuleRuleUpdateResponseError   `json:"errors"`
	Messages []ZoneFirewallAccessRuleRuleUpdateResponseMessage `json:"messages"`
	Result   ZoneFirewallAccessRuleRuleUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallAccessRuleRuleUpdateResponseSuccess `json:"success"`
	JSON    zoneFirewallAccessRuleRuleUpdateResponseJSON    `json:"-"`
}

// zoneFirewallAccessRuleRuleUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallAccessRuleRuleUpdateResponse]
type zoneFirewallAccessRuleRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallAccessRuleRuleUpdateResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneFirewallAccessRuleRuleUpdateResponseErrorJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneFirewallAccessRuleRuleUpdateResponseError]
type zoneFirewallAccessRuleRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallAccessRuleRuleUpdateResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneFirewallAccessRuleRuleUpdateResponseMessageJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneFirewallAccessRuleRuleUpdateResponseMessage]
type zoneFirewallAccessRuleRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallAccessRuleRuleUpdateResponseResult struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []ZoneFirewallAccessRuleRuleUpdateResponseResultAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration ZoneFirewallAccessRuleRuleUpdateResponseResultConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode ZoneFirewallAccessRuleRuleUpdateResponseResultMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                             `json:"notes"`
	JSON  zoneFirewallAccessRuleRuleUpdateResponseResultJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneFirewallAccessRuleRuleUpdateResponseResult]
type zoneFirewallAccessRuleRuleUpdateResponseResultJSON struct {
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

func (r *ZoneFirewallAccessRuleRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type ZoneFirewallAccessRuleRuleUpdateResponseResultAllowedMode string

const (
	ZoneFirewallAccessRuleRuleUpdateResponseResultAllowedModeBlock            ZoneFirewallAccessRuleRuleUpdateResponseResultAllowedMode = "block"
	ZoneFirewallAccessRuleRuleUpdateResponseResultAllowedModeChallenge        ZoneFirewallAccessRuleRuleUpdateResponseResultAllowedMode = "challenge"
	ZoneFirewallAccessRuleRuleUpdateResponseResultAllowedModeWhitelist        ZoneFirewallAccessRuleRuleUpdateResponseResultAllowedMode = "whitelist"
	ZoneFirewallAccessRuleRuleUpdateResponseResultAllowedModeJsChallenge      ZoneFirewallAccessRuleRuleUpdateResponseResultAllowedMode = "js_challenge"
	ZoneFirewallAccessRuleRuleUpdateResponseResultAllowedModeManagedChallenge ZoneFirewallAccessRuleRuleUpdateResponseResultAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfiguration],
// [ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6Configuration],
// [ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfiguration],
// [ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfiguration]
// or
// [ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfiguration].
type ZoneFirewallAccessRuleRuleUpdateResponseResultConfiguration interface {
	implementsZoneFirewallAccessRuleRuleUpdateResponseResultConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallAccessRuleRuleUpdateResponseResultConfiguration)(nil)).Elem(), "")
}

type ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                 `json:"value"`
	JSON  zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTipConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTipConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfiguration]
type zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTipConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfiguration) implementsZoneFirewallAccessRuleRuleUpdateResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfigurationTargetIP ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIPConfigurationTarget = "ip"
)

type ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                                   `json:"value"`
	JSON  zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6Configuration]
type zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6Configuration) implementsZoneFirewallAccessRuleRuleUpdateResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationTargetIp6 ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTIpv6ConfigurationTarget = "ip6"
)

type ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                                   `json:"value"`
	JSON  zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfiguration]
type zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfiguration) implementsZoneFirewallAccessRuleRuleUpdateResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationTargetIPRange ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCidrConfigurationTarget = "ip_range"
)

type ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                                  `json:"value"`
	JSON  zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTasnConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTasnConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfiguration]
type zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTasnConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfiguration) implementsZoneFirewallAccessRuleRuleUpdateResponseResultConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfigurationTargetASN ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTASNConfigurationTarget = "asn"
)

type ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                                      `json:"value"`
	JSON  zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfiguration]
type zoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfiguration) implementsZoneFirewallAccessRuleRuleUpdateResponseResultConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationTargetCountry ZoneFirewallAccessRuleRuleUpdateResponseResultConfigurationDZw70ubTCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type ZoneFirewallAccessRuleRuleUpdateResponseResultMode string

const (
	ZoneFirewallAccessRuleRuleUpdateResponseResultModeBlock            ZoneFirewallAccessRuleRuleUpdateResponseResultMode = "block"
	ZoneFirewallAccessRuleRuleUpdateResponseResultModeChallenge        ZoneFirewallAccessRuleRuleUpdateResponseResultMode = "challenge"
	ZoneFirewallAccessRuleRuleUpdateResponseResultModeWhitelist        ZoneFirewallAccessRuleRuleUpdateResponseResultMode = "whitelist"
	ZoneFirewallAccessRuleRuleUpdateResponseResultModeJsChallenge      ZoneFirewallAccessRuleRuleUpdateResponseResultMode = "js_challenge"
	ZoneFirewallAccessRuleRuleUpdateResponseResultModeManagedChallenge ZoneFirewallAccessRuleRuleUpdateResponseResultMode = "managed_challenge"
)

// Whether the API call was successful
type ZoneFirewallAccessRuleRuleUpdateResponseSuccess bool

const (
	ZoneFirewallAccessRuleRuleUpdateResponseSuccessTrue ZoneFirewallAccessRuleRuleUpdateResponseSuccess = true
)

type ZoneFirewallAccessRuleRuleDeleteResponse struct {
	Errors   []ZoneFirewallAccessRuleRuleDeleteResponseError   `json:"errors"`
	Messages []ZoneFirewallAccessRuleRuleDeleteResponseMessage `json:"messages"`
	Result   ZoneFirewallAccessRuleRuleDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallAccessRuleRuleDeleteResponseSuccess `json:"success"`
	JSON    zoneFirewallAccessRuleRuleDeleteResponseJSON    `json:"-"`
}

// zoneFirewallAccessRuleRuleDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallAccessRuleRuleDeleteResponse]
type zoneFirewallAccessRuleRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallAccessRuleRuleDeleteResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneFirewallAccessRuleRuleDeleteResponseErrorJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleDeleteResponseErrorJSON contains the JSON metadata for
// the struct [ZoneFirewallAccessRuleRuleDeleteResponseError]
type zoneFirewallAccessRuleRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallAccessRuleRuleDeleteResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneFirewallAccessRuleRuleDeleteResponseMessageJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleDeleteResponseMessageJSON contains the JSON metadata
// for the struct [ZoneFirewallAccessRuleRuleDeleteResponseMessage]
type zoneFirewallAccessRuleRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallAccessRuleRuleDeleteResponseResult struct {
	// The unique identifier of the IP Access rule.
	ID   string                                             `json:"id"`
	JSON zoneFirewallAccessRuleRuleDeleteResponseResultJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleDeleteResponseResultJSON contains the JSON metadata
// for the struct [ZoneFirewallAccessRuleRuleDeleteResponseResult]
type zoneFirewallAccessRuleRuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFirewallAccessRuleRuleDeleteResponseSuccess bool

const (
	ZoneFirewallAccessRuleRuleDeleteResponseSuccessTrue ZoneFirewallAccessRuleRuleDeleteResponseSuccess = true
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponse struct {
	Errors   []ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseError   `json:"errors"`
	Messages []ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseMessage `json:"messages"`
	Result   ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseSuccess `json:"success"`
	JSON    zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseJSON    `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponse]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseError struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseErrorJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseError]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseMessage struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseMessageJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseMessage]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResult struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                                                             `json:"notes"`
	JSON  zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResult]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultJSON struct {
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

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultAllowedMode string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultAllowedModeBlock            ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultAllowedMode = "block"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultAllowedModeChallenge        ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultAllowedMode = "challenge"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultAllowedModeWhitelist        ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultAllowedMode = "whitelist"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultAllowedModeJsChallenge      ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultAllowedMode = "js_challenge"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultAllowedModeManagedChallenge ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfiguration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6Configuration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfiguration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfiguration]
// or
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfiguration].
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfiguration interface {
	implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfiguration)(nil)).Elem(), "")
}

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                                                 `json:"value"`
	JSON  zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTipConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTipConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfiguration]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTipConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfigurationTargetIP ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIPConfigurationTarget = "ip"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                                                                   `json:"value"`
	JSON  zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6Configuration]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6Configuration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationTargetIp6 ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTIpv6ConfigurationTarget = "ip6"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                                                                   `json:"value"`
	JSON  zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfiguration]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationTargetIPRange ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCidrConfigurationTarget = "ip_range"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                                                                  `json:"value"`
	JSON  zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTasnConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTasnConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfiguration]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTasnConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfigurationTargetASN ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTASNConfigurationTarget = "asn"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                                                                      `json:"value"`
	JSON  zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfiguration]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationTargetCountry ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultConfigurationDZw70ubTCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultMode string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultModeBlock            ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultMode = "block"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultModeChallenge        ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultMode = "challenge"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultModeWhitelist        ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultMode = "whitelist"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultModeJsChallenge      ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultMode = "js_challenge"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultModeManagedChallenge ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseResultMode = "managed_challenge"
)

// Whether the API call was successful
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseSuccess bool

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseSuccessTrue ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleResponseSuccess = true
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string                                                                       `json:"notes"`
	JSON  zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponse]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseJSON struct {
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

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseAllowedMode string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseAllowedModeBlock            ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseAllowedMode = "block"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseAllowedModeChallenge        ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseAllowedMode = "challenge"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseAllowedModeWhitelist        ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseAllowedMode = "whitelist"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseAllowedModeJsChallenge      ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseAllowedMode = "js_challenge"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseAllowedModeManagedChallenge ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseAllowedMode = "managed_challenge"
)

// The rule configuration.
//
// Union satisfied by
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIPConfiguration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIpv6Configuration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCidrConfiguration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTASNConfiguration]
// or
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCountryConfiguration].
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfiguration interface {
	implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfiguration)(nil)).Elem(), "")
}

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                                           `json:"value"`
	JSON  zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTipConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTipConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIPConfiguration]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTipConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIPConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIPConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIPConfigurationTargetIP ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIPConfigurationTarget = "ip"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                                                                             `json:"value"`
	JSON  zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIpv6Configuration]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIpv6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIpv6Configuration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationTargetIp6 ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTIpv6ConfigurationTarget = "ip6"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                                                                             `json:"value"`
	JSON  zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCidrConfiguration]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCidrConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationTargetIPRange ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCidrConfigurationTarget = "ip_range"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                                                                            `json:"value"`
	JSON  zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTasnConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTasnConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTASNConfiguration]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTasnConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTASNConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTASNConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTASNConfigurationTargetASN ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTASNConfigurationTarget = "asn"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                                                                                `json:"value"`
	JSON  zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationJSON `json:"-"`
}

// zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCountryConfiguration]
type zoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCountryConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationTargetCountry ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseConfigurationDZw70ubTCountryConfigurationTarget = "country"
)

// The action to apply to a matched request.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseMode string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseModeBlock            ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseMode = "block"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseModeChallenge        ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseMode = "challenge"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseModeWhitelist        ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseMode = "whitelist"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseModeJsChallenge      ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseMode = "js_challenge"
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseModeManagedChallenge ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneListIPAccessRulesResponseMode = "managed_challenge"
)

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
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6Configuration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfiguration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfiguration],
// [ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfiguration].
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration interface {
	implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration()
}

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTargetIP ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIPConfigurationTarget = "ip"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6Configuration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6ConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6ConfigurationTargetIp6 ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTIpv6ConfigurationTarget = "ip6"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfigurationTargetIPRange ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCidrConfigurationTarget = "ip_range"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfigurationTargetASN ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTASNConfigurationTarget = "asn"
)

type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfiguration) implementsZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfigurationTargetCountry ZoneFirewallAccessRuleRuleIPAccessRulesForAZoneNewAnIPAccessRuleParamsConfigurationDZw70ubTCountryConfigurationTarget = "country"
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
