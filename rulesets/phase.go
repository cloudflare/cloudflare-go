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

// PhaseService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPhaseService] method instead.
type PhaseService struct {
	Options  []option.RequestOption
	Versions *PhaseVersionService
}

// NewPhaseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPhaseService(opts ...option.RequestOption) (r *PhaseService) {
	r = &PhaseService{}
	r.Options = opts
	r.Versions = NewPhaseVersionService(opts...)
	return
}

// Updates an account or zone entry point ruleset, creating a new version.
func (r *PhaseService) Update(ctx context.Context, rulesetPhase PhaseUpdateParamsRulesetPhase, params PhaseUpdateParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PhaseUpdateResponseEnvelope
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
func (r *PhaseService) Get(ctx context.Context, rulesetPhase PhaseGetParamsRulesetPhase, query PhaseGetParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PhaseGetResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PhaseUpdateParams struct {
	// The unique ID of the ruleset.
	ID param.Field[string] `json:"id,required"`
	// The list of rules in the ruleset.
	Rules param.Field[[]PhaseUpdateParamsRule] `json:"rules,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[PhaseUpdateParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[PhaseUpdateParamsPhase] `json:"phase"`
}

func (r PhaseUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The phase of the ruleset.
type PhaseUpdateParamsRulesetPhase string

const (
	PhaseUpdateParamsRulesetPhaseDDOSL4                         PhaseUpdateParamsRulesetPhase = "ddos_l4"
	PhaseUpdateParamsRulesetPhaseDDOSL7                         PhaseUpdateParamsRulesetPhase = "ddos_l7"
	PhaseUpdateParamsRulesetPhaseHTTPConfigSettings             PhaseUpdateParamsRulesetPhase = "http_config_settings"
	PhaseUpdateParamsRulesetPhaseHTTPCustomErrors               PhaseUpdateParamsRulesetPhase = "http_custom_errors"
	PhaseUpdateParamsRulesetPhaseHTTPLogCustomFields            PhaseUpdateParamsRulesetPhase = "http_log_custom_fields"
	PhaseUpdateParamsRulesetPhaseHTTPRatelimit                  PhaseUpdateParamsRulesetPhase = "http_ratelimit"
	PhaseUpdateParamsRulesetPhaseHTTPRequestCacheSettings       PhaseUpdateParamsRulesetPhase = "http_request_cache_settings"
	PhaseUpdateParamsRulesetPhaseHTTPRequestDynamicRedirect     PhaseUpdateParamsRulesetPhase = "http_request_dynamic_redirect"
	PhaseUpdateParamsRulesetPhaseHTTPRequestFirewallCustom      PhaseUpdateParamsRulesetPhase = "http_request_firewall_custom"
	PhaseUpdateParamsRulesetPhaseHTTPRequestFirewallManaged     PhaseUpdateParamsRulesetPhase = "http_request_firewall_managed"
	PhaseUpdateParamsRulesetPhaseHTTPRequestLateTransform       PhaseUpdateParamsRulesetPhase = "http_request_late_transform"
	PhaseUpdateParamsRulesetPhaseHTTPRequestOrigin              PhaseUpdateParamsRulesetPhase = "http_request_origin"
	PhaseUpdateParamsRulesetPhaseHTTPRequestRedirect            PhaseUpdateParamsRulesetPhase = "http_request_redirect"
	PhaseUpdateParamsRulesetPhaseHTTPRequestSanitize            PhaseUpdateParamsRulesetPhase = "http_request_sanitize"
	PhaseUpdateParamsRulesetPhaseHTTPRequestSbfm                PhaseUpdateParamsRulesetPhase = "http_request_sbfm"
	PhaseUpdateParamsRulesetPhaseHTTPRequestSelectConfiguration PhaseUpdateParamsRulesetPhase = "http_request_select_configuration"
	PhaseUpdateParamsRulesetPhaseHTTPRequestTransform           PhaseUpdateParamsRulesetPhase = "http_request_transform"
	PhaseUpdateParamsRulesetPhaseHTTPResponseCompression        PhaseUpdateParamsRulesetPhase = "http_response_compression"
	PhaseUpdateParamsRulesetPhaseHTTPResponseFirewallManaged    PhaseUpdateParamsRulesetPhase = "http_response_firewall_managed"
	PhaseUpdateParamsRulesetPhaseHTTPResponseHeadersTransform   PhaseUpdateParamsRulesetPhase = "http_response_headers_transform"
	PhaseUpdateParamsRulesetPhaseMagicTransit                   PhaseUpdateParamsRulesetPhase = "magic_transit"
	PhaseUpdateParamsRulesetPhaseMagicTransitIDsManaged         PhaseUpdateParamsRulesetPhase = "magic_transit_ids_managed"
	PhaseUpdateParamsRulesetPhaseMagicTransitManaged            PhaseUpdateParamsRulesetPhase = "magic_transit_managed"
)

// Satisfied by [rulesets.PhaseUpdateParamsRulesRulesetsBlockRule],
// [rulesets.PhaseUpdateParamsRulesRulesetsExecuteRule],
// [rulesets.PhaseUpdateParamsRulesRulesetsLogRule],
// [rulesets.PhaseUpdateParamsRulesRulesetsSkipRule].
type PhaseUpdateParamsRule interface {
	implementsRulesetsPhaseUpdateParamsRule()
}

type PhaseUpdateParamsRulesRulesetsBlockRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsBlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsBlockRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[PhaseUpdateParamsRulesRulesetsBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsBlockRule) implementsRulesetsPhaseUpdateParamsRule() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsBlockRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsBlockRuleActionBlock PhaseUpdateParamsRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[PhaseUpdateParamsRulesRulesetsBlockRuleActionParametersResponse] `json:"response"`
}

func (r PhaseUpdateParamsRulesRulesetsBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type PhaseUpdateParamsRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r PhaseUpdateParamsRulesRulesetsBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type PhaseUpdateParamsRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r PhaseUpdateParamsRulesRulesetsBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhaseUpdateParamsRulesRulesetsExecuteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsExecuteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[PhaseUpdateParamsRulesRulesetsExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsExecuteRule) implementsRulesetsPhaseUpdateParamsRule() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsExecuteRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionExecute PhaseUpdateParamsRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r PhaseUpdateParamsRulesRulesetsExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRule] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for rules in the category.
type PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for the rule.
type PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type PhaseUpdateParamsRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r PhaseUpdateParamsRulesRulesetsExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhaseUpdateParamsRulesRulesetsLogRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsLogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[PhaseUpdateParamsRulesRulesetsLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsLogRule) implementsRulesetsPhaseUpdateParamsRule() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsLogRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsLogRuleActionLog PhaseUpdateParamsRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type PhaseUpdateParamsRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r PhaseUpdateParamsRulesRulesetsLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhaseUpdateParamsRulesRulesetsSkipRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsSkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsSkipRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[PhaseUpdateParamsRulesRulesetsSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSkipRule) implementsRulesetsPhaseUpdateParamsRule() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsSkipRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsSkipRuleActionSkip PhaseUpdateParamsRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r PhaseUpdateParamsRulesRulesetsSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A phase to skip the execution of.
type PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase string

const (
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct string

const (
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductBic           PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "bic"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductHot           PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "hot"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductRateLimit     PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductSecurityLevel PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductUABlock       PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductWAF           PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "waf"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductZoneLockdown  PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset string

const (
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRulesetCurrent PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type PhaseUpdateParamsRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r PhaseUpdateParamsRulesRulesetsSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type PhaseUpdateParamsKind string

const (
	PhaseUpdateParamsKindManaged PhaseUpdateParamsKind = "managed"
	PhaseUpdateParamsKindCustom  PhaseUpdateParamsKind = "custom"
	PhaseUpdateParamsKindRoot    PhaseUpdateParamsKind = "root"
	PhaseUpdateParamsKindZone    PhaseUpdateParamsKind = "zone"
)

// The phase of the ruleset.
type PhaseUpdateParamsPhase string

const (
	PhaseUpdateParamsPhaseDDOSL4                         PhaseUpdateParamsPhase = "ddos_l4"
	PhaseUpdateParamsPhaseDDOSL7                         PhaseUpdateParamsPhase = "ddos_l7"
	PhaseUpdateParamsPhaseHTTPConfigSettings             PhaseUpdateParamsPhase = "http_config_settings"
	PhaseUpdateParamsPhaseHTTPCustomErrors               PhaseUpdateParamsPhase = "http_custom_errors"
	PhaseUpdateParamsPhaseHTTPLogCustomFields            PhaseUpdateParamsPhase = "http_log_custom_fields"
	PhaseUpdateParamsPhaseHTTPRatelimit                  PhaseUpdateParamsPhase = "http_ratelimit"
	PhaseUpdateParamsPhaseHTTPRequestCacheSettings       PhaseUpdateParamsPhase = "http_request_cache_settings"
	PhaseUpdateParamsPhaseHTTPRequestDynamicRedirect     PhaseUpdateParamsPhase = "http_request_dynamic_redirect"
	PhaseUpdateParamsPhaseHTTPRequestFirewallCustom      PhaseUpdateParamsPhase = "http_request_firewall_custom"
	PhaseUpdateParamsPhaseHTTPRequestFirewallManaged     PhaseUpdateParamsPhase = "http_request_firewall_managed"
	PhaseUpdateParamsPhaseHTTPRequestLateTransform       PhaseUpdateParamsPhase = "http_request_late_transform"
	PhaseUpdateParamsPhaseHTTPRequestOrigin              PhaseUpdateParamsPhase = "http_request_origin"
	PhaseUpdateParamsPhaseHTTPRequestRedirect            PhaseUpdateParamsPhase = "http_request_redirect"
	PhaseUpdateParamsPhaseHTTPRequestSanitize            PhaseUpdateParamsPhase = "http_request_sanitize"
	PhaseUpdateParamsPhaseHTTPRequestSbfm                PhaseUpdateParamsPhase = "http_request_sbfm"
	PhaseUpdateParamsPhaseHTTPRequestSelectConfiguration PhaseUpdateParamsPhase = "http_request_select_configuration"
	PhaseUpdateParamsPhaseHTTPRequestTransform           PhaseUpdateParamsPhase = "http_request_transform"
	PhaseUpdateParamsPhaseHTTPResponseCompression        PhaseUpdateParamsPhase = "http_response_compression"
	PhaseUpdateParamsPhaseHTTPResponseFirewallManaged    PhaseUpdateParamsPhase = "http_response_firewall_managed"
	PhaseUpdateParamsPhaseHTTPResponseHeadersTransform   PhaseUpdateParamsPhase = "http_response_headers_transform"
	PhaseUpdateParamsPhaseMagicTransit                   PhaseUpdateParamsPhase = "magic_transit"
	PhaseUpdateParamsPhaseMagicTransitIDsManaged         PhaseUpdateParamsPhase = "magic_transit_ids_managed"
	PhaseUpdateParamsPhaseMagicTransitManaged            PhaseUpdateParamsPhase = "magic_transit_managed"
)

// A response object.
type PhaseUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success PhaseUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    phaseUpdateResponseEnvelopeJSON    `json:"-"`
}

// phaseUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelope]
type phaseUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseUpdateResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON   phaseUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// phaseUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeErrors]
type phaseUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseUpdateResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                      `json:"pointer,required"`
	JSON    phaseUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseEnvelopeErrorsSource]
type phaseUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseUpdateResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON   phaseUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// phaseUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseEnvelopeMessages]
type phaseUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseUpdateResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                        `json:"pointer,required"`
	JSON    phaseUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseEnvelopeMessagesSource]
type phaseUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PhaseUpdateResponseEnvelopeSuccess bool

const (
	PhaseUpdateResponseEnvelopeSuccessTrue PhaseUpdateResponseEnvelopeSuccess = true
)

type PhaseGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// The phase of the ruleset.
type PhaseGetParamsRulesetPhase string

const (
	PhaseGetParamsRulesetPhaseDDOSL4                         PhaseGetParamsRulesetPhase = "ddos_l4"
	PhaseGetParamsRulesetPhaseDDOSL7                         PhaseGetParamsRulesetPhase = "ddos_l7"
	PhaseGetParamsRulesetPhaseHTTPConfigSettings             PhaseGetParamsRulesetPhase = "http_config_settings"
	PhaseGetParamsRulesetPhaseHTTPCustomErrors               PhaseGetParamsRulesetPhase = "http_custom_errors"
	PhaseGetParamsRulesetPhaseHTTPLogCustomFields            PhaseGetParamsRulesetPhase = "http_log_custom_fields"
	PhaseGetParamsRulesetPhaseHTTPRatelimit                  PhaseGetParamsRulesetPhase = "http_ratelimit"
	PhaseGetParamsRulesetPhaseHTTPRequestCacheSettings       PhaseGetParamsRulesetPhase = "http_request_cache_settings"
	PhaseGetParamsRulesetPhaseHTTPRequestDynamicRedirect     PhaseGetParamsRulesetPhase = "http_request_dynamic_redirect"
	PhaseGetParamsRulesetPhaseHTTPRequestFirewallCustom      PhaseGetParamsRulesetPhase = "http_request_firewall_custom"
	PhaseGetParamsRulesetPhaseHTTPRequestFirewallManaged     PhaseGetParamsRulesetPhase = "http_request_firewall_managed"
	PhaseGetParamsRulesetPhaseHTTPRequestLateTransform       PhaseGetParamsRulesetPhase = "http_request_late_transform"
	PhaseGetParamsRulesetPhaseHTTPRequestOrigin              PhaseGetParamsRulesetPhase = "http_request_origin"
	PhaseGetParamsRulesetPhaseHTTPRequestRedirect            PhaseGetParamsRulesetPhase = "http_request_redirect"
	PhaseGetParamsRulesetPhaseHTTPRequestSanitize            PhaseGetParamsRulesetPhase = "http_request_sanitize"
	PhaseGetParamsRulesetPhaseHTTPRequestSbfm                PhaseGetParamsRulesetPhase = "http_request_sbfm"
	PhaseGetParamsRulesetPhaseHTTPRequestSelectConfiguration PhaseGetParamsRulesetPhase = "http_request_select_configuration"
	PhaseGetParamsRulesetPhaseHTTPRequestTransform           PhaseGetParamsRulesetPhase = "http_request_transform"
	PhaseGetParamsRulesetPhaseHTTPResponseCompression        PhaseGetParamsRulesetPhase = "http_response_compression"
	PhaseGetParamsRulesetPhaseHTTPResponseFirewallManaged    PhaseGetParamsRulesetPhase = "http_response_firewall_managed"
	PhaseGetParamsRulesetPhaseHTTPResponseHeadersTransform   PhaseGetParamsRulesetPhase = "http_response_headers_transform"
	PhaseGetParamsRulesetPhaseMagicTransit                   PhaseGetParamsRulesetPhase = "magic_transit"
	PhaseGetParamsRulesetPhaseMagicTransitIDsManaged         PhaseGetParamsRulesetPhase = "magic_transit_ids_managed"
	PhaseGetParamsRulesetPhaseMagicTransitManaged            PhaseGetParamsRulesetPhase = "magic_transit_managed"
)

// A response object.
type PhaseGetResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success PhaseGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    phaseGetResponseEnvelopeJSON    `json:"-"`
}

// phaseGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PhaseGetResponseEnvelope]
type phaseGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   phaseGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// phaseGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PhaseGetResponseEnvelopeErrors]
type phaseGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                   `json:"pointer,required"`
	JSON    phaseGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// phaseGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [PhaseGetResponseEnvelopeErrorsSource]
type phaseGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   phaseGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// phaseGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PhaseGetResponseEnvelopeMessages]
type phaseGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    phaseGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// phaseGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [PhaseGetResponseEnvelopeMessagesSource]
type phaseGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PhaseGetResponseEnvelopeSuccess bool

const (
	PhaseGetResponseEnvelopeSuccessTrue PhaseGetResponseEnvelopeSuccess = true
)
