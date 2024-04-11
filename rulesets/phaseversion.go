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
	Logging Logging `json:"logging"`
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

// Union satisfied by [rulesets.BlockRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsChallengeRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsCompressResponseRule],
// [rulesets.ExecuteRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsJsChallengeRule],
// [rulesets.LogRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsManagedChallengeRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsRedirectRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsRewriteRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsRouteRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsScoreRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsServeErrorRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsSetConfigRule],
// [rulesets.SkipRule] or
// [rulesets.PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRule].
type PhaseVersionGetResponseRulesUnion interface {
	implementsRulesetsPhaseVersionGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsCompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsJsChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsRedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsRewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsRouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
	)
}

type PhaseVersionGetResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsChallengeRuleAction `json:"action"`
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
	JSON phaseVersionGetResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseRulesRulesetsChallengeRule]
type phaseVersionGetResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsChallengeRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsChallengeRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsChallengeRuleActionChallenge PhaseVersionGetResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r PhaseVersionGetResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

type PhaseVersionGetResponseRulesRulesetsCompressResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsCompressResponseRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParameters `json:"action_parameters"`
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
	Ref  string                                                       `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsCompressResponseRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsCompressResponseRuleJSON contains the JSON
// metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsCompressResponseRule]
type phaseVersionGetResponseRulesRulesetsCompressResponseRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsCompressResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsCompressResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsCompressResponseRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsCompressResponseRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionCompressResponse PhaseVersionGetResponseRulesRulesetsCompressResponseRuleAction = "compress_response"
)

func (r PhaseVersionGetResponseRulesRulesetsCompressResponseRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms []PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm `json:"algorithms"`
	JSON       phaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON        `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON struct {
	Algorithms  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Compression algorithm to enable.
type PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName `json:"name"`
	JSON phaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON  `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm]
type phaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON) RawJSON() string {
	return r.raw
}

// Name of compression algorithm to enable.
type PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName string

const (
	PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone    PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "none"
	PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto    PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "auto"
	PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "default"
	PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip    PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "gzip"
	PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli  PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "brotli"
)

func (r PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone, PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto, PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault, PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip, PhaseVersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli:
		return true
	}
	return false
}

type PhaseVersionGetResponseRulesRulesetsJsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsJsChallengeRuleAction `json:"action"`
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
	Ref  string                                                  `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsJsChallengeRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsJsChallengeRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseRulesRulesetsJsChallengeRule]
type phaseVersionGetResponseRulesRulesetsJsChallengeRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsJsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsJsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsJsChallengeRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsJsChallengeRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsJsChallengeRuleActionJsChallenge PhaseVersionGetResponseRulesRulesetsJsChallengeRuleAction = "js_challenge"
)

func (r PhaseVersionGetResponseRulesRulesetsJsChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsJsChallengeRuleActionJsChallenge:
		return true
	}
	return false
}

type PhaseVersionGetResponseRulesRulesetsManagedChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsManagedChallengeRuleAction `json:"action"`
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
	Ref  string                                                       `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsManagedChallengeRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsManagedChallengeRuleJSON contains the JSON
// metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsManagedChallengeRule]
type phaseVersionGetResponseRulesRulesetsManagedChallengeRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsManagedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsManagedChallengeRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsManagedChallengeRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge PhaseVersionGetResponseRulesRulesetsManagedChallengeRuleAction = "managed_challenge"
)

func (r PhaseVersionGetResponseRulesRulesetsManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

type PhaseVersionGetResponseRulesRulesetsRedirectRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsRedirectRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParameters `json:"action_parameters"`
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
	Ref  string                                               `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsRedirectRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRedirectRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseRulesRulesetsRedirectRule]
type phaseVersionGetResponseRulesRulesetsRedirectRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsRedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRedirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsRedirectRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsRedirectRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsRedirectRuleActionRedirect PhaseVersionGetResponseRulesRulesetsRedirectRuleAction = "redirect"
)

func (r PhaseVersionGetResponseRulesRulesetsRedirectRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsRedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromList `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValue `json:"from_value"`
	JSON      phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersJSON      `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersJSON contains
// the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersJSON struct {
	FromList    apijson.Field
	FromValue   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on a bulk list lookup.
type PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key string `json:"key"`
	// The name of the list to match against.
	Name string                                                                       `json:"name"`
	JSON phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromList]
type phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON struct {
	Key         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on the request properties.
type PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString bool `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL `json:"target_url"`
	JSON      phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON      `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValue]
type phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON struct {
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	TargetURL           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON) RawJSON() string {
	return r.raw
}

// The status code to be used for the redirect.
type PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode float64

const (
	PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301 PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 301
	PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302 PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 302
	PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303 PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 303
	PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307 PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 307
	PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308 PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 308
)

func (r PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301, PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302, PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303, PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307, PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL struct {
	// The URL to redirect the request to.
	Value string `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                                 `json:"expression"`
	JSON       phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON `json:"-"`
	union      PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion
}

// phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL]
type phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) AsUnion() PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion {
	return r.union
}

// The URL to redirect the request to.
//
// Union satisfied by
// [rulesets.PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
// or
// [rulesets.PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect].
type PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRulesetsPhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect{}),
		},
	)
}

type PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value string                                                                                                  `json:"value"`
	JSON  phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
type phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRulesetsPhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                                                   `json:"expression"`
	JSON       phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect]
type phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRulesetsPhaseVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type PhaseVersionGetResponseRulesRulesetsRewriteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsRewriteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParameters `json:"action_parameters"`
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
	JSON phaseVersionGetResponseRulesRulesetsRewriteRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseRulesRulesetsRewriteRule]
type phaseVersionGetResponseRulesRulesetsRewriteRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsRewriteRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsRewriteRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsRewriteRuleActionRewrite PhaseVersionGetResponseRulesRulesetsRewriteRuleAction = "rewrite"
)

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsRewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers map[string]PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader `json:"headers"`
	// URI to rewrite the request to.
	URI  PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURI  `json:"uri"`
	JSON phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersJSON struct {
	Headers     apijson.Field
	URI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Remove the header from the request.
type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader struct {
	Operation PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation `json:"operation,required"`
	// Static value for the header.
	Value string `json:"value"`
	// Expression for the header value.
	Expression string                                                                    `json:"expression"`
	JSON       phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON `json:"-"`
	union      PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader]
type phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader) AsUnion() PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion {
	return r.union
}

// Remove the header from the request.
//
// Union satisfied by
// [rulesets.PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader],
// [rulesets.PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
// or
// [rulesets.PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader].
type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion interface {
	implementsRulesetsPhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader{}),
		},
	)
}

// Remove the header from the request.
type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation `json:"operation,required"`
	JSON      phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON      `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader]
type phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON struct {
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) implementsRulesetsPhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Set a request header with a static value.
type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader struct {
	Operation PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation `json:"operation,required"`
	// Static value for the header.
	Value string                                                                                 `json:"value,required"`
	JSON  phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
type phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) implementsRulesetsPhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation string

const (
	PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation = "set"
)

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader struct {
	// Expression for the header value.
	Expression string                                                                                       `json:"expression,required"`
	Operation  PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation `json:"operation,required"`
	JSON       phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON      `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader]
type phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) implementsRulesetsPhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation string

const (
	PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation = "set"
)

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet:
		return true
	}
	return false
}

type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation string

const (
	PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "remove"
	PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet    PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "set"
)

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove, PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath `json:"path"`
	// Query portion rewrite.
	Query PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery `json:"query"`
	JSON  phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON  `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON contains
// the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURI]
type phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON struct {
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON) RawJSON() string {
	return r.raw
}

// Path portion rewrite.
type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                     `json:"expression"`
	JSON       phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON `json:"-"`
	union      PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath]
type phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath) AsUnion() PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion {
	return r.union
}

// Path portion rewrite.
//
// Union satisfied by
// [rulesets.PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
// or
// [rulesets.PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue].
type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion interface {
	implementsRulesetsPhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue{}),
		},
	)
}

type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue struct {
	// Predefined replacement value.
	Value string                                                                                `json:"value,required"`
	JSON  phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
type phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) implementsRulesetsPhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                                 `json:"expression,required"`
	JSON       phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue]
type phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) implementsRulesetsPhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

// Query portion rewrite.
type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                      `json:"expression"`
	JSON       phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON `json:"-"`
	union      PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery]
type phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery) AsUnion() PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion {
	return r.union
}

// Query portion rewrite.
//
// Union satisfied by
// [rulesets.PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
// or
// [rulesets.PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue].
type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion interface {
	implementsRulesetsPhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue{}),
		},
	)
}

type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue struct {
	// Predefined replacement value.
	Value string                                                                                 `json:"value,required"`
	JSON  phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
type phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) implementsRulesetsPhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                                  `json:"expression,required"`
	JSON       phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue]
type phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) implementsRulesetsPhaseVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type PhaseVersionGetResponseRulesRulesetsRouteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsRouteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsRouteRuleActionParameters `json:"action_parameters"`
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
	JSON phaseVersionGetResponseRulesRulesetsRouteRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRouteRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseRulesRulesetsRouteRule]
type phaseVersionGetResponseRulesRulesetsRouteRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsRouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRouteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsRouteRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsRouteRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsRouteRuleActionRoute PhaseVersionGetResponseRulesRulesetsRouteRuleAction = "route"
)

func (r PhaseVersionGetResponseRulesRulesetsRouteRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsRouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsRouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader string `json:"host_header"`
	// Override the IP/TCP destination.
	Origin PhaseVersionGetResponseRulesRulesetsRouteRuleActionParametersOrigin `json:"origin"`
	// Override the Server Name Indication (SNI).
	Sni  PhaseVersionGetResponseRulesRulesetsRouteRuleActionParametersSni  `json:"sni"`
	JSON phaseVersionGetResponseRulesRulesetsRouteRuleActionParametersJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRouteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRouteRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsRouteRuleActionParametersJSON struct {
	HostHeader  apijson.Field
	Origin      apijson.Field
	Sni         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRouteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRouteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Override the IP/TCP destination.
type PhaseVersionGetResponseRulesRulesetsRouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host string `json:"host"`
	// Override the destination port.
	Port float64                                                                 `json:"port"`
	JSON phaseVersionGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON contains
// the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRouteRuleActionParametersOrigin]
type phaseVersionGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON struct {
	Host        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRouteRuleActionParametersOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON) RawJSON() string {
	return r.raw
}

// Override the Server Name Indication (SNI).
type PhaseVersionGetResponseRulesRulesetsRouteRuleActionParametersSni struct {
	// The SNI override.
	Value string                                                               `json:"value,required"`
	JSON  phaseVersionGetResponseRulesRulesetsRouteRuleActionParametersSniJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsRouteRuleActionParametersSniJSON contains
// the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsRouteRuleActionParametersSni]
type phaseVersionGetResponseRulesRulesetsRouteRuleActionParametersSniJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsRouteRuleActionParametersSni) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsRouteRuleActionParametersSniJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseRulesRulesetsScoreRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsScoreRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsScoreRuleActionParameters `json:"action_parameters"`
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
	JSON phaseVersionGetResponseRulesRulesetsScoreRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsScoreRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseRulesRulesetsScoreRule]
type phaseVersionGetResponseRulesRulesetsScoreRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsScoreRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsScoreRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsScoreRuleActionScore PhaseVersionGetResponseRulesRulesetsScoreRuleAction = "score"
)

func (r PhaseVersionGetResponseRulesRulesetsScoreRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment int64                                                             `json:"increment"`
	JSON      phaseVersionGetResponseRulesRulesetsScoreRuleActionParametersJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsScoreRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsScoreRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsScoreRuleActionParametersJSON struct {
	Increment   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsScoreRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsScoreRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseRulesRulesetsServeErrorRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsServeErrorRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParameters `json:"action_parameters"`
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
	Ref  string                                                 `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsServeErrorRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsServeErrorRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseRulesRulesetsServeErrorRule]
type phaseVersionGetResponseRulesRulesetsServeErrorRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsServeErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsServeErrorRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsServeErrorRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionServeError PhaseVersionGetResponseRulesRulesetsServeErrorRuleAction = "serve_error"
)

func (r PhaseVersionGetResponseRulesRulesetsServeErrorRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParameters struct {
	// Error response content.
	Content string `json:"content"`
	// Content-type header to set with the response.
	ContentType PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType `json:"content_type"`
	// The status code to use for the error.
	StatusCode float64                                                                `json:"status_code"`
	JSON       phaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersJSON contains
// the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Content-type header to set with the response.
type PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType string

const (
	PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "application/json"
	PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml         PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/xml"
	PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain       PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/plain"
	PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML        PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/html"
)

func (r PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson, PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml, PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain, PhaseVersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

type PhaseVersionGetResponseRulesRulesetsSetConfigRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsSetConfigRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParameters `json:"action_parameters"`
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
	JSON phaseVersionGetResponseRulesRulesetsSetConfigRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetConfigRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseRulesRulesetsSetConfigRule]
type phaseVersionGetResponseRulesRulesetsSetConfigRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsSetConfigRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetConfigRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetConfigRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsSetConfigRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionSetConfig PhaseVersionGetResponseRulesRulesetsSetConfigRuleAction = "set_config"
)

func (r PhaseVersionGetResponseRulesRulesetsSetConfigRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites bool `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify `json:"autominify"`
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
	Polish PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolish `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader bool `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes bool `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg  bool                                                                  `json:"sxg"`
	JSON phaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersJSON contains
// the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Select which file extensions to minify automatically.
type PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	Css bool `json:"css"`
	// Minify HTML files.
	HTML bool `json:"html"`
	// Minify JS files.
	Js   bool                                                                            `json:"js"`
	JSON phaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify]
type phaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON) RawJSON() string {
	return r.raw
}

// Configure the Polish level.
type PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolish string

const (
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolishOff      PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "off"
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossless"
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy    PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossy"
)

func (r PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolishOff, PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless, PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy:
		return true
	}
	return false
}

// Configure the Security Level.
type PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel string

const (
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff            PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "off"
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "essentially_off"
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow            PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "low"
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium         PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "medium"
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh           PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "high"
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack    PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff, PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff, PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow, PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium, PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh, PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL string

const (
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOff        PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "off"
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible   PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "flexible"
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFull       PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "full"
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict     PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "strict"
	PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "origin_pull"
)

func (r PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOff, PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible, PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFull, PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict, PhaseVersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters `json:"action_parameters"`
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
	Ref  string                                                       `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleJSON contains the JSON
// metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRule]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleAction = "set_cache_settings"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts []int64 `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache bool `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL `json:"edge_ttl"`
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
	ServeStale PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale `json:"serve_stale"`
	JSON       phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON       `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64                                                                                  `json:"default"`
	JSON    phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON struct {
	Mode        apijson.Field
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON) RawJSON() string {
	return r.raw
}

// Determines which browser ttl mode to use.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType bool `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor bool `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder bool                                                                                 `json:"ignore_query_strings_order"`
	JSON                    phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON struct {
	CacheByDeviceType       apijson.Field
	CacheDeceptionArmor     apijson.Field
	CustomKey               apijson.Field
	IgnoreQueryStringsOrder apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON) RawJSON() string {
	return r.raw
}

// Customize which components of the request are included or excluded from the
// cache key.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser `json:"user"`
	JSON phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON) RawJSON() string {
	return r.raw
}

// The cookies to include in building the cache key.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Include these cookies' names and their values.
	Include []string                                                                                            `json:"include"`
	JSON    phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON) RawJSON() string {
	return r.raw
}

// The header names and values to include in building the cache key.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin bool `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include []string                                                                                            `json:"include"`
	JSON    phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON struct {
	CheckPresence apijson.Field
	ExcludeOrigin apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON) RawJSON() string {
	return r.raw
}

// Whether to use the original host or the resolved host in the cache key.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved bool                                                                                              `json:"resolved"`
	JSON     phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON) RawJSON() string {
	return r.raw
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude `json:"include"`
	JSON    phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON    `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON) RawJSON() string {
	return r.raw
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Exclude all query string parameters from use in building the cache key.
	All bool `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List []string                                                                                                        `json:"list"`
	JSON phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON) RawJSON() string {
	return r.raw
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Use all query string parameters in the cache key.
	All bool `json:"all"`
	// A list of query string parameters used to build the cache key.
	List []string                                                                                                        `json:"list"`
	JSON phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON) RawJSON() string {
	return r.raw
}

// Characteristics of the request user agent used in building the cache key.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType bool `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo bool `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang bool                                                                                              `json:"lang"`
	JSON phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON) RawJSON() string {
	return r.raw
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible bool `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize int64                                                                                    `json:"min_file_size,required"`
	JSON        phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON struct {
	Eligible    apijson.Field
	MinFileSize apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON) RawJSON() string {
	return r.raw
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64 `json:"default,required"`
	// edge ttl options
	Mode PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL []PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL `json:"status_code_ttl,required"`
	JSON          phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON            `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON struct {
	Default       apijson.Field
	Mode          apijson.Field
	StatusCodeTTL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON) RawJSON() string {
	return r.raw
}

// edge ttl options
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value int64 `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue int64                                                                                            `json:"status_code_value"`
	JSON            phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON struct {
	Value           apijson.Field
	StatusCodeRange apijson.Field
	StatusCodeValue apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON) RawJSON() string {
	return r.raw
}

// The range of status codes used to apply the selected mode.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// response status code lower bound
	From int64 `json:"from,required"`
	// response status code upper bound
	To   int64                                                                                                           `json:"to,required"`
	JSON phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON) RawJSON() string {
	return r.raw
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating bool                                                                                   `json:"disable_stale_while_updating,required"`
	JSON                      phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale]
type phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON struct {
	DisableStaleWhileUpdating apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesAction string

const (
	PhaseVersionGetResponseRulesActionBlock            PhaseVersionGetResponseRulesAction = "block"
	PhaseVersionGetResponseRulesActionChallenge        PhaseVersionGetResponseRulesAction = "challenge"
	PhaseVersionGetResponseRulesActionCompressResponse PhaseVersionGetResponseRulesAction = "compress_response"
	PhaseVersionGetResponseRulesActionExecute          PhaseVersionGetResponseRulesAction = "execute"
	PhaseVersionGetResponseRulesActionJsChallenge      PhaseVersionGetResponseRulesAction = "js_challenge"
	PhaseVersionGetResponseRulesActionLog              PhaseVersionGetResponseRulesAction = "log"
	PhaseVersionGetResponseRulesActionManagedChallenge PhaseVersionGetResponseRulesAction = "managed_challenge"
	PhaseVersionGetResponseRulesActionRedirect         PhaseVersionGetResponseRulesAction = "redirect"
	PhaseVersionGetResponseRulesActionRewrite          PhaseVersionGetResponseRulesAction = "rewrite"
	PhaseVersionGetResponseRulesActionRoute            PhaseVersionGetResponseRulesAction = "route"
	PhaseVersionGetResponseRulesActionScore            PhaseVersionGetResponseRulesAction = "score"
	PhaseVersionGetResponseRulesActionServeError       PhaseVersionGetResponseRulesAction = "serve_error"
	PhaseVersionGetResponseRulesActionSetConfig        PhaseVersionGetResponseRulesAction = "set_config"
	PhaseVersionGetResponseRulesActionSkip             PhaseVersionGetResponseRulesAction = "skip"
	PhaseVersionGetResponseRulesActionSetCacheSettings PhaseVersionGetResponseRulesAction = "set_cache_settings"
)

func (r PhaseVersionGetResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesActionBlock, PhaseVersionGetResponseRulesActionChallenge, PhaseVersionGetResponseRulesActionCompressResponse, PhaseVersionGetResponseRulesActionExecute, PhaseVersionGetResponseRulesActionJsChallenge, PhaseVersionGetResponseRulesActionLog, PhaseVersionGetResponseRulesActionManagedChallenge, PhaseVersionGetResponseRulesActionRedirect, PhaseVersionGetResponseRulesActionRewrite, PhaseVersionGetResponseRulesActionRoute, PhaseVersionGetResponseRulesActionScore, PhaseVersionGetResponseRulesActionServeError, PhaseVersionGetResponseRulesActionSetConfig, PhaseVersionGetResponseRulesActionSkip, PhaseVersionGetResponseRulesActionSetCacheSettings:
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
