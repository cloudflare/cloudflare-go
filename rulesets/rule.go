// File generated from our OpenAPI spec by Stainless.

package rulesets

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
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
func (r *RuleService) New(ctx context.Context, rulesetID string, params RuleNewParams, opts ...option.RequestOption) (res *RuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
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
func (r *RuleService) Delete(ctx context.Context, rulesetID string, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) (res *RuleDeleteResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing rule in an account ruleset.
func (r *RuleService) Edit(ctx context.Context, rulesetID string, ruleID string, params RuleEditParams, opts ...option.RequestOption) (res *RuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleEditResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/%s", accountOrZone, accountOrZoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A result.
type RuleNewResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RuleNewResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RuleNewResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RuleNewResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string              `json:"description"`
	JSON        ruleNewResponseJSON `json:"-"`
}

// ruleNewResponseJSON contains the JSON metadata for the struct [RuleNewResponse]
type ruleNewResponseJSON struct {
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

func (r *RuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RuleNewResponseKind string

const (
	RuleNewResponseKindManaged RuleNewResponseKind = "managed"
	RuleNewResponseKindCustom  RuleNewResponseKind = "custom"
	RuleNewResponseKindRoot    RuleNewResponseKind = "root"
	RuleNewResponseKindZone    RuleNewResponseKind = "zone"
)

// The phase of the ruleset.
type RuleNewResponsePhase string

const (
	RuleNewResponsePhaseDDOSL4                         RuleNewResponsePhase = "ddos_l4"
	RuleNewResponsePhaseDDOSL7                         RuleNewResponsePhase = "ddos_l7"
	RuleNewResponsePhaseHTTPConfigSettings             RuleNewResponsePhase = "http_config_settings"
	RuleNewResponsePhaseHTTPCustomErrors               RuleNewResponsePhase = "http_custom_errors"
	RuleNewResponsePhaseHTTPLogCustomFields            RuleNewResponsePhase = "http_log_custom_fields"
	RuleNewResponsePhaseHTTPRatelimit                  RuleNewResponsePhase = "http_ratelimit"
	RuleNewResponsePhaseHTTPRequestCacheSettings       RuleNewResponsePhase = "http_request_cache_settings"
	RuleNewResponsePhaseHTTPRequestDynamicRedirect     RuleNewResponsePhase = "http_request_dynamic_redirect"
	RuleNewResponsePhaseHTTPRequestFirewallCustom      RuleNewResponsePhase = "http_request_firewall_custom"
	RuleNewResponsePhaseHTTPRequestFirewallManaged     RuleNewResponsePhase = "http_request_firewall_managed"
	RuleNewResponsePhaseHTTPRequestLateTransform       RuleNewResponsePhase = "http_request_late_transform"
	RuleNewResponsePhaseHTTPRequestOrigin              RuleNewResponsePhase = "http_request_origin"
	RuleNewResponsePhaseHTTPRequestRedirect            RuleNewResponsePhase = "http_request_redirect"
	RuleNewResponsePhaseHTTPRequestSanitize            RuleNewResponsePhase = "http_request_sanitize"
	RuleNewResponsePhaseHTTPRequestSbfm                RuleNewResponsePhase = "http_request_sbfm"
	RuleNewResponsePhaseHTTPRequestSelectConfiguration RuleNewResponsePhase = "http_request_select_configuration"
	RuleNewResponsePhaseHTTPRequestTransform           RuleNewResponsePhase = "http_request_transform"
	RuleNewResponsePhaseHTTPResponseCompression        RuleNewResponsePhase = "http_response_compression"
	RuleNewResponsePhaseHTTPResponseFirewallManaged    RuleNewResponsePhase = "http_response_firewall_managed"
	RuleNewResponsePhaseHTTPResponseHeadersTransform   RuleNewResponsePhase = "http_response_headers_transform"
	RuleNewResponsePhaseMagicTransit                   RuleNewResponsePhase = "magic_transit"
	RuleNewResponsePhaseMagicTransitIDsManaged         RuleNewResponsePhase = "magic_transit_ids_managed"
	RuleNewResponsePhaseMagicTransitManaged            RuleNewResponsePhase = "magic_transit_managed"
)

// Union satisfied by [rulesets.RuleNewResponseRulesRulesetsBlockRule],
// [rulesets.RuleNewResponseRulesRulesetsExecuteRule],
// [rulesets.RuleNewResponseRulesRulesetsLogRule] or
// [rulesets.RuleNewResponseRulesRulesetsSkipRule].
type RuleNewResponseRule interface {
	implementsRulesetsRuleNewResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleNewResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleNewResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleNewResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleNewResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleNewResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RuleNewResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleNewResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RuleNewResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RuleNewResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                    `json:"ref"`
	JSON ruleNewResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for the
// struct [RuleNewResponseRulesRulesetsBlockRule]
type ruleNewResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RuleNewResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleNewResponseRulesRulesetsBlockRule) implementsRulesetsRuleNewResponseRule() {}

// The action to perform when the rule matches.
type RuleNewResponseRulesRulesetsBlockRuleAction string

const (
	RuleNewResponseRulesRulesetsBlockRuleActionBlock RuleNewResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RuleNewResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RuleNewResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     ruleNewResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// ruleNewResponseRulesRulesetsBlockRuleActionParametersJSON contains the JSON
// metadata for the struct [RuleNewResponseRulesRulesetsBlockRuleActionParameters]
type ruleNewResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type RuleNewResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                             `json:"status_code,required"`
	JSON       ruleNewResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsBlockRuleActionParametersResponseJSON contains the
// JSON metadata for the struct
// [RuleNewResponseRulesRulesetsBlockRuleActionParametersResponse]
type ruleNewResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type RuleNewResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                             `json:"enabled,required"`
	JSON    ruleNewResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON metadata for
// the struct [RuleNewResponseRulesRulesetsBlockRuleLogging]
type ruleNewResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RuleNewResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleNewResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RuleNewResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RuleNewResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                      `json:"ref"`
	JSON ruleNewResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata for the
// struct [RuleNewResponseRulesRulesetsExecuteRule]
type ruleNewResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RuleNewResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleNewResponseRulesRulesetsExecuteRule) implementsRulesetsRuleNewResponseRule() {}

// The action to perform when the rule matches.
type RuleNewResponseRulesRulesetsExecuteRuleAction string

const (
	RuleNewResponseRulesRulesetsExecuteRuleActionExecute RuleNewResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RuleNewResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RuleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      ruleNewResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// ruleNewResponseRulesRulesetsExecuteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RuleNewResponseRulesRulesetsExecuteRuleActionParameters]
type ruleNewResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type RuleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                 `json:"public_key,required"`
	JSON      ruleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON contains
// the JSON metadata for the struct
// [RuleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type ruleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             ruleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// ruleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON contains
// the JSON metadata for the struct
// [RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type ruleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             ruleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// ruleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type ruleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             ruleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// ruleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type ruleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RuleNewResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                               `json:"enabled,required"`
	JSON    ruleNewResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON metadata
// for the struct [RuleNewResponseRulesRulesetsExecuteRuleLogging]
type ruleNewResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RuleNewResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleNewResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RuleNewResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                  `json:"ref"`
	JSON ruleNewResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsLogRuleJSON contains the JSON metadata for the
// struct [RuleNewResponseRulesRulesetsLogRule]
type ruleNewResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RuleNewResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleNewResponseRulesRulesetsLogRule) implementsRulesetsRuleNewResponseRule() {}

// The action to perform when the rule matches.
type RuleNewResponseRulesRulesetsLogRuleAction string

const (
	RuleNewResponseRulesRulesetsLogRuleActionLog RuleNewResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RuleNewResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                           `json:"enabled,required"`
	JSON    ruleNewResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsLogRuleLoggingJSON contains the JSON metadata for
// the struct [RuleNewResponseRulesRulesetsLogRuleLogging]
type ruleNewResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RuleNewResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleNewResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RuleNewResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RuleNewResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                   `json:"ref"`
	JSON ruleNewResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for the
// struct [RuleNewResponseRulesRulesetsSkipRule]
type ruleNewResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RuleNewResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleNewResponseRulesRulesetsSkipRule) implementsRulesetsRuleNewResponseRule() {}

// The action to perform when the rule matches.
type RuleNewResponseRulesRulesetsSkipRuleAction string

const (
	RuleNewResponseRulesRulesetsSkipRuleActionSkip RuleNewResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RuleNewResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RuleNewResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RuleNewResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                 `json:"rulesets"`
	JSON     ruleNewResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsSkipRuleActionParametersJSON contains the JSON
// metadata for the struct [RuleNewResponseRulesRulesetsSkipRuleActionParameters]
type ruleNewResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RuleNewResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RuleNewResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RuleNewResponseRulesRulesetsSkipRuleActionParametersProductBic           RuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersProductHot           RuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersProductWAF           RuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RuleNewResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RuleNewResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RuleNewResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RuleNewResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RuleNewResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                            `json:"enabled,required"`
	JSON    ruleNewResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON metadata for
// the struct [RuleNewResponseRulesRulesetsSkipRuleLogging]
type ruleNewResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
}

// A result.
type RuleDeleteResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RuleDeleteResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RuleDeleteResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RuleDeleteResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                 `json:"description"`
	JSON        ruleDeleteResponseJSON `json:"-"`
}

// ruleDeleteResponseJSON contains the JSON metadata for the struct
// [RuleDeleteResponse]
type ruleDeleteResponseJSON struct {
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

func (r *RuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RuleDeleteResponseKind string

const (
	RuleDeleteResponseKindManaged RuleDeleteResponseKind = "managed"
	RuleDeleteResponseKindCustom  RuleDeleteResponseKind = "custom"
	RuleDeleteResponseKindRoot    RuleDeleteResponseKind = "root"
	RuleDeleteResponseKindZone    RuleDeleteResponseKind = "zone"
)

// The phase of the ruleset.
type RuleDeleteResponsePhase string

const (
	RuleDeleteResponsePhaseDDOSL4                         RuleDeleteResponsePhase = "ddos_l4"
	RuleDeleteResponsePhaseDDOSL7                         RuleDeleteResponsePhase = "ddos_l7"
	RuleDeleteResponsePhaseHTTPConfigSettings             RuleDeleteResponsePhase = "http_config_settings"
	RuleDeleteResponsePhaseHTTPCustomErrors               RuleDeleteResponsePhase = "http_custom_errors"
	RuleDeleteResponsePhaseHTTPLogCustomFields            RuleDeleteResponsePhase = "http_log_custom_fields"
	RuleDeleteResponsePhaseHTTPRatelimit                  RuleDeleteResponsePhase = "http_ratelimit"
	RuleDeleteResponsePhaseHTTPRequestCacheSettings       RuleDeleteResponsePhase = "http_request_cache_settings"
	RuleDeleteResponsePhaseHTTPRequestDynamicRedirect     RuleDeleteResponsePhase = "http_request_dynamic_redirect"
	RuleDeleteResponsePhaseHTTPRequestFirewallCustom      RuleDeleteResponsePhase = "http_request_firewall_custom"
	RuleDeleteResponsePhaseHTTPRequestFirewallManaged     RuleDeleteResponsePhase = "http_request_firewall_managed"
	RuleDeleteResponsePhaseHTTPRequestLateTransform       RuleDeleteResponsePhase = "http_request_late_transform"
	RuleDeleteResponsePhaseHTTPRequestOrigin              RuleDeleteResponsePhase = "http_request_origin"
	RuleDeleteResponsePhaseHTTPRequestRedirect            RuleDeleteResponsePhase = "http_request_redirect"
	RuleDeleteResponsePhaseHTTPRequestSanitize            RuleDeleteResponsePhase = "http_request_sanitize"
	RuleDeleteResponsePhaseHTTPRequestSbfm                RuleDeleteResponsePhase = "http_request_sbfm"
	RuleDeleteResponsePhaseHTTPRequestSelectConfiguration RuleDeleteResponsePhase = "http_request_select_configuration"
	RuleDeleteResponsePhaseHTTPRequestTransform           RuleDeleteResponsePhase = "http_request_transform"
	RuleDeleteResponsePhaseHTTPResponseCompression        RuleDeleteResponsePhase = "http_response_compression"
	RuleDeleteResponsePhaseHTTPResponseFirewallManaged    RuleDeleteResponsePhase = "http_response_firewall_managed"
	RuleDeleteResponsePhaseHTTPResponseHeadersTransform   RuleDeleteResponsePhase = "http_response_headers_transform"
	RuleDeleteResponsePhaseMagicTransit                   RuleDeleteResponsePhase = "magic_transit"
	RuleDeleteResponsePhaseMagicTransitIDsManaged         RuleDeleteResponsePhase = "magic_transit_ids_managed"
	RuleDeleteResponsePhaseMagicTransitManaged            RuleDeleteResponsePhase = "magic_transit_managed"
)

// Union satisfied by [rulesets.RuleDeleteResponseRulesRulesetsBlockRule],
// [rulesets.RuleDeleteResponseRulesRulesetsExecuteRule],
// [rulesets.RuleDeleteResponseRulesRulesetsLogRule] or
// [rulesets.RuleDeleteResponseRulesRulesetsSkipRule].
type RuleDeleteResponseRule interface {
	implementsRulesetsRuleDeleteResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleDeleteResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleDeleteResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleDeleteResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleDeleteResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleDeleteResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RuleDeleteResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleDeleteResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RuleDeleteResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RuleDeleteResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                       `json:"ref"`
	JSON ruleDeleteResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for the
// struct [RuleDeleteResponseRulesRulesetsBlockRule]
type ruleDeleteResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RuleDeleteResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleDeleteResponseRulesRulesetsBlockRule) implementsRulesetsRuleDeleteResponseRule() {}

// The action to perform when the rule matches.
type RuleDeleteResponseRulesRulesetsBlockRuleAction string

const (
	RuleDeleteResponseRulesRulesetsBlockRuleActionBlock RuleDeleteResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RuleDeleteResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RuleDeleteResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     ruleDeleteResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// ruleDeleteResponseRulesRulesetsBlockRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RuleDeleteResponseRulesRulesetsBlockRuleActionParameters]
type ruleDeleteResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type RuleDeleteResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                `json:"status_code,required"`
	JSON       ruleDeleteResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsBlockRuleActionParametersResponseJSON contains
// the JSON metadata for the struct
// [RuleDeleteResponseRulesRulesetsBlockRuleActionParametersResponse]
type ruleDeleteResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type RuleDeleteResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                `json:"enabled,required"`
	JSON    ruleDeleteResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON metadata
// for the struct [RuleDeleteResponseRulesRulesetsBlockRuleLogging]
type ruleDeleteResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleDeleteResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RuleDeleteResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RuleDeleteResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                         `json:"ref"`
	JSON ruleDeleteResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata for
// the struct [RuleDeleteResponseRulesRulesetsExecuteRule]
type ruleDeleteResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RuleDeleteResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleDeleteResponseRulesRulesetsExecuteRule) implementsRulesetsRuleDeleteResponseRule() {}

// The action to perform when the rule matches.
type RuleDeleteResponseRulesRulesetsExecuteRuleAction string

const (
	RuleDeleteResponseRulesRulesetsExecuteRuleActionExecute RuleDeleteResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RuleDeleteResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RuleDeleteResponseRulesRulesetsExecuteRuleActionParameters]
type ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                    `json:"public_key,required"`
	JSON      ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON contains
// the JSON metadata for the struct
// [RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RuleDeleteResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                  `json:"enabled,required"`
	JSON    ruleDeleteResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON metadata
// for the struct [RuleDeleteResponseRulesRulesetsExecuteRuleLogging]
type ruleDeleteResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleDeleteResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RuleDeleteResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                     `json:"ref"`
	JSON ruleDeleteResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsLogRuleJSON contains the JSON metadata for the
// struct [RuleDeleteResponseRulesRulesetsLogRule]
type ruleDeleteResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RuleDeleteResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleDeleteResponseRulesRulesetsLogRule) implementsRulesetsRuleDeleteResponseRule() {}

// The action to perform when the rule matches.
type RuleDeleteResponseRulesRulesetsLogRuleAction string

const (
	RuleDeleteResponseRulesRulesetsLogRuleActionLog RuleDeleteResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RuleDeleteResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                              `json:"enabled,required"`
	JSON    ruleDeleteResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsLogRuleLoggingJSON contains the JSON metadata for
// the struct [RuleDeleteResponseRulesRulesetsLogRuleLogging]
type ruleDeleteResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleDeleteResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RuleDeleteResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RuleDeleteResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                      `json:"ref"`
	JSON ruleDeleteResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for the
// struct [RuleDeleteResponseRulesRulesetsSkipRule]
type ruleDeleteResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RuleDeleteResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleDeleteResponseRulesRulesetsSkipRule) implementsRulesetsRuleDeleteResponseRule() {}

// The action to perform when the rule matches.
type RuleDeleteResponseRulesRulesetsSkipRuleAction string

const (
	RuleDeleteResponseRulesRulesetsSkipRuleActionSkip RuleDeleteResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RuleDeleteResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RuleDeleteResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                    `json:"rulesets"`
	JSON     ruleDeleteResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsSkipRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RuleDeleteResponseRulesRulesetsSkipRuleActionParameters]
type ruleDeleteResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RuleDeleteResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductBic           RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductHot           RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductWAF           RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RuleDeleteResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RuleDeleteResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RuleDeleteResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RuleDeleteResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                               `json:"enabled,required"`
	JSON    ruleDeleteResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON metadata
// for the struct [RuleDeleteResponseRulesRulesetsSkipRuleLogging]
type ruleDeleteResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
}

// A result.
type RuleEditResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RuleEditResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RuleEditResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RuleEditResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string               `json:"description"`
	JSON        ruleEditResponseJSON `json:"-"`
}

// ruleEditResponseJSON contains the JSON metadata for the struct
// [RuleEditResponse]
type ruleEditResponseJSON struct {
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

func (r *RuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RuleEditResponseKind string

const (
	RuleEditResponseKindManaged RuleEditResponseKind = "managed"
	RuleEditResponseKindCustom  RuleEditResponseKind = "custom"
	RuleEditResponseKindRoot    RuleEditResponseKind = "root"
	RuleEditResponseKindZone    RuleEditResponseKind = "zone"
)

// The phase of the ruleset.
type RuleEditResponsePhase string

const (
	RuleEditResponsePhaseDDOSL4                         RuleEditResponsePhase = "ddos_l4"
	RuleEditResponsePhaseDDOSL7                         RuleEditResponsePhase = "ddos_l7"
	RuleEditResponsePhaseHTTPConfigSettings             RuleEditResponsePhase = "http_config_settings"
	RuleEditResponsePhaseHTTPCustomErrors               RuleEditResponsePhase = "http_custom_errors"
	RuleEditResponsePhaseHTTPLogCustomFields            RuleEditResponsePhase = "http_log_custom_fields"
	RuleEditResponsePhaseHTTPRatelimit                  RuleEditResponsePhase = "http_ratelimit"
	RuleEditResponsePhaseHTTPRequestCacheSettings       RuleEditResponsePhase = "http_request_cache_settings"
	RuleEditResponsePhaseHTTPRequestDynamicRedirect     RuleEditResponsePhase = "http_request_dynamic_redirect"
	RuleEditResponsePhaseHTTPRequestFirewallCustom      RuleEditResponsePhase = "http_request_firewall_custom"
	RuleEditResponsePhaseHTTPRequestFirewallManaged     RuleEditResponsePhase = "http_request_firewall_managed"
	RuleEditResponsePhaseHTTPRequestLateTransform       RuleEditResponsePhase = "http_request_late_transform"
	RuleEditResponsePhaseHTTPRequestOrigin              RuleEditResponsePhase = "http_request_origin"
	RuleEditResponsePhaseHTTPRequestRedirect            RuleEditResponsePhase = "http_request_redirect"
	RuleEditResponsePhaseHTTPRequestSanitize            RuleEditResponsePhase = "http_request_sanitize"
	RuleEditResponsePhaseHTTPRequestSbfm                RuleEditResponsePhase = "http_request_sbfm"
	RuleEditResponsePhaseHTTPRequestSelectConfiguration RuleEditResponsePhase = "http_request_select_configuration"
	RuleEditResponsePhaseHTTPRequestTransform           RuleEditResponsePhase = "http_request_transform"
	RuleEditResponsePhaseHTTPResponseCompression        RuleEditResponsePhase = "http_response_compression"
	RuleEditResponsePhaseHTTPResponseFirewallManaged    RuleEditResponsePhase = "http_response_firewall_managed"
	RuleEditResponsePhaseHTTPResponseHeadersTransform   RuleEditResponsePhase = "http_response_headers_transform"
	RuleEditResponsePhaseMagicTransit                   RuleEditResponsePhase = "magic_transit"
	RuleEditResponsePhaseMagicTransitIDsManaged         RuleEditResponsePhase = "magic_transit_ids_managed"
	RuleEditResponsePhaseMagicTransitManaged            RuleEditResponsePhase = "magic_transit_managed"
)

// Union satisfied by [rulesets.RuleEditResponseRulesRulesetsBlockRule],
// [rulesets.RuleEditResponseRulesRulesetsExecuteRule],
// [rulesets.RuleEditResponseRulesRulesetsLogRule] or
// [rulesets.RuleEditResponseRulesRulesetsSkipRule].
type RuleEditResponseRule interface {
	implementsRulesetsRuleEditResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleEditResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleEditResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleEditResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleEditResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleEditResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RuleEditResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleEditResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RuleEditResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RuleEditResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                     `json:"ref"`
	JSON ruleEditResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for the
// struct [RuleEditResponseRulesRulesetsBlockRule]
type ruleEditResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RuleEditResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleEditResponseRulesRulesetsBlockRule) implementsRulesetsRuleEditResponseRule() {}

// The action to perform when the rule matches.
type RuleEditResponseRulesRulesetsBlockRuleAction string

const (
	RuleEditResponseRulesRulesetsBlockRuleActionBlock RuleEditResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RuleEditResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RuleEditResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     ruleEditResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// ruleEditResponseRulesRulesetsBlockRuleActionParametersJSON contains the JSON
// metadata for the struct [RuleEditResponseRulesRulesetsBlockRuleActionParameters]
type ruleEditResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type RuleEditResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                              `json:"status_code,required"`
	JSON       ruleEditResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsBlockRuleActionParametersResponseJSON contains the
// JSON metadata for the struct
// [RuleEditResponseRulesRulesetsBlockRuleActionParametersResponse]
type ruleEditResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type RuleEditResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                              `json:"enabled,required"`
	JSON    ruleEditResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON metadata for
// the struct [RuleEditResponseRulesRulesetsBlockRuleLogging]
type ruleEditResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RuleEditResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleEditResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RuleEditResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RuleEditResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                       `json:"ref"`
	JSON ruleEditResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata for the
// struct [RuleEditResponseRulesRulesetsExecuteRule]
type ruleEditResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RuleEditResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleEditResponseRulesRulesetsExecuteRule) implementsRulesetsRuleEditResponseRule() {}

// The action to perform when the rule matches.
type RuleEditResponseRulesRulesetsExecuteRuleAction string

const (
	RuleEditResponseRulesRulesetsExecuteRuleActionExecute RuleEditResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RuleEditResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RuleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      ruleEditResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// ruleEditResponseRulesRulesetsExecuteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RuleEditResponseRulesRulesetsExecuteRuleActionParameters]
type ruleEditResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type RuleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                  `json:"public_key,required"`
	JSON      ruleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON contains
// the JSON metadata for the struct
// [RuleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type ruleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             ruleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// ruleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON contains
// the JSON metadata for the struct
// [RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type ruleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             ruleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// ruleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type ruleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             ruleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// ruleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type ruleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RuleEditResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                `json:"enabled,required"`
	JSON    ruleEditResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON metadata
// for the struct [RuleEditResponseRulesRulesetsExecuteRuleLogging]
type ruleEditResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RuleEditResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleEditResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RuleEditResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                   `json:"ref"`
	JSON ruleEditResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsLogRuleJSON contains the JSON metadata for the
// struct [RuleEditResponseRulesRulesetsLogRule]
type ruleEditResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RuleEditResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleEditResponseRulesRulesetsLogRule) implementsRulesetsRuleEditResponseRule() {}

// The action to perform when the rule matches.
type RuleEditResponseRulesRulesetsLogRuleAction string

const (
	RuleEditResponseRulesRulesetsLogRuleActionLog RuleEditResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RuleEditResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                            `json:"enabled,required"`
	JSON    ruleEditResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsLogRuleLoggingJSON contains the JSON metadata for
// the struct [RuleEditResponseRulesRulesetsLogRuleLogging]
type ruleEditResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RuleEditResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleEditResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RuleEditResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RuleEditResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                    `json:"ref"`
	JSON ruleEditResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for the
// struct [RuleEditResponseRulesRulesetsSkipRule]
type ruleEditResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RuleEditResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleEditResponseRulesRulesetsSkipRule) implementsRulesetsRuleEditResponseRule() {}

// The action to perform when the rule matches.
type RuleEditResponseRulesRulesetsSkipRuleAction string

const (
	RuleEditResponseRulesRulesetsSkipRuleActionSkip RuleEditResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RuleEditResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RuleEditResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RuleEditResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                  `json:"rulesets"`
	JSON     ruleEditResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsSkipRuleActionParametersJSON contains the JSON
// metadata for the struct [RuleEditResponseRulesRulesetsSkipRuleActionParameters]
type ruleEditResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RuleEditResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RuleEditResponseRulesRulesetsSkipRuleActionParametersProductBic           RuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersProductHot           RuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersProductWAF           RuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RuleEditResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RuleEditResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RuleEditResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RuleEditResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RuleEditResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                             `json:"enabled,required"`
	JSON    ruleEditResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON metadata for
// the struct [RuleEditResponseRulesRulesetsSkipRuleLogging]
type ruleEditResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RuleNewParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsPosition] `json:"position"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsPositionBeforePosition],
// [rulesets.RuleNewParamsPositionAfterPosition],
// [rulesets.RuleNewParamsPositionIndexPosition].
type RuleNewParamsPosition interface {
	implementsRulesetsRuleNewParamsPosition()
}

// An object configuring where the rule will be placed.
type RuleNewParamsPositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsPositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsPositionBeforePosition) implementsRulesetsRuleNewParamsPosition() {}

// An object configuring where the rule will be placed.
type RuleNewParamsPositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsPositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsPositionAfterPosition) implementsRulesetsRuleNewParamsPosition() {}

// An object configuring where the rule will be placed.
type RuleNewParamsPositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsPositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsPositionIndexPosition) implementsRulesetsRuleNewParamsPosition() {}

// A response object.
type RuleNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleNewResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleNewResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RuleNewResponse `json:"result,required"`
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
	Result RuleDeleteResponse `json:"result,required"`
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

type RuleEditParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsPosition] `json:"position"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsPositionBeforePosition],
// [rulesets.RuleEditParamsPositionAfterPosition],
// [rulesets.RuleEditParamsPositionIndexPosition].
type RuleEditParamsPosition interface {
	implementsRulesetsRuleEditParamsPosition()
}

// An object configuring where the rule will be placed.
type RuleEditParamsPositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsPositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionBeforePosition) implementsRulesetsRuleEditParamsPosition() {}

// An object configuring where the rule will be placed.
type RuleEditParamsPositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsPositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionAfterPosition) implementsRulesetsRuleEditParamsPosition() {}

// An object configuring where the rule will be placed.
type RuleEditParamsPositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsPositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionIndexPosition) implementsRulesetsRuleEditParamsPosition() {}

// A response object.
type RuleEditResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleEditResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleEditResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RuleEditResponse `json:"result,required"`
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
