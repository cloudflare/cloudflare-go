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

type PhaseUpdateResponseRule struct {
	// The action to perform when the rule matches.
	Action           PhaseUpdateResponseRulesAction `json:"action"`
	ActionParameters interface{}                    `json:"action_parameters,required"`
	Categories       interface{}                    `json:"categories,required"`
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
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                      `json:"version,required"`
	JSON    phaseUpdateResponseRuleJSON `json:"-"`
	union   PhaseUpdateResponseRulesUnion
}

// phaseUpdateResponseRuleJSON contains the JSON metadata for the struct
// [PhaseUpdateResponseRule]
type phaseUpdateResponseRuleJSON struct {
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

func (r phaseUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseUpdateResponseRule) AsUnion() PhaseUpdateResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule],
// [rulesets.PhaseUpdateResponseRulesRulesetsChallengeRule],
// [rulesets.PhaseUpdateResponseRulesRulesetsCompressResponseRule],
// [rulesets.ExecuteRule],
// [rulesets.PhaseUpdateResponseRulesRulesetsJsChallengeRule], [rulesets.LogRule],
// [rulesets.PhaseUpdateResponseRulesRulesetsManagedChallengeRule],
// [rulesets.PhaseUpdateResponseRulesRulesetsRedirectRule],
// [rulesets.PhaseUpdateResponseRulesRulesetsRewriteRule],
// [rulesets.PhaseUpdateResponseRulesRulesetsRouteRule],
// [rulesets.PhaseUpdateResponseRulesRulesetsScoreRule],
// [rulesets.PhaseUpdateResponseRulesRulesetsServeErrorRule],
// [rulesets.PhaseUpdateResponseRulesRulesetsSetConfigRule], [rulesets.SkipRule] or
// [rulesets.PhaseUpdateResponseRulesRulesetsSetCacheSettingsRule].
type PhaseUpdateResponseRulesUnion interface {
	implementsRulesetsPhaseUpdateResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsCompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsJsChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsRedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsRewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsRouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
	)
}

type PhaseUpdateResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsChallengeRuleAction `json:"action"`
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
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                            `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [PhaseUpdateResponseRulesRulesetsChallengeRule]
type phaseUpdateResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsChallengeRule) implementsRulesetsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsChallengeRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsChallengeRuleActionChallenge PhaseUpdateResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r PhaseUpdateResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

type PhaseUpdateResponseRulesRulesetsCompressResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsCompressResponseRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                   `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsCompressResponseRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsCompressResponseRuleJSON contains the JSON
// metadata for the struct [PhaseUpdateResponseRulesRulesetsCompressResponseRule]
type phaseUpdateResponseRulesRulesetsCompressResponseRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsCompressResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsCompressResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsCompressResponseRule) implementsRulesetsPhaseUpdateResponseRule() {
}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsCompressResponseRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionCompressResponse PhaseUpdateResponseRulesRulesetsCompressResponseRuleAction = "compress_response"
)

func (r PhaseUpdateResponseRulesRulesetsCompressResponseRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms []PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm `json:"algorithms"`
	JSON       phaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersJSON        `json:"-"`
}

// phaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParameters]
type phaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersJSON struct {
	Algorithms  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Compression algorithm to enable.
type PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName `json:"name"`
	JSON phaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON  `json:"-"`
}

// phaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm]
type phaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON) RawJSON() string {
	return r.raw
}

// Name of compression algorithm to enable.
type PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName string

const (
	PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone    PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "none"
	PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto    PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "auto"
	PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "default"
	PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip    PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "gzip"
	PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli  PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "brotli"
)

func (r PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone, PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto, PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault, PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip, PhaseUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli:
		return true
	}
	return false
}

type PhaseUpdateResponseRulesRulesetsJsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsJsChallengeRuleAction `json:"action"`
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
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                              `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsJsChallengeRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsJsChallengeRuleJSON contains the JSON metadata
// for the struct [PhaseUpdateResponseRulesRulesetsJsChallengeRule]
type phaseUpdateResponseRulesRulesetsJsChallengeRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsJsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsJsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsJsChallengeRule) implementsRulesetsPhaseUpdateResponseRule() {
}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsJsChallengeRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsJsChallengeRuleActionJsChallenge PhaseUpdateResponseRulesRulesetsJsChallengeRuleAction = "js_challenge"
)

func (r PhaseUpdateResponseRulesRulesetsJsChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsJsChallengeRuleActionJsChallenge:
		return true
	}
	return false
}

type PhaseUpdateResponseRulesRulesetsManagedChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsManagedChallengeRuleAction `json:"action"`
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
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                   `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsManagedChallengeRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsManagedChallengeRuleJSON contains the JSON
// metadata for the struct [PhaseUpdateResponseRulesRulesetsManagedChallengeRule]
type phaseUpdateResponseRulesRulesetsManagedChallengeRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsManagedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsManagedChallengeRule) implementsRulesetsPhaseUpdateResponseRule() {
}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsManagedChallengeRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge PhaseUpdateResponseRulesRulesetsManagedChallengeRuleAction = "managed_challenge"
)

func (r PhaseUpdateResponseRulesRulesetsManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

type PhaseUpdateResponseRulesRulesetsRedirectRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsRedirectRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsRedirectRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsRedirectRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRedirectRuleJSON contains the JSON metadata for
// the struct [PhaseUpdateResponseRulesRulesetsRedirectRule]
type phaseUpdateResponseRulesRulesetsRedirectRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsRedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRedirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsRedirectRule) implementsRulesetsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsRedirectRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsRedirectRuleActionRedirect PhaseUpdateResponseRulesRulesetsRedirectRuleAction = "redirect"
)

func (r PhaseUpdateResponseRulesRulesetsRedirectRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsRedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsRedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromList `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValue `json:"from_value"`
	JSON      phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersJSON      `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRedirectRuleActionParameters]
type phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersJSON struct {
	FromList    apijson.Field
	FromValue   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRedirectRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on a bulk list lookup.
type PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key string `json:"key"`
	// The name of the list to match against.
	Name string                                                                   `json:"name"`
	JSON phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromListJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromListJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromList]
type phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromListJSON struct {
	Key         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromListJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on the request properties.
type PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString bool `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL `json:"target_url"`
	JSON      phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON      `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValue]
type phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON struct {
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	TargetURL           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON) RawJSON() string {
	return r.raw
}

// The status code to be used for the redirect.
type PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode float64

const (
	PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301 PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 301
	PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302 PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 302
	PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303 PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 303
	PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307 PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 307
	PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308 PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 308
)

func (r PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301, PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302, PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303, PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307, PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL struct {
	// The URL to redirect the request to.
	Value string `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                             `json:"expression"`
	JSON       phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON `json:"-"`
	union      PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion
}

// phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL]
type phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) AsUnion() PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion {
	return r.union
}

// The URL to redirect the request to.
//
// Union satisfied by
// [rulesets.PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
// or
// [rulesets.PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect].
type PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRulesetsPhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect{}),
		},
	)
}

type PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value string                                                                                              `json:"value"`
	JSON  phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
type phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRulesetsPhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                                               `json:"expression"`
	JSON       phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect]
type phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRulesetsPhaseUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type PhaseUpdateResponseRulesRulesetsRewriteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsRewriteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsRewriteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                          `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsRewriteRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRewriteRuleJSON contains the JSON metadata for
// the struct [PhaseUpdateResponseRulesRulesetsRewriteRule]
type phaseUpdateResponseRulesRulesetsRewriteRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsRewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsRewriteRule) implementsRulesetsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsRewriteRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsRewriteRuleActionRewrite PhaseUpdateResponseRulesRulesetsRewriteRuleAction = "rewrite"
)

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsRewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers map[string]PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader `json:"headers"`
	// URI to rewrite the request to.
	URI  PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURI  `json:"uri"`
	JSON phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRewriteRuleActionParameters]
type phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersJSON struct {
	Headers     apijson.Field
	URI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRewriteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Remove the header from the request.
type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader struct {
	Operation PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation `json:"operation,required"`
	// Static value for the header.
	Value string `json:"value"`
	// Expression for the header value.
	Expression string                                                                `json:"expression"`
	JSON       phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON `json:"-"`
	union      PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion
}

// phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON contains
// the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader]
type phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader) AsUnion() PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion {
	return r.union
}

// Remove the header from the request.
//
// Union satisfied by
// [rulesets.PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader],
// [rulesets.PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
// or
// [rulesets.PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader].
type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion interface {
	implementsRulesetsPhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader{}),
		},
	)
}

// Remove the header from the request.
type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation `json:"operation,required"`
	JSON      phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON      `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader]
type phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON struct {
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) implementsRulesetsPhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Set a request header with a static value.
type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader struct {
	Operation PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation `json:"operation,required"`
	// Static value for the header.
	Value string                                                                             `json:"value,required"`
	JSON  phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
type phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) implementsRulesetsPhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation string

const (
	PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation = "set"
)

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader struct {
	// Expression for the header value.
	Expression string                                                                                   `json:"expression,required"`
	Operation  PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation `json:"operation,required"`
	JSON       phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON      `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader]
type phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) implementsRulesetsPhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation string

const (
	PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation = "set"
)

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet:
		return true
	}
	return false
}

type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation string

const (
	PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "remove"
	PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet    PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "set"
)

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove, PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath `json:"path"`
	// Query portion rewrite.
	Query PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery `json:"query"`
	JSON  phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIJSON  `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIJSON contains the
// JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURI]
type phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIJSON struct {
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIJSON) RawJSON() string {
	return r.raw
}

// Path portion rewrite.
type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                 `json:"expression"`
	JSON       phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON `json:"-"`
	union      PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion
}

// phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON contains
// the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath]
type phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath) AsUnion() PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion {
	return r.union
}

// Path portion rewrite.
//
// Union satisfied by
// [rulesets.PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
// or
// [rulesets.PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue].
type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion interface {
	implementsRulesetsPhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue{}),
		},
	)
}

type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue struct {
	// Predefined replacement value.
	Value string                                                                            `json:"value,required"`
	JSON  phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
type phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) implementsRulesetsPhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                             `json:"expression,required"`
	JSON       phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue]
type phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) implementsRulesetsPhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

// Query portion rewrite.
type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                  `json:"expression"`
	JSON       phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON `json:"-"`
	union      PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion
}

// phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON contains
// the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery]
type phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery) AsUnion() PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion {
	return r.union
}

// Query portion rewrite.
//
// Union satisfied by
// [rulesets.PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
// or
// [rulesets.PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue].
type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion interface {
	implementsRulesetsPhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue{}),
		},
	)
}

type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue struct {
	// Predefined replacement value.
	Value string                                                                             `json:"value,required"`
	JSON  phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
type phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) implementsRulesetsPhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                              `json:"expression,required"`
	JSON       phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue]
type phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) implementsRulesetsPhaseUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type PhaseUpdateResponseRulesRulesetsRouteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsRouteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsRouteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                        `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsRouteRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRouteRuleJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseRulesRulesetsRouteRule]
type phaseUpdateResponseRulesRulesetsRouteRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsRouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRouteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsRouteRule) implementsRulesetsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsRouteRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsRouteRuleActionRoute PhaseUpdateResponseRulesRulesetsRouteRuleAction = "route"
)

func (r PhaseUpdateResponseRulesRulesetsRouteRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsRouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsRouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader string `json:"host_header"`
	// Override the IP/TCP destination.
	Origin PhaseUpdateResponseRulesRulesetsRouteRuleActionParametersOrigin `json:"origin"`
	// Override the Server Name Indication (SNI).
	Sni  PhaseUpdateResponseRulesRulesetsRouteRuleActionParametersSni  `json:"sni"`
	JSON phaseUpdateResponseRulesRulesetsRouteRuleActionParametersJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRouteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRouteRuleActionParameters]
type phaseUpdateResponseRulesRulesetsRouteRuleActionParametersJSON struct {
	HostHeader  apijson.Field
	Origin      apijson.Field
	Sni         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRouteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRouteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Override the IP/TCP destination.
type PhaseUpdateResponseRulesRulesetsRouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host string `json:"host"`
	// Override the destination port.
	Port float64                                                             `json:"port"`
	JSON phaseUpdateResponseRulesRulesetsRouteRuleActionParametersOriginJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRouteRuleActionParametersOriginJSON contains the
// JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRouteRuleActionParametersOrigin]
type phaseUpdateResponseRulesRulesetsRouteRuleActionParametersOriginJSON struct {
	Host        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRouteRuleActionParametersOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRouteRuleActionParametersOriginJSON) RawJSON() string {
	return r.raw
}

// Override the Server Name Indication (SNI).
type PhaseUpdateResponseRulesRulesetsRouteRuleActionParametersSni struct {
	// The SNI override.
	Value string                                                           `json:"value,required"`
	JSON  phaseUpdateResponseRulesRulesetsRouteRuleActionParametersSniJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsRouteRuleActionParametersSniJSON contains the
// JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsRouteRuleActionParametersSni]
type phaseUpdateResponseRulesRulesetsRouteRuleActionParametersSniJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsRouteRuleActionParametersSni) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsRouteRuleActionParametersSniJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseRulesRulesetsScoreRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsScoreRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsScoreRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                        `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsScoreRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsScoreRuleJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseRulesRulesetsScoreRule]
type phaseUpdateResponseRulesRulesetsScoreRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsScoreRule) implementsRulesetsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsScoreRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsScoreRuleActionScore PhaseUpdateResponseRulesRulesetsScoreRuleAction = "score"
)

func (r PhaseUpdateResponseRulesRulesetsScoreRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment int64                                                         `json:"increment"`
	JSON      phaseUpdateResponseRulesRulesetsScoreRuleActionParametersJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsScoreRuleActionParametersJSON contains the JSON
// metadata for the struct
// [PhaseUpdateResponseRulesRulesetsScoreRuleActionParameters]
type phaseUpdateResponseRulesRulesetsScoreRuleActionParametersJSON struct {
	Increment   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsScoreRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsScoreRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseRulesRulesetsServeErrorRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsServeErrorRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                             `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsServeErrorRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsServeErrorRuleJSON contains the JSON metadata
// for the struct [PhaseUpdateResponseRulesRulesetsServeErrorRule]
type phaseUpdateResponseRulesRulesetsServeErrorRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsServeErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsServeErrorRule) implementsRulesetsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsServeErrorRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsServeErrorRuleActionServeError PhaseUpdateResponseRulesRulesetsServeErrorRuleAction = "serve_error"
)

func (r PhaseUpdateResponseRulesRulesetsServeErrorRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParameters struct {
	// Error response content.
	Content string `json:"content"`
	// Content-type header to set with the response.
	ContentType PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType `json:"content_type"`
	// The status code to use for the error.
	StatusCode float64                                                            `json:"status_code"`
	JSON       phaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParameters]
type phaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Content-type header to set with the response.
type PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType string

const (
	PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType = "application/json"
	PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml         PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/xml"
	PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain       PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/plain"
	PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML        PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/html"
)

func (r PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson, PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml, PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain, PhaseUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

type PhaseUpdateResponseRulesRulesetsSetConfigRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsSetConfigRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                            `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsSetConfigRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetConfigRuleJSON contains the JSON metadata for
// the struct [PhaseUpdateResponseRulesRulesetsSetConfigRule]
type phaseUpdateResponseRulesRulesetsSetConfigRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsSetConfigRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetConfigRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetConfigRule) implementsRulesetsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsSetConfigRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionSetConfig PhaseUpdateResponseRulesRulesetsSetConfigRuleAction = "set_config"
)

func (r PhaseUpdateResponseRulesRulesetsSetConfigRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites bool `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominify `json:"autominify"`
	// Turn on or off Browser Integrity Check.
	Bic bool `json:"bic"`
	// Turn off all active Cloudflare Apps.
	DisableApps bool `json:"disable_apps"`
	// Turn off Zaraz.
	DisableZaraz bool `json:"disable_zaraz"`
	// Turn on or off Email Obfuscation.
	EmailObfuscation bool `json:"email_obfuscation"`
	// Turn on or off the Hotlink Protection.
	HotlinkProtection bool `json:"hotlink_protection"`
	// Turn on or off Mirage.
	Mirage bool `json:"mirage"`
	// Turn on or off Opportunistic Encryption.
	OpportunisticEncryption bool `json:"opportunistic_encryption"`
	// Configure the Polish level.
	Polish PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolish `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader bool `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes bool `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg  bool                                                              `json:"sxg"`
	JSON phaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParameters]
type phaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersJSON struct {
	AutomaticHTTPSRewrites  apijson.Field
	Autominify              apijson.Field
	Bic                     apijson.Field
	DisableApps             apijson.Field
	DisableZaraz            apijson.Field
	EmailObfuscation        apijson.Field
	HotlinkProtection       apijson.Field
	Mirage                  apijson.Field
	OpportunisticEncryption apijson.Field
	Polish                  apijson.Field
	RocketLoader            apijson.Field
	SecurityLevel           apijson.Field
	ServerSideExcludes      apijson.Field
	SSL                     apijson.Field
	Sxg                     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Select which file extensions to minify automatically.
type PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	Css bool `json:"css"`
	// Minify HTML files.
	HTML bool `json:"html"`
	// Minify JS files.
	Js   bool                                                                        `json:"js"`
	JSON phaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominify]
type phaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON) RawJSON() string {
	return r.raw
}

// Configure the Polish level.
type PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolish string

const (
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolishOff      PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolish = "off"
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossless"
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy    PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossy"
)

func (r PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolishOff, PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless, PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy:
		return true
	}
	return false
}

// Configure the Security Level.
type PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel string

const (
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff            PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "off"
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "essentially_off"
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow            PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "low"
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium         PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "medium"
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh           PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "high"
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack    PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff, PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff, PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow, PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium, PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh, PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL string

const (
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLOff        PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL = "off"
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible   PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL = "flexible"
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLFull       PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL = "full"
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict     PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL = "strict"
	PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL = "origin_pull"
)

func (r PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLOff, PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible, PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLFull, PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict, PhaseUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                   `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleJSON contains the JSON
// metadata for the struct [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRule]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheSettingsRule) implementsRulesetsPhaseUpdateResponseRule() {
}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleAction = "set_cache_settings"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts []int64 `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache bool `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL `json:"edge_ttl"`
	// When enabled, Cloudflare will aim to strictly adhere to RFC 7234.
	OriginCacheControl bool `json:"origin_cache_control"`
	// Generate Cloudflare error pages from issues sent from the origin server. When
	// on, error pages will trigger for issues from the origin
	OriginErrorPagePassthru bool `json:"origin_error_page_passthru"`
	// Define a timeout value between two successive read operations to your origin
	// server. Historically, the timeout value between two read options from Cloudflare
	// to an origin server is 100 seconds. If you are attempting to reduce HTTP 524
	// errors because of timeouts from an origin server, try increasing this timeout
	// value.
	ReadTimeout int64 `json:"read_timeout"`
	// Specify whether or not Cloudflare should respect strong ETag (entity tag)
	// headers. When off, Cloudflare converts strong ETag headers to weak ETag headers.
	RespectStrongEtags bool `json:"respect_strong_etags"`
	// Define if Cloudflare should serve stale content while getting the latest content
	// from the origin. If on, Cloudflare will not serve stale content while getting
	// the latest content from the origin.
	ServeStale PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale `json:"serve_stale"`
	JSON       phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON       `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParameters]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON struct {
	AdditionalCacheablePorts apijson.Field
	BrowserTTL               apijson.Field
	Cache                    apijson.Field
	CacheKey                 apijson.Field
	CacheReserve             apijson.Field
	EdgeTTL                  apijson.Field
	OriginCacheControl       apijson.Field
	OriginErrorPagePassthru  apijson.Field
	ReadTimeout              apijson.Field
	RespectStrongEtags       apijson.Field
	ServeStale               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64                                                                              `json:"default"`
	JSON    phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON struct {
	Mode        apijson.Field
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON) RawJSON() string {
	return r.raw
}

// Determines which browser ttl mode to use.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType bool `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor bool `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder bool                                                                             `json:"ignore_query_strings_order"`
	JSON                    phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON struct {
	CacheByDeviceType       apijson.Field
	CacheDeceptionArmor     apijson.Field
	CustomKey               apijson.Field
	IgnoreQueryStringsOrder apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON) RawJSON() string {
	return r.raw
}

// Customize which components of the request are included or excluded from the
// cache key.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser `json:"user"`
	JSON phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON) RawJSON() string {
	return r.raw
}

// The cookies to include in building the cache key.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Include these cookies' names and their values.
	Include []string                                                                                        `json:"include"`
	JSON    phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON) RawJSON() string {
	return r.raw
}

// The header names and values to include in building the cache key.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin bool `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include []string                                                                                        `json:"include"`
	JSON    phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON struct {
	CheckPresence apijson.Field
	ExcludeOrigin apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON) RawJSON() string {
	return r.raw
}

// Whether to use the original host or the resolved host in the cache key.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved bool                                                                                          `json:"resolved"`
	JSON     phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON) RawJSON() string {
	return r.raw
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude `json:"include"`
	JSON    phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON    `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON) RawJSON() string {
	return r.raw
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Exclude all query string parameters from use in building the cache key.
	All bool `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List []string                                                                                                    `json:"list"`
	JSON phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON) RawJSON() string {
	return r.raw
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Use all query string parameters in the cache key.
	All bool `json:"all"`
	// A list of query string parameters used to build the cache key.
	List []string                                                                                                    `json:"list"`
	JSON phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON) RawJSON() string {
	return r.raw
}

// Characteristics of the request user agent used in building the cache key.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType bool `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo bool `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang bool                                                                                          `json:"lang"`
	JSON phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON) RawJSON() string {
	return r.raw
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible bool `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize int64                                                                                `json:"min_file_size,required"`
	JSON        phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON struct {
	Eligible    apijson.Field
	MinFileSize apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON) RawJSON() string {
	return r.raw
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64 `json:"default,required"`
	// edge ttl options
	Mode PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL []PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL `json:"status_code_ttl,required"`
	JSON          phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON            `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON struct {
	Default       apijson.Field
	Mode          apijson.Field
	StatusCodeTTL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON) RawJSON() string {
	return r.raw
}

// edge ttl options
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value int64 `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue int64                                                                                        `json:"status_code_value"`
	JSON            phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON struct {
	Value           apijson.Field
	StatusCodeRange apijson.Field
	StatusCodeValue apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON) RawJSON() string {
	return r.raw
}

// The range of status codes used to apply the selected mode.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// response status code lower bound
	From int64 `json:"from,required"`
	// response status code upper bound
	To   int64                                                                                                       `json:"to,required"`
	JSON phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON) RawJSON() string {
	return r.raw
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating bool                                                                               `json:"disable_stale_while_updating,required"`
	JSON                      phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale]
type phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON struct {
	DisableStaleWhileUpdating apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesAction string

const (
	PhaseUpdateResponseRulesActionBlock            PhaseUpdateResponseRulesAction = "block"
	PhaseUpdateResponseRulesActionChallenge        PhaseUpdateResponseRulesAction = "challenge"
	PhaseUpdateResponseRulesActionCompressResponse PhaseUpdateResponseRulesAction = "compress_response"
	PhaseUpdateResponseRulesActionExecute          PhaseUpdateResponseRulesAction = "execute"
	PhaseUpdateResponseRulesActionJsChallenge      PhaseUpdateResponseRulesAction = "js_challenge"
	PhaseUpdateResponseRulesActionLog              PhaseUpdateResponseRulesAction = "log"
	PhaseUpdateResponseRulesActionManagedChallenge PhaseUpdateResponseRulesAction = "managed_challenge"
	PhaseUpdateResponseRulesActionRedirect         PhaseUpdateResponseRulesAction = "redirect"
	PhaseUpdateResponseRulesActionRewrite          PhaseUpdateResponseRulesAction = "rewrite"
	PhaseUpdateResponseRulesActionRoute            PhaseUpdateResponseRulesAction = "route"
	PhaseUpdateResponseRulesActionScore            PhaseUpdateResponseRulesAction = "score"
	PhaseUpdateResponseRulesActionServeError       PhaseUpdateResponseRulesAction = "serve_error"
	PhaseUpdateResponseRulesActionSetConfig        PhaseUpdateResponseRulesAction = "set_config"
	PhaseUpdateResponseRulesActionSkip             PhaseUpdateResponseRulesAction = "skip"
	PhaseUpdateResponseRulesActionSetCacheSettings PhaseUpdateResponseRulesAction = "set_cache_settings"
)

func (r PhaseUpdateResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesActionBlock, PhaseUpdateResponseRulesActionChallenge, PhaseUpdateResponseRulesActionCompressResponse, PhaseUpdateResponseRulesActionExecute, PhaseUpdateResponseRulesActionJsChallenge, PhaseUpdateResponseRulesActionLog, PhaseUpdateResponseRulesActionManagedChallenge, PhaseUpdateResponseRulesActionRedirect, PhaseUpdateResponseRulesActionRewrite, PhaseUpdateResponseRulesActionRoute, PhaseUpdateResponseRulesActionScore, PhaseUpdateResponseRulesActionServeError, PhaseUpdateResponseRulesActionSetConfig, PhaseUpdateResponseRulesActionSkip, PhaseUpdateResponseRulesActionSetCacheSettings:
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

type PhaseGetResponseRule struct {
	// The action to perform when the rule matches.
	Action           PhaseGetResponseRulesAction `json:"action"`
	ActionParameters interface{}                 `json:"action_parameters,required"`
	Categories       interface{}                 `json:"categories,required"`
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
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                   `json:"version,required"`
	JSON    phaseGetResponseRuleJSON `json:"-"`
	union   PhaseGetResponseRulesUnion
}

// phaseGetResponseRuleJSON contains the JSON metadata for the struct
// [PhaseGetResponseRule]
type phaseGetResponseRuleJSON struct {
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

func (r phaseGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseGetResponseRule) AsUnion() PhaseGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule],
// [rulesets.PhaseGetResponseRulesRulesetsChallengeRule],
// [rulesets.PhaseGetResponseRulesRulesetsCompressResponseRule],
// [rulesets.ExecuteRule], [rulesets.PhaseGetResponseRulesRulesetsJsChallengeRule],
// [rulesets.LogRule],
// [rulesets.PhaseGetResponseRulesRulesetsManagedChallengeRule],
// [rulesets.PhaseGetResponseRulesRulesetsRedirectRule],
// [rulesets.PhaseGetResponseRulesRulesetsRewriteRule],
// [rulesets.PhaseGetResponseRulesRulesetsRouteRule],
// [rulesets.PhaseGetResponseRulesRulesetsScoreRule],
// [rulesets.PhaseGetResponseRulesRulesetsServeErrorRule],
// [rulesets.PhaseGetResponseRulesRulesetsSetConfigRule], [rulesets.SkipRule] or
// [rulesets.PhaseGetResponseRulesRulesetsSetCacheSettingsRule].
type PhaseGetResponseRulesUnion interface {
	implementsRulesetsPhaseGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsCompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsJsChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsRedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsRewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsRouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsSetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
	)
}

type PhaseGetResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsChallengeRuleAction `json:"action"`
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
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                         `json:"ref"`
	JSON phaseGetResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [PhaseGetResponseRulesRulesetsChallengeRule]
type phaseGetResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsChallengeRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsChallengeRuleAction string

const (
	PhaseGetResponseRulesRulesetsChallengeRuleActionChallenge PhaseGetResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r PhaseGetResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

type PhaseGetResponseRulesRulesetsCompressResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsCompressResponseRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsCompressResponseRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                `json:"ref"`
	JSON phaseGetResponseRulesRulesetsCompressResponseRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsCompressResponseRuleJSON contains the JSON metadata
// for the struct [PhaseGetResponseRulesRulesetsCompressResponseRule]
type phaseGetResponseRulesRulesetsCompressResponseRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsCompressResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsCompressResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsCompressResponseRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsCompressResponseRuleAction string

const (
	PhaseGetResponseRulesRulesetsCompressResponseRuleActionCompressResponse PhaseGetResponseRulesRulesetsCompressResponseRuleAction = "compress_response"
)

func (r PhaseGetResponseRulesRulesetsCompressResponseRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsCompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsCompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms []PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm `json:"algorithms"`
	JSON       phaseGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON        `json:"-"`
}

// phaseGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsCompressResponseRuleActionParameters]
type phaseGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON struct {
	Algorithms  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsCompressResponseRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Compression algorithm to enable.
type PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName `json:"name"`
	JSON phaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON  `json:"-"`
}

// phaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm]
type phaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON) RawJSON() string {
	return r.raw
}

// Name of compression algorithm to enable.
type PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName string

const (
	PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone    PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "none"
	PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto    PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "auto"
	PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "default"
	PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip    PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "gzip"
	PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli  PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "brotli"
)

func (r PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone, PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto, PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault, PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip, PhaseGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli:
		return true
	}
	return false
}

type PhaseGetResponseRulesRulesetsJsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsJsChallengeRuleAction `json:"action"`
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
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON phaseGetResponseRulesRulesetsJsChallengeRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsJsChallengeRuleJSON contains the JSON metadata for
// the struct [PhaseGetResponseRulesRulesetsJsChallengeRule]
type phaseGetResponseRulesRulesetsJsChallengeRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsJsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsJsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsJsChallengeRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsJsChallengeRuleAction string

const (
	PhaseGetResponseRulesRulesetsJsChallengeRuleActionJsChallenge PhaseGetResponseRulesRulesetsJsChallengeRuleAction = "js_challenge"
)

func (r PhaseGetResponseRulesRulesetsJsChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsJsChallengeRuleActionJsChallenge:
		return true
	}
	return false
}

type PhaseGetResponseRulesRulesetsManagedChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsManagedChallengeRuleAction `json:"action"`
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
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                `json:"ref"`
	JSON phaseGetResponseRulesRulesetsManagedChallengeRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsManagedChallengeRuleJSON contains the JSON metadata
// for the struct [PhaseGetResponseRulesRulesetsManagedChallengeRule]
type phaseGetResponseRulesRulesetsManagedChallengeRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsManagedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsManagedChallengeRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsManagedChallengeRuleAction string

const (
	PhaseGetResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge PhaseGetResponseRulesRulesetsManagedChallengeRuleAction = "managed_challenge"
)

func (r PhaseGetResponseRulesRulesetsManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

type PhaseGetResponseRulesRulesetsRedirectRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsRedirectRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsRedirectRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                        `json:"ref"`
	JSON phaseGetResponseRulesRulesetsRedirectRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRedirectRuleJSON contains the JSON metadata for the
// struct [PhaseGetResponseRulesRulesetsRedirectRule]
type phaseGetResponseRulesRulesetsRedirectRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsRedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRedirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsRedirectRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsRedirectRuleAction string

const (
	PhaseGetResponseRulesRulesetsRedirectRuleActionRedirect PhaseGetResponseRulesRulesetsRedirectRuleAction = "redirect"
)

func (r PhaseGetResponseRulesRulesetsRedirectRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsRedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsRedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromList `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValue `json:"from_value"`
	JSON      phaseGetResponseRulesRulesetsRedirectRuleActionParametersJSON      `json:"-"`
}

// phaseGetResponseRulesRulesetsRedirectRuleActionParametersJSON contains the JSON
// metadata for the struct
// [PhaseGetResponseRulesRulesetsRedirectRuleActionParameters]
type phaseGetResponseRulesRulesetsRedirectRuleActionParametersJSON struct {
	FromList    apijson.Field
	FromValue   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRedirectRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRedirectRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on a bulk list lookup.
type PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key string `json:"key"`
	// The name of the list to match against.
	Name string                                                                `json:"name"`
	JSON phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromList]
type phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON struct {
	Key         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on the request properties.
type PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString bool `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL `json:"target_url"`
	JSON      phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON      `json:"-"`
}

// phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValue]
type phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON struct {
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	TargetURL           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON) RawJSON() string {
	return r.raw
}

// The status code to be used for the redirect.
type PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode float64

const (
	PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301 PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 301
	PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302 PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 302
	PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303 PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 303
	PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307 PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 307
	PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308 PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 308
)

func (r PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301, PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302, PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303, PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307, PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL struct {
	// The URL to redirect the request to.
	Value string `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                          `json:"expression"`
	JSON       phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON `json:"-"`
	union      PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion
}

// phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL]
type phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) AsUnion() PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion {
	return r.union
}

// The URL to redirect the request to.
//
// Union satisfied by
// [rulesets.PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
// or
// [rulesets.PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect].
type PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRulesetsPhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect{}),
		},
	)
}

type PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value string                                                                                           `json:"value"`
	JSON  phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
type phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRulesetsPhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                                            `json:"expression"`
	JSON       phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect]
type phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRulesetsPhaseGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type PhaseGetResponseRulesRulesetsRewriteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsRewriteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsRewriteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                       `json:"ref"`
	JSON phaseGetResponseRulesRulesetsRewriteRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRewriteRuleJSON contains the JSON metadata for the
// struct [PhaseGetResponseRulesRulesetsRewriteRule]
type phaseGetResponseRulesRulesetsRewriteRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsRewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsRewriteRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsRewriteRuleAction string

const (
	PhaseGetResponseRulesRulesetsRewriteRuleActionRewrite PhaseGetResponseRulesRulesetsRewriteRuleAction = "rewrite"
)

func (r PhaseGetResponseRulesRulesetsRewriteRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsRewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsRewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers map[string]PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeader `json:"headers"`
	// URI to rewrite the request to.
	URI  PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURI  `json:"uri"`
	JSON phaseGetResponseRulesRulesetsRewriteRuleActionParametersJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRewriteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [PhaseGetResponseRulesRulesetsRewriteRuleActionParameters]
type phaseGetResponseRulesRulesetsRewriteRuleActionParametersJSON struct {
	Headers     apijson.Field
	URI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRewriteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRewriteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Remove the header from the request.
type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeader struct {
	Operation PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation `json:"operation,required"`
	// Static value for the header.
	Value string `json:"value"`
	// Expression for the header value.
	Expression string                                                             `json:"expression"`
	JSON       phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON `json:"-"`
	union      PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion
}

// phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON contains the
// JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeader]
type phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeader) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeader) AsUnion() PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion {
	return r.union
}

// Remove the header from the request.
//
// Union satisfied by
// [rulesets.PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader],
// [rulesets.PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
// or
// [rulesets.PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader].
type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion interface {
	implementsRulesetsPhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeader()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader{}),
		},
	)
}

// Remove the header from the request.
type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation `json:"operation,required"`
	JSON      phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON      `json:"-"`
}

// phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader]
type phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON struct {
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) implementsRulesetsPhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Set a request header with a static value.
type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader struct {
	Operation PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation `json:"operation,required"`
	// Static value for the header.
	Value string                                                                          `json:"value,required"`
	JSON  phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
type phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) implementsRulesetsPhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation string

const (
	PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation = "set"
)

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader struct {
	// Expression for the header value.
	Expression string                                                                                `json:"expression,required"`
	Operation  PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation `json:"operation,required"`
	JSON       phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON      `json:"-"`
}

// phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader]
type phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) implementsRulesetsPhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation string

const (
	PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation = "set"
)

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet:
		return true
	}
	return false
}

type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation string

const (
	PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "remove"
	PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet    PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "set"
)

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove, PhaseGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPath `json:"path"`
	// Query portion rewrite.
	Query PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery `json:"query"`
	JSON  phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON  `json:"-"`
}

// phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON contains the
// JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURI]
type phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON struct {
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON) RawJSON() string {
	return r.raw
}

// Path portion rewrite.
type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPath struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                              `json:"expression"`
	JSON       phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON `json:"-"`
	union      PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion
}

// phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON contains the
// JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPath]
type phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPath) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPath) AsUnion() PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion {
	return r.union
}

// Path portion rewrite.
//
// Union satisfied by
// [rulesets.PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
// or
// [rulesets.PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue].
type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion interface {
	implementsRulesetsPhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPath()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue{}),
		},
	)
}

type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue struct {
	// Predefined replacement value.
	Value string                                                                         `json:"value,required"`
	JSON  phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
type phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) implementsRulesetsPhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                          `json:"expression,required"`
	JSON       phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue]
type phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) implementsRulesetsPhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

// Query portion rewrite.
type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                               `json:"expression"`
	JSON       phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON `json:"-"`
	union      PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion
}

// phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery]
type phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery) AsUnion() PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion {
	return r.union
}

// Query portion rewrite.
//
// Union satisfied by
// [rulesets.PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
// or
// [rulesets.PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue].
type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion interface {
	implementsRulesetsPhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue{}),
		},
	)
}

type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue struct {
	// Predefined replacement value.
	Value string                                                                          `json:"value,required"`
	JSON  phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
type phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) implementsRulesetsPhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                           `json:"expression,required"`
	JSON       phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue]
type phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) implementsRulesetsPhaseGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type PhaseGetResponseRulesRulesetsRouteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsRouteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsRouteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                     `json:"ref"`
	JSON phaseGetResponseRulesRulesetsRouteRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRouteRuleJSON contains the JSON metadata for the
// struct [PhaseGetResponseRulesRulesetsRouteRule]
type phaseGetResponseRulesRulesetsRouteRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsRouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRouteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsRouteRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsRouteRuleAction string

const (
	PhaseGetResponseRulesRulesetsRouteRuleActionRoute PhaseGetResponseRulesRulesetsRouteRuleAction = "route"
)

func (r PhaseGetResponseRulesRulesetsRouteRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsRouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsRouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader string `json:"host_header"`
	// Override the IP/TCP destination.
	Origin PhaseGetResponseRulesRulesetsRouteRuleActionParametersOrigin `json:"origin"`
	// Override the Server Name Indication (SNI).
	Sni  PhaseGetResponseRulesRulesetsRouteRuleActionParametersSni  `json:"sni"`
	JSON phaseGetResponseRulesRulesetsRouteRuleActionParametersJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRouteRuleActionParametersJSON contains the JSON
// metadata for the struct [PhaseGetResponseRulesRulesetsRouteRuleActionParameters]
type phaseGetResponseRulesRulesetsRouteRuleActionParametersJSON struct {
	HostHeader  apijson.Field
	Origin      apijson.Field
	Sni         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRouteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRouteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Override the IP/TCP destination.
type PhaseGetResponseRulesRulesetsRouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host string `json:"host"`
	// Override the destination port.
	Port float64                                                          `json:"port"`
	JSON phaseGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON contains the
// JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsRouteRuleActionParametersOrigin]
type phaseGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON struct {
	Host        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRouteRuleActionParametersOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON) RawJSON() string {
	return r.raw
}

// Override the Server Name Indication (SNI).
type PhaseGetResponseRulesRulesetsRouteRuleActionParametersSni struct {
	// The SNI override.
	Value string                                                        `json:"value,required"`
	JSON  phaseGetResponseRulesRulesetsRouteRuleActionParametersSniJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsRouteRuleActionParametersSniJSON contains the JSON
// metadata for the struct
// [PhaseGetResponseRulesRulesetsRouteRuleActionParametersSni]
type phaseGetResponseRulesRulesetsRouteRuleActionParametersSniJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsRouteRuleActionParametersSni) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsRouteRuleActionParametersSniJSON) RawJSON() string {
	return r.raw
}

type PhaseGetResponseRulesRulesetsScoreRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsScoreRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsScoreRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                     `json:"ref"`
	JSON phaseGetResponseRulesRulesetsScoreRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsScoreRuleJSON contains the JSON metadata for the
// struct [PhaseGetResponseRulesRulesetsScoreRule]
type phaseGetResponseRulesRulesetsScoreRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsScoreRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsScoreRuleAction string

const (
	PhaseGetResponseRulesRulesetsScoreRuleActionScore PhaseGetResponseRulesRulesetsScoreRuleAction = "score"
)

func (r PhaseGetResponseRulesRulesetsScoreRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment int64                                                      `json:"increment"`
	JSON      phaseGetResponseRulesRulesetsScoreRuleActionParametersJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsScoreRuleActionParametersJSON contains the JSON
// metadata for the struct [PhaseGetResponseRulesRulesetsScoreRuleActionParameters]
type phaseGetResponseRulesRulesetsScoreRuleActionParametersJSON struct {
	Increment   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsScoreRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsScoreRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

type PhaseGetResponseRulesRulesetsServeErrorRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsServeErrorRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsServeErrorRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                          `json:"ref"`
	JSON phaseGetResponseRulesRulesetsServeErrorRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsServeErrorRuleJSON contains the JSON metadata for
// the struct [PhaseGetResponseRulesRulesetsServeErrorRule]
type phaseGetResponseRulesRulesetsServeErrorRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsServeErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsServeErrorRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsServeErrorRuleAction string

const (
	PhaseGetResponseRulesRulesetsServeErrorRuleActionServeError PhaseGetResponseRulesRulesetsServeErrorRuleAction = "serve_error"
)

func (r PhaseGetResponseRulesRulesetsServeErrorRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsServeErrorRuleActionParameters struct {
	// Error response content.
	Content string `json:"content"`
	// Content-type header to set with the response.
	ContentType PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentType `json:"content_type"`
	// The status code to use for the error.
	StatusCode float64                                                         `json:"status_code"`
	JSON       phaseGetResponseRulesRulesetsServeErrorRuleActionParametersJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsServeErrorRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsServeErrorRuleActionParameters]
type phaseGetResponseRulesRulesetsServeErrorRuleActionParametersJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsServeErrorRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsServeErrorRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Content-type header to set with the response.
type PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentType string

const (
	PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "application/json"
	PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml         PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/xml"
	PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain       PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/plain"
	PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML        PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/html"
)

func (r PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson, PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml, PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain, PhaseGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

type PhaseGetResponseRulesRulesetsSetConfigRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsSetConfigRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsSetConfigRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                         `json:"ref"`
	JSON phaseGetResponseRulesRulesetsSetConfigRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetConfigRuleJSON contains the JSON metadata for
// the struct [PhaseGetResponseRulesRulesetsSetConfigRule]
type phaseGetResponseRulesRulesetsSetConfigRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsSetConfigRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetConfigRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetConfigRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsSetConfigRuleAction string

const (
	PhaseGetResponseRulesRulesetsSetConfigRuleActionSetConfig PhaseGetResponseRulesRulesetsSetConfigRuleAction = "set_config"
)

func (r PhaseGetResponseRulesRulesetsSetConfigRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsSetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites bool `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify `json:"autominify"`
	// Turn on or off Browser Integrity Check.
	Bic bool `json:"bic"`
	// Turn off all active Cloudflare Apps.
	DisableApps bool `json:"disable_apps"`
	// Turn off Zaraz.
	DisableZaraz bool `json:"disable_zaraz"`
	// Turn on or off Email Obfuscation.
	EmailObfuscation bool `json:"email_obfuscation"`
	// Turn on or off the Hotlink Protection.
	HotlinkProtection bool `json:"hotlink_protection"`
	// Turn on or off Mirage.
	Mirage bool `json:"mirage"`
	// Turn on or off Opportunistic Encryption.
	OpportunisticEncryption bool `json:"opportunistic_encryption"`
	// Configure the Polish level.
	Polish PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersPolish `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader bool `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes bool `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSL `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg  bool                                                           `json:"sxg"`
	JSON phaseGetResponseRulesRulesetsSetConfigRuleActionParametersJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetConfigRuleActionParametersJSON contains the JSON
// metadata for the struct
// [PhaseGetResponseRulesRulesetsSetConfigRuleActionParameters]
type phaseGetResponseRulesRulesetsSetConfigRuleActionParametersJSON struct {
	AutomaticHTTPSRewrites  apijson.Field
	Autominify              apijson.Field
	Bic                     apijson.Field
	DisableApps             apijson.Field
	DisableZaraz            apijson.Field
	EmailObfuscation        apijson.Field
	HotlinkProtection       apijson.Field
	Mirage                  apijson.Field
	OpportunisticEncryption apijson.Field
	Polish                  apijson.Field
	RocketLoader            apijson.Field
	SecurityLevel           apijson.Field
	ServerSideExcludes      apijson.Field
	SSL                     apijson.Field
	Sxg                     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetConfigRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetConfigRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Select which file extensions to minify automatically.
type PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	Css bool `json:"css"`
	// Minify HTML files.
	HTML bool `json:"html"`
	// Minify JS files.
	Js   bool                                                                     `json:"js"`
	JSON phaseGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify]
type phaseGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON) RawJSON() string {
	return r.raw
}

// Configure the Polish level.
type PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersPolish string

const (
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersPolishOff      PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "off"
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossless"
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy    PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossy"
)

func (r PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersPolishOff, PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless, PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy:
		return true
	}
	return false
}

// Configure the Security Level.
type PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel string

const (
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff            PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "off"
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "essentially_off"
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow            PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "low"
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium         PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "medium"
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh           PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "high"
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack    PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff, PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff, PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow, PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium, PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh, PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSL string

const (
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOff        PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "off"
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible   PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "flexible"
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFull       PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "full"
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict     PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "strict"
	PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "origin_pull"
)

func (r PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOff, PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible, PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFull, PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict, PhaseGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

type PhaseGetResponseRulesRulesetsSetCacheSettingsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsSetCacheSettingsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                `json:"ref"`
	JSON phaseGetResponseRulesRulesetsSetCacheSettingsRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleJSON contains the JSON metadata
// for the struct [PhaseGetResponseRulesRulesetsSetCacheSettingsRule]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheSettingsRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleAction string

const (
	PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings PhaseGetResponseRulesRulesetsSetCacheSettingsRuleAction = "set_cache_settings"
)

func (r PhaseGetResponseRulesRulesetsSetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts []int64 `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache bool `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL `json:"edge_ttl"`
	// When enabled, Cloudflare will aim to strictly adhere to RFC 7234.
	OriginCacheControl bool `json:"origin_cache_control"`
	// Generate Cloudflare error pages from issues sent from the origin server. When
	// on, error pages will trigger for issues from the origin
	OriginErrorPagePassthru bool `json:"origin_error_page_passthru"`
	// Define a timeout value between two successive read operations to your origin
	// server. Historically, the timeout value between two read options from Cloudflare
	// to an origin server is 100 seconds. If you are attempting to reduce HTTP 524
	// errors because of timeouts from an origin server, try increasing this timeout
	// value.
	ReadTimeout int64 `json:"read_timeout"`
	// Specify whether or not Cloudflare should respect strong ETag (entity tag)
	// headers. When off, Cloudflare converts strong ETag headers to weak ETag headers.
	RespectStrongEtags bool `json:"respect_strong_etags"`
	// Define if Cloudflare should serve stale content while getting the latest content
	// from the origin. If on, Cloudflare will not serve stale content while getting
	// the latest content from the origin.
	ServeStale PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale `json:"serve_stale"`
	JSON       phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON       `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON struct {
	AdditionalCacheablePorts apijson.Field
	BrowserTTL               apijson.Field
	Cache                    apijson.Field
	CacheKey                 apijson.Field
	CacheReserve             apijson.Field
	EdgeTTL                  apijson.Field
	OriginCacheControl       apijson.Field
	OriginErrorPagePassthru  apijson.Field
	ReadTimeout              apijson.Field
	RespectStrongEtags       apijson.Field
	ServeStale               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64                                                                           `json:"default"`
	JSON    phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON struct {
	Mode        apijson.Field
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON) RawJSON() string {
	return r.raw
}

// Determines which browser ttl mode to use.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
)

func (r PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType bool `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor bool `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder bool                                                                          `json:"ignore_query_strings_order"`
	JSON                    phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON struct {
	CacheByDeviceType       apijson.Field
	CacheDeceptionArmor     apijson.Field
	CustomKey               apijson.Field
	IgnoreQueryStringsOrder apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON) RawJSON() string {
	return r.raw
}

// Customize which components of the request are included or excluded from the
// cache key.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser `json:"user"`
	JSON phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON) RawJSON() string {
	return r.raw
}

// The cookies to include in building the cache key.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Include these cookies' names and their values.
	Include []string                                                                                     `json:"include"`
	JSON    phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON) RawJSON() string {
	return r.raw
}

// The header names and values to include in building the cache key.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin bool `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include []string                                                                                     `json:"include"`
	JSON    phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON struct {
	CheckPresence apijson.Field
	ExcludeOrigin apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON) RawJSON() string {
	return r.raw
}

// Whether to use the original host or the resolved host in the cache key.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved bool                                                                                       `json:"resolved"`
	JSON     phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON) RawJSON() string {
	return r.raw
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude `json:"include"`
	JSON    phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON    `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON) RawJSON() string {
	return r.raw
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Exclude all query string parameters from use in building the cache key.
	All bool `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List []string                                                                                                 `json:"list"`
	JSON phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON) RawJSON() string {
	return r.raw
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Use all query string parameters in the cache key.
	All bool `json:"all"`
	// A list of query string parameters used to build the cache key.
	List []string                                                                                                 `json:"list"`
	JSON phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON) RawJSON() string {
	return r.raw
}

// Characteristics of the request user agent used in building the cache key.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType bool `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo bool `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang bool                                                                                       `json:"lang"`
	JSON phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON) RawJSON() string {
	return r.raw
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible bool `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize int64                                                                             `json:"min_file_size,required"`
	JSON        phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON struct {
	Eligible    apijson.Field
	MinFileSize apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON) RawJSON() string {
	return r.raw
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64 `json:"default,required"`
	// edge ttl options
	Mode PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL []PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL `json:"status_code_ttl,required"`
	JSON          phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON            `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON struct {
	Default       apijson.Field
	Mode          apijson.Field
	StatusCodeTTL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON) RawJSON() string {
	return r.raw
}

// edge ttl options
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value int64 `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue int64                                                                                     `json:"status_code_value"`
	JSON            phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON struct {
	Value           apijson.Field
	StatusCodeRange apijson.Field
	StatusCodeValue apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON) RawJSON() string {
	return r.raw
}

// The range of status codes used to apply the selected mode.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// response status code lower bound
	From int64 `json:"from,required"`
	// response status code upper bound
	To   int64                                                                                                    `json:"to,required"`
	JSON phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON) RawJSON() string {
	return r.raw
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating bool                                                                            `json:"disable_stale_while_updating,required"`
	JSON                      phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale]
type phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON struct {
	DisableStaleWhileUpdating apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type PhaseGetResponseRulesAction string

const (
	PhaseGetResponseRulesActionBlock            PhaseGetResponseRulesAction = "block"
	PhaseGetResponseRulesActionChallenge        PhaseGetResponseRulesAction = "challenge"
	PhaseGetResponseRulesActionCompressResponse PhaseGetResponseRulesAction = "compress_response"
	PhaseGetResponseRulesActionExecute          PhaseGetResponseRulesAction = "execute"
	PhaseGetResponseRulesActionJsChallenge      PhaseGetResponseRulesAction = "js_challenge"
	PhaseGetResponseRulesActionLog              PhaseGetResponseRulesAction = "log"
	PhaseGetResponseRulesActionManagedChallenge PhaseGetResponseRulesAction = "managed_challenge"
	PhaseGetResponseRulesActionRedirect         PhaseGetResponseRulesAction = "redirect"
	PhaseGetResponseRulesActionRewrite          PhaseGetResponseRulesAction = "rewrite"
	PhaseGetResponseRulesActionRoute            PhaseGetResponseRulesAction = "route"
	PhaseGetResponseRulesActionScore            PhaseGetResponseRulesAction = "score"
	PhaseGetResponseRulesActionServeError       PhaseGetResponseRulesAction = "serve_error"
	PhaseGetResponseRulesActionSetConfig        PhaseGetResponseRulesAction = "set_config"
	PhaseGetResponseRulesActionSkip             PhaseGetResponseRulesAction = "skip"
	PhaseGetResponseRulesActionSetCacheSettings PhaseGetResponseRulesAction = "set_cache_settings"
)

func (r PhaseGetResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesActionBlock, PhaseGetResponseRulesActionChallenge, PhaseGetResponseRulesActionCompressResponse, PhaseGetResponseRulesActionExecute, PhaseGetResponseRulesActionJsChallenge, PhaseGetResponseRulesActionLog, PhaseGetResponseRulesActionManagedChallenge, PhaseGetResponseRulesActionRedirect, PhaseGetResponseRulesActionRewrite, PhaseGetResponseRulesActionRoute, PhaseGetResponseRulesActionScore, PhaseGetResponseRulesActionServeError, PhaseGetResponseRulesActionSetConfig, PhaseGetResponseRulesActionSkip, PhaseGetResponseRulesActionSetCacheSettings:
		return true
	}
	return false
}

type PhaseUpdateParams struct {
	// The list of rules in the ruleset.
	Rules param.Field[[]PhaseUpdateParamsRuleUnion] `json:"rules,required"`
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

type PhaseUpdateParamsRule struct {
	// The action to perform when the rule matches.
	Action           param.Field[PhaseUpdateParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                  `json:"action_parameters,required"`
	Categories       param.Field[interface{}]                  `json:"categories,required"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam],
// [rulesets.PhaseUpdateParamsRulesRulesetsChallengeRule],
// [rulesets.PhaseUpdateParamsRulesRulesetsCompressResponseRule],
// [rulesets.ExecuteRuleParam],
// [rulesets.PhaseUpdateParamsRulesRulesetsJsChallengeRule],
// [rulesets.LogRuleParam],
// [rulesets.PhaseUpdateParamsRulesRulesetsManagedChallengeRule],
// [rulesets.PhaseUpdateParamsRulesRulesetsRedirectRule],
// [rulesets.PhaseUpdateParamsRulesRulesetsRewriteRule],
// [rulesets.PhaseUpdateParamsRulesRulesetsRouteRule],
// [rulesets.PhaseUpdateParamsRulesRulesetsScoreRule],
// [rulesets.PhaseUpdateParamsRulesRulesetsServeErrorRule],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetConfigRule],
// [rulesets.SkipRuleParam],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheSettingsRule],
// [PhaseUpdateParamsRule].
type PhaseUpdateParamsRuleUnion interface {
	implementsRulesetsPhaseUpdateParamsRuleUnion()
}

type PhaseUpdateParamsRulesRulesetsChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsChallengeRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsChallengeRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsChallengeRuleActionChallenge PhaseUpdateParamsRulesRulesetsChallengeRuleAction = "challenge"
)

func (r PhaseUpdateParamsRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

type PhaseUpdateParamsRulesRulesetsCompressResponseRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsCompressResponseRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsCompressResponseRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsCompressResponseRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsCompressResponseRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionCompressResponse PhaseUpdateParamsRulesRulesetsCompressResponseRuleAction = "compress_response"
)

func (r PhaseUpdateParamsRulesRulesetsCompressResponseRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms param.Field[[]PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithm] `json:"algorithms"`
}

func (r PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Compression algorithm to enable.
type PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name param.Field[PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName] `json:"name"`
}

func (r PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithm) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Name of compression algorithm to enable.
type PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName string

const (
	PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone    PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "none"
	PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto    PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "auto"
	PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "default"
	PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip    PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "gzip"
	PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli  PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "brotli"
)

func (r PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone, PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto, PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault, PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip, PhaseUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli:
		return true
	}
	return false
}

type PhaseUpdateParamsRulesRulesetsJsChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsJsChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsJsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsJsChallengeRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsJsChallengeRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsJsChallengeRuleActionJsChallenge PhaseUpdateParamsRulesRulesetsJsChallengeRuleAction = "js_challenge"
)

func (r PhaseUpdateParamsRulesRulesetsJsChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsJsChallengeRuleActionJsChallenge:
		return true
	}
	return false
}

type PhaseUpdateParamsRulesRulesetsManagedChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsManagedChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsManagedChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsManagedChallengeRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsManagedChallengeRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsManagedChallengeRuleActionManagedChallenge PhaseUpdateParamsRulesRulesetsManagedChallengeRuleAction = "managed_challenge"
)

func (r PhaseUpdateParamsRulesRulesetsManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

type PhaseUpdateParamsRulesRulesetsRedirectRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsRedirectRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsRedirectRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsRedirectRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRedirectRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsRedirectRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsRedirectRuleActionRedirect PhaseUpdateParamsRulesRulesetsRedirectRuleAction = "redirect"
)

func (r PhaseUpdateParamsRulesRulesetsRedirectRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsRedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsRedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList param.Field[PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromList] `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue param.Field[PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValue] `json:"from_value"`
}

func (r PhaseUpdateParamsRulesRulesetsRedirectRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Serve a redirect based on a bulk list lookup.
type PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key param.Field[string] `json:"key"`
	// The name of the list to match against.
	Name param.Field[string] `json:"name"`
}

func (r PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Serve a redirect based on the request properties.
type PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString param.Field[bool] `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode param.Field[PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL param.Field[PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion] `json:"target_url"`
}

func (r PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status code to be used for the redirect.
type PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode float64

const (
	PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301 PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 301
	PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302 PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 302
	PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303 PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 303
	PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307 PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 307
	PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308 PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 308
)

func (r PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301, PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302, PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303, PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307, PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURL struct {
	// The URL to redirect the request to.
	Value param.Field[string] `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression param.Field[string] `json:"expression"`
}

func (r PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) implementsRulesetsPhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion() {
}

// The URL to redirect the request to.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect],
// [rulesets.PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect],
// [PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURL].
type PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRulesetsPhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion()
}

type PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value param.Field[string] `json:"value"`
}

func (r PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRulesetsPhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion() {
}

type PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression param.Field[string] `json:"expression"`
}

func (r PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRulesetsPhaseUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion() {
}

type PhaseUpdateParamsRulesRulesetsRewriteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsRewriteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsRewriteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsRewriteRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsRewriteRuleActionRewrite PhaseUpdateParamsRulesRulesetsRewriteRuleAction = "rewrite"
)

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsRewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers param.Field[map[string]PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion] `json:"headers"`
	// URI to rewrite the request to.
	URI param.Field[PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURI] `json:"uri"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Remove the header from the request.
type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeaders struct {
	Operation param.Field[PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation] `json:"operation,required"`
	// Static value for the header.
	Value param.Field[string] `json:"value"`
	// Expression for the header value.
	Expression param.Field[string] `json:"expression"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeaders) implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion() {
}

// Remove the header from the request.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader],
// [rulesets.PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader],
// [rulesets.PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader],
// [PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeaders].
type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion interface {
	implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion()
}

// Remove the header from the request.
type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation param.Field[PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation] `json:"operation,required"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion() {
}

type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Set a request header with a static value.
type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader struct {
	Operation param.Field[PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation] `json:"operation,required"`
	// Static value for the header.
	Value param.Field[string] `json:"value,required"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion() {
}

type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation string

const (
	PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation = "set"
)

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader struct {
	// Expression for the header value.
	Expression param.Field[string]                                                                                 `json:"expression,required"`
	Operation  param.Field[PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation] `json:"operation,required"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion() {
}

type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation string

const (
	PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation = "set"
)

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet:
		return true
	}
	return false
}

type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation string

const (
	PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation = "remove"
	PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperationSet    PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation = "set"
)

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove, PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path param.Field[PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion] `json:"path"`
	// Query portion rewrite.
	Query param.Field[PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion] `json:"query"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Path portion rewrite.
type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPath struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPath) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPath) implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion() {
}

// Path portion rewrite.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathStaticValue],
// [rulesets.PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue],
// [PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPath].
type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion interface {
	implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion()
}

type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathStaticValue struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value,required"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion() {
}

type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression,required"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion() {
}

// Query portion rewrite.
type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQuery struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQuery) implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion() {
}

// Query portion rewrite.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue],
// [rulesets.PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue],
// [PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQuery].
type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion interface {
	implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion()
}

type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value,required"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion() {
}

type PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression,required"`
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) implementsRulesetsPhaseUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion() {
}

type PhaseUpdateParamsRulesRulesetsRouteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsRouteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsRouteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsRouteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsRouteRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsRouteRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsRouteRuleActionRoute PhaseUpdateParamsRulesRulesetsRouteRuleAction = "route"
)

func (r PhaseUpdateParamsRulesRulesetsRouteRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsRouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsRouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader param.Field[string] `json:"host_header"`
	// Override the IP/TCP destination.
	Origin param.Field[PhaseUpdateParamsRulesRulesetsRouteRuleActionParametersOrigin] `json:"origin"`
	// Override the Server Name Indication (SNI).
	Sni param.Field[PhaseUpdateParamsRulesRulesetsRouteRuleActionParametersSni] `json:"sni"`
}

func (r PhaseUpdateParamsRulesRulesetsRouteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Override the IP/TCP destination.
type PhaseUpdateParamsRulesRulesetsRouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host param.Field[string] `json:"host"`
	// Override the destination port.
	Port param.Field[float64] `json:"port"`
}

func (r PhaseUpdateParamsRulesRulesetsRouteRuleActionParametersOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Override the Server Name Indication (SNI).
type PhaseUpdateParamsRulesRulesetsRouteRuleActionParametersSni struct {
	// The SNI override.
	Value param.Field[string] `json:"value,required"`
}

func (r PhaseUpdateParamsRulesRulesetsRouteRuleActionParametersSni) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhaseUpdateParamsRulesRulesetsScoreRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsScoreRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsScoreRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsScoreRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsScoreRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsScoreRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsScoreRuleActionScore PhaseUpdateParamsRulesRulesetsScoreRuleAction = "score"
)

func (r PhaseUpdateParamsRulesRulesetsScoreRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment param.Field[int64] `json:"increment"`
}

func (r PhaseUpdateParamsRulesRulesetsScoreRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhaseUpdateParamsRulesRulesetsServeErrorRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsServeErrorRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsServeErrorRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsServeErrorRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsServeErrorRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsServeErrorRuleActionServeError PhaseUpdateParamsRulesRulesetsServeErrorRuleAction = "serve_error"
)

func (r PhaseUpdateParamsRulesRulesetsServeErrorRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParameters struct {
	// Error response content.
	Content param.Field[string] `json:"content"`
	// Content-type header to set with the response.
	ContentType param.Field[PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType] `json:"content_type"`
	// The status code to use for the error.
	StatusCode param.Field[float64] `json:"status_code"`
}

func (r PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Content-type header to set with the response.
type PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType string

const (
	PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType = "application/json"
	PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml         PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType = "text/xml"
	PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain       PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType = "text/plain"
	PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML        PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType = "text/html"
)

func (r PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson, PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml, PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain, PhaseUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

type PhaseUpdateParamsRulesRulesetsSetConfigRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsSetConfigRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsSetConfigRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetConfigRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsSetConfigRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionSetConfig PhaseUpdateParamsRulesRulesetsSetConfigRuleAction = "set_config"
)

func (r PhaseUpdateParamsRulesRulesetsSetConfigRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites param.Field[bool] `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify param.Field[PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersAutominify] `json:"autominify"`
	// Turn on or off Browser Integrity Check.
	Bic param.Field[bool] `json:"bic"`
	// Turn off all active Cloudflare Apps.
	DisableApps param.Field[bool] `json:"disable_apps"`
	// Turn off Zaraz.
	DisableZaraz param.Field[bool] `json:"disable_zaraz"`
	// Turn on or off Email Obfuscation.
	EmailObfuscation param.Field[bool] `json:"email_obfuscation"`
	// Turn on or off the Hotlink Protection.
	HotlinkProtection param.Field[bool] `json:"hotlink_protection"`
	// Turn on or off Mirage.
	Mirage param.Field[bool] `json:"mirage"`
	// Turn on or off Opportunistic Encryption.
	OpportunisticEncryption param.Field[bool] `json:"opportunistic_encryption"`
	// Configure the Polish level.
	Polish param.Field[PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolish] `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader param.Field[bool] `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel param.Field[PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel] `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes param.Field[bool] `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL param.Field[PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL] `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg param.Field[bool] `json:"sxg"`
}

func (r PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Select which file extensions to minify automatically.
type PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	Css param.Field[bool] `json:"css"`
	// Minify HTML files.
	HTML param.Field[bool] `json:"html"`
	// Minify JS files.
	Js param.Field[bool] `json:"js"`
}

func (r PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersAutominify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure the Polish level.
type PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolish string

const (
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolishOff      PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolish = "off"
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolishLossless PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolish = "lossless"
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolishLossy    PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolish = "lossy"
)

func (r PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolishOff, PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolishLossless, PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolishLossy:
		return true
	}
	return false
}

// Configure the Security Level.
type PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel string

const (
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff            PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "off"
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "essentially_off"
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow            PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "low"
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium         PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "medium"
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh           PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "high"
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack    PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff, PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff, PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow, PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium, PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh, PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL string

const (
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLOff        PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL = "off"
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLFlexible   PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL = "flexible"
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLFull       PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL = "full"
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLStrict     PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL = "strict"
	PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLOriginPull PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL = "origin_pull"
)

func (r PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLOff, PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLFlexible, PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLFull, PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLStrict, PhaseUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleAction = "set_cache_settings"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts param.Field[[]int64] `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL] `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache param.Field[bool] `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey] `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve] `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL] `json:"edge_ttl"`
	// When enabled, Cloudflare will aim to strictly adhere to RFC 7234.
	OriginCacheControl param.Field[bool] `json:"origin_cache_control"`
	// Generate Cloudflare error pages from issues sent from the origin server. When
	// on, error pages will trigger for issues from the origin
	OriginErrorPagePassthru param.Field[bool] `json:"origin_error_page_passthru"`
	// Define a timeout value between two successive read operations to your origin
	// server. Historically, the timeout value between two read options from Cloudflare
	// to an origin server is 100 seconds. If you are attempting to reduce HTTP 524
	// errors because of timeouts from an origin server, try increasing this timeout
	// value.
	ReadTimeout param.Field[int64] `json:"read_timeout"`
	// Specify whether or not Cloudflare should respect strong ETag (entity tag)
	// headers. When off, Cloudflare converts strong ETag headers to weak ETag headers.
	RespectStrongEtags param.Field[bool] `json:"respect_strong_etags"`
	// Define if Cloudflare should serve stale content while getting the latest content
	// from the origin. If on, Cloudflare will not serve stale content while getting
	// the latest content from the origin.
	ServeStale param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersServeStale] `json:"serve_stale"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode] `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default param.Field[int64] `json:"default"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which browser ttl mode to use.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType param.Field[bool] `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor param.Field[bool] `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey] `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder param.Field[bool] `json:"ignore_query_strings_order"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Customize which components of the request are included or excluded from the
// cache key.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie] `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader] `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost] `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString] `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser] `json:"user"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cookies to include in building the cache key.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// Include these cookies' names and their values.
	Include param.Field[[]string] `json:"include"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The header names and values to include in building the cache key.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin param.Field[bool] `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include param.Field[[]string] `json:"include"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to use the original host or the resolved host in the cache key.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved param.Field[bool] `json:"resolved"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude] `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude] `json:"include"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Exclude all query string parameters from use in building the cache key.
	All param.Field[bool] `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List param.Field[[]string] `json:"list"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Use all query string parameters in the cache key.
	All param.Field[bool] `json:"all"`
	// A list of query string parameters used to build the cache key.
	List param.Field[[]string] `json:"list"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Characteristics of the request user agent used in building the cache key.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType param.Field[bool] `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo param.Field[bool] `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang param.Field[bool] `json:"lang"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible param.Field[bool] `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize param.Field[int64] `json:"min_file_size,required"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default param.Field[int64] `json:"default,required"`
	// edge ttl options
	Mode param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode] `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL param.Field[[]PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL] `json:"status_code_ttl,required"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// edge ttl options
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value param.Field[int64] `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange param.Field[PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange] `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue param.Field[int64] `json:"status_code_value"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The range of status codes used to apply the selected mode.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// response status code lower bound
	From param.Field[int64] `json:"from,required"`
	// response status code upper bound
	To param.Field[int64] `json:"to,required"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating param.Field[bool] `json:"disable_stale_while_updating,required"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersServeStale) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesAction string

const (
	PhaseUpdateParamsRulesActionBlock            PhaseUpdateParamsRulesAction = "block"
	PhaseUpdateParamsRulesActionChallenge        PhaseUpdateParamsRulesAction = "challenge"
	PhaseUpdateParamsRulesActionCompressResponse PhaseUpdateParamsRulesAction = "compress_response"
	PhaseUpdateParamsRulesActionExecute          PhaseUpdateParamsRulesAction = "execute"
	PhaseUpdateParamsRulesActionJsChallenge      PhaseUpdateParamsRulesAction = "js_challenge"
	PhaseUpdateParamsRulesActionLog              PhaseUpdateParamsRulesAction = "log"
	PhaseUpdateParamsRulesActionManagedChallenge PhaseUpdateParamsRulesAction = "managed_challenge"
	PhaseUpdateParamsRulesActionRedirect         PhaseUpdateParamsRulesAction = "redirect"
	PhaseUpdateParamsRulesActionRewrite          PhaseUpdateParamsRulesAction = "rewrite"
	PhaseUpdateParamsRulesActionRoute            PhaseUpdateParamsRulesAction = "route"
	PhaseUpdateParamsRulesActionScore            PhaseUpdateParamsRulesAction = "score"
	PhaseUpdateParamsRulesActionServeError       PhaseUpdateParamsRulesAction = "serve_error"
	PhaseUpdateParamsRulesActionSetConfig        PhaseUpdateParamsRulesAction = "set_config"
	PhaseUpdateParamsRulesActionSkip             PhaseUpdateParamsRulesAction = "skip"
	PhaseUpdateParamsRulesActionSetCacheSettings PhaseUpdateParamsRulesAction = "set_cache_settings"
)

func (r PhaseUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesActionBlock, PhaseUpdateParamsRulesActionChallenge, PhaseUpdateParamsRulesActionCompressResponse, PhaseUpdateParamsRulesActionExecute, PhaseUpdateParamsRulesActionJsChallenge, PhaseUpdateParamsRulesActionLog, PhaseUpdateParamsRulesActionManagedChallenge, PhaseUpdateParamsRulesActionRedirect, PhaseUpdateParamsRulesActionRewrite, PhaseUpdateParamsRulesActionRoute, PhaseUpdateParamsRulesActionScore, PhaseUpdateParamsRulesActionServeError, PhaseUpdateParamsRulesActionSetConfig, PhaseUpdateParamsRulesActionSkip, PhaseUpdateParamsRulesActionSetCacheSettings:
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
