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

// RulesetPhaseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRulesetPhaseService] method
// instead.
type RulesetPhaseService struct {
	Options []option.RequestOption
}

// NewRulesetPhaseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRulesetPhaseService(opts ...option.RequestOption) (r *RulesetPhaseService) {
	r = &RulesetPhaseService{}
	r.Options = opts
	return
}

// Fetches the latest version of the account or zone entry point ruleset for a
// given phase.
func (r *RulesetPhaseService) Get(ctx context.Context, rulesetPhase RulesetPhaseGetParamsRulesetPhase, query RulesetPhaseGetParams, opts ...option.RequestOption) (res *RulesetPhaseGetResponse, err error) {
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

// A result.
type RulesetPhaseGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetPhaseGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetPhaseGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetPhaseGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                      `json:"description"`
	JSON        rulesetPhaseGetResponseJSON `json:"-"`
}

// rulesetPhaseGetResponseJSON contains the JSON metadata for the struct
// [RulesetPhaseGetResponse]
type rulesetPhaseGetResponseJSON struct {
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

func (r *RulesetPhaseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetPhaseGetResponseKind string

const (
	RulesetPhaseGetResponseKindManaged RulesetPhaseGetResponseKind = "managed"
	RulesetPhaseGetResponseKindCustom  RulesetPhaseGetResponseKind = "custom"
	RulesetPhaseGetResponseKindRoot    RulesetPhaseGetResponseKind = "root"
	RulesetPhaseGetResponseKindZone    RulesetPhaseGetResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetPhaseGetResponsePhase string

const (
	RulesetPhaseGetResponsePhaseDDOSL4                         RulesetPhaseGetResponsePhase = "ddos_l4"
	RulesetPhaseGetResponsePhaseDDOSL7                         RulesetPhaseGetResponsePhase = "ddos_l7"
	RulesetPhaseGetResponsePhaseHTTPConfigSettings             RulesetPhaseGetResponsePhase = "http_config_settings"
	RulesetPhaseGetResponsePhaseHTTPCustomErrors               RulesetPhaseGetResponsePhase = "http_custom_errors"
	RulesetPhaseGetResponsePhaseHTTPLogCustomFields            RulesetPhaseGetResponsePhase = "http_log_custom_fields"
	RulesetPhaseGetResponsePhaseHTTPRatelimit                  RulesetPhaseGetResponsePhase = "http_ratelimit"
	RulesetPhaseGetResponsePhaseHTTPRequestCacheSettings       RulesetPhaseGetResponsePhase = "http_request_cache_settings"
	RulesetPhaseGetResponsePhaseHTTPRequestDynamicRedirect     RulesetPhaseGetResponsePhase = "http_request_dynamic_redirect"
	RulesetPhaseGetResponsePhaseHTTPRequestFirewallCustom      RulesetPhaseGetResponsePhase = "http_request_firewall_custom"
	RulesetPhaseGetResponsePhaseHTTPRequestFirewallManaged     RulesetPhaseGetResponsePhase = "http_request_firewall_managed"
	RulesetPhaseGetResponsePhaseHTTPRequestLateTransform       RulesetPhaseGetResponsePhase = "http_request_late_transform"
	RulesetPhaseGetResponsePhaseHTTPRequestOrigin              RulesetPhaseGetResponsePhase = "http_request_origin"
	RulesetPhaseGetResponsePhaseHTTPRequestRedirect            RulesetPhaseGetResponsePhase = "http_request_redirect"
	RulesetPhaseGetResponsePhaseHTTPRequestSanitize            RulesetPhaseGetResponsePhase = "http_request_sanitize"
	RulesetPhaseGetResponsePhaseHTTPRequestSbfm                RulesetPhaseGetResponsePhase = "http_request_sbfm"
	RulesetPhaseGetResponsePhaseHTTPRequestSelectConfiguration RulesetPhaseGetResponsePhase = "http_request_select_configuration"
	RulesetPhaseGetResponsePhaseHTTPRequestTransform           RulesetPhaseGetResponsePhase = "http_request_transform"
	RulesetPhaseGetResponsePhaseHTTPResponseCompression        RulesetPhaseGetResponsePhase = "http_response_compression"
	RulesetPhaseGetResponsePhaseHTTPResponseFirewallManaged    RulesetPhaseGetResponsePhase = "http_response_firewall_managed"
	RulesetPhaseGetResponsePhaseHTTPResponseHeadersTransform   RulesetPhaseGetResponsePhase = "http_response_headers_transform"
	RulesetPhaseGetResponsePhaseMagicTransit                   RulesetPhaseGetResponsePhase = "magic_transit"
	RulesetPhaseGetResponsePhaseMagicTransitIDsManaged         RulesetPhaseGetResponsePhase = "magic_transit_ids_managed"
	RulesetPhaseGetResponsePhaseMagicTransitManaged            RulesetPhaseGetResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetPhaseGetResponseRulesRulesetsBlockRule],
// [RulesetPhaseGetResponseRulesRulesetsExecuteRule],
// [RulesetPhaseGetResponseRulesRulesetsLogRule] or
// [RulesetPhaseGetResponseRulesRulesetsSkipRule].
type RulesetPhaseGetResponseRule interface {
	implementsRulesetPhaseGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetPhaseGetResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetPhaseGetResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetPhaseGetResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetPhaseGetResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetPhaseGetResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetPhaseGetResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetPhaseGetResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetPhaseGetResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetPhaseGetResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                            `json:"ref"`
	JSON rulesetPhaseGetResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for
// the struct [RulesetPhaseGetResponseRulesRulesetsBlockRule]
type rulesetPhaseGetResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RulesetPhaseGetResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetPhaseGetResponseRulesRulesetsBlockRule) implementsRulesetPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetPhaseGetResponseRulesRulesetsBlockRuleAction string

const (
	RulesetPhaseGetResponseRulesRulesetsBlockRuleActionBlock RulesetPhaseGetResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetPhaseGetResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetPhaseGetResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetPhaseGetResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsBlockRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetPhaseGetResponseRulesRulesetsBlockRuleActionParameters]
type rulesetPhaseGetResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response to show when the block is applied.
type RulesetPhaseGetResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                     `json:"status_code,required"`
	JSON       rulesetPhaseGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [RulesetPhaseGetResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetPhaseGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object configuring the rule's logging behavior.
type RulesetPhaseGetResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                     `json:"enabled,required"`
	JSON    rulesetPhaseGetResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetPhaseGetResponseRulesRulesetsBlockRuleLogging]
type rulesetPhaseGetResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetPhaseGetResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetPhaseGetResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetPhaseGetResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                              `json:"ref"`
	JSON rulesetPhaseGetResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata
// for the struct [RulesetPhaseGetResponseRulesRulesetsExecuteRule]
type rulesetPhaseGetResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RulesetPhaseGetResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetPhaseGetResponseRulesRulesetsExecuteRule) implementsRulesetPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetPhaseGetResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionExecute RulesetPhaseGetResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration to use for matched data logging.
type RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                         `json:"public_key,required"`
	JSON      rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A set of overrides to apply to the target ruleset.
type RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A category-level override
type RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for rules in the category.
type RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for the rule.
type RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetPhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetPhaseGetResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                       `json:"enabled,required"`
	JSON    rulesetPhaseGetResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetPhaseGetResponseRulesRulesetsExecuteRuleLogging]
type rulesetPhaseGetResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetPhaseGetResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetPhaseGetResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RulesetPhaseGetResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                          `json:"ref"`
	JSON rulesetPhaseGetResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsLogRuleJSON contains the JSON metadata for
// the struct [RulesetPhaseGetResponseRulesRulesetsLogRule]
type rulesetPhaseGetResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RulesetPhaseGetResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetPhaseGetResponseRulesRulesetsLogRule) implementsRulesetPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetPhaseGetResponseRulesRulesetsLogRuleAction string

const (
	RulesetPhaseGetResponseRulesRulesetsLogRuleActionLog RulesetPhaseGetResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetPhaseGetResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                   `json:"enabled,required"`
	JSON    rulesetPhaseGetResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsLogRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetPhaseGetResponseRulesRulesetsLogRuleLogging]
type rulesetPhaseGetResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetPhaseGetResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetPhaseGetResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetPhaseGetResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON rulesetPhaseGetResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for
// the struct [RulesetPhaseGetResponseRulesRulesetsSkipRule]
type rulesetPhaseGetResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RulesetPhaseGetResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetPhaseGetResponseRulesRulesetsSkipRule) implementsRulesetPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetPhaseGetResponseRulesRulesetsSkipRuleAction string

const (
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionSkip RulesetPhaseGetResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                         `json:"rulesets"`
	JSON     rulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParameters]
type rulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A phase to skip the execution of.
type RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetPhaseGetResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetPhaseGetResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                    `json:"enabled,required"`
	JSON    rulesetPhaseGetResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetPhaseGetResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetPhaseGetResponseRulesRulesetsSkipRuleLogging]
type rulesetPhaseGetResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseGetResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetPhaseGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id,required"`
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
	Result RulesetPhaseGetResponse `json:"result,required"`
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
