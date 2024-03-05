// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RulesetPhaseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRulesetPhaseService] method
// instead.
type RulesetPhaseService struct {
	Options  []option.RequestOption
	Versions *RulesetPhaseVersionService
}

// NewRulesetPhaseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRulesetPhaseService(opts ...option.RequestOption) (r *RulesetPhaseService) {
	r = &RulesetPhaseService{}
	r.Options = opts
	r.Versions = NewRulesetPhaseVersionService(opts...)
	return
}

// Updates an account or zone entry point ruleset, creating a new version.
func (r *RulesetPhaseService) Update(ctx context.Context, rulesetPhase RulesetPhaseUpdateParamsRulesetPhase, params RulesetPhaseUpdateParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetPhaseUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the latest version of the account or zone entry point ruleset for a
// given phase.
func (r *RulesetPhaseService) Get(ctx context.Context, rulesetPhase RulesetPhaseGetParamsRulesetPhase, query RulesetPhaseGetParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetPhaseGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RulesetPhaseUpdateParams struct {
	// The unique ID of the ruleset.
	ID param.Field[string] `json:"id,required"`
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetPhaseUpdateParamsRule] `json:"rules,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[RulesetPhaseUpdateParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[RulesetPhaseUpdateParamsPhase] `json:"phase"`
}

func (r RulesetPhaseUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The phase of the ruleset.
type RulesetPhaseUpdateParamsRulesetPhase string

const (
	RulesetPhaseUpdateParamsRulesetPhaseDDOSL4                         RulesetPhaseUpdateParamsRulesetPhase = "ddos_l4"
	RulesetPhaseUpdateParamsRulesetPhaseDDOSL7                         RulesetPhaseUpdateParamsRulesetPhase = "ddos_l7"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPConfigSettings             RulesetPhaseUpdateParamsRulesetPhase = "http_config_settings"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPCustomErrors               RulesetPhaseUpdateParamsRulesetPhase = "http_custom_errors"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPLogCustomFields            RulesetPhaseUpdateParamsRulesetPhase = "http_log_custom_fields"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPRatelimit                  RulesetPhaseUpdateParamsRulesetPhase = "http_ratelimit"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPRequestCacheSettings       RulesetPhaseUpdateParamsRulesetPhase = "http_request_cache_settings"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPRequestDynamicRedirect     RulesetPhaseUpdateParamsRulesetPhase = "http_request_dynamic_redirect"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPRequestFirewallCustom      RulesetPhaseUpdateParamsRulesetPhase = "http_request_firewall_custom"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPRequestFirewallManaged     RulesetPhaseUpdateParamsRulesetPhase = "http_request_firewall_managed"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPRequestLateTransform       RulesetPhaseUpdateParamsRulesetPhase = "http_request_late_transform"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPRequestOrigin              RulesetPhaseUpdateParamsRulesetPhase = "http_request_origin"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPRequestRedirect            RulesetPhaseUpdateParamsRulesetPhase = "http_request_redirect"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPRequestSanitize            RulesetPhaseUpdateParamsRulesetPhase = "http_request_sanitize"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPRequestSbfm                RulesetPhaseUpdateParamsRulesetPhase = "http_request_sbfm"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPRequestSelectConfiguration RulesetPhaseUpdateParamsRulesetPhase = "http_request_select_configuration"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPRequestTransform           RulesetPhaseUpdateParamsRulesetPhase = "http_request_transform"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPResponseCompression        RulesetPhaseUpdateParamsRulesetPhase = "http_response_compression"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPResponseFirewallManaged    RulesetPhaseUpdateParamsRulesetPhase = "http_response_firewall_managed"
	RulesetPhaseUpdateParamsRulesetPhaseHTTPResponseHeadersTransform   RulesetPhaseUpdateParamsRulesetPhase = "http_response_headers_transform"
	RulesetPhaseUpdateParamsRulesetPhaseMagicTransit                   RulesetPhaseUpdateParamsRulesetPhase = "magic_transit"
	RulesetPhaseUpdateParamsRulesetPhaseMagicTransitIDsManaged         RulesetPhaseUpdateParamsRulesetPhase = "magic_transit_ids_managed"
	RulesetPhaseUpdateParamsRulesetPhaseMagicTransitManaged            RulesetPhaseUpdateParamsRulesetPhase = "magic_transit_managed"
)

// Satisfied by [RulesetPhaseUpdateParamsRulesRulesetsBlockRule],
// [RulesetPhaseUpdateParamsRulesRulesetsExecuteRule],
// [RulesetPhaseUpdateParamsRulesRulesetsLogRule],
// [RulesetPhaseUpdateParamsRulesRulesetsSkipRule].
type RulesetPhaseUpdateParamsRule interface {
	implementsRulesetPhaseUpdateParamsRule()
}

type RulesetPhaseUpdateParamsRulesRulesetsBlockRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetPhaseUpdateParamsRulesRulesetsBlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetPhaseUpdateParamsRulesRulesetsBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetPhaseUpdateParamsRulesRulesetsBlockRule) implementsRulesetPhaseUpdateParamsRule() {}

// The action to perform when the rule matches.
type RulesetPhaseUpdateParamsRulesRulesetsBlockRuleAction string

const (
	RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionBlock RulesetPhaseUpdateParamsRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionParametersResponse] `json:"response"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type RulesetPhaseUpdateParamsRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetPhaseUpdateParamsRulesRulesetsExecuteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetPhaseUpdateParamsRulesRulesetsExecuteRule) implementsRulesetPhaseUpdateParamsRule() {}

// The action to perform when the rule matches.
type RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleAction string

const (
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionExecute RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRule] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for rules in the category.
type RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for the rule.
type RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetPhaseUpdateParamsRulesRulesetsLogRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetPhaseUpdateParamsRulesRulesetsLogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetPhaseUpdateParamsRulesRulesetsLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetPhaseUpdateParamsRulesRulesetsLogRule) implementsRulesetPhaseUpdateParamsRule() {}

// The action to perform when the rule matches.
type RulesetPhaseUpdateParamsRulesRulesetsLogRuleAction string

const (
	RulesetPhaseUpdateParamsRulesRulesetsLogRuleActionLog RulesetPhaseUpdateParamsRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetPhaseUpdateParamsRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetPhaseUpdateParamsRulesRulesetsSkipRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetPhaseUpdateParamsRulesRulesetsSkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetPhaseUpdateParamsRulesRulesetsSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetPhaseUpdateParamsRulesRulesetsSkipRule) implementsRulesetPhaseUpdateParamsRule() {}

// The action to perform when the rule matches.
type RulesetPhaseUpdateParamsRulesRulesetsSkipRuleAction string

const (
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionSkip RulesetPhaseUpdateParamsRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A phase to skip the execution of.
type RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductBic           RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductHot           RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductWAF           RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetPhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetPhaseUpdateParamsRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetPhaseUpdateParamsRulesRulesetsSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type RulesetPhaseUpdateParamsKind string

const (
	RulesetPhaseUpdateParamsKindManaged RulesetPhaseUpdateParamsKind = "managed"
	RulesetPhaseUpdateParamsKindCustom  RulesetPhaseUpdateParamsKind = "custom"
	RulesetPhaseUpdateParamsKindRoot    RulesetPhaseUpdateParamsKind = "root"
	RulesetPhaseUpdateParamsKindZone    RulesetPhaseUpdateParamsKind = "zone"
)

// The phase of the ruleset.
type RulesetPhaseUpdateParamsPhase string

const (
	RulesetPhaseUpdateParamsPhaseDDOSL4                         RulesetPhaseUpdateParamsPhase = "ddos_l4"
	RulesetPhaseUpdateParamsPhaseDDOSL7                         RulesetPhaseUpdateParamsPhase = "ddos_l7"
	RulesetPhaseUpdateParamsPhaseHTTPConfigSettings             RulesetPhaseUpdateParamsPhase = "http_config_settings"
	RulesetPhaseUpdateParamsPhaseHTTPCustomErrors               RulesetPhaseUpdateParamsPhase = "http_custom_errors"
	RulesetPhaseUpdateParamsPhaseHTTPLogCustomFields            RulesetPhaseUpdateParamsPhase = "http_log_custom_fields"
	RulesetPhaseUpdateParamsPhaseHTTPRatelimit                  RulesetPhaseUpdateParamsPhase = "http_ratelimit"
	RulesetPhaseUpdateParamsPhaseHTTPRequestCacheSettings       RulesetPhaseUpdateParamsPhase = "http_request_cache_settings"
	RulesetPhaseUpdateParamsPhaseHTTPRequestDynamicRedirect     RulesetPhaseUpdateParamsPhase = "http_request_dynamic_redirect"
	RulesetPhaseUpdateParamsPhaseHTTPRequestFirewallCustom      RulesetPhaseUpdateParamsPhase = "http_request_firewall_custom"
	RulesetPhaseUpdateParamsPhaseHTTPRequestFirewallManaged     RulesetPhaseUpdateParamsPhase = "http_request_firewall_managed"
	RulesetPhaseUpdateParamsPhaseHTTPRequestLateTransform       RulesetPhaseUpdateParamsPhase = "http_request_late_transform"
	RulesetPhaseUpdateParamsPhaseHTTPRequestOrigin              RulesetPhaseUpdateParamsPhase = "http_request_origin"
	RulesetPhaseUpdateParamsPhaseHTTPRequestRedirect            RulesetPhaseUpdateParamsPhase = "http_request_redirect"
	RulesetPhaseUpdateParamsPhaseHTTPRequestSanitize            RulesetPhaseUpdateParamsPhase = "http_request_sanitize"
	RulesetPhaseUpdateParamsPhaseHTTPRequestSbfm                RulesetPhaseUpdateParamsPhase = "http_request_sbfm"
	RulesetPhaseUpdateParamsPhaseHTTPRequestSelectConfiguration RulesetPhaseUpdateParamsPhase = "http_request_select_configuration"
	RulesetPhaseUpdateParamsPhaseHTTPRequestTransform           RulesetPhaseUpdateParamsPhase = "http_request_transform"
	RulesetPhaseUpdateParamsPhaseHTTPResponseCompression        RulesetPhaseUpdateParamsPhase = "http_response_compression"
	RulesetPhaseUpdateParamsPhaseHTTPResponseFirewallManaged    RulesetPhaseUpdateParamsPhase = "http_response_firewall_managed"
	RulesetPhaseUpdateParamsPhaseHTTPResponseHeadersTransform   RulesetPhaseUpdateParamsPhase = "http_response_headers_transform"
	RulesetPhaseUpdateParamsPhaseMagicTransit                   RulesetPhaseUpdateParamsPhase = "magic_transit"
	RulesetPhaseUpdateParamsPhaseMagicTransitIDsManaged         RulesetPhaseUpdateParamsPhase = "magic_transit_ids_managed"
	RulesetPhaseUpdateParamsPhaseMagicTransitManaged            RulesetPhaseUpdateParamsPhase = "magic_transit_managed"
)

// A response object.
type RulesetPhaseUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetPhaseUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetPhaseUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetPhaseUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetPhaseUpdateResponseEnvelopeJSON    `json:"-"`
}

// rulesetPhaseUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetPhaseUpdateResponseEnvelope]
type rulesetPhaseUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetPhaseUpdateResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetPhaseUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetPhaseUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetPhaseUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetPhaseUpdateResponseEnvelopeErrors]
type rulesetPhaseUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetPhaseUpdateResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                             `json:"pointer,required"`
	JSON    rulesetPhaseUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetPhaseUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RulesetPhaseUpdateResponseEnvelopeErrorsSource]
type rulesetPhaseUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetPhaseUpdateResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetPhaseUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetPhaseUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetPhaseUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RulesetPhaseUpdateResponseEnvelopeMessages]
type rulesetPhaseUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetPhaseUpdateResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                               `json:"pointer,required"`
	JSON    rulesetPhaseUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetPhaseUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RulesetPhaseUpdateResponseEnvelopeMessagesSource]
type rulesetPhaseUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetPhaseUpdateResponseEnvelopeSuccess bool

const (
	RulesetPhaseUpdateResponseEnvelopeSuccessTrue RulesetPhaseUpdateResponseEnvelopeSuccess = true
)

type RulesetPhaseGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// The phase of the ruleset.
type RulesetPhaseGetParamsRulesetPhase string

const (
	RulesetPhaseGetParamsRulesetPhaseDDOSL4                         RulesetPhaseGetParamsRulesetPhase = "ddos_l4"
	RulesetPhaseGetParamsRulesetPhaseDDOSL7                         RulesetPhaseGetParamsRulesetPhase = "ddos_l7"
	RulesetPhaseGetParamsRulesetPhaseHTTPConfigSettings             RulesetPhaseGetParamsRulesetPhase = "http_config_settings"
	RulesetPhaseGetParamsRulesetPhaseHTTPCustomErrors               RulesetPhaseGetParamsRulesetPhase = "http_custom_errors"
	RulesetPhaseGetParamsRulesetPhaseHTTPLogCustomFields            RulesetPhaseGetParamsRulesetPhase = "http_log_custom_fields"
	RulesetPhaseGetParamsRulesetPhaseHTTPRatelimit                  RulesetPhaseGetParamsRulesetPhase = "http_ratelimit"
	RulesetPhaseGetParamsRulesetPhaseHTTPRequestCacheSettings       RulesetPhaseGetParamsRulesetPhase = "http_request_cache_settings"
	RulesetPhaseGetParamsRulesetPhaseHTTPRequestDynamicRedirect     RulesetPhaseGetParamsRulesetPhase = "http_request_dynamic_redirect"
	RulesetPhaseGetParamsRulesetPhaseHTTPRequestFirewallCustom      RulesetPhaseGetParamsRulesetPhase = "http_request_firewall_custom"
	RulesetPhaseGetParamsRulesetPhaseHTTPRequestFirewallManaged     RulesetPhaseGetParamsRulesetPhase = "http_request_firewall_managed"
	RulesetPhaseGetParamsRulesetPhaseHTTPRequestLateTransform       RulesetPhaseGetParamsRulesetPhase = "http_request_late_transform"
	RulesetPhaseGetParamsRulesetPhaseHTTPRequestOrigin              RulesetPhaseGetParamsRulesetPhase = "http_request_origin"
	RulesetPhaseGetParamsRulesetPhaseHTTPRequestRedirect            RulesetPhaseGetParamsRulesetPhase = "http_request_redirect"
	RulesetPhaseGetParamsRulesetPhaseHTTPRequestSanitize            RulesetPhaseGetParamsRulesetPhase = "http_request_sanitize"
	RulesetPhaseGetParamsRulesetPhaseHTTPRequestSbfm                RulesetPhaseGetParamsRulesetPhase = "http_request_sbfm"
	RulesetPhaseGetParamsRulesetPhaseHTTPRequestSelectConfiguration RulesetPhaseGetParamsRulesetPhase = "http_request_select_configuration"
	RulesetPhaseGetParamsRulesetPhaseHTTPRequestTransform           RulesetPhaseGetParamsRulesetPhase = "http_request_transform"
	RulesetPhaseGetParamsRulesetPhaseHTTPResponseCompression        RulesetPhaseGetParamsRulesetPhase = "http_response_compression"
	RulesetPhaseGetParamsRulesetPhaseHTTPResponseFirewallManaged    RulesetPhaseGetParamsRulesetPhase = "http_response_firewall_managed"
	RulesetPhaseGetParamsRulesetPhaseHTTPResponseHeadersTransform   RulesetPhaseGetParamsRulesetPhase = "http_response_headers_transform"
	RulesetPhaseGetParamsRulesetPhaseMagicTransit                   RulesetPhaseGetParamsRulesetPhase = "magic_transit"
	RulesetPhaseGetParamsRulesetPhaseMagicTransitIDsManaged         RulesetPhaseGetParamsRulesetPhase = "magic_transit_ids_managed"
	RulesetPhaseGetParamsRulesetPhaseMagicTransitManaged            RulesetPhaseGetParamsRulesetPhase = "magic_transit_managed"
)

// A response object.
type RulesetPhaseGetResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetPhaseGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetPhaseGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetPhaseGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetPhaseGetResponseEnvelopeJSON    `json:"-"`
}

// rulesetPhaseGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetPhaseGetResponseEnvelope]
type rulesetPhaseGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetPhaseGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetPhaseGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetPhaseGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetPhaseGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetPhaseGetResponseEnvelopeErrors]
type rulesetPhaseGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetPhaseGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    rulesetPhaseGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetPhaseGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RulesetPhaseGetResponseEnvelopeErrorsSource]
type rulesetPhaseGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetPhaseGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetPhaseGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetPhaseGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetPhaseGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetPhaseGetResponseEnvelopeMessages]
type rulesetPhaseGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetPhaseGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    rulesetPhaseGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetPhaseGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RulesetPhaseGetResponseEnvelopeMessagesSource]
type rulesetPhaseGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetPhaseGetResponseEnvelopeSuccess bool

const (
	RulesetPhaseGetResponseEnvelopeSuccessTrue RulesetPhaseGetResponseEnvelopeSuccess = true
)
