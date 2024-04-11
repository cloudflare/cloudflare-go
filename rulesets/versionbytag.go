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

// VersionByTagService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewVersionByTagService] method
// instead.
type VersionByTagService struct {
	Options []option.RequestOption
}

// NewVersionByTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVersionByTagService(opts ...option.RequestOption) (r *VersionByTagService) {
	r = &VersionByTagService{}
	r.Options = opts
	return
}

// Fetches the rules of a managed account ruleset version for a given tag.
func (r *VersionByTagService) Get(ctx context.Context, rulesetID string, rulesetVersion string, ruleTag string, query VersionByTagGetParams, opts ...option.RequestOption) (res *VersionByTagGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VersionByTagGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s/by_tag/%s", query.AccountID, rulesetID, rulesetVersion, ruleTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type VersionByTagGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind VersionByTagGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase VersionByTagGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []VersionByTagGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                      `json:"description"`
	JSON        versionByTagGetResponseJSON `json:"-"`
}

// versionByTagGetResponseJSON contains the JSON metadata for the struct
// [VersionByTagGetResponse]
type versionByTagGetResponseJSON struct {
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

func (r *VersionByTagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type VersionByTagGetResponseKind string

const (
	VersionByTagGetResponseKindManaged VersionByTagGetResponseKind = "managed"
	VersionByTagGetResponseKindCustom  VersionByTagGetResponseKind = "custom"
	VersionByTagGetResponseKindRoot    VersionByTagGetResponseKind = "root"
	VersionByTagGetResponseKindZone    VersionByTagGetResponseKind = "zone"
)

func (r VersionByTagGetResponseKind) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseKindManaged, VersionByTagGetResponseKindCustom, VersionByTagGetResponseKindRoot, VersionByTagGetResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type VersionByTagGetResponsePhase string

const (
	VersionByTagGetResponsePhaseDDoSL4                         VersionByTagGetResponsePhase = "ddos_l4"
	VersionByTagGetResponsePhaseDDoSL7                         VersionByTagGetResponsePhase = "ddos_l7"
	VersionByTagGetResponsePhaseHTTPConfigSettings             VersionByTagGetResponsePhase = "http_config_settings"
	VersionByTagGetResponsePhaseHTTPCustomErrors               VersionByTagGetResponsePhase = "http_custom_errors"
	VersionByTagGetResponsePhaseHTTPLogCustomFields            VersionByTagGetResponsePhase = "http_log_custom_fields"
	VersionByTagGetResponsePhaseHTTPRatelimit                  VersionByTagGetResponsePhase = "http_ratelimit"
	VersionByTagGetResponsePhaseHTTPRequestCacheSettings       VersionByTagGetResponsePhase = "http_request_cache_settings"
	VersionByTagGetResponsePhaseHTTPRequestDynamicRedirect     VersionByTagGetResponsePhase = "http_request_dynamic_redirect"
	VersionByTagGetResponsePhaseHTTPRequestFirewallCustom      VersionByTagGetResponsePhase = "http_request_firewall_custom"
	VersionByTagGetResponsePhaseHTTPRequestFirewallManaged     VersionByTagGetResponsePhase = "http_request_firewall_managed"
	VersionByTagGetResponsePhaseHTTPRequestLateTransform       VersionByTagGetResponsePhase = "http_request_late_transform"
	VersionByTagGetResponsePhaseHTTPRequestOrigin              VersionByTagGetResponsePhase = "http_request_origin"
	VersionByTagGetResponsePhaseHTTPRequestRedirect            VersionByTagGetResponsePhase = "http_request_redirect"
	VersionByTagGetResponsePhaseHTTPRequestSanitize            VersionByTagGetResponsePhase = "http_request_sanitize"
	VersionByTagGetResponsePhaseHTTPRequestSbfm                VersionByTagGetResponsePhase = "http_request_sbfm"
	VersionByTagGetResponsePhaseHTTPRequestSelectConfiguration VersionByTagGetResponsePhase = "http_request_select_configuration"
	VersionByTagGetResponsePhaseHTTPRequestTransform           VersionByTagGetResponsePhase = "http_request_transform"
	VersionByTagGetResponsePhaseHTTPResponseCompression        VersionByTagGetResponsePhase = "http_response_compression"
	VersionByTagGetResponsePhaseHTTPResponseFirewallManaged    VersionByTagGetResponsePhase = "http_response_firewall_managed"
	VersionByTagGetResponsePhaseHTTPResponseHeadersTransform   VersionByTagGetResponsePhase = "http_response_headers_transform"
	VersionByTagGetResponsePhaseMagicTransit                   VersionByTagGetResponsePhase = "magic_transit"
	VersionByTagGetResponsePhaseMagicTransitIDsManaged         VersionByTagGetResponsePhase = "magic_transit_ids_managed"
	VersionByTagGetResponsePhaseMagicTransitManaged            VersionByTagGetResponsePhase = "magic_transit_managed"
)

func (r VersionByTagGetResponsePhase) IsKnown() bool {
	switch r {
	case VersionByTagGetResponsePhaseDDoSL4, VersionByTagGetResponsePhaseDDoSL7, VersionByTagGetResponsePhaseHTTPConfigSettings, VersionByTagGetResponsePhaseHTTPCustomErrors, VersionByTagGetResponsePhaseHTTPLogCustomFields, VersionByTagGetResponsePhaseHTTPRatelimit, VersionByTagGetResponsePhaseHTTPRequestCacheSettings, VersionByTagGetResponsePhaseHTTPRequestDynamicRedirect, VersionByTagGetResponsePhaseHTTPRequestFirewallCustom, VersionByTagGetResponsePhaseHTTPRequestFirewallManaged, VersionByTagGetResponsePhaseHTTPRequestLateTransform, VersionByTagGetResponsePhaseHTTPRequestOrigin, VersionByTagGetResponsePhaseHTTPRequestRedirect, VersionByTagGetResponsePhaseHTTPRequestSanitize, VersionByTagGetResponsePhaseHTTPRequestSbfm, VersionByTagGetResponsePhaseHTTPRequestSelectConfiguration, VersionByTagGetResponsePhaseHTTPRequestTransform, VersionByTagGetResponsePhaseHTTPResponseCompression, VersionByTagGetResponsePhaseHTTPResponseFirewallManaged, VersionByTagGetResponsePhaseHTTPResponseHeadersTransform, VersionByTagGetResponsePhaseMagicTransit, VersionByTagGetResponsePhaseMagicTransitIDsManaged, VersionByTagGetResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type VersionByTagGetResponseRule struct {
	// The action to perform when the rule matches.
	Action           VersionByTagGetResponseRulesAction `json:"action"`
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
	JSON    versionByTagGetResponseRuleJSON `json:"-"`
	union   VersionByTagGetResponseRulesUnion
}

// versionByTagGetResponseRuleJSON contains the JSON metadata for the struct
// [VersionByTagGetResponseRule]
type versionByTagGetResponseRuleJSON struct {
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

func (r versionByTagGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *VersionByTagGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r VersionByTagGetResponseRule) AsUnion() VersionByTagGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsChallengeRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsCompressResponseRule],
// [rulesets.ExecuteRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsJsChallengeRule],
// [rulesets.LogRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsManagedChallengeRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsRedirectRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsRewriteRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsRouteRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsScoreRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsServeErrorRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsSetConfigRule],
// [rulesets.SkipRule] or
// [rulesets.VersionByTagGetResponseRulesRulesetsSetCacheSettingsRule].
type VersionByTagGetResponseRulesUnion interface {
	implementsRulesetsVersionByTagGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionByTagGetResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsCompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsJsChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsRedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsRewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsRouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsSetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsSetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
	)
}

type VersionByTagGetResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsChallengeRuleAction `json:"action"`
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
	JSON versionByTagGetResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata
// for the struct [VersionByTagGetResponseRulesRulesetsChallengeRule]
type versionByTagGetResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsChallengeRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsChallengeRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsChallengeRuleActionChallenge VersionByTagGetResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r VersionByTagGetResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

type VersionByTagGetResponseRulesRulesetsCompressResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsCompressResponseRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParameters `json:"action_parameters"`
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
	JSON versionByTagGetResponseRulesRulesetsCompressResponseRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsCompressResponseRuleJSON contains the JSON
// metadata for the struct
// [VersionByTagGetResponseRulesRulesetsCompressResponseRule]
type versionByTagGetResponseRulesRulesetsCompressResponseRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsCompressResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsCompressResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsCompressResponseRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsCompressResponseRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionCompressResponse VersionByTagGetResponseRulesRulesetsCompressResponseRuleAction = "compress_response"
)

func (r VersionByTagGetResponseRulesRulesetsCompressResponseRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms []VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm `json:"algorithms"`
	JSON       versionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON        `json:"-"`
}

// versionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParameters]
type versionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON struct {
	Algorithms  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Compression algorithm to enable.
type VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName `json:"name"`
	JSON versionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON  `json:"-"`
}

// versionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm]
type versionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON) RawJSON() string {
	return r.raw
}

// Name of compression algorithm to enable.
type VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName string

const (
	VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone    VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "none"
	VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto    VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "auto"
	VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "default"
	VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip    VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "gzip"
	VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli  VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "brotli"
)

func (r VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone, VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto, VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault, VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip, VersionByTagGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli:
		return true
	}
	return false
}

type VersionByTagGetResponseRulesRulesetsJsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsJsChallengeRuleAction `json:"action"`
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
	JSON versionByTagGetResponseRulesRulesetsJsChallengeRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsJsChallengeRuleJSON contains the JSON
// metadata for the struct [VersionByTagGetResponseRulesRulesetsJsChallengeRule]
type versionByTagGetResponseRulesRulesetsJsChallengeRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsJsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsJsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsJsChallengeRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsJsChallengeRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsJsChallengeRuleActionJsChallenge VersionByTagGetResponseRulesRulesetsJsChallengeRuleAction = "js_challenge"
)

func (r VersionByTagGetResponseRulesRulesetsJsChallengeRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsJsChallengeRuleActionJsChallenge:
		return true
	}
	return false
}

type VersionByTagGetResponseRulesRulesetsManagedChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsManagedChallengeRuleAction `json:"action"`
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
	JSON versionByTagGetResponseRulesRulesetsManagedChallengeRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsManagedChallengeRuleJSON contains the JSON
// metadata for the struct
// [VersionByTagGetResponseRulesRulesetsManagedChallengeRule]
type versionByTagGetResponseRulesRulesetsManagedChallengeRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsManagedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsManagedChallengeRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsManagedChallengeRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge VersionByTagGetResponseRulesRulesetsManagedChallengeRuleAction = "managed_challenge"
)

func (r VersionByTagGetResponseRulesRulesetsManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

type VersionByTagGetResponseRulesRulesetsRedirectRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsRedirectRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionByTagGetResponseRulesRulesetsRedirectRuleActionParameters `json:"action_parameters"`
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
	JSON versionByTagGetResponseRulesRulesetsRedirectRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRedirectRuleJSON contains the JSON metadata
// for the struct [VersionByTagGetResponseRulesRulesetsRedirectRule]
type versionByTagGetResponseRulesRulesetsRedirectRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsRedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRedirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsRedirectRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsRedirectRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsRedirectRuleActionRedirect VersionByTagGetResponseRulesRulesetsRedirectRuleAction = "redirect"
)

func (r VersionByTagGetResponseRulesRulesetsRedirectRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsRedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionByTagGetResponseRulesRulesetsRedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromList `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValue `json:"from_value"`
	JSON      versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersJSON      `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersJSON contains
// the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRedirectRuleActionParameters]
type versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersJSON struct {
	FromList    apijson.Field
	FromValue   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRedirectRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on a bulk list lookup.
type VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key string `json:"key"`
	// The name of the list to match against.
	Name string                                                                       `json:"name"`
	JSON versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromList]
type versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON struct {
	Key         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on the request properties.
type VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString bool `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL `json:"target_url"`
	JSON      versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON      `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValue]
type versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON struct {
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	TargetURL           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON) RawJSON() string {
	return r.raw
}

// The status code to be used for the redirect.
type VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode float64

const (
	VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301 VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 301
	VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302 VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 302
	VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303 VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 303
	VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307 VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 307
	VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308 VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 308
)

func (r VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301, VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302, VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303, VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307, VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL struct {
	// The URL to redirect the request to.
	Value string `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                                 `json:"expression"`
	JSON       versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON `json:"-"`
	union      VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion
}

// versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL]
type versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON) RawJSON() string {
	return r.raw
}

func (r *VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) AsUnion() VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion {
	return r.union
}

// The URL to redirect the request to.
//
// Union satisfied by
// [rulesets.VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
// or
// [rulesets.VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect].
type VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRulesetsVersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect{}),
		},
	)
}

type VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value string                                                                                                  `json:"value"`
	JSON  versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
type versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRulesetsVersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                                                   `json:"expression"`
	JSON       versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect]
type versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRulesetsVersionByTagGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type VersionByTagGetResponseRulesRulesetsRewriteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsRewriteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionByTagGetResponseRulesRulesetsRewriteRuleActionParameters `json:"action_parameters"`
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
	JSON versionByTagGetResponseRulesRulesetsRewriteRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRewriteRuleJSON contains the JSON metadata
// for the struct [VersionByTagGetResponseRulesRulesetsRewriteRule]
type versionByTagGetResponseRulesRulesetsRewriteRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsRewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsRewriteRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsRewriteRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsRewriteRuleActionRewrite VersionByTagGetResponseRulesRulesetsRewriteRuleAction = "rewrite"
)

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsRewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers map[string]VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeader `json:"headers"`
	// URI to rewrite the request to.
	URI  VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURI  `json:"uri"`
	JSON versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRewriteRuleActionParameters]
type versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersJSON struct {
	Headers     apijson.Field
	URI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRewriteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Remove the header from the request.
type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeader struct {
	Operation VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation `json:"operation,required"`
	// Static value for the header.
	Value string `json:"value"`
	// Expression for the header value.
	Expression string                                                                    `json:"expression"`
	JSON       versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON `json:"-"`
	union      VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion
}

// versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeader]
type versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON) RawJSON() string {
	return r.raw
}

func (r *VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeader) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeader) AsUnion() VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion {
	return r.union
}

// Remove the header from the request.
//
// Union satisfied by
// [rulesets.VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader],
// [rulesets.VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
// or
// [rulesets.VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader].
type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion interface {
	implementsRulesetsVersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeader()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader{}),
		},
	)
}

// Remove the header from the request.
type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation `json:"operation,required"`
	JSON      versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON      `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader]
type versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON struct {
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) implementsRulesetsVersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Set a request header with a static value.
type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader struct {
	Operation VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation `json:"operation,required"`
	// Static value for the header.
	Value string                                                                                 `json:"value,required"`
	JSON  versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
type versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) implementsRulesetsVersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation string

const (
	VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation = "set"
)

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader struct {
	// Expression for the header value.
	Expression string                                                                                       `json:"expression,required"`
	Operation  VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation `json:"operation,required"`
	JSON       versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON      `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader]
type versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) implementsRulesetsVersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation string

const (
	VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation = "set"
)

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet:
		return true
	}
	return false
}

type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation string

const (
	VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "remove"
	VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet    VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "set"
)

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove, VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPath `json:"path"`
	// Query portion rewrite.
	Query VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery `json:"query"`
	JSON  versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON  `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON contains
// the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURI]
type versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON struct {
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON) RawJSON() string {
	return r.raw
}

// Path portion rewrite.
type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPath struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                     `json:"expression"`
	JSON       versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON `json:"-"`
	union      VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion
}

// versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPath]
type versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON) RawJSON() string {
	return r.raw
}

func (r *VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPath) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPath) AsUnion() VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion {
	return r.union
}

// Path portion rewrite.
//
// Union satisfied by
// [rulesets.VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
// or
// [rulesets.VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue].
type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion interface {
	implementsRulesetsVersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPath()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue{}),
		},
	)
}

type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue struct {
	// Predefined replacement value.
	Value string                                                                                `json:"value,required"`
	JSON  versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
type versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) implementsRulesetsVersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                                 `json:"expression,required"`
	JSON       versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue]
type versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) implementsRulesetsVersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

// Query portion rewrite.
type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                      `json:"expression"`
	JSON       versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON `json:"-"`
	union      VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion
}

// versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery]
type versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON) RawJSON() string {
	return r.raw
}

func (r *VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery) AsUnion() VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion {
	return r.union
}

// Query portion rewrite.
//
// Union satisfied by
// [rulesets.VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
// or
// [rulesets.VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue].
type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion interface {
	implementsRulesetsVersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue{}),
		},
	)
}

type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue struct {
	// Predefined replacement value.
	Value string                                                                                 `json:"value,required"`
	JSON  versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
type versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) implementsRulesetsVersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                                  `json:"expression,required"`
	JSON       versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue]
type versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) implementsRulesetsVersionByTagGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type VersionByTagGetResponseRulesRulesetsRouteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsRouteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionByTagGetResponseRulesRulesetsRouteRuleActionParameters `json:"action_parameters"`
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
	JSON versionByTagGetResponseRulesRulesetsRouteRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRouteRuleJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseRulesRulesetsRouteRule]
type versionByTagGetResponseRulesRulesetsRouteRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsRouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRouteRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsRouteRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsRouteRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsRouteRuleActionRoute VersionByTagGetResponseRulesRulesetsRouteRuleAction = "route"
)

func (r VersionByTagGetResponseRulesRulesetsRouteRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsRouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionByTagGetResponseRulesRulesetsRouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader string `json:"host_header"`
	// Override the IP/TCP destination.
	Origin VersionByTagGetResponseRulesRulesetsRouteRuleActionParametersOrigin `json:"origin"`
	// Override the Server Name Indication (SNI).
	Sni  VersionByTagGetResponseRulesRulesetsRouteRuleActionParametersSni  `json:"sni"`
	JSON versionByTagGetResponseRulesRulesetsRouteRuleActionParametersJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRouteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRouteRuleActionParameters]
type versionByTagGetResponseRulesRulesetsRouteRuleActionParametersJSON struct {
	HostHeader  apijson.Field
	Origin      apijson.Field
	Sni         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRouteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRouteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Override the IP/TCP destination.
type VersionByTagGetResponseRulesRulesetsRouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host string `json:"host"`
	// Override the destination port.
	Port float64                                                                 `json:"port"`
	JSON versionByTagGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON contains
// the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRouteRuleActionParametersOrigin]
type versionByTagGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON struct {
	Host        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRouteRuleActionParametersOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON) RawJSON() string {
	return r.raw
}

// Override the Server Name Indication (SNI).
type VersionByTagGetResponseRulesRulesetsRouteRuleActionParametersSni struct {
	// The SNI override.
	Value string                                                               `json:"value,required"`
	JSON  versionByTagGetResponseRulesRulesetsRouteRuleActionParametersSniJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsRouteRuleActionParametersSniJSON contains
// the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsRouteRuleActionParametersSni]
type versionByTagGetResponseRulesRulesetsRouteRuleActionParametersSniJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsRouteRuleActionParametersSni) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsRouteRuleActionParametersSniJSON) RawJSON() string {
	return r.raw
}

type VersionByTagGetResponseRulesRulesetsScoreRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsScoreRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionByTagGetResponseRulesRulesetsScoreRuleActionParameters `json:"action_parameters"`
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
	JSON versionByTagGetResponseRulesRulesetsScoreRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsScoreRuleJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseRulesRulesetsScoreRule]
type versionByTagGetResponseRulesRulesetsScoreRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsScoreRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsScoreRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsScoreRuleActionScore VersionByTagGetResponseRulesRulesetsScoreRuleAction = "score"
)

func (r VersionByTagGetResponseRulesRulesetsScoreRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionByTagGetResponseRulesRulesetsScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment int64                                                             `json:"increment"`
	JSON      versionByTagGetResponseRulesRulesetsScoreRuleActionParametersJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsScoreRuleActionParametersJSON contains the
// JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsScoreRuleActionParameters]
type versionByTagGetResponseRulesRulesetsScoreRuleActionParametersJSON struct {
	Increment   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsScoreRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsScoreRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

type VersionByTagGetResponseRulesRulesetsServeErrorRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsServeErrorRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParameters `json:"action_parameters"`
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
	JSON versionByTagGetResponseRulesRulesetsServeErrorRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsServeErrorRuleJSON contains the JSON
// metadata for the struct [VersionByTagGetResponseRulesRulesetsServeErrorRule]
type versionByTagGetResponseRulesRulesetsServeErrorRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsServeErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsServeErrorRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsServeErrorRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsServeErrorRuleActionServeError VersionByTagGetResponseRulesRulesetsServeErrorRuleAction = "serve_error"
)

func (r VersionByTagGetResponseRulesRulesetsServeErrorRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParameters struct {
	// Error response content.
	Content string `json:"content"`
	// Content-type header to set with the response.
	ContentType VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentType `json:"content_type"`
	// The status code to use for the error.
	StatusCode float64                                                                `json:"status_code"`
	JSON       versionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersJSON contains
// the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParameters]
type versionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Content-type header to set with the response.
type VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentType string

const (
	VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "application/json"
	VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml         VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/xml"
	VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain       VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/plain"
	VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML        VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/html"
)

func (r VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson, VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml, VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain, VersionByTagGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

type VersionByTagGetResponseRulesRulesetsSetConfigRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsSetConfigRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParameters `json:"action_parameters"`
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
	JSON versionByTagGetResponseRulesRulesetsSetConfigRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetConfigRuleJSON contains the JSON metadata
// for the struct [VersionByTagGetResponseRulesRulesetsSetConfigRule]
type versionByTagGetResponseRulesRulesetsSetConfigRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsSetConfigRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetConfigRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsSetConfigRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsSetConfigRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionSetConfig VersionByTagGetResponseRulesRulesetsSetConfigRuleAction = "set_config"
)

func (r VersionByTagGetResponseRulesRulesetsSetConfigRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsSetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites bool `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify `json:"autominify"`
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
	Polish VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersPolish `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader bool `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes bool `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSL `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg  bool                                                                  `json:"sxg"`
	JSON versionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersJSON contains
// the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParameters]
type versionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Select which file extensions to minify automatically.
type VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	Css bool `json:"css"`
	// Minify HTML files.
	HTML bool `json:"html"`
	// Minify JS files.
	Js   bool                                                                            `json:"js"`
	JSON versionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify]
type versionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON) RawJSON() string {
	return r.raw
}

// Configure the Polish level.
type VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersPolish string

const (
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersPolishOff      VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "off"
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossless"
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy    VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossy"
)

func (r VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersPolishOff, VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless, VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy:
		return true
	}
	return false
}

// Configure the Security Level.
type VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel string

const (
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff            VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "off"
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "essentially_off"
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow            VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "low"
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium         VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "medium"
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh           VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "high"
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack    VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff, VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff, VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow, VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium, VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh, VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSL string

const (
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOff        VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "off"
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible   VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "flexible"
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFull       VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "full"
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict     VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "strict"
	VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "origin_pull"
)

func (r VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOff, VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible, VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFull, VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict, VersionByTagGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters `json:"action_parameters"`
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
	JSON versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleJSON contains the JSON
// metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRule]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsSetCacheSettingsRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleAction = "set_cache_settings"
)

func (r VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts []int64 `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache bool `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL `json:"edge_ttl"`
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
	ServeStale VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale `json:"serve_stale"`
	JSON       versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON       `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64                                                                                  `json:"default"`
	JSON    versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON struct {
	Mode        apijson.Field
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON) RawJSON() string {
	return r.raw
}

// Determines which browser ttl mode to use.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
)

func (r VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType bool `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor bool `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder bool                                                                                 `json:"ignore_query_strings_order"`
	JSON                    versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON struct {
	CacheByDeviceType       apijson.Field
	CacheDeceptionArmor     apijson.Field
	CustomKey               apijson.Field
	IgnoreQueryStringsOrder apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON) RawJSON() string {
	return r.raw
}

// Customize which components of the request are included or excluded from the
// cache key.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser `json:"user"`
	JSON versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON) RawJSON() string {
	return r.raw
}

// The cookies to include in building the cache key.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Include these cookies' names and their values.
	Include []string                                                                                            `json:"include"`
	JSON    versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON) RawJSON() string {
	return r.raw
}

// The header names and values to include in building the cache key.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin bool `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include []string                                                                                            `json:"include"`
	JSON    versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON struct {
	CheckPresence apijson.Field
	ExcludeOrigin apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON) RawJSON() string {
	return r.raw
}

// Whether to use the original host or the resolved host in the cache key.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved bool                                                                                              `json:"resolved"`
	JSON     versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON) RawJSON() string {
	return r.raw
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude `json:"include"`
	JSON    versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON    `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON) RawJSON() string {
	return r.raw
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Exclude all query string parameters from use in building the cache key.
	All bool `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List []string                                                                                                        `json:"list"`
	JSON versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON) RawJSON() string {
	return r.raw
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Use all query string parameters in the cache key.
	All bool `json:"all"`
	// A list of query string parameters used to build the cache key.
	List []string                                                                                                        `json:"list"`
	JSON versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON) RawJSON() string {
	return r.raw
}

// Characteristics of the request user agent used in building the cache key.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType bool `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo bool `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang bool                                                                                              `json:"lang"`
	JSON versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON) RawJSON() string {
	return r.raw
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible bool `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize int64                                                                                    `json:"min_file_size,required"`
	JSON        versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON struct {
	Eligible    apijson.Field
	MinFileSize apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON) RawJSON() string {
	return r.raw
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64 `json:"default,required"`
	// edge ttl options
	Mode VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL []VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL `json:"status_code_ttl,required"`
	JSON          versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON            `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON struct {
	Default       apijson.Field
	Mode          apijson.Field
	StatusCodeTTL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON) RawJSON() string {
	return r.raw
}

// edge ttl options
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value int64 `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue int64                                                                                            `json:"status_code_value"`
	JSON            versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON struct {
	Value           apijson.Field
	StatusCodeRange apijson.Field
	StatusCodeValue apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON) RawJSON() string {
	return r.raw
}

// The range of status codes used to apply the selected mode.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// response status code lower bound
	From int64 `json:"from,required"`
	// response status code upper bound
	To   int64                                                                                                           `json:"to,required"`
	JSON versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON) RawJSON() string {
	return r.raw
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating bool                                                                                   `json:"disable_stale_while_updating,required"`
	JSON                      versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale]
type versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON struct {
	DisableStaleWhileUpdating apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesAction string

const (
	VersionByTagGetResponseRulesActionBlock            VersionByTagGetResponseRulesAction = "block"
	VersionByTagGetResponseRulesActionChallenge        VersionByTagGetResponseRulesAction = "challenge"
	VersionByTagGetResponseRulesActionCompressResponse VersionByTagGetResponseRulesAction = "compress_response"
	VersionByTagGetResponseRulesActionExecute          VersionByTagGetResponseRulesAction = "execute"
	VersionByTagGetResponseRulesActionJsChallenge      VersionByTagGetResponseRulesAction = "js_challenge"
	VersionByTagGetResponseRulesActionLog              VersionByTagGetResponseRulesAction = "log"
	VersionByTagGetResponseRulesActionManagedChallenge VersionByTagGetResponseRulesAction = "managed_challenge"
	VersionByTagGetResponseRulesActionRedirect         VersionByTagGetResponseRulesAction = "redirect"
	VersionByTagGetResponseRulesActionRewrite          VersionByTagGetResponseRulesAction = "rewrite"
	VersionByTagGetResponseRulesActionRoute            VersionByTagGetResponseRulesAction = "route"
	VersionByTagGetResponseRulesActionScore            VersionByTagGetResponseRulesAction = "score"
	VersionByTagGetResponseRulesActionServeError       VersionByTagGetResponseRulesAction = "serve_error"
	VersionByTagGetResponseRulesActionSetConfig        VersionByTagGetResponseRulesAction = "set_config"
	VersionByTagGetResponseRulesActionSkip             VersionByTagGetResponseRulesAction = "skip"
	VersionByTagGetResponseRulesActionSetCacheSettings VersionByTagGetResponseRulesAction = "set_cache_settings"
)

func (r VersionByTagGetResponseRulesAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesActionBlock, VersionByTagGetResponseRulesActionChallenge, VersionByTagGetResponseRulesActionCompressResponse, VersionByTagGetResponseRulesActionExecute, VersionByTagGetResponseRulesActionJsChallenge, VersionByTagGetResponseRulesActionLog, VersionByTagGetResponseRulesActionManagedChallenge, VersionByTagGetResponseRulesActionRedirect, VersionByTagGetResponseRulesActionRewrite, VersionByTagGetResponseRulesActionRoute, VersionByTagGetResponseRulesActionScore, VersionByTagGetResponseRulesActionServeError, VersionByTagGetResponseRulesActionSetConfig, VersionByTagGetResponseRulesActionSkip, VersionByTagGetResponseRulesActionSetCacheSettings:
		return true
	}
	return false
}

type VersionByTagGetParams struct {
	// The unique ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

// A response object.
type VersionByTagGetResponseEnvelope struct {
	// A list of error messages.
	Errors []VersionByTagGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []VersionByTagGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result VersionByTagGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success VersionByTagGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    versionByTagGetResponseEnvelopeJSON    `json:"-"`
}

// versionByTagGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VersionByTagGetResponseEnvelope]
type versionByTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionByTagGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionByTagGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   versionByTagGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// versionByTagGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VersionByTagGetResponseEnvelopeErrors]
type versionByTagGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionByTagGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    versionByTagGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// versionByTagGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseEnvelopeErrorsSource]
type versionByTagGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionByTagGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionByTagGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   versionByTagGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// versionByTagGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VersionByTagGetResponseEnvelopeMessages]
type versionByTagGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionByTagGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    versionByTagGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// versionByTagGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseEnvelopeMessagesSource]
type versionByTagGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type VersionByTagGetResponseEnvelopeSuccess bool

const (
	VersionByTagGetResponseEnvelopeSuccessTrue VersionByTagGetResponseEnvelopeSuccess = true
)

func (r VersionByTagGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
