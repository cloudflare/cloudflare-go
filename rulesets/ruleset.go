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
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// RulesetService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRulesetService] method instead.
type RulesetService struct {
	Options  []option.RequestOption
	Phases   *PhaseService
	Rules    *RuleService
	Versions *VersionService
}

// NewRulesetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRulesetService(opts ...option.RequestOption) (r *RulesetService) {
	r = &RulesetService{}
	r.Options = opts
	r.Phases = NewPhaseService(opts...)
	r.Rules = NewRuleService(opts...)
	r.Versions = NewVersionService(opts...)
	return
}

// Creates a ruleset.
func (r *RulesetService) New(ctx context.Context, params RulesetNewParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
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
func (r *RulesetService) Update(ctx context.Context, rulesetID string, params RulesetUpdateParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
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
func (r *RulesetService) List(ctx context.Context, query RulesetListParams, opts ...option.RequestOption) (res *RulesetsRulesetsResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches the latest version of an account or zone ruleset.
func (r *RulesetService) Get(ctx context.Context, rulesetID string, query RulesetGetParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type RulesetsRulesetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetsRulesetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetsRulesetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetsRulesetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                      `json:"description"`
	JSON        rulesetsRulesetResponseJSON `json:"-"`
}

// rulesetsRulesetResponseJSON contains the JSON metadata for the struct
// [RulesetsRulesetResponse]
type rulesetsRulesetResponseJSON struct {
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

func (r *RulesetsRulesetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RulesetsRulesetResponseKind string

const (
	RulesetsRulesetResponseKindManaged RulesetsRulesetResponseKind = "managed"
	RulesetsRulesetResponseKindCustom  RulesetsRulesetResponseKind = "custom"
	RulesetsRulesetResponseKindRoot    RulesetsRulesetResponseKind = "root"
	RulesetsRulesetResponseKindZone    RulesetsRulesetResponseKind = "zone"
)

func (r RulesetsRulesetResponseKind) IsKnown() bool {
	switch r {
	case RulesetsRulesetResponseKindManaged, RulesetsRulesetResponseKindCustom, RulesetsRulesetResponseKindRoot, RulesetsRulesetResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type RulesetsRulesetResponsePhase string

const (
	RulesetsRulesetResponsePhaseDDOSL4                         RulesetsRulesetResponsePhase = "ddos_l4"
	RulesetsRulesetResponsePhaseDDOSL7                         RulesetsRulesetResponsePhase = "ddos_l7"
	RulesetsRulesetResponsePhaseHTTPConfigSettings             RulesetsRulesetResponsePhase = "http_config_settings"
	RulesetsRulesetResponsePhaseHTTPCustomErrors               RulesetsRulesetResponsePhase = "http_custom_errors"
	RulesetsRulesetResponsePhaseHTTPLogCustomFields            RulesetsRulesetResponsePhase = "http_log_custom_fields"
	RulesetsRulesetResponsePhaseHTTPRatelimit                  RulesetsRulesetResponsePhase = "http_ratelimit"
	RulesetsRulesetResponsePhaseHTTPRequestCacheSettings       RulesetsRulesetResponsePhase = "http_request_cache_settings"
	RulesetsRulesetResponsePhaseHTTPRequestDynamicRedirect     RulesetsRulesetResponsePhase = "http_request_dynamic_redirect"
	RulesetsRulesetResponsePhaseHTTPRequestFirewallCustom      RulesetsRulesetResponsePhase = "http_request_firewall_custom"
	RulesetsRulesetResponsePhaseHTTPRequestFirewallManaged     RulesetsRulesetResponsePhase = "http_request_firewall_managed"
	RulesetsRulesetResponsePhaseHTTPRequestLateTransform       RulesetsRulesetResponsePhase = "http_request_late_transform"
	RulesetsRulesetResponsePhaseHTTPRequestOrigin              RulesetsRulesetResponsePhase = "http_request_origin"
	RulesetsRulesetResponsePhaseHTTPRequestRedirect            RulesetsRulesetResponsePhase = "http_request_redirect"
	RulesetsRulesetResponsePhaseHTTPRequestSanitize            RulesetsRulesetResponsePhase = "http_request_sanitize"
	RulesetsRulesetResponsePhaseHTTPRequestSbfm                RulesetsRulesetResponsePhase = "http_request_sbfm"
	RulesetsRulesetResponsePhaseHTTPRequestSelectConfiguration RulesetsRulesetResponsePhase = "http_request_select_configuration"
	RulesetsRulesetResponsePhaseHTTPRequestTransform           RulesetsRulesetResponsePhase = "http_request_transform"
	RulesetsRulesetResponsePhaseHTTPResponseCompression        RulesetsRulesetResponsePhase = "http_response_compression"
	RulesetsRulesetResponsePhaseHTTPResponseFirewallManaged    RulesetsRulesetResponsePhase = "http_response_firewall_managed"
	RulesetsRulesetResponsePhaseHTTPResponseHeadersTransform   RulesetsRulesetResponsePhase = "http_response_headers_transform"
	RulesetsRulesetResponsePhaseMagicTransit                   RulesetsRulesetResponsePhase = "magic_transit"
	RulesetsRulesetResponsePhaseMagicTransitIDsManaged         RulesetsRulesetResponsePhase = "magic_transit_ids_managed"
	RulesetsRulesetResponsePhaseMagicTransitManaged            RulesetsRulesetResponsePhase = "magic_transit_managed"
)

func (r RulesetsRulesetResponsePhase) IsKnown() bool {
	switch r {
	case RulesetsRulesetResponsePhaseDDOSL4, RulesetsRulesetResponsePhaseDDOSL7, RulesetsRulesetResponsePhaseHTTPConfigSettings, RulesetsRulesetResponsePhaseHTTPCustomErrors, RulesetsRulesetResponsePhaseHTTPLogCustomFields, RulesetsRulesetResponsePhaseHTTPRatelimit, RulesetsRulesetResponsePhaseHTTPRequestCacheSettings, RulesetsRulesetResponsePhaseHTTPRequestDynamicRedirect, RulesetsRulesetResponsePhaseHTTPRequestFirewallCustom, RulesetsRulesetResponsePhaseHTTPRequestFirewallManaged, RulesetsRulesetResponsePhaseHTTPRequestLateTransform, RulesetsRulesetResponsePhaseHTTPRequestOrigin, RulesetsRulesetResponsePhaseHTTPRequestRedirect, RulesetsRulesetResponsePhaseHTTPRequestSanitize, RulesetsRulesetResponsePhaseHTTPRequestSbfm, RulesetsRulesetResponsePhaseHTTPRequestSelectConfiguration, RulesetsRulesetResponsePhaseHTTPRequestTransform, RulesetsRulesetResponsePhaseHTTPResponseCompression, RulesetsRulesetResponsePhaseHTTPResponseFirewallManaged, RulesetsRulesetResponsePhaseHTTPResponseHeadersTransform, RulesetsRulesetResponsePhaseMagicTransit, RulesetsRulesetResponsePhaseMagicTransitIDsManaged, RulesetsRulesetResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

// Union satisfied by [rulesets.RulesetsRulesetResponseRulesRulesetsBlockRule],
// [rulesets.RulesetsRulesetResponseRulesRulesetsExecuteRule],
// [rulesets.RulesetsRulesetResponseRulesRulesetsLogRule] or
// [rulesets.RulesetsRulesetResponseRulesRulesetsSkipRule].
type RulesetsRulesetResponseRule interface {
	implementsRulesetsRulesetsRulesetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetsRulesetResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetsRulesetResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetsRulesetResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetsRulesetResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetsRulesetResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetsRulesetResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetsRulesetResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetsRulesetResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetsRulesetResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                            `json:"ref"`
	JSON rulesetsRulesetResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for
// the struct [RulesetsRulesetResponseRulesRulesetsBlockRule]
type rulesetsRulesetResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RulesetsRulesetResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetsRulesetResponseRulesRulesetsBlockRule) implementsRulesetsRulesetsRulesetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetsRulesetResponseRulesRulesetsBlockRuleAction string

const (
	RulesetsRulesetResponseRulesRulesetsBlockRuleActionBlock RulesetsRulesetResponseRulesRulesetsBlockRuleAction = "block"
)

func (r RulesetsRulesetResponseRulesRulesetsBlockRuleAction) IsKnown() bool {
	switch r {
	case RulesetsRulesetResponseRulesRulesetsBlockRuleActionBlock:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetsRulesetResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetsRulesetResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetsRulesetResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsBlockRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetsRulesetResponseRulesRulesetsBlockRuleActionParameters]
type rulesetsRulesetResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsRulesetResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type RulesetsRulesetResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                     `json:"status_code,required"`
	JSON       rulesetsRulesetResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [RulesetsRulesetResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetsRulesetResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsRulesetResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type RulesetsRulesetResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                     `json:"enabled,required"`
	JSON    rulesetsRulesetResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetsRulesetResponseRulesRulesetsBlockRuleLogging]
type rulesetsRulesetResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsRulesetResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetsRulesetResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetsRulesetResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetsRulesetResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                              `json:"ref"`
	JSON rulesetsRulesetResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata
// for the struct [RulesetsRulesetResponseRulesRulesetsExecuteRule]
type rulesetsRulesetResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RulesetsRulesetResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetsRulesetResponseRulesRulesetsExecuteRule) implementsRulesetsRulesetsRulesetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetsRulesetResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionExecute RulesetsRulesetResponseRulesRulesetsExecuteRuleAction = "execute"
)

func (r RulesetsRulesetResponseRulesRulesetsExecuteRuleAction) IsKnown() bool {
	switch r {
	case RulesetsRulesetResponseRulesRulesetsExecuteRuleActionExecute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                         `json:"public_key,required"`
	JSON      rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

func (r RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

// A rule-level override
type RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

func (r RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault, RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium, RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow, RulesetsRulesetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

// An object configuring the rule's logging behavior.
type RulesetsRulesetResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                       `json:"enabled,required"`
	JSON    rulesetsRulesetResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetsRulesetResponseRulesRulesetsExecuteRuleLogging]
type rulesetsRulesetResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsRulesetResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetsRulesetResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetsRulesetResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RulesetsRulesetResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                          `json:"ref"`
	JSON rulesetsRulesetResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsLogRuleJSON contains the JSON metadata for
// the struct [RulesetsRulesetResponseRulesRulesetsLogRule]
type rulesetsRulesetResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RulesetsRulesetResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetsRulesetResponseRulesRulesetsLogRule) implementsRulesetsRulesetsRulesetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetsRulesetResponseRulesRulesetsLogRuleAction string

const (
	RulesetsRulesetResponseRulesRulesetsLogRuleActionLog RulesetsRulesetResponseRulesRulesetsLogRuleAction = "log"
)

func (r RulesetsRulesetResponseRulesRulesetsLogRuleAction) IsKnown() bool {
	switch r {
	case RulesetsRulesetResponseRulesRulesetsLogRuleActionLog:
		return true
	}
	return false
}

// An object configuring the rule's logging behavior.
type RulesetsRulesetResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                   `json:"enabled,required"`
	JSON    rulesetsRulesetResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsLogRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetsRulesetResponseRulesRulesetsLogRuleLogging]
type rulesetsRulesetResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsRulesetResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetsRulesetResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetsRulesetResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetsRulesetResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetsRulesetResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON rulesetsRulesetResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for
// the struct [RulesetsRulesetResponseRulesRulesetsSkipRule]
type rulesetsRulesetResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RulesetsRulesetResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetsRulesetResponseRulesRulesetsSkipRule) implementsRulesetsRulesetsRulesetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetsRulesetResponseRulesRulesetsSkipRuleAction string

const (
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionSkip RulesetsRulesetResponseRulesRulesetsSkipRuleAction = "skip"
)

func (r RulesetsRulesetResponseRulesRulesetsSkipRuleAction) IsKnown() bool {
	switch r {
	case RulesetsRulesetResponseRulesRulesetsSkipRuleActionSkip:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetsRulesetResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                         `json:"rulesets"`
	JSON     rulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetsRulesetResponseRulesRulesetsSkipRuleActionParameters]
type rulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsRulesetResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

func (r RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

// The name of a legacy security product to skip the execution of.
type RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

func (r RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductBic, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductHot, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductUABlock, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductWAF, RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

func (r RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case RulesetsRulesetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

// An object configuring the rule's logging behavior.
type RulesetsRulesetResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                    `json:"enabled,required"`
	JSON    rulesetsRulesetResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetsRulesetResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetsRulesetResponseRulesRulesetsSkipRuleLogging]
type rulesetsRulesetResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsRulesetResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetsRulesetResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type RulesetsRulesetsResponse []RulesetsRulesetsResponse

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

func (r RulesetListResponseKind) IsKnown() bool {
	switch r {
	case RulesetListResponseKindManaged, RulesetListResponseKindCustom, RulesetListResponseKindRoot, RulesetListResponseKindZone:
		return true
	}
	return false
}

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

func (r RulesetListResponsePhase) IsKnown() bool {
	switch r {
	case RulesetListResponsePhaseDDOSL4, RulesetListResponsePhaseDDOSL7, RulesetListResponsePhaseHTTPConfigSettings, RulesetListResponsePhaseHTTPCustomErrors, RulesetListResponsePhaseHTTPLogCustomFields, RulesetListResponsePhaseHTTPRatelimit, RulesetListResponsePhaseHTTPRequestCacheSettings, RulesetListResponsePhaseHTTPRequestDynamicRedirect, RulesetListResponsePhaseHTTPRequestFirewallCustom, RulesetListResponsePhaseHTTPRequestFirewallManaged, RulesetListResponsePhaseHTTPRequestLateTransform, RulesetListResponsePhaseHTTPRequestOrigin, RulesetListResponsePhaseHTTPRequestRedirect, RulesetListResponsePhaseHTTPRequestSanitize, RulesetListResponsePhaseHTTPRequestSbfm, RulesetListResponsePhaseHTTPRequestSelectConfiguration, RulesetListResponsePhaseHTTPRequestTransform, RulesetListResponsePhaseHTTPResponseCompression, RulesetListResponsePhaseHTTPResponseFirewallManaged, RulesetListResponsePhaseHTTPResponseHeadersTransform, RulesetListResponsePhaseMagicTransit, RulesetListResponsePhaseMagicTransitIDsManaged, RulesetListResponsePhaseMagicTransitManaged:
		return true
	}
	return false
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

func (r RulesetNewParamsKind) IsKnown() bool {
	switch r {
	case RulesetNewParamsKindManaged, RulesetNewParamsKindCustom, RulesetNewParamsKindRoot, RulesetNewParamsKindZone:
		return true
	}
	return false
}

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

func (r RulesetNewParamsPhase) IsKnown() bool {
	switch r {
	case RulesetNewParamsPhaseDDOSL4, RulesetNewParamsPhaseDDOSL7, RulesetNewParamsPhaseHTTPConfigSettings, RulesetNewParamsPhaseHTTPCustomErrors, RulesetNewParamsPhaseHTTPLogCustomFields, RulesetNewParamsPhaseHTTPRatelimit, RulesetNewParamsPhaseHTTPRequestCacheSettings, RulesetNewParamsPhaseHTTPRequestDynamicRedirect, RulesetNewParamsPhaseHTTPRequestFirewallCustom, RulesetNewParamsPhaseHTTPRequestFirewallManaged, RulesetNewParamsPhaseHTTPRequestLateTransform, RulesetNewParamsPhaseHTTPRequestOrigin, RulesetNewParamsPhaseHTTPRequestRedirect, RulesetNewParamsPhaseHTTPRequestSanitize, RulesetNewParamsPhaseHTTPRequestSbfm, RulesetNewParamsPhaseHTTPRequestSelectConfiguration, RulesetNewParamsPhaseHTTPRequestTransform, RulesetNewParamsPhaseHTTPResponseCompression, RulesetNewParamsPhaseHTTPResponseFirewallManaged, RulesetNewParamsPhaseHTTPResponseHeadersTransform, RulesetNewParamsPhaseMagicTransit, RulesetNewParamsPhaseMagicTransitIDsManaged, RulesetNewParamsPhaseMagicTransitManaged:
		return true
	}
	return false
}

// Satisfied by [rulesets.RulesetNewParamsRulesRulesetsBlockRule],
// [rulesets.RulesetNewParamsRulesRulesetsExecuteRule],
// [rulesets.RulesetNewParamsRulesRulesetsLogRule],
// [rulesets.RulesetNewParamsRulesRulesetsSkipRule].
type RulesetNewParamsRule interface {
	implementsRulesetsRulesetNewParamsRule()
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

func (r RulesetNewParamsRulesRulesetsBlockRule) implementsRulesetsRulesetNewParamsRule() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsBlockRuleAction string

const (
	RulesetNewParamsRulesRulesetsBlockRuleActionBlock RulesetNewParamsRulesRulesetsBlockRuleAction = "block"
)

func (r RulesetNewParamsRulesRulesetsBlockRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsBlockRuleActionBlock:
		return true
	}
	return false
}

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

func (r RulesetNewParamsRulesRulesetsExecuteRule) implementsRulesetsRulesetNewParamsRule() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsExecuteRuleAction string

const (
	RulesetNewParamsRulesRulesetsExecuteRuleActionExecute RulesetNewParamsRulesRulesetsExecuteRuleAction = "execute"
)

func (r RulesetNewParamsRulesRulesetsExecuteRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsExecuteRuleActionExecute:
		return true
	}
	return false
}

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

func (r RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

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

func (r RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault, RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium, RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow, RulesetNewParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

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

func (r RulesetNewParamsRulesRulesetsLogRule) implementsRulesetsRulesetNewParamsRule() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsLogRuleAction string

const (
	RulesetNewParamsRulesRulesetsLogRuleActionLog RulesetNewParamsRulesRulesetsLogRuleAction = "log"
)

func (r RulesetNewParamsRulesRulesetsLogRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsLogRuleActionLog:
		return true
	}
	return false
}

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

func (r RulesetNewParamsRulesRulesetsSkipRule) implementsRulesetsRulesetNewParamsRule() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsSkipRuleAction string

const (
	RulesetNewParamsRulesRulesetsSkipRuleActionSkip RulesetNewParamsRulesRulesetsSkipRuleAction = "skip"
)

func (r RulesetNewParamsRulesRulesetsSkipRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSkipRuleActionSkip:
		return true
	}
	return false
}

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

func (r RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL4, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL7, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransit, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged, RulesetNewParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

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

func (r RulesetNewParamsRulesRulesetsSkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductBic, RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductHot, RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductRateLimit, RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductSecurityLevel, RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductUABlock, RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductWAF, RulesetNewParamsRulesRulesetsSkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetNewParamsRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetNewParamsRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetNewParamsRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

func (r RulesetNewParamsRulesRulesetsSkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

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
	Result RulesetsRulesetResponse `json:"result,required"`
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

func (r RulesetNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

// Satisfied by [rulesets.RulesetUpdateParamsRulesRulesetsBlockRule],
// [rulesets.RulesetUpdateParamsRulesRulesetsExecuteRule],
// [rulesets.RulesetUpdateParamsRulesRulesetsLogRule],
// [rulesets.RulesetUpdateParamsRulesRulesetsSkipRule].
type RulesetUpdateParamsRule interface {
	implementsRulesetsRulesetUpdateParamsRule()
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

func (r RulesetUpdateParamsRulesRulesetsBlockRule) implementsRulesetsRulesetUpdateParamsRule() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsBlockRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsBlockRuleActionBlock RulesetUpdateParamsRulesRulesetsBlockRuleAction = "block"
)

func (r RulesetUpdateParamsRulesRulesetsBlockRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsBlockRuleActionBlock:
		return true
	}
	return false
}

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

func (r RulesetUpdateParamsRulesRulesetsExecuteRule) implementsRulesetsRulesetUpdateParamsRule() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsExecuteRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionExecute RulesetUpdateParamsRulesRulesetsExecuteRuleAction = "execute"
)

func (r RulesetUpdateParamsRulesRulesetsExecuteRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsExecuteRuleActionExecute:
		return true
	}
	return false
}

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

func (r RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

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

func (r RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault, RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium, RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow, RulesetUpdateParamsRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

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

func (r RulesetUpdateParamsRulesRulesetsLogRule) implementsRulesetsRulesetUpdateParamsRule() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsLogRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsLogRuleActionLog RulesetUpdateParamsRulesRulesetsLogRuleAction = "log"
)

func (r RulesetUpdateParamsRulesRulesetsLogRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsLogRuleActionLog:
		return true
	}
	return false
}

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

func (r RulesetUpdateParamsRulesRulesetsSkipRule) implementsRulesetsRulesetUpdateParamsRule() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsSkipRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsSkipRuleActionSkip RulesetUpdateParamsRulesRulesetsSkipRuleAction = "skip"
)

func (r RulesetUpdateParamsRulesRulesetsSkipRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSkipRuleActionSkip:
		return true
	}
	return false
}

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

func (r RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL4, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseDDOSL7, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransit, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

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

func (r RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductBic, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductHot, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductRateLimit, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductSecurityLevel, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductUABlock, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductWAF, RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

func (r RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

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

func (r RulesetUpdateParamsKind) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsKindManaged, RulesetUpdateParamsKindCustom, RulesetUpdateParamsKindRoot, RulesetUpdateParamsKindZone:
		return true
	}
	return false
}

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

func (r RulesetUpdateParamsPhase) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsPhaseDDOSL4, RulesetUpdateParamsPhaseDDOSL7, RulesetUpdateParamsPhaseHTTPConfigSettings, RulesetUpdateParamsPhaseHTTPCustomErrors, RulesetUpdateParamsPhaseHTTPLogCustomFields, RulesetUpdateParamsPhaseHTTPRatelimit, RulesetUpdateParamsPhaseHTTPRequestCacheSettings, RulesetUpdateParamsPhaseHTTPRequestDynamicRedirect, RulesetUpdateParamsPhaseHTTPRequestFirewallCustom, RulesetUpdateParamsPhaseHTTPRequestFirewallManaged, RulesetUpdateParamsPhaseHTTPRequestLateTransform, RulesetUpdateParamsPhaseHTTPRequestOrigin, RulesetUpdateParamsPhaseHTTPRequestRedirect, RulesetUpdateParamsPhaseHTTPRequestSanitize, RulesetUpdateParamsPhaseHTTPRequestSbfm, RulesetUpdateParamsPhaseHTTPRequestSelectConfiguration, RulesetUpdateParamsPhaseHTTPRequestTransform, RulesetUpdateParamsPhaseHTTPResponseCompression, RulesetUpdateParamsPhaseHTTPResponseFirewallManaged, RulesetUpdateParamsPhaseHTTPResponseHeadersTransform, RulesetUpdateParamsPhaseMagicTransit, RulesetUpdateParamsPhaseMagicTransitIDsManaged, RulesetUpdateParamsPhaseMagicTransitManaged:
		return true
	}
	return false
}

// A response object.
type RulesetUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
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

func (r RulesetUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
	Result RulesetsRulesetsResponse `json:"result,required"`
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

func (r RulesetListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RulesetListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
	Result RulesetsRulesetResponse `json:"result,required"`
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

func (r RulesetGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RulesetGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
