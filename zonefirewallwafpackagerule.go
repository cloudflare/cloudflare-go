// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneFirewallWafPackageRuleService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneFirewallWafPackageRuleService] method instead.
type ZoneFirewallWafPackageRuleService struct {
	Options []option.RequestOption
}

// NewZoneFirewallWafPackageRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneFirewallWafPackageRuleService(opts ...option.RequestOption) (r *ZoneFirewallWafPackageRuleService) {
	r = &ZoneFirewallWafPackageRuleService{}
	r.Options = opts
	return
}

// Fetches the details of a WAF rule in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageRuleService) Get(ctx context.Context, zoneID string, packageID string, identifier string, opts ...option.RequestOption) (res *ZoneFirewallWafPackageRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", zoneID, packageID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a WAF rule. You can only update the mode/action of the rule.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageRuleService) Update(ctx context.Context, zoneID string, packageID string, identifier string, body ZoneFirewallWafPackageRuleUpdateParams, opts ...option.RequestOption) (res *ZoneFirewallWafPackageRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", zoneID, packageID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches WAF rules in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageRuleService) WafRulesListWafRules(ctx context.Context, zoneID string, packageID string, query ZoneFirewallWafPackageRuleWafRulesListWafRulesParams, opts ...option.RequestOption) (res *shared.Page[ZoneFirewallWafPackageRuleWafRulesListWafRulesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules", zoneID, packageID)
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

type ZoneFirewallWafPackageRuleGetResponse struct {
	Errors   []ZoneFirewallWafPackageRuleGetResponseError   `json:"errors"`
	Messages []ZoneFirewallWafPackageRuleGetResponseMessage `json:"messages"`
	Result   interface{}                                    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallWafPackageRuleGetResponseSuccess `json:"success"`
	JSON    zoneFirewallWafPackageRuleGetResponseJSON    `json:"-"`
}

// zoneFirewallWafPackageRuleGetResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafPackageRuleGetResponse]
type zoneFirewallWafPackageRuleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafPackageRuleGetResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneFirewallWafPackageRuleGetResponseErrorJSON `json:"-"`
}

// zoneFirewallWafPackageRuleGetResponseErrorJSON contains the JSON metadata for
// the struct [ZoneFirewallWafPackageRuleGetResponseError]
type zoneFirewallWafPackageRuleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafPackageRuleGetResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneFirewallWafPackageRuleGetResponseMessageJSON `json:"-"`
}

// zoneFirewallWafPackageRuleGetResponseMessageJSON contains the JSON metadata for
// the struct [ZoneFirewallWafPackageRuleGetResponseMessage]
type zoneFirewallWafPackageRuleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFirewallWafPackageRuleGetResponseSuccess bool

const (
	ZoneFirewallWafPackageRuleGetResponseSuccessTrue ZoneFirewallWafPackageRuleGetResponseSuccess = true
)

type ZoneFirewallWafPackageRuleUpdateResponse struct {
	Errors   []ZoneFirewallWafPackageRuleUpdateResponseError   `json:"errors"`
	Messages []ZoneFirewallWafPackageRuleUpdateResponseMessage `json:"messages"`
	// When triggered, anomaly detection WAF rules contribute to an overall threat
	// score that will determine if a request is considered malicious. You can
	// configure the total scoring threshold through the 'sensitivity' property of the
	// WAF package.
	Result ZoneFirewallWafPackageRuleUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallWafPackageRuleUpdateResponseSuccess `json:"success"`
	JSON    zoneFirewallWafPackageRuleUpdateResponseJSON    `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafPackageRuleUpdateResponse]
type zoneFirewallWafPackageRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafPackageRuleUpdateResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneFirewallWafPackageRuleUpdateResponseErrorJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneFirewallWafPackageRuleUpdateResponseError]
type zoneFirewallWafPackageRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafPackageRuleUpdateResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneFirewallWafPackageRuleUpdateResponseMessageJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneFirewallWafPackageRuleUpdateResponseMessage]
type zoneFirewallWafPackageRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by
// [ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRule],
// [ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRule] or
// [ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRule].
type ZoneFirewallWafPackageRuleUpdateResponseResult interface {
	implementsZoneFirewallWafPackageRuleUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallWafPackageRuleUpdateResponseResult)(nil)).Elem(), "")
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleGroup `json:"group,required"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                `json:"priority,required"`
	JSON     zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsAnomalyRuleJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsAnomalyRuleJSON contains
// the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRule]
type zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsAnomalyRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRule) implementsZoneFirewallWafPackageRuleUpdateResponseResult() {
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleAllowedMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleAllowedModeOn  ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleAllowedMode = "on"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleAllowedModeOff ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                     `json:"name"`
	JSON zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsAnomalyRuleGroupJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsAnomalyRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleGroup]
type zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsAnomalyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleModeOn  ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleMode = "on"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleModeOff ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsAnomalyRuleMode = "off"
)

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The list of possible actions of the WAF rule when it is triggered.
	AllowedModes []ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// The default action/mode of a rule.
	DefaultMode ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleGroup `json:"group,required"`
	// The action that the current WAF rule will perform when triggered. Applies to
	// traditional (deny) WAF rules.
	Mode ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                        `json:"priority,required"`
	JSON     zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsTraditionalDenyRuleJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsTraditionalDenyRuleJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRule]
type zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsTraditionalDenyRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	DefaultMode  apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRule) implementsZoneFirewallWafPackageRuleUpdateResponseResult() {
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleAllowedMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleAllowedModeDefault   ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleAllowedMode = "default"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleAllowedModeDisable   ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleAllowedMode = "disable"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleAllowedModeSimulate  ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleAllowedMode = "simulate"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleAllowedModeBlock     ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleAllowedMode = "block"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleAllowedModeChallenge ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleAllowedMode = "challenge"
)

// The default action/mode of a rule.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleDefaultMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleDefaultModeDisable   ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleDefaultMode = "disable"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleDefaultModeSimulate  ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleDefaultMode = "simulate"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleDefaultModeBlock     ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleDefaultMode = "block"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleDefaultModeChallenge ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleDefaultMode = "challenge"
)

// The rule group to which the current WAF rule belongs.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                             `json:"name"`
	JSON zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsTraditionalDenyRuleGroupJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsTraditionalDenyRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleGroup]
type zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsTraditionalDenyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleModeDefault   ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleMode = "default"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleModeDisable   ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleMode = "disable"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleModeSimulate  ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleMode = "simulate"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleModeBlock     ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleMode = "block"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleModeChallenge ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalDenyRuleMode = "challenge"
)

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                         `json:"priority,required"`
	JSON     zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsTraditionalAllowRuleJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsTraditionalAllowRuleJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRule]
type zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsTraditionalAllowRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRule) implementsZoneFirewallWafPackageRuleUpdateResponseResult() {
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleAllowedMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleAllowedModeOn  ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleAllowedMode = "on"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleAllowedModeOff ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                              `json:"name"`
	JSON zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsTraditionalAllowRuleGroupJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsTraditionalAllowRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleGroup]
type zoneFirewallWafPackageRuleUpdateResponseResultEg9LSmLsTraditionalAllowRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleModeOn  ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleMode = "on"
	ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleModeOff ZoneFirewallWafPackageRuleUpdateResponseResultEG9LSmLsTraditionalAllowRuleMode = "off"
)

// Whether the API call was successful
type ZoneFirewallWafPackageRuleUpdateResponseSuccess bool

const (
	ZoneFirewallWafPackageRuleUpdateResponseSuccessTrue ZoneFirewallWafPackageRuleUpdateResponseSuccess = true
)

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by
// [ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRule],
// [ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRule]
// or
// [ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRule].
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponse interface {
	implementsZoneFirewallWafPackageRuleWafRulesListWafRulesResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallWafPackageRuleWafRulesListWafRulesResponse)(nil)).Elem(), "")
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleGroup `json:"group,required"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                        `json:"priority,required"`
	JSON     zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsAnomalyRuleJSON `json:"-"`
}

// zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsAnomalyRuleJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRule]
type zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsAnomalyRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRule) implementsZoneFirewallWafPackageRuleWafRulesListWafRulesResponse() {
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleAllowedMode string

const (
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleAllowedModeOn  ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleAllowedMode = "on"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleAllowedModeOff ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                             `json:"name"`
	JSON zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsAnomalyRuleGroupJSON `json:"-"`
}

// zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsAnomalyRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleGroup]
type zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsAnomalyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleMode string

const (
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleModeOn  ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleMode = "on"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleModeOff ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsAnomalyRuleMode = "off"
)

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The list of possible actions of the WAF rule when it is triggered.
	AllowedModes []ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// The default action/mode of a rule.
	DefaultMode ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleGroup `json:"group,required"`
	// The action that the current WAF rule will perform when triggered. Applies to
	// traditional (deny) WAF rules.
	Mode ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                                `json:"priority,required"`
	JSON     zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsTraditionalDenyRuleJSON `json:"-"`
}

// zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsTraditionalDenyRuleJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRule]
type zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsTraditionalDenyRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	DefaultMode  apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRule) implementsZoneFirewallWafPackageRuleWafRulesListWafRulesResponse() {
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleAllowedMode string

const (
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleAllowedModeDefault   ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleAllowedMode = "default"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleAllowedModeDisable   ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleAllowedMode = "disable"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleAllowedModeSimulate  ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleAllowedMode = "simulate"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleAllowedModeBlock     ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleAllowedMode = "block"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleAllowedModeChallenge ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleAllowedMode = "challenge"
)

// The default action/mode of a rule.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleDefaultMode string

const (
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleDefaultModeDisable   ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleDefaultMode = "disable"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleDefaultModeSimulate  ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleDefaultMode = "simulate"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleDefaultModeBlock     ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleDefaultMode = "block"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleDefaultModeChallenge ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleDefaultMode = "challenge"
)

// The rule group to which the current WAF rule belongs.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                                     `json:"name"`
	JSON zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsTraditionalDenyRuleGroupJSON `json:"-"`
}

// zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsTraditionalDenyRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleGroup]
type zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsTraditionalDenyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleMode string

const (
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleModeDefault   ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleMode = "default"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleModeDisable   ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleMode = "disable"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleModeSimulate  ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleMode = "simulate"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleModeBlock     ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleMode = "block"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleModeChallenge ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalDenyRuleMode = "challenge"
)

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                                 `json:"priority,required"`
	JSON     zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsTraditionalAllowRuleJSON `json:"-"`
}

// zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsTraditionalAllowRuleJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRule]
type zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsTraditionalAllowRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRule) implementsZoneFirewallWafPackageRuleWafRulesListWafRulesResponse() {
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleAllowedMode string

const (
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleAllowedModeOn  ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleAllowedMode = "on"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleAllowedModeOff ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                                      `json:"name"`
	JSON zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsTraditionalAllowRuleGroupJSON `json:"-"`
}

// zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsTraditionalAllowRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleGroup]
type zoneFirewallWafPackageRuleWafRulesListWafRulesResponseEg9LSmLsTraditionalAllowRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleMode string

const (
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleModeOn  ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleMode = "on"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleModeOff ZoneFirewallWafPackageRuleWafRulesListWafRulesResponseEG9LSmLsTraditionalAllowRuleMode = "off"
)

type ZoneFirewallWafPackageRuleUpdateParams struct {
	// The mode/action of the rule when triggered. You must use a value from the
	// `allowed_modes` array of the current rule.
	Mode param.Field[ZoneFirewallWafPackageRuleUpdateParamsMode] `json:"mode"`
}

func (r ZoneFirewallWafPackageRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The mode/action of the rule when triggered. You must use a value from the
// `allowed_modes` array of the current rule.
type ZoneFirewallWafPackageRuleUpdateParamsMode string

const (
	ZoneFirewallWafPackageRuleUpdateParamsModeDefault   ZoneFirewallWafPackageRuleUpdateParamsMode = "default"
	ZoneFirewallWafPackageRuleUpdateParamsModeDisable   ZoneFirewallWafPackageRuleUpdateParamsMode = "disable"
	ZoneFirewallWafPackageRuleUpdateParamsModeSimulate  ZoneFirewallWafPackageRuleUpdateParamsMode = "simulate"
	ZoneFirewallWafPackageRuleUpdateParamsModeBlock     ZoneFirewallWafPackageRuleUpdateParamsMode = "block"
	ZoneFirewallWafPackageRuleUpdateParamsModeChallenge ZoneFirewallWafPackageRuleUpdateParamsMode = "challenge"
	ZoneFirewallWafPackageRuleUpdateParamsModeOn        ZoneFirewallWafPackageRuleUpdateParamsMode = "on"
	ZoneFirewallWafPackageRuleUpdateParamsModeOff       ZoneFirewallWafPackageRuleUpdateParamsMode = "off"
)

type ZoneFirewallWafPackageRuleWafRulesListWafRulesParams struct {
	// The direction used to sort returned rules.
	Direction param.Field[ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMatch] `query:"match"`
	// The action/mode a rule has been overridden to perform.
	Mode param.Field[ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMode] `query:"mode"`
	// The field used to sort returned rules.
	Order param.Field[ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of rules per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneFirewallWafPackageRuleWafRulesListWafRulesParams]'s
// query parameters as `url.Values`.
func (r ZoneFirewallWafPackageRuleWafRulesListWafRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsDirection string

const (
	ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsDirectionAsc  ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsDirection = "asc"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsDirectionDesc ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsDirection = "desc"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMatch string

const (
	ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMatchAny ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMatch = "any"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMatchAll ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMatch = "all"
)

// The action/mode a rule has been overridden to perform.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMode string

const (
	ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsModeDis ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMode = "DIS"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsModeChl ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMode = "CHL"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsModeBlk ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMode = "BLK"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsModeSim ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsMode = "SIM"
)

// The field used to sort returned rules.
type ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsOrder string

const (
	ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsOrderPriority    ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsOrder = "priority"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsOrderGroupID     ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsOrder = "group_id"
	ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsOrderDescription ZoneFirewallWafPackageRuleWafRulesListWafRulesParamsOrder = "description"
)
