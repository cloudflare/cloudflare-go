// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// WAFPackageRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWAFPackageRuleService] method
// instead.
type WAFPackageRuleService struct {
	Options []option.RequestOption
}

// NewWAFPackageRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWAFPackageRuleService(opts ...option.RequestOption) (r *WAFPackageRuleService) {
	r = &WAFPackageRuleService{}
	r.Options = opts
	return
}

// Fetches WAF rules in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFPackageRuleService) List(ctx context.Context, packageID string, params WAFPackageRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Rule], err error) {
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
func (r *WAFPackageRuleService) ListAutoPaging(ctx context.Context, packageID string, params WAFPackageRuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Rule] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, packageID, params, opts...))
}

// Updates a WAF rule. You can only update the mode/action of the rule.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFPackageRuleService) Edit(ctx context.Context, packageID string, ruleID string, params WAFPackageRuleEditParams, opts ...option.RequestOption) (res *WAFPackageRuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WAFPackageRuleEditResponseEnvelope
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
func (r *WAFPackageRuleService) Get(ctx context.Context, packageID string, ruleID string, query WAFPackageRuleGetParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env WAFPackageRuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", query.ZoneID, packageID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type AllowedModesAnomaly string

const (
	AllowedModesAnomalyOn  AllowedModesAnomaly = "on"
	AllowedModesAnomalyOff AllowedModesAnomaly = "off"
)

func (r AllowedModesAnomaly) IsKnown() bool {
	switch r {
	case AllowedModesAnomalyOn, AllowedModesAnomalyOff:
		return true
	}
	return false
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type Rule struct {
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group UnnamedSchemaRef532d8b97684c9032dd36bae8acddebf5 `json:"group,required"`
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority     string      `json:"priority,required"`
	AllowedModes interface{} `json:"allowed_modes"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode        RuleMode    `json:"mode,required"`
	DefaultMode interface{} `json:"default_mode,required"`
	JSON        ruleJSON    `json:"-"`
	union       RuleUnion
}

// ruleJSON contains the JSON metadata for the struct [Rule]
type ruleJSON struct {
	Description  apijson.Field
	Group        apijson.Field
	ID           apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	AllowedModes apijson.Field
	Mode         apijson.Field
	DefaultMode  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r ruleJSON) RawJSON() string {
	return r.raw
}

func (r *Rule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r Rule) AsUnion() RuleUnion {
	return r.union
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by [firewall.RuleWAFManagedRulesAnomalyRule],
// [firewall.RuleWAFManagedRulesTraditionalDenyRule] or
// [firewall.RuleWAFManagedRulesTraditionalAllowRule].
type RuleUnion interface {
	implementsFirewallRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleWAFManagedRulesAnomalyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleWAFManagedRulesTraditionalDenyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleWAFManagedRulesTraditionalAllowRule{}),
		},
	)
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type RuleWAFManagedRulesAnomalyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []AllowedModesAnomaly `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group UnnamedSchemaRef532d8b97684c9032dd36bae8acddebf5 `json:"group,required"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode RuleWAFManagedRulesAnomalyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                             `json:"priority,required"`
	JSON     ruleWAFManagedRulesAnomalyRuleJSON `json:"-"`
}

// ruleWAFManagedRulesAnomalyRuleJSON contains the JSON metadata for the struct
// [RuleWAFManagedRulesAnomalyRule]
type ruleWAFManagedRulesAnomalyRuleJSON struct {
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

func (r *RuleWAFManagedRulesAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleWAFManagedRulesAnomalyRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleWAFManagedRulesAnomalyRule) implementsFirewallRule() {}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type RuleWAFManagedRulesAnomalyRuleMode string

const (
	RuleWAFManagedRulesAnomalyRuleModeOn  RuleWAFManagedRulesAnomalyRuleMode = "on"
	RuleWAFManagedRulesAnomalyRuleModeOff RuleWAFManagedRulesAnomalyRuleMode = "off"
)

func (r RuleWAFManagedRulesAnomalyRuleMode) IsKnown() bool {
	switch r {
	case RuleWAFManagedRulesAnomalyRuleModeOn, RuleWAFManagedRulesAnomalyRuleModeOff:
		return true
	}
	return false
}

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type RuleWAFManagedRulesTraditionalDenyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The list of possible actions of the WAF rule when it is triggered.
	AllowedModes []RuleWAFManagedRulesTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// The default action/mode of a rule.
	DefaultMode RuleWAFManagedRulesTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group UnnamedSchemaRef532d8b97684c9032dd36bae8acddebf5 `json:"group,required"`
	// The action that the current WAF rule will perform when triggered. Applies to
	// traditional (deny) WAF rules.
	Mode RuleWAFManagedRulesTraditionalDenyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                     `json:"priority,required"`
	JSON     ruleWAFManagedRulesTraditionalDenyRuleJSON `json:"-"`
}

// ruleWAFManagedRulesTraditionalDenyRuleJSON contains the JSON metadata for the
// struct [RuleWAFManagedRulesTraditionalDenyRule]
type ruleWAFManagedRulesTraditionalDenyRuleJSON struct {
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

func (r *RuleWAFManagedRulesTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleWAFManagedRulesTraditionalDenyRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleWAFManagedRulesTraditionalDenyRule) implementsFirewallRule() {}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type RuleWAFManagedRulesTraditionalDenyRuleAllowedMode string

const (
	RuleWAFManagedRulesTraditionalDenyRuleAllowedModeDefault   RuleWAFManagedRulesTraditionalDenyRuleAllowedMode = "default"
	RuleWAFManagedRulesTraditionalDenyRuleAllowedModeDisable   RuleWAFManagedRulesTraditionalDenyRuleAllowedMode = "disable"
	RuleWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate  RuleWAFManagedRulesTraditionalDenyRuleAllowedMode = "simulate"
	RuleWAFManagedRulesTraditionalDenyRuleAllowedModeBlock     RuleWAFManagedRulesTraditionalDenyRuleAllowedMode = "block"
	RuleWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge RuleWAFManagedRulesTraditionalDenyRuleAllowedMode = "challenge"
)

func (r RuleWAFManagedRulesTraditionalDenyRuleAllowedMode) IsKnown() bool {
	switch r {
	case RuleWAFManagedRulesTraditionalDenyRuleAllowedModeDefault, RuleWAFManagedRulesTraditionalDenyRuleAllowedModeDisable, RuleWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate, RuleWAFManagedRulesTraditionalDenyRuleAllowedModeBlock, RuleWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge:
		return true
	}
	return false
}

// The default action/mode of a rule.
type RuleWAFManagedRulesTraditionalDenyRuleDefaultMode string

const (
	RuleWAFManagedRulesTraditionalDenyRuleDefaultModeDisable   RuleWAFManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	RuleWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate  RuleWAFManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	RuleWAFManagedRulesTraditionalDenyRuleDefaultModeBlock     RuleWAFManagedRulesTraditionalDenyRuleDefaultMode = "block"
	RuleWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge RuleWAFManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

func (r RuleWAFManagedRulesTraditionalDenyRuleDefaultMode) IsKnown() bool {
	switch r {
	case RuleWAFManagedRulesTraditionalDenyRuleDefaultModeDisable, RuleWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate, RuleWAFManagedRulesTraditionalDenyRuleDefaultModeBlock, RuleWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge:
		return true
	}
	return false
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type RuleWAFManagedRulesTraditionalDenyRuleMode string

const (
	RuleWAFManagedRulesTraditionalDenyRuleModeDefault   RuleWAFManagedRulesTraditionalDenyRuleMode = "default"
	RuleWAFManagedRulesTraditionalDenyRuleModeDisable   RuleWAFManagedRulesTraditionalDenyRuleMode = "disable"
	RuleWAFManagedRulesTraditionalDenyRuleModeSimulate  RuleWAFManagedRulesTraditionalDenyRuleMode = "simulate"
	RuleWAFManagedRulesTraditionalDenyRuleModeBlock     RuleWAFManagedRulesTraditionalDenyRuleMode = "block"
	RuleWAFManagedRulesTraditionalDenyRuleModeChallenge RuleWAFManagedRulesTraditionalDenyRuleMode = "challenge"
)

func (r RuleWAFManagedRulesTraditionalDenyRuleMode) IsKnown() bool {
	switch r {
	case RuleWAFManagedRulesTraditionalDenyRuleModeDefault, RuleWAFManagedRulesTraditionalDenyRuleModeDisable, RuleWAFManagedRulesTraditionalDenyRuleModeSimulate, RuleWAFManagedRulesTraditionalDenyRuleModeBlock, RuleWAFManagedRulesTraditionalDenyRuleModeChallenge:
		return true
	}
	return false
}

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type RuleWAFManagedRulesTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []RuleWAFManagedRulesTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	DefaultMode  interface{}                                          `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group UnnamedSchemaRef532d8b97684c9032dd36bae8acddebf5 `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode RuleWAFManagedRulesTraditionalAllowRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                      `json:"priority,required"`
	JSON     ruleWAFManagedRulesTraditionalAllowRuleJSON `json:"-"`
}

// ruleWAFManagedRulesTraditionalAllowRuleJSON contains the JSON metadata for the
// struct [RuleWAFManagedRulesTraditionalAllowRule]
type ruleWAFManagedRulesTraditionalAllowRuleJSON struct {
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

func (r *RuleWAFManagedRulesTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleWAFManagedRulesTraditionalAllowRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleWAFManagedRulesTraditionalAllowRule) implementsFirewallRule() {}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type RuleWAFManagedRulesTraditionalAllowRuleAllowedMode string

const (
	RuleWAFManagedRulesTraditionalAllowRuleAllowedModeOn  RuleWAFManagedRulesTraditionalAllowRuleAllowedMode = "on"
	RuleWAFManagedRulesTraditionalAllowRuleAllowedModeOff RuleWAFManagedRulesTraditionalAllowRuleAllowedMode = "off"
)

func (r RuleWAFManagedRulesTraditionalAllowRuleAllowedMode) IsKnown() bool {
	switch r {
	case RuleWAFManagedRulesTraditionalAllowRuleAllowedModeOn, RuleWAFManagedRulesTraditionalAllowRuleAllowedModeOff:
		return true
	}
	return false
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type RuleWAFManagedRulesTraditionalAllowRuleMode string

const (
	RuleWAFManagedRulesTraditionalAllowRuleModeOn  RuleWAFManagedRulesTraditionalAllowRuleMode = "on"
	RuleWAFManagedRulesTraditionalAllowRuleModeOff RuleWAFManagedRulesTraditionalAllowRuleMode = "off"
)

func (r RuleWAFManagedRulesTraditionalAllowRuleMode) IsKnown() bool {
	switch r {
	case RuleWAFManagedRulesTraditionalAllowRuleModeOn, RuleWAFManagedRulesTraditionalAllowRuleModeOff:
		return true
	}
	return false
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type RuleMode string

const (
	RuleModeOn        RuleMode = "on"
	RuleModeOff       RuleMode = "off"
	RuleModeDefault   RuleMode = "default"
	RuleModeDisable   RuleMode = "disable"
	RuleModeSimulate  RuleMode = "simulate"
	RuleModeBlock     RuleMode = "block"
	RuleModeChallenge RuleMode = "challenge"
)

func (r RuleMode) IsKnown() bool {
	switch r {
	case RuleModeOn, RuleModeOff, RuleModeDefault, RuleModeDisable, RuleModeSimulate, RuleModeBlock, RuleModeChallenge:
		return true
	}
	return false
}

// The rule group to which the current WAF rule belongs.
type UnnamedSchemaRef532d8b97684c9032dd36bae8acddebf5 struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                               `json:"name"`
	JSON unnamedSchemaRef532d8b97684c9032dd36bae8acddebf5JSON `json:"-"`
}

// unnamedSchemaRef532d8b97684c9032dd36bae8acddebf5JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef532d8b97684c9032dd36bae8acddebf5]
type unnamedSchemaRef532d8b97684c9032dd36bae8acddebf5JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef532d8b97684c9032dd36bae8acddebf5) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef532d8b97684c9032dd36bae8acddebf5JSON) RawJSON() string {
	return r.raw
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type WAFPackageRuleEditResponse struct {
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group UnnamedSchemaRef532d8b97684c9032dd36bae8acddebf5 `json:"group,required"`
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority     string      `json:"priority,required"`
	AllowedModes interface{} `json:"allowed_modes"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode        WAFPackageRuleEditResponseMode `json:"mode,required"`
	DefaultMode interface{}                    `json:"default_mode,required"`
	JSON        wafPackageRuleEditResponseJSON `json:"-"`
	union       WAFPackageRuleEditResponseUnion
}

// wafPackageRuleEditResponseJSON contains the JSON metadata for the struct
// [WAFPackageRuleEditResponse]
type wafPackageRuleEditResponseJSON struct {
	Description  apijson.Field
	Group        apijson.Field
	ID           apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	AllowedModes apijson.Field
	Mode         apijson.Field
	DefaultMode  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r wafPackageRuleEditResponseJSON) RawJSON() string {
	return r.raw
}

func (r *WAFPackageRuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WAFPackageRuleEditResponse) AsUnion() WAFPackageRuleEditResponseUnion {
	return r.union
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by
// [firewall.WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule],
// [firewall.WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule] or
// [firewall.WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule].
type WAFPackageRuleEditResponseUnion interface {
	implementsFirewallWAFPackageRuleEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageRuleEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule{}),
		},
	)
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []AllowedModesAnomaly `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group UnnamedSchemaRef532d8b97684c9032dd36bae8acddebf5 `json:"group,required"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                   `json:"priority,required"`
	JSON     wafPackageRuleEditResponseWAFManagedRulesAnomalyRuleJSON `json:"-"`
}

// wafPackageRuleEditResponseWAFManagedRulesAnomalyRuleJSON contains the JSON
// metadata for the struct [WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule]
type wafPackageRuleEditResponseWAFManagedRulesAnomalyRuleJSON struct {
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

func (r *WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseWAFManagedRulesAnomalyRuleJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule) implementsFirewallWAFPackageRuleEditResponse() {
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleModeOn  WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleMode = "on"
	WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleModeOff WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleMode = "off"
)

func (r WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleModeOn, WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleModeOff:
		return true
	}
	return false
}

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The list of possible actions of the WAF rule when it is triggered.
	AllowedModes []WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// The default action/mode of a rule.
	DefaultMode WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group UnnamedSchemaRef532d8b97684c9032dd36bae8acddebf5 `json:"group,required"`
	// The action that the current WAF rule will perform when triggered. Applies to
	// traditional (deny) WAF rules.
	Mode WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                           `json:"priority,required"`
	JSON     wafPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleJSON `json:"-"`
}

// wafPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleJSON contains the
// JSON metadata for the struct
// [WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule]
type wafPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleJSON struct {
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

func (r *WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule) implementsFirewallWAFPackageRuleEditResponse() {
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDefault   WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "default"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDisable   WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "disable"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate  WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "simulate"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeBlock     WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "block"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "challenge"
)

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDefault, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDisable, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeBlock, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge:
		return true
	}
	return false
}

// The default action/mode of a rule.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeDisable   WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate  WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeBlock     WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "block"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeDisable, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeBlock, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge:
		return true
	}
	return false
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeDefault   WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "default"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeDisable   WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "disable"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeSimulate  WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "simulate"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeBlock     WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "block"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeChallenge WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "challenge"
)

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeDefault, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeDisable, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeSimulate, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeBlock, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeChallenge:
		return true
	}
	return false
}

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	DefaultMode  interface{}                                                                `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group UnnamedSchemaRef532d8b97684c9032dd36bae8acddebf5 `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                            `json:"priority,required"`
	JSON     wafPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleJSON `json:"-"`
}

// wafPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleJSON contains the
// JSON metadata for the struct
// [WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule]
type wafPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleJSON struct {
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

func (r *WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule) implementsFirewallWAFPackageRuleEditResponse() {
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOn  WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "on"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOff WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "off"
)

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOn, WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOff:
		return true
	}
	return false
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleModeOn  WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode = "on"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleModeOff WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode = "off"
)

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleModeOn, WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleModeOff:
		return true
	}
	return false
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type WAFPackageRuleEditResponseMode string

const (
	WAFPackageRuleEditResponseModeOn        WAFPackageRuleEditResponseMode = "on"
	WAFPackageRuleEditResponseModeOff       WAFPackageRuleEditResponseMode = "off"
	WAFPackageRuleEditResponseModeDefault   WAFPackageRuleEditResponseMode = "default"
	WAFPackageRuleEditResponseModeDisable   WAFPackageRuleEditResponseMode = "disable"
	WAFPackageRuleEditResponseModeSimulate  WAFPackageRuleEditResponseMode = "simulate"
	WAFPackageRuleEditResponseModeBlock     WAFPackageRuleEditResponseMode = "block"
	WAFPackageRuleEditResponseModeChallenge WAFPackageRuleEditResponseMode = "challenge"
)

func (r WAFPackageRuleEditResponseMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseModeOn, WAFPackageRuleEditResponseModeOff, WAFPackageRuleEditResponseModeDefault, WAFPackageRuleEditResponseModeDisable, WAFPackageRuleEditResponseModeSimulate, WAFPackageRuleEditResponseModeBlock, WAFPackageRuleEditResponseModeChallenge:
		return true
	}
	return false
}

type WAFPackageRuleListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The public description of the WAF rule.
	Description param.Field[string] `query:"description"`
	// The direction used to sort returned rules.
	Direction param.Field[WAFPackageRuleListParamsDirection] `query:"direction"`
	// The unique identifier of the rule group.
	GroupID param.Field[string] `query:"group_id"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[WAFPackageRuleListParamsMatch] `query:"match"`
	// The action/mode a rule has been overridden to perform.
	Mode param.Field[WAFPackageRuleListParamsMode] `query:"mode"`
	// The field used to sort returned rules.
	Order param.Field[WAFPackageRuleListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of rules per page.
	PerPage param.Field[float64] `query:"per_page"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority param.Field[string] `query:"priority"`
}

// URLQuery serializes [WAFPackageRuleListParams]'s query parameters as
// `url.Values`.
func (r WAFPackageRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type WAFPackageRuleListParamsDirection string

const (
	WAFPackageRuleListParamsDirectionAsc  WAFPackageRuleListParamsDirection = "asc"
	WAFPackageRuleListParamsDirectionDesc WAFPackageRuleListParamsDirection = "desc"
)

func (r WAFPackageRuleListParamsDirection) IsKnown() bool {
	switch r {
	case WAFPackageRuleListParamsDirectionAsc, WAFPackageRuleListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type WAFPackageRuleListParamsMatch string

const (
	WAFPackageRuleListParamsMatchAny WAFPackageRuleListParamsMatch = "any"
	WAFPackageRuleListParamsMatchAll WAFPackageRuleListParamsMatch = "all"
)

func (r WAFPackageRuleListParamsMatch) IsKnown() bool {
	switch r {
	case WAFPackageRuleListParamsMatchAny, WAFPackageRuleListParamsMatchAll:
		return true
	}
	return false
}

// The action/mode a rule has been overridden to perform.
type WAFPackageRuleListParamsMode string

const (
	WAFPackageRuleListParamsModeDis WAFPackageRuleListParamsMode = "DIS"
	WAFPackageRuleListParamsModeChl WAFPackageRuleListParamsMode = "CHL"
	WAFPackageRuleListParamsModeBlk WAFPackageRuleListParamsMode = "BLK"
	WAFPackageRuleListParamsModeSim WAFPackageRuleListParamsMode = "SIM"
)

func (r WAFPackageRuleListParamsMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleListParamsModeDis, WAFPackageRuleListParamsModeChl, WAFPackageRuleListParamsModeBlk, WAFPackageRuleListParamsModeSim:
		return true
	}
	return false
}

// The field used to sort returned rules.
type WAFPackageRuleListParamsOrder string

const (
	WAFPackageRuleListParamsOrderPriority    WAFPackageRuleListParamsOrder = "priority"
	WAFPackageRuleListParamsOrderGroupID     WAFPackageRuleListParamsOrder = "group_id"
	WAFPackageRuleListParamsOrderDescription WAFPackageRuleListParamsOrder = "description"
)

func (r WAFPackageRuleListParamsOrder) IsKnown() bool {
	switch r {
	case WAFPackageRuleListParamsOrderPriority, WAFPackageRuleListParamsOrderGroupID, WAFPackageRuleListParamsOrderDescription:
		return true
	}
	return false
}

type WAFPackageRuleEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The mode/action of the rule when triggered. You must use a value from the
	// `allowed_modes` array of the current rule.
	Mode param.Field[WAFPackageRuleEditParamsMode] `json:"mode"`
}

func (r WAFPackageRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The mode/action of the rule when triggered. You must use a value from the
// `allowed_modes` array of the current rule.
type WAFPackageRuleEditParamsMode string

const (
	WAFPackageRuleEditParamsModeDefault   WAFPackageRuleEditParamsMode = "default"
	WAFPackageRuleEditParamsModeDisable   WAFPackageRuleEditParamsMode = "disable"
	WAFPackageRuleEditParamsModeSimulate  WAFPackageRuleEditParamsMode = "simulate"
	WAFPackageRuleEditParamsModeBlock     WAFPackageRuleEditParamsMode = "block"
	WAFPackageRuleEditParamsModeChallenge WAFPackageRuleEditParamsMode = "challenge"
	WAFPackageRuleEditParamsModeOn        WAFPackageRuleEditParamsMode = "on"
	WAFPackageRuleEditParamsModeOff       WAFPackageRuleEditParamsMode = "off"
)

func (r WAFPackageRuleEditParamsMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditParamsModeDefault, WAFPackageRuleEditParamsModeDisable, WAFPackageRuleEditParamsModeSimulate, WAFPackageRuleEditParamsModeBlock, WAFPackageRuleEditParamsModeChallenge, WAFPackageRuleEditParamsModeOn, WAFPackageRuleEditParamsModeOff:
		return true
	}
	return false
}

type WAFPackageRuleEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// When triggered, anomaly detection WAF rules contribute to an overall threat
	// score that will determine if a request is considered malicious. You can
	// configure the total scoring threshold through the 'sensitivity' property of the
	// WAF package.
	Result WAFPackageRuleEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success WAFPackageRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    wafPackageRuleEditResponseEnvelopeJSON    `json:"-"`
}

// wafPackageRuleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFPackageRuleEditResponseEnvelope]
type wafPackageRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WAFPackageRuleEditResponseEnvelopeSuccess bool

const (
	WAFPackageRuleEditResponseEnvelopeSuccessTrue WAFPackageRuleEditResponseEnvelopeSuccess = true
)

func (r WAFPackageRuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WAFPackageRuleGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WAFPackageRuleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                                        `json:"errors,required"`
	Messages []shared.ResponseInfo                                        `json:"messages,required"`
	Result   shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion `json:"result,required"`
	// Whether the API call was successful
	Success WAFPackageRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    wafPackageRuleGetResponseEnvelopeJSON    `json:"-"`
}

// wafPackageRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFPackageRuleGetResponseEnvelope]
type wafPackageRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WAFPackageRuleGetResponseEnvelopeSuccess bool

const (
	WAFPackageRuleGetResponseEnvelopeSuccessTrue WAFPackageRuleGetResponseEnvelopeSuccess = true
)

func (r WAFPackageRuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFPackageRuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
