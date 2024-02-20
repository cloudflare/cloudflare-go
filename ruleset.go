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

// RulesetService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRulesetService] method instead.
type RulesetService struct {
	Options  []option.RequestOption
	Phases   *RulesetPhaseService
	Rules    *RulesetRuleService
	Versions *RulesetVersionService
}

// NewRulesetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRulesetService(opts ...option.RequestOption) (r *RulesetService) {
	r = &RulesetService{}
	r.Options = opts
	r.Phases = NewRulesetPhaseService(opts...)
	r.Rules = NewRulesetRuleService(opts...)
	r.Versions = NewRulesetVersionService(opts...)
	return
}

// Creates a ruleset.
func (r *RulesetService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, body RulesetNewParams, opts ...option.RequestOption) (res *RulesetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetNewResponseEnvelope
	path := fmt.Sprintf("%s/%s/rulesets", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all rulesets.
func (r *RulesetService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]RulesetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetListResponseEnvelope
	path := fmt.Sprintf("%s/%s/rulesets", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes all versions of an existing account or zone ruleset.
func (r *RulesetService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, rulesetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches the latest version of an account or zone ruleset.
func (r *RulesetService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, rulesetID string, opts ...option.RequestOption) (res *RulesetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetGetResponseEnvelope
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an account or zone ruleset, creating a new version.
func (r *RulesetService) Replace(ctx context.Context, accountOrZone string, accountOrZoneID string, rulesetID string, body RulesetReplaceParams, opts ...option.RequestOption) (res *RulesetReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetReplaceResponseEnvelope
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A result.
type RulesetNewResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetNewResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetNewResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetNewResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                 `json:"description"`
	JSON        rulesetNewResponseJSON `json:"-"`
}

// rulesetNewResponseJSON contains the JSON metadata for the struct
// [RulesetNewResponse]
type rulesetNewResponseJSON struct {
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

func (r *RulesetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetNewResponseKind string

const (
	RulesetNewResponseKindManaged RulesetNewResponseKind = "managed"
	RulesetNewResponseKindCustom  RulesetNewResponseKind = "custom"
	RulesetNewResponseKindRoot    RulesetNewResponseKind = "root"
	RulesetNewResponseKindZone    RulesetNewResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetNewResponsePhase string

const (
	RulesetNewResponsePhaseDDOSL4                         RulesetNewResponsePhase = "ddos_l4"
	RulesetNewResponsePhaseDDOSL7                         RulesetNewResponsePhase = "ddos_l7"
	RulesetNewResponsePhaseHTTPConfigSettings             RulesetNewResponsePhase = "http_config_settings"
	RulesetNewResponsePhaseHTTPCustomErrors               RulesetNewResponsePhase = "http_custom_errors"
	RulesetNewResponsePhaseHTTPLogCustomFields            RulesetNewResponsePhase = "http_log_custom_fields"
	RulesetNewResponsePhaseHTTPRatelimit                  RulesetNewResponsePhase = "http_ratelimit"
	RulesetNewResponsePhaseHTTPRequestCacheSettings       RulesetNewResponsePhase = "http_request_cache_settings"
	RulesetNewResponsePhaseHTTPRequestDynamicRedirect     RulesetNewResponsePhase = "http_request_dynamic_redirect"
	RulesetNewResponsePhaseHTTPRequestFirewallCustom      RulesetNewResponsePhase = "http_request_firewall_custom"
	RulesetNewResponsePhaseHTTPRequestFirewallManaged     RulesetNewResponsePhase = "http_request_firewall_managed"
	RulesetNewResponsePhaseHTTPRequestLateTransform       RulesetNewResponsePhase = "http_request_late_transform"
	RulesetNewResponsePhaseHTTPRequestOrigin              RulesetNewResponsePhase = "http_request_origin"
	RulesetNewResponsePhaseHTTPRequestRedirect            RulesetNewResponsePhase = "http_request_redirect"
	RulesetNewResponsePhaseHTTPRequestSanitize            RulesetNewResponsePhase = "http_request_sanitize"
	RulesetNewResponsePhaseHTTPRequestSbfm                RulesetNewResponsePhase = "http_request_sbfm"
	RulesetNewResponsePhaseHTTPRequestSelectConfiguration RulesetNewResponsePhase = "http_request_select_configuration"
	RulesetNewResponsePhaseHTTPRequestTransform           RulesetNewResponsePhase = "http_request_transform"
	RulesetNewResponsePhaseHTTPResponseCompression        RulesetNewResponsePhase = "http_response_compression"
	RulesetNewResponsePhaseHTTPResponseFirewallManaged    RulesetNewResponsePhase = "http_response_firewall_managed"
	RulesetNewResponsePhaseHTTPResponseHeadersTransform   RulesetNewResponsePhase = "http_response_headers_transform"
	RulesetNewResponsePhaseMagicTransit                   RulesetNewResponsePhase = "magic_transit"
	RulesetNewResponsePhaseMagicTransitIDsManaged         RulesetNewResponsePhase = "magic_transit_ids_managed"
	RulesetNewResponsePhaseMagicTransitManaged            RulesetNewResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetNewResponseRulesRulesetsBlockRule],
// [RulesetNewResponseRulesRulesetsExecuteRule],
// [RulesetNewResponseRulesRulesetsLogRule] or
// [RulesetNewResponseRulesRulesetsSkipRule].
type RulesetNewResponseRule interface {
	implementsRulesetNewResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetNewResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetNewResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                       `json:"ref"`
	JSON rulesetNewResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for the
// struct [RulesetNewResponseRulesRulesetsBlockRule]
type rulesetNewResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetNewResponseRulesRulesetsBlockRule) implementsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsBlockRuleAction string

const (
	RulesetNewResponseRulesRulesetsBlockRuleActionBlock RulesetNewResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetNewResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetNewResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetNewResponseRulesRulesetsBlockRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetNewResponseRulesRulesetsBlockRuleActionParameters]
type rulesetNewResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response to show when the block is applied.
type RulesetNewResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                `json:"status_code,required"`
	JSON       rulesetNewResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsBlockRuleActionParametersResponseJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetNewResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object configuring the rule's logging behavior.
type RulesetNewResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                `json:"enabled,required"`
	JSON    rulesetNewResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON metadata
// for the struct [RulesetNewResponseRulesRulesetsBlockRuleLogging]
type rulesetNewResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetNewResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetNewResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                         `json:"ref"`
	JSON rulesetNewResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata for
// the struct [RulesetNewResponseRulesRulesetsExecuteRule]
type rulesetNewResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetNewResponseRulesRulesetsExecuteRule) implementsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetNewResponseRulesRulesetsExecuteRuleActionExecute RulesetNewResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetNewResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetNewResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetNewResponseRulesRulesetsExecuteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetNewResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetNewResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration to use for matched data logging.
type RulesetNewResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                    `json:"public_key,required"`
	JSON      rulesetNewResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetNewResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A set of overrides to apply to the target ruleset.
type RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A category-level override
type RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for rules in the category.
type RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for the rule.
type RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetNewResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                  `json:"enabled,required"`
	JSON    rulesetNewResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON metadata
// for the struct [RulesetNewResponseRulesRulesetsExecuteRuleLogging]
type rulesetNewResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetNewResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RulesetNewResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                     `json:"ref"`
	JSON rulesetNewResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsLogRuleJSON contains the JSON metadata for the
// struct [RulesetNewResponseRulesRulesetsLogRule]
type rulesetNewResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetNewResponseRulesRulesetsLogRule) implementsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsLogRuleAction string

const (
	RulesetNewResponseRulesRulesetsLogRuleActionLog RulesetNewResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetNewResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                              `json:"enabled,required"`
	JSON    rulesetNewResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsLogRuleLoggingJSON contains the JSON metadata for
// the struct [RulesetNewResponseRulesRulesetsLogRuleLogging]
type rulesetNewResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetNewResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetNewResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                      `json:"ref"`
	JSON rulesetNewResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for the
// struct [RulesetNewResponseRulesRulesetsSkipRule]
type rulesetNewResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetNewResponseRulesRulesetsSkipRule) implementsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsSkipRuleAction string

const (
	RulesetNewResponseRulesRulesetsSkipRuleActionSkip RulesetNewResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetNewResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetNewResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                    `json:"rulesets"`
	JSON     rulesetNewResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSkipRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetNewResponseRulesRulesetsSkipRuleActionParameters]
type rulesetNewResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A phase to skip the execution of.
type RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetNewResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetNewResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetNewResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetNewResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetNewResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetNewResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersProductUaBlock       RulesetNewResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetNewResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetNewResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetNewResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetNewResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetNewResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                               `json:"enabled,required"`
	JSON    rulesetNewResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON metadata
// for the struct [RulesetNewResponseRulesRulesetsSkipRuleLogging]
type rulesetNewResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ruleset object.
type RulesetListResponse struct {
	// The kind of the ruleset.
	Kind RulesetListResponseKind `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetListResponsePhase `json:"phase,required"`
	// The unique ID of the ruleset.
	ID string `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the ruleset.
	Version string                  `json:"version"`
	JSON    rulesetListResponseJSON `json:"-"`
}

// rulesetListResponseJSON contains the JSON metadata for the struct
// [RulesetListResponse]
type rulesetListResponseJSON struct {
	Kind        apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	ID          apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetListResponseKind string

const (
	RulesetListResponseKindManaged RulesetListResponseKind = "managed"
	RulesetListResponseKindCustom  RulesetListResponseKind = "custom"
	RulesetListResponseKindRoot    RulesetListResponseKind = "root"
	RulesetListResponseKindZone    RulesetListResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetListResponsePhase string

const (
	RulesetListResponsePhaseDDOSL4                         RulesetListResponsePhase = "ddos_l4"
	RulesetListResponsePhaseDDOSL7                         RulesetListResponsePhase = "ddos_l7"
	RulesetListResponsePhaseHTTPConfigSettings             RulesetListResponsePhase = "http_config_settings"
	RulesetListResponsePhaseHTTPCustomErrors               RulesetListResponsePhase = "http_custom_errors"
	RulesetListResponsePhaseHTTPLogCustomFields            RulesetListResponsePhase = "http_log_custom_fields"
	RulesetListResponsePhaseHTTPRatelimit                  RulesetListResponsePhase = "http_ratelimit"
	RulesetListResponsePhaseHTTPRequestCacheSettings       RulesetListResponsePhase = "http_request_cache_settings"
	RulesetListResponsePhaseHTTPRequestDynamicRedirect     RulesetListResponsePhase = "http_request_dynamic_redirect"
	RulesetListResponsePhaseHTTPRequestFirewallCustom      RulesetListResponsePhase = "http_request_firewall_custom"
	RulesetListResponsePhaseHTTPRequestFirewallManaged     RulesetListResponsePhase = "http_request_firewall_managed"
	RulesetListResponsePhaseHTTPRequestLateTransform       RulesetListResponsePhase = "http_request_late_transform"
	RulesetListResponsePhaseHTTPRequestOrigin              RulesetListResponsePhase = "http_request_origin"
	RulesetListResponsePhaseHTTPRequestRedirect            RulesetListResponsePhase = "http_request_redirect"
	RulesetListResponsePhaseHTTPRequestSanitize            RulesetListResponsePhase = "http_request_sanitize"
	RulesetListResponsePhaseHTTPRequestSbfm                RulesetListResponsePhase = "http_request_sbfm"
	RulesetListResponsePhaseHTTPRequestSelectConfiguration RulesetListResponsePhase = "http_request_select_configuration"
	RulesetListResponsePhaseHTTPRequestTransform           RulesetListResponsePhase = "http_request_transform"
	RulesetListResponsePhaseHTTPResponseCompression        RulesetListResponsePhase = "http_response_compression"
	RulesetListResponsePhaseHTTPResponseFirewallManaged    RulesetListResponsePhase = "http_response_firewall_managed"
	RulesetListResponsePhaseHTTPResponseHeadersTransform   RulesetListResponsePhase = "http_response_headers_transform"
	RulesetListResponsePhaseMagicTransit                   RulesetListResponsePhase = "magic_transit"
	RulesetListResponsePhaseMagicTransitIDsManaged         RulesetListResponsePhase = "magic_transit_ids_managed"
	RulesetListResponsePhaseMagicTransitManaged            RulesetListResponsePhase = "magic_transit_managed"
)

// A result.
type RulesetGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                 `json:"description"`
	JSON        rulesetGetResponseJSON `json:"-"`
}

// rulesetGetResponseJSON contains the JSON metadata for the struct
// [RulesetGetResponse]
type rulesetGetResponseJSON struct {
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

func (r *RulesetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetGetResponseKind string

const (
	RulesetGetResponseKindManaged RulesetGetResponseKind = "managed"
	RulesetGetResponseKindCustom  RulesetGetResponseKind = "custom"
	RulesetGetResponseKindRoot    RulesetGetResponseKind = "root"
	RulesetGetResponseKindZone    RulesetGetResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetGetResponsePhase string

const (
	RulesetGetResponsePhaseDDOSL4                         RulesetGetResponsePhase = "ddos_l4"
	RulesetGetResponsePhaseDDOSL7                         RulesetGetResponsePhase = "ddos_l7"
	RulesetGetResponsePhaseHTTPConfigSettings             RulesetGetResponsePhase = "http_config_settings"
	RulesetGetResponsePhaseHTTPCustomErrors               RulesetGetResponsePhase = "http_custom_errors"
	RulesetGetResponsePhaseHTTPLogCustomFields            RulesetGetResponsePhase = "http_log_custom_fields"
	RulesetGetResponsePhaseHTTPRatelimit                  RulesetGetResponsePhase = "http_ratelimit"
	RulesetGetResponsePhaseHTTPRequestCacheSettings       RulesetGetResponsePhase = "http_request_cache_settings"
	RulesetGetResponsePhaseHTTPRequestDynamicRedirect     RulesetGetResponsePhase = "http_request_dynamic_redirect"
	RulesetGetResponsePhaseHTTPRequestFirewallCustom      RulesetGetResponsePhase = "http_request_firewall_custom"
	RulesetGetResponsePhaseHTTPRequestFirewallManaged     RulesetGetResponsePhase = "http_request_firewall_managed"
	RulesetGetResponsePhaseHTTPRequestLateTransform       RulesetGetResponsePhase = "http_request_late_transform"
	RulesetGetResponsePhaseHTTPRequestOrigin              RulesetGetResponsePhase = "http_request_origin"
	RulesetGetResponsePhaseHTTPRequestRedirect            RulesetGetResponsePhase = "http_request_redirect"
	RulesetGetResponsePhaseHTTPRequestSanitize            RulesetGetResponsePhase = "http_request_sanitize"
	RulesetGetResponsePhaseHTTPRequestSbfm                RulesetGetResponsePhase = "http_request_sbfm"
	RulesetGetResponsePhaseHTTPRequestSelectConfiguration RulesetGetResponsePhase = "http_request_select_configuration"
	RulesetGetResponsePhaseHTTPRequestTransform           RulesetGetResponsePhase = "http_request_transform"
	RulesetGetResponsePhaseHTTPResponseCompression        RulesetGetResponsePhase = "http_response_compression"
	RulesetGetResponsePhaseHTTPResponseFirewallManaged    RulesetGetResponsePhase = "http_response_firewall_managed"
	RulesetGetResponsePhaseHTTPResponseHeadersTransform   RulesetGetResponsePhase = "http_response_headers_transform"
	RulesetGetResponsePhaseMagicTransit                   RulesetGetResponsePhase = "magic_transit"
	RulesetGetResponsePhaseMagicTransitIDsManaged         RulesetGetResponsePhase = "magic_transit_ids_managed"
	RulesetGetResponsePhaseMagicTransitManaged            RulesetGetResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetGetResponseRulesRulesetsBlockRule],
// [RulesetGetResponseRulesRulesetsExecuteRule],
// [RulesetGetResponseRulesRulesetsLogRule] or
// [RulesetGetResponseRulesRulesetsSkipRule].
type RulesetGetResponseRule interface {
	implementsRulesetGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetGetResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetGetResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                       `json:"ref"`
	JSON rulesetGetResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for the
// struct [RulesetGetResponseRulesRulesetsBlockRule]
type rulesetGetResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetGetResponseRulesRulesetsBlockRule) implementsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsBlockRuleAction string

const (
	RulesetGetResponseRulesRulesetsBlockRuleActionBlock RulesetGetResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetGetResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetGetResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetGetResponseRulesRulesetsBlockRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetGetResponseRulesRulesetsBlockRuleActionParameters]
type rulesetGetResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response to show when the block is applied.
type RulesetGetResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                `json:"status_code,required"`
	JSON       rulesetGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object configuring the rule's logging behavior.
type RulesetGetResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                `json:"enabled,required"`
	JSON    rulesetGetResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON metadata
// for the struct [RulesetGetResponseRulesRulesetsBlockRuleLogging]
type rulesetGetResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetGetResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetGetResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                         `json:"ref"`
	JSON rulesetGetResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata for
// the struct [RulesetGetResponseRulesRulesetsExecuteRule]
type rulesetGetResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetGetResponseRulesRulesetsExecuteRule) implementsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetGetResponseRulesRulesetsExecuteRuleActionExecute RulesetGetResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetGetResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetGetResponseRulesRulesetsExecuteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetGetResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetGetResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration to use for matched data logging.
type RulesetGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                    `json:"public_key,required"`
	JSON      rulesetGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A set of overrides to apply to the target ruleset.
type RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A category-level override
type RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for rules in the category.
type RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for the rule.
type RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetGetResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                  `json:"enabled,required"`
	JSON    rulesetGetResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON metadata
// for the struct [RulesetGetResponseRulesRulesetsExecuteRuleLogging]
type rulesetGetResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetGetResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RulesetGetResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                     `json:"ref"`
	JSON rulesetGetResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsLogRuleJSON contains the JSON metadata for the
// struct [RulesetGetResponseRulesRulesetsLogRule]
type rulesetGetResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetGetResponseRulesRulesetsLogRule) implementsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsLogRuleAction string

const (
	RulesetGetResponseRulesRulesetsLogRuleActionLog RulesetGetResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetGetResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                              `json:"enabled,required"`
	JSON    rulesetGetResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsLogRuleLoggingJSON contains the JSON metadata for
// the struct [RulesetGetResponseRulesRulesetsLogRuleLogging]
type rulesetGetResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetGetResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetGetResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                      `json:"ref"`
	JSON rulesetGetResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for the
// struct [RulesetGetResponseRulesRulesetsSkipRule]
type rulesetGetResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetGetResponseRulesRulesetsSkipRule) implementsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsSkipRuleAction string

const (
	RulesetGetResponseRulesRulesetsSkipRuleActionSkip RulesetGetResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetGetResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetGetResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                    `json:"rulesets"`
	JSON     rulesetGetResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSkipRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetGetResponseRulesRulesetsSkipRuleActionParameters]
type rulesetGetResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A phase to skip the execution of.
type RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetGetResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetGetResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetGetResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetGetResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetGetResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersProductUaBlock       RulesetGetResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetGetResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetGetResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetGetResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetGetResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetGetResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                               `json:"enabled,required"`
	JSON    rulesetGetResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON metadata
// for the struct [RulesetGetResponseRulesRulesetsSkipRuleLogging]
type rulesetGetResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A result.
type RulesetReplaceResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetReplaceResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetReplaceResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetReplaceResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                     `json:"description"`
	JSON        rulesetReplaceResponseJSON `json:"-"`
}

// rulesetReplaceResponseJSON contains the JSON metadata for the struct
// [RulesetReplaceResponse]
type rulesetReplaceResponseJSON struct {
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

func (r *RulesetReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetReplaceResponseKind string

const (
	RulesetReplaceResponseKindManaged RulesetReplaceResponseKind = "managed"
	RulesetReplaceResponseKindCustom  RulesetReplaceResponseKind = "custom"
	RulesetReplaceResponseKindRoot    RulesetReplaceResponseKind = "root"
	RulesetReplaceResponseKindZone    RulesetReplaceResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetReplaceResponsePhase string

const (
	RulesetReplaceResponsePhaseDDOSL4                         RulesetReplaceResponsePhase = "ddos_l4"
	RulesetReplaceResponsePhaseDDOSL7                         RulesetReplaceResponsePhase = "ddos_l7"
	RulesetReplaceResponsePhaseHTTPConfigSettings             RulesetReplaceResponsePhase = "http_config_settings"
	RulesetReplaceResponsePhaseHTTPCustomErrors               RulesetReplaceResponsePhase = "http_custom_errors"
	RulesetReplaceResponsePhaseHTTPLogCustomFields            RulesetReplaceResponsePhase = "http_log_custom_fields"
	RulesetReplaceResponsePhaseHTTPRatelimit                  RulesetReplaceResponsePhase = "http_ratelimit"
	RulesetReplaceResponsePhaseHTTPRequestCacheSettings       RulesetReplaceResponsePhase = "http_request_cache_settings"
	RulesetReplaceResponsePhaseHTTPRequestDynamicRedirect     RulesetReplaceResponsePhase = "http_request_dynamic_redirect"
	RulesetReplaceResponsePhaseHTTPRequestFirewallCustom      RulesetReplaceResponsePhase = "http_request_firewall_custom"
	RulesetReplaceResponsePhaseHTTPRequestFirewallManaged     RulesetReplaceResponsePhase = "http_request_firewall_managed"
	RulesetReplaceResponsePhaseHTTPRequestLateTransform       RulesetReplaceResponsePhase = "http_request_late_transform"
	RulesetReplaceResponsePhaseHTTPRequestOrigin              RulesetReplaceResponsePhase = "http_request_origin"
	RulesetReplaceResponsePhaseHTTPRequestRedirect            RulesetReplaceResponsePhase = "http_request_redirect"
	RulesetReplaceResponsePhaseHTTPRequestSanitize            RulesetReplaceResponsePhase = "http_request_sanitize"
	RulesetReplaceResponsePhaseHTTPRequestSbfm                RulesetReplaceResponsePhase = "http_request_sbfm"
	RulesetReplaceResponsePhaseHTTPRequestSelectConfiguration RulesetReplaceResponsePhase = "http_request_select_configuration"
	RulesetReplaceResponsePhaseHTTPRequestTransform           RulesetReplaceResponsePhase = "http_request_transform"
	RulesetReplaceResponsePhaseHTTPResponseCompression        RulesetReplaceResponsePhase = "http_response_compression"
	RulesetReplaceResponsePhaseHTTPResponseFirewallManaged    RulesetReplaceResponsePhase = "http_response_firewall_managed"
	RulesetReplaceResponsePhaseHTTPResponseHeadersTransform   RulesetReplaceResponsePhase = "http_response_headers_transform"
	RulesetReplaceResponsePhaseMagicTransit                   RulesetReplaceResponsePhase = "magic_transit"
	RulesetReplaceResponsePhaseMagicTransitIDsManaged         RulesetReplaceResponsePhase = "magic_transit_ids_managed"
	RulesetReplaceResponsePhaseMagicTransitManaged            RulesetReplaceResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetReplaceResponseRulesRulesetsBlockRule],
// [RulesetReplaceResponseRulesRulesetsExecuteRule],
// [RulesetReplaceResponseRulesRulesetsLogRule] or
// [RulesetReplaceResponseRulesRulesetsSkipRule].
type RulesetReplaceResponseRule interface {
	implementsRulesetReplaceResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetReplaceResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetReplaceResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetReplaceResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetReplaceResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetReplaceResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetReplaceResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetReplaceResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetReplaceResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetReplaceResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON rulesetReplaceResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for
// the struct [RulesetReplaceResponseRulesRulesetsBlockRule]
type rulesetReplaceResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RulesetReplaceResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetReplaceResponseRulesRulesetsBlockRule) implementsRulesetReplaceResponseRule() {}

// The action to perform when the rule matches.
type RulesetReplaceResponseRulesRulesetsBlockRuleAction string

const (
	RulesetReplaceResponseRulesRulesetsBlockRuleActionBlock RulesetReplaceResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetReplaceResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetReplaceResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetReplaceResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsBlockRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetReplaceResponseRulesRulesetsBlockRuleActionParameters]
type rulesetReplaceResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response to show when the block is applied.
type RulesetReplaceResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                    `json:"status_code,required"`
	JSON       rulesetReplaceResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [RulesetReplaceResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetReplaceResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object configuring the rule's logging behavior.
type RulesetReplaceResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                    `json:"enabled,required"`
	JSON    rulesetReplaceResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetReplaceResponseRulesRulesetsBlockRuleLogging]
type rulesetReplaceResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetReplaceResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetReplaceResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetReplaceResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetReplaceResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                             `json:"ref"`
	JSON rulesetReplaceResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata
// for the struct [RulesetReplaceResponseRulesRulesetsExecuteRule]
type rulesetReplaceResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RulesetReplaceResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetReplaceResponseRulesRulesetsExecuteRule) implementsRulesetReplaceResponseRule() {}

// The action to perform when the rule matches.
type RulesetReplaceResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionExecute RulesetReplaceResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetReplaceResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetReplaceResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration to use for matched data logging.
type RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                        `json:"public_key,required"`
	JSON      rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A set of overrides to apply to the target ruleset.
type RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A category-level override
type RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for rules in the category.
type RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for the rule.
type RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetReplaceResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetReplaceResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                      `json:"enabled,required"`
	JSON    rulesetReplaceResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetReplaceResponseRulesRulesetsExecuteRuleLogging]
type rulesetReplaceResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetReplaceResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetReplaceResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RulesetReplaceResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                         `json:"ref"`
	JSON rulesetReplaceResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsLogRuleJSON contains the JSON metadata for
// the struct [RulesetReplaceResponseRulesRulesetsLogRule]
type rulesetReplaceResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RulesetReplaceResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetReplaceResponseRulesRulesetsLogRule) implementsRulesetReplaceResponseRule() {}

// The action to perform when the rule matches.
type RulesetReplaceResponseRulesRulesetsLogRuleAction string

const (
	RulesetReplaceResponseRulesRulesetsLogRuleActionLog RulesetReplaceResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetReplaceResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                  `json:"enabled,required"`
	JSON    rulesetReplaceResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsLogRuleLoggingJSON contains the JSON metadata
// for the struct [RulesetReplaceResponseRulesRulesetsLogRuleLogging]
type rulesetReplaceResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetReplaceResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetReplaceResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetReplaceResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetReplaceResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                          `json:"ref"`
	JSON rulesetReplaceResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for
// the struct [RulesetReplaceResponseRulesRulesetsSkipRule]
type rulesetReplaceResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RulesetReplaceResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetReplaceResponseRulesRulesetsSkipRule) implementsRulesetReplaceResponseRule() {}

// The action to perform when the rule matches.
type RulesetReplaceResponseRulesRulesetsSkipRuleAction string

const (
	RulesetReplaceResponseRulesRulesetsSkipRuleActionSkip RulesetReplaceResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetReplaceResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                        `json:"rulesets"`
	JSON     rulesetReplaceResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsSkipRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetReplaceResponseRulesRulesetsSkipRuleActionParameters]
type rulesetReplaceResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A phase to skip the execution of.
type RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProductUaBlock       RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetReplaceResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetReplaceResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                   `json:"enabled,required"`
	JSON    rulesetReplaceResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetReplaceResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetReplaceResponseRulesRulesetsSkipRuleLogging]
type rulesetReplaceResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetNewParams struct {
	// The kind of the ruleset.
	Kind param.Field[RulesetNewParamsKind] `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name,required"`
	// The phase of the ruleset.
	Phase param.Field[RulesetNewParamsPhase] `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetNewParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r RulesetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type RulesetNewParamsKind string

const (
	RulesetNewParamsKindManaged RulesetNewParamsKind = "managed"
	RulesetNewParamsKindCustom  RulesetNewParamsKind = "custom"
	RulesetNewParamsKindRoot    RulesetNewParamsKind = "root"
	RulesetNewParamsKindZone    RulesetNewParamsKind = "zone"
)

// The phase of the ruleset.
type RulesetNewParamsPhase string

const (
	RulesetNewParamsPhaseDDOSL4                         RulesetNewParamsPhase = "ddos_l4"
	RulesetNewParamsPhaseDDOSL7                         RulesetNewParamsPhase = "ddos_l7"
	RulesetNewParamsPhaseHTTPConfigSettings             RulesetNewParamsPhase = "http_config_settings"
	RulesetNewParamsPhaseHTTPCustomErrors               RulesetNewParamsPhase = "http_custom_errors"
	RulesetNewParamsPhaseHTTPLogCustomFields            RulesetNewParamsPhase = "http_log_custom_fields"
	RulesetNewParamsPhaseHTTPRatelimit                  RulesetNewParamsPhase = "http_ratelimit"
	RulesetNewParamsPhaseHTTPRequestCacheSettings       RulesetNewParamsPhase = "http_request_cache_settings"
	RulesetNewParamsPhaseHTTPRequestDynamicRedirect     RulesetNewParamsPhase = "http_request_dynamic_redirect"
	RulesetNewParamsPhaseHTTPRequestFirewallCustom      RulesetNewParamsPhase = "http_request_firewall_custom"
	RulesetNewParamsPhaseHTTPRequestFirewallManaged     RulesetNewParamsPhase = "http_request_firewall_managed"
	RulesetNewParamsPhaseHTTPRequestLateTransform       RulesetNewParamsPhase = "http_request_late_transform"
	RulesetNewParamsPhaseHTTPRequestOrigin              RulesetNewParamsPhase = "http_request_origin"
	RulesetNewParamsPhaseHTTPRequestRedirect            RulesetNewParamsPhase = "http_request_redirect"
	RulesetNewParamsPhaseHTTPRequestSanitize            RulesetNewParamsPhase = "http_request_sanitize"
	RulesetNewParamsPhaseHTTPRequestSbfm                RulesetNewParamsPhase = "http_request_sbfm"
	RulesetNewParamsPhaseHTTPRequestSelectConfiguration RulesetNewParamsPhase = "http_request_select_configuration"
	RulesetNewParamsPhaseHTTPRequestTransform           RulesetNewParamsPhase = "http_request_transform"
	RulesetNewParamsPhaseHTTPResponseCompression        RulesetNewParamsPhase = "http_response_compression"
	RulesetNewParamsPhaseHTTPResponseFirewallManaged    RulesetNewParamsPhase = "http_response_firewall_managed"
	RulesetNewParamsPhaseHTTPResponseHeadersTransform   RulesetNewParamsPhase = "http_response_headers_transform"
	RulesetNewParamsPhaseMagicTransit                   RulesetNewParamsPhase = "magic_transit"
	RulesetNewParamsPhaseMagicTransitIDsManaged         RulesetNewParamsPhase = "magic_transit_ids_managed"
	RulesetNewParamsPhaseMagicTransitManaged            RulesetNewParamsPhase = "magic_transit_managed"
)

// Satisfied by [RulesetNewParamsRulesRulesetsBlockRule],
// [RulesetNewParamsRulesRulesetsExecuteRule],
// [RulesetNewParamsRulesRulesetsLogRule], [RulesetNewParamsRulesRulesetsSkipRule].
type RulesetNewParamsRule interface {
	implementsRulesetNewParamsRule()
}

type RulesetNewParamsRulesRulesetsBlockRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsBlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsBlockRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetNewParamsRulesRulesetsBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRulesRulesetsBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsBlockRule) implementsRulesetNewParamsRule() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsBlockRuleAction string

const (
	RulesetNewParamsRulesRulesetsBlockRuleActionBlock RulesetNewParamsRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[RulesetNewParamsRulesRulesetsBlockRuleActionParametersResponse] `json:"response"`
}

func (r RulesetNewParamsRulesRulesetsBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type RulesetNewParamsRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r RulesetNewParamsRulesRulesetsBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type RulesetNewParamsRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetNewParamsRulesRulesetsBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetNewParamsRulesRulesetsExecuteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsExecuteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetNewParamsRulesRulesetsExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRulesRulesetsExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsExecuteRule) implementsRulesetNewParamsRule() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsExecuteRuleAction string

const (
	RulesetNewParamsRulesRulesetsExecuteRuleActionExecute RulesetNewParamsRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[RulesetNewParamsRulesRulesetsExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r RulesetNewParamsRulesRulesetsExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type RulesetNewParamsRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r RulesetNewParamsRulesRulesetsExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRule] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for rules in the category.
type RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for the rule.
type RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetNewParamsRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetNewParamsRulesRulesetsExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetNewParamsRulesRulesetsLogRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsLogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetNewParamsRulesRulesetsLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRulesRulesetsLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsLogRule) implementsRulesetNewParamsRule() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsLogRuleAction string

const (
	RulesetNewParamsRulesRulesetsLogRuleActionLog RulesetNewParamsRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetNewParamsRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetNewParamsRulesRulesetsLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetNewParamsRulesRulesetsSkipRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsSkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsSkipRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetNewParamsRulesRulesetsSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRulesRulesetsSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSkipRule) implementsRulesetNewParamsRule() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsSkipRuleAction string

const (
	RulesetNewParamsRulesRulesetsSkipRuleActionSkip RulesetNewParamsRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]RulesetNewParamsRulesRulesetsSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[RulesetNewParamsRulesRulesetsSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r RulesetNewParamsRulesRulesetsSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A phase to skip the execution of.
type RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetNewParamsRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductBic           RulesetNewParamsRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductHot           RulesetNewParamsRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetNewParamsRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetNewParamsRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductUaBlock       RulesetNewParamsRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductWAF           RulesetNewParamsRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetNewParamsRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetNewParamsRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetNewParamsRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetNewParamsRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetNewParamsRulesRulesetsSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A response object.
type RulesetNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetNewResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetNewResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetNewResponseEnvelopeJSON    `json:"-"`
}

// rulesetNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetNewResponseEnvelope]
type rulesetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetNewResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetNewResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeErrors]
type rulesetNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetNewResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    rulesetNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RulesetNewResponseEnvelopeErrorsSource]
type rulesetNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetNewResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetNewResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeMessages]
type rulesetNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetNewResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    rulesetNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RulesetNewResponseEnvelopeMessagesSource]
type rulesetNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetNewResponseEnvelopeSuccess bool

const (
	RulesetNewResponseEnvelopeSuccessTrue RulesetNewResponseEnvelopeSuccess = true
)

// A response object.
type RulesetListResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetListResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetListResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result []RulesetListResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetListResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetListResponseEnvelopeJSON    `json:"-"`
}

// rulesetListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetListResponseEnvelope]
type rulesetListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetListResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetListResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetListResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RulesetListResponseEnvelopeErrors]
type rulesetListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetListResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                      `json:"pointer,required"`
	JSON    rulesetListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetListResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RulesetListResponseEnvelopeErrorsSource]
type rulesetListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetListResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetListResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetListResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetListResponseEnvelopeMessages]
type rulesetListResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetListResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                        `json:"pointer,required"`
	JSON    rulesetListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetListResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RulesetListResponseEnvelopeMessagesSource]
type rulesetListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetListResponseEnvelopeSuccess bool

const (
	RulesetListResponseEnvelopeSuccessTrue RulesetListResponseEnvelopeSuccess = true
)

// A response object.
type RulesetGetResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetGetResponseEnvelopeJSON    `json:"-"`
}

// rulesetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetGetResponseEnvelope]
type rulesetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RulesetGetResponseEnvelopeErrors]
type rulesetGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    rulesetGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RulesetGetResponseEnvelopeErrorsSource]
type rulesetGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RulesetGetResponseEnvelopeMessages]
type rulesetGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    rulesetGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RulesetGetResponseEnvelopeMessagesSource]
type rulesetGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetGetResponseEnvelopeSuccess bool

const (
	RulesetGetResponseEnvelopeSuccessTrue RulesetGetResponseEnvelopeSuccess = true
)

type RulesetReplaceParams struct {
	// The unique ID of the ruleset.
	ID param.Field[string] `json:"id,required"`
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetReplaceParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[RulesetReplaceParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[RulesetReplaceParamsPhase] `json:"phase"`
}

func (r RulesetReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [RulesetReplaceParamsRulesRulesetsBlockRule],
// [RulesetReplaceParamsRulesRulesetsExecuteRule],
// [RulesetReplaceParamsRulesRulesetsLogRule],
// [RulesetReplaceParamsRulesRulesetsSkipRule].
type RulesetReplaceParamsRule interface {
	implementsRulesetReplaceParamsRule()
}

type RulesetReplaceParamsRulesRulesetsBlockRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetReplaceParamsRulesRulesetsBlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetReplaceParamsRulesRulesetsBlockRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetReplaceParamsRulesRulesetsBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetReplaceParamsRulesRulesetsBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetReplaceParamsRulesRulesetsBlockRule) implementsRulesetReplaceParamsRule() {}

// The action to perform when the rule matches.
type RulesetReplaceParamsRulesRulesetsBlockRuleAction string

const (
	RulesetReplaceParamsRulesRulesetsBlockRuleActionBlock RulesetReplaceParamsRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetReplaceParamsRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[RulesetReplaceParamsRulesRulesetsBlockRuleActionParametersResponse] `json:"response"`
}

func (r RulesetReplaceParamsRulesRulesetsBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type RulesetReplaceParamsRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r RulesetReplaceParamsRulesRulesetsBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type RulesetReplaceParamsRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetReplaceParamsRulesRulesetsBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetReplaceParamsRulesRulesetsExecuteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetReplaceParamsRulesRulesetsExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetReplaceParamsRulesRulesetsExecuteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetReplaceParamsRulesRulesetsExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetReplaceParamsRulesRulesetsExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetReplaceParamsRulesRulesetsExecuteRule) implementsRulesetReplaceParamsRule() {}

// The action to perform when the rule matches.
type RulesetReplaceParamsRulesRulesetsExecuteRuleAction string

const (
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionExecute RulesetReplaceParamsRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetReplaceParamsRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r RulesetReplaceParamsRulesRulesetsExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRule] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for rules in the category.
type RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for the rule.
type RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetReplaceParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetReplaceParamsRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetReplaceParamsRulesRulesetsExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetReplaceParamsRulesRulesetsLogRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetReplaceParamsRulesRulesetsLogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetReplaceParamsRulesRulesetsLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetReplaceParamsRulesRulesetsLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetReplaceParamsRulesRulesetsLogRule) implementsRulesetReplaceParamsRule() {}

// The action to perform when the rule matches.
type RulesetReplaceParamsRulesRulesetsLogRuleAction string

const (
	RulesetReplaceParamsRulesRulesetsLogRuleActionLog RulesetReplaceParamsRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetReplaceParamsRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetReplaceParamsRulesRulesetsLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetReplaceParamsRulesRulesetsSkipRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetReplaceParamsRulesRulesetsSkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetReplaceParamsRulesRulesetsSkipRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetReplaceParamsRulesRulesetsSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetReplaceParamsRulesRulesetsSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetReplaceParamsRulesRulesetsSkipRule) implementsRulesetReplaceParamsRule() {}

// The action to perform when the rule matches.
type RulesetReplaceParamsRulesRulesetsSkipRuleAction string

const (
	RulesetReplaceParamsRulesRulesetsSkipRuleActionSkip RulesetReplaceParamsRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetReplaceParamsRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r RulesetReplaceParamsRulesRulesetsSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A phase to skip the execution of.
type RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProductBic           RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProductHot           RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProductUaBlock       RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProductWAF           RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetReplaceParamsRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetReplaceParamsRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetReplaceParamsRulesRulesetsSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type RulesetReplaceParamsKind string

const (
	RulesetReplaceParamsKindManaged RulesetReplaceParamsKind = "managed"
	RulesetReplaceParamsKindCustom  RulesetReplaceParamsKind = "custom"
	RulesetReplaceParamsKindRoot    RulesetReplaceParamsKind = "root"
	RulesetReplaceParamsKindZone    RulesetReplaceParamsKind = "zone"
)

// The phase of the ruleset.
type RulesetReplaceParamsPhase string

const (
	RulesetReplaceParamsPhaseDDOSL4                         RulesetReplaceParamsPhase = "ddos_l4"
	RulesetReplaceParamsPhaseDDOSL7                         RulesetReplaceParamsPhase = "ddos_l7"
	RulesetReplaceParamsPhaseHTTPConfigSettings             RulesetReplaceParamsPhase = "http_config_settings"
	RulesetReplaceParamsPhaseHTTPCustomErrors               RulesetReplaceParamsPhase = "http_custom_errors"
	RulesetReplaceParamsPhaseHTTPLogCustomFields            RulesetReplaceParamsPhase = "http_log_custom_fields"
	RulesetReplaceParamsPhaseHTTPRatelimit                  RulesetReplaceParamsPhase = "http_ratelimit"
	RulesetReplaceParamsPhaseHTTPRequestCacheSettings       RulesetReplaceParamsPhase = "http_request_cache_settings"
	RulesetReplaceParamsPhaseHTTPRequestDynamicRedirect     RulesetReplaceParamsPhase = "http_request_dynamic_redirect"
	RulesetReplaceParamsPhaseHTTPRequestFirewallCustom      RulesetReplaceParamsPhase = "http_request_firewall_custom"
	RulesetReplaceParamsPhaseHTTPRequestFirewallManaged     RulesetReplaceParamsPhase = "http_request_firewall_managed"
	RulesetReplaceParamsPhaseHTTPRequestLateTransform       RulesetReplaceParamsPhase = "http_request_late_transform"
	RulesetReplaceParamsPhaseHTTPRequestOrigin              RulesetReplaceParamsPhase = "http_request_origin"
	RulesetReplaceParamsPhaseHTTPRequestRedirect            RulesetReplaceParamsPhase = "http_request_redirect"
	RulesetReplaceParamsPhaseHTTPRequestSanitize            RulesetReplaceParamsPhase = "http_request_sanitize"
	RulesetReplaceParamsPhaseHTTPRequestSbfm                RulesetReplaceParamsPhase = "http_request_sbfm"
	RulesetReplaceParamsPhaseHTTPRequestSelectConfiguration RulesetReplaceParamsPhase = "http_request_select_configuration"
	RulesetReplaceParamsPhaseHTTPRequestTransform           RulesetReplaceParamsPhase = "http_request_transform"
	RulesetReplaceParamsPhaseHTTPResponseCompression        RulesetReplaceParamsPhase = "http_response_compression"
	RulesetReplaceParamsPhaseHTTPResponseFirewallManaged    RulesetReplaceParamsPhase = "http_response_firewall_managed"
	RulesetReplaceParamsPhaseHTTPResponseHeadersTransform   RulesetReplaceParamsPhase = "http_response_headers_transform"
	RulesetReplaceParamsPhaseMagicTransit                   RulesetReplaceParamsPhase = "magic_transit"
	RulesetReplaceParamsPhaseMagicTransitIDsManaged         RulesetReplaceParamsPhase = "magic_transit_ids_managed"
	RulesetReplaceParamsPhaseMagicTransitManaged            RulesetReplaceParamsPhase = "magic_transit_managed"
)

// A response object.
type RulesetReplaceResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetReplaceResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetReplaceResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetReplaceResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetReplaceResponseEnvelopeJSON    `json:"-"`
}

// rulesetReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetReplaceResponseEnvelope]
type rulesetReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetReplaceResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetReplaceResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetReplaceResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetReplaceResponseEnvelopeErrors]
type rulesetReplaceResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetReplaceResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                         `json:"pointer,required"`
	JSON    rulesetReplaceResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetReplaceResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RulesetReplaceResponseEnvelopeErrorsSource]
type rulesetReplaceResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetReplaceResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetReplaceResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetReplaceResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetReplaceResponseEnvelopeMessages]
type rulesetReplaceResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetReplaceResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                           `json:"pointer,required"`
	JSON    rulesetReplaceResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetReplaceResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RulesetReplaceResponseEnvelopeMessagesSource]
type rulesetReplaceResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetReplaceResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetReplaceResponseEnvelopeSuccess bool

const (
	RulesetReplaceResponseEnvelopeSuccessTrue RulesetReplaceResponseEnvelopeSuccess = true
)
