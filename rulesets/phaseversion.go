// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// PhaseVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPhaseVersionService] method
// instead.
type PhaseVersionService struct {
	Options []option.RequestOption
}

// NewPhaseVersionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPhaseVersionService(opts ...option.RequestOption) (r *PhaseVersionService) {
	r = &PhaseVersionService{}
	r.Options = opts
	return
}

// Fetches the versions of an account or zone entry point ruleset.
func (r *PhaseVersionService) List(ctx context.Context, rulesetPhase PhaseVersionListParamsRulesetPhase, query PhaseVersionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Ruleset], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint/versions", accountOrZone, accountOrZoneID, rulesetPhase)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches the versions of an account or zone entry point ruleset.
func (r *PhaseVersionService) ListAutoPaging(ctx context.Context, rulesetPhase PhaseVersionListParamsRulesetPhase, query PhaseVersionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Ruleset] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, rulesetPhase, query, opts...))
}

// Fetches a specific version of an account or zone entry point ruleset.
func (r *PhaseVersionService) Get(ctx context.Context, rulesetPhase PhaseVersionGetParamsRulesetPhase, rulesetVersion string, query PhaseVersionGetParams, opts ...option.RequestOption) (res *PhaseVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PhaseVersionGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint/versions/%s", accountOrZone, accountOrZoneID, rulesetPhase, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type PhaseVersionGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind PhaseVersionGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase PhaseVersionGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []PhaseVersionGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                      `json:"description"`
	JSON        phaseVersionGetResponseJSON `json:"-"`
}

// phaseVersionGetResponseJSON contains the JSON metadata for the struct
// [PhaseVersionGetResponse]
type phaseVersionGetResponseJSON struct {
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

func (r *PhaseVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type PhaseVersionGetResponseKind string

const (
	PhaseVersionGetResponseKindManaged PhaseVersionGetResponseKind = "managed"
	PhaseVersionGetResponseKindCustom  PhaseVersionGetResponseKind = "custom"
	PhaseVersionGetResponseKindRoot    PhaseVersionGetResponseKind = "root"
	PhaseVersionGetResponseKindZone    PhaseVersionGetResponseKind = "zone"
)

func (r PhaseVersionGetResponseKind) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseKindManaged, PhaseVersionGetResponseKindCustom, PhaseVersionGetResponseKindRoot, PhaseVersionGetResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type PhaseVersionGetResponsePhase string

const (
	PhaseVersionGetResponsePhaseDDoSL4                         PhaseVersionGetResponsePhase = "ddos_l4"
	PhaseVersionGetResponsePhaseDDoSL7                         PhaseVersionGetResponsePhase = "ddos_l7"
	PhaseVersionGetResponsePhaseHTTPConfigSettings             PhaseVersionGetResponsePhase = "http_config_settings"
	PhaseVersionGetResponsePhaseHTTPCustomErrors               PhaseVersionGetResponsePhase = "http_custom_errors"
	PhaseVersionGetResponsePhaseHTTPLogCustomFields            PhaseVersionGetResponsePhase = "http_log_custom_fields"
	PhaseVersionGetResponsePhaseHTTPRatelimit                  PhaseVersionGetResponsePhase = "http_ratelimit"
	PhaseVersionGetResponsePhaseHTTPRequestCacheSettings       PhaseVersionGetResponsePhase = "http_request_cache_settings"
	PhaseVersionGetResponsePhaseHTTPRequestDynamicRedirect     PhaseVersionGetResponsePhase = "http_request_dynamic_redirect"
	PhaseVersionGetResponsePhaseHTTPRequestFirewallCustom      PhaseVersionGetResponsePhase = "http_request_firewall_custom"
	PhaseVersionGetResponsePhaseHTTPRequestFirewallManaged     PhaseVersionGetResponsePhase = "http_request_firewall_managed"
	PhaseVersionGetResponsePhaseHTTPRequestLateTransform       PhaseVersionGetResponsePhase = "http_request_late_transform"
	PhaseVersionGetResponsePhaseHTTPRequestOrigin              PhaseVersionGetResponsePhase = "http_request_origin"
	PhaseVersionGetResponsePhaseHTTPRequestRedirect            PhaseVersionGetResponsePhase = "http_request_redirect"
	PhaseVersionGetResponsePhaseHTTPRequestSanitize            PhaseVersionGetResponsePhase = "http_request_sanitize"
	PhaseVersionGetResponsePhaseHTTPRequestSbfm                PhaseVersionGetResponsePhase = "http_request_sbfm"
	PhaseVersionGetResponsePhaseHTTPRequestSelectConfiguration PhaseVersionGetResponsePhase = "http_request_select_configuration"
	PhaseVersionGetResponsePhaseHTTPRequestTransform           PhaseVersionGetResponsePhase = "http_request_transform"
	PhaseVersionGetResponsePhaseHTTPResponseCompression        PhaseVersionGetResponsePhase = "http_response_compression"
	PhaseVersionGetResponsePhaseHTTPResponseFirewallManaged    PhaseVersionGetResponsePhase = "http_response_firewall_managed"
	PhaseVersionGetResponsePhaseHTTPResponseHeadersTransform   PhaseVersionGetResponsePhase = "http_response_headers_transform"
	PhaseVersionGetResponsePhaseMagicTransit                   PhaseVersionGetResponsePhase = "magic_transit"
	PhaseVersionGetResponsePhaseMagicTransitIDsManaged         PhaseVersionGetResponsePhase = "magic_transit_ids_managed"
	PhaseVersionGetResponsePhaseMagicTransitManaged            PhaseVersionGetResponsePhase = "magic_transit_managed"
)

func (r PhaseVersionGetResponsePhase) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponsePhaseDDoSL4, PhaseVersionGetResponsePhaseDDoSL7, PhaseVersionGetResponsePhaseHTTPConfigSettings, PhaseVersionGetResponsePhaseHTTPCustomErrors, PhaseVersionGetResponsePhaseHTTPLogCustomFields, PhaseVersionGetResponsePhaseHTTPRatelimit, PhaseVersionGetResponsePhaseHTTPRequestCacheSettings, PhaseVersionGetResponsePhaseHTTPRequestDynamicRedirect, PhaseVersionGetResponsePhaseHTTPRequestFirewallCustom, PhaseVersionGetResponsePhaseHTTPRequestFirewallManaged, PhaseVersionGetResponsePhaseHTTPRequestLateTransform, PhaseVersionGetResponsePhaseHTTPRequestOrigin, PhaseVersionGetResponsePhaseHTTPRequestRedirect, PhaseVersionGetResponsePhaseHTTPRequestSanitize, PhaseVersionGetResponsePhaseHTTPRequestSbfm, PhaseVersionGetResponsePhaseHTTPRequestSelectConfiguration, PhaseVersionGetResponsePhaseHTTPRequestTransform, PhaseVersionGetResponsePhaseHTTPResponseCompression, PhaseVersionGetResponsePhaseHTTPResponseFirewallManaged, PhaseVersionGetResponsePhaseHTTPResponseHeadersTransform, PhaseVersionGetResponsePhaseMagicTransit, PhaseVersionGetResponsePhaseMagicTransitIDsManaged, PhaseVersionGetResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type PhaseVersionGetResponseRule struct {
	// The action to perform when the rule matches.
	Action           PhaseVersionGetResponseRulesAction `json:"action"`
	ActionParameters interface{}                        `json:"action_parameters,required"`
	Categories       interface{}                        `json:"categories,required"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// An object configuring the rule's logging behavior.
	Logging shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                          `json:"version,required"`
	JSON    phaseVersionGetResponseRuleJSON `json:"-"`
	union   PhaseVersionGetResponseRulesUnion
}

// phaseVersionGetResponseRuleJSON contains the JSON metadata for the struct
// [PhaseVersionGetResponseRule]
type phaseVersionGetResponseRuleJSON struct {
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	ID               apijson.Field
	LastUpdated      apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r phaseVersionGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseVersionGetResponseRule) AsUnion() PhaseVersionGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.PhaseVersionGetResponseRulesRulesetsBlockRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsExecuteRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsLogRule] or
// [rulesets.PhaseVersionGetResponseRulesRulesetsSkipRule].
type PhaseVersionGetResponseRulesUnion interface {
	implementsRulesetsPhaseVersionGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type PhaseVersionGetResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                            `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseRulesRulesetsBlockRule]
type phaseVersionGetResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsBlockRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsBlockRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsBlockRuleActionBlock PhaseVersionGetResponseRulesRulesetsBlockRuleAction = "block"
)

func (r PhaseVersionGetResponseRulesRulesetsBlockRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsBlockRuleActionBlock:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response PhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     phaseVersionGetResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsBlockRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsBlockRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type PhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                     `json:"status_code,required"`
	JSON       phaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponse]
type phaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                              `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseRulesRulesetsExecuteRule]
type phaseVersionGetResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsExecuteRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsExecuteRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionExecute PhaseVersionGetResponseRulesRulesetsExecuteRuleAction = "execute"
)

func (r PhaseVersionGetResponseRulesRulesetsExecuteRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsExecuteRuleActionExecute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                         `json:"public_key,required"`
	JSON      phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

func (r PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

// A rule-level override
type PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

func (r PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault, PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium, PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow, PhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

type PhaseVersionGetResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                          `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsLogRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseRulesRulesetsLogRule]
type phaseVersionGetResponseRulesRulesetsLogRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsLogRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsLogRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsLogRuleActionLog PhaseVersionGetResponseRulesRulesetsLogRuleAction = "log"
)

func (r PhaseVersionGetResponseRulesRulesetsLogRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsLogRuleActionLog:
		return true
	}
	return false
}

type PhaseVersionGetResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseRulesRulesetsSkipRule]
type phaseVersionGetResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSkipRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsSkipRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionSkip PhaseVersionGetResponseRulesRulesetsSkipRuleAction = "skip"
)

func (r PhaseVersionGetResponseRulesRulesetsSkipRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSkipRuleActionSkip:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                         `json:"rulesets"`
	JSON     phaseVersionGetResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSkipRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSkipRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL4                         PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL7                         PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

func (r PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL4, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL7, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

// The name of a legacy security product to skip the execution of.
type PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductBic           PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductHot           PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductUABlock       PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductWAF           PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

func (r PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductBic, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductHot, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductUABlock, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductWAF, PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

func (r PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesAction string

const (
	PhaseVersionGetResponseRulesActionBlock   PhaseVersionGetResponseRulesAction = "block"
	PhaseVersionGetResponseRulesActionExecute PhaseVersionGetResponseRulesAction = "execute"
	PhaseVersionGetResponseRulesActionLog     PhaseVersionGetResponseRulesAction = "log"
	PhaseVersionGetResponseRulesActionSkip    PhaseVersionGetResponseRulesAction = "skip"
)

func (r PhaseVersionGetResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesActionBlock, PhaseVersionGetResponseRulesActionExecute, PhaseVersionGetResponseRulesActionLog, PhaseVersionGetResponseRulesActionSkip:
		return true
	}
	return false
}

type PhaseVersionListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// The phase of the ruleset.
type PhaseVersionListParamsRulesetPhase string

const (
	PhaseVersionListParamsRulesetPhaseDDoSL4                         PhaseVersionListParamsRulesetPhase = "ddos_l4"
	PhaseVersionListParamsRulesetPhaseDDoSL7                         PhaseVersionListParamsRulesetPhase = "ddos_l7"
	PhaseVersionListParamsRulesetPhaseHTTPConfigSettings             PhaseVersionListParamsRulesetPhase = "http_config_settings"
	PhaseVersionListParamsRulesetPhaseHTTPCustomErrors               PhaseVersionListParamsRulesetPhase = "http_custom_errors"
	PhaseVersionListParamsRulesetPhaseHTTPLogCustomFields            PhaseVersionListParamsRulesetPhase = "http_log_custom_fields"
	PhaseVersionListParamsRulesetPhaseHTTPRatelimit                  PhaseVersionListParamsRulesetPhase = "http_ratelimit"
	PhaseVersionListParamsRulesetPhaseHTTPRequestCacheSettings       PhaseVersionListParamsRulesetPhase = "http_request_cache_settings"
	PhaseVersionListParamsRulesetPhaseHTTPRequestDynamicRedirect     PhaseVersionListParamsRulesetPhase = "http_request_dynamic_redirect"
	PhaseVersionListParamsRulesetPhaseHTTPRequestFirewallCustom      PhaseVersionListParamsRulesetPhase = "http_request_firewall_custom"
	PhaseVersionListParamsRulesetPhaseHTTPRequestFirewallManaged     PhaseVersionListParamsRulesetPhase = "http_request_firewall_managed"
	PhaseVersionListParamsRulesetPhaseHTTPRequestLateTransform       PhaseVersionListParamsRulesetPhase = "http_request_late_transform"
	PhaseVersionListParamsRulesetPhaseHTTPRequestOrigin              PhaseVersionListParamsRulesetPhase = "http_request_origin"
	PhaseVersionListParamsRulesetPhaseHTTPRequestRedirect            PhaseVersionListParamsRulesetPhase = "http_request_redirect"
	PhaseVersionListParamsRulesetPhaseHTTPRequestSanitize            PhaseVersionListParamsRulesetPhase = "http_request_sanitize"
	PhaseVersionListParamsRulesetPhaseHTTPRequestSbfm                PhaseVersionListParamsRulesetPhase = "http_request_sbfm"
	PhaseVersionListParamsRulesetPhaseHTTPRequestSelectConfiguration PhaseVersionListParamsRulesetPhase = "http_request_select_configuration"
	PhaseVersionListParamsRulesetPhaseHTTPRequestTransform           PhaseVersionListParamsRulesetPhase = "http_request_transform"
	PhaseVersionListParamsRulesetPhaseHTTPResponseCompression        PhaseVersionListParamsRulesetPhase = "http_response_compression"
	PhaseVersionListParamsRulesetPhaseHTTPResponseFirewallManaged    PhaseVersionListParamsRulesetPhase = "http_response_firewall_managed"
	PhaseVersionListParamsRulesetPhaseHTTPResponseHeadersTransform   PhaseVersionListParamsRulesetPhase = "http_response_headers_transform"
	PhaseVersionListParamsRulesetPhaseMagicTransit                   PhaseVersionListParamsRulesetPhase = "magic_transit"
	PhaseVersionListParamsRulesetPhaseMagicTransitIDsManaged         PhaseVersionListParamsRulesetPhase = "magic_transit_ids_managed"
	PhaseVersionListParamsRulesetPhaseMagicTransitManaged            PhaseVersionListParamsRulesetPhase = "magic_transit_managed"
)

func (r PhaseVersionListParamsRulesetPhase) IsKnown() bool {
	switch r {
	case PhaseVersionListParamsRulesetPhaseDDoSL4, PhaseVersionListParamsRulesetPhaseDDoSL7, PhaseVersionListParamsRulesetPhaseHTTPConfigSettings, PhaseVersionListParamsRulesetPhaseHTTPCustomErrors, PhaseVersionListParamsRulesetPhaseHTTPLogCustomFields, PhaseVersionListParamsRulesetPhaseHTTPRatelimit, PhaseVersionListParamsRulesetPhaseHTTPRequestCacheSettings, PhaseVersionListParamsRulesetPhaseHTTPRequestDynamicRedirect, PhaseVersionListParamsRulesetPhaseHTTPRequestFirewallCustom, PhaseVersionListParamsRulesetPhaseHTTPRequestFirewallManaged, PhaseVersionListParamsRulesetPhaseHTTPRequestLateTransform, PhaseVersionListParamsRulesetPhaseHTTPRequestOrigin, PhaseVersionListParamsRulesetPhaseHTTPRequestRedirect, PhaseVersionListParamsRulesetPhaseHTTPRequestSanitize, PhaseVersionListParamsRulesetPhaseHTTPRequestSbfm, PhaseVersionListParamsRulesetPhaseHTTPRequestSelectConfiguration, PhaseVersionListParamsRulesetPhaseHTTPRequestTransform, PhaseVersionListParamsRulesetPhaseHTTPResponseCompression, PhaseVersionListParamsRulesetPhaseHTTPResponseFirewallManaged, PhaseVersionListParamsRulesetPhaseHTTPResponseHeadersTransform, PhaseVersionListParamsRulesetPhaseMagicTransit, PhaseVersionListParamsRulesetPhaseMagicTransitIDsManaged, PhaseVersionListParamsRulesetPhaseMagicTransitManaged:
		return true
	}
	return false
}

type PhaseVersionGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// The phase of the ruleset.
type PhaseVersionGetParamsRulesetPhase string

const (
	PhaseVersionGetParamsRulesetPhaseDDoSL4                         PhaseVersionGetParamsRulesetPhase = "ddos_l4"
	PhaseVersionGetParamsRulesetPhaseDDoSL7                         PhaseVersionGetParamsRulesetPhase = "ddos_l7"
	PhaseVersionGetParamsRulesetPhaseHTTPConfigSettings             PhaseVersionGetParamsRulesetPhase = "http_config_settings"
	PhaseVersionGetParamsRulesetPhaseHTTPCustomErrors               PhaseVersionGetParamsRulesetPhase = "http_custom_errors"
	PhaseVersionGetParamsRulesetPhaseHTTPLogCustomFields            PhaseVersionGetParamsRulesetPhase = "http_log_custom_fields"
	PhaseVersionGetParamsRulesetPhaseHTTPRatelimit                  PhaseVersionGetParamsRulesetPhase = "http_ratelimit"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestCacheSettings       PhaseVersionGetParamsRulesetPhase = "http_request_cache_settings"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestDynamicRedirect     PhaseVersionGetParamsRulesetPhase = "http_request_dynamic_redirect"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestFirewallCustom      PhaseVersionGetParamsRulesetPhase = "http_request_firewall_custom"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestFirewallManaged     PhaseVersionGetParamsRulesetPhase = "http_request_firewall_managed"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestLateTransform       PhaseVersionGetParamsRulesetPhase = "http_request_late_transform"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestOrigin              PhaseVersionGetParamsRulesetPhase = "http_request_origin"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestRedirect            PhaseVersionGetParamsRulesetPhase = "http_request_redirect"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestSanitize            PhaseVersionGetParamsRulesetPhase = "http_request_sanitize"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestSbfm                PhaseVersionGetParamsRulesetPhase = "http_request_sbfm"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestSelectConfiguration PhaseVersionGetParamsRulesetPhase = "http_request_select_configuration"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestTransform           PhaseVersionGetParamsRulesetPhase = "http_request_transform"
	PhaseVersionGetParamsRulesetPhaseHTTPResponseCompression        PhaseVersionGetParamsRulesetPhase = "http_response_compression"
	PhaseVersionGetParamsRulesetPhaseHTTPResponseFirewallManaged    PhaseVersionGetParamsRulesetPhase = "http_response_firewall_managed"
	PhaseVersionGetParamsRulesetPhaseHTTPResponseHeadersTransform   PhaseVersionGetParamsRulesetPhase = "http_response_headers_transform"
	PhaseVersionGetParamsRulesetPhaseMagicTransit                   PhaseVersionGetParamsRulesetPhase = "magic_transit"
	PhaseVersionGetParamsRulesetPhaseMagicTransitIDsManaged         PhaseVersionGetParamsRulesetPhase = "magic_transit_ids_managed"
	PhaseVersionGetParamsRulesetPhaseMagicTransitManaged            PhaseVersionGetParamsRulesetPhase = "magic_transit_managed"
)

func (r PhaseVersionGetParamsRulesetPhase) IsKnown() bool {
	switch r {
	case PhaseVersionGetParamsRulesetPhaseDDoSL4, PhaseVersionGetParamsRulesetPhaseDDoSL7, PhaseVersionGetParamsRulesetPhaseHTTPConfigSettings, PhaseVersionGetParamsRulesetPhaseHTTPCustomErrors, PhaseVersionGetParamsRulesetPhaseHTTPLogCustomFields, PhaseVersionGetParamsRulesetPhaseHTTPRatelimit, PhaseVersionGetParamsRulesetPhaseHTTPRequestCacheSettings, PhaseVersionGetParamsRulesetPhaseHTTPRequestDynamicRedirect, PhaseVersionGetParamsRulesetPhaseHTTPRequestFirewallCustom, PhaseVersionGetParamsRulesetPhaseHTTPRequestFirewallManaged, PhaseVersionGetParamsRulesetPhaseHTTPRequestLateTransform, PhaseVersionGetParamsRulesetPhaseHTTPRequestOrigin, PhaseVersionGetParamsRulesetPhaseHTTPRequestRedirect, PhaseVersionGetParamsRulesetPhaseHTTPRequestSanitize, PhaseVersionGetParamsRulesetPhaseHTTPRequestSbfm, PhaseVersionGetParamsRulesetPhaseHTTPRequestSelectConfiguration, PhaseVersionGetParamsRulesetPhaseHTTPRequestTransform, PhaseVersionGetParamsRulesetPhaseHTTPResponseCompression, PhaseVersionGetParamsRulesetPhaseHTTPResponseFirewallManaged, PhaseVersionGetParamsRulesetPhaseHTTPResponseHeadersTransform, PhaseVersionGetParamsRulesetPhaseMagicTransit, PhaseVersionGetParamsRulesetPhaseMagicTransitIDsManaged, PhaseVersionGetParamsRulesetPhaseMagicTransitManaged:
		return true
	}
	return false
}

// A response object.
type PhaseVersionGetResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseVersionGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseVersionGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result PhaseVersionGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success PhaseVersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    phaseVersionGetResponseEnvelopeJSON    `json:"-"`
}

// phaseVersionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PhaseVersionGetResponseEnvelope]
type phaseVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseVersionGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseVersionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   phaseVersionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// phaseVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseEnvelopeErrors]
type phaseVersionGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseVersionGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    phaseVersionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// phaseVersionGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseEnvelopeErrorsSource]
type phaseVersionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseVersionGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseVersionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   phaseVersionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// phaseVersionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseEnvelopeMessages]
type phaseVersionGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseVersionGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    phaseVersionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// phaseVersionGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseEnvelopeMessagesSource]
type phaseVersionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PhaseVersionGetResponseEnvelopeSuccess bool

const (
	PhaseVersionGetResponseEnvelopeSuccessTrue PhaseVersionGetResponseEnvelopeSuccess = true
)

func (r PhaseVersionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
