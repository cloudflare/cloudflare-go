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
func (r *RulesetPhaseService) Update(ctx context.Context, rulesetPhase RulesetPhaseUpdateParamsRulesetPhase, params RulesetPhaseUpdateParams, opts ...option.RequestOption) (res *RulesetPhaseUpdateResponse, err error) {
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
type RulesetPhaseUpdateResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetPhaseUpdateResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetPhaseUpdateResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetPhaseUpdateResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                         `json:"description"`
	JSON        rulesetPhaseUpdateResponseJSON `json:"-"`
}

// rulesetPhaseUpdateResponseJSON contains the JSON metadata for the struct
// [RulesetPhaseUpdateResponse]
type rulesetPhaseUpdateResponseJSON struct {
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

func (r *RulesetPhaseUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetPhaseUpdateResponseKind string

const (
	RulesetPhaseUpdateResponseKindManaged RulesetPhaseUpdateResponseKind = "managed"
	RulesetPhaseUpdateResponseKindCustom  RulesetPhaseUpdateResponseKind = "custom"
	RulesetPhaseUpdateResponseKindRoot    RulesetPhaseUpdateResponseKind = "root"
	RulesetPhaseUpdateResponseKindZone    RulesetPhaseUpdateResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetPhaseUpdateResponsePhase string

const (
	RulesetPhaseUpdateResponsePhaseDDOSL4                         RulesetPhaseUpdateResponsePhase = "ddos_l4"
	RulesetPhaseUpdateResponsePhaseDDOSL7                         RulesetPhaseUpdateResponsePhase = "ddos_l7"
	RulesetPhaseUpdateResponsePhaseHTTPConfigSettings             RulesetPhaseUpdateResponsePhase = "http_config_settings"
	RulesetPhaseUpdateResponsePhaseHTTPCustomErrors               RulesetPhaseUpdateResponsePhase = "http_custom_errors"
	RulesetPhaseUpdateResponsePhaseHTTPLogCustomFields            RulesetPhaseUpdateResponsePhase = "http_log_custom_fields"
	RulesetPhaseUpdateResponsePhaseHTTPRatelimit                  RulesetPhaseUpdateResponsePhase = "http_ratelimit"
	RulesetPhaseUpdateResponsePhaseHTTPRequestCacheSettings       RulesetPhaseUpdateResponsePhase = "http_request_cache_settings"
	RulesetPhaseUpdateResponsePhaseHTTPRequestDynamicRedirect     RulesetPhaseUpdateResponsePhase = "http_request_dynamic_redirect"
	RulesetPhaseUpdateResponsePhaseHTTPRequestFirewallCustom      RulesetPhaseUpdateResponsePhase = "http_request_firewall_custom"
	RulesetPhaseUpdateResponsePhaseHTTPRequestFirewallManaged     RulesetPhaseUpdateResponsePhase = "http_request_firewall_managed"
	RulesetPhaseUpdateResponsePhaseHTTPRequestLateTransform       RulesetPhaseUpdateResponsePhase = "http_request_late_transform"
	RulesetPhaseUpdateResponsePhaseHTTPRequestOrigin              RulesetPhaseUpdateResponsePhase = "http_request_origin"
	RulesetPhaseUpdateResponsePhaseHTTPRequestRedirect            RulesetPhaseUpdateResponsePhase = "http_request_redirect"
	RulesetPhaseUpdateResponsePhaseHTTPRequestSanitize            RulesetPhaseUpdateResponsePhase = "http_request_sanitize"
	RulesetPhaseUpdateResponsePhaseHTTPRequestSbfm                RulesetPhaseUpdateResponsePhase = "http_request_sbfm"
	RulesetPhaseUpdateResponsePhaseHTTPRequestSelectConfiguration RulesetPhaseUpdateResponsePhase = "http_request_select_configuration"
	RulesetPhaseUpdateResponsePhaseHTTPRequestTransform           RulesetPhaseUpdateResponsePhase = "http_request_transform"
	RulesetPhaseUpdateResponsePhaseHTTPResponseCompression        RulesetPhaseUpdateResponsePhase = "http_response_compression"
	RulesetPhaseUpdateResponsePhaseHTTPResponseFirewallManaged    RulesetPhaseUpdateResponsePhase = "http_response_firewall_managed"
	RulesetPhaseUpdateResponsePhaseHTTPResponseHeadersTransform   RulesetPhaseUpdateResponsePhase = "http_response_headers_transform"
	RulesetPhaseUpdateResponsePhaseMagicTransit                   RulesetPhaseUpdateResponsePhase = "magic_transit"
	RulesetPhaseUpdateResponsePhaseMagicTransitIDsManaged         RulesetPhaseUpdateResponsePhase = "magic_transit_ids_managed"
	RulesetPhaseUpdateResponsePhaseMagicTransitManaged            RulesetPhaseUpdateResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetPhaseUpdateResponseRulesRulesetsBlockRule],
// [RulesetPhaseUpdateResponseRulesRulesetsExecuteRule],
// [RulesetPhaseUpdateResponseRulesRulesetsLogRule] or
// [RulesetPhaseUpdateResponseRulesRulesetsSkipRule].
type RulesetPhaseUpdateResponseRule interface {
	implementsRulesetPhaseUpdateResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetPhaseUpdateResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetPhaseUpdateResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetPhaseUpdateResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetPhaseUpdateResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetPhaseUpdateResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetPhaseUpdateResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetPhaseUpdateResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetPhaseUpdateResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                               `json:"ref"`
	JSON rulesetPhaseUpdateResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsBlockRuleJSON contains the JSON metadata
// for the struct [RulesetPhaseUpdateResponseRulesRulesetsBlockRule]
type rulesetPhaseUpdateResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RulesetPhaseUpdateResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetPhaseUpdateResponseRulesRulesetsBlockRule) implementsRulesetPhaseUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetPhaseUpdateResponseRulesRulesetsBlockRuleAction string

const (
	RulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionBlock RulesetPhaseUpdateResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParameters]
type rulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response to show when the block is applied.
type RulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                        `json:"status_code,required"`
	JSON       rulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [RulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object configuring the rule's logging behavior.
type RulesetPhaseUpdateResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                        `json:"enabled,required"`
	JSON    rulesetPhaseUpdateResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON
// metadata for the struct
// [RulesetPhaseUpdateResponseRulesRulesetsBlockRuleLogging]
type rulesetPhaseUpdateResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetPhaseUpdateResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                 `json:"ref"`
	JSON rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleJSON contains the JSON
// metadata for the struct [RulesetPhaseUpdateResponseRulesRulesetsExecuteRule]
type rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RulesetPhaseUpdateResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetPhaseUpdateResponseRulesRulesetsExecuteRule) implementsRulesetPhaseUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionExecute RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration to use for matched data logging.
type RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                            `json:"public_key,required"`
	JSON      rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A set of overrides to apply to the target ruleset.
type RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A category-level override
type RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for rules in the category.
type RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for the rule.
type RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                          `json:"enabled,required"`
	JSON    rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON
// metadata for the struct
// [RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleLogging]
type rulesetPhaseUpdateResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetPhaseUpdateResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetPhaseUpdateResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RulesetPhaseUpdateResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                             `json:"ref"`
	JSON rulesetPhaseUpdateResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsLogRuleJSON contains the JSON metadata
// for the struct [RulesetPhaseUpdateResponseRulesRulesetsLogRule]
type rulesetPhaseUpdateResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RulesetPhaseUpdateResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetPhaseUpdateResponseRulesRulesetsLogRule) implementsRulesetPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetPhaseUpdateResponseRulesRulesetsLogRuleAction string

const (
	RulesetPhaseUpdateResponseRulesRulesetsLogRuleActionLog RulesetPhaseUpdateResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetPhaseUpdateResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                      `json:"enabled,required"`
	JSON    rulesetPhaseUpdateResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsLogRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetPhaseUpdateResponseRulesRulesetsLogRuleLogging]
type rulesetPhaseUpdateResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetPhaseUpdateResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetPhaseUpdateResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetPhaseUpdateResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                              `json:"ref"`
	JSON rulesetPhaseUpdateResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsSkipRuleJSON contains the JSON metadata
// for the struct [RulesetPhaseUpdateResponseRulesRulesetsSkipRule]
type rulesetPhaseUpdateResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RulesetPhaseUpdateResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetPhaseUpdateResponseRulesRulesetsSkipRule) implementsRulesetPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetPhaseUpdateResponseRulesRulesetsSkipRuleAction string

const (
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionSkip RulesetPhaseUpdateResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                            `json:"rulesets"`
	JSON     rulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParameters]
type rulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A phase to skip the execution of.
type RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetPhaseUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetPhaseUpdateResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                       `json:"enabled,required"`
	JSON    rulesetPhaseUpdateResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetPhaseUpdateResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetPhaseUpdateResponseRulesRulesetsSkipRuleLogging]
type rulesetPhaseUpdateResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseUpdateResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type RulesetPhaseUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The unique ID of the ruleset.
	ID param.Field[string] `json:"id,required"`
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetPhaseUpdateParamsRule] `json:"rules,required"`
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
	Result RulesetPhaseUpdateResponse `json:"result,required"`
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
