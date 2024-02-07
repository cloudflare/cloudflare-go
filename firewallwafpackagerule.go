// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// FirewallWAFPackageRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallWAFPackageRuleService]
// method instead.
type FirewallWAFPackageRuleService struct {
	Options []option.RequestOption
}

// NewFirewallWAFPackageRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallWAFPackageRuleService(opts ...option.RequestOption) (r *FirewallWAFPackageRuleService) {
	r = &FirewallWAFPackageRuleService{}
	r.Options = opts
	return
}

// Fetches the details of a WAF rule in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageRuleService) Get(ctx context.Context, zoneID string, packageID string, ruleID string, opts ...option.RequestOption) (res *FirewallWAFPackageRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", zoneID, packageID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a WAF rule. You can only update the mode/action of the rule.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageRuleService) Update(ctx context.Context, zoneID string, packageID string, ruleID string, body FirewallWAFPackageRuleUpdateParams, opts ...option.RequestOption) (res *FirewallWAFPackageRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", zoneID, packageID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type FirewallWAFPackageRuleGetResponse struct {
	Errors   []FirewallWAFPackageRuleGetResponseError   `json:"errors"`
	Messages []FirewallWAFPackageRuleGetResponseMessage `json:"messages"`
	Result   interface{}                                `json:"result"`
	// Whether the API call was successful
	Success FirewallWAFPackageRuleGetResponseSuccess `json:"success"`
	JSON    firewallWAFPackageRuleGetResponseJSON    `json:"-"`
}

// firewallWAFPackageRuleGetResponseJSON contains the JSON metadata for the struct
// [FirewallWAFPackageRuleGetResponse]
type firewallWAFPackageRuleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleGetResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    firewallWAFPackageRuleGetResponseErrorJSON `json:"-"`
}

// firewallWAFPackageRuleGetResponseErrorJSON contains the JSON metadata for the
// struct [FirewallWAFPackageRuleGetResponseError]
type firewallWAFPackageRuleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleGetResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    firewallWAFPackageRuleGetResponseMessageJSON `json:"-"`
}

// firewallWAFPackageRuleGetResponseMessageJSON contains the JSON metadata for the
// struct [FirewallWAFPackageRuleGetResponseMessage]
type firewallWAFPackageRuleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFPackageRuleGetResponseSuccess bool

const (
	FirewallWAFPackageRuleGetResponseSuccessTrue FirewallWAFPackageRuleGetResponseSuccess = true
)

type FirewallWAFPackageRuleUpdateResponse struct {
	Errors   []FirewallWAFPackageRuleUpdateResponseError   `json:"errors"`
	Messages []FirewallWAFPackageRuleUpdateResponseMessage `json:"messages"`
	// When triggered, anomaly detection WAF rules contribute to an overall threat
	// score that will determine if a request is considered malicious. You can
	// configure the total scoring threshold through the 'sensitivity' property of the
	// WAF package.
	Result FirewallWAFPackageRuleUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success FirewallWAFPackageRuleUpdateResponseSuccess `json:"success"`
	JSON    firewallWAFPackageRuleUpdateResponseJSON    `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseJSON contains the JSON metadata for the
// struct [FirewallWAFPackageRuleUpdateResponse]
type firewallWAFPackageRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleUpdateResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    firewallWAFPackageRuleUpdateResponseErrorJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseErrorJSON contains the JSON metadata for the
// struct [FirewallWAFPackageRuleUpdateResponseError]
type firewallWAFPackageRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleUpdateResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    firewallWAFPackageRuleUpdateResponseMessageJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseMessageJSON contains the JSON metadata for
// the struct [FirewallWAFPackageRuleUpdateResponseMessage]
type firewallWAFPackageRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by
// [FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRule],
// [FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRule]
// or
// [FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRule].
type FirewallWAFPackageRuleUpdateResponseResult interface {
	implementsFirewallWAFPackageRuleUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallWAFPackageRuleUpdateResponseResult)(nil)).Elem(), "")
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleGroup `json:"group,required"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                   `json:"priority,required"`
	JSON     firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRule]
type firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleJSON struct {
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

func (r *FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRule) implementsFirewallWAFPackageRuleUpdateResponseResult() {
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleAllowedMode string

const (
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleAllowedModeOn  FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleAllowedMode = "on"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleAllowedModeOff FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                        `json:"name"`
	JSON firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleGroupJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleGroup]
type firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleMode string

const (
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleModeOn  FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleMode = "on"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleModeOff FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesAnomalyRuleMode = "off"
)

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The list of possible actions of the WAF rule when it is triggered.
	AllowedModes []FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// The default action/mode of a rule.
	DefaultMode FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleGroup `json:"group,required"`
	// The action that the current WAF rule will perform when triggered. Applies to
	// traditional (deny) WAF rules.
	Mode FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                           `json:"priority,required"`
	JSON     firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRule]
type firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleJSON struct {
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

func (r *FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRule) implementsFirewallWAFPackageRuleUpdateResponseResult() {
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleAllowedMode string

const (
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleAllowedModeDefault   FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleAllowedMode = "default"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleAllowedModeDisable   FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleAllowedMode = "disable"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate  FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleAllowedMode = "simulate"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleAllowedModeBlock     FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleAllowedMode = "block"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleAllowedMode = "challenge"
)

// The default action/mode of a rule.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleDefaultMode string

const (
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleDefaultModeDisable   FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate  FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleDefaultModeBlock     FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleDefaultMode = "block"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                                `json:"name"`
	JSON firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleGroupJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleGroup]
type firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleMode string

const (
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleModeDefault   FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleMode = "default"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleModeDisable   FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleMode = "disable"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleModeSimulate  FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleMode = "simulate"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleModeBlock     FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleMode = "block"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleModeChallenge FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalDenyRuleMode = "challenge"
)

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                            `json:"priority,required"`
	JSON     firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRule]
type firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleJSON struct {
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

func (r *FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRule) implementsFirewallWAFPackageRuleUpdateResponseResult() {
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleAllowedMode string

const (
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleAllowedModeOn  FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleAllowedMode = "on"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleAllowedModeOff FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                                 `json:"name"`
	JSON firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleGroupJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleGroup]
type firewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleMode string

const (
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleModeOn  FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleMode = "on"
	FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleModeOff FirewallWAFPackageRuleUpdateResponseResultWAFManagedRulesTraditionalAllowRuleMode = "off"
)

// Whether the API call was successful
type FirewallWAFPackageRuleUpdateResponseSuccess bool

const (
	FirewallWAFPackageRuleUpdateResponseSuccessTrue FirewallWAFPackageRuleUpdateResponseSuccess = true
)

type FirewallWAFPackageRuleUpdateParams struct {
	// The mode/action of the rule when triggered. You must use a value from the
	// `allowed_modes` array of the current rule.
	Mode param.Field[FirewallWAFPackageRuleUpdateParamsMode] `json:"mode"`
}

func (r FirewallWAFPackageRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The mode/action of the rule when triggered. You must use a value from the
// `allowed_modes` array of the current rule.
type FirewallWAFPackageRuleUpdateParamsMode string

const (
	FirewallWAFPackageRuleUpdateParamsModeDefault   FirewallWAFPackageRuleUpdateParamsMode = "default"
	FirewallWAFPackageRuleUpdateParamsModeDisable   FirewallWAFPackageRuleUpdateParamsMode = "disable"
	FirewallWAFPackageRuleUpdateParamsModeSimulate  FirewallWAFPackageRuleUpdateParamsMode = "simulate"
	FirewallWAFPackageRuleUpdateParamsModeBlock     FirewallWAFPackageRuleUpdateParamsMode = "block"
	FirewallWAFPackageRuleUpdateParamsModeChallenge FirewallWAFPackageRuleUpdateParamsMode = "challenge"
	FirewallWAFPackageRuleUpdateParamsModeOn        FirewallWAFPackageRuleUpdateParamsMode = "on"
	FirewallWAFPackageRuleUpdateParamsModeOff       FirewallWAFPackageRuleUpdateParamsMode = "off"
)
