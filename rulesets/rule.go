// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRuleService] method instead.
type RuleService struct {
	Options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	return
}

// Adds a new rule to an account or zone ruleset. The rule will be added to the end
// of the existing list of rules in the ruleset by default.
func (r *RuleService) New(ctx context.Context, rulesetID string, params RuleNewParams, opts ...option.RequestOption) (res *Ruleset, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.getAccountID().Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.getAccountID()
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.getZoneID()
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing rule from an account or zone ruleset.
func (r *RuleService) Delete(ctx context.Context, rulesetID string, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) (res *Ruleset, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/%s", accountOrZone, accountOrZoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing rule in an account or zone ruleset.
func (r *RuleService) Edit(ctx context.Context, rulesetID string, ruleID string, params RuleEditParams, opts ...option.RequestOption) (res *Ruleset, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleEditResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.getAccountID().Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.getAccountID()
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.getZoneID()
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/%s", accountOrZone, accountOrZoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// This interface is a union satisfied by one of the following:
// [RuleNewParamsRulesetsBlockRule], [RuleNewParamsRulesetsExecuteRule],
// [RuleNewParamsRulesetsLogRule], [RuleNewParamsRulesetsSkipRule].
type RuleNewParams interface {
	ImplementsRuleNewParams()

	getAccountID() param.Field[string]

	getZoneID() param.Field[string]
}

type RuleNewParamsRulesetsBlockRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleNewParamsRulesetsBlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RuleNewParamsRulesetsBlockRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RuleNewParamsRulesetsBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleNewParamsRulesetsBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsRulesetsBlockRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleNewParamsRulesetsBlockRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleNewParamsRulesetsBlockRule) ImplementsRuleNewParams() {

}

// The action to perform when the rule matches.
type RuleNewParamsRulesetsBlockRuleAction string

const (
	RuleNewParamsRulesetsBlockRuleActionBlock RuleNewParamsRulesetsBlockRuleAction = "block"
)

func (r RuleNewParamsRulesetsBlockRuleAction) IsKnown() bool {
	switch r {
	case RuleNewParamsRulesetsBlockRuleActionBlock:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleNewParamsRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[RuleNewParamsRulesetsBlockRuleActionParametersResponse] `json:"response"`
}

func (r RuleNewParamsRulesetsBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type RuleNewParamsRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r RuleNewParamsRulesetsBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type RuleNewParamsRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RuleNewParamsRulesetsBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewParamsRulesetsExecuteRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleNewParamsRulesetsExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RuleNewParamsRulesetsExecuteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RuleNewParamsRulesetsExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleNewParamsRulesetsExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsRulesetsExecuteRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleNewParamsRulesetsExecuteRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleNewParamsRulesetsExecuteRule) ImplementsRuleNewParams() {

}

// The action to perform when the rule matches.
type RuleNewParamsRulesetsExecuteRuleAction string

const (
	RuleNewParamsRulesetsExecuteRuleActionExecute RuleNewParamsRulesetsExecuteRuleAction = "execute"
)

func (r RuleNewParamsRulesetsExecuteRuleAction) IsKnown() bool {
	switch r {
	case RuleNewParamsRulesetsExecuteRuleActionExecute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleNewParamsRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[RuleNewParamsRulesetsExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[RuleNewParamsRulesetsExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r RuleNewParamsRulesetsExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type RuleNewParamsRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r RuleNewParamsRulesetsExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type RuleNewParamsRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRule] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RuleNewParamsRulesetsExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for rules in the category.
type RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

func (r RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, RuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

// A rule-level override
type RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for the rule.
type RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

func (r RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, RuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault, RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium, RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow, RuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

// An object configuring the rule's logging behavior.
type RuleNewParamsRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RuleNewParamsRulesetsExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewParamsRulesetsLogRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleNewParamsRulesetsLogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RuleNewParamsRulesetsLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleNewParamsRulesetsLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsRulesetsLogRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleNewParamsRulesetsLogRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleNewParamsRulesetsLogRule) ImplementsRuleNewParams() {

}

// The action to perform when the rule matches.
type RuleNewParamsRulesetsLogRuleAction string

const (
	RuleNewParamsRulesetsLogRuleActionLog RuleNewParamsRulesetsLogRuleAction = "log"
)

func (r RuleNewParamsRulesetsLogRuleAction) IsKnown() bool {
	switch r {
	case RuleNewParamsRulesetsLogRuleActionLog:
		return true
	}
	return false
}

// An object configuring the rule's logging behavior.
type RuleNewParamsRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RuleNewParamsRulesetsLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewParamsRulesetsSkipRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleNewParamsRulesetsSkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RuleNewParamsRulesetsSkipRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RuleNewParamsRulesetsSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleNewParamsRulesetsSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsRulesetsSkipRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleNewParamsRulesetsSkipRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleNewParamsRulesetsSkipRule) ImplementsRuleNewParams() {

}

// The action to perform when the rule matches.
type RuleNewParamsRulesetsSkipRuleAction string

const (
	RuleNewParamsRulesetsSkipRuleActionSkip RuleNewParamsRulesetsSkipRuleAction = "skip"
)

func (r RuleNewParamsRulesetsSkipRuleAction) IsKnown() bool {
	switch r {
	case RuleNewParamsRulesetsSkipRuleActionSkip:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleNewParamsRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]RuleNewParamsRulesetsSkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]RuleNewParamsRulesetsSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[RuleNewParamsRulesetsSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r RuleNewParamsRulesetsSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A phase to skip the execution of.
type RuleNewParamsRulesetsSkipRuleActionParametersPhase string

const (
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseDDoSL4                         RuleNewParamsRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseDDoSL7                         RuleNewParamsRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseMagicTransit                   RuleNewParamsRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RuleNewParamsRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RuleNewParamsRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RuleNewParamsRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

func (r RuleNewParamsRulesetsSkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case RuleNewParamsRulesetsSkipRuleActionParametersPhaseDDoSL4, RuleNewParamsRulesetsSkipRuleActionParametersPhaseDDoSL7, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRatelimit, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged, RuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform, RuleNewParamsRulesetsSkipRuleActionParametersPhaseMagicTransit, RuleNewParamsRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged, RuleNewParamsRulesetsSkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

// The name of a legacy security product to skip the execution of.
type RuleNewParamsRulesetsSkipRuleActionParametersProduct string

const (
	RuleNewParamsRulesetsSkipRuleActionParametersProductBic           RuleNewParamsRulesetsSkipRuleActionParametersProduct = "bic"
	RuleNewParamsRulesetsSkipRuleActionParametersProductHot           RuleNewParamsRulesetsSkipRuleActionParametersProduct = "hot"
	RuleNewParamsRulesetsSkipRuleActionParametersProductRateLimit     RuleNewParamsRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RuleNewParamsRulesetsSkipRuleActionParametersProductSecurityLevel RuleNewParamsRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RuleNewParamsRulesetsSkipRuleActionParametersProductUABlock       RuleNewParamsRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RuleNewParamsRulesetsSkipRuleActionParametersProductWAF           RuleNewParamsRulesetsSkipRuleActionParametersProduct = "waf"
	RuleNewParamsRulesetsSkipRuleActionParametersProductZoneLockdown  RuleNewParamsRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

func (r RuleNewParamsRulesetsSkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case RuleNewParamsRulesetsSkipRuleActionParametersProductBic, RuleNewParamsRulesetsSkipRuleActionParametersProductHot, RuleNewParamsRulesetsSkipRuleActionParametersProductRateLimit, RuleNewParamsRulesetsSkipRuleActionParametersProductSecurityLevel, RuleNewParamsRulesetsSkipRuleActionParametersProductUABlock, RuleNewParamsRulesetsSkipRuleActionParametersProductWAF, RuleNewParamsRulesetsSkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RuleNewParamsRulesetsSkipRuleActionParametersRuleset string

const (
	RuleNewParamsRulesetsSkipRuleActionParametersRulesetCurrent RuleNewParamsRulesetsSkipRuleActionParametersRuleset = "current"
)

func (r RuleNewParamsRulesetsSkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case RuleNewParamsRulesetsSkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

// An object configuring the rule's logging behavior.
type RuleNewParamsRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RuleNewParamsRulesetsSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A response object.
type RuleNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleNewResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleNewResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result Ruleset `json:"result,required"`
	// Whether the API call was successful.
	Success RuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleNewResponseEnvelopeJSON    `json:"-"`
}

// ruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelope]
type ruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleNewResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleNewResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeErrors]
type ruleNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleNewResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                  `json:"pointer,required"`
	JSON    ruleNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleNewResponseEnvelopeErrorsSource]
type ruleNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleNewResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleNewResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeMessages]
type ruleNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleNewResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                    `json:"pointer,required"`
	JSON    ruleNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleNewResponseEnvelopeMessagesSource]
type ruleNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleNewResponseEnvelopeSuccess bool

const (
	RuleNewResponseEnvelopeSuccessTrue RuleNewResponseEnvelopeSuccess = true
)

func (r RuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type RuleDeleteResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleDeleteResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result Ruleset `json:"result,required"`
	// Whether the API call was successful.
	Success RuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelope]
type ruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleDeleteResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeErrors]
type ruleDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleDeleteResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    ruleDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleDeleteResponseEnvelopeErrorsSource]
type ruleDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleDeleteResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeMessages]
type ruleDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleDeleteResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    ruleDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleDeleteResponseEnvelopeMessagesSource]
type ruleDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleDeleteResponseEnvelopeSuccess bool

const (
	RuleDeleteResponseEnvelopeSuccessTrue RuleDeleteResponseEnvelopeSuccess = true
)

func (r RuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

// This interface is a union satisfied by one of the following:
// [RuleEditParamsRulesetsBlockRule], [RuleEditParamsRulesetsExecuteRule],
// [RuleEditParamsRulesetsLogRule], [RuleEditParamsRulesetsSkipRule].
type RuleEditParams interface {
	ImplementsRuleEditParams()

	getAccountID() param.Field[string]

	getZoneID() param.Field[string]
}

type RuleEditParamsRulesetsBlockRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleEditParamsRulesetsBlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RuleEditParamsRulesetsBlockRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RuleEditParamsRulesetsBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleEditParamsRulesetsBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsRulesetsBlockRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleEditParamsRulesetsBlockRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleEditParamsRulesetsBlockRule) ImplementsRuleEditParams() {

}

// The action to perform when the rule matches.
type RuleEditParamsRulesetsBlockRuleAction string

const (
	RuleEditParamsRulesetsBlockRuleActionBlock RuleEditParamsRulesetsBlockRuleAction = "block"
)

func (r RuleEditParamsRulesetsBlockRuleAction) IsKnown() bool {
	switch r {
	case RuleEditParamsRulesetsBlockRuleActionBlock:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleEditParamsRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[RuleEditParamsRulesetsBlockRuleActionParametersResponse] `json:"response"`
}

func (r RuleEditParamsRulesetsBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type RuleEditParamsRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r RuleEditParamsRulesetsBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type RuleEditParamsRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RuleEditParamsRulesetsBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleEditParamsRulesetsExecuteRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleEditParamsRulesetsExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RuleEditParamsRulesetsExecuteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RuleEditParamsRulesetsExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleEditParamsRulesetsExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsRulesetsExecuteRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleEditParamsRulesetsExecuteRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleEditParamsRulesetsExecuteRule) ImplementsRuleEditParams() {

}

// The action to perform when the rule matches.
type RuleEditParamsRulesetsExecuteRuleAction string

const (
	RuleEditParamsRulesetsExecuteRuleActionExecute RuleEditParamsRulesetsExecuteRuleAction = "execute"
)

func (r RuleEditParamsRulesetsExecuteRuleAction) IsKnown() bool {
	switch r {
	case RuleEditParamsRulesetsExecuteRuleActionExecute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleEditParamsRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[RuleEditParamsRulesetsExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[RuleEditParamsRulesetsExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r RuleEditParamsRulesetsExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type RuleEditParamsRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r RuleEditParamsRulesetsExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type RuleEditParamsRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRule] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RuleEditParamsRulesetsExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for rules in the category.
type RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

func (r RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, RuleEditParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

// A rule-level override
type RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for the rule.
type RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

func (r RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, RuleEditParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault, RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium, RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow, RuleEditParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

// An object configuring the rule's logging behavior.
type RuleEditParamsRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RuleEditParamsRulesetsExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleEditParamsRulesetsLogRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleEditParamsRulesetsLogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RuleEditParamsRulesetsLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleEditParamsRulesetsLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsRulesetsLogRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleEditParamsRulesetsLogRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleEditParamsRulesetsLogRule) ImplementsRuleEditParams() {

}

// The action to perform when the rule matches.
type RuleEditParamsRulesetsLogRuleAction string

const (
	RuleEditParamsRulesetsLogRuleActionLog RuleEditParamsRulesetsLogRuleAction = "log"
)

func (r RuleEditParamsRulesetsLogRuleAction) IsKnown() bool {
	switch r {
	case RuleEditParamsRulesetsLogRuleActionLog:
		return true
	}
	return false
}

// An object configuring the rule's logging behavior.
type RuleEditParamsRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RuleEditParamsRulesetsLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleEditParamsRulesetsSkipRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleEditParamsRulesetsSkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RuleEditParamsRulesetsSkipRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RuleEditParamsRulesetsSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleEditParamsRulesetsSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsRulesetsSkipRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleEditParamsRulesetsSkipRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleEditParamsRulesetsSkipRule) ImplementsRuleEditParams() {

}

// The action to perform when the rule matches.
type RuleEditParamsRulesetsSkipRuleAction string

const (
	RuleEditParamsRulesetsSkipRuleActionSkip RuleEditParamsRulesetsSkipRuleAction = "skip"
)

func (r RuleEditParamsRulesetsSkipRuleAction) IsKnown() bool {
	switch r {
	case RuleEditParamsRulesetsSkipRuleActionSkip:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleEditParamsRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]RuleEditParamsRulesetsSkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]RuleEditParamsRulesetsSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[RuleEditParamsRulesetsSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r RuleEditParamsRulesetsSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A phase to skip the execution of.
type RuleEditParamsRulesetsSkipRuleActionParametersPhase string

const (
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseDDoSL4                         RuleEditParamsRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseDDoSL7                         RuleEditParamsRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RuleEditParamsRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseMagicTransit                   RuleEditParamsRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RuleEditParamsRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RuleEditParamsRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RuleEditParamsRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

func (r RuleEditParamsRulesetsSkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case RuleEditParamsRulesetsSkipRuleActionParametersPhaseDDoSL4, RuleEditParamsRulesetsSkipRuleActionParametersPhaseDDoSL7, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRatelimit, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged, RuleEditParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform, RuleEditParamsRulesetsSkipRuleActionParametersPhaseMagicTransit, RuleEditParamsRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged, RuleEditParamsRulesetsSkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

// The name of a legacy security product to skip the execution of.
type RuleEditParamsRulesetsSkipRuleActionParametersProduct string

const (
	RuleEditParamsRulesetsSkipRuleActionParametersProductBic           RuleEditParamsRulesetsSkipRuleActionParametersProduct = "bic"
	RuleEditParamsRulesetsSkipRuleActionParametersProductHot           RuleEditParamsRulesetsSkipRuleActionParametersProduct = "hot"
	RuleEditParamsRulesetsSkipRuleActionParametersProductRateLimit     RuleEditParamsRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RuleEditParamsRulesetsSkipRuleActionParametersProductSecurityLevel RuleEditParamsRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RuleEditParamsRulesetsSkipRuleActionParametersProductUABlock       RuleEditParamsRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RuleEditParamsRulesetsSkipRuleActionParametersProductWAF           RuleEditParamsRulesetsSkipRuleActionParametersProduct = "waf"
	RuleEditParamsRulesetsSkipRuleActionParametersProductZoneLockdown  RuleEditParamsRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

func (r RuleEditParamsRulesetsSkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case RuleEditParamsRulesetsSkipRuleActionParametersProductBic, RuleEditParamsRulesetsSkipRuleActionParametersProductHot, RuleEditParamsRulesetsSkipRuleActionParametersProductRateLimit, RuleEditParamsRulesetsSkipRuleActionParametersProductSecurityLevel, RuleEditParamsRulesetsSkipRuleActionParametersProductUABlock, RuleEditParamsRulesetsSkipRuleActionParametersProductWAF, RuleEditParamsRulesetsSkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RuleEditParamsRulesetsSkipRuleActionParametersRuleset string

const (
	RuleEditParamsRulesetsSkipRuleActionParametersRulesetCurrent RuleEditParamsRulesetsSkipRuleActionParametersRuleset = "current"
)

func (r RuleEditParamsRulesetsSkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case RuleEditParamsRulesetsSkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

// An object configuring the rule's logging behavior.
type RuleEditParamsRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RuleEditParamsRulesetsSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A response object.
type RuleEditResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleEditResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleEditResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result Ruleset `json:"result,required"`
	// Whether the API call was successful.
	Success RuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleEditResponseEnvelopeJSON    `json:"-"`
}

// ruleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelope]
type ruleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleEditResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleEditResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeErrors]
type ruleEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleEditResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                   `json:"pointer,required"`
	JSON    ruleEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleEditResponseEnvelopeErrorsSource]
type ruleEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleEditResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleEditResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeMessages]
type ruleEditResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleEditResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    ruleEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleEditResponseEnvelopeMessagesSource]
type ruleEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleEditResponseEnvelopeSuccess bool

const (
	RuleEditResponseEnvelopeSuccessTrue RuleEditResponseEnvelopeSuccess = true
)

func (r RuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
