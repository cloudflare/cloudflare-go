// File generated from our OpenAPI spec by Stainless.

package cloudflare

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
func (r *RulesetRuleService) New(ctx context.Context, rulesetID string, params RulesetRuleNewParams, opts ...option.RequestOption) (res *RulesetRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetRuleNewResponseEnvelope
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
func (r *RulesetRuleService) Delete(ctx context.Context, rulesetID string, ruleID string, body RulesetRuleDeleteParams, opts ...option.RequestOption) (res *RulesetRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetRuleDeleteResponseEnvelope
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
func (r *RulesetRuleService) Edit(ctx context.Context, rulesetID string, ruleID string, params RulesetRuleEditParams, opts ...option.RequestOption) (res *RulesetRuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetRuleEditResponseEnvelope
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

func (r rulesetRuleNewResponseJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
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
	RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetRuleNewResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
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

func (r rulesetRuleNewResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
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
	RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetRuleDeleteResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
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

func (r rulesetRuleDeleteResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
}

// A result.
type RulesetRuleEditResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetRuleEditResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetRuleEditResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetRuleEditResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                      `json:"description"`
	JSON        rulesetRuleEditResponseJSON `json:"-"`
}

// rulesetRuleEditResponseJSON contains the JSON metadata for the struct
// [RulesetRuleEditResponse]
type rulesetRuleEditResponseJSON struct {
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

func (r *RulesetRuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RulesetRuleEditResponseKind string

const (
	RulesetRuleEditResponseKindManaged RulesetRuleEditResponseKind = "managed"
	RulesetRuleEditResponseKindCustom  RulesetRuleEditResponseKind = "custom"
	RulesetRuleEditResponseKindRoot    RulesetRuleEditResponseKind = "root"
	RulesetRuleEditResponseKindZone    RulesetRuleEditResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetRuleEditResponsePhase string

const (
	RulesetRuleEditResponsePhaseDDOSL4                         RulesetRuleEditResponsePhase = "ddos_l4"
	RulesetRuleEditResponsePhaseDDOSL7                         RulesetRuleEditResponsePhase = "ddos_l7"
	RulesetRuleEditResponsePhaseHTTPConfigSettings             RulesetRuleEditResponsePhase = "http_config_settings"
	RulesetRuleEditResponsePhaseHTTPCustomErrors               RulesetRuleEditResponsePhase = "http_custom_errors"
	RulesetRuleEditResponsePhaseHTTPLogCustomFields            RulesetRuleEditResponsePhase = "http_log_custom_fields"
	RulesetRuleEditResponsePhaseHTTPRatelimit                  RulesetRuleEditResponsePhase = "http_ratelimit"
	RulesetRuleEditResponsePhaseHTTPRequestCacheSettings       RulesetRuleEditResponsePhase = "http_request_cache_settings"
	RulesetRuleEditResponsePhaseHTTPRequestDynamicRedirect     RulesetRuleEditResponsePhase = "http_request_dynamic_redirect"
	RulesetRuleEditResponsePhaseHTTPRequestFirewallCustom      RulesetRuleEditResponsePhase = "http_request_firewall_custom"
	RulesetRuleEditResponsePhaseHTTPRequestFirewallManaged     RulesetRuleEditResponsePhase = "http_request_firewall_managed"
	RulesetRuleEditResponsePhaseHTTPRequestLateTransform       RulesetRuleEditResponsePhase = "http_request_late_transform"
	RulesetRuleEditResponsePhaseHTTPRequestOrigin              RulesetRuleEditResponsePhase = "http_request_origin"
	RulesetRuleEditResponsePhaseHTTPRequestRedirect            RulesetRuleEditResponsePhase = "http_request_redirect"
	RulesetRuleEditResponsePhaseHTTPRequestSanitize            RulesetRuleEditResponsePhase = "http_request_sanitize"
	RulesetRuleEditResponsePhaseHTTPRequestSbfm                RulesetRuleEditResponsePhase = "http_request_sbfm"
	RulesetRuleEditResponsePhaseHTTPRequestSelectConfiguration RulesetRuleEditResponsePhase = "http_request_select_configuration"
	RulesetRuleEditResponsePhaseHTTPRequestTransform           RulesetRuleEditResponsePhase = "http_request_transform"
	RulesetRuleEditResponsePhaseHTTPResponseCompression        RulesetRuleEditResponsePhase = "http_response_compression"
	RulesetRuleEditResponsePhaseHTTPResponseFirewallManaged    RulesetRuleEditResponsePhase = "http_response_firewall_managed"
	RulesetRuleEditResponsePhaseHTTPResponseHeadersTransform   RulesetRuleEditResponsePhase = "http_response_headers_transform"
	RulesetRuleEditResponsePhaseMagicTransit                   RulesetRuleEditResponsePhase = "magic_transit"
	RulesetRuleEditResponsePhaseMagicTransitIDsManaged         RulesetRuleEditResponsePhase = "magic_transit_ids_managed"
	RulesetRuleEditResponsePhaseMagicTransitManaged            RulesetRuleEditResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetRuleEditResponseRulesRulesetsBlockRule],
// [RulesetRuleEditResponseRulesRulesetsExecuteRule],
// [RulesetRuleEditResponseRulesRulesetsLogRule] or
// [RulesetRuleEditResponseRulesRulesetsSkipRule].
type RulesetRuleEditResponseRule interface {
	implementsRulesetRuleEditResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetRuleEditResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleEditResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleEditResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleEditResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetRuleEditResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetRuleEditResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleEditResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleEditResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleEditResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                            `json:"ref"`
	JSON rulesetRuleEditResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for
// the struct [RulesetRuleEditResponseRulesRulesetsBlockRule]
type rulesetRuleEditResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RulesetRuleEditResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetRuleEditResponseRulesRulesetsBlockRule) implementsRulesetRuleEditResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleEditResponseRulesRulesetsBlockRuleAction string

const (
	RulesetRuleEditResponseRulesRulesetsBlockRuleActionBlock RulesetRuleEditResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetRuleEditResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetRuleEditResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetRuleEditResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsBlockRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetRuleEditResponseRulesRulesetsBlockRuleActionParameters]
type rulesetRuleEditResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type RulesetRuleEditResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                     `json:"status_code,required"`
	JSON       rulesetRuleEditResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [RulesetRuleEditResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetRuleEditResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type RulesetRuleEditResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                     `json:"enabled,required"`
	JSON    rulesetRuleEditResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleEditResponseRulesRulesetsBlockRuleLogging]
type rulesetRuleEditResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetRuleEditResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleEditResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleEditResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                              `json:"ref"`
	JSON rulesetRuleEditResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata
// for the struct [RulesetRuleEditResponseRulesRulesetsExecuteRule]
type rulesetRuleEditResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RulesetRuleEditResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetRuleEditResponseRulesRulesetsExecuteRule) implementsRulesetRuleEditResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleEditResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionExecute RulesetRuleEditResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                         `json:"public_key,required"`
	JSON      rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetRuleEditResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetRuleEditResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                       `json:"enabled,required"`
	JSON    rulesetRuleEditResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleEditResponseRulesRulesetsExecuteRuleLogging]
type rulesetRuleEditResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetRuleEditResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleEditResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RulesetRuleEditResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                          `json:"ref"`
	JSON rulesetRuleEditResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsLogRuleJSON contains the JSON metadata for
// the struct [RulesetRuleEditResponseRulesRulesetsLogRule]
type rulesetRuleEditResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RulesetRuleEditResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetRuleEditResponseRulesRulesetsLogRule) implementsRulesetRuleEditResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleEditResponseRulesRulesetsLogRuleAction string

const (
	RulesetRuleEditResponseRulesRulesetsLogRuleActionLog RulesetRuleEditResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetRuleEditResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                   `json:"enabled,required"`
	JSON    rulesetRuleEditResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsLogRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleEditResponseRulesRulesetsLogRuleLogging]
type rulesetRuleEditResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetRuleEditResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetRuleEditResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetRuleEditResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetRuleEditResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON rulesetRuleEditResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for
// the struct [RulesetRuleEditResponseRulesRulesetsSkipRule]
type rulesetRuleEditResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RulesetRuleEditResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetRuleEditResponseRulesRulesetsSkipRule) implementsRulesetRuleEditResponseRule() {}

// The action to perform when the rule matches.
type RulesetRuleEditResponseRulesRulesetsSkipRuleAction string

const (
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionSkip RulesetRuleEditResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetRuleEditResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                         `json:"rulesets"`
	JSON     rulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetRuleEditResponseRulesRulesetsSkipRuleActionParameters]
type rulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetRuleEditResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetRuleEditResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                    `json:"enabled,required"`
	JSON    rulesetRuleEditResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetRuleEditResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetRuleEditResponseRulesRulesetsSkipRuleLogging]
type rulesetRuleEditResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetRuleNewParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An object configuring where the rule will be placed.
	Position param.Field[RulesetRuleNewParamsPosition] `json:"position"`
}

func (r RulesetRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
//
// Satisfied by [RulesetRuleNewParamsPositionBeforePosition],
// [RulesetRuleNewParamsPositionAfterPosition],
// [RulesetRuleNewParamsPositionIndexPosition].
type RulesetRuleNewParamsPosition interface {
	implementsRulesetRuleNewParamsPosition()
}

// An object configuring where the rule will be placed.
type RulesetRuleNewParamsPositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RulesetRuleNewParamsPositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetRuleNewParamsPositionBeforePosition) implementsRulesetRuleNewParamsPosition() {}

// An object configuring where the rule will be placed.
type RulesetRuleNewParamsPositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RulesetRuleNewParamsPositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetRuleNewParamsPositionAfterPosition) implementsRulesetRuleNewParamsPosition() {}

// An object configuring where the rule will be placed.
type RulesetRuleNewParamsPositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RulesetRuleNewParamsPositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetRuleNewParamsPositionIndexPosition) implementsRulesetRuleNewParamsPosition() {}

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

func (r rulesetRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetRuleNewResponseEnvelopeSuccess bool

const (
	RulesetRuleNewResponseEnvelopeSuccessTrue RulesetRuleNewResponseEnvelopeSuccess = true
)

type RulesetRuleDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

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

func (r rulesetRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r rulesetRuleDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetRuleDeleteResponseEnvelopeSuccess bool

const (
	RulesetRuleDeleteResponseEnvelopeSuccessTrue RulesetRuleDeleteResponseEnvelopeSuccess = true
)

type RulesetRuleEditParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An object configuring where the rule will be placed.
	Position param.Field[RulesetRuleEditParamsPosition] `json:"position"`
}

func (r RulesetRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
//
// Satisfied by [RulesetRuleEditParamsPositionBeforePosition],
// [RulesetRuleEditParamsPositionAfterPosition],
// [RulesetRuleEditParamsPositionIndexPosition].
type RulesetRuleEditParamsPosition interface {
	implementsRulesetRuleEditParamsPosition()
}

// An object configuring where the rule will be placed.
type RulesetRuleEditParamsPositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RulesetRuleEditParamsPositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetRuleEditParamsPositionBeforePosition) implementsRulesetRuleEditParamsPosition() {}

// An object configuring where the rule will be placed.
type RulesetRuleEditParamsPositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RulesetRuleEditParamsPositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetRuleEditParamsPositionAfterPosition) implementsRulesetRuleEditParamsPosition() {}

// An object configuring where the rule will be placed.
type RulesetRuleEditParamsPositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RulesetRuleEditParamsPositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetRuleEditParamsPositionIndexPosition) implementsRulesetRuleEditParamsPosition() {}

// A response object.
type RulesetRuleEditResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetRuleEditResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetRuleEditResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetRuleEditResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetRuleEditResponseEnvelopeJSON    `json:"-"`
}

// rulesetRuleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetRuleEditResponseEnvelope]
type rulesetRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetRuleEditResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleEditResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetRuleEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetRuleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetRuleEditResponseEnvelopeErrors]
type rulesetRuleEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetRuleEditResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    rulesetRuleEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetRuleEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RulesetRuleEditResponseEnvelopeErrorsSource]
type rulesetRuleEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetRuleEditResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleEditResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetRuleEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetRuleEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetRuleEditResponseEnvelopeMessages]
type rulesetRuleEditResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetRuleEditResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    rulesetRuleEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetRuleEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RulesetRuleEditResponseEnvelopeMessagesSource]
type rulesetRuleEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetRuleEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetRuleEditResponseEnvelopeSuccess bool

const (
	RulesetRuleEditResponseEnvelopeSuccessTrue RulesetRuleEditResponseEnvelopeSuccess = true
)
