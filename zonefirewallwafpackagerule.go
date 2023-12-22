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
func (r *ZoneFirewallWafPackageRuleService) Get(ctx context.Context, zoneID string, packageID string, identifier string, opts ...option.RequestOption) (res *RuleResponseSingleDIRsTgpk, err error) {
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
func (r *ZoneFirewallWafPackageRuleService) WafRulesListWafRules(ctx context.Context, zoneID string, packageID string, query ZoneFirewallWafPackageRuleWafRulesListWafRulesParams, opts ...option.RequestOption) (res *RuleResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules", zoneID, packageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RuleResponseCollection struct {
	Errors     []RuleResponseCollectionError    `json:"errors"`
	Messages   []RuleResponseCollectionMessage  `json:"messages"`
	Result     []RuleResponseCollectionResult   `json:"result"`
	ResultInfo RuleResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success RuleResponseCollectionSuccess `json:"success"`
	JSON    ruleResponseCollectionJSON    `json:"-"`
}

// ruleResponseCollectionJSON contains the JSON metadata for the struct
// [RuleResponseCollection]
type ruleResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleResponseCollectionError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    ruleResponseCollectionErrorJSON `json:"-"`
}

// ruleResponseCollectionErrorJSON contains the JSON metadata for the struct
// [RuleResponseCollectionError]
type ruleResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleResponseCollectionMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    ruleResponseCollectionMessageJSON `json:"-"`
}

// ruleResponseCollectionMessageJSON contains the JSON metadata for the struct
// [RuleResponseCollectionMessage]
type ruleResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by [RuleResponseCollectionResultAnomalyRule],
// [RuleResponseCollectionResultTraditionalDenyRule] or
// [RuleResponseCollectionResultTraditionalAllowRule].
type RuleResponseCollectionResult interface {
	implementsRuleResponseCollectionResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*RuleResponseCollectionResult)(nil)).Elem(), "")
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type RuleResponseCollectionResultAnomalyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []RuleResponseCollectionResultAnomalyRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group RuleResponseCollectionResultAnomalyRuleGroup `json:"group,required"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode RuleResponseCollectionResultAnomalyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                      `json:"priority,required"`
	JSON     ruleResponseCollectionResultAnomalyRuleJSON `json:"-"`
}

// ruleResponseCollectionResultAnomalyRuleJSON contains the JSON metadata for the
// struct [RuleResponseCollectionResultAnomalyRule]
type ruleResponseCollectionResultAnomalyRuleJSON struct {
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

func (r *RuleResponseCollectionResultAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleResponseCollectionResultAnomalyRule) implementsRuleResponseCollectionResult() {}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type RuleResponseCollectionResultAnomalyRuleAllowedMode string

const (
	RuleResponseCollectionResultAnomalyRuleAllowedModeOn  RuleResponseCollectionResultAnomalyRuleAllowedMode = "on"
	RuleResponseCollectionResultAnomalyRuleAllowedModeOff RuleResponseCollectionResultAnomalyRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type RuleResponseCollectionResultAnomalyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                           `json:"name"`
	JSON ruleResponseCollectionResultAnomalyRuleGroupJSON `json:"-"`
}

// ruleResponseCollectionResultAnomalyRuleGroupJSON contains the JSON metadata for
// the struct [RuleResponseCollectionResultAnomalyRuleGroup]
type ruleResponseCollectionResultAnomalyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseCollectionResultAnomalyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type RuleResponseCollectionResultAnomalyRuleMode string

const (
	RuleResponseCollectionResultAnomalyRuleModeOn  RuleResponseCollectionResultAnomalyRuleMode = "on"
	RuleResponseCollectionResultAnomalyRuleModeOff RuleResponseCollectionResultAnomalyRuleMode = "off"
)

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type RuleResponseCollectionResultTraditionalDenyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The list of possible actions of the WAF rule when it is triggered.
	AllowedModes []RuleResponseCollectionResultTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// The default action/mode of a rule.
	DefaultMode RuleResponseCollectionResultTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group RuleResponseCollectionResultTraditionalDenyRuleGroup `json:"group,required"`
	// The action that the current WAF rule will perform when triggered. Applies to
	// traditional (deny) WAF rules.
	Mode RuleResponseCollectionResultTraditionalDenyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                              `json:"priority,required"`
	JSON     ruleResponseCollectionResultTraditionalDenyRuleJSON `json:"-"`
}

// ruleResponseCollectionResultTraditionalDenyRuleJSON contains the JSON metadata
// for the struct [RuleResponseCollectionResultTraditionalDenyRule]
type ruleResponseCollectionResultTraditionalDenyRuleJSON struct {
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

func (r *RuleResponseCollectionResultTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleResponseCollectionResultTraditionalDenyRule) implementsRuleResponseCollectionResult() {}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type RuleResponseCollectionResultTraditionalDenyRuleAllowedMode string

const (
	RuleResponseCollectionResultTraditionalDenyRuleAllowedModeDefault   RuleResponseCollectionResultTraditionalDenyRuleAllowedMode = "default"
	RuleResponseCollectionResultTraditionalDenyRuleAllowedModeDisable   RuleResponseCollectionResultTraditionalDenyRuleAllowedMode = "disable"
	RuleResponseCollectionResultTraditionalDenyRuleAllowedModeSimulate  RuleResponseCollectionResultTraditionalDenyRuleAllowedMode = "simulate"
	RuleResponseCollectionResultTraditionalDenyRuleAllowedModeBlock     RuleResponseCollectionResultTraditionalDenyRuleAllowedMode = "block"
	RuleResponseCollectionResultTraditionalDenyRuleAllowedModeChallenge RuleResponseCollectionResultTraditionalDenyRuleAllowedMode = "challenge"
)

// The default action/mode of a rule.
type RuleResponseCollectionResultTraditionalDenyRuleDefaultMode string

const (
	RuleResponseCollectionResultTraditionalDenyRuleDefaultModeDisable   RuleResponseCollectionResultTraditionalDenyRuleDefaultMode = "disable"
	RuleResponseCollectionResultTraditionalDenyRuleDefaultModeSimulate  RuleResponseCollectionResultTraditionalDenyRuleDefaultMode = "simulate"
	RuleResponseCollectionResultTraditionalDenyRuleDefaultModeBlock     RuleResponseCollectionResultTraditionalDenyRuleDefaultMode = "block"
	RuleResponseCollectionResultTraditionalDenyRuleDefaultModeChallenge RuleResponseCollectionResultTraditionalDenyRuleDefaultMode = "challenge"
)

// The rule group to which the current WAF rule belongs.
type RuleResponseCollectionResultTraditionalDenyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                   `json:"name"`
	JSON ruleResponseCollectionResultTraditionalDenyRuleGroupJSON `json:"-"`
}

// ruleResponseCollectionResultTraditionalDenyRuleGroupJSON contains the JSON
// metadata for the struct [RuleResponseCollectionResultTraditionalDenyRuleGroup]
type ruleResponseCollectionResultTraditionalDenyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseCollectionResultTraditionalDenyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type RuleResponseCollectionResultTraditionalDenyRuleMode string

const (
	RuleResponseCollectionResultTraditionalDenyRuleModeDefault   RuleResponseCollectionResultTraditionalDenyRuleMode = "default"
	RuleResponseCollectionResultTraditionalDenyRuleModeDisable   RuleResponseCollectionResultTraditionalDenyRuleMode = "disable"
	RuleResponseCollectionResultTraditionalDenyRuleModeSimulate  RuleResponseCollectionResultTraditionalDenyRuleMode = "simulate"
	RuleResponseCollectionResultTraditionalDenyRuleModeBlock     RuleResponseCollectionResultTraditionalDenyRuleMode = "block"
	RuleResponseCollectionResultTraditionalDenyRuleModeChallenge RuleResponseCollectionResultTraditionalDenyRuleMode = "challenge"
)

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type RuleResponseCollectionResultTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []RuleResponseCollectionResultTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group RuleResponseCollectionResultTraditionalAllowRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode RuleResponseCollectionResultTraditionalAllowRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                               `json:"priority,required"`
	JSON     ruleResponseCollectionResultTraditionalAllowRuleJSON `json:"-"`
}

// ruleResponseCollectionResultTraditionalAllowRuleJSON contains the JSON metadata
// for the struct [RuleResponseCollectionResultTraditionalAllowRule]
type ruleResponseCollectionResultTraditionalAllowRuleJSON struct {
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

func (r *RuleResponseCollectionResultTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RuleResponseCollectionResultTraditionalAllowRule) implementsRuleResponseCollectionResult() {}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type RuleResponseCollectionResultTraditionalAllowRuleAllowedMode string

const (
	RuleResponseCollectionResultTraditionalAllowRuleAllowedModeOn  RuleResponseCollectionResultTraditionalAllowRuleAllowedMode = "on"
	RuleResponseCollectionResultTraditionalAllowRuleAllowedModeOff RuleResponseCollectionResultTraditionalAllowRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type RuleResponseCollectionResultTraditionalAllowRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                    `json:"name"`
	JSON ruleResponseCollectionResultTraditionalAllowRuleGroupJSON `json:"-"`
}

// ruleResponseCollectionResultTraditionalAllowRuleGroupJSON contains the JSON
// metadata for the struct [RuleResponseCollectionResultTraditionalAllowRuleGroup]
type ruleResponseCollectionResultTraditionalAllowRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseCollectionResultTraditionalAllowRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type RuleResponseCollectionResultTraditionalAllowRuleMode string

const (
	RuleResponseCollectionResultTraditionalAllowRuleModeOn  RuleResponseCollectionResultTraditionalAllowRuleMode = "on"
	RuleResponseCollectionResultTraditionalAllowRuleModeOff RuleResponseCollectionResultTraditionalAllowRuleMode = "off"
)

type RuleResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       ruleResponseCollectionResultInfoJSON `json:"-"`
}

// ruleResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [RuleResponseCollectionResultInfo]
type ruleResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleResponseCollectionSuccess bool

const (
	RuleResponseCollectionSuccessTrue RuleResponseCollectionSuccess = true
)

type RuleResponseSingleDIRsTgpk struct {
	Errors   []RuleResponseSingleDIRsTgpkError   `json:"errors"`
	Messages []RuleResponseSingleDIRsTgpkMessage `json:"messages"`
	Result   interface{}                         `json:"result"`
	// Whether the API call was successful
	Success RuleResponseSingleDIRsTgpkSuccess `json:"success"`
	JSON    ruleResponseSingleDiRsTgpkJSON    `json:"-"`
}

// ruleResponseSingleDiRsTgpkJSON contains the JSON metadata for the struct
// [RuleResponseSingleDIRsTgpk]
type ruleResponseSingleDiRsTgpkJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseSingleDIRsTgpk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleResponseSingleDIRsTgpkError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    ruleResponseSingleDiRsTgpkErrorJSON `json:"-"`
}

// ruleResponseSingleDiRsTgpkErrorJSON contains the JSON metadata for the struct
// [RuleResponseSingleDIRsTgpkError]
type ruleResponseSingleDiRsTgpkErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseSingleDIRsTgpkError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleResponseSingleDIRsTgpkMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    ruleResponseSingleDiRsTgpkMessageJSON `json:"-"`
}

// ruleResponseSingleDiRsTgpkMessageJSON contains the JSON metadata for the struct
// [RuleResponseSingleDIRsTgpkMessage]
type ruleResponseSingleDiRsTgpkMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleResponseSingleDIRsTgpkMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleResponseSingleDIRsTgpkSuccess bool

const (
	RuleResponseSingleDIRsTgpkSuccessTrue RuleResponseSingleDIRsTgpkSuccess = true
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
// Union satisfied by [ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRule],
// [ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRule] or
// [ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRule].
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
type ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleGroup `json:"group,required"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                        `json:"priority,required"`
	JSON     zoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleJSON contains the JSON
// metadata for the struct
// [ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRule]
type zoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleJSON struct {
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

func (r *ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRule) implementsZoneFirewallWafPackageRuleUpdateResponseResult() {
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleAllowedMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleAllowedModeOn  ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleAllowedMode = "on"
	ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleAllowedModeOff ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                             `json:"name"`
	JSON zoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleGroupJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleGroupJSON contains the
// JSON metadata for the struct
// [ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleGroup]
type zoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleModeOn  ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleMode = "on"
	ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleModeOff ZoneFirewallWafPackageRuleUpdateResponseResultAnomalyRuleMode = "off"
)

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The list of possible actions of the WAF rule when it is triggered.
	AllowedModes []ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// The default action/mode of a rule.
	DefaultMode ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleGroup `json:"group,required"`
	// The action that the current WAF rule will perform when triggered. Applies to
	// traditional (deny) WAF rules.
	Mode ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                `json:"priority,required"`
	JSON     zoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleJSON contains
// the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRule]
type zoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleJSON struct {
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

func (r *ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRule) implementsZoneFirewallWafPackageRuleUpdateResponseResult() {
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleAllowedMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleAllowedModeDefault   ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleAllowedMode = "default"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleAllowedModeDisable   ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleAllowedMode = "disable"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleAllowedModeSimulate  ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleAllowedMode = "simulate"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleAllowedModeBlock     ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleAllowedMode = "block"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleAllowedModeChallenge ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleAllowedMode = "challenge"
)

// The default action/mode of a rule.
type ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleDefaultMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleDefaultModeDisable   ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleDefaultMode = "disable"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleDefaultModeSimulate  ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleDefaultMode = "simulate"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleDefaultModeBlock     ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleDefaultMode = "block"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleDefaultModeChallenge ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleDefaultMode = "challenge"
)

// The rule group to which the current WAF rule belongs.
type ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                     `json:"name"`
	JSON zoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleGroupJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleGroup]
type zoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleModeDefault   ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleMode = "default"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleModeDisable   ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleMode = "disable"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleModeSimulate  ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleMode = "simulate"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleModeBlock     ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleMode = "block"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleModeChallenge ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalDenyRuleMode = "challenge"
)

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                 `json:"priority,required"`
	JSON     zoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleJSON contains
// the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRule]
type zoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleJSON struct {
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

func (r *ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRule) implementsZoneFirewallWafPackageRuleUpdateResponseResult() {
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleAllowedMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleAllowedModeOn  ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleAllowedMode = "on"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleAllowedModeOff ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                      `json:"name"`
	JSON zoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleGroupJSON `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleGroup]
type zoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleModeOn  ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleMode = "on"
	ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleModeOff ZoneFirewallWafPackageRuleUpdateResponseResultTraditionalAllowRuleMode = "off"
)

// Whether the API call was successful
type ZoneFirewallWafPackageRuleUpdateResponseSuccess bool

const (
	ZoneFirewallWafPackageRuleUpdateResponseSuccessTrue ZoneFirewallWafPackageRuleUpdateResponseSuccess = true
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
