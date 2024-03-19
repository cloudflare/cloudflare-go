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
func (r *WAFPackageRuleService) List(ctx context.Context, packageID string, params WAFPackageRuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[WAFManagedRulesRule], err error) {
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
func (r *WAFPackageRuleService) ListAutoPaging(ctx context.Context, packageID string, params WAFPackageRuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[WAFManagedRulesRule] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, packageID, params, opts...))
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
func (r *WAFPackageRuleService) Get(ctx context.Context, packageID string, ruleID string, query WAFPackageRuleGetParams, opts ...option.RequestOption) (res *WAFPackageRuleGetResponse, err error) {
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

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by [firewall.WAFManagedRulesRuleWAFManagedRulesAnomalyRule],
// [firewall.WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRule] or
// [firewall.WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRule].
type WAFManagedRulesRule interface {
	implementsFirewallWAFManagedRulesRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFManagedRulesRule)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFManagedRulesRuleWAFManagedRulesAnomalyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRule{}),
		},
	)
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type WAFManagedRulesRuleWAFManagedRulesAnomalyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []WAFManagedRulesRuleWAFManagedRulesAnomalyRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group WAFManagedRulesRuleWAFManagedRulesAnomalyRuleGroup `json:"group,required"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode WAFManagedRulesRuleWAFManagedRulesAnomalyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                            `json:"priority,required"`
	JSON     wafManagedRulesRuleWAFManagedRulesAnomalyRuleJSON `json:"-"`
}

// wafManagedRulesRuleWAFManagedRulesAnomalyRuleJSON contains the JSON metadata for
// the struct [WAFManagedRulesRuleWAFManagedRulesAnomalyRule]
type wafManagedRulesRuleWAFManagedRulesAnomalyRuleJSON struct {
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

func (r *WAFManagedRulesRuleWAFManagedRulesAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesRuleWAFManagedRulesAnomalyRuleJSON) RawJSON() string {
	return r.raw
}

func (r WAFManagedRulesRuleWAFManagedRulesAnomalyRule) implementsFirewallWAFManagedRulesRule() {}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type WAFManagedRulesRuleWAFManagedRulesAnomalyRuleAllowedMode string

const (
	WAFManagedRulesRuleWAFManagedRulesAnomalyRuleAllowedModeOn  WAFManagedRulesRuleWAFManagedRulesAnomalyRuleAllowedMode = "on"
	WAFManagedRulesRuleWAFManagedRulesAnomalyRuleAllowedModeOff WAFManagedRulesRuleWAFManagedRulesAnomalyRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type WAFManagedRulesRuleWAFManagedRulesAnomalyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                 `json:"name"`
	JSON wafManagedRulesRuleWAFManagedRulesAnomalyRuleGroupJSON `json:"-"`
}

// wafManagedRulesRuleWAFManagedRulesAnomalyRuleGroupJSON contains the JSON
// metadata for the struct [WAFManagedRulesRuleWAFManagedRulesAnomalyRuleGroup]
type wafManagedRulesRuleWAFManagedRulesAnomalyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFManagedRulesRuleWAFManagedRulesAnomalyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesRuleWAFManagedRulesAnomalyRuleGroupJSON) RawJSON() string {
	return r.raw
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type WAFManagedRulesRuleWAFManagedRulesAnomalyRuleMode string

const (
	WAFManagedRulesRuleWAFManagedRulesAnomalyRuleModeOn  WAFManagedRulesRuleWAFManagedRulesAnomalyRuleMode = "on"
	WAFManagedRulesRuleWAFManagedRulesAnomalyRuleModeOff WAFManagedRulesRuleWAFManagedRulesAnomalyRuleMode = "off"
)

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The list of possible actions of the WAF rule when it is triggered.
	AllowedModes []WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// The default action/mode of a rule.
	DefaultMode WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleGroup `json:"group,required"`
	// The action that the current WAF rule will perform when triggered. Applies to
	// traditional (deny) WAF rules.
	Mode WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                    `json:"priority,required"`
	JSON     wafManagedRulesRuleWAFManagedRulesTraditionalDenyRuleJSON `json:"-"`
}

// wafManagedRulesRuleWAFManagedRulesTraditionalDenyRuleJSON contains the JSON
// metadata for the struct [WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRule]
type wafManagedRulesRuleWAFManagedRulesTraditionalDenyRuleJSON struct {
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

func (r *WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesRuleWAFManagedRulesTraditionalDenyRuleJSON) RawJSON() string {
	return r.raw
}

func (r WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRule) implementsFirewallWAFManagedRulesRule() {
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedMode string

const (
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedModeDefault   WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedMode = "default"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedModeDisable   WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedMode = "disable"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate  WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedMode = "simulate"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedModeBlock     WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedMode = "block"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedMode = "challenge"
)

// The default action/mode of a rule.
type WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultMode string

const (
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultModeDisable   WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate  WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultModeBlock     WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultMode = "block"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

// The rule group to which the current WAF rule belongs.
type WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                         `json:"name"`
	JSON wafManagedRulesRuleWAFManagedRulesTraditionalDenyRuleGroupJSON `json:"-"`
}

// wafManagedRulesRuleWAFManagedRulesTraditionalDenyRuleGroupJSON contains the JSON
// metadata for the struct
// [WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleGroup]
type wafManagedRulesRuleWAFManagedRulesTraditionalDenyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesRuleWAFManagedRulesTraditionalDenyRuleGroupJSON) RawJSON() string {
	return r.raw
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleMode string

const (
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleModeDefault   WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleMode = "default"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleModeDisable   WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleMode = "disable"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleModeSimulate  WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleMode = "simulate"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleModeBlock     WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleMode = "block"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleModeChallenge WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleMode = "challenge"
)

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                     `json:"priority,required"`
	JSON     wafManagedRulesRuleWAFManagedRulesTraditionalAllowRuleJSON `json:"-"`
}

// wafManagedRulesRuleWAFManagedRulesTraditionalAllowRuleJSON contains the JSON
// metadata for the struct [WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRule]
type wafManagedRulesRuleWAFManagedRulesTraditionalAllowRuleJSON struct {
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

func (r *WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesRuleWAFManagedRulesTraditionalAllowRuleJSON) RawJSON() string {
	return r.raw
}

func (r WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRule) implementsFirewallWAFManagedRulesRule() {
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleAllowedMode string

const (
	WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleAllowedModeOn  WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleAllowedMode = "on"
	WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleAllowedModeOff WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                          `json:"name"`
	JSON wafManagedRulesRuleWAFManagedRulesTraditionalAllowRuleGroupJSON `json:"-"`
}

// wafManagedRulesRuleWAFManagedRulesTraditionalAllowRuleGroupJSON contains the
// JSON metadata for the struct
// [WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleGroup]
type wafManagedRulesRuleWAFManagedRulesTraditionalAllowRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesRuleWAFManagedRulesTraditionalAllowRuleGroupJSON) RawJSON() string {
	return r.raw
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleMode string

const (
	WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleModeOn  WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleMode = "on"
	WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleModeOff WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleMode = "off"
)

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by
// [firewall.WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule],
// [firewall.WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule] or
// [firewall.WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule].
type WAFPackageRuleEditResponse interface {
	implementsFirewallWAFPackageRuleEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageRuleEditResponse)(nil)).Elem(),
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
	AllowedModes []WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroup `json:"group,required"`
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
type WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedModeOn  WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedMode = "on"
	WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedModeOff WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                        `json:"name"`
	JSON wafPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroupJSON `json:"-"`
}

// wafPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroupJSON contains the JSON
// metadata for the struct
// [WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroup]
type wafPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseWAFManagedRulesAnomalyRuleGroupJSON) RawJSON() string {
	return r.raw
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleModeOn  WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleMode = "on"
	WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleModeOff WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleMode = "off"
)

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
	Group WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroup `json:"group,required"`
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

// The default action/mode of a rule.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeDisable   WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate  WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeBlock     WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "block"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

// The rule group to which the current WAF rule belongs.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                `json:"name"`
	JSON wafPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroupJSON `json:"-"`
}

// wafPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroupJSON contains
// the JSON metadata for the struct
// [WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroup]
type wafPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleGroupJSON) RawJSON() string {
	return r.raw
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

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroup `json:"group,required"`
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

// The rule group to which the current WAF rule belongs.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                 `json:"name"`
	JSON wafPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroupJSON `json:"-"`
}

// wafPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroupJSON contains
// the JSON metadata for the struct
// [WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroup]
type wafPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleGroupJSON) RawJSON() string {
	return r.raw
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleModeOn  WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode = "on"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleModeOff WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode = "off"
)

// Union satisfied by [firewall.WAFPackageRuleGetResponseUnknown],
// [firewall.WAFPackageRuleGetResponseArray] or [shared.UnionString].
type WAFPackageRuleGetResponse interface {
	ImplementsFirewallWAFPackageRuleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageRuleGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageRuleGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WAFPackageRuleGetResponseArray []interface{}

func (r WAFPackageRuleGetResponseArray) ImplementsFirewallWAFPackageRuleGetResponse() {}

type WAFPackageRuleListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned rules.
	Direction param.Field[WAFPackageRuleListParamsDirection] `query:"direction"`
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
}

// URLQuery serializes [WAFPackageRuleListParams]'s query parameters as
// `url.Values`.
func (r WAFPackageRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type WAFPackageRuleListParamsDirection string

const (
	WAFPackageRuleListParamsDirectionAsc  WAFPackageRuleListParamsDirection = "asc"
	WAFPackageRuleListParamsDirectionDesc WAFPackageRuleListParamsDirection = "desc"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type WAFPackageRuleListParamsMatch string

const (
	WAFPackageRuleListParamsMatchAny WAFPackageRuleListParamsMatch = "any"
	WAFPackageRuleListParamsMatchAll WAFPackageRuleListParamsMatch = "all"
)

// The action/mode a rule has been overridden to perform.
type WAFPackageRuleListParamsMode string

const (
	WAFPackageRuleListParamsModeDis WAFPackageRuleListParamsMode = "DIS"
	WAFPackageRuleListParamsModeChl WAFPackageRuleListParamsMode = "CHL"
	WAFPackageRuleListParamsModeBlk WAFPackageRuleListParamsMode = "BLK"
	WAFPackageRuleListParamsModeSim WAFPackageRuleListParamsMode = "SIM"
)

// The field used to sort returned rules.
type WAFPackageRuleListParamsOrder string

const (
	WAFPackageRuleListParamsOrderPriority    WAFPackageRuleListParamsOrder = "priority"
	WAFPackageRuleListParamsOrderGroupID     WAFPackageRuleListParamsOrder = "group_id"
	WAFPackageRuleListParamsOrderDescription WAFPackageRuleListParamsOrder = "description"
)

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

type WAFPackageRuleEditResponseEnvelope struct {
	Errors   []WAFPackageRuleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WAFPackageRuleEditResponseEnvelopeMessages `json:"messages,required"`
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

type WAFPackageRuleEditResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    wafPackageRuleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// wafPackageRuleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WAFPackageRuleEditResponseEnvelopeErrors]
type wafPackageRuleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WAFPackageRuleEditResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    wafPackageRuleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// wafPackageRuleEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WAFPackageRuleEditResponseEnvelopeMessages]
type wafPackageRuleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WAFPackageRuleEditResponseEnvelopeSuccess bool

const (
	WAFPackageRuleEditResponseEnvelopeSuccessTrue WAFPackageRuleEditResponseEnvelopeSuccess = true
)

type WAFPackageRuleGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WAFPackageRuleGetResponseEnvelope struct {
	Errors   []WAFPackageRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WAFPackageRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WAFPackageRuleGetResponse                   `json:"result,required"`
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

type WAFPackageRuleGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    wafPackageRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// wafPackageRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WAFPackageRuleGetResponseEnvelopeErrors]
type wafPackageRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WAFPackageRuleGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    wafPackageRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// wafPackageRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WAFPackageRuleGetResponseEnvelopeMessages]
type wafPackageRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WAFPackageRuleGetResponseEnvelopeSuccess bool

const (
	WAFPackageRuleGetResponseEnvelopeSuccessTrue WAFPackageRuleGetResponseEnvelopeSuccess = true
)
