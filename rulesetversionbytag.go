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

// RulesetVersionByTagService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRulesetVersionByTagService]
// method instead.
type RulesetVersionByTagService struct {
	Options []option.RequestOption
}

// NewRulesetVersionByTagService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRulesetVersionByTagService(opts ...option.RequestOption) (r *RulesetVersionByTagService) {
	r = &RulesetVersionByTagService{}
	r.Options = opts
	return
}

// Fetches the rules of a managed account ruleset version for a given tag.
func (r *RulesetVersionByTagService) Get(ctx context.Context, rulesetID string, rulesetVersion string, ruleTag string, query RulesetVersionByTagGetParams, opts ...option.RequestOption) (res *RulesetVersionByTagGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetVersionByTagGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s/by_tag/%s", query.AccountID, rulesetID, rulesetVersion, ruleTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A result.
type RulesetVersionByTagGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetVersionByTagGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetVersionByTagGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetVersionByTagGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                             `json:"description"`
	JSON        rulesetVersionByTagGetResponseJSON `json:"-"`
}

// rulesetVersionByTagGetResponseJSON contains the JSON metadata for the struct
// [RulesetVersionByTagGetResponse]
type rulesetVersionByTagGetResponseJSON struct {
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

func (r *RulesetVersionByTagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RulesetVersionByTagGetResponseKind string

const (
	RulesetVersionByTagGetResponseKindManaged RulesetVersionByTagGetResponseKind = "managed"
	RulesetVersionByTagGetResponseKindCustom  RulesetVersionByTagGetResponseKind = "custom"
	RulesetVersionByTagGetResponseKindRoot    RulesetVersionByTagGetResponseKind = "root"
	RulesetVersionByTagGetResponseKindZone    RulesetVersionByTagGetResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetVersionByTagGetResponsePhase string

const (
	RulesetVersionByTagGetResponsePhaseDDOSL4                         RulesetVersionByTagGetResponsePhase = "ddos_l4"
	RulesetVersionByTagGetResponsePhaseDDOSL7                         RulesetVersionByTagGetResponsePhase = "ddos_l7"
	RulesetVersionByTagGetResponsePhaseHTTPConfigSettings             RulesetVersionByTagGetResponsePhase = "http_config_settings"
	RulesetVersionByTagGetResponsePhaseHTTPCustomErrors               RulesetVersionByTagGetResponsePhase = "http_custom_errors"
	RulesetVersionByTagGetResponsePhaseHTTPLogCustomFields            RulesetVersionByTagGetResponsePhase = "http_log_custom_fields"
	RulesetVersionByTagGetResponsePhaseHTTPRatelimit                  RulesetVersionByTagGetResponsePhase = "http_ratelimit"
	RulesetVersionByTagGetResponsePhaseHTTPRequestCacheSettings       RulesetVersionByTagGetResponsePhase = "http_request_cache_settings"
	RulesetVersionByTagGetResponsePhaseHTTPRequestDynamicRedirect     RulesetVersionByTagGetResponsePhase = "http_request_dynamic_redirect"
	RulesetVersionByTagGetResponsePhaseHTTPRequestFirewallCustom      RulesetVersionByTagGetResponsePhase = "http_request_firewall_custom"
	RulesetVersionByTagGetResponsePhaseHTTPRequestFirewallManaged     RulesetVersionByTagGetResponsePhase = "http_request_firewall_managed"
	RulesetVersionByTagGetResponsePhaseHTTPRequestLateTransform       RulesetVersionByTagGetResponsePhase = "http_request_late_transform"
	RulesetVersionByTagGetResponsePhaseHTTPRequestOrigin              RulesetVersionByTagGetResponsePhase = "http_request_origin"
	RulesetVersionByTagGetResponsePhaseHTTPRequestRedirect            RulesetVersionByTagGetResponsePhase = "http_request_redirect"
	RulesetVersionByTagGetResponsePhaseHTTPRequestSanitize            RulesetVersionByTagGetResponsePhase = "http_request_sanitize"
	RulesetVersionByTagGetResponsePhaseHTTPRequestSbfm                RulesetVersionByTagGetResponsePhase = "http_request_sbfm"
	RulesetVersionByTagGetResponsePhaseHTTPRequestSelectConfiguration RulesetVersionByTagGetResponsePhase = "http_request_select_configuration"
	RulesetVersionByTagGetResponsePhaseHTTPRequestTransform           RulesetVersionByTagGetResponsePhase = "http_request_transform"
	RulesetVersionByTagGetResponsePhaseHTTPResponseCompression        RulesetVersionByTagGetResponsePhase = "http_response_compression"
	RulesetVersionByTagGetResponsePhaseHTTPResponseFirewallManaged    RulesetVersionByTagGetResponsePhase = "http_response_firewall_managed"
	RulesetVersionByTagGetResponsePhaseHTTPResponseHeadersTransform   RulesetVersionByTagGetResponsePhase = "http_response_headers_transform"
	RulesetVersionByTagGetResponsePhaseMagicTransit                   RulesetVersionByTagGetResponsePhase = "magic_transit"
	RulesetVersionByTagGetResponsePhaseMagicTransitIDsManaged         RulesetVersionByTagGetResponsePhase = "magic_transit_ids_managed"
	RulesetVersionByTagGetResponsePhaseMagicTransitManaged            RulesetVersionByTagGetResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetVersionByTagGetResponseRulesRulesetsBlockRule],
// [RulesetVersionByTagGetResponseRulesRulesetsExecuteRule],
// [RulesetVersionByTagGetResponseRulesRulesetsLogRule] or
// [RulesetVersionByTagGetResponseRulesRulesetsSkipRule].
type RulesetVersionByTagGetResponseRule interface {
	implementsRulesetVersionByTagGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetVersionByTagGetResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetVersionByTagGetResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetVersionByTagGetResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetVersionByTagGetResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetVersionByTagGetResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetVersionByTagGetResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetVersionByTagGetResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetVersionByTagGetResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                   `json:"ref"`
	JSON rulesetVersionByTagGetResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsBlockRuleJSON contains the JSON
// metadata for the struct [RulesetVersionByTagGetResponseRulesRulesetsBlockRule]
type rulesetVersionByTagGetResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RulesetVersionByTagGetResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetVersionByTagGetResponseRulesRulesetsBlockRule) implementsRulesetVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetVersionByTagGetResponseRulesRulesetsBlockRuleAction string

const (
	RulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionBlock RulesetVersionByTagGetResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParameters]
type rulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type RulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                            `json:"status_code,required"`
	JSON       rulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [RulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type RulesetVersionByTagGetResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                            `json:"enabled,required"`
	JSON    rulesetVersionByTagGetResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsBlockRuleLoggingJSON contains the
// JSON metadata for the struct
// [RulesetVersionByTagGetResponseRulesRulesetsBlockRuleLogging]
type rulesetVersionByTagGetResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetVersionByTagGetResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                     `json:"ref"`
	JSON rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleJSON contains the JSON
// metadata for the struct [RulesetVersionByTagGetResponseRulesRulesetsExecuteRule]
type rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RulesetVersionByTagGetResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetVersionByTagGetResponseRulesRulesetsExecuteRule) implementsRulesetVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionExecute RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                                `json:"public_key,required"`
	JSON      rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                              `json:"enabled,required"`
	JSON    rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleLoggingJSON contains the
// JSON metadata for the struct
// [RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleLogging]
type rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetVersionByTagGetResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetVersionByTagGetResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RulesetVersionByTagGetResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                 `json:"ref"`
	JSON rulesetVersionByTagGetResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsLogRuleJSON contains the JSON
// metadata for the struct [RulesetVersionByTagGetResponseRulesRulesetsLogRule]
type rulesetVersionByTagGetResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RulesetVersionByTagGetResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetVersionByTagGetResponseRulesRulesetsLogRule) implementsRulesetVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetVersionByTagGetResponseRulesRulesetsLogRuleAction string

const (
	RulesetVersionByTagGetResponseRulesRulesetsLogRuleActionLog RulesetVersionByTagGetResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetVersionByTagGetResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                          `json:"enabled,required"`
	JSON    rulesetVersionByTagGetResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsLogRuleLoggingJSON contains the JSON
// metadata for the struct
// [RulesetVersionByTagGetResponseRulesRulesetsLogRuleLogging]
type rulesetVersionByTagGetResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetVersionByTagGetResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetVersionByTagGetResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetVersionByTagGetResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                  `json:"ref"`
	JSON rulesetVersionByTagGetResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsSkipRuleJSON contains the JSON
// metadata for the struct [RulesetVersionByTagGetResponseRulesRulesetsSkipRule]
type rulesetVersionByTagGetResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RulesetVersionByTagGetResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetVersionByTagGetResponseRulesRulesetsSkipRule) implementsRulesetVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetVersionByTagGetResponseRulesRulesetsSkipRuleAction string

const (
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionSkip RulesetVersionByTagGetResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                                `json:"rulesets"`
	JSON     rulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParameters]
type rulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetVersionByTagGetResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetVersionByTagGetResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                           `json:"enabled,required"`
	JSON    rulesetVersionByTagGetResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetVersionByTagGetResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON
// metadata for the struct
// [RulesetVersionByTagGetResponseRulesRulesetsSkipRuleLogging]
type rulesetVersionByTagGetResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetVersionByTagGetParams struct {
	// The unique ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

// A response object.
type RulesetVersionByTagGetResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetVersionByTagGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetVersionByTagGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetVersionByTagGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetVersionByTagGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetVersionByTagGetResponseEnvelopeJSON    `json:"-"`
}

// rulesetVersionByTagGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RulesetVersionByTagGetResponseEnvelope]
type rulesetVersionByTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetVersionByTagGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetVersionByTagGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetVersionByTagGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetVersionByTagGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RulesetVersionByTagGetResponseEnvelopeErrors]
type rulesetVersionByTagGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetVersionByTagGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                 `json:"pointer,required"`
	JSON    rulesetVersionByTagGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetVersionByTagGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [RulesetVersionByTagGetResponseEnvelopeErrorsSource]
type rulesetVersionByTagGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetVersionByTagGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetVersionByTagGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetVersionByTagGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetVersionByTagGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RulesetVersionByTagGetResponseEnvelopeMessages]
type rulesetVersionByTagGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetVersionByTagGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                   `json:"pointer,required"`
	JSON    rulesetVersionByTagGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetVersionByTagGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RulesetVersionByTagGetResponseEnvelopeMessagesSource]
type rulesetVersionByTagGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetVersionByTagGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetVersionByTagGetResponseEnvelopeSuccess bool

const (
	RulesetVersionByTagGetResponseEnvelopeSuccessTrue RulesetVersionByTagGetResponseEnvelopeSuccess = true
)
