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
	"github.com/tidwall/gjson"
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

// Updates a WAF rule. You can only update the mode/action of the rule.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageRuleService) Edit(ctx context.Context, zoneID string, packageID string, ruleID string, body FirewallWAFPackageRuleEditParams, opts ...option.RequestOption) (res *FirewallWAFPackageRuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFPackageRuleEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", zoneID, packageID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a WAF rule in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageRuleService) Get(ctx context.Context, zoneID string, packageID string, ruleID string, opts ...option.RequestOption) (res *FirewallWAFPackageRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFPackageRuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", zoneID, packageID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches WAF rules in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageRuleService) WAFRulesListWAFRules(ctx context.Context, zoneID string, packageID string, query FirewallWAFPackageRuleWAFRulesListWAFRulesParams, opts ...option.RequestOption) (res *[]FirewallWAFPackageRuleWAFRulesListWAFRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules", zoneID, packageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by
// [FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRule],
// [FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule] or
// [FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule].
type FirewallWAFPackageRuleEditResponse interface {
	implementsFirewallWAFPackageRuleEditResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallWAFPackageRuleEditResponse)(nil)).Elem(), "")
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroup `json:"group,required"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                           `json:"priority,required"`
	JSON     firewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleJSON `json:"-"`
}

// firewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleJSON contains the
// JSON metadata for the struct
// [FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRule]
type firewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleJSON struct {
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

func (r *FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRule) implementsFirewallWAFPackageRuleEditResponse() {
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedMode string

const (
	FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedModeOn  FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedMode = "on"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedModeOff FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                `json:"name"`
	JSON firewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroupJSON contains
// the JSON metadata for the struct
// [FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroup]
type firewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleMode string

const (
	FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleModeOn  FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleMode = "on"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleModeOff FirewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleMode = "off"
)

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The list of possible actions of the WAF rule when it is triggered.
	AllowedModes []FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// The default action/mode of a rule.
	DefaultMode FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroup `json:"group,required"`
	// The action that the current WAF rule will perform when triggered. Applies to
	// traditional (deny) WAF rules.
	Mode FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                   `json:"priority,required"`
	JSON     firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleJSON `json:"-"`
}

// firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule]
type firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleJSON struct {
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

func (r *FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule) implementsFirewallWAFPackageRuleEditResponse() {
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode string

const (
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDefault   FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "default"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDisable   FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "disable"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate  FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "simulate"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeBlock     FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "block"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "challenge"
)

// The default action/mode of a rule.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode string

const (
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeDisable   FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate  FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeBlock     FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "block"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                        `json:"name"`
	JSON firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroupJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroup]
type firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode string

const (
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeDefault   FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "default"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeDisable   FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "disable"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeSimulate  FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "simulate"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeBlock     FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "block"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeChallenge FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "challenge"
)

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                    `json:"priority,required"`
	JSON     firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleJSON `json:"-"`
}

// firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule]
type firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleJSON struct {
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

func (r *FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule) implementsFirewallWAFPackageRuleEditResponse() {
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode string

const (
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOn  FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "on"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOff FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                         `json:"name"`
	JSON firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroupJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroup]
type firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode string

const (
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleModeOn  FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode = "on"
	FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleModeOff FirewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode = "off"
)

// Union satisfied by [FirewallWAFPackageRuleGetResponseUnknown],
// [FirewallWAFPackageRuleGetResponseArray] or [shared.UnionString].
type FirewallWAFPackageRuleGetResponse interface {
	ImplementsFirewallWAFPackageRuleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallWAFPackageRuleGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallWAFPackageRuleGetResponseArray []interface{}

func (r FirewallWAFPackageRuleGetResponseArray) ImplementsFirewallWAFPackageRuleGetResponse() {}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRule],
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRule]
// or
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRule].
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponse interface {
	implementsFirewallWAFPackageRuleWAFRulesListWAFRulesResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallWAFPackageRuleWAFRulesListWAFRulesResponse)(nil)).Elem(), "")
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleGroup `json:"group,required"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                           `json:"priority,required"`
	JSON     firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleJSON `json:"-"`
}

// firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRule]
type firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleJSON struct {
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

func (r *FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRule) implementsFirewallWAFPackageRuleWAFRulesListWAFRulesResponse() {
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleAllowedMode string

const (
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleAllowedModeOn  FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleAllowedMode = "on"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleAllowedModeOff FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                                `json:"name"`
	JSON firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleGroupJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleGroup]
type firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleMode string

const (
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleModeOn  FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleMode = "on"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleModeOff FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesAnomalyRuleMode = "off"
)

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The list of possible actions of the WAF rule when it is triggered.
	AllowedModes []FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// The default action/mode of a rule.
	DefaultMode FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleGroup `json:"group,required"`
	// The action that the current WAF rule will perform when triggered. Applies to
	// traditional (deny) WAF rules.
	Mode FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                                   `json:"priority,required"`
	JSON     firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleJSON `json:"-"`
}

// firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRule]
type firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleJSON struct {
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

func (r *FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRule) implementsFirewallWAFPackageRuleWAFRulesListWAFRulesResponse() {
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleAllowedMode string

const (
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDefault   FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "default"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDisable   FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "disable"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate  FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "simulate"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleAllowedModeBlock     FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "block"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "challenge"
)

// The default action/mode of a rule.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleDefaultMode string

const (
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleDefaultModeDisable   FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate  FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleDefaultModeBlock     FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "block"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                                        `json:"name"`
	JSON firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleGroupJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleGroup]
type firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleMode string

const (
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleModeDefault   FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleMode = "default"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleModeDisable   FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleMode = "disable"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleModeSimulate  FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleMode = "simulate"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleModeBlock     FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleMode = "block"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleModeChallenge FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalDenyRuleMode = "challenge"
)

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                                    `json:"priority,required"`
	JSON     firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleJSON `json:"-"`
}

// firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRule]
type firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleJSON struct {
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

func (r *FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRule) implementsFirewallWAFPackageRuleWAFRulesListWAFRulesResponse() {
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleAllowedMode string

const (
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOn  FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "on"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOff FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                                         `json:"name"`
	JSON firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleGroupJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleGroup]
type firewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleMode string

const (
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleModeOn  FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleMode = "on"
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleModeOff FirewallWAFPackageRuleWAFRulesListWAFRulesResponseWAFManagedRulesTraditionalAllowRuleMode = "off"
)

type FirewallWAFPackageRuleEditParams struct {
	// The mode/action of the rule when triggered. You must use a value from the
	// `allowed_modes` array of the current rule.
	Mode param.Field[FirewallWAFPackageRuleEditParamsMode] `json:"mode"`
}

func (r FirewallWAFPackageRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The mode/action of the rule when triggered. You must use a value from the
// `allowed_modes` array of the current rule.
type FirewallWAFPackageRuleEditParamsMode string

const (
	FirewallWAFPackageRuleEditParamsModeDefault   FirewallWAFPackageRuleEditParamsMode = "default"
	FirewallWAFPackageRuleEditParamsModeDisable   FirewallWAFPackageRuleEditParamsMode = "disable"
	FirewallWAFPackageRuleEditParamsModeSimulate  FirewallWAFPackageRuleEditParamsMode = "simulate"
	FirewallWAFPackageRuleEditParamsModeBlock     FirewallWAFPackageRuleEditParamsMode = "block"
	FirewallWAFPackageRuleEditParamsModeChallenge FirewallWAFPackageRuleEditParamsMode = "challenge"
	FirewallWAFPackageRuleEditParamsModeOn        FirewallWAFPackageRuleEditParamsMode = "on"
	FirewallWAFPackageRuleEditParamsModeOff       FirewallWAFPackageRuleEditParamsMode = "off"
)

type FirewallWAFPackageRuleEditResponseEnvelope struct {
	Errors   []FirewallWAFPackageRuleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFPackageRuleEditResponseEnvelopeMessages `json:"messages,required"`
	// When triggered, anomaly detection WAF rules contribute to an overall threat
	// score that will determine if a request is considered malicious. You can
	// configure the total scoring threshold through the 'sensitivity' property of the
	// WAF package.
	Result FirewallWAFPackageRuleEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success FirewallWAFPackageRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFPackageRuleEditResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFPackageRuleEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [FirewallWAFPackageRuleEditResponseEnvelope]
type firewallWAFPackageRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleEditResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    firewallWAFPackageRuleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFPackageRuleEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FirewallWAFPackageRuleEditResponseEnvelopeErrors]
type firewallWAFPackageRuleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleEditResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    firewallWAFPackageRuleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFPackageRuleEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [FirewallWAFPackageRuleEditResponseEnvelopeMessages]
type firewallWAFPackageRuleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFPackageRuleEditResponseEnvelopeSuccess bool

const (
	FirewallWAFPackageRuleEditResponseEnvelopeSuccessTrue FirewallWAFPackageRuleEditResponseEnvelopeSuccess = true
)

type FirewallWAFPackageRuleGetResponseEnvelope struct {
	Errors   []FirewallWAFPackageRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFPackageRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallWAFPackageRuleGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FirewallWAFPackageRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFPackageRuleGetResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFPackageRuleGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallWAFPackageRuleGetResponseEnvelope]
type firewallWAFPackageRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    firewallWAFPackageRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFPackageRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FirewallWAFPackageRuleGetResponseEnvelopeErrors]
type firewallWAFPackageRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    firewallWAFPackageRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFPackageRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallWAFPackageRuleGetResponseEnvelopeMessages]
type firewallWAFPackageRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFPackageRuleGetResponseEnvelopeSuccess bool

const (
	FirewallWAFPackageRuleGetResponseEnvelopeSuccessTrue FirewallWAFPackageRuleGetResponseEnvelopeSuccess = true
)

type FirewallWAFPackageRuleWAFRulesListWAFRulesParams struct {
	// The direction used to sort returned rules.
	Direction param.Field[FirewallWAFPackageRuleWAFRulesListWAFRulesParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMatch] `query:"match"`
	// The action/mode a rule has been overridden to perform.
	Mode param.Field[FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMode] `query:"mode"`
	// The field used to sort returned rules.
	Order param.Field[FirewallWAFPackageRuleWAFRulesListWAFRulesParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of rules per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [FirewallWAFPackageRuleWAFRulesListWAFRulesParams]'s query
// parameters as `url.Values`.
func (r FirewallWAFPackageRuleWAFRulesListWAFRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type FirewallWAFPackageRuleWAFRulesListWAFRulesParamsDirection string

const (
	FirewallWAFPackageRuleWAFRulesListWAFRulesParamsDirectionAsc  FirewallWAFPackageRuleWAFRulesListWAFRulesParamsDirection = "asc"
	FirewallWAFPackageRuleWAFRulesListWAFRulesParamsDirectionDesc FirewallWAFPackageRuleWAFRulesListWAFRulesParamsDirection = "desc"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMatch string

const (
	FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMatchAny FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMatch = "any"
	FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMatchAll FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMatch = "all"
)

// The action/mode a rule has been overridden to perform.
type FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMode string

const (
	FirewallWAFPackageRuleWAFRulesListWAFRulesParamsModeDis FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMode = "DIS"
	FirewallWAFPackageRuleWAFRulesListWAFRulesParamsModeChl FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMode = "CHL"
	FirewallWAFPackageRuleWAFRulesListWAFRulesParamsModeBlk FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMode = "BLK"
	FirewallWAFPackageRuleWAFRulesListWAFRulesParamsModeSim FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMode = "SIM"
)

// The field used to sort returned rules.
type FirewallWAFPackageRuleWAFRulesListWAFRulesParamsOrder string

const (
	FirewallWAFPackageRuleWAFRulesListWAFRulesParamsOrderPriority    FirewallWAFPackageRuleWAFRulesListWAFRulesParamsOrder = "priority"
	FirewallWAFPackageRuleWAFRulesListWAFRulesParamsOrderGroupID     FirewallWAFPackageRuleWAFRulesListWAFRulesParamsOrder = "group_id"
	FirewallWAFPackageRuleWAFRulesListWAFRulesParamsOrderDescription FirewallWAFPackageRuleWAFRulesListWAFRulesParamsOrder = "description"
)

type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelope struct {
	Errors   []FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []FirewallWAFPackageRuleWAFRulesListWAFRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeJSON       `json:"-"`
}

// firewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelope]
type firewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeErrors struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    firewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeErrors]
type firewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeMessages struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    firewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeMessages]
type firewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeSuccess bool

const (
	FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeSuccessTrue FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeSuccess = true
)

type FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                  `json:"total_count"`
	JSON       firewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeResultInfo]
type firewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleWAFRulesListWAFRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
