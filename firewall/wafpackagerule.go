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
func (r *WAFPackageRuleService) List(ctx context.Context, packageID string, params WAFPackageRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WAFManagedRulesRule], err error) {
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
func (r *WAFPackageRuleService) ListAutoPaging(ctx context.Context, packageID string, params WAFPackageRuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WAFManagedRulesRule] {
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
	Group shared.UnnamedSchemaRef120 `json:"group,required"`
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

func (r WAFManagedRulesRuleWAFManagedRulesAnomalyRuleAllowedMode) IsKnown() bool {
	switch r {
	case WAFManagedRulesRuleWAFManagedRulesAnomalyRuleAllowedModeOn, WAFManagedRulesRuleWAFManagedRulesAnomalyRuleAllowedModeOff:
		return true
	}
	return false
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type WAFManagedRulesRuleWAFManagedRulesAnomalyRuleMode string

const (
	WAFManagedRulesRuleWAFManagedRulesAnomalyRuleModeOn  WAFManagedRulesRuleWAFManagedRulesAnomalyRuleMode = "on"
	WAFManagedRulesRuleWAFManagedRulesAnomalyRuleModeOff WAFManagedRulesRuleWAFManagedRulesAnomalyRuleMode = "off"
)

func (r WAFManagedRulesRuleWAFManagedRulesAnomalyRuleMode) IsKnown() bool {
	switch r {
	case WAFManagedRulesRuleWAFManagedRulesAnomalyRuleModeOn, WAFManagedRulesRuleWAFManagedRulesAnomalyRuleModeOff:
		return true
	}
	return false
}

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
	Group shared.UnnamedSchemaRef120 `json:"group,required"`
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

func (r WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedMode) IsKnown() bool {
	switch r {
	case WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedModeDefault, WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedModeDisable, WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate, WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedModeBlock, WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge:
		return true
	}
	return false
}

// The default action/mode of a rule.
type WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultMode string

const (
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultModeDisable   WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate  WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultModeBlock     WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultMode = "block"
	WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

func (r WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultMode) IsKnown() bool {
	switch r {
	case WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultModeDisable, WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate, WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultModeBlock, WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge:
		return true
	}
	return false
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

func (r WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleMode) IsKnown() bool {
	switch r {
	case WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleModeDefault, WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleModeDisable, WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleModeSimulate, WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleModeBlock, WAFManagedRulesRuleWAFManagedRulesTraditionalDenyRuleModeChallenge:
		return true
	}
	return false
}

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	DefaultMode  interface{}                                                         `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group shared.UnnamedSchemaRef120 `json:"group,required"`
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
	DefaultMode  apijson.Field
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

func (r WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleAllowedMode) IsKnown() bool {
	switch r {
	case WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleAllowedModeOn, WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleAllowedModeOff:
		return true
	}
	return false
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleMode string

const (
	WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleModeOn  WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleMode = "on"
	WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleModeOff WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleMode = "off"
)

func (r WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleMode) IsKnown() bool {
	switch r {
	case WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleModeOn, WAFManagedRulesRuleWAFManagedRulesTraditionalAllowRuleModeOff:
		return true
	}
	return false
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
	Group shared.UnnamedSchemaRef120 `json:"group,required"`
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

func (r WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedModeOn, WAFPackageRuleEditResponseWAFManagedRulesAnomalyRuleAllowedModeOff:
		return true
	}
	return false
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
	Group shared.UnnamedSchemaRef120 `json:"group,required"`
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
	Group shared.UnnamedSchemaRef120 `json:"group,required"`
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

// Union satisfied by [firewall.WAFPackageRuleGetResponseUnknown] or
// [shared.UnionString].
type WAFPackageRuleGetResponse interface {
	ImplementsFirewallWAFPackageRuleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageRuleGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   WAFPackageRuleGetResponse    `json:"result,required"`
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
