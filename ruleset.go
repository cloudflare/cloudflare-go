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
func (r *RulesetService) New(ctx context.Context, params RulesetNewParams, opts ...option.RequestOption) (res *RulesetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an account or zone ruleset, creating a new version.
func (r *RulesetService) Update(ctx context.Context, rulesetID string, params RulesetUpdateParams, opts ...option.RequestOption) (res *RulesetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all rulesets.
func (r *RulesetService) List(ctx context.Context, query RulesetListParams, opts ...option.RequestOption) (res *[]RulesetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes all versions of an existing account or zone ruleset.
func (r *RulesetService) Delete(ctx context.Context, rulesetID string, body RulesetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Fetches the latest version of an account or zone ruleset.
func (r *RulesetService) Get(ctx context.Context, rulesetID string, query RulesetGetParams, opts ...option.RequestOption) (res *RulesetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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

func (r rulesetNewResponseJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
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
	RulesetNewResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetNewResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
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

func (r rulesetNewResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
}

// A result.
type RulesetUpdateResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetUpdateResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetUpdateResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetUpdateResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                    `json:"description"`
	JSON        rulesetUpdateResponseJSON `json:"-"`
}

// rulesetUpdateResponseJSON contains the JSON metadata for the struct
// [RulesetUpdateResponse]
type rulesetUpdateResponseJSON struct {
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

func (r *RulesetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RulesetUpdateResponseKind string

const (
	RulesetUpdateResponseKindManaged RulesetUpdateResponseKind = "managed"
	RulesetUpdateResponseKindCustom  RulesetUpdateResponseKind = "custom"
	RulesetUpdateResponseKindRoot    RulesetUpdateResponseKind = "root"
	RulesetUpdateResponseKindZone    RulesetUpdateResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetUpdateResponsePhase string

const (
	RulesetUpdateResponsePhaseDDOSL4                         RulesetUpdateResponsePhase = "ddos_l4"
	RulesetUpdateResponsePhaseDDOSL7                         RulesetUpdateResponsePhase = "ddos_l7"
	RulesetUpdateResponsePhaseHTTPConfigSettings             RulesetUpdateResponsePhase = "http_config_settings"
	RulesetUpdateResponsePhaseHTTPCustomErrors               RulesetUpdateResponsePhase = "http_custom_errors"
	RulesetUpdateResponsePhaseHTTPLogCustomFields            RulesetUpdateResponsePhase = "http_log_custom_fields"
	RulesetUpdateResponsePhaseHTTPRatelimit                  RulesetUpdateResponsePhase = "http_ratelimit"
	RulesetUpdateResponsePhaseHTTPRequestCacheSettings       RulesetUpdateResponsePhase = "http_request_cache_settings"
	RulesetUpdateResponsePhaseHTTPRequestDynamicRedirect     RulesetUpdateResponsePhase = "http_request_dynamic_redirect"
	RulesetUpdateResponsePhaseHTTPRequestFirewallCustom      RulesetUpdateResponsePhase = "http_request_firewall_custom"
	RulesetUpdateResponsePhaseHTTPRequestFirewallManaged     RulesetUpdateResponsePhase = "http_request_firewall_managed"
	RulesetUpdateResponsePhaseHTTPRequestLateTransform       RulesetUpdateResponsePhase = "http_request_late_transform"
	RulesetUpdateResponsePhaseHTTPRequestOrigin              RulesetUpdateResponsePhase = "http_request_origin"
	RulesetUpdateResponsePhaseHTTPRequestRedirect            RulesetUpdateResponsePhase = "http_request_redirect"
	RulesetUpdateResponsePhaseHTTPRequestSanitize            RulesetUpdateResponsePhase = "http_request_sanitize"
	RulesetUpdateResponsePhaseHTTPRequestSbfm                RulesetUpdateResponsePhase = "http_request_sbfm"
	RulesetUpdateResponsePhaseHTTPRequestSelectConfiguration RulesetUpdateResponsePhase = "http_request_select_configuration"
	RulesetUpdateResponsePhaseHTTPRequestTransform           RulesetUpdateResponsePhase = "http_request_transform"
	RulesetUpdateResponsePhaseHTTPResponseCompression        RulesetUpdateResponsePhase = "http_response_compression"
	RulesetUpdateResponsePhaseHTTPResponseFirewallManaged    RulesetUpdateResponsePhase = "http_response_firewall_managed"
	RulesetUpdateResponsePhaseHTTPResponseHeadersTransform   RulesetUpdateResponsePhase = "http_response_headers_transform"
	RulesetUpdateResponsePhaseMagicTransit                   RulesetUpdateResponsePhase = "magic_transit"
	RulesetUpdateResponsePhaseMagicTransitIDsManaged         RulesetUpdateResponsePhase = "magic_transit_ids_managed"
	RulesetUpdateResponsePhaseMagicTransitManaged            RulesetUpdateResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetUpdateResponseRulesRulesetsBlockRule],
// [RulesetUpdateResponseRulesRulesetsExecuteRule],
// [RulesetUpdateResponseRulesRulesetsLogRule] or
// [RulesetUpdateResponseRulesRulesetsSkipRule].
type RulesetUpdateResponseRule interface {
	implementsRulesetUpdateResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetUpdateResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetUpdateResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                          `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for
// the struct [RulesetUpdateResponseRulesRulesetsBlockRule]
type rulesetUpdateResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsBlockRule) implementsRulesetUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsBlockRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsBlockRuleActionBlock RulesetUpdateResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetUpdateResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetUpdateResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsBlockRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsBlockRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type RulesetUpdateResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                   `json:"status_code,required"`
	JSON       rulesetUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON contains
// the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type RulesetUpdateResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                   `json:"enabled,required"`
	JSON    rulesetUpdateResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetUpdateResponseRulesRulesetsBlockRuleLogging]
type rulesetUpdateResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetUpdateResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                            `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata for
// the struct [RulesetUpdateResponseRulesRulesetsExecuteRule]
type rulesetUpdateResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsExecuteRule) implementsRulesetUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionExecute RulesetUpdateResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                       `json:"public_key,required"`
	JSON      rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetUpdateResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                     `json:"enabled,required"`
	JSON    rulesetUpdateResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetUpdateResponseRulesRulesetsExecuteRuleLogging]
type rulesetUpdateResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RulesetUpdateResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                        `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsLogRuleJSON contains the JSON metadata for the
// struct [RulesetUpdateResponseRulesRulesetsLogRule]
type rulesetUpdateResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsLogRule) implementsRulesetUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsLogRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsLogRuleActionLog RulesetUpdateResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetUpdateResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                 `json:"enabled,required"`
	JSON    rulesetUpdateResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsLogRuleLoggingJSON contains the JSON metadata
// for the struct [RulesetUpdateResponseRulesRulesetsLogRuleLogging]
type rulesetUpdateResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetUpdateResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                         `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for
// the struct [RulesetUpdateResponseRulesRulesetsSkipRule]
type rulesetUpdateResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSkipRule) implementsRulesetUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsSkipRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsSkipRuleActionSkip RulesetUpdateResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                       `json:"rulesets"`
	JSON     rulesetUpdateResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSkipRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSkipRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetUpdateResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                  `json:"enabled,required"`
	JSON    rulesetUpdateResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON metadata
// for the struct [RulesetUpdateResponseRulesRulesetsSkipRuleLogging]
type rulesetUpdateResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetListResponseJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
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
	RulesetGetResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetGetResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
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

func (r rulesetGetResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
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
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
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
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetNewParamsRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
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

func (r rulesetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r rulesetNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetNewResponseEnvelopeSuccess bool

const (
	RulesetNewResponseEnvelopeSuccessTrue RulesetNewResponseEnvelopeSuccess = true
)

type RulesetUpdateParams struct {
	// The unique ID of the ruleset.
	ID param.Field[string] `json:"id,required"`
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetUpdateParamsRule] `json:"rules,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[RulesetUpdateParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[RulesetUpdateParamsPhase] `json:"phase"`
}

func (r RulesetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [RulesetUpdateParamsRulesRulesetsBlockRule],
// [RulesetUpdateParamsRulesRulesetsExecuteRule],
// [RulesetUpdateParamsRulesRulesetsLogRule],
// [RulesetUpdateParamsRulesRulesetsSkipRule].
type RulesetUpdateParamsRule interface {
	implementsRulesetUpdateParamsRule()
}

type RulesetUpdateParamsRulesRulesetsBlockRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsBlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsBlockRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetUpdateParamsRulesRulesetsBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRulesRulesetsBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsBlockRule) implementsRulesetUpdateParamsRule() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsBlockRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsBlockRuleActionBlock RulesetUpdateParamsRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[RulesetUpdateParamsRulesRulesetsBlockRuleActionParametersResponse] `json:"response"`
}

func (r RulesetUpdateParamsRulesRulesetsBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type RulesetUpdateParamsRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r RulesetUpdateParamsRulesRulesetsBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type RulesetUpdateParamsRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetUpdateParamsRulesRulesetsBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetUpdateParamsRulesRulesetsExecuteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsExecuteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetUpdateParamsRulesRulesetsExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRulesRulesetsExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsExecuteRule) implementsRulesetUpdateParamsRule() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsExecuteRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionExecute RulesetUpdateParamsRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r RulesetUpdateParamsRulesRulesetsExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRule] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for rules in the category.
type RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for the rule.
type RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetUpdateParamsRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetUpdateParamsRulesRulesetsExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetUpdateParamsRulesRulesetsLogRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsLogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetUpdateParamsRulesRulesetsLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRulesRulesetsLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsLogRule) implementsRulesetUpdateParamsRule() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsLogRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsLogRuleActionLog RulesetUpdateParamsRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetUpdateParamsRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetUpdateParamsRulesRulesetsLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetUpdateParamsRulesRulesetsSkipRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsSkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsSkipRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[RulesetUpdateParamsRulesRulesetsSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRulesRulesetsSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSkipRule) implementsRulesetUpdateParamsRule() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsSkipRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsSkipRuleActionSkip RulesetUpdateParamsRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r RulesetUpdateParamsRulesRulesetsSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A phase to skip the execution of.
type RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductBic           RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductHot           RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductWAF           RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetUpdateParamsRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RulesetUpdateParamsRulesRulesetsSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type RulesetUpdateParamsKind string

const (
	RulesetUpdateParamsKindManaged RulesetUpdateParamsKind = "managed"
	RulesetUpdateParamsKindCustom  RulesetUpdateParamsKind = "custom"
	RulesetUpdateParamsKindRoot    RulesetUpdateParamsKind = "root"
	RulesetUpdateParamsKindZone    RulesetUpdateParamsKind = "zone"
)

// The phase of the ruleset.
type RulesetUpdateParamsPhase string

const (
	RulesetUpdateParamsPhaseDDOSL4                         RulesetUpdateParamsPhase = "ddos_l4"
	RulesetUpdateParamsPhaseDDOSL7                         RulesetUpdateParamsPhase = "ddos_l7"
	RulesetUpdateParamsPhaseHTTPConfigSettings             RulesetUpdateParamsPhase = "http_config_settings"
	RulesetUpdateParamsPhaseHTTPCustomErrors               RulesetUpdateParamsPhase = "http_custom_errors"
	RulesetUpdateParamsPhaseHTTPLogCustomFields            RulesetUpdateParamsPhase = "http_log_custom_fields"
	RulesetUpdateParamsPhaseHTTPRatelimit                  RulesetUpdateParamsPhase = "http_ratelimit"
	RulesetUpdateParamsPhaseHTTPRequestCacheSettings       RulesetUpdateParamsPhase = "http_request_cache_settings"
	RulesetUpdateParamsPhaseHTTPRequestDynamicRedirect     RulesetUpdateParamsPhase = "http_request_dynamic_redirect"
	RulesetUpdateParamsPhaseHTTPRequestFirewallCustom      RulesetUpdateParamsPhase = "http_request_firewall_custom"
	RulesetUpdateParamsPhaseHTTPRequestFirewallManaged     RulesetUpdateParamsPhase = "http_request_firewall_managed"
	RulesetUpdateParamsPhaseHTTPRequestLateTransform       RulesetUpdateParamsPhase = "http_request_late_transform"
	RulesetUpdateParamsPhaseHTTPRequestOrigin              RulesetUpdateParamsPhase = "http_request_origin"
	RulesetUpdateParamsPhaseHTTPRequestRedirect            RulesetUpdateParamsPhase = "http_request_redirect"
	RulesetUpdateParamsPhaseHTTPRequestSanitize            RulesetUpdateParamsPhase = "http_request_sanitize"
	RulesetUpdateParamsPhaseHTTPRequestSbfm                RulesetUpdateParamsPhase = "http_request_sbfm"
	RulesetUpdateParamsPhaseHTTPRequestSelectConfiguration RulesetUpdateParamsPhase = "http_request_select_configuration"
	RulesetUpdateParamsPhaseHTTPRequestTransform           RulesetUpdateParamsPhase = "http_request_transform"
	RulesetUpdateParamsPhaseHTTPResponseCompression        RulesetUpdateParamsPhase = "http_response_compression"
	RulesetUpdateParamsPhaseHTTPResponseFirewallManaged    RulesetUpdateParamsPhase = "http_response_firewall_managed"
	RulesetUpdateParamsPhaseHTTPResponseHeadersTransform   RulesetUpdateParamsPhase = "http_response_headers_transform"
	RulesetUpdateParamsPhaseMagicTransit                   RulesetUpdateParamsPhase = "magic_transit"
	RulesetUpdateParamsPhaseMagicTransitIDsManaged         RulesetUpdateParamsPhase = "magic_transit_ids_managed"
	RulesetUpdateParamsPhaseMagicTransitManaged            RulesetUpdateParamsPhase = "magic_transit_managed"
)

// A response object.
type RulesetUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetUpdateResponseEnvelopeJSON    `json:"-"`
}

// rulesetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelope]
type rulesetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetUpdateResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetUpdateResponseEnvelopeErrors]
type rulesetUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetUpdateResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                        `json:"pointer,required"`
	JSON    rulesetUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RulesetUpdateResponseEnvelopeErrorsSource]
type rulesetUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetUpdateResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetUpdateResponseEnvelopeMessages]
type rulesetUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetUpdateResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    rulesetUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RulesetUpdateResponseEnvelopeMessagesSource]
type rulesetUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetUpdateResponseEnvelopeSuccess bool

const (
	RulesetUpdateResponseEnvelopeSuccessTrue RulesetUpdateResponseEnvelopeSuccess = true
)

type RulesetListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

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

func (r rulesetListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r rulesetListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r rulesetListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
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

func (r rulesetListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r rulesetListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetListResponseEnvelopeSuccess bool

const (
	RulesetListResponseEnvelopeSuccessTrue RulesetListResponseEnvelopeSuccess = true
)

type RulesetDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type RulesetGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

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

func (r rulesetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r rulesetGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetGetResponseEnvelopeSuccess bool

const (
	RulesetGetResponseEnvelopeSuccessTrue RulesetGetResponseEnvelopeSuccess = true
)
