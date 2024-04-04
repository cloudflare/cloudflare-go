// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
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
func (r *PhaseService) Update(ctx context.Context, rulesetPhase PhaseUpdateParamsRulesetPhase, params PhaseUpdateParams, opts ...option.RequestOption) (res *PhaseUpdateResponse, err error) {
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
func (r *PhaseService) Get(ctx context.Context, rulesetPhase PhaseGetParamsRulesetPhase, query PhaseGetParams, opts ...option.RequestOption) (res *PhaseGetResponse, err error) {
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

// A ruleset object.
type PhaseUpdateResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind PhaseUpdateResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase PhaseUpdateResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []PhaseUpdateResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                  `json:"description"`
	JSON        phaseUpdateResponseJSON `json:"-"`
}

// phaseUpdateResponseJSON contains the JSON metadata for the struct
// [PhaseUpdateResponse]
type phaseUpdateResponseJSON struct {
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

func (r *PhaseUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type PhaseUpdateResponseKind string

const (
	PhaseUpdateResponseKindManaged PhaseUpdateResponseKind = "managed"
	PhaseUpdateResponseKindCustom  PhaseUpdateResponseKind = "custom"
	PhaseUpdateResponseKindRoot    PhaseUpdateResponseKind = "root"
	PhaseUpdateResponseKindZone    PhaseUpdateResponseKind = "zone"
)

func (r PhaseUpdateResponseKind) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseKindManaged, PhaseUpdateResponseKindCustom, PhaseUpdateResponseKindRoot, PhaseUpdateResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type PhaseUpdateResponsePhase string

const (
	PhaseUpdateResponsePhaseDDoSL4                         PhaseUpdateResponsePhase = "ddos_l4"
	PhaseUpdateResponsePhaseDDoSL7                         PhaseUpdateResponsePhase = "ddos_l7"
	PhaseUpdateResponsePhaseHTTPConfigSettings             PhaseUpdateResponsePhase = "http_config_settings"
	PhaseUpdateResponsePhaseHTTPCustomErrors               PhaseUpdateResponsePhase = "http_custom_errors"
	PhaseUpdateResponsePhaseHTTPLogCustomFields            PhaseUpdateResponsePhase = "http_log_custom_fields"
	PhaseUpdateResponsePhaseHTTPRatelimit                  PhaseUpdateResponsePhase = "http_ratelimit"
	PhaseUpdateResponsePhaseHTTPRequestCacheSettings       PhaseUpdateResponsePhase = "http_request_cache_settings"
	PhaseUpdateResponsePhaseHTTPRequestDynamicRedirect     PhaseUpdateResponsePhase = "http_request_dynamic_redirect"
	PhaseUpdateResponsePhaseHTTPRequestFirewallCustom      PhaseUpdateResponsePhase = "http_request_firewall_custom"
	PhaseUpdateResponsePhaseHTTPRequestFirewallManaged     PhaseUpdateResponsePhase = "http_request_firewall_managed"
	PhaseUpdateResponsePhaseHTTPRequestLateTransform       PhaseUpdateResponsePhase = "http_request_late_transform"
	PhaseUpdateResponsePhaseHTTPRequestOrigin              PhaseUpdateResponsePhase = "http_request_origin"
	PhaseUpdateResponsePhaseHTTPRequestRedirect            PhaseUpdateResponsePhase = "http_request_redirect"
	PhaseUpdateResponsePhaseHTTPRequestSanitize            PhaseUpdateResponsePhase = "http_request_sanitize"
	PhaseUpdateResponsePhaseHTTPRequestSbfm                PhaseUpdateResponsePhase = "http_request_sbfm"
	PhaseUpdateResponsePhaseHTTPRequestSelectConfiguration PhaseUpdateResponsePhase = "http_request_select_configuration"
	PhaseUpdateResponsePhaseHTTPRequestTransform           PhaseUpdateResponsePhase = "http_request_transform"
	PhaseUpdateResponsePhaseHTTPResponseCompression        PhaseUpdateResponsePhase = "http_response_compression"
	PhaseUpdateResponsePhaseHTTPResponseFirewallManaged    PhaseUpdateResponsePhase = "http_response_firewall_managed"
	PhaseUpdateResponsePhaseHTTPResponseHeadersTransform   PhaseUpdateResponsePhase = "http_response_headers_transform"
	PhaseUpdateResponsePhaseMagicTransit                   PhaseUpdateResponsePhase = "magic_transit"
	PhaseUpdateResponsePhaseMagicTransitIDsManaged         PhaseUpdateResponsePhase = "magic_transit_ids_managed"
	PhaseUpdateResponsePhaseMagicTransitManaged            PhaseUpdateResponsePhase = "magic_transit_managed"
)

func (r PhaseUpdateResponsePhase) IsKnown() bool {
	switch r {
	case PhaseUpdateResponsePhaseDDoSL4, PhaseUpdateResponsePhaseDDoSL7, PhaseUpdateResponsePhaseHTTPConfigSettings, PhaseUpdateResponsePhaseHTTPCustomErrors, PhaseUpdateResponsePhaseHTTPLogCustomFields, PhaseUpdateResponsePhaseHTTPRatelimit, PhaseUpdateResponsePhaseHTTPRequestCacheSettings, PhaseUpdateResponsePhaseHTTPRequestDynamicRedirect, PhaseUpdateResponsePhaseHTTPRequestFirewallCustom, PhaseUpdateResponsePhaseHTTPRequestFirewallManaged, PhaseUpdateResponsePhaseHTTPRequestLateTransform, PhaseUpdateResponsePhaseHTTPRequestOrigin, PhaseUpdateResponsePhaseHTTPRequestRedirect, PhaseUpdateResponsePhaseHTTPRequestSanitize, PhaseUpdateResponsePhaseHTTPRequestSbfm, PhaseUpdateResponsePhaseHTTPRequestSelectConfiguration, PhaseUpdateResponsePhaseHTTPRequestTransform, PhaseUpdateResponsePhaseHTTPResponseCompression, PhaseUpdateResponsePhaseHTTPResponseFirewallManaged, PhaseUpdateResponsePhaseHTTPResponseHeadersTransform, PhaseUpdateResponsePhaseMagicTransit, PhaseUpdateResponsePhaseMagicTransitIDsManaged, PhaseUpdateResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

// Union satisfied by [rulesets.PhaseUpdateResponseRulesRulesetsBlockRule],
// [rulesets.PhaseUpdateResponseRulesRulesetsExecuteRule],
// [rulesets.PhaseUpdateResponseRulesRulesetsLogRule] or
// [rulesets.PhaseUpdateResponseRulesRulesetsSkipRule].
type PhaseUpdateResponseRule interface {
	implementsRulesetsPhaseUpdateResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type PhaseUpdateResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging shared.Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                        `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseRulesRulesetsBlockRule]
type phaseUpdateResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsBlockRule) implementsRulesetsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsBlockRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsBlockRuleActionBlock PhaseUpdateResponseRulesRulesetsBlockRuleAction = "block"
)

func (r PhaseUpdateResponseRulesRulesetsBlockRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsBlockRuleActionBlock:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response PhaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     phaseUpdateResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// phaseUpdateResponseRulesRulesetsBlockRuleActionParametersJSON contains the JSON
// metadata for the struct
// [PhaseUpdateResponseRulesRulesetsBlockRuleActionParameters]
type phaseUpdateResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type PhaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                 `json:"status_code,required"`
	JSON       phaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON contains
// the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponse]
type phaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging shared.Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                          `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata for
// the struct [PhaseUpdateResponseRulesRulesetsExecuteRule]
type phaseUpdateResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsExecuteRule) implementsRulesetsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsExecuteRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionExecute PhaseUpdateResponseRulesRulesetsExecuteRuleAction = "execute"
)

func (r PhaseUpdateResponseRulesRulesetsExecuteRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsExecuteRuleActionExecute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsExecuteRuleActionParameters]
type phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                     `json:"public_key,required"`
	JSON      phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

func (r PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

// A rule-level override
type PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

func (r PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault, PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium, PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow, PhaseUpdateResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

type PhaseUpdateResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging shared.Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                      `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsLogRuleJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseRulesRulesetsLogRule]
type phaseUpdateResponseRulesRulesetsLogRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsLogRule) implementsRulesetsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsLogRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsLogRuleActionLog PhaseUpdateResponseRulesRulesetsLogRuleAction = "log"
)

func (r PhaseUpdateResponseRulesRulesetsLogRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsLogRuleActionLog:
		return true
	}
	return false
}

type PhaseUpdateResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging shared.Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                       `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseRulesRulesetsSkipRule]
type phaseUpdateResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSkipRule) implementsRulesetsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsSkipRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsSkipRuleActionSkip PhaseUpdateResponseRulesRulesetsSkipRuleAction = "skip"
)

func (r PhaseUpdateResponseRulesRulesetsSkipRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSkipRuleActionSkip:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                     `json:"rulesets"`
	JSON     phaseUpdateResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSkipRuleActionParametersJSON contains the JSON
// metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSkipRuleActionParameters]
type phaseUpdateResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL4                         PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL7                         PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

func (r PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL4, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL7, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

// The name of a legacy security product to skip the execution of.
type PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductBic           PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductHot           PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductUABlock       PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductWAF           PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

func (r PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductBic, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductHot, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductRateLimit, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductUABlock, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductWAF, PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

func (r PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

// A ruleset object.
type PhaseGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind PhaseGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase PhaseGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []PhaseGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string               `json:"description"`
	JSON        phaseGetResponseJSON `json:"-"`
}

// phaseGetResponseJSON contains the JSON metadata for the struct
// [PhaseGetResponse]
type phaseGetResponseJSON struct {
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

func (r *PhaseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type PhaseGetResponseKind string

const (
	PhaseGetResponseKindManaged PhaseGetResponseKind = "managed"
	PhaseGetResponseKindCustom  PhaseGetResponseKind = "custom"
	PhaseGetResponseKindRoot    PhaseGetResponseKind = "root"
	PhaseGetResponseKindZone    PhaseGetResponseKind = "zone"
)

func (r PhaseGetResponseKind) IsKnown() bool {
	switch r {
	case PhaseGetResponseKindManaged, PhaseGetResponseKindCustom, PhaseGetResponseKindRoot, PhaseGetResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type PhaseGetResponsePhase string

const (
	PhaseGetResponsePhaseDDoSL4                         PhaseGetResponsePhase = "ddos_l4"
	PhaseGetResponsePhaseDDoSL7                         PhaseGetResponsePhase = "ddos_l7"
	PhaseGetResponsePhaseHTTPConfigSettings             PhaseGetResponsePhase = "http_config_settings"
	PhaseGetResponsePhaseHTTPCustomErrors               PhaseGetResponsePhase = "http_custom_errors"
	PhaseGetResponsePhaseHTTPLogCustomFields            PhaseGetResponsePhase = "http_log_custom_fields"
	PhaseGetResponsePhaseHTTPRatelimit                  PhaseGetResponsePhase = "http_ratelimit"
	PhaseGetResponsePhaseHTTPRequestCacheSettings       PhaseGetResponsePhase = "http_request_cache_settings"
	PhaseGetResponsePhaseHTTPRequestDynamicRedirect     PhaseGetResponsePhase = "http_request_dynamic_redirect"
	PhaseGetResponsePhaseHTTPRequestFirewallCustom      PhaseGetResponsePhase = "http_request_firewall_custom"
	PhaseGetResponsePhaseHTTPRequestFirewallManaged     PhaseGetResponsePhase = "http_request_firewall_managed"
	PhaseGetResponsePhaseHTTPRequestLateTransform       PhaseGetResponsePhase = "http_request_late_transform"
	PhaseGetResponsePhaseHTTPRequestOrigin              PhaseGetResponsePhase = "http_request_origin"
	PhaseGetResponsePhaseHTTPRequestRedirect            PhaseGetResponsePhase = "http_request_redirect"
	PhaseGetResponsePhaseHTTPRequestSanitize            PhaseGetResponsePhase = "http_request_sanitize"
	PhaseGetResponsePhaseHTTPRequestSbfm                PhaseGetResponsePhase = "http_request_sbfm"
	PhaseGetResponsePhaseHTTPRequestSelectConfiguration PhaseGetResponsePhase = "http_request_select_configuration"
	PhaseGetResponsePhaseHTTPRequestTransform           PhaseGetResponsePhase = "http_request_transform"
	PhaseGetResponsePhaseHTTPResponseCompression        PhaseGetResponsePhase = "http_response_compression"
	PhaseGetResponsePhaseHTTPResponseFirewallManaged    PhaseGetResponsePhase = "http_response_firewall_managed"
	PhaseGetResponsePhaseHTTPResponseHeadersTransform   PhaseGetResponsePhase = "http_response_headers_transform"
	PhaseGetResponsePhaseMagicTransit                   PhaseGetResponsePhase = "magic_transit"
	PhaseGetResponsePhaseMagicTransitIDsManaged         PhaseGetResponsePhase = "magic_transit_ids_managed"
	PhaseGetResponsePhaseMagicTransitManaged            PhaseGetResponsePhase = "magic_transit_managed"
)

func (r PhaseGetResponsePhase) IsKnown() bool {
	switch r {
	case PhaseGetResponsePhaseDDoSL4, PhaseGetResponsePhaseDDoSL7, PhaseGetResponsePhaseHTTPConfigSettings, PhaseGetResponsePhaseHTTPCustomErrors, PhaseGetResponsePhaseHTTPLogCustomFields, PhaseGetResponsePhaseHTTPRatelimit, PhaseGetResponsePhaseHTTPRequestCacheSettings, PhaseGetResponsePhaseHTTPRequestDynamicRedirect, PhaseGetResponsePhaseHTTPRequestFirewallCustom, PhaseGetResponsePhaseHTTPRequestFirewallManaged, PhaseGetResponsePhaseHTTPRequestLateTransform, PhaseGetResponsePhaseHTTPRequestOrigin, PhaseGetResponsePhaseHTTPRequestRedirect, PhaseGetResponsePhaseHTTPRequestSanitize, PhaseGetResponsePhaseHTTPRequestSbfm, PhaseGetResponsePhaseHTTPRequestSelectConfiguration, PhaseGetResponsePhaseHTTPRequestTransform, PhaseGetResponsePhaseHTTPResponseCompression, PhaseGetResponsePhaseHTTPResponseFirewallManaged, PhaseGetResponsePhaseHTTPResponseHeadersTransform, PhaseGetResponsePhaseMagicTransit, PhaseGetResponsePhaseMagicTransitIDsManaged, PhaseGetResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

// Union satisfied by [rulesets.PhaseGetResponseRulesRulesetsBlockRule],
// [rulesets.PhaseGetResponseRulesRulesetsExecuteRule],
// [rulesets.PhaseGetResponseRulesRulesetsLogRule] or
// [rulesets.PhaseGetResponseRulesRulesetsSkipRule].
type PhaseGetResponseRule interface {
	implementsRulesetsPhaseGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type PhaseGetResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging shared.Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                     `json:"ref"`
	JSON phaseGetResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for the
// struct [PhaseGetResponseRulesRulesetsBlockRule]
type phaseGetResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsBlockRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsBlockRuleAction string

const (
	PhaseGetResponseRulesRulesetsBlockRuleActionBlock PhaseGetResponseRulesRulesetsBlockRuleAction = "block"
)

func (r PhaseGetResponseRulesRulesetsBlockRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsBlockRuleActionBlock:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response PhaseGetResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     phaseGetResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// phaseGetResponseRulesRulesetsBlockRuleActionParametersJSON contains the JSON
// metadata for the struct [PhaseGetResponseRulesRulesetsBlockRuleActionParameters]
type phaseGetResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type PhaseGetResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                              `json:"status_code,required"`
	JSON       phaseGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON contains the
// JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsBlockRuleActionParametersResponse]
type phaseGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

type PhaseGetResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging shared.Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                       `json:"ref"`
	JSON phaseGetResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata for the
// struct [PhaseGetResponseRulesRulesetsExecuteRule]
type phaseGetResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsExecuteRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsExecuteRuleAction string

const (
	PhaseGetResponseRulesRulesetsExecuteRuleActionExecute PhaseGetResponseRulesRulesetsExecuteRuleAction = "execute"
)

func (r PhaseGetResponseRulesRulesetsExecuteRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsExecuteRuleActionExecute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData PhaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      phaseGetResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// phaseGetResponseRulesRulesetsExecuteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [PhaseGetResponseRulesRulesetsExecuteRuleActionParameters]
type phaseGetResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type PhaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                  `json:"public_key,required"`
	JSON      phaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type phaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             phaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// phaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type phaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             phaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// phaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type phaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

func (r PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

// A rule-level override
type PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             phaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// phaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type phaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

func (r PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault, PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium, PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow, PhaseGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

type PhaseGetResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging shared.Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                   `json:"ref"`
	JSON phaseGetResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsLogRuleJSON contains the JSON metadata for the
// struct [PhaseGetResponseRulesRulesetsLogRule]
type phaseGetResponseRulesRulesetsLogRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsLogRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsLogRuleAction string

const (
	PhaseGetResponseRulesRulesetsLogRuleActionLog PhaseGetResponseRulesRulesetsLogRuleAction = "log"
)

func (r PhaseGetResponseRulesRulesetsLogRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsLogRuleActionLog:
		return true
	}
	return false
}

type PhaseGetResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging shared.Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                    `json:"ref"`
	JSON phaseGetResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for the
// struct [PhaseGetResponseRulesRulesetsSkipRule]
type phaseGetResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSkipRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsSkipRuleAction string

const (
	PhaseGetResponseRulesRulesetsSkipRuleActionSkip PhaseGetResponseRulesRulesetsSkipRuleAction = "skip"
)

func (r PhaseGetResponseRulesRulesetsSkipRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSkipRuleActionSkip:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []PhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset PhaseGetResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                  `json:"rulesets"`
	JSON     phaseGetResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSkipRuleActionParametersJSON contains the JSON
// metadata for the struct [PhaseGetResponseRulesRulesetsSkipRuleActionParameters]
type phaseGetResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL4                         PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL7                         PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

func (r PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL4, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL7, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged, PhaseGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

// The name of a legacy security product to skip the execution of.
type PhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductBic           PhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductHot           PhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     PhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel PhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductUABlock       PhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductWAF           PhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  PhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

func (r PhaseGetResponseRulesRulesetsSkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductBic, PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductHot, PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit, PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel, PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductUABlock, PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductWAF, PhaseGetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type PhaseGetResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	PhaseGetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent PhaseGetResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

func (r PhaseGetResponseRulesRulesetsSkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

type PhaseUpdateParams struct {
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
	PhaseUpdateParamsRulesetPhaseDDoSL4                         PhaseUpdateParamsRulesetPhase = "ddos_l4"
	PhaseUpdateParamsRulesetPhaseDDoSL7                         PhaseUpdateParamsRulesetPhase = "ddos_l7"
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

func (r PhaseUpdateParamsRulesetPhase) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesetPhaseDDoSL4, PhaseUpdateParamsRulesetPhaseDDoSL7, PhaseUpdateParamsRulesetPhaseHTTPConfigSettings, PhaseUpdateParamsRulesetPhaseHTTPCustomErrors, PhaseUpdateParamsRulesetPhaseHTTPLogCustomFields, PhaseUpdateParamsRulesetPhaseHTTPRatelimit, PhaseUpdateParamsRulesetPhaseHTTPRequestCacheSettings, PhaseUpdateParamsRulesetPhaseHTTPRequestDynamicRedirect, PhaseUpdateParamsRulesetPhaseHTTPRequestFirewallCustom, PhaseUpdateParamsRulesetPhaseHTTPRequestFirewallManaged, PhaseUpdateParamsRulesetPhaseHTTPRequestLateTransform, PhaseUpdateParamsRulesetPhaseHTTPRequestOrigin, PhaseUpdateParamsRulesetPhaseHTTPRequestRedirect, PhaseUpdateParamsRulesetPhaseHTTPRequestSanitize, PhaseUpdateParamsRulesetPhaseHTTPRequestSbfm, PhaseUpdateParamsRulesetPhaseHTTPRequestSelectConfiguration, PhaseUpdateParamsRulesetPhaseHTTPRequestTransform, PhaseUpdateParamsRulesetPhaseHTTPResponseCompression, PhaseUpdateParamsRulesetPhaseHTTPResponseFirewallManaged, PhaseUpdateParamsRulesetPhaseHTTPResponseHeadersTransform, PhaseUpdateParamsRulesetPhaseMagicTransit, PhaseUpdateParamsRulesetPhaseMagicTransitIDsManaged, PhaseUpdateParamsRulesetPhaseMagicTransitManaged:
		return true
	}
	return false
}

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
	Logging param.Field[shared.LoggingParam] `json:"logging"`
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

func (r PhaseUpdateParamsRulesRulesetsBlockRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsBlockRuleActionBlock:
		return true
	}
	return false
}

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
	Logging param.Field[shared.LoggingParam] `json:"logging"`
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

func (r PhaseUpdateParamsRulesRulesetsExecuteRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsExecuteRuleActionExecute:
		return true
	}
	return false
}

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

func (r PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

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

func (r PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault, PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium, PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow, PhaseUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
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
	Logging param.Field[shared.LoggingParam] `json:"logging"`
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

func (r PhaseUpdateParamsRulesRulesetsLogRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsLogRuleActionLog:
		return true
	}
	return false
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
	Logging param.Field[shared.LoggingParam] `json:"logging"`
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

func (r PhaseUpdateParamsRulesRulesetsSkipRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSkipRuleActionSkip:
		return true
	}
	return false
}

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
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseDDoSL4                         PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseDDoSL7                         PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
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

func (r PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseDDoSL4, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseDDoSL7, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransit, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

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

func (r PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductBic, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductHot, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductRateLimit, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductSecurityLevel, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductUABlock, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductWAF, PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset string

const (
	PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRulesetCurrent PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

func (r PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

// The kind of the ruleset.
type PhaseUpdateParamsKind string

const (
	PhaseUpdateParamsKindManaged PhaseUpdateParamsKind = "managed"
	PhaseUpdateParamsKindCustom  PhaseUpdateParamsKind = "custom"
	PhaseUpdateParamsKindRoot    PhaseUpdateParamsKind = "root"
	PhaseUpdateParamsKindZone    PhaseUpdateParamsKind = "zone"
)

func (r PhaseUpdateParamsKind) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsKindManaged, PhaseUpdateParamsKindCustom, PhaseUpdateParamsKindRoot, PhaseUpdateParamsKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type PhaseUpdateParamsPhase string

const (
	PhaseUpdateParamsPhaseDDoSL4                         PhaseUpdateParamsPhase = "ddos_l4"
	PhaseUpdateParamsPhaseDDoSL7                         PhaseUpdateParamsPhase = "ddos_l7"
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

func (r PhaseUpdateParamsPhase) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsPhaseDDoSL4, PhaseUpdateParamsPhaseDDoSL7, PhaseUpdateParamsPhaseHTTPConfigSettings, PhaseUpdateParamsPhaseHTTPCustomErrors, PhaseUpdateParamsPhaseHTTPLogCustomFields, PhaseUpdateParamsPhaseHTTPRatelimit, PhaseUpdateParamsPhaseHTTPRequestCacheSettings, PhaseUpdateParamsPhaseHTTPRequestDynamicRedirect, PhaseUpdateParamsPhaseHTTPRequestFirewallCustom, PhaseUpdateParamsPhaseHTTPRequestFirewallManaged, PhaseUpdateParamsPhaseHTTPRequestLateTransform, PhaseUpdateParamsPhaseHTTPRequestOrigin, PhaseUpdateParamsPhaseHTTPRequestRedirect, PhaseUpdateParamsPhaseHTTPRequestSanitize, PhaseUpdateParamsPhaseHTTPRequestSbfm, PhaseUpdateParamsPhaseHTTPRequestSelectConfiguration, PhaseUpdateParamsPhaseHTTPRequestTransform, PhaseUpdateParamsPhaseHTTPResponseCompression, PhaseUpdateParamsPhaseHTTPResponseFirewallManaged, PhaseUpdateParamsPhaseHTTPResponseHeadersTransform, PhaseUpdateParamsPhaseMagicTransit, PhaseUpdateParamsPhaseMagicTransitIDsManaged, PhaseUpdateParamsPhaseMagicTransitManaged:
		return true
	}
	return false
}

// A response object.
type PhaseUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result PhaseUpdateResponse `json:"result,required"`
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

func (r PhaseUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PhaseGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// The phase of the ruleset.
type PhaseGetParamsRulesetPhase string

const (
	PhaseGetParamsRulesetPhaseDDoSL4                         PhaseGetParamsRulesetPhase = "ddos_l4"
	PhaseGetParamsRulesetPhaseDDoSL7                         PhaseGetParamsRulesetPhase = "ddos_l7"
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

func (r PhaseGetParamsRulesetPhase) IsKnown() bool {
	switch r {
	case PhaseGetParamsRulesetPhaseDDoSL4, PhaseGetParamsRulesetPhaseDDoSL7, PhaseGetParamsRulesetPhaseHTTPConfigSettings, PhaseGetParamsRulesetPhaseHTTPCustomErrors, PhaseGetParamsRulesetPhaseHTTPLogCustomFields, PhaseGetParamsRulesetPhaseHTTPRatelimit, PhaseGetParamsRulesetPhaseHTTPRequestCacheSettings, PhaseGetParamsRulesetPhaseHTTPRequestDynamicRedirect, PhaseGetParamsRulesetPhaseHTTPRequestFirewallCustom, PhaseGetParamsRulesetPhaseHTTPRequestFirewallManaged, PhaseGetParamsRulesetPhaseHTTPRequestLateTransform, PhaseGetParamsRulesetPhaseHTTPRequestOrigin, PhaseGetParamsRulesetPhaseHTTPRequestRedirect, PhaseGetParamsRulesetPhaseHTTPRequestSanitize, PhaseGetParamsRulesetPhaseHTTPRequestSbfm, PhaseGetParamsRulesetPhaseHTTPRequestSelectConfiguration, PhaseGetParamsRulesetPhaseHTTPRequestTransform, PhaseGetParamsRulesetPhaseHTTPResponseCompression, PhaseGetParamsRulesetPhaseHTTPResponseFirewallManaged, PhaseGetParamsRulesetPhaseHTTPResponseHeadersTransform, PhaseGetParamsRulesetPhaseMagicTransit, PhaseGetParamsRulesetPhaseMagicTransitIDsManaged, PhaseGetParamsRulesetPhaseMagicTransitManaged:
		return true
	}
	return false
}

// A response object.
type PhaseGetResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result PhaseGetResponse `json:"result,required"`
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

func (r PhaseGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PhaseGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
