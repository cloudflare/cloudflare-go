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

// Adds a new rule to an account or zone ruleset. The rule will be added to the end
// of the existing list of rules in the ruleset by default.
func (r *RulesetRuleService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, rulesetID string, body RulesetRuleNewParams, opts ...option.RequestOption) (res *RulesetRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetRuleNewResponseEnvelope
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

// A result.
type RulesetRuleNewResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetRuleNewResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetRuleNewResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetRuleNewResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                     `json:"description"`
	JSON        rulesetRuleNewResponseJSON `json:"-"`
}

// rulesetRuleNewResponseJSON contains the JSON metadata for the struct
// [RulesetRuleNewResponse]
type rulesetRuleNewResponseJSON struct {
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

func (r *RulesetRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetRuleNewResponseKind string

const (
	RulesetRuleNewResponseKindManaged RulesetRuleNewResponseKind = "managed"
	RulesetRuleNewResponseKindCustom  RulesetRuleNewResponseKind = "custom"
	RulesetRuleNewResponseKindRoot    RulesetRuleNewResponseKind = "root"
	RulesetRuleNewResponseKindZone    RulesetRuleNewResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetRuleNewResponsePhase string

const (
	RulesetRuleNewResponsePhaseDDOSL4                         RulesetRuleNewResponsePhase = "ddos_l4"
	RulesetRuleNewResponsePhaseDDOSL7                         RulesetRuleNewResponsePhase = "ddos_l7"
	RulesetRuleNewResponsePhaseHTTPConfigSettings             RulesetRuleNewResponsePhase = "http_config_settings"
	RulesetRuleNewResponsePhaseHTTPCustomErrors               RulesetRuleNewResponsePhase = "http_custom_errors"
	RulesetRuleNewResponsePhaseHTTPLogCustomFields            RulesetRuleNewResponsePhase = "http_log_custom_fields"
	RulesetRuleNewResponsePhaseHTTPRatelimit                  RulesetRuleNewResponsePhase = "http_ratelimit"
	RulesetRuleNewResponsePhaseHTTPRequestCacheSettings       RulesetRuleNewResponsePhase = "http_request_cache_settings"
	RulesetRuleNewResponsePhaseHTTPRequestDynamicRedirect     RulesetRuleNewResponsePhase = "http_request_dynamic_redirect"
	RulesetRuleNewResponsePhaseHTTPRequestFirewallCustom      RulesetRuleNewResponsePhase = "http_request_firewall_custom"
	RulesetRuleNewResponsePhaseHTTPRequestFirewallManaged     RulesetRuleNewResponsePhase = "http_request_firewall_managed"
	RulesetRuleNewResponsePhaseHTTPRequestLateTransform       RulesetRuleNewResponsePhase = "http_request_late_transform"
	RulesetRuleNewResponsePhaseHTTPRequestOrigin              RulesetRuleNewResponsePhase = "http_request_origin"
	RulesetRuleNewResponsePhaseHTTPRequestRedirect            RulesetRuleNewResponsePhase = "http_request_redirect"
	RulesetRuleNewResponsePhaseHTTPRequestSanitize            RulesetRuleNewResponsePhase = "http_request_sanitize"
	RulesetRuleNewResponsePhaseHTTPRequestSbfm                RulesetRuleNewResponsePhase = "http_request_sbfm"
	RulesetRuleNewResponsePhaseHTTPRequestSelectConfiguration RulesetRuleNewResponsePhase = "http_request_select_configuration"
	RulesetRuleNewResponsePhaseHTTPRequestTransform           RulesetRuleNewResponsePhase = "http_request_transform"
	RulesetRuleNewResponsePhaseHTTPResponseCompression        RulesetRuleNewResponsePhase = "http_response_compression"
	RulesetRuleNewResponsePhaseHTTPResponseFirewallManaged    RulesetRuleNewResponsePhase = "http_response_firewall_managed"
	RulesetRuleNewResponsePhaseHTTPResponseHeadersTransform   RulesetRuleNewResponsePhase = "http_response_headers_transform"
	RulesetRuleNewResponsePhaseMagicTransit                   RulesetRuleNewResponsePhase = "magic_transit"
	RulesetRuleNewResponsePhaseMagicTransitIDsManaged         RulesetRuleNewResponsePhase = "magic_transit_ids_managed"
	RulesetRuleNewResponsePhaseMagicTransitManaged            RulesetRuleNewResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetRuleNewResponseRulesRulesetsBlockRule],
// [RulesetRuleNewResponseRulesRulesetsExecuteRule],
// [RulesetRuleNewResponseRulesRulesetsLogRule] or
// [RulesetRuleNewResponseRulesRulesetsSkipRule].
type RulesetRuleNewResponseRule interface {
	implementsRulesetRuleNewResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetRuleNewResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleNewResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleNewResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleNewResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleNewResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetRuleNewResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleNewResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleNewResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleNewResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON rulesetRuleNewResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for
// the struct [RulesetRuleNewResponseRulesRulesetsBlockRule]
type rulesetRuleNewResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RulesetRuleNewResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleNewResponseRulesRulesetsBlockRule) implementsRulesetRuleNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleNewResponseRulesRulesetsBlockRuleAction string

const (
	RulesetRuleNewResponseRulesRulesetsBlockRuleActionBlock RulesetRuleNewResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetRuleNewResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetRuleNewResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetRuleNewResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsBlockRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetRuleNewResponseRulesRulesetsBlockRuleActionParameters]
type rulesetRuleNewResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response to show when the block is applied.
type RulesetRuleNewResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                    `json:"status_code,required"`
	JSON       rulesetRuleNewResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [RulesetRuleNewResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetRuleNewResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object configuring the rule's logging behavior.
type RulesetRuleNewResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                    `json:"enabled,required"`
	JSON    rulesetRuleNewResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleNewResponseRulesRulesetsBlockRuleLogging]
type rulesetRuleNewResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRuleNewResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleNewResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleNewResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                             `json:"ref"`
	JSON rulesetRuleNewResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata
// for the struct [RulesetRuleNewResponseRulesRulesetsExecuteRule]
type rulesetRuleNewResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RulesetRuleNewResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleNewResponseRulesRulesetsExecuteRule) implementsRulesetRuleNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleNewResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionExecute RulesetRuleNewResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration to use for matched data logging.
type RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                        `json:"public_key,required"`
	JSON      rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A set of overrides to apply to the target ruleset.
type RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A category-level override
type RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for rules in the category.
type RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for the rule.
type RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetRuleNewResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                      `json:"enabled,required"`
	JSON    rulesetRuleNewResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleNewResponseRulesRulesetsExecuteRuleLogging]
type rulesetRuleNewResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRuleNewResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleNewResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RulesetRuleNewResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                         `json:"ref"`
	JSON rulesetRuleNewResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsLogRuleJSON contains the JSON metadata for
// the struct [RulesetRuleNewResponseRulesRulesetsLogRule]
type rulesetRuleNewResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RulesetRuleNewResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleNewResponseRulesRulesetsLogRule) implementsRulesetRuleNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleNewResponseRulesRulesetsLogRuleAction string

const (
	RulesetRuleNewResponseRulesRulesetsLogRuleActionLog RulesetRuleNewResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetRuleNewResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                  `json:"enabled,required"`
	JSON    rulesetRuleNewResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsLogRuleLoggingJSON contains the JSON metadata
// for the struct [RulesetRuleNewResponseRulesRulesetsLogRuleLogging]
type rulesetRuleNewResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRuleNewResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleNewResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleNewResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleNewResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                          `json:"ref"`
	JSON rulesetRuleNewResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for
// the struct [RulesetRuleNewResponseRulesRulesetsSkipRule]
type rulesetRuleNewResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RulesetRuleNewResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRuleNewResponseRulesRulesetsSkipRule) implementsRulesetRuleNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleNewResponseRulesRulesetsSkipRuleAction string

const (
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionSkip RulesetRuleNewResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetRuleNewResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                        `json:"rulesets"`
	JSON     rulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetRuleNewResponseRulesRulesetsSkipRuleActionParameters]
type rulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A phase to skip the execution of.
type RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProductUaBlock       RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetRuleNewResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                   `json:"enabled,required"`
	JSON    rulesetRuleNewResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetRuleNewResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleNewResponseRulesRulesetsSkipRuleLogging]
type rulesetRuleNewResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// This interface is a union satisfied by one of the following:
// [RulesetRuleNewParamsRulesetsBlockRule],
// [RulesetRuleNewParamsRulesetsExecuteRule],
// [RulesetRuleNewParamsRulesetsLogRule], [RulesetRuleNewParamsRulesetsSkipRule].
type RulesetRuleNewParams interface {
	ImplementsRulesetRuleNewParams()
}

type RulesetRuleNewParamsRulesetsBlockRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetRuleNewParamsRulesetsBlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetRuleNewParamsRulesetsBlockRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetRuleNewParamsRulesetsBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetRuleNewParamsRulesetsBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (RulesetRuleNewParamsRulesetsBlockRule) ImplementsRulesetRuleNewParams() {

}

// The action to perform when the rule matches.
type RulesetRuleNewParamsRulesetsBlockRuleAction string

const (
	RulesetRuleNewParamsRulesetsBlockRuleActionBlock RulesetRuleNewParamsRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetRuleNewParamsRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[RulesetRuleNewParamsRulesetsBlockRuleActionParametersResponse] `json:"response"`
}

func (r RulesetRuleNewParamsRulesetsBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type RulesetRuleNewParamsRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r RulesetRuleNewParamsRulesetsBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type RulesetRuleNewParamsRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetRuleNewParamsRulesetsBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetRuleNewParamsRulesetsExecuteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetRuleNewParamsRulesetsExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetRuleNewParamsRulesetsExecuteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetRuleNewParamsRulesetsExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetRuleNewParamsRulesetsExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (RulesetRuleNewParamsRulesetsExecuteRule) ImplementsRulesetRuleNewParams() {

}

// The action to perform when the rule matches.
type RulesetRuleNewParamsRulesetsExecuteRuleAction string

const (
	RulesetRuleNewParamsRulesetsExecuteRuleActionExecute RulesetRuleNewParamsRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetRuleNewParamsRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[RulesetRuleNewParamsRulesetsExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r RulesetRuleNewParamsRulesetsExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type RulesetRuleNewParamsRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r RulesetRuleNewParamsRulesetsExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRule] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for rules in the category.
type RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for the rule.
type RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetRuleNewParamsRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetRuleNewParamsRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetRuleNewParamsRulesetsExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetRuleNewParamsRulesetsLogRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetRuleNewParamsRulesetsLogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetRuleNewParamsRulesetsLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetRuleNewParamsRulesetsLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (RulesetRuleNewParamsRulesetsLogRule) ImplementsRulesetRuleNewParams() {

}

// The action to perform when the rule matches.
type RulesetRuleNewParamsRulesetsLogRuleAction string

const (
	RulesetRuleNewParamsRulesetsLogRuleActionLog RulesetRuleNewParamsRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetRuleNewParamsRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetRuleNewParamsRulesetsLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetRuleNewParamsRulesetsSkipRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetRuleNewParamsRulesetsSkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetRuleNewParamsRulesetsSkipRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetRuleNewParamsRulesetsSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetRuleNewParamsRulesetsSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (RulesetRuleNewParamsRulesetsSkipRule) ImplementsRulesetRuleNewParams() {

}

// The action to perform when the rule matches.
type RulesetRuleNewParamsRulesetsSkipRuleAction string

const (
	RulesetRuleNewParamsRulesetsSkipRuleActionSkip RulesetRuleNewParamsRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetRuleNewParamsRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]RulesetRuleNewParamsRulesetsSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[RulesetRuleNewParamsRulesetsSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r RulesetRuleNewParamsRulesetsSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A phase to skip the execution of.
type RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase string

const (
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetRuleNewParamsRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetRuleNewParamsRulesetsSkipRuleActionParametersProduct string

const (
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersProductBic           RulesetRuleNewParamsRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersProductHot           RulesetRuleNewParamsRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersProductRateLimit     RulesetRuleNewParamsRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersProductSecurityLevel RulesetRuleNewParamsRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersProductUaBlock       RulesetRuleNewParamsRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersProductWAF           RulesetRuleNewParamsRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetRuleNewParamsRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetRuleNewParamsRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetRuleNewParamsRulesetsSkipRuleActionParametersRulesetCurrent RulesetRuleNewParamsRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetRuleNewParamsRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetRuleNewParamsRulesetsSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A response object.
type RulesetRuleNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetRuleNewResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetRuleNewResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetRuleNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetRuleNewResponseEnvelopeJSON    `json:"-"`
}

// rulesetRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetRuleNewResponseEnvelope]
type rulesetRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleNewResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleNewResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetRuleNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetRuleNewResponseEnvelopeErrors]
type rulesetRuleNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleNewResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                         `json:"pointer,required"`
	JSON    rulesetRuleNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetRuleNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RulesetRuleNewResponseEnvelopeErrorsSource]
type rulesetRuleNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleNewResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleNewResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetRuleNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetRuleNewResponseEnvelopeMessages]
type rulesetRuleNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleNewResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                           `json:"pointer,required"`
	JSON    rulesetRuleNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetRuleNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RulesetRuleNewResponseEnvelopeMessagesSource]
type rulesetRuleNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetRuleNewResponseEnvelopeSuccess bool

const (
	RulesetRuleNewResponseEnvelopeSuccessTrue RulesetRuleNewResponseEnvelopeSuccess = true
)

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
