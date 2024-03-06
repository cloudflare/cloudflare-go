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

// Fetches WAF rules in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageRuleService) List(ctx context.Context, packageID string, params FirewallWAFPackageRuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[FirewallWAFPackageRuleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules", params.ZoneID, packageID)
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

// Fetches WAF rules in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageRuleService) ListAutoPaging(ctx context.Context, packageID string, params FirewallWAFPackageRuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[FirewallWAFPackageRuleListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, packageID, params, opts...))
}

// Updates a WAF rule. You can only update the mode/action of the rule.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageRuleService) Edit(ctx context.Context, packageID string, ruleID string, params FirewallWAFPackageRuleEditParams, opts ...option.RequestOption) (res *FirewallWAFPackageRuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFPackageRuleEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", params.ZoneID, packageID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
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
func (r *FirewallWAFPackageRuleService) Get(ctx context.Context, packageID string, ruleID string, query FirewallWAFPackageRuleGetParams, opts ...option.RequestOption) (res *FirewallWAFPackageRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFPackageRuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", query.ZoneID, packageID, ruleID)
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
// [FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRule],
// [FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRule] or
// [FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRule].
type FirewallWAFPackageRuleListResponse interface {
	implementsFirewallWAFPackageRuleListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallWAFPackageRuleListResponse)(nil)).Elem(), "")
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleGroup `json:"group,required"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                           `json:"priority,required"`
	JSON     firewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleJSON `json:"-"`
}

// firewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleJSON contains the
// JSON metadata for the struct
// [FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRule]
type firewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleJSON struct {
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

func (r *FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleJSON) RawJSON() string {
	return r.raw
}

func (r FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRule) implementsFirewallWAFPackageRuleListResponse() {
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleAllowedMode string

const (
	FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleAllowedModeOn  FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleAllowedMode = "on"
	FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleAllowedModeOff FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                `json:"name"`
	JSON firewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleGroupJSON contains
// the JSON metadata for the struct
// [FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleGroup]
type firewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleGroupJSON) RawJSON() string {
	return r.raw
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleMode string

const (
	FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleModeOn  FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleMode = "on"
	FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleModeOff FirewallWAFPackageRuleListResponseWAFManagedRulesAnomalyRuleMode = "off"
)

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The list of possible actions of the WAF rule when it is triggered.
	AllowedModes []FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// The default action/mode of a rule.
	DefaultMode FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleGroup `json:"group,required"`
	// The action that the current WAF rule will perform when triggered. Applies to
	// traditional (deny) WAF rules.
	Mode FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                   `json:"priority,required"`
	JSON     firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleJSON `json:"-"`
}

// firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRule]
type firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleJSON struct {
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

func (r *FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleJSON) RawJSON() string {
	return r.raw
}

func (r FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRule) implementsFirewallWAFPackageRuleListResponse() {
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode string

const (
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDefault   FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "default"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDisable   FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "disable"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate  FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "simulate"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeBlock     FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "block"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "challenge"
)

// The default action/mode of a rule.
type FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode string

const (
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultModeDisable   FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate  FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultModeBlock     FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "block"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                        `json:"name"`
	JSON firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleGroupJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleGroup]
type firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleGroupJSON) RawJSON() string {
	return r.raw
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode string

const (
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeDefault   FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode = "default"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeDisable   FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode = "disable"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeSimulate  FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode = "simulate"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeBlock     FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode = "block"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeChallenge FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode = "challenge"
)

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                    `json:"priority,required"`
	JSON     firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleJSON `json:"-"`
}

// firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRule]
type firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleJSON struct {
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

func (r *FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleJSON) RawJSON() string {
	return r.raw
}

func (r FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRule) implementsFirewallWAFPackageRuleListResponse() {
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedMode string

const (
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOn  FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "on"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOff FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                         `json:"name"`
	JSON firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleGroupJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleGroup]
type firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleGroupJSON) RawJSON() string {
	return r.raw
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleMode string

const (
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleModeOn  FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleMode = "on"
	FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleModeOff FirewallWAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleMode = "off"
)

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

func (r firewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleJSON) RawJSON() string {
	return r.raw
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

func (r firewallWAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleJSON) RawJSON() string {
	return r.raw
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

func (r firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleJSON) RawJSON() string {
	return r.raw
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

func (r firewallWAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroupJSON) RawJSON() string {
	return r.raw
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

type FirewallWAFPackageRuleListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned rules.
	Direction param.Field[FirewallWAFPackageRuleListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[FirewallWAFPackageRuleListParamsMatch] `query:"match"`
	// The action/mode a rule has been overridden to perform.
	Mode param.Field[FirewallWAFPackageRuleListParamsMode] `query:"mode"`
	// The field used to sort returned rules.
	Order param.Field[FirewallWAFPackageRuleListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of rules per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [FirewallWAFPackageRuleListParams]'s query parameters as
// `url.Values`.
func (r FirewallWAFPackageRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type FirewallWAFPackageRuleListParamsDirection string

const (
	FirewallWAFPackageRuleListParamsDirectionAsc  FirewallWAFPackageRuleListParamsDirection = "asc"
	FirewallWAFPackageRuleListParamsDirectionDesc FirewallWAFPackageRuleListParamsDirection = "desc"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type FirewallWAFPackageRuleListParamsMatch string

const (
	FirewallWAFPackageRuleListParamsMatchAny FirewallWAFPackageRuleListParamsMatch = "any"
	FirewallWAFPackageRuleListParamsMatchAll FirewallWAFPackageRuleListParamsMatch = "all"
)

// The action/mode a rule has been overridden to perform.
type FirewallWAFPackageRuleListParamsMode string

const (
	FirewallWAFPackageRuleListParamsModeDis FirewallWAFPackageRuleListParamsMode = "DIS"
	FirewallWAFPackageRuleListParamsModeChl FirewallWAFPackageRuleListParamsMode = "CHL"
	FirewallWAFPackageRuleListParamsModeBlk FirewallWAFPackageRuleListParamsMode = "BLK"
	FirewallWAFPackageRuleListParamsModeSim FirewallWAFPackageRuleListParamsMode = "SIM"
)

// The field used to sort returned rules.
type FirewallWAFPackageRuleListParamsOrder string

const (
	FirewallWAFPackageRuleListParamsOrderPriority    FirewallWAFPackageRuleListParamsOrder = "priority"
	FirewallWAFPackageRuleListParamsOrderGroupID     FirewallWAFPackageRuleListParamsOrder = "group_id"
	FirewallWAFPackageRuleListParamsOrderDescription FirewallWAFPackageRuleListParamsOrder = "description"
)

type FirewallWAFPackageRuleEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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

func (r firewallWAFPackageRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r firewallWAFPackageRuleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r firewallWAFPackageRuleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallWAFPackageRuleEditResponseEnvelopeSuccess bool

const (
	FirewallWAFPackageRuleEditResponseEnvelopeSuccessTrue FirewallWAFPackageRuleEditResponseEnvelopeSuccess = true
)

type FirewallWAFPackageRuleGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

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

func (r firewallWAFPackageRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r firewallWAFPackageRuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r firewallWAFPackageRuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallWAFPackageRuleGetResponseEnvelopeSuccess bool

const (
	FirewallWAFPackageRuleGetResponseEnvelopeSuccessTrue FirewallWAFPackageRuleGetResponseEnvelopeSuccess = true
)
