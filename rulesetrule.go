// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// RulesetRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRulesetRuleService] method
// instead.
type RulesetRuleService struct {
	Options []option.RequestOption
}

// NewRulesetRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRulesetRuleService(opts ...option.RequestOption) (r *RulesetRuleService) {
	r = &RulesetRuleService{}
	r.Options = opts
	return
}

// Updates an existing rule in an account ruleset.
func (r *RulesetRuleService) Update(ctx context.Context, rulesetID string, ruleID string, params RulesetRuleUpdateParams, opts ...option.RequestOption) (res *RulesetRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetRuleUpdateResponseEnvelope
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/:rule_id", params.AccountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing rule from an account or zone ruleset.
func (r *RulesetRuleService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, rulesetID string, ruleID string, opts ...option.RequestOption) (res *RulesetRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetRuleDeleteResponseEnvelope
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/%s", accountOrZone, accountOrZoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Adds a new rule to an account or zone ruleset. The rule will be added to the end
// of the existing list of rules in the ruleset by default.
func (r *RulesetRuleService) AccountRulesetsNewAnAccountRulesetRule(ctx context.Context, accountOrZone string, accountOrZoneID string, rulesetID string, body RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams, opts ...option.RequestOption) (res *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelope
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A result.
type RulesetRuleUpdateResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetRuleUpdateResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetRuleUpdateResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetRuleUpdateResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                        `json:"description"`
	JSON        rulesetRuleUpdateResponseJSON `json:"-"`
}

// rulesetRuleUpdateResponseJSON contains the JSON metadata for the struct
// [RulesetRuleUpdateResponse]
type rulesetRuleUpdateResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetRuleUpdateResponseKind string

const (
	RulesetRuleUpdateResponseKindManaged RulesetRuleUpdateResponseKind = "managed"
	RulesetRuleUpdateResponseKindCustom  RulesetRuleUpdateResponseKind = "custom"
	RulesetRuleUpdateResponseKindRoot    RulesetRuleUpdateResponseKind = "root"
	RulesetRuleUpdateResponseKindZone    RulesetRuleUpdateResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetRuleUpdateResponsePhase string

const (
	RulesetRuleUpdateResponsePhaseDDOSL4                         RulesetRuleUpdateResponsePhase = "ddos_l4"
	RulesetRuleUpdateResponsePhaseDDOSL7                         RulesetRuleUpdateResponsePhase = "ddos_l7"
	RulesetRuleUpdateResponsePhaseHTTPConfigSettings             RulesetRuleUpdateResponsePhase = "http_config_settings"
	RulesetRuleUpdateResponsePhaseHTTPCustomErrors               RulesetRuleUpdateResponsePhase = "http_custom_errors"
	RulesetRuleUpdateResponsePhaseHTTPLogCustomFields            RulesetRuleUpdateResponsePhase = "http_log_custom_fields"
	RulesetRuleUpdateResponsePhaseHTTPRatelimit                  RulesetRuleUpdateResponsePhase = "http_ratelimit"
	RulesetRuleUpdateResponsePhaseHTTPRequestCacheSettings       RulesetRuleUpdateResponsePhase = "http_request_cache_settings"
	RulesetRuleUpdateResponsePhaseHTTPRequestDynamicRedirect     RulesetRuleUpdateResponsePhase = "http_request_dynamic_redirect"
	RulesetRuleUpdateResponsePhaseHTTPRequestFirewallCustom      RulesetRuleUpdateResponsePhase = "http_request_firewall_custom"
	RulesetRuleUpdateResponsePhaseHTTPRequestFirewallManaged     RulesetRuleUpdateResponsePhase = "http_request_firewall_managed"
	RulesetRuleUpdateResponsePhaseHTTPRequestLateTransform       RulesetRuleUpdateResponsePhase = "http_request_late_transform"
	RulesetRuleUpdateResponsePhaseHTTPRequestOrigin              RulesetRuleUpdateResponsePhase = "http_request_origin"
	RulesetRuleUpdateResponsePhaseHTTPRequestRedirect            RulesetRuleUpdateResponsePhase = "http_request_redirect"
	RulesetRuleUpdateResponsePhaseHTTPRequestSanitize            RulesetRuleUpdateResponsePhase = "http_request_sanitize"
	RulesetRuleUpdateResponsePhaseHTTPRequestSbfm                RulesetRuleUpdateResponsePhase = "http_request_sbfm"
	RulesetRuleUpdateResponsePhaseHTTPRequestSelectConfiguration RulesetRuleUpdateResponsePhase = "http_request_select_configuration"
	RulesetRuleUpdateResponsePhaseHTTPRequestTransform           RulesetRuleUpdateResponsePhase = "http_request_transform"
	RulesetRuleUpdateResponsePhaseHTTPResponseCompression        RulesetRuleUpdateResponsePhase = "http_response_compression"
	RulesetRuleUpdateResponsePhaseHTTPResponseFirewallManaged    RulesetRuleUpdateResponsePhase = "http_response_firewall_managed"
	RulesetRuleUpdateResponsePhaseHTTPResponseHeadersTransform   RulesetRuleUpdateResponsePhase = "http_response_headers_transform"
	RulesetRuleUpdateResponsePhaseMagicTransit                   RulesetRuleUpdateResponsePhase = "magic_transit"
	RulesetRuleUpdateResponsePhaseMagicTransitIDsManaged         RulesetRuleUpdateResponsePhase = "magic_transit_ids_managed"
	RulesetRuleUpdateResponsePhaseMagicTransitManaged            RulesetRuleUpdateResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetRuleUpdateResponseRulesRulesetsBlockRule],
// [RulesetRuleUpdateResponseRulesRulesetsExecuteRule],
// [RulesetRuleUpdateResponseRulesRulesetsLogRule] or
// [RulesetRuleUpdateResponseRulesRulesetsSkipRule].
type RulesetRuleUpdateResponseRule interface {
	implementsRulesetRuleUpdateResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetRuleUpdateResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleUpdateResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleUpdateResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleUpdateResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleUpdateResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetRuleUpdateResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleUpdateResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleUpdateResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                              `json:"ref"`
	JSON rulesetRuleUpdateResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsBlockRuleJSON contains the JSON metadata
// for the struct [RulesetRuleUpdateResponseRulesRulesetsBlockRule]
type rulesetRuleUpdateResponseRulesRulesetsBlockRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleUpdateResponseRulesRulesetsBlockRule) implementsRulesetRuleUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleUpdateResponseRulesRulesetsBlockRuleAction string

const (
	RulesetRuleUpdateResponseRulesRulesetsBlockRuleActionBlock RulesetRuleUpdateResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParameters]
type rulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response to show when the block is applied.
type RulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                       `json:"status_code,required"`
	JSON       rulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [RulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object configuring the rule's logging behavior.
type RulesetRuleUpdateResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                       `json:"enabled,required"`
	JSON    rulesetRuleUpdateResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleUpdateResponseRulesRulesetsBlockRuleLogging]
type rulesetRuleUpdateResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRuleUpdateResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleUpdateResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleUpdateResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                `json:"ref"`
	JSON rulesetRuleUpdateResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata
// for the struct [RulesetRuleUpdateResponseRulesRulesetsExecuteRule]
type rulesetRuleUpdateResponseRulesRulesetsExecuteRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleUpdateResponseRulesRulesetsExecuteRule) implementsRulesetRuleUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetRuleUpdateResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionExecute RulesetRuleUpdateResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration to use for matched data logging.
type RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                           `json:"public_key,required"`
	JSON      rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A set of overrides to apply to the target ruleset.
type RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A category-level override
type RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for rules in the category.
type RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for the rule.
type RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetRuleUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetRuleUpdateResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                         `json:"enabled,required"`
	JSON    rulesetRuleUpdateResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON
// metadata for the struct
// [RulesetRuleUpdateResponseRulesRulesetsExecuteRuleLogging]
type rulesetRuleUpdateResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRuleUpdateResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleUpdateResponseRulesRulesetsLogRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleUpdateResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                            `json:"ref"`
	JSON rulesetRuleUpdateResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsLogRuleJSON contains the JSON metadata for
// the struct [RulesetRuleUpdateResponseRulesRulesetsLogRule]
type rulesetRuleUpdateResponseRulesRulesetsLogRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleUpdateResponseRulesRulesetsLogRule) implementsRulesetRuleUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleUpdateResponseRulesRulesetsLogRuleAction string

const (
	RulesetRuleUpdateResponseRulesRulesetsLogRuleActionLog RulesetRuleUpdateResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetRuleUpdateResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                     `json:"enabled,required"`
	JSON    rulesetRuleUpdateResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsLogRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleUpdateResponseRulesRulesetsLogRuleLogging]
type rulesetRuleUpdateResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRuleUpdateResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleUpdateResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleUpdateResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                             `json:"ref"`
	JSON rulesetRuleUpdateResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsSkipRuleJSON contains the JSON metadata
// for the struct [RulesetRuleUpdateResponseRulesRulesetsSkipRule]
type rulesetRuleUpdateResponseRulesRulesetsSkipRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleUpdateResponseRulesRulesetsSkipRule) implementsRulesetRuleUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleUpdateResponseRulesRulesetsSkipRuleAction string

const (
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionSkip RulesetRuleUpdateResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                           `json:"rulesets"`
	JSON     rulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParameters]
type rulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A phase to skip the execution of.
type RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProductUaBlock       RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetRuleUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetRuleUpdateResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                      `json:"enabled,required"`
	JSON    rulesetRuleUpdateResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetRuleUpdateResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleUpdateResponseRulesRulesetsSkipRuleLogging]
type rulesetRuleUpdateResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A result.
type RulesetRuleDeleteResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetRuleDeleteResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetRuleDeleteResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetRuleDeleteResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                        `json:"description"`
	JSON        rulesetRuleDeleteResponseJSON `json:"-"`
}

// rulesetRuleDeleteResponseJSON contains the JSON metadata for the struct
// [RulesetRuleDeleteResponse]
type rulesetRuleDeleteResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetRuleDeleteResponseKind string

const (
	RulesetRuleDeleteResponseKindManaged RulesetRuleDeleteResponseKind = "managed"
	RulesetRuleDeleteResponseKindCustom  RulesetRuleDeleteResponseKind = "custom"
	RulesetRuleDeleteResponseKindRoot    RulesetRuleDeleteResponseKind = "root"
	RulesetRuleDeleteResponseKindZone    RulesetRuleDeleteResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetRuleDeleteResponsePhase string

const (
	RulesetRuleDeleteResponsePhaseDDOSL4                         RulesetRuleDeleteResponsePhase = "ddos_l4"
	RulesetRuleDeleteResponsePhaseDDOSL7                         RulesetRuleDeleteResponsePhase = "ddos_l7"
	RulesetRuleDeleteResponsePhaseHTTPConfigSettings             RulesetRuleDeleteResponsePhase = "http_config_settings"
	RulesetRuleDeleteResponsePhaseHTTPCustomErrors               RulesetRuleDeleteResponsePhase = "http_custom_errors"
	RulesetRuleDeleteResponsePhaseHTTPLogCustomFields            RulesetRuleDeleteResponsePhase = "http_log_custom_fields"
	RulesetRuleDeleteResponsePhaseHTTPRatelimit                  RulesetRuleDeleteResponsePhase = "http_ratelimit"
	RulesetRuleDeleteResponsePhaseHTTPRequestCacheSettings       RulesetRuleDeleteResponsePhase = "http_request_cache_settings"
	RulesetRuleDeleteResponsePhaseHTTPRequestDynamicRedirect     RulesetRuleDeleteResponsePhase = "http_request_dynamic_redirect"
	RulesetRuleDeleteResponsePhaseHTTPRequestFirewallCustom      RulesetRuleDeleteResponsePhase = "http_request_firewall_custom"
	RulesetRuleDeleteResponsePhaseHTTPRequestFirewallManaged     RulesetRuleDeleteResponsePhase = "http_request_firewall_managed"
	RulesetRuleDeleteResponsePhaseHTTPRequestLateTransform       RulesetRuleDeleteResponsePhase = "http_request_late_transform"
	RulesetRuleDeleteResponsePhaseHTTPRequestOrigin              RulesetRuleDeleteResponsePhase = "http_request_origin"
	RulesetRuleDeleteResponsePhaseHTTPRequestRedirect            RulesetRuleDeleteResponsePhase = "http_request_redirect"
	RulesetRuleDeleteResponsePhaseHTTPRequestSanitize            RulesetRuleDeleteResponsePhase = "http_request_sanitize"
	RulesetRuleDeleteResponsePhaseHTTPRequestSbfm                RulesetRuleDeleteResponsePhase = "http_request_sbfm"
	RulesetRuleDeleteResponsePhaseHTTPRequestSelectConfiguration RulesetRuleDeleteResponsePhase = "http_request_select_configuration"
	RulesetRuleDeleteResponsePhaseHTTPRequestTransform           RulesetRuleDeleteResponsePhase = "http_request_transform"
	RulesetRuleDeleteResponsePhaseHTTPResponseCompression        RulesetRuleDeleteResponsePhase = "http_response_compression"
	RulesetRuleDeleteResponsePhaseHTTPResponseFirewallManaged    RulesetRuleDeleteResponsePhase = "http_response_firewall_managed"
	RulesetRuleDeleteResponsePhaseHTTPResponseHeadersTransform   RulesetRuleDeleteResponsePhase = "http_response_headers_transform"
	RulesetRuleDeleteResponsePhaseMagicTransit                   RulesetRuleDeleteResponsePhase = "magic_transit"
	RulesetRuleDeleteResponsePhaseMagicTransitIDsManaged         RulesetRuleDeleteResponsePhase = "magic_transit_ids_managed"
	RulesetRuleDeleteResponsePhaseMagicTransitManaged            RulesetRuleDeleteResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetRuleDeleteResponseRulesRulesetsBlockRule],
// [RulesetRuleDeleteResponseRulesRulesetsExecuteRule],
// [RulesetRuleDeleteResponseRulesRulesetsLogRule] or
// [RulesetRuleDeleteResponseRulesRulesetsSkipRule].
type RulesetRuleDeleteResponseRule interface {
	implementsRulesetRuleDeleteResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetRuleDeleteResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleDeleteResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleDeleteResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleDeleteResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleDeleteResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetRuleDeleteResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleDeleteResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleDeleteResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                              `json:"ref"`
	JSON rulesetRuleDeleteResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsBlockRuleJSON contains the JSON metadata
// for the struct [RulesetRuleDeleteResponseRulesRulesetsBlockRule]
type rulesetRuleDeleteResponseRulesRulesetsBlockRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleDeleteResponseRulesRulesetsBlockRule) implementsRulesetRuleDeleteResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleDeleteResponseRulesRulesetsBlockRuleAction string

const (
	RulesetRuleDeleteResponseRulesRulesetsBlockRuleActionBlock RulesetRuleDeleteResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParameters]
type rulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response to show when the block is applied.
type RulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                       `json:"status_code,required"`
	JSON       rulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [RulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object configuring the rule's logging behavior.
type RulesetRuleDeleteResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                       `json:"enabled,required"`
	JSON    rulesetRuleDeleteResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleDeleteResponseRulesRulesetsBlockRuleLogging]
type rulesetRuleDeleteResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRuleDeleteResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleDeleteResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleDeleteResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                `json:"ref"`
	JSON rulesetRuleDeleteResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata
// for the struct [RulesetRuleDeleteResponseRulesRulesetsExecuteRule]
type rulesetRuleDeleteResponseRulesRulesetsExecuteRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleDeleteResponseRulesRulesetsExecuteRule) implementsRulesetRuleDeleteResponseRule() {
}

// The action to perform when the rule matches.
type RulesetRuleDeleteResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionExecute RulesetRuleDeleteResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration to use for matched data logging.
type RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                           `json:"public_key,required"`
	JSON      rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A set of overrides to apply to the target ruleset.
type RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A category-level override
type RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for rules in the category.
type RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for the rule.
type RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetRuleDeleteResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                         `json:"enabled,required"`
	JSON    rulesetRuleDeleteResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON
// metadata for the struct
// [RulesetRuleDeleteResponseRulesRulesetsExecuteRuleLogging]
type rulesetRuleDeleteResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRuleDeleteResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleDeleteResponseRulesRulesetsLogRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleDeleteResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                            `json:"ref"`
	JSON rulesetRuleDeleteResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsLogRuleJSON contains the JSON metadata for
// the struct [RulesetRuleDeleteResponseRulesRulesetsLogRule]
type rulesetRuleDeleteResponseRulesRulesetsLogRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleDeleteResponseRulesRulesetsLogRule) implementsRulesetRuleDeleteResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleDeleteResponseRulesRulesetsLogRuleAction string

const (
	RulesetRuleDeleteResponseRulesRulesetsLogRuleActionLog RulesetRuleDeleteResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetRuleDeleteResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                     `json:"enabled,required"`
	JSON    rulesetRuleDeleteResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsLogRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleDeleteResponseRulesRulesetsLogRuleLogging]
type rulesetRuleDeleteResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRuleDeleteResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleDeleteResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleDeleteResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                             `json:"ref"`
	JSON rulesetRuleDeleteResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsSkipRuleJSON contains the JSON metadata
// for the struct [RulesetRuleDeleteResponseRulesRulesetsSkipRule]
type rulesetRuleDeleteResponseRulesRulesetsSkipRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleDeleteResponseRulesRulesetsSkipRule) implementsRulesetRuleDeleteResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleDeleteResponseRulesRulesetsSkipRuleAction string

const (
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionSkip RulesetRuleDeleteResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                           `json:"rulesets"`
	JSON     rulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParameters]
type rulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A phase to skip the execution of.
type RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductUaBlock       RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetRuleDeleteResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                      `json:"enabled,required"`
	JSON    rulesetRuleDeleteResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetRuleDeleteResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleDeleteResponseRulesRulesetsSkipRuleLogging]
type rulesetRuleDeleteResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A result.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                                                        `json:"description"`
	JSON        rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseJSON contains the JSON
// metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponse]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseKind string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseKindManaged RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseKind = "managed"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseKindCustom  RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseKind = "custom"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseKindRoot    RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseKind = "root"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseKindZone    RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseDDOSL4                         RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "ddos_l4"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseDDOSL7                         RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "ddos_l7"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPConfigSettings             RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_config_settings"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPCustomErrors               RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_custom_errors"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPLogCustomFields            RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_log_custom_fields"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPRatelimit                  RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_ratelimit"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPRequestCacheSettings       RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_request_cache_settings"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPRequestDynamicRedirect     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_request_dynamic_redirect"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPRequestFirewallCustom      RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_request_firewall_custom"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPRequestFirewallManaged     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_request_firewall_managed"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPRequestLateTransform       RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_request_late_transform"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPRequestOrigin              RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_request_origin"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPRequestRedirect            RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_request_redirect"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPRequestSanitize            RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_request_sanitize"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPRequestSbfm                RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_request_sbfm"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPRequestSelectConfiguration RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_request_select_configuration"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPRequestTransform           RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_request_transform"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPResponseCompression        RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_response_compression"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPResponseFirewallManaged    RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_response_firewall_managed"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseHTTPResponseHeadersTransform   RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "http_response_headers_transform"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseMagicTransit                   RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "magic_transit"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseMagicTransitIDsManaged         RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "magic_transit_ids_managed"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhaseMagicTransitManaged            RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponsePhase = "magic_transit_managed"
)

// Union satisfied by
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRule],
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRule],
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRule]
// or
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRule].
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRule interface {
	implementsRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                                              `json:"ref"`
	JSON rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRule]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRule) implementsRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRule() {
}

// The action to perform when the rule matches.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleAction string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionBlock RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParameters]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response to show when the block is applied.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                                                       `json:"status_code,required"`
	JSON       rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object configuring the rule's logging behavior.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                                                       `json:"enabled,required"`
	JSON    rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleLoggingJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleLogging]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                                                `json:"ref"`
	JSON rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRule]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRule) implementsRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRule() {
}

// The action to perform when the rule matches.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionExecute RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration to use for matched data logging.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                                                           `json:"public_key,required"`
	JSON      rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A set of overrides to apply to the target ruleset.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A category-level override
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for rules in the category.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for the rule.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                                                         `json:"enabled,required"`
	JSON    rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleLoggingJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleLogging]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                                            `json:"ref"`
	JSON rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRule]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRule) implementsRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRule() {
}

// The action to perform when the rule matches.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleAction string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleActionLog RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                                                     `json:"enabled,required"`
	JSON    rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleLoggingJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleLogging]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                                             `json:"ref"`
	JSON rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRule]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRule) implementsRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRule() {
}

// The action to perform when the rule matches.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleAction string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionSkip RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                                                           `json:"rulesets"`
	JSON     rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParameters]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A phase to skip the execution of.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProductUaBlock       RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                                                      `json:"enabled,required"`
	JSON    rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleLoggingJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleLogging]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This interface is a union satisfied by one of the following:
// [RulesetRuleUpdateParamsRulesetsBlockRule],
// [RulesetRuleUpdateParamsRulesetsExecuteRule],
// [RulesetRuleUpdateParamsRulesetsLogRule],
// [RulesetRuleUpdateParamsRulesetsSkipRule].
type RulesetRuleUpdateParams interface {
	ImplementsRulesetRuleUpdateParams()
}

type RulesetRuleUpdateParamsRulesetsBlockRule struct {
	// The unique ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetRuleUpdateParamsRulesetsBlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetRuleUpdateParamsRulesetsBlockRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetRuleUpdateParamsRulesetsBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetRuleUpdateParamsRulesetsBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (RulesetRuleUpdateParamsRulesetsBlockRule) ImplementsRulesetRuleUpdateParams() {

}

// The action to perform when the rule matches.
type RulesetRuleUpdateParamsRulesetsBlockRuleAction string

const (
	RulesetRuleUpdateParamsRulesetsBlockRuleActionBlock RulesetRuleUpdateParamsRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetRuleUpdateParamsRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[RulesetRuleUpdateParamsRulesetsBlockRuleActionParametersResponse] `json:"response"`
}

func (r RulesetRuleUpdateParamsRulesetsBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type RulesetRuleUpdateParamsRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r RulesetRuleUpdateParamsRulesetsBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type RulesetRuleUpdateParamsRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetRuleUpdateParamsRulesetsBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetRuleUpdateParamsRulesetsExecuteRule struct {
	// The unique ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetRuleUpdateParamsRulesetsExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetRuleUpdateParamsRulesetsExecuteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetRuleUpdateParamsRulesetsExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetRuleUpdateParamsRulesetsExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (RulesetRuleUpdateParamsRulesetsExecuteRule) ImplementsRulesetRuleUpdateParams() {

}

// The action to perform when the rule matches.
type RulesetRuleUpdateParamsRulesetsExecuteRuleAction string

const (
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionExecute RulesetRuleUpdateParamsRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetRuleUpdateParamsRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r RulesetRuleUpdateParamsRulesetsExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRule] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for rules in the category.
type RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for the rule.
type RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetRuleUpdateParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetRuleUpdateParamsRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetRuleUpdateParamsRulesetsExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetRuleUpdateParamsRulesetsLogRule struct {
	// The unique ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetRuleUpdateParamsRulesetsLogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetRuleUpdateParamsRulesetsLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetRuleUpdateParamsRulesetsLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (RulesetRuleUpdateParamsRulesetsLogRule) ImplementsRulesetRuleUpdateParams() {

}

// The action to perform when the rule matches.
type RulesetRuleUpdateParamsRulesetsLogRuleAction string

const (
	RulesetRuleUpdateParamsRulesetsLogRuleActionLog RulesetRuleUpdateParamsRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetRuleUpdateParamsRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetRuleUpdateParamsRulesetsLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetRuleUpdateParamsRulesetsSkipRule struct {
	// The unique ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetRuleUpdateParamsRulesetsSkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetRuleUpdateParamsRulesetsSkipRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetRuleUpdateParamsRulesetsSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetRuleUpdateParamsRulesetsSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (RulesetRuleUpdateParamsRulesetsSkipRule) ImplementsRulesetRuleUpdateParams() {

}

// The action to perform when the rule matches.
type RulesetRuleUpdateParamsRulesetsSkipRuleAction string

const (
	RulesetRuleUpdateParamsRulesetsSkipRuleActionSkip RulesetRuleUpdateParamsRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetRuleUpdateParamsRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r RulesetRuleUpdateParamsRulesetsSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A phase to skip the execution of.
type RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase string

const (
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProduct string

const (
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProductBic           RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProductHot           RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProductRateLimit     RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProductSecurityLevel RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProductUaBlock       RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProductWAF           RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersRulesetCurrent RulesetRuleUpdateParamsRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetRuleUpdateParamsRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetRuleUpdateParamsRulesetsSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A response object.
type RulesetRuleUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetRuleUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetRuleUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// rulesetRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetRuleUpdateResponseEnvelope]
type rulesetRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleUpdateResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetRuleUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetRuleUpdateResponseEnvelopeErrors]
type rulesetRuleUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleUpdateResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    rulesetRuleUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetRuleUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RulesetRuleUpdateResponseEnvelopeErrorsSource]
type rulesetRuleUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleUpdateResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetRuleUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetRuleUpdateResponseEnvelopeMessages]
type rulesetRuleUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleUpdateResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                              `json:"pointer,required"`
	JSON    rulesetRuleUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetRuleUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RulesetRuleUpdateResponseEnvelopeMessagesSource]
type rulesetRuleUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetRuleUpdateResponseEnvelopeSuccess bool

const (
	RulesetRuleUpdateResponseEnvelopeSuccessTrue RulesetRuleUpdateResponseEnvelopeSuccess = true
)

// A response object.
type RulesetRuleDeleteResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetRuleDeleteResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetRuleDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// rulesetRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetRuleDeleteResponseEnvelope]
type rulesetRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleDeleteResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetRuleDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetRuleDeleteResponseEnvelopeErrors]
type rulesetRuleDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleDeleteResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    rulesetRuleDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetRuleDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RulesetRuleDeleteResponseEnvelopeErrorsSource]
type rulesetRuleDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleDeleteResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetRuleDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetRuleDeleteResponseEnvelopeMessages]
type rulesetRuleDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleDeleteResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                              `json:"pointer,required"`
	JSON    rulesetRuleDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetRuleDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RulesetRuleDeleteResponseEnvelopeMessagesSource]
type rulesetRuleDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetRuleDeleteResponseEnvelopeSuccess bool

const (
	RulesetRuleDeleteResponseEnvelopeSuccessTrue RulesetRuleDeleteResponseEnvelopeSuccess = true
)

// This interface is a union satisfied by one of the following:
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRule],
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRule],
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsLogRule],
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRule].
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams interface {
	ImplementsRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams()
}

type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRule) ImplementsRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams() {

}

// The action to perform when the rule matches.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleAction string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleActionBlock RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleActionParametersResponse] `json:"response"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRule) ImplementsRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams() {

}

// The action to perform when the rule matches.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleAction string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionExecute RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRule] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for rules in the category.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for the rule.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsLogRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsLogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsLogRule) ImplementsRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams() {

}

// The action to perform when the rule matches.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsLogRuleAction string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsLogRuleActionLog RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRule) ImplementsRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams() {

}

// The action to perform when the rule matches.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleAction string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionSkip RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A phase to skip the execution of.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProduct string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProductBic           RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProductHot           RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProductRateLimit     RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProductSecurityLevel RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProductUaBlock       RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProductWAF           RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersRulesetCurrent RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsRulesetsSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A response object.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeJSON    `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelope]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrors]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                            `json:"pointer,required"`
	JSON    rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrorsSource]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessages]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                              `json:"pointer,required"`
	JSON    rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessagesSource]
type rulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeSuccess bool

const (
	RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeSuccessTrue RulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseEnvelopeSuccess = true
)
