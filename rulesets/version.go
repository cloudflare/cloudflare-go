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

// VersionService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewVersionService] method instead.
type VersionService struct {
	Options []option.RequestOption
	ByTag   *VersionByTagService
}

// NewVersionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVersionService(opts ...option.RequestOption) (r *VersionService) {
	r = &VersionService{}
	r.Options = opts
	r.ByTag = NewVersionByTagService(opts...)
	return
}

// Fetches the versions of an account or zone ruleset.
func (r *VersionService) List(ctx context.Context, rulesetID string, query VersionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Ruleset], err error) {
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
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions", accountOrZone, accountOrZoneID, rulesetID)
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

// Fetches the versions of an account or zone ruleset.
func (r *VersionService) ListAutoPaging(ctx context.Context, rulesetID string, query VersionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Ruleset] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, rulesetID, query, opts...))
}

// Deletes an existing version of an account or zone ruleset.
func (r *VersionService) Delete(ctx context.Context, rulesetID string, rulesetVersion string, body VersionDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions/%s", accountOrZone, accountOrZoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches a specific version of an account or zone ruleset.
func (r *VersionService) Get(ctx context.Context, rulesetID string, rulesetVersion string, query VersionGetParams, opts ...option.RequestOption) (res *VersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VersionGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions/%s", accountOrZone, accountOrZoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type VersionGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind VersionGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase VersionGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []VersionGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                 `json:"description"`
	JSON        versionGetResponseJSON `json:"-"`
}

// versionGetResponseJSON contains the JSON metadata for the struct
// [VersionGetResponse]
type versionGetResponseJSON struct {
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

func (r *VersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type VersionGetResponseKind string

const (
	VersionGetResponseKindManaged VersionGetResponseKind = "managed"
	VersionGetResponseKindCustom  VersionGetResponseKind = "custom"
	VersionGetResponseKindRoot    VersionGetResponseKind = "root"
	VersionGetResponseKindZone    VersionGetResponseKind = "zone"
)

func (r VersionGetResponseKind) IsKnown() bool {
	switch r {
	case VersionGetResponseKindManaged, VersionGetResponseKindCustom, VersionGetResponseKindRoot, VersionGetResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type VersionGetResponsePhase string

const (
	VersionGetResponsePhaseDDoSL4                         VersionGetResponsePhase = "ddos_l4"
	VersionGetResponsePhaseDDoSL7                         VersionGetResponsePhase = "ddos_l7"
	VersionGetResponsePhaseHTTPConfigSettings             VersionGetResponsePhase = "http_config_settings"
	VersionGetResponsePhaseHTTPCustomErrors               VersionGetResponsePhase = "http_custom_errors"
	VersionGetResponsePhaseHTTPLogCustomFields            VersionGetResponsePhase = "http_log_custom_fields"
	VersionGetResponsePhaseHTTPRatelimit                  VersionGetResponsePhase = "http_ratelimit"
	VersionGetResponsePhaseHTTPRequestCacheSettings       VersionGetResponsePhase = "http_request_cache_settings"
	VersionGetResponsePhaseHTTPRequestDynamicRedirect     VersionGetResponsePhase = "http_request_dynamic_redirect"
	VersionGetResponsePhaseHTTPRequestFirewallCustom      VersionGetResponsePhase = "http_request_firewall_custom"
	VersionGetResponsePhaseHTTPRequestFirewallManaged     VersionGetResponsePhase = "http_request_firewall_managed"
	VersionGetResponsePhaseHTTPRequestLateTransform       VersionGetResponsePhase = "http_request_late_transform"
	VersionGetResponsePhaseHTTPRequestOrigin              VersionGetResponsePhase = "http_request_origin"
	VersionGetResponsePhaseHTTPRequestRedirect            VersionGetResponsePhase = "http_request_redirect"
	VersionGetResponsePhaseHTTPRequestSanitize            VersionGetResponsePhase = "http_request_sanitize"
	VersionGetResponsePhaseHTTPRequestSbfm                VersionGetResponsePhase = "http_request_sbfm"
	VersionGetResponsePhaseHTTPRequestSelectConfiguration VersionGetResponsePhase = "http_request_select_configuration"
	VersionGetResponsePhaseHTTPRequestTransform           VersionGetResponsePhase = "http_request_transform"
	VersionGetResponsePhaseHTTPResponseCompression        VersionGetResponsePhase = "http_response_compression"
	VersionGetResponsePhaseHTTPResponseFirewallManaged    VersionGetResponsePhase = "http_response_firewall_managed"
	VersionGetResponsePhaseHTTPResponseHeadersTransform   VersionGetResponsePhase = "http_response_headers_transform"
	VersionGetResponsePhaseMagicTransit                   VersionGetResponsePhase = "magic_transit"
	VersionGetResponsePhaseMagicTransitIDsManaged         VersionGetResponsePhase = "magic_transit_ids_managed"
	VersionGetResponsePhaseMagicTransitManaged            VersionGetResponsePhase = "magic_transit_managed"
)

func (r VersionGetResponsePhase) IsKnown() bool {
	switch r {
	case VersionGetResponsePhaseDDoSL4, VersionGetResponsePhaseDDoSL7, VersionGetResponsePhaseHTTPConfigSettings, VersionGetResponsePhaseHTTPCustomErrors, VersionGetResponsePhaseHTTPLogCustomFields, VersionGetResponsePhaseHTTPRatelimit, VersionGetResponsePhaseHTTPRequestCacheSettings, VersionGetResponsePhaseHTTPRequestDynamicRedirect, VersionGetResponsePhaseHTTPRequestFirewallCustom, VersionGetResponsePhaseHTTPRequestFirewallManaged, VersionGetResponsePhaseHTTPRequestLateTransform, VersionGetResponsePhaseHTTPRequestOrigin, VersionGetResponsePhaseHTTPRequestRedirect, VersionGetResponsePhaseHTTPRequestSanitize, VersionGetResponsePhaseHTTPRequestSbfm, VersionGetResponsePhaseHTTPRequestSelectConfiguration, VersionGetResponsePhaseHTTPRequestTransform, VersionGetResponsePhaseHTTPResponseCompression, VersionGetResponsePhaseHTTPResponseFirewallManaged, VersionGetResponsePhaseHTTPResponseHeadersTransform, VersionGetResponsePhaseMagicTransit, VersionGetResponsePhaseMagicTransitIDsManaged, VersionGetResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type VersionGetResponseRule struct {
	// The action to perform when the rule matches.
	Action           VersionGetResponseRulesAction `json:"action"`
	ActionParameters interface{}                   `json:"action_parameters,required"`
	Categories       interface{}                   `json:"categories,required"`
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
	Version string                     `json:"version,required"`
	JSON    versionGetResponseRuleJSON `json:"-"`
	union   VersionGetResponseRulesUnion
}

// versionGetResponseRuleJSON contains the JSON metadata for the struct
// [VersionGetResponseRule]
type versionGetResponseRuleJSON struct {
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

func (r versionGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r VersionGetResponseRule) AsUnion() VersionGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule],
// [rulesets.VersionGetResponseRulesRulesetsChallengeRule],
// [rulesets.VersionGetResponseRulesRulesetsCompressResponseRule],
// [rulesets.ExecuteRule],
// [rulesets.VersionGetResponseRulesRulesetsJsChallengeRule], [rulesets.LogRule],
// [rulesets.VersionGetResponseRulesRulesetsManagedChallengeRule],
// [rulesets.VersionGetResponseRulesRulesetsRedirectRule],
// [rulesets.VersionGetResponseRulesRulesetsRewriteRule],
// [rulesets.VersionGetResponseRulesRulesetsRouteRule],
// [rulesets.VersionGetResponseRulesRulesetsScoreRule],
// [rulesets.VersionGetResponseRulesRulesetsServeErrorRule],
// [rulesets.VersionGetResponseRulesRulesetsSetConfigRule], [rulesets.SkipRule] or
// [rulesets.VersionGetResponseRulesRulesetsSetCacheSettingsRule].
type VersionGetResponseRulesUnion interface {
	implementsRulesetsVersionGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsCompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsJsChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsRedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsRewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsRouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsSetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsSetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
	)
}

type VersionGetResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsChallengeRuleAction `json:"action"`
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
	JSON versionGetResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseRulesRulesetsChallengeRule]
type versionGetResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsChallengeRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsChallengeRuleAction string

const (
	VersionGetResponseRulesRulesetsChallengeRuleActionChallenge VersionGetResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r VersionGetResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

type VersionGetResponseRulesRulesetsCompressResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsCompressResponseRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesRulesetsCompressResponseRuleActionParameters `json:"action_parameters"`
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
	JSON versionGetResponseRulesRulesetsCompressResponseRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsCompressResponseRuleJSON contains the JSON
// metadata for the struct [VersionGetResponseRulesRulesetsCompressResponseRule]
type versionGetResponseRulesRulesetsCompressResponseRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsCompressResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsCompressResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsCompressResponseRule) implementsRulesetsVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsCompressResponseRuleAction string

const (
	VersionGetResponseRulesRulesetsCompressResponseRuleActionCompressResponse VersionGetResponseRulesRulesetsCompressResponseRuleAction = "compress_response"
)

func (r VersionGetResponseRulesRulesetsCompressResponseRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsCompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesRulesetsCompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms []VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm `json:"algorithms"`
	JSON       versionGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON        `json:"-"`
}

// versionGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsCompressResponseRuleActionParameters]
type versionGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON struct {
	Algorithms  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsCompressResponseRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Compression algorithm to enable.
type VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName `json:"name"`
	JSON versionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON  `json:"-"`
}

// versionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm]
type versionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON) RawJSON() string {
	return r.raw
}

// Name of compression algorithm to enable.
type VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName string

const (
	VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone    VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "none"
	VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto    VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "auto"
	VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "default"
	VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip    VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "gzip"
	VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli  VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "brotli"
)

func (r VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone, VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto, VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault, VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip, VersionGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli:
		return true
	}
	return false
}

type VersionGetResponseRulesRulesetsJsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsJsChallengeRuleAction `json:"action"`
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
	Ref  string                                             `json:"ref"`
	JSON versionGetResponseRulesRulesetsJsChallengeRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsJsChallengeRuleJSON contains the JSON metadata
// for the struct [VersionGetResponseRulesRulesetsJsChallengeRule]
type versionGetResponseRulesRulesetsJsChallengeRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsJsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsJsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsJsChallengeRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsJsChallengeRuleAction string

const (
	VersionGetResponseRulesRulesetsJsChallengeRuleActionJsChallenge VersionGetResponseRulesRulesetsJsChallengeRuleAction = "js_challenge"
)

func (r VersionGetResponseRulesRulesetsJsChallengeRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsJsChallengeRuleActionJsChallenge:
		return true
	}
	return false
}

type VersionGetResponseRulesRulesetsManagedChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsManagedChallengeRuleAction `json:"action"`
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
	JSON versionGetResponseRulesRulesetsManagedChallengeRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsManagedChallengeRuleJSON contains the JSON
// metadata for the struct [VersionGetResponseRulesRulesetsManagedChallengeRule]
type versionGetResponseRulesRulesetsManagedChallengeRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsManagedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsManagedChallengeRule) implementsRulesetsVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsManagedChallengeRuleAction string

const (
	VersionGetResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge VersionGetResponseRulesRulesetsManagedChallengeRuleAction = "managed_challenge"
)

func (r VersionGetResponseRulesRulesetsManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

type VersionGetResponseRulesRulesetsRedirectRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsRedirectRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesRulesetsRedirectRuleActionParameters `json:"action_parameters"`
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
	JSON versionGetResponseRulesRulesetsRedirectRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRedirectRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseRulesRulesetsRedirectRule]
type versionGetResponseRulesRulesetsRedirectRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsRedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRedirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsRedirectRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsRedirectRuleAction string

const (
	VersionGetResponseRulesRulesetsRedirectRuleActionRedirect VersionGetResponseRulesRulesetsRedirectRuleAction = "redirect"
)

func (r VersionGetResponseRulesRulesetsRedirectRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsRedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesRulesetsRedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromList `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValue `json:"from_value"`
	JSON      versionGetResponseRulesRulesetsRedirectRuleActionParametersJSON      `json:"-"`
}

// versionGetResponseRulesRulesetsRedirectRuleActionParametersJSON contains the
// JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRedirectRuleActionParameters]
type versionGetResponseRulesRulesetsRedirectRuleActionParametersJSON struct {
	FromList    apijson.Field
	FromValue   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRedirectRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRedirectRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on a bulk list lookup.
type VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key string `json:"key"`
	// The name of the list to match against.
	Name string                                                                  `json:"name"`
	JSON versionGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromList]
type versionGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON struct {
	Key         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on the request properties.
type VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString bool `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL `json:"target_url"`
	JSON      versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON      `json:"-"`
}

// versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValue]
type versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON struct {
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	TargetURL           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON) RawJSON() string {
	return r.raw
}

// The status code to be used for the redirect.
type VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode float64

const (
	VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301 VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 301
	VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302 VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 302
	VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303 VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 303
	VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307 VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 307
	VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308 VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 308
)

func (r VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301, VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302, VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303, VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307, VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL struct {
	// The URL to redirect the request to.
	Value string `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                            `json:"expression"`
	JSON       versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON `json:"-"`
	union      VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion
}

// versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL]
type versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) AsUnion() VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion {
	return r.union
}

// The URL to redirect the request to.
//
// Union satisfied by
// [rulesets.VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
// or
// [rulesets.VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect].
type VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRulesetsVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect{}),
		},
	)
}

type VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value string                                                                                             `json:"value"`
	JSON  versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
type versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRulesetsVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                                              `json:"expression"`
	JSON       versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect]
type versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRulesetsVersionGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type VersionGetResponseRulesRulesetsRewriteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsRewriteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesRulesetsRewriteRuleActionParameters `json:"action_parameters"`
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
	JSON versionGetResponseRulesRulesetsRewriteRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRewriteRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseRulesRulesetsRewriteRule]
type versionGetResponseRulesRulesetsRewriteRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsRewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsRewriteRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsRewriteRuleAction string

const (
	VersionGetResponseRulesRulesetsRewriteRuleActionRewrite VersionGetResponseRulesRulesetsRewriteRuleAction = "rewrite"
)

func (r VersionGetResponseRulesRulesetsRewriteRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsRewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesRulesetsRewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers map[string]VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader `json:"headers"`
	// URI to rewrite the request to.
	URI  VersionGetResponseRulesRulesetsRewriteRuleActionParametersURI  `json:"uri"`
	JSON versionGetResponseRulesRulesetsRewriteRuleActionParametersJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRewriteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [VersionGetResponseRulesRulesetsRewriteRuleActionParameters]
type versionGetResponseRulesRulesetsRewriteRuleActionParametersJSON struct {
	Headers     apijson.Field
	URI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRewriteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRewriteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Remove the header from the request.
type VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader struct {
	Operation VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation `json:"operation,required"`
	// Static value for the header.
	Value string `json:"value"`
	// Expression for the header value.
	Expression string                                                               `json:"expression"`
	JSON       versionGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON `json:"-"`
	union      VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion
}

// versionGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader]
type versionGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r versionGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader) AsUnion() VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion {
	return r.union
}

// Remove the header from the request.
//
// Union satisfied by
// [rulesets.VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader],
// [rulesets.VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
// or
// [rulesets.VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader].
type VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion interface {
	implementsRulesetsVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader{}),
		},
	)
}

// Remove the header from the request.
type VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation `json:"operation,required"`
	JSON      versionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON      `json:"-"`
}

// versionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader]
type versionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON struct {
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) implementsRulesetsVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Set a request header with a static value.
type VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader struct {
	Operation VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation `json:"operation,required"`
	// Static value for the header.
	Value string                                                                            `json:"value,required"`
	JSON  versionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
type versionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) implementsRulesetsVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation string

const (
	VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation = "set"
)

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader struct {
	// Expression for the header value.
	Expression string                                                                                  `json:"expression,required"`
	Operation  VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation `json:"operation,required"`
	JSON       versionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON      `json:"-"`
}

// versionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader]
type versionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) implementsRulesetsVersionGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation string

const (
	VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation = "set"
)

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet:
		return true
	}
	return false
}

type VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation string

const (
	VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "remove"
	VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet    VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "set"
)

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove, VersionGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type VersionGetResponseRulesRulesetsRewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath `json:"path"`
	// Query portion rewrite.
	Query VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery `json:"query"`
	JSON  versionGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON  `json:"-"`
}

// versionGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON contains the
// JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRewriteRuleActionParametersURI]
type versionGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON struct {
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRewriteRuleActionParametersURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON) RawJSON() string {
	return r.raw
}

// Path portion rewrite.
type VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                `json:"expression"`
	JSON       versionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON `json:"-"`
	union      VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion
}

// versionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath]
type versionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r versionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath) AsUnion() VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion {
	return r.union
}

// Path portion rewrite.
//
// Union satisfied by
// [rulesets.VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
// or
// [rulesets.VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue].
type VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion interface {
	implementsRulesetsVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue{}),
		},
	)
}

type VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue struct {
	// Predefined replacement value.
	Value string                                                                           `json:"value,required"`
	JSON  versionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
type versionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) implementsRulesetsVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

type VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                            `json:"expression,required"`
	JSON       versionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue]
type versionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) implementsRulesetsVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

// Query portion rewrite.
type VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                 `json:"expression"`
	JSON       versionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON `json:"-"`
	union      VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion
}

// versionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery]
type versionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r versionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery) AsUnion() VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion {
	return r.union
}

// Query portion rewrite.
//
// Union satisfied by
// [rulesets.VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
// or
// [rulesets.VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue].
type VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion interface {
	implementsRulesetsVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue{}),
		},
	)
}

type VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue struct {
	// Predefined replacement value.
	Value string                                                                            `json:"value,required"`
	JSON  versionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
type versionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) implementsRulesetsVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                             `json:"expression,required"`
	JSON       versionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue]
type versionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) implementsRulesetsVersionGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type VersionGetResponseRulesRulesetsRouteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsRouteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesRulesetsRouteRuleActionParameters `json:"action_parameters"`
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
	JSON versionGetResponseRulesRulesetsRouteRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRouteRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseRulesRulesetsRouteRule]
type versionGetResponseRulesRulesetsRouteRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsRouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRouteRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsRouteRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsRouteRuleAction string

const (
	VersionGetResponseRulesRulesetsRouteRuleActionRoute VersionGetResponseRulesRulesetsRouteRuleAction = "route"
)

func (r VersionGetResponseRulesRulesetsRouteRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsRouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesRulesetsRouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader string `json:"host_header"`
	// Override the IP/TCP destination.
	Origin VersionGetResponseRulesRulesetsRouteRuleActionParametersOrigin `json:"origin"`
	// Override the Server Name Indication (SNI).
	Sni  VersionGetResponseRulesRulesetsRouteRuleActionParametersSni  `json:"sni"`
	JSON versionGetResponseRulesRulesetsRouteRuleActionParametersJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRouteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [VersionGetResponseRulesRulesetsRouteRuleActionParameters]
type versionGetResponseRulesRulesetsRouteRuleActionParametersJSON struct {
	HostHeader  apijson.Field
	Origin      apijson.Field
	Sni         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRouteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRouteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Override the IP/TCP destination.
type VersionGetResponseRulesRulesetsRouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host string `json:"host"`
	// Override the destination port.
	Port float64                                                            `json:"port"`
	JSON versionGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON contains the
// JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRouteRuleActionParametersOrigin]
type versionGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON struct {
	Host        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRouteRuleActionParametersOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON) RawJSON() string {
	return r.raw
}

// Override the Server Name Indication (SNI).
type VersionGetResponseRulesRulesetsRouteRuleActionParametersSni struct {
	// The SNI override.
	Value string                                                          `json:"value,required"`
	JSON  versionGetResponseRulesRulesetsRouteRuleActionParametersSniJSON `json:"-"`
}

// versionGetResponseRulesRulesetsRouteRuleActionParametersSniJSON contains the
// JSON metadata for the struct
// [VersionGetResponseRulesRulesetsRouteRuleActionParametersSni]
type versionGetResponseRulesRulesetsRouteRuleActionParametersSniJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsRouteRuleActionParametersSni) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsRouteRuleActionParametersSniJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseRulesRulesetsScoreRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsScoreRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesRulesetsScoreRuleActionParameters `json:"action_parameters"`
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
	JSON versionGetResponseRulesRulesetsScoreRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsScoreRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseRulesRulesetsScoreRule]
type versionGetResponseRulesRulesetsScoreRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsScoreRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsScoreRuleAction string

const (
	VersionGetResponseRulesRulesetsScoreRuleActionScore VersionGetResponseRulesRulesetsScoreRuleAction = "score"
)

func (r VersionGetResponseRulesRulesetsScoreRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesRulesetsScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment int64                                                        `json:"increment"`
	JSON      versionGetResponseRulesRulesetsScoreRuleActionParametersJSON `json:"-"`
}

// versionGetResponseRulesRulesetsScoreRuleActionParametersJSON contains the JSON
// metadata for the struct
// [VersionGetResponseRulesRulesetsScoreRuleActionParameters]
type versionGetResponseRulesRulesetsScoreRuleActionParametersJSON struct {
	Increment   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsScoreRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsScoreRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseRulesRulesetsServeErrorRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsServeErrorRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesRulesetsServeErrorRuleActionParameters `json:"action_parameters"`
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
	JSON versionGetResponseRulesRulesetsServeErrorRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsServeErrorRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseRulesRulesetsServeErrorRule]
type versionGetResponseRulesRulesetsServeErrorRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsServeErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsServeErrorRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsServeErrorRuleAction string

const (
	VersionGetResponseRulesRulesetsServeErrorRuleActionServeError VersionGetResponseRulesRulesetsServeErrorRuleAction = "serve_error"
)

func (r VersionGetResponseRulesRulesetsServeErrorRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesRulesetsServeErrorRuleActionParameters struct {
	// Error response content.
	Content string `json:"content"`
	// Content-type header to set with the response.
	ContentType VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType `json:"content_type"`
	// The status code to use for the error.
	StatusCode float64                                                           `json:"status_code"`
	JSON       versionGetResponseRulesRulesetsServeErrorRuleActionParametersJSON `json:"-"`
}

// versionGetResponseRulesRulesetsServeErrorRuleActionParametersJSON contains the
// JSON metadata for the struct
// [VersionGetResponseRulesRulesetsServeErrorRuleActionParameters]
type versionGetResponseRulesRulesetsServeErrorRuleActionParametersJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsServeErrorRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsServeErrorRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Content-type header to set with the response.
type VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType string

const (
	VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "application/json"
	VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml         VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/xml"
	VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain       VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/plain"
	VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML        VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/html"
)

func (r VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson, VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml, VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain, VersionGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

type VersionGetResponseRulesRulesetsSetConfigRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsSetConfigRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesRulesetsSetConfigRuleActionParameters `json:"action_parameters"`
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
	JSON versionGetResponseRulesRulesetsSetConfigRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetConfigRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseRulesRulesetsSetConfigRule]
type versionGetResponseRulesRulesetsSetConfigRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsSetConfigRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetConfigRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsSetConfigRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsSetConfigRuleAction string

const (
	VersionGetResponseRulesRulesetsSetConfigRuleActionSetConfig VersionGetResponseRulesRulesetsSetConfigRuleAction = "set_config"
)

func (r VersionGetResponseRulesRulesetsSetConfigRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsSetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesRulesetsSetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites bool `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify VersionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify `json:"autominify"`
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
	Polish VersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolish `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader bool `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes bool `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg  bool                                                             `json:"sxg"`
	JSON versionGetResponseRulesRulesetsSetConfigRuleActionParametersJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetConfigRuleActionParametersJSON contains the
// JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetConfigRuleActionParameters]
type versionGetResponseRulesRulesetsSetConfigRuleActionParametersJSON struct {
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

func (r *VersionGetResponseRulesRulesetsSetConfigRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetConfigRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Select which file extensions to minify automatically.
type VersionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	Css bool `json:"css"`
	// Minify HTML files.
	HTML bool `json:"html"`
	// Minify JS files.
	Js   bool                                                                       `json:"js"`
	JSON versionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify]
type versionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON) RawJSON() string {
	return r.raw
}

// Configure the Polish level.
type VersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolish string

const (
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolishOff      VersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "off"
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless VersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossless"
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy    VersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossy"
)

func (r VersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolishOff, VersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless, VersionGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy:
		return true
	}
	return false
}

// Configure the Security Level.
type VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel string

const (
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff            VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "off"
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "essentially_off"
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow            VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "low"
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium         VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "medium"
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh           VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "high"
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack    VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff, VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff, VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow, VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium, VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh, VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL string

const (
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOff        VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "off"
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible   VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "flexible"
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFull       VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "full"
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict     VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "strict"
	VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "origin_pull"
)

func (r VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOff, VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible, VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFull, VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict, VersionGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

type VersionGetResponseRulesRulesetsSetCacheSettingsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsSetCacheSettingsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters `json:"action_parameters"`
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
	JSON versionGetResponseRulesRulesetsSetCacheSettingsRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleJSON contains the JSON
// metadata for the struct [VersionGetResponseRulesRulesetsSetCacheSettingsRule]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsSetCacheSettingsRule) implementsRulesetsVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleAction string

const (
	VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings VersionGetResponseRulesRulesetsSetCacheSettingsRuleAction = "set_cache_settings"
)

func (r VersionGetResponseRulesRulesetsSetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts []int64 `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache bool `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL `json:"edge_ttl"`
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
	ServeStale VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale `json:"serve_stale"`
	JSON       versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON       `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON struct {
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

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64                                                                             `json:"default"`
	JSON    versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON struct {
	Mode        apijson.Field
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON) RawJSON() string {
	return r.raw
}

// Determines which browser ttl mode to use.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
)

func (r VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType bool `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor bool `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder bool                                                                            `json:"ignore_query_strings_order"`
	JSON                    versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON struct {
	CacheByDeviceType       apijson.Field
	CacheDeceptionArmor     apijson.Field
	CustomKey               apijson.Field
	IgnoreQueryStringsOrder apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON) RawJSON() string {
	return r.raw
}

// Customize which components of the request are included or excluded from the
// cache key.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser `json:"user"`
	JSON versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON) RawJSON() string {
	return r.raw
}

// The cookies to include in building the cache key.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Include these cookies' names and their values.
	Include []string                                                                                       `json:"include"`
	JSON    versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON) RawJSON() string {
	return r.raw
}

// The header names and values to include in building the cache key.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin bool `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include []string                                                                                       `json:"include"`
	JSON    versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON struct {
	CheckPresence apijson.Field
	ExcludeOrigin apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON) RawJSON() string {
	return r.raw
}

// Whether to use the original host or the resolved host in the cache key.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved bool                                                                                         `json:"resolved"`
	JSON     versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON) RawJSON() string {
	return r.raw
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude `json:"include"`
	JSON    versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON    `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON) RawJSON() string {
	return r.raw
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Exclude all query string parameters from use in building the cache key.
	All bool `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List []string                                                                                                   `json:"list"`
	JSON versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON) RawJSON() string {
	return r.raw
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Use all query string parameters in the cache key.
	All bool `json:"all"`
	// A list of query string parameters used to build the cache key.
	List []string                                                                                                   `json:"list"`
	JSON versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON) RawJSON() string {
	return r.raw
}

// Characteristics of the request user agent used in building the cache key.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType bool `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo bool `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang bool                                                                                         `json:"lang"`
	JSON versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON) RawJSON() string {
	return r.raw
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible bool `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize int64                                                                               `json:"min_file_size,required"`
	JSON        versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON struct {
	Eligible    apijson.Field
	MinFileSize apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON) RawJSON() string {
	return r.raw
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64 `json:"default,required"`
	// edge ttl options
	Mode VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL []VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL `json:"status_code_ttl,required"`
	JSON          versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON            `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON struct {
	Default       apijson.Field
	Mode          apijson.Field
	StatusCodeTTL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON) RawJSON() string {
	return r.raw
}

// edge ttl options
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value int64 `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue int64                                                                                       `json:"status_code_value"`
	JSON            versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON struct {
	Value           apijson.Field
	StatusCodeRange apijson.Field
	StatusCodeValue apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON) RawJSON() string {
	return r.raw
}

// The range of status codes used to apply the selected mode.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// response status code lower bound
	From int64 `json:"from,required"`
	// response status code upper bound
	To   int64                                                                                                      `json:"to,required"`
	JSON versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON) RawJSON() string {
	return r.raw
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating bool                                                                              `json:"disable_stale_while_updating,required"`
	JSON                      versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale]
type versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON struct {
	DisableStaleWhileUpdating apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type VersionGetResponseRulesAction string

const (
	VersionGetResponseRulesActionBlock            VersionGetResponseRulesAction = "block"
	VersionGetResponseRulesActionChallenge        VersionGetResponseRulesAction = "challenge"
	VersionGetResponseRulesActionCompressResponse VersionGetResponseRulesAction = "compress_response"
	VersionGetResponseRulesActionExecute          VersionGetResponseRulesAction = "execute"
	VersionGetResponseRulesActionJsChallenge      VersionGetResponseRulesAction = "js_challenge"
	VersionGetResponseRulesActionLog              VersionGetResponseRulesAction = "log"
	VersionGetResponseRulesActionManagedChallenge VersionGetResponseRulesAction = "managed_challenge"
	VersionGetResponseRulesActionRedirect         VersionGetResponseRulesAction = "redirect"
	VersionGetResponseRulesActionRewrite          VersionGetResponseRulesAction = "rewrite"
	VersionGetResponseRulesActionRoute            VersionGetResponseRulesAction = "route"
	VersionGetResponseRulesActionScore            VersionGetResponseRulesAction = "score"
	VersionGetResponseRulesActionServeError       VersionGetResponseRulesAction = "serve_error"
	VersionGetResponseRulesActionSetConfig        VersionGetResponseRulesAction = "set_config"
	VersionGetResponseRulesActionSkip             VersionGetResponseRulesAction = "skip"
	VersionGetResponseRulesActionSetCacheSettings VersionGetResponseRulesAction = "set_cache_settings"
)

func (r VersionGetResponseRulesAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesActionBlock, VersionGetResponseRulesActionChallenge, VersionGetResponseRulesActionCompressResponse, VersionGetResponseRulesActionExecute, VersionGetResponseRulesActionJsChallenge, VersionGetResponseRulesActionLog, VersionGetResponseRulesActionManagedChallenge, VersionGetResponseRulesActionRedirect, VersionGetResponseRulesActionRewrite, VersionGetResponseRulesActionRoute, VersionGetResponseRulesActionScore, VersionGetResponseRulesActionServeError, VersionGetResponseRulesActionSetConfig, VersionGetResponseRulesActionSkip, VersionGetResponseRulesActionSetCacheSettings:
		return true
	}
	return false
}

type VersionListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type VersionDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type VersionGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type VersionGetResponseEnvelope struct {
	// A list of error messages.
	Errors []VersionGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []VersionGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result VersionGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success VersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    versionGetResponseEnvelopeJSON    `json:"-"`
}

// versionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VersionGetResponseEnvelope]
type versionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   versionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// versionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [VersionGetResponseEnvelopeErrors]
type versionGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    versionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// versionGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [VersionGetResponseEnvelopeErrorsSource]
type versionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   versionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// versionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [VersionGetResponseEnvelopeMessages]
type versionGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    versionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// versionGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [VersionGetResponseEnvelopeMessagesSource]
type versionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type VersionGetResponseEnvelopeSuccess bool

const (
	VersionGetResponseEnvelopeSuccessTrue VersionGetResponseEnvelopeSuccess = true
)

func (r VersionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VersionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
